package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	poSwordSlot, err := poSwordC2S.ToSlot()
	if err != nil {
		panic(err)
	}

	capturedPackets[&packets.C2SSetCreativeModeSlot{
		Slot:        36, // first slot in hotbar from left
		ClickedItem: poSwordSlot,
	}] = hexToBytesMust("002401a307040004000605080002706f107702022e6d696e6563726166743a32313231663762342d353938352d343361302d616133612d353737313764376231356334408f400000000000020000042e6d696e6563726166743a31646631393962322d333834392d343131322d623966342d3766313664393864396433384059000000000000000000120400020410")
}
