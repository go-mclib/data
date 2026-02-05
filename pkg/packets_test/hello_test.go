package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	capturedPackets[&packets.C2SHello{
		Name:       GoMclibPlayerName,
		PlayerUuid: GoMclibPlayerUUID,
	}] = hexToBytesMust("07476f4d636c6962f8ccd41b3ab832d1a575afb9913101d6")
}
