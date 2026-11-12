package entities

import (
	"testing"

	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestResolvableProfileMetadataKeepsLaterEntries(t *testing.T) {
	w := ns.NewWriter()
	w.WriteByte(17)
	w.WriteVarInt(41)  // resolvable_profile
	w.WriteBool(false) // partial profile
	w.WriteBool(true)
	w.WriteString("Alex")
	w.WriteBool(false) // no UUID
	w.WriteVarInt(0)   // no properties
	w.WriteBool(false) // body
	w.WriteBool(true)  // cape
	w.WriteIdentifier("minecraft:cape")
	w.WriteBool(false) // elytra
	w.WriteBool(true)  // model
	w.WriteBool(false) // wide
	w.WriteByte(18)
	w.WriteVarInt(0) // byte serializer
	w.WriteByte(7)
	w.WriteByte(0xFF)

	entries, err := ReadMetadata(ns.NewReader(w.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 2 || entries[1].Index != 18 || len(entries[1].Data) != 1 || entries[1].Data[0] != 7 {
		t.Fatalf("entries = %+v", entries)
	}
}
