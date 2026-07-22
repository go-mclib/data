package storage_test

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-mclib/data/blocks"
	"github.com/go-mclib/data/chunks"
	"github.com/go-mclib/data/storage"
	"github.com/go-mclib/nbt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegionFileRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "r.0.0.mca")

	r, err := storage.OpenRegion(path)
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

	r, err := storage.OpenRegion(path)
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
	r, err := storage.OpenRegion(path)
	require.NoError(t, err)
	require.NoError(t, r.WriteChunk(10, 20, []byte("persistent")))
	r.Close()

	// reopen and read
	r2, err := storage.OpenRegion(path)
	require.NoError(t, err)
	defer r2.Close()

	got, err := r2.ReadChunk(10, 20)
	require.NoError(t, err)
	assert.Equal(t, []byte("persistent"), got)
}

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
	col := &chunks.ChunkColumn{
		X:          5,
		Z:          -3,
		MinY:       chunks.MinY,
		Sections:   make([]*chunks.ChunkSection, chunks.SectionCount),
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
	data, err := storage.ChunkToNBT(col)
	require.NoError(t, err)
	require.NotEmpty(t, data)

	// deserialize
	loaded, err := storage.NBTToChunk(data)
	require.NoError(t, err)
	assert.Equal(t, col.X, loaded.X)
	assert.Equal(t, col.Z, loaded.Z)

	// verify blocks
	assert.Equal(t, stoneID, loaded.GetBlockState(0, -64, 0))
	assert.Equal(t, stoneID, loaded.GetBlockState(1, -64, 0))
	assert.Equal(t, int32(0), loaded.GetBlockState(2, -64, 0)) // air
	assert.Equal(t, int16(2), loaded.Sections[0].BlockCount)
}

func TestChunkNBTDimensionHeight(t *testing.T) {
	const sectionCount = 16
	col := &chunks.ChunkColumn{
		X: 2, Z: 3, MinY: 0,
		Sections: make([]*chunks.ChunkSection, sectionCount), Heightmaps: make(map[int32][]int64),
	}
	for i := range sectionCount {
		col.Sections[i] = chunks.NewEmptySection()
	}
	stoneID := blocks.DefaultStateID(blocks.BlockID("minecraft:stone"))
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

func TestCurrentDataVersion(t *testing.T) {
	assert.Equal(t, 4903, storage.DataVersion)
}

func TestChunkNBTWithRegionStorage(t *testing.T) {
	dir := t.TempDir()
	regionDir := filepath.Join(dir, "region")
	rs, err := storage.NewRegionStorage(regionDir)
	require.NoError(t, err)
	defer rs.Close()

	// generate a simple superflat chunk
	col := &chunks.ChunkColumn{
		X:          0,
		Z:          0,
		MinY:       chunks.MinY,
		Sections:   make([]*chunks.ChunkSection, chunks.SectionCount),
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

func TestPlayerDataRoundTrip(t *testing.T) {
	dir := t.TempDir()
	pdDir := filepath.Join(dir, "playerdata")

	pd := &storage.PlayerData{
		X: 100.5, Y: 65.0, Z: -200.3,
		Yaw: 45.0, Pitch: -10.0,
		Dimension:        "minecraft:overworld",
		Gamemode:         1,
		HeldSlot:         3,
		Tags:             []string{"builder", "quest_started"},
		FireTicks:        157,
		AirSupply:        123,
		TicksFrozen:      91,
		FallDistance:     6.25,
		AbsorptionAmount: 5.5,
		FallFlying:       true,
		Food:             17,
		Saturation:       2.5,
		FoodExhaustion:   3.75,
		FoodTickTimer:    61,
		Inventory: []storage.InventorySlot{
			{Slot: 0, ID: "minecraft:stone", Count: 64},
			{Slot: 9, ID: "minecraft:diamond_sword", Count: 1, Components: nbt.Compound{"minecraft:damage": nbt.Int(17)}},
		},
		Equipment: []storage.EquipmentSlot{
			{Slot: "head", ID: "minecraft:iron_helmet", Count: 1},
		},
		Effects: []storage.EffectData{{
			Name: "minecraft:speed", Duration: 200, Amplifier: 1, ShowParticles: true, ShowIcon: true,
			Hidden: &storage.EffectData{Name: "minecraft:speed", Duration: 400, ShowParticles: true, ShowIcon: true},
		}},
		RecipeBook: storage.RecipeBookData{
			CraftingOpen: true, FurnaceFiltering: true,
			Recipes: []string{"minecraft:oak_planks"}, ToBeDisplayed: []string{"minecraft:oak_planks"},
		},
		HasRespawn: true, RespawnDimension: "minecraft:overworld",
		RespawnX: 12, RespawnY: 70, RespawnZ: -4, RespawnYaw: 90, RespawnPitch: 10,
	}

	uuid := "12345678-1234-1234-1234-123456789abc"
	require.NoError(t, storage.SavePlayer(pdDir, uuid, pd))

	// verify the 26.2 inventory/equipment layout on disk
	path := filepath.Join(pdDir, uuid+".dat")
	file, err := os.Open(path)
	require.NoError(t, err)
	compressed, err := gzip.NewReader(file)
	require.NoError(t, err)
	raw, err := io.ReadAll(compressed)
	require.NoError(t, err)
	require.NoError(t, compressed.Close())
	require.NoError(t, file.Close())
	tag, _, err := nbt.DecodeFile(raw, nbt.WithMaxBytes(0))
	require.NoError(t, err)
	root, ok := tag.(nbt.Compound)
	require.True(t, ok)
	for _, entry := range root.GetList("Inventory").Elements {
		item := entry.(nbt.Compound)
		assert.Less(t, item.GetByte("Slot"), int8(36))
		if item.GetString("id") == "minecraft:diamond_sword" {
			assert.Equal(t, int32(17), item.GetCompound("components").GetInt("minecraft:damage"))
		}
	}
	head := root.GetCompound("equipment").GetCompound("head")
	assert.Equal(t, "minecraft:iron_helmet", head.GetString("id"))
	assert.Equal(t, int16(157), root.GetShort("Fire"))
	assert.Equal(t, int8(1), root.GetByte("FallFlying"))
	effect := root.GetList("active_effects").Get(0).(nbt.Compound)
	assert.Equal(t, "minecraft:speed", effect.GetString("id"))
	assert.Equal(t, int32(200), effect.GetInt("duration"))
	hiddenEffect := effect.GetCompound("hidden_effect")
	assert.Equal(t, int32(400), hiddenEffect.GetInt("duration"))
	assert.Nil(t, hiddenEffect.Get("id"))
	recipeBook := root.GetCompound("recipeBook")
	assert.Equal(t, int8(1), recipeBook.GetByte("isGuiOpen"))
	assert.Equal(t, int8(1), recipeBook.GetByte("isFurnaceFilteringCraftable"))
	assert.Equal(t, "minecraft:oak_planks", string(recipeBook.GetList("recipes").Get(0).(nbt.String)))
	respawn := root.GetCompound("respawn")
	assert.Equal(t, "minecraft:overworld", respawn.GetString("dimension"))
	assert.Equal(t, []int32{12, 70, -4}, respawn.GetIntArray("pos"))

	// load back
	loaded, err := storage.LoadPlayer(pdDir, uuid)
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
	assert.Equal(t, pd.Tags, loaded.Tags)
	assert.Equal(t, pd.FireTicks, loaded.FireTicks)
	assert.Equal(t, pd.AirSupply, loaded.AirSupply)
	assert.Equal(t, pd.TicksFrozen, loaded.TicksFrozen)
	assert.Equal(t, pd.FallDistance, loaded.FallDistance)
	assert.Equal(t, pd.AbsorptionAmount, loaded.AbsorptionAmount)
	assert.Equal(t, pd.FallFlying, loaded.FallFlying)
	assert.Equal(t, pd.Food, loaded.Food)
	assert.Equal(t, pd.Saturation, loaded.Saturation)
	assert.Equal(t, pd.FoodExhaustion, loaded.FoodExhaustion)
	assert.Equal(t, pd.FoodTickTimer, loaded.FoodTickTimer)
	assert.Len(t, loaded.Inventory, 2)
	assert.Equal(t, "minecraft:stone", loaded.Inventory[0].ID)
	assert.Equal(t, int32(64), loaded.Inventory[0].Count)
	assert.Equal(t, int32(17), loaded.Inventory[1].Components.GetInt("minecraft:damage"))
	require.Len(t, loaded.Equipment, 1)
	assert.Equal(t, "head", loaded.Equipment[0].Slot)
	assert.Equal(t, "minecraft:iron_helmet", loaded.Equipment[0].ID)
	assert.Equal(t, pd.Effects, loaded.Effects)
	assert.Equal(t, pd.RecipeBook, loaded.RecipeBook)
	assert.True(t, loaded.HasRespawn)
	assert.Equal(t, pd.RespawnDimension, loaded.RespawnDimension)
	assert.Equal(t, pd.RespawnX, loaded.RespawnX)
	assert.Equal(t, pd.RespawnY, loaded.RespawnY)
	assert.Equal(t, pd.RespawnZ, loaded.RespawnZ)
	assert.Equal(t, pd.RespawnYaw, loaded.RespawnYaw)
	assert.Equal(t, pd.RespawnPitch, loaded.RespawnPitch)
}

func TestLoadPlayerNotFound(t *testing.T) {
	dir := t.TempDir()
	pd, err := storage.LoadPlayer(dir, "nonexistent")
	require.NoError(t, err)
	assert.Nil(t, pd)
}

func TestLevelDataBorderRoundTrip(t *testing.T) {
	dir := t.TempDir()
	want := &storage.LevelData{
		Seed: 42, BorderSet: true,
		BorderCenterX: 12.5, BorderCenterZ: -31.25,
		BorderSize: 400, BorderSizeLerpTime: 2750, BorderSizeLerpTarget: 125,
		BorderSafeZone: 3.5, BorderDamagePerBlock: 0.75,
		BorderWarningBlocks: 9, BorderWarningTime: 21,
	}
	require.NoError(t, storage.SaveLevel(dir, want))

	got, err := storage.LoadLevel(dir)
	require.NoError(t, err)
	assert.True(t, got.BorderSet)
	assert.Equal(t, want.BorderCenterX, got.BorderCenterX)
	assert.Equal(t, want.BorderCenterZ, got.BorderCenterZ)
	assert.Equal(t, want.BorderSize, got.BorderSize)
	assert.Equal(t, want.BorderSizeLerpTime, got.BorderSizeLerpTime)
	assert.Equal(t, want.BorderSizeLerpTarget, got.BorderSizeLerpTarget)
	assert.Equal(t, want.BorderSafeZone, got.BorderSafeZone)
	assert.Equal(t, want.BorderDamagePerBlock, got.BorderDamagePerBlock)
	assert.Equal(t, want.BorderWarningBlocks, got.BorderWarningBlocks)
	assert.Equal(t, want.BorderWarningTime, got.BorderWarningTime)
}
