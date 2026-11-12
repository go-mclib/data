package items

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

var swingAnimationNames = [3]string{"none", "whack", "stab"}

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

func signTextNBT(value *SignText) nbt.Tag {
	if value == nil {
		return nil
	}
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
		stack **Stack
	}{
		{"back", &result.Back}, {"left", &result.Left}, {"right", &result.Right}, {"front", &result.Front},
	} {
		if value, present := compound[side.name]; present {
			slot, err := TemplateFromNBT(value)
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

// TemplateFromNBT decodes an ItemStackTemplate: an item identifier or
// {id, count, components}. Components are a DataComponentPatch; `!component`
// removals (value {}) are kept even when the item has no such default, as in
// vanilla's ItemStackTemplate.
func TemplateFromNBT(value nbt.Tag) (ns.Slot, error) {
	if name, ok := value.(nbt.String); ok {
		id := ID(string(name))
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
	id := ID(compound.GetString("id"))
	if id < 0 || count < 1 || count > 99 {
		return ns.Slot{}, fmt.Errorf("invalid template item/count %q/%d", compound.GetString("id"), count)
	}
	components := compound.GetCompound("components")
	patch := make(nbt.Compound, len(components))
	var removals []ns.VarInt
	for key, value := range components {
		name, removed := strings.CutPrefix(key, "!")
		if !removed {
			patch[key] = value
			continue
		}
		componentID := registries.DataComponentType.Get(name)
		if componentID < 0 {
			return ns.Slot{}, fmt.Errorf("unknown removed component %q", name)
		}
		if unit, ok := value.(nbt.Compound); !ok || len(unit) != 0 {
			return ns.Slot{}, fmt.Errorf("removed component %q must have value {}", name)
		}
		removals = append(removals, ns.VarInt(componentID))
	}
	slot, err := ApplyNBTComponents(ns.NewSlot(ns.VarInt(id), ns.VarInt(count)), patch)
	if err != nil || len(removals) == 0 {
		return slot, err
	}
	// The template is a raw DataComponentPatch: removals stay even when the
	// item has no default for the component (unlike an applied stack patch).
	slices.Sort(removals)
	slot.Components.Remove = slices.Compact(removals)
	return slot, nil
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
