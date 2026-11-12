package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func generateBlockStates(blocks map[string]BlockJSON, registries map[string]RegistryJSON, outPath string) {
	blockRegistry := registries["minecraft:block"]

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("blocks"))

	// generate block state data map
	sb.WriteString("var blockStates = map[int32]*blockStateData{\n")

	type stateRange struct {
		baseID  int32
		endID   int32
		blockID int32
	}
	var ranges []stateRange
	for _, blockName := range sortedKeys(blocks) {
		block := blocks[blockName]
		if len(block.States) == 0 {
			continue
		}

		// find base and default state IDs
		baseID := block.States[0].ID
		defaultID := baseID
		for _, state := range block.States {
			if state.Default {
				defaultID = state.ID
				break
			}
		}

		blockID := blockRegistry.Entries[blockName].ProtocolID
		ranges = append(ranges, stateRange{baseID, block.States[len(block.States)-1].ID + 1, blockID})
		fmt.Fprintf(&sb, "\t%d: { // %s\n", blockID, blockName)
		fmt.Fprintf(&sb, "\t\tBaseID:    %d,\n", baseID)
		fmt.Fprintf(&sb, "\t\tDefaultID: %d,\n", defaultID)

		// vanilla keeps properties in an ImmutableSortedMap and enumerates states
		// as their cartesian product, so the order is always alphabetical
		if len(block.Properties) > 0 {
			sb.WriteString("\t\tProperties: []blockProperty{\n")
			for _, prop := range sortedKeys(block.Properties) {
				values := block.Properties[prop]
				fmt.Fprintf(&sb, "\t\t\t{%q, []string{", prop)
				for i, v := range values {
					if i > 0 {
						sb.WriteString(", ")
					}
					fmt.Fprintf(&sb, "%q", v)
				}
				sb.WriteString("}},\n")
			}
			sb.WriteString("\t\t},\n")
		}

		sb.WriteString("\t},\n")
	}
	sb.WriteString("}\n\n")

	// state range index for O(log n) lookup
	slices.SortFunc(ranges, func(a, b stateRange) int { return cmp.Compare(a.baseID, b.baseID) })
	sb.WriteString("// stateRanges is sorted by BaseID for binary search.\n")
	sb.WriteString("var stateRanges = []stateRangeEntry{\n")
	for _, r := range ranges {
		fmt.Fprintf(&sb, "\t{%d, %d, %d},\n", r.baseID, r.endID, r.blockID)
	}
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
}
