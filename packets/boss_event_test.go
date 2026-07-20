package packets_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/go-mclib/data/packets"
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func TestBossEventActionsRoundTrip(t *testing.T) {
	uuid := ns.UUID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	tests := []struct {
		name   string
		action packets.BossEventActionEnum
		data   interface{ Write(*ns.PacketBuffer) error }
		check  func(*testing.T, *packets.S2CBossEvent)
	}{
		{
			name: "add", action: packets.BossEventActionAdd,
			data: &packets.BossEventActionAddData{
				Title: ns.TextComponent{Text: "Boss", Color: "red"}, Health: 0.75,
				Color: 2, Division: 3, Flags: 7,
			},
			check: func(t *testing.T, packet *packets.S2CBossEvent) {
				got, err := packet.DataActionAdd()
				if err != nil {
					t.Fatal(err)
				}
				if got.Title.PlainText() != "Boss" || got.Title.Color != "red" || got.Health != 0.75 || got.Color != 2 || got.Division != 3 || got.Flags != 7 {
					t.Fatalf("add data = %+v", got)
				}
			},
		},
		{
			name: "progress", action: packets.BossEventActionUpdateHealth,
			data: &packets.BossEventActionUpdateHealthData{Health: 0.25},
			check: func(t *testing.T, packet *packets.S2CBossEvent) {
				got, err := packet.DataActionUpdateHealth()
				if err != nil || got.Health != 0.25 {
					t.Fatalf("progress = %+v, %v", got, err)
				}
			},
		},
		{
			name: "title", action: packets.BossEventActionUpdateTitle,
			data: &packets.BossEventActionUpdateTitleData{Title: ns.TextComponent{Text: "New"}},
			check: func(t *testing.T, packet *packets.S2CBossEvent) {
				got, err := packet.DataActionUpdateTitle()
				if err != nil || got.Title.PlainText() != "New" {
					t.Fatalf("title = %+v, %v", got, err)
				}
			},
		},
		{
			name: "style", action: packets.BossEventActionUpdateStyle,
			data: &packets.BossEventActionUpdateStyleData{Color: 5, Division: 4},
			check: func(t *testing.T, packet *packets.S2CBossEvent) {
				got, err := packet.DataActionUpdateStyle()
				if err != nil || got.Color != 5 || got.Division != 4 {
					t.Fatalf("style = %+v, %v", got, err)
				}
			},
		},
		{
			name: "flags", action: packets.BossEventActionUpdateFlags,
			data: &packets.BossEventActionUpdateFlagsData{Flags: 5},
			check: func(t *testing.T, packet *packets.S2CBossEvent) {
				got, err := packet.DataActionUpdateFlags()
				if err != nil || got.Flags != 5 {
					t.Fatalf("flags = %+v, %v", got, err)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := ns.NewWriter()
			if err := test.data.Write(data); err != nil {
				t.Fatal(err)
			}
			original := &packets.S2CBossEvent{Uuid: uuid, Action: test.action, Data: data.Bytes()}
			wire, err := jp.ToWire(original)
			if err != nil {
				t.Fatal(err)
			}
			var decoded packets.S2CBossEvent
			if err := decoded.Read(ns.NewReaderFrom(bytes.NewReader(wire.Data))); err != nil {
				t.Fatal(err)
			}
			if decoded.Uuid != uuid || decoded.Action != test.action || !reflect.DeepEqual(decoded.Data, original.Data) {
				t.Fatalf("decoded packet = %+v", decoded)
			}
			test.check(t, &decoded)
		})
	}
}

func TestBossEventAddVanillaWireLayout(t *testing.T) {
	uuid := ns.UUID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	data := ns.NewWriter()
	if err := (&packets.BossEventActionAddData{
		Title: ns.TextComponent{Text: "Boss"}, Health: 0.75, Color: 2, Division: 3, Flags: 7,
	}).Write(data); err != nil {
		t.Fatal(err)
	}
	wire, err := jp.ToWire(&packets.S2CBossEvent{Uuid: uuid, Action: packets.BossEventActionAdd, Data: data.Bytes()})
	if err != nil {
		t.Fatal(err)
	}
	// ClientboundBossEventPacket.AddOperation: UUID, operation VarInt,
	// trusted component NBT, progress float, color/overlay VarInts, flags byte.
	want := []byte{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		0,
		8, 0, 4, 'B', 'o', 's', 's',
		0x3f, 0x40, 0x00, 0x00,
		2, 3, 7,
	}
	if !bytes.Equal(wire.Data, want) {
		t.Fatalf("wire data = % x, want % x", wire.Data, want)
	}
}

func TestBossEventRemoveHasNoPayload(t *testing.T) {
	original := &packets.S2CBossEvent{Uuid: ns.UUID{1}, Action: packets.BossEventActionRemove}
	wire, err := jp.ToWire(original)
	if err != nil {
		t.Fatal(err)
	}
	var decoded packets.S2CBossEvent
	if err := decoded.Read(ns.NewReaderFrom(bytes.NewReader(wire.Data))); err != nil {
		t.Fatal(err)
	}
	if decoded.Action != packets.BossEventActionRemove || len(decoded.Data) != 0 {
		t.Fatalf("remove packet = %+v", decoded)
	}
}
