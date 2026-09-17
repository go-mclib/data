package items

import (
	"fmt"
	"math"

	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

var swingAnimationNames = [3]string{"none", "whack", "stab"}

func componentPayloadNBT(name string, c *Components) nbt.Tag {
	switch name {
	case "minecraft:attack_animation", "minecraft:interact_animation":
		value := c.AttackAnimation
		if name == "minecraft:interact_animation" {
			value = c.InteractAnimation
		}
		if value != nil && value.Type >= 0 && value.Type < 3 {
			return nbt.Compound{"type": nbt.String(swingAnimationNames[value.Type]), "duration": nbt.Int(value.Duration)}
		}
	case "minecraft:block_transformer":
		if c.BlockTransformer != "" {
			return nbt.String(c.BlockTransformer)
		}
	case "minecraft:provides_pottery_pattern":
		if c.ProvidesPotteryPattern != "" {
			return nbt.String(c.ProvidesPotteryPattern)
		}
	case "minecraft:villager_food":
		if c.VillagerFood != nil {
			return nbt.Compound{"nutrition": nbt.Int(c.VillagerFood.Nutrition)}
		}
	case "minecraft:compostable":
		if c.Compostable != nil {
			return nbt.Compound{"layers": resolvableIntNBT(c.Compostable.Layers)}
		}
	case "minecraft:cooking_fuel":
		if c.CookingFuel != nil {
			return nbt.Compound{"burn_time": resolvableIntNBT(c.CookingFuel.BurnTime), "speed_multiplier": resolvableFloatNBT(c.CookingFuel.SpeedMultiplier)}
		}
	case "minecraft:brewing_fuel":
		if c.BrewingFuel != nil {
			return nbt.Compound{"uses": resolvableIntNBT(c.BrewingFuel.Uses), "speed_multiplier": resolvableFloatNBT(c.BrewingFuel.SpeedMultiplier)}
		}
	case "minecraft:mob_visibility":
		if c.MobVisibility != nil {
			return nbt.Compound{"targeting_entity_types": holderSetNBT(c.MobVisibility.TargetingEntityTypes), "visibility": nbt.Float(c.MobVisibility.Visibility)}
		}
	case "minecraft:sign_text_front":
		if c.SignTextFront != nil {
			return signTextNBT(c.SignTextFront)
		}
	case "minecraft:sign_text_back":
		if c.SignTextBack != nil {
			return signTextNBT(c.SignTextBack)
		}
	case "minecraft:waxed":
		return nbt.Compound{}
	case "minecraft:cushion/color":
		if c.CushionColor != nil && *c.CushionColor >= 0 && *c.CushionColor < 16 {
			return nbt.String(dyeNames[*c.CushionColor])
		}
	case "minecraft:pot_decorations":
		if c.PotDecorations != nil {
			result := nbt.Compound{}
			for _, side := range []struct {
				name  string
				stack *ItemStack
			}{
				{"back", c.PotDecorations.Back}, {"left", c.PotDecorations.Left},
				{"right", c.PotDecorations.Right}, {"front", c.PotDecorations.Front},
			} {
				if side.stack != nil {
					if slot, err := side.stack.ToSlot(); err == nil {
						result[side.name] = componentSlotNBT(slot)
					}
				}
			}
			return result
		}
	}
	return nil
}

func applyComponentPayloadNBT(c *Components, name string, value nbt.Tag) (bool, error) {
	switch name {
	case "minecraft:block_transformer", "minecraft:provides_pottery_pattern":
		identifier, ok := value.(nbt.String)
		if !ok || identifier == "" {
			return true, fmt.Errorf("%s must be a registry identifier", name)
		}
		if name == "minecraft:block_transformer" {
			c.BlockTransformer = string(identifier)
		} else {
			c.ProvidesPotteryPattern = string(identifier)
		}
		return true, nil
	case "minecraft:waxed":
		if _, ok := value.(nbt.Compound); !ok {
			return true, fmt.Errorf("waxed must be a unit compound")
		}
		c.Waxed = true
		return true, nil
	case "minecraft:cushion/color":
		color, ok := value.(nbt.String)
		if !ok || dyeID(string(color)) < 0 {
			return true, fmt.Errorf("invalid cushion color %v", value)
		}
		c.CushionColor = int32Pointer(dyeID(string(color)))
		return true, nil
	case "minecraft:sign_text_front", "minecraft:sign_text_back":
		text, err := signTextFromNBT(value)
		if err != nil {
			return true, err
		}
		if name == "minecraft:sign_text_front" {
			c.SignTextFront = text
		} else {
			c.SignTextBack = text
		}
		return true, nil
	case "minecraft:pot_decorations":
		decorations, err := potDecorationsFromNBT(value)
		if err == nil {
			c.PotDecorations = decorations
		}
		return true, err
	case "minecraft:attack_animation", "minecraft:interact_animation", "minecraft:villager_food",
		"minecraft:compostable", "minecraft:cooking_fuel", "minecraft:brewing_fuel", "minecraft:mob_visibility":
	default:
		return false, nil
	}
	compound, ok := value.(nbt.Compound)
	if !ok {
		return true, fmt.Errorf("%s must be a compound", name)
	}
	switch name {
	case "minecraft:attack_animation", "minecraft:interact_animation":
		animation := &SwingAnimation{Type: 1, Duration: 6}
		if kind, present := compound["type"]; present {
			text, ok := kind.(nbt.String)
			animation.Type = -1
			if ok {
				for id, name := range swingAnimationNames {
					if string(text) == name {
						animation.Type = int32(id)
					}
				}
			}
			if animation.Type < 0 {
				return true, fmt.Errorf("invalid animation type %v", kind)
			}
		}
		if duration, present := compound["duration"]; present {
			var err error
			animation.Duration, err = componentIntNBT(duration)
			if err != nil {
				return true, err
			}
		}
		if animation.Duration < 0 {
			return true, fmt.Errorf("negative animation duration")
		}
		if name == "minecraft:attack_animation" {
			c.AttackAnimation = animation
		} else {
			c.InteractAnimation = animation
		}
	case "minecraft:villager_food":
		nutrition, err := componentIntNBT(compound["nutrition"])
		if err != nil {
			return true, err
		}
		if nutrition <= 0 {
			return true, fmt.Errorf("villager food nutrition must be positive")
		}
		c.VillagerFood = &VillagerFood{Nutrition: nutrition}
	case "minecraft:compostable":
		layers, err := resolvableIntFromNBT(compound["layers"])
		if err != nil {
			return true, err
		}
		c.Compostable = &Compostable{Layers: layers}
	case "minecraft:cooking_fuel", "minecraft:brewing_fuel":
		field := "burn_time"
		if name == "minecraft:brewing_fuel" {
			field = "uses"
		}
		amount, err := resolvableIntFromNBT(compound[field])
		if err != nil {
			return true, err
		}
		speed, err := resolvableFloatFromNBT(compound["speed_multiplier"])
		if err != nil {
			return true, err
		}
		if name == "minecraft:cooking_fuel" {
			c.CookingFuel = &CookingFuel{BurnTime: amount, SpeedMultiplier: speed}
		} else {
			c.BrewingFuel = &BrewingFuel{Uses: amount, SpeedMultiplier: speed}
		}
	case "minecraft:mob_visibility":
		types, err := holderSetFromNBT(compound["targeting_entity_types"])
		if err != nil {
			return true, err
		}
		visibility, err := componentFloatNBT(compound["visibility"])
		if err != nil {
			return true, err
		}
		if visibility < 0 || visibility > 10 || math.IsNaN(float64(visibility)) {
			return true, fmt.Errorf("mob visibility outside [0,10]")
		}
		c.MobVisibility = &MobVisibility{TargetingEntityTypes: types, Visibility: visibility}
	}
	return true, nil
}

func componentIntNBT(value nbt.Tag) (int32, error) {
	switch value.(type) {
	case nbt.Byte, nbt.Short, nbt.Int:
		return intValue(value), nil
	case nbt.Long:
		v := int64(value.(nbt.Long))
		if v >= math.MinInt32 && v <= math.MaxInt32 {
			return int32(v), nil
		}
	}
	return 0, fmt.Errorf("expected a 32-bit integer, got %v", value)
}
func componentFloatNBT(value nbt.Tag) (float32, error) {
	switch value.(type) {
	case nbt.Byte, nbt.Short, nbt.Int, nbt.Long, nbt.Float, nbt.Double:
		return float32(floatValue(value)), nil
	}
	return 0, fmt.Errorf("expected a number, got %v", value)
}
func resolvableIntNBT(value ResolvableInt) nbt.Tag {
	if value.Reference != "" {
		return nbt.String(value.Reference)
	}
	return nbt.Int(value.Value)
}
func resolvableFloatNBT(value ResolvableFloat) nbt.Tag {
	if value.Reference != "" {
		return nbt.String(value.Reference)
	}
	return nbt.Float(value.Value)
}
func resolvableIntFromNBT(value nbt.Tag) (ResolvableInt, error) {
	if reference, ok := value.(nbt.String); ok {
		if reference == "" {
			return ResolvableInt{}, fmt.Errorf("empty integer provider reference")
		}
		return ResolvableInt{Reference: string(reference)}, nil
	}
	constant, err := componentIntNBT(value)
	return ResolvableInt{Value: constant}, err
}
func resolvableFloatFromNBT(value nbt.Tag) (ResolvableFloat, error) {
	if reference, ok := value.(nbt.String); ok {
		if reference == "" {
			return ResolvableFloat{}, fmt.Errorf("empty float provider reference")
		}
		return ResolvableFloat{Reference: string(reference)}, nil
	}
	constant, err := componentFloatNBT(value)
	return ResolvableFloat{Value: constant}, err
}

func signTextNBT(value *SignText) nbt.Compound {
	result := nbt.Compound{"messages": signMessagesNBT(&value.Messages), "color": nbt.String(dyeNames[value.Color]), "has_glowing_text": nbt.Byte(0)}
	if value.HasGlowingText {
		result["has_glowing_text"] = nbt.Byte(1)
	}
	if value.FilteredMessages != nil {
		result["filtered_messages"] = signMessagesNBT(value.FilteredMessages)
	}
	return result
}
func signMessagesNBT(messages *[4]ItemNameComponent) nbt.List {
	// Java ListTag uses a single empty-name compound field to wrap
	// heterogeneous entries. Preserve the original text tag exactly.
	result := nbt.List{ElementType: nbt.TagCompound, Elements: make([]nbt.Tag, 4)}
	for i := range messages {
		value := textComponentNBT(&messages[i])
		if compound, ok := value.(nbt.Compound); !ok || len(compound) == 1 && compound[""] != nil {
			value = nbt.Compound{"": value}
		}
		result.Elements[i] = value
	}
	return result
}
func signTextFromNBT(value nbt.Tag) (*SignText, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("sign text must be a compound")
	}
	result := &SignText{Color: 15, HasGlowingText: nbtBoolOr(compound, "has_glowing_text", false)}
	if err := signMessagesFromNBT(compound["messages"], &result.Messages); err != nil {
		return nil, err
	}
	if filtered, ok := compound["filtered_messages"]; ok {
		result.FilteredMessages = new([4]ItemNameComponent)
		if err := signMessagesFromNBT(filtered, result.FilteredMessages); err != nil {
			return nil, err
		}
	}
	if color, present := compound["color"]; present {
		name, ok := color.(nbt.String)
		if !ok || dyeID(string(name)) < 0 {
			return nil, fmt.Errorf("invalid sign dye color %v", color)
		}
		result.Color = dyeID(string(name))
	}
	return result, nil
}
func signMessagesFromNBT(value nbt.Tag, messages *[4]ItemNameComponent) error {
	list, ok := value.(nbt.List)
	if !ok || len(list.Elements) != 4 {
		return fmt.Errorf("sign messages must contain exactly four text components")
	}
	for i, value := range list.Elements {
		if compound, ok := value.(nbt.Compound); ok && len(compound) == 1 && compound[""] != nil {
			value = compound[""]
		}
		text := textComponentFromNBT(value)
		if text == nil {
			return fmt.Errorf("invalid sign message %d", i)
		}
		messages[i] = *text
	}
	return nil
}
func potDecorationsFromNBT(value nbt.Tag) (*PotDecorations, error) {
	compound, ok := value.(nbt.Compound)
	if !ok {
		return nil, fmt.Errorf("pot decorations must be a compound")
	}
	result := &PotDecorations{}
	for _, side := range []struct {
		name  string
		stack **ItemStack
	}{
		{"back", &result.Back}, {"left", &result.Left}, {"right", &result.Right}, {"front", &result.Front},
	} {
		if value, present := compound[side.name]; present {
			slot, err := itemTemplateFromNBT(value)
			if err != nil {
				return nil, fmt.Errorf("pot decoration %s: %w", side.name, err)
			}
			*side.stack, err = FromSlot(slot)
			if err != nil {
				return nil, err
			}
		}
	}
	return result, nil
}
func itemTemplateFromNBT(value nbt.Tag) (ns.Slot, error) {
	if name, ok := value.(nbt.String); ok {
		id := ItemID(string(name))
		if id < 0 {
			return ns.Slot{}, fmt.Errorf("unknown template item %q", name)
		}
		return ns.NewSlot(ns.VarInt(id), 1), nil
	}
	compound, ok := value.(nbt.Compound)
	if !ok {
		return ns.Slot{}, fmt.Errorf("item template must be an identifier or compound")
	}
	count := int32(1)
	if value, present := compound["count"]; present {
		var err error
		count, err = componentIntNBT(value)
		if err != nil {
			return ns.Slot{}, err
		}
	}
	id := ItemID(compound.GetString("id"))
	if id < 0 || count < 1 || count > 99 {
		return ns.Slot{}, fmt.Errorf("invalid template item/count %q/%d", compound.GetString("id"), count)
	}
	return ApplyNBTComponents(ns.NewSlot(ns.VarInt(id), ns.VarInt(count)), compound.GetCompound("components"))
}

// defaultPotDecorations decodes generated defaults only after codec registration
// and the item default table are initialized, so nested patches inherit defaults.
func defaultPotDecorations(data []byte) *PotDecorations {
	value, err := rawNBTComponent(data)
	if err != nil {
		panic(err)
	}
	decorations, err := potDecorationsFromNBT(value)
	if err != nil {
		panic(err)
	}
	return decorations
}
