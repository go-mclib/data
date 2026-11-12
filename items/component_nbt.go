package items

import (
	"cmp"
	"fmt"
	"maps"
	"math"
	"slices"
	"strings"

	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// nbtBridge converts one component between vanilla item NBT and its wire payload.
type nbtBridge struct {
	// toNBT renders a patched component. c holds the decoded stack and data
	// the component's wire payload; a nil tag omits the component.
	toNBT func(c *Components, data []byte) (nbt.Tag, error)
	// fromNBT encodes a vanilla NBT value as the component's wire payload.
	fromNBT func(value nbt.Tag) ([]byte, error)
}

// fieldBridge bridges a typed component through its Components field: fromNBT
// fills a scratch Components that the component's wire codec then encodes.
func fieldBridge(id int32, toNBT func(*Components) nbt.Tag, fromNBT func(*Components, nbt.Tag) error) nbtBridge {
	return nbtBridge{
		toNBT: func(c *Components, _ []byte) (nbt.Tag, error) { return toNBT(c), nil },
		fromNBT: func(value nbt.Tag) ([]byte, error) {
			var c Components
			if err := fromNBT(&c, value); err != nil {
				return nil, err
			}
			return encodeComponent(&c, id)
		},
	}
}

// wireBridge bridges a component directly between NBT and its wire payload.
func wireBridge(toNBT func(*ns.PacketBuffer) (nbt.Tag, error), fromNBT func(*ns.PacketBuffer, nbt.Tag) error) nbtBridge {
	return nbtBridge{
		toNBT: func(_ *Components, data []byte) (nbt.Tag, error) { return toNBT(ns.NewReader(data)) },
		fromNBT: func(value nbt.Tag) ([]byte, error) {
			w := ns.NewWriter()
			if err := fromNBT(w, value); err != nil {
				return nil, err
			}
			return w.Bytes(), nil
		},
	}
}

func intFieldBridge(id int32, field func(*Components) *int32) nbtBridge {
	return fieldBridge(id,
		func(c *Components) nbt.Tag { return nbt.Int(*field(c)) },
		func(c *Components, value nbt.Tag) error { *field(c) = intValue(value); return nil })
}

func floatFieldBridge(id int32, field func(*Components) *float64) nbtBridge {
	return fieldBridge(id,
		func(c *Components) nbt.Tag { return nbt.Float(*field(c)) },
		func(c *Components, value nbt.Tag) error { *field(c) = floatValue(value); return nil })
}

// stringFieldBridge bridges a string component; valid rejects unknown values.
func stringFieldBridge(id int32, name string, field func(*Components) *string, valid func(string) bool) nbtBridge {
	return fieldBridge(id,
		func(c *Components) nbt.Tag {
			if *field(c) == "" {
				return nil
			}
			return nbt.String(*field(c))
		},
		func(c *Components, value nbt.Tag) error {
			text, ok := value.(nbt.String)
			if !ok || text == "" || valid != nil && !valid(string(text)) {
				return fmt.Errorf("%s must be a valid identifier, got %v", name, value)
			}
			*field(c) = string(text)
			return nil
		})
}

// markerFieldBridge bridges a unit component rendered as an empty compound.
func markerFieldBridge(id int32, field func(*Components) *bool) nbtBridge {
	return fieldBridge(id,
		func(*Components) nbt.Tag { return nbt.Compound{} },
		func(c *Components, value nbt.Tag) error {
			if _, ok := value.(nbt.Compound); !ok {
				return fmt.Errorf("unit component must be an empty compound")
			}
			*field(c) = true
			return nil
		})
}

// rawNBTBridge bridges a component whose wire payload is a single NBT tag.
func rawNBTBridge(valid func(nbt.Tag) bool) nbtBridge {
	return wireBridge(
		func(r *ns.PacketBuffer) (nbt.Tag, error) { return r.ReadNBT() },
		func(w *ns.PacketBuffer, value nbt.Tag) error {
			if !valid(value) {
				return fmt.Errorf("invalid NBT component value %v", value)
			}
			return w.WriteNBT(value)
		})
}

func isCompound(value nbt.Tag) bool {
	_, ok := value.(nbt.Compound)
	return ok
}

func textFieldBridge(id int32, field func(*Components) **ItemNameComponent) nbtBridge {
	return fieldBridge(id,
		func(c *Components) nbt.Tag { return textComponentNBT(*field(c)) },
		func(c *Components, value nbt.Tag) error {
			if *field(c) = textComponentFromNBT(value); *field(c) == nil {
				return fmt.Errorf("invalid text component %v", value)
			}
			return nil
		})
}

func enchantmentsFieldBridge(id int32, field func(*Components) *map[string]int32) nbtBridge {
	return fieldBridge(id,
		func(c *Components) nbt.Tag { return enchantmentsNBT(*field(c)) },
		func(c *Components, value nbt.Tag) error { *field(c) = enchantmentsFromNBT(value); return nil })
}

func dyeFieldBridge(id int32, field func(*Components) **int32) nbtBridge {
	return fieldBridge(id,
		func(c *Components) nbt.Tag {
			if color := *field(c); color != nil && *color >= 0 && int(*color) < len(dyeNames) {
				return nbt.String(dyeNames[*color])
			}
			return nil
		},
		func(c *Components, value nbt.Tag) error {
			name, ok := value.(nbt.String)
			color := dyeID(string(name))
			if !ok || color < 0 {
				return fmt.Errorf("unknown dye color %v", value)
			}
			*field(c) = &color
			return nil
		})
}

// aquaticVariantBridge bridges an enum entity-variant component stored as a VarInt.
func aquaticVariantBridge(names map[int32]string) nbtBridge {
	return wireBridge(
		func(r *ns.PacketBuffer) (nbt.Tag, error) {
			id, err := r.ReadVarInt()
			if err != nil {
				return nil, err
			}
			name, ok := names[int32(id)]
			if !ok {
				return nil, fmt.Errorf("unknown variant %d", id)
			}
			return nbt.String(name), nil
		},
		func(w *ns.PacketBuffer, value nbt.Tag) error {
			name, ok := value.(nbt.String)
			if !ok {
				return fmt.Errorf("variant must be a string")
			}
			for id, candidate := range names {
				if candidate == string(name) {
					return w.WriteVarInt(ns.VarInt(id))
				}
			}
			return fmt.Errorf("unknown variant %q", name)
		})
}

func dyeNameMap() map[int32]string {
	names := make(map[int32]string, len(dyeNames))
	for id, name := range dyeNames {
		names[int32(id)] = name
	}
	return names
}

// nbtBridges maps component IDs to their NBT bridges. It is built in init
// because bridges for item lists recurse into ApplyNBTComponents.
var nbtBridges map[int32]nbtBridge

func init() {
	nbtBridges = map[int32]nbtBridge{
		ComponentAttackAnimation: fieldBridge(ComponentAttackAnimation,
			func(c *Components) nbt.Tag { return swingAnimationNBT(c.AttackAnimation) },
			func(c *Components, value nbt.Tag) (err error) {
				c.AttackAnimation, err = swingAnimationFromNBT(value)
				return err
			}),
		ComponentInteractAnimation: fieldBridge(ComponentInteractAnimation,
			func(c *Components) nbt.Tag { return swingAnimationNBT(c.InteractAnimation) },
			func(c *Components, value nbt.Tag) (err error) {
				c.InteractAnimation, err = swingAnimationFromNBT(value)
				return err
			}),
		ComponentBlockTransformer: stringFieldBridge(ComponentBlockTransformer, "minecraft:block_transformer",
			func(c *Components) *string { return &c.BlockTransformer }, nil),
		ComponentProvidesPotteryPattern: stringFieldBridge(ComponentProvidesPotteryPattern, "minecraft:provides_pottery_pattern",
			func(c *Components) *string { return &c.ProvidesPotteryPattern }, nil),
		ComponentVillagerFood: fieldBridge(ComponentVillagerFood,
			func(c *Components) nbt.Tag {
				if c.VillagerFood == nil {
					return nil
				}
				return nbt.Compound{"nutrition": nbt.Int(c.VillagerFood.Nutrition)}
			},
			func(c *Components, value nbt.Tag) error {
				compound, ok := value.(nbt.Compound)
				if !ok {
					return fmt.Errorf("villager_food must be a compound")
				}
				nutrition, err := componentIntNBT(compound["nutrition"])
				if err != nil {
					return err
				}
				if nutrition <= 0 {
					return fmt.Errorf("villager food nutrition must be positive")
				}
				c.VillagerFood = &VillagerFood{Nutrition: nutrition}
				return nil
			}),
		ComponentCompostable: fieldBridge(ComponentCompostable,
			func(c *Components) nbt.Tag {
				if c.Compostable == nil {
					return nil
				}
				return nbt.Compound{"layers": resolvableIntNBT(c.Compostable.Layers)}
			},
			func(c *Components, value nbt.Tag) error {
				compound, ok := value.(nbt.Compound)
				if !ok {
					return fmt.Errorf("compostable must be a compound")
				}
				layers, err := resolvableIntFromNBT(compound["layers"])
				if err != nil {
					return err
				}
				c.Compostable = &Compostable{Layers: layers}
				return nil
			}),
		ComponentCookingFuel: fieldBridge(ComponentCookingFuel,
			func(c *Components) nbt.Tag {
				if c.CookingFuel == nil {
					return nil
				}
				return nbt.Compound{"burn_time": resolvableIntNBT(c.CookingFuel.BurnTime), "speed_multiplier": resolvableFloatNBT(c.CookingFuel.SpeedMultiplier)}
			},
			func(c *Components, value nbt.Tag) error {
				amount, speed, err := fuelFromNBT(value, "burn_time")
				c.CookingFuel = &CookingFuel{BurnTime: amount, SpeedMultiplier: speed}
				return err
			}),
		ComponentBrewingFuel: fieldBridge(ComponentBrewingFuel,
			func(c *Components) nbt.Tag {
				if c.BrewingFuel == nil {
					return nil
				}
				return nbt.Compound{"uses": resolvableIntNBT(c.BrewingFuel.Uses), "speed_multiplier": resolvableFloatNBT(c.BrewingFuel.SpeedMultiplier)}
			},
			func(c *Components, value nbt.Tag) error {
				amount, speed, err := fuelFromNBT(value, "uses")
				c.BrewingFuel = &BrewingFuel{Uses: amount, SpeedMultiplier: speed}
				return err
			}),
		ComponentMobVisibility: fieldBridge(ComponentMobVisibility,
			func(c *Components) nbt.Tag {
				if c.MobVisibility == nil {
					return nil
				}
				return nbt.Compound{"targeting_entity_types": holderSetNBT(c.MobVisibility.TargetingEntityTypes), "visibility": nbt.Float(c.MobVisibility.Visibility)}
			},
			func(c *Components, value nbt.Tag) (err error) {
				c.MobVisibility, err = mobVisibilityFromNBT(value)
				return err
			}),
		ComponentSignTextFront: fieldBridge(ComponentSignTextFront,
			func(c *Components) nbt.Tag { return signTextNBT(c.SignTextFront) },
			func(c *Components, value nbt.Tag) (err error) {
				c.SignTextFront, err = signTextFromNBT(value)
				return err
			}),
		ComponentSignTextBack: fieldBridge(ComponentSignTextBack,
			func(c *Components) nbt.Tag { return signTextNBT(c.SignTextBack) },
			func(c *Components, value nbt.Tag) (err error) {
				c.SignTextBack, err = signTextFromNBT(value)
				return err
			}),
		ComponentWaxed:        markerFieldBridge(ComponentWaxed, func(c *Components) *bool { return &c.Waxed }),
		ComponentCushionColor: dyeFieldBridge(ComponentCushionColor, func(c *Components) **int32 { return &c.CushionColor }),
		ComponentPotDecorations: fieldBridge(ComponentPotDecorations,
			func(c *Components) nbt.Tag { return potDecorationsNBT(c.PotDecorations) },
			func(c *Components, value nbt.Tag) (err error) {
				c.PotDecorations, err = potDecorationsFromNBT(value)
				return err
			}),

		ComponentAxolotlVariant: aquaticVariantBridge(map[int32]string{
			0: "lucy", 1: "wild", 2: "gold", 3: "cyan", 4: "blue",
		}),
		ComponentTropicalFishPattern: aquaticVariantBridge(map[int32]string{
			0: "kob", 256: "sunstreak", 512: "snooper", 768: "dasher", 1024: "brinely", 1280: "spotty",
			1: "flopper", 257: "stripey", 513: "glitter", 769: "blockfish", 1025: "betty", 1281: "clayfish",
		}),
		ComponentTropicalFishBaseColor:    aquaticVariantBridge(dyeNameMap()),
		ComponentTropicalFishPatternColor: aquaticVariantBridge(dyeNameMap()),

		ComponentDamage:                 intFieldBridge(ComponentDamage, func(c *Components) *int32 { return &c.Damage }),
		ComponentMaxStackSize:           intFieldBridge(ComponentMaxStackSize, func(c *Components) *int32 { return &c.MaxStackSize }),
		ComponentMaxDamage:              intFieldBridge(ComponentMaxDamage, func(c *Components) *int32 { return &c.MaxDamage }),
		ComponentAdditionalTradeCost:    intFieldBridge(ComponentAdditionalTradeCost, func(c *Components) *int32 { return &c.AdditionalTradeCost }),
		ComponentRepairCost:             intFieldBridge(ComponentRepairCost, func(c *Components) *int32 { return &c.RepairCost }),
		ComponentOminousBottleAmplifier: intFieldBridge(ComponentOminousBottleAmplifier, func(c *Components) *int32 { return &c.OminousBottleAmplifier }),
		ComponentMapId: fieldBridge(ComponentMapId,
			func(c *Components) nbt.Tag {
				if c.MapID == nil {
					return nil
				}
				return nbt.Int(*c.MapID)
			},
			func(c *Components, value nbt.Tag) error { c.MapID = int32Pointer(intValue(value)); return nil }),
		ComponentPotionDurationScale: floatFieldBridge(ComponentPotionDurationScale, func(c *Components) *float64 { return &c.PotionDurationScale }),
		ComponentMinimumAttackCharge: floatFieldBridge(ComponentMinimumAttackCharge, func(c *Components) *float64 { return &c.MinimumAttackCharge }),
		ComponentItemModel:           stringFieldBridge(ComponentItemModel, "minecraft:item_model", func(c *Components) *string { return &c.ItemModel }, nil),
		ComponentRarity: stringFieldBridge(ComponentRarity, "minecraft:rarity", func(c *Components) *string { return &c.Rarity },
			func(name string) bool { return slices.Contains(rarityNames, name) }),
		ComponentAttributeModifiers: fieldBridge(ComponentAttributeModifiers,
			func(c *Components) nbt.Tag { return attributeModifiersNBT(c.AttributeModifiers) },
			func(c *Components, value nbt.Tag) (err error) {
				c.AttributeModifiers, err = attributeModifiersFromNBT(value)
				return err
			}),
		ComponentEquippable: fieldBridge(ComponentEquippable,
			func(c *Components) nbt.Tag { return equippableNBT(c.Equippable) },
			func(c *Components, value nbt.Tag) (err error) {
				c.Equippable, err = equippableFromNBT(value)
				return err
			}),

		ComponentBreakSound:           wireBridge(soundEventHolderNBT, writeSoundEventHolderNBT),
		ComponentInstrument:           wireBridge(instrumentHolderNBT, writeInstrumentHolderNBT),
		ComponentJukeboxPlayable:      wireBridge(jukeboxSongHolderNBT, writeJukeboxSongHolderNBT),
		ComponentProvidesTrimMaterial: wireBridge(trimMaterialHolderNBT, writeTrimMaterialHolderNBT),

		ComponentFood: fieldBridge(ComponentFood,
			func(c *Components) nbt.Tag {
				if c.Food == nil {
					return nil
				}
				return nbt.Compound{
					"nutrition": nbt.Int(c.Food.Nutrition), "saturation": nbt.Float(c.Food.Saturation),
					"can_always_eat": nbtBool(c.Food.CanAlwaysEat),
				}
			},
			func(c *Components, value nbt.Tag) error {
				compound, ok := value.(nbt.Compound)
				if !ok {
					return fmt.Errorf("food component must be a compound")
				}
				nutrition := intValue(compound["nutrition"])
				if nutrition < 0 {
					return fmt.Errorf("food nutrition must be non-negative")
				}
				c.Food = &Food{
					Nutrition: nutrition, Saturation: floatValue(compound["saturation"]),
					CanAlwaysEat: nbtBoolOr(compound, "can_always_eat", false),
				}
				return nil
			}),
		ComponentBlocksAttacks: fieldBridge(ComponentBlocksAttacks,
			func(c *Components) nbt.Tag { return blocksAttacksNBT(c.BlocksAttacks) },
			func(c *Components, value nbt.Tag) (err error) {
				c.BlocksAttacks, err = blocksAttacksFromNBT(value)
				return err
			}),
		ComponentConsumable: fieldBridge(ComponentConsumable,
			func(c *Components) nbt.Tag { return consumableNBT(c.Consumable) },
			func(c *Components, value nbt.Tag) (err error) {
				c.Consumable, err = consumableFromNBT(value)
				return err
			}),
		ComponentDeathProtection: fieldBridge(ComponentDeathProtection,
			func(c *Components) nbt.Tag {
				if c.DeathProtection == nil {
					return nil
				}
				return nbt.Compound{"death_effects": consumeEffectsNBT(c.DeathProtection.DeathEffects)}
			},
			func(c *Components, value nbt.Tag) error {
				compound, ok := value.(nbt.Compound)
				if !ok {
					return fmt.Errorf("death_protection component must be a compound")
				}
				effects, err := consumeEffectsFromNBT(compound.GetList("death_effects"))
				c.DeathProtection = &DeathProtection{DeathEffects: effects}
				return err
			}),
		ComponentWeapon: fieldBridge(ComponentWeapon,
			func(c *Components) nbt.Tag {
				if c.Weapon == nil {
					return nil
				}
				return nbt.Compound{
					"item_damage_per_attack":       nbt.Int(c.Weapon.ItemDamagePerAttack),
					"disable_blocking_for_seconds": nbt.Float(c.Weapon.DisableBlockingForSeconds),
				}
			},
			func(c *Components, value nbt.Tag) error {
				compound, ok := value.(nbt.Compound)
				if !ok {
					return fmt.Errorf("weapon component must be a compound")
				}
				damage := int32(1)
				if tag, present := compound["item_damage_per_attack"]; present {
					damage = intValue(tag)
				}
				disable := floatValue(compound["disable_blocking_for_seconds"])
				if damage < 0 || disable < 0 {
					return fmt.Errorf("weapon values must be non-negative")
				}
				c.Weapon = &Weapon{ItemDamagePerAttack: damage, DisableBlockingForSeconds: disable}
				return nil
			}),
		ComponentAttackRange: fieldBridge(ComponentAttackRange,
			func(c *Components) nbt.Tag { return attackRangeNBT(c.AttackRange) },
			func(c *Components, value nbt.Tag) (err error) {
				c.AttackRange, err = attackRangeFromNBT(value)
				return err
			}),
		ComponentTool: fieldBridge(ComponentTool,
			func(c *Components) nbt.Tag { return toolNBT(c.Tool) },
			func(c *Components, value nbt.Tag) (err error) { c.Tool, err = toolFromNBT(value); return err }),
		ComponentEnchantable: fieldBridge(ComponentEnchantable,
			func(c *Components) nbt.Tag {
				if c.Enchantable == nil {
					return nil
				}
				return nbt.Compound{"value": nbt.Int(c.Enchantable.Value)}
			},
			func(c *Components, value nbt.Tag) error {
				compound, ok := value.(nbt.Compound)
				if !ok || intValue(compound["value"]) <= 0 {
					return fmt.Errorf("enchantable value must be positive")
				}
				c.Enchantable = &Enchantable{Value: intValue(compound["value"])}
				return nil
			}),
		ComponentTooltipDisplay: fieldBridge(ComponentTooltipDisplay,
			func(c *Components) nbt.Tag { return tooltipDisplayNBT(c.TooltipDisplay) },
			func(c *Components, value nbt.Tag) (err error) {
				c.TooltipDisplay, err = tooltipDisplayFromNBT(value)
				return err
			}),
		ComponentUseEffects: fieldBridge(ComponentUseEffects,
			func(c *Components) nbt.Tag {
				if c.UseEffects == nil {
					return nil
				}
				return nbt.Compound{
					"can_sprint": nbtBool(c.UseEffects.CanSprint), "interact_vibrations": nbtBool(c.UseEffects.InteractVibrations),
					"speed_multiplier": nbt.Float(c.UseEffects.SpeedMultiplier),
				}
			},
			func(c *Components, value nbt.Tag) error {
				compound, ok := value.(nbt.Compound)
				if !ok {
					return fmt.Errorf("use_effects component must be a compound")
				}
				speed := 0.2
				if tag, present := compound["speed_multiplier"]; present {
					speed = floatValue(tag)
				}
				if speed < 0 || speed > 1 {
					return fmt.Errorf("use_effects speed_multiplier must be between 0 and 1")
				}
				c.UseEffects = &UseEffects{
					CanSprint:          nbtBoolOr(compound, "can_sprint", false),
					InteractVibrations: nbtBoolOr(compound, "interact_vibrations", true),
					SpeedMultiplier:    speed,
				}
				return nil
			}),
		ComponentUseCooldown: fieldBridge(ComponentUseCooldown,
			func(c *Components) nbt.Tag {
				if c.UseCooldown == nil {
					return nil
				}
				value := nbt.Compound{"seconds": nbt.Float(c.UseCooldown.Seconds)}
				if c.UseCooldown.CooldownGroup != "" {
					value["cooldown_group"] = nbt.String(c.UseCooldown.CooldownGroup)
				}
				return value
			},
			func(c *Components, value nbt.Tag) error {
				compound, ok := value.(nbt.Compound)
				seconds := floatValue(compound["seconds"])
				if !ok || seconds <= 0 {
					return fmt.Errorf("use_cooldown seconds must be positive")
				}
				c.UseCooldown = &UseCooldown{Seconds: seconds, CooldownGroup: compound.GetString("cooldown_group")}
				return nil
			}),
		ComponentGlider:      markerFieldBridge(ComponentGlider, func(c *Components) *bool { return &c.Glider }),
		ComponentUnbreakable: markerFieldBridge(ComponentUnbreakable, func(c *Components) *bool { return &c.Unbreakable }),
		ComponentCreativeSlotLock: wireBridge(
			func(*ns.PacketBuffer) (nbt.Tag, error) { return nbt.Compound{}, nil },
			func(*ns.PacketBuffer, nbt.Tag) error { return nil }),
		ComponentIntangibleProjectile: rawNBTBridge(isCompound),
		ComponentEnchantmentGlintOverride: wireBridge(
			func(r *ns.PacketBuffer) (nbt.Tag, error) {
				value, err := r.ReadBool()
				return nbtBool(bool(value)), err
			},
			func(w *ns.PacketBuffer, value nbt.Tag) error { return w.WriteBool(intValue(value) != 0) }),
		ComponentNoteBlockSound: stringWireBridge(),
		ComponentTooltipStyle:   stringWireBridge(),

		ComponentBundleContents:     wireBridge(func(r *ns.PacketBuffer) (nbt.Tag, error) { return slotListComponentNBT(r, false) }, writeSlotListNBT),
		ComponentChargedProjectiles: wireBridge(func(r *ns.PacketBuffer) (nbt.Tag, error) { return slotListComponentNBT(r, false) }, writeSlotListNBT),
		ComponentContainer:          wireBridge(func(r *ns.PacketBuffer) (nbt.Tag, error) { return slotListComponentNBT(r, true) }, writeContainerNBT),
		ComponentSulfurCubeContent:  wireBridge(itemTemplateNBT, writeItemTemplateNBT),
		ComponentUseRemainder:       wireBridge(itemTemplateNBT, writeItemTemplateNBT),
		ComponentEntityData: wireBridge(
			func(r *ns.PacketBuffer) (nbt.Tag, error) { return readTypedEntityDataNBT(r, registries.EntityType) },
			func(w *ns.PacketBuffer, value nbt.Tag) error {
				return writeTypedEntityDataNBT(w, value, registries.EntityType)
			}),
		ComponentBlockEntityData: wireBridge(
			func(r *ns.PacketBuffer) (nbt.Tag, error) {
				return readTypedEntityDataNBT(r, registries.BlockEntityType)
			},
			func(w *ns.PacketBuffer, value nbt.Tag) error {
				return writeTypedEntityDataNBT(w, value, registries.BlockEntityType)
			}),
		ComponentCustomData:       rawNBTBridge(isCompound),
		ComponentBucketEntityData: rawNBTBridge(isCompound),
		ComponentDebugStickState:  rawNBTBridge(isCompound),
		ComponentMapDecorations:   rawNBTBridge(isCompound),
		ComponentLock:             rawNBTBridge(isCompound),
		ComponentContainerLoot:    rawNBTBridge(isCompound),
		ComponentRecipes: rawNBTBridge(func(value nbt.Tag) bool {
			list, ok := value.(nbt.List)
			return ok && list.ElementType == nbt.TagString && len(list.Elements) <= maxPrealloc
		}),
		ComponentBlockState:          wireBridge(stringMapComponentNBT, writeStringMapNBT),
		ComponentWritableBookContent: wireBridge(writableBookNBT, writeWritableBookNBT),
		ComponentWrittenBookContent:  wireBridge(writtenBookNBT, writeWrittenBookNBT),
		ComponentCustomName:          textFieldBridge(ComponentCustomName, func(c *Components) **ItemNameComponent { return &c.CustomName }),
		ComponentItemName:            textFieldBridge(ComponentItemName, func(c *Components) **ItemNameComponent { return &c.ItemName }),
		ComponentLore:                wireBridge(textComponentListNBT, writeTextComponentListNBT),
		ComponentEnchantments: enchantmentsFieldBridge(ComponentEnchantments,
			func(c *Components) *map[string]int32 { return &c.Enchantments }),
		ComponentStoredEnchantments: enchantmentsFieldBridge(ComponentStoredEnchantments,
			func(c *Components) *map[string]int32 { return &c.StoredEnchantments }),
		ComponentTrim: fieldBridge(ComponentTrim,
			func(c *Components) nbt.Tag {
				if c.Trim == nil {
					return nil
				}
				return nbt.Compound{"material": nbt.String(c.Trim.Material), "pattern": nbt.String(c.Trim.Pattern)}
			},
			func(c *Components, value nbt.Tag) error {
				compound, ok := value.(nbt.Compound)
				if !ok {
					return fmt.Errorf("trim component must be a compound")
				}
				c.Trim = &ArmorTrim{Material: compound.GetString("material"), Pattern: compound.GetString("pattern")}
				if c.Trim.Material == "" || c.Trim.Pattern == "" {
					return fmt.Errorf("trim component requires material and pattern")
				}
				return nil
			}),
		ComponentSuspiciousStewEffects: wireBridge(suspiciousStewNBT, writeSuspiciousStewNBT),
		ComponentPotionContents:        wireBridge(potionContentsNBT, writePotionContentsNBT),
		ComponentPaintingVariant: stringFieldBridge(ComponentPaintingVariant, "minecraft:painting/variant",
			func(c *Components) *string { return &c.PaintingVariant },
			func(name string) bool {
				return registries.SynchronizedEntries["minecraft:painting_variant"].Get(name) >= 0
			}),
		ComponentDye:       dyeFieldBridge(ComponentDye, func(c *Components) **int32 { return &c.Dye }),
		ComponentBaseColor: dyeFieldBridge(ComponentBaseColor, func(c *Components) **int32 { return &c.BaseColor }),
		ComponentDyedColor: fieldBridge(ComponentDyedColor,
			func(c *Components) nbt.Tag {
				if c.DyedColor == nil {
					return nil
				}
				return nbt.Int(*c.DyedColor)
			},
			func(c *Components, value nbt.Tag) error {
				color := intValue(value)
				if color < 0 || color > 0xffffff {
					return fmt.Errorf("dyed_color %d is outside RGB range", color)
				}
				c.DyedColor = &color
				return nil
			}),
		ComponentBannerPatterns: fieldBridge(ComponentBannerPatterns,
			func(c *Components) nbt.Tag { return bannerPatternsNBT(c.BannerPatterns) },
			func(c *Components, value nbt.Tag) (err error) {
				c.BannerPatterns, err = bannerPatternsFromNBT(value)
				return err
			}),
		ComponentFireworkExplosion: fieldBridge(ComponentFireworkExplosion,
			func(c *Components) nbt.Tag {
				if c.FireworkExplosion == nil {
					return nil
				}
				return fireworkExplosionNBT(*c.FireworkExplosion)
			},
			func(c *Components, value nbt.Tag) error {
				explosion, err := fireworkExplosionFromNBT(value)
				c.FireworkExplosion = &explosion
				return err
			}),
		ComponentFireworks: fieldBridge(ComponentFireworks,
			func(c *Components) nbt.Tag { return fireworksNBT(c) },
			fireworksFromNBT),
	}
}

// SupportsNBTComponent reports whether a component has an NBT/wire bridge.
func SupportsNBTComponent(name string) bool {
	_, ok := nbtBridges[registries.DataComponentType.Get(name)]
	return ok
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
// Components without an NBT bridge or with an invalid payload are omitted.
func ComponentsToNBT(slot ns.Slot) nbt.Compound {
	result := nbt.Compound{}
	stack, err := FromSlot(slot)
	if err != nil {
		stack = &Stack{Components: &Components{}}
	}
	for _, raw := range slot.Components.Add {
		name := registries.DataComponentType.ByID(int32(raw.ID))
		bridge, ok := nbtBridges[int32(raw.ID)]
		if name == "" || !ok {
			continue
		}
		if value, err := bridge.toNBT(stack.Components, raw.Data); err == nil && value != nil {
			result[name] = value
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
// Components without an NBT bridge are ignored.
func ApplyNBTComponents(slot ns.Slot, patch nbt.Compound) (ns.Slot, error) {
	stack, err := FromSlot(slot)
	if err != nil {
		return ns.Slot{}, err
	}
	var removals []int32
	for key, value := range patch {
		if name, removed := strings.CutPrefix(key, "!"); removed {
			if id := registries.DataComponentType.Get(name); id >= 0 {
				removals = append(removals, id)
			}
			continue
		}
		id := registries.DataComponentType.Get(key)
		bridge, ok := nbtBridges[id]
		if !ok {
			continue
		}
		data, err := bridge.fromNBT(value)
		if err != nil {
			return ns.Slot{}, fmt.Errorf("%s: %w", key, err)
		}
		if err := applyComponent(stack.Components, id, data); err != nil {
			return ns.Slot{}, fmt.Errorf("%s: %w", key, err)
		}
	}
	for _, id := range removals {
		clearComponent(stack.Components, id)
	}
	encoded, err := stack.ToSlot()
	if err != nil {
		return ns.Slot{}, fmt.Errorf("encode item components: %w", err)
	}
	return encoded, nil
}

func nbtBool(value bool) nbt.Byte {
	if value {
		return 1
	}
	return 0
}

func stringWireBridge() nbtBridge {
	return wireBridge(
		func(r *ns.PacketBuffer) (nbt.Tag, error) {
			value, err := r.ReadString(maxStringLen)
			return nbt.String(value), err
		},
		func(w *ns.PacketBuffer, value nbt.Tag) error {
			text, ok := value.(nbt.String)
			if !ok {
				return fmt.Errorf("component must be a string")
			}
			return w.WriteString(ns.String(text))
		})
}

func swingAnimationNBT(value *SwingAnimation) nbt.Tag {
	if value == nil || value.Type < 0 || int(value.Type) >= len(swingAnimationNames) {
		return nil
	}
	return nbt.Compound{"type": nbt.String(swingAnimationNames[value.Type]), "duration": nbt.Int(value.Duration)}
}

func swingAnimationFromNBT(value nbt.Tag) (*SwingAnimation, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("swing animation must be a compound")
	}
	animation := &SwingAnimation{Type: 1, Duration: 6}
	if kind, present := compound["type"]; present {
		text, _ := kind.(nbt.String)
		animation.Type = int32(slices.Index(swingAnimationNames[:], string(text)))
		if animation.Type < 0 {
			return nil, fmt.Errorf("invalid animation type %v", kind)
		}
	}
	if duration, present := compound["duration"]; present {
		var err error
		if animation.Duration, err = componentIntNBT(duration); err != nil {
			return nil, err
		}
	}
	if animation.Duration < 0 {
		return nil, fmt.Errorf("negative animation duration")
	}
	return animation, nil
}

func fuelFromNBT(value nbt.Tag, amountField string) (ResolvableInt, ResolvableFloat, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return ResolvableInt{}, ResolvableFloat{}, fmt.Errorf("fuel component must be a compound")
	}
	amount, err := resolvableIntFromNBT(compound[amountField])
	if err != nil {
		return ResolvableInt{}, ResolvableFloat{}, err
	}
	speed, err := resolvableFloatFromNBT(compound["speed_multiplier"])
	return amount, speed, err
}

func mobVisibilityFromNBT(value nbt.Tag) (*MobVisibility, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("mob_visibility must be a compound")
	}
	types, err := holderSetFromNBT(compound["targeting_entity_types"])
	if err != nil {
		return nil, err
	}
	visibility, err := componentFloatNBT(compound["visibility"])
	if err != nil {
		return nil, err
	}
	if visibility < 0 || visibility > 10 || math.IsNaN(float64(visibility)) {
		return nil, fmt.Errorf("mob visibility outside [0,10]")
	}
	return &MobVisibility{TargetingEntityTypes: types, Visibility: visibility}, nil
}

func consumeEffectsNBT(effects []ConsumeEffect) nbt.List {
	list := nbt.List{ElementType: nbt.TagCompound}
	for _, effect := range effects {
		list.Elements = append(list.Elements, consumeEffectNBT(effect))
	}
	return list
}

func consumeEffectsFromNBT(list nbt.List) ([]ConsumeEffect, error) {
	effects := make([]ConsumeEffect, 0, len(list.Elements))
	for _, element := range list.Elements {
		effect, err := consumeEffectFromNBT(element)
		if err != nil {
			return nil, err
		}
		effects = append(effects, effect)
	}
	return effects, nil
}

func consumableNBT(consumable *Consumable) nbt.Tag {
	if consumable == nil {
		return nil
	}
	return nbt.Compound{
		"consume_seconds": nbt.Float(consumable.ConsumeSeconds), "animation": nbt.String(consumable.Animation),
		"sound":                 componentSoundNBT(consumable.Sound, consumable.SoundDirect, consumable.SoundRange),
		"has_consume_particles": nbtBool(consumable.HasParticles), "on_consume_effects": consumeEffectsNBT(consumable.OnConsumeEffects),
	}
}

func consumableFromNBT(value nbt.Tag) (*Consumable, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("consumable component must be a compound")
	}
	seconds := defaultConsumeSeconds
	if tag, present := compound["consume_seconds"]; present {
		seconds = floatValue(tag)
	}
	if seconds < 0 {
		return nil, fmt.Errorf("consume_seconds must be non-negative")
	}
	animation := cmp.Or(compound.GetString("animation"), defaultConsumeAnimation)
	if !slices.Contains(consumeAnimations, animation) {
		return nil, fmt.Errorf("unknown consume animation %q", animation)
	}
	sound := soundHolder{name: defaultConsumeSound}
	if soundTag, present := compound["sound"]; present {
		var err error
		if sound, err = soundHolderFromNBT(soundTag); err != nil {
			return nil, err
		}
	}
	effects, err := consumeEffectsFromNBT(compound.GetList("on_consume_effects"))
	if err != nil {
		return nil, err
	}
	return &Consumable{
		ConsumeSeconds: seconds, Animation: animation, Sound: sound.name, SoundDirect: sound.direct, SoundRange: sound.fixedRange,
		HasParticles: nbtBoolOr(compound, "has_consume_particles", true), OnConsumeEffects: effects,
	}, nil
}

func attackRangeNBT(attackRange *AttackRange) nbt.Tag {
	if attackRange == nil {
		return nil
	}
	return nbt.Compound{
		"min_reach": nbt.Float(attackRange.MinReach), "max_reach": nbt.Float(attackRange.MaxReach),
		"min_creative_reach": nbt.Float(attackRange.MinCreativeReach), "max_creative_reach": nbt.Float(attackRange.MaxCreativeReach),
		"hitbox_margin": nbt.Float(attackRange.HitboxMargin), "mob_factor": nbt.Float(attackRange.MobFactor),
	}
}

func attackRangeFromNBT(value nbt.Tag) (*AttackRange, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("attack_range component must be a compound")
	}
	attackRange := &AttackRange{
		MinReach: 0, MaxReach: 3, MinCreativeReach: 0,
		MaxCreativeReach: 5, HitboxMargin: 0.3, MobFactor: 1,
	}
	for key, field := range map[string]*float64{
		"min_reach": &attackRange.MinReach, "max_reach": &attackRange.MaxReach,
		"min_creative_reach": &attackRange.MinCreativeReach, "max_creative_reach": &attackRange.MaxCreativeReach,
		"hitbox_margin": &attackRange.HitboxMargin, "mob_factor": &attackRange.MobFactor,
	} {
		if tag, present := compound[key]; present {
			*field = floatValue(tag)
		}
	}
	if attackRange.MinReach < 0 || attackRange.MinReach > 64 ||
		attackRange.MaxReach < 0 || attackRange.MaxReach > 64 ||
		attackRange.MinCreativeReach < 0 || attackRange.MinCreativeReach > 64 ||
		attackRange.MaxCreativeReach < 0 || attackRange.MaxCreativeReach > 64 ||
		attackRange.HitboxMargin < 0 || attackRange.HitboxMargin > 1 ||
		attackRange.MobFactor < 0 || attackRange.MobFactor > 2 {
		return nil, fmt.Errorf("attack_range values are outside vanilla bounds")
	}
	return attackRange, nil
}

func tooltipDisplayNBT(tooltip *TooltipDisplay) nbt.Tag {
	if tooltip == nil {
		return nil
	}
	hidden := nbt.List{ElementType: nbt.TagString}
	for _, componentID := range tooltip.HiddenComponents {
		if name := registries.DataComponentType.ByID(componentID); name != "" {
			hidden.Elements = append(hidden.Elements, nbt.String(name))
		}
	}
	return nbt.Compound{"hide_tooltip": nbtBool(tooltip.HideTooltip), "hidden_components": hidden}
}

func tooltipDisplayFromNBT(value nbt.Tag) (*TooltipDisplay, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("tooltip_display component must be a compound")
	}
	hiddenList := compound.GetList("hidden_components")
	hidden := make([]int32, 0, len(hiddenList.Elements))
	for _, element := range hiddenList.Elements {
		name, ok := element.(nbt.String)
		componentID := registries.DataComponentType.Get(string(name))
		if !ok || componentID < 0 {
			return nil, fmt.Errorf("unknown hidden component %q", name)
		}
		if !slices.Contains(hidden, componentID) {
			hidden = append(hidden, componentID)
		}
	}
	return &TooltipDisplay{HideTooltip: nbtBoolOr(compound, "hide_tooltip", false), HiddenComponents: hidden}, nil
}

func writeSlotListNBT(w *ns.PacketBuffer, value nbt.Tag) error {
	list, ok := value.(nbt.List)
	if !ok || len(list.Elements) > 0 && list.ElementType != nbt.TagCompound || len(list.Elements) > maxPrealloc {
		return fmt.Errorf("item list must be a compound list of at most %d items", maxPrealloc)
	}
	w.WriteVarInt(ns.VarInt(len(list.Elements)))
	for _, element := range list.Elements {
		item, err := SlotFromNBT(element)
		if err != nil {
			return err
		}
		if err := WriteItemStackTemplate(w, item); err != nil {
			return err
		}
	}
	return nil
}

func writeContainerNBT(w *ns.PacketBuffer, value nbt.Tag) error {
	list, ok := value.(nbt.List)
	if !ok || len(list.Elements) > 0 && list.ElementType != nbt.TagCompound || len(list.Elements) > maxPrealloc {
		return fmt.Errorf("container must be a compound list of at most %d items", maxPrealloc)
	}
	slots := make(map[int]ns.Slot, len(list.Elements))
	maxSlot := -1
	for _, element := range list.Elements {
		entry, ok := element.(nbt.Compound)
		if !ok {
			return fmt.Errorf("container entry must be a compound")
		}
		slotIndex := int(entry.GetInt("slot"))
		if slotIndex < 0 || slotIndex >= maxPrealloc {
			return fmt.Errorf("container slot %d is invalid", slotIndex)
		}
		item, err := SlotFromNBT(entry["item"])
		if err != nil {
			return err
		}
		slots[slotIndex] = item
		maxSlot = max(maxSlot, slotIndex)
	}
	w.WriteVarInt(ns.VarInt(maxSlot + 1))
	for slotIndex := range maxSlot + 1 {
		item, present := slots[slotIndex]
		w.WriteBool(ns.Boolean(present))
		if present {
			if err := WriteItemStackTemplate(w, item); err != nil {
				return err
			}
		}
	}
	return nil
}

func itemTemplateNBT(r *ns.PacketBuffer) (nbt.Tag, error) {
	slot, err := ReadItemStackTemplate(r)
	if err != nil {
		return nil, err
	}
	return SlotToNBT(slot), nil
}

func writeItemTemplateNBT(w *ns.PacketBuffer, value nbt.Tag) error {
	item, err := SlotFromNBT(value)
	if err != nil {
		return err
	}
	return WriteItemStackTemplate(w, item)
}

func writeStringMapNBT(w *ns.PacketBuffer, value nbt.Tag) error {
	compound, ok := value.(nbt.Compound)
	if !ok || len(compound) > maxPrealloc {
		return fmt.Errorf("block_state component must be a compound")
	}
	w.WriteVarInt(ns.VarInt(len(compound)))
	for _, property := range slices.Sorted(maps.Keys(compound)) {
		propertyValue, ok := compound[property].(nbt.String)
		if !ok {
			return fmt.Errorf("block state property %q must be a string", property)
		}
		w.WriteString(ns.String(property))
		w.WriteString(ns.String(propertyValue))
	}
	return nil
}

func writeWritableBookNBT(w *ns.PacketBuffer, value nbt.Tag) error {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return fmt.Errorf("writable_book_content component must be a compound")
	}
	pages := compound.GetList("pages")
	if len(pages.Elements) > 100 {
		return fmt.Errorf("writable book has too many pages")
	}
	w.WriteVarInt(ns.VarInt(len(pages.Elements)))
	for _, page := range pages.Elements {
		raw, filtered, hasFiltered, err := filterableStringNBT(page, 1024)
		if err != nil {
			return err
		}
		w.WriteString(ns.String(raw))
		w.WriteBool(ns.Boolean(hasFiltered))
		if hasFiltered {
			w.WriteString(ns.String(filtered))
		}
	}
	return nil
}

func writeWrittenBookNBT(w *ns.PacketBuffer, value nbt.Tag) error {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return fmt.Errorf("written_book_content component must be a compound")
	}
	titleRaw, titleFiltered, titleHasFiltered, err := filterableStringNBT(compound["title"], 32)
	if err != nil {
		return err
	}
	generation := compound.GetInt("generation")
	if generation < 0 || generation > 3 {
		return fmt.Errorf("written book generation %d is invalid", generation)
	}
	w.WriteString(ns.String(titleRaw))
	w.WriteBool(ns.Boolean(titleHasFiltered))
	if titleHasFiltered {
		w.WriteString(ns.String(titleFiltered))
	}
	w.WriteString(ns.String(compound.GetString("author")))
	w.WriteVarInt(ns.VarInt(generation))
	pages := compound.GetList("pages")
	if len(pages.Elements) > 100 {
		return fmt.Errorf("written book has too many pages")
	}
	w.WriteVarInt(ns.VarInt(len(pages.Elements)))
	for _, page := range pages.Elements {
		raw, filtered, hasFiltered := filterableTagNBT(page)
		if raw == nil {
			return fmt.Errorf("written book page is invalid")
		}
		if err := w.WriteNBT(raw); err != nil {
			return err
		}
		w.WriteBool(ns.Boolean(hasFiltered))
		if hasFiltered {
			if err := w.WriteNBT(filtered); err != nil {
				return err
			}
		}
	}
	return w.WriteBool(ns.Boolean(nbtBoolOr(compound, "resolved", false)))
}

func writeTextComponentListNBT(w *ns.PacketBuffer, value nbt.Tag) error {
	list, ok := value.(nbt.List)
	if !ok || len(list.Elements) > maxLoreLines {
		return fmt.Errorf("lore component must be a list of at most %d components", maxLoreLines)
	}
	w.WriteVarInt(ns.VarInt(len(list.Elements)))
	for _, line := range list.Elements {
		if err := w.WriteNBT(line); err != nil {
			return err
		}
	}
	return nil
}

func suspiciousStewNBT(r *ns.PacketBuffer) (nbt.Tag, error) {
	count, err := r.ReadVarInt()
	if err != nil || count < 0 || count > maxPotionEffects {
		return nil, fmt.Errorf("invalid suspicious stew effect count %d", count)
	}
	list := nbt.List{ElementType: nbt.TagCompound}
	for range int(count) {
		id, err := r.ReadVarInt()
		if err != nil {
			return nil, err
		}
		name := registries.MobEffect.ByID(int32(id))
		if name == "" {
			return nil, fmt.Errorf("unknown suspicious stew effect %d", id)
		}
		duration, err := r.ReadVarInt()
		if err != nil {
			return nil, err
		}
		list.Elements = append(list.Elements, nbt.Compound{"id": nbt.String(name), "duration": nbt.Int(duration)})
	}
	return list, nil
}

func writeSuspiciousStewNBT(w *ns.PacketBuffer, value nbt.Tag) error {
	list, ok := value.(nbt.List)
	if !ok || len(list.Elements) > 0 && list.ElementType != nbt.TagCompound || len(list.Elements) > maxPotionEffects {
		return fmt.Errorf("suspicious_stew_effects must be a compound list")
	}
	w.WriteVarInt(ns.VarInt(len(list.Elements)))
	for _, element := range list.Elements {
		entry, ok := element.(nbt.Compound)
		if !ok {
			return fmt.Errorf("suspicious stew effect must be a compound")
		}
		effectID := registries.MobEffect.Get(entry.GetString("id"))
		if effectID < 0 {
			return fmt.Errorf("unknown suspicious stew effect %q", entry.GetString("id"))
		}
		duration := int32(160)
		if tag, present := entry["duration"]; present {
			duration = intValue(tag)
		}
		w.WriteVarInt(ns.VarInt(effectID))
		w.WriteVarInt(ns.VarInt(duration))
	}
	return nil
}

func potionContentsNBT(r *ns.PacketBuffer) (nbt.Tag, error) {
	contents, err := readPotionContents(r)
	if err != nil {
		return nil, err
	}
	value := nbt.Compound{}
	if contents.Potion != "" {
		value["potion"] = nbt.String(contents.Potion)
	}
	if contents.CustomColor != nil {
		value["custom_color"] = nbt.Int(*contents.CustomColor)
	}
	if len(contents.Effects) > 0 {
		effects := nbt.List{ElementType: nbt.TagCompound}
		for _, effect := range contents.Effects {
			effects.Elements = append(effects.Elements, statusEffectNBT(effect, true))
		}
		value["custom_effects"] = effects
	}
	if contents.CustomName != "" {
		value["custom_name"] = nbt.String(contents.CustomName)
	}
	return value, nil
}

func writePotionContentsNBT(w *ns.PacketBuffer, value nbt.Tag) error {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return fmt.Errorf("potion_contents component must be a compound")
	}
	contents := PotionContentsData{Potion: compound.GetString("potion"), CustomName: compound.GetString("custom_name")}
	if color, present := compound["custom_color"]; present {
		contents.CustomColor = int32Pointer(intValue(color))
	}
	effects := compound.GetList("custom_effects")
	if len(effects.Elements) > maxPotionEffects {
		return fmt.Errorf("too many custom potion effects")
	}
	for _, element := range effects.Elements {
		effect, err := statusEffectFromNBT(element, "")
		if err != nil {
			return err
		}
		contents.Effects = append(contents.Effects, effect)
	}
	return writePotionContents(w, contents)
}

func readTypedEntityDataNBT(r *ns.PacketBuffer, registry *registries.Registry) (nbt.Tag, error) {
	return typedEntityDataNBTFrom(r, registry)
}

func writeTypedEntityDataNBT(w *ns.PacketBuffer, value nbt.Tag, registry *registries.Registry) error {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return fmt.Errorf("typed entity data must be a compound")
	}
	typeName := compound.GetString("id")
	typeID := registry.Get(typeName)
	if typeID < 0 {
		return fmt.Errorf("unknown typed entity data type %q", typeName)
	}
	payload := maps.Clone(compound)
	delete(payload, "id")
	w.WriteVarInt(ns.VarInt(typeID))
	return w.WriteNBT(payload)
}

func bannerPatternsNBT(patterns *[]BannerPatternLayer) nbt.Tag {
	if patterns == nil {
		return nil
	}
	layers := nbt.List{ElementType: nbt.TagCompound}
	for _, layer := range *patterns {
		var pattern nbt.Tag = nbt.String(layer.Pattern)
		if layer.Pattern == "" {
			pattern = nbt.Compound{"asset_id": nbt.String(layer.AssetID), "translation_key": nbt.String(layer.TranslationKey)}
		}
		layers.Elements = append(layers.Elements, nbt.Compound{"pattern": pattern, "color": nbt.String(dyeNames[layer.Color])})
	}
	return layers
}

func bannerPatternsFromNBT(value nbt.Tag) (*[]BannerPatternLayer, error) {
	list, ok := value.(nbt.List)
	if !ok || len(list.Elements) > 0 && list.ElementType != nbt.TagCompound || len(list.Elements) > maxPrealloc {
		return nil, fmt.Errorf("banner_patterns component must be a compound list of at most %d layers", maxPrealloc)
	}
	layers := make([]BannerPatternLayer, 0, len(list.Elements))
	for _, element := range list.Elements {
		layer, ok := element.(nbt.Compound)
		if !ok {
			return nil, fmt.Errorf("banner pattern layer must be a compound")
		}
		pattern := layer.GetString("pattern")
		patternData := BannerPatternLayer{Pattern: pattern}
		if direct, ok := layer["pattern"].(nbt.Compound); ok {
			patternData.Pattern = ""
			patternData.AssetID = direct.GetString("asset_id")
			patternData.TranslationKey = direct.GetString("translation_key")
		}
		color := dyeID(layer.GetString("color"))
		if color < 0 || patternData.Pattern != "" && registries.SynchronizedEntries["minecraft:banner_pattern"].Get(patternData.Pattern) < 0 || patternData.Pattern == "" && (patternData.AssetID == "" || patternData.TranslationKey == "") {
			return nil, fmt.Errorf("invalid banner pattern %q color %q", pattern, layer.GetString("color"))
		}
		patternData.Color = color
		layers = append(layers, patternData)
	}
	return &layers, nil
}

func fireworksNBT(c *Components) nbt.Tag {
	if c.Fireworks == nil {
		return nil
	}
	explosions := nbt.List{ElementType: nbt.TagCompound}
	for _, explosion := range c.FireworkExplosions {
		explosions.Elements = append(explosions.Elements, fireworkExplosionNBT(explosion))
	}
	return nbt.Compound{"flight_duration": nbt.Int(c.Fireworks.FlightDuration), "explosions": explosions}
}

func fireworksFromNBT(c *Components, value nbt.Tag) error {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return fmt.Errorf("fireworks component must be a compound")
	}
	flight := intValue(compound["flight_duration"])
	if flight < 0 || flight > 255 {
		return fmt.Errorf("firework flight duration %d is invalid", flight)
	}
	list := compound.GetList("explosions")
	if len(list.Elements) > maxPrealloc {
		return fmt.Errorf("too many firework explosions")
	}
	explosions := make([]FireworkExplosion, 0, len(list.Elements))
	for _, element := range list.Elements {
		explosion, err := fireworkExplosionFromNBT(element)
		if err != nil {
			return err
		}
		explosions = append(explosions, explosion)
	}
	c.Fireworks = &Fireworks{FlightDuration: flight}
	c.FireworkExplosions = explosions
	return nil
}

func potDecorationsNBT(decorations *PotDecorations) nbt.Tag {
	if decorations == nil {
		return nil
	}
	result := nbt.Compound{}
	for name, stack := range map[string]*Stack{
		"back": decorations.Back, "left": decorations.Left, "right": decorations.Right, "front": decorations.Front,
	} {
		if stack != nil {
			if slot, err := stack.ToSlot(); err == nil {
				result[name] = SlotToNBT(slot)
			}
		}
	}
	return result
}

func equippableNBT(component *Equippable) nbt.Tag {
	if component == nil {
		return nil
	}
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
		sound, err := soundHolderFromNBT(soundTag)
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
		sound, err := soundHolderFromNBT(soundTag)
		if err != nil {
			return nil, err
		}
		component.ShearingSound, component.ShearingSoundDirect, component.ShearingSoundRange = sound.name, sound.direct, sound.fixedRange
	}
	return component, nil
}

func blocksAttacksNBT(component *BlocksAttacks) nbt.Tag {
	if component == nil {
		return nil
	}
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
		sound, err := soundHolderFromNBT(soundTag)
		if err != nil {
			return nil, err
		}
		component.BlockSound, component.BlockSoundDirect, component.BlockSoundRange = sound.name, sound.direct, sound.fixedRange
	}
	if soundTag, present := compound["disabled_sound"]; present {
		sound, err := soundHolderFromNBT(soundTag)
		if err != nil {
			return nil, err
		}
		component.DisabledSound, component.DisabledSoundDirect, component.DisabledSoundRange = sound.name, sound.direct, sound.fixedRange
	}
	return component, nil
}

func toolNBT(tool *Tool) nbt.Tag {
	if tool == nil {
		return nil
	}
	rules := nbt.List{ElementType: nbt.TagCompound}
	for _, rule := range tool.Rules {
		entry := nbt.Compound{"blocks": holderSetNBT(rule.Blocks)}
		if rule.HasSpeed {
			entry["speed"] = nbt.Float(rule.Speed)
		}
		if rule.HasCorrectForDrops {
			entry["correct_for_drops"] = nbtBool(rule.CorrectForDrops)
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
	if _, ok := singleTagRef(entries); ok {
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
	encoded := ns.NewWriter()
	if value == nil || encoded.WriteNBT(value) != nil {
		return nil
	}
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
			Slot:      normalizedEquipmentSlotGroup(entry.GetString("slot")),
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
		if !slices.Contains(attributeModifierOperations, modifier.Operation) {
			return nil, fmt.Errorf("unknown attribute modifier operation %q", modifier.Operation)
		}
		if !slices.Contains(equipmentSlotGroups, modifier.Slot) {
			return nil, fmt.Errorf("unknown equipment slot group %q", modifier.Slot)
		}
		if displayTag, present := entry["display"]; present {
			display, ok := displayTag.(nbt.Compound)
			if !ok {
				return nil, fmt.Errorf("attribute modifier display must be a compound")
			}
			modifier.Display.Type = display.GetString("type")
			if !slices.Contains(attributeModifierDisplayTypes, modifier.Display.Type) {
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
	name := registries.SynchronizedEntries[registry].HolderName(int32(holderID))
	if name == "" {
		return "", fmt.Errorf("unknown %s holder %d", registry, holderID)
	}
	return name, nil
}

func writeSynchronizedHolderReference(writer *ns.PacketBuffer, registry, name string) error {
	holderID := registries.SynchronizedEntries[registry].HolderID(name)
	if holderID < 0 {
		return fmt.Errorf("unknown %s entry %q", registry, name)
	}
	writer.WriteVarInt(ns.VarInt(holderID))
	return nil
}

func soundEventHolderNBT(reader *ns.PacketBuffer) (nbt.Tag, error) {
	sound, err := readSoundHolder(reader)
	if err != nil {
		return nil, err
	}
	return componentSoundNBT(sound.name, sound.direct, sound.fixedRange), nil
}

func writeSoundEventHolderNBT(writer *ns.PacketBuffer, value nbt.Tag) error {
	if name, ok := value.(nbt.String); ok {
		holderID := registries.SoundEvent.HolderID(string(name))
		if holderID < 0 {
			return fmt.Errorf("unknown sound event %q", name)
		}
		writer.WriteVarInt(ns.VarInt(holderID))
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
	description, err := reader.ReadNBT()
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
	return writer.WriteNBT(description)
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
	description, err := reader.ReadNBT()
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
	if err := writer.WriteNBT(description); err != nil {
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
	description, err := reader.ReadNBT()
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
	keys := slices.Sorted(maps.Keys(overrides))
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
	return writer.WriteNBT(description)
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

func writableBookNBT(reader *ns.PacketBuffer) (nbt.Tag, error) {
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

func writtenBookNBT(reader *ns.PacketBuffer) (nbt.Tag, error) {
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
		raw, err := reader.ReadNBT()
		if err != nil {
			return nil, err
		}
		hasFiltered, err := reader.ReadBool()
		if err != nil {
			return nil, err
		}
		page := pageData{raw: raw, hasFiltered: bool(hasFiltered)}
		if hasFiltered {
			page.filtered, err = reader.ReadNBT()
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
		"pages": pageList, "resolved": nbtBool(bool(resolved)),
	}, nil
}

func textComponentListNBT(reader *ns.PacketBuffer) (nbt.Tag, error) {
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > maxLoreLines {
		return nil, fmt.Errorf("invalid text component list length %d", count)
	}
	values := make([]nbt.Tag, 0, int(count))
	elementType := byte(nbt.TagEnd)
	for range int(count) {
		value, err := reader.ReadNBT()
		if err != nil {
			return nil, err
		}
		if elementType == nbt.TagEnd {
			elementType = value.ID()
		} else if value.ID() != elementType {
			return nil, fmt.Errorf("mixed text component tag types cannot be represented in NBT list")
		}
		values = append(values, value)
	}
	if elementType == nbt.TagEnd {
		elementType = nbt.TagCompound
	}
	return nbt.List{ElementType: elementType, Elements: values}, nil
}

func stringMapComponentNBT(reader *ns.PacketBuffer) (nbt.Tag, error) {
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > maxPrealloc {
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
	return ns.NewReader(data).ReadNBT()
}

func typedEntityDataNBT(data []byte, registry *registries.Registry) (nbt.Compound, error) {
	return typedEntityDataNBTFrom(ns.NewReader(data), registry)
}

func typedEntityDataNBTFrom(reader *ns.PacketBuffer, registry *registries.Registry) (nbt.Compound, error) {
	typeID, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	}
	typeName := registry.ByID(int32(typeID))
	if typeName == "" {
		return nil, fmt.Errorf("unknown typed entity data id %d", typeID)
	}
	tag, err := reader.ReadNBT()
	if err != nil {
		return nil, err
	}
	compound, ok := tag.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("typed entity data payload must be a compound")
	}
	result := make(nbt.Compound, len(compound)+1)
	maps.Copy(result, compound)
	result["id"] = nbt.String(typeName)
	return result, nil
}

// SlotToNBT converts a non-empty slot to vanilla's item-stack compound.
func SlotToNBT(slot ns.Slot) nbt.Compound {
	item := nbt.Compound{"id": nbt.String(Name(int32(slot.ItemID))), "count": nbt.Int(slot.Count)}
	if components := ComponentsToNBT(slot); len(components) > 0 {
		item["components"] = components
	}
	return item
}

func slotListComponentNBT(reader *ns.PacketBuffer, optional bool) (nbt.Tag, error) {
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > maxPrealloc {
		return nil, fmt.Errorf("invalid item component list length %d", count)
	}
	result := nbt.List{ElementType: nbt.TagCompound}
	for slotIndex := range int(count) {
		if optional {
			hasItem, err := reader.ReadBool()
			if err != nil {
				return nil, err
			}
			if !hasItem {
				continue
			}
		}
		slot, err := ReadItemStackTemplate(reader)
		if err != nil {
			return nil, err
		}
		item := SlotToNBT(slot)
		if optional {
			result.Elements = append(result.Elements, nbt.Compound{"slot": nbt.Int(slotIndex), "item": item})
		} else {
			result.Elements = append(result.Elements, item)
		}
	}
	return result, nil
}

// SlotFromNBT decodes vanilla's item-stack compound.
func SlotFromNBT(value nbt.Tag) (ns.Slot, error) {
	item, ok := value.(nbt.Compound)
	if !ok {
		return ns.Slot{}, fmt.Errorf("item component list entry must be a compound")
	}
	id := ID(item.GetString("id"))
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

func soundHolderFromNBT(value nbt.Tag) (soundHolder, error) {
	writer := ns.NewWriter()
	if err := writeSoundEventHolderNBT(writer, value); err != nil {
		return soundHolder{}, err
	}
	return readSoundHolder(ns.NewReader(writer.Bytes()))
}

// statusEffectNBT renders a vanilla MobEffectInstance; the amplifier is an
// unsigned byte (MobEffectInstance.Details uses ExtraCodecs.UNSIGNED_BYTE).
func statusEffectNBT(effect StatusEffect, includeID bool) nbt.Compound {
	result := nbt.Compound{
		"amplifier": nbt.Byte(int8(uint8(effect.Amplifier))), "duration": nbt.Int(effect.Duration),
		"ambient": nbtBool(effect.Ambient), "show_particles": nbtBool(effect.ShowParticles), "show_icon": nbtBool(effect.ShowIcon),
	}
	if includeID {
		result["id"] = nbt.String(effect.ID)
	}
	if effect.Hidden != nil {
		result["hidden_effect"] = statusEffectNBT(*effect.Hidden, false)
	}
	return result
}

// statusEffectFromNBT parses a vanilla MobEffectInstance. Hidden effects
// inherit the outer effect ID; the amplifier accepts any integer tag and is
// read as an unsigned byte.
func statusEffectFromNBT(value nbt.Tag, inheritedID string) (StatusEffect, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return StatusEffect{}, fmt.Errorf("status effect must be a compound")
	}
	effectID := cmp.Or(compound.GetString("id"), inheritedID)
	if registries.MobEffect.Get(effectID) < 0 {
		return StatusEffect{}, fmt.Errorf("unknown status effect %q", effectID)
	}
	status := StatusEffect{
		ID: effectID, Amplifier: intValue(compound["amplifier"]) & 0xFF, Duration: intValue(compound["duration"]),
		Ambient: nbtBoolOr(compound, "ambient", false), ShowParticles: nbtBoolOr(compound, "show_particles", true),
	}
	status.ShowIcon = nbtBoolOr(compound, "show_icon", status.ShowParticles)
	if hidden, present := compound["hidden_effect"]; present {
		parsed, err := statusEffectFromNBT(hidden, effectID)
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
			effects.Elements = append(effects.Elements, statusEffectNBT(status, true))
		}
		result["effects"] = effects
		result["probability"] = nbt.Float(effect.Probability)
	case "minecraft:remove_effects":
		result["effects"] = holderSetNBT(effect.EffectIDs)
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
			status, err := statusEffectFromNBT(element, "")
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
		sound, err := soundHolderFromNBT(compound["sound"])
		if err != nil {
			return ConsumeEffect{}, err
		}
		effect.Sound, effect.SoundDirect, effect.SoundRange = sound.name, sound.direct, sound.fixedRange
	default:
		return ConsumeEffect{}, fmt.Errorf("unknown consume effect type %q", typeName)
	}
	return effect, nil
}

func nbtBoolOr(compound nbt.Compound, key string, fallback bool) bool {
	value, present := compound[key]
	if !present {
		return fallback
	}
	return intValue(value) != 0
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
		"fade_colors": nbt.IntArray(explosion.FadeColors), "has_trail": nbtBool(explosion.HasTrail),
		"has_twinkle": nbtBool(explosion.HasTwinkle),
	}
}

func fireworkExplosionFromNBT(value nbt.Tag) (FireworkExplosion, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return FireworkExplosion{}, fmt.Errorf("firework explosion must be a compound")
	}
	shape := int32(slices.Index(fireworkShapeNames[:], compound.GetString("shape")))
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
		Shape: shape, Colors: slices.Clone([]int32(colors)), FadeColors: slices.Clone([]int32(fadeColors)),
		HasTrail: compound.GetByte("has_trail") != 0, HasTwinkle: compound.GetByte("has_twinkle") != 0,
	}, nil
}

var dyeNames = [16]string{
	"white", "orange", "magenta", "light_blue", "yellow", "lime", "pink", "gray",
	"light_gray", "cyan", "purple", "blue", "brown", "green", "red", "black",
}

func dyeID(name string) int32 { return int32(slices.Index(dyeNames[:], name)) }
