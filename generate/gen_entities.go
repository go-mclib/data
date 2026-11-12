package main

import (
	"cmp"
	"fmt"
	"strings"
)

// generateEntities emits the per-entity mob category map from the mod's
// entities.json (the game's own EntityType.getCategory()).
func generateEntities(ents map[string]mcdumpEntity, outPath string) {
	names := sortedKeys(ents)
	checkMinCount("mob categories", len(names), 100)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("entities"))
	sb.WriteString("var entityCategory = map[string]string{\n")
	for _, name := range names {
		fmt.Fprintf(&sb, "\t%q: %q,\n", name, cmp.Or(ents[name].Category, "misc"))
	}
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
	fmt.Printf("mob categories: %d entities\n", len(names))
}
