package packets_test

import (
	"bytes"
	"testing"

	"github.com/go-mclib/data/packetids"
	"github.com/go-mclib/data/packets"
	"github.com/go-mclib/nbt"
	jp "github.com/go-mclib/protocol/javaproto"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// buildRegistryWire builds a registry_data payload exactly as the server sends
// it: registry id, entry count, then per-entry [id][hasData bool][inline NBT].
func buildRegistryWire(t *testing.T) []byte {
	t.Helper()
	buf := ns.NewWriter()
	if err := buf.WriteIdentifier("minecraft:worldgen/biome"); err != nil {
		t.Fatal(err)
	}
	if err := buf.WriteVarInt(2); err != nil {
		t.Fatal(err)
	}

	entries := []struct {
		id   ns.Identifier
		data nbt.Compound
	}{
		{"minecraft:plains", nbt.Compound{
			"has_precipitation": nbt.Byte(1),
			"temperature":       nbt.Float(0.8),
			"spawn_light_level": nbt.Compound{"type": nbt.String("minecraft:uniform"), "max_inclusive": nbt.Int(7), "min_inclusive": nbt.Int(0)},
			"infiniburn":        nbt.String("#minecraft:infiniburn_overworld"),
		}},
		{"minecraft:desert", nbt.Compound{
			"has_precipitation": nbt.Byte(0),
			"temperature":       nbt.Float(2.0),
		}},
	}

	for _, e := range entries {
		if err := buf.WriteIdentifier(e.id); err != nil {
			t.Fatal(err)
		}
		if err := buf.WriteBool(true); err != nil {
			t.Fatal(err)
		}
		raw, err := nbt.EncodeNetwork(e.data)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := buf.Write(raw); err != nil {
			t.Fatal(err)
		}
	}
	return buf.Bytes()
}

func TestRegistryDataRoundTrip(t *testing.T) {
	wire := buildRegistryWire(t)

	in := &jp.WirePacket{PacketID: ns.VarInt(packetids.S2CRegistryDataID), Data: wire}
	var d packets.S2CRegistryData
	if err := in.ReadInto(&d); err != nil {
		t.Fatalf("ReadInto failed: %v", err)
	}
	t.Logf("parsed registry=%s entries=%d", d.RegistryID, len(d.Entries))

	out, err := jp.ToWire(&d)
	if err != nil {
		t.Fatalf("ToWire failed: %v", err)
	}

	if !bytes.Equal(wire, out.Data) {
		t.Fatalf("round-trip mismatch:\n in (%d): %x\nout (%d): %x", len(wire), wire, len(out.Data), out.Data)
	}
}

// TestInlineNBTPacketsRoundTrip covers the other packets that carry a trailing
// inline-NBT field, ensuring read consumes exactly the tag (no length prefix).
func TestInlineNBTPacketsRoundTrip(t *testing.T) {
	be := &packets.S2CBlockEntityData{
		Location: ns.Position{X: 10, Y: 64, Z: -20},
		Type:     7,
		NBTData:  nbt.Compound{"Items": nbt.Compound{}, "id": nbt.String("minecraft:chest")},
	}
	roundTrip(t, be, &packets.S2CBlockEntityData{})

	tq := &packets.S2CTagQuery{TransactionID: 42, NBT: nbt.Compound{"foo": nbt.Int(1)}}
	roundTrip(t, tq, &packets.S2CTagQuery{})

	// nullable tag (server sends TAG_End when absent)
	tqEmpty := &packets.S2CTagQuery{TransactionID: 1, NBT: nbt.End{}}
	roundTrip(t, tqEmpty, &packets.S2CTagQuery{})
}

func roundTrip(t *testing.T, in, out jp.Packet) {
	t.Helper()
	wire, err := jp.ToWire(in)
	if err != nil {
		t.Fatalf("ToWire: %v", err)
	}
	if err := wire.ReadInto(out); err != nil {
		t.Fatalf("ReadInto: %v", err)
	}
	again, err := jp.ToWire(out)
	if err != nil {
		t.Fatalf("ToWire(out): %v", err)
	}
	if !bytes.Equal(wire.Data, again.Data) {
		t.Fatalf("%T mismatch:\n in: %x\nout: %x", in, wire.Data, again.Data)
	}
}
