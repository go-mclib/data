// go test -bench=. -benchmem ./...
package blocks_test

import (
	"testing"

	"github.com/go-mclib/data/pkg/data/blocks"
)

func BenchmarkStateIDCached(b *testing.B) {
	blockID := blocks.RedstoneWire
	props := map[string]string{
		"east":  "side",
		"north": "up",
		"power": "8",
		"south": "none",
		"west":  "side",
	}

	// warm up cache
	blocks.StateID(blockID, props)

	b.ResetTimer()
	for b.Loop() {
		blocks.StateID(blockID, props)
	}
}

func BenchmarkStateIDUncached(b *testing.B) {
	blockID := blocks.RedstoneWire
	props := map[string]string{
		"east":  "side",
		"north": "up",
		"power": "8",
		"south": "none",
		"west":  "side",
	}

	// disable caching
	blocks.SetCacheSize(0)
	defer blocks.SetCacheSize(4096)

	for b.Loop() {
		blocks.StateID(blockID, props)
	}
}

func BenchmarkStateProperties(b *testing.B) {
	stateID := 4020 // redstone_wire state

	for b.Loop() {
		blocks.StateProperties(stateID)
	}
}
