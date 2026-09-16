package entities

import (
	"bytes"
	"testing"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func TestParticleMetadataPreservesFollowingUseFlags(t *testing.T) {
	particles := ns.NewWriter()
	particles.WriteVarInt(12)
	particle := func(name string) { particles.WriteVarInt(ns.VarInt(registries.ParticleType.Get("minecraft:" + name))) }
	particle("cloud")
	particle("entity_effect")
	particles.WriteInt32(0x00123456)
	particle("effect")
	particles.WriteInt32(0x00010203)
	particles.WriteFloat32(1.5)
	particle("dust_color_transition")
	particles.WriteInt32(0xabcdef)
	particles.WriteInt32(0x123456)
	particles.WriteFloat32(.75)
	particle("block")
	particles.WriteVarInt(300)
	particle("dragon_breath")
	particles.WriteFloat32(.5)
	particle("shriek")
	particles.WriteVarInt(200)
	particle("item")
	if err := items.WriteItemStackTemplate(particles, ns.NewSlot(ns.VarInt(items.ItemID("minecraft:diamond")), 2)); err != nil {
		t.Fatal(err)
	}
	particle("trail")
	particles.WriteFloat64(1)
	particles.WriteFloat64(-2)
	particles.WriteFloat64(3)
	particles.WriteInt32(0x102030)
	particles.WriteVarInt(25)
	particle("vibration")
	particles.WriteVarInt(ns.VarInt(registries.PositionSourceType.Get("minecraft:block")))
	particles.WritePosition(ns.Position{X: -17, Y: 65, Z: 32})
	particles.WriteVarInt(10)
	particle("vibration")
	particles.WriteVarInt(ns.VarInt(registries.PositionSourceType.Get("minecraft:entity")))
	particles.WriteVarInt(42)
	particles.WriteFloat32(1.62)
	particles.WriteVarInt(15)
	particle("geyser_base")
	particles.WriteInt32(4)
	particles.WriteFloat32(1.25)
	expected := Metadata{{Index: 10, Serializer: 17, Data: particles.Bytes()}, {Index: 8, Serializer: 0, Data: []byte{0}}, {Index: 11, Serializer: 8, Data: []byte{1}}}
	encoded := ns.NewWriter()
	if err := WriteMetadata(encoded, expected); err != nil {
		t.Fatal(err)
	}
	frame := append(append([]byte(nil), encoded.Bytes()...), 0x6a)
	reader := ns.NewReader(frame)
	actual, err := ReadMetadata(reader)
	if err != nil {
		t.Fatal(err)
	}
	if use := actual.Get(8); len(use) != 1 || use[0] != 0 {
		t.Fatalf("particle data hid the authoritative use-ended flag: %+v", actual)
	}
	sentinel, err := reader.ReadByte()
	if err != nil || sentinel != 0x6a {
		t.Fatalf("particle payload crossed metadata boundary: byte=%x err=%v", sentinel, err)
	}
	reencoded := ns.NewWriter()
	if err := WriteMetadata(reencoded, actual); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(reencoded.Bytes(), encoded.Bytes()) {
		t.Fatal("particle payload or following metadata changed on read")
	}
}

func TestParticleMetadataRejectsMalformedPayload(t *testing.T) {
	for _, data := range [][]byte{{0xff, 0xff, 0xff, 0xff, 0x0f}, {1, 0xff, 0xff, 0xff, 0xff, 0x07}, {1, byte(registries.ParticleType.Get("minecraft:entity_effect")), 0}} {
		wire := ns.NewWriter()
		if err := WriteMetadata(wire, Metadata{{Index: 10, Serializer: 17, Data: data}}); err != nil {
			t.Fatal(err)
		}
		if _, err := ReadMetadata(ns.NewReader(wire.Bytes())); err == nil {
			t.Fatalf("malformed particle list accepted: %x", data)
		}
	}
}
