package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type recipeJSON struct {
	Type string `json:"type"`
	// pattern is per-type: []string for crafting_shaped, a trim-pattern id
	// string for smithing — decoded lazily
	Pattern     json.RawMessage `json:"pattern"`
	Key         map[string]any  `json:"key"`
	Ingredients []any           `json:"ingredients"`
	Ingredient  any             `json:"ingredient"`
	Result      json.RawMessage `json:"result"`
	Experience  float64         `json:"experience"`
	CookingTime int             `json:"cookingtime"`
}

func (r *recipeJSON) patternRows() []string {
	var rows []string
	_ = json.Unmarshal(r.Pattern, &rows)
	return rows
}

type recipeResult struct {
	ID    string `json:"id"`
	Count int    `json:"count"`
}

// generateRecipes emits shaped/shapeless crafting and smelting recipes from
// the datapack recipe JSONs. Ingredients keep tag references ("#minecraft:planks")
// for runtime resolution via items.ItemTag.
func generateRecipes(recipeDir, outPath string) {
	entries, err := os.ReadDir(recipeDir)
	if err != nil {
		panic(fmt.Sprintf("failed to read recipe dir %s: %v", recipeDir, err))
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".json") {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("recipes"))

	var shaped, shapeless, smelting strings.Builder
	nShaped, nShapeless, nSmelting := 0, 0, 0

	for _, name := range names {
		r := loadJSON[recipeJSON](filepath.Join(recipeDir, name))
		var res recipeResult
		if len(r.Result) > 0 {
			// smelting results may be a bare string in older formats; modern is an object
			if err := json.Unmarshal(r.Result, &res); err != nil {
				var s string
				if json.Unmarshal(r.Result, &s) == nil {
					res.ID = s
				}
			}
		}
		if res.Count == 0 {
			res.Count = 1
		}
		switch r.Type {
		case "minecraft:crafting_shaped":
			w, h, cells := shapedCells(r)
			fmt.Fprintf(&shaped, "\t{Width: %d, Height: %d, Cells: %s, Result: %q, Count: %d},\n",
				w, h, ingredientsLiteral(cells), res.ID, res.Count)
			nShaped++
		case "minecraft:crafting_shapeless":
			cells := make([]ingredient, 0, len(r.Ingredients))
			for _, ing := range r.Ingredients {
				cells = append(cells, parseIngredient(ing))
			}
			fmt.Fprintf(&shapeless, "\t{Ingredients: %s, Result: %q, Count: %d},\n",
				ingredientsLiteral(cells), res.ID, res.Count)
			nShapeless++
		case "minecraft:smelting":
			fmt.Fprintf(&smelting, "\t{Ingredient: %s, Result: %q, Experience: %g, CookingTime: %d},\n",
				ingredientLiteral(parseIngredient(r.Ingredient)), res.ID, r.Experience, r.CookingTime)
			nSmelting++
		}
	}

	sb.WriteString("var ShapedRecipes = []Shaped{\n")
	sb.WriteString(shaped.String())
	sb.WriteString("}\n\nvar ShapelessRecipes = []Shapeless{\n")
	sb.WriteString(shapeless.String())
	sb.WriteString("}\n\nvar SmeltingRecipes = []Smelting{\n")
	sb.WriteString(smelting.String())
	sb.WriteString("}\n")

	checkMinCount("shaped recipes", nShaped, 500)
	checkMinCount("shapeless recipes", nShapeless, 100)
	checkMinCount("smelting recipes", nSmelting, 50)
	writeFile(outPath, sb.String())
	fmt.Printf("generated %d shaped, %d shapeless, %d smelting recipes\n", nShaped, nShapeless, nSmelting)
}

type ingredient []string

// parseIngredient normalizes a JSON ingredient (string or array) into
// alternative matchers.
func parseIngredient(v any) ingredient {
	switch t := v.(type) {
	case string:
		return ingredient{t}
	case []any:
		var out ingredient
		for _, e := range t {
			if s, ok := e.(string); ok {
				out = append(out, s)
			}
		}
		return out
	}
	return nil
}

// shapedCells flattens the pattern+key into row-major cells.
func shapedCells(r recipeJSON) (w, h int, cells []ingredient) {
	rows := r.patternRows()
	h = len(rows)
	for _, row := range rows {
		w = max(w, len(row))
	}
	for _, row := range rows {
		for x := range w {
			if x >= len(row) || row[x] == ' ' {
				cells = append(cells, nil)
				continue
			}
			cells = append(cells, parseIngredient(r.Key[string(row[x])]))
		}
	}
	return w, h, cells
}

func ingredientLiteral(ing ingredient) string {
	if len(ing) == 0 {
		return "nil"
	}
	parts := make([]string, len(ing))
	for i, s := range ing {
		parts[i] = fmt.Sprintf("%q", s)
	}
	return "Ingredient{" + strings.Join(parts, ", ") + "}"
}

func ingredientsLiteral(cells []ingredient) string {
	parts := make([]string, len(cells))
	for i, c := range cells {
		parts[i] = ingredientLiteral(c)
	}
	return "[]Ingredient{" + strings.Join(parts, ", ") + "}"
}
