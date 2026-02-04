package packets_test

import (
	"github.com/go-mclib/data/pkg/data/entities"
	"github.com/go-mclib/data/pkg/packets"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	entityUUID, err := ns.UUIDFromString("0001d206-ae48-0592-62b4-4703b0aeeacc")
	if err != nil {
		panic(err)
	}

	capturedPackets[&packets.S2CAddEntity{
		EntityId:   55,
		EntityUuid: entityUUID,
		Type:       entities.Item,
		X:          0,
		Y:          -58.68000000715256,
		Z:          0,
		Velocity: ns.LpVec3{
			X: -0.51416015625,
			Y: 0.02520751953125,
			Z: -0.43597412109375,
		},
		Pitch:   0,
		Yaw:     195,
		HeadYaw: 0,
		Data:    0,
	}] = hexToBytesMust("370001d206ae48059262b44703b0aeeacc24e16a3147c01571e220cc8410c04c12066fe16bfe402d4616bd4cf8e239856b93117d00d60000")
}
