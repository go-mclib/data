package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	capturedPackets[&packets.C2SStatusRequest{}] = capturedBytes["c2s_status_request"]
}
