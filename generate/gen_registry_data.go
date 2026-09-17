package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// generateRegistryData embeds vanilla registry entries as raw JSON. Synchronized
// entries become configuration packets; reloadable entries support item components
// whose values depend on their block context.
func generateRegistryData(datapackDir, outPath string) {
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("registries"))
	sb.WriteString("import \"encoding/json\"\n\n")

	sb.WriteString("// SynchronizedRegistryData contains the raw JSON data for each entry\n")
	sb.WriteString("// in each synchronized registry, from the vanilla datapack.\n")
	sb.WriteString("// The server converts these to NBT at startup for S2CRegistryData packets.\n")
	sb.WriteString("var SynchronizedRegistryData = map[string]map[string]json.RawMessage{\n")

	writeRegistryDataMap(&sb, datapackDir, synchronizedRegistryIDs)
	sb.WriteString("}\n\n")
	sb.WriteString("// ReloadableRegistryData contains contextual number providers and their predicates.\n")
	sb.WriteString("// These server-side registries are not sent in configuration packets.\n")
	sb.WriteString("var ReloadableRegistryData = map[string]map[string]json.RawMessage{\n")
	writeRegistryDataMap(&sb, datapackDir, []string{
		"minecraft:context_int_provider",
		"minecraft:context_float_provider",
		"minecraft:predicate",
	})
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
}

func writeRegistryDataMap(sb *strings.Builder, datapackDir string, registryIDs []string) {
	for _, regID := range registryIDs {
		dir := filepath.Join(datapackDir, strings.TrimPrefix(regID, "minecraft:"))
		sb.WriteString(fmt.Sprintf("\t%q: {\n", regID))
		err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
				return nil
			}
			rel, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			var compact bytes.Buffer
			if err := json.Compact(&compact, data); err != nil {
				return fmt.Errorf("%s: %w", path, err)
			}
			name := "minecraft:" + strings.TrimSuffix(filepath.ToSlash(rel), ".json")
			sb.WriteString(fmt.Sprintf("\t\t%q: json.RawMessage(%q),\n", name, compact.String()))
			return nil
		})
		if err != nil {
			panic(fmt.Errorf("load registry %s: %w", regID, err))
		}
		sb.WriteString("\t},\n")
	}
}
