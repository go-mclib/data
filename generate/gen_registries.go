package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// synchronizedRegistryIDs lists the registry identifiers that are sent over the
// network from server to client during the configuration phase.
// Source: decompiled RegistryDataLoader.SYNCHRONIZED_REGISTRIES (Minecraft 26.3).
var synchronizedRegistryIDs = []string{
	"minecraft:worldgen/biome",
	"minecraft:chat_type",
	"minecraft:trim_pattern",
	"minecraft:trim_material",
	"minecraft:wolf_variant",
	"minecraft:wolf_sound_variant",
	"minecraft:pig_variant",
	"minecraft:pig_sound_variant",
	"minecraft:frog_variant",
	"minecraft:cat_variant",
	"minecraft:cat_sound_variant",
	"minecraft:cow_variant",
	"minecraft:cow_sound_variant",
	"minecraft:chicken_variant",
	"minecraft:chicken_sound_variant",
	"minecraft:zombie_nautilus_variant",
	"minecraft:painting_variant",
	"minecraft:dimension_type",
	"minecraft:damage_type",
	"minecraft:banner_pattern",
	"minecraft:enchantment",
	"minecraft:jukebox_song",
	"minecraft:instrument",
	"minecraft:test_environment",
	"minecraft:test_instance",
	"minecraft:dialog",
	"minecraft:world_clock",
	"minecraft:timeline",
	"minecraft:sulfur_cube_archetype",
	"minecraft:decorated_pot_pattern",
	"minecraft:block_transformer",
	"minecraft:worldgen/block_state_provider",
}

// loadSynchronizedEntries walks each synchronized registry's datapack directory
// once; the sorted entries decide both the protocol IDs (SynchronizedEntries)
// and the embedded data (SynchronizedRegistryData).
func loadSynchronizedEntries(datapackDir string) map[string][]datapackEntry {
	result := make(map[string][]datapackEntry, len(synchronizedRegistryIDs))
	for _, regID := range synchronizedRegistryIDs {
		result[regID] = datapackEntries(filepath.Join(datapackDir, strings.TrimPrefix(regID, "minecraft:")))
	}
	return result
}

func generateRegistries(registries map[string]RegistryJSON, synchronized map[string][]datapackEntry, outPath string) {
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("registries"))

	// generate registry variables
	sb.WriteString("// Registry instances\nvar (\n")
	for _, name := range sortedKeys(registries) {
		reg := registries[name]
		goName := toGoVarName(name)
		entriesVar := strings.ToLower(goName[:1]) + goName[1:] + "Entries"
		fmt.Fprintf(&sb, "\t%s = newRegistry(%q, %d, %s)\n", goName, name, reg.ProtocolID, entriesVar)
	}
	sb.WriteString(")\n\n")

	// generate ByIdentifier lookup map
	sb.WriteString("// ByIdentifier maps registry identifier strings to registry instances.\n")
	sb.WriteString("var ByIdentifier = map[string]*Registry{\n")
	for _, name := range sortedKeys(registries) {
		fmt.Fprintf(&sb, "\t%q: %s,\n", name, toGoVarName(name))
	}
	sb.WriteString("}\n\n")

	// generate SynchronizedRegistryIDs
	sb.WriteString("// SynchronizedRegistryIDs lists registry identifiers sent over the network during configuration.\n")
	sb.WriteString("var SynchronizedRegistryIDs = [...]string{\n")
	for _, id := range synchronizedRegistryIDs {
		fmt.Fprintf(&sb, "\t%q,\n", id)
	}
	sb.WriteString("}\n\n")

	// generate entry maps
	for _, name := range sortedKeys(registries) {
		reg := registries[name]
		goName := toGoVarName(name)
		varName := strings.ToLower(goName[:1]) + goName[1:] + "Entries"

		fmt.Fprintf(&sb, "var %s = map[string]int32{\n", varName)
		for _, entryName := range sortedKeys(reg.Entries) {
			fmt.Fprintf(&sb, "\t%q: %d,\n", entryName, reg.Entries[entryName].ProtocolID)
		}
		sb.WriteString("}\n\n")
	}

	// generate SynchronizedEntries: registries for data-driven entries from the
	// vanilla datapack (data/minecraft/<registry>/**.json), in protocol ID order
	sb.WriteString("// SynchronizedEntries maps synchronized registry identifiers to their\n")
	sb.WriteString("// vanilla datapack registries; protocol IDs follow the sorted entry names.\n")
	sb.WriteString("var SynchronizedEntries = map[string]*Registry{\n")
	for _, regID := range synchronizedRegistryIDs {
		entries := synchronized[regID]
		if len(entries) == 0 {
			continue
		}
		fmt.Fprintf(&sb, "\t%q: newSynchronizedRegistry(%q, []string{\n", regID, regID)
		for _, entry := range entries {
			fmt.Fprintf(&sb, "\t\t%q,\n", entry.ID)
		}
		sb.WriteString("\t}),\n")
	}
	sb.WriteString("}\n\n")

	writeFile(outPath, sb.String())
}
