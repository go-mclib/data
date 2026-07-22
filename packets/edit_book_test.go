package packets_test

import (
	"strings"
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func TestEditBookRejectsOversizedPayloads(t *testing.T) {
	tests := []struct {
		name  string
		write func(*ns.PacketBuffer)
	}{
		{
			name: "page count",
			write: func(writer *ns.PacketBuffer) {
				writer.WriteVarInt(0)
				writer.WriteVarInt(101)
			},
		},
		{
			name: "page length",
			write: func(writer *ns.PacketBuffer) {
				writer.WriteVarInt(0)
				writer.WriteVarInt(1)
				writer.WriteString(ns.String(strings.Repeat("a", 1025)))
			},
		},
		{
			name: "title length",
			write: func(writer *ns.PacketBuffer) {
				writer.WriteVarInt(0)
				writer.WriteVarInt(0)
				writer.WriteBool(true)
				writer.WriteString(ns.String(strings.Repeat("a", 33)))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := ns.NewWriter()
			test.write(writer)
			var packet packets.C2SEditBook
			if err := packet.Read(ns.NewReader(writer.Bytes())); err == nil {
				t.Fatal("oversized edit-book payload was accepted")
			}
		})
	}
}
