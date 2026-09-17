package storage

import (
	"fmt"
	"math/bits"
	"sync"

	"github.com/go-mclib/data/blocks"
	"github.com/go-mclib/data/chunks"
	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// DataVersion is the 26.3 world data version written to saved data.
const DataVersion = 5023

// ChunkToNBT serializes a ChunkColumn to Anvil-format NBT bytes (uncompressed).
func ChunkToNBT(col *chunks.ChunkColumn) ([]byte, error) {
	root := nbt.Compound{
		"DataVersion": nbt.Int(DataVersion),
		"xPos":        nbt.Int(col.X),
		"yPos":        nbt.Int(col.MinY >> 4),
		"zPos":        nbt.Int(col.Z),
		"Status":      nbt.String("minecraft:full"),
		"LastUpdate":  nbt.Long(0),
	}

	// sections
	var sectionTags []nbt.Tag
	if tag := lightSectionToNBT(col, int8((col.MinY>>4)-1), 0); tag != nil {
		sectionTags = append(sectionTags, tag)
	}
	for i := range len(col.Sections) {
		sec := col.Sections[i]
		if sec == nil {
			continue
		}
		sectionY := int8(i + (col.MinY >> 4))
		tag, err := sectionToNBT(sec, sectionY, col, i)
		if err != nil {
			return nil, fmt.Errorf("section %d: %w", i, err)
		}
		if tag != nil {
			sectionTags = append(sectionTags, tag)
		}
	}
	if tag := lightSectionToNBT(col, int8((col.MinY>>4)+len(col.Sections)), len(col.Sections)+1); tag != nil {
		sectionTags = append(sectionTags, tag)
	}
	root["sections"] = nbt.List{ElementType: nbt.TagCompound, Elements: sectionTags}

	// heightmaps (pass through)
	if len(col.Heightmaps) > 0 {
		hm := nbt.Compound{}
		for typeID, longs := range col.Heightmaps {
			name := heightmapTypeName(typeID)
			if name == "" {
				continue
			}
			la := make(nbt.LongArray, len(longs))
			copy(la, longs)
			hm[name] = la
		}
		root["Heightmaps"] = hm
	}

	// light
	if col.Light != nil {
		root["isLightOn"] = nbt.Byte(1)
	}

	// empty required fields
	root["block_entities"] = nbt.List{ElementType: nbt.TagCompound}
	root["block_ticks"] = nbt.List{ElementType: nbt.TagCompound}
	root["fluid_ticks"] = nbt.List{ElementType: nbt.TagCompound}

	return nbt.EncodeFile(root, "")
}

func sectionToNBT(sec *chunks.ChunkSection, sectionY int8, col *chunks.ChunkColumn, sectionIdx int) (nbt.Compound, error) {
	tag := nbt.Compound{
		"Y": nbt.Byte(sectionY),
	}

	// block states
	bsTag, err := blockStatesToNBT(sec)
	if err != nil {
		return nil, err
	}
	tag["block_states"] = bsTag

	// biomes
	tag["biomes"] = biomesToNBT(sec)

	// light data from the column
	if col.Light != nil {
		lightIdx := sectionIdx + 1 // light sections are offset by 1
		if skyArr := getLightArray(col.Light.SkyLightMask, col.Light.SkyLightArrays, lightIdx); skyArr != nil {
			tag["SkyLight"] = nbt.ByteArray(skyArr)
		}
		if blockArr := getLightArray(col.Light.BlockLightMask, col.Light.BlockLightArrays, lightIdx); blockArr != nil {
			tag["BlockLight"] = nbt.ByteArray(blockArr)
		}
	}

	return tag, nil
}

func lightSectionToNBT(col *chunks.ChunkColumn, sectionY int8, lightIdx int) nbt.Compound {
	if col.Light == nil {
		return nil
	}
	tag := nbt.Compound{"Y": nbt.Byte(sectionY)}
	if skyArr := getLightArray(col.Light.SkyLightMask, col.Light.SkyLightArrays, lightIdx); skyArr != nil {
		tag["SkyLight"] = nbt.ByteArray(skyArr)
	}
	if blockArr := getLightArray(col.Light.BlockLightMask, col.Light.BlockLightArrays, lightIdx); blockArr != nil {
		tag["BlockLight"] = nbt.ByteArray(blockArr)
	}
	if len(tag) == 1 {
		return nil
	}
	return tag
}

func getLightArray(mask ns.BitSet, arrays [][]byte, sectionIdx int) []byte {
	if !mask.Get(sectionIdx) {
		return nil
	}
	// count how many bits are set before this index to find the array position
	arrIdx := 0
	for i := range sectionIdx {
		if mask.Get(i) {
			arrIdx++
		}
	}
	if arrIdx >= len(arrays) {
		return nil
	}
	return arrays[arrIdx]
}

func blockStatesToNBT(sec *chunks.ChunkSection) (nbt.Compound, error) {
	// collect unique states and build palette
	type paletteEntry struct {
		name  string
		props map[string]string
	}
	stateToIdx := make(map[int32]int32) // stateID -> palette index
	var palette []paletteEntry

	for i := range 4096 {
		stateID := sec.BlockStates.Get(i)
		if _, ok := stateToIdx[stateID]; ok {
			continue
		}
		idx := int32(len(palette))
		stateToIdx[stateID] = idx

		if stateID == 0 {
			palette = append(palette, paletteEntry{name: "minecraft:air"})
		} else {
			blockID, props := blocks.StateProperties(int(stateID))
			name := blocks.BlockName(blockID)
			if name == "" {
				name = "minecraft:air"
			}
			palette = append(palette, paletteEntry{name: name, props: props})
		}
	}

	// build NBT palette
	paletteTags := make([]nbt.Tag, len(palette))
	for i, entry := range palette {
		comp := nbt.Compound{"id": nbt.String(entry.name)}
		if len(entry.props) > 0 {
			propsComp := nbt.Compound{}
			for k, v := range entry.props {
				propsComp[k] = nbt.String(v)
			}
			comp["properties"] = propsComp
		}
		paletteTags[i] = comp
	}

	result := nbt.Compound{
		"palette": nbt.List{ElementType: nbt.TagCompound, Elements: paletteTags},
	}

	// pack data array (only if palette size > 1)
	if len(palette) > 1 {
		bpe := max(4, bits.Len(uint(len(palette)-1)))
		result["data"] = packLongArray(sec.BlockStates, stateToIdx, 4096, bpe)
	}

	return result, nil
}

func biomesToNBT(sec *chunks.ChunkSection) nbt.Compound {
	biomeNames := registries.SynchronizedEntries["minecraft:worldgen/biome"]

	// collect unique biome IDs
	biomeToIdx := make(map[int32]int32)
	type biomeEntry struct {
		id   int32
		name string
	}
	var palette []biomeEntry

	for i := range 64 {
		biomeID := sec.Biomes.Get(i)
		if _, ok := biomeToIdx[biomeID]; ok {
			continue
		}
		idx := int32(len(palette))
		biomeToIdx[biomeID] = idx

		name := "minecraft:plains"
		if int(biomeID) < len(biomeNames) {
			name = biomeNames[biomeID]
		}
		palette = append(palette, biomeEntry{id: biomeID, name: name})
	}

	paletteTags := make([]nbt.Tag, len(palette))
	for i, entry := range palette {
		paletteTags[i] = nbt.String(entry.name)
	}

	result := nbt.Compound{
		"palette": nbt.List{ElementType: nbt.TagString, Elements: paletteTags},
	}

	if len(palette) > 1 {
		bpe := max(1, bits.Len(uint(len(palette)-1)))
		result["data"] = packLongArray(sec.Biomes, biomeToIdx, 64, bpe)
	}

	return result
}

// packLongArray packs container entries into an nbt.LongArray using SimpleBitStorage format.
func packLongArray(container *chunks.PalettedContainer, idMap map[int32]int32, entryCount, bpe int) nbt.LongArray {
	valuesPerLong := 64 / bpe
	longCount := (entryCount + valuesPerLong - 1) / valuesPerLong
	longs := make(nbt.LongArray, longCount)

	for i := range entryCount {
		raw := container.Get(i)
		paletteIdx := idMap[raw]
		longIdx := i / valuesPerLong
		bitIdx := (i % valuesPerLong) * bpe
		longs[longIdx] |= int64(paletteIdx) << bitIdx
	}

	return longs
}

// NBTToChunk deserializes an overworld-height ChunkColumn from Anvil-format NBT bytes.
func NBTToChunk(data []byte) (*chunks.ChunkColumn, error) {
	return NBTToChunkWithHeight(data, chunks.MinY, chunks.SectionCount)
}

// NBTToChunkWithHeight deserializes a ChunkColumn for a dimension height.
func NBTToChunkWithHeight(data []byte, minY, sectionCount int) (*chunks.ChunkColumn, error) {
	if minY%16 != 0 || sectionCount <= 0 {
		return nil, fmt.Errorf("invalid chunk height minY=%d sections=%d", minY, sectionCount)
	}
	tag, _, err := nbt.DecodeFile(data, nbt.WithMaxBytes(0))
	if err != nil {
		return nil, fmt.Errorf("decode NBT: %w", err)
	}
	root, ok := tag.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("root is not compound")
	}

	col := &chunks.ChunkColumn{
		X:          root.GetInt("xPos"),
		Z:          root.GetInt("zPos"),
		MinY:       minY,
		Sections:   make([]*chunks.ChunkSection, sectionCount),
		Heightmaps: make(map[int32][]int64),
	}

	// heightmaps
	if hmComp := root.GetCompound("Heightmaps"); hmComp != nil {
		for name, tag := range hmComp {
			if la, ok := tag.(nbt.LongArray); ok {
				typeID := heightmapTypeID(name)
				if typeID >= 0 {
					col.Heightmaps[typeID] = []int64(la)
				}
			}
		}
	}

	// sections
	for i := range sectionCount {
		col.Sections[i] = chunks.NewEmptySection()
	}

	sectionList := root.GetList("sections")
	var skyLightArrays [][]byte
	var blockLightArrays [][]byte
	skyMask := ns.NewBitSet(sectionCount + 2)
	emptySkyMask := ns.NewBitSet(sectionCount + 2)
	blockMask := ns.NewBitSet(sectionCount + 2)
	emptyBlockMask := ns.NewBitSet(sectionCount + 2)

	for _, elem := range sectionList.Elements {
		secTag, ok := elem.(nbt.Compound)
		if !ok {
			continue
		}
		sectionY := secTag.GetByte("Y")
		secIdx := int(sectionY) - (minY >> 4)
		lightIdx := secIdx + 1
		if lightIdx >= 0 && lightIdx < sectionCount+2 {
			if skyData := secTag.GetByteArray("SkyLight"); len(skyData) == 2048 {
				skyMask.Set(lightIdx)
				skyLightArrays = append(skyLightArrays, skyData)
			} else if secIdx >= 0 && secIdx < sectionCount {
				emptySkyMask.Set(lightIdx)
			}
			if blockData := secTag.GetByteArray("BlockLight"); len(blockData) == 2048 {
				blockMask.Set(lightIdx)
				blockLightArrays = append(blockLightArrays, blockData)
			} else if secIdx >= 0 && secIdx < sectionCount {
				emptyBlockMask.Set(lightIdx)
			}
		}
		if secIdx < 0 || secIdx >= sectionCount {
			continue
		}

		sec := col.Sections[secIdx]

		// block states
		if bsComp := secTag.GetCompound("block_states"); bsComp != nil {
			if err := loadBlockStates(sec, bsComp); err != nil {
				return nil, fmt.Errorf("section Y=%d block_states: %w", sectionY, err)
			}
		}

		// biomes
		if biomeComp := secTag.GetCompound("biomes"); biomeComp != nil {
			loadBiomes(sec, biomeComp)
		}

	}

	// build light data if any arrays were found
	if len(skyLightArrays) > 0 || len(blockLightArrays) > 0 {
		col.Light = &ns.LightData{
			SkyLightMask:        *skyMask,
			BlockLightMask:      *blockMask,
			EmptySkyLightMask:   *emptySkyMask,
			EmptyBlockLightMask: *emptyBlockMask,
			SkyLightArrays:      skyLightArrays,
			BlockLightArrays:    blockLightArrays,
		}
	}

	return col, nil
}

func loadBlockStates(sec *chunks.ChunkSection, comp nbt.Compound) error {
	paletteList := comp.GetList("palette")
	if paletteList.Len() == 0 {
		return nil
	}

	// build state ID palette
	stateIDs := make([]int32, paletteList.Len())
	nonAir := int16(0)
	for i, elem := range paletteList.Elements {
		entry, ok := elem.(nbt.Compound)
		if !ok {
			continue
		}
		name := entry.GetString("id")
		if name == "minecraft:air" || name == "" {
			stateIDs[i] = 0
			continue
		}

		blockID := blocks.BlockID(name)
		if blockID < 0 {
			continue
		}

		propsComp := entry.GetCompound("properties")
		if len(propsComp) > 0 {
			props := make(map[string]string, len(propsComp))
			for k, v := range propsComp {
				if s, ok := v.(nbt.String); ok {
					props[k] = string(s)
				}
			}
			stateIDs[i] = blocks.StateID(int(blockID), props)
		} else {
			stateIDs[i] = blocks.DefaultStateID(blockID)
		}
	}

	if paletteList.Len() == 1 {
		// single-value: all blocks are the same
		sid := stateIDs[0]
		if sid != 0 {
			nonAir = 4096
		}
		for i := range 4096 {
			sec.BlockStates.Set(i, sid)
		}
		sec.BlockCount = nonAir
		return nil
	}

	// unpack data array
	dataArr := comp.GetLongArray("data")
	if dataArr == nil {
		return nil
	}

	bpe := max(4, bits.Len(uint(paletteList.Len()-1)))
	valuesPerLong := 64 / bpe
	mask := int64((1 << bpe) - 1)

	for i := range 4096 {
		longIdx := i / valuesPerLong
		if longIdx >= len(dataArr) {
			break
		}
		bitIdx := (i % valuesPerLong) * bpe
		paletteIdx := int((dataArr[longIdx] >> bitIdx) & mask)
		if paletteIdx >= len(stateIDs) {
			continue
		}
		sid := stateIDs[paletteIdx]
		sec.BlockStates.Set(i, sid)
		if sid != 0 {
			nonAir++
		}
	}

	sec.BlockCount = nonAir
	return nil
}

// biomeNameToID maps biome names to protocol IDs (built lazily).
var (
	biomeNameMap     map[string]int32
	biomeNameMapOnce sync.Once
)

func getBiomeNameMap() map[string]int32 {
	biomeNameMapOnce.Do(func() {
		biomeNameMap = make(map[string]int32)
		for id, name := range registries.SynchronizedEntries["minecraft:worldgen/biome"] {
			biomeNameMap[name] = int32(id)
		}
	})
	return biomeNameMap
}

func loadBiomes(sec *chunks.ChunkSection, comp nbt.Compound) {
	paletteList := comp.GetList("palette")
	if paletteList.Len() == 0 {
		return
	}

	nameMap := getBiomeNameMap()

	// build biome ID palette
	biomeIDs := make([]int32, paletteList.Len())
	for i, elem := range paletteList.Elements {
		if s, ok := elem.(nbt.String); ok {
			biomeIDs[i] = nameMap[string(s)]
		}
	}

	if paletteList.Len() == 1 {
		for i := range 64 {
			sec.Biomes.Set(i, biomeIDs[0])
		}
		return
	}

	dataArr := comp.GetLongArray("data")
	if dataArr == nil {
		return
	}

	bpe := max(1, bits.Len(uint(paletteList.Len()-1)))
	valuesPerLong := 64 / bpe
	mask := int64((1 << bpe) - 1)

	for i := range 64 {
		longIdx := i / valuesPerLong
		if longIdx >= len(dataArr) {
			break
		}
		bitIdx := (i % valuesPerLong) * bpe
		paletteIdx := int((dataArr[longIdx] >> bitIdx) & mask)
		if paletteIdx >= len(biomeIDs) {
			continue
		}
		sec.Biomes.Set(i, biomeIDs[paletteIdx])
	}
}

// heightmap type ID <-> name mapping (matching vanilla)
func heightmapTypeName(id int32) string {
	switch id {
	case 1:
		return "WORLD_SURFACE"
	case 4:
		return "MOTION_BLOCKING"
	case 5:
		return "MOTION_BLOCKING_NO_LEAVES"
	default:
		return ""
	}
}

func heightmapTypeID(name string) int32 {
	switch name {
	case "WORLD_SURFACE":
		return 1
	case "MOTION_BLOCKING":
		return 4
	case "MOTION_BLOCKING_NO_LEAVES":
		return 5
	default:
		return -1
	}
}
