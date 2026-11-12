package packets_test

import (
	"bytes"
	"testing"

	"github.com/go-mclib/data/packets"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

func TestAwardStatsPacketRoundTrip(t *testing.T) {
	packet := &packets.S2CAwardStats{Statistics: []packets.Statistic{
		{CategoryID: 2, StatisticID: 967, Value: 3},
		{CategoryID: 8, StatisticID: 47, Value: 2400},
	}}
	writer := ns.NewWriter()
	if err := packet.Write(writer); err != nil {
		t.Fatalf("Write: %v", err)
	}
	want := []byte{2, 2, 0xc7, 7, 3, 8, 47, 0xe0, 0x12}
	if !bytes.Equal(writer.Bytes(), want) {
		t.Fatalf("wire = %x, want %x", writer.Bytes(), want)
	}
	var decoded packets.S2CAwardStats
	if err := decoded.Read(ns.NewReader(writer.Bytes())); err != nil {
		t.Fatalf("Read: %v", err)
	}
	if len(decoded.Statistics) != 2 || decoded.Statistics[1].Value != 2400 {
		t.Fatalf("decoded statistics = %+v", decoded.Statistics)
	}
}
