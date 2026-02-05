package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	capturedPackets[&packets.S2CLoginFinished{
		Profile: packets.GameProfile{
			UUID:       GoMclibPlayerUUID,
			Name:       GoMclibPlayerName,
			Properties: []packets.GameProfileProperty{}, // offline-mode, so nothing
		},
	}] = hexToBytesMust("f8ccd41b3ab832d1a575afb9913101d607476f4d636c696200")
}
