package packets_test

import (
	"bytes"
	"testing"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func TestPositionedAdvancementsWireAlignment(t *testing.T) {
	// Construct the Java field order independently of the packet encoder. The
	// first holder has no display, but must still consume coordinates. The
	// second has a template icon, requirements and telemetry before coordinates.
	wire := ns.NewWriter()
	wire.WriteBool(true)
	wire.WriteVarInt(2)
	wire.WriteIdentifier("minecraft:hidden")
	wire.WriteBool(false) // parent
	wire.WriteBool(false) // display
	wire.WriteVarInt(1)   // requirement groups
	wire.WriteVarInt(1)
	wire.WriteString("first")
	wire.WriteBool(true) // telemetry
	wire.WriteFloat32(-1.5)
	wire.WriteFloat32(2.25)
	wire.WriteIdentifier("minecraft:shown")
	wire.WriteBool(true)
	wire.WriteIdentifier("minecraft:hidden")
	wire.WriteBool(true) // display
	wire.WriteTextComponent(ns.TextComponent{Text: "Title"})
	wire.WriteTextComponent(ns.TextComponent{Text: "Description"})
	itemID := ns.VarInt(items.ItemID("minecraft:stone"))
	wire.WriteVarInt(itemID) // template holder precedes count (not slot framing)
	wire.WriteVarInt(2)
	wire.WriteVarInt(0) // component additions
	wire.WriteVarInt(0) // component removals
	wire.WriteVarInt(1) // challenge
	wire.WriteInt32(7)  // background, toast, hidden
	wire.WriteIdentifier("minecraft:gui/advancements/backgrounds/stone")
	wire.WriteVarInt(2) // requirements follow the complete display
	wire.WriteVarInt(2)
	wire.WriteString("a")
	wire.WriteString("b")
	wire.WriteVarInt(1)
	wire.WriteString("c")
	wire.WriteBool(false)
	wire.WriteFloat32(7.5)
	wire.WriteFloat32(-8.25)
	wire.WriteVarInt(1)
	wire.WriteIdentifier("minecraft:removed")
	wire.WriteVarInt(1) // progress map
	wire.WriteIdentifier("minecraft:shown")
	wire.WriteVarInt(2)
	wire.WriteString("a")
	wire.WriteBool(true)
	wire.WriteInt64(1700000000123)
	wire.WriteString("b")
	wire.WriteBool(false)
	wire.WriteBool(true) // show advancements

	reader := bytes.NewReader(append(append([]byte(nil), wire.Bytes()...), 0x5a))
	var packet packets.S2CUpdateAdvancements
	if err := packet.Read(ns.NewReaderFrom(reader)); err != nil {
		t.Fatal(err)
	}
	if !packet.Reset || !packet.ShowAdvancements || len(packet.Added) != 2 || len(packet.Removed) != 1 || packet.Removed[0] != "minecraft:removed" || len(packet.Progress) != 1 {
		t.Fatalf("advancement packet fields misaligned: %+v", packet)
	}
	first, second := packet.Added[0], packet.Added[1]
	if first.Display != nil || first.X != -1.5 || first.Y != 2.25 || !first.SendsTelemetryEvent || first.Requirements[0][0] != "first" {
		t.Fatalf("no-display holder = %+v", first)
	}
	if second.Id != "minecraft:shown" || !second.Parent.Present || second.Parent.Value != first.Id || second.X != 7.5 || second.Y != -8.25 || bool(second.SendsTelemetryEvent) || len(second.Requirements) != 2 || second.Requirements[0][1] != "b" || second.Requirements[1][0] != "c" {
		t.Fatalf("display holder = %+v", second)
	}
	display := second.Display
	if display == nil || display.Title.PlainText() != "Title" || display.Description.PlainText() != "Description" || display.Icon.ItemID != itemID || display.Icon.Count != 2 || display.Frame != packets.AdvancementFrameChallenge || !bool(display.ShowToast) || !bool(display.Hidden) || !display.Background.Present || display.Background.Value != "minecraft:gui/advancements/backgrounds/stone" {
		t.Fatalf("display = %+v", display)
	}
	progress := packet.Progress[0]
	if progress.Id != second.Id || len(progress.Criteria) != 2 || progress.Criteria[0].Name != "a" || !progress.Criteria[0].Obtained.Present || progress.Criteria[0].Obtained.Value != 1700000000123 || progress.Criteria[1].Name != "b" || progress.Criteria[1].Obtained.Present {
		t.Fatalf("progress = %+v", progress)
	}
	if sentinel, err := reader.ReadByte(); err != nil || sentinel != 0x5a || reader.Len() != 0 {
		t.Fatalf("advancement payload boundary: sentinel=%x err=%v", sentinel, err)
	}
	encoded := ns.NewWriter()
	if err := packet.Write(encoded); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(encoded.Bytes(), wire.Bytes()) {
		t.Fatalf("encoded=%x; want=%x", encoded.Bytes(), wire.Bytes())
	}
}

func TestAdvancementRejectsInvalidCollectionCounts(t *testing.T) {
	for _, wire := range [][]byte{
		{0, 0xff, 0xff, 0xff, 0xff, 0x0f},                  // negative holder count
		{0, 0xff, 0xff, 0xff, 0xff, 0x07},                  // enormous count without a holder
		{0, 0, 0, 1, 1, 'a', 0xff, 0xff, 0xff, 0xff, 0x0f}, // negative criterion count
	} {
		var packet packets.S2CUpdateAdvancements
		if err := packet.Read(ns.NewReader(wire)); err == nil {
			t.Fatalf("accepted malformed advancement packet %x", wire)
		}
	}
}
