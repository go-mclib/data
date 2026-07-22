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
	Slots                        []string
	EquipmentDropEffects         []EquipmentDropEffect
}

type EquipmentDropEffect struct {
	Enchanted              string
	Affected               string
	Operation              string
	Base                   float32
	PerLevel               float32
	RequiresAttackerPlayer bool
}

func (e EquipmentDropEffect) Apply(level int, chance float32) float32 {
	value := e.Base + e.PerLevel*float32(level-1)
	switch e.Operation {
	case "minecraft:add":
		return chance + value
	case "minecraft:multiply":
		return chance * value
	case "minecraft:set":
		return value
	default:
		return chance
	}
}

func (e *Enchantment) MatchesSlot(slot string) bool {
	for _, group := range e.Slots {
		switch group {
		case "any":
			return true
		case slot:
			return true
		case "hand":
			if slot == "mainhand" || slot == "offhand" {
				return true
			}
		case "armor":
			if slot == "feet" || slot == "legs" || slot == "chest" || slot == "head" {
				return true
			}
		}
	}
	return false
}

var byName = func() map[string]*Enchantment {
	result := make(map[string]*Enchantment, len(Enchantments))
	for index := range Enchantments {
		result[Enchantments[index].Name] = &Enchantments[index]
	}
	return result
}()

func Get(name string) *Enchantment {
	return byName[name]
}

// MinCost is the lowest enchanting power that yields the enchantment at level.
func (e *Enchantment) MinCost(level int) int {
	return e.MinCostBase + e.MinCostPerLevel*(level-1)
}

// MaxCost is the highest enchanting power that yields the enchantment at level.
func (e *Enchantment) MaxCost(level int) int {
	return e.MaxCostBase + e.MaxCostPerLevel*(level-1)
}
