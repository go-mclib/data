package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func generateAdvancements(advancementsDir, outPath string) {
	entries := datapackEntries(advancementsDir)
	checkMinCount("advancements", len(entries), 500)

	var output strings.Builder
	output.WriteString(generatedFileHeader("advancements"))
	output.WriteString("import \"encoding/json\"\n\n")
	output.WriteString("// rawDefinitions contains the vanilla advancement datapack definitions.\n")
	output.WriteString("var rawDefinitions = map[string]json.RawMessage{\n")
	for _, entry := range entries {
		data, err := os.ReadFile(entry.Path)
		if err != nil {
			panic(fmt.Sprintf("load advancement %s: %v", entry.Path, err))
		}
		var compact bytes.Buffer
		if err := json.Compact(&compact, data); err != nil {
			panic(fmt.Sprintf("compact advancement %s: %v", entry.Path, err))
		}
		fmt.Fprintf(&output, "\t%q: json.RawMessage(`%s`),\n", entry.ID, compact.Bytes())
	}
	output.WriteString("}\n")
	writeFile(outPath, output.String())
}
