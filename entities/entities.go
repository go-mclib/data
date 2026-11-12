package entities

import "github.com/go-mclib/data/registries"

// TypeID returns the protocol ID for an entity type identifier, or -1 if not found.
func TypeID(name string) int32 {
	return registries.EntityType.Get(name)
}

// TypeName returns the identifier for an entity type protocol ID, or empty string if not found.
func TypeName(protocolID int32) string {
	return registries.EntityType.ByID(protocolID)
}

// misc entities that extend LivingEntity and can be attacked
var attackableMisc = map[string]bool{
	"minecraft:player":       true,
	"minecraft:armor_stand":  true,
	"minecraft:mannequin":    true,
	"minecraft:copper_golem": true,
	"minecraft:iron_golem":   true,
	"minecraft:snow_golem":   true,
}

// IsAttackable returns true if the entity type is a LivingEntity (has health and can be attacked).
func IsAttackable(name string) bool {
	cat := entityCategory[name]
	if cat == "" {
		return false
	}
	if cat != "misc" {
		return true
	}
	return attackableMisc[name]
}
