package chunks

func (p *PalettedContainer) GetXYZ(x, y, z int) int32 { return p.getXYZ(x, y, z) }

func (p *PalettedContainer) SetXYZ(x, y, z int, value int32) { p.setXYZ(x, y, z, value) }

var (
	NewSingleValue  = newSingleValue
	BlockStatesKind = blockStatesKind
	BiomesKind      = biomesKind
)
