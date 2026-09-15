package packets

import (
	"bytes"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
	"testing"
)

func TestCommandsPayloadHasNoByteArrayPrefix(t *testing.T) {
	// One executable root node, zero children, root index zero.
	raw := []byte{1, 4, 0, 0}
	var p S2CCommands
	if err := p.Read(ns.NewReader(raw)); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(p.Data, raw) {
		t.Fatalf("command graph prefix was consumed: %x", p.Data)
	}
	b := ns.NewWriter()
	if err := p.Write(b); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b.Bytes(), raw) {
		t.Fatalf("extra prefix on command graph: %x", b.Bytes())
	}
}
func TestSignedCommandRejectsOversizedSignatureArray(t *testing.T) {
	for _, count := range []ns.VarInt{-1, 9} {
		b := ns.NewWriter()
		b.WriteString("msg alice hi")
		b.WriteInt64(0)
		b.WriteInt64(0)
		b.WriteVarInt(count)
		var p C2SChatCommandSigned
		if err := p.Read(ns.NewReader(b.Bytes())); err == nil {
			t.Fatalf("accepted signature count %d", count)
		}
	}
	p := C2SChatCommandSigned{Command: "msg alice hi", ArgumentSignatures: make([]CommandArgumentSignature, 9), Acknowledged: ns.NewFixedBitSet(20)}
	if err := p.Write(ns.NewWriter()); err == nil {
		t.Fatal("serialized oversized signature array")
	}
}
