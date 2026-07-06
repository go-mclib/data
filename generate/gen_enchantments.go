package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type enchantmentJSON struct {
	AnvilCost int `json:"anvil_cost"`
	MaxLevel  int `json:"max_level"`
	Weight    int `json:"weight"`
	MinCost   struct {
		Base     int `json:"base"`
		PerLevel int `json:"per_level_above_first"`
	} `json:"min_cost"`
	MaxCost struct {
		Base     int `json:"base"`
		PerLevel int `json:"per_level_above_first"`
	} `json:"max_cost"`
	SupportedItems any `json:"supported_items"` // item, #tag, or list
	PrimaryItems   any `json:"primary_items"`
	ExclusiveSet   any `json:"exclusive_set"`
}

// generateEnchantments emits the enchantment definitions from the datapack.
// Item/tag matchers stay symbolic for runtime resolution.
func generateEnchantments(enchDir, outPath string) {
	entries, err := os.ReadDir(enchDir)
	if err != nil {
		panic(fmt.Sprintf("failed to read enchantment dir %s: %v", enchDir, err))
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".json") {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("enchantments"))
	sb.WriteString("var Enchantments = []Enchantment{\n")
	for _, name := range names {
		e := loadJSON[enchantmentJSON](filepath.Join(enchDir, name))
		id := "minecraft:" + strings.TrimSuffix(name, ".json")
		fmt.Fprintf(&sb, "\t{Name: %q, MaxLevel: %d, Weight: %d, AnvilCost: %d, MinCostBase: %d, MinCostPerLevel: %d, MaxCostBase: %d, MaxCostPerLevel: %d, SupportedItems: %s, PrimaryItems: %s, ExclusiveSet: %s},\n",
			id, e.MaxLevel, e.Weight, e.AnvilCost,
			e.MinCost.Base, e.MinCost.PerLevel, e.MaxCost.Base, e.MaxCost.PerLevel,
			matcherListLiteral(e.SupportedItems), matcherListLiteral(e.PrimaryItems), matcherListLiteral(e.ExclusiveSet))
	}
	sb.WriteString("}\n")

	checkMinCount("enchantments", len(names), 40)
	writeFile(outPath, sb.String())
	fmt.Printf("generated %d enchantments\n", len(names))
}

// matcherListLiteral renders an item/tag matcher field (string or list) as a
// []string literal.
func matcherListLiteral(v any) string {
	var out []string
	switch t := v.(type) {
	case string:
		out = []string{t}
	case []any:
		for _, e := range t {
			if s, ok := e.(string); ok {
				out = append(out, s)
			}
		}
	}
	if len(out) == 0 {
		return "nil"
	}
	parts := make([]string, len(out))
	for i, s := range out {
		parts[i] = fmt.Sprintf("%q", s)
	}
	return "[]string{" + strings.Join(parts, ", ") + "}"
}
