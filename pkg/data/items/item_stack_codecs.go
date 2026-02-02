package items

// Item component decoders and encoders.
//
// Wire format reference:
//   - https://minecraft.wiki/w/Java_Edition_protocol/Slot_data
//   - https://minecraft.wiki/w/Data_component_format
//
// Component IDs from minecraft:data_component_type registry.

import (
	"fmt"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
	"github.com/go-mclib/protocol/nbt"

	"github.com/go-mclib/data/pkg/data/registries"
)

// Component type constants are generated in component_types_gen.go

// ComponentID returns the protocol ID for a component name, or -1 if not found.
// Example: ComponentID("minecraft:damage") returns 3.
func ComponentID(name string) int32 {
	return registries.DataComponentType.Get(name)
}

// ComponentName returns the name for a component protocol ID, or empty string if not found.
// Example: ComponentName(3) returns "minecraft:damage".
func ComponentName(id int32) string {
	return registries.DataComponentType.ByID(id)
}

// max string length for identifiers/strings in protocol
const maxStringLen = 32767

// decodeComponentWire reads a component from wire format and returns raw bytes.
// This is the SlotDecoder passed to Slot.Decode.
func decodeComponentWire(buf *ns.PacketBuffer, id ns.VarInt) ([]byte, error) {
	w := ns.NewWriter()

	switch id {
	// === Simple types ===

	case ComponentMaxStackSize, ComponentMaxDamage, ComponentDamage,
		ComponentRepairCost, ComponentMapColor, ComponentMapId,
		ComponentOminousBottleAmplifier:
		// VarInt
		v, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(v)

	case ComponentMinimumAttackCharge, ComponentPotionDurationScale:
		// Float
		v, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		w.WriteFloat32(v)

	case ComponentUnbreakable, ComponentCreativeSlotLock,
		ComponentIntangibleProjectile, ComponentGlider:
		// empty marker component - no data

	case ComponentEnchantmentGlintOverride:
		// Boolean
		v, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(v)

	case ComponentRarity:
		// VarInt enum (0=common, 1=uncommon, 2=rare, 3=epic)
		v, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(v)

	case ComponentDamageType, ComponentItemModel, ComponentInstrument,
		ComponentProvidesTrimMaterial, ComponentJukeboxPlayable,
		ComponentProvidesBannerPatterns, ComponentTooltipStyle,
		ComponentNoteBlockSound, ComponentBreakSound:
		// Identifier (string)
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		w.WriteString(v)

	// === Text components (NBT) ===

	case ComponentCustomName, ComponentItemName:
		// Text Component encoded as NBT
		if err := copyNBT(buf, w); err != nil {
			return nil, err
		}

	case ComponentLore:
		// VarInt count + Text Component[] (NBT)
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for i := 0; i < int(count); i++ {
			if err := copyNBT(buf, w); err != nil {
				return nil, err
			}
		}

	// === Enchantments ===

	case ComponentEnchantments, ComponentStoredEnchantments:
		// https://minecraft.wiki/w/Data_component_format/enchantments
		// VarInt count, then [VarInt enchantID, VarInt level][], Bool showInTooltip
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for i := 0; i < int(count); i++ {
			enchantID, err := buf.ReadVarInt()
			if err != nil {
				return nil, err
			}
			level, err := buf.ReadVarInt()
			if err != nil {
				return nil, err
			}
			w.WriteVarInt(enchantID)
			w.WriteVarInt(level)
		}
		showInTooltip, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(showInTooltip)

	case ComponentEnchantable:
		// VarInt value
		v, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(v)

	// === Food & Consumable ===

	case ComponentFood:
		// https://minecraft.wiki/w/Data_component_format/food
		// VarInt nutrition, Float saturation
		nutrition, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		saturation, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(nutrition)
		w.WriteFloat32(saturation)

	case ComponentConsumable:
		// https://minecraft.wiki/w/Data_component_format/consumable
		// Float consumeSeconds, VarInt animation, SoundEvent sound,
		// Bool hasParticles, VarInt effectCount, effects[]
		consumeSeconds, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		w.WriteFloat32(consumeSeconds)

		animation, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(animation)

		if err := copySoundEvent(buf, w); err != nil {
			return nil, err
		}

		hasParticles, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(hasParticles)

		effectCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(effectCount)
		for i := 0; i < int(effectCount); i++ {
			if err := copyConsumeEffect(buf, w); err != nil {
				return nil, err
			}
		}

	case ComponentUseRemainder:
		// https://minecraft.wiki/w/Data_component_format/use_remainder
		// Slot (recursive - the remainder item)
		if err := copySlot(buf, w); err != nil {
			return nil, err
		}

	case ComponentUseCooldown:
		// https://minecraft.wiki/w/Data_component_format/use_cooldown
		// Float seconds, Optional<Identifier> cooldownGroup
		seconds, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		w.WriteFloat32(seconds)

		hasGroup, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(hasGroup)
		if hasGroup {
			group, err := buf.ReadString(maxStringLen)
			if err != nil {
				return nil, err
			}
			w.WriteString(group)
		}

	// === Combat ===

	case ComponentTool:
		// https://minecraft.wiki/w/Data_component_format/tool
		// VarInt ruleCount, rules[], VarInt damagePerBlock, Bool canDestroyInCreative
		ruleCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(ruleCount)
		for i := 0; i < int(ruleCount); i++ {
			if err := copyToolRule(buf, w); err != nil {
				return nil, err
			}
		}
		damagePerBlock, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(damagePerBlock)

		canDestroy, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(canDestroy)

	case ComponentWeapon:
		// https://minecraft.wiki/w/Data_component_format/weapon
		// VarInt itemDamagePerAttack, Float disableBlockingForSeconds
		itemDamage, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(itemDamage)

		disableBlocking, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		w.WriteFloat32(disableBlocking)

	case ComponentAttackRange:
		// https://minecraft.wiki/w/Data_component_format/attack_range
		// Float minReach, Float maxReach, Float minCreativeReach, Float maxCreativeReach,
		// Float hitboxMargin, Float mobFactor
		for i := 0; i < 6; i++ {
			v, err := buf.ReadFloat32()
			if err != nil {
				return nil, err
			}
			w.WriteFloat32(v)
		}

	case ComponentDamageResistant:
		// Identifier (damage type tag)
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		w.WriteString(v)

	case ComponentRepairable:
		// https://minecraft.wiki/w/Data_component_format/repairable
		// HolderSet (item tag or list)
		if err := copyHolderSet(buf, w); err != nil {
			return nil, err
		}

	case ComponentBlocksAttacks:
		// https://minecraft.wiki/w/Data_component_format/blocks_attacks
		// Float blockDelaySeconds, Float disableCooldownScale,
		// VarInt[] damageReductions, ItemDamageFunction itemDamage,
		// Optional<Identifier> bypassedBy, Optional<SoundEvent> blockSound,
		// Optional<SoundEvent> disableSound
		blockDelay, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		w.WriteFloat32(blockDelay)

		disableCooldown, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		w.WriteFloat32(disableCooldown)

		reductionCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(reductionCount)
		for i := 0; i < int(reductionCount); i++ {
			if err := copyDamageReduction(buf, w); err != nil {
				return nil, err
			}
		}

		if err := copyItemDamageFunction(buf, w); err != nil {
			return nil, err
		}

		if err := copyOptionalIdentifier(buf, w); err != nil { // bypassedBy
			return nil, err
		}
		if err := copyOptionalSoundEvent(buf, w); err != nil { // blockSound
			return nil, err
		}
		if err := copyOptionalSoundEvent(buf, w); err != nil { // disableSound
			return nil, err
		}

	case ComponentPiercingWeapon:
		// https://minecraft.wiki/w/Data_component_format/piercing_weapon
		// Optional<SoundEvent> hitSound, Optional<SoundEvent> sound
		if err := copyOptionalSoundEvent(buf, w); err != nil {
			return nil, err
		}
		if err := copyOptionalSoundEvent(buf, w); err != nil {
			return nil, err
		}

	case ComponentKineticWeapon:
		// https://minecraft.wiki/w/Data_component_format/kinetic_weapon
		// VarInt delayTicks, Float forwardMovement, Float damageMultiplier,
		// Optional<KineticConditions> damageConditions, knockbackConditions, dismountConditions,
		// Optional<SoundEvent> hitSound, Optional<SoundEvent> sound
		delayTicks, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(delayTicks)

		forwardMovement, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		w.WriteFloat32(forwardMovement)

		damageMultiplier, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		w.WriteFloat32(damageMultiplier)

		// damage, knockback, dismount conditions
		for i := 0; i < 3; i++ {
			if err := copyOptionalKineticConditions(buf, w); err != nil {
				return nil, err
			}
		}

		if err := copyOptionalSoundEvent(buf, w); err != nil {
			return nil, err
		}
		if err := copyOptionalSoundEvent(buf, w); err != nil {
			return nil, err
		}

	case ComponentSwingAnimation:
		// VarInt animation enum
		v, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(v)

	case ComponentDeathProtection:
		// https://minecraft.wiki/w/Data_component_format/death_protection
		// VarInt effectCount, ConsumeEffect[] effects
		effectCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(effectCount)
		for i := 0; i < int(effectCount); i++ {
			if err := copyConsumeEffect(buf, w); err != nil {
				return nil, err
			}
		}

	// === Equipment ===

	case ComponentEquippable:
		// https://minecraft.wiki/w/Data_component_format/equippable
		// VarInt slot, SoundEvent equipSound, Optional<Identifier> model,
		// Optional<Identifier> cameraOverlay, Optional<HolderSet> allowedEntities,
		// Bool dispensable, Bool swappable, Bool damageOnHurt,
		// Optional<Identifier> equipOnInteract
		slot, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(slot)

		if err := copySoundEvent(buf, w); err != nil {
			return nil, err
		}

		if err := copyOptionalIdentifier(buf, w); err != nil { // model
			return nil, err
		}
		if err := copyOptionalIdentifier(buf, w); err != nil { // cameraOverlay
			return nil, err
		}
		if err := copyOptionalHolderSet(buf, w); err != nil { // allowedEntities
			return nil, err
		}

		dispensable, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(dispensable)

		swappable, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(swappable)

		damageOnHurt, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(damageOnHurt)

		if err := copyOptionalIdentifier(buf, w); err != nil { // equipOnInteract
			return nil, err
		}

	case ComponentAttributeModifiers:
		// https://minecraft.wiki/w/Data_component_format/attribute_modifiers
		// VarInt count, AttributeModifier[] (each entry includes Display field)
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for i := 0; i < int(count); i++ {
			if err := copyAttributeModifier(buf, w); err != nil {
				return nil, fmt.Errorf("failed to read modifier %d: %w", i, err)
			}
		}

	case ComponentDyedColor:
		// https://minecraft.wiki/w/Data_component_format/dyed_color
		// Int RGB color
		v, err := buf.ReadInt32()
		if err != nil {
			return nil, err
		}
		w.WriteInt32(v)

	// === Tooltip ===

	case ComponentTooltipDisplay:
		// https://minecraft.wiki/w/Data_component_format/tooltip_display
		// Bool hideTooltip, VarInt hiddenComponentCount, VarInt[] hiddenComponents
		hideTooltip, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(hideTooltip)

		hiddenCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(hiddenCount)
		for i := 0; i < int(hiddenCount); i++ {
			hidden, err := buf.ReadVarInt()
			if err != nil {
				return nil, err
			}
			w.WriteVarInt(hidden)
		}

	case ComponentCustomModelData:
		// https://minecraft.wiki/w/Data_component_format/custom_model_data
		// VarInt floatCount, Float[], VarInt flagCount, Bool[],
		// VarInt stringCount, String[], VarInt colorCount, Int[]
		floatCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(floatCount)
		for i := 0; i < int(floatCount); i++ {
			v, err := buf.ReadFloat32()
			if err != nil {
				return nil, err
			}
			w.WriteFloat32(v)
		}

		flagCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(flagCount)
		for i := 0; i < int(flagCount); i++ {
			v, err := buf.ReadBool()
			if err != nil {
				return nil, err
			}
			w.WriteBool(v)
		}

		stringCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(stringCount)
		for i := 0; i < int(stringCount); i++ {
			v, err := buf.ReadString(maxStringLen)
			if err != nil {
				return nil, err
			}
			w.WriteString(v)
		}

		colorCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(colorCount)
		for i := 0; i < int(colorCount); i++ {
			v, err := buf.ReadInt32()
			if err != nil {
				return nil, err
			}
			w.WriteInt32(v)
		}

	// === Adventure mode ===

	case ComponentCanPlaceOn, ComponentCanBreak:
		// https://minecraft.wiki/w/Data_component_format/can_place_on
		// VarInt predicateCount, BlockPredicate[]
		predicateCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(predicateCount)
		for i := 0; i < int(predicateCount); i++ {
			if err := copyBlockPredicate(buf, w); err != nil {
				return nil, err
			}
		}

	// === Map ===

	case ComponentMapDecorations:
		// https://minecraft.wiki/w/Data_component_format/map_decorations
		// VarInt count, [String key, MapDecoration value][]
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for i := 0; i < int(count); i++ {
			key, err := buf.ReadString(maxStringLen)
			if err != nil {
				return nil, err
			}
			w.WriteString(key)
			if err := copyMapDecoration(buf, w); err != nil {
				return nil, err
			}
		}

	case ComponentMapPostProcessing:
		// VarInt enum (0=lock, 1=scale)
		v, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(v)

	// === Containers ===

	case ComponentChargedProjectiles, ComponentBundleContents:
		// https://minecraft.wiki/w/Data_component_format/bundle_contents
		// VarInt itemCount, Slot[]
		itemCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(itemCount)
		for i := 0; i < int(itemCount); i++ {
			if err := copySlot(buf, w); err != nil {
				return nil, err
			}
		}

	case ComponentContainer:
		// https://minecraft.wiki/w/Data_component_format/container
		// VarInt slotCount, [VarInt slotIndex, Slot item][]
		slotCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(slotCount)
		for i := 0; i < int(slotCount); i++ {
			slotIndex, err := buf.ReadVarInt()
			if err != nil {
				return nil, err
			}
			w.WriteVarInt(slotIndex)
			if err := copySlot(buf, w); err != nil {
				return nil, err
			}
		}

	case ComponentContainerLoot:
		// https://minecraft.wiki/w/Data_component_format/container_loot
		// Identifier lootTable, Long seed
		lootTable, err := buf.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		w.WriteString(lootTable)

		seed, err := buf.ReadInt64()
		if err != nil {
			return nil, err
		}
		w.WriteInt64(seed)

	// === Potions ===

	case ComponentPotionContents:
		// https://minecraft.wiki/w/Data_component_format/potion_contents
		// Optional<VarInt> potionID, Optional<Int> customColor,
		// VarInt customEffectCount, StatusEffect[], Optional<Identifier> customName
		if err := copyOptionalVarInt(buf, w); err != nil { // potionID
			return nil, err
		}
		if err := copyOptionalInt(buf, w); err != nil { // customColor
			return nil, err
		}

		effectCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(effectCount)
		for i := 0; i < int(effectCount); i++ {
			if err := copyStatusEffect(buf, w); err != nil {
				return nil, err
			}
		}

		if err := copyOptionalIdentifier(buf, w); err != nil { // customName
			return nil, err
		}

	case ComponentSuspiciousStewEffects:
		// https://minecraft.wiki/w/Data_component_format/suspicious_stew_effects
		// VarInt count, [VarInt effectID, VarInt duration][]
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for i := 0; i < int(count); i++ {
			effectID, err := buf.ReadVarInt()
			if err != nil {
				return nil, err
			}
			w.WriteVarInt(effectID)
			duration, err := buf.ReadVarInt()
			if err != nil {
				return nil, err
			}
			w.WriteVarInt(duration)
		}

	// === Effects ===

	case ComponentUseEffects:
		// https://minecraft.wiki/w/Data_component_format/use_effects
		// VarInt effectCount, ConsumeEffect[]
		effectCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(effectCount)
		for i := 0; i < int(effectCount); i++ {
			if err := copyConsumeEffect(buf, w); err != nil {
				return nil, err
			}
		}

	// === Books ===

	case ComponentWritableBookContent:
		// https://minecraft.wiki/w/Data_component_format/writable_book_content
		// VarInt pageCount, [String rawText, Optional<String> filteredText][]
		pageCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(pageCount)
		for i := 0; i < int(pageCount); i++ {
			raw, err := buf.ReadString(maxStringLen)
			if err != nil {
				return nil, err
			}
			w.WriteString(raw)
			if err := copyOptionalString(buf, w); err != nil {
				return nil, err
			}
		}

	case ComponentWrittenBookContent:
		// https://minecraft.wiki/w/Data_component_format/written_book_content
		// FilteredText title, String author, VarInt generation,
		// VarInt pageCount, [TextComponent rawText, Optional<TextComponent> filtered][],
		// Bool resolved
		rawTitle, err := buf.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		w.WriteString(rawTitle)
		if err := copyOptionalString(buf, w); err != nil { // filtered title
			return nil, err
		}

		author, err := buf.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		w.WriteString(author)

		generation, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(generation)

		pageCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(pageCount)
		for i := 0; i < int(pageCount); i++ {
			if err := copyNBT(buf, w); err != nil {
				return nil, err
			}
			if err := copyOptionalNBT(buf, w); err != nil {
				return nil, err
			}
		}

		resolved, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(resolved)

	// === Fireworks ===

	case ComponentFireworkExplosion:
		// https://minecraft.wiki/w/Data_component_format/firework_explosion
		// VarInt shape, VarInt colorCount, Int[] colors, VarInt fadeColorCount, Int[] fadeColors,
		// Bool hasTrail, Bool hasTwinkle
		shape, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(shape)

		colorCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(colorCount)
		for i := 0; i < int(colorCount); i++ {
			color, err := buf.ReadInt32()
			if err != nil {
				return nil, err
			}
			w.WriteInt32(color)
		}

		fadeColorCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(fadeColorCount)
		for i := 0; i < int(fadeColorCount); i++ {
			color, err := buf.ReadInt32()
			if err != nil {
				return nil, err
			}
			w.WriteInt32(color)
		}

		hasTrail, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(hasTrail)

		hasTwinkle, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(hasTwinkle)

	case ComponentFireworks:
		// https://minecraft.wiki/w/Data_component_format/fireworks
		// VarInt flightDuration, VarInt explosionCount, FireworkExplosion[]
		flightDuration, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(flightDuration)

		explosionCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(explosionCount)
		for i := 0; i < int(explosionCount); i++ {
			if err := copyFireworkExplosion(buf, w); err != nil {
				return nil, err
			}
		}

	// === Decorations ===

	case ComponentTrim:
		// https://minecraft.wiki/w/Data_component_format/trim
		// IDOrInline material, IDOrInline pattern
		if err := copyTrimMaterial(buf, w); err != nil {
			return nil, err
		}
		if err := copyTrimPattern(buf, w); err != nil {
			return nil, err
		}

	case ComponentBannerPatterns:
		// https://minecraft.wiki/w/Data_component_format/banner_patterns
		// VarInt count, [IDOrInline pattern, VarInt color][]
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for i := 0; i < int(count); i++ {
			if err := copyBannerPattern(buf, w); err != nil {
				return nil, err
			}
			color, err := buf.ReadVarInt()
			if err != nil {
				return nil, err
			}
			w.WriteVarInt(color)
		}

	case ComponentBaseColor:
		// VarInt dye color
		v, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(v)

	case ComponentPotDecorations:
		// https://minecraft.wiki/w/Data_component_format/pot_decorations
		// 4x Optional<VarInt> (item IDs for back, left, right, front)
		for i := 0; i < 4; i++ {
			if err := copyOptionalVarInt(buf, w); err != nil {
				return nil, err
			}
		}

	// === Player data ===

	case ComponentProfile:
		// https://minecraft.wiki/w/Data_component_format/profile
		// Optional<String> name, Optional<UUID> uuid, VarInt propertyCount, GameProfileProperty[]
		if err := copyOptionalString(buf, w); err != nil { // name
			return nil, err
		}
		if err := copyOptionalUUID(buf, w); err != nil { // uuid
			return nil, err
		}

		propertyCount, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(propertyCount)
		for i := 0; i < int(propertyCount); i++ {
			if err := copyGameProfileProperty(buf, w); err != nil {
				return nil, err
			}
		}

	case ComponentRecipes:
		// VarInt count, Identifier[]
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for i := 0; i < int(count); i++ {
			recipe, err := buf.ReadString(maxStringLen)
			if err != nil {
				return nil, err
			}
			w.WriteString(recipe)
		}

	case ComponentLodestoneTracker:
		// https://minecraft.wiki/w/Data_component_format/lodestone_tracker
		// Optional<GlobalPos> target, Bool tracked
		if err := copyOptionalGlobalPos(buf, w); err != nil {
			return nil, err
		}
		tracked, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		w.WriteBool(tracked)

	// === Block/Entity data (NBT passthrough) ===

	case ComponentCustomData, ComponentEntityData, ComponentBucketEntityData,
		ComponentBlockEntityData, ComponentDebugStickState, ComponentLock:
		// NBT compound
		if err := copyNBT(buf, w); err != nil {
			return nil, err
		}

	case ComponentBlockState:
		// https://minecraft.wiki/w/Data_component_format/block_state
		// VarInt count, [String key, String value][]
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for i := 0; i < int(count); i++ {
			key, err := buf.ReadString(maxStringLen)
			if err != nil {
				return nil, err
			}
			w.WriteString(key)
			value, err := buf.ReadString(maxStringLen)
			if err != nil {
				return nil, err
			}
			w.WriteString(value)
		}

	case ComponentBees:
		// https://minecraft.wiki/w/Data_component_format/bees
		// VarInt count, [NBT entityData, VarInt ticksInHive, VarInt minTicksInHive][]
		count, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(count)
		for i := 0; i < int(count); i++ {
			if err := copyNBT(buf, w); err != nil {
				return nil, err
			}
			ticksInHive, err := buf.ReadVarInt()
			if err != nil {
				return nil, err
			}
			w.WriteVarInt(ticksInHive)
			minTicks, err := buf.ReadVarInt()
			if err != nil {
				return nil, err
			}
			w.WriteVarInt(minTicks)
		}

	// === Entity variant components ===
	// used for bucket items and spawn eggs to store entity variants

	case ComponentVillagerVariant, ComponentWolfVariant, ComponentWolfSoundVariant,
		ComponentFoxVariant, ComponentParrotVariant, ComponentTropicalFishPattern,
		ComponentMooshroomVariant, ComponentRabbitVariant, ComponentPigVariant,
		ComponentCowVariant, ComponentChickenVariant, ComponentZombieNautilusVariant,
		ComponentFrogVariant, ComponentHorseVariant, ComponentPaintingVariant,
		ComponentLlamaVariant, ComponentAxolotlVariant, ComponentCatVariant:
		// IDOrInline<RegistryEntry> - VarInt for registry ID or inline definition
		if err := copyIDOrTag(buf, w); err != nil {
			return nil, err
		}

	case ComponentWolfCollar, ComponentCatCollar, ComponentSheepColor, ComponentShulkerColor,
		ComponentTropicalFishBaseColor, ComponentTropicalFishPatternColor:
		// VarInt dye color
		v, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(v)

	case ComponentSalmonSize:
		// VarInt size enum
		v, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		w.WriteVarInt(v)

	default:
		return nil, fmt.Errorf("unknown component ID: %d", id)
	}

	return w.Bytes(), nil
}

// === Helper functions for copying complex structures ===

func copyNBT(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// read NBT tag from buffer and write to writer
	// use network format (nameless root)
	reader := nbt.NewReaderFrom(buf.Reader())
	tag, _, err := reader.ReadTag(true) // network format
	if err != nil {
		return err
	}

	writer := nbt.NewWriterTo(w.Writer())
	return writer.WriteTag(tag, "", true) // network format
}

// decodeItemName reads an NBT text component and returns an ItemNameComponent.
func decodeItemName(buf *ns.PacketBuffer) (*ItemNameComponent, error) {
	reader := nbt.NewReaderFrom(buf.Reader())
	tag, _, err := reader.ReadTag(true) // network format
	if err != nil {
		return nil, err
	}

	// text component can be a string (literal) or compound (with translate/etc)
	switch v := tag.(type) {
	case nbt.String:
		return &ItemNameComponent{Text: string(v)}, nil
	case nbt.Compound:
		name := &ItemNameComponent{}
		if t, ok := v["text"].(nbt.String); ok {
			name.Text = string(t)
		}
		if t, ok := v["translate"].(nbt.String); ok {
			name.Translate = string(t)
		}
		return name, nil
	default:
		return &ItemNameComponent{}, nil
	}
}

// decodeAttributeModifier reads an attribute modifier entry.
func decodeAttributeModifier(buf *ns.PacketBuffer) (AttributeModifier, error) {
	var mod AttributeModifier

	// attribute ID (registry reference)
	attrID, err := buf.ReadVarInt()
	if err != nil {
		return mod, err
	}
	mod.Type = registries.Attribute.ByID(int32(attrID))

	// modifier ID (Identifier string)
	modID, err := buf.ReadString(maxStringLen)
	if err != nil {
		return mod, err
	}
	mod.ID = string(modID)

	// amount (Double)
	amount, err := buf.ReadFloat64()
	if err != nil {
		return mod, err
	}
	mod.Amount = float64(amount)

	// operation (VarInt)
	operation, err := buf.ReadVarInt()
	if err != nil {
		return mod, err
	}
	operations := []string{"add_value", "add_multiplied_base", "add_multiplied_total"}
	if int(operation) < len(operations) {
		mod.Operation = operations[operation]
	}

	// slot (VarInt - equipment slot group)
	slot, err := buf.ReadVarInt()
	if err != nil {
		return mod, err
	}
	slots := []string{"any", "hand", "mainhand", "offhand", "armor", "feet", "legs", "chest", "head", "body"}
	if int(slot) < len(slots) {
		mod.Slot = slots[slot]
	}

	// display type (VarInt)
	displayType, err := buf.ReadVarInt()
	if err != nil {
		return mod, err
	}
	if displayType == 2 {
		// OVERRIDE includes a Component (NBT) - skip for now
		reader := nbt.NewReaderFrom(buf.Reader())
		_, _, _ = reader.ReadTag(true)
	}

	return mod, nil
}

func copySlot(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// recursive slot read - count, then item data if count > 0
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

	for i := 0; i < int(addCount); i++ {
		compID, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(compID)

		compData, err := decodeComponentWire(buf, compID)
		if err != nil {
			return err
		}
		w.Write(compData)
	}

	for i := 0; i < int(removeCount); i++ {
		removeID, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(removeID)
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
		return copyNBT(buf, w)
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
	// HolderSet: VarInt type (0 = tag, >0 = direct list of size type-1)
	typeOrSize, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeOrSize)

	if typeOrSize == 0 {
		// tag reference
		tag, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(tag)
	} else {
		// direct list
		for i := 0; i < int(typeOrSize)-1; i++ {
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

func copyIDOrTag(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// IDOrTag: VarInt (0 = inline tag follows, >0 = registry ID + 1)
	typeOrID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeOrID)

	if typeOrID == 0 {
		// inline tag - identifier
		tag, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(tag)
	}
	// else: typeOrID - 1 is the registry ID, no more data
	return nil
}

func copySoundEvent(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// SoundEvent: VarInt (0 = inline, >0 = registry ID + 1)
	typeOrID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeOrID)

	if typeOrID == 0 {
		// inline: identifier, Optional<Float> fixedRange
		id, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(id)

		hasRange, err := buf.ReadBool()
		if err != nil {
			return err
		}
		w.WriteBool(hasRange)
		if hasRange {
			fixedRange, err := buf.ReadFloat32()
			if err != nil {
				return err
			}
			w.WriteFloat32(fixedRange)
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
	// ConsumeEffect: VarInt type, then type-specific data
	effectType, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(effectType)

	switch effectType {
	case 0: // apply_effects
		effectCount, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(effectCount)
		for i := 0; i < int(effectCount); i++ {
			if err := copyStatusEffect(buf, w); err != nil {
				return err
			}
		}
		probability, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(probability)

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
	// StatusEffectInstance: VarInt effectID, StatusEffectDetails
	effectID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(effectID)

	return copyStatusEffectDetails(buf, w)
}

func copyStatusEffectDetails(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// StatusEffectDetails: VarInt amplifier, VarInt duration, Bool ambient,
	// Bool showParticles, Bool showIcon, Optional<StatusEffectDetails> hiddenEffect
	amplifier, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(amplifier)

	duration, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(duration)

	ambient, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(ambient)

	showParticles, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(showParticles)

	showIcon, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(showIcon)

	// recursive hidden effect
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
	// ToolRule: HolderSet blocks, Optional<Float> speed, Optional<Bool> correctForDrops
	if err := copyHolderSet(buf, w); err != nil {
		return err
	}

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

func copyAttributeModifier(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// AttributeModifier: VarInt attributeID, AttributeModifierData, VarInt slot, Display
	attributeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(attributeID)

	// AttributeModifierData: Identifier id, Double amount, VarInt operation
	modID, err := buf.ReadString(maxStringLen)
	if err != nil {
		return fmt.Errorf("failed to read string data: %w", err)
	}
	w.WriteString(modID)

	amount, err := buf.ReadFloat64()
	if err != nil {
		return err
	}
	w.WriteFloat64(amount)

	operation, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(operation)

	slot, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(slot)

	// Display: type-dispatched (0=DEFAULT, 1=HIDDEN, 2=OVERRIDE with Component)
	displayType, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(displayType)

	if displayType == 2 { // OVERRIDE includes a Component (NBT)
		if err := copyNBT(buf, w); err != nil {
			return err
		}
	}

	return nil
}

func copyBlockPredicate(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// BlockPredicate: Optional<HolderSet> blocks, Optional<PropertyMatcher[]> properties, Optional<NBT> nbt
	if err := copyOptionalHolderSet(buf, w); err != nil { // blocks
		return err
	}

	hasProperties, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasProperties)
	if hasProperties {
		propCount, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(propCount)
		for i := 0; i < int(propCount); i++ {
			if err := copyPropertyMatcher(buf, w); err != nil {
				return err
			}
		}
	}

	if err := copyOptionalNBT(buf, w); err != nil { // nbt
		return err
	}
	return nil
}

func copyPropertyMatcher(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// PropertyMatcher: String name, Bool isExact, if exact: String value, else: String min, String max
	name, err := buf.ReadString(maxStringLen)
	if err != nil {
		return err
	}
	w.WriteString(name)

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
		minVal, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(minVal)
		maxVal, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(maxVal)
	}
	return nil
}

func copyMapDecoration(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// MapDecoration: VarInt type, Double x, Double z, Float rotation
	decorationType, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(decorationType)

	x, err := buf.ReadFloat64()
	if err != nil {
		return err
	}
	w.WriteFloat64(x)

	z, err := buf.ReadFloat64()
	if err != nil {
		return err
	}
	w.WriteFloat64(z)

	rotation, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(rotation)

	return nil
}

func copyFireworkExplosion(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	shape, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(shape)

	colorCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(colorCount)
	for j := 0; j < int(colorCount); j++ {
		color, err := buf.ReadInt32()
		if err != nil {
			return err
		}
		w.WriteInt32(color)
	}

	fadeColorCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(fadeColorCount)
	for j := 0; j < int(fadeColorCount); j++ {
		color, err := buf.ReadInt32()
		if err != nil {
			return err
		}
		w.WriteInt32(color)
	}

	hasTrail, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasTrail)

	hasTwinkle, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(hasTwinkle)

	return nil
}

func copyTrimMaterial(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// IDOrInline<TrimMaterial>
	typeOrID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeOrID)

	if typeOrID == 0 {
		// inline TrimMaterial: Identifier assetName, VarInt ingredient,
		// Float itemModelIndex, Map<ArmorMaterial, Identifier> overrideArmorMaterials, TextComponent description
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
		for i := 0; i < int(overrideCount); i++ {
			material, err := buf.ReadVarInt()
			if err != nil {
				return err
			}
			w.WriteVarInt(material)
			override, err := buf.ReadString(maxStringLen)
			if err != nil {
				return err
			}
			w.WriteString(override)
		}

		if err := copyNBT(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func copyTrimPattern(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// IDOrInline<TrimPattern>
	typeOrID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeOrID)

	if typeOrID == 0 {
		// inline TrimPattern: Identifier assetID, VarInt templateItem, TextComponent description, Bool decal
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
	// IDOrInline<BannerPattern>
	typeOrID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeOrID)

	if typeOrID == 0 {
		// inline BannerPattern: Identifier assetID, String translationKey
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
	return nil
}

func copyGameProfileProperty(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// GameProfileProperty: String name, String value, Optional<String> signature
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

	return copyOptionalString(buf, w)
}

func copyOptionalGlobalPos(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		// GlobalPos: Identifier dimension, Position pos
		dimension, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		w.WriteString(dimension)

		pos, err := buf.ReadInt64() // packed position
		if err != nil {
			return err
		}
		w.WriteInt64(pos)
	}
	return nil
}

func copyDamageReduction(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// DamageReduction: HolderSet types, Float base, Float factor, Float horizontalBlockingAngle
	if err := copyHolderSet(buf, w); err != nil {
		return err
	}

	base, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(base)

	factor, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(factor)

	angle, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(angle)

	return nil
}

func copyItemDamageFunction(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	// ItemDamageFunction: Float threshold, Float base, Float factor
	threshold, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(threshold)

	base, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	w.WriteFloat32(base)

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
		// KineticConditions: Float minSpeed, Float minRelativeSpeed, VarInt maxDurationTicks
		minSpeed, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(minSpeed)

		minRelativeSpeed, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		w.WriteFloat32(minRelativeSpeed)

		maxDuration, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(maxDuration)
	}
	return nil
}

// === Component application and encoding ===

// applyComponent applies raw component bytes to a Components struct.
// Currently only implements commonly used components.
func applyComponent(c *Components, id int32, data []byte) error {
	if c == nil {
		return nil
	}

	buf := ns.NewReader(data)

	switch id {
	case ComponentMaxStackSize:
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		c.MaxStackSize = int32(v)

	case ComponentMaxDamage:
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		c.MaxDamage = int32(v)

	case ComponentDamage:
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		c.Damage = int32(v)

	case ComponentRepairCost:
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		c.RepairCost = int32(v)

	case ComponentMapColor:
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		c.MapColor = int32(v)

	case ComponentOminousBottleAmplifier:
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		c.OminousBottleAmplifier = int32(v)

	case ComponentMinimumAttackCharge:
		v, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		c.MinimumAttackCharge = float64(v)

	case ComponentPotionDurationScale:
		v, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		c.PotionDurationScale = float64(v)

	case ComponentGlider:
		c.Glider = true

	case ComponentRarity:
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		rarities := []string{"common", "uncommon", "rare", "epic"}
		if int(v) < len(rarities) {
			c.Rarity = rarities[v]
		}

	case ComponentDamageType:
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		c.DamageType = string(v)

	case ComponentItemModel:
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		c.ItemModel = string(v)

	case ComponentInstrument:
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		c.Instrument = string(v)

	case ComponentProvidesTrimMaterial:
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		c.ProvidesTrimMaterial = string(v)

	case ComponentJukeboxPlayable:
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		c.JukeboxPlayable = string(v)

	case ComponentProvidesBannerPatterns:
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		c.ProvidesBannerPatterns = string(v)

	case ComponentBreakSound:
		v, err := buf.ReadString(maxStringLen)
		if err != nil {
			return err
		}
		c.BreakSound = string(v)

	case ComponentFood:
		nutrition, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		saturation, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		c.Food = &Food{
			Nutrition:  int32(nutrition),
			Saturation: float64(saturation),
		}

	case ComponentWeapon:
		itemDamage, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		disableBlocking, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		c.Weapon = &Weapon{
			ItemDamagePerAttack:       int32(itemDamage),
			DisableBlockingForSeconds: float64(disableBlocking),
		}

	case ComponentEnchantable:
		v, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		c.Enchantable = &Enchantable{Value: int32(v)}

	case ComponentUseCooldown:
		seconds, err := buf.ReadFloat32()
		if err != nil {
			return err
		}
		c.UseCooldown = &UseCooldown{Seconds: float64(seconds)}

	case ComponentFireworks:
		flightDuration, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		c.Fireworks = &Fireworks{FlightDuration: int32(flightDuration)}

	case ComponentUnbreakable:
		// empty marker component
		c.Unbreakable = true

	case ComponentCustomName:
		// NBT text component
		name, err := decodeItemName(buf)
		if err != nil {
			return err
		}
		c.CustomName = name

	case ComponentAttributeModifiers:
		count, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		modifiers := make([]AttributeModifier, 0, count)
		for i := 0; i < int(count); i++ {
			mod, err := decodeAttributeModifier(buf)
			if err != nil {
				return fmt.Errorf("modifier %d: %w", i, err)
			}
			modifiers = append(modifiers, mod)
		}
		c.AttributeModifiers = modifiers

	case ComponentTooltipDisplay:
		hideTooltip, err := buf.ReadBool()
		if err != nil {
			return err
		}
		hiddenCount, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		hidden := make([]int32, 0, hiddenCount)
		for i := 0; i < int(hiddenCount); i++ {
			compID, err := buf.ReadVarInt()
			if err != nil {
				return err
			}
			hidden = append(hidden, int32(compID))
		}
		c.TooltipDisplay = &TooltipDisplay{
			HideTooltip:      bool(hideTooltip),
			HiddenComponents: hidden,
		}

	// complex components are passed through as raw data for now
	default:
		// unknown or complex component - ignore for typed access
	}

	return nil
}

// clearComponent sets a component to its zero value.
func clearComponent(c *Components, id int32) {
	if c == nil {
		return
	}

	switch id {
	case ComponentMaxStackSize:
		c.MaxStackSize = 0
	case ComponentMaxDamage:
		c.MaxDamage = 0
	case ComponentDamage:
		c.Damage = 0
	case ComponentRepairCost:
		c.RepairCost = 0
	case ComponentMapColor:
		c.MapColor = 0
	case ComponentOminousBottleAmplifier:
		c.OminousBottleAmplifier = 0
	case ComponentMinimumAttackCharge:
		c.MinimumAttackCharge = 0
	case ComponentPotionDurationScale:
		c.PotionDurationScale = 0
	case ComponentGlider:
		c.Glider = false
	case ComponentRarity:
		c.Rarity = ""
	case ComponentDamageType:
		c.DamageType = ""
	case ComponentItemModel:
		c.ItemModel = ""
	case ComponentInstrument:
		c.Instrument = ""
	case ComponentProvidesTrimMaterial:
		c.ProvidesTrimMaterial = ""
	case ComponentJukeboxPlayable:
		c.JukeboxPlayable = ""
	case ComponentProvidesBannerPatterns:
		c.ProvidesBannerPatterns = ""
	case ComponentBreakSound:
		c.BreakSound = ""
	case ComponentFood:
		c.Food = nil
	case ComponentWeapon:
		c.Weapon = nil
	case ComponentEnchantable:
		c.Enchantable = nil
	case ComponentUseCooldown:
		c.UseCooldown = nil
	case ComponentFireworks:
		c.Fireworks = nil
	case ComponentTool:
		c.Tool = nil
	case ComponentConsumable:
		c.Consumable = nil
	case ComponentRepairable:
		c.Repairable = nil
	case ComponentDamageResistant:
		c.DamageResistant = nil
	case ComponentEquippable:
		c.Equippable = nil
	case ComponentAttributeModifiers:
		c.AttributeModifiers = nil
	case ComponentEnchantments:
		c.Enchantments = nil
	case ComponentStoredEnchantments:
		c.StoredEnchantments = nil
	case ComponentLore:
		c.Lore = nil
	case ComponentItemName:
		c.ItemName = nil
	case ComponentBlocksAttacks:
		c.BlocksAttacks = nil
	case ComponentDeathProtection:
		c.DeathProtection = nil
	case ComponentKineticWeapon:
		c.KineticWeapon = nil
	case ComponentPiercingWeapon:
		c.PiercingWeapon = nil
	case ComponentTooltipDisplay:
		c.TooltipDisplay = nil
	case ComponentUnbreakable:
		c.Unbreakable = false
	case ComponentCustomName:
		c.CustomName = nil
	case ComponentUseEffects:
		c.UseEffects = nil
	case ComponentPotionContents:
		c.PotionContents = nil
	case ComponentUseRemainder:
		c.UseRemainder = nil
	case ComponentContainer:
		c.Container = nil
	case ComponentRecipes:
		c.Recipes = nil
	}
}

// componentDiffers checks if a component in c differs from defaults.
// Returns (differs bool, hasValue bool).
func componentDiffers(c *Components, defaults *Components, id int32) (bool, bool) {
	if c == nil && defaults == nil {
		return false, false
	}
	if c == nil {
		return defaults != nil, false
	}
	if defaults == nil {
		defaults = &Components{}
	}

	switch id {
	case ComponentMaxStackSize:
		return c.MaxStackSize != defaults.MaxStackSize, c.MaxStackSize != 0
	case ComponentMaxDamage:
		return c.MaxDamage != defaults.MaxDamage, c.MaxDamage != 0
	case ComponentDamage:
		return c.Damage != defaults.Damage, c.Damage != 0
	case ComponentRepairCost:
		return c.RepairCost != defaults.RepairCost, c.RepairCost != 0
	case ComponentMapColor:
		return c.MapColor != defaults.MapColor, c.MapColor != 0
	case ComponentOminousBottleAmplifier:
		return c.OminousBottleAmplifier != defaults.OminousBottleAmplifier, c.OminousBottleAmplifier != 0
	case ComponentMinimumAttackCharge:
		return c.MinimumAttackCharge != defaults.MinimumAttackCharge, c.MinimumAttackCharge != 0
	case ComponentPotionDurationScale:
		return c.PotionDurationScale != defaults.PotionDurationScale, c.PotionDurationScale != 0
	case ComponentGlider:
		return c.Glider != defaults.Glider, c.Glider
	case ComponentRarity:
		return c.Rarity != defaults.Rarity, c.Rarity != ""
	case ComponentDamageType:
		return c.DamageType != defaults.DamageType, c.DamageType != ""
	case ComponentItemModel:
		return c.ItemModel != defaults.ItemModel, c.ItemModel != ""
	case ComponentInstrument:
		return c.Instrument != defaults.Instrument, c.Instrument != ""
	case ComponentProvidesTrimMaterial:
		return c.ProvidesTrimMaterial != defaults.ProvidesTrimMaterial, c.ProvidesTrimMaterial != ""
	case ComponentJukeboxPlayable:
		return c.JukeboxPlayable != defaults.JukeboxPlayable, c.JukeboxPlayable != ""
	case ComponentProvidesBannerPatterns:
		return c.ProvidesBannerPatterns != defaults.ProvidesBannerPatterns, c.ProvidesBannerPatterns != ""
	case ComponentBreakSound:
		return c.BreakSound != defaults.BreakSound, c.BreakSound != ""
	case ComponentFood:
		cHas := c.Food != nil
		dHas := defaults.Food != nil
		if cHas != dHas {
			return true, cHas
		}
		if cHas && dHas {
			return *c.Food != *defaults.Food, true
		}
		return false, false
	case ComponentWeapon:
		cHas := c.Weapon != nil
		dHas := defaults.Weapon != nil
		if cHas != dHas {
			return true, cHas
		}
		if cHas && dHas {
			return *c.Weapon != *defaults.Weapon, true
		}
		return false, false
	case ComponentEnchantable:
		cHas := c.Enchantable != nil
		dHas := defaults.Enchantable != nil
		if cHas != dHas {
			return true, cHas
		}
		if cHas && dHas {
			return *c.Enchantable != *defaults.Enchantable, true
		}
		return false, false
	case ComponentUseCooldown:
		cHas := c.UseCooldown != nil
		dHas := defaults.UseCooldown != nil
		if cHas != dHas {
			return true, cHas
		}
		if cHas && dHas {
			return *c.UseCooldown != *defaults.UseCooldown, true
		}
		return false, false
	case ComponentFireworks:
		cHas := c.Fireworks != nil
		dHas := defaults.Fireworks != nil
		if cHas != dHas {
			return true, cHas
		}
		if cHas && dHas {
			return *c.Fireworks != *defaults.Fireworks, true
		}
		return false, false
	default:
		// for complex components we don't track differences yet
		return false, false
	}
}

// encodeComponent encodes a component to raw bytes.
func encodeComponent(c *Components, id int32) ([]byte, error) {
	if c == nil {
		return nil, nil
	}

	w := ns.NewWriter()

	switch id {
	case ComponentMaxStackSize:
		w.WriteVarInt(ns.VarInt(c.MaxStackSize))
	case ComponentMaxDamage:
		w.WriteVarInt(ns.VarInt(c.MaxDamage))
	case ComponentDamage:
		w.WriteVarInt(ns.VarInt(c.Damage))
	case ComponentRepairCost:
		w.WriteVarInt(ns.VarInt(c.RepairCost))
	case ComponentMapColor:
		w.WriteVarInt(ns.VarInt(c.MapColor))
	case ComponentOminousBottleAmplifier:
		w.WriteVarInt(ns.VarInt(c.OminousBottleAmplifier))
	case ComponentMinimumAttackCharge:
		w.WriteFloat32(ns.Float32(c.MinimumAttackCharge))
	case ComponentPotionDurationScale:
		w.WriteFloat32(ns.Float32(c.PotionDurationScale))
	case ComponentGlider:
		// empty marker - no data
	case ComponentRarity:
		rarities := map[string]int32{"common": 0, "uncommon": 1, "rare": 2, "epic": 3}
		w.WriteVarInt(ns.VarInt(rarities[c.Rarity]))
	case ComponentDamageType:
		w.WriteString(ns.String(c.DamageType))
	case ComponentItemModel:
		w.WriteString(ns.String(c.ItemModel))
	case ComponentInstrument:
		w.WriteString(ns.String(c.Instrument))
	case ComponentProvidesTrimMaterial:
		w.WriteString(ns.String(c.ProvidesTrimMaterial))
	case ComponentJukeboxPlayable:
		w.WriteString(ns.String(c.JukeboxPlayable))
	case ComponentProvidesBannerPatterns:
		w.WriteString(ns.String(c.ProvidesBannerPatterns))
	case ComponentBreakSound:
		w.WriteString(ns.String(c.BreakSound))
	case ComponentFood:
		if c.Food != nil {
			w.WriteVarInt(ns.VarInt(c.Food.Nutrition))
			w.WriteFloat32(ns.Float32(c.Food.Saturation))
		}
	case ComponentWeapon:
		if c.Weapon != nil {
			w.WriteVarInt(ns.VarInt(c.Weapon.ItemDamagePerAttack))
			w.WriteFloat32(ns.Float32(c.Weapon.DisableBlockingForSeconds))
		}
	case ComponentEnchantable:
		if c.Enchantable != nil {
			w.WriteVarInt(ns.VarInt(c.Enchantable.Value))
		}
	case ComponentUseCooldown:
		if c.UseCooldown != nil {
			w.WriteFloat32(ns.Float32(c.UseCooldown.Seconds))
			w.WriteBool(false) // no cooldown group
		}
	case ComponentFireworks:
		if c.Fireworks != nil {
			w.WriteVarInt(ns.VarInt(c.Fireworks.FlightDuration))
			w.WriteVarInt(0) // no explosions
		}
	default:
		return nil, fmt.Errorf("cannot encode component %d", id)
	}

	return w.Bytes(), nil
}
