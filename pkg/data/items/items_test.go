package items_test

import (
	"testing"

	"github.com/go-mclib/data/pkg/data/items"
)

func TestItemIDLookup(t *testing.T) {
	tests := []struct {
		name string
		id   int32
	}{
		{"minecraft:diamond_sword", items.DiamondSword},
		{"minecraft:iron_sword", items.IronSword},
		{"minecraft:apple", items.Apple},
		{"minecraft:golden_apple", items.GoldenApple},
		{"minecraft:diamond_pickaxe", items.DiamondPickaxe},
		{"minecraft:iron_pickaxe", items.IronPickaxe},
		{"minecraft:stick", items.Stick},
		{"minecraft:diamond", items.Diamond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := items.ItemID(tt.name); got != tt.id {
				t.Errorf("ItemID(%q) = %d, want %d", tt.name, got, tt.id)
			}
			if got := items.ItemName(tt.id); got != tt.name {
				t.Errorf("ItemName(%d) = %q, want %q", tt.id, got, tt.name)
			}
		})
	}
}

func TestItemIDNotFound(t *testing.T) {
	if got := items.ItemID("minecraft:nonexistent_item"); got != -1 {
		t.Errorf("ItemID for nonexistent item = %d, want -1", got)
	}
	if got := items.ItemName(-999); got != "" {
		t.Errorf("ItemName for invalid ID = %q, want empty string", got)
	}
}

func TestDefaultComponents(t *testing.T) {
	// apple has food component
	apple := items.DefaultComponents(items.Apple)
	if apple == nil {
		t.Fatal("DefaultComponents(Apple) = nil")
	}
	if apple.Food == nil {
		t.Error("Apple should have Food component")
	} else if apple.Food.Nutrition != 4 {
		t.Errorf("Apple nutrition = %d, want 4", apple.Food.Nutrition)
	}

	// diamond sword has max damage (durability)
	sword := items.DefaultComponents(items.DiamondSword)
	if sword == nil {
		t.Fatal("DefaultComponents(DiamondSword) = nil")
	}
	if sword.MaxDamage != 1561 {
		t.Errorf("DiamondSword max damage = %d, want 1561", sword.MaxDamage)
	}
}

func TestComponentConstants(t *testing.T) {
	// verify some well-known component IDs exist
	if items.ComponentDamage < 0 {
		t.Error("ComponentDamage should be non-negative")
	}
	if items.ComponentMaxDamage < 0 {
		t.Error("ComponentMaxDamage should be non-negative")
	}
	if items.ComponentFood < 0 {
		t.Error("ComponentFood should be non-negative")
	}
	if items.MaxComponentID < items.ComponentFood {
		t.Error("MaxComponentID should be >= ComponentFood")
	}
}
