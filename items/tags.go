package items

import (
	"slices"
	"sync"

	"github.com/go-mclib/data/registries"
)

// itemTagItems maps each non-empty item tag to its sorted item IDs, built once
// from registries.TagData; names unknown to the item registry are dropped.
var itemTagItems = sync.OnceValue(func() map[string][]int32 {
	tags := registries.TagData["minecraft:item"]
	result := make(map[string][]int32, len(tags))
	for tag, names := range tags {
		var ids []int32
		for _, name := range names {
			if id := registries.Item.Get(name); id >= 0 {
				ids = append(ids, id)
			}
		}
		if len(ids) > 0 {
			slices.Sort(ids)
			result[tag] = ids
		}
	}
	return result
})

// Tag returns all item IDs belonging to the given tag, or nil if the tag doesn't exist.
// Tag names should be fully qualified (e.g. "minecraft:swords").
func Tag(tag string) []int32 {
	return itemTagItems()[tag]
}
