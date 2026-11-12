package registries_test

import (
	"testing"

	"github.com/go-mclib/data/registries"
)

func TestBlockRegistry(t *testing.T) {
	tests := []struct {
		name string
		id   int32
	}{
		{"minecraft:air", 0},
		{"minecraft:stone", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := registries.Block.Get(tt.name); got != tt.id {
				t.Errorf("Block.Get(%q) = %d, want %d", tt.name, got, tt.id)
			}
			if got := registries.Block.ByID(tt.id); got != tt.name {
				t.Errorf("Block.ByID(%d) = %q, want %q", tt.id, got, tt.name)
			}
		})
	}
}

func TestItemRegistry(t *testing.T) {
	// item registry should have entries
	if got := registries.Item.Get("minecraft:diamond"); got == -1 {
		t.Error("Item.Get(diamond) should return valid ID")
	}
	if got := registries.Item.Get("minecraft:stick"); got == -1 {
		t.Error("Item.Get(stick) should return valid ID")
	}
}

func TestEntityTypeRegistry(t *testing.T) {
	// check some well-known entity types
	if got := registries.EntityType.Get("minecraft:player"); got == -1 {
		t.Error("EntityType.Get(player) should return valid ID")
	}
	if got := registries.EntityType.Get("minecraft:zombie"); got == -1 {
		t.Error("EntityType.Get(zombie) should return valid ID")
	}
	if got := registries.EntityType.Get("minecraft:creeper"); got == -1 {
		t.Error("EntityType.Get(creeper) should return valid ID")
	}
}

func TestRegistryNotFound(t *testing.T) {
	if got := registries.Block.Get("minecraft:nonexistent"); got != -1 {
		t.Errorf("Block.Get for nonexistent = %d, want -1", got)
	}
	if got := registries.Block.ByID(-999); got != "" {
		t.Errorf("Block.ByID for invalid ID = %q, want empty string", got)
	}
}

func TestDataComponentTypeRegistry(t *testing.T) {
	// data component type registry should exist and have entries
	if got := registries.DataComponentType.Get("minecraft:damage"); got == -1 {
		t.Error("DataComponentType.Get(damage) should return valid ID")
	}
	if got := registries.DataComponentType.Get("minecraft:max_damage"); got == -1 {
		t.Error("DataComponentType.Get(max_damage) should return valid ID")
	}
	if got := registries.DataComponentType.Get("minecraft:food"); got == -1 {
		t.Error("DataComponentType.Get(food) should return valid ID")
	}
}

func TestSynchronizedRegistry(t *testing.T) {
	enchantments := registries.SynchronizedEntries["minecraft:enchantment"]
	if enchantments.ProtocolID != -1 {
		t.Errorf("synchronized registry ProtocolID = %d, want -1", enchantments.ProtocolID)
	}
	// datapack entries are sorted, so protocol IDs follow name order
	if got := enchantments.ByID(0); got != "minecraft:aqua_affinity" {
		t.Errorf("ByID(0) = %q, want minecraft:aqua_affinity", got)
	}
	for id := range int32(enchantments.Size()) {
		name := enchantments.ByID(id)
		if got := enchantments.Get(name); got != id {
			t.Fatalf("Get(ByID(%d) = %q) = %d", id, name, got)
		}
		if got := enchantments.HolderID(name); got != id+1 {
			t.Fatalf("HolderID(%q) = %d, want %d", name, got, id+1)
		}
		if got := enchantments.HolderName(id + 1); got != name {
			t.Fatalf("HolderName(%d) = %q, want %q", id+1, got, name)
		}
	}
}

func TestRegistryHolderEncoding(t *testing.T) {
	id := registries.SoundEvent.Get("minecraft:music_disc.cat")
	if got := registries.SoundEvent.HolderID("minecraft:music_disc.cat"); got != id+1 {
		t.Errorf("HolderID = %d, want %d", got, id+1)
	}
	if got := registries.SoundEvent.HolderName(id + 1); got != "minecraft:music_disc.cat" {
		t.Errorf("HolderName(%d) = %q", id+1, got)
	}
	if got := registries.SoundEvent.HolderID("minecraft:nonexistent"); got != -1 {
		t.Errorf("HolderID(unknown) = %d, want -1", got)
	}
	size := int32(registries.SoundEvent.Size())
	for _, holder := range []int32{0, -1, size + 1} {
		if got := registries.SoundEvent.HolderName(holder); got != "" {
			t.Errorf("HolderName(%d) = %q, want empty (inline or unknown)", holder, got)
		}
	}
}
