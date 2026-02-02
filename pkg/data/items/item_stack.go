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
func (s *ItemStack) ToSlot() (ns.Slot, error) {
	if s.IsEmpty() {
		return ns.EmptySlot(), nil
	}

	slot := ns.NewSlot(ns.VarInt(s.ID), ns.VarInt(s.Count))
	defaults := DefaultComponents(s.ID)

	// encode components that differ from defaults
	for id := int32(0); id <= MaxComponentID; id++ {
		differs, hasValue := componentDiffers(s.Components, defaults, id)
		if differs {
			if hasValue {
				data, err := encodeComponent(s.Components, id)
				if err != nil {
					return ns.Slot{}, fmt.Errorf("encode component %d: %w", id, err)
				}
				slot.AddComponent(ns.VarInt(id), data)
			} else {
				// component was removed (default has it, we don't)
				slot.RemoveComponent(ns.VarInt(id))
			}
		}
	}

	return slot, nil
}

// Decoder returns a SlotDecoder function that can be passed to Slot.Decode.
// This reads component data from the wire format.
func Decoder() ns.SlotDecoder {
	return decodeComponentWire
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
