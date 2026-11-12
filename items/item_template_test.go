package items_test

import (
	"slices"
	"testing"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// A template is a raw DataComponentPatch: a removal the item has no default
// for (stick has no custom_name) must survive, unlike ApplyNBTComponents.
func TestItemTemplateKeepsRemovalsWithoutDefault(t *testing.T) {
	slot, err := items.TemplateFromNBT(nbt.Compound{
		"id":    nbt.String("minecraft:stick"),
		"count": nbt.Int(2),
		"components": nbt.Compound{
			"!minecraft:custom_name":    nbt.Compound{},
			"!minecraft:max_stack_size": nbt.Compound{},
			"minecraft:damage":          nbt.Int(3),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	want := []ns.VarInt{ns.VarInt(items.ComponentMaxStackSize), ns.VarInt(items.ComponentCustomName)}
	slices.Sort(want)
	if slot.Count != 2 || !slices.Equal(slot.Components.Remove, want) {
		t.Fatalf("template = %+v, want removals %v", slot, want)
	}
	stack, err := items.FromSlot(slot)
	if err != nil || stack.Components.Damage != 3 {
		t.Fatalf("template addition = %+v, %v", stack, err)
	}

	for _, bad := range []nbt.Compound{
		{"!minecraft:custom_name": nbt.Int(1)},
		{"!minecraft:no_such_component": nbt.Compound{}},
	} {
		if _, err := items.TemplateFromNBT(nbt.Compound{"id": nbt.String("minecraft:stick"), "components": bad}); err == nil {
			t.Fatalf("components %v accepted", bad)
		}
	}
}
