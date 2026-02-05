package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	capturedPackets[&packets.C2SIntention{
		ProtocolVersion: 774,
		ServerAddress:   "localhost",
		ServerPort:      25565,
		Intent:          1,
	}] = hexToBytesMust("8606096c6f63616c686f737463dd01")
}
