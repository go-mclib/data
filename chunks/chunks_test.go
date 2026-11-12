package chunks_test

import (
	"testing"

	"github.com/go-mclib/data/blocks"
	"github.com/go-mclib/data/chunks"
	ns "github.com/go-mclib/protocol/javaproto/wire"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPalettedContainerSingleValue(t *testing.T) {
	c := chunks.NewSingleValue(chunks.BlockStatesKind, 42)
	assert.Equal(t, 0, c.BitsPerEntry())
	// all positions should return 42
	for x := range 16 {
		for y := range 16 {
			for z := range 16 {
				assert.Equal(t, int32(42), c.GetXYZ(x, y, z))
			}
		}
	}
}

func TestPalettedContainerSetExpandsFromSingleValue(t *testing.T) {
	c := chunks.NewSingleValue(chunks.BlockStatesKind, 10)
	// setting same value should be a no-op
	c.SetXYZ(0, 0, 0, 10)
	assert.Equal(t, 0, c.BitsPerEntry())

	// setting different value should expand to indirect palette
	c.SetXYZ(5, 3, 7, 99)
	assert.Equal(t, 4, c.BitsPerEntry()) // minIndirectBits for blocks
	assert.Equal(t, int32(99), c.GetXYZ(5, 3, 7))
	// original values should still be accessible
	assert.Equal(t, int32(10), c.GetXYZ(0, 0, 0))
	assert.Equal(t, int32(10), c.GetXYZ(15, 15, 15))
}

func TestPalettedContainerDecodeEncodeSingleValue(t *testing.T) {
	// encode a single-value container, then decode and verify
	original := chunks.NewSingleValue(chunks.BlockStatesKind, 7)
	buf := ns.NewWriter()
	require.NoError(t, original.Encode(buf))

	decoded := &chunks.PalettedContainer{}
	*decoded = chunks.PalettedContainer{} // init with kind via Decode
	// we need to set the kind before decoding — use a helper approach
	roundTripped := chunks.NewSingleValue(chunks.BlockStatesKind, 0)
	require.NoError(t, roundTripped.Decode(ns.NewReader(buf.Bytes())))

	assert.Equal(t, 0, roundTripped.BitsPerEntry())
	assert.Equal(t, int32(7), roundTripped.GetXYZ(0, 0, 0))
}

func TestPalettedContainerDecodeEncodeIndirect(t *testing.T) {
	// build a container with a few different block states
	c := chunks.NewSingleValue(chunks.BlockStatesKind, 1) // stone
	c.SetXYZ(0, 0, 0, 2)                                  // granite
	c.SetXYZ(1, 0, 0, 3)                                  // polished granite
	c.SetXYZ(0, 1, 0, 4)                                  // diorite

	assert.Equal(t, 4, c.BitsPerEntry())

	// encode
	buf := ns.NewWriter()
	require.NoError(t, c.Encode(buf))

	// decode into new container
	decoded := chunks.NewSingleValue(chunks.BlockStatesKind, 0)
	require.NoError(t, decoded.Decode(ns.NewReader(buf.Bytes())))

	// verify all positions
	assert.Equal(t, int32(2), decoded.GetXYZ(0, 0, 0))
	assert.Equal(t, int32(3), decoded.GetXYZ(1, 0, 0))
	assert.Equal(t, int32(4), decoded.GetXYZ(0, 1, 0))
	assert.Equal(t, int32(1), decoded.GetXYZ(5, 5, 5)) // untouched = stone
}

func TestPalettedContainerBiomes(t *testing.T) {
	c := chunks.NewSingleValue(chunks.BiomesKind, 0)
	c.SetXYZ(0, 0, 0, 5)
	assert.Equal(t, 1, c.BitsPerEntry()) // minIndirectBits for biomes
	assert.Equal(t, int32(5), c.GetXYZ(0, 0, 0))
	assert.Equal(t, int32(0), c.GetXYZ(3, 3, 3))

	// round-trip
	buf := ns.NewWriter()
	require.NoError(t, c.Encode(buf))
	decoded := chunks.NewSingleValue(chunks.BiomesKind, 0)
	require.NoError(t, decoded.Decode(ns.NewReader(buf.Bytes())))
	assert.Equal(t, int32(5), decoded.GetXYZ(0, 0, 0))
	assert.Equal(t, int32(0), decoded.GetXYZ(3, 3, 3))
}

// TestPalettedContainerBiomesDirect covers a direct (global-palette) biome
// container. With >8 distinct biomes a section overflows the indirect palette
// and packs with the global-palette width (ceil(log2(66 biomes)) = 7 bits in
// 26.2). Regression: the decoder used to hardcode 6 bits here, reading one long
// too few and desyncing the rest of the column.
func TestPalettedContainerBiomesDirect(t *testing.T) {
	c := chunks.NewSingleValue(chunks.BiomesKind, 0)
	for i := range 9 { // 9 distinct values force the indirect palette to direct
		c.SetXYZ(i, 0, 0, int32(10+i))
	}
	assert.Equal(t, 7, c.BitsPerEntry(), "9 distinct biomes should use the 7-bit global palette")

	buf := ns.NewWriter()
	require.NoError(t, c.Encode(buf))

	decoded := chunks.NewSingleValue(chunks.BiomesKind, 0)
	require.NoError(t, decoded.Decode(ns.NewReader(buf.Bytes())))
	assert.Equal(t, 7, decoded.BitsPerEntry())
	for i := range 9 {
		assert.Equal(t, int32(10+i), decoded.GetXYZ(i, 0, 0))
	}
}

// TestChunkSectionDirectBiomeNoDesync proves a section whose biome container is
// direct does not desync the next section read from the same buffer.
func TestChunkSectionDirectBiomeNoDesync(t *testing.T) {
	first := chunks.NewEmptySection()
	for i := range 9 {
		first.SetBiome(i, 0, 0, int32(20+i)) // force direct biome palette
	}
	assert.Equal(t, 7, first.Biomes.BitsPerEntry())

	second := chunks.NewEmptySection()
	second.BlockCount = 3
	second.FluidCount = 7
	second.SetBlockState(2, 2, 2, 1234)

	buf := ns.NewWriter()
	require.NoError(t, first.Encode(buf))
	require.NoError(t, second.Encode(buf))

	r := ns.NewReader(buf.Bytes())
	d1, d2 := &chunks.Section{}, &chunks.Section{}
	require.NoError(t, d1.Decode(r))
	require.NoError(t, d2.Decode(r), "second section must stay aligned after a direct biome container")
	assert.Equal(t, int32(20), d1.GetBiome(0, 0, 0))
	assert.Equal(t, int16(3), d2.BlockCount)
	assert.Equal(t, int16(7), d2.FluidCount)
	assert.Equal(t, int32(1234), d2.GetBlockState(2, 2, 2))
}

func TestChunkSectionDecodeEncode(t *testing.T) {
	sec := chunks.NewEmptySection()
	sec.BlockCount = 2
	sec.BlockStates.SetXYZ(0, 0, 0, 1) // stone
	sec.BlockStates.SetXYZ(1, 0, 0, 2) // granite

	buf := ns.NewWriter()
	require.NoError(t, sec.Encode(buf))

	decoded := &chunks.Section{}
	require.NoError(t, decoded.Decode(ns.NewReader(buf.Bytes())))

	assert.Equal(t, int16(2), decoded.BlockCount)
	assert.Equal(t, int32(1), decoded.GetBlockState(0, 0, 0))
	assert.Equal(t, int32(2), decoded.GetBlockState(1, 0, 0))
	assert.Equal(t, int32(0), decoded.GetBlockState(5, 5, 5)) // air
}

func TestChunkColumnGetSetBlockState(t *testing.T) {
	col := chunks.NewColumn(0, 0)
	// Y=-64 is section index 0
	col.SetBlockState(0, -64, 0, 1)
	assert.Equal(t, int32(1), col.GetBlockState(0, -64, 0))

	// Y=0 is section index 4
	col.SetBlockState(3, 0, 5, 42)
	assert.Equal(t, int32(42), col.GetBlockState(3, 0, 5))

	// unset section returns air
	assert.Equal(t, int32(0), col.GetBlockState(0, 100, 0))

	// out of range returns 0
	assert.Equal(t, int32(0), col.GetBlockState(0, -65, 0))
	assert.Equal(t, int32(0), col.GetBlockState(0, 320, 0))
}

func TestChunkColumnEncodeDecodeSections(t *testing.T) {
	col := chunks.NewColumn(0, 0)
	col.SetBlockState(5, 10, 3, 100)
	col.SetBlockState(0, -64, 0, 1)

	data, err := col.EncodeSections()
	require.NoError(t, err)

	// parse it back via a ChunkData-like flow
	col2 := chunks.NewColumn(0, 0)
	buf := ns.NewReader(data)
	for i := range chunks.SectionCount {
		sec := &chunks.Section{}
		require.NoError(t, sec.Decode(buf))
		col2.Sections[i] = sec
	}

	assert.Equal(t, int32(100), col2.GetBlockState(5, 10, 3))
	assert.Equal(t, int32(1), col2.GetBlockState(0, -64, 0))
}

func TestSectionIndex(t *testing.T) {
	tests := []struct {
		y    int
		want int
	}{
		{-64, 0},
		{-49, 0},
		{-48, 1},
		{0, 4},
		{255, 19},
		{319, 23},
		{-65, -1},
		{320, -1},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, chunks.SectionIndex(tt.y), "SectionIndex(%d)", tt.y)
	}
}

func TestLocalCoords(t *testing.T) {
	lx, ly, lz := chunks.LocalCoords(17, -60, 35)
	assert.Equal(t, 1, lx)
	assert.Equal(t, 4, ly)
	assert.Equal(t, 3, lz)
}

func TestChunkPos(t *testing.T) {
	cx, cz := chunks.ChunkPos(100, -200)
	assert.Equal(t, int32(6), cx)
	assert.Equal(t, int32(-13), cz)
}

func TestDecodeSectionPosition(t *testing.T) {
	// encode a known section position and decode it
	// sectionX=3, sectionY=-1, sectionZ=5
	// packed = (3 << 42) | (5 << 20) | ((-1) & 0xFFFFF)
	packed := (int64(3) << 42) | (int64(5) << 20) | (int64(0xFFFFF))
	sx, sy, sz := chunks.DecodeSectionPosition(packed)
	assert.Equal(t, int32(3), sx)
	assert.Equal(t, int32(-1), sy)
	assert.Equal(t, int32(5), sz)
}

func TestDecodeSectionPositionNegative(t *testing.T) {
	// negative X and Z
	// sectionX=-2, sectionY=3, sectionZ=-4
	x := int64(-2) & 0x3FFFFF // 22-bit mask
	z := int64(-4) & 0x3FFFFF
	y := int64(3) & 0xFFFFF // 20-bit mask
	packed := (x << 42) | (z << 20) | y
	sx, sy, sz := chunks.DecodeSectionPosition(packed)
	assert.Equal(t, int32(-2), sx)
	assert.Equal(t, int32(3), sy)
	assert.Equal(t, int32(-4), sz)
}

func TestDecodeBlockEntry(t *testing.T) {
	// blockState=42, localX=3, localZ=7, localY=12
	// entry = (42 << 12) | (3 << 8) | (7 << 4) | 12
	entry := int64(42<<12) | int64(3<<8) | int64(7<<4) | 12
	stateID, lx, ly, lz := chunks.DecodeBlockEntry(entry)
	assert.Equal(t, int32(42), stateID)
	assert.Equal(t, 3, lx)
	assert.Equal(t, 12, ly)
	assert.Equal(t, 7, lz)
}

func TestChunkBlockLightPropagatesAndClears(t *testing.T) {
	col := chunks.NewColumn(0, 0)
	torch := blocks.DefaultStateID(blocks.ID("minecraft:torch"))
	stone := blocks.DefaultStateID(blocks.ID("minecraft:stone"))
	col.SetBlockState(8, 64, 8, torch)
	col.SetBlockState(9, 64, 8, stone)
	col.ComputeSkylight()

	assert.Equal(t, byte(14), lightAt(col.Light.BlockLightMask, col.Light.BlockLightArrays, 8, 64, 8))
	assert.Equal(t, byte(13), lightAt(col.Light.BlockLightMask, col.Light.BlockLightArrays, 7, 64, 8))
	assert.Equal(t, byte(0), lightAt(col.Light.BlockLightMask, col.Light.BlockLightArrays, 9, 64, 8))
	assert.Equal(t, byte(1), lightAt(col.Light.BlockLightMask, col.Light.BlockLightArrays, 8, 77, 8))

	col.SetBlockState(8, 64, 8, 0)
	col.ComputeSkylight()
	assert.Empty(t, col.Light.BlockLightArrays)
}

func TestChunkBlockLightUsesDynamicStateEmission(t *testing.T) {
	col := chunks.NewColumn(0, 0)
	furnaceID := blocks.ID("minecraft:furnace")
	lit := blocks.StateID(furnaceID, map[string]string{"facing": "north", "lit": "true"})
	unlit := blocks.StateID(furnaceID, map[string]string{"facing": "north", "lit": "false"})
	col.SetBlockState(4, 20, 4, lit)
	col.ComputeSkylight()
	assert.Equal(t, byte(13), lightAt(col.Light.BlockLightMask, col.Light.BlockLightArrays, 4, 20, 4))

	col.SetBlockState(4, 20, 4, unlit)
	col.ComputeSkylight()
	assert.Empty(t, col.Light.BlockLightArrays)
}

func lightAt(mask ns.BitSet, arrays [][]byte, x, y, z int) byte {
	section := chunks.SectionIndex(y) + 1
	arrayIndex := 0
	for i := range section {
		if mask.Get(i) {
			arrayIndex++
		}
	}
	if !mask.Get(section) {
		return 0
	}
	index := ((y&15)*16+z)*16 + x
	value := arrays[arrayIndex][index/2]
	if index&1 == 0 {
		return value & 15
	}
	return value >> 4
}

// TestPalettedContainerBlockStatesDirectHighIDs covers direct block-state
// storage with state IDs above 32767, which need the 16-bit global palette.
func TestPalettedContainerBlockStatesDirectHighIDs(t *testing.T) {
	c := chunks.NewSingleValue(chunks.BlockStatesKind, 0)
	for i := range 300 { // >256 distinct values force the direct palette
		c.Set(i, int32(35000+i))
	}
	assert.Equal(t, 16, c.BitsPerEntry())

	buf := ns.NewWriter()
	require.NoError(t, c.Encode(buf))
	decoded := chunks.NewSingleValue(chunks.BlockStatesKind, 0)
	require.NoError(t, decoded.Decode(ns.NewReader(buf.Bytes())))
	for i := range 300 {
		assert.Equal(t, int32(35000+i), decoded.Get(i))
	}
	assert.Equal(t, int32(0), decoded.Get(300))
}

// TestPalettedContainerRejectsBadPaletteLength feeds crafted palette lengths
// that would otherwise panic or allocate gigabytes.
func TestPalettedContainerRejectsBadPaletteLength(t *testing.T) {
	for _, length := range []ns.VarInt{-1, 0, 17, 1 << 30} {
		w := ns.NewWriter()
		require.NoError(t, w.WriteUint8(4))
		require.NoError(t, w.WriteVarInt(length))
		c := chunks.NewSingleValue(chunks.BlockStatesKind, 0)
		assert.Error(t, c.Decode(ns.NewReader(w.Bytes())), "palette length %d", length)
	}
}
