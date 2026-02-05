package packets_test

import (
	"encoding/json"

	"github.com/go-mclib/data/pkg/data/misc"
	"github.com/go-mclib/data/pkg/packets"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	statusWant := misc.ServerStatusResponse{
		Description: "go-mclib/data test server",
		Players: misc.ServerStatusPlayers{
			Max:    5,
			Online: 0,
		},
		Version: misc.ServerStatusVersion{
			Name:     "1.21.11",
			Protocol: 774,
		},
	}

	statusWantBytes, err := json.Marshal(statusWant)
	if err != nil {
		panic(err)
	}
	capturedPackets[&packets.S2CStatusResponse{
		JsonResponse: ns.String(statusWantBytes),
	}] = hexToBytesMust("767b226465736372697074696f6e223a22676f2d6d636c69622f64617461207465737420736572766572222c22706c6179657273223a7b226d6178223a352c226f6e6c696e65223a307d2c2276657273696f6e223a7b226e616d65223a22312e32312e3131222c2270726f746f636f6c223a3737347d7d")
}
