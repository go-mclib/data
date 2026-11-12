package chunks

import (
	"github.com/go-mclib/data/blocks"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// ComputeSkylight computes sky and emitted block light for the column.
func (c *Column) ComputeSkylight() {
	sectionCount := len(c.Sections)
	lightSectionCount := sectionCount + 2
	height := sectionCount * 16
	size := height * 16 * 16
	skyValues := make([]byte, size)
	blockValues := make([]byte, size)
	opacity := make([]byte, size)

	for x := range 16 {
		for z := range 16 {
			sky := byte(15)
			for y := height - 1; y >= 0; y-- {
				index := lightIndex(x, y, z)
				state := c.GetBlockState(x, c.MinY+y, z)
				dampening := blocks.LightDampening(state)
				opacity[index] = dampening
				sky -= min(sky, dampening)
				skyValues[index] = sky
				blockValues[index] = blocks.LightEmission(state)
			}
		}
	}

	propagateLight(skyValues, opacity, height)
	propagateLight(blockValues, opacity, height)

	skyLight := make([][]byte, lightSectionCount)
	blockLight := make([][]byte, lightSectionCount)
	skyLight[lightSectionCount-1] = fullSkyLightArray()
	for y := range height {
		section := y/16 + 1
		if skyLight[section] == nil {
			skyLight[section] = make([]byte, 2048)
			blockLight[section] = make([]byte, 2048)
		}
		for z := range 16 {
			for x := range 16 {
				index := lightIndex(x, y, z)
				setLightNibble(skyLight[section], x, y&15, z, skyValues[index])
				setLightNibble(blockLight[section], x, y&15, z, blockValues[index])
			}
		}
	}

	skyMask, emptySkyMask, skyArrays := packLight(skyLight)
	blockMask, emptyBlockMask, blockArrays := packLight(blockLight)
	c.Light = &ns.LightData{
		SkyLightMask:        *skyMask,
		BlockLightMask:      *blockMask,
		EmptySkyLightMask:   *emptySkyMask,
		EmptyBlockLightMask: *emptyBlockMask,
		SkyLightArrays:      skyArrays,
		BlockLightArrays:    blockArrays,
	}
}

func propagateLight(values, opacity []byte, height int) {
	queue := make([]int, 0, len(values))
	for index, value := range values {
		if value > 1 {
			queue = append(queue, index)
		}
	}
	for head := 0; head < len(queue); head++ {
		index := queue[head]
		value := values[index]
		x := index & 15
		z := index / 16 & 15
		y := index / 256
		for _, offset := range [...]struct{ x, y, z int }{
			{x: 1}, {x: -1}, {y: 1}, {y: -1}, {z: 1}, {z: -1},
		} {
			nx, ny, nz := x+offset.x, y+offset.y, z+offset.z
			if nx < 0 || nx >= 16 || ny < 0 || ny >= height || nz < 0 || nz >= 16 {
				continue
			}
			next := lightIndex(nx, ny, nz)
			candidate := value - min(value, max(byte(1), opacity[next]))
			if candidate > values[next] {
				values[next] = candidate
				if candidate > 1 {
					queue = append(queue, next)
				}
			}
		}
	}
}

func lightIndex(x, y, z int) int {
	return y*256 + z*16 + x
}

func packLight(light [][]byte) (*ns.BitSet, *ns.BitSet, [][]byte) {
	mask := ns.NewBitSet(len(light))
	emptyMask := ns.NewBitSet(len(light))
	var arrays [][]byte
	for i, values := range light {
		if values != nil && !isAllZero(values) {
			mask.Set(i)
			arrays = append(arrays, values)
		} else {
			emptyMask.Set(i)
		}
	}
	return mask, emptyMask, arrays
}

func fullSkyLightArray() []byte {
	arr := make([]byte, 2048)
	for i := range arr {
		arr[i] = 0xFF
	}
	return arr
}

func setLightNibble(arr []byte, x, y, z int, value byte) {
	idx := (y*16+z)*16 + x
	byteIdx := idx / 2
	if idx%2 == 0 {
		arr[byteIdx] = (arr[byteIdx] & 0xF0) | (value & 0x0F)
	} else {
		arr[byteIdx] = (arr[byteIdx] & 0x0F) | ((value & 0x0F) << 4)
	}
}

func isAllZero(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}
