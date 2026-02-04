package packets_test

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/go-mclib/data/pkg/data/items"
	jp "github.com/go-mclib/protocol/java_protocol"
	"github.com/stretchr/testify/assert"
)

type packetsToBytes map[jp.Packet][]byte

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

// NOTE: setter order determines wire encoding order, must match captured data
// S2C order: attribute_modifiers -> unbreakable -> tooltip_display -> custom_name
var poSwordS2C = items.NewStack(items.IronSword, 1).
	SetAttributeModifiers(poSwordAttribs).
	SetUnbreakable(true).
	SetTooltipDisplay(&items.TooltipDisplay{
		HideTooltip:      false,
		HiddenComponents: []int32{4, 16},
	}).
	SetCustomName(&items.ItemNameComponent{
		Text: "po",
	})

// C2S order: tooltip_display -> custom_name -> attribute_modifiers -> unbreakable
var poSwordC2S = items.NewStack(items.IronSword, 1).
	SetTooltipDisplay(&items.TooltipDisplay{
		HideTooltip:      false,
		HiddenComponents: []int32{4, 16},
	}).
	SetCustomName(&items.ItemNameComponent{
		Text: "po",
	}).
	SetAttributeModifiers(poSwordAttribs).
	SetUnbreakable(true)

func validatePackets(t *testing.T, packets packetsToBytes) {
	for packet, capture := range packets {
		validatePacket(t, packet, capture)
	}
}

func validatePacket(t *testing.T, packet jp.Packet, capture []byte) {
	const compressionThreshold = 256 // default vanilla threshold

	// encode
	wirePacket, err := jp.ToWire(packet)
	if err != nil {
		t.Fatalf("failed to convert packet %T to wire: %v", packet, err)
	}
	buf := bytes.NewBuffer(nil)
	wirePacket.WriteTo(buf, compressionThreshold)
	if !assert.Equal(t, capture, buf.Bytes()) {
		t.Fatalf("packet `%T` does not match captured bytes", packet)
	}

	// decode: read captured bytes back into a new packet instance
	wire, err := jp.ReadWirePacketFrom(bytes.NewReader(capture), compressionThreshold)
	if err != nil {
		t.Fatalf("failed to read wire packet from captured bytes: %v", err)
	}

	// create new instance of same type and decode into it
	decoded := reflect.New(reflect.TypeOf(packet).Elem()).Interface().(jp.Packet)
	if err := wire.ReadInto(decoded); err != nil {
		t.Fatalf("failed to decode packet %T: %v", packet, err)
	}

	// compare by re-encoding decoded packet (avoids nil vs empty slice differences)
	reEncoded, err := jp.ToWire(decoded)
	if err != nil {
		t.Fatalf("failed to re-encode decoded packet %T: %v", decoded, err)
	}
	reBuf := bytes.NewBuffer(nil)
	reEncoded.WriteTo(reBuf, compressionThreshold)
	if !assert.Equal(t, capture, reBuf.Bytes()) {
		t.Fatalf("re-encoded packet `%T` does not match original bytes", decoded)
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
