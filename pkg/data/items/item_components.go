package items

// Components holds all item component data.
type Components struct {
	AttributeModifiers     []AttributeModifier
	BlocksAttacks          *BlocksAttacks
	BreakSound             string
	Consumable             *Consumable
	Container              []any
	Damage                 int32
	DamageResistant        *DamageResistant
	DamageType             string
	DeathProtection        *DeathProtection
	Enchantable            *Enchantable
	Enchantments           map[string]int32
	Equippable             *Equippable
	Fireworks              *Fireworks
	Food                   *Food
	Glider                 bool
	Instrument             string
	ItemModel              string
	ItemName               *ItemNameComponent
	JukeboxPlayable        string
	KineticWeapon          *KineticWeapon
	Lore                   []string
	MapColor               int32
	MaxDamage              int32
	MaxStackSize           int32
	MinimumAttackCharge    float64
	OminousBottleAmplifier int32
	PiercingWeapon         *PiercingWeapon
	PotionContents         *PotionContents
	PotionDurationScale    float64
	ProvidesBannerPatterns string
	ProvidesTrimMaterial   string
	Rarity                 string
	Recipes                []any
	Repairable             *Repairable
	RepairCost             int32
	StoredEnchantments     map[string]int32
	Tool                   *Tool
	TooltipDisplay         *TooltipDisplay
	UseCooldown            *UseCooldown
	UseEffects             *UseEffects
	UseRemainder           *UseRemainder
	Weapon                 *Weapon
}

type AttributeModifier struct {
	Type      string
	Amount    float64
	ID        string
	Operation string
	Slot      string
}

type BlocksAttacks struct {
	BlockDelaySeconds float64
	BlockSound        string
	BypassedBy        string
	DisabledSound     string
	ItemDamage        *DamageSpec
}

type DamageSpec struct {
	Base      float64
	Factor    float64
	Threshold float64
}

type Consumable struct {
	ConsumeSeconds   float64
	Animation        string
	OnConsumeEffects []ConsumeEffect
}

type ConsumeEffect struct {
	Type        string
	Effects     []StatusEffect
	Probability float64
}

type StatusEffect struct {
	Duration  int32
	ID        string
	ShowIcon  bool
	Amplifier int32
}

type DamageResistant struct {
	Types string
}

type DeathProtection struct {
	DeathEffects []DeathEffect
}

type DeathEffect struct {
	Type    string
	Effects []StatusEffect
}

type Enchantable struct {
	Value int32
}

type Equippable struct {
	Slot            string
	EquipSound      string
	AssetID         string
	AllowedEntities []string
	Swappable       bool
	CanBeSheared    bool
	ShearingSound   string
}

type Fireworks struct {
	FlightDuration int32
}

type Food struct {
	Nutrition  int32
	Saturation float64
}

type ItemNameComponent struct {
	Translate string
}

type KineticWeapon struct {
	DamageConditions    *KineticConditions
	DamageMultiplier    float64
	DelayTicks          int32
	DismountConditions  *KineticConditions
	ForwardMovement     float64
	HitSound            string
	KnockbackConditions *KineticConditions
	Sound               string
}

type KineticConditions struct {
	MaxDurationTicks int32
	MinRelativeSpeed float64
	MinSpeed         float64
}

type PiercingWeapon struct {
	HitSound string
	Sound    string
}

type PotionContents struct {
	// contents vary by potion type
}

type Repairable struct {
	Items string
}

type Tool struct {
	Rules                      []ToolRule
	DamagePerBlock             int32
	CanDestroyBlocksInCreative bool
}

type ToolRule struct {
	Blocks          string
	Speed           float64
	CorrectForDrops bool
}

type TooltipDisplay struct {
	// marker component
}

type UseCooldown struct {
	Seconds float64
}

type UseEffects struct {
	CanSprint          bool
	InteractVibrations bool
	SpeedMultiplier    float64
}

type UseRemainder struct {
	Count int32
	ID    string
}

type Weapon struct {
	DisableBlockingForSeconds float64
	ItemDamagePerAttack       int32
}

type AttackRange struct {
	HitboxMargin     float64
	MaxCreativeReach float64
	MaxReach         float64
	MinCreativeReach float64
	MinReach         float64
	MobFactor        float64
}
