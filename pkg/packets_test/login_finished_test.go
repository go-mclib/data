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
	}] = capturedBytes["s2c_login_finished"]
}
