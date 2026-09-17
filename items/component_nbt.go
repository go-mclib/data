package items

import (
	"bytes"
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"

	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

var aquaticVariantComponentNames = map[string]map[int32]string{
	"minecraft:axolotl/variant": {
		0: "lucy", 1: "wild", 2: "gold", 3: "cyan", 4: "blue",
	},
	"minecraft:tropical_fish/pattern": {
		0: "kob", 256: "sunstreak", 512: "snooper", 768: "dasher", 1024: "brinely", 1280: "spotty",
		1: "flopper", 257: "stripey", 513: "glitter", 769: "blockfish", 1025: "betty", 1281: "clayfish",
	},
	"minecraft:tropical_fish/base_color": {
		0: "white", 1: "orange", 2: "magenta", 3: "light_blue", 4: "yellow", 5: "lime", 6: "pink", 7: "gray",
		8: "light_gray", 9: "cyan", 10: "purple", 11: "blue", 12: "brown", 13: "green", 14: "red", 15: "black",
	},
	"minecraft:tropical_fish/pattern_color": {
		0: "white", 1: "orange", 2: "magenta", 3: "light_blue", 4: "yellow", 5: "lime", 6: "pink", 7: "gray",
		8: "light_gray", 9: "cyan", 10: "purple", 11: "blue", 12: "brown", 13: "green", 14: "red", 15: "black",
	},
}

func aquaticVariantComponentID(component, name string) (int32, bool) {
	for id, candidate := range aquaticVariantComponentNames[component] {
		if candidate == name {
			return id, true
		}
	}
	return 0, false
}

// SupportsNBTComponent reports whether a component has an NBT/wire bridge.
func SupportsNBTComponent(name string) bool {
	switch name {
	case "minecraft:attack_animation", "minecraft:interact_animation", "minecraft:block_transformer",
		"minecraft:villager_food", "minecraft:compostable", "minecraft:cooking_fuel", "minecraft:brewing_fuel",
		"minecraft:mob_visibility", "minecraft:provides_pottery_pattern", "minecraft:sign_text_front",
		"minecraft:sign_text_back", "minecraft:waxed", "minecraft:cushion/color":
		return true
	case "minecraft:damage", "minecraft:repair_cost", "minecraft:ominous_bottle_amplifier", "minecraft:unbreakable",
		"minecraft:max_stack_size", "minecraft:max_damage", "minecraft:additional_trade_cost",
		"minecraft:potion_duration_scale", "minecraft:minimum_attack_charge", "minecraft:enchantment_glint_override", "minecraft:suspicious_stew_effects",
		"minecraft:item_model", "minecraft:break_sound", "minecraft:instrument", "minecraft:jukebox_playable", "minecraft:provides_trim_material", "minecraft:rarity",
		"minecraft:note_block_sound", "minecraft:tooltip_style", "minecraft:glider", "minecraft:creative_slot_lock", "minecraft:intangible_projectile",
		"minecraft:bundle_contents", "minecraft:charged_projectiles", "minecraft:container", "minecraft:sulfur_cube_content", "minecraft:writable_book_content", "minecraft:written_book_content",
		"minecraft:food", "minecraft:consumable", "minecraft:death_protection", "minecraft:weapon", "minecraft:attack_range", "minecraft:blocks_attacks", "minecraft:tool", "minecraft:enchantable", "minecraft:equippable", "minecraft:tooltip_display", "minecraft:use_cooldown", "minecraft:use_remainder", "minecraft:use_effects",
		"minecraft:entity_data", "minecraft:block_entity_data", "minecraft:bucket_entity_data", "minecraft:debug_stick_state", "minecraft:map_decorations", "minecraft:lock", "minecraft:container_loot", "minecraft:recipes", "minecraft:block_state",
		"minecraft:attribute_modifiers", "minecraft:custom_name", "minecraft:item_name", "minecraft:lore", "minecraft:enchantments", "minecraft:stored_enchantments",
		"minecraft:trim", "minecraft:potion_contents", "minecraft:painting/variant", "minecraft:dye", "minecraft:dyed_color", "minecraft:base_color", "minecraft:banner_patterns", "minecraft:pot_decorations", "minecraft:firework_explosion", "minecraft:fireworks", "minecraft:custom_data",
		"minecraft:axolotl/variant", "minecraft:tropical_fish/pattern", "minecraft:tropical_fish/base_color", "minecraft:tropical_fish/pattern_color":
		return true
	default:
		return false
	}
}

// HasWritableBookContent reports whether the effective stack is editable.
func HasWritableBookContent(slot ns.Slot) bool {
	stack, err := FromSlot(slot)
	if err != nil || stack.Components == nil {
		return false
	}
	_, present := stack.Components.rawPassthrough[ComponentWritableBookContent]
	return present
}

// ComponentsToNBT converts a slot's component patch to vanilla item NBT.
func ComponentsToNBT(slot ns.Slot) nbt.Compound {
	result := nbt.Compound{}
	stack, err := FromSlot(slot)
	if err != nil {
		stack = &ItemStack{Components: &Components{}}
	}
	for _, raw := range slot.Components.Add {
		name := registries.DataComponentType.ByID(int32(raw.ID))
		if name == "" {
			continue
		}
		if value := componentPayloadNBT(name, stack.Components); value != nil {
			result[name] = value
			continue
		}
		switch name {
		case "minecraft:axolotl/variant", "minecraft:tropical_fish/pattern",
			"minecraft:tropical_fish/base_color", "minecraft:tropical_fish/pattern_color":
			if enumID, err := ns.NewReader(raw.Data).ReadVarInt(); err == nil {
				if enumName := aquaticVariantComponentNames[name][int32(enumID)]; enumName != "" {
					result[name] = nbt.String(enumName)
				}
			}
		case "minecraft:damage":
			result[name] = nbt.Int(stack.Components.Damage)
		case "minecraft:max_stack_size":
			result[name] = nbt.Int(stack.Components.MaxStackSize)
		case "minecraft:max_damage":
			result[name] = nbt.Int(stack.Components.MaxDamage)
		case "minecraft:additional_trade_cost":
			result[name] = nbt.Int(stack.Components.AdditionalTradeCost)
		case "minecraft:attribute_modifiers":
			result[name] = attributeModifiersNBT(stack.Components.AttributeModifiers)
		case "minecraft:equippable":
			if stack.Components.Equippable != nil {
				result[name] = equippableNBT(stack.Components.Equippable)
			}
		case "minecraft:potion_duration_scale":
			result[name] = nbt.Float(stack.Components.PotionDurationScale)
		case "minecraft:minimum_attack_charge":
			result[name] = nbt.Float(stack.Components.MinimumAttackCharge)
		case "minecraft:item_model":
			result[name] = nbt.String(stack.Components.ItemModel)
		case "minecraft:break_sound":
			if value, err := soundEventHolderNBT(ns.NewReader(raw.Data)); err == nil {
				result[name] = value
			}
		case "minecraft:instrument":
			if value, err := instrumentHolderNBT(ns.NewReader(raw.Data)); err == nil {
				result[name] = value
			}
		case "minecraft:jukebox_playable":
			if value, err := jukeboxSongHolderNBT(ns.NewReader(raw.Data)); err == nil {
				result[name] = value
			}
		case "minecraft:provides_trim_material":
			if value, err := trimMaterialHolderNBT(ns.NewReader(raw.Data)); err == nil {
				result[name] = value
			}
		case "minecraft:rarity":
			result[name] = nbt.String(stack.Components.Rarity)
		case "minecraft:food":
			if food := stack.Components.Food; food != nil {
				result[name] = nbt.Compound{
					"nutrition": nbt.Int(food.Nutrition), "saturation": nbt.Float(food.Saturation),
					"can_always_eat": nbt.Byte(map[bool]int8{true: 1}[food.CanAlwaysEat]),
				}
			}
		case "minecraft:blocks_attacks":
			if component := stack.Components.BlocksAttacks; component != nil {
				result[name] = blocksAttacksNBT(component)
			}
		case "minecraft:consumable":
			if consumable := stack.Components.Consumable; consumable != nil {
				effects := nbt.List{ElementType: nbt.TagCompound}
				for _, effect := range consumable.OnConsumeEffects {
					if effectNBT := consumeEffectNBT(effect); effectNBT != nil {
						effects.Elements = append(effects.Elements, effectNBT)
					}
				}
				result[name] = nbt.Compound{
					"consume_seconds": nbt.Float(consumable.ConsumeSeconds), "animation": nbt.String(consumable.Animation),
					"sound":                 componentSoundNBT(consumable.Sound, consumable.SoundDirect, consumable.SoundRange),
					"has_consume_particles": nbt.Byte(map[bool]int8{true: 1}[consumable.HasParticles]), "on_consume_effects": effects,
				}
			}
		case "minecraft:death_protection":
			if protection := stack.Components.DeathProtection; protection != nil {
				effects := nbt.List{ElementType: nbt.TagCompound}
				for _, effect := range protection.DeathEffects {
					effects.Elements = append(effects.Elements, consumeEffectNBT(effect))
				}
				result[name] = nbt.Compound{"death_effects": effects}
			}
		case "minecraft:weapon":
			if weapon := stack.Components.Weapon; weapon != nil {
				result[name] = nbt.Compound{
					"item_damage_per_attack":       nbt.Int(weapon.ItemDamagePerAttack),
					"disable_blocking_for_seconds": nbt.Float(weapon.DisableBlockingForSeconds),
				}
			}
		case "minecraft:attack_range":
			if attackRange := stack.Components.AttackRange; attackRange != nil {
				result[name] = nbt.Compound{
					"min_reach": nbt.Float(attackRange.MinReach), "max_reach": nbt.Float(attackRange.MaxReach),
					"min_creative_reach": nbt.Float(attackRange.MinCreativeReach), "max_creative_reach": nbt.Float(attackRange.MaxCreativeReach),
					"hitbox_margin": nbt.Float(attackRange.HitboxMargin), "mob_factor": nbt.Float(attackRange.MobFactor),
				}
			}
		case "minecraft:tool":
			if tool := stack.Components.Tool; tool != nil {
				result[name] = toolNBT(tool)
			}
		case "minecraft:enchantable":
			if enchantable := stack.Components.Enchantable; enchantable != nil {
				result[name] = nbt.Compound{"value": nbt.Int(enchantable.Value)}
			}
		case "minecraft:tooltip_display":
			if tooltip := stack.Components.TooltipDisplay; tooltip != nil {
				hidden := nbt.List{ElementType: nbt.TagString}
				for _, componentID := range tooltip.HiddenComponents {
					componentName := registries.DataComponentType.ByID(componentID)
					if componentName != "" {
						hidden.Elements = append(hidden.Elements, nbt.String(componentName))
					}
				}
				result[name] = nbt.Compound{
					"hide_tooltip": nbt.Byte(map[bool]int8{true: 1}[tooltip.HideTooltip]), "hidden_components": hidden,
				}
			}
		case "minecraft:use_effects":
			reader := ns.NewReader(raw.Data)
			canSprint, sprintErr := reader.ReadBool()
			vibrations, vibrationErr := reader.ReadBool()
			speed, speedErr := reader.ReadFloat32()
			if sprintErr == nil && vibrationErr == nil && speedErr == nil {
				result[name] = nbt.Compound{
					"can_sprint":          nbt.Byte(map[bool]int8{true: 1}[bool(canSprint)]),
					"interact_vibrations": nbt.Byte(map[bool]int8{true: 1}[bool(vibrations)]), "speed_multiplier": nbt.Float(speed),
				}
			}
		case "minecraft:use_cooldown":
			if cooldown := stack.Components.UseCooldown; cooldown != nil {
				value := nbt.Compound{"seconds": nbt.Float(cooldown.Seconds)}
				if cooldown.CooldownGroup != "" {
					value["cooldown_group"] = nbt.String(cooldown.CooldownGroup)
				}
				result[name] = value
			}
		case "minecraft:glider", "minecraft:creative_slot_lock", "minecraft:intangible_projectile":
			result[name] = nbt.Compound{}
		case "minecraft:enchantment_glint_override":
			if value, err := ns.NewReader(raw.Data).ReadBool(); err == nil {
				result[name] = nbt.Byte(map[bool]int8{true: 1}[bool(value)])
			}
		case "minecraft:note_block_sound", "minecraft:tooltip_style":
			if value, err := ns.NewReader(raw.Data).ReadString(32767); err == nil {
				result[name] = nbt.String(value)
			}
		case "minecraft:bundle_contents", "minecraft:charged_projectiles", "minecraft:container":
			if value, err := slotListComponentNBT(raw.Data, name == "minecraft:container"); err == nil {
				result[name] = value
			}
		case "minecraft:sulfur_cube_content", "minecraft:use_remainder":
			if slot, err := ReadItemStackTemplate(ns.NewReader(raw.Data)); err == nil {
				result[name] = componentSlotNBT(slot)
			}
		case "minecraft:entity_data":
			if value, err := typedEntityDataNBT(raw.Data, registries.EntityType); err == nil {
				result[name] = value
			}
		case "minecraft:block_entity_data":
			if value, err := typedEntityDataNBT(raw.Data, registries.BlockEntityType); err == nil {
				result[name] = value
			}
		case "minecraft:bucket_entity_data", "minecraft:debug_stick_state", "minecraft:map_decorations", "minecraft:lock", "minecraft:container_loot":
			if value, err := rawNBTComponent(raw.Data); err == nil {
				result[name] = value
			}
		case "minecraft:recipes":
			if value, err := rawNBTComponent(raw.Data); err == nil {
				result[name] = value
			}
		case "minecraft:block_state":
			if value, err := stringMapComponentNBT(raw.Data); err == nil {
				result[name] = value
			}
		case "minecraft:writable_book_content":

			if value, err := writableBookNBT(raw.Data); err == nil {
				result[name] = value
			}
		case "minecraft:written_book_content":
			if value, err := writtenBookNBT(raw.Data); err == nil {
				result[name] = value
			}
		case "minecraft:repair_cost":
			result[name] = nbt.Int(stack.Components.RepairCost)
		case "minecraft:ominous_bottle_amplifier":
			result[name] = nbt.Int(stack.Components.OminousBottleAmplifier)
		case "minecraft:unbreakable":
			result[name] = nbt.Compound{}
		case "minecraft:custom_name":
			if value := textComponentNBT(stack.Components.CustomName); value != nil {
				result[name] = value
			}
		case "minecraft:item_name":
			if value := textComponentNBT(stack.Components.ItemName); value != nil {
				result[name] = value
			}
		case "minecraft:lore":
			if value, err := textComponentListNBT(raw.Data); err == nil {
				result[name] = value
			}
		case "minecraft:enchantments":
			result[name] = enchantmentsNBT(stack.Components.Enchantments)
		case "minecraft:stored_enchantments":
			result[name] = enchantmentsNBT(stack.Components.StoredEnchantments)
		case "minecraft:trim":
			if trim := stack.Components.Trim; trim != nil {
				result[name] = nbt.Compound{"material": nbt.String(trim.Material), "pattern": nbt.String(trim.Pattern)}
			}
		case "minecraft:suspicious_stew_effects":
			if effects, err := ReadSuspiciousStewEffects(slot); err == nil {
				list := nbt.List{ElementType: nbt.TagCompound}
				for _, effect := range effects {
					list.Elements = append(list.Elements, nbt.Compound{"id": nbt.String(effect.Effect), "duration": nbt.Int(effect.Duration)})
				}
				result[name] = list
			}
		case "minecraft:potion_contents":
			if contents, err := ReadPotionContents(slot); err == nil {
				value := nbt.Compound{}
				if contents.Potion != "" {
					value["potion"] = nbt.String(contents.Potion)
				}
				if contents.CustomColor != nil {
					value["custom_color"] = nbt.Int(*contents.CustomColor)
				}
				if len(contents.Effects) > 0 {
					effects := nbt.List{ElementType: nbt.TagCompound}
					for _, potionEffect := range contents.Effects {
						effects.Elements = append(effects.Elements, potionEffectNBT(potionEffect, true))
					}
					value["custom_effects"] = effects
				}
				if contents.CustomName != "" {
					value["custom_name"] = nbt.String(contents.CustomName)
				}
				result[name] = value
			}
		case "minecraft:painting/variant":
			if stack.Components.PaintingVariant != "" {
				result[name] = nbt.String(stack.Components.PaintingVariant)
			}
		case "minecraft:dye":
			if stack.Components.Dye != nil {
				result[name] = nbt.String(dyeNames[*stack.Components.Dye])
			}
		case "minecraft:dyed_color":
			if stack.Components.DyedColor != nil {
				result[name] = nbt.Int(*stack.Components.DyedColor)
			}
		case "minecraft:base_color":
			if stack.Components.BaseColor != nil {
				result[name] = nbt.String(dyeNames[*stack.Components.BaseColor])
			}
		case "minecraft:banner_patterns":
			if stack.Components.BannerPatterns != nil {
				layers := nbt.List{ElementType: nbt.TagCompound}
				for _, layer := range *stack.Components.BannerPatterns {
					var pattern nbt.Tag = nbt.String(layer.Pattern)
					if layer.Pattern == "" {
						pattern = nbt.Compound{"asset_id": nbt.String(layer.AssetID), "translation_key": nbt.String(layer.TranslationKey)}
					}
					layers.Elements = append(layers.Elements, nbt.Compound{
						"pattern": pattern, "color": nbt.String(dyeNames[layer.Color]),
					})
				}
				result[name] = layers
			}
		case "minecraft:firework_explosion":
			if explosion := stack.Components.FireworkExplosion; explosion != nil {
				result[name] = fireworkExplosionNBT(*explosion)
			}
		case "minecraft:fireworks":
			if fireworks := stack.Components.Fireworks; fireworks != nil {
				explosions := nbt.List{ElementType: nbt.TagCompound}
				for _, explosion := range stack.Components.FireworkExplosions {
					explosions.Elements = append(explosions.Elements, fireworkExplosionNBT(explosion))
				}
				result[name] = nbt.Compound{"flight_duration": nbt.Int(fireworks.FlightDuration), "explosions": explosions}
			}
		case "minecraft:custom_data":
			reader := nbt.NewReaderFrom(bytes.NewReader(raw.Data))
			if value, _, err := reader.ReadTag(true); err == nil {
				result[name] = value
			}
		}
	}
	for _, id := range slot.Components.Remove {
		if name := registries.DataComponentType.ByID(int32(id)); name != "" {
			result["!"+name] = nbt.Compound{}
		}
	}
	return result
}

// ApplyNBTComponents applies a vanilla item component patch to a slot.
func ApplyNBTComponents(slot ns.Slot, patch nbt.Compound) (ns.Slot, error) {
	stack, err := FromSlot(slot)
	if err != nil {
		return ns.Slot{}, err
	}
	var rawComponents []ns.RawSlotComponent
	var removals []ns.VarInt
	for key, value := range patch {
		if name, removed := strings.CutPrefix(key, "!"); removed {
			id := registries.DataComponentType.Get(name)
			if id >= 0 {
				removals = append(removals, ns.VarInt(id))
			}
			continue
		}
		id := registries.DataComponentType.Get(key)
		if id < 0 {
			continue
		}
		if handled, err := applyComponentPayloadNBT(stack.Components, key, value); handled {
			if err != nil {
				return ns.Slot{}, err
			}
			stack.Components.SetPresent(id)
			continue
		}
		switch key {
		case "minecraft:axolotl/variant", "minecraft:tropical_fish/pattern",
			"minecraft:tropical_fish/base_color", "minecraft:tropical_fish/pattern_color":
			enumName, ok := value.(nbt.String)
			if !ok {
				return ns.Slot{}, fmt.Errorf("%s component must be a string", key)
			}
			enumID, ok := aquaticVariantComponentID(key, string(enumName))
			if !ok {
				return ns.Slot{}, fmt.Errorf("unknown %s value %q", key, enumName)
			}
			writer := ns.NewWriter()
			writer.WriteVarInt(ns.VarInt(enumID))
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:damage":
			stack.Components.Damage = intValue(value)
			stack.Components.SetPresent(id)
		case "minecraft:max_stack_size":
			stack.Components.MaxStackSize = intValue(value)
			stack.Components.SetPresent(id)
		case "minecraft:max_damage":
			stack.Components.MaxDamage = intValue(value)
			stack.Components.SetPresent(id)
		case "minecraft:additional_trade_cost":
			stack.Components.AdditionalTradeCost = intValue(value)
			stack.Components.SetPresent(id)
		case "minecraft:attribute_modifiers":
			modifiers, err := attributeModifiersFromNBT(value)
			if err != nil {
				return ns.Slot{}, err
			}
			stack.Components.AttributeModifiers = modifiers
			stack.Components.SetPresent(id)
		case "minecraft:equippable":
			equippable, err := equippableFromNBT(value)
			if err != nil {
				return ns.Slot{}, err
			}
			stack.Components.Equippable = equippable
			stack.Components.SetPresent(id)
		case "minecraft:potion_duration_scale":
			stack.Components.PotionDurationScale = floatValue(value)
			stack.Components.SetPresent(id)
		case "minecraft:minimum_attack_charge":
			stack.Components.MinimumAttackCharge = floatValue(value)
			stack.Components.SetPresent(id)
		case "minecraft:food":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("food component must be a compound")
			}
			nutrition := compound.GetInt("nutrition")
			if nutrition < 0 {
				return ns.Slot{}, fmt.Errorf("food nutrition must be non-negative")
			}
			stack.Components.Food = &Food{
				Nutrition: nutrition, Saturation: floatValue(compound["saturation"]),
				CanAlwaysEat: nbtBoolOr(compound, "can_always_eat", false),
			}
			stack.Components.SetPresent(id)
		case "minecraft:blocks_attacks":
			component, err := blocksAttacksFromNBT(value)
			if err != nil {
				return ns.Slot{}, err
			}
			stack.Components.BlocksAttacks = component
			stack.Components.SetPresent(id)
		case "minecraft:consumable":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("consumable component must be a compound")
			}
			seconds := 1.6
			if _, present := compound["consume_seconds"]; present {
				seconds = floatValue(compound["consume_seconds"])
			}
			if seconds < 0 {
				return ns.Slot{}, fmt.Errorf("consume_seconds must be non-negative")
			}
			animation := compound.GetString("animation")
			if animation == "" {
				animation = "eat"
			}
			if !slices.Contains(consumeAnimations, animation) {
				return ns.Slot{}, fmt.Errorf("unknown consume animation %q", animation)
			}
			soundName := "minecraft:entity.generic.eat"
			soundDirect := false
			var soundRange *float64
			if soundTag, present := compound["sound"]; present {
				parsedSound, err := consumableSoundFromNBT(soundTag)
				if err != nil {
					return ns.Slot{}, err
				}
				soundName, soundDirect, soundRange = parsedSound.name, parsedSound.direct, parsedSound.fixedRange
			}
			effectList := compound.GetList("on_consume_effects")
			effects := make([]ConsumeEffect, 0, len(effectList.Elements))
			for _, element := range effectList.Elements {
				effect, err := consumeEffectFromNBT(element)
				if err != nil {
					return ns.Slot{}, err
				}
				effects = append(effects, effect)
			}
			stack.Components.Consumable = &Consumable{
				ConsumeSeconds: seconds, Animation: animation, Sound: soundName, SoundDirect: soundDirect, SoundRange: soundRange,
				HasParticles: nbtBoolOr(compound, "has_consume_particles", true), OnConsumeEffects: effects,
			}
			stack.Components.SetPresent(id)
		case "minecraft:death_protection":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("death_protection component must be a compound")
			}
			effectList := compound.GetList("death_effects")
			effects := make([]ConsumeEffect, 0, len(effectList.Elements))
			for _, element := range effectList.Elements {
				effect, err := consumeEffectFromNBT(element)
				if err != nil {
					return ns.Slot{}, err
				}
				effects = append(effects, effect)
			}
			stack.Components.DeathProtection = &DeathProtection{DeathEffects: effects}
			stack.Components.SetPresent(id)
		case "minecraft:weapon":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("weapon component must be a compound")
			}
			damage := int32(1)
			if _, present := compound["item_damage_per_attack"]; present {
				damage = compound.GetInt("item_damage_per_attack")
			}
			disable := floatValue(compound["disable_blocking_for_seconds"])
			if damage < 0 || disable < 0 {
				return ns.Slot{}, fmt.Errorf("weapon values must be non-negative")
			}
			stack.Components.Weapon = &Weapon{ItemDamagePerAttack: damage, DisableBlockingForSeconds: disable}
			stack.Components.SetPresent(id)
		case "minecraft:attack_range":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("attack_range component must be a compound")
			}
			attackRange := &AttackRange{
				MinReach: 0, MaxReach: 3, MinCreativeReach: 0,
				MaxCreativeReach: 5, HitboxMargin: 0.3, MobFactor: 1,
			}
			if tag, present := compound["min_reach"]; present {
				attackRange.MinReach = floatValue(tag)
			}
			if tag, present := compound["max_reach"]; present {
				attackRange.MaxReach = floatValue(tag)
			}
			if tag, present := compound["min_creative_reach"]; present {
				attackRange.MinCreativeReach = floatValue(tag)
			}
			if tag, present := compound["max_creative_reach"]; present {
				attackRange.MaxCreativeReach = floatValue(tag)
			}
			if tag, present := compound["hitbox_margin"]; present {
				attackRange.HitboxMargin = floatValue(tag)
			}
			if tag, present := compound["mob_factor"]; present {
				attackRange.MobFactor = floatValue(tag)
			}
			if attackRange.MinReach < 0 || attackRange.MinReach > 64 ||
				attackRange.MaxReach < 0 || attackRange.MaxReach > 64 ||
				attackRange.MinCreativeReach < 0 || attackRange.MinCreativeReach > 64 ||
				attackRange.MaxCreativeReach < 0 || attackRange.MaxCreativeReach > 64 ||
				attackRange.HitboxMargin < 0 || attackRange.HitboxMargin > 1 ||
				attackRange.MobFactor < 0 || attackRange.MobFactor > 2 {
				return ns.Slot{}, fmt.Errorf("attack_range values are outside vanilla bounds")
			}
			stack.Components.AttackRange = attackRange
			stack.Components.SetPresent(id)
		case "minecraft:tool":
			tool, err := toolFromNBT(value)
			if err != nil {
				return ns.Slot{}, err
			}
			stack.Components.Tool = tool
			stack.Components.SetPresent(id)
		case "minecraft:enchantable":
			compound, ok := value.(nbt.Compound)
			if !ok || compound.GetInt("value") <= 0 {
				return ns.Slot{}, fmt.Errorf("enchantable value must be positive")
			}
			stack.Components.Enchantable = &Enchantable{Value: compound.GetInt("value")}
			stack.Components.SetPresent(id)
		case "minecraft:tooltip_display":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("tooltip_display component must be a compound")
			}
			hiddenList := compound.GetList("hidden_components")
			hidden := make([]int32, 0, len(hiddenList.Elements))
			seen := make(map[int32]bool)
			for _, element := range hiddenList.Elements {
				name, ok := element.(nbt.String)
				componentID := registries.DataComponentType.Get(string(name))
				if !ok || componentID < 0 {
					return ns.Slot{}, fmt.Errorf("unknown hidden component %q", name)
				}
				if !seen[componentID] {
					hidden = append(hidden, componentID)
					seen[componentID] = true
				}
			}
			stack.Components.TooltipDisplay = &TooltipDisplay{HideTooltip: nbtBoolOr(compound, "hide_tooltip", false), HiddenComponents: hidden}
			stack.Components.SetPresent(id)
		case "minecraft:use_effects":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("use_effects component must be a compound")
			}
			speed := 0.2
			if speedTag, present := compound["speed_multiplier"]; present {
				speed = floatValue(speedTag)
			}
			if speed < 0 || speed > 1 {
				return ns.Slot{}, fmt.Errorf("use_effects speed_multiplier must be between 0 and 1")
			}
			writer := ns.NewWriter()
			writer.WriteBool(ns.Boolean(nbtBoolOr(compound, "can_sprint", false)))
			writer.WriteBool(ns.Boolean(nbtBoolOr(compound, "interact_vibrations", true)))
			writer.WriteFloat32(ns.Float32(speed))
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:use_cooldown":
			compound, ok := value.(nbt.Compound)
			seconds := floatValue(compound["seconds"])
			if !ok || seconds <= 0 {
				return ns.Slot{}, fmt.Errorf("use_cooldown seconds must be positive")
			}
			stack.Components.UseCooldown = &UseCooldown{Seconds: seconds, CooldownGroup: compound.GetString("cooldown_group")}
			stack.Components.SetPresent(id)
		case "minecraft:item_model", "minecraft:rarity":
			stringValue, ok := value.(nbt.String)
			if !ok {
				return ns.Slot{}, fmt.Errorf("%s component must be a string", key)
			}
			if key == "minecraft:item_model" {
				stack.Components.ItemModel = string(stringValue)
			} else {
				stack.Components.Rarity = string(stringValue)
			}
			stack.Components.SetPresent(id)
		case "minecraft:break_sound", "minecraft:instrument", "minecraft:jukebox_playable", "minecraft:provides_trim_material":
			if stringValue, ok := value.(nbt.String); ok {
				switch key {
				case "minecraft:break_sound":
					stack.Components.BreakSound = string(stringValue)
				case "minecraft:instrument":
					stack.Components.Instrument = string(stringValue)
				case "minecraft:jukebox_playable":
					stack.Components.JukeboxPlayable = string(stringValue)
				case "minecraft:provides_trim_material":
					stack.Components.ProvidesTrimMaterial = string(stringValue)
				}
				stack.Components.SetPresent(id)
				break
			}
			writer := ns.NewWriter()
			var err error
			switch key {
			case "minecraft:break_sound":
				err = writeSoundEventHolderNBT(writer, value)
			case "minecraft:instrument":
				err = writeInstrumentHolderNBT(writer, value)
			case "minecraft:jukebox_playable":
				err = writeJukeboxSongHolderNBT(writer, value)
			case "minecraft:provides_trim_material":
				err = writeTrimMaterialHolderNBT(writer, value)
			}
			if err != nil {
				return ns.Slot{}, fmt.Errorf("%s: %w", key, err)
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:glider":
			stack.Components.Glider = true
			stack.Components.SetPresent(id)
		case "minecraft:enchantment_glint_override":
			writer := ns.NewWriter()
			writer.WriteBool(ns.Boolean(intValue(value) != 0))
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:note_block_sound", "minecraft:tooltip_style":
			stringValue, ok := value.(nbt.String)
			if !ok {
				return ns.Slot{}, fmt.Errorf("%s component must be a string", key)
			}
			writer := ns.NewWriter()
			writer.WriteString(ns.String(stringValue))
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:creative_slot_lock":
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id)})
		case "minecraft:intangible_projectile":
			var encoded bytes.Buffer
			if err := nbt.NewWriterTo(&encoded).WriteTag(nbt.Compound{}, "", true); err != nil {
				return ns.Slot{}, err
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: encoded.Bytes()})
		case "minecraft:bundle_contents", "minecraft:charged_projectiles", "minecraft:container":
			list, ok := value.(nbt.List)
			if !ok || list.ElementType != nbt.TagCompound || len(list.Elements) > 256 {
				return ns.Slot{}, fmt.Errorf("%s component must be a compound list of at most 256 items", key)
			}
			writer := ns.NewWriter()
			if key == "minecraft:container" {
				slots := make(map[int]ns.Slot, len(list.Elements))
				maxSlot := -1
				for _, element := range list.Elements {
					entry, ok := element.(nbt.Compound)
					if !ok {
						return ns.Slot{}, fmt.Errorf("container entry must be a compound")
					}
					slotIndex := int(entry.GetInt("slot"))
					if slotIndex < 0 || slotIndex > 255 {
						return ns.Slot{}, fmt.Errorf("container slot %d is invalid", slotIndex)
					}
					item, err := slotFromComponentNBT(entry["item"])
					if err != nil {
						return ns.Slot{}, err
					}
					slots[slotIndex] = item
					maxSlot = max(maxSlot, slotIndex)
				}
				writer.WriteVarInt(ns.VarInt(maxSlot + 1))
				for slotIndex := range maxSlot + 1 {
					item, present := slots[slotIndex]
					writer.WriteBool(ns.Boolean(present))
					if present {
						if err := WriteItemStackTemplate(writer, item); err != nil {
							return ns.Slot{}, err
						}
					}
				}
			} else {
				writer.WriteVarInt(ns.VarInt(len(list.Elements)))
				for _, element := range list.Elements {
					item, err := slotFromComponentNBT(element)
					if err != nil {
						return ns.Slot{}, err
					}
					if err := WriteItemStackTemplate(writer, item); err != nil {
						return ns.Slot{}, err
					}
				}
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:sulfur_cube_content", "minecraft:use_remainder":
			item, err := slotFromComponentNBT(value)
			if err != nil {
				return ns.Slot{}, err
			}
			writer := ns.NewWriter()
			if err := WriteItemStackTemplate(writer, item); err != nil {
				return ns.Slot{}, err
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:writable_book_content":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("writable_book_content component must be a compound")
			}
			pages := compound.GetList("pages")
			if len(pages.Elements) > 100 {
				return ns.Slot{}, fmt.Errorf("writable book has too many pages")
			}
			writer := ns.NewWriter()
			writer.WriteVarInt(ns.VarInt(len(pages.Elements)))
			for _, page := range pages.Elements {
				raw, filtered, hasFiltered, err := filterableStringNBT(page, 1024)
				if err != nil {
					return ns.Slot{}, err
				}
				writer.WriteString(ns.String(raw))
				writer.WriteBool(ns.Boolean(hasFiltered))
				if hasFiltered {
					writer.WriteString(ns.String(filtered))
				}
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:written_book_content":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("written_book_content component must be a compound")
			}
			titleRaw, titleFiltered, titleHasFiltered, err := filterableStringNBT(compound["title"], 32)
			if err != nil {
				return ns.Slot{}, err
			}
			generation := compound.GetInt("generation")
			if generation < 0 || generation > 3 {
				return ns.Slot{}, fmt.Errorf("written book generation %d is invalid", generation)
			}
			writer := ns.NewWriter()
			writer.WriteString(ns.String(titleRaw))
			writer.WriteBool(ns.Boolean(titleHasFiltered))
			if titleHasFiltered {
				writer.WriteString(ns.String(titleFiltered))
			}
			writer.WriteString(ns.String(compound.GetString("author")))
			writer.WriteVarInt(ns.VarInt(generation))
			pages := compound.GetList("pages")
			if len(pages.Elements) > 100 {
				return ns.Slot{}, fmt.Errorf("written book has too many pages")
			}
			writer.WriteVarInt(ns.VarInt(len(pages.Elements)))
			for _, page := range pages.Elements {
				raw, filtered, hasFiltered := filterableTagNBT(page)
				if raw == nil {
					return ns.Slot{}, fmt.Errorf("written book page is invalid")
				}
				if err := nbt.NewWriterTo(writer.Writer()).WriteTag(raw, "", true); err != nil {
					return ns.Slot{}, err
				}
				writer.WriteBool(ns.Boolean(hasFiltered))
				if hasFiltered {
					if err := nbt.NewWriterTo(writer.Writer()).WriteTag(filtered, "", true); err != nil {
						return ns.Slot{}, err
					}
				}
			}
			writer.WriteBool(ns.Boolean(nbtBoolOr(compound, "resolved", false)))
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:repair_cost":
			stack.Components.RepairCost = intValue(value)
			stack.Components.SetPresent(id)
		case "minecraft:ominous_bottle_amplifier":
			stack.Components.OminousBottleAmplifier = intValue(value)
			stack.Components.SetPresent(id)
		case "minecraft:unbreakable":
			stack.Components.Unbreakable = true
			stack.Components.SetPresent(id)
		case "minecraft:custom_name":
			stack.Components.CustomName = textComponentFromNBT(value)
			stack.Components.SetPresent(id)
		case "minecraft:item_name":
			stack.Components.ItemName = textComponentFromNBT(value)
			stack.Components.SetPresent(id)
		case "minecraft:lore":
			list, ok := value.(nbt.List)
			if !ok || len(list.Elements) > 256 {
				return ns.Slot{}, fmt.Errorf("lore component must be a list of at most 256 components")
			}
			writer := ns.NewWriter()
			writer.WriteVarInt(ns.VarInt(len(list.Elements)))
			for _, line := range list.Elements {
				if err := nbt.NewWriterTo(writer.Writer()).WriteTag(line, "", true); err != nil {
					return ns.Slot{}, err
				}
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:enchantments":
			stack.Components.Enchantments = enchantmentsFromNBT(value)
			stack.Components.SetPresent(id)
		case "minecraft:stored_enchantments":
			stack.Components.StoredEnchantments = enchantmentsFromNBT(value)
			stack.Components.SetPresent(id)
		case "minecraft:trim":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("trim component must be a compound")
			}
			trim := &ArmorTrim{Material: compound.GetString("material"), Pattern: compound.GetString("pattern")}
			if trim.Material == "" || trim.Pattern == "" {
				return ns.Slot{}, fmt.Errorf("trim component requires material and pattern")
			}
			stack.Components.Trim = trim
			stack.Components.SetPresent(id)
		case "minecraft:dye":
			name, ok := value.(nbt.String)
			dye := dyeID(string(name))
			if !ok || dye < 0 {
				return ns.Slot{}, fmt.Errorf("unknown dye %q", name)
			}
			stack.Components.Dye = &dye
			stack.Components.SetPresent(id)
		case "minecraft:base_color":
			name, ok := value.(nbt.String)
			color := dyeID(string(name))
			if !ok || color < 0 {
				return ns.Slot{}, fmt.Errorf("unknown base color %q", name)
			}
			stack.Components.BaseColor = &color
			stack.Components.SetPresent(id)
		case "minecraft:dyed_color":
			color := intValue(value)
			if color < 0 || color > 0xffffff {
				return ns.Slot{}, fmt.Errorf("dyed_color %d is outside RGB range", color)
			}
			stack.Components.DyedColor = &color
			stack.Components.SetPresent(id)
		case "minecraft:firework_explosion":
			explosion, err := fireworkExplosionFromNBT(value)
			if err != nil {
				return ns.Slot{}, err
			}
			stack.Components.FireworkExplosion = &explosion
			stack.Components.SetPresent(id)
		case "minecraft:fireworks":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("fireworks component must be a compound")
			}
			flight := intValue(compound["flight_duration"])
			if flight < 0 || flight > 255 {
				return ns.Slot{}, fmt.Errorf("firework flight duration %d is invalid", flight)
			}
			list := compound.GetList("explosions")
			if len(list.Elements) > 256 {
				return ns.Slot{}, fmt.Errorf("too many firework explosions")
			}
			explosions := make([]FireworkExplosion, 0, len(list.Elements))
			for _, element := range list.Elements {
				explosion, err := fireworkExplosionFromNBT(element)
				if err != nil {
					return ns.Slot{}, err
				}
				explosions = append(explosions, explosion)
			}
			stack.Components.Fireworks = &Fireworks{FlightDuration: flight}
			stack.Components.FireworkExplosions = explosions
			stack.Components.SetPresent(id)
		case "minecraft:banner_patterns":
			list, ok := value.(nbt.List)
			if !ok || list.ElementType != nbt.TagCompound || len(list.Elements) > 256 {
				return ns.Slot{}, fmt.Errorf("banner_patterns component must be a compound list of at most 256 layers")
			}
			layers := make([]BannerPatternLayer, 0, len(list.Elements))
			for _, element := range list.Elements {
				layer, ok := element.(nbt.Compound)
				if !ok {
					return ns.Slot{}, fmt.Errorf("banner pattern layer must be a compound")
				}
				pattern := layer.GetString("pattern")
				patternData := BannerPatternLayer{Pattern: pattern}
				if direct, ok := layer["pattern"].(nbt.Compound); ok {
					patternData.Pattern = ""
					patternData.AssetID = direct.GetString("asset_id")
					patternData.TranslationKey = direct.GetString("translation_key")
				}
				color := dyeID(layer.GetString("color"))
				if color < 0 || patternData.Pattern != "" && registryEntryID("minecraft:banner_pattern", patternData.Pattern) < 0 || patternData.Pattern == "" && (patternData.AssetID == "" || patternData.TranslationKey == "") {
					return ns.Slot{}, fmt.Errorf("invalid banner pattern %q color %q", pattern, layer.GetString("color"))
				}
				patternData.Color = color
				layers = append(layers, patternData)
			}
			stack.Components.BannerPatterns = &layers
			stack.Components.SetPresent(id)
		case "minecraft:painting/variant":
			variant, ok := value.(nbt.String)
			if !ok || registryEntryID("minecraft:painting_variant", string(variant)) < 0 {
				return ns.Slot{}, fmt.Errorf("unknown painting variant %q", variant)
			}
			stack.Components.PaintingVariant = string(variant)
			stack.Components.SetPresent(id)
		case "minecraft:suspicious_stew_effects":
			list, ok := value.(nbt.List)
			if !ok || list.ElementType != nbt.TagCompound || len(list.Elements) > 256 {
				return ns.Slot{}, fmt.Errorf("suspicious_stew_effects must be a compound list")
			}
			writer := ns.NewWriter()
			writer.WriteVarInt(ns.VarInt(len(list.Elements)))
			for _, element := range list.Elements {
				entry, ok := element.(nbt.Compound)
				if !ok {
					return ns.Slot{}, fmt.Errorf("suspicious stew effect must be a compound")
				}
				effectID := registries.MobEffect.Get(entry.GetString("id"))
				if effectID < 0 {
					return ns.Slot{}, fmt.Errorf("unknown suspicious stew effect %q", entry.GetString("id"))
				}
				duration := int32(160)
				if _, present := entry["duration"]; present {
					duration = entry.GetInt("duration")
				}
				writer.WriteVarInt(ns.VarInt(effectID))
				writer.WriteVarInt(ns.VarInt(duration))
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:potion_contents":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("potion_contents component must be a compound")
			}
			potion := compound.GetString("potion")
			if potion != "" && registries.Potion.Get(potion) < 0 {
				return ns.Slot{}, fmt.Errorf("unknown potion %q", potion)
			}
			writer := ns.NewWriter()
			writer.WriteBool(ns.Boolean(potion != ""))
			if potion != "" {
				writer.WriteVarInt(ns.VarInt(registries.Potion.Get(potion)))
			}
			color, hasColor := compound["custom_color"]
			writer.WriteBool(ns.Boolean(hasColor))
			if hasColor {
				writer.WriteInt32(ns.Int32(intValue(color)))
			}
			effects := compound.GetList("custom_effects")
			if len(effects.Elements) > 256 {
				return ns.Slot{}, fmt.Errorf("too many custom potion effects")
			}
			writer.WriteVarInt(ns.VarInt(len(effects.Elements)))
			for _, element := range effects.Elements {
				potionEffect, ok := element.(nbt.Compound)
				if !ok {
					return ns.Slot{}, fmt.Errorf("custom potion effect must be a compound")
				}
				effectName := potionEffect.GetString("id")
				effectID := registries.MobEffect.Get(effectName)
				if effectID < 0 {
					return ns.Slot{}, fmt.Errorf("unknown potion effect %q", effectName)
				}
				writer.WriteVarInt(ns.VarInt(effectID))
				if err := writePotionEffectNBT(writer, potionEffect); err != nil {
					return ns.Slot{}, err
				}
			}
			customName := compound.GetString("custom_name")
			writer.WriteBool(ns.Boolean(customName != ""))
			if customName != "" {
				writer.WriteString(ns.String(customName))
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:entity_data", "minecraft:block_entity_data":
			compound, ok := value.(nbt.Compound)
			if !ok {
				return ns.Slot{}, fmt.Errorf("%s component must be a compound", key)
			}
			registry := registries.EntityType
			if key == "minecraft:block_entity_data" {
				registry = registries.BlockEntityType
			}
			typeName := compound.GetString("id")
			typeID := registry.Get(typeName)
			if typeID < 0 {
				return ns.Slot{}, fmt.Errorf("unknown %s type %q", key, typeName)
			}
			payload := make(nbt.Compound, len(compound)-1)
			for field, fieldValue := range compound {
				if field != "id" {
					payload[field] = fieldValue
				}
			}
			writer := ns.NewWriter()
			writer.WriteVarInt(ns.VarInt(typeID))
			if err := nbt.NewWriterTo(writer.Writer()).WriteTag(payload, "", true); err != nil {
				return ns.Slot{}, err
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:recipes":
			list, ok := value.(nbt.List)
			if !ok || list.ElementType != nbt.TagString || len(list.Elements) > 256 {
				return ns.Slot{}, fmt.Errorf("recipes component must be a string list")
			}
			var encoded bytes.Buffer
			if err := nbt.NewWriterTo(&encoded).WriteTag(list, "", true); err != nil {
				return ns.Slot{}, err
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: encoded.Bytes()})
		case "minecraft:block_state":
			compound, ok := value.(nbt.Compound)
			if !ok || len(compound) > 256 {
				return ns.Slot{}, fmt.Errorf("block_state component must be a compound")
			}
			keys := make([]string, 0, len(compound))
			for property := range compound {
				keys = append(keys, property)
			}
			sort.Strings(keys)
			writer := ns.NewWriter()
			writer.WriteVarInt(ns.VarInt(len(keys)))
			for _, property := range keys {
				propertyValue, ok := compound[property].(nbt.String)
				if !ok {
					return ns.Slot{}, fmt.Errorf("block state property %q must be a string", property)
				}
				writer.WriteString(ns.String(property))
				writer.WriteString(ns.String(propertyValue))
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: writer.Bytes()})
		case "minecraft:custom_data", "minecraft:bucket_entity_data", "minecraft:debug_stick_state", "minecraft:map_decorations", "minecraft:lock", "minecraft:container_loot":
			if _, ok := value.(nbt.Compound); !ok {
				return ns.Slot{}, fmt.Errorf("%s component must be a compound", key)
			}
			var encoded bytes.Buffer
			if err := nbt.NewWriterTo(&encoded).WriteTag(value, "", true); err != nil {
				return ns.Slot{}, err
			}
			rawComponents = append(rawComponents, ns.RawSlotComponent{ID: ns.VarInt(id), Data: encoded.Bytes()})
		}
	}
	encoded, err := stack.ToSlot()
	if err != nil {
		return ns.Slot{}, fmt.Errorf("encode item components: %w", err)
	}
	for _, raw := range rawComponents {
		encoded.Components.Add = replaceRawComponent(encoded.Components.Add, raw)
	}
	for _, id := range removals {
		encoded.Components.Add = removeRawComponent(encoded.Components.Add, id)
		encoded.Components.Remove = append(encoded.Components.Remove, id)
	}
	return encoded, nil
}

func equippableNBT(component *Equippable) nbt.Tag {
	result := nbt.Compound{"slot": nbt.String(component.Slot)}
	if component.EquipSound != "" && component.EquipSound != "minecraft:item.armor.equip_generic" ||
		component.EquipSoundDirect || component.EquipSoundRange != nil {
		result["equip_sound"] = componentSoundNBT(component.EquipSound, component.EquipSoundDirect, component.EquipSoundRange)
	}
	if component.AssetID != "" {
		result["asset_id"] = nbt.String(component.AssetID)
	}
	if component.CameraOverlay != "" {
		result["camera_overlay"] = nbt.String(component.CameraOverlay)
	}
	if component.AllowedEntities != nil {
		result["allowed_entities"] = holderSetNBT(component.AllowedEntities)
	}
	if !component.Dispensable {
		result["dispensable"] = nbt.Byte(0)
	}
	if !component.Swappable {
		result["swappable"] = nbt.Byte(0)
	}
	if !component.DamageOnHurt {
		result["damage_on_hurt"] = nbt.Byte(0)
	}
	if component.EquipOnInteract {
		result["equip_on_interact"] = nbt.Byte(1)
	}
	if component.CanBeSheared {
		result["can_be_sheared"] = nbt.Byte(1)
	}
	if component.ShearingSound != "" && component.ShearingSound != "minecraft:item.shears.snip" ||
		component.ShearingSoundDirect || component.ShearingSoundRange != nil {
		result["shearing_sound"] = componentSoundNBT(component.ShearingSound, component.ShearingSoundDirect, component.ShearingSoundRange)
	}
	return result
}

func equippableFromNBT(value nbt.Tag) (*Equippable, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("equippable component must be a compound")
	}
	slot := compound.GetString("slot")
	if !slices.Contains(equipmentSlots, slot) {
		return nil, fmt.Errorf("unknown equipment slot %q", slot)
	}
	component := &Equippable{
		Slot: slot, EquipSound: "minecraft:item.armor.equip_generic",
		Dispensable: true, Swappable: true, DamageOnHurt: true,
		ShearingSound: "minecraft:item.shears.snip",
	}
	if soundTag, present := compound["equip_sound"]; present {
		sound, err := consumableSoundFromNBT(soundTag)
		if err != nil {
			return nil, err
		}
		component.EquipSound, component.EquipSoundDirect, component.EquipSoundRange = sound.name, sound.direct, sound.fixedRange
	}
	component.AssetID = compound.GetString("asset_id")
	component.CameraOverlay = compound.GetString("camera_overlay")
	if allowed, present := compound["allowed_entities"]; present {
		var err error
		component.AllowedEntities, err = holderSetFromNBT(allowed)
		if err != nil {
			return nil, err
		}
	}
	component.Dispensable = nbtBoolOr(compound, "dispensable", true)
	component.Swappable = nbtBoolOr(compound, "swappable", true)
	component.DamageOnHurt = nbtBoolOr(compound, "damage_on_hurt", true)
	component.EquipOnInteract = nbtBoolOr(compound, "equip_on_interact", false)
	component.CanBeSheared = nbtBoolOr(compound, "can_be_sheared", false)
	if soundTag, present := compound["shearing_sound"]; present {
		sound, err := consumableSoundFromNBT(soundTag)
		if err != nil {
			return nil, err
		}
		component.ShearingSound, component.ShearingSoundDirect, component.ShearingSoundRange = sound.name, sound.direct, sound.fixedRange
	}
	return component, nil
}

func blocksAttacksNBT(component *BlocksAttacks) nbt.Tag {
	reductions := nbt.List{ElementType: nbt.TagCompound}
	for _, reduction := range component.DamageReductions {
		entry := nbt.Compound{
			"horizontal_blocking_angle": nbt.Float(reduction.HorizontalBlockingAngle),
			"base":                      nbt.Float(reduction.Base),
			"factor":                    nbt.Float(reduction.Factor),
		}
		if reduction.Types != nil {
			entry["type"] = holderSetNBT(reduction.Types)
		}
		reductions.Elements = append(reductions.Elements, entry)
	}
	result := nbt.Compound{
		"block_delay_seconds":    nbt.Float(component.BlockDelaySeconds),
		"disable_cooldown_scale": nbt.Float(component.DisableCooldownScale),
		"damage_reductions":      reductions,
		"item_damage": nbt.Compound{
			"threshold": nbt.Float(component.ItemDamage.Threshold),
			"base":      nbt.Float(component.ItemDamage.Base),
			"factor":    nbt.Float(component.ItemDamage.Factor),
		},
	}
	if component.BypassedBy != nil {
		result["bypassed_by"] = holderSetNBT(component.BypassedBy)
	}
	if component.BlockSound != "" {
		result["block_sound"] = componentSoundNBT(component.BlockSound, component.BlockSoundDirect, component.BlockSoundRange)
	}
	if component.DisabledSound != "" {
		result["disabled_sound"] = componentSoundNBT(component.DisabledSound, component.DisabledSoundDirect, component.DisabledSoundRange)
	}
	return result
}

func blocksAttacksFromNBT(value nbt.Tag) (*BlocksAttacks, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("blocks_attacks component must be a compound")
	}
	component := &BlocksAttacks{
		DisableCooldownScale: 1,
		DamageReductions:     []DamageReduction{{HorizontalBlockingAngle: 90, Factor: 1}},
		ItemDamage:           DamageSpec{Threshold: 1, Factor: 1},
	}
	if _, present := compound["block_delay_seconds"]; present {
		component.BlockDelaySeconds = floatValue(compound["block_delay_seconds"])
	}
	if _, present := compound["disable_cooldown_scale"]; present {
		component.DisableCooldownScale = floatValue(compound["disable_cooldown_scale"])
	}
	if component.BlockDelaySeconds < 0 || component.DisableCooldownScale < 0 {
		return nil, fmt.Errorf("blocks_attacks timing values must be non-negative")
	}
	if raw, present := compound["damage_reductions"]; present {
		list, ok := raw.(nbt.List)
		if !ok || len(list.Elements) > 0 && list.ElementType != nbt.TagCompound {
			return nil, fmt.Errorf("blocks_attacks damage_reductions must be a compound list")
		}
		component.DamageReductions = make([]DamageReduction, 0, len(list.Elements))
		for _, element := range list.Elements {
			entry, ok := element.(nbt.Compound)
			if !ok {
				return nil, fmt.Errorf("blocks_attacks damage reduction must be a compound")
			}
			angle := 90.0
			if _, present := entry["horizontal_blocking_angle"]; present {
				angle = floatValue(entry["horizontal_blocking_angle"])
			}
			if angle <= 0 {
				return nil, fmt.Errorf("horizontal blocking angle must be positive")
			}
			var types []string
			if typeTag, present := entry["type"]; present {
				var err error
				types, err = holderSetFromNBT(typeTag)
				if err != nil {
					return nil, err
				}
			}
			baseTag, hasBase := entry["base"]
			factorTag, hasFactor := entry["factor"]
			if !hasBase || !hasFactor {
				return nil, fmt.Errorf("blocks_attacks damage reduction requires base and factor")
			}
			base, factor := floatValue(baseTag), floatValue(factorTag)
			component.DamageReductions = append(component.DamageReductions, DamageReduction{
				HorizontalBlockingAngle: angle, Types: types, Base: base, Factor: factor,
			})
		}
	}
	if raw, present := compound["item_damage"]; present {
		damage, ok := raw.(nbt.Compound)
		if !ok {
			return nil, fmt.Errorf("blocks_attacks item_damage must be a compound")
		}
		thresholdTag, hasThreshold := damage["threshold"]
		baseTag, hasBase := damage["base"]
		factorTag, hasFactor := damage["factor"]
		if !hasThreshold || !hasBase || !hasFactor {
			return nil, fmt.Errorf("blocks_attacks item_damage requires threshold, base, and factor")
		}
		component.ItemDamage = DamageSpec{
			Threshold: floatValue(thresholdTag), Base: floatValue(baseTag), Factor: floatValue(factorTag),
		}
		if component.ItemDamage.Threshold < 0 {
			return nil, fmt.Errorf("item damage threshold must be non-negative")
		}
	}
	if bypassed, present := compound["bypassed_by"]; present {
		var err error
		component.BypassedBy, err = holderSetFromNBT(bypassed)
		if err != nil {
			return nil, err
		}
	}
	if soundTag, present := compound["block_sound"]; present {
		sound, err := consumableSoundFromNBT(soundTag)
		if err != nil {
			return nil, err
		}
		component.BlockSound, component.BlockSoundDirect, component.BlockSoundRange = sound.name, sound.direct, sound.fixedRange
	}
	if soundTag, present := compound["disabled_sound"]; present {
		sound, err := consumableSoundFromNBT(soundTag)
		if err != nil {
			return nil, err
		}
		component.DisabledSound, component.DisabledSoundDirect, component.DisabledSoundRange = sound.name, sound.direct, sound.fixedRange
	}
	return component, nil
}

func toolNBT(tool *Tool) nbt.Tag {
	rules := nbt.List{ElementType: nbt.TagCompound}
	for _, rule := range tool.Rules {
		entry := nbt.Compound{"blocks": holderSetNBT(rule.Blocks)}
		if rule.HasSpeed {
			entry["speed"] = nbt.Float(rule.Speed)
		}
		if rule.HasCorrectForDrops {
			entry["correct_for_drops"] = nbt.Byte(map[bool]int8{true: 1}[rule.CorrectForDrops])
		}
		rules.Elements = append(rules.Elements, entry)
	}
	result := nbt.Compound{"rules": rules}
	if tool.DefaultMiningSpeed != 1 {
		result["default_mining_speed"] = nbt.Float(tool.DefaultMiningSpeed)
	}
	if tool.DamagePerBlock != 1 {
		result["damage_per_block"] = nbt.Int(tool.DamagePerBlock)
	}
	if !tool.CanDestroyBlocksInCreative {
		result["can_destroy_blocks_in_creative"] = nbt.Byte(0)
	}
	return result
}

func toolFromNBT(value nbt.Tag) (*Tool, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("tool component must be a compound")
	}
	ruleList, ok := compound["rules"].(nbt.List)
	if !ok || len(ruleList.Elements) > 0 && ruleList.ElementType != nbt.TagCompound {
		return nil, fmt.Errorf("tool rules must be a compound list")
	}
	tool := &Tool{DefaultMiningSpeed: 1, DamagePerBlock: 1, CanDestroyBlocksInCreative: true}
	if tag, present := compound["default_mining_speed"]; present {
		tool.DefaultMiningSpeed = floatValue(tag)
	}
	if tag, present := compound["damage_per_block"]; present {
		tool.DamagePerBlock = intValue(tag)
	}
	if _, present := compound["can_destroy_blocks_in_creative"]; present {
		tool.CanDestroyBlocksInCreative = nbtBoolOr(compound, "can_destroy_blocks_in_creative", true)
	}
	if tool.DamagePerBlock < 0 {
		return nil, fmt.Errorf("tool damage_per_block must be non-negative")
	}
	tool.Rules = make([]ToolRule, 0, len(ruleList.Elements))
	for _, element := range ruleList.Elements {
		entry, ok := element.(nbt.Compound)
		if !ok {
			return nil, fmt.Errorf("tool rule must be a compound")
		}
		blocksTag, present := entry["blocks"]
		if !present {
			return nil, fmt.Errorf("tool rule is missing blocks")
		}
		rule := ToolRule{}
		var err error
		rule.Blocks, err = holderSetFromNBT(blocksTag)
		if err != nil {
			return nil, err
		}
		if len(rule.Blocks) == 0 {
			return nil, fmt.Errorf("tool rule block holder set is empty")
		}
		if speed, present := entry["speed"]; present {
			rule.Speed = floatValue(speed)
			rule.HasSpeed = true
			if rule.Speed <= 0 || rule.Speed > math.MaxFloat32 {
				return nil, fmt.Errorf("tool rule speed must be positive")
			}
		}
		if _, present := entry["correct_for_drops"]; present {
			rule.CorrectForDrops = nbtBoolOr(entry, "correct_for_drops", false)
			rule.HasCorrectForDrops = true
		}
		tool.Rules = append(tool.Rules, rule)
	}
	return tool, nil
}

func holderSetNBT(entries []string) nbt.Tag {
	if len(entries) == 1 && strings.HasPrefix(entries[0], "#") {
		return nbt.String(entries[0])
	}
	list := nbt.List{ElementType: nbt.TagString}
	for _, entry := range entries {
		list.Elements = append(list.Elements, nbt.String(entry))
	}
	return list
}

func holderSetFromNBT(value nbt.Tag) ([]string, error) {
	if name, ok := value.(nbt.String); ok {
		return []string{string(name)}, nil
	}
	list, ok := value.(nbt.List)
	if !ok || len(list.Elements) > 0 && list.ElementType != nbt.TagString {
		return nil, fmt.Errorf("holder set must be a tag string or string list")
	}
	result := make([]string, 0, len(list.Elements))
	for _, element := range list.Elements {
		name, ok := element.(nbt.String)
		if !ok {
			return nil, fmt.Errorf("holder set entry must be a string")
		}
		result = append(result, string(name))
	}
	return result, nil
}

func textComponentNBT(component *ItemNameComponent) nbt.Tag {
	if component == nil {
		return nil
	}
	if component.Raw != nil {
		if value, err := rawNBTComponent(component.Raw); err == nil {
			return value
		}
	}
	if component.Translate != "" {
		return nbt.Compound{"translate": nbt.String(component.Translate)}
	}
	return nbt.String(component.Text)
}

func textComponentFromNBT(value nbt.Tag) *ItemNameComponent {
	var encoded bytes.Buffer
	_ = nbt.NewWriterTo(&encoded).WriteTag(value, "", true)
	switch value := value.(type) {
	case nbt.String:
		return &ItemNameComponent{Text: string(value), Raw: encoded.Bytes()}
	case nbt.Compound:
		return &ItemNameComponent{Text: value.GetString("text"), Translate: value.GetString("translate"), Raw: encoded.Bytes()}
	case nbt.List:
		if len(value.Elements) > 0 {
			return &ItemNameComponent{Raw: encoded.Bytes()}
		}
	}
	return nil
}

func attributeModifiersNBT(modifiers []AttributeModifier) nbt.List {
	result := nbt.List{ElementType: nbt.TagCompound, Elements: make([]nbt.Tag, 0, len(modifiers))}
	for _, modifier := range modifiers {
		entry := nbt.Compound{
			"type":      nbt.String(modifier.Type),
			"id":        nbt.String(modifier.ID),
			"amount":    nbt.Double(modifier.Amount),
			"operation": nbt.String(modifier.Operation),
		}
		if slot := normalizedEquipmentSlotGroup(modifier.Slot); slot != "any" {
			entry["slot"] = nbt.String(slot)
		}
		switch displayType := normalizedAttributeModifierDisplayType(modifier.Display.Type); displayType {
		case "hidden":
			entry["display"] = nbt.Compound{"type": nbt.String(displayType)}
		case "override":
			entry["display"] = nbt.Compound{"type": nbt.String(displayType), "value": textComponentNBT(modifier.Display.Value)}
		}
		result.Elements = append(result.Elements, entry)
	}
	return result
}

func attributeModifiersFromNBT(value nbt.Tag) ([]AttributeModifier, error) {
	list, ok := value.(nbt.List)
	if !ok || len(list.Elements) > 0 && list.ElementType != nbt.TagCompound {
		return nil, fmt.Errorf("attribute_modifiers component must be a compound list")
	}
	modifiers := make([]AttributeModifier, 0, len(list.Elements))
	for _, element := range list.Elements {
		entry, ok := element.(nbt.Compound)
		if !ok {
			return nil, fmt.Errorf("attribute modifier must be a compound")
		}
		modifier := AttributeModifier{
			Type:      entry.GetString("type"),
			ID:        entry.GetString("id"),
			Amount:    floatValue(entry["amount"]),
			Operation: entry.GetString("operation"),
			Slot:      entry.GetString("slot"),
		}
		if modifier.Slot == "" {
			modifier.Slot = "any"
		}
		if registries.Attribute.Get(modifier.Type) < 0 {
			return nil, fmt.Errorf("unknown attribute type %q", modifier.Type)
		}
		if modifier.ID == "" {
			return nil, fmt.Errorf("attribute modifier requires an id")
		}
		if _, present := entry["amount"]; !present {
			return nil, fmt.Errorf("attribute modifier %q requires an amount", modifier.ID)
		}
		if _, ok := attributeModifierOperationIDs[modifier.Operation]; !ok {
			return nil, fmt.Errorf("unknown attribute modifier operation %q", modifier.Operation)
		}
		if _, ok := equipmentSlotGroupIDs[modifier.Slot]; !ok {
			return nil, fmt.Errorf("unknown equipment slot group %q", modifier.Slot)
		}
		if displayTag, present := entry["display"]; present {
			display, ok := displayTag.(nbt.Compound)
			if !ok {
				return nil, fmt.Errorf("attribute modifier display must be a compound")
			}
			modifier.Display.Type = display.GetString("type")
			if _, ok := attributeModifierDisplayTypeIDs[modifier.Display.Type]; !ok {
				return nil, fmt.Errorf("unknown attribute modifier display type %q", modifier.Display.Type)
			}
			if modifier.Display.Type == "override" {
				override, present := display["value"]
				if !present {
					return nil, fmt.Errorf("override attribute modifier display requires a value")
				}
				modifier.Display.Value = textComponentFromNBT(override)
				if modifier.Display.Value == nil {
					return nil, fmt.Errorf("override attribute modifier display value is invalid")
				}
			}
		}
		modifiers = append(modifiers, modifier)
	}
	return modifiers, nil
}

func enchantmentsNBT(enchantments map[string]int32) nbt.Compound {
	result := nbt.Compound{}
	for name, level := range enchantments {
		result[name] = nbt.Int(level)
	}
	return result
}

func enchantmentsFromNBT(value nbt.Tag) map[string]int32 {
	result := map[string]int32{}
	if compound, ok := value.(nbt.Compound); ok {
		for name, level := range compound {
			result[name] = intValue(level)
		}
	}
	return result
}

func synchronizedHolderName(registry string, holderID ns.VarInt) (string, error) {
	entries := registries.SynchronizedEntries[registry]
	index := int(holderID) - 1
	if index < 0 || index >= len(entries) {
		return "", fmt.Errorf("unknown %s holder %d", registry, holderID)
	}
	return entries[index], nil
}

func writeSynchronizedHolderReference(writer *ns.PacketBuffer, registry, name string) error {
	entryID := registryEntryID(registry, name)
	if entryID < 0 {
		return fmt.Errorf("unknown %s entry %q", registry, name)
	}
	writer.WriteVarInt(ns.VarInt(entryID + 1))
	return nil
}

func soundEventHolderNBT(reader *ns.PacketBuffer) (nbt.Tag, error) {
	holderID, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if holderID > 0 {
		name := registries.SoundEvent.ByID(int32(holderID - 1))
		if name == "" {
			return nil, fmt.Errorf("unknown sound event holder %d", holderID)
		}
		return nbt.String(name), nil
	}
	if holderID < 0 {
		return nil, fmt.Errorf("invalid sound event holder %d", holderID)
	}
	soundID, err := reader.ReadString(maxStringLen)
	if err != nil {
		return nil, err
	}
	result := nbt.Compound{"sound_id": nbt.String(soundID)}
	hasRange, err := reader.ReadBool()
	if err != nil {
		return nil, err
	}
	if hasRange {
		rangeValue, err := reader.ReadFloat32()
		if err != nil {
			return nil, err
		}
		result["range"] = nbt.Float(rangeValue)
	}
	return result, nil
}

func writeSoundEventHolderNBT(writer *ns.PacketBuffer, value nbt.Tag) error {
	if name, ok := value.(nbt.String); ok {
		entryID := registries.SoundEvent.Get(string(name))
		if entryID < 0 {
			return fmt.Errorf("unknown sound event %q", name)
		}
		writer.WriteVarInt(ns.VarInt(entryID + 1))
		return nil
	}
	direct, ok := value.(nbt.Compound)
	if !ok || direct.GetString("sound_id") == "" {
		return fmt.Errorf("sound event must be an identifier or direct compound")
	}
	writer.WriteVarInt(0)
	writer.WriteString(ns.String(direct.GetString("sound_id")))
	rangeTag, hasRange := direct["range"]
	writer.WriteBool(ns.Boolean(hasRange))
	if hasRange {
		writer.WriteFloat32(ns.Float32(floatValue(rangeTag)))
	}
	return nil
}

func instrumentHolderNBT(reader *ns.PacketBuffer) (nbt.Tag, error) {
	holderID, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if holderID > 0 {
		name, err := synchronizedHolderName("minecraft:instrument", holderID)
		return nbt.String(name), err
	}
	if holderID < 0 {
		return nil, fmt.Errorf("invalid instrument holder %d", holderID)
	}
	soundEvent, err := soundEventHolderNBT(reader)
	if err != nil {
		return nil, err
	}
	useDuration, err := reader.ReadFloat32()
	if err != nil {
		return nil, err
	}
	rangeValue, err := reader.ReadFloat32()
	if err != nil {
		return nil, err
	}
	description, err := readNetworkTag(reader)
	if err != nil {
		return nil, err
	}
	return nbt.Compound{
		"sound_event": soundEvent, "use_duration": nbt.Float(useDuration),
		"range": nbt.Float(rangeValue), "description": description,
	}, nil
}

func writeInstrumentHolderNBT(writer *ns.PacketBuffer, value nbt.Tag) error {
	if name, ok := value.(nbt.String); ok {
		return writeSynchronizedHolderReference(writer, "minecraft:instrument", string(name))
	}
	direct, ok := value.(nbt.Compound)
	if !ok {
		return fmt.Errorf("instrument must be an identifier or direct compound")
	}
	writer.WriteVarInt(0)
	if err := writeSoundEventHolderNBT(writer, direct["sound_event"]); err != nil {
		return err
	}
	writer.WriteFloat32(ns.Float32(floatValue(direct["use_duration"])))
	writer.WriteFloat32(ns.Float32(floatValue(direct["range"])))
	description := direct["description"]
	if description == nil {
		return fmt.Errorf("direct instrument requires description")
	}
	return nbt.NewWriterTo(writer.Writer()).WriteTag(description, "", true)
}

func jukeboxSongHolderNBT(reader *ns.PacketBuffer) (nbt.Tag, error) {
	holderID, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if holderID > 0 {
		name, err := synchronizedHolderName("minecraft:jukebox_song", holderID)
		return nbt.String(name), err
	}
	if holderID < 0 {
		return nil, fmt.Errorf("invalid jukebox song holder %d", holderID)
	}
	soundEvent, err := soundEventHolderNBT(reader)
	if err != nil {
		return nil, err
	}
	description, err := readNetworkTag(reader)
	if err != nil {
		return nil, err
	}
	length, err := reader.ReadFloat32()
	if err != nil {
		return nil, err
	}
	output, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	return nbt.Compound{
		"sound_event": soundEvent, "description": description,
		"length_in_seconds": nbt.Float(length), "comparator_output": nbt.Int(output),
	}, nil
}

func writeJukeboxSongHolderNBT(writer *ns.PacketBuffer, value nbt.Tag) error {
	if name, ok := value.(nbt.String); ok {
		return writeSynchronizedHolderReference(writer, "minecraft:jukebox_song", string(name))
	}
	direct, ok := value.(nbt.Compound)
	if !ok {
		return fmt.Errorf("jukebox song must be an identifier or direct compound")
	}
	writer.WriteVarInt(0)
	if err := writeSoundEventHolderNBT(writer, direct["sound_event"]); err != nil {
		return err
	}
	description := direct["description"]
	if description == nil {
		return fmt.Errorf("direct jukebox song requires description")
	}
	if err := nbt.NewWriterTo(writer.Writer()).WriteTag(description, "", true); err != nil {
		return err
	}
	writer.WriteFloat32(ns.Float32(floatValue(direct["length_in_seconds"])))
	writer.WriteVarInt(ns.VarInt(intValue(direct["comparator_output"])))
	return nil
}

func trimMaterialHolderNBT(reader *ns.PacketBuffer) (nbt.Tag, error) {
	holderID, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if holderID > 0 {
		name, err := synchronizedHolderName("minecraft:trim_material", holderID)
		return nbt.String(name), err
	}
	if holderID < 0 {
		return nil, fmt.Errorf("invalid trim material holder %d", holderID)
	}
	assetName, err := reader.ReadString(maxStringLen)
	if err != nil {
		return nil, err
	}
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return nil, fmt.Errorf("invalid trim material override count %d", count)
	}
	overrides := nbt.Compound{}
	for range int(count) {
		asset, err := reader.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		suffix, err := reader.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		overrides[string(asset)] = nbt.String(suffix)
	}
	description, err := readNetworkTag(reader)
	if err != nil {
		return nil, err
	}
	result := nbt.Compound{"asset_name": nbt.String(assetName), "description": description}
	if len(overrides) > 0 {
		result["override_armor_assets"] = overrides
	}
	return result, nil
}

func writeTrimMaterialHolderNBT(writer *ns.PacketBuffer, value nbt.Tag) error {
	if name, ok := value.(nbt.String); ok {
		return writeSynchronizedHolderReference(writer, "minecraft:trim_material", string(name))
	}
	direct, ok := value.(nbt.Compound)
	if !ok {
		return fmt.Errorf("trim material must be an identifier or direct compound")
	}
	assetName := direct.GetString("asset_name")
	if assetName == "" {
		return fmt.Errorf("direct trim material requires asset_name")
	}
	writer.WriteVarInt(0)
	writer.WriteString(ns.String(assetName))
	overrides := direct.GetCompound("override_armor_assets")
	keys := make([]string, 0, len(overrides))
	for key := range overrides {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	writer.WriteVarInt(ns.VarInt(len(keys)))
	for _, key := range keys {
		suffix, ok := overrides[key].(nbt.String)
		if !ok {
			return fmt.Errorf("trim material override %q must be a string", key)
		}
		writer.WriteString(ns.String(key))
		writer.WriteString(ns.String(suffix))
	}
	description := direct["description"]
	if description == nil {
		return fmt.Errorf("direct trim material requires description")
	}
	return nbt.NewWriterTo(writer.Writer()).WriteTag(description, "", true)
}

type filterableString struct {
	raw, filtered string
	hasFiltered   bool
}

func filterableStringNBT(value nbt.Tag, maxLength int) (string, string, bool, error) {
	if raw, ok := value.(nbt.String); ok {
		if len(raw) > maxLength {
			return "", "", false, fmt.Errorf("filterable string exceeds %d bytes", maxLength)
		}
		return string(raw), "", false, nil
	}
	compound, ok := value.(nbt.Compound)
	if !ok {
		return "", "", false, fmt.Errorf("filterable string must be a string or compound")
	}
	raw := compound.GetString("raw")
	filtered, hasFiltered := compound["filtered"].(nbt.String)
	if len(raw) > maxLength || hasFiltered && len(filtered) > maxLength {
		return "", "", false, fmt.Errorf("filterable string exceeds %d bytes", maxLength)
	}
	return raw, string(filtered), hasFiltered, nil
}

func writableBookNBT(data []byte) (nbt.Compound, error) {
	reader := ns.NewReader(data)
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 100 {
		return nil, fmt.Errorf("invalid writable book page count %d", count)
	}
	pages := make([]filterableString, 0, int(count))
	hasFilteredPages := false
	for range int(count) {
		raw, err := reader.ReadString(1024)
		if err != nil {
			return nil, err
		}
		hasFiltered, err := reader.ReadBool()
		if err != nil {
			return nil, err
		}
		page := filterableString{raw: string(raw), hasFiltered: bool(hasFiltered)}
		if hasFiltered {
			filtered, err := reader.ReadString(1024)
			if err != nil {
				return nil, err
			}
			page.filtered = string(filtered)
			hasFilteredPages = true
		}
		pages = append(pages, page)
	}
	list := nbt.List{ElementType: nbt.TagString}
	if hasFilteredPages {
		list.ElementType = nbt.TagCompound
	}
	for _, page := range pages {
		if hasFilteredPages {
			value := nbt.Compound{"raw": nbt.String(page.raw)}
			if page.hasFiltered {
				value["filtered"] = nbt.String(page.filtered)
			}
			list.Elements = append(list.Elements, value)
		} else {
			list.Elements = append(list.Elements, nbt.String(page.raw))
		}
	}
	return nbt.Compound{"pages": list}, nil
}

func filterableTagNBT(value nbt.Tag) (nbt.Tag, nbt.Tag, bool) {
	if compound, ok := value.(nbt.Compound); ok {
		if raw, present := compound["raw"]; present {
			filtered, hasFiltered := compound["filtered"]
			return raw, filtered, hasFiltered
		}
	}
	return value, nil, false
}

func readNetworkTag(reader *ns.PacketBuffer) (nbt.Tag, error) {
	tag, _, err := nbt.NewReaderFrom(reader.Reader()).ReadTag(true)
	return tag, err
}

func writtenBookNBT(data []byte) (nbt.Compound, error) {
	reader := ns.NewReader(data)
	titleRaw, err := reader.ReadString(32)
	if err != nil {
		return nil, err
	}
	titleHasFiltered, err := reader.ReadBool()
	if err != nil {
		return nil, err
	}
	title := nbt.Tag(nbt.String(titleRaw))
	if titleHasFiltered {
		filtered, err := reader.ReadString(32)
		if err != nil {
			return nil, err
		}
		title = nbt.Compound{"raw": nbt.String(titleRaw), "filtered": nbt.String(filtered)}
	}
	author, err := reader.ReadString(32767)
	if err != nil {
		return nil, err
	}
	generation, err := reader.ReadVarInt()
	if err != nil || generation < 0 || generation > 3 {
		return nil, fmt.Errorf("invalid written book generation %d", generation)
	}
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 100 {
		return nil, fmt.Errorf("invalid written book page count %d", count)
	}
	type pageData struct {
		raw, filtered nbt.Tag
		hasFiltered   bool
	}
	pages := make([]pageData, 0, int(count))
	hasFilteredPages := false
	for range int(count) {
		raw, err := readNetworkTag(reader)
		if err != nil {
			return nil, err
		}
		hasFiltered, err := reader.ReadBool()
		if err != nil {
			return nil, err
		}
		page := pageData{raw: raw, hasFiltered: bool(hasFiltered)}
		if hasFiltered {
			page.filtered, err = readNetworkTag(reader)
			if err != nil {
				return nil, err
			}
			hasFilteredPages = true
		}
		pages = append(pages, page)
	}
	resolved, err := reader.ReadBool()
	if err != nil {
		return nil, err
	}
	pageType := nbt.TagString
	if hasFilteredPages {
		pageType = nbt.TagCompound
	} else if len(pages) > 0 {
		pageType = pages[0].raw.ID()
	}
	pageList := nbt.List{ElementType: pageType}
	for _, page := range pages {
		if hasFilteredPages {
			value := nbt.Compound{"raw": page.raw}
			if page.hasFiltered {
				value["filtered"] = page.filtered
			}
			pageList.Elements = append(pageList.Elements, value)
		} else {
			pageList.Elements = append(pageList.Elements, page.raw)
		}
	}
	return nbt.Compound{
		"title": title, "author": nbt.String(author), "generation": nbt.Int(generation),
		"pages": pageList, "resolved": nbt.Byte(map[bool]int8{true: 1}[bool(resolved)]),
	}, nil
}

func textComponentListNBT(data []byte) (nbt.List, error) {
	reader := ns.NewReader(data)
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return nbt.List{}, fmt.Errorf("invalid text component list length %d", count)
	}
	values := make([]nbt.Tag, 0, int(count))
	elementType := byte(nbt.TagEnd)
	for range int(count) {
		value, err := readNetworkTag(reader)
		if err != nil {
			return nbt.List{}, err
		}
		if elementType == nbt.TagEnd {
			elementType = value.ID()
		} else if value.ID() != elementType {
			return nbt.List{}, fmt.Errorf("mixed text component tag types cannot be represented in NBT list")
		}
		values = append(values, value)
	}
	if elementType == nbt.TagEnd {
		elementType = nbt.TagCompound
	}
	return nbt.List{ElementType: elementType, Elements: values}, nil
}

func stringListComponentNBT(data []byte) (nbt.List, error) {
	reader := ns.NewReader(data)
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return nbt.List{}, fmt.Errorf("invalid string list length %d", count)
	}
	result := nbt.List{ElementType: nbt.TagString}
	for range int(count) {
		value, err := reader.ReadString(maxStringLen)
		if err != nil {
			return nbt.List{}, err
		}
		result.Elements = append(result.Elements, nbt.String(value))
	}
	return result, nil
}

func stringMapComponentNBT(data []byte) (nbt.Compound, error) {
	reader := ns.NewReader(data)
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return nil, fmt.Errorf("invalid string map length %d", count)
	}
	result := nbt.Compound{}
	for range int(count) {
		key, err := reader.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		value, err := reader.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		result[string(key)] = nbt.String(value)
	}
	return result, nil
}

func rawNBTComponent(data []byte) (nbt.Tag, error) {
	tag, _, err := nbt.NewReaderFrom(bytes.NewReader(data)).ReadTag(true)
	return tag, err
}

func typedEntityDataNBT(data []byte, registry *registries.Registry) (nbt.Compound, error) {
	reader := ns.NewReader(data)
	typeID, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	typeName := registry.ByID(int32(typeID))
	if typeName == "" {
		return nil, fmt.Errorf("unknown typed entity data id %d", typeID)
	}
	tag, err := readNetworkTag(reader)
	if err != nil {
		return nil, err
	}
	compound, ok := tag.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("typed entity data payload must be a compound")
	}
	result := make(nbt.Compound, len(compound)+1)
	for key, value := range compound {
		result[key] = value
	}
	result["id"] = nbt.String(typeName)
	return result, nil
}

// SlotToNBT converts a non-empty slot to vanilla's item-stack compound.
func SlotToNBT(slot ns.Slot) nbt.Compound {
	return componentSlotNBT(slot)
}

// SlotFromNBT decodes vanilla's item-stack compound.
func SlotFromNBT(value nbt.Tag) (ns.Slot, error) {
	return slotFromComponentNBT(value)
}

func componentSlotNBT(slot ns.Slot) nbt.Compound {
	item := nbt.Compound{"id": nbt.String(ItemName(int32(slot.ItemID))), "count": nbt.Int(slot.Count)}
	if components := ComponentsToNBT(slot); len(components) > 0 {
		item["components"] = components
	}
	return item
}

func slotListComponentNBT(data []byte, optional bool) (nbt.List, error) {
	reader := ns.NewReader(data)
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > 256 {
		return nbt.List{}, fmt.Errorf("invalid item component list length %d", count)
	}
	result := nbt.List{ElementType: nbt.TagCompound}
	for slotIndex := range int(count) {
		if optional {
			hasItem, err := reader.ReadBool()
			if err != nil {
				return nbt.List{}, err
			}
			if !hasItem {
				continue
			}
		}
		slot, err := ReadItemStackTemplate(reader)
		if err != nil {
			return nbt.List{}, err
		}
		item := componentSlotNBT(slot)
		if optional {
			result.Elements = append(result.Elements, nbt.Compound{"slot": nbt.Int(slotIndex), "item": item})
		} else {
			result.Elements = append(result.Elements, item)
		}
	}
	return result, nil
}

func slotFromComponentNBT(value nbt.Tag) (ns.Slot, error) {
	item, ok := value.(nbt.Compound)
	if !ok {
		return ns.Slot{}, fmt.Errorf("item component list entry must be a compound")
	}
	id := ItemID(item.GetString("id"))
	count := item.GetInt("count")
	if id < 0 || count <= 0 {
		return ns.Slot{}, fmt.Errorf("invalid component item %q count %d", item.GetString("id"), count)
	}
	return ApplyNBTComponents(ns.NewSlot(ns.VarInt(id), ns.VarInt(count)), item.GetCompound("components"))
}

func componentSoundNBT(name string, direct bool, fixedRange *float64) nbt.Tag {
	if !direct && registries.SoundEvent.Get(name) >= 0 {
		return nbt.String(name)
	}
	value := nbt.Compound{"sound_id": nbt.String(name)}
	if fixedRange != nil {
		value["range"] = nbt.Float(*fixedRange)
	}
	return value
}

func consumableSoundFromNBT(value nbt.Tag) (consumableSound, error) {
	writer := ns.NewWriter()
	if err := writeSoundEventHolderNBT(writer, value); err != nil {
		return consumableSound{}, err
	}
	return readConsumableSound(ns.NewReader(writer.Bytes()))
}

func statusEffectComponentNBT(effect StatusEffect, includeID bool) nbt.Compound {
	result := nbt.Compound{
		"amplifier": nbt.Int(effect.Amplifier), "duration": nbt.Int(effect.Duration),
		"ambient":        nbt.Byte(map[bool]int8{true: 1}[effect.Ambient]),
		"show_particles": nbt.Byte(map[bool]int8{true: 1}[effect.ShowParticles]),
		"show_icon":      nbt.Byte(map[bool]int8{true: 1}[effect.ShowIcon]),
	}
	if includeID {
		result["id"] = nbt.String(effect.ID)
	}
	if effect.Hidden != nil {
		result["hidden_effect"] = statusEffectComponentNBT(*effect.Hidden, false)
	}
	return result
}

func statusEffectComponentFromNBT(value nbt.Tag, inheritedID string) (StatusEffect, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return StatusEffect{}, fmt.Errorf("status effect must be a compound")
	}
	effectID := compound.GetString("id")
	if effectID == "" {
		effectID = inheritedID
	}
	if registries.MobEffect.Get(effectID) < 0 {
		return StatusEffect{}, fmt.Errorf("unknown status effect %q", effectID)
	}
	status := StatusEffect{
		ID: effectID, Amplifier: compound.GetInt("amplifier"), Duration: compound.GetInt("duration"),
		Ambient: nbtBoolOr(compound, "ambient", false), ShowParticles: nbtBoolOr(compound, "show_particles", true),
	}
	status.ShowIcon = nbtBoolOr(compound, "show_icon", status.ShowParticles)
	if hidden, present := compound["hidden_effect"]; present {
		parsed, err := statusEffectComponentFromNBT(hidden, effectID)
		if err != nil {
			return StatusEffect{}, err
		}
		status.Hidden = &parsed
	}
	return status, nil
}

func consumeEffectNBT(effect ConsumeEffect) nbt.Compound {
	result := nbt.Compound{"type": nbt.String(effect.Type)}
	switch effect.Type {
	case "minecraft:apply_effects":
		effects := nbt.List{ElementType: nbt.TagCompound}
		for _, status := range effect.Effects {
			effects.Elements = append(effects.Elements, statusEffectComponentNBT(status, true))
		}
		result["effects"] = effects
		result["probability"] = nbt.Float(effect.Probability)
	case "minecraft:remove_effects":
		if len(effect.EffectIDs) == 1 && strings.HasPrefix(effect.EffectIDs[0], "#") {
			result["effects"] = nbt.String(effect.EffectIDs[0])
		} else {
			effects := nbt.List{ElementType: nbt.TagString}
			for _, effectID := range effect.EffectIDs {
				effects.Elements = append(effects.Elements, nbt.String(effectID))
			}
			result["effects"] = effects
		}
	case "minecraft:teleport_randomly":
		result["diameter"] = nbt.Float(effect.Diameter)
	case "minecraft:play_sound":
		result["sound"] = componentSoundNBT(effect.Sound, effect.SoundDirect, effect.SoundRange)
	}
	return result
}

func consumeEffectFromNBT(value nbt.Tag) (ConsumeEffect, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return ConsumeEffect{}, fmt.Errorf("consume effect must be a compound")
	}
	typeName := compound.GetString("type")
	effect := ConsumeEffect{Type: typeName, Probability: 1}
	switch typeName {
	case "minecraft:apply_effects":
		list := compound.GetList("effects")
		for _, element := range list.Elements {
			status, err := statusEffectComponentFromNBT(element, "")
			if err != nil {
				return ConsumeEffect{}, err
			}
			effect.Effects = append(effect.Effects, status)
		}
		if probability, present := compound["probability"]; present {
			effect.Probability = floatValue(probability)
		}
		if effect.Probability < 0 || effect.Probability > 1 {
			return ConsumeEffect{}, fmt.Errorf("consume effect probability must be between 0 and 1")
		}
	case "minecraft:remove_effects":
		switch effects := compound["effects"].(type) {
		case nbt.String:
			effect.EffectIDs = []string{string(effects)}
		case nbt.List:
			for _, element := range effects.Elements {
				name, ok := element.(nbt.String)
				if !ok || registries.MobEffect.Get(string(name)) < 0 {
					return ConsumeEffect{}, fmt.Errorf("unknown removed status effect %q", name)
				}
				effect.EffectIDs = append(effect.EffectIDs, string(name))
			}
		default:
			return ConsumeEffect{}, fmt.Errorf("remove_effects requires an effect list or tag")
		}
	case "minecraft:clear_all_effects":
	case "minecraft:teleport_randomly":
		effect.Diameter = 16
		if diameter, present := compound["diameter"]; present {
			effect.Diameter = floatValue(diameter)
		}
		if effect.Diameter <= 0 {
			return ConsumeEffect{}, fmt.Errorf("teleport diameter must be positive")
		}
	case "minecraft:play_sound":
		sound, err := consumableSoundFromNBT(compound["sound"])
		if err != nil {
			return ConsumeEffect{}, err
		}
		effect.Sound, effect.SoundDirect, effect.SoundRange = sound.name, sound.direct, sound.fixedRange
	default:
		return ConsumeEffect{}, fmt.Errorf("unknown consume effect type %q", typeName)
	}
	return effect, nil
}

func potionEffectNBT(effect PotionEffect, includeID bool) nbt.Compound {
	result := nbt.Compound{
		"amplifier": nbt.Int(effect.Amplifier), "duration": nbt.Int(effect.Duration),
		"ambient":        nbt.Byte(map[bool]int8{true: 1}[effect.Ambient]),
		"show_particles": nbt.Byte(map[bool]int8{true: 1}[effect.ShowParticles]),
		"show_icon":      nbt.Byte(map[bool]int8{true: 1}[effect.ShowIcon]),
	}
	if includeID {
		result["id"] = nbt.String(effect.Effect)
	}
	if effect.Hidden != nil {
		result["hidden_effect"] = potionEffectNBT(*effect.Hidden, false)
	}
	return result
}

func nbtBoolOr(compound nbt.Compound, key string, fallback bool) bool {
	value, present := compound[key]
	if !present {
		return fallback
	}
	return intValue(value) != 0
}

func writePotionEffectNBT(writer *ns.PacketBuffer, effect nbt.Compound) error {
	writer.WriteVarInt(ns.VarInt(effect.GetInt("amplifier")))
	writer.WriteVarInt(ns.VarInt(effect.GetInt("duration")))
	writer.WriteBool(ns.Boolean(nbtBoolOr(effect, "ambient", false)))
	showParticles := nbtBoolOr(effect, "show_particles", true)
	writer.WriteBool(ns.Boolean(showParticles))
	writer.WriteBool(ns.Boolean(nbtBoolOr(effect, "show_icon", showParticles)))
	hidden, hasHidden := effect["hidden_effect"].(nbt.Compound)
	writer.WriteBool(ns.Boolean(hasHidden))
	if hasHidden {
		return writePotionEffectNBT(writer, hidden)
	}
	return nil
}

func floatValue(value nbt.Tag) float64 {
	switch value := value.(type) {
	case nbt.Float:
		return float64(value)
	case nbt.Double:
		return float64(value)
	case nbt.Byte:
		return float64(value)
	case nbt.Short:
		return float64(value)
	case nbt.Int:
		return float64(value)
	case nbt.Long:
		return float64(value)
	default:
		return 0
	}
}

func intValue(value nbt.Tag) int32 {
	switch value := value.(type) {
	case nbt.Byte:
		return int32(value)
	case nbt.Short:
		return int32(value)
	case nbt.Int:
		return int32(value)
	case nbt.Long:
		return int32(value)
	default:
		return 0
	}
}

var fireworkShapeNames = [5]string{"small_ball", "large_ball", "star", "creeper", "burst"}

func fireworkExplosionNBT(explosion FireworkExplosion) nbt.Compound {
	return nbt.Compound{
		"shape": nbt.String(fireworkShapeNames[explosion.Shape]), "colors": nbt.IntArray(explosion.Colors),
		"fade_colors": nbt.IntArray(explosion.FadeColors), "has_trail": nbt.Byte(map[bool]int8{true: 1}[explosion.HasTrail]),
		"has_twinkle": nbt.Byte(map[bool]int8{true: 1}[explosion.HasTwinkle]),
	}
}

func fireworkExplosionFromNBT(value nbt.Tag) (FireworkExplosion, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return FireworkExplosion{}, fmt.Errorf("firework explosion must be a compound")
	}
	shape := int32(-1)
	for id, name := range fireworkShapeNames {
		if compound.GetString("shape") == name {
			shape = int32(id)
			break
		}
	}
	if shape < 0 {
		return FireworkExplosion{}, fmt.Errorf("unknown firework shape %q", compound.GetString("shape"))
	}
	colors, ok := compound["colors"].(nbt.IntArray)
	if !ok && compound["colors"] != nil {
		return FireworkExplosion{}, fmt.Errorf("firework colors must be an int array")
	}
	fadeColors, ok := compound["fade_colors"].(nbt.IntArray)
	if !ok && compound["fade_colors"] != nil {
		return FireworkExplosion{}, fmt.Errorf("firework fade_colors must be an int array")
	}
	if len(colors) > 256 || len(fadeColors) > 256 {
		return FireworkExplosion{}, fmt.Errorf("too many firework colors")
	}
	return FireworkExplosion{
		Shape: shape, Colors: append([]int32(nil), colors...), FadeColors: append([]int32(nil), fadeColors...),
		HasTrail: compound.GetByte("has_trail") != 0, HasTwinkle: compound.GetByte("has_twinkle") != 0,
	}, nil
}

var dyeNames = [16]string{
	"white", "orange", "magenta", "light_blue", "yellow", "lime", "pink", "gray",
	"light_gray", "cyan", "purple", "blue", "brown", "green", "red", "black",
}

func dyeID(name string) int32 {
	for id, entry := range dyeNames {
		if entry == name {
			return int32(id)
		}
	}
	return -1
}

func registryEntryID(registry, name string) int {
	for id, entry := range registries.SynchronizedEntries[registry] {
		if entry == name {
			return id
		}
	}
	return -1
}

func replaceRawComponent(components []ns.RawSlotComponent, replacement ns.RawSlotComponent) []ns.RawSlotComponent {
	for i := range components {
		if components[i].ID == replacement.ID {
			components[i] = replacement
			return components
		}
	}
	return append(components, replacement)
}

func removeRawComponent(components []ns.RawSlotComponent, id ns.VarInt) []ns.RawSlotComponent {
	for i := range components {
		if components[i].ID == id {
			return append(components[:i], components[i+1:]...)
		}
	}
	return components
}
