package items

import (
	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

var (
	damageComponentID       = ns.VarInt(registries.DataComponentType.Get("minecraft:damage"))
	enchantmentsComponentID = ns.VarInt(registries.DataComponentType.Get("minecraft:enchantments"))
	storedEnchantmentsID    = ns.VarInt(registries.DataComponentType.Get("minecraft:stored_enchantments"))
	equippableComponentID   = ns.VarInt(registries.DataComponentType.Get("minecraft:equippable"))
)

// StackEnchantment is one enchantment on an item stack (wire id + level).
type StackEnchantment struct {
	ID    int32
	Level int32
}

// StackDamage reads the stack's damage component (0 when absent).
func StackDamage(s ns.Slot) int32 {
	for _, c := range s.Components.Add {
		if c.ID != damageComponentID {
			continue
		}
		v, err := ns.NewReader(c.Data).ReadVarInt()
		if err != nil {
			return 0
		}
		return int32(v)
	}
	return 0
}

// WithStackDamage copies the stack with the damage component replaced.
func WithStackDamage(s ns.Slot, dmg int32) ns.Slot {
	buf := ns.NewWriter()
	buf.WriteVarInt(ns.VarInt(dmg))
	return withComponent(s, damageComponentID, buf.Bytes())
}

// StackEnchantments reads the stack's enchantments component.
func StackEnchantments(s ns.Slot) []StackEnchantment {
	return readEnchantmentList(s, enchantmentsComponentID)
}

// StackStoredEnchantments reads an enchanted book's stored enchantments.
func StackStoredEnchantments(s ns.Slot) []StackEnchantment {
	return readEnchantmentList(s, storedEnchantmentsID)
}

// StackEquippableSlot returns the effective equipment slot for a stack.
func StackEquippableSlot(slot ns.Slot) (string, bool) {
	if slot.IsEmpty() {
		return "", false
	}
	equippable := DefaultComponents(int32(slot.ItemID)).Equippable
	for _, removed := range slot.Components.Remove {
		if removed == equippableComponentID {
			equippable = nil
			break
		}
	}
	for _, component := range slot.Components.Add {
		if component.ID != equippableComponentID {
			continue
		}
		id, err := ns.NewReader(component.Data).ReadVarInt()
		if err != nil {
			return "", false
		}
		slots := [...]string{"mainhand", "offhand", "feet", "legs", "chest", "head", "body", "saddle"}
		if id < 0 || int(id) >= len(slots) {
			return "", false
		}
		return slots[id], true
	}
	if equippable == nil {
		return "", false
	}
	return equippable.Slot, true
}

// WithStackStoredEnchantments copies the stack with stored enchantments
// replaced (enchanted books).
func WithStackStoredEnchantments(s ns.Slot, enchs []StackEnchantment) ns.Slot {
	return withEnchantmentList(s, storedEnchantmentsID, enchs)
}

func readEnchantmentList(s ns.Slot, componentID ns.VarInt) []StackEnchantment {
	for _, c := range s.Components.Add {
		if c.ID != componentID {
			continue
		}
		buf := ns.NewReader(c.Data)
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil
		}
		out := make([]StackEnchantment, 0, int(count))
		for range int(count) {
			id, err1 := buf.ReadVarInt()
			lvl, err2 := buf.ReadVarInt()
			if err1 != nil || err2 != nil {
				return nil
			}
			out = append(out, StackEnchantment{ID: int32(id), Level: int32(lvl)})
		}
		return out
	}
	return nil
}

// WithStackEnchantments copies the stack with the enchantments component
// replaced (an empty list removes it).
func WithStackEnchantments(s ns.Slot, enchs []StackEnchantment) ns.Slot {
	return withEnchantmentList(s, enchantmentsComponentID, enchs)
}

func withEnchantmentList(s ns.Slot, componentID ns.VarInt, enchs []StackEnchantment) ns.Slot {
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

// withComponent copies the stack replacing one component (nil data removes it).
func withComponent(s ns.Slot, id ns.VarInt, data []byte) ns.Slot {
	out := s
	out.Components.Add = nil
	for _, c := range s.Components.Add {
		if c.ID != id {
			out.Components.Add = append(out.Components.Add, c)
		}
	}
	if data != nil {
		out.Components.Add = append(out.Components.Add, ns.RawSlotComponent{ID: id, Data: data})
	}
	return out
}
