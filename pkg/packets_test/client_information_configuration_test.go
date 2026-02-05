package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	capturedPackets[&packets.C2SClientInformationConfiguration{}] = hexToBytesMust("000000000000000000")
}
