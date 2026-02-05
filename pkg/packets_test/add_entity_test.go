package packets_test

import (
	"github.com/go-mclib/data/pkg/data/entities"
	"github.com/go-mclib/data/pkg/packets"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	poSwordItem()
	immortalSnailEntity()
}

func poSwordItem() {
	poItemUUID, err := ns.UUIDFromString("884f5b21-b5ee-41da-b734-409451e3bc78")
	if err != nil {
		panic(err)
	}
	capturedPackets[&packets.S2CAddEntity{
		EntityId:   2,
		EntityUuid: poItemUUID,
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
	}] = hexToBytesMust("02884f5b21b5ee41dab734409451e3bc78470000000000000000c04d570a3d8000000000000000000000c1f880cee41900c30000")
}

func immortalSnailEntity() {

}
