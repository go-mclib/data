package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	capturedPackets[&packets.C2SClientInformationConfiguration{}] = capturedBytes["c2s_client_information_configuration"]
}
