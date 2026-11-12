package items

// CraftingRemainder returns the item left behind when an ingredient is crafted.
func CraftingRemainder(itemID int32) (int32, bool) {
	remainder, ok := craftingRemainders[itemID]
	return remainder, ok
}
