package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-mclib/data/pkg/data/blocks"
	"github.com/go-mclib/data/pkg/data/chunks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegionFileRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "r.0.0.mca")

	r, err := OpenRegion(path)
	require.NoError(t, err)
	defer r.Close()

	// write some data
	data := []byte("hello chunk")
	require.NoError(t, r.WriteChunk(3, 5, data))

	// verify it can be read back
	assert.True(t, r.HasChunk(3, 5))
	assert.False(t, r.HasChunk(0, 0))

	got, err := r.ReadChunk(3, 5)
	require.NoError(t, err)
	assert.Equal(t, data, got)

	// absent chunk returns nil
	got, err = r.ReadChunk(0, 0)
	require.NoError(t, err)
	assert.Nil(t, got)
}

func TestRegionFileOverwrite(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "r.0.0.mca")

	r, err := OpenRegion(path)
	require.NoError(t, err)
	defer r.Close()

	require.NoError(t, r.WriteChunk(1, 1, []byte("first")))
	require.NoError(t, r.WriteChunk(1, 1, []byte("second")))

	got, err := r.ReadChunk(1, 1)
	require.NoError(t, err)
	assert.Equal(t, []byte("second"), got)
}

func TestRegionFilePersistence(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "r.0.0.mca")

	// write and close
	r, err := OpenRegion(path)
	require.NoError(t, err)
	require.NoError(t, r.WriteChunk(10, 20, []byte("persistent")))
	r.Close()

	// reopen and read
	r2, err := OpenRegion(path)
	require.NoError(t, err)
	defer r2.Close()

	got, err := r2.ReadChunk(10, 20)
	require.NoError(t, err)
	assert.Equal(t, []byte("persistent"), got)
}

func TestRegionStorage(t *testing.T) {
	dir := t.TempDir()
	regionDir := filepath.Join(dir, "region")

	rs, err := NewRegionStorage(regionDir)
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
	col := &chunks.ChunkColumn{
		X:          5,
		Z:          -3,
		Heightmaps: make(map[int32][]int64),
	}
	for i := range chunks.SectionCount {
		col.Sections[i] = chunks.NewEmptySection()
	}

	// place some blocks in section 0 (Y=-64 to -49)
	stoneID := blocks.DefaultStateID(blocks.BlockID("minecraft:stone"))
	require.Greater(t, stoneID, int32(0))

	col.Sections[0].SetBlockState(0, 0, 0, stoneID)
	col.Sections[0].SetBlockState(1, 0, 0, stoneID)
	col.Sections[0].BlockCount = 2

	// compute sky light so it round-trips
	col.ComputeSkylight()

	// serialize to NBT
	data, err := ChunkToNBT(col)
	require.NoError(t, err)
	require.NotEmpty(t, data)

	// deserialize
	loaded, err := NBTToChunk(data)
	require.NoError(t, err)
	assert.Equal(t, col.X, loaded.X)
	assert.Equal(t, col.Z, loaded.Z)

	// verify blocks
	assert.Equal(t, stoneID, loaded.GetBlockState(0, -64, 0))
	assert.Equal(t, stoneID, loaded.GetBlockState(1, -64, 0))
	assert.Equal(t, int32(0), loaded.GetBlockState(2, -64, 0)) // air
	assert.Equal(t, int16(2), loaded.Sections[0].BlockCount)
}

func TestChunkNBTWithRegionStorage(t *testing.T) {
	dir := t.TempDir()
	regionDir := filepath.Join(dir, "region")
	rs, err := NewRegionStorage(regionDir)
	require.NoError(t, err)
	defer rs.Close()

	// generate a simple superflat chunk
	col := &chunks.ChunkColumn{
		X:          0,
		Z:          0,
		Heightmaps: make(map[int32][]int64),
	}
	for i := range chunks.SectionCount {
		col.Sections[i] = chunks.NewEmptySection()
	}
	bedrockID := blocks.DefaultStateID(blocks.BlockID("minecraft:bedrock"))
	for x := range 16 {
		for z := range 16 {
			col.Sections[0].SetBlockState(x, 0, z, bedrockID)
		}
	}
	col.Sections[0].BlockCount = 256
	col.ComputeSkylight()

	// save via region storage
	data, err := ChunkToNBT(col)
	require.NoError(t, err)
	require.NoError(t, rs.WriteChunk(0, 0, data))

	// load back
	loadedData, err := rs.ReadChunk(0, 0)
	require.NoError(t, err)
	loaded, err := NBTToChunk(loadedData)
	require.NoError(t, err)

	// verify bedrock layer
	assert.Equal(t, bedrockID, loaded.GetBlockState(0, -64, 0))
	assert.Equal(t, bedrockID, loaded.GetBlockState(15, -64, 15))
	assert.Equal(t, int32(0), loaded.GetBlockState(0, -63, 0)) // air above
}

func TestPlayerDataRoundTrip(t *testing.T) {
	dir := t.TempDir()
	pdDir := filepath.Join(dir, "playerdata")

	pd := &PlayerData{
		X: 100.5, Y: 65.0, Z: -200.3,
		Yaw: 45.0, Pitch: -10.0,
		Dimension: "minecraft:overworld",
		Gamemode:  1,
		HeldSlot:  3,
		Inventory: []InventorySlot{
			{Slot: 0, ID: "minecraft:stone", Count: 64},
			{Slot: 36, ID: "minecraft:diamond_sword", Count: 1},
		},
	}

	uuid := "12345678-1234-1234-1234-123456789abc"
	require.NoError(t, SavePlayer(pdDir, uuid, pd))

	// verify file exists
	_, err := os.Stat(filepath.Join(pdDir, uuid+".dat"))
	require.NoError(t, err)

	// load back
	loaded, err := LoadPlayer(pdDir, uuid)
	require.NoError(t, err)
	require.NotNil(t, loaded)

	assert.InDelta(t, pd.X, loaded.X, 0.001)
	assert.InDelta(t, pd.Y, loaded.Y, 0.001)
	assert.InDelta(t, pd.Z, loaded.Z, 0.001)
	assert.InDelta(t, pd.Yaw, loaded.Yaw, 0.001)
	assert.InDelta(t, pd.Pitch, loaded.Pitch, 0.001)
	assert.Equal(t, pd.Gamemode, loaded.Gamemode)
	assert.Equal(t, pd.HeldSlot, loaded.HeldSlot)
	assert.Equal(t, pd.Dimension, loaded.Dimension)
	assert.Len(t, loaded.Inventory, 2)
	assert.Equal(t, "minecraft:stone", loaded.Inventory[0].ID)
	assert.Equal(t, int32(64), loaded.Inventory[0].Count)
}

func TestLoadPlayerNotFound(t *testing.T) {
	dir := t.TempDir()
	pd, err := LoadPlayer(dir, "nonexistent")
	require.NoError(t, err)
	assert.Nil(t, pd)
}
