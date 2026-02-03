package items

import (
	"fmt"
	"slices"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// ItemStack represents a fully decoded item stack with typed components.
// It acts as middleware over net_structures.Slot, merging default
// components with slot-specific modifications.
type ItemStack struct {
	ID         int32
	Count      int32
	Components *Components

	// componentOrder preserves the order components appeared in the original packet.
	// Used to maintain byte-for-byte compatibility when re-encoding.
	componentOrder []int32
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

	// capture component order for re-encoding
	var order []int32
	for _, raw := range slot.Components.Add {
		order = append(order, int32(raw.ID))
	}

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
		ID:             int32(slot.ItemID),
		Count:          int32(slot.Count),
		Components:     components,
		componentOrder: order,
	}, nil
}

// ToSlot converts the ItemStack back to a raw net_structures.Slot.
// Only components that differ from the item's defaults are encoded.
// If the ItemStack was created from FromSlot, the original component order is preserved.
// Otherwise, components are written in ascending ID order.
func (s *ItemStack) ToSlot() (ns.Slot, error) {
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

	// add components in preserved order first (from original packet)
	for _, id := range s.componentOrder {
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

// recordComponentOrder adds a component ID to the order list if not already present.
// component order is DETERMINISTIC - follows insertion order into Reference2ObjectArrayMap in Java source code
// the order depends on how the item was created:
// - for /give commands: order components appear in the command string
// - for UI modifications: order player modified components in creative mode
// - for item initialization: order Item.Properties.component() was called
func (s *ItemStack) recordComponentOrder(id int32) {
	if slices.Contains(s.componentOrder, id) {
		return
	}
	s.componentOrder = append(s.componentOrder, id)
}

// Builder methods - each sets a component and records its order for encoding.

func (s *ItemStack) SetAttributeModifiers(v []AttributeModifier) *ItemStack {
	s.Components.AttributeModifiers = v
	s.recordComponentOrder(ComponentAttributeModifiers)
	return s
}

func (s *ItemStack) SetBlocksAttacks(v *BlocksAttacks) *ItemStack {
	s.Components.BlocksAttacks = v
	s.recordComponentOrder(ComponentBlocksAttacks)
	return s
}

func (s *ItemStack) SetBreakSound(v string) *ItemStack {
	s.Components.BreakSound = v
	s.recordComponentOrder(ComponentBreakSound)
	return s
}

func (s *ItemStack) SetConsumable(v *Consumable) *ItemStack {
	s.Components.Consumable = v
	s.recordComponentOrder(ComponentConsumable)
	return s
}

func (s *ItemStack) SetCustomName(v *ItemNameComponent) *ItemStack {
	s.Components.CustomName = v
	s.recordComponentOrder(ComponentCustomName)
	return s
}

func (s *ItemStack) SetDamage(v int32) *ItemStack {
	s.Components.Damage = v
	s.recordComponentOrder(ComponentDamage)
	return s
}

func (s *ItemStack) SetDamageResistant(v *DamageResistant) *ItemStack {
	s.Components.DamageResistant = v
	s.recordComponentOrder(ComponentDamageResistant)
	return s
}

func (s *ItemStack) SetDamageType(v string) *ItemStack {
	s.Components.DamageType = v
	s.recordComponentOrder(ComponentDamageType)
	return s
}

func (s *ItemStack) SetDeathProtection(v *DeathProtection) *ItemStack {
	s.Components.DeathProtection = v
	s.recordComponentOrder(ComponentDeathProtection)
	return s
}

func (s *ItemStack) SetEnchantable(v *Enchantable) *ItemStack {
	s.Components.Enchantable = v
	s.recordComponentOrder(ComponentEnchantable)
	return s
}

func (s *ItemStack) SetEnchantments(v map[string]int32) *ItemStack {
	s.Components.Enchantments = v
	s.recordComponentOrder(ComponentEnchantments)
	return s
}

func (s *ItemStack) SetEquippable(v *Equippable) *ItemStack {
	s.Components.Equippable = v
	s.recordComponentOrder(ComponentEquippable)
	return s
}

func (s *ItemStack) SetFireworks(v *Fireworks) *ItemStack {
	s.Components.Fireworks = v
	s.recordComponentOrder(ComponentFireworks)
	return s
}

func (s *ItemStack) SetFood(v *Food) *ItemStack {
	s.Components.Food = v
	s.recordComponentOrder(ComponentFood)
	return s
}

func (s *ItemStack) SetGlider(v bool) *ItemStack {
	s.Components.Glider = v
	s.recordComponentOrder(ComponentGlider)
	return s
}

func (s *ItemStack) SetInstrument(v string) *ItemStack {
	s.Components.Instrument = v
	s.recordComponentOrder(ComponentInstrument)
	return s
}

func (s *ItemStack) SetItemModel(v string) *ItemStack {
	s.Components.ItemModel = v
	s.recordComponentOrder(ComponentItemModel)
	return s
}

func (s *ItemStack) SetItemName(v *ItemNameComponent) *ItemStack {
	s.Components.ItemName = v
	s.recordComponentOrder(ComponentItemName)
	return s
}

func (s *ItemStack) SetJukeboxPlayable(v string) *ItemStack {
	s.Components.JukeboxPlayable = v
	s.recordComponentOrder(ComponentJukeboxPlayable)
	return s
}

func (s *ItemStack) SetKineticWeapon(v *KineticWeapon) *ItemStack {
	s.Components.KineticWeapon = v
	s.recordComponentOrder(ComponentKineticWeapon)
	return s
}

func (s *ItemStack) SetLore(v []string) *ItemStack {
	s.Components.Lore = v
	s.recordComponentOrder(ComponentLore)
	return s
}

func (s *ItemStack) SetMapColor(v int32) *ItemStack {
	s.Components.MapColor = v
	s.recordComponentOrder(ComponentMapColor)
	return s
}

func (s *ItemStack) SetMaxDamage(v int32) *ItemStack {
	s.Components.MaxDamage = v
	s.recordComponentOrder(ComponentMaxDamage)
	return s
}

func (s *ItemStack) SetMaxStackSize(v int32) *ItemStack {
	s.Components.MaxStackSize = v
	s.recordComponentOrder(ComponentMaxStackSize)
	return s
}

func (s *ItemStack) SetMinimumAttackCharge(v float64) *ItemStack {
	s.Components.MinimumAttackCharge = v
	s.recordComponentOrder(ComponentMinimumAttackCharge)
	return s
}

func (s *ItemStack) SetOminousBottleAmplifier(v int32) *ItemStack {
	s.Components.OminousBottleAmplifier = v
	s.recordComponentOrder(ComponentOminousBottleAmplifier)
	return s
}

func (s *ItemStack) SetPiercingWeapon(v *PiercingWeapon) *ItemStack {
	s.Components.PiercingWeapon = v
	s.recordComponentOrder(ComponentPiercingWeapon)
	return s
}

func (s *ItemStack) SetPotionContents(v *PotionContents) *ItemStack {
	s.Components.PotionContents = v
	s.recordComponentOrder(ComponentPotionContents)
	return s
}

func (s *ItemStack) SetPotionDurationScale(v float64) *ItemStack {
	s.Components.PotionDurationScale = v
	s.recordComponentOrder(ComponentPotionDurationScale)
	return s
}

func (s *ItemStack) SetProvidesBannerPatterns(v string) *ItemStack {
	s.Components.ProvidesBannerPatterns = v
	s.recordComponentOrder(ComponentProvidesBannerPatterns)
	return s
}

func (s *ItemStack) SetProvidesTrimMaterial(v string) *ItemStack {
	s.Components.ProvidesTrimMaterial = v
	s.recordComponentOrder(ComponentProvidesTrimMaterial)
	return s
}

func (s *ItemStack) SetRarity(v string) *ItemStack {
	s.Components.Rarity = v
	s.recordComponentOrder(ComponentRarity)
	return s
}

func (s *ItemStack) SetRepairable(v *Repairable) *ItemStack {
	s.Components.Repairable = v
	s.recordComponentOrder(ComponentRepairable)
	return s
}

func (s *ItemStack) SetRepairCost(v int32) *ItemStack {
	s.Components.RepairCost = v
	s.recordComponentOrder(ComponentRepairCost)
	return s
}

func (s *ItemStack) SetStoredEnchantments(v map[string]int32) *ItemStack {
	s.Components.StoredEnchantments = v
	s.recordComponentOrder(ComponentStoredEnchantments)
	return s
}

func (s *ItemStack) SetTool(v *Tool) *ItemStack {
	s.Components.Tool = v
	s.recordComponentOrder(ComponentTool)
	return s
}

func (s *ItemStack) SetTooltipDisplay(v *TooltipDisplay) *ItemStack {
	s.Components.TooltipDisplay = v
	s.recordComponentOrder(ComponentTooltipDisplay)
	return s
}

func (s *ItemStack) SetUnbreakable(v bool) *ItemStack {
	s.Components.Unbreakable = v
	s.recordComponentOrder(ComponentUnbreakable)
	return s
}

func (s *ItemStack) SetUseCooldown(v *UseCooldown) *ItemStack {
	s.Components.UseCooldown = v
	s.recordComponentOrder(ComponentUseCooldown)
	return s
}

func (s *ItemStack) SetUseEffects(v *UseEffects) *ItemStack {
	s.Components.UseEffects = v
	s.recordComponentOrder(ComponentUseEffects)
	return s
}

func (s *ItemStack) SetUseRemainder(v *UseRemainder) *ItemStack {
	s.Components.UseRemainder = v
	s.recordComponentOrder(ComponentUseRemainder)
	return s
}

func (s *ItemStack) SetWeapon(v *Weapon) *ItemStack {
	s.Components.Weapon = v
	s.recordComponentOrder(ComponentWeapon)
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
