package entities_test

import (
	"testing"

	"github.com/go-mclib/data/pkg/data/entities"
)

func TestEntityTypeIDLookup(t *testing.T) {
	tests := []struct {
		name string
		id   int32
	}{
		{"minecraft:player", entities.Player},
		{"minecraft:zombie", entities.Zombie},
		{"minecraft:creeper", entities.Creeper},
		{"minecraft:skeleton", entities.Skeleton},
		{"minecraft:spider", entities.Spider},
		{"minecraft:pig", entities.Pig},
		{"minecraft:cow", entities.Cow},
		{"minecraft:sheep", entities.Sheep},
		{"minecraft:chicken", entities.Chicken},
		{"minecraft:villager", entities.Villager},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := entities.EntityTypeID(tt.name); got != tt.id {
				t.Errorf("EntityTypeID(%q) = %d, want %d", tt.name, got, tt.id)
			}
			if got := entities.EntityTypeName(tt.id); got != tt.name {
				t.Errorf("EntityTypeName(%d) = %q, want %q", tt.id, got, tt.name)
			}
		})
	}
}

func TestEntityTypeIDNotFound(t *testing.T) {
	if got := entities.EntityTypeID("minecraft:nonexistent_entity"); got != -1 {
		t.Errorf("EntityTypeID for nonexistent entity = %d, want -1", got)
	}
	if got := entities.EntityTypeName(-999); got != "" {
		t.Errorf("EntityTypeName for invalid ID = %q, want empty string", got)
	}
}

func TestCommonEntityTypes(t *testing.T) {
	// verify some well-known entity types have non-negative IDs
	commonTypes := []int32{
		entities.Player,
		entities.Zombie,
		entities.Creeper,
		entities.Skeleton,
		entities.EnderDragon,
		entities.Wither,
		entities.Item,
		entities.ExperienceOrb,
		entities.Arrow,
		entities.Fireball,
	}

	for _, id := range commonTypes {
		if id < 0 {
			t.Errorf("entity type ID %d should be non-negative", id)
		}
		name := entities.EntityTypeName(id)
		if name == "" {
			t.Errorf("entity type ID %d should have a name", id)
		}
	}
}

func TestMetadataSerializerConstants(t *testing.T) {
	// verify well-known serializer constants
	tests := []struct {
		name string
		id   int32
	}{
		{"BYTE", entities.SerializerBYTE},
		{"INT", entities.SerializerINT},
		{"FLOAT", entities.SerializerFLOAT},
		{"STRING", entities.SerializerSTRING},
		{"BOOLEAN", entities.SerializerBOOLEAN},
		{"ROTATIONS", entities.SerializerROTATIONS},
		{"BLOCK_POS", entities.SerializerBLOCK_POS},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.id < 0 {
				t.Errorf("serializer %s should have non-negative ID", tt.name)
			}
		})
	}
}

func TestMetadataFieldIndices(t *testing.T) {
	// verify base Entity metadata indices
	if entities.EntityIndexFlags != 0 {
		t.Errorf("EntityIndexFlags = %d, want 0", entities.EntityIndexFlags)
	}
	if entities.EntityIndexAirSupply != 1 {
		t.Errorf("EntityIndexAirSupply = %d, want 1", entities.EntityIndexAirSupply)
	}

	// verify Player-specific indices start after LivingEntity
	if entities.PlayerIndexAdditionalHearts < 8 {
		t.Errorf("PlayerIndexAdditionalHearts = %d, should be >= 8", entities.PlayerIndexAdditionalHearts)
	}

	// verify Creeper-specific indices
	if entities.CreeperIndexSwellDir < 8 {
		t.Errorf("CreeperIndexSwellDir = %d, should be >= 8", entities.CreeperIndexSwellDir)
	}
}

func TestMetadataOperations(t *testing.T) {
	var m entities.Metadata

	// test Set and Get
	m.Set(0, entities.SerializerBYTE, []byte{0x01})
	m.Set(1, entities.SerializerINT, []byte{0x64}) // VarInt 100

	if data := m.Get(0); data == nil || data[0] != 0x01 {
		t.Errorf("Get(0) failed")
	}
	if data := m.Get(1); data == nil || data[0] != 0x64 {
		t.Errorf("Get(1) failed")
	}
	if data := m.Get(99); data != nil {
		t.Errorf("Get(99) should return nil for missing index")
	}

	// test update
	m.Set(0, entities.SerializerBYTE, []byte{0x02})
	if data := m.Get(0); data == nil || data[0] != 0x02 {
		t.Errorf("Set update failed")
	}
}
