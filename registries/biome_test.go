package registries_test

import (
	"testing"

	"github.com/go-mclib/data/registries"
)

func TestBiomeHasPrecipitation(t *testing.T) {
	biomes := registries.SynchronizedEntries["minecraft:worldgen/biome"]
	if !registries.BiomeHasPrecipitation(biomes.Get("minecraft:plains")) {
		t.Fatal("plains should have precipitation")
	}
	if registries.BiomeHasPrecipitation(biomes.Get("minecraft:desert")) {
		t.Fatal("desert should not have precipitation")
	}
	if temperature := registries.BiomeTemperature(biomes.Get("minecraft:plains")); temperature != 0.8 {
		t.Fatalf("plains temperature = %v, want 0.8", temperature)
	}
	if color := registries.MobEffectColor("minecraft:speed"); color == 0 {
		t.Fatal("speed effect color is absent")
	}
}
