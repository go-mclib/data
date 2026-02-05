package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	capturedPackets[&packets.S2CLoginCompression{
		Threshold: ns.VarInt(256),
	}] = hexToBytesMust("8002")
}
