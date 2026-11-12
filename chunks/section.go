package chunks

import (
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// Section represents a 16x16x16 section of blocks and a 4x4x4 grid of biomes.
type Section struct {
	BlockCount int16
	// FluidCount is the number of fluid blocks in the section (added in 26.1).
	FluidCount  int16
	BlockStates *PalettedContainer
	Biomes      *PalettedContainer
}

// NewEmptySection creates a section with all air blocks and a single-value biome (ID 0).
func NewEmptySection() *Section {
	return &Section{
		BlockCount:  0,
		FluidCount:  0,
		BlockStates: newSingleValue(blockStatesKind, 0),
		Biomes:      newSingleValue(biomesKind, 0),
	}
}

// GetBlockState returns the block state ID at local coordinates (0-15 each).
func (s *Section) GetBlockState(x, y, z int) int32 {
	return s.BlockStates.getXYZ(x, y, z)
}

// SetBlockState sets the block state ID at local coordinates (0-15 each).
func (s *Section) SetBlockState(x, y, z int, stateID int32) {
	s.BlockStates.setXYZ(x, y, z, stateID)
}

// GetBiome returns the biome ID at local biome coordinates (0-3 each).
func (s *Section) GetBiome(x, y, z int) int32 {
	return s.Biomes.getXYZ(x, y, z)
}

// SetBiome sets the biome ID at local biome coordinates (0-3 each).
func (s *Section) SetBiome(x, y, z int, biomeID int32) {
	s.Biomes.setXYZ(x, y, z, biomeID)
}

// Decode reads a Section from the buffer.
func (s *Section) Decode(buf *ns.PacketBuffer) error {
	bc, err := buf.ReadInt16()
	if err != nil {
		return err
	}
	s.BlockCount = int16(bc)

	fc, err := buf.ReadInt16()
	if err != nil {
		return err
	}
	s.FluidCount = int16(fc)

	s.BlockStates = &PalettedContainer{kind: blockStatesKind}
	if err := s.BlockStates.Decode(buf); err != nil {
		return err
	}

	s.Biomes = &PalettedContainer{kind: biomesKind}
	return s.Biomes.Decode(buf)
}

// Encode writes a Section to the buffer.
func (s *Section) Encode(buf *ns.PacketBuffer) error {
	if err := buf.WriteInt16(ns.Int16(s.BlockCount)); err != nil {
		return err
	}
	if err := buf.WriteInt16(ns.Int16(s.FluidCount)); err != nil {
		return err
	}
	if err := s.BlockStates.Encode(buf); err != nil {
		return err
	}
	return s.Biomes.Encode(buf)
}
