package registries_test

import (
	"slices"
	"testing"

	"github.com/go-mclib/data/registries"
)

func TestBiomeHasPrecipitation(t *testing.T) {
	entries := registries.SynchronizedEntries["minecraft:worldgen/biome"]
	if !registries.BiomeHasPrecipitation(int32(slices.Index(entries, "minecraft:plains"))) {
		t.Fatal("plains should have precipitation")
	}
	if registries.BiomeHasPrecipitation(int32(slices.Index(entries, "minecraft:desert"))) {
		t.Fatal("desert should not have precipitation")
	}
	if temperature := registries.BiomeTemperature(int32(slices.Index(entries, "minecraft:plains"))); temperature != 0.8 {
		t.Fatalf("plains temperature = %v, want 0.8", temperature)
	}
	if color := registries.MobEffectColor("minecraft:speed"); color == 0 {
		t.Fatal("speed effect color is absent")
	}
}
