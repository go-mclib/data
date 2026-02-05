// Code generated for Minecraft 1.21.11 (Protocol 774); DO NOT EDIT.

package entities

// Entity type protocol IDs
const (
	AcaciaBoat           = 0
	AcaciaChestBoat      = 1
	Allay                = 2
	AreaEffectCloud      = 3
	Armadillo            = 4
	ArmorStand           = 5
	Arrow                = 6
	Axolotl              = 7
	BambooChestRaft      = 8
	BambooRaft           = 9
	Bat                  = 10
	Bee                  = 11
	BirchBoat            = 12
	BirchChestBoat       = 13
	Blaze                = 14
	BlockDisplay         = 15
	Bogged               = 16
	Breeze               = 17
	BreezeWindCharge     = 18
	Camel                = 19
	CamelHusk            = 20
	Cat                  = 21
	CaveSpider           = 22
	CherryBoat           = 23
	CherryChestBoat      = 24
	ChestMinecart        = 25
	Chicken              = 26
	Cod                  = 27
	CommandBlockMinecart = 29
	CopperGolem          = 28
	Cow                  = 30
	Creaking             = 31
	Creeper              = 32
	DarkOakBoat          = 33
	DarkOakChestBoat     = 34
	Dolphin              = 35
	Donkey               = 36
	DragonFireball       = 37
	Drowned              = 38
	Egg                  = 39
	ElderGuardian        = 40
	EndCrystal           = 45
	EnderDragon          = 43
	EnderPearl           = 44
	Enderman             = 41
	Endermite            = 42
	Evoker               = 46
	EvokerFangs          = 47
	ExperienceBottle     = 48
	ExperienceOrb        = 49
	EyeOfEnder           = 50
	FallingBlock         = 51
	Fireball             = 52
	FireworkRocket       = 53
	FishingBobber        = 156
	Fox                  = 54
	Frog                 = 55
	FurnaceMinecart      = 56
	Ghast                = 57
	Giant                = 59
	GlowItemFrame        = 60
	GlowSquid            = 61
	Goat                 = 62
	Guardian             = 63
	HappyGhast           = 58
	Hoglin               = 64
	HopperMinecart       = 65
	Horse                = 66
	Husk                 = 67
	Illusioner           = 68
	Interaction          = 69
	IronGolem            = 70
	Item                 = 71
	ItemDisplay          = 72
	ItemFrame            = 73
	JungleBoat           = 74
	JungleChestBoat      = 75
	LeashKnot            = 76
	LightningBolt        = 77
	LingeringPotion      = 106
	Llama                = 78
	LlamaSpit            = 79
	MagmaCube            = 80
	MangroveBoat         = 81
	MangroveChestBoat    = 82
	Mannequin            = 83
	Marker               = 84
	Minecart             = 85
	Mooshroom            = 86
	Mule                 = 87
	Nautilus             = 88
	OakBoat              = 89
	OakChestBoat         = 90
	Ocelot               = 91
	OminousItemSpawner   = 92
	Painting             = 93
	PaleOakBoat          = 94
	PaleOakChestBoat     = 95
	Panda                = 96
	Parched              = 97
	Parrot               = 98
	Phantom              = 99
	Pig                  = 100
	Piglin               = 101
	PiglinBrute          = 102
	Pillager             = 103
	Player               = 155
	PolarBear            = 104
	Pufferfish           = 107
	Rabbit               = 108
	Ravager              = 109
	Salmon               = 110
	Sheep                = 111
	Shulker              = 112
	ShulkerBullet        = 113
	Silverfish           = 114
	Skeleton             = 115
	SkeletonHorse        = 116
	Slime                = 117
	SmallFireball        = 118
	Sniffer              = 119
	SnowGolem            = 121
	Snowball             = 120
	SpawnerMinecart      = 122
	SpectralArrow        = 123
	Spider               = 124
	SplashPotion         = 105
	SpruceBoat           = 125
	SpruceChestBoat      = 126
	Squid                = 127
	Stray                = 128
	Strider              = 129
	Tadpole              = 130
	TextDisplay          = 131
	Tnt                  = 132
	TntMinecart          = 133
	TraderLlama          = 134
	Trident              = 135
	TropicalFish         = 136
	Turtle               = 137
	Vex                  = 138
	Villager             = 139
	Vindicator           = 140
	WanderingTrader      = 141
	Warden               = 142
	WindCharge           = 143
	Witch                = 144
	Wither               = 145
	WitherSkeleton       = 146
	WitherSkull          = 147
	Wolf                 = 148
	Zoglin               = 149
	Zombie               = 150
	ZombieHorse          = 151
	ZombieNautilus       = 152
	ZombieVillager       = 153
	ZombifiedPiglin      = 154
)

var entityByName = map[string]int32{
	"minecraft:acacia_boat":            0,
	"minecraft:acacia_chest_boat":      1,
	"minecraft:allay":                  2,
	"minecraft:area_effect_cloud":      3,
	"minecraft:armadillo":              4,
	"minecraft:armor_stand":            5,
	"minecraft:arrow":                  6,
	"minecraft:axolotl":                7,
	"minecraft:bamboo_chest_raft":      8,
	"minecraft:bamboo_raft":            9,
	"minecraft:bat":                    10,
	"minecraft:bee":                    11,
	"minecraft:birch_boat":             12,
	"minecraft:birch_chest_boat":       13,
	"minecraft:blaze":                  14,
	"minecraft:block_display":          15,
	"minecraft:bogged":                 16,
	"minecraft:breeze":                 17,
	"minecraft:breeze_wind_charge":     18,
	"minecraft:camel":                  19,
	"minecraft:camel_husk":             20,
	"minecraft:cat":                    21,
	"minecraft:cave_spider":            22,
	"minecraft:cherry_boat":            23,
	"minecraft:cherry_chest_boat":      24,
	"minecraft:chest_minecart":         25,
	"minecraft:chicken":                26,
	"minecraft:cod":                    27,
	"minecraft:command_block_minecart": 29,
	"minecraft:copper_golem":           28,
	"minecraft:cow":                    30,
	"minecraft:creaking":               31,
	"minecraft:creeper":                32,
	"minecraft:dark_oak_boat":          33,
	"minecraft:dark_oak_chest_boat":    34,
	"minecraft:dolphin":                35,
	"minecraft:donkey":                 36,
	"minecraft:dragon_fireball":        37,
	"minecraft:drowned":                38,
	"minecraft:egg":                    39,
	"minecraft:elder_guardian":         40,
	"minecraft:end_crystal":            45,
	"minecraft:ender_dragon":           43,
	"minecraft:ender_pearl":            44,
	"minecraft:enderman":               41,
	"minecraft:endermite":              42,
	"minecraft:evoker":                 46,
	"minecraft:evoker_fangs":           47,
	"minecraft:experience_bottle":      48,
	"minecraft:experience_orb":         49,
	"minecraft:eye_of_ender":           50,
	"minecraft:falling_block":          51,
	"minecraft:fireball":               52,
	"minecraft:firework_rocket":        53,
	"minecraft:fishing_bobber":         156,
	"minecraft:fox":                    54,
	"minecraft:frog":                   55,
	"minecraft:furnace_minecart":       56,
	"minecraft:ghast":                  57,
	"minecraft:giant":                  59,
	"minecraft:glow_item_frame":        60,
	"minecraft:glow_squid":             61,
	"minecraft:goat":                   62,
	"minecraft:guardian":               63,
	"minecraft:happy_ghast":            58,
	"minecraft:hoglin":                 64,
	"minecraft:hopper_minecart":        65,
	"minecraft:horse":                  66,
	"minecraft:husk":                   67,
	"minecraft:illusioner":             68,
	"minecraft:interaction":            69,
	"minecraft:iron_golem":             70,
	"minecraft:item":                   71,
	"minecraft:item_display":           72,
	"minecraft:item_frame":             73,
	"minecraft:jungle_boat":            74,
	"minecraft:jungle_chest_boat":      75,
	"minecraft:leash_knot":             76,
	"minecraft:lightning_bolt":         77,
	"minecraft:lingering_potion":       106,
	"minecraft:llama":                  78,
	"minecraft:llama_spit":             79,
	"minecraft:magma_cube":             80,
	"minecraft:mangrove_boat":          81,
	"minecraft:mangrove_chest_boat":    82,
	"minecraft:mannequin":              83,
	"minecraft:marker":                 84,
	"minecraft:minecart":               85,
	"minecraft:mooshroom":              86,
	"minecraft:mule":                   87,
	"minecraft:nautilus":               88,
	"minecraft:oak_boat":               89,
	"minecraft:oak_chest_boat":         90,
	"minecraft:ocelot":                 91,
	"minecraft:ominous_item_spawner":   92,
	"minecraft:painting":               93,
	"minecraft:pale_oak_boat":          94,
	"minecraft:pale_oak_chest_boat":    95,
	"minecraft:panda":                  96,
	"minecraft:parched":                97,
	"minecraft:parrot":                 98,
	"minecraft:phantom":                99,
	"minecraft:pig":                    100,
	"minecraft:piglin":                 101,
	"minecraft:piglin_brute":           102,
	"minecraft:pillager":               103,
	"minecraft:player":                 155,
	"minecraft:polar_bear":             104,
	"minecraft:pufferfish":             107,
	"minecraft:rabbit":                 108,
	"minecraft:ravager":                109,
	"minecraft:salmon":                 110,
	"minecraft:sheep":                  111,
	"minecraft:shulker":                112,
	"minecraft:shulker_bullet":         113,
	"minecraft:silverfish":             114,
	"minecraft:skeleton":               115,
	"minecraft:skeleton_horse":         116,
	"minecraft:slime":                  117,
	"minecraft:small_fireball":         118,
	"minecraft:sniffer":                119,
	"minecraft:snow_golem":             121,
	"minecraft:snowball":               120,
	"minecraft:spawner_minecart":       122,
	"minecraft:spectral_arrow":         123,
	"minecraft:spider":                 124,
	"minecraft:splash_potion":          105,
	"minecraft:spruce_boat":            125,
	"minecraft:spruce_chest_boat":      126,
	"minecraft:squid":                  127,
	"minecraft:stray":                  128,
	"minecraft:strider":                129,
	"minecraft:tadpole":                130,
	"minecraft:text_display":           131,
	"minecraft:tnt":                    132,
	"minecraft:tnt_minecart":           133,
	"minecraft:trader_llama":           134,
	"minecraft:trident":                135,
	"minecraft:tropical_fish":          136,
	"minecraft:turtle":                 137,
	"minecraft:vex":                    138,
	"minecraft:villager":               139,
	"minecraft:vindicator":             140,
	"minecraft:wandering_trader":       141,
	"minecraft:warden":                 142,
	"minecraft:wind_charge":            143,
	"minecraft:witch":                  144,
	"minecraft:wither":                 145,
	"minecraft:wither_skeleton":        146,
	"minecraft:wither_skull":           147,
	"minecraft:wolf":                   148,
	"minecraft:zoglin":                 149,
	"minecraft:zombie":                 150,
	"minecraft:zombie_horse":           151,
	"minecraft:zombie_nautilus":        152,
	"minecraft:zombie_villager":        153,
	"minecraft:zombified_piglin":       154,
}

var entityByID = map[int32]string{
	0:   "minecraft:acacia_boat",
	1:   "minecraft:acacia_chest_boat",
	2:   "minecraft:allay",
	3:   "minecraft:area_effect_cloud",
	4:   "minecraft:armadillo",
	5:   "minecraft:armor_stand",
	6:   "minecraft:arrow",
	7:   "minecraft:axolotl",
	8:   "minecraft:bamboo_chest_raft",
	9:   "minecraft:bamboo_raft",
	10:  "minecraft:bat",
	11:  "minecraft:bee",
	12:  "minecraft:birch_boat",
	13:  "minecraft:birch_chest_boat",
	14:  "minecraft:blaze",
	15:  "minecraft:block_display",
	16:  "minecraft:bogged",
	17:  "minecraft:breeze",
	18:  "minecraft:breeze_wind_charge",
	19:  "minecraft:camel",
	20:  "minecraft:camel_husk",
	21:  "minecraft:cat",
	22:  "minecraft:cave_spider",
	23:  "minecraft:cherry_boat",
	24:  "minecraft:cherry_chest_boat",
	25:  "minecraft:chest_minecart",
	26:  "minecraft:chicken",
	27:  "minecraft:cod",
	29:  "minecraft:command_block_minecart",
	28:  "minecraft:copper_golem",
	30:  "minecraft:cow",
	31:  "minecraft:creaking",
	32:  "minecraft:creeper",
	33:  "minecraft:dark_oak_boat",
	34:  "minecraft:dark_oak_chest_boat",
	35:  "minecraft:dolphin",
	36:  "minecraft:donkey",
	37:  "minecraft:dragon_fireball",
	38:  "minecraft:drowned",
	39:  "minecraft:egg",
	40:  "minecraft:elder_guardian",
	45:  "minecraft:end_crystal",
	43:  "minecraft:ender_dragon",
	44:  "minecraft:ender_pearl",
	41:  "minecraft:enderman",
	42:  "minecraft:endermite",
	46:  "minecraft:evoker",
	47:  "minecraft:evoker_fangs",
	48:  "minecraft:experience_bottle",
	49:  "minecraft:experience_orb",
	50:  "minecraft:eye_of_ender",
	51:  "minecraft:falling_block",
	52:  "minecraft:fireball",
	53:  "minecraft:firework_rocket",
	156: "minecraft:fishing_bobber",
	54:  "minecraft:fox",
	55:  "minecraft:frog",
	56:  "minecraft:furnace_minecart",
	57:  "minecraft:ghast",
	59:  "minecraft:giant",
	60:  "minecraft:glow_item_frame",
	61:  "minecraft:glow_squid",
	62:  "minecraft:goat",
	63:  "minecraft:guardian",
	58:  "minecraft:happy_ghast",
	64:  "minecraft:hoglin",
	65:  "minecraft:hopper_minecart",
	66:  "minecraft:horse",
	67:  "minecraft:husk",
	68:  "minecraft:illusioner",
	69:  "minecraft:interaction",
	70:  "minecraft:iron_golem",
	71:  "minecraft:item",
	72:  "minecraft:item_display",
	73:  "minecraft:item_frame",
	74:  "minecraft:jungle_boat",
	75:  "minecraft:jungle_chest_boat",
	76:  "minecraft:leash_knot",
	77:  "minecraft:lightning_bolt",
	106: "minecraft:lingering_potion",
	78:  "minecraft:llama",
	79:  "minecraft:llama_spit",
	80:  "minecraft:magma_cube",
	81:  "minecraft:mangrove_boat",
	82:  "minecraft:mangrove_chest_boat",
	83:  "minecraft:mannequin",
	84:  "minecraft:marker",
	85:  "minecraft:minecart",
	86:  "minecraft:mooshroom",
	87:  "minecraft:mule",
	88:  "minecraft:nautilus",
	89:  "minecraft:oak_boat",
	90:  "minecraft:oak_chest_boat",
	91:  "minecraft:ocelot",
	92:  "minecraft:ominous_item_spawner",
	93:  "minecraft:painting",
	94:  "minecraft:pale_oak_boat",
	95:  "minecraft:pale_oak_chest_boat",
	96:  "minecraft:panda",
	97:  "minecraft:parched",
	98:  "minecraft:parrot",
	99:  "minecraft:phantom",
	100: "minecraft:pig",
	101: "minecraft:piglin",
	102: "minecraft:piglin_brute",
	103: "minecraft:pillager",
	155: "minecraft:player",
	104: "minecraft:polar_bear",
	107: "minecraft:pufferfish",
	108: "minecraft:rabbit",
	109: "minecraft:ravager",
	110: "minecraft:salmon",
	111: "minecraft:sheep",
	112: "minecraft:shulker",
	113: "minecraft:shulker_bullet",
	114: "minecraft:silverfish",
	115: "minecraft:skeleton",
	116: "minecraft:skeleton_horse",
	117: "minecraft:slime",
	118: "minecraft:small_fireball",
	119: "minecraft:sniffer",
	121: "minecraft:snow_golem",
	120: "minecraft:snowball",
	122: "minecraft:spawner_minecart",
	123: "minecraft:spectral_arrow",
	124: "minecraft:spider",
	105: "minecraft:splash_potion",
	125: "minecraft:spruce_boat",
	126: "minecraft:spruce_chest_boat",
	127: "minecraft:squid",
	128: "minecraft:stray",
	129: "minecraft:strider",
	130: "minecraft:tadpole",
	131: "minecraft:text_display",
	132: "minecraft:tnt",
	133: "minecraft:tnt_minecart",
	134: "minecraft:trader_llama",
	135: "minecraft:trident",
	136: "minecraft:tropical_fish",
	137: "minecraft:turtle",
	138: "minecraft:vex",
	139: "minecraft:villager",
	140: "minecraft:vindicator",
	141: "minecraft:wandering_trader",
	142: "minecraft:warden",
	143: "minecraft:wind_charge",
	144: "minecraft:witch",
	145: "minecraft:wither",
	146: "minecraft:wither_skeleton",
	147: "minecraft:wither_skull",
	148: "minecraft:wolf",
	149: "minecraft:zoglin",
	150: "minecraft:zombie",
	151: "minecraft:zombie_horse",
	152: "minecraft:zombie_nautilus",
	153: "minecraft:zombie_villager",
	154: "minecraft:zombified_piglin",
}
