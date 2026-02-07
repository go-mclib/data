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
	}] = capturedBytes["c2s_intention_status"]
}
