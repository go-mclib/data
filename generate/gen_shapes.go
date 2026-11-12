package main

import (
	"cmp"
	"fmt"
	"math/bits"
	"slices"
	"strconv"
	"strings"
)

// mcdumpState is one block state of mcdump/blocks.json.
type mcdumpState struct {
	ID                  int32 `json:"id"`
	Shape               int   `json:"shape"`
	Replaceable         bool  `json:"replaceable"`
	IgniteOdds          uint8 `json:"igniteOdds"`
	BurnOdds            uint8 `json:"burnOdds"`
	BlocksMotion        bool  `json:"blocksMotion"`
	BlocksFluidFlow     bool  `json:"blocksFluidFlow"`
	LargeCollisionShape bool  `json:"largeCollisionShape"`
	Suffocating         bool  `json:"suffocating"`
	SturdyFaces         uint8 `json:"sturdyFaces"`
	FluidKind           uint8 `json:"fluidKind"`
	FluidAmount         uint8 `json:"fluidAmount"`
	FluidFalling        bool  `json:"fluidFalling"`
	LightEmission       uint8 `json:"lightEmission"`
	LightDampening      uint8 `json:"lightDampening"`
	MapColor            uint8 `json:"mapColor"`
}

// mcdumpBlock is one block of mcdump/blocks.json.
type mcdumpBlock struct {
	ProtocolID   int32         `json:"protocolId"`
	Hardness     float64       `json:"hardness"`
	RequiresTool bool          `json:"requiresTool"`
	Friction     float64       `json:"friction"`
	States       []mcdumpState `json:"states"`
}

// mcdumpBlocks is mcdump/blocks.json, loaded once and shared by the block
// generators.
type mcdumpBlocks struct {
	Blocks           map[string]mcdumpBlock `json:"blocks"`
	Shapes           [][][]float64          `json:"shapes"`           // shape -> list of [minX,minY,minZ,maxX,maxY,maxZ]
	ShapeCoordinates [][][]float64          `json:"shapeCoordinates"` // shape -> axis X/Y/Z -> voxel planes
}

// states returns every block state ordered by state ID, so generated tables
// (including the interned shape order) do not depend on map iteration.
func (d *mcdumpBlocks) states() []mcdumpState {
	var states []mcdumpState
	for _, block := range d.Blocks {
		states = append(states, block.States...)
	}
	slices.SortFunc(states, func(a, b mcdumpState) int { return cmp.Compare(a.ID, b.ID) })
	for i, state := range states {
		if state.ID != int32(i) {
			panic(fmt.Sprintf("mcdump block states are not contiguous: state %d at index %d", state.ID, i))
		}
	}
	return states
}

// Bit layouts of the generated block shape tables; emitted as constants so
// blocks/shapes.go reads them by name.
const (
	flagBlocksMotion = 1 << iota
	flagLargeCollisionShape
	flagSuffocating
	flagFullBlock
	flagBlocksFluidFlow
)

const (
	fluidKindMask    = 3
	fluidAmountShift = 2
	fluidAmountMask  = 15
	fluidFalling     = 1 << 6
)

// generateBlockShapes emits a compact shapes table + state->shape index from
// the game's collision shapes, plus per-state collision flags, sturdy faces
// and fluid states. Shape index 0 is always the empty shape so
// HasCollision/IsFullBlock stay valid.
func generateBlockShapes(dump *mcdumpBlocks, states []mcdumpState, outPath string) {
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

	shapeByState := make([]uint16, len(states))
	collisionFlags := make([]uint8, len(states))
	sturdyFaces := make([]uint8, len(states))
	fluidStates := make([]uint8, len(states))
	for _, s := range states {
		shapeByState[s.ID] = uint16(intern(dump.Shapes[s.Shape], dump.ShapeCoordinates[s.Shape]))
		var flags uint8
		if s.BlocksMotion {
			flags |= flagBlocksMotion
		}
		if s.LargeCollisionShape {
			flags |= flagLargeCollisionShape
		}
		if s.Suffocating {
			flags |= flagSuffocating
		}
		if isFullBlockShape(dump.Shapes[s.Shape]) {
			flags |= flagFullBlock
		}
		if s.BlocksFluidFlow {
			flags |= flagBlocksFluidFlow
		}
		collisionFlags[s.ID] = flags
		sturdyFaces[s.ID] = s.SturdyFaces
		fluidStates[s.ID] = s.FluidKind | s.FluidAmount<<fluidAmountShift
		if s.FluidFalling {
			fluidStates[s.ID] |= fluidFalling
		}
	}
	checkMinCount("block shape states", len(shapeByState), 1000)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("blocks"))

	sb.WriteString("// collisionFlags bits.\nconst (\n")
	fmt.Fprintf(&sb, "\tflagBlocksMotion = %d\n", flagBlocksMotion)
	fmt.Fprintf(&sb, "\tflagLargeCollisionShape = %d\n", flagLargeCollisionShape)
	fmt.Fprintf(&sb, "\tflagSuffocating = %d\n", flagSuffocating)
	fmt.Fprintf(&sb, "\tflagFullBlock = %d\n", flagFullBlock)
	fmt.Fprintf(&sb, "\tflagBlocksFluidFlow = %d\n", flagBlocksFluidFlow)
	sb.WriteString(")\n\n")
	sb.WriteString("// fluidStates layout: kind | amount<<fluidAmountShift | fluidFalling.\nconst (\n")
	fmt.Fprintf(&sb, "\tfluidKindMask = %d\n", fluidKindMask)
	fmt.Fprintf(&sb, "\tfluidAmountShift = %d\n", fluidAmountShift)
	fmt.Fprintf(&sb, "\tfluidAmountMask = %d\n", fluidAmountMask)
	fmt.Fprintf(&sb, "\tfluidFalling = %d\n", fluidFalling)
	sb.WriteString(")\n\n")

	fmt.Fprintf(&sb, "// shapes contains %d unique collision shapes (shape 0 is empty).\n", len(compact))
	sb.WriteString("var shapes = [...][]AABB{\n")
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
			fmt.Fprintf(&sb, "{MinX: %s, MinY: %s, MinZ: %s, MaxX: %s, MaxY: %s, MaxZ: %s}",
				formatFloat(aabb[0]), formatFloat(aabb[1]), formatFloat(aabb[2]),
				formatFloat(aabb[3]), formatFloat(aabb[4]), formatFloat(aabb[5]))
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

	fmt.Fprintf(&sb, "// shapeByState maps each of %d block state IDs to a shape index.\n", len(shapeByState))
	sb.WriteString("var shapeByState = [...]uint16{\n\t")
	for i, idx := range shapeByState {
		if i > 0 && i%20 == 0 {
			sb.WriteString("\n\t")
		}
		fmt.Fprintf(&sb, "%d, ", idx)
	}
	sb.WriteString("\n}\n")
	writeUint8Array(&sb, "collisionFlags", collisionFlags)
	writeUint8Array(&sb, "sturdyFaces", sturdyFaces)
	writeUint8Array(&sb, "fluidStates", fluidStates)

	writeFile(outPath, sb.String())
	fmt.Printf("block shapes: %d states, %d unique shapes\n", len(shapeByState), len(compact))
}

// generateBlockSupport emits the per-state fire, map-color, replaceability and
// light tables plus per-block friction into the blocks package.
func generateBlockSupport(dump *mcdumpBlocks, states []mcdumpState, outPath string) {
	igniteOdds := make([]uint8, len(states))
	burnOdds := make([]uint8, len(states))
	mapColors := make([]uint8, len(states))
	lightEmission := make([]uint8, len(states))
	lightDampening := make([]uint8, len(states))
	replaceable := make([]uint64, (len(states)+63)/64)
	for _, state := range states {
		igniteOdds[state.ID] = state.IgniteOdds
		burnOdds[state.ID] = state.BurnOdds
		mapColors[state.ID] = state.MapColor
		lightEmission[state.ID] = state.LightEmission
		lightDampening[state.ID] = state.LightDampening
		if state.Replaceable {
			replaceable[state.ID/64] |= 1 << (state.ID % 64)
		}
	}

	var maxBlockID int32
	for _, block := range dump.Blocks {
		maxBlockID = max(maxBlockID, block.ProtocolID)
	}
	friction := make([]float64, maxBlockID+1)
	for _, block := range dump.Blocks {
		friction[block.ProtocolID] = block.Friction
	}

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("blocks"))
	writeUint8Array(&sb, "fireIgniteOdds", igniteOdds)
	writeUint8Array(&sb, "fireBurnOdds", burnOdds)
	writeUint8Array(&sb, "mapColors", mapColors)
	writeUint8Array(&sb, "lightEmission", lightEmission)
	writeUint8Array(&sb, "lightDampening", lightDampening)
	writeUint64Array(&sb, "replaceableStates", replaceable)
	sb.WriteString("// blockFriction is indexed by block protocol ID.\n")
	sb.WriteString("var blockFriction = [...]float64{\n\t")
	for i, value := range friction {
		if i > 0 && i%16 == 0 {
			sb.WriteString("\n\t")
		}
		sb.WriteString(formatFloat(value) + ", ")
	}
	sb.WriteString("\n}\n")
	writeFile(outPath, sb.String())
	fmt.Printf("block support: %d flammable, %d replaceable states, %d mapped states, %d light-emitting states\n", countNonzero(igniteOdds), countSetBits(replaceable), countNonzero(mapColors), countNonzero(lightEmission))
}

func writeUint8Array(sb *strings.Builder, name string, values []uint8) {
	sb.WriteString("var " + name + " = [...]uint8{\n\t")
	for i, value := range values {
		if i > 0 && i%32 == 0 {
			sb.WriteString("\n\t")
		}
		fmt.Fprintf(sb, "%d, ", value)
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
		fmt.Fprintf(sb, "0x%x, ", value)
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
