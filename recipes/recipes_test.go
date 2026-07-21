package recipes_test

import (
	"testing"

	"github.com/go-mclib/data/recipes"
)

func TestGeneratedRecipeCounts(t *testing.T) {
	if len(recipes.ShapedRecipes) != 733 || len(recipes.ShapelessRecipes) != 323 || len(recipes.DyeRecipes) != 6 || len(recipes.ImbueRecipes) != 1 || len(recipes.TransmuteRecipes) != 33 || len(recipes.SmeltingRecipes) != 73 ||
		len(recipes.BlastingRecipes) != 25 || len(recipes.SmokingRecipes) != 9 || len(recipes.CampfireRecipes) != 9 ||
		len(recipes.StonecuttingRecipes) != 319 || len(recipes.SmithingRecipes) != 30 {
		t.Fatalf("recipe counts = shaped %d shapeless %d dye %d imbue %d transmute %d smelting %d blasting %d smoking %d campfire %d stonecutting %d smithing %d",
			len(recipes.ShapedRecipes), len(recipes.ShapelessRecipes), len(recipes.DyeRecipes), len(recipes.ImbueRecipes), len(recipes.TransmuteRecipes), len(recipes.SmeltingRecipes), len(recipes.BlastingRecipes),
			len(recipes.SmokingRecipes), len(recipes.CampfireRecipes), len(recipes.StonecuttingRecipes), len(recipes.SmithingRecipes))
	}
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
	if craftingTable == nil || craftingTable.DisplayID != 318 || craftingTable.ShowNotification || craftingTable.Category != "minecraft:crafting_misc" {
		t.Fatalf("crafting table info = %+v", craftingTable)
	}
	if oakPlanks == nil || oakPlanks.DisplayID != 894 || oakPlanks.Group != 7 || oakPlanks.Category != "minecraft:crafting_building_blocks" {
		t.Fatalf("oak planks info = %+v", oakPlanks)
	}
}
