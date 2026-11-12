// Code generated for Minecraft 26.3 (Protocol 777); DO NOT EDIT.

package entities

// Entity metadata field indices.

// AbstractCubeMob metadata indices
const (
	AbstractCubeMobIndexSize = 18
)

// AbstractFish metadata indices
const (
	AbstractFishIndexFromBucket = 16
)

// AbstractPiglin metadata indices
const (
	AbstractPiglinIndexImmuneToZombification = 16
)

// AbstractVillager metadata indices
const (
	AbstractVillagerIndexUnhappyCounter = 18
)

// AgeableMob metadata indices
const (
	AgeableMobIndexIsBaby    = 16
	AgeableMobIndexAgeLocked = 17
)

// ArmorStand metadata indices
const (
	ArmorStandIndexArmorStandFlags = 15
	ArmorStandIndexHeadPose        = 16
	ArmorStandIndexBodyPose        = 17
	ArmorStandIndexLeftArmPose     = 18
	ArmorStandIndexRightArmPose    = 19
	ArmorStandIndexLeftLegPose     = 20
	ArmorStandIndexRightLegPose    = 21
)

// Arrow metadata indices
const (
	ArrowIndexArrowFlags  = 8
	ArrowIndexPierceLevel = 9
	ArrowIndexInGround    = 10
)

// Avatar metadata indices
const (
	AvatarIndexMainHand  = 15
	AvatarIndexSkinParts = 16
)

// Bat metadata indices
const (
	BatIndexBatFlags = 16
)

// Blaze metadata indices
const (
	BlazeIndexBlazeFlags = 16
)

// Boat metadata indices
const (
	BoatIndexHurtTime    = 8
	BoatIndexHurtDir     = 9
	BoatIndexDamage      = 10
	BoatIndexPaddleLeft  = 11
	BoatIndexPaddleRight = 12
	BoatIndexBubbleTime  = 13
)

// Cat metadata indices
const (
	CatIndexVariant      = 20
	CatIndexIsLying      = 21
	CatIndexIsRelaxed    = 22
	CatIndexCollarColor  = 23
	CatIndexSoundVariant = 24
)

// Chicken metadata indices
const (
	ChickenIndexVariant      = 18
	ChickenIndexSoundVariant = 19
)

// CommandBlockMinecart metadata indices
const (
	CommandBlockMinecartIndexCommand    = 14
	CommandBlockMinecartIndexLastOutput = 15
)

// Cow metadata indices
const (
	CowIndexVariant      = 18
	CowIndexSoundVariant = 19
)

// Creeper metadata indices
const (
	CreeperIndexSwellDir  = 16
	CreeperIndexIsPowered = 17
	CreeperIndexIsIgnited = 18
)

// Cushion metadata indices
const (
	CushionIndexColor = 8
)

// EnderDragon metadata indices
const (
	EnderDragonIndexPhase = 16
)

// Enderman metadata indices
const (
	EndermanIndexCarriedBlock = 16
	EndermanIndexIsScreaming  = 17
	EndermanIndexIsStaring    = 18
)

// Entity metadata indices
const (
	EntityIndexFlags             = 0
	EntityIndexAirSupply         = 1
	EntityIndexCustomName        = 2
	EntityIndexCustomNameVisible = 3
	EntityIndexSilent            = 4
	EntityIndexNoGravity         = 5
	EntityIndexPose              = 6
	EntityIndexTicksFrozen       = 7
)

// FallingBlock metadata indices
const (
	FallingBlockIndexSpawnPos = 8
)

// Fireball metadata indices
const (
	FireballIndexItem = 8
)

// Ghast metadata indices
const (
	GhastIndexIsCharging = 16
)

// HappyGhast metadata indices
const (
	HappyGhastIndexIsLeashHolder = 18
	HappyGhastIndexStaysStill    = 19
)

// ItemEntity metadata indices
const (
	ItemEntityIndexItem = 8
)

// ItemFrame metadata indices
const (
	ItemFrameIndexItem     = 8
	ItemFrameIndexRotation = 9
)

// LivingEntity metadata indices
const (
	LivingEntityIndexLivingFlags     = 8
	LivingEntityIndexHealth          = 9
	LivingEntityIndexEffectParticles = 10
	LivingEntityIndexEffectAmbient   = 11
	LivingEntityIndexArrowCount      = 12
	LivingEntityIndexStingerCount    = 13
	LivingEntityIndexSleepingPos     = 14
)

// Minecart metadata indices
const (
	MinecartIndexHurtTime      = 8
	MinecartIndexHurtDir       = 9
	MinecartIndexDamage        = 10
	MinecartIndexDisplayBlock  = 11
	MinecartIndexDisplayOffset = 12
	MinecartIndexCustomDisplay = 13
)

// Mob metadata indices
const (
	MobIndexMobFlags = 15
)

// Phantom metadata indices
const (
	PhantomIndexSize = 16
)

// Pig metadata indices
const (
	PigIndexBoostTime    = 18
	PigIndexVariant      = 19
	PigIndexSoundVariant = 20
)

// Piglin metadata indices
const (
	PiglinIndexIsBaby             = 17
	PiglinIndexIsChargingCrossbow = 18
	PiglinIndexIsDancing          = 19
)

// Player metadata indices
const (
	PlayerIndexAdditionalHearts    = 17
	PlayerIndexScore               = 18
	PlayerIndexLeftShoulderEntity  = 19
	PlayerIndexRightShoulderEntity = 20
)

// Pufferfish metadata indices
const (
	PufferfishIndexPuffState = 17
)

// Salmon metadata indices
const (
	SalmonIndexVariant = 17
)

// Sheep metadata indices
const (
	SheepIndexColorFlags = 18
)

// Shulker metadata indices
const (
	ShulkerIndexAttachFace = 16
	ShulkerIndexPeek       = 17
	ShulkerIndexColor      = 18
)

// Skeleton metadata indices
const (
	SkeletonIndexIsConverting = 16
)

// Sniffer metadata indices
const (
	SnifferIndexState          = 18
	SnifferIndexDropSeedAtTick = 19
)

// Spider metadata indices
const (
	SpiderIndexSpiderFlags = 16
)

// TamableAnimal metadata indices
const (
	TamableAnimalIndexTamableFlags = 18
	TamableAnimalIndexOwner        = 19
)

// Trident metadata indices
const (
	TridentIndexLoyalty = 11
	TridentIndexFoil    = 12
)

// Villager metadata indices
const (
	VillagerIndexVillagerData          = 19
	VillagerIndexVillagerDataFinalized = 20
)

// Wither metadata indices
const (
	WitherIndexTargetA          = 16
	WitherIndexTargetB          = 17
	WitherIndexTargetC          = 18
	WitherIndexInvulnerableTime = 19
)

// Wolf metadata indices
const (
	WolfIndexIsInterested = 20
	WolfIndexCollarColor  = 21
	WolfIndexAngerEndTime = 22
	WolfIndexVariant      = 23
	WolfIndexSoundVariant = 24
)

// Zombie metadata indices
const (
	ZombieIndexIsBaby          = 16
	ZombieIndexSpecialType     = 17
	ZombieIndexBecomingDrowned = 18
)

// MetadataEntry represents a single entity metadata entry.
type MetadataEntry struct {
	Index      byte
	Serializer int32
	Data       []byte // raw wire data
}
