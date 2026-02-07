package packets_test

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/go-mclib/data/pkg/data/items"
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
	"github.com/stretchr/testify/assert"
)

type packetsToBytes map[jp.Packet][]byte

// maps packet structs to their data bytes (excluding length and id)
var capturedPackets packetsToBytes = make(packetsToBytes)

// attribute modifiers used in test items
var poSwordAttribs = []items.AttributeModifier{
	{
		Type:      "minecraft:attack_damage",
		Amount:    1000,
		ID:        "minecraft:2121f7b4-5985-43a0-aa3a-57717d7b15c4",
		Operation: "add_multiplied_total",
		Slot:      "any",
	},
	{
		Type:      "minecraft:attack_speed",
		Amount:    100,
		ID:        "minecraft:1df199b2-3849-4112-b9f4-7f16d98d9d38",
		Operation: "add_value",
		Slot:      "any",
	},
}

var poSword = items.NewStackWithComponents(items.IronSword, 1, &items.Components{
	AttributeModifiers: poSwordAttribs,
	Unbreakable:        true,
	TooltipDisplay: &items.TooltipDisplay{
		HideTooltip:      false,
		HiddenComponents: []int32{4, 16},
	},
	CustomName: &items.ItemNameComponent{
		Text: "po",
	},
})

var GoMclibPlayerName = ns.String("GoMclib")
var GoMclibPlayerUUID, _ = ns.UUIDFromString("f8ccd41b-3ab8-32d1-a575-afb9913101d6")

func validatePackets(t *testing.T, packets packetsToBytes) {
	for packet, capture := range packets {
		validatePacket(t, packet, capture)
	}
}

// validatePacket validates that a packet decodes and encodes correctly via struct-level comparison.
// we compare decoded structs (not encoded bytes) because encoding can be non-deterministic:
// Go maps have random iteration order (e.g. heightmaps), and NBT field order depends on
// how the data was originally created. reflect.DeepEqual handles maps by key-value equality,
// making struct comparison robust against all ordering variations.
// both sides are normalized through encode→decode to eliminate nil vs empty slice differences.
func validatePacket(t *testing.T, packet jp.Packet, capture []byte) {
	// decode captured bytes (validates decoder handles real traffic)
	decoded := newPacketLike(packet)
	if err := decoded.Read(ns.NewReader(capture)); err != nil {
		t.Fatalf("failed to decode captured %T: %v", packet, err)
	}

	// normalize both through encode→decode to handle nil vs empty slice differences
	expected := encodeDecodePacket(t, packet)
	actual := encodeDecodePacket(t, decoded)

	// struct-level comparison handles non-deterministic encoding (map ordering, NBT field order)
	if !assert.Equal(t, expected, actual) {
		t.Fatalf("decoded packet `%T` does not match expected", packet)
	}

	// verify round-trip is idempotent: encode→decode→encode→decode produces the same struct
	reRoundTripped := encodeDecodePacket(t, expected)
	if !assert.Equal(t, expected, reRoundTripped) {
		t.Fatalf("round-trip not idempotent for `%T`", packet)
	}
}

// encodeDecodePacket normalizes a packet through an encode→decode cycle.
func encodeDecodePacket(t *testing.T, packet jp.Packet) jp.Packet {
	t.Helper()
	wire, err := jp.ToWire(packet)
	if err != nil {
		t.Fatalf("failed to encode %T: %v", packet, err)
	}
	result := newPacketLike(packet)
	if err := result.Read(ns.NewReader(wire.Data)); err != nil {
		t.Fatalf("failed to decode encoded %T: %v", packet, err)
	}
	return result
}

func newPacketLike(packet jp.Packet) jp.Packet {
	return reflect.New(reflect.TypeOf(packet).Elem()).Interface().(jp.Packet)
}

func hexToBytesMust(hexData string) []byte {
	bytes, err := hex.DecodeString(hexData)
	if err != nil {
		panic(err)
	}
	return bytes
}

func TestPackets(t *testing.T) {
	validatePackets(t, capturedPackets)
}
