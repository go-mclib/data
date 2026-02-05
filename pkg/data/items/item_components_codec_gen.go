// Code generated for Minecraft 1.21.11 (Protocol 774); DO NOT EDIT.

package items

// Auto-generated codec registrations.

func init() {
	// Simple VarInt codecs
	RegisterCodec(ComponentDamage, &varIntCodec{
		get: func(c *Components) int32 { return c.Damage },
		set: func(c *Components, v int32) { c.Damage = v },
	})
	RegisterCodec(ComponentMapColor, &varIntCodec{
		get: func(c *Components) int32 { return c.MapColor },
		set: func(c *Components, v int32) { c.MapColor = v },
	})
	RegisterCodec(ComponentMaxDamage, &varIntCodec{
		get: func(c *Components) int32 { return c.MaxDamage },
		set: func(c *Components, v int32) { c.MaxDamage = v },
	})
	RegisterCodec(ComponentMaxStackSize, &varIntCodec{
		get: func(c *Components) int32 { return c.MaxStackSize },
		set: func(c *Components, v int32) { c.MaxStackSize = v },
	})
	RegisterCodec(ComponentOminousBottleAmplifier, &varIntCodec{
		get: func(c *Components) int32 { return c.OminousBottleAmplifier },
		set: func(c *Components, v int32) { c.OminousBottleAmplifier = v },
	})
	RegisterCodec(ComponentRepairCost, &varIntCodec{
		get: func(c *Components) int32 { return c.RepairCost },
		set: func(c *Components, v int32) { c.RepairCost = v },
	})

	// Simple Float32 codecs
	RegisterCodec(ComponentMinimumAttackCharge, &float32Codec{
		get: func(c *Components) float64 { return c.MinimumAttackCharge },
		set: func(c *Components, v float64) { c.MinimumAttackCharge = v },
	})
	RegisterCodec(ComponentPotionDurationScale, &float32Codec{
		get: func(c *Components) float64 { return c.PotionDurationScale },
		set: func(c *Components, v float64) { c.PotionDurationScale = v },
	})

	// Simple String/Identifier codecs
	RegisterCodec(ComponentBreakSound, &stringCodec{
		get: func(c *Components) string { return c.BreakSound },
		set: func(c *Components, v string) { c.BreakSound = v },
	})
	RegisterCodec(ComponentDamageType, &stringCodec{
		get: func(c *Components) string { return c.DamageType },
		set: func(c *Components, v string) { c.DamageType = v },
	})
	RegisterCodec(ComponentInstrument, &stringCodec{
		get: func(c *Components) string { return c.Instrument },
		set: func(c *Components, v string) { c.Instrument = v },
	})
	RegisterCodec(ComponentItemModel, &stringCodec{
		get: func(c *Components) string { return c.ItemModel },
		set: func(c *Components, v string) { c.ItemModel = v },
	})
	RegisterCodec(ComponentJukeboxPlayable, &stringCodec{
		get: func(c *Components) string { return c.JukeboxPlayable },
		set: func(c *Components, v string) { c.JukeboxPlayable = v },
	})
	RegisterCodec(ComponentProvidesBannerPatterns, &stringCodec{
		get: func(c *Components) string { return c.ProvidesBannerPatterns },
		set: func(c *Components, v string) { c.ProvidesBannerPatterns = v },
	})
	RegisterCodec(ComponentProvidesTrimMaterial, &stringCodec{
		get: func(c *Components) string { return c.ProvidesTrimMaterial },
		set: func(c *Components, v string) { c.ProvidesTrimMaterial = v },
	})

	// Empty marker codecs (bool flags)
	RegisterCodec(ComponentGlider, &emptyMarkerCodec{
		get: func(c *Components) bool { return c.Glider },
		set: func(c *Components, v bool) { c.Glider = v },
	})
	RegisterCodec(ComponentUnbreakable, &emptyMarkerCodec{
		get: func(c *Components) bool { return c.Unbreakable },
		set: func(c *Components, v bool) { c.Unbreakable = v },
	})

	// VarInt passthrough
	for _, id := range []int32{
		ComponentBaseColor,
		ComponentMapId,
		ComponentMapPostProcessing,
		ComponentSwingAnimation,
	} {
		registerVarIntPassthrough(id)
	}

	// Bool passthrough
	for _, id := range []int32{
		ComponentEnchantmentGlintOverride,
	} {
		registerBoolPassthrough(id)
	}

	// String passthrough
	for _, id := range []int32{
		ComponentNoteBlockSound,
		ComponentTooltipStyle,
	} {
		registerStringPassthrough(id)
	}

	// Empty passthrough
	for _, id := range []int32{
		ComponentCreativeSlotLock,
		ComponentIntangibleProjectile,
	} {
		registerEmptyPassthrough(id)
	}

	// Int32 passthrough
	for _, id := range []int32{
		ComponentDyedColor,
	} {
		registerInt32Passthrough(id)
	}

	// NBT passthrough
	for _, id := range []int32{
		ComponentBlockEntityData,
		ComponentBucketEntityData,
		ComponentContainerLoot,
		ComponentCustomData,
		ComponentDebugStickState,
		ComponentEntityData,
		ComponentLock,
		ComponentMapDecorations,
	} {
		registerNBTPassthrough(id)
	}

	// HolderSet passthrough
	for _, id := range []int32{
		ComponentDamageResistant,
		ComponentRepairable,
	} {
		registerHolderSetPassthrough(id)
	}

	// SlotList passthrough
	for _, id := range []int32{
		ComponentBundleContents,
		ComponentChargedProjectiles,
		ComponentContainer,
	} {
		registerSlotListPassthrough(id)
	}

	// Slot passthrough
	for _, id := range []int32{
		ComponentUseRemainder,
	} {
		registerSlotPassthrough(id)
	}

	// Entity variant (VarInt) passthrough
	for _, id := range []int32{
		ComponentAxolotlVariant,
		ComponentCatCollar,
		ComponentCatVariant,
		ComponentChickenVariant,
		ComponentCowVariant,
		ComponentFoxVariant,
		ComponentFrogVariant,
		ComponentHorseVariant,
		ComponentLlamaVariant,
		ComponentMooshroomVariant,
		ComponentPaintingVariant,
		ComponentParrotVariant,
		ComponentPigVariant,
		ComponentRabbitVariant,
		ComponentSalmonSize,
		ComponentSheepColor,
		ComponentShulkerColor,
		ComponentTropicalFishBaseColor,
		ComponentTropicalFishPattern,
		ComponentTropicalFishPatternColor,
		ComponentVillagerVariant,
		ComponentWolfCollar,
		ComponentWolfSoundVariant,
		ComponentWolfVariant,
		ComponentZombieNautilusVariant,
	} {
		registerVarIntPassthrough(id)
	}

}
