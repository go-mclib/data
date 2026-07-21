package items

import (
	"testing"

	"github.com/go-mclib/data/registries"
)

func TestEveryDataComponentHasWireCodec(t *testing.T) {
	for id := int32(0); id <= MaxComponentID; id++ {
		name := registries.DataComponentType.ByID(id)
		if name != "" && componentCodecs[id] == nil {
			t.Errorf("component %d (%s) has no wire codec", id, name)
		}
	}
}
