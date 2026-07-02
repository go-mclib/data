package pgen

import (
	"bytes"
	"testing"

	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// rw verifies a generated packet's Read/Write round-trip byte-for-byte.
func rw(t *testing.T, write func(*ns.PacketBuffer) error, read func(*ns.PacketBuffer) error) {
	t.Helper()
	w := ns.NewWriter()
	if err := write(w); err != nil {
		t.Fatalf("write: %v", err)
	}
	wire := w.Bytes()
	if err := read(ns.NewReader(wire)); err != nil {
		t.Fatalf("read: %v", err)
	}
	w2 := ns.NewWriter()
	if err := write(w2); err != nil {
		t.Fatalf("re-write: %v", err)
	}
	if !bytes.Equal(wire, w2.Bytes()) {
		t.Fatalf("round-trip mismatch:\n in: %x\nout: %x", wire, w2.Bytes())
	}
}

func TestAddEntityRoundTrip(t *testing.T) {
	in := &S2CAddEntity{Id: 1, Type: 5, X: 1.5, Y: 64, Z: -3.5, Movement: ns.LpVec3{X: 0.1, Y: 0.2, Z: 0.3}, XRot: 10, YRot: -20, Data: 7}
	out := &S2CAddEntity{}
	rw(t, in.Write, out.Read)
	if out.Id != in.Id || out.X != in.X || out.Data != in.Data {
		t.Fatalf("decoded fields differ: %+v vs %+v", out, in)
	}
}

func TestSetPassengersRoundTrip(t *testing.T) {
	in := &S2CSetPassengers{Vehicle: 42, Passengers: ns.PrefixedArray[ns.VarInt]{1, 2, 3}}
	out := &S2CSetPassengers{}
	rw(t, in.Write, out.Read)
	if len(out.Passengers) != 3 || out.Passengers[2] != 3 {
		t.Fatalf("list decoded wrong: %+v", out.Passengers)
	}
}

// exercises the slot alias (buf.ReadSlot(items.Decoder())) with an empty slot
func TestContainerSetSlotRoundTrip(t *testing.T) {
	in := &S2CContainerSetSlot{ContainerId: 1, StateId: 2, Slot: 9, ItemStack: ns.Slot{Count: 0}}
	out := &S2CContainerSetSlot{}
	rw(t, in.Write, out.Read)
	if out.Slot != 9 || out.ItemStack.Count != 0 {
		t.Fatalf("decoded wrong: %+v", out)
	}
}

// exercises the inline-nbt path (streaming reader + EncodeNetwork)
func TestBlockEntityDataRoundTrip(t *testing.T) {
	in := &S2CBlockEntityData{Pos: ns.Position{X: 1, Y: 2, Z: 3}, Type: 4, Tag: nbt.Compound{"id": nbt.String("minecraft:chest")}}
	out := &S2CBlockEntityData{}
	rw(t, in.Write, out.Read)
	if c, ok := out.Tag.(nbt.Compound); !ok || c["id"] != nbt.String("minecraft:chest") {
		t.Fatalf("nbt decoded wrong: %+v", out.Tag)
	}
}
