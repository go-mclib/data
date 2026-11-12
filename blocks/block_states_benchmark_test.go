// go test -bench=. -benchmem ./...
package blocks_test

import (
	"testing"

	"github.com/go-mclib/data/blocks"
)

func BenchmarkStateID(b *testing.B) {
	blockID := int(blocks.ID("minecraft:redstone_wire"))
	props := map[string]string{
		"east":  "side",
		"north": "up",
		"power": "8",
		"south": "none",
		"west":  "side",
	}

	for b.Loop() {
		blocks.StateID(int32(blockID), props)
	}
}

func BenchmarkStateProperties(b *testing.B) {
	stateID := 4020 // redstone_wire state

	for b.Loop() {
		blocks.StateProperties(int32(stateID))
	}
}
