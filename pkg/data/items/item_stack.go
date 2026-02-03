package items

import (
	"fmt"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// ItemStack represents a fully decoded item stack with typed components.
// It acts as middleware over net_structures.Slot, merging default
// components with slot-specific modifications.
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
// Components are initialized to the item's defaults.
func NewStack(itemID int32, count int32) *ItemStack {
	return &ItemStack{
		ID:         itemID,
		Count:      count,
		Components: DefaultComponents(itemID).Clone(),
	}
}

// IsEmpty returns true if the stack is empty.
func (s *ItemStack) IsEmpty() bool {
	return s == nil || s.Count <= 0
}

// FromSlot creates an ItemStack from a raw net_structures.Slot.
// It applies the slot's component modifications on top of the item's defaults.
func FromSlot(slot ns.Slot) (*ItemStack, error) {
	if slot.IsEmpty() {
		return EmptyStack(), nil
	}

	// start with default components for this item
	defaults := DefaultComponents(int32(slot.ItemID))
	components := defaults.Clone()

	// apply added components
	for _, raw := range slot.Components.Add {
		if err := applyComponent(components, int32(raw.ID), raw.Data); err != nil {
			return nil, fmt.Errorf("component %d: %w", raw.ID, err)
		}
	}

	// apply removals (set to zero values)
	for _, id := range slot.Components.Remove {
		clearComponent(components, int32(id))
	}

	return &ItemStack{
		ID:         int32(slot.ItemID),
		Count:      int32(slot.Count),
		Components: components,
	}, nil
}

// ToSlot converts the ItemStack back to a raw net_structures.Slot.
// Only components that differ from the item's defaults are encoded.
// Components are written in ascending ID order.
func (s *ItemStack) ToSlot() (ns.Slot, error) {
	return s.ToSlotOrdered(nil)
}

// ToSlotOrdered converts the ItemStack to a Slot with components in the specified order.
// If order is nil, components are written in ascending ID order.
// The order slice should contain component IDs in the desired output order.
// Components not in the order slice are appended in ascending ID order.
func (s *ItemStack) ToSlotOrdered(order []int32) (ns.Slot, error) {
	if s.IsEmpty() {
		return ns.EmptySlot(), nil
	}

	slot := ns.NewSlot(ns.VarInt(s.ID), ns.VarInt(s.Count))
	defaults := DefaultComponents(s.ID)

	// collect which components differ
	differing := make(map[int32]bool)
	hasValue := make(map[int32]bool)
	for id := int32(0); id <= MaxComponentID; id++ {
		differs, hv := componentDiffers(s.Components, defaults, id)
		if differs {
			differing[id] = true
			hasValue[id] = hv
		}
	}

	// helper to add a component
	addComponent := func(id int32) error {
		if !differing[id] {
			return nil
		}
		delete(differing, id) // mark as processed
		if hasValue[id] {
			data, err := encodeComponent(s.Components, id)
			if err != nil {
				return fmt.Errorf("encode component %d: %w", id, err)
			}
			slot.AddComponent(ns.VarInt(id), data)
		} else {
			slot.RemoveComponent(ns.VarInt(id))
		}
		return nil
	}

	// add components in specified order first
	for _, id := range order {
		if err := addComponent(id); err != nil {
			return ns.Slot{}, err
		}
	}

	// add remaining components in ID order
	for id := int32(0); id <= MaxComponentID; id++ {
		if err := addComponent(id); err != nil {
			return ns.Slot{}, err
		}
	}

	return slot, nil
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

// ReadSlotDelimited reads a slot with length-prefixed component data.
// Used for packets with OPTIONAL_UNTRUSTED_STREAM_CODEC like creative mode.
func ReadSlotDelimited(buf *ns.PacketBuffer) (*ItemStack, error) {
	slot, err := buf.ReadSlot(DecoderDelimited())
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

// WriteSlotDelimited writes the ItemStack with length-prefixed component data.
// Used for packets with OPTIONAL_UNTRUSTED_STREAM_CODEC like creative mode.
func (s *ItemStack) WriteSlotDelimited(buf *ns.PacketBuffer) error {
	slot, err := s.ToSlot()
	if err != nil {
		return err
	}
	return writeSlotDelimited(buf, slot)
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
