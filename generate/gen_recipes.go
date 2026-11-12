package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"strings"
)

type recipeJSON struct {
	Type string `json:"type"`
	// pattern is per-type: []string for crafting_shaped, a trim-pattern id
	// string for smithing — decoded lazily
	Pattern                  json.RawMessage `json:"pattern"`
	Key                      map[string]any  `json:"key"`
	Ingredients              []any           `json:"ingredients"`
	Ingredient               any             `json:"ingredient"`
	Target                   any             `json:"target"`
	Dye                      any             `json:"dye"`
	Input                    any             `json:"input"`
	Reagent                  any             `json:"reagent"`
	Output                   json.RawMessage `json:"output"`
	Source                   any             `json:"source"`
	Material                 any             `json:"material"`
	Template                 any             `json:"template"`
	Base                     any             `json:"base"`
	Addition                 any             `json:"addition"`
	Result                   json.RawMessage `json:"result"`
	Experience               float64         `json:"experience"`
	CookingTime              *int            `json:"cookingtime"`
	Category                 string          `json:"category"`
	Group                    string          `json:"group"`
	MaterialCount            *recipeIntRange `json:"material_count"`
	AddMaterialCountToResult bool            `json:"add_material_count_to_result"`
	ShowNotification         *bool           `json:"show_notification"`
}

type recipeIntRange struct {
	Min *int `json:"min"`
	Max *int `json:"max"`
}

func (r *recipeJSON) patternRows() []string {
	var rows []string
	_ = json.Unmarshal(r.Pattern, &rows)
	return rows
}

type recipeResult struct {
	ID         string         `json:"id"`
	Count      int            `json:"count"`
	Components map[string]any `json:"components"`
}

// generateRecipes emits crafting, cooking, brewing, stonecutting, and smithing recipes
// from the datapack JSONs. Ingredients keep tag references ("#minecraft:planks")
// for runtime resolution via items.Tag.
func generateRecipes(recipeDir, outPath string) {
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("recipes"))

	var shaped, shapeless, dye, imbue, transmute, brewing, smelting, blasting, smoking, campfire, stonecutting, smithing strings.Builder
	nShaped, nShapeless, nDye, nImbue, nTransmute, nSmelting := 0, 0, 0, 0, 0, 0
	nBlasting, nSmoking, nCampfire, nStonecutting, nSmithing, nBrewing := 0, 0, 0, 0, 0, 0
	displayID := 0
	groups := make(map[string]int)

	for _, entry := range datapackEntries(recipeDir) {
		recipeName := entry.ID
		r := loadJSON[recipeJSON](entry.Path)
		result := r.Result
		if r.Type == "minecraft:brewing" {
			result = r.Output
		}
		res := recipeResult{Count: 1}
		if len(result) > 0 {
			// Both item templates and transmute results accept a bare item name.
			if err := json.Unmarshal(result, &res); err != nil {
				if err := json.Unmarshal(result, &res.ID); err != nil {
					panic(fmt.Sprintf("failed to parse recipe result %s: %v", entry.Path, err))
				}
			}
		}
		if r.Type == "minecraft:brewing" {
			fmt.Fprintf(&brewing, "\t{Name: %q, Input: %s, Reagent: %s, Result: %q, Count: %d, Components: %s},\n",
				recipeName, potionIngredientLiteral(r.Input), potionIngredientLiteral(r.Reagent), res.ID, res.Count, recipeJSONLiteral(res.Components))
			nBrewing++
			continue
		}
		groupID := -1
		if r.Group != "" {
			var exists bool
			groupID, exists = groups[r.Group]
			if !exists {
				groupID = len(groups)
				groups[r.Group] = groupID
			}
		}
		showNotification := true
		if r.ShowNotification != nil {
			showNotification = *r.ShowNotification
		}
		info := fmt.Sprintf("Info: Info{Name: %q, DisplayID: %d, Group: %d, Category: %q, ShowNotification: %t}", recipeName, displayID, groupID, recipeBookCategory(r), showNotification)
		displayID += recipeDisplayCount(r)
		switch r.Type {
		case "minecraft:crafting_shaped":
			w, h, cells := shapedCells(r)
			fmt.Fprintf(&shaped, "\t{%s, Width: %d, Height: %d, Cells: %s, Result: %q, Count: %d},\n",
				info, w, h, ingredientsLiteral(cells), res.ID, res.Count)
			nShaped++
		case "minecraft:crafting_shapeless":
			cells := make([]ingredient, 0, len(r.Ingredients))
			for _, ing := range r.Ingredients {
				cells = append(cells, parseIngredient(ing))
			}
			fmt.Fprintf(&shapeless, "\t{%s, Ingredients: %s, Result: %q, Count: %d},\n",
				info, ingredientsLiteral(cells), res.ID, res.Count)
			nShapeless++
		case "minecraft:crafting_dye":
			fmt.Fprintf(&dye, "\t{%s, Target: %s, Dye: %s, Result: %q},\n",
				info, ingredientLiteral(parseIngredient(r.Target)), ingredientLiteral(parseIngredient(r.Dye)), res.ID)
			nDye++
		case "minecraft:crafting_imbue":
			fmt.Fprintf(&imbue, "\t{%s, Source: %s, Material: %s, Result: %q, Count: %d},\n",
				info, ingredientLiteral(parseIngredient(r.Source)), ingredientLiteral(parseIngredient(r.Material)), res.ID, res.Count)
			nImbue++
		case "minecraft:crafting_transmute":
			minCount, maxCount := 1, 1
			if r.MaterialCount != nil {
				if r.MaterialCount.Min != nil {
					minCount = *r.MaterialCount.Min
				}
				if r.MaterialCount.Max != nil {
					maxCount = *r.MaterialCount.Max
				}
			}
			fmt.Fprintf(&transmute, "\t{%s, Input: %s, Material: %s, MinMaterialCount: %d, MaxMaterialCount: %d, Result: %q, Count: %d, KeepInput: %t, Components: %s, AddMaterialCountToResult: %t},\n",
				info, ingredientLiteral(parseIngredient(r.Input)), ingredientLiteral(parseIngredient(r.Material)), minCount, maxCount, res.ID, res.Count, res.ID == "", recipeJSONLiteral(res.Components), r.AddMaterialCountToResult)
			nTransmute++
		case "minecraft:smelting":
			writeCookingRecipe(&smelting, info, r, res, 200)
			nSmelting++
		case "minecraft:blasting":
			writeCookingRecipe(&blasting, info, r, res, 100)
			nBlasting++
		case "minecraft:smoking":
			writeCookingRecipe(&smoking, info, r, res, 100)
			nSmoking++
		case "minecraft:campfire_cooking":
			writeCookingRecipe(&campfire, info, r, res, 100)
			nCampfire++
		case "minecraft:stonecutting":
			fmt.Fprintf(&stonecutting, "\t{%s, Ingredient: %s, Result: %q, Count: %d},\n",
				info, ingredientLiteral(parseIngredient(r.Ingredient)), res.ID, res.Count)
			nStonecutting++
		case "minecraft:smithing_transform", "minecraft:smithing_trim":
			var pattern string
			_ = json.Unmarshal(r.Pattern, &pattern)
			fmt.Fprintf(&smithing, "\t{%s, Type: %q, Template: %s, Base: %s, Addition: %s, Result: %q, Pattern: %q},\n",
				info, r.Type, ingredientLiteral(parseIngredient(r.Template)), ingredientLiteral(parseIngredient(r.Base)), ingredientLiteral(parseIngredient(r.Addition)), res.ID, pattern)
			nSmithing++
		}
	}

	sb.WriteString("var ShapedRecipes = []Shaped{\n")
	sb.WriteString(shaped.String())
	sb.WriteString("}\n\nvar ShapelessRecipes = []Shapeless{\n")
	sb.WriteString(shapeless.String())
	sb.WriteString("}\n\nvar DyeRecipes = []Dye{\n")
	sb.WriteString(dye.String())
	sb.WriteString("}\n\nvar ImbueRecipes = []Imbue{\n")
	sb.WriteString(imbue.String())
	sb.WriteString("}\n\nvar TransmuteRecipes = []Transmute{\n")
	sb.WriteString(transmute.String())
	sb.WriteString("}\n\nvar BrewingRecipes = []Brewing{\n")
	sb.WriteString(brewing.String())
	sb.WriteString("}\n\nvar SmeltingRecipes = []Smelting{\n")
	sb.WriteString(smelting.String())
	sb.WriteString("}\n\nvar BlastingRecipes = []Smelting{\n")
	sb.WriteString(blasting.String())
	sb.WriteString("}\n\nvar SmokingRecipes = []Smelting{\n")
	sb.WriteString(smoking.String())
	sb.WriteString("}\n\nvar CampfireRecipes = []Smelting{\n")
	sb.WriteString(campfire.String())
	sb.WriteString("}\n\nvar StonecuttingRecipes = []Stonecutting{\n")
	sb.WriteString(stonecutting.String())
	sb.WriteString("}\n\nvar SmithingRecipes = []Smithing{\n")
	sb.WriteString(smithing.String())
	sb.WriteString("}\n")

	checkMinCount("shaped recipes", nShaped, 500)
	checkMinCount("dye recipes", nDye, 6)
	checkMinCount("imbue recipes", nImbue, 1)
	checkMinCount("transmute recipes", nTransmute, 33)
	checkMinCount("brewing recipes", nBrewing, 1)
	checkMinCount("shapeless recipes", nShapeless, 100)
	checkMinCount("smelting recipes", nSmelting, 50)
	checkMinCount("blasting recipes", nBlasting, 20)
	checkMinCount("smoking recipes", nSmoking, 8)
	checkMinCount("campfire recipes", nCampfire, 8)
	checkMinCount("stonecutting recipes", nStonecutting, 300)
	checkMinCount("smithing recipes", nSmithing, 25)
	writeFile(outPath, sb.String())
	fmt.Printf("generated %d shaped, %d shapeless, %d dye, %d imbue, %d transmute, %d brewing, %d cooking, %d stonecutting, %d smithing recipes\n",
		nShaped, nShapeless, nDye, nImbue, nTransmute, nBrewing, nSmelting+nBlasting+nSmoking+nCampfire, nStonecutting, nSmithing)
}

func recipeDisplayCount(recipe recipeJSON) int {
	switch recipe.Type {
	case "minecraft:crafting_shaped", "minecraft:crafting_shapeless", "minecraft:crafting_dye", "minecraft:crafting_imbue",
		"minecraft:smelting", "minecraft:blasting", "minecraft:smoking", "minecraft:campfire_cooking",
		"minecraft:stonecutting", "minecraft:smithing_transform", "minecraft:smithing_trim":
		return 1
	case "minecraft:crafting_transmute":
		minCount, maxCount := 1, 1
		if recipe.MaterialCount != nil {
			if recipe.MaterialCount.Min != nil {
				minCount = *recipe.MaterialCount.Min
			}
			if recipe.MaterialCount.Max != nil {
				maxCount = *recipe.MaterialCount.Max
			}
		}
		return max(maxCount-minCount+1, 0)
	default:
		return 0
	}
}

func recipeBookCategory(recipe recipeJSON) string {
	switch recipe.Type {
	case "minecraft:crafting_shaped", "minecraft:crafting_shapeless", "minecraft:crafting_dye", "minecraft:crafting_imbue", "minecraft:crafting_transmute":
		category := cmp.Or(recipe.Category, "misc")
		if category == "building" {
			category = "building_blocks"
		}
		return "minecraft:crafting_" + category
	case "minecraft:smelting":
		category := cmp.Or(recipe.Category, "misc")
		return "minecraft:furnace_" + category
	case "minecraft:blasting":
		if recipe.Category == "blocks" {
			return "minecraft:blast_furnace_blocks"
		}
		return "minecraft:blast_furnace_misc"
	case "minecraft:smoking":
		return "minecraft:smoker_food"
	case "minecraft:campfire_cooking":
		return "minecraft:campfire"
	case "minecraft:stonecutting":
		return "minecraft:stonecutter"
	case "minecraft:smithing_transform", "minecraft:smithing_trim":
		return "minecraft:smithing"
	default:
		return ""
	}
}

func writeCookingRecipe(out *strings.Builder, info string, recipe recipeJSON, result recipeResult, defaultCookingTime int) {
	cookingTime := defaultCookingTime
	if recipe.CookingTime != nil {
		cookingTime = *recipe.CookingTime
	}
	fmt.Fprintf(out, "\t{%s, Ingredient: %s, Result: %q, Experience: %g, CookingTime: %d},\n",
		info, ingredientLiteral(parseIngredient(recipe.Ingredient)), result.ID, recipe.Experience, cookingTime)
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

func potionIngredientLiteral(v any) string {
	predicate, ok := v.(map[string]any)
	if !ok {
		panic(fmt.Sprintf("invalid potion ingredient: %v", v))
	}
	return fmt.Sprintf("PotionIngredient{Item: %s, PotionContents: %s}",
		ingredientLiteral(parseIngredient(predicate["item"])), recipeJSONLiteral(predicate["potion_contents"]))
}

// recipeJSONLiteral preserves JSON values, including empty maps and lists, in
// predicates and component patches without decoding JSON at package startup.
func recipeJSONLiteral(v any) string {
	switch value := v.(type) {
	case nil:
		return "nil"
	case string:
		return fmt.Sprintf("%q", value)
	case bool:
		return fmt.Sprintf("%t", value)
	case float64:
		return fmt.Sprintf("float64(%g)", value)
	case []any:
		if value == nil {
			return "nil"
		}
		parts := make([]string, len(value))
		for i, entry := range value {
			parts[i] = recipeJSONLiteral(entry)
		}
		return "[]any{" + strings.Join(parts, ", ") + "}"
	case map[string]any:
		if value == nil {
			return "nil"
		}
		parts := make([]string, 0, len(value))
		for _, key := range sortedKeys(value) {
			parts = append(parts, fmt.Sprintf("%q: %s", key, recipeJSONLiteral(value[key])))
		}
		return "map[string]any{" + strings.Join(parts, ", ") + "}"
	default:
		panic(fmt.Sprintf("unsupported recipe JSON value %T", v))
	}
}
