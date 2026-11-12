package main

import (
	"fmt"
	"strings"
)

// generateMobEffectColors emits each mob effect's color (MobEffect.getColor)
// from mcdump/effects.json.
func generateMobEffectColors(mcdumpPath, outPath string) {
	effects := loadJSON[map[string]struct {
		Color int32 `json:"color"`
	}](mcdumpPath)
	checkMinCount("mob effect colors", len(effects), 30)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("registries"))
	sb.WriteString("var mobEffectColors = map[string]int32{\n")
	for _, name := range sortedKeys(effects) {
		fmt.Fprintf(&sb, "\t%q: %d,\n", name, effects[name].Color)
	}
	sb.WriteString("}\n")
	writeFile(outPath, sb.String())
}

// generateCraftingRemainders emits the item each ingredient leaves behind when
// crafted (Item.getCraftingRemainder) from mcdump/items.json.
func generateCraftingRemainders(mcdumpPath string, registries map[string]RegistryJSON, outPath string) {
	items := loadJSON[map[string]struct {
		CraftingRemainder string `json:"craftingRemainder"`
	}](mcdumpPath)
	itemRegistry := registries["minecraft:item"]
	itemID := func(name string) int32 {
		entry, ok := itemRegistry.Entries[name]
		if !ok {
			panic(fmt.Sprintf("crafting remainder: unknown item %s", name))
		}
		return entry.ProtocolID
	}
	checkMinCount("crafting remainders", len(items), 3)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("items"))
	sb.WriteString("var craftingRemainders = map[int32]int32{\n")
	for _, name := range sortedKeys(items) {
		remainder := items[name].CraftingRemainder
		fmt.Fprintf(&sb, "\t%d: %d, // %s -> %s\n", itemID(name), itemID(remainder), name, remainder)
	}
	sb.WriteString("}\n")
	writeFile(outPath, sb.String())
}
