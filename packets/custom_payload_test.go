package packets_test

import (
	"bytes"
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func TestCustomPayloadIsInline(t *testing.T) {
	payload := ns.ByteArray{7, 'v', 'a', 'n', 'i', 'l', 'l', 'a'}
	writers := map[string]interface{ Write(*ns.PacketBuffer) error }{
		"c2s_configuration": &packets.C2SCustomPayloadConfiguration{Channel: "minecraft:brand", Data: payload},
		"s2c_configuration": &packets.S2CCustomPayloadConfiguration{Channel: "minecraft:brand", Data: payload},
		"c2s_play":          &packets.C2SCustomPayloadPlay{Channel: "minecraft:brand", Data: payload},
		"s2c_play":          &packets.S2CCustomPayloadPlay{Channel: "minecraft:brand", Data: payload},
	}
	want := append([]byte{15}, []byte("minecraft:brand")...)
	want = append(want, payload...)
	for name, packet := range writers {
		t.Run(name, func(t *testing.T) {
			buf := ns.NewWriter()
			if err := packet.Write(buf); err != nil {
				t.Fatalf("Write: %v", err)
			}
			if !bytes.Equal(buf.Bytes(), want) {
				t.Fatalf("wire = %x, want %x", buf.Bytes(), want)
			}
		})
	}
}

func TestCustomPayloadReadsRemainder(t *testing.T) {
	wire := append([]byte{15}, []byte("minecraft:brand")...)
	wire = append(wire, 7, 'v', 'a', 'n', 'i', 'l', 'l', 'a')
	var packet packets.C2SCustomPayloadConfiguration
	if err := packet.Read(ns.NewReader(wire)); err != nil {
		t.Fatalf("Read: %v", err)
	}
	if !bytes.Equal(packet.Data, wire[16:]) {
		t.Fatalf("payload = %x, want %x", packet.Data, wire[16:])
	}
}
