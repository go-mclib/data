package blocks

import "github.com/go-mclib/data/registries"

// ID returns the protocol ID for a block identifier, or -1 if not found.
func ID(id string) int32 {
	return registries.Block.Get(id)
}

// Name returns the identifier for a block protocol ID, or empty string if not found.
func Name(protocolID int32) string {
	return registries.Block.ByID(protocolID)
}
