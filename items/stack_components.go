package items

import (
	"slices"
	"strings"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// StackEnchantment is one enchantment on an item stack (wire id + level).
type StackEnchantment struct {
	ID    int32
	Level int32
}

// patchedComponent looks up a component in the slot's patch. As in vanilla's
// DataComponentPatch decoding, a removal overrides an addition of the same ID,
// and a later addition overrides an earlier one.
func patchedComponent(slot ns.Slot, id int32) (data []byte, added, removed bool) {
	if slices.Contains(slot.Components.Remove, ns.VarInt(id)) {
		return nil, false, true
	}
	for _, component := range slot.Components.Add {
		if int32(component.ID) == id {
			data, added = component.Data, true
		}
	}
	return data, added, false
}

// stackHasComponent reports whether a component is present after applying the
// slot's patch over a default presence.
func stackHasComponent(s ns.Slot, id int32, defaultPresent bool) bool {
	_, added, removed := patchedComponent(s, id)
	return added || !removed && defaultPresent
}

// StackCanStoreEnchantments reports whether the stack's crafting enchantment component is present.
func StackCanStoreEnchantments(s ns.Slot) bool {
	if s.IsEmpty() {
		return false
	}
	if int32(s.ItemID) == ID("minecraft:enchanted_book") {
		return stackHasComponent(s, ComponentStoredEnchantments, true)
	}
	return stackHasComponent(s, ComponentEnchantments, DefaultComponents(int32(s.ItemID)) != nil)
}

// StackHasStoredEnchantments reports whether the stored-enchantments component is present.
func StackHasStoredEnchantments(s ns.Slot) bool {
	defaultPresent := int32(s.ItemID) == ID("minecraft:enchanted_book")
	return stackHasComponent(s, ComponentStoredEnchantments, defaultPresent)
}

// StackDamageable reports whether max-damage and damage are present and unbreakable is absent.
func StackDamageable(s ns.Slot) bool {
	defaults := DefaultComponents(int32(s.ItemID))
	defaultDamageable := defaults != nil && defaults.MaxDamage > 0
	defaultUnbreakable := defaults != nil && defaults.Unbreakable
	return stackHasComponent(s, ComponentMaxDamage, defaultDamageable) &&
		stackHasComponent(s, ComponentDamage, defaultDamageable) &&
		!stackHasComponent(s, ComponentUnbreakable, defaultUnbreakable)
}

// patchedVarInt reads a VarInt component from the slot's patch (0 when absent or invalid).
func patchedVarInt(s ns.Slot, id int32) int32 {
	data, added, _ := patchedComponent(s, id)
	if !added {
		return 0
	}
	v, err := ns.NewReader(data).ReadVarInt()
	if err != nil {
		return 0
	}
	return int32(v)
}

// StackDamage reads the stack's damage component (0 when absent).
func StackDamage(s ns.Slot) int32 { return patchedVarInt(s, ComponentDamage) }

// WithStackDamage copies the stack with the damage component replaced.
func WithStackDamage(s ns.Slot, dmg int32) ns.Slot {
	buf := ns.NewWriter()
	buf.WriteVarInt(ns.VarInt(dmg))
	return withComponent(s, ComponentDamage, buf.Bytes())
}

// StackRepairCost reads the stack's prior-work penalty.
func StackRepairCost(s ns.Slot) int32 { return patchedVarInt(s, ComponentRepairCost) }

// WithStackRepairCost copies the stack with its prior-work penalty replaced.
func WithStackRepairCost(s ns.Slot, cost int32) ns.Slot {
	buf := ns.NewWriter()
	buf.WriteVarInt(ns.VarInt(cost))
	return withComponent(s, ComponentRepairCost, buf.Bytes())
}

// StackRepairableWith reports whether material matches the stack's effective repairable component.
func StackRepairableWith(stack, material ns.Slot) bool {
	if stack.IsEmpty() || material.IsEmpty() {
		return false
	}
	data, added, removed := patchedComponent(stack, ComponentRepairable)
	if removed {
		return false
	}
	var items []string
	if added {
		var err error
		if items, err = readHolderSet(ns.NewReader(data), "item", registries.Item.ByID); err != nil {
			return false
		}
	} else if defaults := DefaultComponents(int32(stack.ItemID)); defaults != nil && defaults.Repairable != nil {
		items = []string{defaults.Repairable.Items}
	}
	return holderSetContainsItem(items, int32(material.ItemID))
}

// holderSetContainsItem reports whether an item holder set (names or a single
// "#tag") contains the item.
func holderSetContainsItem(entries []string, itemID int32) bool {
	if tag, ok := singleTagRef(entries); ok {
		return slices.Contains(Tag(tag), itemID)
	}
	return slices.Contains(entries, Name(itemID))
}

// StackEnchantments reads the stack's enchantments component.
func StackEnchantments(s ns.Slot) []StackEnchantment {
	return readEnchantmentList(s, ComponentEnchantments)
}

// StackStoredEnchantments reads an enchanted book's stored enchantments.
func StackStoredEnchantments(s ns.Slot) []StackEnchantment {
	return readEnchantmentList(s, ComponentStoredEnchantments)
}

// StackEquippableSlot returns the effective equipment slot for a stack.
func StackEquippableSlot(slot ns.Slot) (string, bool) {
	if slot.IsEmpty() {
		return "", false
	}
	data, added, removed := patchedComponent(slot, ComponentEquippable)
	if removed {
		return "", false
	}
	if added {
		equippable, err := readEquippable(ns.NewReader(data))
		if err != nil {
			return "", false
		}
		return equippable.Slot, true
	}
	defaults := DefaultComponents(int32(slot.ItemID))
	if defaults == nil || defaults.Equippable == nil {
		return "", false
	}
	return defaults.Equippable.Slot, true
}

// WithStackStoredEnchantments copies the stack with stored enchantments
// replaced (enchanted books).
func WithStackStoredEnchantments(s ns.Slot, enchs []StackEnchantment) ns.Slot {
	return withEnchantmentList(s, ComponentStoredEnchantments, enchs)
}

func readEnchantmentList(s ns.Slot, componentID int32) []StackEnchantment {
	data, added, _ := patchedComponent(s, componentID)
	if !added {
		return nil
	}
	enchantments, err := readEnchantmentEntries(ns.NewReader(data))
	if err != nil {
		return nil
	}
	return enchantments
}

// WithStackEnchantments copies the stack with the enchantments component
// replaced (an empty list removes it).
func WithStackEnchantments(s ns.Slot, enchs []StackEnchantment) ns.Slot {
	return withEnchantmentList(s, ComponentEnchantments, enchs)
}

func withEnchantmentList(s ns.Slot, componentID int32, enchs []StackEnchantment) ns.Slot {
	if len(enchs) == 0 {
		return withComponent(s, componentID, nil)
	}
	buf := ns.NewWriter()
	buf.WriteVarInt(ns.VarInt(len(enchs)))
	for _, e := range enchs {
		buf.WriteVarInt(ns.VarInt(e.ID))
		buf.WriteVarInt(ns.VarInt(e.Level))
	}
	return withComponent(s, componentID, buf.Bytes())
}

// withComponent copies the stack replacing one component (nil data drops it
// from the patch, so the item default applies).
func withComponent(s ns.Slot, id int32, data []byte) ns.Slot {
	out := s
	out.Components.Add = nil
	for _, c := range s.Components.Add {
		if int32(c.ID) != id {
			out.Components.Add = append(out.Components.Add, c)
		}
	}
	out.Components.Remove = slices.DeleteFunc(slices.Clone(s.Components.Remove), func(removed ns.VarInt) bool {
		return int32(removed) == id
	})
	if data != nil {
		out.Components.Add = append(out.Components.Add, ns.RawSlotComponent{ID: ns.VarInt(id), Data: data})
	}
	return out
}

// slotDamageResistantTypes returns the effective damage-resistant holder set.
func slotDamageResistantTypes(slot ns.Slot) []string {
	data, added, removed := patchedComponent(slot, ComponentDamageResistant)
	if removed {
		return nil
	}
	if added {
		types, err := readHolderSet(ns.NewReader(data), "damage type", damageTypes.ByID)
		if err != nil {
			return nil
		}
		return types
	}
	if defaults := DefaultComponents(int32(slot.ItemID)); defaults != nil && defaults.DamageResistant != nil {
		return []string{defaults.DamageResistant.Types}
	}
	return nil
}

// SlotResistsDamage reports whether a slot's effective damage-resistant component contains a damage type.
func SlotResistsDamage(slot ns.Slot, damageType string) bool {
	for _, entry := range slotDamageResistantTypes(slot) {
		if entry == damageType {
			return true
		}
		if tag, ok := strings.CutPrefix(entry, "#"); ok &&
			slices.Contains(registries.TagData["minecraft:damage_type"][tag], damageType) {
			return true
		}
	}
	return false
}
