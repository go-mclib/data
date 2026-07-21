package registries

import "encoding/json"

// BiomeHasPrecipitation reports the synchronized biome's precipitation flag.
func BiomeHasPrecipitation(id int32) bool {
	entries := SynchronizedEntries["minecraft:worldgen/biome"]
	if id < 0 || int(id) >= len(entries) {
		return false
	}
	name := entries[id]
	raw := SynchronizedRegistryData["minecraft:worldgen/biome"][name]
	if len(raw) == 0 {
		return false
	}
	var value struct {
		HasPrecipitation bool `json:"has_precipitation"`
	}
	return json.Unmarshal(raw, &value) == nil && value.HasPrecipitation
}

// BiomeTemperature returns the synchronized biome's base temperature.
func BiomeTemperature(id int32) float64 {
	entries := SynchronizedEntries["minecraft:worldgen/biome"]
	if id < 0 || int(id) >= len(entries) {
		return 0
	}
	var value struct {
		Temperature float64 `json:"temperature"`
	}
	raw := SynchronizedRegistryData["minecraft:worldgen/biome"][entries[id]]
	if json.Unmarshal(raw, &value) != nil {
		return 0
	}
	return value.Temperature
}

var mobEffectColors = map[string]int32{
	"minecraft:speed": 3402751, "minecraft:slowness": 9154528, "minecraft:haste": 14270531,
	"minecraft:mining_fatigue": 4866583, "minecraft:strength": 16762624, "minecraft:instant_health": 16262179,
	"minecraft:instant_damage": 11101546, "minecraft:jump_boost": 16646020, "minecraft:nausea": 5578058,
	"minecraft:regeneration": 13458603, "minecraft:resistance": 9520880, "minecraft:fire_resistance": 16750848,
	"minecraft:water_breathing": 10017472, "minecraft:invisibility": 16185078, "minecraft:blindness": 2039587,
	"minecraft:night_vision": 12779366, "minecraft:hunger": 5797459, "minecraft:weakness": 4738376,
	"minecraft:poison": 8889187, "minecraft:wither": 7561558, "minecraft:health_boost": 16284963,
	"minecraft:absorption": 2445989, "minecraft:saturation": 16262179, "minecraft:glowing": 9740385,
	"minecraft:levitation": 13565951, "minecraft:luck": 5882118, "minecraft:unluck": 12624973,
	"minecraft:slow_falling": 15978425, "minecraft:conduit_power": 1950417, "minecraft:dolphins_grace": 8954814,
	"minecraft:bad_omen": 745784, "minecraft:hero_of_the_village": 4521796, "minecraft:darkness": 2696993,
	"minecraft:trial_omen": 1484454, "minecraft:raid_omen": 14565464, "minecraft:wind_charged": 12438015,
	"minecraft:weaving": 7891290, "minecraft:oozing": 10092451, "minecraft:infested": 9214860,
	"minecraft:breath_of_the_nautilus": 65518,
}

// MobEffectColor returns the vanilla effect's RGB color.
func MobEffectColor(name string) int32 { return mobEffectColors[name] }
