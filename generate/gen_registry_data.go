package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// generateRegistryData embeds vanilla registry entries as raw JSON. Synchronized
// entries become configuration packets; reloadable entries support item components
// whose values depend on their block context.
func generateRegistryData(datapackDir string, synchronized map[string][]datapackEntry, outPath string) {
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("registries"))
	sb.WriteString("import \"encoding/json\"\n\n")

	sb.WriteString("// SynchronizedRegistryData contains the raw JSON data for each entry\n")
	sb.WriteString("// in each synchronized registry, from the vanilla datapack.\n")
	sb.WriteString("// The server converts these to NBT at startup for S2CRegistryData packets.\n")
	sb.WriteString("var SynchronizedRegistryData = map[string]map[string]json.RawMessage{\n")
	for _, regID := range synchronizedRegistryIDs {
		writeRegistryData(&sb, regID, synchronized[regID])
	}
	sb.WriteString("}\n\n")
	sb.WriteString("// ReloadableRegistryData contains contextual number providers and their predicates.\n")
	sb.WriteString("// These server-side registries are not sent in configuration packets.\n")
	sb.WriteString("var ReloadableRegistryData = map[string]map[string]json.RawMessage{\n")
	for _, regID := range []string{
		"minecraft:context_int_provider",
		"minecraft:context_float_provider",
		"minecraft:predicate",
	} {
		writeRegistryData(&sb, regID, datapackEntries(filepath.Join(datapackDir, strings.TrimPrefix(regID, "minecraft:"))))
	}
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
}

func writeRegistryData(sb *strings.Builder, regID string, entries []datapackEntry) {
	fmt.Fprintf(sb, "\t%q: {\n", regID)
	for _, entry := range entries {
		data, err := os.ReadFile(entry.Path)
		if err != nil {
			panic(fmt.Errorf("load registry %s: %w", regID, err))
		}
		var compact bytes.Buffer
		if err := json.Compact(&compact, data); err != nil {
			panic(fmt.Errorf("load registry %s: %s: %w", regID, entry.Path, err))
		}
		fmt.Fprintf(sb, "\t\t%q: json.RawMessage(%q),\n", entry.ID, compact.String())
	}
	sb.WriteString("\t},\n")
}
