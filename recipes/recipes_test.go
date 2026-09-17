package recipes_test

import (
	"testing"

	"github.com/go-mclib/data/recipes"
)

func TestGeneratedRecipeMetadata(t *testing.T) {
	var foundTransform, foundTrim bool
	for _, recipe := range recipes.SmithingRecipes {
		if recipe.Type == "minecraft:smithing_transform" && recipe.Result == "minecraft:netherite_sword" {
			foundTransform = true
		}
		if recipe.Type == "minecraft:smithing_trim" && recipe.Pattern == "minecraft:sentry" {
			foundTrim = true
		}
	}
	if !foundTransform || !foundTrim {
		t.Fatalf("smithing metadata: transform=%v trim=%v", foundTransform, foundTrim)
	}
	var craftingTable, oakPlanks *recipes.Info
	for i := range recipes.ShapedRecipes {
		if recipes.ShapedRecipes[i].Name == "minecraft:crafting_table" {
			craftingTable = &recipes.ShapedRecipes[i].Info
		}
	}
	for i := range recipes.ShapelessRecipes {
		if recipes.ShapelessRecipes[i].Name == "minecraft:oak_planks" {
			oakPlanks = &recipes.ShapelessRecipes[i].Info
		}
	}
	if craftingTable == nil || craftingTable.ShowNotification || craftingTable.Category != "minecraft:crafting_misc" {
		t.Fatalf("crafting table info = %+v", craftingTable)
	}
	if oakPlanks == nil || oakPlanks.Category != "minecraft:crafting_building_blocks" {
		t.Fatalf("oak planks info = %+v", oakPlanks)
	}
}
