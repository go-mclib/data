package packets_test

import (
	"github.com/go-mclib/data/pkg/data/items"
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	poSword := items.NewStack(items.IronSword, 1)
	poSword.Components.AttributeModifiers = []items.AttributeModifier{
		{
			Type:      "minecraft:attack_damage",
			Amount:    1000,
			ID:        "minecraft:2121f7b4-5985-43a0-aa3a-57717d7b15c4",
			Operation: "add_multiplied_total",
			Slot:      "any",
		},
		{
			Type:      "minecraft:attack_speed",
			Amount:    100,
			ID:        "minecraft:1df199b2-3849-4112-b9f4-7f16d98d9d38",
			Operation: "add_value",
			Slot:      "any",
		},
	}
	poSword.Components.CustomName = &items.ItemNameComponent{
		Text: "po",
	}
	poSword.Components.TooltipDisplay = &items.TooltipDisplay{
		HideTooltip:      false,
		HiddenComponents: []int32{4, 16},
	}
	poSword.Components.Unbreakable = true

	// component order as captured from Minecraft client (insertion order preserved)
	// TODO: is the order in the MC client actually deterministic? or is it random for every packet?
	// if its random, we could just reorder the bits manually to be alphabetic and make the Go implementation always alphabetic, so its deterministic and can be validated
	// if its not random & we can predict the order, we reverse engineer the algorithm that client uses to order the components
	poSword.SetComponentOrder([]int32{
		items.ComponentTooltipDisplay,     // 18
		items.ComponentCustomName,         // 6
		items.ComponentAttributeModifiers, // 16
		items.ComponentUnbreakable,        // 4
	})
	poSwordSlot, err := poSword.ToSlot()
	if err != nil {
		panic(err)
	}

	capturedPackets[&packets.C2SSetCreativeModeSlot{
		Slot:        36, // first slot in hotbar from left
		ClickedItem: poSwordSlot,
	}] = hexToBytesMust("91010037002401a30704001204000204100605080002706f107702022e6d696e6563726166743a32313231663762342d353938352d343361302d616133612d353737313764376231356334408f400000000000020000042e6d696e6563726166743a31646631393962322d333834392d343131322d623966342d37663136643938643964333840590000000000000000000400")
}
