package packets_test

import (
	"github.com/go-mclib/data/pkg/data/entities"
	"github.com/go-mclib/data/pkg/packets"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	poSwordItem()
}

func poSwordItem() {
	poItemUUID, err := ns.UUIDFromString("884f5b21-b5ee-41da-b734-409451e3bc78")
	if err != nil {
		panic(err)
	}
	capturedPackets[&packets.S2CAddEntity{
		EntityId:   2,
		EntityUuid: poItemUUID,
		Type:       ns.VarInt(entities.EntityTypeID("minecraft:item")),
		X:          0,
		Y:          -58.68000000715256,
		Z:          0,
		Velocity: ns.LpVec3{
			X: -0.014099981688335483,
			Y: -0.10895440395531952,
			Z: 0.006348043703839457,
		},
		Pitch:   0,
		Yaw:     195,
		HeadYaw: 0,
		Data:    0,
	}] = capturedBytes["s2c_add_entity_item"]
}
