package main

import (
	"fmt"
	"strings"
)

// synchronizedRegistryIDs lists the registry identifiers that are sent over the
// network from server to client during the configuration phase.
// Source: decompiled RegistryDataLoader.SYNCHRONIZED_REGISTRIES (Minecraft 1.21.11).
var synchronizedRegistryIDs = []string{
	"minecraft:worldgen/biome",
	"minecraft:chat_type",
	"minecraft:trim_pattern",
	"minecraft:trim_material",
	"minecraft:wolf_variant",
	"minecraft:wolf_sound_variant",
	"minecraft:pig_variant",
	"minecraft:frog_variant",
	"minecraft:cat_variant",
	"minecraft:cow_variant",
	"minecraft:chicken_variant",
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
	"minecraft:timeline",
}

func generateRegistries(registries map[string]RegistryJSON, outPath string) {
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("registries"))

	// generate registry variables
	sb.WriteString("// Registry instances\nvar (\n")
	for _, name := range sortedKeys(registries) {
		reg := registries[name]
		goName := toGoVarName(name)
		entriesVar := strings.ToLower(goName[:1]) + goName[1:] + "Entries"
		sb.WriteString(fmt.Sprintf("\t%s = newRegistry(%q, %d, %s)\n", goName, name, reg.ProtocolID, entriesVar))
	}
	sb.WriteString(")\n\n")

	// generate ByIdentifier lookup map
	sb.WriteString("// ByIdentifier maps registry identifier strings to registry instances.\n")
	sb.WriteString("var ByIdentifier = map[string]*Registry{\n")
	for _, name := range sortedKeys(registries) {
		goName := toGoVarName(name)
		sb.WriteString(fmt.Sprintf("\t%q: %s,\n", name, goName))
	}
	sb.WriteString("}\n\n")

	// generate SynchronizedRegistryIDs
	sb.WriteString("// SynchronizedRegistryIDs lists registry identifiers sent over the network during configuration.\n")
	sb.WriteString("var SynchronizedRegistryIDs = [...]string{\n")
	for _, id := range synchronizedRegistryIDs {
		sb.WriteString(fmt.Sprintf("\t%q,\n", id))
	}
	sb.WriteString("}\n\n")

	// generate entry maps
	for _, name := range sortedKeys(registries) {
		reg := registries[name]
		goName := toGoVarName(name)
		varName := strings.ToLower(goName[:1]) + goName[1:] + "Entries"

		sb.WriteString(fmt.Sprintf("var %s = map[string]int32{\n", varName))
		for _, entryName := range sortedKeys(reg.Entries) {
			entry := reg.Entries[entryName]
			sb.WriteString(fmt.Sprintf("\t%q: %d,\n", entryName, entry.ProtocolID))
		}
		sb.WriteString("}\n\n")
	}

	writeFile(outPath, sb.String())
}
