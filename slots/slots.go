// Package slots provides the slot indices of the player inventory menu (window ID 0).
package slots

// Player inventory slot indices for the survival inventory GUI (window ID 0).
// See https://minecraft.wiki/w/Java_Edition_protocol/Inventory#Player_Inventory
const (
	PlayerCraftingResult = 0
	PlayerCrafting1      = 1 // top-left
	PlayerCrafting2      = 2 // top-right
	PlayerCrafting3      = 3 // bottom-left
	PlayerCrafting4      = 4 // bottom-right
	PlayerArmorHead      = 5
	PlayerArmorChest     = 6
	PlayerArmorLegs      = 7
	PlayerArmorFeet      = 8
	PlayerMainStart      = 9  // top-left of the 27-slot main inventory
	PlayerMainEnd        = 36 // exclusive
	PlayerHotbarStart    = 36 // leftmost hotbar slot
	PlayerHotbarEnd      = 45 // exclusive
	PlayerOffhand        = 45

	PlayerInvSize = 46
)

// PlayerPickupScanOrder is the slot scan order for item pickup:
// hotbar (36-44) left to right, then main inventory (9-35) left to right.
var PlayerPickupScanOrder = func() [36]int {
	var order [36]int
	// hotbar first
	for i := range PlayerHotbarEnd - PlayerHotbarStart {
		order[i] = PlayerHotbarStart + i
	}
	// then main inventory
	for i := range PlayerMainEnd - PlayerMainStart {
		order[9+i] = PlayerMainStart + i
	}
	return order
}()
