package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type prismarineShapesJSON struct {
	Blocks map[string]json.RawMessage `json:"blocks"`
	Shapes map[string][][]float64     `json:"shapes"`
}

func generateBlockShapes(blocks map[string]BlockJSON, prismarinePath, outPath string) {
	data, err := os.ReadFile(prismarinePath)
	if err != nil {
		fmt.Printf("skipping block shapes generation: %v\n", err)
		return
	}
	var prismarine prismarineShapesJSON
	if err := json.Unmarshal(data, &prismarine); err != nil {
		panic(fmt.Sprintf("failed to parse %s: %v", prismarinePath, err))
	}

	// remap sparse PrismarineJS shape IDs to compact sequential indices
	// collect and sort original IDs
	origIDs := make([]int, 0, len(prismarine.Shapes))
	for idStr := range prismarine.Shapes {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(fmt.Sprintf("invalid shape ID %q", idStr))
		}
		origIDs = append(origIDs, id)
	}
	sort.Ints(origIDs)

	// build compact shapes list and ID mapping
	type shapeEntry struct {
		aabbs [][]float64
	}
	compactShapes := make([]shapeEntry, len(origIDs))
	origToCompact := make(map[int]uint16, len(origIDs))
	for i, origID := range origIDs {
		compactShapes[i] = shapeEntry{prismarine.Shapes[strconv.Itoa(origID)]}
		origToCompact[origID] = uint16(i)
	}

	// find the full block shape index (single AABB of 0,0,0 -> 1,1,1)
	fullBlockIdx := uint16(0)
	for i, s := range compactShapes {
		if len(s.aabbs) == 1 && len(s.aabbs[0]) == 6 &&
			s.aabbs[0][0] == 0 && s.aabbs[0][1] == 0 && s.aabbs[0][2] == 0 &&
			s.aabbs[0][3] == 1 && s.aabbs[0][4] == 1 && s.aabbs[0][5] == 1 {
			fullBlockIdx = uint16(i)
			break
		}
	}

	// build per-block shape info from PrismarineJS
	// each block is either a single int (all states same shape) or an array of ints
	type blockShapeInfo struct {
		single   int  // shape ID if all states share it
		isSingle bool // true if single
		perState []int
	}
	prismarineBlocks := make(map[string]blockShapeInfo, len(prismarine.Blocks))
	for name, raw := range prismarine.Blocks {
		var singleID int
		if err := json.Unmarshal(raw, &singleID); err == nil {
			prismarineBlocks[name] = blockShapeInfo{single: singleID, isSingle: true}
			continue
		}
		var ids []int
		if err := json.Unmarshal(raw, &ids); err != nil {
			panic(fmt.Sprintf("block %q: unexpected shape format: %s", name, string(raw)))
		}
		prismarineBlocks[name] = blockShapeInfo{perState: ids}
	}

	// find max state ID to size the flat array
	maxStateID := int32(0)
	for _, block := range blocks {
		for _, state := range block.States {
			if state.ID > maxStateID {
				maxStateID = state.ID
			}
		}
	}

	// build flat shapeByState array: state ID -> compact shape index
	// default to full block (most common), override with actual data
	shapeByState := make([]uint16, maxStateID+1)
	for i := range shapeByState {
		shapeByState[i] = fullBlockIdx
	}

	matched, unmatched := 0, 0
	for blockName, block := range blocks {
		if len(block.States) == 0 {
			continue
		}
		// strip "minecraft:" to match PrismarineJS naming
		shortName := strings.TrimPrefix(blockName, "minecraft:")

		info, ok := prismarineBlocks[shortName]
		if !ok {
			unmatched++
			continue
		}
		matched++

		baseID := block.States[0].ID
		if info.isSingle {
			compactIdx := origToCompact[info.single]
			for _, state := range block.States {
				shapeByState[state.ID] = compactIdx
			}
		} else {
			numStates := len(block.States)
			if len(info.perState) != numStates {
				fmt.Printf("warning: block %s: PrismarineJS has %d shape entries but blocks.json has %d states\n",
					blockName, len(info.perState), numStates)
			}
			for _, state := range block.States {
				offset := int(state.ID - baseID)
				if offset < len(info.perState) {
					shapeByState[state.ID] = origToCompact[info.perState[offset]]
				}
			}
		}
	}
	fmt.Printf("block shapes: %d matched, %d unmatched (defaulting to full block), %d unique shapes\n",
		matched, unmatched, len(compactShapes))

	// generate Go code
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("blocks"))

	sb.WriteString("import \"github.com/go-mclib/data/pkg/data/hitboxes\"\n\n")

	sb.WriteString(fmt.Sprintf("const fullBlockShapeIdx = %d\n\n", fullBlockIdx))

	// shapes table
	sb.WriteString(fmt.Sprintf("// shapes contains %d unique collision shapes.\n", len(compactShapes)))
	sb.WriteString("var shapes = [...][]hitboxes.AABB{\n")
	for _, s := range compactShapes {
		if len(s.aabbs) == 0 {
			sb.WriteString("\t{},\n")
			continue
		}
		sb.WriteString("\t{")
		for j, aabb := range s.aabbs {
			if j > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("{%s, %s, %s, %s, %s, %s}",
				formatFloat(aabb[0]), formatFloat(aabb[1]), formatFloat(aabb[2]),
				formatFloat(aabb[3]), formatFloat(aabb[4]), formatFloat(aabb[5])))
		}
		sb.WriteString("},\n")
	}
	sb.WriteString("}\n\n")

	// flat state -> shape index array
	sb.WriteString(fmt.Sprintf("// shapeByState maps each of %d block state IDs to a shape index.\n", len(shapeByState)))
	sb.WriteString("var shapeByState = [...]uint16{\n\t")
	for i, idx := range shapeByState {
		if i > 0 && i%20 == 0 {
			sb.WriteString("\n\t")
		}
		sb.WriteString(fmt.Sprintf("%d, ", idx))
	}
	sb.WriteString("\n}\n")

	writeFile(outPath, sb.String())
}

func formatFloat(f float64) string {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	// ensure there's a decimal point for Go float literal
	if !strings.Contains(s, ".") {
		s += ".0"
	}
	return s
}
