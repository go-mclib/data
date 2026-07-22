package packets_test

import (
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func TestUpdateAttributesRoundTrip(t *testing.T) {
	roundTrip(t, &packets.S2CUpdateAttributes{
		EntityId: 42,
		Attributes: []packets.AttributeSnapshot{{
			Type: 17,
			Base: 30,
			Modifiers: []packets.AttributeModifier{{
				ID:        "minecraft:test_health",
				Amount:    0.5,
				Operation: 2,
			}},
		}},
	}, &packets.S2CUpdateAttributes{})

	writer := ns.NewWriter()
	if err := (&packets.S2CUpdateAttributes{EntityId: 42}).Write(writer); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if got := writer.Bytes(); len(got) != 2 || got[0] != 42 || got[1] != 0 {
		t.Fatalf("empty attribute packet wire = %x, want 2a00", got)
	}
}
