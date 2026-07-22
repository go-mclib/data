package items

import (
	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
	"slices"
	"strings"
)

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

// SlotMaxStackSize returns the effective max-stack-size component for a slot.
func SlotMaxStackSize(slot ns.Slot) int32 {
	limit := MaxStackSize(int32(slot.ItemID))
	for _, component := range slot.Components.Add {
		if int32(component.ID) != ComponentMaxStackSize {
			continue
		}
		value, err := ns.NewReader(component.Data).ReadVarInt()
		if err == nil {
			limit = max(int32(value), 1)
		}
	}
	for _, component := range slot.Components.Remove {
		if int32(component) == ComponentMaxStackSize {
			return 1
		}
	}
	return limit
}

// SlotResistsDamage reports whether a slot's effective damage-resistant component contains a damage type.
func SlotResistsDamage(slot ns.Slot, damageType string) bool {
	var types []string
	if defaults := DefaultComponents(int32(slot.ItemID)); defaults != nil && defaults.DamageResistant != nil {
		types = []string{defaults.DamageResistant.Types}
	}
	for _, component := range slot.Components.Add {
		if int32(component.ID) != ComponentDamageResistant {
			continue
		}
		decoded, err := readDamageTypeHolderSet(ns.NewReader(component.Data))
		if err == nil {
			types = decoded
		}
	}
	for _, component := range slot.Components.Remove {
		if int32(component) == ComponentDamageResistant {
			return false
		}
	}
	for _, entry := range types {
		if entry == damageType {
			return true
		}
		if tag := strings.TrimPrefix(entry, "#"); tag != entry &&
			slices.Contains(registries.TagData["minecraft:damage_type"][tag], damageType) {
			return true
		}
	}
	return false
}
