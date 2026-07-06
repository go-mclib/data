// Package enchantments provides the vanilla enchantment definitions,
// generated from the datapack enchantment JSONs.
package enchantments

// Enchantment mirrors a datapack enchantment definition. Item matchers are
// item names or tag references ("#minecraft:enchantable/armor"); resolve tags
// with items.ItemTag. ExclusiveSet entries are enchantment names or
// "#minecraft:exclusive_set/..." tags (registries.TagData resolves them).
type Enchantment struct {
	Name      string
	MaxLevel  int
	Weight    int
	AnvilCost int
	// cost bounds for a given level: base + perLevel*(level-1)
	MinCostBase, MinCostPerLevel int
	MaxCostBase, MaxCostPerLevel int
	SupportedItems               []string
	PrimaryItems                 []string
	ExclusiveSet                 []string
}

// MinCost is the lowest enchanting power that yields the enchantment at level.
func (e *Enchantment) MinCost(level int) int {
	return e.MinCostBase + e.MinCostPerLevel*(level-1)
}

// MaxCost is the highest enchanting power that yields the enchantment at level.
func (e *Enchantment) MaxCost(level int) int {
	return e.MaxCostBase + e.MaxCostPerLevel*(level-1)
}
