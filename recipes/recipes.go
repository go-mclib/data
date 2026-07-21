// Package recipes provides vanilla crafting and smelting recipe tables,
// generated from the datapack recipe JSONs.
package recipes

// Ingredient lists the acceptable matchers for one slot: item names
// ("minecraft:wheat") or tag references ("#minecraft:planks"), any of which
// fit. Resolve tags with items.ItemTag. An empty Ingredient matches only an
// empty slot.
type Ingredient []string

type Info struct {
	Name             string
	DisplayID        int
	Group            int
	Category         string
	ShowNotification bool
}

// Shaped is a pattern-based crafting recipe; Cells is row-major Width×Height.
type Shaped struct {
	Info
	Width, Height int
	Cells         []Ingredient
	Result        string
	Count         int
}

// Shapeless is an order-independent crafting recipe.
type Shapeless struct {
	Info
	Ingredients []Ingredient
	Result      string
	Count       int
}

// Dye is a component-aware shapeless dyeing recipe.
type Dye struct {
	Info
	Target Ingredient
	Dye    Ingredient
	Result string
}

// Imbue copies potion contents from a centered source onto a result.
type Imbue struct {
	Info
	Source   Ingredient
	Material Ingredient
	Result   string
	Count    int
}

// Transmute changes one input while preserving its component patch.
type Transmute struct {
	Info
	Input                    Ingredient
	Material                 Ingredient
	MinMaterialCount         int
	MaxMaterialCount         int
	Result                   string
	Count                    int
	AddMaterialCountToResult bool
}

// Smelting is a furnace recipe.
type Smelting struct {
	Info
	Ingredient  Ingredient
	Result      string
	Experience  float64
	CookingTime int
}

// Stonecutting is a single-input stonecutter recipe.
type Stonecutting struct {
	Info
	Ingredient Ingredient
	Result     string
	Count      int
}

// Smithing contributes ingredients to the synchronized smithing property sets.
type Smithing struct {
	Info
	Type     string
	Template Ingredient
	Base     Ingredient
	Addition Ingredient
	Result   string
	Pattern  string
}
