package items_test

import (
	"testing"

	"github.com/go-mclib/data/items"
)

func TestZeroValuedOptionalComponentsRoundTrip(t *testing.T) {
	zero := int32(0)
	stack := items.NewStack(items.ID("minecraft:filled_map"), 1)
	stack.Components.MapID = &zero
	stack.Components.MapPostProcessing = &zero
	slot, err := stack.ToSlot()
	if err != nil {
		t.Fatal(err)
	}
	var mapID, postProcessing bool
	for _, component := range slot.Components.Add {
		switch int32(component.ID) {
		case items.ComponentMapId:
			mapID = true
		case items.ComponentMapPostProcessing:
			postProcessing = true
		}
	}
	if !mapID || !postProcessing {
		t.Fatalf("zero-valued component patch = %+v", slot.Components)
	}
	decoded, err := items.FromSlot(slot)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Components.MapID == nil || *decoded.Components.MapID != 0 {
		t.Fatalf("decoded map ID = %v", decoded.Components.MapID)
	}
	if decoded.Components.MapPostProcessing == nil || *decoded.Components.MapPostProcessing != 0 {
		t.Fatalf("decoded post-processing = %v", decoded.Components.MapPostProcessing)
	}
	decoded.Components.MapPostProcessing = nil
	withoutPostProcessing, err := decoded.ToSlot()
	if err != nil {
		t.Fatal(err)
	}
	for _, component := range withoutPostProcessing.Components.Add {
		if int32(component.ID) == items.ComponentMapPostProcessing {
			t.Fatal("cleared map post-processing component was encoded")
		}
	}
}
