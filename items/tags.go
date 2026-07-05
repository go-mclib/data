package items

// ItemTag returns all item IDs belonging to the given tag, or nil if the tag doesn't exist.
// Tag names should be fully qualified (e.g. "minecraft:swords").
func ItemTag(tag string) []int32 {
	return itemTagItems[tag]
}
