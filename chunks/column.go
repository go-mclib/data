package chunks

import (
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

const (
	// SectionCount is the vanilla overworld section count (Y -64..319). Used as
	// a fallback when the dimension's real height is unknown.
	SectionCount = 24
	// MinY is the vanilla overworld minimum world Y (fallback default).
	MinY = -64
	// MaxY is the vanilla overworld maximum world Y, exclusive (fallback default).
	MaxY = MinY + SectionCount*16
)

// Column represents a vertical column of chunk sections with associated
// data. Height is per-column: MinY and len(Sections) come from the dimension
// type, so datapack-changed world heights parse correctly.
type Column struct {
	X, Z          int32
	MinY          int // world Y of the bottom section
	Sections      []*Section
	Heightmaps    map[int32][]int64
	BlockEntities []ns.BlockEntity
	Light         *ns.LightData
}

// NewColumn returns an empty column at the vanilla overworld height with its
// section slice allocated (all nil). Use for tests and chunk generation.
func NewColumn(x, z int32) *Column {
	return &Column{
		X:          x,
		Z:          z,
		MinY:       MinY,
		Sections:   make([]*Section, SectionCount),
		Heightmaps: make(map[int32][]int64),
	}
}

// ParseColumn decodes a Column. minY and sectionCount describe the
// dimension's vertical extent (from the dimension type); pass chunks.MinY and
// chunks.SectionCount for the vanilla overworld.
func ParseColumn(x, z int32, minY, sectionCount int, data ns.ChunkData, light *ns.LightData) (*Column, error) {
	col := &Column{
		X:             x,
		Z:             z,
		MinY:          minY,
		Sections:      make([]*Section, sectionCount),
		Heightmaps:    data.Heightmaps,
		BlockEntities: data.BlockEntities,
		Light:         light,
	}

	buf := ns.NewReader(data.Data)
	for i := range sectionCount {
		sec := &Section{}
		if err := sec.Decode(buf); err != nil {
			return nil, err
		}
		col.Sections[i] = sec
	}
	return col, nil
}

// GetBlockState returns the block state ID at absolute world coordinates.
// Returns 0 (air) if the coordinates are out of range or the section is nil.
func (c *Column) GetBlockState(x, y, z int) int32 {
	idx := c.SectionIndex(y)
	if idx < 0 {
		return 0
	}
	sec := c.Sections[idx]
	if sec == nil {
		return 0
	}
	lx, ly, lz := LocalCoords(x, y, z)
	return sec.GetBlockState(lx, ly, lz)
}

// SetBlockState sets the block state ID at absolute world coordinates.
// Creates an empty section if the target section is nil.
func (c *Column) SetBlockState(x, y, z int, stateID int32) {
	idx := c.SectionIndex(y)
	if idx < 0 {
		return
	}
	if c.Sections[idx] == nil {
		c.Sections[idx] = NewEmptySection()
	}
	lx, ly, lz := LocalCoords(x, y, z)
	c.Sections[idx].SetBlockState(lx, ly, lz, stateID)
}

// EncodeSections encodes all chunk sections back to raw bytes (the ChunkData.Data portion).
func (c *Column) EncodeSections() ([]byte, error) {
	buf := ns.NewWriter()
	for _, sec := range c.Sections {
		if sec == nil {
			sec = NewEmptySection()
		}
		if err := sec.Encode(buf); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// SectionIndex returns the section array index for the given world Y, or -1 if
// Y is outside this column's vertical range.
func (c *Column) SectionIndex(worldY int) int {
	idx := (worldY - c.MinY) >> 4
	if idx < 0 || idx >= len(c.Sections) {
		return -1
	}
	return idx
}

// SectionIndex returns the section index for a world Y assuming the vanilla
// overworld range. Prefer (*Column).SectionIndex for dimension-correct
// bounds; this exists for callers without a column.
func SectionIndex(worldY int) int {
	idx := (worldY - MinY) >> 4
	if idx < 0 || idx >= SectionCount {
		return -1
	}
	return idx
}

// LocalCoords extracts the local 0-15 coordinates from world coordinates.
func LocalCoords(worldX, worldY, worldZ int) (int, int, int) {
	return worldX & 0xF, worldY & 0xF, worldZ & 0xF
}

// ChunkPos returns the chunk coordinates for the given world X, Z.
func ChunkPos(worldX, worldZ int) (int32, int32) {
	return int32(worldX >> 4), int32(worldZ >> 4)
}
