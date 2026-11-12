package items_test

import (
	"testing"

	"github.com/go-mclib/data/blocks"
	"github.com/go-mclib/data/items"
)

func TestContextFuelProviders(t *testing.T) {
	coal := items.DefaultComponents(items.ID("minecraft:coal")).CookingFuel
	for _, test := range []struct {
		block string
		burn  int32
		speed float32
	}{
		{"minecraft:furnace", 1600, 1}, {"minecraft:blast_furnace", 800, 2}, {"minecraft:smoker", 800, 2},
	} {
		state := blocks.DefaultStateID(blocks.ID(test.block))
		if got := coal.BurnTime.Resolve(state, -1); got != test.burn {
			t.Fatalf("%s coal duration = %d", test.block, got)
		}
		if got := coal.SpeedMultiplier.Resolve(state, -1); got != test.speed {
			t.Fatalf("%s coal speed = %g", test.block, got)
		}
	}
	fast := blocks.DefaultStateID(blocks.ID("minecraft:smoker"))
	kelp := items.DefaultComponents(items.ID("minecraft:dried_kelp_block")).CookingFuel
	if got := kelp.BurnTime.Resolve(fast, -1); got != 2000 {
		t.Fatalf("integer provider division = %d, want truncation", got)
	}
	blaze := items.DefaultComponents(items.ID("minecraft:blaze_powder")).BrewingFuel
	if got := blaze.Uses.Resolve(fast, -1); got != 20 {
		t.Fatalf("brewing uses = %d", got)
	}
}

func TestContextMissingReferenceFallbackAndConstants(t *testing.T) {
	if got := (items.ResolvableInt{Reference: "test:missing"}).Resolve(0, 17); got != 17 {
		t.Fatalf("missing int = %d", got)
	}
	if got := (items.ResolvableFloat{Reference: "test:missing"}).Resolve(0, 1.5); got != 1.5 {
		t.Fatalf("missing float = %g", got)
	}
	if got := (items.ResolvableInt{Value: 0}).Resolve(0, 17); got != 0 {
		t.Fatalf("explicit zero = %d", got)
	}
	if got := (items.ResolvableFloat{Value: -2}).Resolve(0, 1); got != -2 {
		t.Fatalf("explicit negative speed = %g", got)
	}
}
