package blocks

import (
	"slices"
	"strings"
)

// StateID calculates the state protocol ID from a block ID and property values.
// Returns -1 if the block ID is invalid or properties are missing/invalid.
func StateID(blockID int32, props map[string]string) int32 {
	state := blockStates[blockID]
	if state == nil {
		return -1
	}

	offset := int32(0)
	multiplier := int32(1)
	for i := len(state.Properties) - 1; i >= 0; i-- {
		p := &state.Properties[i]
		propValue, ok := props[p.Name]
		if !ok {
			return -1
		}
		valueIdx := slices.Index(p.Values, propValue)
		if valueIdx < 0 {
			return -1
		}
		offset += int32(valueIdx) * multiplier
		multiplier *= int32(len(p.Values))
	}

	return state.BaseID + offset
}

// stateRange returns the stateRanges entry containing stateID.
func stateRange(stateID int32) (*stateRangeEntry, bool) {
	i, found := slices.BinarySearchFunc(stateRanges, stateID, func(e stateRangeEntry, id int32) int {
		switch {
		case id < e.BaseID:
			return 1
		case id >= e.EndID:
			return -1
		}
		return 0
	})
	if !found {
		return nil, false
	}
	return &stateRanges[i], true
}

// StateBlockID returns the block protocol ID for a state ID without decoding its properties.
func StateBlockID(stateID int32) int32 {
	if r, ok := stateRange(stateID); ok {
		return r.BlockID
	}
	return -1
}

// Friction returns the vanilla friction coefficient (Block.getFriction) for a block state.
func Friction(stateID int32) float64 {
	blockID := StateBlockID(stateID)
	if blockID < 0 || int(blockID) >= len(blockFriction) {
		return 0.6 // vanilla BlockBehaviour.Properties default
	}
	return blockFriction[blockID]
}

// BounceRestitution returns the vanilla bounce coefficient for a block state.
func BounceRestitution(stateID int32) float64 {
	name := Name(StateBlockID(stateID))
	if name == "minecraft:slime_block" {
		return 1
	}
	if strings.HasSuffix(name, "_bed") {
		return 0.75
	}
	return 0
}

// StateProperties returns the block ID and property values for a given state ID.
// Uses binary search for O(log n) block lookup. Returns -1 and nil if the state ID is invalid.
func StateProperties(stateID int32) (blockID int32, props map[string]string) {
	r, ok := stateRange(stateID)
	if !ok {
		return -1, nil
	}
	blockID = r.BlockID
	state := blockStates[blockID]
	offset := stateID - state.BaseID

	props = make(map[string]string, len(state.Properties))
	for i := len(state.Properties) - 1; i >= 0; i-- {
		p := &state.Properties[i]
		cardinality := int32(len(p.Values))
		props[p.Name] = p.Values[offset%cardinality]
		offset /= cardinality
	}
	return blockID, props
}

// DefaultStateID returns the default state ID for a block, or -1 if invalid.
func DefaultStateID(blockID int32) int32 {
	state := blockStates[blockID]
	if state == nil {
		return -1
	}
	return state.DefaultID
}

// IsFlammable reports whether vanilla fire can burn the block state.
func IsFlammable(stateID int32) bool {
	return FireIgniteOdds(stateID) > 0
}

// FireIgniteOdds returns the state's vanilla chance to catch fire.
func FireIgniteOdds(stateID int32) int {
	if stateID < 0 || int(stateID) >= len(fireIgniteOdds) {
		return 0
	}
	return int(fireIgniteOdds[stateID])
}

// FireBurnOdds returns the state's vanilla chance to burn away.
func FireBurnOdds(stateID int32) int {
	if stateID < 0 || int(stateID) >= len(fireBurnOdds) {
		return 0
	}
	return int(fireBurnOdds[stateID])
}

// IsReplaceable reports whether vanilla permits replacing the block state.
func IsReplaceable(stateID int32) bool {
	if stateID < 0 || int(stateID/64) >= len(replaceableStates) {
		return false
	}
	return replaceableStates[stateID/64]&(uint64(1)<<(stateID%64)) != 0
}

// MapColor returns the vanilla base map-color ID for a block state.
func MapColor(stateID int32) uint8 {
	if stateID < 0 || int(stateID) >= len(mapColors) {
		return 0
	}
	return mapColors[stateID]
}

// LightEmission returns the light level (0-15) the block state emits
// (BlockState.getLightEmission).
func LightEmission(stateID int32) uint8 {
	if stateID < 0 || int(stateID) >= len(lightEmission) {
		return 0
	}
	return lightEmission[stateID]
}

// LightDampening returns how much light the block state absorbs (0-15,
// BlockState.getLightDampening): 15 for solid-render blocks, 0 for blocks
// that let skylight pass down, 1 otherwise.
func LightDampening(stateID int32) uint8 {
	if stateID < 0 || int(stateID) >= len(lightDampening) {
		return 0
	}
	return lightDampening[stateID]
}

// blockStateData holds the data needed to calculate state IDs.
type blockStateData struct {
	BaseID     int32
	DefaultID  int32
	Properties []blockProperty
}

type blockProperty struct {
	Name   string
	Values []string
}

// stateRangeEntry maps a range of state IDs to a block ID for O(log n) lookup.
type stateRangeEntry struct {
	BaseID, EndID, BlockID int32
}
