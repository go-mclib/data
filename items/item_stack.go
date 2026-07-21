package items

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// ItemStack represents a fully decoded item stack with typed components.
// It acts as middleware over net_structures.Slot, encoding components
// as a DataComponentPatch (adds/removes) relative to the item's defaults.
type ItemStack struct {
	ID         int32
	Count      int32
	Components *Components
}

// EmptyStack returns an empty item stack.
func EmptyStack() *ItemStack {
	return &ItemStack{}
}

// NewStack creates a new item stack with the given item ID and count.
// Components are initialized to the item's defaults, all marked as present.
func NewStack(itemID int32, count int32) *ItemStack {
	components := DefaultComponents(itemID).Clone()
	markAllPresent(components)
	return &ItemStack{
		ID:         itemID,
		Count:      count,
		Components: components,
	}
}

// SameStack reports item-and-component-patch equality, ignoring patch entry order.
func SameStack(a, b ns.Slot) bool {
	if a.ItemID != b.ItemID || len(a.Components.Add) != len(b.Components.Add) || len(a.Components.Remove) != len(b.Components.Remove) {
		return false
	}
	added := make(map[ns.VarInt][]byte, len(a.Components.Add))
	for _, component := range a.Components.Add {
		added[component.ID] = component.Data
	}
	for _, component := range b.Components.Add {
		data, ok := added[component.ID]
		if !ok || !componentDataEqual(component.ID, data, component.Data) {
			return false
		}
	}
	removed := make(map[ns.VarInt]struct{}, len(a.Components.Remove))
	for _, component := range a.Components.Remove {
		removed[component] = struct{}{}
	}
	for _, component := range b.Components.Remove {
		if _, ok := removed[component]; !ok {
			return false
		}
	}
	return true
}

func componentDataEqual(id ns.VarInt, a, b []byte) bool {
	if bytes.Equal(a, b) {
		return true
	}
	if codec := componentCodecs[int32(id)]; codec != nil {
		if _, passthrough := codec.(*passthroughCodec); !passthrough {
			var left, right Components
			if codec.Apply(&left, a) == nil && codec.Apply(&right, b) == nil {
				return reflect.DeepEqual(left, right)
			}
		}
	}
	switch int32(id) {
	case ComponentBlockEntityData, ComponentBucketEntityData, ComponentContainerLoot,
		ComponentCustomData, ComponentDebugStickState, ComponentEntityData, ComponentLock, ComponentMapDecorations:
		left, _, leftErr := nbt.NewReaderFrom(bytes.NewReader(a)).ReadTag(true)
		right, _, rightErr := nbt.NewReaderFrom(bytes.NewReader(b)).ReadTag(true)
		return leftErr == nil && rightErr == nil && reflect.DeepEqual(left, right)
	}
	return false
}

// SetDefaultComponents overwrites the stack's components with the item's
// defaults and marks all of them as present.
func (s *ItemStack) SetDefaultComponents() {
	defaults := DefaultComponents(s.ID)
	if defaults != nil {
		s.Components = defaults.Clone()
	} else {
		s.Components = &Components{}
	}
	markAllPresent(s.Components)
}

// IsEmpty returns true if the stack is empty.
func (s *ItemStack) IsEmpty() bool {
	return s == nil || s.Count <= 0
}

// FromSlot creates an ItemStack from a raw net_structures.Slot.
// It applies the slot's component patch on top of the item's defaults.
// All default and patched components are marked as present.
func FromSlot(slot ns.Slot) (*ItemStack, error) {
	if slot.IsEmpty() {
		return EmptyStack(), nil
	}

	defaults := DefaultComponents(int32(slot.ItemID))
	components := defaults.Clone()
	markAllPresent(components)

	// apply added components
	for _, raw := range slot.Components.Add {
		if err := applyComponent(components, int32(raw.ID), raw.Data); err != nil {
			return nil, fmt.Errorf("component %d: %w", raw.ID, err)
		}
		components.SetPresent(int32(raw.ID))
	}

	// apply removals (set to zero values, keep present so they encode as removes)
	for _, id := range slot.Components.Remove {
		clearComponent(components, int32(id))
		components.SetPresent(int32(id))
	}

	return &ItemStack{
		ID:         int32(slot.ItemID),
		Count:      int32(slot.Count),
		Components: components,
	}, nil
}

// ToSlot converts the ItemStack back to a raw net_structures.Slot.
// Only present components that differ from the item's defaults are encoded,
// mirroring Java's DataComponentPatch serialization.
func (s *ItemStack) ToSlot() (ns.Slot, error) {
	if s.IsEmpty() {
		return ns.EmptySlot(), nil
	}

	slot := ns.NewSlot(ns.VarInt(s.ID), ns.VarInt(s.Count))
	defaults := DefaultComponents(s.ID)

	for id := int32(0); id <= MaxComponentID; id++ {
		if !s.Components.HasComponent(id) {
			continue
		}
		differs, hv := componentDiffers(s.Components, defaults, id)
		if !differs {
			continue
		}
		if hv {
			data, err := encodeComponent(s.Components, id)
			if err != nil {
				return ns.Slot{}, fmt.Errorf("encode component %d: %w", id, err)
			}
			slot.AddComponent(ns.VarInt(id), data)
		} else {
			slot.RemoveComponent(ns.VarInt(id))
		}
	}

	return slot, nil
}

// markAllPresent marks all registered component IDs as present.
func markAllPresent(c *Components) {
	for id := int32(0); id <= MaxComponentID; id++ {
		if componentCodecs[id] != nil {
			c.SetPresent(id)
		}
	}
}

// Decoder returns a SlotDecoder function that can be passed to Slot.Decode.
// This reads component data from the wire format.
func Decoder() ns.SlotDecoder {
	return decodeComponentWire
}

// DecoderDelimited returns a SlotDecoder that handles length-prefixed
// component data, as used by OPTIONAL_UNTRUSTED_STREAM_CODEC (e.g. creative mode slots).
func DecoderDelimited() ns.SlotDecoder {
	return func(buf *ns.PacketBuffer, id ns.VarInt) ([]byte, error) {
		// read length prefix
		length, err := buf.ReadVarInt()
		if err != nil {
			return nil, fmt.Errorf("failed to read component length: %w", err)
		}

		if length == 0 {
			// empty component (e.g. Unbreakable)
			return nil, nil
		}

		// read exactly 'length' bytes
		rawData, err := buf.ReadFixedByteArray(int(length))
		if err != nil {
			return nil, fmt.Errorf("failed to read component data: %w", err)
		}

		// decode from the raw bytes
		limitedBuf := ns.NewReader(rawData)
		return decodeComponentWire(limitedBuf, id)
	}
}

// ReadSlot is a convenience function that reads a Slot from the buffer
// and converts it to an ItemStack.
func ReadSlot(buf *ns.PacketBuffer) (*ItemStack, error) {
	slot, err := buf.ReadSlot(Decoder())
	if err != nil {
		return nil, err
	}
	return FromSlot(slot)
}

// WriteSlot writes the ItemStack to the buffer as a Slot.
func (s *ItemStack) WriteSlot(buf *ns.PacketBuffer) error {
	slot, err := s.ToSlot()
	if err != nil {
		return err
	}
	return buf.WriteSlot(slot)
}

// UseRemainderSlot returns the component-defined remainder template.
func UseRemainderSlot(slot ns.Slot) (ns.Slot, bool) {
	for _, removed := range slot.Components.Remove {
		if int32(removed) == ComponentUseRemainder {
			return ns.Slot{}, false
		}
	}
	for _, component := range slot.Components.Add {
		if int32(component.ID) == ComponentUseRemainder {
			remainder, err := ReadItemStackTemplate(ns.NewReader(component.Data))
			return remainder, err == nil && !remainder.IsEmpty()
		}
	}
	defaults := DefaultComponents(int32(slot.ItemID))
	if defaults == nil || defaults.UseRemainder == nil {
		return ns.Slot{}, false
	}
	id := ItemID(defaults.UseRemainder.ID)
	if id < 0 {
		return ns.Slot{}, false
	}
	return ns.NewSlot(ns.VarInt(id), ns.VarInt(max(defaults.UseRemainder.Count, 1))), true
}

// ReadItemStackTemplate reads an item holder, count, and unframed component patch.
func ReadItemStackTemplate(buf *ns.PacketBuffer) (ns.Slot, error) {
	itemID, err := buf.ReadVarInt()
	if err != nil {
		return ns.Slot{}, err
	}
	count, err := buf.ReadVarInt()
	if err != nil {
		return ns.Slot{}, err
	}
	if count < 1 || count > 99 {
		return ns.Slot{}, fmt.Errorf("item template count %d is outside [1,99]", count)
	}
	added, err := buf.ReadVarInt()
	if err != nil {
		return ns.Slot{}, err
	}
	removed, err := buf.ReadVarInt()
	if err != nil {
		return ns.Slot{}, err
	}
	if added < 0 || removed < 0 || int64(added)+int64(removed) > 256 {
		return ns.Slot{}, fmt.Errorf("item template patch size %d/%d is invalid", added, removed)
	}
	slot := ns.NewSlot(itemID, count)
	for range int(added) {
		id, err := buf.ReadVarInt()
		if err != nil {
			return ns.Slot{}, err
		}
		data, err := decodeComponentWire(buf, id)
		if err != nil {
			return ns.Slot{}, err
		}
		slot.Components.Add = append(slot.Components.Add, ns.RawSlotComponent{ID: id, Data: data})
	}
	for range int(removed) {
		id, err := buf.ReadVarInt()
		if err != nil {
			return ns.Slot{}, err
		}
		slot.Components.Remove = append(slot.Components.Remove, id)
	}
	return slot, nil
}

// WriteItemStackTemplate writes an item holder, count, and unframed component patch.
func WriteItemStackTemplate(buf *ns.PacketBuffer, slot ns.Slot) error {
	if slot.IsEmpty() || slot.Count > 99 {
		return fmt.Errorf("item template count %d is outside [1,99]", slot.Count)
	}
	if err := buf.WriteVarInt(slot.ItemID); err != nil {
		return err
	}
	if err := buf.WriteVarInt(slot.Count); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(len(slot.Components.Add))); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(len(slot.Components.Remove))); err != nil {
		return err
	}
	for _, component := range slot.Components.Add {
		if err := buf.WriteVarInt(component.ID); err != nil {
			return err
		}
		if _, err := buf.Write(component.Data); err != nil {
			return err
		}
	}
	for _, id := range slot.Components.Remove {
		if err := buf.WriteVarInt(id); err != nil {
			return err
		}
	}
	return nil
}

// WriteRawSlotDelimited writes a raw ns.Slot with length-prefixed component data.
// Used for packets with OPTIONAL_UNTRUSTED_STREAM_CODEC like creative mode.
func WriteRawSlotDelimited(buf *ns.PacketBuffer, slot ns.Slot) error {
	return writeSlotDelimited(buf, slot)
}

// writeSlotDelimited writes a slot with length-prefixed component data.
func writeSlotDelimited(buf *ns.PacketBuffer, slot ns.Slot) error {
	// write count
	if err := buf.WriteVarInt(slot.Count); err != nil {
		return err
	}

	if slot.Count <= 0 {
		return nil
	}

	// write item ID
	if err := buf.WriteVarInt(slot.ItemID); err != nil {
		return err
	}

	// write add count and remove count
	if err := buf.WriteVarInt(ns.VarInt(len(slot.Components.Add))); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(len(slot.Components.Remove))); err != nil {
		return err
	}

	// write components with length prefix
	for _, comp := range slot.Components.Add {
		// write component ID
		if err := buf.WriteVarInt(comp.ID); err != nil {
			return err
		}
		// write component data length
		if err := buf.WriteVarInt(ns.VarInt(len(comp.Data))); err != nil {
			return err
		}
		// write component data
		if _, err := buf.Write(comp.Data); err != nil {
			return err
		}
	}

	// write removed component IDs (no length prefix for these)
	for _, id := range slot.Components.Remove {
		if err := buf.WriteVarInt(id); err != nil {
			return err
		}
	}

	return nil
}
