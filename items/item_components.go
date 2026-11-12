package items

import (
	"maps"
	"slices"
)

// componentWords is the number of uint64 words in the presence bitset.
const componentWords = MaxComponentID/64 + 1

// Components holds all item component data. Optional components are pointers
// (nil = absent). present tracks which component IDs ToSlot considers when
// diffing against the item defaults.
type Components struct {
	present            [componentWords]uint64
	rawPassthrough     map[int32][]byte
	removedPassthrough map[int32]bool

	AdditionalTradeCost    int32
	AttributeModifiers     []AttributeModifier
	AttackRange            *AttackRange
	AttackAnimation        *SwingAnimation
	InteractAnimation      *SwingAnimation
	BlockTransformer       string
	VillagerFood           *VillagerFood
	Compostable            *Compostable
	CookingFuel            *CookingFuel
	BrewingFuel            *BrewingFuel
	MobVisibility          *MobVisibility
	ProvidesPotteryPattern string
	SignTextFront          *SignText
	SignTextBack           *SignText
	Waxed                  bool
	CushionColor           *int32
	BannerPatterns         *[]BannerPatternLayer
	BaseColor              *int32
	BlocksAttacks          *BlocksAttacks
	BreakSound             string
	Consumable             *Consumable
	CustomName             *ItemNameComponent
	Damage                 int32
	DamageResistant        *DamageResistant
	DamageType             string
	DeathProtection        *DeathProtection
	Dye                    *int32
	DyedColor              *int32
	Enchantable            *Enchantable
	Enchantments           map[string]int32
	EntityData             *TypedEntityData
	Equippable             *Equippable
	FireworkExplosion      *FireworkExplosion
	Fireworks              *Fireworks
	FireworkExplosions     []FireworkExplosion
	Food                   *Food
	Glider                 bool
	Instrument             string
	ItemModel              string
	ItemName               *ItemNameComponent
	JukeboxPlayable        string
	Lore                   []string
	loreRaw                []byte
	MapID                  *int32
	MapPostProcessing      *int32
	MaxDamage              int32
	MaxStackSize           int32
	MinimumAttackCharge    float64
	OminousBottleAmplifier int32
	PaintingVariant        string
	PotDecorations         *PotDecorations
	PotionDurationScale    float64
	ProvidesBannerPatterns string
	ProvidesTrimMaterial   string
	Rarity                 string
	Repairable             *Repairable
	RepairCost             int32
	StoredEnchantments     map[string]int32
	Tool                   *Tool
	Trim                   *ArmorTrim
	TooltipDisplay         *TooltipDisplay
	Unbreakable            bool
	UseCooldown            *UseCooldown
	UseEffects             *UseEffects
	UseRemainder           *UseRemainder
	Weapon                 *Weapon
}

type BannerPatternLayer struct {
	Pattern        string
	AssetID        string
	TranslationKey string
	Color          int32
}

type ArmorTrim struct {
	Material string
	Pattern  string
}

type AttributeModifier struct {
	Type      string
	Amount    float64
	ID        string
	Operation string
	Slot      string
	Display   AttributeModifierDisplay
}

type AttributeModifierDisplay struct {
	Type  string
	Value *ItemNameComponent
}

type BlocksAttacks struct {
	BlockDelaySeconds    float64
	DisableCooldownScale float64
	DamageReductions     []DamageReduction
	ItemDamage           DamageSpec
	BypassedBy           []string
	BlockSound           string
	BlockSoundDirect     bool
	BlockSoundRange      *float64
	DisabledSound        string
	DisabledSoundDirect  bool
	DisabledSoundRange   *float64
}

type DamageReduction struct {
	HorizontalBlockingAngle float64
	Types                   []string
	Base                    float64
	Factor                  float64
}

type DamageSpec struct {
	Threshold float64
	Base      float64
	Factor    float64
}

type Consumable struct {
	ConsumeSeconds   float64
	Animation        string
	Sound            string
	SoundDirect      bool
	SoundRange       *float64
	HasParticles     bool
	OnConsumeEffects []ConsumeEffect
}

type ConsumeEffect struct {
	Type        string
	Effects     []StatusEffect
	EffectIDs   []string
	Sound       string
	SoundDirect bool
	SoundRange  *float64
	Probability float64
	Diameter    float64
}

type StatusEffect struct {
	Duration      int32
	ID            string
	ShowIcon      bool
	ShowParticles bool
	Ambient       bool
	Amplifier     int32
	Hidden        *StatusEffect
}

type DamageResistant struct {
	Types string
}

type DeathProtection struct {
	DeathEffects []ConsumeEffect
}

type Enchantable struct {
	Value int32
}

type TypedEntityData struct {
	Type string
	Raw  []byte
}

type Equippable struct {
	Slot                string
	EquipSound          string
	EquipSoundDirect    bool
	EquipSoundRange     *float64
	AssetID             string
	CameraOverlay       string
	AllowedEntities     []string
	Dispensable         bool
	Swappable           bool
	DamageOnHurt        bool
	EquipOnInteract     bool
	CanBeSheared        bool
	ShearingSound       string
	ShearingSoundDirect bool
	ShearingSoundRange  *float64
}

type FireworkExplosion struct {
	Shape      int32
	Colors     []int32
	FadeColors []int32
	HasTrail   bool
	HasTwinkle bool
}

type Fireworks struct {
	FlightDuration int32
}

type Food struct {
	Nutrition    int32
	Saturation   float64
	CanAlwaysEat bool
}

type ItemNameComponent struct {
	Text      string
	Translate string
	Raw       []byte
}

type PotDecorations struct {
	Back  *Stack
	Left  *Stack
	Right *Stack
	Front *Stack
}

type Repairable struct {
	Items string
}

type Tool struct {
	Rules                      []ToolRule
	DefaultMiningSpeed         float64
	DamagePerBlock             int32
	CanDestroyBlocksInCreative bool
}

type ToolRule struct {
	Blocks             []string
	Speed              float64
	HasSpeed           bool
	CorrectForDrops    bool
	HasCorrectForDrops bool
}

type TooltipDisplay struct {
	HideTooltip      bool
	HiddenComponents []int32
}

type UseCooldown struct {
	Seconds       float64
	CooldownGroup string
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

// hasComponent reports whether a component ID is tracked for patch encoding.
func (c *Components) hasComponent(id int32) bool {
	if id < 0 || id > MaxComponentID {
		return false
	}
	return c.present[id/64]&(1<<(id%64)) != 0
}

// SetPresent marks a component ID for patch encoding. Stacks from NewStack and
// FromSlot already track every component, so SetPresent is only needed for
// Components built by hand (e.g. &Components{} literals).
func (c *Components) SetPresent(id int32) {
	if id >= 0 && id <= MaxComponentID {
		c.present[id/64] |= 1 << (id % 64)
	}
}

// ClearPresent excludes a component ID from patch encoding: ToSlot neither
// adds nor removes it.
func (c *Components) ClearPresent(id int32) {
	if id >= 0 && id <= MaxComponentID {
		c.present[id/64] &^= 1 << (id % 64)
	}
}

func cloneFireworkExplosion(explosion FireworkExplosion) FireworkExplosion {
	explosion.Colors = slices.Clone(explosion.Colors)
	explosion.FadeColors = slices.Clone(explosion.FadeColors)
	return explosion
}

func cloneStatusEffect(effect StatusEffect) StatusEffect {
	if effect.Hidden != nil {
		hidden := cloneStatusEffect(*effect.Hidden)
		effect.Hidden = &hidden
	}
	return effect
}

func cloneConsumeEffects(source []ConsumeEffect) []ConsumeEffect {
	result := slices.Clone(source)
	for i, effect := range source {
		result[i].Effects = slices.Clone(effect.Effects)
		for j, status := range effect.Effects {
			result[i].Effects[j] = cloneStatusEffect(status)
		}
		result[i].EffectIDs = slices.Clone(effect.EffectIDs)
		result[i].SoundRange = cloneValue(effect.SoundRange)
	}
	return result
}

func cloneItemName(value *ItemNameComponent) *ItemNameComponent {
	clone := cloneValue(value)
	if clone != nil {
		clone.Raw = slices.Clone(value.Raw)
	}
	return clone
}

// Clone returns a deep copy of the Components struct.
func (c *Components) Clone() *Components {
	if c == nil {
		return &Components{}
	}

	clone := *c
	clone.AttackAnimation = cloneValue(c.AttackAnimation)
	clone.InteractAnimation = cloneValue(c.InteractAnimation)
	clone.VillagerFood = cloneValue(c.VillagerFood)
	clone.Compostable = cloneValue(c.Compostable)
	clone.CookingFuel = cloneValue(c.CookingFuel)
	clone.BrewingFuel = cloneValue(c.BrewingFuel)
	clone.CushionColor = cloneValue(c.CushionColor)
	if clone.MobVisibility = cloneValue(c.MobVisibility); clone.MobVisibility != nil {
		clone.MobVisibility.TargetingEntityTypes = slices.Clone(c.MobVisibility.TargetingEntityTypes)
	}
	clone.SignTextFront = cloneSignText(c.SignTextFront)
	clone.SignTextBack = cloneSignText(c.SignTextBack)

	clone.AttributeModifiers = slices.Clone(c.AttributeModifiers)
	for i := range clone.AttributeModifiers {
		clone.AttributeModifiers[i].Display.Value = cloneItemName(c.AttributeModifiers[i].Display.Value)
	}
	clone.Lore = slices.Clone(c.Lore)
	clone.loreRaw = slices.Clone(c.loreRaw)
	clone.Enchantments = maps.Clone(c.Enchantments)
	clone.StoredEnchantments = maps.Clone(c.StoredEnchantments)
	clone.removedPassthrough = maps.Clone(c.removedPassthrough)
	if c.rawPassthrough != nil {
		clone.rawPassthrough = make(map[int32][]byte, len(c.rawPassthrough))
		for id, data := range c.rawPassthrough {
			clone.rawPassthrough[id] = slices.Clone(data)
		}
	}

	clone.AttackRange = cloneValue(c.AttackRange)
	if c.BannerPatterns != nil {
		layers := slices.Clone(*c.BannerPatterns)
		clone.BannerPatterns = &layers
	}
	clone.BaseColor = cloneValue(c.BaseColor)
	if clone.BlocksAttacks = cloneValue(c.BlocksAttacks); clone.BlocksAttacks != nil {
		v := clone.BlocksAttacks
		v.DamageReductions = slices.Clone(v.DamageReductions)
		for i := range v.DamageReductions {
			v.DamageReductions[i].Types = slices.Clone(v.DamageReductions[i].Types)
		}
		v.BypassedBy = slices.Clone(v.BypassedBy)
		v.BlockSoundRange = cloneValue(v.BlockSoundRange)
		v.DisabledSoundRange = cloneValue(v.DisabledSoundRange)
	}
	if clone.Consumable = cloneValue(c.Consumable); clone.Consumable != nil {
		clone.Consumable.SoundRange = cloneValue(c.Consumable.SoundRange)
		clone.Consumable.OnConsumeEffects = cloneConsumeEffects(c.Consumable.OnConsumeEffects)
	}
	clone.DamageResistant = cloneValue(c.DamageResistant)
	clone.CustomName = cloneItemName(c.CustomName)
	if clone.DeathProtection = cloneValue(c.DeathProtection); clone.DeathProtection != nil {
		clone.DeathProtection.DeathEffects = cloneConsumeEffects(c.DeathProtection.DeathEffects)
	}
	clone.Dye = cloneValue(c.Dye)
	clone.DyedColor = cloneValue(c.DyedColor)
	clone.Enchantable = cloneValue(c.Enchantable)
	if clone.EntityData = cloneValue(c.EntityData); clone.EntityData != nil {
		clone.EntityData.Raw = slices.Clone(c.EntityData.Raw)
	}
	if clone.Equippable = cloneValue(c.Equippable); clone.Equippable != nil {
		clone.Equippable.AllowedEntities = slices.Clone(c.Equippable.AllowedEntities)
		clone.Equippable.EquipSoundRange = cloneValue(c.Equippable.EquipSoundRange)
		clone.Equippable.ShearingSoundRange = cloneValue(c.Equippable.ShearingSoundRange)
	}
	if c.FireworkExplosion != nil {
		v := cloneFireworkExplosion(*c.FireworkExplosion)
		clone.FireworkExplosion = &v
	}
	clone.Fireworks = cloneValue(c.Fireworks)
	if c.FireworkExplosions != nil {
		clone.FireworkExplosions = make([]FireworkExplosion, len(c.FireworkExplosions))
		for i, explosion := range c.FireworkExplosions {
			clone.FireworkExplosions[i] = cloneFireworkExplosion(explosion)
		}
	}
	clone.Food = cloneValue(c.Food)
	clone.ItemName = cloneItemName(c.ItemName)
	clone.MapID = cloneValue(c.MapID)
	clone.MapPostProcessing = cloneValue(c.MapPostProcessing)
	if clone.PotDecorations = cloneValue(c.PotDecorations); clone.PotDecorations != nil {
		v := clone.PotDecorations
		v.Back = cloneDecoration(v.Back)
		v.Left = cloneDecoration(v.Left)
		v.Right = cloneDecoration(v.Right)
		v.Front = cloneDecoration(v.Front)
	}
	clone.Repairable = cloneValue(c.Repairable)
	if clone.Tool = cloneValue(c.Tool); clone.Tool != nil {
		clone.Tool.Rules = slices.Clone(c.Tool.Rules)
		for i := range clone.Tool.Rules {
			clone.Tool.Rules[i].Blocks = slices.Clone(clone.Tool.Rules[i].Blocks)
		}
	}
	clone.Trim = cloneValue(c.Trim)
	if clone.TooltipDisplay = cloneValue(c.TooltipDisplay); clone.TooltipDisplay != nil {
		clone.TooltipDisplay.HiddenComponents = slices.Clone(c.TooltipDisplay.HiddenComponents)
	}
	clone.UseCooldown = cloneValue(c.UseCooldown)
	clone.UseEffects = cloneValue(c.UseEffects)
	clone.UseRemainder = cloneValue(c.UseRemainder)
	clone.Weapon = cloneValue(c.Weapon)

	return &clone
}

func int32Pointer(value int32) *int32 {
	return &value
}
