package packets_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestSetPassengersWireLayout(t *testing.T) {
	packet := &packets.S2CSetPassengers{EntityID: 300, Passengers: []ns.VarInt{1, 400}}
	buf := ns.NewWriter()
	if err := packet.Write(buf); err != nil {
		t.Fatalf("Write: %v", err)
	}
	want := []byte{0xac, 0x02, 0x02, 0x01, 0x90, 0x03}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Fatalf("wire = %x, want %x", buf.Bytes(), want)
	}
	var decoded packets.S2CSetPassengers
	reader := ns.NewReader(want)
	if err := decoded.Read(reader); err != nil {
		t.Fatalf("Read: %v", err)
	}
	if decoded.EntityID != 300 || len(decoded.Passengers) != 2 || decoded.Passengers[1] != 400 {
		t.Fatalf("decoded = %+v", decoded)
	}
}

func TestSetEquipmentWireLayout(t *testing.T) {
	packet := &packets.S2CSetEquipment{
		EntityID: 300,
		Slots: []packets.EquipmentEntry{
			{Slot: 2, Item: ns.NewSlot(5, 1)},
			{Slot: 4, Item: ns.EmptySlot()},
		},
	}
	buf := ns.NewWriter()
	if err := packet.Write(buf); err != nil {
		t.Fatalf("Write: %v", err)
	}
	want := []byte{0xac, 0x02, 0x82, 0x01, 0x05, 0x00, 0x00, 0x04, 0x00}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Fatalf("wire = %x, want %x", buf.Bytes(), want)
	}
	var decoded packets.S2CSetEquipment
	reader := ns.NewReader(want)
	if err := decoded.Read(reader); err != nil {
		t.Fatalf("Read: %v", err)
	}
	if decoded.EntityID != 300 || len(decoded.Slots) != 2 || decoded.Slots[0].Slot != 2 || decoded.Slots[0].Item.ItemID != 5 || decoded.Slots[1].Slot != 4 || !decoded.Slots[1].Item.IsEmpty() {
		t.Fatalf("decoded = %+v", decoded)
	}
	if _, err := reader.ReadByte(); err != io.EOF {
		t.Fatalf("trailing data: %v", err)
	}
}
