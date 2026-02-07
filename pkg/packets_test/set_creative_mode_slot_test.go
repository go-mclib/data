package packets_test

import (
	"github.com/go-mclib/data/pkg/packets"
)

func init() {
	poSwordSlot, err := poSword.ToSlot()
	if err != nil {
		panic(err)
	}

	capturedPackets[&packets.C2SSetCreativeModeSlot{
		Slot:        36, // first slot in hotbar from left
		ClickedItem: poSwordSlot,
	}] = capturedBytes["c2s_set_creative_mode_slot"]
}
