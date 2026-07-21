package items

import "github.com/go-mclib/data/registries"

func init() {
	for id, color := range dyeNames {
		components := defaultComponents[ItemID("minecraft:"+color+"_dye")]
		if components != nil {
			value := int32(id)
			components.Dye = &value
		}
		banner := defaultComponents[ItemID("minecraft:"+color+"_banner")]
		if banner != nil {
			layers := []BannerPatternLayer{}
			banner.BannerPatterns = &layers
		}
	}
	if shield := defaultComponents[ItemID("minecraft:shield")]; shield != nil {
		layers := []BannerPatternLayer{}
		shield.BannerPatterns = &layers
	}
	if pot := defaultComponents[ItemID("minecraft:decorated_pot")]; pot != nil {
		pot.PotDecorations = &PotDecorations{
			Back: "minecraft:brick", Left: "minecraft:brick", Right: "minecraft:brick", Front: "minecraft:brick",
		}
	}
}

// ItemID returns the protocol ID for an item identifier, or -1 if not found.
func ItemID(id string) int32 {
	return registries.Item.Get(id)
}

// ItemName returns the identifier for an item protocol ID, or empty string if not found.
func ItemName(protocolID int32) string {
	return registries.Item.ByID(protocolID)
}

// DefaultComponents returns the default components for an item, or nil if not found.
func DefaultComponents(itemID int32) *Components {
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
