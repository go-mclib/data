package items

import (
	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

var potionContentsComponent = ns.VarInt(registries.DataComponentType.Get("minecraft:potion_contents"))

// PotionSlot builds a potion item stack carrying a potion_contents component
// (PotionContents codec: optional potion + optional color + custom effects +
// optional name).
func PotionSlot(itemID ns.VarInt, potion string) ns.Slot {
	buf := ns.NewWriter()
	buf.WriteBool(true)
	buf.WriteVarInt(ns.VarInt(registries.Potion.Get(potion)))
	buf.WriteBool(false) // no custom color
	buf.WriteVarInt(0)   // no custom effects
	buf.WriteBool(false) // no custom name
	s := ns.NewSlot(itemID, 1)
	s.Components.Add = append(s.Components.Add, ns.RawSlotComponent{
		ID: potionContentsComponent, Data: buf.Bytes(),
	})
	return s
}

// PotionName reads the potion id from a stack's potion_contents component.
// A potion item without the component reads as water (lenient for /give).
func PotionName(s ns.Slot) string {
	for _, c := range s.Components.Add {
		if c.ID != potionContentsComponent {
			continue
		}
		buf := ns.NewReader(c.Data)
		has, err := buf.ReadBool()
		if err != nil || !bool(has) {
			return ""
		}
		id, err := buf.ReadVarInt()
		if err != nil {
			return ""
		}
		return registries.Potion.ByID(int32(id))
	}
	return "minecraft:water"
}
