package packets_test

import (
	"github.com/go-mclib/data/pkg/data/items"
	"github.com/go-mclib/data/pkg/packets"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	oakPlanks := items.NewStack(items.OakPlanks, 4)
	oakPlanksSlot, err := oakPlanks.ToSlot()
	if err != nil {
		panic(err)
	}

	emeralds55 := items.NewStack(items.Emerald, 55)
	emeralds55Slot, err := emeralds55.ToSlot()
	if err != nil {
		panic(err)
	}

	// S2CSetHeldSlot: held slot 0
	capturedPackets[&packets.S2CSetHeldSlot{
		Slot: 0,
	}] = capturedBytes["s2c_set_held_slot"]

	// S2CRecipeBookSettings: crafting open, rest closed
	capturedPackets[&packets.S2CRecipeBookSettings{
		CraftingRecipeBookOpen:             true,
		CraftingRecipeBookFilterActive:     false,
		SmeltingRecipeBookOpen:             false,
		SmeltingRecipeBookFilterActive:     false,
		BlastFurnaceRecipeBookOpen:         false,
		BlastFurnaceRecipeBookFilterActive: false,
		SmokerRecipeBookOpen:               false,
		SmokerRecipeBookFilterActive:       false,
	}] = capturedBytes["s2c_recipe_book_settings"]

	// S2COpenScreen: open double chest
	capturedPackets[&packets.S2COpenScreen{
		WindowId:   1,
		WindowType: 5,
		WindowTitle: ns.TextComponent{
			Translate: "container.chestDouble",
		},
	}] = capturedBytes["s2c_open_screen_chest"]

	// S2COpenScreen: open crafting table
	capturedPackets[&packets.S2COpenScreen{
		WindowId:   2,
		WindowType: 12,
		WindowTitle: ns.TextComponent{
			Translate: "container.crafting",
		},
	}] = capturedBytes["s2c_open_screen_crafting"]

	// S2COpenScreen: open beacon
	capturedPackets[&packets.S2COpenScreen{
		WindowId:   6,
		WindowType: 9,
		WindowTitle: ns.TextComponent{
			Translate: "container.beacon",
		},
	}] = capturedBytes["s2c_open_screen_beacon"]

	// S2CContainerSetContent: empty player inventory (window 0, 46 empty slots)
	emptyInventorySlots := make([]ns.Slot, 46)
	for i := range emptyInventorySlots {
		emptyInventorySlots[i] = ns.EmptySlot()
	}
	capturedPackets[&packets.S2CContainerSetContent{
		WindowId:    0,
		StateId:     1,
		Slots:       emptyInventorySlots,
		CarriedItem: ns.EmptySlot(),
	}] = capturedBytes["s2c_container_set_content_empty"]

	// S2CContainerSetSlot: crafting result with 4 oak planks
	capturedPackets[&packets.S2CContainerSetSlot{
		WindowId: 2,
		StateId:  2,
		Slot:     0,
		SlotData: oakPlanksSlot,
	}] = capturedBytes["s2c_container_set_slot_crafting_result"]

	// S2CContainerSetSlot: clear slot to empty
	capturedPackets[&packets.S2CContainerSetSlot{
		WindowId: 2,
		StateId:  3,
		Slot:     0,
		SlotData: ns.EmptySlot(),
	}] = capturedBytes["s2c_container_set_slot_clear"]

	// S2CContainerSetData: beacon power level = 1
	capturedPackets[&packets.S2CContainerSetData{
		WindowId: 6,
		Property: 0,
		Value:    1,
	}] = capturedBytes["s2c_container_set_data_beacon"]

	// S2CSetPlayerInventory: 55 emeralds in cursor slot
	capturedPackets[&packets.S2CSetPlayerInventory{
		Slot:     0,
		SlotData: emeralds55Slot,
	}] = capturedBytes["s2c_set_player_inventory_emeralds"]

	// C2SContainerClose: close window 1
	capturedPackets[&packets.C2SContainerClose{
		WindowId: 1,
	}] = capturedBytes["c2s_container_close"]

	// C2SPlaceRecipe: place oak planks recipe in crafting table
	capturedPackets[&packets.C2SPlaceRecipe{
		WindowId: 3,
		RecipeId: 838,
		MakeAll:  false,
	}] = capturedBytes["c2s_place_recipe"]

	// C2SRecipeBookSeenRecipe: seen recipe 838
	capturedPackets[&packets.C2SRecipeBookSeenRecipe{
		RecipeId: 838, // ID of recipe previously defined in Recipe Book Add.
	}] = capturedBytes["c2s_recipe_book_seen_recipe"]

	// C2SSetBeacon: primary effect = haste (2), no secondary
	capturedPackets[&packets.C2SSetBeacon{
		PrimaryEffect:   ns.Some[ns.VarInt](2),
		SecondaryEffect: ns.None[ns.VarInt](),
	}] = capturedBytes["c2s_set_beacon"]

	// C2SSetCarriedItem: select hotbar slot 1
	capturedPackets[&packets.C2SSetCarriedItem{
		Slot: 1,
	}] = capturedBytes["c2s_set_carried_item"]

	// C2SSelectTrade: select first trade slot
	capturedPackets[&packets.C2SSelectTrade{
		SelectedSlot: 0, // index of the first trade slot
	}] = capturedBytes["c2s_select_trade"]

	// C2SContainerClick: pick up stack from double chest slot 42
	capturedPackets[&packets.C2SContainerClick{
		WindowId: 1,
		StateId:  1,
		Slot:     42,
		Button:   0,
		Mode:     0,
		ChangedSlots: []packets.ChangedSlot{
			{SlotNum: 42, Item: ns.EmptyHashedSlot()},
		},
		CarriedItem: ns.NewHashedSlot(134, 64),
	}] = capturedBytes["c2s_container_click_pickup_chest"]

	// C2SContainerClick: place stack into hotbar slot 81
	capturedPackets[&packets.C2SContainerClick{
		WindowId: 1,
		StateId:  1,
		Slot:     81,
		Button:   0,
		Mode:     0,
		ChangedSlots: []packets.ChangedSlot{
			{SlotNum: 81, Item: ns.NewHashedSlot(134, 64)},
		},
		CarriedItem: ns.EmptyHashedSlot(),
	}] = capturedBytes["c2s_container_click_place_hotbar"]

	// C2SContainerClick: pick up crafting result (2 changed slots emptied)
	capturedPackets[&packets.C2SContainerClick{
		WindowId: 2,
		StateId:  2,
		Slot:     0,
		Button:   0,
		Mode:     0,
		ChangedSlots: []packets.ChangedSlot{
			{SlotNum: 0, Item: ns.EmptyHashedSlot()},
			{SlotNum: 5, Item: ns.EmptyHashedSlot()},
		},
		CarriedItem: ns.NewHashedSlot(36, 4),
	}] = capturedBytes["c2s_container_click_crafting_result"]

	// C2SContainerClick: painting/drag start (no changed slots)
	capturedPackets[&packets.C2SContainerClick{
		WindowId:     2,
		StateId:      2,
		Slot:         -999,
		Button:       0,
		Mode:         5,
		ChangedSlots: []packets.ChangedSlot{},
		CarriedItem:  ns.NewHashedSlot(134, 63),
	}] = capturedBytes["c2s_container_click_drag_start"]

	// C2SContainerClick: painting/drag end (1 slot filled, cursor emptied)
	capturedPackets[&packets.C2SContainerClick{
		WindowId: 2,
		StateId:  2,
		Slot:     -999,
		Button:   2,
		Mode:     5,
		ChangedSlots: []packets.ChangedSlot{
			{SlotNum: 13, Item: ns.NewHashedSlot(134, 63)},
		},
		CarriedItem: ns.EmptyHashedSlot(),
	}] = capturedBytes["c2s_container_click_drag_end"]

	// C2SContainerClick: place crafted item into inventory slot 14
	capturedPackets[&packets.C2SContainerClick{
		WindowId: 2,
		StateId:  3,
		Slot:     14,
		Button:   0,
		Mode:     0,
		ChangedSlots: []packets.ChangedSlot{
			{SlotNum: 14, Item: ns.NewHashedSlot(36, 4)},
		},
		CarriedItem: ns.EmptyHashedSlot(),
	}] = capturedBytes["c2s_container_click_place_crafted"]

	// C2SContainerClick: shift-click item from slot 31 to crafting output
	capturedPackets[&packets.C2SContainerClick{
		WindowId: 4,
		StateId:  1,
		Slot:     31,
		Button:   0,
		Mode:     1,
		ChangedSlots: []packets.ChangedSlot{
			{SlotNum: 0, Item: ns.NewHashedSlot(36, 8)},
			{SlotNum: 31, Item: ns.EmptyHashedSlot()},
		},
		CarriedItem: ns.EmptyHashedSlot(),
	}] = capturedBytes["c2s_container_click_shift_click"]
}
