// Package recipes provides vanilla crafting and smelting recipe tables,
// generated from the datapack recipe JSONs.
package recipes

// Ingredient lists the acceptable matchers for one slot: item names
// ("minecraft:wheat") or tag references ("#minecraft:planks"), any of which
// fit. Resolve tags with items.ItemTag. An empty Ingredient matches only an
// empty slot.
type Ingredient []string

// Shaped is a pattern-based crafting recipe; Cells is row-major Width×Height.
type Shaped struct {
	Width, Height int
	Cells         []Ingredient
	Result        string
	Count         int
}

// Shapeless is an order-independent crafting recipe.
type Shapeless struct {
	Ingredients []Ingredient
	Result      string
	Count       int
}

// Smelting is a furnace recipe.
type Smelting struct {
	Ingredient  Ingredient
	Result      string
	Experience  float64
	CookingTime int
}
