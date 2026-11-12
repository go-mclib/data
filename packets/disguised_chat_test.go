package packets_test

import (
	"bytes"
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestDisguisedChatUsesBoundChatTypeWireLayout(t *testing.T) {
	// Vanilla sends message + ChatType.Bound, not a length-prefixed opaque blob.
	wire := ns.NewWriter()
	if err := wire.WriteTextComponent(ns.TextComponent{Text: "hello"}); err != nil {
		t.Fatal(err)
	}
	if err := wire.WriteVarInt(4); err != nil {
		t.Fatal(err)
	} // logical registry ID 3 + 1
	if err := wire.WriteTextComponent(ns.TextComponent{Text: "Alice"}); err != nil {
		t.Fatal(err)
	}
	if err := wire.WriteBool(true); err != nil {
		t.Fatal(err)
	}
	if err := wire.WriteTextComponent(ns.TextComponent{Text: "SteedsGate"}); err != nil {
		t.Fatal(err)
	}
	fixture := bytes.Clone(wire.Bytes())
	reader := ns.NewReader(append(bytes.Clone(fixture), 0x55))
	var decoded packets.S2CDisguisedChat
	if err := decoded.Read(reader); err != nil {
		t.Fatal(err)
	}
	if decoded.Message.String() != "hello" {
		t.Fatalf("wrong message: %s", decoded.Message.String())
	}
	if sentinel, err := reader.ReadByte(); err != nil || sentinel != 0x55 {
		t.Fatalf("chat codec consumed wrong number of bytes: %x %v", sentinel, err)
	}
	encoded := ns.NewWriter()
	if err := decoded.Write(encoded); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(encoded.Bytes(), fixture) {
		t.Fatalf("bound chat type lost on replay: got %x want %x", encoded.Bytes(), fixture)
	}
}
