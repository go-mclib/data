package items

var craftingRemainders = map[int32]int32{
	ItemID("minecraft:water_bucket"):  ItemID("minecraft:bucket"),
	ItemID("minecraft:lava_bucket"):   ItemID("minecraft:bucket"),
	ItemID("minecraft:milk_bucket"):   ItemID("minecraft:bucket"),
	ItemID("minecraft:dragon_breath"): ItemID("minecraft:glass_bottle"),
	ItemID("minecraft:honey_bottle"):  ItemID("minecraft:glass_bottle"),
}

// CraftingRemainder returns the item left behind when an ingredient is crafted.
func CraftingRemainder(itemID int32) (int32, bool) {
	remainder, ok := craftingRemainders[itemID]
	return remainder, ok
}
