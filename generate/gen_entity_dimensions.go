package main

import (
	"fmt"
	"strconv"
	"strings"
)

// mcdumpEntity is one entry of mcdump/entities.json.
type mcdumpEntity struct {
	ProtocolID    int     `json:"protocolId"`
	Width         float64 `json:"width"`
	Height        float64 `json:"height"`
	EyeHeight     float64 `json:"eyeHeight"`
	Category      string  `json:"category"`
	TrackingRange int     `json:"trackingRange"`
}

// generateEntityDimensions emits standing hitbox dimensions (width, height, real eye
// height) for every entity type, straight from the mod's entities.json.
func generateEntityDimensions(ents map[string]mcdumpEntity, outPath string) {
	names := sortedKeys(ents)
	checkMinCount("entity hitboxes", len(names), 100)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("entities"))
	fmt.Fprintf(&sb, "// dimensions contains the standing hitbox dimensions for %d entity types.\n", len(names))
	sb.WriteString("var dimensions = map[string]struct{ Width, Height, EyeHeight float32 }{\n")
	for _, name := range names {
		e := ents[name]
		fmt.Fprintf(&sb, "\t%q: {%s, %s, %s},\n",
			name, formatFloat32(e.Width), formatFloat32(e.Height), formatFloat32(e.EyeHeight))
	}
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
	fmt.Printf("entity hitboxes: %d entities\n", len(names))
}

func formatFloat32(f float64) string {
	s := strconv.FormatFloat(f, 'f', -1, 32)
	if !strings.Contains(s, ".") {
		s += ".0"
	}
	return s
}
