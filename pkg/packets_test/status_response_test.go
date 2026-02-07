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
	}] = capturedBytes["s2c_status_response"]
}
