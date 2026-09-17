package main

import (
	"fmt"
	"math/bits"
	"strconv"
	"strings"
)

// mcdumpBlocks is the relevant slice of mcdump/blocks.json.
type mcdumpBlocks struct {
	Blocks map[string]struct {
		States []struct {
			ID                  int32 `json:"id"`
			Shape               int   `json:"shape"`
			Replaceable         bool  `json:"replaceable"`
			IgniteOdds          uint8 `json:"igniteOdds"`
			BurnOdds            uint8 `json:"burnOdds"`
			SturdyTop           bool  `json:"sturdyTop"`
			BlocksMotion        bool  `json:"blocksMotion"`
			BlocksFluidFlow     bool  `json:"blocksFluidFlow"`
			LargeCollisionShape bool  `json:"largeCollisionShape"`
			Suffocating         bool  `json:"suffocating"`
			SturdyFaces         uint8 `json:"sturdyFaces"`
			FluidKind           uint8 `json:"fluidKind"`
			FluidAmount         uint8 `json:"fluidAmount"`
			FluidFalling        bool  `json:"fluidFalling"`
		} `json:"states"`
	} `json:"blocks"`
	Shapes           [][][]float64 `json:"shapes"`           // shape -> list of [minX,minY,minZ,maxX,maxY,maxZ]
	ShapeCoordinates [][][]float64 `json:"shapeCoordinates"` // shape -> axis X/Y/Z -> voxel planes
}

// generateBlockShapes reads the mod's blocks.json (real game collision shapes per
// block state) and emits a compact shapes table + state->shape index. Shape index
// 0 is always the empty shape so HasCollision/IsFullBlock stay valid.
func generateBlockShapes(mcdumpPath, outPath string) {
	dump := loadJSON[mcdumpBlocks](mcdumpPath)

	// re-intern shapes with the empty shape pinned at index 0.
	compact := [][][]float64{{}}
	coordinates := [][][]float64{{nil, nil, nil}}
	indexByKey := map[string]int{}
	intern := func(aabbs, coords [][]float64) int {
		if len(aabbs) == 0 {
			return 0
		}
		key := shapeKey(aabbs) + ";" + shapeKey(coords)
		if i, ok := indexByKey[key]; ok {
			return i
		}
		i := len(compact)
		compact = append(compact, aabbs)
		coordinates = append(coordinates, coords)
		indexByKey[key] = i
		return i
	}

	maxStateID := int32(0)
	for _, b := range dump.Blocks {
		for _, s := range b.States {
			if s.ID > maxStateID {
				maxStateID = s.ID
			}
		}
	}

	shapeByState := make([]uint16, maxStateID+1)
	collisionFlags := make([]uint8, maxStateID+1)
	sturdyFaces := make([]uint8, maxStateID+1)
	fluidStates := make([]uint8, maxStateID+1)
	for _, b := range dump.Blocks {
		for _, s := range b.States {
			shapeByState[s.ID] = uint16(intern(dump.Shapes[s.Shape], dump.ShapeCoordinates[s.Shape]))
			if s.BlocksMotion {
				collisionFlags[s.ID] |= 1
			}
			if s.LargeCollisionShape {
				collisionFlags[s.ID] |= 2
			}
			if s.Suffocating {
				collisionFlags[s.ID] |= 4
			}
			if isFullBlockShape(dump.Shapes[s.Shape]) {
				collisionFlags[s.ID] |= 8
			}
			if s.BlocksFluidFlow {
				collisionFlags[s.ID] |= 16
			}
			sturdyFaces[s.ID] = s.SturdyFaces
			fluidStates[s.ID] = s.FluidKind | s.FluidAmount<<2
			if s.FluidFalling {
				fluidStates[s.ID] |= 1 << 6
			}
		}
	}
	checkMinCount("block shape states", len(shapeByState), 1000)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("block_shapes"))
	sb.WriteString("import \"github.com/go-mclib/data/hitboxes\"\n\n")

	sb.WriteString(fmt.Sprintf("// shapes contains %d unique collision shapes (shape 0 is empty).\n", len(compact)))
	sb.WriteString("var shapes = [...][]hitboxes.AABB{\n")
	for _, shape := range compact {
		if len(shape) == 0 {
			sb.WriteString("\t{},\n")
			continue
		}
		sb.WriteString("\t{")
		for j, aabb := range shape {
			if j > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("{MinX: %s, MinY: %s, MinZ: %s, MaxX: %s, MaxY: %s, MaxZ: %s}",
				formatFloat(aabb[0]), formatFloat(aabb[1]), formatFloat(aabb[2]),
				formatFloat(aabb[3]), formatFloat(aabb[4]), formatFloat(aabb[5])))
		}
		sb.WriteString("},\n")
	}
	sb.WriteString("}\n\n")
	sb.WriteString("var shapeCoordinates = [...][3][]float64{\n")
	for _, axes := range coordinates {
		sb.WriteString("\t{")
		for _, coords := range axes {
			sb.WriteString("{")
			for _, coord := range coords {
				sb.WriteString(formatFloat(coord))
				sb.WriteByte(',')
			}
			sb.WriteString("},")
		}
		sb.WriteString("},\n")
	}
	sb.WriteString("}\n\n")

	sb.WriteString(fmt.Sprintf("// shapeByState maps each of %d block state IDs to a shape index.\n", len(shapeByState)))
	sb.WriteString("var shapeByState = [...]uint16{\n\t")
	for i, idx := range shapeByState {
		if i > 0 && i%20 == 0 {
			sb.WriteString("\n\t")
		}
		sb.WriteString(fmt.Sprintf("%d, ", idx))
	}
	sb.WriteString("\n}\n")
	writeUint8Array(&sb, "collisionFlags", collisionFlags)
	writeUint8Array(&sb, "sturdyFaces", sturdyFaces)
	writeUint8Array(&sb, "fluidStates", fluidStates)

	writeFile(outPath, sb.String())
	fmt.Printf("block shapes: %d states, %d unique shapes\n", len(shapeByState), len(compact))
}
func generateBlockSupport(mcdumpPath, outPath string) {
	dump := loadJSON[mcdumpBlocks](mcdumpPath)
	maxStateID := int32(0)
	for _, block := range dump.Blocks {
		for _, state := range block.States {
			maxStateID = max(maxStateID, state.ID)
		}
	}
	igniteOdds := make([]uint8, maxStateID+1)
	burnOdds := make([]uint8, maxStateID+1)
	sturdyTop := make([]uint64, maxStateID/64+1)
	replaceable := make([]uint64, maxStateID/64+1)
	for _, block := range dump.Blocks {
		for _, state := range block.States {
			igniteOdds[state.ID] = state.IgniteOdds
			burnOdds[state.ID] = state.BurnOdds
			if state.SturdyTop {
				sturdyTop[state.ID/64] |= 1 << (state.ID % 64)
			}
			if state.Replaceable {
				replaceable[state.ID/64] |= 1 << (state.ID % 64)
			}
		}
	}

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("blocks"))
	writeUint8Array(&sb, "fireIgniteOdds", igniteOdds)
	writeUint8Array(&sb, "fireBurnOdds", burnOdds)
	writeUint64Array(&sb, "sturdyTopStates", sturdyTop)
	writeUint64Array(&sb, "replaceableStates", replaceable)
	writeFile(outPath, sb.String())
	fmt.Printf("block support: %d flammable, %d sturdy-top, %d replaceable states\n", countNonzero(igniteOdds), countSetBits(sturdyTop), countSetBits(replaceable))
}

func writeUint8Array(sb *strings.Builder, name string, values []uint8) {
	sb.WriteString("var " + name + " = [...]uint8{\n\t")
	for i, value := range values {
		if i > 0 && i%32 == 0 {
			sb.WriteString("\n\t")
		}
		sb.WriteString(fmt.Sprintf("%d, ", value))
	}
	sb.WriteString("\n}\n\n")
}

func countNonzero(values []uint8) int {
	result := 0
	for _, value := range values {
		if value != 0 {
			result++
		}
	}
	return result
}

func writeUint64Array(sb *strings.Builder, name string, values []uint64) {
	sb.WriteString("var " + name + " = [...]uint64{\n\t")
	for i, value := range values {
		if i > 0 && i%8 == 0 {
			sb.WriteString("\n\t")
		}
		sb.WriteString(fmt.Sprintf("0x%x, ", value))
	}
	sb.WriteString("\n}\n\n")
}

func countSetBits(values []uint64) int {
	result := 0
	for _, value := range values {
		result += bits.OnesCount64(value)
	}
	return result
}

func shapeKey(aabbs [][]float64) string {
	var sb strings.Builder
	for _, a := range aabbs {
		for _, v := range a {
			sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
			sb.WriteByte(',')
		}
		sb.WriteByte('|')
	}
	return sb.String()
}

func isFullBlockShape(aabbs [][]float64) bool {
	return len(aabbs) == 1 && len(aabbs[0]) == 6 &&
		aabbs[0][0] == 0 && aabbs[0][1] == 0 && aabbs[0][2] == 0 &&
		aabbs[0][3] == 1 && aabbs[0][4] == 1 && aabbs[0][5] == 1
}

func formatFloat(f float64) string {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	if !strings.Contains(s, ".") {
		s += ".0"
	}
	return s
}
