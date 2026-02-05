package packets_test

import (
	"github.com/go-mclib/data/pkg/data/entities"
	"github.com/go-mclib/data/pkg/packets"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	poSwordSlot, err := poSwordS2C.ToSlot()
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
	}] = hexToBytesMust("02080701a30704001002022e6d696e6563726166743a32313231663762342d353938352d343361302d616133612d353737313764376231356334408f400000000000020000042e6d696e6563726166743a31646631393962322d333834392d343131322d623966342d376631366439386439643338405900000000000000000004120002041006080002706fff")
}
