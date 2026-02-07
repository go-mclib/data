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
	}] = hexToBytesMust("00")

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
	}] = hexToBytesMust("0100000000000000")

	// S2COpenScreen: open double chest
	capturedPackets[&packets.S2COpenScreen{
		WindowId:   1,
		WindowType: 5,
		WindowTitle: ns.TextComponent{
			Translate: "container.chestDouble",
		},
	}] = hexToBytesMust("01050a0800097472616e736c6174650015636f6e7461696e65722e6368657374446f75626c6500")

	// S2COpenScreen: open crafting table
	capturedPackets[&packets.S2COpenScreen{
		WindowId:   2,
		WindowType: 12,
		WindowTitle: ns.TextComponent{
			Translate: "container.crafting",
		},
	}] = hexToBytesMust("020c0a0800097472616e736c6174650012636f6e7461696e65722e6372616674696e6700")

	// S2COpenScreen: open beacon
	capturedPackets[&packets.S2COpenScreen{
		WindowId:   6,
		WindowType: 9,
		WindowTitle: ns.TextComponent{
			Translate: "container.beacon",
		},
	}] = hexToBytesMust("06090a0800097472616e736c6174650010636f6e7461696e65722e626561636f6e00")

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
	}] = hexToBytesMust("00012e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

	// S2CContainerSetSlot: crafting result with 4 oak planks
	capturedPackets[&packets.S2CContainerSetSlot{
		WindowId: 2,
		StateId:  2,
		Slot:     0,
		SlotData: oakPlanksSlot,
	}] = hexToBytesMust("0202000004240000")

	// S2CContainerSetSlot: clear slot to empty
	capturedPackets[&packets.S2CContainerSetSlot{
		WindowId: 2,
		StateId:  3,
		Slot:     0,
		SlotData: ns.EmptySlot(),
	}] = hexToBytesMust("0203000000")

	// S2CContainerSetData: beacon power level = 1
	capturedPackets[&packets.S2CContainerSetData{
		WindowId: 6,
		Property: 0,
		Value:    1,
	}] = hexToBytesMust("0600000001")

	// S2CSetPlayerInventory: 55 emeralds in cursor slot
	capturedPackets[&packets.S2CSetPlayerInventory{
		Slot:     0,
		SlotData: emeralds55Slot,
	}] = hexToBytesMust("003783070000")

	// C2SContainerClose: close window 1
	capturedPackets[&packets.C2SContainerClose{
		WindowId: 1,
	}] = hexToBytesMust("01")

	// C2SPlaceRecipe: place oak planks recipe in crafting table
	capturedPackets[&packets.C2SPlaceRecipe{
		WindowId: 3,
		RecipeId: 838,
		MakeAll:  false,
	}] = hexToBytesMust("03c60600")

	// C2SRecipeBookSeenRecipe: seen recipe 838
	capturedPackets[&packets.C2SRecipeBookSeenRecipe{
		RecipeId: 838, // ID of recipe previously defined in Recipe Book Add.
	}] = hexToBytesMust("c606")

	// C2SSetBeacon: primary effect = haste (2), no secondary
	capturedPackets[&packets.C2SSetBeacon{
		PrimaryEffect:   ns.Some[ns.VarInt](2),
		SecondaryEffect: ns.None[ns.VarInt](),
	}] = hexToBytesMust("010200")

	// C2SSetCarriedItem: select hotbar slot 1
	capturedPackets[&packets.C2SSetCarriedItem{
		Slot: 1,
	}] = hexToBytesMust("0001")

	// C2SSelectTrade: select first trade slot
	capturedPackets[&packets.C2SSelectTrade{
		SelectedSlot: 0, // index of the first trade slot
	}] = hexToBytesMust("00")

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
	}] = hexToBytesMust("0101002a000001002a00018601400000")

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
	}] = hexToBytesMust("01010051000001005101860140000000")

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
	}] = hexToBytesMust("020200000000020000000005000124040000")

	// C2SContainerClick: painting/drag start (no changed slots)
	capturedPackets[&packets.C2SContainerClick{
		WindowId:     2,
		StateId:      2,
		Slot:         -999,
		Button:       0,
		Mode:         5,
		ChangedSlots: []packets.ChangedSlot{},
		CarriedItem:  ns.NewHashedSlot(134, 63),
	}] = hexToBytesMust("0202fc190005000186013f0000")

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
	}] = hexToBytesMust("0202fc19020501000d0186013f000000")

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
	}] = hexToBytesMust("0203000e000001000e012404000000")

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
	}] = hexToBytesMust("0401001f00010200000124080000001f0000")
}
