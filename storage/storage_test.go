package storage_test

import (
	"path/filepath"
	"testing"

	"github.com/go-mclib/data/blocks"
	"github.com/go-mclib/data/chunks"
	"github.com/go-mclib/data/storage"
	"github.com/go-mclib/nbt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegionStorage(t *testing.T) {
	dir := t.TempDir()
	regionDir := filepath.Join(dir, "region")

	rs, err := storage.NewRegionStorage(regionDir)
	require.NoError(t, err)
	defer rs.Close()

	require.NoError(t, rs.WriteChunk(0, 0, []byte("chunk-0-0")))
	require.NoError(t, rs.WriteChunk(33, 33, []byte("chunk-33-33"))) // different region

	got, err := rs.ReadChunk(0, 0)
	require.NoError(t, err)
	assert.Equal(t, []byte("chunk-0-0"), got)

	got, err = rs.ReadChunk(33, 33)
	require.NoError(t, err)
	assert.Equal(t, []byte("chunk-33-33"), got)
}

func TestChunkNBTRoundTrip(t *testing.T) {
	// build a simple chunk with some blocks
	col := &chunks.Column{
		X:          5,
		Z:          -3,
		MinY:       chunks.MinY,
		Sections:   make([]*chunks.Section, chunks.SectionCount),
		Heightmaps: make(map[int32][]int64),
	}
	for i := range chunks.SectionCount {
		col.Sections[i] = chunks.NewEmptySection()
	}

	// place some blocks in section 0 (Y=-64 to -49)
	stoneID := blocks.DefaultStateID(blocks.ID("minecraft:stone"))
	require.Greater(t, stoneID, int32(0))

	col.Sections[0].SetBlockState(0, 0, 0, stoneID)
	col.Sections[0].SetBlockState(1, 0, 0, stoneID)
	logID := blocks.StateID(blocks.ID("minecraft:oak_log"), map[string]string{"axis": "x"})
	col.Sections[0].SetBlockState(3, 0, 0, logID)
	col.Sections[0].BlockCount = 3

	// compute sky light so it round-trips
	col.ComputeSkylight()

	// serialize to NBT
	data, err := storage.ChunkToNBT(col)
	require.NoError(t, err)
	tag, _, err := nbt.DecodeFile(data, nbt.WithMaxBytes(0))
	require.NoError(t, err)
	root := tag.(nbt.Compound)
	assert.Equal(t, int32(5023), root.GetInt("DataVersion"))
	var savedLog nbt.Compound
	for _, sectionTag := range root.GetList("sections").Elements {
		section := sectionTag.(nbt.Compound)
		for _, paletteTag := range section.GetCompound("block_states").GetList("palette").Elements {
			entry := paletteTag.(nbt.Compound)
			assert.NotContains(t, entry, "Name")
			assert.NotContains(t, entry, "Properties")
			if entry.GetString("id") == "minecraft:oak_log" {
				savedLog = entry
			}
		}
	}
	require.NotNil(t, savedLog)
	assert.Equal(t, "x", savedLog.GetCompound("properties").GetString("axis"))

	// deserialize
	loaded, err := storage.NBTToChunk(data)
	require.NoError(t, err)
	assert.Equal(t, col.X, loaded.X)
	assert.Equal(t, col.Z, loaded.Z)

	// verify blocks
	assert.Equal(t, stoneID, loaded.GetBlockState(0, -64, 0))
	assert.Equal(t, stoneID, loaded.GetBlockState(1, -64, 0))
	assert.Equal(t, int32(0), loaded.GetBlockState(2, -64, 0)) // air
	assert.Equal(t, logID, loaded.GetBlockState(3, -64, 0))
	assert.Equal(t, int16(3), loaded.Sections[0].BlockCount)
}

func TestChunkNBTDimensionHeight(t *testing.T) {
	const sectionCount = 16
	col := &chunks.Column{
		X: 2, Z: 3, MinY: 0,
		Sections: make([]*chunks.Section, sectionCount), Heightmaps: make(map[int32][]int64),
	}
	for i := range sectionCount {
		col.Sections[i] = chunks.NewEmptySection()
	}
	stoneID := blocks.DefaultStateID(blocks.ID("minecraft:stone"))
	col.SetBlockState(1, 0, 1, stoneID)
	col.Sections[0].BlockCount = 1
	col.ComputeSkylight()

	data, err := storage.ChunkToNBT(col)
	require.NoError(t, err)
	loaded, err := storage.NBTToChunkWithHeight(data, 0, sectionCount)
	require.NoError(t, err)
	assert.Equal(t, 0, loaded.MinY)
	assert.Len(t, loaded.Sections, sectionCount)
	assert.Equal(t, stoneID, loaded.GetBlockState(1, 0, 1))
	require.NotNil(t, loaded.Light)
	assert.True(t, loaded.Light.SkyLightMask.Get(sectionCount+1))
	assert.Len(t, loaded.Light.SkyLightArrays, sectionCount+1)
}

func TestChunkNBTWithRegionStorage(t *testing.T) {
	dir := t.TempDir()
	regionDir := filepath.Join(dir, "region")
	rs, err := storage.NewRegionStorage(regionDir)
	require.NoError(t, err)
	defer rs.Close()

	// generate a simple superflat chunk
	col := &chunks.Column{
		X:          0,
		Z:          0,
		MinY:       chunks.MinY,
		Sections:   make([]*chunks.Section, chunks.SectionCount),
		Heightmaps: make(map[int32][]int64),
	}
	for i := range chunks.SectionCount {
		col.Sections[i] = chunks.NewEmptySection()
	}
	bedrockID := blocks.DefaultStateID(blocks.ID("minecraft:bedrock"))
	for x := range 16 {
		for z := range 16 {
			col.Sections[0].SetBlockState(x, 0, z, bedrockID)
		}
	}
	col.Sections[0].BlockCount = 256
	col.ComputeSkylight()

	// save via region storage
	data, err := storage.ChunkToNBT(col)
	require.NoError(t, err)
	require.NoError(t, rs.WriteChunk(0, 0, data))

	// load back
	loadedData, err := rs.ReadChunk(0, 0)
	require.NoError(t, err)
	loaded, err := storage.NBTToChunk(loadedData)
	require.NoError(t, err)

	// verify bedrock layer
	assert.Equal(t, bedrockID, loaded.GetBlockState(0, -64, 0))
	assert.Equal(t, bedrockID, loaded.GetBlockState(15, -64, 15))
	assert.Equal(t, int32(0), loaded.GetBlockState(0, -63, 0)) // air above
}
