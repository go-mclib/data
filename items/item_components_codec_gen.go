// Code generated for Minecraft 26.3 (Protocol 777); DO NOT EDIT.

package items

import (
	"slices"

	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// auto-generated codec registrations.

func init() {
	// simple VarInt codecs
	registerCodec(ComponentAdditionalTradeCost, &varIntCodec{
		get: func(c *Components) int32 { return c.AdditionalTradeCost },
		set: func(c *Components, v int32) { c.AdditionalTradeCost = v },
	})
	registerCodec(ComponentDamage, &varIntCodec{
		get: func(c *Components) int32 { return c.Damage },
		set: func(c *Components, v int32) { c.Damage = v },
	})
	registerCodec(ComponentMaxDamage, &varIntCodec{
		get: func(c *Components) int32 { return c.MaxDamage },
		set: func(c *Components, v int32) { c.MaxDamage = v },
	})
	registerCodec(ComponentMaxStackSize, &varIntCodec{
		get: func(c *Components) int32 { return c.MaxStackSize },
		set: func(c *Components, v int32) { c.MaxStackSize = v },
	})
	registerCodec(ComponentOminousBottleAmplifier, &varIntCodec{
		get: func(c *Components) int32 { return c.OminousBottleAmplifier },
		set: func(c *Components, v int32) { c.OminousBottleAmplifier = v },
	})
	registerCodec(ComponentRepairCost, &varIntCodec{
		get: func(c *Components) int32 { return c.RepairCost },
		set: func(c *Components, v int32) { c.RepairCost = v },
	})

	// zero-valued optional VarInt codecs
	registerCodec(ComponentMapId, pointerComponentCodec[int32]{
		get:  func(c *Components) *int32 { return c.MapID },
		set:  func(c *Components, v *int32) { c.MapID = v },
		read: readVarIntPointer, write: writeVarIntPointer, equal: equalValue[int32],
	})
	registerCodec(ComponentMapPostProcessing, pointerComponentCodec[int32]{
		get:  func(c *Components) *int32 { return c.MapPostProcessing },
		set:  func(c *Components, v *int32) { c.MapPostProcessing = v },
		read: readVarIntPointer, write: writeVarIntPointer, equal: equalValue[int32],
	})

	// simple Float32 codecs
	registerCodec(ComponentMinimumAttackCharge, &float32Codec{
		get: func(c *Components) float64 { return c.MinimumAttackCharge },
		set: func(c *Components, v float64) { c.MinimumAttackCharge = v },
	})
	registerCodec(ComponentPotionDurationScale, &float32Codec{
		get: func(c *Components) float64 { return c.PotionDurationScale },
		set: func(c *Components, v float64) { c.PotionDurationScale = v },
	})

	// simple string/identifier codecs
	registerCodec(ComponentItemModel, &stringCodec{
		get: func(c *Components) string { return c.ItemModel },
		set: func(c *Components, v string) { c.ItemModel = v },
	})

	// empty marker codecs (bool flags)
	registerCodec(ComponentGlider, &emptyMarkerCodec{
		get: func(c *Components) bool { return c.Glider },
		set: func(c *Components, v bool) { c.Glider = v },
	})
	registerCodec(ComponentUnbreakable, &emptyMarkerCodec{
		get: func(c *Components) bool { return c.Unbreakable },
		set: func(c *Components, v bool) { c.Unbreakable = v },
	})
	registerCodec(ComponentWaxed, &emptyMarkerCodec{
		get: func(c *Components) bool { return c.Waxed },
		set: func(c *Components, v bool) { c.Waxed = v },
	})

	// struct codecs
	registerCodec(ComponentEnchantable, genEnchantableCodec{})
	registerCodec(ComponentFood, genFoodCodec{})
	registerCodec(ComponentTooltipDisplay, genTooltipDisplayCodec{})
	registerCodec(ComponentUseCooldown, genUseCooldownCodec{})
	registerCodec(ComponentVillagerFood, genVillagerFoodCodec{})
	registerCodec(ComponentWeapon, genWeaponCodec{})

	// VarInt passthrough
	registerPassthrough(ComponentDamageType, decodeVarIntWire)

	// Bool passthrough
	registerPassthrough(ComponentEnchantmentGlintOverride, decodeBoolWire)

	// String passthrough
	registerPassthrough(ComponentNoteBlockSound, decodeStringWire)
	registerPassthrough(ComponentTooltipStyle, decodeStringWire)

	// Empty passthrough
	registerPassthrough(ComponentCreativeSlotLock, decodeEmptyWire)

	// NBT passthrough
	registerPassthrough(ComponentBucketEntityData, decodeNBTWire)
	registerPassthrough(ComponentContainerLoot, decodeNBTWire)
	registerPassthrough(ComponentCustomData, decodeNBTWire)
	registerPassthrough(ComponentDebugStickState, decodeNBTWire)
	registerPassthrough(ComponentIntangibleProjectile, decodeNBTWire)
	registerPassthrough(ComponentLock, decodeNBTWire)
	registerPassthrough(ComponentMapDecorations, decodeNBTWire)

	// HolderSet passthrough
	registerPassthrough(ComponentDamageResistant, decodeHolderSetWire)
	registerPassthrough(ComponentProvidesBannerPatterns, decodeHolderSetWire)
	registerPassthrough(ComponentRepairable, decodeHolderSetWire)

	// Entity variant (VarInt) passthrough
	registerPassthrough(ComponentAxolotlVariant, decodeVarIntWire)
	registerPassthrough(ComponentCatCollar, decodeVarIntWire)
	registerPassthrough(ComponentCatSoundVariant, decodeVarIntWire)
	registerPassthrough(ComponentCatVariant, decodeVarIntWire)
	registerPassthrough(ComponentChickenSoundVariant, decodeVarIntWire)
	registerPassthrough(ComponentChickenVariant, decodeVarIntWire)
	registerPassthrough(ComponentCowSoundVariant, decodeVarIntWire)
	registerPassthrough(ComponentCowVariant, decodeVarIntWire)
	registerPassthrough(ComponentFoxVariant, decodeVarIntWire)
	registerPassthrough(ComponentFrogVariant, decodeVarIntWire)
	registerPassthrough(ComponentHorseVariant, decodeVarIntWire)
	registerPassthrough(ComponentLlamaVariant, decodeVarIntWire)
	registerPassthrough(ComponentMooshroomVariant, decodeVarIntWire)
	registerPassthrough(ComponentParrotVariant, decodeVarIntWire)
	registerPassthrough(ComponentPigSoundVariant, decodeVarIntWire)
	registerPassthrough(ComponentPigVariant, decodeVarIntWire)
	registerPassthrough(ComponentRabbitVariant, decodeVarIntWire)
	registerPassthrough(ComponentSalmonSize, decodeVarIntWire)
	registerPassthrough(ComponentSheepColor, decodeVarIntWire)
	registerPassthrough(ComponentShulkerColor, decodeVarIntWire)
	registerPassthrough(ComponentTropicalFishBaseColor, decodeVarIntWire)
	registerPassthrough(ComponentTropicalFishPattern, decodeVarIntWire)
	registerPassthrough(ComponentTropicalFishPatternColor, decodeVarIntWire)
	registerPassthrough(ComponentVillagerVariant, decodeVarIntWire)
	registerPassthrough(ComponentWolfCollar, decodeVarIntWire)
	registerPassthrough(ComponentWolfSoundVariant, decodeVarIntWire)
	registerPassthrough(ComponentWolfVariant, decodeVarIntWire)
	registerPassthrough(ComponentZombieNautilusVariant, decodeVarIntWire)

}

type genEnchantableCodec struct{}

func (genEnchantableCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := w.CopyVarInt(buf); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (genEnchantableCodec) Apply(c *Components, data []byte) error {
	buf := ns.NewReader(data)
	s := &Enchantable{}
	{
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		s.Value = int32(v)
	}
	c.Enchantable = s
	return nil
}

func (genEnchantableCodec) Clear(c *Components) { c.Enchantable = nil }

func (genEnchantableCodec) Differs(c, defaults *Components) (bool, bool) {
	cHas := c.Enchantable != nil
	dHas := defaults.Enchantable != nil
	if cHas != dHas {
		return true, cHas
	}
	if cHas && dHas {
		return *c.Enchantable != *defaults.Enchantable, true
	}
	return false, false
}

func (genEnchantableCodec) Encode(c *Components) ([]byte, error) {
	if c.Enchantable == nil {
		return nil, nil
	}
	w := ns.NewWriter()
	w.WriteVarInt(ns.VarInt(c.Enchantable.Value))
	return w.Bytes(), nil
}

type genFoodCodec struct{}

func (genFoodCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := w.CopyVarInt(buf); err != nil {
		return nil, err
	}
	if err := w.CopyFloat32(buf); err != nil {
		return nil, err
	}
	if err := w.CopyBool(buf); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (genFoodCodec) Apply(c *Components, data []byte) error {
	buf := ns.NewReader(data)
	s := &Food{}
	{
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		s.Nutrition = int32(v)
	}
	{
		v, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		s.Saturation = float64(v)
	}
	{
		v, err := buf.ReadBool()
		if err != nil {
			return err
		}
		s.CanAlwaysEat = bool(v)
	}
	c.Food = s
	return nil
}

func (genFoodCodec) Clear(c *Components) { c.Food = nil }

func (genFoodCodec) Differs(c, defaults *Components) (bool, bool) {
	cHas := c.Food != nil
	dHas := defaults.Food != nil
	if cHas != dHas {
		return true, cHas
	}
	if cHas && dHas {
		return *c.Food != *defaults.Food, true
	}
	return false, false
}

func (genFoodCodec) Encode(c *Components) ([]byte, error) {
	if c.Food == nil {
		return nil, nil
	}
	w := ns.NewWriter()
	w.WriteVarInt(ns.VarInt(c.Food.Nutrition))
	w.WriteFloat32(ns.Float32(c.Food.Saturation))
	w.WriteBool(ns.Boolean(c.Food.CanAlwaysEat))
	return w.Bytes(), nil
}

type genTooltipDisplayCodec struct{}

func (genTooltipDisplayCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := w.CopyBool(buf); err != nil {
		return nil, err
	}
	{
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for range int(count) {
			if err := w.CopyVarInt(buf); err != nil {
				return nil, err
			}
		}
	}
	return w.Bytes(), nil
}

func (genTooltipDisplayCodec) Apply(c *Components, data []byte) error {
	buf := ns.NewReader(data)
	s := &TooltipDisplay{}
	{
		v, err := buf.ReadBool()
		if err != nil {
			return err
		}
		s.HideTooltip = bool(v)
	}
	{
		count, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		arr := make([]int32, 0, count)
		for range int(count) {
			v, err := buf.ReadVarInt()
			if err != nil {
				return err
			}
			arr = append(arr, int32(v))
		}
		s.HiddenComponents = arr
	}
	c.TooltipDisplay = s
	return nil
}

func (genTooltipDisplayCodec) Clear(c *Components) { c.TooltipDisplay = nil }

func (genTooltipDisplayCodec) Differs(c, defaults *Components) (bool, bool) {
	cHas := c.TooltipDisplay != nil
	dHas := defaults.TooltipDisplay != nil
	if cHas != dHas {
		return true, cHas
	}
	if cHas && dHas {
		if c.TooltipDisplay.HideTooltip != defaults.TooltipDisplay.HideTooltip {
			return true, true
		}
		if !slices.Equal(c.TooltipDisplay.HiddenComponents, defaults.TooltipDisplay.HiddenComponents) {
			return true, true
		}
		return false, true
	}
	return false, false
}

func (genTooltipDisplayCodec) Encode(c *Components) ([]byte, error) {
	if c.TooltipDisplay == nil {
		return nil, nil
	}
	w := ns.NewWriter()
	w.WriteBool(ns.Boolean(c.TooltipDisplay.HideTooltip))
	w.WriteVarInt(ns.VarInt(len(c.TooltipDisplay.HiddenComponents)))
	for _, v := range c.TooltipDisplay.HiddenComponents {
		w.WriteVarInt(ns.VarInt(v))
	}
	return w.Bytes(), nil
}

type genUseCooldownCodec struct{}

func (genUseCooldownCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := w.CopyFloat32(buf); err != nil {
		return nil, err
	}
	{
		present, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(present)
		if present {
			if err := w.CopyString(buf, maxStringLen); err != nil {
				return nil, err
			}
		}
	}
	return w.Bytes(), nil
}

func (genUseCooldownCodec) Apply(c *Components, data []byte) error {
	buf := ns.NewReader(data)
	s := &UseCooldown{}
	{
		v, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		s.Seconds = float64(v)
	}
	{
		present, err := buf.ReadBool()
		if err != nil {
			return err
		}
		if present {
			v, err := buf.ReadIdentifier()
			if err != nil {
				return err
			}
			s.CooldownGroup = string(v)
		}
	}
	c.UseCooldown = s
	return nil
}

func (genUseCooldownCodec) Clear(c *Components) { c.UseCooldown = nil }

func (genUseCooldownCodec) Differs(c, defaults *Components) (bool, bool) {
	cHas := c.UseCooldown != nil
	dHas := defaults.UseCooldown != nil
	if cHas != dHas {
		return true, cHas
	}
	if cHas && dHas {
		return *c.UseCooldown != *defaults.UseCooldown, true
	}
	return false, false
}

func (genUseCooldownCodec) Encode(c *Components) ([]byte, error) {
	if c.UseCooldown == nil {
		return nil, nil
	}
	w := ns.NewWriter()
	w.WriteFloat32(ns.Float32(c.UseCooldown.Seconds))
	w.WriteBool(ns.Boolean(c.UseCooldown.CooldownGroup != ""))
	if c.UseCooldown.CooldownGroup != "" {
		w.WriteIdentifier(ns.Identifier(c.UseCooldown.CooldownGroup))
	}
	return w.Bytes(), nil
}

type genVillagerFoodCodec struct{}

func (genVillagerFoodCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := w.CopyVarInt(buf); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (genVillagerFoodCodec) Apply(c *Components, data []byte) error {
	buf := ns.NewReader(data)
	s := &VillagerFood{}
	{
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		s.Nutrition = int32(v)
	}
	c.VillagerFood = s
	return nil
}

func (genVillagerFoodCodec) Clear(c *Components) { c.VillagerFood = nil }

func (genVillagerFoodCodec) Differs(c, defaults *Components) (bool, bool) {
	cHas := c.VillagerFood != nil
	dHas := defaults.VillagerFood != nil
	if cHas != dHas {
		return true, cHas
	}
	if cHas && dHas {
		return *c.VillagerFood != *defaults.VillagerFood, true
	}
	return false, false
}

func (genVillagerFoodCodec) Encode(c *Components) ([]byte, error) {
	if c.VillagerFood == nil {
		return nil, nil
	}
	w := ns.NewWriter()
	w.WriteVarInt(ns.VarInt(c.VillagerFood.Nutrition))
	return w.Bytes(), nil
}

type genWeaponCodec struct{}

func (genWeaponCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := w.CopyVarInt(buf); err != nil {
		return nil, err
	}
	if err := w.CopyFloat32(buf); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (genWeaponCodec) Apply(c *Components, data []byte) error {
	buf := ns.NewReader(data)
	s := &Weapon{}
	{
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		s.ItemDamagePerAttack = int32(v)
	}
	{
		v, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		s.DisableBlockingForSeconds = float64(v)
	}
	c.Weapon = s
	return nil
}

func (genWeaponCodec) Clear(c *Components) { c.Weapon = nil }

func (genWeaponCodec) Differs(c, defaults *Components) (bool, bool) {
	cHas := c.Weapon != nil
	dHas := defaults.Weapon != nil
	if cHas != dHas {
		return true, cHas
	}
	if cHas && dHas {
		return *c.Weapon != *defaults.Weapon, true
	}
	return false, false
}

func (genWeaponCodec) Encode(c *Components) ([]byte, error) {
	if c.Weapon == nil {
		return nil, nil
	}
	w := ns.NewWriter()
	w.WriteVarInt(ns.VarInt(c.Weapon.ItemDamagePerAttack))
	w.WriteFloat32(ns.Float32(c.Weapon.DisableBlockingForSeconds))
	return w.Bytes(), nil
}
