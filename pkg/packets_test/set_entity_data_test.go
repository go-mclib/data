package packets_test

import (
	"github.com/go-mclib/data/pkg/data/entities"
	"github.com/go-mclib/data/pkg/packets"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	poSwordSlot, err := poSword.ToSlot()
	if err != nil {
		panic(err)
	}

	poSwordSlotBuffer := ns.NewWriter()
	if err = poSwordSlot.Encode(poSwordSlotBuffer); err != nil {
		panic(err)
	}

	capturedPackets[&packets.S2CSetEntityData{
		EntityId: 2,
		Metadata: entities.Metadata{
			entities.MetadataEntry{
				Index:      8,
				Serializer: entities.SerializerITEM_STACK,
				Data:       poSwordSlotBuffer.Bytes(),
			},
		},
	}] = capturedBytes["s2c_set_entity_data_item"]
}
