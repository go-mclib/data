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
	for i := 0; i < b.N; i++ {
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

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blocks.StateID(blockID, props)
	}
}

func BenchmarkStateProperties(b *testing.B) {
	stateID := int32(4020) // redstone_wire state

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blocks.StateProperties(stateID)
	}
}
