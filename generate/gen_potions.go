package main

import (
	"fmt"
	"strings"
)

type potionEffectJSON struct {
	ID            string            `json:"id"`
	Amplifier     int               `json:"amplifier"`
	Duration      int               `json:"duration"`
	Ambient       bool              `json:"ambient"`
	ShowParticles *bool             `json:"show_particles"`
	ShowIcon      *bool             `json:"show_icon"`
	Hidden        *potionEffectJSON `json:"hidden_effect"`
}

func generatePotions(input, output string) {
	potions := loadJSON[map[string][]potionEffectJSON](input)
	var code strings.Builder
	code.WriteString(generatedFileHeader("items"))
	code.WriteString("// PotionEffects contains each built-in potion's complete vanilla effects.\n")
	code.WriteString("// Treat the map and its effect lists as read-only. Empty potions are included.\n")
	code.WriteString("var PotionEffects = map[string][]PotionEffect{\n")
	for _, name := range sortedKeys(potions) {
		fmt.Fprintf(&code, "\t%q: {", name)
		for _, effect := range potions[name] {
			code.WriteString(potionEffectLiteral(effect, effect.ID))
			code.WriteString(",")
		}
		code.WriteString("},\n")
	}
	code.WriteString("}\n")
	checkMinCount("potions", len(potions), 40)
	writeFile(output, code.String())
}

func potionEffectLiteral(effect potionEffectJSON, id string) string {
	particles := true
	if effect.ShowParticles != nil {
		particles = *effect.ShowParticles
	}
	icon := particles
	if effect.ShowIcon != nil {
		icon = *effect.ShowIcon
	}
	hidden := "nil"
	if effect.Hidden != nil {
		hidden = "&PotionEffect" + potionEffectLiteral(*effect.Hidden, id)
	}
	return fmt.Sprintf("{Effect:%q,Amplifier:%d,Duration:%d,Ambient:%t,ShowParticles:%t,ShowIcon:%t,Hidden:%s}",
		id, effect.Amplifier, effect.Duration, effect.Ambient, particles, icon, hidden)
}
