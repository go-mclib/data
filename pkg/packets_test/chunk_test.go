package packets_test

import (
	"fmt"

	"github.com/go-mclib/data/pkg/data/blocks"
	"github.com/go-mclib/data/pkg/data/chunks"
	"github.com/go-mclib/data/pkg/data/registries"
	"github.com/go-mclib/data/pkg/packets"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
	"github.com/go-mclib/protocol/nbt"
)

func init() {
	// S2CSetChunkCacheCenter: set center chunk to (1, 1)
	capturedPackets[&packets.S2CSetChunkCacheCenter{
		ChunkX: 1,
		ChunkZ: 1,
	}] = capturedBytes["s2c_set_chunk_cache_center"]

	// S2CChunkBatchStart: empty packet
	capturedPackets[&packets.S2CChunkBatchStart{}] = capturedBytes["s2c_chunk_batch_start"]

	// S2CChunkBatchFinished: batch size = 9
	capturedPackets[&packets.S2CChunkBatchFinished{
		BatchSize: 9,
	}] = capturedBytes["s2c_chunk_batch_finished"]

	// S2CLevelChunkWithLight: chunk at (1, 1) with 50 hay bales and a sign
	chunkRaw := capturedBytes["s2c_level_chunk_with_light"]
	var chunkPkt packets.S2CLevelChunkWithLight
	if err := chunkPkt.Read(ns.NewReader(chunkRaw)); err != nil {
		panic(fmt.Errorf("failed to decode chunk packet: %w", err))
	}
	capturedPackets[&chunkPkt] = chunkRaw

	validateChunkContents(&chunkPkt)
}

func validateChunkContents(pkt *packets.S2CLevelChunkWithLight) {
	if pkt.ChunkX != 1 || pkt.ChunkZ != 1 {
		panic(fmt.Errorf("chunk position: got (%d, %d), want (1, 1)", pkt.ChunkX, pkt.ChunkZ))
	}

	// parse chunk sections and count blocks
	col, err := chunks.ParseChunkColumn(int32(pkt.ChunkX), int32(pkt.ChunkZ), pkt.ChunkData, nil)
	if err != nil {
		panic(fmt.Errorf("failed to parse chunk column: %w", err))
	}

	hayBlockState := blocks.StateID(blocks.HayBlock, map[string]string{"axis": "y"})
	signBlockState := blocks.StateID(blocks.PaleOakWallSign, map[string]string{"facing": "west", "waterlogged": "false"})

	hayCount := 0
	signCount := 0
	for y := chunks.MinY; y < chunks.MaxY; y++ {
		for x := range 16 {
			for z := range 16 {
				state := col.GetBlockState(x, y, z)
				switch state {
				case hayBlockState:
					hayCount++
				case signBlockState:
					signCount++
				}
			}
		}
	}
	if hayCount != 50 {
		panic(fmt.Errorf("hay block count: got %d, want 50", hayCount))
	}
	if signCount != 1 {
		panic(fmt.Errorf("sign block count: got %d, want 1", signCount))
	}

	// validate sign block entity
	if len(pkt.ChunkData.BlockEntities) != 1 {
		panic(fmt.Errorf("block entity count: got %d, want 1", len(pkt.ChunkData.BlockEntities)))
	}
	be := pkt.ChunkData.BlockEntities[0]
	// has to be minecraft:sign even if it's pale_oak_wall_sign - they're generic in the registry
	signEntityType := ns.VarInt(registries.BlockEntityType.Get("minecraft:sign"))
	if be.Type != signEntityType {
		panic(fmt.Errorf("block entity type: got %d, want %d (sign)", be.Type, signEntityType))
	}

	// validate sign front text messages: "needle", "in", "a", "haystack"
	compound, ok := be.Data.(nbt.Compound)
	if !ok {
		panic(fmt.Errorf("block entity data is %T, want nbt.Compound", be.Data))
	}
	frontText, ok := compound["front_text"].(nbt.Compound)
	if !ok {
		panic("missing front_text compound in sign block entity")
	}
	messages, ok := frontText["messages"].(nbt.List)
	if !ok {
		panic("missing messages list in sign front_text")
	}
	expected := []string{"needle", "in", "a", "haystack"}
	if len(messages.Elements) != len(expected) {
		panic(fmt.Errorf("sign messages count: got %d, want %d", len(messages.Elements), len(expected)))
	}
	for i, want := range expected {
		got := string(messages.Elements[i].(nbt.String))
		if got != want {
			panic(fmt.Errorf("sign message %d: got %q, want %q", i, got, want))
		}
	}
}
