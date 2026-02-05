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
// The component order from the slot is preserved for byte-for-byte re-encoding.
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
// Components are always written in ascending ID order.
func (s *ItemStack) ToSlot() (ns.Slot, error) {
	if s.IsEmpty() {
		return ns.EmptySlot(), nil
	}

	slot := ns.NewSlot(ns.VarInt(s.ID), ns.VarInt(s.Count))
	defaults := DefaultComponents(s.ID)

	for id := int32(0); id <= MaxComponentID; id++ {
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

// Builder methods for setting components.

func (s *ItemStack) SetAttributeModifiers(v []AttributeModifier) *ItemStack {
	s.Components.AttributeModifiers = v
	return s
}

func (s *ItemStack) SetBlocksAttacks(v *BlocksAttacks) *ItemStack {
	s.Components.BlocksAttacks = v
	return s
}

func (s *ItemStack) SetBreakSound(v string) *ItemStack {
	s.Components.BreakSound = v
	return s
}

func (s *ItemStack) SetConsumable(v *Consumable) *ItemStack {
	s.Components.Consumable = v
	return s
}

func (s *ItemStack) SetCustomName(v *ItemNameComponent) *ItemStack {
	s.Components.CustomName = v
	return s
}

func (s *ItemStack) SetDamage(v int32) *ItemStack {
	s.Components.Damage = v
	return s
}

func (s *ItemStack) SetDamageResistant(v *DamageResistant) *ItemStack {
	s.Components.DamageResistant = v
	return s
}

func (s *ItemStack) SetDamageType(v string) *ItemStack {
	s.Components.DamageType = v
	return s
}

func (s *ItemStack) SetDeathProtection(v *DeathProtection) *ItemStack {
	s.Components.DeathProtection = v
	return s
}

func (s *ItemStack) SetEnchantable(v *Enchantable) *ItemStack {
	s.Components.Enchantable = v
	return s
}

func (s *ItemStack) SetEnchantments(v map[string]int32) *ItemStack {
	s.Components.Enchantments = v
	return s
}

func (s *ItemStack) SetEquippable(v *Equippable) *ItemStack {
	s.Components.Equippable = v
	return s
}

func (s *ItemStack) SetFireworks(v *Fireworks) *ItemStack {
	s.Components.Fireworks = v
	return s
}

func (s *ItemStack) SetFood(v *Food) *ItemStack {
	s.Components.Food = v
	return s
}

func (s *ItemStack) SetGlider(v bool) *ItemStack {
	s.Components.Glider = v
	return s
}

func (s *ItemStack) SetInstrument(v string) *ItemStack {
	s.Components.Instrument = v
	return s
}

func (s *ItemStack) SetItemModel(v string) *ItemStack {
	s.Components.ItemModel = v
	return s
}

func (s *ItemStack) SetItemName(v *ItemNameComponent) *ItemStack {
	s.Components.ItemName = v
	return s
}

func (s *ItemStack) SetJukeboxPlayable(v string) *ItemStack {
	s.Components.JukeboxPlayable = v
	return s
}

func (s *ItemStack) SetKineticWeapon(v *KineticWeapon) *ItemStack {
	s.Components.KineticWeapon = v
	return s
}

func (s *ItemStack) SetLore(v []string) *ItemStack {
	s.Components.Lore = v
	return s
}

func (s *ItemStack) SetMapColor(v int32) *ItemStack {
	s.Components.MapColor = v
	return s
}

func (s *ItemStack) SetMaxDamage(v int32) *ItemStack {
	s.Components.MaxDamage = v
	return s
}

func (s *ItemStack) SetMaxStackSize(v int32) *ItemStack {
	s.Components.MaxStackSize = v
	return s
}

func (s *ItemStack) SetMinimumAttackCharge(v float64) *ItemStack {
	s.Components.MinimumAttackCharge = v
	return s
}

func (s *ItemStack) SetOminousBottleAmplifier(v int32) *ItemStack {
	s.Components.OminousBottleAmplifier = v
	return s
}

func (s *ItemStack) SetPiercingWeapon(v *PiercingWeapon) *ItemStack {
	s.Components.PiercingWeapon = v
	return s
}

func (s *ItemStack) SetPotionContents(v *PotionContents) *ItemStack {
	s.Components.PotionContents = v
	return s
}

func (s *ItemStack) SetPotionDurationScale(v float64) *ItemStack {
	s.Components.PotionDurationScale = v
	return s
}

func (s *ItemStack) SetProvidesBannerPatterns(v string) *ItemStack {
	s.Components.ProvidesBannerPatterns = v
	return s
}

func (s *ItemStack) SetProvidesTrimMaterial(v string) *ItemStack {
	s.Components.ProvidesTrimMaterial = v
	return s
}

func (s *ItemStack) SetRarity(v string) *ItemStack {
	s.Components.Rarity = v
	return s
}

func (s *ItemStack) SetRepairable(v *Repairable) *ItemStack {
	s.Components.Repairable = v
	return s
}

func (s *ItemStack) SetRepairCost(v int32) *ItemStack {
	s.Components.RepairCost = v
	return s
}

func (s *ItemStack) SetStoredEnchantments(v map[string]int32) *ItemStack {
	s.Components.StoredEnchantments = v
	return s
}

func (s *ItemStack) SetTool(v *Tool) *ItemStack {
	s.Components.Tool = v
	return s
}

func (s *ItemStack) SetTooltipDisplay(v *TooltipDisplay) *ItemStack {
	s.Components.TooltipDisplay = v
	return s
}

func (s *ItemStack) SetUnbreakable(v bool) *ItemStack {
	s.Components.Unbreakable = v
	return s
}

func (s *ItemStack) SetUseCooldown(v *UseCooldown) *ItemStack {
	s.Components.UseCooldown = v
	return s
}

func (s *ItemStack) SetUseEffects(v *UseEffects) *ItemStack {
	s.Components.UseEffects = v
	return s
}

func (s *ItemStack) SetUseRemainder(v *UseRemainder) *ItemStack {
	s.Components.UseRemainder = v
	return s
}

func (s *ItemStack) SetWeapon(v *Weapon) *ItemStack {
	s.Components.Weapon = v
	return s
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
