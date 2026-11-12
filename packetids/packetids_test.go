package packetids_test

import (
	"testing"

	"github.com/go-mclib/data/packetids"
)

func TestPlayPacketIDs(t *testing.T) {
	// verify some well-known packet IDs are non-negative
	// exact values may change between versions, but they should exist
	tests := []struct {
		name string
		id   int32
	}{
		{"C2SClientInformationPlayID", packetids.C2SClientInformationPlayID},
		{"C2SMovePlayerPosID", packetids.C2SMovePlayerPosID},
		{"C2SChatID", packetids.C2SChatID},
		{"S2CLoginID", packetids.S2CLoginID},
		{"S2CPlayerPositionID", packetids.S2CPlayerPositionID},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.id < 0 {
				t.Errorf("%s = %d, want non-negative", tt.name, tt.id)
			}
		})
	}
}

func TestLoginPacketIDs(t *testing.T) {
	// login phase packets
	if packetids.C2SHelloID < 0 {
		t.Error("C2SHelloID should be non-negative")
	}
	if packetids.S2CHelloID < 0 {
		t.Error("S2CHelloID should be non-negative")
	}
	if packetids.S2CLoginFinishedID < 0 {
		t.Error("S2CLoginFinishedID should be non-negative")
	}
}

func TestConfigurationPacketIDs(t *testing.T) {
	// configuration phase packets
	if packetids.C2SClientInformationConfigurationID < 0 {
		t.Error("C2SClientInformationConfigurationID should be non-negative")
	}
	if packetids.S2CFinishConfigurationID < 0 {
		t.Error("S2CFinishConfigurationID should be non-negative")
	}
}

func TestHandshakingPacketIDs(t *testing.T) {
	// handshaking phase has only client intention
	if packetids.C2SIntentionID < 0 {
		t.Error("C2SIntentionID should be non-negative")
	}
}

func TestStatusPacketIDs(t *testing.T) {
	// status phase packets
	if packetids.C2SStatusRequestID < 0 {
		t.Error("C2SStatusRequestID should be non-negative")
	}
	if packetids.S2CStatusResponseID < 0 {
		t.Error("S2CStatusResponseID should be non-negative")
	}
}
