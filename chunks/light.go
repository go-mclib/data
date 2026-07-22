package chunks

import (
	"strconv"
	"strings"

	"github.com/go-mclib/data/block_shapes"
	"github.com/go-mclib/data/blocks"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// ComputeSkylight computes sky and emitted block light for the column.
func (c *ChunkColumn) ComputeSkylight() {
	sectionCount := len(c.Sections)
	lightSectionCount := sectionCount + 2
	height := sectionCount * 16
	size := height * 16 * 16
	skyValues := make([]byte, size)
	blockValues := make([]byte, size)
	opacity := make([]byte, size)
	cache := make(map[int32]lightTraits)

	for x := range 16 {
		for z := range 16 {
			sky := byte(15)
			for y := height - 1; y >= 0; y-- {
				index := lightIndex(x, y, z)
				state := c.GetBlockState(x, c.MinY+y, z)
				traits := traitsForState(state, cache)
				opacity[index] = traits.opacity
				if traits.opacity >= 15 {
					sky = 0
				} else if traits.opacity > 0 {
					sky -= min(sky, traits.opacity)
				}
				skyValues[index] = sky
				blockValues[index] = traits.emission
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

type lightTraits struct {
	opacity  byte
	emission byte
}

func traitsForState(state int32, cache map[int32]lightTraits) lightTraits {
	if state == 0 {
		return lightTraits{}
	}
	if traits, ok := cache[state]; ok {
		return traits
	}
	id, properties := blocks.StateProperties(int(state))
	name := blocks.BlockName(id)
	traits := lightTraits{
		opacity:  blockOpacity(state, name),
		emission: blockEmission(name, properties),
	}
	cache[state] = traits
	return traits
}

func blockOpacity(state int32, name string) byte {
	switch {
	case name == "minecraft:water":
		return 1
	case strings.HasSuffix(name, "_leaves"):
		return 1
	case name == "minecraft:tinted_glass":
		return 15
	case strings.Contains(name, "glass"), strings.Contains(name, "ice"):
		return 0
	case name == "minecraft:slime_block", name == "minecraft:honey_block":
		return 0
	case block_shapes.IsFullBlock(state):
		return 15
	default:
		return 0
	}
}

func blockEmission(name string, properties map[string]string) byte {
	lit := properties["lit"] == "true"
	switch {
	case name == "minecraft:lava", name == "minecraft:fire":
		return 15
	case name == "minecraft:soul_fire":
		return 10
	case name == "minecraft:torch", name == "minecraft:wall_torch",
		name == "minecraft:copper_torch", name == "minecraft:copper_wall_torch":
		return 14
	case name == "minecraft:soul_torch", name == "minecraft:soul_wall_torch":
		return 10
	case name == "minecraft:redstone_torch" || name == "minecraft:redstone_wall_torch":
		if lit {
			return 7
		}
	case name == "minecraft:redstone_ore" || name == "minecraft:deepslate_redstone_ore":
		if lit {
			return 9
		}
	case name == "minecraft:furnace" || name == "minecraft:smoker" || name == "minecraft:blast_furnace":
		if lit {
			return 13
		}
	case name == "minecraft:campfire":
		if lit {
			return 15
		}
	case name == "minecraft:soul_campfire":
		if lit {
			return 10
		}
	case name == "minecraft:redstone_lamp" || name == "minecraft:jack_o_lantern":
		if name == "minecraft:jack_o_lantern" || lit {
			return 15
		}
	case name == "minecraft:glowstone", name == "minecraft:sea_lantern",
		name == "minecraft:beacon", name == "minecraft:conduit",
		name == "minecraft:shroomlight", name == "minecraft:ochre_froglight",
		name == "minecraft:verdant_froglight", name == "minecraft:pearlescent_froglight",
		name == "minecraft:lantern",
		name == "minecraft:lava_cauldron", name == "minecraft:end_portal",
		name == "minecraft:end_gateway":
		return 15
	case name == "minecraft:soul_lantern":
		return 10
	case name == "minecraft:end_rod":
		return 14
	case name == "minecraft:nether_portal":
		return 11
	case name == "minecraft:crying_obsidian":
		return 10
	case name == "minecraft:magma_block":
		return 3
	case name == "minecraft:brewing_stand", name == "minecraft:dragon_egg",
		name == "minecraft:end_portal_frame", name == "minecraft:brown_mushroom":
		return 1
	case name == "minecraft:firefly_bush":
		return 2
	case name == "minecraft:enchanting_table":
		return 7
	case name == "minecraft:ender_chest":
		return 7
	case name == "minecraft:sculk_catalyst":
		return 6
	case name == "minecraft:sculk_sensor", name == "minecraft:calibrated_sculk_sensor":
		return 1
	case name == "minecraft:amethyst_cluster":
		return 5
	case name == "minecraft:large_amethyst_bud":
		return 4
	case name == "minecraft:medium_amethyst_bud":
		return 2
	case name == "minecraft:small_amethyst_bud":
		return 1
	case strings.HasSuffix(name, "copper_lantern"):
		return 15
	case name == "minecraft:trial_spawner":
		switch properties["trial_spawner_state"] {
		case "waiting_for_players":
			return 4
		case "active", "waiting_for_reward_ejection", "ejecting_reward":
			return 8
		}
	case name == "minecraft:vault":
		if properties["vault_state"] == "inactive" {
			return 6
		}
		return 12
	case name == "minecraft:respawn_anchor":
		charges, _ := strconv.Atoi(properties["charges"])
		return byte(charges * 15 / 4)
	case name == "minecraft:light":
		level, _ := strconv.Atoi(properties["level"])
		return byte(level)
	case name == "minecraft:sea_pickle":
		if properties["waterlogged"] == "true" {
			pickles, _ := strconv.Atoi(properties["pickles"])
			return byte(3 + 3*pickles)
		}
	case name == "minecraft:glow_lichen":
		for _, face := range [...]string{"up", "down", "north", "south", "east", "west"} {
			if properties[face] == "true" {
				return 7
			}
		}
	case name == "minecraft:cave_vines" || name == "minecraft:cave_vines_plant":
		if properties["berries"] == "true" {
			return 14
		}
	case strings.HasSuffix(name, "_candle") || strings.HasSuffix(name, "_candle_cake"):
		if lit {
			if strings.HasSuffix(name, "_candle_cake") {
				return 3
			}
			candles, _ := strconv.Atoi(properties["candles"])
			return byte(candles * 3)
		}
	case strings.Contains(name, "copper_bulb"):
		if lit {
			switch {
			case strings.Contains(name, "oxidized"):
				return 4
			case strings.Contains(name, "weathered"):
				return 8
			case strings.Contains(name, "exposed"):
				return 12
			default:
				return 15
			}
		}
	}
	return 0
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
