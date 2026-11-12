package items

import (
	"testing"

	"github.com/go-mclib/data/registries"
)

func TestEveryDataComponentHasWireCodec(t *testing.T) {
	for id := range int32(MaxComponentID + 1) {
		name := registries.DataComponentType.ByID(id)
		if name != "" && componentCodecs[id] == nil {
			t.Errorf("component %d (%s) has no wire codec", id, name)
		}
	}
}

func TestEveryNBTBridgeHasWireCodec(t *testing.T) {
	for id, bridge := range nbtBridges {
		if componentCodecs[id] == nil {
			t.Errorf("component %d (%s) has an NBT bridge but no wire codec", id, registries.DataComponentType.ByID(id))
		}
		if bridge.toNBT == nil || bridge.fromNBT == nil {
			t.Errorf("component %d (%s) has an incomplete NBT bridge", id, registries.DataComponentType.ByID(id))
		}
	}
}
