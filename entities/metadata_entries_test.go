package entities_test

import (
	"reflect"
	"testing"

	"github.com/go-mclib/data/entities"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestMetadataEntriesRoundTrip(t *testing.T) {
	name := ns.NewTextComponent("Bob")
	written := entities.Metadata{
		entities.ByteEntry(0, 0x20),
		entities.IntEntry(1, 300),
		entities.VarIntEntry(2, entities.SerializerPOSE, 5),
		entities.StringEntry(3, "text"),
		entities.ComponentEntry(4, name),
		entities.OptionalComponentEntry(5, &name),
		entities.OptionalComponentEntry(6, nil),
		entities.BoolEntry(7, true),
		entities.FloatEntry(8, 1.5),
		entities.RotationsEntry(9, 1, 2, 3),
		entities.SlotEntry(10, ns.NewSlot(1, 2)),
	}
	buf := ns.NewWriter()
	if err := entities.WriteMetadata(buf, written); err != nil {
		t.Fatal(err)
	}
	read, err := entities.ReadMetadata(ns.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(read, written) {
		t.Fatalf("read %+v\nwant %+v", read, written)
	}
}
