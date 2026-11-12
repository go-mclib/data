package items

import (
	"bytes"
	"fmt"
	"reflect"

	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// Stack represents a fully decoded item stack with typed components.
// It acts as middleware over wire.Slot, encoding components
// as a DataComponentPatch (adds/removes) relative to the item's defaults.
type Stack struct {
	ID         int32
	Count      int32
	Components *Components
}

// EmptyStack returns an empty item stack.
func EmptyStack() *Stack {
	return &Stack{}
}

// NewStack creates a new item stack with the given item ID and count.
// Components are initialized to the item's defaults, all marked as present.
func NewStack(itemID int32, count int32) *Stack {
	components := DefaultComponents(itemID).Clone()
	markAllPresent(components)
	return &Stack{
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
		left, leftErr := ns.NewReader(a).ReadNBT()
		right, rightErr := ns.NewReader(b).ReadNBT()
		return leftErr == nil && rightErr == nil && reflect.DeepEqual(left, right)
	}
	return false
}

// IsEmpty returns true if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s == nil || s.Count <= 0
}

// FromSlot creates a Stack from a raw wire.Slot.
// It applies the slot's component patch on top of the item's defaults.
// All default and patched components are marked as present.
func FromSlot(slot ns.Slot) (*Stack, error) {
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

	// apply removals (set to zero values, keep tracked so they encode as removes)
	for _, id := range slot.Components.Remove {
		clearComponent(components, int32(id))
		components.SetPresent(int32(id))
	}

	return &Stack{
		ID:         int32(slot.ItemID),
		Count:      int32(slot.Count),
		Components: components,
	}, nil
}

// ToSlot converts the Stack back to a raw wire.Slot.
// Only present components that differ from the item's defaults are encoded,
// mirroring Java's DataComponentPatch serialization.
func (s *Stack) ToSlot() (ns.Slot, error) {
	if s.IsEmpty() {
		return ns.EmptySlot(), nil
	}

	slot := ns.NewSlot(ns.VarInt(s.ID), ns.VarInt(s.Count))
	defaults := DefaultComponents(s.ID)

	for id := range int32(MaxComponentID + 1) {
		if !s.Components.hasComponent(id) {
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

// markAllPresent marks all registered component IDs for diffing.
func markAllPresent(c *Components) {
	for id := range int32(MaxComponentID + 1) {
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

// maxDelimitedComponentLength bounds a length-prefixed component payload; it
// is vanilla's packet size limit (2 MiB).
const maxDelimitedComponentLength = 2097152

// DecoderDelimited returns a SlotDecoder that handles length-prefixed
// component data, as used by OPTIONAL_UNTRUSTED_STREAM_CODEC (e.g. creative mode slots).
func DecoderDelimited() ns.SlotDecoder {
	return func(buf *ns.PacketBuffer, id ns.VarInt) ([]byte, error) {
		// read length prefix
		length, err := buf.ReadVarInt()
		if err != nil {
			return nil, fmt.Errorf("failed to read component length: %w", err)
		}
		if length < 0 || length > maxDelimitedComponentLength {
			return nil, fmt.Errorf("component length %d out of range", length)
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
// and converts it to a Stack.
func ReadSlot(buf *ns.PacketBuffer) (*Stack, error) {
	slot, err := buf.ReadSlot(Decoder())
	if err != nil {
		return nil, err
	}
	return FromSlot(slot)
}

// UseRemainderSlot returns the component-defined remainder template.
func UseRemainderSlot(slot ns.Slot) (ns.Slot, bool) {
	data, added, removed := patchedComponent(slot, ComponentUseRemainder)
	if removed {
		return ns.Slot{}, false
	}
	if added {
		remainder, err := ReadItemStackTemplate(ns.NewReader(data))
		return remainder, err == nil && !remainder.IsEmpty()
	}
	defaults := DefaultComponents(int32(slot.ItemID))
	if defaults == nil || defaults.UseRemainder == nil {
		return ns.Slot{}, false
	}
	id := ID(defaults.UseRemainder.ID)
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
