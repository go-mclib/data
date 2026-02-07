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

var poSword = items.NewStack(items.IronSword, 1).
	SetAttributeModifiers(poSwordAttribs).
	SetUnbreakable(true).
	SetTooltipDisplay(&items.TooltipDisplay{
		HideTooltip:      false,
		HiddenComponents: []int32{4, 16},
	}).
	SetCustomName(&items.ItemNameComponent{
		Text: "po",
	})

var GoMclibPlayerName = ns.String("GoMclib")
var GoMclibPlayerUUID, _ = ns.UUIDFromString("f8ccd41b-3ab8-32d1-a575-afb9913101d6")

func validatePackets(t *testing.T, packets packetsToBytes) {
	for packet, capture := range packets {
		validatePacket(t, packet, capture)
	}
}

// validatePacket uses round-trip validation to ensure the packet decoder and encoder are working correctly
// why not compare against the captured bytes directly? because the field order e.g. in NBT tags, text components etc.
// is not consistent (its order is determined e.g. by the order the components appear in /give command), and
// our library cannot predict that. we could add ordered encoding via builder pattern for these fields, but that adds
// complexity, so instead we just simplify the tests to validate that the packet encodes into something meaningful.
// NB the protocol will still understand the packet regardless of the order *in the packet fields* (NBT, etc.) -
// the order of the packet fields themselves (in packet "root") is important!
func validatePacket(t *testing.T, packet jp.Packet, capture []byte) {
	// decode captured bytes (validates decoder handles real traffic)
	decoded := reflect.New(reflect.TypeOf(packet).Elem()).Interface().(jp.Packet)
	if err := decoded.Read(ns.NewReader(capture)); err != nil {
		t.Fatalf("failed to decode captured %T: %v", packet, err)
	}

	// encode both structs through our encoder and compare bytes;
	// this validates decoding correctness while normalizing
	// nil vs empty slice differences
	expectedBytes, err := jp.ToWire(packet)
	if err != nil {
		t.Fatalf("failed to encode expected %T: %v", packet, err)
	}
	decodedBytes, err := jp.ToWire(decoded)
	if err != nil {
		t.Fatalf("failed to encode decoded %T: %v", decoded, err)
	}
	if !assert.Equal(t, expectedBytes.Data, decodedBytes.Data) {
		t.Fatalf("decoded packet `%T` does not match expected", packet)
	}

	// round-trip: encode -> decode -> re-encode (validates encoder determinism)
	roundTripped := reflect.New(reflect.TypeOf(packet).Elem()).Interface().(jp.Packet)
	if err := roundTripped.Read(ns.NewReader(expectedBytes.Data)); err != nil {
		t.Fatalf("failed to decode round-trip %T: %v", packet, err)
	}
	reEncoded, err := jp.ToWire(roundTripped)
	if err != nil {
		t.Fatalf("failed to re-encode round-trip %T: %v", roundTripped, err)
	}
	if !assert.Equal(t, expectedBytes.Data, reEncoded.Data) {
		t.Fatalf("round-trip failed for packet `%T`", packet)
	}
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
