package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/go-mclib/nbt"
)

func generateItems(items map[string]ItemJSON, registries map[string]RegistryJSON, metadataPath, outPath string) {
	itemRegistry := registries["minecraft:item"]
	metadata := loadJSON[ComponentMetadataFile](metadataPath)

	var sb, potDecorations strings.Builder
	sb.WriteString(generatedFileHeader("items"))

	// generate default components data
	sb.WriteString("// defaultComponents maps item IDs to their default components.\n")
	sb.WriteString("var defaultComponents = map[int32]*Components{\n")

	for _, itemName := range sortedKeys(items) {
		item := items[itemName]
		itemID := itemRegistry.Entries[itemName].ProtocolID

		fmt.Fprintf(&sb, "\t%d: { // %s\n", itemID, itemName)
		generateComponentsLiteral(&sb, item.Components, metadata, "\t\t")
		sb.WriteString("\t},\n")

		if value, ok := item.Components["minecraft:pot_decorations"].(map[string]any); ok && len(value) > 0 {
			data := strings.TrimPrefix(fmt.Sprintf("%#v", componentDefaultNBT(value)), "[]byte")
			fmt.Fprintf(&potDecorations, "\t%d: %s, // %s\n", itemID, data, itemName)
		}
	}
	sb.WriteString("}\n\n")

	sb.WriteString("// defaultPotDecorationData holds the NBT of non-empty default pot decorations by item ID.\n")
	sb.WriteString("var defaultPotDecorationData = map[int32][]byte{\n")
	sb.WriteString(potDecorations.String())
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
}

// componentLiteralWriter writes the Components field literal for one vanilla default component value.
type componentLiteralWriter func(sb *strings.Builder, indent, field string, value any)

// componentLiteralWriters lists the components emitted into default item literals; the Go
// field name comes from the component's goField in component_metadata.include.json.
var componentLiteralWriters = map[string]componentLiteralWriter{
	"minecraft:attack_animation":         writeSwingAnimationLiteral,
	"minecraft:attack_range":             writeAttackRangeLiteral,
	"minecraft:attribute_modifiers":      writeAttributeModifiersLiteral,
	"minecraft:banner_patterns":          writeBannerPatternsLiteral,
	"minecraft:block_transformer":        writeStringLiteral,
	"minecraft:blocks_attacks":           writeBlocksAttacksLiteral,
	"minecraft:break_sound":              writeStringLiteral,
	"minecraft:brewing_fuel":             writeBrewingFuelLiteral,
	"minecraft:compostable":              writeCompostableLiteral,
	"minecraft:consumable":               writeConsumableLiteral,
	"minecraft:cooking_fuel":             writeCookingFuelLiteral,
	"minecraft:cushion/color":            writeDyeColorLiteral,
	"minecraft:damage":                   writeInt32Literal,
	"minecraft:damage_resistant":         writeDamageResistantLiteral,
	"minecraft:damage_type":              writeStringLiteral,
	"minecraft:death_protection":         writeDeathProtectionLiteral,
	"minecraft:dye":                      writeDyeColorLiteral,
	"minecraft:enchantable":              writeEnchantableLiteral,
	"minecraft:entity_data":              writeEntityDataLiteral,
	"minecraft:equippable":               writeEquippableLiteral,
	"minecraft:fireworks":                writeFireworksLiteral,
	"minecraft:food":                     writeFoodLiteral,
	"minecraft:glider":                   writeMarkerLiteral,
	"minecraft:instrument":               writeStringLiteral,
	"minecraft:interact_animation":       writeSwingAnimationLiteral,
	"minecraft:item_model":               writeStringLiteral,
	"minecraft:item_name":                writeItemNameLiteral,
	"minecraft:jukebox_playable":         writeStringLiteral,
	"minecraft:max_damage":               writeInt32Literal,
	"minecraft:max_stack_size":           writeInt32Literal,
	"minecraft:minimum_attack_charge":    writeFloatLiteral,
	"minecraft:mob_visibility":           writeMobVisibilityLiteral,
	"minecraft:ominous_bottle_amplifier": writeInt32Literal,
	"minecraft:pot_decorations":          writePotDecorationsLiteral,
	"minecraft:potion_duration_scale":    writeFloatLiteral,
	"minecraft:provides_banner_patterns": writeStringLiteral,
	"minecraft:provides_pottery_pattern": writeStringLiteral,
	"minecraft:provides_trim_material":   writeStringLiteral,
	"minecraft:rarity":                   writeStringLiteral,
	"minecraft:repair_cost":              writeInt32Literal,
	"minecraft:repairable":               writeRepairableLiteral,
	"minecraft:sign_text_back":           writeSignTextLiteral,
	"minecraft:sign_text_front":          writeSignTextLiteral,
	"minecraft:tool":                     writeToolLiteral,
	"minecraft:use_cooldown":             writeUseCooldownLiteral,
	"minecraft:use_effects":              writeUseEffectsLiteral,
	"minecraft:use_remainder":            writeUseRemainderLiteral,
	"minecraft:villager_food":            writeVillagerFoodLiteral,
	"minecraft:waxed":                    writeMarkerLiteral,
	"minecraft:weapon":                   writeWeaponLiteral,
	"minecraft:writable_book_content":    writeWritableBookContentLiteral,
}

func generateComponentsLiteral(sb *strings.Builder, components map[string]any, metadata ComponentMetadataFile, indent string) {
	for _, key := range sortedKeys(components) {
		write, ok := componentLiteralWriters[key]
		if !ok {
			continue
		}
		field := metadata.Components[key].GoField
		if field == "" {
			panic("component " + key + " has a literal writer but no goField metadata")
		}
		write(sb, indent, field, components[key])
	}
}

func writeStringLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: %q,\n", indent, field, value.(string))
}

func writeInt32Literal(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: %d,\n", indent, field, int32(value.(float64)))
}

func writeFloatLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: %v,\n", indent, field, value.(float64))
}

func writeMarkerLiteral(sb *strings.Builder, indent, field string, _ any) {
	fmt.Fprintf(sb, "%s%s: true,\n", indent, field)
}

func writeDyeColorLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: int32Pointer(%d),\n", indent, field, dyeColorIndex(value.(string)))
}

func writeSwingAnimationLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	kind, duration := int32(1), int32(6)
	if name, ok := m["type"].(string); ok {
		switch name {
		case "none":
			kind = 0
		case "whack":
			kind = 1
		case "stab":
			kind = 2
		default:
			panic("unknown swing animation " + name)
		}
	}
	if v, ok := m["duration"].(float64); ok {
		duration = int32(v)
	}
	fmt.Fprintf(sb, "%s%s: &SwingAnimation{Type: %d, Duration: %d},\n", indent, field, kind, duration)
}

func writeVillagerFoodLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &VillagerFood{Nutrition: %d},\n", indent, field, int32(value.(map[string]any)["nutrition"].(float64)))
}

func writeCompostableLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &Compostable{Layers: %s},\n", indent, field, resolvableNumberLiteral(value.(map[string]any)["layers"], false))
}

func writeCookingFuelLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	fmt.Fprintf(sb, "%s%s: &CookingFuel{BurnTime: %s, SpeedMultiplier: %s},\n", indent, field, resolvableNumberLiteral(m["burn_time"], false), resolvableNumberLiteral(m["speed_multiplier"], true))
}

func writeBrewingFuelLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	fmt.Fprintf(sb, "%s%s: &BrewingFuel{Uses: %s, SpeedMultiplier: %s},\n", indent, field, resolvableNumberLiteral(m["uses"], false), resolvableNumberLiteral(m["speed_multiplier"], true))
}

func writeMobVisibilityLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	fmt.Fprintf(sb, "%s%s: &MobVisibility{Visibility: %v, TargetingEntityTypes: []string{", indent, field, m["visibility"])
	generateStringValues(sb, m["targeting_entity_types"])
	sb.WriteString("}},\n")
}

func writePotDecorationsLiteral(sb *strings.Builder, indent, field string, _ any) {
	fmt.Fprintf(sb, "%s%s: &PotDecorations{},\n", indent, field)
}

func writeBannerPatternsLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &[]BannerPatternLayer{", indent, field)
	for _, raw := range value.([]any) {
		layer := raw.(map[string]any)
		sb.WriteString("{")
		switch pattern := layer["pattern"].(type) {
		case string:
			fmt.Fprintf(sb, "Pattern: %q, ", pattern)
		case map[string]any:
			fmt.Fprintf(sb, "AssetID: %q, TranslationKey: %q, ", pattern["asset_id"], pattern["translation_key"])
		default:
			panic(fmt.Sprintf("unsupported banner pattern %v", pattern))
		}
		fmt.Fprintf(sb, "Color: %d}, ", dyeColorIndex(layer["color"].(string)))
	}
	sb.WriteString("},\n")
}

func writeUseEffectsLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	// vanilla UseEffects.DEFAULT
	canSprint, interactVibrations, speedMultiplier := false, true, 0.2
	if v, ok := m["can_sprint"].(bool); ok {
		canSprint = v
	}
	if v, ok := m["interact_vibrations"].(bool); ok {
		interactVibrations = v
	}
	if v, ok := m["speed_multiplier"].(float64); ok {
		speedMultiplier = v
	}
	fmt.Fprintf(sb, "%s%s: &UseEffects{CanSprint: %t, InteractVibrations: %t, SpeedMultiplier: %v},\n", indent, field, canSprint, interactVibrations, speedMultiplier)
}

func writeFoodLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	fmt.Fprintf(sb, "%s%s: &Food{\n", indent, field)
	if n, ok := m["nutrition"].(float64); ok {
		fmt.Fprintf(sb, "%s\tNutrition: %d,\n", indent, int32(n))
	}
	if s, ok := m["saturation"].(float64); ok {
		fmt.Fprintf(sb, "%s\tSaturation: %v,\n", indent, s)
	}
	if canAlwaysEat, ok := m["can_always_eat"].(bool); ok && canAlwaysEat {
		fmt.Fprintf(sb, "%s\tCanAlwaysEat: true,\n", indent)
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

// writeConsumableLiteral refers to the items package defaults when vanilla omits a field.
func writeConsumableLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	seconds := "defaultConsumeSeconds"
	if configured, ok := m["consume_seconds"].(float64); ok {
		seconds = fmt.Sprintf("%v", configured)
	}
	animation := "defaultConsumeAnimation"
	if configured, ok := m["animation"].(string); ok {
		animation = fmt.Sprintf("%q", configured)
	}
	sound := "defaultConsumeSound"
	if configured, ok := m["sound"].(string); ok {
		sound = fmt.Sprintf("%q", configured)
	}
	hasParticles := true
	if configured, ok := m["has_consume_particles"].(bool); ok {
		hasParticles = configured
	}
	fmt.Fprintf(sb, "%s%s: &Consumable{\n", indent, field)
	fmt.Fprintf(sb, "%s\tConsumeSeconds: %s,\n", indent, seconds)
	fmt.Fprintf(sb, "%s\tAnimation: %s,\n", indent, animation)
	fmt.Fprintf(sb, "%s\tSound: %s,\n", indent, sound)
	fmt.Fprintf(sb, "%s\tHasParticles: %t,\n", indent, hasParticles)
	writeConsumeEffects(sb, indent+"\t", "OnConsumeEffects", m["on_consume_effects"])
	fmt.Fprintf(sb, "%s},\n", indent)
}

func writeToolLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	defaultMiningSpeed := 1.0
	if v, ok := m["default_mining_speed"].(float64); ok {
		defaultMiningSpeed = v
	}
	damagePerBlock := int32(1)
	if v, ok := m["damage_per_block"].(float64); ok {
		damagePerBlock = int32(v)
	}
	canDestroyBlocksInCreative := true
	if v, ok := m["can_destroy_blocks_in_creative"].(bool); ok {
		canDestroyBlocksInCreative = v
	}
	fmt.Fprintf(sb, "%s%s: &Tool{\n", indent, field)
	fmt.Fprintf(sb, "%s\tDefaultMiningSpeed: %v, DamagePerBlock: %d, CanDestroyBlocksInCreative: %v,\n",
		indent, defaultMiningSpeed, damagePerBlock, canDestroyBlocksInCreative)
	if rules, ok := m["rules"].([]any); ok && len(rules) > 0 {
		fmt.Fprintf(sb, "%s\tRules: []ToolRule{\n", indent)
		for _, raw := range rules {
			rule := raw.(map[string]any)
			fmt.Fprintf(sb, "%s\t\t{", indent)
			if b, ok := rule["blocks"].(string); ok {
				fmt.Fprintf(sb, "Blocks: []string{%q}, ", b)
			}
			if s, ok := rule["speed"].(float64); ok {
				fmt.Fprintf(sb, "Speed: %v, HasSpeed: true, ", s)
			}
			if c, ok := rule["correct_for_drops"].(bool); ok {
				fmt.Fprintf(sb, "CorrectForDrops: %v, HasCorrectForDrops: true", c)
			}
			sb.WriteString("},\n")
		}
		fmt.Fprintf(sb, "%s\t},\n", indent)
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

func writeBlocksAttacksLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	blockDelay, _ := m["block_delay_seconds"].(float64)
	disableScale := 1.0
	if v, ok := m["disable_cooldown_scale"].(float64); ok {
		disableScale = v
	}
	fmt.Fprintf(sb, "%s%s: &BlocksAttacks{\n", indent, field)
	fmt.Fprintf(sb, "%s\tBlockDelaySeconds: %v,\n", indent, blockDelay)
	fmt.Fprintf(sb, "%s\tDisableCooldownScale: %v,\n", indent, disableScale)
	writeDamageReductions(sb, indent+"\t", m["damage_reductions"])
	threshold, damageBase, damageFactor := 1.0, 0.0, 1.0
	if damage, ok := m["item_damage"].(map[string]any); ok {
		if v, ok := damage["threshold"].(float64); ok {
			threshold = v
		}
		if v, ok := damage["base"].(float64); ok {
			damageBase = v
		}
		if v, ok := damage["factor"].(float64); ok {
			damageFactor = v
		}
	}
	fmt.Fprintf(sb, "%s\tItemDamage: DamageSpec{Threshold: %v, Base: %v, Factor: %v},\n", indent, threshold, damageBase, damageFactor)
	if bypassedBy, ok := m["bypassed_by"]; ok {
		fmt.Fprintf(sb, "%s\tBypassedBy: []string{", indent)
		generateStringValues(sb, bypassedBy)
		sb.WriteString("},\n")
	}
	if sound, ok := m["block_sound"].(string); ok {
		fmt.Fprintf(sb, "%s\tBlockSound: %q,\n", indent, sound)
	}
	if sound, ok := m["disabled_sound"].(string); ok {
		fmt.Fprintf(sb, "%s\tDisabledSound: %q,\n", indent, sound)
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

func writeDamageReductions(sb *strings.Builder, indent string, value any) {
	fmt.Fprintf(sb, "%sDamageReductions: []DamageReduction{\n", indent)
	reductions, ok := value.([]any)
	if !ok {
		fmt.Fprintf(sb, "%s\t{HorizontalBlockingAngle: 90, Factor: 1},\n", indent)
	}
	for _, raw := range reductions {
		reduction := raw.(map[string]any)
		angle, base, factor := 90.0, 0.0, 1.0
		if v, ok := reduction["horizontal_blocking_angle"].(float64); ok {
			angle = v
		}
		if v, ok := reduction["base"].(float64); ok {
			base = v
		}
		if v, ok := reduction["factor"].(float64); ok {
			factor = v
		}
		fmt.Fprintf(sb, "%s\t{HorizontalBlockingAngle: %v, Base: %v, Factor: %v", indent, angle, base, factor)
		generateHolderSetLiteral(sb, reduction["type"])
		sb.WriteString("},\n")
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

func writeWeaponLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	fmt.Fprintf(sb, "%s%s: &Weapon{\n", indent, field)
	if d, ok := m["disable_blocking_for_seconds"].(float64); ok {
		fmt.Fprintf(sb, "%s\tDisableBlockingForSeconds: %v,\n", indent, d)
	}
	if i, ok := m["item_damage_per_attack"].(float64); ok {
		fmt.Fprintf(sb, "%s\tItemDamagePerAttack: %d,\n", indent, int32(i))
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

func writeAttackRangeLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	fmt.Fprintf(sb, "%s%s: &AttackRange{\n", indent, field)
	for _, f := range []struct{ json, goName string }{
		{"min_reach", "MinReach"},
		{"max_reach", "MaxReach"},
		{"min_creative_reach", "MinCreativeReach"},
		{"max_creative_reach", "MaxCreativeReach"},
		{"hitbox_margin", "HitboxMargin"},
		{"mob_factor", "MobFactor"},
	} {
		if amount, ok := m[f.json].(float64); ok {
			fmt.Fprintf(sb, "%s\t%s: %v,\n", indent, f.goName, amount)
		}
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

func writeEquippableLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	slot, _ := m["slot"].(string)
	swappable, dispensable, damageOnHurt, equipOnInteract := true, true, true, false
	if v, ok := m["swappable"].(bool); ok {
		swappable = v
	}
	if v, ok := m["dispensable"].(bool); ok {
		dispensable = v
	}
	if v, ok := m["damage_on_hurt"].(bool); ok {
		damageOnHurt = v
	}
	if v, ok := m["equip_on_interact"].(bool); ok {
		equipOnInteract = v
	}
	fmt.Fprintf(sb, "%s%s: &Equippable{Slot: %q, Dispensable: %t, Swappable: %t, DamageOnHurt: %t, EquipOnInteract: %t", indent, field, slot, dispensable, swappable, damageOnHurt, equipOnInteract)
	for _, f := range []struct{ json, goName string }{
		{"equip_sound", "EquipSound"},
		{"asset_id", "AssetID"},
		{"camera_overlay", "CameraOverlay"},
		{"shearing_sound", "ShearingSound"},
	} {
		if v, ok := m[f.json].(string); ok {
			fmt.Fprintf(sb, ", %s: %q", f.goName, v)
		}
	}
	if v, ok := m["can_be_sheared"].(bool); ok {
		fmt.Fprintf(sb, ", CanBeSheared: %t", v)
	}
	if allowed, ok := m["allowed_entities"]; ok {
		sb.WriteString(", AllowedEntities: []string{")
		generateStringValues(sb, allowed)
		sb.WriteString("}")
	}
	sb.WriteString("},\n")
}

func writeEnchantableLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &Enchantable{Value: %d},\n", indent, field, int32(value.(map[string]any)["value"].(float64)))
}

func writeRepairableLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &Repairable{Items: %q},\n", indent, field, value.(map[string]any)["items"].(string))
}

func writeDamageResistantLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &DamageResistant{Types: %q},\n", indent, field, value.(map[string]any)["types"].(string))
}

func writeEntityDataLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &TypedEntityData{Type: %q},\n", indent, field, value.(map[string]any)["id"].(string))
}

func writeItemNameLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &ItemNameComponent{Translate: %q},\n", indent, field, value.(map[string]any)["translate"].(string))
}

func writeFireworksLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &Fireworks{FlightDuration: %d},\n", indent, field, int32(value.(map[string]any)["flight_duration"].(float64)))
}

// writeWritableBookContentLiteral marks the component present through its raw passthrough payload (empty page list).
func writeWritableBookContentLiteral(sb *strings.Builder, indent, field string, _ any) {
	fmt.Fprintf(sb, "%s%s: map[int32][]byte{ComponentWritableBookContent: {0}},\n", indent, field)
}

func writeUseCooldownLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	fmt.Fprintf(sb, "%s%s: &UseCooldown{Seconds: %v", indent, field, m["seconds"].(float64))
	if group, ok := m["cooldown_group"].(string); ok {
		fmt.Fprintf(sb, ", CooldownGroup: %q", group)
	}
	sb.WriteString("},\n")
}

func writeUseRemainderLiteral(sb *strings.Builder, indent, field string, value any) {
	m := value.(map[string]any)
	fmt.Fprintf(sb, "%s%s: &UseRemainder{\n", indent, field)
	if c, ok := m["count"].(float64); ok {
		fmt.Fprintf(sb, "%s\tCount: %d,\n", indent, int32(c))
	}
	if id, ok := m["id"].(string); ok {
		fmt.Fprintf(sb, "%s\tID: %q,\n", indent, id)
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

func writeDeathProtectionLiteral(sb *strings.Builder, indent, field string, value any) {
	fmt.Fprintf(sb, "%s%s: &DeathProtection{\n", indent, field)
	writeConsumeEffects(sb, indent+"\t", "DeathEffects", value.(map[string]any)["death_effects"])
	fmt.Fprintf(sb, "%s},\n", indent)
}

func writeAttributeModifiersLiteral(sb *strings.Builder, indent, field string, value any) {
	modifiers := value.([]any)
	if len(modifiers) == 0 {
		return
	}
	fmt.Fprintf(sb, "%s%s: []AttributeModifier{\n", indent, field)
	for _, entry := range modifiers {
		m := entry.(map[string]any)
		fields := make([]string, 0, 6)
		if t, ok := m["type"].(string); ok {
			fields = append(fields, fmt.Sprintf("Type: %q", t))
		}
		if a, ok := m["amount"].(float64); ok {
			fields = append(fields, fmt.Sprintf("Amount: %v", a))
		}
		if id, ok := m["id"].(string); ok {
			fields = append(fields, fmt.Sprintf("ID: %q", id))
		}
		if op, ok := m["operation"].(string); ok {
			fields = append(fields, fmt.Sprintf("Operation: %q", op))
		}
		if s, ok := m["slot"].(string); ok {
			fields = append(fields, fmt.Sprintf("Slot: %q", s))
		}
		if display, ok := m["display"].(map[string]any); ok {
			if displayType, ok := display["type"].(string); ok && displayType != "default" {
				fields = append(fields, fmt.Sprintf("Display: AttributeModifierDisplay{Type: %q}", displayType))
			}
		}
		fmt.Fprintf(sb, "%s\t{%s},\n", indent, strings.Join(fields, ", "))
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

func generateHolderSetLiteral(sb *strings.Builder, value any) {
	if value == nil {
		return
	}
	sb.WriteString(", Types: []string{")
	generateStringValues(sb, value)
	sb.WriteString("}")
}

func generateStringValues(sb *strings.Builder, value any) {
	switch value := value.(type) {
	case string:
		fmt.Fprintf(sb, "%q", value)
	case []any:
		for _, raw := range value {
			if entry, ok := raw.(string); ok {
				fmt.Fprintf(sb, "%q, ", entry)
			}
		}
	}
}

func generateComponentTypes(registries map[string]RegistryJSON, outPath string) {
	componentRegistry := registries["minecraft:data_component_type"]

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("items"))

	// component type constants
	sb.WriteString("// Component type protocol IDs from minecraft:data_component_type registry.\n")
	sb.WriteString("// https://minecraft.wiki/w/Java_Edition_protocol/Slot_data#Structured_components\n")
	sb.WriteString("const (\n")

	// find max ID for the constant
	var maxID int32
	for _, name := range sortedKeys(componentRegistry.Entries) {
		entry := componentRegistry.Entries[name]
		goName := "Component" + toGoVarName(name)
		// use untyped constants so they work with VarInt comparisons
		fmt.Fprintf(&sb, "\t%s = %d\n", goName, entry.ProtocolID)
		maxID = max(maxID, entry.ProtocolID)
	}
	fmt.Fprintf(&sb, "\n\tMaxComponentID = %d\n", maxID)
	sb.WriteString(")\n")

	writeFile(outPath, sb.String())
}

// entityVariantGroup collects registry components with a "/" in their name that the metadata does not describe.
const entityVariantGroup = "entityVariant"

// componentCodecGroups renders registrations per codec group in output order. Format verbs:
// %[1]s component constant, %[2]s Components field, %[3]s Go type without pointer.
var componentCodecGroups = []struct{ group, comment, format string }{
	{"varint", "simple VarInt codecs", "\tregisterCodec(%[1]s, &varIntCodec{\n\t\tget: func(c *Components) int32 { return c.%[2]s },\n\t\tset: func(c *Components, v int32) { c.%[2]s = v },\n\t})\n"},
	{"optionalVarint", "zero-valued optional VarInt codecs", "\tregisterCodec(%[1]s, pointerComponentCodec[int32]{\n\t\tget: func(c *Components) *int32 { return c.%[2]s },\n\t\tset: func(c *Components, v *int32) { c.%[2]s = v },\n\t\tread: readVarIntPointer, write: writeVarIntPointer, equal: equalValue[int32],\n\t})\n"},
	{"float32", "simple Float32 codecs", "\tregisterCodec(%[1]s, &float32Codec{\n\t\tget: func(c *Components) float64 { return c.%[2]s },\n\t\tset: func(c *Components, v float64) { c.%[2]s = v },\n\t})\n"},
	{"identifier", "simple string/identifier codecs", "\tregisterCodec(%[1]s, &stringCodec{\n\t\tget: func(c *Components) string { return c.%[2]s },\n\t\tset: func(c *Components, v string) { c.%[2]s = v },\n\t})\n"},
	{"empty", "empty marker codecs (bool flags)", "\tregisterCodec(%[1]s, &emptyMarkerCodec{\n\t\tget: func(c *Components) bool { return c.%[2]s },\n\t\tset: func(c *Components, v bool) { c.%[2]s = v },\n\t})\n"},
	{"struct", "struct codecs", "\tregisterCodec(%[1]s, gen%[3]sCodec{})\n"},
	{"passthrough:varint", "VarInt passthrough", "\tregisterPassthrough(%[1]s, decodeVarIntWire)\n"},
	{"passthrough:bool", "Bool passthrough", "\tregisterPassthrough(%[1]s, decodeBoolWire)\n"},
	{"passthrough:identifier", "String passthrough", "\tregisterPassthrough(%[1]s, decodeStringWire)\n"},
	{"passthrough:empty", "Empty passthrough", "\tregisterPassthrough(%[1]s, decodeEmptyWire)\n"},
	{"passthrough:int32", "Int32 passthrough", "\tregisterPassthrough(%[1]s, decodeInt32Wire)\n"},
	{"passthrough:nbt", "NBT passthrough", "\tregisterPassthrough(%[1]s, decodeNBTWire)\n"},
	{"passthrough:holderSet", "HolderSet passthrough", "\tregisterPassthrough(%[1]s, decodeHolderSetWire)\n"},
	{"passthrough:slotList", "SlotList passthrough", "\tregisterPassthrough(%[1]s, decodeSlotListWire)\n"},
	{"passthrough:slot", "Slot passthrough", "\tregisterPassthrough(%[1]s, decodeSlotWire)\n"},
	{entityVariantGroup, "Entity variant (VarInt) passthrough", "\tregisterPassthrough(%[1]s, decodeVarIntWire)\n"},
}

// componentCodecGroup returns the generated codec group of a component, or "" when items registers it by hand.
func componentCodecGroup(meta ComponentMetadata) string {
	switch {
	case meta.Passthrough:
		return "passthrough:" + meta.WireType
	case meta.GoField == "":
		return ""
	case meta.WireType == "struct" && len(meta.WireFormat) > 0:
		return "struct"
	case meta.WireType == "varint" && meta.AllowZero:
		return "optionalVarint"
	default:
		return meta.WireType
	}
}

func generateComponentCodecs(registries map[string]RegistryJSON, metadataPath, outPath string) {
	componentRegistry := registries["minecraft:data_component_type"]
	metadata := loadJSON[ComponentMetadataFile](metadataPath)

	// component registry names by codec group
	groups := map[string][]string{}
	for name, meta := range metadata.Components {
		if _, ok := componentRegistry.Entries[name]; ok {
			group := componentCodecGroup(meta)
			groups[group] = append(groups[group], name)
		}
	}
	for name := range componentRegistry.Entries {
		if _, described := metadata.Components[name]; !described && strings.Contains(strings.TrimPrefix(name, "minecraft:"), "/") {
			groups[entityVariantGroup] = append(groups[entityVariantGroup], name)
		}
	}
	for _, names := range groups {
		slices.Sort(names)
	}

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("items"))

	structs := groups["struct"]
	if len(structs) > 0 {
		sb.WriteString("import (\n")
		if slices.ContainsFunc(structs, func(name string) bool { return hasSliceField(metadata.Components[name].WireFormat) }) {
			sb.WriteString("\t\"slices\"\n\n")
		}
		sb.WriteString("\tns \"github.com/go-mclib/protocol/javaproto/wire\"\n")
		sb.WriteString(")\n\n")
	}

	sb.WriteString("// auto-generated codec registrations.\n\n")
	sb.WriteString("func init() {\n")
	for _, group := range componentCodecGroups {
		names := groups[group.group]
		if len(names) == 0 {
			continue
		}
		fmt.Fprintf(&sb, "\t// %s\n", group.comment)
		for _, name := range names {
			meta := metadata.Components[name]
			fmt.Fprintf(&sb, group.format, "Component"+toGoVarName(name), meta.GoField, strings.TrimPrefix(meta.GoType, "*"))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("}\n\n")

	// struct codec implementations
	for _, name := range structs {
		meta := metadata.Components[name]
		generateStructCodec(&sb, meta.GoField, strings.TrimPrefix(meta.GoType, "*"), meta.WireFormat)
	}

	writeFile(outPath, sb.String())
}

func writeConsumeEffects(sb *strings.Builder, indent, field string, value any) {
	effects, ok := value.([]any)
	if !ok || len(effects) == 0 {
		return
	}
	fmt.Fprintf(sb, "%s%s: []ConsumeEffect{\n", indent, field)
	for _, raw := range effects {
		effect, ok := raw.(map[string]any)
		if !ok {
			continue
		}
		fmt.Fprintf(sb, "%s\t{Type: %q", indent, effect["type"])
		probability := 1.0
		if configured, ok := effect["probability"].(float64); ok {
			probability = configured
		}
		fmt.Fprintf(sb, ", Probability: %v", probability)
		if effect["type"] == "minecraft:teleport_randomly" {
			diameter := 16.0
			if configured, ok := effect["diameter"].(float64); ok {
				diameter = configured
			}
			fmt.Fprintf(sb, ", Diameter: %v", diameter)
		}
		if sound, ok := effect["sound"].(string); ok {
			fmt.Fprintf(sb, ", Sound: %q", sound)
		}
		switch entries := effect["effects"].(type) {
		case string:
			fmt.Fprintf(sb, ", EffectIDs: []string{%q}", entries)
		case []any:
			sb.WriteString(", Effects: []StatusEffect{")
			for _, rawStatus := range entries {
				status, ok := rawStatus.(map[string]any)
				if !ok {
					continue
				}
				sb.WriteString("{")
				if id, ok := status["id"].(string); ok {
					fmt.Fprintf(sb, "ID: %q, ", id)
				}
				if duration, ok := status["duration"].(float64); ok {
					fmt.Fprintf(sb, "Duration: %d, ", int32(duration))
				}
				if amplifier, ok := status["amplifier"].(float64); ok {
					fmt.Fprintf(sb, "Amplifier: %d, ", int32(amplifier))
				}
				showParticles := true
				if configured, ok := status["show_particles"].(bool); ok {
					showParticles = configured
				}
				if showParticles {
					sb.WriteString("ShowParticles: true, ")
				}
				showIcon := true
				if configured, ok := status["show_icon"].(bool); ok {
					showIcon = configured
				}
				if showIcon {
					sb.WriteString("ShowIcon: true, ")
				}
				if ambient, ok := status["ambient"].(bool); ok && ambient {
					sb.WriteString("Ambient: true, ")
				}
				sb.WriteString("},")
			}
			sb.WriteString("}")
		}
		sb.WriteString("},\n")
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

func hasSliceField(fields []WireField) bool {
	return slices.ContainsFunc(fields, func(f WireField) bool { return f.GoField != "" && f.Type == "varintArray" })
}

// generateStructCodec generates a full componentCodec implementation for a pointer-to-struct component.
func generateStructCodec(sb *strings.Builder, goField, typeName string, fields []WireField) {
	codec := "gen" + typeName + "Codec"

	fmt.Fprintf(sb, "type %s struct{}\n\n", codec)

	// DecodeWire
	fmt.Fprintf(sb, "func (%s) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {\n", codec)
	sb.WriteString("\tw := ns.NewWriter()\n")
	for _, f := range fields {
		writeDecodeWireField(sb, f)
	}
	sb.WriteString("\treturn w.Bytes(), nil\n}\n\n")

	// Apply
	fmt.Fprintf(sb, "func (%s) Apply(c *Components, data []byte) error {\n", codec)
	sb.WriteString("\tbuf := ns.NewReader(data)\n")
	fmt.Fprintf(sb, "\ts := &%s{}\n", typeName)
	for _, f := range fields {
		if f.GoField == "" {
			break // trailing unmapped fields are skipped
		}
		writeApplyField(sb, f)
	}
	fmt.Fprintf(sb, "\tc.%s = s\n", goField)
	sb.WriteString("\treturn nil\n}\n\n")

	// Clear
	fmt.Fprintf(sb, "func (%s) Clear(c *Components) { c.%s = nil }\n\n", codec, goField)

	// Differs
	fmt.Fprintf(sb, "func (%s) Differs(c, defaults *Components) (bool, bool) {\n", codec)
	fmt.Fprintf(sb, "\tcHas := c.%s != nil\n\tdHas := defaults.%s != nil\n", goField, goField)
	sb.WriteString("\tif cHas != dHas {\n\t\treturn true, cHas\n\t}\n")
	sb.WriteString("\tif cHas && dHas {\n")
	if hasSliceField(fields) {
		for _, f := range fields {
			if f.GoField == "" {
				continue
			}
			if f.Type == "varintArray" {
				fmt.Fprintf(sb, "\t\tif !slices.Equal(c.%s.%s, defaults.%s.%s) {\n\t\t\treturn true, true\n\t\t}\n",
					goField, f.GoField, goField, f.GoField)
			} else {
				fmt.Fprintf(sb, "\t\tif c.%s.%s != defaults.%s.%s {\n\t\t\treturn true, true\n\t\t}\n",
					goField, f.GoField, goField, f.GoField)
			}
		}
		sb.WriteString("\t\treturn false, true\n")
	} else {
		fmt.Fprintf(sb, "\t\treturn *c.%s != *defaults.%s, true\n", goField, goField)
	}
	sb.WriteString("\t}\n\treturn false, false\n}\n\n")

	// Encode
	fmt.Fprintf(sb, "func (%s) Encode(c *Components) ([]byte, error) {\n", codec)
	fmt.Fprintf(sb, "\tif c.%s == nil {\n\t\treturn nil, nil\n\t}\n", goField)
	sb.WriteString("\tw := ns.NewWriter()\n")
	for _, f := range fields {
		writeEncodeField(sb, goField, f)
	}
	sb.WriteString("\treturn w.Bytes(), nil\n}\n\n")
}

func writeDecodeWireField(sb *strings.Builder, f WireField) {
	switch f.Type {
	case "varint":
		sb.WriteString("\tif err := w.CopyVarInt(buf); err != nil {\n\t\treturn nil, err\n\t}\n")
	case "float32":
		sb.WriteString("\tif err := w.CopyFloat32(buf); err != nil {\n\t\treturn nil, err\n\t}\n")
	case "bool":
		sb.WriteString("\tif err := w.CopyBool(buf); err != nil {\n\t\treturn nil, err\n\t}\n")
	case "varintArray":
		sb.WriteString("\t{\n")
		sb.WriteString("\t\tcount, err := buf.ReadVarInt()\n")
		sb.WriteString("\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n")
		sb.WriteString("\t\tw.WriteVarInt(count)\n")
		sb.WriteString("\t\tfor range int(count) {\n")
		sb.WriteString("\t\t\tif err := w.CopyVarInt(buf); err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n")
		sb.WriteString("\t\t}\n\t}\n")
	case "optionalIdentifier":
		sb.WriteString("\t{\n")
		sb.WriteString("\t\tpresent, err := buf.ReadBool()\n")
		sb.WriteString("\t\tif err != nil {\n\t\t\treturn nil, err\n\t\t}\n")
		sb.WriteString("\t\tw.WriteBool(present)\n")
		sb.WriteString("\t\tif present {\n")
		sb.WriteString("\t\t\tif err := w.CopyString(buf, maxStringLen); err != nil {\n\t\t\t\treturn nil, err\n\t\t\t}\n")
		sb.WriteString("\t\t}\n\t}\n")
	}
}

func writeApplyField(sb *strings.Builder, f WireField) {
	switch f.Type {
	case "varint":
		sb.WriteString("\t{\n\t\tv, err := buf.ReadVarInt()\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n")
		fmt.Fprintf(sb, "\t\ts.%s = int32(v)\n\t}\n", f.GoField)
	case "float32":
		sb.WriteString("\t{\n\t\tv, err := buf.ReadFloat32()\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n")
		fmt.Fprintf(sb, "\t\ts.%s = float64(v)\n\t}\n", f.GoField)
	case "bool":
		sb.WriteString("\t{\n\t\tv, err := buf.ReadBool()\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n")
		fmt.Fprintf(sb, "\t\ts.%s = bool(v)\n\t}\n", f.GoField)
	case "varintArray":
		sb.WriteString("\t{\n\t\tcount, err := buf.ReadVarInt()\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n")
		sb.WriteString("\t\tarr := make([]int32, 0, count)\n")
		sb.WriteString("\t\tfor range int(count) {\n")
		sb.WriteString("\t\t\tv, err := buf.ReadVarInt()\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n")
		sb.WriteString("\t\t\tarr = append(arr, int32(v))\n\t\t}\n")
		fmt.Fprintf(sb, "\t\ts.%s = arr\n\t}\n", f.GoField)
	case "optionalIdentifier":
		sb.WriteString("\t{\n\t\tpresent, err := buf.ReadBool()\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n")
		sb.WriteString("\t\tif present {\n\t\t\tv, err := buf.ReadIdentifier()\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}\n")
		fmt.Fprintf(sb, "\t\t\ts.%s = string(v)\n\t\t}\n\t}\n", f.GoField)
	}
}

func writeEncodeField(sb *strings.Builder, goField string, f WireField) {
	if f.GoField != "" {
		switch f.Type {
		case "varint":
			fmt.Fprintf(sb, "\tw.WriteVarInt(ns.VarInt(c.%s.%s))\n", goField, f.GoField)
		case "float32":
			fmt.Fprintf(sb, "\tw.WriteFloat32(ns.Float32(c.%s.%s))\n", goField, f.GoField)
		case "bool":
			fmt.Fprintf(sb, "\tw.WriteBool(ns.Boolean(c.%s.%s))\n", goField, f.GoField)
		case "varintArray":
			fmt.Fprintf(sb, "\tw.WriteVarInt(ns.VarInt(len(c.%s.%s)))\n", goField, f.GoField)
			fmt.Fprintf(sb, "\tfor _, v := range c.%s.%s {\n\t\tw.WriteVarInt(ns.VarInt(v))\n\t}\n", goField, f.GoField)
		case "optionalIdentifier":
			fmt.Fprintf(sb, "\tw.WriteBool(ns.Boolean(c.%s.%s != \"\"))\n", goField, f.GoField)
			fmt.Fprintf(sb, "\tif c.%s.%s != \"\" {\n\t\tw.WriteIdentifier(ns.Identifier(c.%s.%s))\n\t}\n", goField, f.GoField, goField, f.GoField)
		}
	} else {
		// unmapped field: write zero/default value
		switch f.Type {
		case "varint":
			sb.WriteString("\tw.WriteVarInt(0)\n")
		case "float32":
			sb.WriteString("\tw.WriteFloat32(0)\n")
		case "bool", "optionalIdentifier":
			sb.WriteString("\tw.WriteBool(false)\n")
		}
	}
}

var dyeColorNames = [...]string{
	"white", "orange", "magenta", "light_blue", "yellow", "lime", "pink", "gray",
	"light_gray", "cyan", "purple", "blue", "brown", "green", "red", "black",
}

func dyeColorIndex(name string) int {
	index := slices.Index(dyeColorNames[:], name)
	if index < 0 {
		panic("unknown dye color " + name)
	}
	return index
}

func resolvableNumberLiteral(value any, floating bool) string {
	typeName := "ResolvableInt"
	if floating {
		typeName = "ResolvableFloat"
	}
	if reference, ok := value.(string); ok {
		return fmt.Sprintf("%s{Reference: %q}", typeName, reference)
	}
	return fmt.Sprintf("%s{Value: %v}", typeName, value.(float64))
}

func writeSignTextLiteral(sb *strings.Builder, indent, field string, raw any) {
	value := raw.(map[string]any)
	color := 15 // black
	if name, ok := value["color"].(string); ok {
		color = dyeColorIndex(name)
	}
	glowing, _ := value["has_glowing_text"].(bool)
	fmt.Fprintf(sb, "%s%s: &SignText{Color: %d, HasGlowingText: %t,\n", indent, field, color, glowing)
	for _, messages := range []struct {
		json, field string
		optional    bool
	}{
		{"messages", "Messages", false}, {"filtered_messages", "FilteredMessages", true},
	} {
		raw, present := value[messages.json]
		if !present && messages.optional {
			continue
		}
		lines, ok := raw.([]any)
		if !ok || len(lines) != 4 {
			panic("sign text requires exactly four lines")
		}
		pointer := ""
		if messages.optional {
			pointer = "&"
		}
		fmt.Fprintf(sb, "%s\t%s: %s[4]ItemNameComponent{\n", indent, messages.field, pointer)
		for _, line := range lines {
			fmt.Fprintf(sb, "%s\t\t{Raw: %#v},\n", indent, componentDefaultNBT(line))
		}
		fmt.Fprintf(sb, "%s\t},\n", indent)
	}
	fmt.Fprintf(sb, "%s},\n", indent)
}

func componentDefaultNBT(value any) []byte {
	data, err := nbt.Encode(componentDefaultTag(value), "", nbt.FormatNetwork)
	if err != nil {
		panic(err)
	}
	return data
}

func componentDefaultTag(value any) nbt.Tag {
	switch v := value.(type) {
	case map[string]any:
		result := nbt.Compound{}
		for name, child := range v {
			result[name] = componentDefaultTag(child)
		}
		return result
	case []any:
		result := nbt.List{ElementType: nbt.TagEnd, Elements: make([]nbt.Tag, len(v))}
		for i, child := range v {
			result.Elements[i] = componentDefaultTag(child)
			if i == 0 {
				result.ElementType = result.Elements[i].ID()
			} else if result.ElementType != result.Elements[i].ID() {
				result.ElementType = nbt.TagCompound
			}
		}
		if result.ElementType == nbt.TagCompound {
			for i, child := range result.Elements {
				compound, ok := child.(nbt.Compound)
				if !ok || len(compound) == 1 && compound[""] != nil {
					result.Elements[i] = nbt.Compound{"": child}
				}
			}
		}
		return result
	case string:
		return nbt.String(v)
	case bool:
		if v {
			return nbt.Byte(1)
		}
		return nbt.Byte(0)
	case float64:
		if v >= math.MinInt32 && v <= math.MaxInt32 && v == float64(int32(v)) {
			return nbt.Int(int32(v))
		}
		return nbt.Double(v)
	default:
		panic(fmt.Sprintf("unsupported component default value %T", value))
	}
}
