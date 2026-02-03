package items

// Component codec registration - registers all component codecs at init time.

import (
	"fmt"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

func init() {
	// Simple VarInt components
	RegisterCodec(ComponentMaxStackSize, &varIntCodec{
		get: func(c *Components) int32 { return c.MaxStackSize },
		set: func(c *Components, v int32) { c.MaxStackSize = v },
	})
	RegisterCodec(ComponentMaxDamage, &varIntCodec{
		get: func(c *Components) int32 { return c.MaxDamage },
		set: func(c *Components, v int32) { c.MaxDamage = v },
	})
	RegisterCodec(ComponentDamage, &varIntCodec{
		get: func(c *Components) int32 { return c.Damage },
		set: func(c *Components, v int32) { c.Damage = v },
	})
	RegisterCodec(ComponentRepairCost, &varIntCodec{
		get: func(c *Components) int32 { return c.RepairCost },
		set: func(c *Components, v int32) { c.RepairCost = v },
	})
	RegisterCodec(ComponentMapColor, &varIntCodec{
		get: func(c *Components) int32 { return c.MapColor },
		set: func(c *Components, v int32) { c.MapColor = v },
	})
	RegisterCodec(ComponentOminousBottleAmplifier, &varIntCodec{
		get: func(c *Components) int32 { return c.OminousBottleAmplifier },
		set: func(c *Components, v int32) { c.OminousBottleAmplifier = v },
	})

	// Simple Float32 components
	RegisterCodec(ComponentMinimumAttackCharge, &float32Codec{
		get: func(c *Components) float64 { return c.MinimumAttackCharge },
		set: func(c *Components, v float64) { c.MinimumAttackCharge = v },
	})
	RegisterCodec(ComponentPotionDurationScale, &float32Codec{
		get: func(c *Components) float64 { return c.PotionDurationScale },
		set: func(c *Components, v float64) { c.PotionDurationScale = v },
	})

	// Simple string/identifier components
	RegisterCodec(ComponentDamageType, &stringCodec{
		get: func(c *Components) string { return c.DamageType },
		set: func(c *Components, v string) { c.DamageType = v },
	})
	RegisterCodec(ComponentItemModel, &stringCodec{
		get: func(c *Components) string { return c.ItemModel },
		set: func(c *Components, v string) { c.ItemModel = v },
	})
	RegisterCodec(ComponentInstrument, &stringCodec{
		get: func(c *Components) string { return c.Instrument },
		set: func(c *Components, v string) { c.Instrument = v },
	})
	RegisterCodec(ComponentProvidesTrimMaterial, &stringCodec{
		get: func(c *Components) string { return c.ProvidesTrimMaterial },
		set: func(c *Components, v string) { c.ProvidesTrimMaterial = v },
	})
	RegisterCodec(ComponentJukeboxPlayable, &stringCodec{
		get: func(c *Components) string { return c.JukeboxPlayable },
		set: func(c *Components, v string) { c.JukeboxPlayable = v },
	})
	RegisterCodec(ComponentProvidesBannerPatterns, &stringCodec{
		get: func(c *Components) string { return c.ProvidesBannerPatterns },
		set: func(c *Components, v string) { c.ProvidesBannerPatterns = v },
	})
	RegisterCodec(ComponentBreakSound, &stringCodec{
		get: func(c *Components) string { return c.BreakSound },
		set: func(c *Components, v string) { c.BreakSound = v },
	})

	// Empty marker components (bool flags)
	RegisterCodec(ComponentUnbreakable, &emptyMarkerCodec{
		get: func(c *Components) bool { return c.Unbreakable },
		set: func(c *Components, v bool) { c.Unbreakable = v },
	})
	RegisterCodec(ComponentGlider, &emptyMarkerCodec{
		get: func(c *Components) bool { return c.Glider },
		set: func(c *Components, v bool) { c.Glider = v },
	})

	// Struct components with dedicated codecs
	RegisterCodec(ComponentCustomName, &customNameCodec{})
	RegisterCodec(ComponentItemName, &itemNameCodec{})
	RegisterCodec(ComponentTooltipDisplay, &tooltipDisplayCodec{})
	RegisterCodec(ComponentAttributeModifiers, &attributeModifiersCodec{})
	RegisterCodec(ComponentFood, &foodCodec{})
	RegisterCodec(ComponentWeapon, &weaponCodec{})
	RegisterCodec(ComponentEnchantable, &enchantableCodec{})
	RegisterCodec(ComponentUseCooldown, &useCooldownCodec{})
	RegisterCodec(ComponentFireworks, &fireworksCodec{})
	RegisterCodec(ComponentRarity, &rarityCodec{})

	// Passthrough codecs for components we don't fully decode yet
	// These just copy the wire format bytes without interpreting them

	// VarInt passthrough
	registerVarIntPassthrough(ComponentMapId)

	// Bool passthrough
	registerBoolPassthrough(ComponentEnchantmentGlintOverride)

	// String/identifier passthrough
	registerStringPassthrough(ComponentTooltipStyle)
	registerStringPassthrough(ComponentNoteBlockSound)

	// Empty marker passthrough
	registerEmptyPassthrough(ComponentCreativeSlotLock)
	registerEmptyPassthrough(ComponentIntangibleProjectile)

	// Lore (VarInt count + NBT[])
	RegisterCodec(ComponentLore, &passthroughCodec{decode: decodeLoreWire})

	// Enchantments (VarInt count + entries + bool)
	RegisterCodec(ComponentEnchantments, &passthroughCodec{decode: decodeEnchantmentsWire})
	RegisterCodec(ComponentStoredEnchantments, &passthroughCodec{decode: decodeEnchantmentsWire})

	// Complex components - passthrough for now
	RegisterCodec(ComponentCustomData, &passthroughCodec{decode: decodeNBTWire})
	RegisterCodec(ComponentCanBreak, &passthroughCodec{decode: decodeBlockPredicatesWire})
	RegisterCodec(ComponentCanPlaceOn, &passthroughCodec{decode: decodeBlockPredicatesWire})
	RegisterCodec(ComponentCustomModelData, &passthroughCodec{decode: decodeCustomModelDataWire})
	RegisterCodec(ComponentConsumable, &passthroughCodec{decode: decodeConsumableWire})
	RegisterCodec(ComponentUseRemainder, &passthroughCodec{decode: decodeSlotWire})
	RegisterCodec(ComponentUseEffects, &passthroughCodec{decode: decodeUseEffectsWire})
	RegisterCodec(ComponentDamageResistant, &passthroughCodec{decode: decodeHolderSetWire})
	RegisterCodec(ComponentTool, &passthroughCodec{decode: decodeToolWire})
	RegisterCodec(ComponentAttackRange, &passthroughCodec{decode: decodeAttackRangeWire})
	RegisterCodec(ComponentEquippable, &passthroughCodec{decode: decodeEquippableWire})
	RegisterCodec(ComponentRepairable, &passthroughCodec{decode: decodeHolderSetWire})
	RegisterCodec(ComponentDeathProtection, &passthroughCodec{decode: decodeDeathProtectionWire})
	RegisterCodec(ComponentBlocksAttacks, &passthroughCodec{decode: decodeBlocksAttacksWire})
	RegisterCodec(ComponentKineticWeapon, &passthroughCodec{decode: decodeKineticWeaponWire})
	RegisterCodec(ComponentPiercingWeapon, &passthroughCodec{decode: decodePiercingWeaponWire})
	RegisterCodec(ComponentSwingAnimation, &passthroughCodec{decode: decodeVarIntWire})
	RegisterCodec(ComponentDyedColor, &passthroughCodec{decode: decodeInt32Wire})
	RegisterCodec(ComponentMapDecorations, &passthroughCodec{decode: decodeNBTWire})
	RegisterCodec(ComponentMapPostProcessing, &passthroughCodec{decode: decodeVarIntWire})
	RegisterCodec(ComponentChargedProjectiles, &passthroughCodec{decode: decodeSlotListWire})
	RegisterCodec(ComponentBundleContents, &passthroughCodec{decode: decodeSlotListWire})
	RegisterCodec(ComponentPotionContents, &passthroughCodec{decode: decodePotionContentsWire})
	RegisterCodec(ComponentSuspiciousStewEffects, &passthroughCodec{decode: decodeSuspiciousStewWire})
	RegisterCodec(ComponentWritableBookContent, &passthroughCodec{decode: decodeWritableBookWire})
	RegisterCodec(ComponentWrittenBookContent, &passthroughCodec{decode: decodeWrittenBookWire})
	RegisterCodec(ComponentTrim, &passthroughCodec{decode: decodeTrimWire})
	RegisterCodec(ComponentDebugStickState, &passthroughCodec{decode: decodeNBTWire})
	RegisterCodec(ComponentEntityData, &passthroughCodec{decode: decodeNBTWire})
	RegisterCodec(ComponentBucketEntityData, &passthroughCodec{decode: decodeNBTWire})
	RegisterCodec(ComponentBlockEntityData, &passthroughCodec{decode: decodeNBTWire})
	RegisterCodec(ComponentRecipes, &passthroughCodec{decode: decodeRecipesWire})
	RegisterCodec(ComponentLodestoneTracker, &passthroughCodec{decode: decodeLodestoneWire})
	RegisterCodec(ComponentFireworkExplosion, &passthroughCodec{decode: decodeFireworkExplosionWire})
	RegisterCodec(ComponentProfile, &passthroughCodec{decode: decodeProfileWire})
	RegisterCodec(ComponentBannerPatterns, &passthroughCodec{decode: decodeBannerPatternsWire})
	RegisterCodec(ComponentBaseColor, &passthroughCodec{decode: decodeVarIntWire})
	RegisterCodec(ComponentPotDecorations, &passthroughCodec{decode: decodePotDecorationsWire})
	RegisterCodec(ComponentContainer, &passthroughCodec{decode: decodeSlotListWire})
	RegisterCodec(ComponentBlockState, &passthroughCodec{decode: decodeBlockStateWire})
	RegisterCodec(ComponentBees, &passthroughCodec{decode: decodeBeesWire})
	RegisterCodec(ComponentLock, &passthroughCodec{decode: decodeNBTWire})
	RegisterCodec(ComponentContainerLoot, &passthroughCodec{decode: decodeNBTWire})

	// Entity variant components (VarInt)
	registerVarIntPassthrough(ComponentVillagerVariant)
	registerVarIntPassthrough(ComponentWolfVariant)
	registerVarIntPassthrough(ComponentWolfSoundVariant)
	registerVarIntPassthrough(ComponentWolfCollar)
	registerVarIntPassthrough(ComponentFoxVariant)
	registerVarIntPassthrough(ComponentSalmonSize)
	registerVarIntPassthrough(ComponentParrotVariant)
	registerVarIntPassthrough(ComponentTropicalFishPattern)
	registerVarIntPassthrough(ComponentTropicalFishBaseColor)
	registerVarIntPassthrough(ComponentTropicalFishPatternColor)
	registerVarIntPassthrough(ComponentMooshroomVariant)
	registerVarIntPassthrough(ComponentRabbitVariant)
	registerVarIntPassthrough(ComponentPigVariant)
	registerVarIntPassthrough(ComponentCowVariant)
	registerVarIntPassthrough(ComponentChickenVariant)
	registerVarIntPassthrough(ComponentZombieNautilusVariant)
	registerVarIntPassthrough(ComponentFrogVariant)
	registerVarIntPassthrough(ComponentHorseVariant)
	registerVarIntPassthrough(ComponentPaintingVariant)
	registerVarIntPassthrough(ComponentLlamaVariant)
	registerVarIntPassthrough(ComponentAxolotlVariant)
	registerVarIntPassthrough(ComponentCatVariant)
	registerVarIntPassthrough(ComponentCatCollar)
	registerVarIntPassthrough(ComponentSheepColor)
	registerVarIntPassthrough(ComponentShulkerColor)
}

// Helper functions for registering simple passthrough codecs

func registerVarIntPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeVarIntWire})
}

func registerBoolPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeBoolWire})
}

func registerStringPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeStringWire})
}

func registerEmptyPassthrough(id int32) {
	RegisterCodec(id, &passthroughCodec{decode: decodeEmptyWire})
}

// Wire format decode functions for passthrough codecs

func decodeVarIntWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	v, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(v)
	return nil
}

func decodeInt32Wire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	v, err := buf.ReadInt32()
	if err != nil {
		return err
	}
	w.WriteInt32(v)
	return nil
}

func decodeBoolWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	v, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(v)
	return nil
}

func decodeStringWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	v, err := buf.ReadString(maxStringLen)
	if err != nil {
		return err
	}
	w.WriteString(v)
	return nil
}

func decodeEmptyWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return nil
}

func decodeNBTWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyNBT(buf, w)
}

func decodeLoreWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		if err := copyNBT(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodeEnchantmentsWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		id, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		level, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(id)
		w.WriteVarInt(level)
	}
	show, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(show)
	return nil
}

func decodeBlockPredicatesWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		if err := copyBlockPredicate(buf, w); err != nil {
			return err
		}
	}
	show, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(show)
	return nil
}

func decodeCustomModelDataWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// floats
	floatCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(floatCount)
	for range int(floatCount) {
		v, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(v)
	}
	// flags
	flagCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(flagCount)
	for range int(flagCount) {
		v, err := buf.ReadBool()
		if err != nil {
			return err
		}
		w.WriteBool(v)
	}
	// strings
	stringCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(stringCount)
	for range int(stringCount) {
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(v)
	}
	// colors
	colorCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(colorCount)
	for range int(colorCount) {
		v, err := buf.ReadInt32()
		if err != nil {
			return err
		}
		w.WriteInt32(v)
	}
	return nil
}

func decodeConsumableWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// consume seconds
	seconds, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(seconds)
	// animation
	animation, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(animation)
	// sound
	if err := copySoundEvent(buf, w); err != nil {
		return err
	}
	// has particles
	hasParticles, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasParticles)
	// on consume effects
	effectCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(effectCount)
	for range int(effectCount) {
		if err := copyConsumeEffect(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodeSlotWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copySlot(buf, w)
}

func decodeSlotListWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		if err := copySlot(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodeUseEffectsWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// can sprint
	canSprint, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(canSprint)
	// interact vibrations
	interact, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(interact)
	// speed multiplier
	speed, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(speed)
	return nil
}

func decodeHolderSetWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyHolderSet(buf, w)
}

func decodeToolWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// rules
	ruleCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(ruleCount)
	for range int(ruleCount) {
		if err := copyToolRule(buf, w); err != nil {
			return err
		}
	}
	// default mining speed
	speed, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(speed)
	// damage per block
	damage, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(damage)
	// can destroy blocks in creative
	canDestroy, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(canDestroy)
	return nil
}

func decodeAttackRangeWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	for range 6 {
		v, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(v)
	}
	return nil
}

func decodeEquippableWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// slot
	slot, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(slot)
	// equip sound
	if err := copySoundEvent(buf, w); err != nil {
		return err
	}
	// optional asset
	if err := copyOptionalString(buf, w); err != nil {
		return err
	}
	// optional camera overlay
	if err := copyOptionalIdentifier(buf, w); err != nil {
		return err
	}
	// optional allowed entities
	if err := copyOptionalHolderSet(buf, w); err != nil {
		return err
	}
	// dispensable
	dispensable, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(dispensable)
	// swappable
	swappable, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(swappable)
	// damages on hurt
	damagesOnHurt, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(damagesOnHurt)
	// can be sheared
	canBeSheared, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(canBeSheared)
	if canBeSheared {
		if err := copySoundEvent(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodeDeathProtectionWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		if err := copyConsumeEffect(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodeBlocksAttacksWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// block delay seconds
	delay, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(delay)
	// disable cooldown scale
	cooldown, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(cooldown)
	// damage reductions
	reductionCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(reductionCount)
	for range int(reductionCount) {
		if err := copyDamageReduction(buf, w); err != nil {
			return err
		}
	}
	// item damage
	if err := copyItemDamageFunction(buf, w); err != nil {
		return err
	}
	// optional bypassed by
	if err := copyOptionalHolderSet(buf, w); err != nil {
		return err
	}
	// optional block sound
	if err := copyOptionalSoundEvent(buf, w); err != nil {
		return err
	}
	// optional disable sound
	if err := copyOptionalSoundEvent(buf, w); err != nil {
		return err
	}
	return nil
}

func decodeKineticWeaponWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// damage multiplier
	multiplier, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(multiplier)
	// optional damage conditions
	if err := copyOptionalKineticConditions(buf, w); err != nil {
		return err
	}
	// optional dismount conditions
	if err := copyOptionalKineticConditions(buf, w); err != nil {
		return err
	}
	// optional knockback conditions
	if err := copyOptionalKineticConditions(buf, w); err != nil {
		return err
	}
	// forward movement
	forward, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(forward)
	// delay ticks
	delay, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(delay)
	// optional sound
	if err := copyOptionalSoundEvent(buf, w); err != nil {
		return err
	}
	// optional hit sound
	if err := copyOptionalSoundEvent(buf, w); err != nil {
		return err
	}
	return nil
}

func decodePiercingWeaponWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// optional sound
	if err := copyOptionalSoundEvent(buf, w); err != nil {
		return err
	}
	// optional hit sound
	if err := copyOptionalSoundEvent(buf, w); err != nil {
		return err
	}
	return nil
}

func decodePotionContentsWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// optional potion
	if err := copyOptionalVarInt(buf, w); err != nil {
		return err
	}
	// optional custom color
	if err := copyOptionalInt(buf, w); err != nil {
		return err
	}
	// optional custom name
	if err := copyOptionalString(buf, w); err != nil {
		return err
	}
	// custom effects
	effectCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(effectCount)
	for range int(effectCount) {
		if err := copyStatusEffect(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodeSuspiciousStewWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		// effect ID
		id, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(id)
		// duration
		duration, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(duration)
	}
	return nil
}

func decodeWritableBookWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// pages
	pageCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(pageCount)
	for range int(pageCount) {
		// raw content
		content, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(content)
		// optional filtered content
		if err := copyOptionalString(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodeWrittenBookWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// raw title
	title, err := buf.ReadString(maxStringLen)
	if err != nil {
		return err
	}
	w.WriteString(title)
	// optional filtered title
	if err := copyOptionalString(buf, w); err != nil {
		return err
	}
	// author
	author, err := buf.ReadString(maxStringLen)
	if err != nil {
		return err
	}
	w.WriteString(author)
	// generation
	generation, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(generation)
	// pages
	pageCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(pageCount)
	for range int(pageCount) {
		// raw (NBT text component)
		if err := copyNBT(buf, w); err != nil {
			return err
		}
		// optional filtered (NBT text component)
		if err := copyOptionalNBT(buf, w); err != nil {
			return err
		}
	}
	// resolved
	resolved, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(resolved)
	return nil
}

func decodeTrimWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// material
	if err := copyTrimMaterial(buf, w); err != nil {
		return err
	}
	// pattern
	if err := copyTrimPattern(buf, w); err != nil {
		return err
	}
	return nil
}

func decodeRecipesWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		recipe, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(recipe)
	}
	return nil
}

func decodeLodestoneWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// optional global pos
	if err := copyOptionalGlobalPos(buf, w); err != nil {
		return err
	}
	// tracked
	tracked, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(tracked)
	return nil
}

func decodeFireworkExplosionWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyFireworkExplosion(buf, w)
}

func decodeProfileWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// optional name
	if err := copyOptionalString(buf, w); err != nil {
		return err
	}
	// optional UUID
	if err := copyOptionalUUID(buf, w); err != nil {
		return err
	}
	// properties
	propCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(propCount)
	for range int(propCount) {
		if err := copyGameProfileProperty(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodeBannerPatternsWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		if err := copyBannerPattern(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodePotDecorationsWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	for range 4 {
		// optional item ID
		if err := copyOptionalVarInt(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func decodeBlockStateWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		key, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(key)
		value, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(value)
	}
	return nil
}

func decodeBeesWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)
	for range int(count) {
		// entity data (NBT)
		if err := copyNBT(buf, w); err != nil {
			return err
		}
		// ticks in hive
		ticks, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(ticks)
		// min ticks in hive
		minTicks, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(minTicks)
	}
	return nil
}

// Copy helper functions

func copySlot(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(count)

	if count <= 0 {
		return nil
	}

	itemID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(itemID)

	addCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(addCount)

	removeCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(removeCount)

	for range int(addCount) {
		compID, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(compID)

		codec := componentCodecs[int32(compID)]
		if codec == nil {
			return fmt.Errorf("unknown component %d in slot", compID)
		}
		data, err := codec.DecodeWire(buf)
		if err != nil {
			return err
		}
		if _, err := w.Write(data); err != nil {
			return err
		}
	}

	for range int(removeCount) {
		compID, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(compID)
	}

	return nil
}

func copyOptionalVarInt(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(v)
	}
	return nil
}

func copyOptionalInt(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		v, err := buf.ReadInt32()
		if err != nil {
			return err
		}
		w.WriteInt32(v)
	}
	return nil
}

func copyOptionalString(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(v)
	}
	return nil
}

func copyOptionalIdentifier(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyOptionalString(buf, w)
}

func copyOptionalNBT(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		if err := copyNBT(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func copyOptionalUUID(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		uuid, err := buf.ReadUUID()
		if err != nil {
			return err
		}
		w.WriteUUID(uuid)
	}
	return nil
}

func copyHolderSet(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	if typeID == 0 {
		tag, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(tag)
	} else {
		count := int(typeID) - 1
		for range count {
			id, err := buf.ReadVarInt()
			if err != nil {
				return err
			}
			w.WriteVarInt(id)
		}
	}
	return nil
}

func copyOptionalHolderSet(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return copyHolderSet(buf, w)
	}
	return nil
}

func copySoundEvent(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	if typeID == 0 {
		id, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(id)
		if err := copyOptionalVarInt(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func copyOptionalSoundEvent(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		return copySoundEvent(buf, w)
	}
	return nil
}

func copyConsumeEffect(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	switch typeID {
	case 0: // apply_effects
		effectCount, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(effectCount)
		for range int(effectCount) {
			if err := copyStatusEffect(buf, w); err != nil {
				return err
			}
		}
		prob, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(prob)
	case 1: // remove_effects
		if err := copyHolderSet(buf, w); err != nil {
			return err
		}
	case 2: // clear_all_effects
		// no data
	case 3: // teleport_randomly
		diameter, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(diameter)
	case 4: // play_sound
		if err := copySoundEvent(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func copyStatusEffect(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	id, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(id)
	return copyStatusEffectDetails(buf, w)
}

func copyStatusEffectDetails(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// amplifier
	amplifier, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(amplifier)
	// duration
	duration, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(duration)
	// ambient
	ambient, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(ambient)
	// show particles
	showParticles, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(showParticles)
	// show icon
	showIcon, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(showIcon)
	// hidden effect (recursive optional)
	hasHidden, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasHidden)
	if hasHidden {
		return copyStatusEffectDetails(buf, w)
	}
	return nil
}

func copyToolRule(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// blocks
	if err := copyHolderSet(buf, w); err != nil {
		return err
	}
	// optional speed
	hasSpeed, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasSpeed)
	if hasSpeed {
		speed, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(speed)
	}
	// optional correct for drops
	hasCorrect, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasCorrect)
	if hasCorrect {
		correct, err := buf.ReadBool()
		if err != nil {
			return err
		}
		w.WriteBool(correct)
	}
	return nil
}

func copyBlockPredicate(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// optional blocks
	if err := copyOptionalHolderSet(buf, w); err != nil {
		return err
	}
	// optional properties
	propCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(propCount)
	for range int(propCount) {
		if err := copyPropertyMatcher(buf, w); err != nil {
			return err
		}
	}
	// optional NBT
	if err := copyOptionalNBT(buf, w); err != nil {
		return err
	}
	return nil
}

func copyPropertyMatcher(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// property name
	name, err := buf.ReadString(maxStringLen)
	if err != nil {
		return err
	}
	w.WriteString(name)
	// is exact match
	isExact, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(isExact)
	if isExact {
		value, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(value)
	} else {
		// range
		if err := copyOptionalString(buf, w); err != nil {
			return err
		}
		if err := copyOptionalString(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func copyDamageReduction(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// horizontal angle
	angle, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(angle)
	// type
	if err := copyHolderSet(buf, w); err != nil {
		return err
	}
	// base
	base, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(base)
	// factor
	factor, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(factor)
	// horizontal limit
	limit, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(limit)
	return nil
}

func copyItemDamageFunction(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// threshold
	threshold, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(threshold)
	// base
	base, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(base)
	// factor
	factor, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(factor)
	return nil
}

func copyOptionalKineticConditions(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		// max duration ticks
		maxDuration, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(maxDuration)
		// min speed
		minSpeed, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(minSpeed)
		// min relative speed
		minRelative, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(minRelative)
	}
	return nil
}

func copyTrimMaterial(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	if typeID == 0 {
		// inline
		assetName, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(assetName)
		ingredient, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(ingredient)
		itemModelIndex, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(itemModelIndex)
		overrideCount, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(overrideCount)
		for range int(overrideCount) {
			armorType, err := buf.ReadVarInt()
			if err != nil {
				return err
			}
			w.WriteVarInt(armorType)
			assetName, err := buf.ReadString(maxStringLen)
			if err != nil {
				return err
			}
			w.WriteString(assetName)
		}
		if err := copyNBT(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func copyTrimPattern(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	if typeID == 0 {
		// inline
		assetID, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(assetID)
		templateItem, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(templateItem)
		if err := copyNBT(buf, w); err != nil {
			return err
		}
		decal, err := buf.ReadBool()
		if err != nil {
			return err
		}
		w.WriteBool(decal)
	}
	return nil
}

func copyBannerPattern(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)

	if typeID == 0 {
		assetID, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(assetID)
		translationKey, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(translationKey)
	}

	color, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(color)
	return nil
}

func copyGameProfileProperty(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	name, err := buf.ReadString(maxStringLen)
	if err != nil {
		return err
	}
	w.WriteString(name)
	value, err := buf.ReadString(maxStringLen)
	if err != nil {
		return err
	}
	w.WriteString(value)
	if err := copyOptionalString(buf, w); err != nil {
		return err
	}
	return nil
}

func copyOptionalGlobalPos(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		// dimension
		dimension, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(dimension)
		// position
		pos, err := buf.ReadPosition()
		if err != nil {
			return err
		}
		w.WritePosition(pos)
	}
	return nil
}
