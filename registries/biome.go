package registries

import (
	"encoding/json"
	"sync"
)

type biomeClimateData struct {
	HasPrecipitation bool    `json:"has_precipitation"`
	Temperature      float64 `json:"temperature"`
}

// biomeClimates holds the climate of every synchronized biome by protocol ID,
// decoded once from SynchronizedRegistryData.
var biomeClimates = sync.OnceValue(func() []biomeClimateData {
	biomes := SynchronizedEntries["minecraft:worldgen/biome"]
	data := SynchronizedRegistryData["minecraft:worldgen/biome"]
	result := make([]biomeClimateData, biomes.Size())
	for name, id := range biomes.Entries() {
		_ = json.Unmarshal(data[name], &result[id])
	}
	return result
})

func biomeClimate(id int32) biomeClimateData {
	climates := biomeClimates()
	if id < 0 || int(id) >= len(climates) {
		return biomeClimateData{}
	}
	return climates[id]
}

// BiomeHasPrecipitation reports the synchronized biome's precipitation flag.
func BiomeHasPrecipitation(id int32) bool {
	return biomeClimate(id).HasPrecipitation
}

// BiomeTemperature returns the synchronized biome's base temperature.
func BiomeTemperature(id int32) float64 {
	return biomeClimate(id).Temperature
}
