package packets_test

import (
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestMapItemDataRoundTrip(t *testing.T) {
	name := ns.TextComponent{Text: "home"}
	decorations := []packets.MapDecoration{
		{Type: 0, X: -12, Y: 34, Rotation: 15},
		{Type: 1, X: 127, Y: -128, Rotation: 4, Name: &name},
	}
	roundTrip(t, &packets.S2CMapItemData{
		MapID:       17,
		Scale:       2,
		Locked:      true,
		Decorations: &decorations,
		ColorPatch: &packets.MapPatch{
			Width: 2, Height: 2, StartX: 63, StartY: 64,
			Colors: ns.ByteArray{1, 2, 3, 4},
		},
	}, &packets.S2CMapItemData{})
}

func TestMapItemDataEmptyUpdateRoundTrip(t *testing.T) {
	roundTrip(t, &packets.S2CMapItemData{MapID: 3, Scale: 4}, &packets.S2CMapItemData{})
}
