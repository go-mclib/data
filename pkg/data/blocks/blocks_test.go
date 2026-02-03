package blocks_test

import (
	"testing"

	"github.com/go-mclib/data/pkg/data/blocks"
)

func TestBlockIDLookup(t *testing.T) {
	tests := []struct {
		name string
		id   int32
	}{
		{"minecraft:air", blocks.Air},
		{"minecraft:stone", blocks.Stone},
		{"minecraft:dirt", blocks.Dirt},
		{"minecraft:oak_planks", blocks.OakPlanks},
		{"minecraft:diamond_block", blocks.DiamondBlock},
		{"minecraft:iron_block", blocks.IronBlock},
		{"minecraft:gold_block", blocks.GoldBlock},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := blocks.BlockID(tt.name); got != tt.id {
				t.Errorf("BlockID(%q) = %d, want %d", tt.name, got, tt.id)
			}
			if got := blocks.BlockName(tt.id); got != tt.name {
				t.Errorf("BlockName(%d) = %q, want %q", tt.id, got, tt.name)
			}
		})
	}
}

func TestBlockIDNotFound(t *testing.T) {
	if got := blocks.BlockID("minecraft:nonexistent_block"); got != -1 {
		t.Errorf("BlockID for nonexistent block = %d, want -1", got)
	}
	if got := blocks.BlockName(-999); got != "" {
		t.Errorf("BlockName for invalid ID = %q, want empty string", got)
	}
}
