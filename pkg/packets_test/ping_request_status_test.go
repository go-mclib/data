package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	capturedPackets[&packets.C2SPingRequestStatus{
		Timestamp: ns.Int64(518236),
	}] = hexToBytesMust("000000000007e85c")
}
