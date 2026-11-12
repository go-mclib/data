package blocks_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/go-mclib/data/blocks"
)

type blockJSON struct {
	Properties map[string][]string `json:"properties"`
	States     []stateJSON         `json:"states"`
}

type stateJSON struct {
	ID         int32             `json:"id"`
	Default    bool              `json:"default,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

func loadBlocksJSON(t *testing.T) map[string]blockJSON {
	data, err := os.ReadFile("../generate/blocks.json")
	if err != nil {
		t.Fatalf("failed to read blocks.json: %v", err)
	}
	var result map[string]blockJSON
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("failed to parse blocks.json: %v", err)
	}
	return result
}

func TestAllBlockStates(t *testing.T) {
	blocksData := loadBlocksJSON(t)

	totalStates := 0
	failedStates := 0

	for blockName, blockData := range blocksData {
		blockID := blocks.ID(blockName)
		if blockID == -1 {
			t.Errorf("block %q not found in registry", blockName)
			continue
		}

		for _, state := range blockData.States {
			totalStates++
			calculated := blocks.StateID(blockID, state.Properties)
			if calculated != state.ID {
				failedStates++
				if failedStates <= 10 {
					t.Errorf("%s: expected state %d, got %d for props %v",
						blockName, state.ID, calculated, state.Properties)
				}
			}
		}
	}

	if failedStates > 10 {
		t.Errorf("... and %d more failures (total: %d/%d states failed)",
			failedStates-10, failedStates, totalStates)
	}

	t.Logf("tested %d block states", totalStates)
}

func TestDefaultStates(t *testing.T) {
	blocksData := loadBlocksJSON(t)

	for blockName, blockData := range blocksData {
		blockID := blocks.ID(blockName)
		if blockID == -1 {
			continue
		}

		// find the default state in JSON
		var expectedDefaultID int32 = -1
		for _, state := range blockData.States {
			if state.Default {
				expectedDefaultID = state.ID
				break
			}
		}

		// blocks without explicit default use first state
		if expectedDefaultID == -1 && len(blockData.States) > 0 {
			expectedDefaultID = blockData.States[0].ID
		}

		actualDefaultID := blocks.DefaultStateID(blockID)
		if actualDefaultID != expectedDefaultID {
			t.Errorf("%s: expected default state %d, got %d",
				blockName, expectedDefaultID, actualDefaultID)
		}
	}
}

func TestStateProperties(t *testing.T) {
	blocksData := loadBlocksJSON(t)

	// test a sample of states
	testedCount := 0
	for blockName, blockData := range blocksData {
		if len(blockData.Properties) == 0 {
			continue
		}

		blockID := blocks.ID(blockName)
		if blockID == -1 {
			continue
		}

		// test first and last state of each block with properties
		for _, idx := range []int{0, len(blockData.States) - 1} {
			if idx >= len(blockData.States) {
				continue
			}
			state := blockData.States[idx]

			gotBlockID, gotProps := blocks.StateProperties(state.ID)
			if gotBlockID != blockID {
				t.Errorf("StateProperties(%d): expected block %d (%s), got %d",
					state.ID, blockID, blockName, gotBlockID)
				continue
			}

			for prop, expectedValue := range state.Properties {
				if gotProps[prop] != expectedValue {
					t.Errorf("StateProperties(%d): %s expected %q=%q, got %q",
						state.ID, blockName, prop, expectedValue, gotProps[prop])
				}
			}
			testedCount++
		}

		// limit to avoid extremely long tests
		if testedCount > 2000 {
			break
		}
	}

	t.Logf("tested %d state reverse lookups", testedCount)
}

func TestBlockSupportMetadata(t *testing.T) {
	type supportDump struct {
		Blocks map[string]struct {
			States []struct {
				ID          int32 `json:"id"`
				Replaceable bool  `json:"replaceable"`
				IgniteOdds  int   `json:"igniteOdds"`
				BurnOdds    int   `json:"burnOdds"`
				SturdyTop   bool  `json:"sturdyTop"`
			} `json:"states"`
		} `json:"blocks"`
	}
	data, err := os.ReadFile("../generate/mcdump/blocks.json")
	if err != nil {
		t.Fatal(err)
	}
	var dump supportDump
	if err := json.Unmarshal(data, &dump); err != nil {
		t.Fatal(err)
	}
	for name, block := range dump.Blocks {
		for _, state := range block.States {
			if got := blocks.FireIgniteOdds(state.ID); got != state.IgniteOdds {
				t.Errorf("FireIgniteOdds(%s, %d) = %d, want %d", name, state.ID, got, state.IgniteOdds)
			}
			if got := blocks.FireBurnOdds(state.ID); got != state.BurnOdds {
				t.Errorf("FireBurnOdds(%s, %d) = %d, want %d", name, state.ID, got, state.BurnOdds)
			}
			if got := blocks.IsFaceSturdy(state.ID, blocks.DirectionUp); got != state.SturdyTop {
				t.Errorf("IsFaceSturdy(%s, %d, up) = %v, want %v", name, state.ID, got, state.SturdyTop)
			}
			if got := blocks.IsReplaceable(state.ID); got != state.Replaceable {
				t.Errorf("IsReplaceable(%s, %d) = %v, want %v", name, state.ID, got, state.Replaceable)
			}
		}
	}
}

func TestBlockFriction(t *testing.T) {
	tests := map[string]float64{
		"minecraft:stone":       0.6,
		"minecraft:ice":         0.98,
		"minecraft:packed_ice":  0.98,
		"minecraft:frosted_ice": 0.98,
		"minecraft:blue_ice":    0.989,
		"minecraft:slime_block": 0.8,
	}
	for name, want := range tests {
		stateID := blocks.DefaultStateID(blocks.ID(name))
		if got := blocks.Friction(stateID); got != want {
			t.Errorf("Friction(%s) = %v, want %v", name, got, want)
		}
		if got := blocks.StateBlockID(stateID); got != blocks.ID(name) {
			t.Errorf("StateBlockID(%s) = %d, want %d", name, got, blocks.ID(name))
		}
	}
}

func TestBlockBounceRestitution(t *testing.T) {
	tests := map[string]float64{
		"minecraft:stone":       0,
		"minecraft:slime_block": 1,
		"minecraft:red_bed":     0.75,
	}
	for name, want := range tests {
		stateID := blocks.DefaultStateID(blocks.ID(name))
		if got := blocks.BounceRestitution(stateID); got != want {
			t.Errorf("BounceRestitution(%s) = %v, want %v", name, got, want)
		}
	}
}

func TestBlockLight(t *testing.T) {
	stateOf := func(name string, props map[string]string) int32 {
		if props == nil {
			return blocks.DefaultStateID(blocks.ID(name))
		}
		return blocks.StateID(blocks.ID(name), props)
	}
	tests := []struct {
		name                string
		props               map[string]string
		emission, dampening uint8
	}{
		{"minecraft:air", nil, 0, 0},
		{"minecraft:stone", nil, 0, 15},
		{"minecraft:packed_ice", nil, 0, 15},
		{"minecraft:ice", nil, 0, 1},
		{"minecraft:glass", nil, 0, 0},
		{"minecraft:oak_leaves", nil, 0, 1},
		{"minecraft:water", nil, 0, 1},
		{"minecraft:torch", nil, 14, 0},
		{"minecraft:glowstone", nil, 15, 15},
		{"minecraft:furnace", map[string]string{"facing": "north", "lit": "true"}, 13, 15},
		{"minecraft:furnace", map[string]string{"facing": "north", "lit": "false"}, 0, 15},
	}
	for _, tt := range tests {
		state := stateOf(tt.name, tt.props)
		if got := blocks.LightEmission(state); got != tt.emission {
			t.Errorf("LightEmission(%s %v) = %d, want %d", tt.name, tt.props, got, tt.emission)
		}
		if got := blocks.LightDampening(state); got != tt.dampening {
			t.Errorf("LightDampening(%s %v) = %d, want %d", tt.name, tt.props, got, tt.dampening)
		}
	}
}
