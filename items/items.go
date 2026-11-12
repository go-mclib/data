package items

import (
	"sync"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// ID returns the protocol ID for an item identifier, or -1 if not found.
func ID(id string) int32 {
	return registries.Item.Get(id)
}

// Name returns the identifier for an item protocol ID, or empty string if not found.
func Name(protocolID int32) string {
	return registries.Item.ByID(protocolID)
}

// defaultPotDecorationLoaders decode the generated default pot decorations on
// first use. Decoding resolves the decoration items' own defaults, so it cannot
// run during package initialization; each item has its own once so decoding
// one item's decorations may look up other items' defaults.
var defaultPotDecorationLoaders = make(map[int32]func(), len(defaultPotDecorationData))

func init() {
	for itemID, data := range defaultPotDecorationData {
		defaultPotDecorationLoaders[itemID] = sync.OnceFunc(func() {
			defaultComponents[itemID].PotDecorations = defaultPotDecorations(data)
		})
	}
}

// DefaultComponents returns the default components for an item, or nil if not found.
// The result is shared and must not be modified; Clone it first.
func DefaultComponents(itemID int32) *Components {
	if load := defaultPotDecorationLoaders[itemID]; load != nil {
		load()
	}
	return defaultComponents[itemID]
}

// MaxStackSize returns an item's stack limit from its default components
// (64 when the item is unknown).
func MaxStackSize(itemID int32) int32 {
	if c := DefaultComponents(itemID); c != nil && c.MaxStackSize > 0 {
		return c.MaxStackSize
	}
	return 64
}

// SlotMaxStackSize returns the effective max-stack-size component for a slot.
func SlotMaxStackSize(slot ns.Slot) int32 {
	data, added, removed := patchedComponent(slot, ComponentMaxStackSize)
	if removed {
		return 1
	}
	if added {
		if value, err := ns.NewReader(data).ReadVarInt(); err == nil {
			return max(int32(value), 1)
		}
	}
	return MaxStackSize(int32(slot.ItemID))
}
