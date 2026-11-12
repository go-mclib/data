// Code generated for Minecraft 26.3 (Protocol 777); DO NOT EDIT.

package blocks

var blockStates = map[int32]*blockStateData{
	499: { // minecraft:acacia_button
		BaseID:    12514,
		DefaultID: 12523,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	705: { // minecraft:acacia_door
		BaseID:    16163,
		DefaultID: 16174,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	695: { // minecraft:acacia_fence
		BaseID:    15747,
		DefaultID: 15778,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	685: { // minecraft:acacia_fence_gate
		BaseID:    15427,
		DefaultID: 15434,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	283: { // minecraft:acacia_hanging_sign
		BaseID:    7695,
		DefaultID: 7744,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	98: { // minecraft:acacia_leaves
		BaseID:    379,
		DefaultID: 406,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	55: { // minecraft:acacia_log
		BaseID:    151,
		DefaultID: 152,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	17: { // minecraft:acacia_planks
		BaseID:    19,
		DefaultID: 19,
	},
	313: { // minecraft:acacia_pressure_plate
		BaseID:    8537,
		DefaultID: 8538,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	30: { // minecraft:acacia_sapling
		BaseID:    38,
		DefaultID: 38,
		Properties: []blockProperty{
			{"stage", []string{"0", "1"}},
		},
	},
	223: { // minecraft:acacia_shelf
		BaseID:    4092,
		DefaultID: 4101,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	257: { // minecraft:acacia_sign
		BaseID:    6987,
		DefaultID: 7004,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	657: { // minecraft:acacia_slab
		BaseID:    15201,
		DefaultID: 15204,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	569: { // minecraft:acacia_stairs
		BaseID:    13739,
		DefaultID: 13750,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	369: { // minecraft:acacia_trapdoor
		BaseID:    9040,
		DefaultID: 9055,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	296: { // minecraft:acacia_wall_hanging_sign
		BaseID:    8359,
		DefaultID: 8360,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	272: { // minecraft:acacia_wall_sign
		BaseID:    7439,
		DefaultID: 7440,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	79: { // minecraft:acacia_wood
		BaseID:    222,
		DefaultID: 223,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	535: { // minecraft:activator_rail
		BaseID:    13175,
		DefaultID: 13188,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"shape", []string{"north_south", "east_west", "ascending_east", "ascending_west", "ascending_north", "ascending_south"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	0: { // minecraft:air
		BaseID:    0,
		DefaultID: 0,
	},
	205: { // minecraft:allium
		BaseID:    3818,
		DefaultID: 3818,
	},
	1068: { // minecraft:amethyst_block
		BaseID:    26759,
		DefaultID: 26759,
	},
	1070: { // minecraft:amethyst_cluster
		BaseID:    26761,
		DefaultID: 26770,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1006: { // minecraft:ancient_debris
		BaseID:    25176,
		DefaultID: 25176,
	},
	6: { // minecraft:andesite
		BaseID:    6,
		DefaultID: 6,
	},
	910: { // minecraft:andesite_slab
		BaseID:    19827,
		DefaultID: 19830,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	897: { // minecraft:andesite_stairs
		BaseID:    19453,
		DefaultID: 19464,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	922: { // minecraft:andesite_wall
		BaseID:    22443,
		DefaultID: 22446,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	520: { // minecraft:anvil
		BaseID:    12962,
		DefaultID: 12962,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	413: { // minecraft:attached_melon_stem
		BaseID:    10072,
		DefaultID: 10072,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	412: { // minecraft:attached_pumpkin_stem
		BaseID:    10068,
		DefaultID: 10068,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	1228: { // minecraft:azalea
		BaseID:    33661,
		DefaultID: 33661,
	},
	106: { // minecraft:azalea_leaves
		BaseID:    603,
		DefaultID: 630,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	206: { // minecraft:azure_bluet
		BaseID:    3819,
		DefaultID: 3819,
	},
	882: { // minecraft:bamboo
		BaseID:    18636,
		DefaultID: 18636,
		Properties: []blockProperty{
			{"age", []string{"0", "1"}},
			{"leaves", []string{"none", "small", "large"}},
			{"stage", []string{"0", "1"}},
		},
	},
	63: { // minecraft:bamboo_block
		BaseID:    174,
		DefaultID: 175,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	505: { // minecraft:bamboo_button
		BaseID:    12658,
		DefaultID: 12667,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	711: { // minecraft:bamboo_door
		BaseID:    16547,
		DefaultID: 16558,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	701: { // minecraft:bamboo_fence
		BaseID:    15939,
		DefaultID: 15970,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	691: { // minecraft:bamboo_fence_gate
		BaseID:    15619,
		DefaultID: 15626,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	292: { // minecraft:bamboo_hanging_sign
		BaseID:    8271,
		DefaultID: 8320,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	25: { // minecraft:bamboo_mosaic
		BaseID:    29,
		DefaultID: 29,
	},
	664: { // minecraft:bamboo_mosaic_slab
		BaseID:    15243,
		DefaultID: 15246,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	576: { // minecraft:bamboo_mosaic_stairs
		BaseID:    14299,
		DefaultID: 14310,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	24: { // minecraft:bamboo_planks
		BaseID:    28,
		DefaultID: 28,
	},
	319: { // minecraft:bamboo_pressure_plate
		BaseID:    8549,
		DefaultID: 8550,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	881: { // minecraft:bamboo_sapling
		BaseID:    18635,
		DefaultID: 18635,
	},
	224: { // minecraft:bamboo_shelf
		BaseID:    4156,
		DefaultID: 4165,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	264: { // minecraft:bamboo_sign
		BaseID:    7211,
		DefaultID: 7228,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	663: { // minecraft:bamboo_slab
		BaseID:    15237,
		DefaultID: 15240,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	575: { // minecraft:bamboo_stairs
		BaseID:    14219,
		DefaultID: 14230,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	375: { // minecraft:bamboo_trapdoor
		BaseID:    9424,
		DefaultID: 9439,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	305: { // minecraft:bamboo_wall_hanging_sign
		BaseID:    8431,
		DefaultID: 8432,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	279: { // minecraft:bamboo_wall_sign
		BaseID:    7495,
		DefaultID: 7496,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	929: { // minecraft:barrel
		BaseID:    24099,
		DefaultID: 24100,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"open", []string{"true", "false"}},
		},
	},
	578: { // minecraft:barrier
		BaseID:    14380,
		DefaultID: 14381,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	337: { // minecraft:basalt
		BaseID:    8670,
		DefaultID: 8671,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	459: { // minecraft:beacon
		BaseID:    11722,
		DefaultID: 11722,
	},
	36: { // minecraft:bedrock
		BaseID:    88,
		DefaultID: 88,
	},
	1001: { // minecraft:bee_nest
		BaseID:    25125,
		DefaultID: 25125,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"honey_level", []string{"0", "1", "2", "3", "4", "5"}},
		},
	},
	1002: { // minecraft:beehive
		BaseID:    25149,
		DefaultID: 25149,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"honey_level", []string{"0", "1", "2", "3", "4", "5"}},
		},
	},
	723: { // minecraft:beetroots
		BaseID:    16792,
		DefaultID: 16792,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3"}},
		},
	},
	938: { // minecraft:bell
		BaseID:    24162,
		DefaultID: 24163,
		Properties: []blockProperty{
			{"attachment", []string{"floor", "ceiling", "single_wall", "double_wall"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1235: { // minecraft:big_dripleaf
		BaseID:    33713,
		DefaultID: 33714,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"tilt", []string{"none", "unstable", "partial", "full"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1236: { // minecraft:big_dripleaf_stem
		BaseID:    33745,
		DefaultID: 33746,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	497: { // minecraft:birch_button
		BaseID:    12466,
		DefaultID: 12475,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	703: { // minecraft:birch_door
		BaseID:    16035,
		DefaultID: 16046,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	693: { // minecraft:birch_fence
		BaseID:    15683,
		DefaultID: 15714,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	683: { // minecraft:birch_fence_gate
		BaseID:    15363,
		DefaultID: 15370,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	282: { // minecraft:birch_hanging_sign
		BaseID:    7631,
		DefaultID: 7680,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	96: { // minecraft:birch_leaves
		BaseID:    323,
		DefaultID: 350,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	53: { // minecraft:birch_log
		BaseID:    145,
		DefaultID: 146,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	15: { // minecraft:birch_planks
		BaseID:    17,
		DefaultID: 17,
	},
	311: { // minecraft:birch_pressure_plate
		BaseID:    8533,
		DefaultID: 8534,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	28: { // minecraft:birch_sapling
		BaseID:    34,
		DefaultID: 34,
		Properties: []blockProperty{
			{"stage", []string{"0", "1"}},
		},
	},
	225: { // minecraft:birch_shelf
		BaseID:    4220,
		DefaultID: 4229,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	256: { // minecraft:birch_sign
		BaseID:    6955,
		DefaultID: 6972,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	655: { // minecraft:birch_slab
		BaseID:    15189,
		DefaultID: 15192,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	456: { // minecraft:birch_stairs
		BaseID:    11550,
		DefaultID: 11561,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	367: { // minecraft:birch_trapdoor
		BaseID:    8912,
		DefaultID: 8927,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	295: { // minecraft:birch_wall_hanging_sign
		BaseID:    8351,
		DefaultID: 8352,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	271: { // minecraft:birch_wall_sign
		BaseID:    7431,
		DefaultID: 7432,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	77: { // minecraft:birch_wood
		BaseID:    216,
		DefaultID: 217,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	632: { // minecraft:black_banner
		BaseID:    15014,
		DefaultID: 15022,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	134: { // minecraft:black_bed
		BaseID:    2270,
		DefaultID: 2273,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1050: { // minecraft:black_candle
		BaseID:    26709,
		DefaultID: 26712,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1067: { // minecraft:black_candle_cake
		BaseID:    26757,
		DefaultID: 26758,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	607: { // minecraft:black_carpet
		BaseID:    14758,
		DefaultID: 14758,
	},
	783: { // minecraft:black_concrete
		BaseID:    17026,
		DefaultID: 17026,
	},
	831: { // minecraft:black_concrete_powder
		BaseID:    18418,
		DefaultID: 18418,
	},
	815: { // minecraft:black_concrete_slab
		BaseID:    18397,
		DefaultID: 18400,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	799: { // minecraft:black_concrete_stairs
		BaseID:    18227,
		DefaultID: 18238,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	767: { // minecraft:black_glazed_terracotta
		BaseID:    17007,
		DefaultID: 17007,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	751: { // minecraft:black_shulker_box
		BaseID:    16941,
		DefaultID: 16945,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	364: { // minecraft:black_stained_glass
		BaseID:    8783,
		DefaultID: 8783,
	},
	568: { // minecraft:black_stained_glass_pane
		BaseID:    13707,
		DefaultID: 13738,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	552: { // minecraft:black_terracotta
		BaseID:    13226,
		DefaultID: 13226,
	},
	648: { // minecraft:black_wall_banner
		BaseID:    15090,
		DefaultID: 15090,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	166: { // minecraft:black_wool
		BaseID:    2424,
		DefaultID: 2424,
	},
	198: { // minecraft:black_wool_slab
		BaseID:    3795,
		DefaultID: 3798,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	182: { // minecraft:black_wool_stairs
		BaseID:    3625,
		DefaultID: 3636,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1014: { // minecraft:blackstone
		BaseID:    25188,
		DefaultID: 25188,
	},
	1017: { // minecraft:blackstone_slab
		BaseID:    25593,
		DefaultID: 25596,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1015: { // minecraft:blackstone_stairs
		BaseID:    25189,
		DefaultID: 25200,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1016: { // minecraft:blackstone_wall
		BaseID:    25269,
		DefaultID: 25272,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	931: { // minecraft:blast_furnace
		BaseID:    24119,
		DefaultID: 24120,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"lit", []string{"true", "false"}},
		},
	},
	628: { // minecraft:blue_banner
		BaseID:    14950,
		DefaultID: 14958,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	130: { // minecraft:blue_bed
		BaseID:    2206,
		DefaultID: 2209,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1046: { // minecraft:blue_candle
		BaseID:    26645,
		DefaultID: 26648,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1063: { // minecraft:blue_candle_cake
		BaseID:    26749,
		DefaultID: 26750,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	603: { // minecraft:blue_carpet
		BaseID:    14754,
		DefaultID: 14754,
	},
	779: { // minecraft:blue_concrete
		BaseID:    17022,
		DefaultID: 17022,
	},
	827: { // minecraft:blue_concrete_powder
		BaseID:    18414,
		DefaultID: 18414,
	},
	811: { // minecraft:blue_concrete_slab
		BaseID:    18373,
		DefaultID: 18376,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	795: { // minecraft:blue_concrete_stairs
		BaseID:    17907,
		DefaultID: 17918,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	763: { // minecraft:blue_glazed_terracotta
		BaseID:    16991,
		DefaultID: 16991,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	879: { // minecraft:blue_ice
		BaseID:    18632,
		DefaultID: 18632,
	},
	204: { // minecraft:blue_orchid
		BaseID:    3817,
		DefaultID: 3817,
	},
	747: { // minecraft:blue_shulker_box
		BaseID:    16917,
		DefaultID: 16921,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	360: { // minecraft:blue_stained_glass
		BaseID:    8779,
		DefaultID: 8779,
	},
	564: { // minecraft:blue_stained_glass_pane
		BaseID:    13579,
		DefaultID: 13610,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	548: { // minecraft:blue_terracotta
		BaseID:    13222,
		DefaultID: 13222,
	},
	644: { // minecraft:blue_wall_banner
		BaseID:    15074,
		DefaultID: 15074,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	162: { // minecraft:blue_wool
		BaseID:    2420,
		DefaultID: 2420,
	},
	194: { // minecraft:blue_wool_slab
		BaseID:    3771,
		DefaultID: 3774,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	178: { // minecraft:blue_wool_stairs
		BaseID:    3305,
		DefaultID: 3316,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	732: { // minecraft:bone_block
		BaseID:    16829,
		DefaultID: 16830,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	221: { // minecraft:bookshelf
		BaseID:    3835,
		DefaultID: 3835,
	},
	854: { // minecraft:brain_coral
		BaseID:    18516,
		DefaultID: 18516,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	844: { // minecraft:brain_coral_block
		BaseID:    18500,
		DefaultID: 18500,
	},
	864: { // minecraft:brain_coral_fan
		BaseID:    18536,
		DefaultID: 18536,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	874: { // minecraft:brain_coral_wall_fan
		BaseID:    18592,
		DefaultID: 18592,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	436: { // minecraft:brewing_stand
		BaseID:    11186,
		DefaultID: 11193,
		Properties: []blockProperty{
			{"has_bottle_0", []string{"true", "false"}},
			{"has_bottle_1", []string{"true", "false"}},
			{"has_bottle_2", []string{"true", "false"}},
		},
	},
	670: { // minecraft:brick_slab
		BaseID:    15279,
		DefaultID: 15282,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	420: { // minecraft:brick_stairs
		BaseID:    10412,
		DefaultID: 10423,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	914: { // minecraft:brick_wall
		BaseID:    19851,
		DefaultID: 19854,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	219: { // minecraft:bricks
		BaseID:    3832,
		DefaultID: 3832,
	},
	629: { // minecraft:brown_banner
		BaseID:    14966,
		DefaultID: 14974,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	131: { // minecraft:brown_bed
		BaseID:    2222,
		DefaultID: 2225,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1047: { // minecraft:brown_candle
		BaseID:    26661,
		DefaultID: 26664,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1064: { // minecraft:brown_candle_cake
		BaseID:    26751,
		DefaultID: 26752,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	604: { // minecraft:brown_carpet
		BaseID:    14755,
		DefaultID: 14755,
	},
	780: { // minecraft:brown_concrete
		BaseID:    17023,
		DefaultID: 17023,
	},
	828: { // minecraft:brown_concrete_powder
		BaseID:    18415,
		DefaultID: 18415,
	},
	812: { // minecraft:brown_concrete_slab
		BaseID:    18379,
		DefaultID: 18382,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	796: { // minecraft:brown_concrete_stairs
		BaseID:    17987,
		DefaultID: 17998,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	764: { // minecraft:brown_glazed_terracotta
		BaseID:    16995,
		DefaultID: 16995,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	215: { // minecraft:brown_mushroom
		BaseID:    3828,
		DefaultID: 3828,
	},
	388: { // minecraft:brown_mushroom_block
		BaseID:    9500,
		DefaultID: 9500,
		Properties: []blockProperty{
			{"down", []string{"true", "false"}},
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"up", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	748: { // minecraft:brown_shulker_box
		BaseID:    16923,
		DefaultID: 16927,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	361: { // minecraft:brown_stained_glass
		BaseID:    8780,
		DefaultID: 8780,
	},
	565: { // minecraft:brown_stained_glass_pane
		BaseID:    13611,
		DefaultID: 13642,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	549: { // minecraft:brown_terracotta
		BaseID:    13223,
		DefaultID: 13223,
	},
	645: { // minecraft:brown_wall_banner
		BaseID:    15078,
		DefaultID: 15078,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	163: { // minecraft:brown_wool
		BaseID:    2421,
		DefaultID: 2421,
	},
	195: { // minecraft:brown_wool_slab
		BaseID:    3777,
		DefaultID: 3780,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	179: { // minecraft:brown_wool_stairs
		BaseID:    3385,
		DefaultID: 3396,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	886: { // minecraft:bubble_column
		BaseID:    18651,
		DefaultID: 18651,
		Properties: []blockProperty{
			{"drag", []string{"true", "false"}},
		},
	},
	855: { // minecraft:bubble_coral
		BaseID:    18518,
		DefaultID: 18518,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	845: { // minecraft:bubble_coral_block
		BaseID:    18501,
		DefaultID: 18501,
	},
	865: { // minecraft:bubble_coral_fan
		BaseID:    18538,
		DefaultID: 18538,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	875: { // minecraft:bubble_coral_wall_fan
		BaseID:    18600,
		DefaultID: 18600,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1069: { // minecraft:budding_amethyst
		BaseID:    26760,
		DefaultID: 26760,
	},
	143: { // minecraft:bush
		BaseID:    2366,
		DefaultID: 2366,
	},
	328: { // minecraft:cactus
		BaseID:    8599,
		DefaultID: 8599,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	329: { // minecraft:cactus_flower
		BaseID:    8615,
		DefaultID: 8615,
	},
	347: { // minecraft:cake
		BaseID:    8697,
		DefaultID: 8697,
		Properties: []blockProperty{
			{"bites", []string{"0", "1", "2", "3", "4", "5", "6"}},
		},
	},
	1115: { // minecraft:calcite
		BaseID:    30517,
		DefaultID: 30517,
	},
	1119: { // minecraft:calibrated_sculk_sensor
		BaseID:    30616,
		DefaultID: 30617,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"power", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"sculk_sensor_phase", []string{"inactive", "active", "cooldown"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	949: { // minecraft:campfire
		BaseID:    24234,
		DefaultID: 24237,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"lit", []string{"true", "false"}},
			{"signal_fire", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1034: { // minecraft:candle
		BaseID:    26453,
		DefaultID: 26456,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1051: { // minecraft:candle_cake
		BaseID:    26725,
		DefaultID: 26726,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	493: { // minecraft:carrots
		BaseID:    12402,
		DefaultID: 12402,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7"}},
		},
	},
	932: { // minecraft:cartography_table
		BaseID:    24127,
		DefaultID: 24127,
	},
	345: { // minecraft:carved_pumpkin
		BaseID:    8689,
		DefaultID: 8689,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	437: { // minecraft:cauldron
		BaseID:    11194,
		DefaultID: 11194,
	},
	885: { // minecraft:cave_air
		BaseID:    18650,
		DefaultID: 18650,
	},
	1225: { // minecraft:cave_vines
		BaseID:    33606,
		DefaultID: 33607,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25"}},
			{"berries", []string{"true", "false"}},
		},
	},
	1226: { // minecraft:cave_vines_plant
		BaseID:    33658,
		DefaultID: 33659,
		Properties: []blockProperty{
			{"berries", []string{"true", "false"}},
		},
	},
	727: { // minecraft:chain_command_block
		BaseID:    16810,
		DefaultID: 16816,
		Properties: []blockProperty{
			{"conditional", []string{"true", "false"}},
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	500: { // minecraft:cherry_button
		BaseID:    12538,
		DefaultID: 12547,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	706: { // minecraft:cherry_door
		BaseID:    16227,
		DefaultID: 16238,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	696: { // minecraft:cherry_fence
		BaseID:    15779,
		DefaultID: 15810,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	686: { // minecraft:cherry_fence_gate
		BaseID:    15459,
		DefaultID: 15466,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	284: { // minecraft:cherry_hanging_sign
		BaseID:    7759,
		DefaultID: 7808,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	99: { // minecraft:cherry_leaves
		BaseID:    407,
		DefaultID: 434,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	56: { // minecraft:cherry_log
		BaseID:    154,
		DefaultID: 155,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	18: { // minecraft:cherry_planks
		BaseID:    20,
		DefaultID: 20,
	},
	314: { // minecraft:cherry_pressure_plate
		BaseID:    8539,
		DefaultID: 8540,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	31: { // minecraft:cherry_sapling
		BaseID:    40,
		DefaultID: 40,
		Properties: []blockProperty{
			{"stage", []string{"0", "1"}},
		},
	},
	226: { // minecraft:cherry_shelf
		BaseID:    4284,
		DefaultID: 4293,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	258: { // minecraft:cherry_sign
		BaseID:    7019,
		DefaultID: 7036,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	658: { // minecraft:cherry_slab
		BaseID:    15207,
		DefaultID: 15210,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	570: { // minecraft:cherry_stairs
		BaseID:    13819,
		DefaultID: 13830,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	370: { // minecraft:cherry_trapdoor
		BaseID:    9104,
		DefaultID: 9119,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	297: { // minecraft:cherry_wall_hanging_sign
		BaseID:    8367,
		DefaultID: 8368,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	273: { // minecraft:cherry_wall_sign
		BaseID:    7447,
		DefaultID: 7448,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	80: { // minecraft:cherry_wood
		BaseID:    225,
		DefaultID: 226,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	245: { // minecraft:chest
		BaseID:    5543,
		DefaultID: 5544,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	521: { // minecraft:chipped_anvil
		BaseID:    12966,
		DefaultID: 12966,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	222: { // minecraft:chiseled_bookshelf
		BaseID:    3836,
		DefaultID: 3899,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"slot_0_occupied", []string{"true", "false"}},
			{"slot_1_occupied", []string{"true", "false"}},
			{"slot_2_occupied", []string{"true", "false"}},
			{"slot_3_occupied", []string{"true", "false"}},
			{"slot_4_occupied", []string{"true", "false"}},
			{"slot_5_occupied", []string{"true", "false"}},
		},
	},
	1114: { // minecraft:chiseled_cinnabar
		BaseID:    30516,
		DefaultID: 30516,
	},
	1142: { // minecraft:chiseled_copper
		BaseID:    31157,
		DefaultID: 31157,
	},
	1258: { // minecraft:chiseled_deepslate
		BaseID:    35420,
		DefaultID: 35420,
	},
	1031: { // minecraft:chiseled_nether_bricks
		BaseID:    26450,
		DefaultID: 26450,
	},
	1021: { // minecraft:chiseled_polished_blackstone
		BaseID:    25602,
		DefaultID: 25602,
	},
	532: { // minecraft:chiseled_quartz_block
		BaseID:    13091,
		DefaultID: 13091,
	},
	650: { // minecraft:chiseled_red_sandstone
		BaseID:    15095,
		DefaultID: 15095,
	},
	430: { // minecraft:chiseled_resin_bricks
		BaseID:    11067,
		DefaultID: 11067,
	},
	116: { // minecraft:chiseled_sandstone
		BaseID:    678,
		DefaultID: 678,
	},
	379: { // minecraft:chiseled_stone_bricks
		BaseID:    9491,
		DefaultID: 9491,
	},
	1101: { // minecraft:chiseled_sulfur
		BaseID:    29282,
		DefaultID: 29282,
	},
	1082: { // minecraft:chiseled_tuff
		BaseID:    27631,
		DefaultID: 27631,
	},
	1087: { // minecraft:chiseled_tuff_bricks
		BaseID:    28043,
		DefaultID: 28043,
	},
	714: { // minecraft:chorus_flower
		BaseID:    16681,
		DefaultID: 16681,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5"}},
		},
	},
	713: { // minecraft:chorus_plant
		BaseID:    16617,
		DefaultID: 16680,
		Properties: []blockProperty{
			{"down", []string{"true", "false"}},
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"up", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1102: { // minecraft:cinnabar
		BaseID:    29283,
		DefaultID: 29283,
	},
	1111: { // minecraft:cinnabar_brick_slab
		BaseID:    30106,
		DefaultID: 30109,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1112: { // minecraft:cinnabar_brick_stairs
		BaseID:    30112,
		DefaultID: 30123,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1113: { // minecraft:cinnabar_brick_wall
		BaseID:    30192,
		DefaultID: 30195,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	1110: { // minecraft:cinnabar_bricks
		BaseID:    30105,
		DefaultID: 30105,
	},
	1103: { // minecraft:cinnabar_slab
		BaseID:    29284,
		DefaultID: 29287,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1104: { // minecraft:cinnabar_stairs
		BaseID:    29290,
		DefaultID: 29301,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1105: { // minecraft:cinnabar_wall
		BaseID:    29370,
		DefaultID: 29373,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	330: { // minecraft:clay
		BaseID:    8616,
		DefaultID: 8616,
	},
	1282: { // minecraft:closed_eyeblossom
		BaseID:    35719,
		DefaultID: 35719,
	},
	609: { // minecraft:coal_block
		BaseID:    14760,
		DefaultID: 14760,
	},
	48: { // minecraft:coal_ore
		BaseID:    136,
		DefaultID: 136,
	},
	10: { // minecraft:coarse_dirt
		BaseID:    11,
		DefaultID: 11,
	},
	1242: { // minecraft:cobbled_deepslate
		BaseID:    33776,
		DefaultID: 33776,
	},
	1244: { // minecraft:cobbled_deepslate_slab
		BaseID:    33857,
		DefaultID: 33860,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1243: { // minecraft:cobbled_deepslate_stairs
		BaseID:    33777,
		DefaultID: 33788,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1245: { // minecraft:cobbled_deepslate_wall
		BaseID:    33863,
		DefaultID: 33866,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	12: { // minecraft:cobblestone
		BaseID:    14,
		DefaultID: 14,
	},
	669: { // minecraft:cobblestone_slab
		BaseID:    15273,
		DefaultID: 15276,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	268: { // minecraft:cobblestone_stairs
		BaseID:    7335,
		DefaultID: 7346,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	460: { // minecraft:cobblestone_wall
		BaseID:    11723,
		DefaultID: 11726,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	139: { // minecraft:cobweb
		BaseID:    2362,
		DefaultID: 2362,
	},
	446: { // minecraft:cocoa
		BaseID:    11215,
		DefaultID: 11215,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2"}},
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	458: { // minecraft:command_block
		BaseID:    11710,
		DefaultID: 11716,
		Properties: []blockProperty{
			{"conditional", []string{"true", "false"}},
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	526: { // minecraft:comparator
		BaseID:    13030,
		DefaultID: 13031,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"mode", []string{"compare", "subtract"}},
			{"powered", []string{"true", "false"}},
		},
	},
	999: { // minecraft:composter
		BaseID:    25100,
		DefaultID: 25100,
		Properties: []blockProperty{
			{"level", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}},
		},
	},
	880: { // minecraft:conduit
		BaseID:    18633,
		DefaultID: 18633,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	392: { // minecraft:copper_bars
		BaseID:    9724,
		DefaultID: 9755,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1124: { // minecraft:copper_block
		BaseID:    31139,
		DefaultID: 31139,
	},
	1190: { // minecraft:copper_bulb
		BaseID:    32893,
		DefaultID: 32896,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	401: { // minecraft:copper_chain
		BaseID:    9986,
		DefaultID: 9989,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1198: { // minecraft:copper_chest
		BaseID:    32925,
		DefaultID: 32926,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1166: { // minecraft:copper_door
		BaseID:    31853,
		DefaultID: 31864,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1206: { // minecraft:copper_golem_statue
		BaseID:    33117,
		DefaultID: 33118,
		Properties: []blockProperty{
			{"copper_golem_pose", []string{"standing", "sitting", "running", "star"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1182: { // minecraft:copper_grate
		BaseID:    32877,
		DefaultID: 32878,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	941: { // minecraft:copper_lantern
		BaseID:    24202,
		DefaultID: 24205,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1132: { // minecraft:copper_ore
		BaseID:    31147,
		DefaultID: 31147,
	},
	341: { // minecraft:copper_torch
		BaseID:    8681,
		DefaultID: 8681,
	},
	1174: { // minecraft:copper_trapdoor
		BaseID:    32365,
		DefaultID: 32380,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	342: { // minecraft:copper_wall_torch
		BaseID:    8682,
		DefaultID: 8682,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	212: { // minecraft:cornflower
		BaseID:    3825,
		DefaultID: 3825,
	},
	1259: { // minecraft:cracked_deepslate_bricks
		BaseID:    35421,
		DefaultID: 35421,
	},
	1260: { // minecraft:cracked_deepslate_tiles
		BaseID:    35422,
		DefaultID: 35422,
	},
	1032: { // minecraft:cracked_nether_bricks
		BaseID:    26451,
		DefaultID: 26451,
	},
	1020: { // minecraft:cracked_polished_blackstone_bricks
		BaseID:    25601,
		DefaultID: 25601,
	},
	378: { // minecraft:cracked_stone_bricks
		BaseID:    9490,
		DefaultID: 9490,
	},
	1274: { // minecraft:crafter
		BaseID:    35459,
		DefaultID: 35504,
		Properties: []blockProperty{
			{"crafting", []string{"true", "false"}},
			{"orientation", []string{"down_east", "down_north", "down_south", "down_west", "up_east", "up_north", "up_south", "up_west", "west_up", "east_up", "north_up", "south_up"}},
			{"triggered", []string{"true", "false"}},
		},
	},
	250: { // minecraft:crafting_table
		BaseID:    6866,
		DefaultID: 6866,
	},
	243: { // minecraft:creaking_heart
		BaseID:    5445,
		DefaultID: 5452,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"creaking_heart_state", []string{"uprooted", "dormant", "awake"}},
			{"natural", []string{"true", "false"}},
		},
	},
	514: { // minecraft:creeper_head
		BaseID:    12842,
		DefaultID: 12858,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	515: { // minecraft:creeper_wall_head
		BaseID:    12874,
		DefaultID: 12875,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	987: { // minecraft:crimson_button
		BaseID:    24823,
		DefaultID: 24832,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	989: { // minecraft:crimson_door
		BaseID:    24871,
		DefaultID: 24882,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	979: { // minecraft:crimson_fence
		BaseID:    24407,
		DefaultID: 24438,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	983: { // minecraft:crimson_fence_gate
		BaseID:    24599,
		DefaultID: 24606,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	966: { // minecraft:crimson_fungus
		BaseID:    24332,
		DefaultID: 24332,
	},
	288: { // minecraft:crimson_hanging_sign
		BaseID:    8015,
		DefaultID: 8064,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	963: { // minecraft:crimson_hyphae
		BaseID:    24325,
		DefaultID: 24326,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	965: { // minecraft:crimson_nylium
		BaseID:    24331,
		DefaultID: 24331,
	},
	973: { // minecraft:crimson_planks
		BaseID:    24389,
		DefaultID: 24389,
	},
	977: { // minecraft:crimson_pressure_plate
		BaseID:    24403,
		DefaultID: 24404,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	972: { // minecraft:crimson_roots
		BaseID:    24388,
		DefaultID: 24388,
	},
	227: { // minecraft:crimson_shelf
		BaseID:    4348,
		DefaultID: 4357,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	991: { // minecraft:crimson_sign
		BaseID:    24999,
		DefaultID: 25016,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	975: { // minecraft:crimson_slab
		BaseID:    24391,
		DefaultID: 24394,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	985: { // minecraft:crimson_stairs
		BaseID:    24663,
		DefaultID: 24674,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	961: { // minecraft:crimson_stem
		BaseID:    24319,
		DefaultID: 24320,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	981: { // minecraft:crimson_trapdoor
		BaseID:    24471,
		DefaultID: 24486,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	303: { // minecraft:crimson_wall_hanging_sign
		BaseID:    8415,
		DefaultID: 8416,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	993: { // minecraft:crimson_wall_sign
		BaseID:    25063,
		DefaultID: 25064,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1007: { // minecraft:crying_obsidian
		BaseID:    25177,
		DefaultID: 25177,
	},
	1134: { // minecraft:cut_copper
		BaseID:    31149,
		DefaultID: 31149,
	},
	1158: { // minecraft:cut_copper_slab
		BaseID:    31805,
		DefaultID: 31808,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1150: { // minecraft:cut_copper_stairs
		BaseID:    31165,
		DefaultID: 31176,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	651: { // minecraft:cut_red_sandstone
		BaseID:    15096,
		DefaultID: 15096,
	},
	676: { // minecraft:cut_red_sandstone_slab
		BaseID:    15315,
		DefaultID: 15318,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	117: { // minecraft:cut_sandstone
		BaseID:    679,
		DefaultID: 679,
	},
	667: { // minecraft:cut_sandstone_slab
		BaseID:    15261,
		DefaultID: 15264,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	626: { // minecraft:cyan_banner
		BaseID:    14918,
		DefaultID: 14926,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	128: { // minecraft:cyan_bed
		BaseID:    2174,
		DefaultID: 2177,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1044: { // minecraft:cyan_candle
		BaseID:    26613,
		DefaultID: 26616,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1061: { // minecraft:cyan_candle_cake
		BaseID:    26745,
		DefaultID: 26746,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	601: { // minecraft:cyan_carpet
		BaseID:    14752,
		DefaultID: 14752,
	},
	777: { // minecraft:cyan_concrete
		BaseID:    17020,
		DefaultID: 17020,
	},
	825: { // minecraft:cyan_concrete_powder
		BaseID:    18412,
		DefaultID: 18412,
	},
	809: { // minecraft:cyan_concrete_slab
		BaseID:    18361,
		DefaultID: 18364,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	793: { // minecraft:cyan_concrete_stairs
		BaseID:    17747,
		DefaultID: 17758,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	761: { // minecraft:cyan_glazed_terracotta
		BaseID:    16983,
		DefaultID: 16983,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	745: { // minecraft:cyan_shulker_box
		BaseID:    16905,
		DefaultID: 16909,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	358: { // minecraft:cyan_stained_glass
		BaseID:    8777,
		DefaultID: 8777,
	},
	562: { // minecraft:cyan_stained_glass_pane
		BaseID:    13515,
		DefaultID: 13546,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	546: { // minecraft:cyan_terracotta
		BaseID:    13220,
		DefaultID: 13220,
	},
	642: { // minecraft:cyan_wall_banner
		BaseID:    15066,
		DefaultID: 15066,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	160: { // minecraft:cyan_wool
		BaseID:    2418,
		DefaultID: 2418,
	},
	192: { // minecraft:cyan_wool_slab
		BaseID:    3759,
		DefaultID: 3762,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	176: { // minecraft:cyan_wool_stairs
		BaseID:    3145,
		DefaultID: 3156,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	522: { // minecraft:damaged_anvil
		BaseID:    12970,
		DefaultID: 12970,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	200: { // minecraft:dandelion
		BaseID:    3813,
		DefaultID: 3813,
	},
	501: { // minecraft:dark_oak_button
		BaseID:    12562,
		DefaultID: 12571,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	707: { // minecraft:dark_oak_door
		BaseID:    16291,
		DefaultID: 16302,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	697: { // minecraft:dark_oak_fence
		BaseID:    15811,
		DefaultID: 15842,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	687: { // minecraft:dark_oak_fence_gate
		BaseID:    15491,
		DefaultID: 15498,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	286: { // minecraft:dark_oak_hanging_sign
		BaseID:    7887,
		DefaultID: 7936,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	100: { // minecraft:dark_oak_leaves
		BaseID:    435,
		DefaultID: 462,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	57: { // minecraft:dark_oak_log
		BaseID:    157,
		DefaultID: 158,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	19: { // minecraft:dark_oak_planks
		BaseID:    21,
		DefaultID: 21,
	},
	315: { // minecraft:dark_oak_pressure_plate
		BaseID:    8541,
		DefaultID: 8542,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	32: { // minecraft:dark_oak_sapling
		BaseID:    42,
		DefaultID: 42,
		Properties: []blockProperty{
			{"stage", []string{"0", "1"}},
		},
	},
	228: { // minecraft:dark_oak_shelf
		BaseID:    4412,
		DefaultID: 4421,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	260: { // minecraft:dark_oak_sign
		BaseID:    7083,
		DefaultID: 7100,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	659: { // minecraft:dark_oak_slab
		BaseID:    15213,
		DefaultID: 15216,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	571: { // minecraft:dark_oak_stairs
		BaseID:    13899,
		DefaultID: 13910,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	371: { // minecraft:dark_oak_trapdoor
		BaseID:    9168,
		DefaultID: 9183,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	299: { // minecraft:dark_oak_wall_hanging_sign
		BaseID:    8383,
		DefaultID: 8384,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	275: { // minecraft:dark_oak_wall_sign
		BaseID:    7463,
		DefaultID: 7464,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	81: { // minecraft:dark_oak_wood
		BaseID:    228,
		DefaultID: 229,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	583: { // minecraft:dark_prismarine
		BaseID:    14480,
		DefaultID: 14480,
	},
	589: { // minecraft:dark_prismarine_slab
		BaseID:    14733,
		DefaultID: 14736,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	586: { // minecraft:dark_prismarine_stairs
		BaseID:    14641,
		DefaultID: 14652,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	527: { // minecraft:daylight_detector
		BaseID:    13046,
		DefaultID: 13062,
		Properties: []blockProperty{
			{"inverted", []string{"true", "false"}},
			{"power", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	849: { // minecraft:dead_brain_coral
		BaseID:    18506,
		DefaultID: 18506,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	839: { // minecraft:dead_brain_coral_block
		BaseID:    18495,
		DefaultID: 18495,
	},
	859: { // minecraft:dead_brain_coral_fan
		BaseID:    18526,
		DefaultID: 18526,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	869: { // minecraft:dead_brain_coral_wall_fan
		BaseID:    18552,
		DefaultID: 18552,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	850: { // minecraft:dead_bubble_coral
		BaseID:    18508,
		DefaultID: 18508,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	840: { // minecraft:dead_bubble_coral_block
		BaseID:    18496,
		DefaultID: 18496,
	},
	860: { // minecraft:dead_bubble_coral_fan
		BaseID:    18528,
		DefaultID: 18528,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	870: { // minecraft:dead_bubble_coral_wall_fan
		BaseID:    18560,
		DefaultID: 18560,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	142: { // minecraft:dead_bush
		BaseID:    2365,
		DefaultID: 2365,
	},
	851: { // minecraft:dead_fire_coral
		BaseID:    18510,
		DefaultID: 18510,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	841: { // minecraft:dead_fire_coral_block
		BaseID:    18497,
		DefaultID: 18497,
	},
	861: { // minecraft:dead_fire_coral_fan
		BaseID:    18530,
		DefaultID: 18530,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	871: { // minecraft:dead_fire_coral_wall_fan
		BaseID:    18568,
		DefaultID: 18568,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	852: { // minecraft:dead_horn_coral
		BaseID:    18512,
		DefaultID: 18512,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	842: { // minecraft:dead_horn_coral_block
		BaseID:    18498,
		DefaultID: 18498,
	},
	862: { // minecraft:dead_horn_coral_fan
		BaseID:    18532,
		DefaultID: 18532,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	872: { // minecraft:dead_horn_coral_wall_fan
		BaseID:    18576,
		DefaultID: 18576,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	848: { // minecraft:dead_tube_coral
		BaseID:    18504,
		DefaultID: 18504,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	838: { // minecraft:dead_tube_coral_block
		BaseID:    18494,
		DefaultID: 18494,
	},
	858: { // minecraft:dead_tube_coral_fan
		BaseID:    18524,
		DefaultID: 18524,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	868: { // minecraft:dead_tube_coral_wall_fan
		BaseID:    18544,
		DefaultID: 18544,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1273: { // minecraft:decorated_pot
		BaseID:    35443,
		DefaultID: 35452,
		Properties: []blockProperty{
			{"cracked", []string{"true", "false"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1241: { // minecraft:deepslate
		BaseID:    33773,
		DefaultID: 33774,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	1256: { // minecraft:deepslate_brick_slab
		BaseID:    35090,
		DefaultID: 35093,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1255: { // minecraft:deepslate_brick_stairs
		BaseID:    35010,
		DefaultID: 35021,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1257: { // minecraft:deepslate_brick_wall
		BaseID:    35096,
		DefaultID: 35099,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	1254: { // minecraft:deepslate_bricks
		BaseID:    35009,
		DefaultID: 35009,
	},
	49: { // minecraft:deepslate_coal_ore
		BaseID:    137,
		DefaultID: 137,
	},
	1133: { // minecraft:deepslate_copper_ore
		BaseID:    31148,
		DefaultID: 31148,
	},
	248: { // minecraft:deepslate_diamond_ore
		BaseID:    6864,
		DefaultID: 6864,
	},
	450: { // minecraft:deepslate_emerald_ore
		BaseID:    11316,
		DefaultID: 11316,
	},
	45: { // minecraft:deepslate_gold_ore
		BaseID:    133,
		DefaultID: 133,
	},
	47: { // minecraft:deepslate_iron_ore
		BaseID:    135,
		DefaultID: 135,
	},
	112: { // minecraft:deepslate_lapis_ore
		BaseID:    663,
		DefaultID: 663,
	},
	321: { // minecraft:deepslate_redstone_ore
		BaseID:    8553,
		DefaultID: 8554,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	1252: { // minecraft:deepslate_tile_slab
		BaseID:    34679,
		DefaultID: 34682,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1251: { // minecraft:deepslate_tile_stairs
		BaseID:    34599,
		DefaultID: 34610,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1253: { // minecraft:deepslate_tile_wall
		BaseID:    34685,
		DefaultID: 34688,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	1250: { // minecraft:deepslate_tiles
		BaseID:    34598,
		DefaultID: 34598,
	},
	137: { // minecraft:detector_rail
		BaseID:    2326,
		DefaultID: 2339,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"shape", []string{"north_south", "east_west", "ascending_east", "ascending_west", "ascending_north", "ascending_south"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	249: { // minecraft:diamond_block
		BaseID:    6865,
		DefaultID: 6865,
	},
	247: { // minecraft:diamond_ore
		BaseID:    6863,
		DefaultID: 6863,
	},
	4: { // minecraft:diorite
		BaseID:    4,
		DefaultID: 4,
	},
	913: { // minecraft:diorite_slab
		BaseID:    19845,
		DefaultID: 19848,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	900: { // minecraft:diorite_stairs
		BaseID:    19693,
		DefaultID: 19704,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	926: { // minecraft:diorite_wall
		BaseID:    23739,
		DefaultID: 23742,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	9: { // minecraft:dirt
		BaseID:    10,
		DefaultID: 10,
	},
	724: { // minecraft:dirt_path
		BaseID:    16796,
		DefaultID: 16796,
	},
	114: { // minecraft:dispenser
		BaseID:    665,
		DefaultID: 666,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"triggered", []string{"true", "false"}},
		},
	},
	444: { // minecraft:dragon_egg
		BaseID:    11212,
		DefaultID: 11212,
	},
	516: { // minecraft:dragon_head
		BaseID:    12882,
		DefaultID: 12898,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	517: { // minecraft:dragon_wall_head
		BaseID:    12914,
		DefaultID: 12915,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	837: { // minecraft:dried_ghast
		BaseID:    18462,
		DefaultID: 18463,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"hydration", []string{"0", "1", "2", "3"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	834: { // minecraft:dried_kelp_block
		BaseID:    18446,
		DefaultID: 18446,
	},
	1222: { // minecraft:dripstone_block
		BaseID:    33565,
		DefaultID: 33565,
	},
	536: { // minecraft:dropper
		BaseID:    13199,
		DefaultID: 13200,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"triggered", []string{"true", "false"}},
		},
	},
	454: { // minecraft:emerald_block
		BaseID:    11469,
		DefaultID: 11469,
	},
	449: { // minecraft:emerald_ore
		BaseID:    11315,
		DefaultID: 11315,
	},
	435: { // minecraft:enchanting_table
		BaseID:    11185,
		DefaultID: 11185,
	},
	725: { // minecraft:end_gateway
		BaseID:    16797,
		DefaultID: 16797,
	},
	441: { // minecraft:end_portal
		BaseID:    11202,
		DefaultID: 11202,
	},
	442: { // minecraft:end_portal_frame
		BaseID:    11203,
		DefaultID: 11207,
		Properties: []blockProperty{
			{"eye", []string{"true", "false"}},
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	712: { // minecraft:end_rod
		BaseID:    16611,
		DefaultID: 16615,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	443: { // minecraft:end_stone
		BaseID:    11211,
		DefaultID: 11211,
	},
	906: { // minecraft:end_stone_brick_slab
		BaseID:    19803,
		DefaultID: 19806,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	892: { // minecraft:end_stone_brick_stairs
		BaseID:    19053,
		DefaultID: 19064,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	925: { // minecraft:end_stone_brick_wall
		BaseID:    23415,
		DefaultID: 23418,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	719: { // minecraft:end_stone_bricks
		BaseID:    16777,
		DefaultID: 16777,
	},
	451: { // minecraft:ender_chest
		BaseID:    11317,
		DefaultID: 11318,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1143: { // minecraft:exposed_chiseled_copper
		BaseID:    31158,
		DefaultID: 31158,
	},
	1125: { // minecraft:exposed_copper
		BaseID:    31140,
		DefaultID: 31140,
	},
	393: { // minecraft:exposed_copper_bars
		BaseID:    9756,
		DefaultID: 9787,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1191: { // minecraft:exposed_copper_bulb
		BaseID:    32897,
		DefaultID: 32900,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	402: { // minecraft:exposed_copper_chain
		BaseID:    9992,
		DefaultID: 9995,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1199: { // minecraft:exposed_copper_chest
		BaseID:    32949,
		DefaultID: 32950,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1167: { // minecraft:exposed_copper_door
		BaseID:    31917,
		DefaultID: 31928,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1207: { // minecraft:exposed_copper_golem_statue
		BaseID:    33149,
		DefaultID: 33150,
		Properties: []blockProperty{
			{"copper_golem_pose", []string{"standing", "sitting", "running", "star"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1183: { // minecraft:exposed_copper_grate
		BaseID:    32879,
		DefaultID: 32880,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	942: { // minecraft:exposed_copper_lantern
		BaseID:    24206,
		DefaultID: 24209,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1175: { // minecraft:exposed_copper_trapdoor
		BaseID:    32429,
		DefaultID: 32444,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1135: { // minecraft:exposed_cut_copper
		BaseID:    31150,
		DefaultID: 31150,
	},
	1159: { // minecraft:exposed_cut_copper_slab
		BaseID:    31811,
		DefaultID: 31814,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1151: { // minecraft:exposed_cut_copper_stairs
		BaseID:    31245,
		DefaultID: 31256,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1215: { // minecraft:exposed_lightning_rod
		BaseID:    33397,
		DefaultID: 33416,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	252: { // minecraft:farmland
		BaseID:    6875,
		DefaultID: 6875,
		Properties: []blockProperty{
			{"moisture", []string{"0", "1", "2", "3", "4", "5", "6", "7"}},
		},
	},
	141: { // minecraft:fern
		BaseID:    2364,
		DefaultID: 2364,
	},
	240: { // minecraft:fire
		BaseID:    4931,
		DefaultID: 4962,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"up", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	856: { // minecraft:fire_coral
		BaseID:    18520,
		DefaultID: 18520,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	846: { // minecraft:fire_coral_block
		BaseID:    18502,
		DefaultID: 18502,
	},
	866: { // minecraft:fire_coral_fan
		BaseID:    18540,
		DefaultID: 18540,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	876: { // minecraft:fire_coral_wall_fan
		BaseID:    18608,
		DefaultID: 18608,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1285: { // minecraft:firefly_bush
		BaseID:    35722,
		DefaultID: 35722,
	},
	933: { // minecraft:fletching_table
		BaseID:    24128,
		DefaultID: 24128,
	},
	462: { // minecraft:flower_pot
		BaseID:    12371,
		DefaultID: 12371,
	},
	1229: { // minecraft:flowering_azalea
		BaseID:    33662,
		DefaultID: 33662,
	},
	107: { // minecraft:flowering_azalea_leaves
		BaseID:    631,
		DefaultID: 658,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1271: { // minecraft:frogspawn
		BaseID:    35441,
		DefaultID: 35441,
	},
	728: { // minecraft:frosted_ice
		BaseID:    16822,
		DefaultID: 16822,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3"}},
		},
	},
	253: { // minecraft:furnace
		BaseID:    6883,
		DefaultID: 6884,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"lit", []string{"true", "false"}},
		},
	},
	1025: { // minecraft:gilded_blackstone
		BaseID:    26013,
		DefaultID: 26013,
	},
	110: { // minecraft:glass
		BaseID:    661,
		DefaultID: 661,
	},
	409: { // minecraft:glass_pane
		BaseID:    10034,
		DefaultID: 10065,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	417: { // minecraft:glow_lichen
		BaseID:    10124,
		DefaultID: 10251,
		Properties: []blockProperty{
			{"down", []string{"true", "false"}},
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	343: { // minecraft:glowstone
		BaseID:    8686,
		DefaultID: 8686,
	},
	217: { // minecraft:gold_block
		BaseID:    3830,
		DefaultID: 3830,
	},
	44: { // minecraft:gold_ore
		BaseID:    132,
		DefaultID: 132,
	},
	201: { // minecraft:golden_dandelion
		BaseID:    3814,
		DefaultID: 3814,
	},
	2: { // minecraft:granite
		BaseID:    2,
		DefaultID: 2,
	},
	909: { // minecraft:granite_slab
		BaseID:    19821,
		DefaultID: 19824,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	896: { // minecraft:granite_stairs
		BaseID:    19373,
		DefaultID: 19384,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	918: { // minecraft:granite_wall
		BaseID:    21147,
		DefaultID: 21150,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	8: { // minecraft:grass_block
		BaseID:    8,
		DefaultID: 9,
		Properties: []blockProperty{
			{"snowy", []string{"true", "false"}},
		},
	},
	42: { // minecraft:gravel
		BaseID:    127,
		DefaultID: 127,
	},
	624: { // minecraft:gray_banner
		BaseID:    14886,
		DefaultID: 14894,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	126: { // minecraft:gray_bed
		BaseID:    2142,
		DefaultID: 2145,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1042: { // minecraft:gray_candle
		BaseID:    26581,
		DefaultID: 26584,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1059: { // minecraft:gray_candle_cake
		BaseID:    26741,
		DefaultID: 26742,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	599: { // minecraft:gray_carpet
		BaseID:    14750,
		DefaultID: 14750,
	},
	775: { // minecraft:gray_concrete
		BaseID:    17018,
		DefaultID: 17018,
	},
	823: { // minecraft:gray_concrete_powder
		BaseID:    18410,
		DefaultID: 18410,
	},
	807: { // minecraft:gray_concrete_slab
		BaseID:    18349,
		DefaultID: 18352,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	791: { // minecraft:gray_concrete_stairs
		BaseID:    17587,
		DefaultID: 17598,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	759: { // minecraft:gray_glazed_terracotta
		BaseID:    16975,
		DefaultID: 16975,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	743: { // minecraft:gray_shulker_box
		BaseID:    16893,
		DefaultID: 16897,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	356: { // minecraft:gray_stained_glass
		BaseID:    8775,
		DefaultID: 8775,
	},
	560: { // minecraft:gray_stained_glass_pane
		BaseID:    13451,
		DefaultID: 13482,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	544: { // minecraft:gray_terracotta
		BaseID:    13218,
		DefaultID: 13218,
	},
	640: { // minecraft:gray_wall_banner
		BaseID:    15058,
		DefaultID: 15058,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	158: { // minecraft:gray_wool
		BaseID:    2416,
		DefaultID: 2416,
	},
	190: { // minecraft:gray_wool_slab
		BaseID:    3747,
		DefaultID: 3750,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	174: { // minecraft:gray_wool_stairs
		BaseID:    2985,
		DefaultID: 2996,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	630: { // minecraft:green_banner
		BaseID:    14982,
		DefaultID: 14990,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	132: { // minecraft:green_bed
		BaseID:    2238,
		DefaultID: 2241,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1048: { // minecraft:green_candle
		BaseID:    26677,
		DefaultID: 26680,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1065: { // minecraft:green_candle_cake
		BaseID:    26753,
		DefaultID: 26754,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	605: { // minecraft:green_carpet
		BaseID:    14756,
		DefaultID: 14756,
	},
	781: { // minecraft:green_concrete
		BaseID:    17024,
		DefaultID: 17024,
	},
	829: { // minecraft:green_concrete_powder
		BaseID:    18416,
		DefaultID: 18416,
	},
	813: { // minecraft:green_concrete_slab
		BaseID:    18385,
		DefaultID: 18388,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	797: { // minecraft:green_concrete_stairs
		BaseID:    18067,
		DefaultID: 18078,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	765: { // minecraft:green_glazed_terracotta
		BaseID:    16999,
		DefaultID: 16999,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	749: { // minecraft:green_shulker_box
		BaseID:    16929,
		DefaultID: 16933,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	362: { // minecraft:green_stained_glass
		BaseID:    8781,
		DefaultID: 8781,
	},
	566: { // minecraft:green_stained_glass_pane
		BaseID:    13643,
		DefaultID: 13674,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	550: { // minecraft:green_terracotta
		BaseID:    13224,
		DefaultID: 13224,
	},
	646: { // minecraft:green_wall_banner
		BaseID:    15082,
		DefaultID: 15082,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	164: { // minecraft:green_wool
		BaseID:    2422,
		DefaultID: 2422,
	},
	196: { // minecraft:green_wool_slab
		BaseID:    3783,
		DefaultID: 3786,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	180: { // minecraft:green_wool_stairs
		BaseID:    3465,
		DefaultID: 3476,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	934: { // minecraft:grindstone
		BaseID:    24129,
		DefaultID: 24133,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	1238: { // minecraft:hanging_roots
		BaseID:    33769,
		DefaultID: 33770,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	591: { // minecraft:hay_block
		BaseID:    14740,
		DefaultID: 14741,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	1277: { // minecraft:heavy_core
		BaseID:    35551,
		DefaultID: 35552,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	525: { // minecraft:heavy_weighted_pressure_plate
		BaseID:    13014,
		DefaultID: 13014,
		Properties: []blockProperty{
			{"power", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	1003: { // minecraft:honey_block
		BaseID:    25173,
		DefaultID: 25173,
	},
	1004: { // minecraft:honeycomb_block
		BaseID:    25174,
		DefaultID: 25174,
	},
	530: { // minecraft:hopper
		BaseID:    13080,
		DefaultID: 13080,
		Properties: []blockProperty{
			{"enabled", []string{"true", "false"}},
			{"facing", []string{"down", "north", "south", "west", "east"}},
		},
	},
	857: { // minecraft:horn_coral
		BaseID:    18522,
		DefaultID: 18522,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	847: { // minecraft:horn_coral_block
		BaseID:    18503,
		DefaultID: 18503,
	},
	867: { // minecraft:horn_coral_fan
		BaseID:    18542,
		DefaultID: 18542,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	877: { // minecraft:horn_coral_wall_fan
		BaseID:    18616,
		DefaultID: 18616,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	326: { // minecraft:ice
		BaseID:    8597,
		DefaultID: 8597,
	},
	387: { // minecraft:infested_chiseled_stone_bricks
		BaseID:    9499,
		DefaultID: 9499,
	},
	383: { // minecraft:infested_cobblestone
		BaseID:    9495,
		DefaultID: 9495,
	},
	386: { // minecraft:infested_cracked_stone_bricks
		BaseID:    9498,
		DefaultID: 9498,
	},
	1261: { // minecraft:infested_deepslate
		BaseID:    35423,
		DefaultID: 35424,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	385: { // minecraft:infested_mossy_stone_bricks
		BaseID:    9497,
		DefaultID: 9497,
	},
	382: { // minecraft:infested_stone
		BaseID:    9494,
		DefaultID: 9494,
	},
	384: { // minecraft:infested_stone_bricks
		BaseID:    9496,
		DefaultID: 9496,
	},
	391: { // minecraft:iron_bars
		BaseID:    9692,
		DefaultID: 9723,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	218: { // minecraft:iron_block
		BaseID:    3831,
		DefaultID: 3831,
	},
	400: { // minecraft:iron_chain
		BaseID:    9980,
		DefaultID: 9983,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	308: { // minecraft:iron_door
		BaseID:    8465,
		DefaultID: 8476,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	46: { // minecraft:iron_ore
		BaseID:    134,
		DefaultID: 134,
	},
	580: { // minecraft:iron_trapdoor
		BaseID:    14414,
		DefaultID: 14429,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	346: { // minecraft:jack_o_lantern
		BaseID:    8693,
		DefaultID: 8693,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	996: { // minecraft:jigsaw
		BaseID:    25083,
		DefaultID: 25093,
		Properties: []blockProperty{
			{"orientation", []string{"down_east", "down_north", "down_south", "down_west", "up_east", "up_north", "up_south", "up_west", "west_up", "east_up", "north_up", "south_up"}},
		},
	},
	332: { // minecraft:jukebox
		BaseID:    8633,
		DefaultID: 8634,
		Properties: []blockProperty{
			{"has_record", []string{"true", "false"}},
		},
	},
	498: { // minecraft:jungle_button
		BaseID:    12490,
		DefaultID: 12499,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	704: { // minecraft:jungle_door
		BaseID:    16099,
		DefaultID: 16110,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	694: { // minecraft:jungle_fence
		BaseID:    15715,
		DefaultID: 15746,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	684: { // minecraft:jungle_fence_gate
		BaseID:    15395,
		DefaultID: 15402,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	285: { // minecraft:jungle_hanging_sign
		BaseID:    7823,
		DefaultID: 7872,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	97: { // minecraft:jungle_leaves
		BaseID:    351,
		DefaultID: 378,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	54: { // minecraft:jungle_log
		BaseID:    148,
		DefaultID: 149,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	16: { // minecraft:jungle_planks
		BaseID:    18,
		DefaultID: 18,
	},
	312: { // minecraft:jungle_pressure_plate
		BaseID:    8535,
		DefaultID: 8536,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	29: { // minecraft:jungle_sapling
		BaseID:    36,
		DefaultID: 36,
		Properties: []blockProperty{
			{"stage", []string{"0", "1"}},
		},
	},
	229: { // minecraft:jungle_shelf
		BaseID:    4476,
		DefaultID: 4485,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	259: { // minecraft:jungle_sign
		BaseID:    7051,
		DefaultID: 7068,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	656: { // minecraft:jungle_slab
		BaseID:    15195,
		DefaultID: 15198,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	457: { // minecraft:jungle_stairs
		BaseID:    11630,
		DefaultID: 11641,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	368: { // minecraft:jungle_trapdoor
		BaseID:    8976,
		DefaultID: 8991,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	298: { // minecraft:jungle_wall_hanging_sign
		BaseID:    8375,
		DefaultID: 8376,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	274: { // minecraft:jungle_wall_sign
		BaseID:    7455,
		DefaultID: 7456,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	78: { // minecraft:jungle_wood
		BaseID:    219,
		DefaultID: 220,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	832: { // minecraft:kelp
		BaseID:    18419,
		DefaultID: 18419,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25"}},
		},
	},
	833: { // minecraft:kelp_plant
		BaseID:    18445,
		DefaultID: 18445,
	},
	266: { // minecraft:ladder
		BaseID:    7307,
		DefaultID: 7308,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	939: { // minecraft:lantern
		BaseID:    24194,
		DefaultID: 24197,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	113: { // minecraft:lapis_block
		BaseID:    664,
		DefaultID: 664,
	},
	111: { // minecraft:lapis_ore
		BaseID:    662,
		DefaultID: 662,
	},
	1071: { // minecraft:large_amethyst_bud
		BaseID:    26773,
		DefaultID: 26782,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	616: { // minecraft:large_fern
		BaseID:    14772,
		DefaultID: 14773,
		Properties: []blockProperty{
			{"half", []string{"upper", "lower"}},
		},
	},
	38: { // minecraft:lava
		BaseID:    105,
		DefaultID: 105,
		Properties: []blockProperty{
			{"level", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	439: { // minecraft:lava_cauldron
		BaseID:    11198,
		DefaultID: 11198,
	},
	1233: { // minecraft:leaf_litter
		BaseID:    33696,
		DefaultID: 33696,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"segment_amount", []string{"1", "2", "3", "4"}},
		},
	},
	935: { // minecraft:lectern
		BaseID:    24141,
		DefaultID: 24144,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"has_book", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	306: { // minecraft:lever
		BaseID:    8439,
		DefaultID: 8448,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	579: { // minecraft:light
		BaseID:    14382,
		DefaultID: 14413,
		Properties: []blockProperty{
			{"level", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	620: { // minecraft:light_blue_banner
		BaseID:    14822,
		DefaultID: 14830,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	122: { // minecraft:light_blue_bed
		BaseID:    2078,
		DefaultID: 2081,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1038: { // minecraft:light_blue_candle
		BaseID:    26517,
		DefaultID: 26520,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1055: { // minecraft:light_blue_candle_cake
		BaseID:    26733,
		DefaultID: 26734,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	595: { // minecraft:light_blue_carpet
		BaseID:    14746,
		DefaultID: 14746,
	},
	771: { // minecraft:light_blue_concrete
		BaseID:    17014,
		DefaultID: 17014,
	},
	819: { // minecraft:light_blue_concrete_powder
		BaseID:    18406,
		DefaultID: 18406,
	},
	803: { // minecraft:light_blue_concrete_slab
		BaseID:    18325,
		DefaultID: 18328,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	787: { // minecraft:light_blue_concrete_stairs
		BaseID:    17267,
		DefaultID: 17278,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	755: { // minecraft:light_blue_glazed_terracotta
		BaseID:    16959,
		DefaultID: 16959,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	739: { // minecraft:light_blue_shulker_box
		BaseID:    16869,
		DefaultID: 16873,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	352: { // minecraft:light_blue_stained_glass
		BaseID:    8771,
		DefaultID: 8771,
	},
	556: { // minecraft:light_blue_stained_glass_pane
		BaseID:    13323,
		DefaultID: 13354,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	540: { // minecraft:light_blue_terracotta
		BaseID:    13214,
		DefaultID: 13214,
	},
	636: { // minecraft:light_blue_wall_banner
		BaseID:    15042,
		DefaultID: 15042,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	154: { // minecraft:light_blue_wool
		BaseID:    2412,
		DefaultID: 2412,
	},
	186: { // minecraft:light_blue_wool_slab
		BaseID:    3723,
		DefaultID: 3726,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	170: { // minecraft:light_blue_wool_stairs
		BaseID:    2665,
		DefaultID: 2676,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	625: { // minecraft:light_gray_banner
		BaseID:    14902,
		DefaultID: 14910,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	127: { // minecraft:light_gray_bed
		BaseID:    2158,
		DefaultID: 2161,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1043: { // minecraft:light_gray_candle
		BaseID:    26597,
		DefaultID: 26600,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1060: { // minecraft:light_gray_candle_cake
		BaseID:    26743,
		DefaultID: 26744,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	600: { // minecraft:light_gray_carpet
		BaseID:    14751,
		DefaultID: 14751,
	},
	776: { // minecraft:light_gray_concrete
		BaseID:    17019,
		DefaultID: 17019,
	},
	824: { // minecraft:light_gray_concrete_powder
		BaseID:    18411,
		DefaultID: 18411,
	},
	808: { // minecraft:light_gray_concrete_slab
		BaseID:    18355,
		DefaultID: 18358,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	792: { // minecraft:light_gray_concrete_stairs
		BaseID:    17667,
		DefaultID: 17678,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	760: { // minecraft:light_gray_glazed_terracotta
		BaseID:    16979,
		DefaultID: 16979,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	744: { // minecraft:light_gray_shulker_box
		BaseID:    16899,
		DefaultID: 16903,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	357: { // minecraft:light_gray_stained_glass
		BaseID:    8776,
		DefaultID: 8776,
	},
	561: { // minecraft:light_gray_stained_glass_pane
		BaseID:    13483,
		DefaultID: 13514,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	545: { // minecraft:light_gray_terracotta
		BaseID:    13219,
		DefaultID: 13219,
	},
	641: { // minecraft:light_gray_wall_banner
		BaseID:    15062,
		DefaultID: 15062,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	159: { // minecraft:light_gray_wool
		BaseID:    2417,
		DefaultID: 2417,
	},
	191: { // minecraft:light_gray_wool_slab
		BaseID:    3753,
		DefaultID: 3756,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	175: { // minecraft:light_gray_wool_stairs
		BaseID:    3065,
		DefaultID: 3076,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	524: { // minecraft:light_weighted_pressure_plate
		BaseID:    12998,
		DefaultID: 12998,
		Properties: []blockProperty{
			{"power", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	1214: { // minecraft:lightning_rod
		BaseID:    33373,
		DefaultID: 33392,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	612: { // minecraft:lilac
		BaseID:    14764,
		DefaultID: 14765,
		Properties: []blockProperty{
			{"half", []string{"upper", "lower"}},
		},
	},
	214: { // minecraft:lily_of_the_valley
		BaseID:    3827,
		DefaultID: 3827,
	},
	424: { // minecraft:lily_pad
		BaseID:    10654,
		DefaultID: 10654,
	},
	622: { // minecraft:lime_banner
		BaseID:    14854,
		DefaultID: 14862,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	124: { // minecraft:lime_bed
		BaseID:    2110,
		DefaultID: 2113,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1040: { // minecraft:lime_candle
		BaseID:    26549,
		DefaultID: 26552,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1057: { // minecraft:lime_candle_cake
		BaseID:    26737,
		DefaultID: 26738,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	597: { // minecraft:lime_carpet
		BaseID:    14748,
		DefaultID: 14748,
	},
	773: { // minecraft:lime_concrete
		BaseID:    17016,
		DefaultID: 17016,
	},
	821: { // minecraft:lime_concrete_powder
		BaseID:    18408,
		DefaultID: 18408,
	},
	805: { // minecraft:lime_concrete_slab
		BaseID:    18337,
		DefaultID: 18340,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	789: { // minecraft:lime_concrete_stairs
		BaseID:    17427,
		DefaultID: 17438,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	757: { // minecraft:lime_glazed_terracotta
		BaseID:    16967,
		DefaultID: 16967,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	741: { // minecraft:lime_shulker_box
		BaseID:    16881,
		DefaultID: 16885,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	354: { // minecraft:lime_stained_glass
		BaseID:    8773,
		DefaultID: 8773,
	},
	558: { // minecraft:lime_stained_glass_pane
		BaseID:    13387,
		DefaultID: 13418,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	542: { // minecraft:lime_terracotta
		BaseID:    13216,
		DefaultID: 13216,
	},
	638: { // minecraft:lime_wall_banner
		BaseID:    15050,
		DefaultID: 15050,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	156: { // minecraft:lime_wool
		BaseID:    2414,
		DefaultID: 2414,
	},
	188: { // minecraft:lime_wool_slab
		BaseID:    3735,
		DefaultID: 3738,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	172: { // minecraft:lime_wool_stairs
		BaseID:    2825,
		DefaultID: 2836,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1013: { // minecraft:lodestone
		BaseID:    25187,
		DefaultID: 25187,
	},
	928: { // minecraft:loom
		BaseID:    24095,
		DefaultID: 24095,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	619: { // minecraft:magenta_banner
		BaseID:    14806,
		DefaultID: 14814,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	121: { // minecraft:magenta_bed
		BaseID:    2062,
		DefaultID: 2065,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1037: { // minecraft:magenta_candle
		BaseID:    26501,
		DefaultID: 26504,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1054: { // minecraft:magenta_candle_cake
		BaseID:    26731,
		DefaultID: 26732,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	594: { // minecraft:magenta_carpet
		BaseID:    14745,
		DefaultID: 14745,
	},
	770: { // minecraft:magenta_concrete
		BaseID:    17013,
		DefaultID: 17013,
	},
	818: { // minecraft:magenta_concrete_powder
		BaseID:    18405,
		DefaultID: 18405,
	},
	802: { // minecraft:magenta_concrete_slab
		BaseID:    18319,
		DefaultID: 18322,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	786: { // minecraft:magenta_concrete_stairs
		BaseID:    17187,
		DefaultID: 17198,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	754: { // minecraft:magenta_glazed_terracotta
		BaseID:    16955,
		DefaultID: 16955,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	738: { // minecraft:magenta_shulker_box
		BaseID:    16863,
		DefaultID: 16867,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	351: { // minecraft:magenta_stained_glass
		BaseID:    8770,
		DefaultID: 8770,
	},
	555: { // minecraft:magenta_stained_glass_pane
		BaseID:    13291,
		DefaultID: 13322,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	539: { // minecraft:magenta_terracotta
		BaseID:    13213,
		DefaultID: 13213,
	},
	635: { // minecraft:magenta_wall_banner
		BaseID:    15038,
		DefaultID: 15038,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	153: { // minecraft:magenta_wool
		BaseID:    2411,
		DefaultID: 2411,
	},
	185: { // minecraft:magenta_wool_slab
		BaseID:    3717,
		DefaultID: 3720,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	169: { // minecraft:magenta_wool_stairs
		BaseID:    2585,
		DefaultID: 2596,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	729: { // minecraft:magma_block
		BaseID:    16826,
		DefaultID: 16826,
	},
	503: { // minecraft:mangrove_button
		BaseID:    12610,
		DefaultID: 12619,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	709: { // minecraft:mangrove_door
		BaseID:    16419,
		DefaultID: 16430,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	699: { // minecraft:mangrove_fence
		BaseID:    15875,
		DefaultID: 15906,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	689: { // minecraft:mangrove_fence_gate
		BaseID:    15555,
		DefaultID: 15562,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	290: { // minecraft:mangrove_hanging_sign
		BaseID:    8143,
		DefaultID: 8192,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	102: { // minecraft:mangrove_leaves
		BaseID:    491,
		DefaultID: 518,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	59: { // minecraft:mangrove_log
		BaseID:    163,
		DefaultID: 164,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	22: { // minecraft:mangrove_planks
		BaseID:    26,
		DefaultID: 26,
	},
	317: { // minecraft:mangrove_pressure_plate
		BaseID:    8545,
		DefaultID: 8546,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	34: { // minecraft:mangrove_propagule
		BaseID:    46,
		DefaultID: 51,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4"}},
			{"hanging", []string{"true", "false"}},
			{"stage", []string{"0", "1"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	61: { // minecraft:mangrove_roots
		BaseID:    169,
		DefaultID: 170,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	230: { // minecraft:mangrove_shelf
		BaseID:    4540,
		DefaultID: 4549,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	262: { // minecraft:mangrove_sign
		BaseID:    7147,
		DefaultID: 7164,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	661: { // minecraft:mangrove_slab
		BaseID:    15225,
		DefaultID: 15228,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	573: { // minecraft:mangrove_stairs
		BaseID:    14059,
		DefaultID: 14070,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	373: { // minecraft:mangrove_trapdoor
		BaseID:    9296,
		DefaultID: 9311,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	301: { // minecraft:mangrove_wall_hanging_sign
		BaseID:    8399,
		DefaultID: 8400,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	277: { // minecraft:mangrove_wall_sign
		BaseID:    7479,
		DefaultID: 7480,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	82: { // minecraft:mangrove_wood
		BaseID:    231,
		DefaultID: 232,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	1072: { // minecraft:medium_amethyst_bud
		BaseID:    26785,
		DefaultID: 26794,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	411: { // minecraft:melon
		BaseID:    10067,
		DefaultID: 10067,
	},
	415: { // minecraft:melon_stem
		BaseID:    10084,
		DefaultID: 10084,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7"}},
		},
	},
	1234: { // minecraft:moss_block
		BaseID:    33712,
		DefaultID: 33712,
	},
	1230: { // minecraft:moss_carpet
		BaseID:    33663,
		DefaultID: 33663,
	},
	236: { // minecraft:mossy_cobblestone
		BaseID:    4924,
		DefaultID: 4924,
	},
	905: { // minecraft:mossy_cobblestone_slab
		BaseID:    19797,
		DefaultID: 19800,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	891: { // minecraft:mossy_cobblestone_stairs
		BaseID:    18973,
		DefaultID: 18984,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	461: { // minecraft:mossy_cobblestone_wall
		BaseID:    12047,
		DefaultID: 12050,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	903: { // minecraft:mossy_stone_brick_slab
		BaseID:    19785,
		DefaultID: 19788,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	889: { // minecraft:mossy_stone_brick_stairs
		BaseID:    18813,
		DefaultID: 18824,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	917: { // minecraft:mossy_stone_brick_wall
		BaseID:    20823,
		DefaultID: 20826,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	377: { // minecraft:mossy_stone_bricks
		BaseID:    9489,
		DefaultID: 9489,
	},
	199: { // minecraft:moving_piston
		BaseID:    3801,
		DefaultID: 3801,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"type", []string{"normal", "sticky"}},
		},
	},
	1240: { // minecraft:mud
		BaseID:    33772,
		DefaultID: 33772,
	},
	672: { // minecraft:mud_brick_slab
		BaseID:    15291,
		DefaultID: 15294,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	422: { // minecraft:mud_brick_stairs
		BaseID:    10572,
		DefaultID: 10583,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	920: { // minecraft:mud_brick_wall
		BaseID:    21795,
		DefaultID: 21798,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	381: { // minecraft:mud_bricks
		BaseID:    9493,
		DefaultID: 9493,
	},
	62: { // minecraft:muddy_mangrove_roots
		BaseID:    171,
		DefaultID: 172,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	390: { // minecraft:mushroom_stem
		BaseID:    9628,
		DefaultID: 9628,
		Properties: []blockProperty{
			{"down", []string{"true", "false"}},
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"up", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	423: { // minecraft:mycelium
		BaseID:    10652,
		DefaultID: 10653,
		Properties: []blockProperty{
			{"snowy", []string{"true", "false"}},
		},
	},
	432: { // minecraft:nether_brick_fence
		BaseID:    11069,
		DefaultID: 11100,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	673: { // minecraft:nether_brick_slab
		BaseID:    15297,
		DefaultID: 15300,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	433: { // minecraft:nether_brick_stairs
		BaseID:    11101,
		DefaultID: 11112,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	921: { // minecraft:nether_brick_wall
		BaseID:    22119,
		DefaultID: 22122,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	431: { // minecraft:nether_bricks
		BaseID:    11068,
		DefaultID: 11068,
	},
	50: { // minecraft:nether_gold_ore
		BaseID:    138,
		DefaultID: 138,
	},
	344: { // minecraft:nether_portal
		BaseID:    8687,
		DefaultID: 8687,
		Properties: []blockProperty{
			{"axis", []string{"x", "z"}},
		},
	},
	529: { // minecraft:nether_quartz_ore
		BaseID:    13079,
		DefaultID: 13079,
	},
	960: { // minecraft:nether_sprouts
		BaseID:    24318,
		DefaultID: 24318,
	},
	434: { // minecraft:nether_wart
		BaseID:    11181,
		DefaultID: 11181,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3"}},
		},
	},
	730: { // minecraft:nether_wart_block
		BaseID:    16827,
		DefaultID: 16827,
	},
	1005: { // minecraft:netherite_block
		BaseID:    25175,
		DefaultID: 25175,
	},
	334: { // minecraft:netherrack
		BaseID:    8667,
		DefaultID: 8667,
	},
	118: { // minecraft:note_block
		BaseID:    680,
		DefaultID: 681,
		Properties: []blockProperty{
			{"instrument", []string{"harp", "basedrum", "snare", "hat", "bass", "flute", "bell", "guitar", "chime", "xylophone", "iron_xylophone", "cow_bell", "didgeridoo", "bit", "banjo", "pling", "trumpet", "trumpet_exposed", "trumpet_oxidized", "trumpet_weathered", "zombie", "skeleton", "creeper", "dragon", "wither_skeleton", "piglin", "custom_head"}},
			{"note", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24"}},
			{"powered", []string{"true", "false"}},
		},
	},
	495: { // minecraft:oak_button
		BaseID:    12418,
		DefaultID: 12427,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	265: { // minecraft:oak_door
		BaseID:    7243,
		DefaultID: 7254,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	333: { // minecraft:oak_fence
		BaseID:    8635,
		DefaultID: 8666,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	419: { // minecraft:oak_fence_gate
		BaseID:    10380,
		DefaultID: 10387,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	280: { // minecraft:oak_hanging_sign
		BaseID:    7503,
		DefaultID: 7552,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	94: { // minecraft:oak_leaves
		BaseID:    267,
		DefaultID: 294,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	51: { // minecraft:oak_log
		BaseID:    139,
		DefaultID: 140,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	13: { // minecraft:oak_planks
		BaseID:    15,
		DefaultID: 15,
	},
	309: { // minecraft:oak_pressure_plate
		BaseID:    8529,
		DefaultID: 8530,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	26: { // minecraft:oak_sapling
		BaseID:    30,
		DefaultID: 30,
		Properties: []blockProperty{
			{"stage", []string{"0", "1"}},
		},
	},
	231: { // minecraft:oak_shelf
		BaseID:    4604,
		DefaultID: 4613,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	254: { // minecraft:oak_sign
		BaseID:    6891,
		DefaultID: 6908,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	653: { // minecraft:oak_slab
		BaseID:    15177,
		DefaultID: 15180,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	244: { // minecraft:oak_stairs
		BaseID:    5463,
		DefaultID: 5474,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	365: { // minecraft:oak_trapdoor
		BaseID:    8784,
		DefaultID: 8799,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	293: { // minecraft:oak_wall_hanging_sign
		BaseID:    8335,
		DefaultID: 8336,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	269: { // minecraft:oak_wall_sign
		BaseID:    7415,
		DefaultID: 7416,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	75: { // minecraft:oak_wood
		BaseID:    210,
		DefaultID: 211,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	734: { // minecraft:observer
		BaseID:    16833,
		DefaultID: 16838,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"powered", []string{"true", "false"}},
		},
	},
	237: { // minecraft:obsidian
		BaseID:    4925,
		DefaultID: 4925,
	},
	1268: { // minecraft:ochre_froglight
		BaseID:    35432,
		DefaultID: 35433,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	1281: { // minecraft:open_eyeblossom
		BaseID:    35718,
		DefaultID: 35718,
	},
	618: { // minecraft:orange_banner
		BaseID:    14790,
		DefaultID: 14798,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	120: { // minecraft:orange_bed
		BaseID:    2046,
		DefaultID: 2049,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1036: { // minecraft:orange_candle
		BaseID:    26485,
		DefaultID: 26488,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1053: { // minecraft:orange_candle_cake
		BaseID:    26729,
		DefaultID: 26730,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	593: { // minecraft:orange_carpet
		BaseID:    14744,
		DefaultID: 14744,
	},
	769: { // minecraft:orange_concrete
		BaseID:    17012,
		DefaultID: 17012,
	},
	817: { // minecraft:orange_concrete_powder
		BaseID:    18404,
		DefaultID: 18404,
	},
	801: { // minecraft:orange_concrete_slab
		BaseID:    18313,
		DefaultID: 18316,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	785: { // minecraft:orange_concrete_stairs
		BaseID:    17107,
		DefaultID: 17118,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	753: { // minecraft:orange_glazed_terracotta
		BaseID:    16951,
		DefaultID: 16951,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	104: { // minecraft:orange_poplar_leaves
		BaseID:    547,
		DefaultID: 574,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	737: { // minecraft:orange_shulker_box
		BaseID:    16857,
		DefaultID: 16861,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	350: { // minecraft:orange_stained_glass
		BaseID:    8769,
		DefaultID: 8769,
	},
	554: { // minecraft:orange_stained_glass_pane
		BaseID:    13259,
		DefaultID: 13290,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	538: { // minecraft:orange_terracotta
		BaseID:    13212,
		DefaultID: 13212,
	},
	208: { // minecraft:orange_tulip
		BaseID:    3821,
		DefaultID: 3821,
	},
	634: { // minecraft:orange_wall_banner
		BaseID:    15034,
		DefaultID: 15034,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	152: { // minecraft:orange_wool
		BaseID:    2410,
		DefaultID: 2410,
	},
	184: { // minecraft:orange_wool_slab
		BaseID:    3711,
		DefaultID: 3714,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	168: { // minecraft:orange_wool_stairs
		BaseID:    2505,
		DefaultID: 2516,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	211: { // minecraft:oxeye_daisy
		BaseID:    3824,
		DefaultID: 3824,
	},
	1145: { // minecraft:oxidized_chiseled_copper
		BaseID:    31160,
		DefaultID: 31160,
	},
	1127: { // minecraft:oxidized_copper
		BaseID:    31142,
		DefaultID: 31142,
	},
	395: { // minecraft:oxidized_copper_bars
		BaseID:    9820,
		DefaultID: 9851,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1193: { // minecraft:oxidized_copper_bulb
		BaseID:    32905,
		DefaultID: 32908,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	404: { // minecraft:oxidized_copper_chain
		BaseID:    10004,
		DefaultID: 10007,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1201: { // minecraft:oxidized_copper_chest
		BaseID:    32997,
		DefaultID: 32998,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1169: { // minecraft:oxidized_copper_door
		BaseID:    32045,
		DefaultID: 32056,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1209: { // minecraft:oxidized_copper_golem_statue
		BaseID:    33213,
		DefaultID: 33214,
		Properties: []blockProperty{
			{"copper_golem_pose", []string{"standing", "sitting", "running", "star"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1185: { // minecraft:oxidized_copper_grate
		BaseID:    32883,
		DefaultID: 32884,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	944: { // minecraft:oxidized_copper_lantern
		BaseID:    24214,
		DefaultID: 24217,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1177: { // minecraft:oxidized_copper_trapdoor
		BaseID:    32557,
		DefaultID: 32572,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1137: { // minecraft:oxidized_cut_copper
		BaseID:    31152,
		DefaultID: 31152,
	},
	1161: { // minecraft:oxidized_cut_copper_slab
		BaseID:    31823,
		DefaultID: 31826,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1153: { // minecraft:oxidized_cut_copper_stairs
		BaseID:    31405,
		DefaultID: 31416,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1217: { // minecraft:oxidized_lightning_rod
		BaseID:    33445,
		DefaultID: 33464,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	610: { // minecraft:packed_ice
		BaseID:    14761,
		DefaultID: 14761,
	},
	380: { // minecraft:packed_mud
		BaseID:    9492,
		DefaultID: 9492,
	},
	1280: { // minecraft:pale_hanging_moss
		BaseID:    35716,
		DefaultID: 35716,
		Properties: []blockProperty{
			{"tip", []string{"true", "false"}},
		},
	},
	1278: { // minecraft:pale_moss_block
		BaseID:    35553,
		DefaultID: 35553,
	},
	1279: { // minecraft:pale_moss_carpet
		BaseID:    35554,
		DefaultID: 35554,
		Properties: []blockProperty{
			{"bottom", []string{"true", "false"}},
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	502: { // minecraft:pale_oak_button
		BaseID:    12586,
		DefaultID: 12595,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	708: { // minecraft:pale_oak_door
		BaseID:    16355,
		DefaultID: 16366,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	698: { // minecraft:pale_oak_fence
		BaseID:    15843,
		DefaultID: 15874,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	688: { // minecraft:pale_oak_fence_gate
		BaseID:    15523,
		DefaultID: 15530,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	287: { // minecraft:pale_oak_hanging_sign
		BaseID:    7951,
		DefaultID: 8000,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	101: { // minecraft:pale_oak_leaves
		BaseID:    463,
		DefaultID: 490,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	58: { // minecraft:pale_oak_log
		BaseID:    160,
		DefaultID: 161,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	21: { // minecraft:pale_oak_planks
		BaseID:    25,
		DefaultID: 25,
	},
	316: { // minecraft:pale_oak_pressure_plate
		BaseID:    8543,
		DefaultID: 8544,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	33: { // minecraft:pale_oak_sapling
		BaseID:    44,
		DefaultID: 44,
		Properties: []blockProperty{
			{"stage", []string{"0", "1"}},
		},
	},
	232: { // minecraft:pale_oak_shelf
		BaseID:    4668,
		DefaultID: 4677,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	261: { // minecraft:pale_oak_sign
		BaseID:    7115,
		DefaultID: 7132,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	660: { // minecraft:pale_oak_slab
		BaseID:    15219,
		DefaultID: 15222,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	572: { // minecraft:pale_oak_stairs
		BaseID:    13979,
		DefaultID: 13990,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	372: { // minecraft:pale_oak_trapdoor
		BaseID:    9232,
		DefaultID: 9247,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	300: { // minecraft:pale_oak_wall_hanging_sign
		BaseID:    8391,
		DefaultID: 8392,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	276: { // minecraft:pale_oak_wall_sign
		BaseID:    7471,
		DefaultID: 7472,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	20: { // minecraft:pale_oak_wood
		BaseID:    22,
		DefaultID: 23,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	1270: { // minecraft:pearlescent_froglight
		BaseID:    35438,
		DefaultID: 35439,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	614: { // minecraft:peony
		BaseID:    14768,
		DefaultID: 14769,
		Properties: []blockProperty{
			{"half", []string{"upper", "lower"}},
		},
	},
	668: { // minecraft:petrified_oak_slab
		BaseID:    15267,
		DefaultID: 15270,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	518: { // minecraft:piglin_head
		BaseID:    12922,
		DefaultID: 12938,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	519: { // minecraft:piglin_wall_head
		BaseID:    12954,
		DefaultID: 12955,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	623: { // minecraft:pink_banner
		BaseID:    14870,
		DefaultID: 14878,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	125: { // minecraft:pink_bed
		BaseID:    2126,
		DefaultID: 2129,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1041: { // minecraft:pink_candle
		BaseID:    26565,
		DefaultID: 26568,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1058: { // minecraft:pink_candle_cake
		BaseID:    26739,
		DefaultID: 26740,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	598: { // minecraft:pink_carpet
		BaseID:    14749,
		DefaultID: 14749,
	},
	774: { // minecraft:pink_concrete
		BaseID:    17017,
		DefaultID: 17017,
	},
	822: { // minecraft:pink_concrete_powder
		BaseID:    18409,
		DefaultID: 18409,
	},
	806: { // minecraft:pink_concrete_slab
		BaseID:    18343,
		DefaultID: 18346,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	790: { // minecraft:pink_concrete_stairs
		BaseID:    17507,
		DefaultID: 17518,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	758: { // minecraft:pink_glazed_terracotta
		BaseID:    16971,
		DefaultID: 16971,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	1231: { // minecraft:pink_petals
		BaseID:    33664,
		DefaultID: 33664,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"flower_amount", []string{"1", "2", "3", "4"}},
		},
	},
	742: { // minecraft:pink_shulker_box
		BaseID:    16887,
		DefaultID: 16891,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	355: { // minecraft:pink_stained_glass
		BaseID:    8774,
		DefaultID: 8774,
	},
	559: { // minecraft:pink_stained_glass_pane
		BaseID:    13419,
		DefaultID: 13450,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	543: { // minecraft:pink_terracotta
		BaseID:    13217,
		DefaultID: 13217,
	},
	210: { // minecraft:pink_tulip
		BaseID:    3823,
		DefaultID: 3823,
	},
	639: { // minecraft:pink_wall_banner
		BaseID:    15054,
		DefaultID: 15054,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	157: { // minecraft:pink_wool
		BaseID:    2415,
		DefaultID: 2415,
	},
	189: { // minecraft:pink_wool_slab
		BaseID:    3741,
		DefaultID: 3744,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	173: { // minecraft:pink_wool_stairs
		BaseID:    2905,
		DefaultID: 2916,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	149: { // minecraft:piston
		BaseID:    2373,
		DefaultID: 2379,
		Properties: []blockProperty{
			{"extended", []string{"true", "false"}},
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	150: { // minecraft:piston_head
		BaseID:    2385,
		DefaultID: 2387,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"short", []string{"true", "false"}},
			{"type", []string{"normal", "sticky"}},
		},
	},
	721: { // minecraft:pitcher_crop
		BaseID:    16780,
		DefaultID: 16781,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4"}},
			{"half", []string{"upper", "lower"}},
		},
	},
	722: { // minecraft:pitcher_plant
		BaseID:    16790,
		DefaultID: 16791,
		Properties: []blockProperty{
			{"half", []string{"upper", "lower"}},
		},
	},
	512: { // minecraft:player_head
		BaseID:    12802,
		DefaultID: 12818,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	513: { // minecraft:player_wall_head
		BaseID:    12834,
		DefaultID: 12835,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	11: { // minecraft:podzol
		BaseID:    12,
		DefaultID: 13,
		Properties: []blockProperty{
			{"snowy", []string{"true", "false"}},
		},
	},
	1223: { // minecraft:pointed_dripstone
		BaseID:    33566,
		DefaultID: 33571,
		Properties: []blockProperty{
			{"thickness", []string{"tip_merge", "tip", "frustum", "middle", "base"}},
			{"vertical_direction", []string{"up", "down"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	7: { // minecraft:polished_andesite
		BaseID:    7,
		DefaultID: 7,
	},
	912: { // minecraft:polished_andesite_slab
		BaseID:    19839,
		DefaultID: 19842,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	899: { // minecraft:polished_andesite_stairs
		BaseID:    19613,
		DefaultID: 19624,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	338: { // minecraft:polished_basalt
		BaseID:    8673,
		DefaultID: 8674,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	1018: { // minecraft:polished_blackstone
		BaseID:    25599,
		DefaultID: 25599,
	},
	1022: { // minecraft:polished_blackstone_brick_slab
		BaseID:    25603,
		DefaultID: 25606,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1023: { // minecraft:polished_blackstone_brick_stairs
		BaseID:    25609,
		DefaultID: 25620,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1024: { // minecraft:polished_blackstone_brick_wall
		BaseID:    25689,
		DefaultID: 25692,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	1019: { // minecraft:polished_blackstone_bricks
		BaseID:    25600,
		DefaultID: 25600,
	},
	1029: { // minecraft:polished_blackstone_button
		BaseID:    26102,
		DefaultID: 26111,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1028: { // minecraft:polished_blackstone_pressure_plate
		BaseID:    26100,
		DefaultID: 26101,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	1027: { // minecraft:polished_blackstone_slab
		BaseID:    26094,
		DefaultID: 26097,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1026: { // minecraft:polished_blackstone_stairs
		BaseID:    26014,
		DefaultID: 26025,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1030: { // minecraft:polished_blackstone_wall
		BaseID:    26126,
		DefaultID: 26129,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	1106: { // minecraft:polished_cinnabar
		BaseID:    29694,
		DefaultID: 29694,
	},
	1107: { // minecraft:polished_cinnabar_slab
		BaseID:    29695,
		DefaultID: 29698,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1108: { // minecraft:polished_cinnabar_stairs
		BaseID:    29701,
		DefaultID: 29712,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1109: { // minecraft:polished_cinnabar_wall
		BaseID:    29781,
		DefaultID: 29784,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	1246: { // minecraft:polished_deepslate
		BaseID:    34187,
		DefaultID: 34187,
	},
	1248: { // minecraft:polished_deepslate_slab
		BaseID:    34268,
		DefaultID: 34271,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1247: { // minecraft:polished_deepslate_stairs
		BaseID:    34188,
		DefaultID: 34199,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1249: { // minecraft:polished_deepslate_wall
		BaseID:    34274,
		DefaultID: 34277,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	5: { // minecraft:polished_diorite
		BaseID:    5,
		DefaultID: 5,
	},
	904: { // minecraft:polished_diorite_slab
		BaseID:    19791,
		DefaultID: 19794,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	890: { // minecraft:polished_diorite_stairs
		BaseID:    18893,
		DefaultID: 18904,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	3: { // minecraft:polished_granite
		BaseID:    3,
		DefaultID: 3,
	},
	901: { // minecraft:polished_granite_slab
		BaseID:    19773,
		DefaultID: 19776,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	887: { // minecraft:polished_granite_stairs
		BaseID:    18653,
		DefaultID: 18664,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1093: { // minecraft:polished_sulfur
		BaseID:    28460,
		DefaultID: 28460,
	},
	1094: { // minecraft:polished_sulfur_slab
		BaseID:    28461,
		DefaultID: 28464,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1095: { // minecraft:polished_sulfur_stairs
		BaseID:    28467,
		DefaultID: 28478,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1096: { // minecraft:polished_sulfur_wall
		BaseID:    28547,
		DefaultID: 28550,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	1078: { // minecraft:polished_tuff
		BaseID:    27220,
		DefaultID: 27220,
	},
	1079: { // minecraft:polished_tuff_slab
		BaseID:    27221,
		DefaultID: 27224,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1080: { // minecraft:polished_tuff_stairs
		BaseID:    27227,
		DefaultID: 27238,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1081: { // minecraft:polished_tuff_wall
		BaseID:    27307,
		DefaultID: 27310,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	504: { // minecraft:poplar_button
		BaseID:    12634,
		DefaultID: 12643,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	710: { // minecraft:poplar_door
		BaseID:    16483,
		DefaultID: 16494,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	700: { // minecraft:poplar_fence
		BaseID:    15907,
		DefaultID: 15938,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	690: { // minecraft:poplar_fence_gate
		BaseID:    15587,
		DefaultID: 15594,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	291: { // minecraft:poplar_hanging_sign
		BaseID:    8207,
		DefaultID: 8256,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	60: { // minecraft:poplar_log
		BaseID:    166,
		DefaultID: 167,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	23: { // minecraft:poplar_planks
		BaseID:    27,
		DefaultID: 27,
	},
	318: { // minecraft:poplar_pressure_plate
		BaseID:    8547,
		DefaultID: 8548,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	35: { // minecraft:poplar_sapling
		BaseID:    86,
		DefaultID: 86,
		Properties: []blockProperty{
			{"stage", []string{"0", "1"}},
		},
	},
	233: { // minecraft:poplar_shelf
		BaseID:    4732,
		DefaultID: 4741,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	263: { // minecraft:poplar_sign
		BaseID:    7179,
		DefaultID: 7196,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	662: { // minecraft:poplar_slab
		BaseID:    15231,
		DefaultID: 15234,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	574: { // minecraft:poplar_stairs
		BaseID:    14139,
		DefaultID: 14150,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	374: { // minecraft:poplar_trapdoor
		BaseID:    9360,
		DefaultID: 9375,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	302: { // minecraft:poplar_wall_hanging_sign
		BaseID:    8407,
		DefaultID: 8408,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	278: { // minecraft:poplar_wall_sign
		BaseID:    7487,
		DefaultID: 7488,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	83: { // minecraft:poplar_wood
		BaseID:    234,
		DefaultID: 235,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	203: { // minecraft:poppy
		BaseID:    3816,
		DefaultID: 3816,
	},
	494: { // minecraft:potatoes
		BaseID:    12410,
		DefaultID: 12410,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7"}},
		},
	},
	1089: { // minecraft:potent_sulfur
		BaseID:    28045,
		DefaultID: 28045,
		Properties: []blockProperty{
			{"potent_sulfur_state", []string{"dry", "wet", "dormant", "erupting", "continuous"}},
		},
	},
	468: { // minecraft:potted_acacia_sapling
		BaseID:    12377,
		DefaultID: 12377,
	},
	479: { // minecraft:potted_allium
		BaseID:    12388,
		DefaultID: 12388,
	},
	1266: { // minecraft:potted_azalea_bush
		BaseID:    35430,
		DefaultID: 35430,
	},
	480: { // minecraft:potted_azure_bluet
		BaseID:    12389,
		DefaultID: 12389,
	},
	883: { // minecraft:potted_bamboo
		BaseID:    18648,
		DefaultID: 18648,
	},
	466: { // minecraft:potted_birch_sapling
		BaseID:    12375,
		DefaultID: 12375,
	},
	478: { // minecraft:potted_blue_orchid
		BaseID:    12387,
		DefaultID: 12387,
	},
	490: { // minecraft:potted_brown_mushroom
		BaseID:    12399,
		DefaultID: 12399,
	},
	492: { // minecraft:potted_cactus
		BaseID:    12401,
		DefaultID: 12401,
	},
	469: { // minecraft:potted_cherry_sapling
		BaseID:    12378,
		DefaultID: 12378,
	},
	1284: { // minecraft:potted_closed_eyeblossom
		BaseID:    35721,
		DefaultID: 35721,
	},
	486: { // minecraft:potted_cornflower
		BaseID:    12395,
		DefaultID: 12395,
	},
	1009: { // minecraft:potted_crimson_fungus
		BaseID:    25183,
		DefaultID: 25183,
	},
	1011: { // minecraft:potted_crimson_roots
		BaseID:    25185,
		DefaultID: 25185,
	},
	475: { // minecraft:potted_dandelion
		BaseID:    12384,
		DefaultID: 12384,
	},
	470: { // minecraft:potted_dark_oak_sapling
		BaseID:    12379,
		DefaultID: 12379,
	},
	491: { // minecraft:potted_dead_bush
		BaseID:    12400,
		DefaultID: 12400,
	},
	474: { // minecraft:potted_fern
		BaseID:    12383,
		DefaultID: 12383,
	},
	1267: { // minecraft:potted_flowering_azalea_bush
		BaseID:    35431,
		DefaultID: 35431,
	},
	476: { // minecraft:potted_golden_dandelion
		BaseID:    12385,
		DefaultID: 12385,
	},
	467: { // minecraft:potted_jungle_sapling
		BaseID:    12376,
		DefaultID: 12376,
	},
	487: { // minecraft:potted_lily_of_the_valley
		BaseID:    12396,
		DefaultID: 12396,
	},
	473: { // minecraft:potted_mangrove_propagule
		BaseID:    12382,
		DefaultID: 12382,
	},
	464: { // minecraft:potted_oak_sapling
		BaseID:    12373,
		DefaultID: 12373,
	},
	1283: { // minecraft:potted_open_eyeblossom
		BaseID:    35720,
		DefaultID: 35720,
	},
	482: { // minecraft:potted_orange_tulip
		BaseID:    12391,
		DefaultID: 12391,
	},
	485: { // minecraft:potted_oxeye_daisy
		BaseID:    12394,
		DefaultID: 12394,
	},
	471: { // minecraft:potted_pale_oak_sapling
		BaseID:    12380,
		DefaultID: 12380,
	},
	484: { // minecraft:potted_pink_tulip
		BaseID:    12393,
		DefaultID: 12393,
	},
	472: { // minecraft:potted_poplar_sapling
		BaseID:    12381,
		DefaultID: 12381,
	},
	477: { // minecraft:potted_poppy
		BaseID:    12386,
		DefaultID: 12386,
	},
	489: { // minecraft:potted_red_mushroom
		BaseID:    12398,
		DefaultID: 12398,
	},
	481: { // minecraft:potted_red_tulip
		BaseID:    12390,
		DefaultID: 12390,
	},
	465: { // minecraft:potted_spruce_sapling
		BaseID:    12374,
		DefaultID: 12374,
	},
	463: { // minecraft:potted_torchflower
		BaseID:    12372,
		DefaultID: 12372,
	},
	1010: { // minecraft:potted_warped_fungus
		BaseID:    25184,
		DefaultID: 25184,
	},
	1012: { // minecraft:potted_warped_roots
		BaseID:    25186,
		DefaultID: 25186,
	},
	483: { // minecraft:potted_white_tulip
		BaseID:    12392,
		DefaultID: 12392,
	},
	488: { // minecraft:potted_wither_rose
		BaseID:    12397,
		DefaultID: 12397,
	},
	1117: { // minecraft:powder_snow
		BaseID:    30519,
		DefaultID: 30519,
	},
	440: { // minecraft:powder_snow_cauldron
		BaseID:    11199,
		DefaultID: 11199,
		Properties: []blockProperty{
			{"level", []string{"1", "2", "3"}},
		},
	},
	136: { // minecraft:powered_rail
		BaseID:    2302,
		DefaultID: 2315,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"shape", []string{"north_south", "east_west", "ascending_east", "ascending_west", "ascending_north", "ascending_south"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	581: { // minecraft:prismarine
		BaseID:    14478,
		DefaultID: 14478,
	},
	588: { // minecraft:prismarine_brick_slab
		BaseID:    14727,
		DefaultID: 14730,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	585: { // minecraft:prismarine_brick_stairs
		BaseID:    14561,
		DefaultID: 14572,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	582: { // minecraft:prismarine_bricks
		BaseID:    14479,
		DefaultID: 14479,
	},
	587: { // minecraft:prismarine_slab
		BaseID:    14721,
		DefaultID: 14724,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	584: { // minecraft:prismarine_stairs
		BaseID:    14481,
		DefaultID: 14492,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	915: { // minecraft:prismarine_wall
		BaseID:    20175,
		DefaultID: 20178,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	410: { // minecraft:pumpkin
		BaseID:    10066,
		DefaultID: 10066,
	},
	414: { // minecraft:pumpkin_stem
		BaseID:    10076,
		DefaultID: 10076,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7"}},
		},
	},
	627: { // minecraft:purple_banner
		BaseID:    14934,
		DefaultID: 14942,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	129: { // minecraft:purple_bed
		BaseID:    2190,
		DefaultID: 2193,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1045: { // minecraft:purple_candle
		BaseID:    26629,
		DefaultID: 26632,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1062: { // minecraft:purple_candle_cake
		BaseID:    26747,
		DefaultID: 26748,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	602: { // minecraft:purple_carpet
		BaseID:    14753,
		DefaultID: 14753,
	},
	778: { // minecraft:purple_concrete
		BaseID:    17021,
		DefaultID: 17021,
	},
	826: { // minecraft:purple_concrete_powder
		BaseID:    18413,
		DefaultID: 18413,
	},
	810: { // minecraft:purple_concrete_slab
		BaseID:    18367,
		DefaultID: 18370,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	794: { // minecraft:purple_concrete_stairs
		BaseID:    17827,
		DefaultID: 17838,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	762: { // minecraft:purple_glazed_terracotta
		BaseID:    16987,
		DefaultID: 16987,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	746: { // minecraft:purple_shulker_box
		BaseID:    16911,
		DefaultID: 16915,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	359: { // minecraft:purple_stained_glass
		BaseID:    8778,
		DefaultID: 8778,
	},
	563: { // minecraft:purple_stained_glass_pane
		BaseID:    13547,
		DefaultID: 13578,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	547: { // minecraft:purple_terracotta
		BaseID:    13221,
		DefaultID: 13221,
	},
	643: { // minecraft:purple_wall_banner
		BaseID:    15070,
		DefaultID: 15070,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	161: { // minecraft:purple_wool
		BaseID:    2419,
		DefaultID: 2419,
	},
	193: { // minecraft:purple_wool_slab
		BaseID:    3765,
		DefaultID: 3768,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	177: { // minecraft:purple_wool_stairs
		BaseID:    3225,
		DefaultID: 3236,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	715: { // minecraft:purpur_block
		BaseID:    16687,
		DefaultID: 16687,
	},
	717: { // minecraft:purpur_pillar
		BaseID:    16694,
		DefaultID: 16695,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	716: { // minecraft:purpur_slab
		BaseID:    16688,
		DefaultID: 16691,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	718: { // minecraft:purpur_stairs
		BaseID:    16697,
		DefaultID: 16708,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	531: { // minecraft:quartz_block
		BaseID:    13090,
		DefaultID: 13090,
	},
	1033: { // minecraft:quartz_bricks
		BaseID:    26452,
		DefaultID: 26452,
	},
	533: { // minecraft:quartz_pillar
		BaseID:    13092,
		DefaultID: 13093,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	674: { // minecraft:quartz_slab
		BaseID:    15303,
		DefaultID: 15306,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	534: { // minecraft:quartz_stairs
		BaseID:    13095,
		DefaultID: 13106,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	267: { // minecraft:rail
		BaseID:    7315,
		DefaultID: 7316,
		Properties: []blockProperty{
			{"shape", []string{"north_south", "east_west", "ascending_east", "ascending_west", "ascending_north", "ascending_south", "south_east", "south_west", "north_west", "north_east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1264: { // minecraft:raw_copper_block
		BaseID:    35428,
		DefaultID: 35428,
	},
	1265: { // minecraft:raw_gold_block
		BaseID:    35429,
		DefaultID: 35429,
	},
	1263: { // minecraft:raw_iron_block
		BaseID:    35427,
		DefaultID: 35427,
	},
	631: { // minecraft:red_banner
		BaseID:    14998,
		DefaultID: 15006,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	133: { // minecraft:red_bed
		BaseID:    2254,
		DefaultID: 2257,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1049: { // minecraft:red_candle
		BaseID:    26693,
		DefaultID: 26696,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1066: { // minecraft:red_candle_cake
		BaseID:    26755,
		DefaultID: 26756,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	606: { // minecraft:red_carpet
		BaseID:    14757,
		DefaultID: 14757,
	},
	782: { // minecraft:red_concrete
		BaseID:    17025,
		DefaultID: 17025,
	},
	830: { // minecraft:red_concrete_powder
		BaseID:    18417,
		DefaultID: 18417,
	},
	814: { // minecraft:red_concrete_slab
		BaseID:    18391,
		DefaultID: 18394,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	798: { // minecraft:red_concrete_stairs
		BaseID:    18147,
		DefaultID: 18158,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	766: { // minecraft:red_glazed_terracotta
		BaseID:    17003,
		DefaultID: 17003,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	216: { // minecraft:red_mushroom
		BaseID:    3829,
		DefaultID: 3829,
	},
	389: { // minecraft:red_mushroom_block
		BaseID:    9564,
		DefaultID: 9564,
		Properties: []blockProperty{
			{"down", []string{"true", "false"}},
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"up", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	911: { // minecraft:red_nether_brick_slab
		BaseID:    19833,
		DefaultID: 19836,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	898: { // minecraft:red_nether_brick_stairs
		BaseID:    19533,
		DefaultID: 19544,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	923: { // minecraft:red_nether_brick_wall
		BaseID:    22767,
		DefaultID: 22770,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	731: { // minecraft:red_nether_bricks
		BaseID:    16828,
		DefaultID: 16828,
	},
	103: { // minecraft:red_poplar_leaves
		BaseID:    519,
		DefaultID: 546,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	41: { // minecraft:red_sand
		BaseID:    126,
		DefaultID: 126,
	},
	649: { // minecraft:red_sandstone
		BaseID:    15094,
		DefaultID: 15094,
	},
	675: { // minecraft:red_sandstone_slab
		BaseID:    15309,
		DefaultID: 15312,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	652: { // minecraft:red_sandstone_stairs
		BaseID:    15097,
		DefaultID: 15108,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	916: { // minecraft:red_sandstone_wall
		BaseID:    20499,
		DefaultID: 20502,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	144: { // minecraft:red_shrub
		BaseID:    2367,
		DefaultID: 2367,
	},
	750: { // minecraft:red_shulker_box
		BaseID:    16935,
		DefaultID: 16939,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	363: { // minecraft:red_stained_glass
		BaseID:    8782,
		DefaultID: 8782,
	},
	567: { // minecraft:red_stained_glass_pane
		BaseID:    13675,
		DefaultID: 13706,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	551: { // minecraft:red_terracotta
		BaseID:    13225,
		DefaultID: 13225,
	},
	207: { // minecraft:red_tulip
		BaseID:    3820,
		DefaultID: 3820,
	},
	647: { // minecraft:red_wall_banner
		BaseID:    15086,
		DefaultID: 15086,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	165: { // minecraft:red_wool
		BaseID:    2423,
		DefaultID: 2423,
	},
	197: { // minecraft:red_wool_slab
		BaseID:    3789,
		DefaultID: 3792,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	181: { // minecraft:red_wool_stairs
		BaseID:    3545,
		DefaultID: 3556,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	528: { // minecraft:redstone_block
		BaseID:    13078,
		DefaultID: 13078,
	},
	445: { // minecraft:redstone_lamp
		BaseID:    11213,
		DefaultID: 11214,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	320: { // minecraft:redstone_ore
		BaseID:    8551,
		DefaultID: 8552,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	322: { // minecraft:redstone_torch
		BaseID:    8555,
		DefaultID: 8555,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	323: { // minecraft:redstone_wall_torch
		BaseID:    8557,
		DefaultID: 8557,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"lit", []string{"true", "false"}},
		},
	},
	246: { // minecraft:redstone_wire
		BaseID:    5567,
		DefaultID: 6727,
		Properties: []blockProperty{
			{"east", []string{"up", "side", "none"}},
			{"north", []string{"up", "side", "none"}},
			{"power", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"south", []string{"up", "side", "none"}},
			{"west", []string{"up", "side", "none"}},
		},
	},
	1272: { // minecraft:reinforced_deepslate
		BaseID:    35442,
		DefaultID: 35442,
	},
	348: { // minecraft:repeater
		BaseID:    8704,
		DefaultID: 8707,
		Properties: []blockProperty{
			{"delay", []string{"1", "2", "3", "4"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"locked", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	726: { // minecraft:repeating_command_block
		BaseID:    16798,
		DefaultID: 16804,
		Properties: []blockProperty{
			{"conditional", []string{"true", "false"}},
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	425: { // minecraft:resin_block
		BaseID:    10655,
		DefaultID: 10655,
	},
	428: { // minecraft:resin_brick_slab
		BaseID:    10737,
		DefaultID: 10740,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	427: { // minecraft:resin_brick_stairs
		BaseID:    10657,
		DefaultID: 10668,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	429: { // minecraft:resin_brick_wall
		BaseID:    10743,
		DefaultID: 10746,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	426: { // minecraft:resin_bricks
		BaseID:    10656,
		DefaultID: 10656,
	},
	418: { // minecraft:resin_clump
		BaseID:    10252,
		DefaultID: 10379,
		Properties: []blockProperty{
			{"down", []string{"true", "false"}},
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1008: { // minecraft:respawn_anchor
		BaseID:    25178,
		DefaultID: 25178,
		Properties: []blockProperty{
			{"charges", []string{"0", "1", "2", "3", "4"}},
		},
	},
	1239: { // minecraft:rooted_dirt
		BaseID:    33771,
		DefaultID: 33771,
	},
	613: { // minecraft:rose_bush
		BaseID:    14766,
		DefaultID: 14767,
		Properties: []blockProperty{
			{"half", []string{"upper", "lower"}},
		},
	},
	39: { // minecraft:sand
		BaseID:    121,
		DefaultID: 121,
	},
	115: { // minecraft:sandstone
		BaseID:    677,
		DefaultID: 677,
	},
	666: { // minecraft:sandstone_slab
		BaseID:    15255,
		DefaultID: 15258,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	448: { // minecraft:sandstone_stairs
		BaseID:    11235,
		DefaultID: 11246,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	924: { // minecraft:sandstone_wall
		BaseID:    23091,
		DefaultID: 23094,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	927: { // minecraft:scaffolding
		BaseID:    24063,
		DefaultID: 24094,
		Properties: []blockProperty{
			{"bottom", []string{"true", "false"}},
			{"distance", []string{"0", "1", "2", "3", "4", "5", "6", "7"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1120: { // minecraft:sculk
		BaseID:    31000,
		DefaultID: 31000,
	},
	1122: { // minecraft:sculk_catalyst
		BaseID:    31129,
		DefaultID: 31130,
		Properties: []blockProperty{
			{"bloom", []string{"true", "false"}},
		},
	},
	1118: { // minecraft:sculk_sensor
		BaseID:    30520,
		DefaultID: 30521,
		Properties: []blockProperty{
			{"power", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"sculk_sensor_phase", []string{"inactive", "active", "cooldown"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1123: { // minecraft:sculk_shrieker
		BaseID:    31131,
		DefaultID: 31138,
		Properties: []blockProperty{
			{"can_summon", []string{"true", "false"}},
			{"shrieking", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1121: { // minecraft:sculk_vein
		BaseID:    31001,
		DefaultID: 31128,
		Properties: []blockProperty{
			{"down", []string{"true", "false"}},
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	590: { // minecraft:sea_lantern
		BaseID:    14739,
		DefaultID: 14739,
	},
	878: { // minecraft:sea_pickle
		BaseID:    18624,
		DefaultID: 18624,
		Properties: []blockProperty{
			{"pickles", []string{"1", "2", "3", "4"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	147: { // minecraft:seagrass
		BaseID:    2370,
		DefaultID: 2370,
	},
	447: { // minecraft:shelf_mushroom
		BaseID:    11227,
		DefaultID: 11227,
		Properties: []blockProperty{
			{"age", []string{"0", "1"}},
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	145: { // minecraft:short_dry_grass
		BaseID:    2368,
		DefaultID: 2368,
	},
	140: { // minecraft:short_grass
		BaseID:    2363,
		DefaultID: 2363,
	},
	967: { // minecraft:shroomlight
		BaseID:    24333,
		DefaultID: 24333,
	},
	735: { // minecraft:shulker_box
		BaseID:    16845,
		DefaultID: 16849,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	506: { // minecraft:skeleton_skull
		BaseID:    12682,
		DefaultID: 12698,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	507: { // minecraft:skeleton_wall_skull
		BaseID:    12714,
		DefaultID: 12715,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	577: { // minecraft:slime_block
		BaseID:    14379,
		DefaultID: 14379,
	},
	1073: { // minecraft:small_amethyst_bud
		BaseID:    26797,
		DefaultID: 26806,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1237: { // minecraft:small_dripleaf
		BaseID:    33753,
		DefaultID: 33756,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	936: { // minecraft:smithing_table
		BaseID:    24157,
		DefaultID: 24157,
	},
	930: { // minecraft:smoker
		BaseID:    24111,
		DefaultID: 24112,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"lit", []string{"true", "false"}},
		},
	},
	1262: { // minecraft:smooth_basalt
		BaseID:    35426,
		DefaultID: 35426,
	},
	680: { // minecraft:smooth_quartz
		BaseID:    15329,
		DefaultID: 15329,
	},
	908: { // minecraft:smooth_quartz_slab
		BaseID:    19815,
		DefaultID: 19818,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	895: { // minecraft:smooth_quartz_stairs
		BaseID:    19293,
		DefaultID: 19304,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	681: { // minecraft:smooth_red_sandstone
		BaseID:    15330,
		DefaultID: 15330,
	},
	902: { // minecraft:smooth_red_sandstone_slab
		BaseID:    19779,
		DefaultID: 19782,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	888: { // minecraft:smooth_red_sandstone_stairs
		BaseID:    18733,
		DefaultID: 18744,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	679: { // minecraft:smooth_sandstone
		BaseID:    15328,
		DefaultID: 15328,
	},
	907: { // minecraft:smooth_sandstone_slab
		BaseID:    19809,
		DefaultID: 19812,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	894: { // minecraft:smooth_sandstone_stairs
		BaseID:    19213,
		DefaultID: 19224,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	677: { // minecraft:smooth_stone
		BaseID:    15321,
		DefaultID: 15321,
	},
	678: { // minecraft:smooth_stone_slab
		BaseID:    15322,
		DefaultID: 15325,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	836: { // minecraft:sniffer_egg
		BaseID:    18459,
		DefaultID: 18459,
		Properties: []blockProperty{
			{"hatch", []string{"0", "1", "2"}},
		},
	},
	325: { // minecraft:snow
		BaseID:    8589,
		DefaultID: 8589,
		Properties: []blockProperty{
			{"layers", []string{"1", "2", "3", "4", "5", "6", "7", "8"}},
		},
	},
	327: { // minecraft:snow_block
		BaseID:    8598,
		DefaultID: 8598,
	},
	950: { // minecraft:soul_campfire
		BaseID:    24266,
		DefaultID: 24269,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"lit", []string{"true", "false"}},
			{"signal_fire", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	241: { // minecraft:soul_fire
		BaseID:    5443,
		DefaultID: 5443,
	},
	940: { // minecraft:soul_lantern
		BaseID:    24198,
		DefaultID: 24201,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	335: { // minecraft:soul_sand
		BaseID:    8668,
		DefaultID: 8668,
	},
	336: { // minecraft:soul_soil
		BaseID:    8669,
		DefaultID: 8669,
	},
	339: { // minecraft:soul_torch
		BaseID:    8676,
		DefaultID: 8676,
	},
	340: { // minecraft:soul_wall_torch
		BaseID:    8677,
		DefaultID: 8677,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	242: { // minecraft:spawner
		BaseID:    5444,
		DefaultID: 5444,
	},
	108: { // minecraft:sponge
		BaseID:    659,
		DefaultID: 659,
	},
	1227: { // minecraft:spore_blossom
		BaseID:    33660,
		DefaultID: 33660,
	},
	496: { // minecraft:spruce_button
		BaseID:    12442,
		DefaultID: 12451,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	702: { // minecraft:spruce_door
		BaseID:    15971,
		DefaultID: 15982,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	692: { // minecraft:spruce_fence
		BaseID:    15651,
		DefaultID: 15682,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	682: { // minecraft:spruce_fence_gate
		BaseID:    15331,
		DefaultID: 15338,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	281: { // minecraft:spruce_hanging_sign
		BaseID:    7567,
		DefaultID: 7616,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	95: { // minecraft:spruce_leaves
		BaseID:    295,
		DefaultID: 322,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	52: { // minecraft:spruce_log
		BaseID:    142,
		DefaultID: 143,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	14: { // minecraft:spruce_planks
		BaseID:    16,
		DefaultID: 16,
	},
	310: { // minecraft:spruce_pressure_plate
		BaseID:    8531,
		DefaultID: 8532,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	27: { // minecraft:spruce_sapling
		BaseID:    32,
		DefaultID: 32,
		Properties: []blockProperty{
			{"stage", []string{"0", "1"}},
		},
	},
	234: { // minecraft:spruce_shelf
		BaseID:    4796,
		DefaultID: 4805,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	255: { // minecraft:spruce_sign
		BaseID:    6923,
		DefaultID: 6940,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	654: { // minecraft:spruce_slab
		BaseID:    15183,
		DefaultID: 15186,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	455: { // minecraft:spruce_stairs
		BaseID:    11470,
		DefaultID: 11481,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	366: { // minecraft:spruce_trapdoor
		BaseID:    8848,
		DefaultID: 8863,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	294: { // minecraft:spruce_wall_hanging_sign
		BaseID:    8343,
		DefaultID: 8344,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	270: { // minecraft:spruce_wall_sign
		BaseID:    7423,
		DefaultID: 7424,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	76: { // minecraft:spruce_wood
		BaseID:    213,
		DefaultID: 214,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	138: { // minecraft:sticky_piston
		BaseID:    2350,
		DefaultID: 2356,
		Properties: []blockProperty{
			{"extended", []string{"true", "false"}},
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	1: { // minecraft:stone
		BaseID:    1,
		DefaultID: 1,
	},
	671: { // minecraft:stone_brick_slab
		BaseID:    15285,
		DefaultID: 15288,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	421: { // minecraft:stone_brick_stairs
		BaseID:    10492,
		DefaultID: 10503,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	919: { // minecraft:stone_brick_wall
		BaseID:    21471,
		DefaultID: 21474,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	376: { // minecraft:stone_bricks
		BaseID:    9488,
		DefaultID: 9488,
	},
	324: { // minecraft:stone_button
		BaseID:    8565,
		DefaultID: 8574,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	307: { // minecraft:stone_pressure_plate
		BaseID:    8463,
		DefaultID: 8464,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	665: { // minecraft:stone_slab
		BaseID:    15249,
		DefaultID: 15252,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	893: { // minecraft:stone_stairs
		BaseID:    19133,
		DefaultID: 19144,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	937: { // minecraft:stonecutter
		BaseID:    24158,
		DefaultID: 24158,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	135: { // minecraft:straw_bed
		BaseID:    2286,
		DefaultID: 2289,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	67: { // minecraft:stripped_acacia_log
		BaseID:    186,
		DefaultID: 187,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	88: { // minecraft:stripped_acacia_wood
		BaseID:    249,
		DefaultID: 250,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	74: { // minecraft:stripped_bamboo_block
		BaseID:    207,
		DefaultID: 208,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	65: { // minecraft:stripped_birch_log
		BaseID:    180,
		DefaultID: 181,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	86: { // minecraft:stripped_birch_wood
		BaseID:    243,
		DefaultID: 244,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	68: { // minecraft:stripped_cherry_log
		BaseID:    189,
		DefaultID: 190,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	89: { // minecraft:stripped_cherry_wood
		BaseID:    252,
		DefaultID: 253,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	964: { // minecraft:stripped_crimson_hyphae
		BaseID:    24328,
		DefaultID: 24329,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	962: { // minecraft:stripped_crimson_stem
		BaseID:    24322,
		DefaultID: 24323,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	69: { // minecraft:stripped_dark_oak_log
		BaseID:    192,
		DefaultID: 193,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	90: { // minecraft:stripped_dark_oak_wood
		BaseID:    255,
		DefaultID: 256,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	66: { // minecraft:stripped_jungle_log
		BaseID:    183,
		DefaultID: 184,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	87: { // minecraft:stripped_jungle_wood
		BaseID:    246,
		DefaultID: 247,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	72: { // minecraft:stripped_mangrove_log
		BaseID:    201,
		DefaultID: 202,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	92: { // minecraft:stripped_mangrove_wood
		BaseID:    261,
		DefaultID: 262,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	71: { // minecraft:stripped_oak_log
		BaseID:    198,
		DefaultID: 199,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	84: { // minecraft:stripped_oak_wood
		BaseID:    237,
		DefaultID: 238,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	70: { // minecraft:stripped_pale_oak_log
		BaseID:    195,
		DefaultID: 196,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	91: { // minecraft:stripped_pale_oak_wood
		BaseID:    258,
		DefaultID: 259,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	73: { // minecraft:stripped_poplar_log
		BaseID:    204,
		DefaultID: 205,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	93: { // minecraft:stripped_poplar_wood
		BaseID:    264,
		DefaultID: 265,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	64: { // minecraft:stripped_spruce_log
		BaseID:    177,
		DefaultID: 178,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	85: { // minecraft:stripped_spruce_wood
		BaseID:    240,
		DefaultID: 241,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	955: { // minecraft:stripped_warped_hyphae
		BaseID:    24311,
		DefaultID: 24312,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	953: { // minecraft:stripped_warped_stem
		BaseID:    24305,
		DefaultID: 24306,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	995: { // minecraft:structure_block
		BaseID:    25079,
		DefaultID: 25080,
		Properties: []blockProperty{
			{"mode", []string{"save", "load", "corner", "data"}},
		},
	},
	733: { // minecraft:structure_void
		BaseID:    16832,
		DefaultID: 16832,
	},
	331: { // minecraft:sugar_cane
		BaseID:    8617,
		DefaultID: 8617,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	1088: { // minecraft:sulfur
		BaseID:    28044,
		DefaultID: 28044,
	},
	1098: { // minecraft:sulfur_brick_slab
		BaseID:    28872,
		DefaultID: 28875,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1099: { // minecraft:sulfur_brick_stairs
		BaseID:    28878,
		DefaultID: 28889,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1100: { // minecraft:sulfur_brick_wall
		BaseID:    28958,
		DefaultID: 28961,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	1097: { // minecraft:sulfur_bricks
		BaseID:    28871,
		DefaultID: 28871,
	},
	1090: { // minecraft:sulfur_slab
		BaseID:    28050,
		DefaultID: 28053,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1224: { // minecraft:sulfur_spike
		BaseID:    33586,
		DefaultID: 33591,
		Properties: []blockProperty{
			{"thickness", []string{"tip_merge", "tip", "frustum", "middle", "base"}},
			{"vertical_direction", []string{"up", "down"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1091: { // minecraft:sulfur_stairs
		BaseID:    28056,
		DefaultID: 28067,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1092: { // minecraft:sulfur_wall
		BaseID:    28136,
		DefaultID: 28139,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	611: { // minecraft:sunflower
		BaseID:    14762,
		DefaultID: 14763,
		Properties: []blockProperty{
			{"half", []string{"upper", "lower"}},
		},
	},
	43: { // minecraft:suspicious_gravel
		BaseID:    128,
		DefaultID: 128,
		Properties: []blockProperty{
			{"dusted", []string{"0", "1", "2", "3"}},
		},
	},
	40: { // minecraft:suspicious_sand
		BaseID:    122,
		DefaultID: 122,
		Properties: []blockProperty{
			{"dusted", []string{"0", "1", "2", "3"}},
		},
	},
	951: { // minecraft:sweet_berry_bush
		BaseID:    24298,
		DefaultID: 24298,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3"}},
		},
	},
	146: { // minecraft:tall_dry_grass
		BaseID:    2369,
		DefaultID: 2369,
	},
	615: { // minecraft:tall_grass
		BaseID:    14770,
		DefaultID: 14771,
		Properties: []blockProperty{
			{"half", []string{"upper", "lower"}},
		},
	},
	148: { // minecraft:tall_seagrass
		BaseID:    2371,
		DefaultID: 2372,
		Properties: []blockProperty{
			{"half", []string{"upper", "lower"}},
		},
	},
	1000: { // minecraft:target
		BaseID:    25109,
		DefaultID: 25109,
		Properties: []blockProperty{
			{"power", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	608: { // minecraft:terracotta
		BaseID:    14759,
		DefaultID: 14759,
	},
	997: { // minecraft:test_block
		BaseID:    25095,
		DefaultID: 25095,
		Properties: []blockProperty{
			{"mode", []string{"start", "log", "fail", "accept"}},
		},
	},
	998: { // minecraft:test_instance_block
		BaseID:    25099,
		DefaultID: 25099,
	},
	1116: { // minecraft:tinted_glass
		BaseID:    30518,
		DefaultID: 30518,
	},
	220: { // minecraft:tnt
		BaseID:    3833,
		DefaultID: 3834,
		Properties: []blockProperty{
			{"unstable", []string{"true", "false"}},
		},
	},
	238: { // minecraft:torch
		BaseID:    4926,
		DefaultID: 4926,
	},
	202: { // minecraft:torchflower
		BaseID:    3815,
		DefaultID: 3815,
	},
	720: { // minecraft:torchflower_crop
		BaseID:    16778,
		DefaultID: 16778,
		Properties: []blockProperty{
			{"age", []string{"0", "1"}},
		},
	},
	523: { // minecraft:trapped_chest
		BaseID:    12974,
		DefaultID: 12975,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1275: { // minecraft:trial_spawner
		BaseID:    35507,
		DefaultID: 35513,
		Properties: []blockProperty{
			{"ominous", []string{"true", "false"}},
			{"trial_spawner_state", []string{"inactive", "waiting_for_players", "active", "waiting_for_reward_ejection", "ejecting_reward", "cooldown"}},
		},
	},
	453: { // minecraft:tripwire
		BaseID:    11341,
		DefaultID: 11468,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"disarmed", []string{"true", "false"}},
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	452: { // minecraft:tripwire_hook
		BaseID:    11325,
		DefaultID: 11334,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	853: { // minecraft:tube_coral
		BaseID:    18514,
		DefaultID: 18514,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	843: { // minecraft:tube_coral_block
		BaseID:    18499,
		DefaultID: 18499,
	},
	863: { // minecraft:tube_coral_fan
		BaseID:    18534,
		DefaultID: 18534,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	873: { // minecraft:tube_coral_wall_fan
		BaseID:    18584,
		DefaultID: 18584,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1074: { // minecraft:tuff
		BaseID:    26809,
		DefaultID: 26809,
	},
	1084: { // minecraft:tuff_brick_slab
		BaseID:    27633,
		DefaultID: 27636,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1085: { // minecraft:tuff_brick_stairs
		BaseID:    27639,
		DefaultID: 27650,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1086: { // minecraft:tuff_brick_wall
		BaseID:    27719,
		DefaultID: 27722,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	1083: { // minecraft:tuff_bricks
		BaseID:    27632,
		DefaultID: 27632,
	},
	1075: { // minecraft:tuff_slab
		BaseID:    26810,
		DefaultID: 26813,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1076: { // minecraft:tuff_stairs
		BaseID:    26816,
		DefaultID: 26827,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1077: { // minecraft:tuff_wall
		BaseID:    26896,
		DefaultID: 26899,
		Properties: []blockProperty{
			{"east", []string{"none", "low", "tall"}},
			{"north", []string{"none", "low", "tall"}},
			{"south", []string{"none", "low", "tall"}},
			{"up", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"none", "low", "tall"}},
		},
	},
	835: { // minecraft:turtle_egg
		BaseID:    18447,
		DefaultID: 18447,
		Properties: []blockProperty{
			{"eggs", []string{"1", "2", "3", "4"}},
			{"hatch", []string{"0", "1", "2"}},
		},
	},
	970: { // minecraft:twisting_vines
		BaseID:    24361,
		DefaultID: 24361,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25"}},
		},
	},
	971: { // minecraft:twisting_vines_plant
		BaseID:    24387,
		DefaultID: 24387,
	},
	1276: { // minecraft:vault
		BaseID:    35519,
		DefaultID: 35523,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"ominous", []string{"true", "false"}},
			{"vault_state", []string{"inactive", "active", "unlocking", "ejecting"}},
		},
	},
	1269: { // minecraft:verdant_froglight
		BaseID:    35435,
		DefaultID: 35436,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	416: { // minecraft:vine
		BaseID:    10092,
		DefaultID: 10123,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"up", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	884: { // minecraft:void_air
		BaseID:    18649,
		DefaultID: 18649,
	},
	239: { // minecraft:wall_torch
		BaseID:    4927,
		DefaultID: 4927,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	988: { // minecraft:warped_button
		BaseID:    24847,
		DefaultID: 24856,
		Properties: []blockProperty{
			{"face", []string{"floor", "wall", "ceiling"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	990: { // minecraft:warped_door
		BaseID:    24935,
		DefaultID: 24946,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	980: { // minecraft:warped_fence
		BaseID:    24439,
		DefaultID: 24470,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	984: { // minecraft:warped_fence_gate
		BaseID:    24631,
		DefaultID: 24638,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"in_wall", []string{"true", "false"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	957: { // minecraft:warped_fungus
		BaseID:    24315,
		DefaultID: 24315,
	},
	289: { // minecraft:warped_hanging_sign
		BaseID:    8079,
		DefaultID: 8128,
		Properties: []blockProperty{
			{"attached", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	954: { // minecraft:warped_hyphae
		BaseID:    24308,
		DefaultID: 24309,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	956: { // minecraft:warped_nylium
		BaseID:    24314,
		DefaultID: 24314,
	},
	974: { // minecraft:warped_planks
		BaseID:    24390,
		DefaultID: 24390,
	},
	978: { // minecraft:warped_pressure_plate
		BaseID:    24405,
		DefaultID: 24406,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
		},
	},
	959: { // minecraft:warped_roots
		BaseID:    24317,
		DefaultID: 24317,
	},
	235: { // minecraft:warped_shelf
		BaseID:    4860,
		DefaultID: 4869,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
			{"side_chain", []string{"unconnected", "right", "center", "left"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	992: { // minecraft:warped_sign
		BaseID:    25031,
		DefaultID: 25048,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	976: { // minecraft:warped_slab
		BaseID:    24397,
		DefaultID: 24400,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	986: { // minecraft:warped_stairs
		BaseID:    24743,
		DefaultID: 24754,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	952: { // minecraft:warped_stem
		BaseID:    24302,
		DefaultID: 24303,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
		},
	},
	982: { // minecraft:warped_trapdoor
		BaseID:    24535,
		DefaultID: 24550,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	304: { // minecraft:warped_wall_hanging_sign
		BaseID:    8423,
		DefaultID: 8424,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	994: { // minecraft:warped_wall_sign
		BaseID:    25071,
		DefaultID: 25072,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	958: { // minecraft:warped_wart_block
		BaseID:    24316,
		DefaultID: 24316,
	},
	37: { // minecraft:water
		BaseID:    89,
		DefaultID: 89,
		Properties: []blockProperty{
			{"level", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	438: { // minecraft:water_cauldron
		BaseID:    11195,
		DefaultID: 11195,
		Properties: []blockProperty{
			{"level", []string{"1", "2", "3"}},
		},
	},
	1146: { // minecraft:waxed_chiseled_copper
		BaseID:    31161,
		DefaultID: 31161,
	},
	396: { // minecraft:waxed_copper_bars
		BaseID:    9852,
		DefaultID: 9883,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1128: { // minecraft:waxed_copper_block
		BaseID:    31143,
		DefaultID: 31143,
	},
	1194: { // minecraft:waxed_copper_bulb
		BaseID:    32909,
		DefaultID: 32912,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	405: { // minecraft:waxed_copper_chain
		BaseID:    10010,
		DefaultID: 10013,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1202: { // minecraft:waxed_copper_chest
		BaseID:    33021,
		DefaultID: 33022,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1170: { // minecraft:waxed_copper_door
		BaseID:    32109,
		DefaultID: 32120,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1210: { // minecraft:waxed_copper_golem_statue
		BaseID:    33245,
		DefaultID: 33246,
		Properties: []blockProperty{
			{"copper_golem_pose", []string{"standing", "sitting", "running", "star"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1186: { // minecraft:waxed_copper_grate
		BaseID:    32885,
		DefaultID: 32886,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	945: { // minecraft:waxed_copper_lantern
		BaseID:    24218,
		DefaultID: 24221,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1178: { // minecraft:waxed_copper_trapdoor
		BaseID:    32621,
		DefaultID: 32636,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1138: { // minecraft:waxed_cut_copper
		BaseID:    31153,
		DefaultID: 31153,
	},
	1162: { // minecraft:waxed_cut_copper_slab
		BaseID:    31829,
		DefaultID: 31832,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1154: { // minecraft:waxed_cut_copper_stairs
		BaseID:    31485,
		DefaultID: 31496,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1147: { // minecraft:waxed_exposed_chiseled_copper
		BaseID:    31162,
		DefaultID: 31162,
	},
	1129: { // minecraft:waxed_exposed_copper
		BaseID:    31144,
		DefaultID: 31144,
	},
	397: { // minecraft:waxed_exposed_copper_bars
		BaseID:    9884,
		DefaultID: 9915,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1195: { // minecraft:waxed_exposed_copper_bulb
		BaseID:    32913,
		DefaultID: 32916,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	406: { // minecraft:waxed_exposed_copper_chain
		BaseID:    10016,
		DefaultID: 10019,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1203: { // minecraft:waxed_exposed_copper_chest
		BaseID:    33045,
		DefaultID: 33046,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1171: { // minecraft:waxed_exposed_copper_door
		BaseID:    32173,
		DefaultID: 32184,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1211: { // minecraft:waxed_exposed_copper_golem_statue
		BaseID:    33277,
		DefaultID: 33278,
		Properties: []blockProperty{
			{"copper_golem_pose", []string{"standing", "sitting", "running", "star"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1187: { // minecraft:waxed_exposed_copper_grate
		BaseID:    32887,
		DefaultID: 32888,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	946: { // minecraft:waxed_exposed_copper_lantern
		BaseID:    24222,
		DefaultID: 24225,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1179: { // minecraft:waxed_exposed_copper_trapdoor
		BaseID:    32685,
		DefaultID: 32700,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1139: { // minecraft:waxed_exposed_cut_copper
		BaseID:    31154,
		DefaultID: 31154,
	},
	1163: { // minecraft:waxed_exposed_cut_copper_slab
		BaseID:    31835,
		DefaultID: 31838,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1155: { // minecraft:waxed_exposed_cut_copper_stairs
		BaseID:    31565,
		DefaultID: 31576,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1219: { // minecraft:waxed_exposed_lightning_rod
		BaseID:    33493,
		DefaultID: 33512,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1218: { // minecraft:waxed_lightning_rod
		BaseID:    33469,
		DefaultID: 33488,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1149: { // minecraft:waxed_oxidized_chiseled_copper
		BaseID:    31164,
		DefaultID: 31164,
	},
	1131: { // minecraft:waxed_oxidized_copper
		BaseID:    31146,
		DefaultID: 31146,
	},
	399: { // minecraft:waxed_oxidized_copper_bars
		BaseID:    9948,
		DefaultID: 9979,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1197: { // minecraft:waxed_oxidized_copper_bulb
		BaseID:    32921,
		DefaultID: 32924,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	408: { // minecraft:waxed_oxidized_copper_chain
		BaseID:    10028,
		DefaultID: 10031,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1205: { // minecraft:waxed_oxidized_copper_chest
		BaseID:    33093,
		DefaultID: 33094,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1173: { // minecraft:waxed_oxidized_copper_door
		BaseID:    32301,
		DefaultID: 32312,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1213: { // minecraft:waxed_oxidized_copper_golem_statue
		BaseID:    33341,
		DefaultID: 33342,
		Properties: []blockProperty{
			{"copper_golem_pose", []string{"standing", "sitting", "running", "star"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1189: { // minecraft:waxed_oxidized_copper_grate
		BaseID:    32891,
		DefaultID: 32892,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	948: { // minecraft:waxed_oxidized_copper_lantern
		BaseID:    24230,
		DefaultID: 24233,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1181: { // minecraft:waxed_oxidized_copper_trapdoor
		BaseID:    32813,
		DefaultID: 32828,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1141: { // minecraft:waxed_oxidized_cut_copper
		BaseID:    31156,
		DefaultID: 31156,
	},
	1165: { // minecraft:waxed_oxidized_cut_copper_slab
		BaseID:    31847,
		DefaultID: 31850,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1157: { // minecraft:waxed_oxidized_cut_copper_stairs
		BaseID:    31725,
		DefaultID: 31736,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1221: { // minecraft:waxed_oxidized_lightning_rod
		BaseID:    33541,
		DefaultID: 33560,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1148: { // minecraft:waxed_weathered_chiseled_copper
		BaseID:    31163,
		DefaultID: 31163,
	},
	1130: { // minecraft:waxed_weathered_copper
		BaseID:    31145,
		DefaultID: 31145,
	},
	398: { // minecraft:waxed_weathered_copper_bars
		BaseID:    9916,
		DefaultID: 9947,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1196: { // minecraft:waxed_weathered_copper_bulb
		BaseID:    32917,
		DefaultID: 32920,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	407: { // minecraft:waxed_weathered_copper_chain
		BaseID:    10022,
		DefaultID: 10025,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1204: { // minecraft:waxed_weathered_copper_chest
		BaseID:    33069,
		DefaultID: 33070,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1172: { // minecraft:waxed_weathered_copper_door
		BaseID:    32237,
		DefaultID: 32248,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1212: { // minecraft:waxed_weathered_copper_golem_statue
		BaseID:    33309,
		DefaultID: 33310,
		Properties: []blockProperty{
			{"copper_golem_pose", []string{"standing", "sitting", "running", "star"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1188: { // minecraft:waxed_weathered_copper_grate
		BaseID:    32889,
		DefaultID: 32890,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	947: { // minecraft:waxed_weathered_copper_lantern
		BaseID:    24226,
		DefaultID: 24229,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1180: { // minecraft:waxed_weathered_copper_trapdoor
		BaseID:    32749,
		DefaultID: 32764,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1140: { // minecraft:waxed_weathered_cut_copper
		BaseID:    31155,
		DefaultID: 31155,
	},
	1164: { // minecraft:waxed_weathered_cut_copper_slab
		BaseID:    31841,
		DefaultID: 31844,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1156: { // minecraft:waxed_weathered_cut_copper_stairs
		BaseID:    31645,
		DefaultID: 31656,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1220: { // minecraft:waxed_weathered_lightning_rod
		BaseID:    33517,
		DefaultID: 33536,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1144: { // minecraft:weathered_chiseled_copper
		BaseID:    31159,
		DefaultID: 31159,
	},
	1126: { // minecraft:weathered_copper
		BaseID:    31141,
		DefaultID: 31141,
	},
	394: { // minecraft:weathered_copper_bars
		BaseID:    9788,
		DefaultID: 9819,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	1192: { // minecraft:weathered_copper_bulb
		BaseID:    32901,
		DefaultID: 32904,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	403: { // minecraft:weathered_copper_chain
		BaseID:    9998,
		DefaultID: 10001,
		Properties: []blockProperty{
			{"axis", []string{"x", "y", "z"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1200: { // minecraft:weathered_copper_chest
		BaseID:    32973,
		DefaultID: 32974,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"type", []string{"single", "left", "right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1168: { // minecraft:weathered_copper_door
		BaseID:    31981,
		DefaultID: 31992,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"upper", "lower"}},
			{"hinge", []string{"left", "right"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
		},
	},
	1208: { // minecraft:weathered_copper_golem_statue
		BaseID:    33181,
		DefaultID: 33182,
		Properties: []blockProperty{
			{"copper_golem_pose", []string{"standing", "sitting", "running", "star"}},
			{"facing", []string{"north", "south", "west", "east"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1184: { // minecraft:weathered_copper_grate
		BaseID:    32881,
		DefaultID: 32882,
		Properties: []blockProperty{
			{"waterlogged", []string{"true", "false"}},
		},
	},
	943: { // minecraft:weathered_copper_lantern
		BaseID:    24210,
		DefaultID: 24213,
		Properties: []blockProperty{
			{"hanging", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1176: { // minecraft:weathered_copper_trapdoor
		BaseID:    32493,
		DefaultID: 32508,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"open", []string{"true", "false"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1136: { // minecraft:weathered_cut_copper
		BaseID:    31151,
		DefaultID: 31151,
	},
	1160: { // minecraft:weathered_cut_copper_slab
		BaseID:    31817,
		DefaultID: 31820,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1152: { // minecraft:weathered_cut_copper_stairs
		BaseID:    31325,
		DefaultID: 31336,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1216: { // minecraft:weathered_lightning_rod
		BaseID:    33421,
		DefaultID: 33440,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
			{"powered", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	968: { // minecraft:weeping_vines
		BaseID:    24334,
		DefaultID: 24334,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25"}},
		},
	},
	969: { // minecraft:weeping_vines_plant
		BaseID:    24360,
		DefaultID: 24360,
	},
	109: { // minecraft:wet_sponge
		BaseID:    660,
		DefaultID: 660,
	},
	251: { // minecraft:wheat
		BaseID:    6867,
		DefaultID: 6867,
		Properties: []blockProperty{
			{"age", []string{"0", "1", "2", "3", "4", "5", "6", "7"}},
		},
	},
	617: { // minecraft:white_banner
		BaseID:    14774,
		DefaultID: 14782,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	119: { // minecraft:white_bed
		BaseID:    2030,
		DefaultID: 2033,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1035: { // minecraft:white_candle
		BaseID:    26469,
		DefaultID: 26472,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1052: { // minecraft:white_candle_cake
		BaseID:    26727,
		DefaultID: 26728,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	592: { // minecraft:white_carpet
		BaseID:    14743,
		DefaultID: 14743,
	},
	768: { // minecraft:white_concrete
		BaseID:    17011,
		DefaultID: 17011,
	},
	816: { // minecraft:white_concrete_powder
		BaseID:    18403,
		DefaultID: 18403,
	},
	800: { // minecraft:white_concrete_slab
		BaseID:    18307,
		DefaultID: 18310,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	784: { // minecraft:white_concrete_stairs
		BaseID:    17027,
		DefaultID: 17038,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	752: { // minecraft:white_glazed_terracotta
		BaseID:    16947,
		DefaultID: 16947,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	736: { // minecraft:white_shulker_box
		BaseID:    16851,
		DefaultID: 16855,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	349: { // minecraft:white_stained_glass
		BaseID:    8768,
		DefaultID: 8768,
	},
	553: { // minecraft:white_stained_glass_pane
		BaseID:    13227,
		DefaultID: 13258,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	537: { // minecraft:white_terracotta
		BaseID:    13211,
		DefaultID: 13211,
	},
	209: { // minecraft:white_tulip
		BaseID:    3822,
		DefaultID: 3822,
	},
	633: { // minecraft:white_wall_banner
		BaseID:    15030,
		DefaultID: 15030,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	151: { // minecraft:white_wool
		BaseID:    2409,
		DefaultID: 2409,
	},
	183: { // minecraft:white_wool_slab
		BaseID:    3705,
		DefaultID: 3708,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	167: { // minecraft:white_wool_stairs
		BaseID:    2425,
		DefaultID: 2436,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1232: { // minecraft:wildflowers
		BaseID:    33680,
		DefaultID: 33680,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"flower_amount", []string{"1", "2", "3", "4"}},
		},
	},
	213: { // minecraft:wither_rose
		BaseID:    3826,
		DefaultID: 3826,
	},
	508: { // minecraft:wither_skeleton_skull
		BaseID:    12722,
		DefaultID: 12738,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	509: { // minecraft:wither_skeleton_wall_skull
		BaseID:    12754,
		DefaultID: 12755,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
	621: { // minecraft:yellow_banner
		BaseID:    14838,
		DefaultID: 14846,
		Properties: []blockProperty{
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	123: { // minecraft:yellow_bed
		BaseID:    2094,
		DefaultID: 2097,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"occupied", []string{"true", "false"}},
			{"part", []string{"head", "foot"}},
		},
	},
	1039: { // minecraft:yellow_candle
		BaseID:    26533,
		DefaultID: 26536,
		Properties: []blockProperty{
			{"candles", []string{"1", "2", "3", "4"}},
			{"lit", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	1056: { // minecraft:yellow_candle_cake
		BaseID:    26735,
		DefaultID: 26736,
		Properties: []blockProperty{
			{"lit", []string{"true", "false"}},
		},
	},
	596: { // minecraft:yellow_carpet
		BaseID:    14747,
		DefaultID: 14747,
	},
	772: { // minecraft:yellow_concrete
		BaseID:    17015,
		DefaultID: 17015,
	},
	820: { // minecraft:yellow_concrete_powder
		BaseID:    18407,
		DefaultID: 18407,
	},
	804: { // minecraft:yellow_concrete_slab
		BaseID:    18331,
		DefaultID: 18334,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	788: { // minecraft:yellow_concrete_stairs
		BaseID:    17347,
		DefaultID: 17358,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	756: { // minecraft:yellow_glazed_terracotta
		BaseID:    16963,
		DefaultID: 16963,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	105: { // minecraft:yellow_poplar_leaves
		BaseID:    575,
		DefaultID: 602,
		Properties: []blockProperty{
			{"distance", []string{"1", "2", "3", "4", "5", "6", "7"}},
			{"persistent", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	740: { // minecraft:yellow_shulker_box
		BaseID:    16875,
		DefaultID: 16879,
		Properties: []blockProperty{
			{"facing", []string{"north", "east", "south", "west", "up", "down"}},
		},
	},
	353: { // minecraft:yellow_stained_glass
		BaseID:    8772,
		DefaultID: 8772,
	},
	557: { // minecraft:yellow_stained_glass_pane
		BaseID:    13355,
		DefaultID: 13386,
		Properties: []blockProperty{
			{"east", []string{"true", "false"}},
			{"north", []string{"true", "false"}},
			{"south", []string{"true", "false"}},
			{"waterlogged", []string{"true", "false"}},
			{"west", []string{"true", "false"}},
		},
	},
	541: { // minecraft:yellow_terracotta
		BaseID:    13215,
		DefaultID: 13215,
	},
	637: { // minecraft:yellow_wall_banner
		BaseID:    15046,
		DefaultID: 15046,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
		},
	},
	155: { // minecraft:yellow_wool
		BaseID:    2413,
		DefaultID: 2413,
	},
	187: { // minecraft:yellow_wool_slab
		BaseID:    3729,
		DefaultID: 3732,
		Properties: []blockProperty{
			{"type", []string{"top", "bottom", "double"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	171: { // minecraft:yellow_wool_stairs
		BaseID:    2745,
		DefaultID: 2756,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"half", []string{"top", "bottom"}},
			{"shape", []string{"straight", "inner_left", "inner_right", "outer_left", "outer_right"}},
			{"waterlogged", []string{"true", "false"}},
		},
	},
	510: { // minecraft:zombie_head
		BaseID:    12762,
		DefaultID: 12778,
		Properties: []blockProperty{
			{"powered", []string{"true", "false"}},
			{"rotation", []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15"}},
		},
	},
	511: { // minecraft:zombie_wall_head
		BaseID:    12794,
		DefaultID: 12795,
		Properties: []blockProperty{
			{"facing", []string{"north", "south", "west", "east"}},
			{"powered", []string{"true", "false"}},
		},
	},
}

// stateRanges is sorted by BaseID for binary search.
var stateRanges = []stateRangeEntry{
	{0, 1, 0},
	{1, 2, 1},
	{2, 3, 2},
	{3, 4, 3},
	{4, 5, 4},
	{5, 6, 5},
	{6, 7, 6},
	{7, 8, 7},
	{8, 10, 8},
	{10, 11, 9},
	{11, 12, 10},
	{12, 14, 11},
	{14, 15, 12},
	{15, 16, 13},
	{16, 17, 14},
	{17, 18, 15},
	{18, 19, 16},
	{19, 20, 17},
	{20, 21, 18},
	{21, 22, 19},
	{22, 25, 20},
	{25, 26, 21},
	{26, 27, 22},
	{27, 28, 23},
	{28, 29, 24},
	{29, 30, 25},
	{30, 32, 26},
	{32, 34, 27},
	{34, 36, 28},
	{36, 38, 29},
	{38, 40, 30},
	{40, 42, 31},
	{42, 44, 32},
	{44, 46, 33},
	{46, 86, 34},
	{86, 88, 35},
	{88, 89, 36},
	{89, 105, 37},
	{105, 121, 38},
	{121, 122, 39},
	{122, 126, 40},
	{126, 127, 41},
	{127, 128, 42},
	{128, 132, 43},
	{132, 133, 44},
	{133, 134, 45},
	{134, 135, 46},
	{135, 136, 47},
	{136, 137, 48},
	{137, 138, 49},
	{138, 139, 50},
	{139, 142, 51},
	{142, 145, 52},
	{145, 148, 53},
	{148, 151, 54},
	{151, 154, 55},
	{154, 157, 56},
	{157, 160, 57},
	{160, 163, 58},
	{163, 166, 59},
	{166, 169, 60},
	{169, 171, 61},
	{171, 174, 62},
	{174, 177, 63},
	{177, 180, 64},
	{180, 183, 65},
	{183, 186, 66},
	{186, 189, 67},
	{189, 192, 68},
	{192, 195, 69},
	{195, 198, 70},
	{198, 201, 71},
	{201, 204, 72},
	{204, 207, 73},
	{207, 210, 74},
	{210, 213, 75},
	{213, 216, 76},
	{216, 219, 77},
	{219, 222, 78},
	{222, 225, 79},
	{225, 228, 80},
	{228, 231, 81},
	{231, 234, 82},
	{234, 237, 83},
	{237, 240, 84},
	{240, 243, 85},
	{243, 246, 86},
	{246, 249, 87},
	{249, 252, 88},
	{252, 255, 89},
	{255, 258, 90},
	{258, 261, 91},
	{261, 264, 92},
	{264, 267, 93},
	{267, 295, 94},
	{295, 323, 95},
	{323, 351, 96},
	{351, 379, 97},
	{379, 407, 98},
	{407, 435, 99},
	{435, 463, 100},
	{463, 491, 101},
	{491, 519, 102},
	{519, 547, 103},
	{547, 575, 104},
	{575, 603, 105},
	{603, 631, 106},
	{631, 659, 107},
	{659, 660, 108},
	{660, 661, 109},
	{661, 662, 110},
	{662, 663, 111},
	{663, 664, 112},
	{664, 665, 113},
	{665, 677, 114},
	{677, 678, 115},
	{678, 679, 116},
	{679, 680, 117},
	{680, 2030, 118},
	{2030, 2046, 119},
	{2046, 2062, 120},
	{2062, 2078, 121},
	{2078, 2094, 122},
	{2094, 2110, 123},
	{2110, 2126, 124},
	{2126, 2142, 125},
	{2142, 2158, 126},
	{2158, 2174, 127},
	{2174, 2190, 128},
	{2190, 2206, 129},
	{2206, 2222, 130},
	{2222, 2238, 131},
	{2238, 2254, 132},
	{2254, 2270, 133},
	{2270, 2286, 134},
	{2286, 2302, 135},
	{2302, 2326, 136},
	{2326, 2350, 137},
	{2350, 2362, 138},
	{2362, 2363, 139},
	{2363, 2364, 140},
	{2364, 2365, 141},
	{2365, 2366, 142},
	{2366, 2367, 143},
	{2367, 2368, 144},
	{2368, 2369, 145},
	{2369, 2370, 146},
	{2370, 2371, 147},
	{2371, 2373, 148},
	{2373, 2385, 149},
	{2385, 2409, 150},
	{2409, 2410, 151},
	{2410, 2411, 152},
	{2411, 2412, 153},
	{2412, 2413, 154},
	{2413, 2414, 155},
	{2414, 2415, 156},
	{2415, 2416, 157},
	{2416, 2417, 158},
	{2417, 2418, 159},
	{2418, 2419, 160},
	{2419, 2420, 161},
	{2420, 2421, 162},
	{2421, 2422, 163},
	{2422, 2423, 164},
	{2423, 2424, 165},
	{2424, 2425, 166},
	{2425, 2505, 167},
	{2505, 2585, 168},
	{2585, 2665, 169},
	{2665, 2745, 170},
	{2745, 2825, 171},
	{2825, 2905, 172},
	{2905, 2985, 173},
	{2985, 3065, 174},
	{3065, 3145, 175},
	{3145, 3225, 176},
	{3225, 3305, 177},
	{3305, 3385, 178},
	{3385, 3465, 179},
	{3465, 3545, 180},
	{3545, 3625, 181},
	{3625, 3705, 182},
	{3705, 3711, 183},
	{3711, 3717, 184},
	{3717, 3723, 185},
	{3723, 3729, 186},
	{3729, 3735, 187},
	{3735, 3741, 188},
	{3741, 3747, 189},
	{3747, 3753, 190},
	{3753, 3759, 191},
	{3759, 3765, 192},
	{3765, 3771, 193},
	{3771, 3777, 194},
	{3777, 3783, 195},
	{3783, 3789, 196},
	{3789, 3795, 197},
	{3795, 3801, 198},
	{3801, 3813, 199},
	{3813, 3814, 200},
	{3814, 3815, 201},
	{3815, 3816, 202},
	{3816, 3817, 203},
	{3817, 3818, 204},
	{3818, 3819, 205},
	{3819, 3820, 206},
	{3820, 3821, 207},
	{3821, 3822, 208},
	{3822, 3823, 209},
	{3823, 3824, 210},
	{3824, 3825, 211},
	{3825, 3826, 212},
	{3826, 3827, 213},
	{3827, 3828, 214},
	{3828, 3829, 215},
	{3829, 3830, 216},
	{3830, 3831, 217},
	{3831, 3832, 218},
	{3832, 3833, 219},
	{3833, 3835, 220},
	{3835, 3836, 221},
	{3836, 4092, 222},
	{4092, 4156, 223},
	{4156, 4220, 224},
	{4220, 4284, 225},
	{4284, 4348, 226},
	{4348, 4412, 227},
	{4412, 4476, 228},
	{4476, 4540, 229},
	{4540, 4604, 230},
	{4604, 4668, 231},
	{4668, 4732, 232},
	{4732, 4796, 233},
	{4796, 4860, 234},
	{4860, 4924, 235},
	{4924, 4925, 236},
	{4925, 4926, 237},
	{4926, 4927, 238},
	{4927, 4931, 239},
	{4931, 5443, 240},
	{5443, 5444, 241},
	{5444, 5445, 242},
	{5445, 5463, 243},
	{5463, 5543, 244},
	{5543, 5567, 245},
	{5567, 6863, 246},
	{6863, 6864, 247},
	{6864, 6865, 248},
	{6865, 6866, 249},
	{6866, 6867, 250},
	{6867, 6875, 251},
	{6875, 6883, 252},
	{6883, 6891, 253},
	{6891, 6923, 254},
	{6923, 6955, 255},
	{6955, 6987, 256},
	{6987, 7019, 257},
	{7019, 7051, 258},
	{7051, 7083, 259},
	{7083, 7115, 260},
	{7115, 7147, 261},
	{7147, 7179, 262},
	{7179, 7211, 263},
	{7211, 7243, 264},
	{7243, 7307, 265},
	{7307, 7315, 266},
	{7315, 7335, 267},
	{7335, 7415, 268},
	{7415, 7423, 269},
	{7423, 7431, 270},
	{7431, 7439, 271},
	{7439, 7447, 272},
	{7447, 7455, 273},
	{7455, 7463, 274},
	{7463, 7471, 275},
	{7471, 7479, 276},
	{7479, 7487, 277},
	{7487, 7495, 278},
	{7495, 7503, 279},
	{7503, 7567, 280},
	{7567, 7631, 281},
	{7631, 7695, 282},
	{7695, 7759, 283},
	{7759, 7823, 284},
	{7823, 7887, 285},
	{7887, 7951, 286},
	{7951, 8015, 287},
	{8015, 8079, 288},
	{8079, 8143, 289},
	{8143, 8207, 290},
	{8207, 8271, 291},
	{8271, 8335, 292},
	{8335, 8343, 293},
	{8343, 8351, 294},
	{8351, 8359, 295},
	{8359, 8367, 296},
	{8367, 8375, 297},
	{8375, 8383, 298},
	{8383, 8391, 299},
	{8391, 8399, 300},
	{8399, 8407, 301},
	{8407, 8415, 302},
	{8415, 8423, 303},
	{8423, 8431, 304},
	{8431, 8439, 305},
	{8439, 8463, 306},
	{8463, 8465, 307},
	{8465, 8529, 308},
	{8529, 8531, 309},
	{8531, 8533, 310},
	{8533, 8535, 311},
	{8535, 8537, 312},
	{8537, 8539, 313},
	{8539, 8541, 314},
	{8541, 8543, 315},
	{8543, 8545, 316},
	{8545, 8547, 317},
	{8547, 8549, 318},
	{8549, 8551, 319},
	{8551, 8553, 320},
	{8553, 8555, 321},
	{8555, 8557, 322},
	{8557, 8565, 323},
	{8565, 8589, 324},
	{8589, 8597, 325},
	{8597, 8598, 326},
	{8598, 8599, 327},
	{8599, 8615, 328},
	{8615, 8616, 329},
	{8616, 8617, 330},
	{8617, 8633, 331},
	{8633, 8635, 332},
	{8635, 8667, 333},
	{8667, 8668, 334},
	{8668, 8669, 335},
	{8669, 8670, 336},
	{8670, 8673, 337},
	{8673, 8676, 338},
	{8676, 8677, 339},
	{8677, 8681, 340},
	{8681, 8682, 341},
	{8682, 8686, 342},
	{8686, 8687, 343},
	{8687, 8689, 344},
	{8689, 8693, 345},
	{8693, 8697, 346},
	{8697, 8704, 347},
	{8704, 8768, 348},
	{8768, 8769, 349},
	{8769, 8770, 350},
	{8770, 8771, 351},
	{8771, 8772, 352},
	{8772, 8773, 353},
	{8773, 8774, 354},
	{8774, 8775, 355},
	{8775, 8776, 356},
	{8776, 8777, 357},
	{8777, 8778, 358},
	{8778, 8779, 359},
	{8779, 8780, 360},
	{8780, 8781, 361},
	{8781, 8782, 362},
	{8782, 8783, 363},
	{8783, 8784, 364},
	{8784, 8848, 365},
	{8848, 8912, 366},
	{8912, 8976, 367},
	{8976, 9040, 368},
	{9040, 9104, 369},
	{9104, 9168, 370},
	{9168, 9232, 371},
	{9232, 9296, 372},
	{9296, 9360, 373},
	{9360, 9424, 374},
	{9424, 9488, 375},
	{9488, 9489, 376},
	{9489, 9490, 377},
	{9490, 9491, 378},
	{9491, 9492, 379},
	{9492, 9493, 380},
	{9493, 9494, 381},
	{9494, 9495, 382},
	{9495, 9496, 383},
	{9496, 9497, 384},
	{9497, 9498, 385},
	{9498, 9499, 386},
	{9499, 9500, 387},
	{9500, 9564, 388},
	{9564, 9628, 389},
	{9628, 9692, 390},
	{9692, 9724, 391},
	{9724, 9756, 392},
	{9756, 9788, 393},
	{9788, 9820, 394},
	{9820, 9852, 395},
	{9852, 9884, 396},
	{9884, 9916, 397},
	{9916, 9948, 398},
	{9948, 9980, 399},
	{9980, 9986, 400},
	{9986, 9992, 401},
	{9992, 9998, 402},
	{9998, 10004, 403},
	{10004, 10010, 404},
	{10010, 10016, 405},
	{10016, 10022, 406},
	{10022, 10028, 407},
	{10028, 10034, 408},
	{10034, 10066, 409},
	{10066, 10067, 410},
	{10067, 10068, 411},
	{10068, 10072, 412},
	{10072, 10076, 413},
	{10076, 10084, 414},
	{10084, 10092, 415},
	{10092, 10124, 416},
	{10124, 10252, 417},
	{10252, 10380, 418},
	{10380, 10412, 419},
	{10412, 10492, 420},
	{10492, 10572, 421},
	{10572, 10652, 422},
	{10652, 10654, 423},
	{10654, 10655, 424},
	{10655, 10656, 425},
	{10656, 10657, 426},
	{10657, 10737, 427},
	{10737, 10743, 428},
	{10743, 11067, 429},
	{11067, 11068, 430},
	{11068, 11069, 431},
	{11069, 11101, 432},
	{11101, 11181, 433},
	{11181, 11185, 434},
	{11185, 11186, 435},
	{11186, 11194, 436},
	{11194, 11195, 437},
	{11195, 11198, 438},
	{11198, 11199, 439},
	{11199, 11202, 440},
	{11202, 11203, 441},
	{11203, 11211, 442},
	{11211, 11212, 443},
	{11212, 11213, 444},
	{11213, 11215, 445},
	{11215, 11227, 446},
	{11227, 11235, 447},
	{11235, 11315, 448},
	{11315, 11316, 449},
	{11316, 11317, 450},
	{11317, 11325, 451},
	{11325, 11341, 452},
	{11341, 11469, 453},
	{11469, 11470, 454},
	{11470, 11550, 455},
	{11550, 11630, 456},
	{11630, 11710, 457},
	{11710, 11722, 458},
	{11722, 11723, 459},
	{11723, 12047, 460},
	{12047, 12371, 461},
	{12371, 12372, 462},
	{12372, 12373, 463},
	{12373, 12374, 464},
	{12374, 12375, 465},
	{12375, 12376, 466},
	{12376, 12377, 467},
	{12377, 12378, 468},
	{12378, 12379, 469},
	{12379, 12380, 470},
	{12380, 12381, 471},
	{12381, 12382, 472},
	{12382, 12383, 473},
	{12383, 12384, 474},
	{12384, 12385, 475},
	{12385, 12386, 476},
	{12386, 12387, 477},
	{12387, 12388, 478},
	{12388, 12389, 479},
	{12389, 12390, 480},
	{12390, 12391, 481},
	{12391, 12392, 482},
	{12392, 12393, 483},
	{12393, 12394, 484},
	{12394, 12395, 485},
	{12395, 12396, 486},
	{12396, 12397, 487},
	{12397, 12398, 488},
	{12398, 12399, 489},
	{12399, 12400, 490},
	{12400, 12401, 491},
	{12401, 12402, 492},
	{12402, 12410, 493},
	{12410, 12418, 494},
	{12418, 12442, 495},
	{12442, 12466, 496},
	{12466, 12490, 497},
	{12490, 12514, 498},
	{12514, 12538, 499},
	{12538, 12562, 500},
	{12562, 12586, 501},
	{12586, 12610, 502},
	{12610, 12634, 503},
	{12634, 12658, 504},
	{12658, 12682, 505},
	{12682, 12714, 506},
	{12714, 12722, 507},
	{12722, 12754, 508},
	{12754, 12762, 509},
	{12762, 12794, 510},
	{12794, 12802, 511},
	{12802, 12834, 512},
	{12834, 12842, 513},
	{12842, 12874, 514},
	{12874, 12882, 515},
	{12882, 12914, 516},
	{12914, 12922, 517},
	{12922, 12954, 518},
	{12954, 12962, 519},
	{12962, 12966, 520},
	{12966, 12970, 521},
	{12970, 12974, 522},
	{12974, 12998, 523},
	{12998, 13014, 524},
	{13014, 13030, 525},
	{13030, 13046, 526},
	{13046, 13078, 527},
	{13078, 13079, 528},
	{13079, 13080, 529},
	{13080, 13090, 530},
	{13090, 13091, 531},
	{13091, 13092, 532},
	{13092, 13095, 533},
	{13095, 13175, 534},
	{13175, 13199, 535},
	{13199, 13211, 536},
	{13211, 13212, 537},
	{13212, 13213, 538},
	{13213, 13214, 539},
	{13214, 13215, 540},
	{13215, 13216, 541},
	{13216, 13217, 542},
	{13217, 13218, 543},
	{13218, 13219, 544},
	{13219, 13220, 545},
	{13220, 13221, 546},
	{13221, 13222, 547},
	{13222, 13223, 548},
	{13223, 13224, 549},
	{13224, 13225, 550},
	{13225, 13226, 551},
	{13226, 13227, 552},
	{13227, 13259, 553},
	{13259, 13291, 554},
	{13291, 13323, 555},
	{13323, 13355, 556},
	{13355, 13387, 557},
	{13387, 13419, 558},
	{13419, 13451, 559},
	{13451, 13483, 560},
	{13483, 13515, 561},
	{13515, 13547, 562},
	{13547, 13579, 563},
	{13579, 13611, 564},
	{13611, 13643, 565},
	{13643, 13675, 566},
	{13675, 13707, 567},
	{13707, 13739, 568},
	{13739, 13819, 569},
	{13819, 13899, 570},
	{13899, 13979, 571},
	{13979, 14059, 572},
	{14059, 14139, 573},
	{14139, 14219, 574},
	{14219, 14299, 575},
	{14299, 14379, 576},
	{14379, 14380, 577},
	{14380, 14382, 578},
	{14382, 14414, 579},
	{14414, 14478, 580},
	{14478, 14479, 581},
	{14479, 14480, 582},
	{14480, 14481, 583},
	{14481, 14561, 584},
	{14561, 14641, 585},
	{14641, 14721, 586},
	{14721, 14727, 587},
	{14727, 14733, 588},
	{14733, 14739, 589},
	{14739, 14740, 590},
	{14740, 14743, 591},
	{14743, 14744, 592},
	{14744, 14745, 593},
	{14745, 14746, 594},
	{14746, 14747, 595},
	{14747, 14748, 596},
	{14748, 14749, 597},
	{14749, 14750, 598},
	{14750, 14751, 599},
	{14751, 14752, 600},
	{14752, 14753, 601},
	{14753, 14754, 602},
	{14754, 14755, 603},
	{14755, 14756, 604},
	{14756, 14757, 605},
	{14757, 14758, 606},
	{14758, 14759, 607},
	{14759, 14760, 608},
	{14760, 14761, 609},
	{14761, 14762, 610},
	{14762, 14764, 611},
	{14764, 14766, 612},
	{14766, 14768, 613},
	{14768, 14770, 614},
	{14770, 14772, 615},
	{14772, 14774, 616},
	{14774, 14790, 617},
	{14790, 14806, 618},
	{14806, 14822, 619},
	{14822, 14838, 620},
	{14838, 14854, 621},
	{14854, 14870, 622},
	{14870, 14886, 623},
	{14886, 14902, 624},
	{14902, 14918, 625},
	{14918, 14934, 626},
	{14934, 14950, 627},
	{14950, 14966, 628},
	{14966, 14982, 629},
	{14982, 14998, 630},
	{14998, 15014, 631},
	{15014, 15030, 632},
	{15030, 15034, 633},
	{15034, 15038, 634},
	{15038, 15042, 635},
	{15042, 15046, 636},
	{15046, 15050, 637},
	{15050, 15054, 638},
	{15054, 15058, 639},
	{15058, 15062, 640},
	{15062, 15066, 641},
	{15066, 15070, 642},
	{15070, 15074, 643},
	{15074, 15078, 644},
	{15078, 15082, 645},
	{15082, 15086, 646},
	{15086, 15090, 647},
	{15090, 15094, 648},
	{15094, 15095, 649},
	{15095, 15096, 650},
	{15096, 15097, 651},
	{15097, 15177, 652},
	{15177, 15183, 653},
	{15183, 15189, 654},
	{15189, 15195, 655},
	{15195, 15201, 656},
	{15201, 15207, 657},
	{15207, 15213, 658},
	{15213, 15219, 659},
	{15219, 15225, 660},
	{15225, 15231, 661},
	{15231, 15237, 662},
	{15237, 15243, 663},
	{15243, 15249, 664},
	{15249, 15255, 665},
	{15255, 15261, 666},
	{15261, 15267, 667},
	{15267, 15273, 668},
	{15273, 15279, 669},
	{15279, 15285, 670},
	{15285, 15291, 671},
	{15291, 15297, 672},
	{15297, 15303, 673},
	{15303, 15309, 674},
	{15309, 15315, 675},
	{15315, 15321, 676},
	{15321, 15322, 677},
	{15322, 15328, 678},
	{15328, 15329, 679},
	{15329, 15330, 680},
	{15330, 15331, 681},
	{15331, 15363, 682},
	{15363, 15395, 683},
	{15395, 15427, 684},
	{15427, 15459, 685},
	{15459, 15491, 686},
	{15491, 15523, 687},
	{15523, 15555, 688},
	{15555, 15587, 689},
	{15587, 15619, 690},
	{15619, 15651, 691},
	{15651, 15683, 692},
	{15683, 15715, 693},
	{15715, 15747, 694},
	{15747, 15779, 695},
	{15779, 15811, 696},
	{15811, 15843, 697},
	{15843, 15875, 698},
	{15875, 15907, 699},
	{15907, 15939, 700},
	{15939, 15971, 701},
	{15971, 16035, 702},
	{16035, 16099, 703},
	{16099, 16163, 704},
	{16163, 16227, 705},
	{16227, 16291, 706},
	{16291, 16355, 707},
	{16355, 16419, 708},
	{16419, 16483, 709},
	{16483, 16547, 710},
	{16547, 16611, 711},
	{16611, 16617, 712},
	{16617, 16681, 713},
	{16681, 16687, 714},
	{16687, 16688, 715},
	{16688, 16694, 716},
	{16694, 16697, 717},
	{16697, 16777, 718},
	{16777, 16778, 719},
	{16778, 16780, 720},
	{16780, 16790, 721},
	{16790, 16792, 722},
	{16792, 16796, 723},
	{16796, 16797, 724},
	{16797, 16798, 725},
	{16798, 16810, 726},
	{16810, 16822, 727},
	{16822, 16826, 728},
	{16826, 16827, 729},
	{16827, 16828, 730},
	{16828, 16829, 731},
	{16829, 16832, 732},
	{16832, 16833, 733},
	{16833, 16845, 734},
	{16845, 16851, 735},
	{16851, 16857, 736},
	{16857, 16863, 737},
	{16863, 16869, 738},
	{16869, 16875, 739},
	{16875, 16881, 740},
	{16881, 16887, 741},
	{16887, 16893, 742},
	{16893, 16899, 743},
	{16899, 16905, 744},
	{16905, 16911, 745},
	{16911, 16917, 746},
	{16917, 16923, 747},
	{16923, 16929, 748},
	{16929, 16935, 749},
	{16935, 16941, 750},
	{16941, 16947, 751},
	{16947, 16951, 752},
	{16951, 16955, 753},
	{16955, 16959, 754},
	{16959, 16963, 755},
	{16963, 16967, 756},
	{16967, 16971, 757},
	{16971, 16975, 758},
	{16975, 16979, 759},
	{16979, 16983, 760},
	{16983, 16987, 761},
	{16987, 16991, 762},
	{16991, 16995, 763},
	{16995, 16999, 764},
	{16999, 17003, 765},
	{17003, 17007, 766},
	{17007, 17011, 767},
	{17011, 17012, 768},
	{17012, 17013, 769},
	{17013, 17014, 770},
	{17014, 17015, 771},
	{17015, 17016, 772},
	{17016, 17017, 773},
	{17017, 17018, 774},
	{17018, 17019, 775},
	{17019, 17020, 776},
	{17020, 17021, 777},
	{17021, 17022, 778},
	{17022, 17023, 779},
	{17023, 17024, 780},
	{17024, 17025, 781},
	{17025, 17026, 782},
	{17026, 17027, 783},
	{17027, 17107, 784},
	{17107, 17187, 785},
	{17187, 17267, 786},
	{17267, 17347, 787},
	{17347, 17427, 788},
	{17427, 17507, 789},
	{17507, 17587, 790},
	{17587, 17667, 791},
	{17667, 17747, 792},
	{17747, 17827, 793},
	{17827, 17907, 794},
	{17907, 17987, 795},
	{17987, 18067, 796},
	{18067, 18147, 797},
	{18147, 18227, 798},
	{18227, 18307, 799},
	{18307, 18313, 800},
	{18313, 18319, 801},
	{18319, 18325, 802},
	{18325, 18331, 803},
	{18331, 18337, 804},
	{18337, 18343, 805},
	{18343, 18349, 806},
	{18349, 18355, 807},
	{18355, 18361, 808},
	{18361, 18367, 809},
	{18367, 18373, 810},
	{18373, 18379, 811},
	{18379, 18385, 812},
	{18385, 18391, 813},
	{18391, 18397, 814},
	{18397, 18403, 815},
	{18403, 18404, 816},
	{18404, 18405, 817},
	{18405, 18406, 818},
	{18406, 18407, 819},
	{18407, 18408, 820},
	{18408, 18409, 821},
	{18409, 18410, 822},
	{18410, 18411, 823},
	{18411, 18412, 824},
	{18412, 18413, 825},
	{18413, 18414, 826},
	{18414, 18415, 827},
	{18415, 18416, 828},
	{18416, 18417, 829},
	{18417, 18418, 830},
	{18418, 18419, 831},
	{18419, 18445, 832},
	{18445, 18446, 833},
	{18446, 18447, 834},
	{18447, 18459, 835},
	{18459, 18462, 836},
	{18462, 18494, 837},
	{18494, 18495, 838},
	{18495, 18496, 839},
	{18496, 18497, 840},
	{18497, 18498, 841},
	{18498, 18499, 842},
	{18499, 18500, 843},
	{18500, 18501, 844},
	{18501, 18502, 845},
	{18502, 18503, 846},
	{18503, 18504, 847},
	{18504, 18506, 848},
	{18506, 18508, 849},
	{18508, 18510, 850},
	{18510, 18512, 851},
	{18512, 18514, 852},
	{18514, 18516, 853},
	{18516, 18518, 854},
	{18518, 18520, 855},
	{18520, 18522, 856},
	{18522, 18524, 857},
	{18524, 18526, 858},
	{18526, 18528, 859},
	{18528, 18530, 860},
	{18530, 18532, 861},
	{18532, 18534, 862},
	{18534, 18536, 863},
	{18536, 18538, 864},
	{18538, 18540, 865},
	{18540, 18542, 866},
	{18542, 18544, 867},
	{18544, 18552, 868},
	{18552, 18560, 869},
	{18560, 18568, 870},
	{18568, 18576, 871},
	{18576, 18584, 872},
	{18584, 18592, 873},
	{18592, 18600, 874},
	{18600, 18608, 875},
	{18608, 18616, 876},
	{18616, 18624, 877},
	{18624, 18632, 878},
	{18632, 18633, 879},
	{18633, 18635, 880},
	{18635, 18636, 881},
	{18636, 18648, 882},
	{18648, 18649, 883},
	{18649, 18650, 884},
	{18650, 18651, 885},
	{18651, 18653, 886},
	{18653, 18733, 887},
	{18733, 18813, 888},
	{18813, 18893, 889},
	{18893, 18973, 890},
	{18973, 19053, 891},
	{19053, 19133, 892},
	{19133, 19213, 893},
	{19213, 19293, 894},
	{19293, 19373, 895},
	{19373, 19453, 896},
	{19453, 19533, 897},
	{19533, 19613, 898},
	{19613, 19693, 899},
	{19693, 19773, 900},
	{19773, 19779, 901},
	{19779, 19785, 902},
	{19785, 19791, 903},
	{19791, 19797, 904},
	{19797, 19803, 905},
	{19803, 19809, 906},
	{19809, 19815, 907},
	{19815, 19821, 908},
	{19821, 19827, 909},
	{19827, 19833, 910},
	{19833, 19839, 911},
	{19839, 19845, 912},
	{19845, 19851, 913},
	{19851, 20175, 914},
	{20175, 20499, 915},
	{20499, 20823, 916},
	{20823, 21147, 917},
	{21147, 21471, 918},
	{21471, 21795, 919},
	{21795, 22119, 920},
	{22119, 22443, 921},
	{22443, 22767, 922},
	{22767, 23091, 923},
	{23091, 23415, 924},
	{23415, 23739, 925},
	{23739, 24063, 926},
	{24063, 24095, 927},
	{24095, 24099, 928},
	{24099, 24111, 929},
	{24111, 24119, 930},
	{24119, 24127, 931},
	{24127, 24128, 932},
	{24128, 24129, 933},
	{24129, 24141, 934},
	{24141, 24157, 935},
	{24157, 24158, 936},
	{24158, 24162, 937},
	{24162, 24194, 938},
	{24194, 24198, 939},
	{24198, 24202, 940},
	{24202, 24206, 941},
	{24206, 24210, 942},
	{24210, 24214, 943},
	{24214, 24218, 944},
	{24218, 24222, 945},
	{24222, 24226, 946},
	{24226, 24230, 947},
	{24230, 24234, 948},
	{24234, 24266, 949},
	{24266, 24298, 950},
	{24298, 24302, 951},
	{24302, 24305, 952},
	{24305, 24308, 953},
	{24308, 24311, 954},
	{24311, 24314, 955},
	{24314, 24315, 956},
	{24315, 24316, 957},
	{24316, 24317, 958},
	{24317, 24318, 959},
	{24318, 24319, 960},
	{24319, 24322, 961},
	{24322, 24325, 962},
	{24325, 24328, 963},
	{24328, 24331, 964},
	{24331, 24332, 965},
	{24332, 24333, 966},
	{24333, 24334, 967},
	{24334, 24360, 968},
	{24360, 24361, 969},
	{24361, 24387, 970},
	{24387, 24388, 971},
	{24388, 24389, 972},
	{24389, 24390, 973},
	{24390, 24391, 974},
	{24391, 24397, 975},
	{24397, 24403, 976},
	{24403, 24405, 977},
	{24405, 24407, 978},
	{24407, 24439, 979},
	{24439, 24471, 980},
	{24471, 24535, 981},
	{24535, 24599, 982},
	{24599, 24631, 983},
	{24631, 24663, 984},
	{24663, 24743, 985},
	{24743, 24823, 986},
	{24823, 24847, 987},
	{24847, 24871, 988},
	{24871, 24935, 989},
	{24935, 24999, 990},
	{24999, 25031, 991},
	{25031, 25063, 992},
	{25063, 25071, 993},
	{25071, 25079, 994},
	{25079, 25083, 995},
	{25083, 25095, 996},
	{25095, 25099, 997},
	{25099, 25100, 998},
	{25100, 25109, 999},
	{25109, 25125, 1000},
	{25125, 25149, 1001},
	{25149, 25173, 1002},
	{25173, 25174, 1003},
	{25174, 25175, 1004},
	{25175, 25176, 1005},
	{25176, 25177, 1006},
	{25177, 25178, 1007},
	{25178, 25183, 1008},
	{25183, 25184, 1009},
	{25184, 25185, 1010},
	{25185, 25186, 1011},
	{25186, 25187, 1012},
	{25187, 25188, 1013},
	{25188, 25189, 1014},
	{25189, 25269, 1015},
	{25269, 25593, 1016},
	{25593, 25599, 1017},
	{25599, 25600, 1018},
	{25600, 25601, 1019},
	{25601, 25602, 1020},
	{25602, 25603, 1021},
	{25603, 25609, 1022},
	{25609, 25689, 1023},
	{25689, 26013, 1024},
	{26013, 26014, 1025},
	{26014, 26094, 1026},
	{26094, 26100, 1027},
	{26100, 26102, 1028},
	{26102, 26126, 1029},
	{26126, 26450, 1030},
	{26450, 26451, 1031},
	{26451, 26452, 1032},
	{26452, 26453, 1033},
	{26453, 26469, 1034},
	{26469, 26485, 1035},
	{26485, 26501, 1036},
	{26501, 26517, 1037},
	{26517, 26533, 1038},
	{26533, 26549, 1039},
	{26549, 26565, 1040},
	{26565, 26581, 1041},
	{26581, 26597, 1042},
	{26597, 26613, 1043},
	{26613, 26629, 1044},
	{26629, 26645, 1045},
	{26645, 26661, 1046},
	{26661, 26677, 1047},
	{26677, 26693, 1048},
	{26693, 26709, 1049},
	{26709, 26725, 1050},
	{26725, 26727, 1051},
	{26727, 26729, 1052},
	{26729, 26731, 1053},
	{26731, 26733, 1054},
	{26733, 26735, 1055},
	{26735, 26737, 1056},
	{26737, 26739, 1057},
	{26739, 26741, 1058},
	{26741, 26743, 1059},
	{26743, 26745, 1060},
	{26745, 26747, 1061},
	{26747, 26749, 1062},
	{26749, 26751, 1063},
	{26751, 26753, 1064},
	{26753, 26755, 1065},
	{26755, 26757, 1066},
	{26757, 26759, 1067},
	{26759, 26760, 1068},
	{26760, 26761, 1069},
	{26761, 26773, 1070},
	{26773, 26785, 1071},
	{26785, 26797, 1072},
	{26797, 26809, 1073},
	{26809, 26810, 1074},
	{26810, 26816, 1075},
	{26816, 26896, 1076},
	{26896, 27220, 1077},
	{27220, 27221, 1078},
	{27221, 27227, 1079},
	{27227, 27307, 1080},
	{27307, 27631, 1081},
	{27631, 27632, 1082},
	{27632, 27633, 1083},
	{27633, 27639, 1084},
	{27639, 27719, 1085},
	{27719, 28043, 1086},
	{28043, 28044, 1087},
	{28044, 28045, 1088},
	{28045, 28050, 1089},
	{28050, 28056, 1090},
	{28056, 28136, 1091},
	{28136, 28460, 1092},
	{28460, 28461, 1093},
	{28461, 28467, 1094},
	{28467, 28547, 1095},
	{28547, 28871, 1096},
	{28871, 28872, 1097},
	{28872, 28878, 1098},
	{28878, 28958, 1099},
	{28958, 29282, 1100},
	{29282, 29283, 1101},
	{29283, 29284, 1102},
	{29284, 29290, 1103},
	{29290, 29370, 1104},
	{29370, 29694, 1105},
	{29694, 29695, 1106},
	{29695, 29701, 1107},
	{29701, 29781, 1108},
	{29781, 30105, 1109},
	{30105, 30106, 1110},
	{30106, 30112, 1111},
	{30112, 30192, 1112},
	{30192, 30516, 1113},
	{30516, 30517, 1114},
	{30517, 30518, 1115},
	{30518, 30519, 1116},
	{30519, 30520, 1117},
	{30520, 30616, 1118},
	{30616, 31000, 1119},
	{31000, 31001, 1120},
	{31001, 31129, 1121},
	{31129, 31131, 1122},
	{31131, 31139, 1123},
	{31139, 31140, 1124},
	{31140, 31141, 1125},
	{31141, 31142, 1126},
	{31142, 31143, 1127},
	{31143, 31144, 1128},
	{31144, 31145, 1129},
	{31145, 31146, 1130},
	{31146, 31147, 1131},
	{31147, 31148, 1132},
	{31148, 31149, 1133},
	{31149, 31150, 1134},
	{31150, 31151, 1135},
	{31151, 31152, 1136},
	{31152, 31153, 1137},
	{31153, 31154, 1138},
	{31154, 31155, 1139},
	{31155, 31156, 1140},
	{31156, 31157, 1141},
	{31157, 31158, 1142},
	{31158, 31159, 1143},
	{31159, 31160, 1144},
	{31160, 31161, 1145},
	{31161, 31162, 1146},
	{31162, 31163, 1147},
	{31163, 31164, 1148},
	{31164, 31165, 1149},
	{31165, 31245, 1150},
	{31245, 31325, 1151},
	{31325, 31405, 1152},
	{31405, 31485, 1153},
	{31485, 31565, 1154},
	{31565, 31645, 1155},
	{31645, 31725, 1156},
	{31725, 31805, 1157},
	{31805, 31811, 1158},
	{31811, 31817, 1159},
	{31817, 31823, 1160},
	{31823, 31829, 1161},
	{31829, 31835, 1162},
	{31835, 31841, 1163},
	{31841, 31847, 1164},
	{31847, 31853, 1165},
	{31853, 31917, 1166},
	{31917, 31981, 1167},
	{31981, 32045, 1168},
	{32045, 32109, 1169},
	{32109, 32173, 1170},
	{32173, 32237, 1171},
	{32237, 32301, 1172},
	{32301, 32365, 1173},
	{32365, 32429, 1174},
	{32429, 32493, 1175},
	{32493, 32557, 1176},
	{32557, 32621, 1177},
	{32621, 32685, 1178},
	{32685, 32749, 1179},
	{32749, 32813, 1180},
	{32813, 32877, 1181},
	{32877, 32879, 1182},
	{32879, 32881, 1183},
	{32881, 32883, 1184},
	{32883, 32885, 1185},
	{32885, 32887, 1186},
	{32887, 32889, 1187},
	{32889, 32891, 1188},
	{32891, 32893, 1189},
	{32893, 32897, 1190},
	{32897, 32901, 1191},
	{32901, 32905, 1192},
	{32905, 32909, 1193},
	{32909, 32913, 1194},
	{32913, 32917, 1195},
	{32917, 32921, 1196},
	{32921, 32925, 1197},
	{32925, 32949, 1198},
	{32949, 32973, 1199},
	{32973, 32997, 1200},
	{32997, 33021, 1201},
	{33021, 33045, 1202},
	{33045, 33069, 1203},
	{33069, 33093, 1204},
	{33093, 33117, 1205},
	{33117, 33149, 1206},
	{33149, 33181, 1207},
	{33181, 33213, 1208},
	{33213, 33245, 1209},
	{33245, 33277, 1210},
	{33277, 33309, 1211},
	{33309, 33341, 1212},
	{33341, 33373, 1213},
	{33373, 33397, 1214},
	{33397, 33421, 1215},
	{33421, 33445, 1216},
	{33445, 33469, 1217},
	{33469, 33493, 1218},
	{33493, 33517, 1219},
	{33517, 33541, 1220},
	{33541, 33565, 1221},
	{33565, 33566, 1222},
	{33566, 33586, 1223},
	{33586, 33606, 1224},
	{33606, 33658, 1225},
	{33658, 33660, 1226},
	{33660, 33661, 1227},
	{33661, 33662, 1228},
	{33662, 33663, 1229},
	{33663, 33664, 1230},
	{33664, 33680, 1231},
	{33680, 33696, 1232},
	{33696, 33712, 1233},
	{33712, 33713, 1234},
	{33713, 33745, 1235},
	{33745, 33753, 1236},
	{33753, 33769, 1237},
	{33769, 33771, 1238},
	{33771, 33772, 1239},
	{33772, 33773, 1240},
	{33773, 33776, 1241},
	{33776, 33777, 1242},
	{33777, 33857, 1243},
	{33857, 33863, 1244},
	{33863, 34187, 1245},
	{34187, 34188, 1246},
	{34188, 34268, 1247},
	{34268, 34274, 1248},
	{34274, 34598, 1249},
	{34598, 34599, 1250},
	{34599, 34679, 1251},
	{34679, 34685, 1252},
	{34685, 35009, 1253},
	{35009, 35010, 1254},
	{35010, 35090, 1255},
	{35090, 35096, 1256},
	{35096, 35420, 1257},
	{35420, 35421, 1258},
	{35421, 35422, 1259},
	{35422, 35423, 1260},
	{35423, 35426, 1261},
	{35426, 35427, 1262},
	{35427, 35428, 1263},
	{35428, 35429, 1264},
	{35429, 35430, 1265},
	{35430, 35431, 1266},
	{35431, 35432, 1267},
	{35432, 35435, 1268},
	{35435, 35438, 1269},
	{35438, 35441, 1270},
	{35441, 35442, 1271},
	{35442, 35443, 1272},
	{35443, 35459, 1273},
	{35459, 35507, 1274},
	{35507, 35519, 1275},
	{35519, 35551, 1276},
	{35551, 35553, 1277},
	{35553, 35554, 1278},
	{35554, 35716, 1279},
	{35716, 35718, 1280},
	{35718, 35719, 1281},
	{35719, 35720, 1282},
	{35720, 35721, 1283},
	{35721, 35722, 1284},
	{35722, 35723, 1285},
}
