// Package recipes provides vanilla crafting, cooking, and brewing recipe tables,
// generated from the datapack recipe JSONs.
package recipes

// Ingredient lists the acceptable matchers for one slot: item names
// ("minecraft:wheat") or tag references ("#minecraft:planks"), any of which
// fit. Resolve tags with items.Tag. An empty Ingredient matches only an
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

// Transmute changes one input while preserving its component patch, then applies
// Components. Keys prefixed with "!" remove components.
type Transmute struct {
	Info
	Input            Ingredient
	Material         Ingredient
	MinMaterialCount int
	MaxMaterialCount int
	Result           string
	Count            int
	// KeepInput selects the input item instead of Result.
	KeepInput                bool
	Components               map[string]any
	AddMaterialCountToResult bool
}

// PotionIngredient matches an item and optionally its potion contents.
type PotionIngredient struct {
	Item Ingredient
	// PotionContents is the persisted PotionsPredicate: optional "potions"
	// (a holder name, #tag, or list of names) and "effects" (the complete
	// collection predicate, including contains, count, and size). Nil means
	// no potion-contents constraint; an empty map requires potion contents.
	PotionContents map[string]any
}

// Brewing has no recipe-book display or common recipe-book metadata.
type Brewing struct {
	Name    string
	Input   PotionIngredient
	Reagent PotionIngredient
	Result  string
	Count   int
	// Components is the output template's persisted component patch.
	// Keys prefixed with "!" remove components.
	Components map[string]any
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
