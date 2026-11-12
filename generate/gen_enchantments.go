package main

import (
	"fmt"
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
	SupportedItems any                    `json:"supported_items"` // item, #tag, or list
	PrimaryItems   any                    `json:"primary_items"`
	ExclusiveSet   any                    `json:"exclusive_set"`
	Slots          any                    `json:"slots"`
	Effects        enchantmentEffectsJSON `json:"effects"`
}

type enchantmentEffectsJSON struct {
	EquipmentDrops []equipmentDropEffectJSON `json:"minecraft:equipment_drops"`
}

type equipmentDropEffectJSON struct {
	Enchanted string `json:"enchanted"`
	Effect    struct {
		Type  string              `json:"type"`
		Value levelBasedValueJSON `json:"value"`
	} `json:"effect"`
	Requirements *struct {
		Type      string `json:"type"`
		Entity    string `json:"entity"`
		Predicate struct {
			EntityType string `json:"minecraft:entity_type"`
		} `json:"predicate"`
	} `json:"requirements"`
}

type levelBasedValueJSON struct {
	Type     string  `json:"type"`
	Base     float32 `json:"base"`
	PerLevel float32 `json:"per_level_above_first"`
}

// generateEnchantments emits the enchantment definitions from the datapack.
// Item/tag matchers stay symbolic for runtime resolution.
func generateEnchantments(enchDir, outPath string) {
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("enchantments"))
	sb.WriteString("var Enchantments = []Enchantment{\n")
	entries := datapackEntries(enchDir)
	for _, entry := range entries {
		e := loadJSON[enchantmentJSON](entry.Path)
		id := entry.ID
		fmt.Fprintf(&sb, "\t{Name: %q, MaxLevel: %d, Weight: %d, AnvilCost: %d, MinCostBase: %d, MinCostPerLevel: %d, MaxCostBase: %d, MaxCostPerLevel: %d, SupportedItems: %s, PrimaryItems: %s, ExclusiveSet: %s, Slots: %s, EquipmentDropEffects: %s},\n",
			id, e.MaxLevel, e.Weight, e.AnvilCost,
			e.MinCost.Base, e.MinCost.PerLevel, e.MaxCost.Base, e.MaxCost.PerLevel,
			matcherListLiteral(e.SupportedItems), matcherListLiteral(e.PrimaryItems), matcherListLiteral(e.ExclusiveSet),
			matcherListLiteral(e.Slots), equipmentDropEffectsLiteral(id, e.Effects.EquipmentDrops))
	}
	sb.WriteString("}\n")

	checkMinCount("enchantments", len(entries), 40)
	writeFile(outPath, sb.String())
	fmt.Printf("generated %d enchantments\n", len(entries))
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

func equipmentDropEffectsLiteral(enchantment string, effects []equipmentDropEffectJSON) string {
	if len(effects) == 0 {
		return "nil"
	}
	parts := make([]string, len(effects))
	for index, effect := range effects {
		if effect.Enchanted != "attacker" && effect.Enchanted != "victim" {
			panic(fmt.Sprintf("%s has unsupported equipment drop target %q", enchantment, effect.Enchanted))
		}
		if effect.Effect.Type != "minecraft:add" || effect.Effect.Value.Type != "minecraft:linear" {
			panic(fmt.Sprintf("%s has unsupported equipment drop effect %q/%q", enchantment, effect.Effect.Type, effect.Effect.Value.Type))
		}
		requiresAttackerPlayer := false
		if requirement := effect.Requirements; requirement != nil {
			if requirement.Type != "minecraft:entity_properties" || requirement.Entity != "attacker" ||
				requirement.Predicate.EntityType != "minecraft:player" {
				panic(fmt.Sprintf("%s has unsupported equipment drop requirement", enchantment))
			}
			requiresAttackerPlayer = true
		}
		parts[index] = fmt.Sprintf(
			"{Enchanted: %q, Affected: \"victim\", Operation: %q, Base: %g, PerLevel: %g, RequiresAttackerPlayer: %t}",
			effect.Enchanted, effect.Effect.Type, effect.Effect.Value.Base, effect.Effect.Value.PerLevel, requiresAttackerPlayer,
		)
	}
	return "[]EquipmentDropEffect{" + strings.Join(parts, ", ") + "}"
}
