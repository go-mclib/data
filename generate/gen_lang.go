package main

import (
	"fmt"
	"strings"
)

func generateLang(langPath, outPath string) {
	translations := loadJSON[map[string]string](langPath)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("lang"))

	// generate the map
	sb.WriteString("var translations = map[string]string{\n")
	for _, key := range sortedKeys(translations) {
		fmt.Fprintf(&sb, "\t%q: %q,\n", key, translations[key])
	}
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
}
