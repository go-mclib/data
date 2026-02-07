package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	capturedPackets[&packets.C2SHello{
		Name:       GoMclibPlayerName,
		PlayerUuid: GoMclibPlayerUUID,
	}] = capturedBytes["c2s_hello"]
}
