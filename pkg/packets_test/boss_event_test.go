package packets_test

import (
	"fmt"

	"github.com/go-mclib/data/pkg/packets"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	bossUuid, err := ns.UUIDFromString("8135db86-12a7-4d3a-8bc4-7e1456c21390")
	if err != nil {
		panic(err)
	}
	bossEventPacket := &packets.S2CBossEvent{
		Uuid:   bossUuid,
		Action: packets.BossEventActionUpdateTitle,
		Data:   hexToBytesMust("0a080009696e73657274696f6e002435646664313336632d666638632d346437392d613239312d6662623164626263656161360800047465787400104861707079205769746865726c696e670a000b686f7665725f6576656e740800046e616d6500104861707079205769746865726c696e67080006616374696f6e000b73686f775f656e74697479080002696400106d696e6563726166743a7769746865720b000475756964000000045dfd136cff8c4d79a291fbb1dbbceaa60000"),
	}
	capturedPackets[bossEventPacket] = hexToBytesMust("8135db8612a74d3a8bc47e1456c21390030a080009696e73657274696f6e002435646664313336632d666638632d346437392d613239312d6662623164626263656161360800047465787400104861707079205769746865726c696e670a000b686f7665725f6576656e740800046e616d6500104861707079205769746865726c696e67080006616374696f6e000b73686f775f656e74697479080002696400106d696e6563726166743a7769746865720b000475756964000000045dfd136cff8c4d79a291fbb1dbbceaa60000")

	// test data decode
	updateTitleData, err := bossEventPacket.DataActionUpdateTitle()
	if err != nil {
		panic(err)
	}
	if updateTitleData.Title.Text != "Happy Witherling" {
		panic(fmt.Errorf("updateTitleData.Title is not 'Happy Witherling', but '%s'", updateTitleData.Title.Text))
	}

	// test data roundtrip: decode then re-encode should produce valid data
	roundtripBuf := ns.NewWriter()
	if err := updateTitleData.Write(roundtripBuf); err != nil {
		panic(fmt.Errorf("roundtrip write failed: %w", err))
	}
	var roundtripped packets.BossEventActionUpdateTitleData
	if err := roundtripped.Read(ns.NewReader(roundtripBuf.Bytes())); err != nil {
		panic(fmt.Errorf("roundtrip read failed: %w", err))
	}
	if roundtripped.Title.Text != "Happy Witherling" {
		panic(fmt.Errorf("roundtripped title is %q, expected 'Happy Witherling'", roundtripped.Title.Text))
	}
}
