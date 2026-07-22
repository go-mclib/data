package items

import (
	"maps"
	"slices"
)

// Components holds all item component data.
// The present bitset tracks which component IDs are explicitly set,
// mirroring Java's DataComponentPatch — only present components are
// encoded on the wire as adds or removes relative to the item's defaults.
type Components struct {
	present            [2]uint64 // bitset: which component IDs are explicitly set
	rawPassthrough     map[int32][]byte
	removedPassthrough map[int32]bool

	AdditionalTradeCost    int32
	AttributeModifiers     []AttributeModifier
	AttackRange            *AttackRange
	BannerPatterns         *[]BannerPatternLayer
	BaseColor              *int32
	BlocksAttacks          *BlocksAttacks
	BreakSound             string
	Consumable             *Consumable
	Container              []any
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
	KineticWeapon          *KineticWeapon
	Lore                   []string
	loreRaw                []byte
	MapColor               int32
	MaxDamage              int32
	MaxStackSize           int32
	MinimumAttackCharge    float64
	OminousBottleAmplifier int32
	PaintingVariant        string
	PiercingWeapon         *PiercingWeapon
	PotionContents         *PotionContents
	PotDecorations         *PotDecorations
	PotionDurationScale    float64
	ProvidesBannerPatterns string
	ProvidesTrimMaterial   string
	Rarity                 string
	Recipes                []any
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
	Slot            string
	EquipSound      string
	AssetID         string
	CameraOverlay   string
	AllowedEntities []string
	Dispensable     bool
	Swappable       bool
	DamageOnHurt    bool
	CanBeSheared    bool
	ShearingSound   string
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

type PotDecorations struct {
	Back  string
	Left  string
	Right string
	Front string
}

type PotionContents struct {
	// contents vary by potion type
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

// HasComponent returns true if the given component ID is marked as present.
func (c *Components) HasComponent(id int32) bool {
	if id < 0 || id >= 128 {
		return false
	}
	return c.present[id/64]&(1<<uint(id%64)) != 0
}

// SetPresent marks the given component ID as present.
func (c *Components) SetPresent(id int32) {
	if id >= 0 && id < 128 {
		c.present[id/64] |= 1 << uint(id%64)
	}
}

func cloneFireworkExplosion(explosion FireworkExplosion) FireworkExplosion {
	explosion.Colors = append([]int32(nil), explosion.Colors...)
	explosion.FadeColors = append([]int32(nil), explosion.FadeColors...)
	return explosion
}

func cloneStatusEffects(source []StatusEffect) []StatusEffect {
	result := make([]StatusEffect, len(source))
	for i, effect := range source {
		result[i] = effect
		if effect.Hidden != nil {
			hidden := cloneStatusEffects([]StatusEffect{*effect.Hidden})
			result[i].Hidden = &hidden[0]
		}
	}
	return result
}

func cloneConsumeEffects(source []ConsumeEffect) []ConsumeEffect {
	result := make([]ConsumeEffect, len(source))
	for i, effect := range source {
		result[i] = effect
		result[i].Effects = cloneStatusEffects(effect.Effects)
		result[i].EffectIDs = append([]string(nil), effect.EffectIDs...)
		if effect.SoundRange != nil {
			value := *effect.SoundRange
			result[i].SoundRange = &value
		}
	}
	return result
}

// Clone returns a deep copy of the Components struct.
func (c *Components) Clone() *Components {
	if c == nil {
		return &Components{}
	}

	clone := *c

	// clone slices
	if c.AttributeModifiers != nil {
		clone.AttributeModifiers = make([]AttributeModifier, len(c.AttributeModifiers))
		for i, modifier := range c.AttributeModifiers {
			clone.AttributeModifiers[i] = modifier
			if modifier.Display.Value != nil {
				value := *modifier.Display.Value
				value.Raw = append([]byte(nil), modifier.Display.Value.Raw...)
				clone.AttributeModifiers[i].Display.Value = &value
			}
		}
	}
	if c.Container != nil {
		clone.Container = make([]any, len(c.Container))
		copy(clone.Container, c.Container)
	}
	if c.Lore != nil {
		clone.Lore = make([]string, len(c.Lore))
		copy(clone.Lore, c.Lore)
	}
	clone.loreRaw = append([]byte(nil), c.loreRaw...)
	if c.Recipes != nil {
		clone.Recipes = make([]any, len(c.Recipes))
		copy(clone.Recipes, c.Recipes)
	}

	// clone maps
	if c.Enchantments != nil {
		clone.Enchantments = make(map[string]int32, len(c.Enchantments))
		maps.Copy(clone.Enchantments, c.Enchantments)
	}
	if c.StoredEnchantments != nil {
		clone.StoredEnchantments = make(map[string]int32, len(c.StoredEnchantments))
		maps.Copy(clone.StoredEnchantments, c.StoredEnchantments)
	}
	clone.removedPassthrough = maps.Clone(c.removedPassthrough)
	if c.rawPassthrough != nil {
		clone.rawPassthrough = make(map[int32][]byte, len(c.rawPassthrough))
		for id, data := range c.rawPassthrough {
			clone.rawPassthrough[id] = append([]byte(nil), data...)
		}
	}

	// clone pointer types
	if c.AttackRange != nil {
		v := *c.AttackRange
		clone.AttackRange = &v
	}
	if c.BannerPatterns != nil {
		layers := append([]BannerPatternLayer(nil), (*c.BannerPatterns)...)
		clone.BannerPatterns = &layers
	}
	if c.BaseColor != nil {
		v := *c.BaseColor
		clone.BaseColor = &v
	}
	if c.BlocksAttacks != nil {
		v := *c.BlocksAttacks
		v.DamageReductions = slices.Clone(c.BlocksAttacks.DamageReductions)
		for i := range v.DamageReductions {
			v.DamageReductions[i].Types = slices.Clone(v.DamageReductions[i].Types)
		}
		v.BypassedBy = slices.Clone(c.BlocksAttacks.BypassedBy)
		if c.BlocksAttacks.BlockSoundRange != nil {
			rangeValue := *c.BlocksAttacks.BlockSoundRange
			v.BlockSoundRange = &rangeValue
		}
		if c.BlocksAttacks.DisabledSoundRange != nil {
			rangeValue := *c.BlocksAttacks.DisabledSoundRange
			v.DisabledSoundRange = &rangeValue
		}
		clone.BlocksAttacks = &v
	}
	if c.Consumable != nil {
		v := *c.Consumable
		if c.Consumable.SoundRange != nil {
			rangeValue := *c.Consumable.SoundRange
			v.SoundRange = &rangeValue
		}
		v.OnConsumeEffects = cloneConsumeEffects(c.Consumable.OnConsumeEffects)
		clone.Consumable = &v
	}
	if c.DamageResistant != nil {
		v := *c.DamageResistant
		clone.DamageResistant = &v
	}
	if c.CustomName != nil {
		v := *c.CustomName
		v.Raw = append([]byte(nil), c.CustomName.Raw...)
		clone.CustomName = &v
	}
	if c.DeathProtection != nil {
		v := *c.DeathProtection
		v.DeathEffects = cloneConsumeEffects(c.DeathProtection.DeathEffects)
		clone.DeathProtection = &v
	}
	if c.Dye != nil {
		v := *c.Dye
		clone.Dye = &v
	}
	if c.DyedColor != nil {
		v := *c.DyedColor
		clone.DyedColor = &v
	}
	if c.Enchantable != nil {
		v := *c.Enchantable
		clone.Enchantable = &v
	}
	if c.EntityData != nil {
		v := *c.EntityData
		v.Raw = append([]byte(nil), c.EntityData.Raw...)
		clone.EntityData = &v
	}
	if c.Equippable != nil {
		v := *c.Equippable
		v.AllowedEntities = append([]string(nil), c.Equippable.AllowedEntities...)
		clone.Equippable = &v
	}
	if c.FireworkExplosion != nil {
		v := cloneFireworkExplosion(*c.FireworkExplosion)
		clone.FireworkExplosion = &v
	}
	if c.Fireworks != nil {
		v := *c.Fireworks
		clone.Fireworks = &v
	}
	if c.FireworkExplosions != nil {
		clone.FireworkExplosions = make([]FireworkExplosion, len(c.FireworkExplosions))
		for i, explosion := range c.FireworkExplosions {
			clone.FireworkExplosions[i] = cloneFireworkExplosion(explosion)
		}
	}
	if c.Food != nil {
		v := *c.Food
		clone.Food = &v
	}
	if c.ItemName != nil {
		v := *c.ItemName
		v.Raw = append([]byte(nil), c.ItemName.Raw...)
		clone.ItemName = &v
	}
	if c.KineticWeapon != nil {
		v := *c.KineticWeapon
		if v.DamageConditions != nil {
			conditions := *v.DamageConditions
			v.DamageConditions = &conditions
		}
		if v.DismountConditions != nil {
			conditions := *v.DismountConditions
			v.DismountConditions = &conditions
		}
		if v.KnockbackConditions != nil {
			conditions := *v.KnockbackConditions
			v.KnockbackConditions = &conditions
		}
		clone.KineticWeapon = &v
	}
	if c.PiercingWeapon != nil {
		v := *c.PiercingWeapon
		clone.PiercingWeapon = &v
	}
	if c.PotionContents != nil {
		v := *c.PotionContents
		clone.PotionContents = &v
	}
	if c.PotDecorations != nil {
		v := *c.PotDecorations
		clone.PotDecorations = &v
	}
	if c.Repairable != nil {
		v := *c.Repairable
		clone.Repairable = &v
	}
	if c.Tool != nil {
		v := *c.Tool
		v.Rules = append([]ToolRule(nil), c.Tool.Rules...)
		for i := range v.Rules {
			v.Rules[i].Blocks = append([]string(nil), v.Rules[i].Blocks...)
		}
		clone.Tool = &v
	}
	if c.Trim != nil {
		v := *c.Trim
		clone.Trim = &v
	}
	if c.TooltipDisplay != nil {
		v := *c.TooltipDisplay
		v.HiddenComponents = append([]int32(nil), c.TooltipDisplay.HiddenComponents...)
		clone.TooltipDisplay = &v
	}
	if c.UseCooldown != nil {
		v := *c.UseCooldown
		clone.UseCooldown = &v
	}
	if c.UseEffects != nil {
		v := *c.UseEffects
		clone.UseEffects = &v
	}
	if c.UseRemainder != nil {
		v := *c.UseRemainder
		clone.UseRemainder = &v
	}
	if c.Weapon != nil {
		v := *c.Weapon
		clone.Weapon = &v
	}

	return &clone
}
