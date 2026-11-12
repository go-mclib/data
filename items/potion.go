package items

import (
	"fmt"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// maxPotionEffects bounds the custom effect and stew effect lists.
const maxPotionEffects = 256

// PotionContentsData contains the gameplay-relevant potion component fields.
type PotionContentsData struct {
	Potion      string
	CustomColor *int32
	Effects     []StatusEffect
	CustomName  string
}

// PotionSlot builds a potion item stack carrying a potion_contents component.
func PotionSlot(itemID ns.VarInt, potion string) ns.Slot {
	buf := ns.NewWriter()
	buf.WriteBool(true)
	buf.WriteVarInt(ns.VarInt(registries.Potion.Get(potion)))
	buf.WriteBool(false)
	buf.WriteVarInt(0)
	buf.WriteBool(false)
	s := ns.NewSlot(itemID, 1)
	s.Components.Add = append(s.Components.Add, ns.RawSlotComponent{
		ID: ns.VarInt(ComponentPotionContents), Data: buf.Bytes(),
	})
	return s
}

// SuspiciousStewEffect contains one component-defined stew effect.
type SuspiciousStewEffect struct {
	Effect   string
	Duration int
}

// ReadSuspiciousStewEffects returns the stack's patched stew effects.
func ReadSuspiciousStewEffects(stack ns.Slot) ([]SuspiciousStewEffect, error) {
	data, added, _ := patchedComponent(stack, ComponentSuspiciousStewEffects)
	if !added {
		return nil, nil
	}
	reader := ns.NewReader(data)
	count, err := reader.ReadVarInt()
	if err != nil || count < 0 || count > maxPotionEffects {
		return nil, fmt.Errorf("invalid suspicious stew effect count %d", count)
	}
	effects := make([]SuspiciousStewEffect, 0, int(count))
	for range int(count) {
		id, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		name := registries.MobEffect.ByID(int32(id))
		if name == "" {
			return nil, fmt.Errorf("unknown suspicious stew effect %d", id)
		}
		duration, err := reader.ReadVarInt()
		if err != nil {
			return nil, err
		}
		effects = append(effects, SuspiciousStewEffect{Effect: name, Duration: int(duration)})
	}
	return effects, nil
}

// PotionDurationScale returns the component-defined regular-effect multiplier.
func PotionDurationScale(stack ns.Slot) float64 {
	data, added, removed := patchedComponent(stack, ComponentPotionDurationScale)
	if removed {
		return 1
	}
	if added {
		if value, err := ns.NewReader(data).ReadFloat32(); err == nil {
			return float64(value)
		}
	}
	if defaults := DefaultComponents(int32(stack.ItemID)); defaults != nil && defaults.PotionDurationScale != 0 {
		return defaults.PotionDurationScale
	}
	return 1
}

// ReadPotionContents decodes a stack's potion_contents component.
func ReadPotionContents(stack ns.Slot) (PotionContentsData, error) {
	data, added, _ := patchedComponent(stack, ComponentPotionContents)
	if !added {
		return PotionContentsData{}, nil
	}
	return readPotionContents(ns.NewReader(data))
}

func readPotionContents(buf *ns.PacketBuffer) (PotionContentsData, error) {
	var contents PotionContentsData
	hasPotion, err := buf.ReadBool()
	if err != nil {
		return contents, err
	}
	if hasPotion {
		id, err := buf.ReadVarInt()
		if err != nil {
			return contents, err
		}
		contents.Potion = registries.Potion.ByID(int32(id))
		if contents.Potion == "" {
			return contents, fmt.Errorf("unknown potion %d", id)
		}
	}
	hasColor, err := buf.ReadBool()
	if err != nil {
		return contents, err
	}
	if hasColor {
		color, err := buf.ReadInt32()
		if err != nil {
			return contents, err
		}
		contents.CustomColor = int32Pointer(int32(color))
	}
	count, err := buf.ReadVarInt()
	if err != nil || count < 0 || count > maxPotionEffects {
		return contents, fmt.Errorf("invalid custom potion effect count %d", count)
	}
	contents.Effects = make([]StatusEffect, 0, int(count))
	for range int(count) {
		effect, err := readStatusEffect(buf)
		if err != nil {
			return contents, err
		}
		contents.Effects = append(contents.Effects, effect)
	}
	hasName, err := buf.ReadBool()
	if err != nil {
		return contents, err
	}
	if hasName {
		name, err := buf.ReadString(maxStringLen)
		if err != nil {
			return contents, err
		}
		contents.CustomName = string(name)
	}
	return contents, nil
}

func writePotionContents(w *ns.PacketBuffer, contents PotionContentsData) error {
	w.WriteBool(ns.Boolean(contents.Potion != ""))
	if contents.Potion != "" {
		id := registries.Potion.Get(contents.Potion)
		if id < 0 {
			return fmt.Errorf("unknown potion %q", contents.Potion)
		}
		w.WriteVarInt(ns.VarInt(id))
	}
	w.WriteBool(ns.Boolean(contents.CustomColor != nil))
	if contents.CustomColor != nil {
		w.WriteInt32(ns.Int32(*contents.CustomColor))
	}
	if len(contents.Effects) > maxPotionEffects {
		return fmt.Errorf("too many custom potion effects")
	}
	w.WriteVarInt(ns.VarInt(len(contents.Effects)))
	for _, effect := range contents.Effects {
		if err := writeStatusEffect(w, effect); err != nil {
			return err
		}
	}
	w.WriteBool(ns.Boolean(contents.CustomName != ""))
	if contents.CustomName != "" {
		return w.WriteString(ns.String(contents.CustomName))
	}
	return nil
}

// PotionName reads the potion id from a stack's potion_contents component.
func PotionName(stack ns.Slot) string {
	contents, err := ReadPotionContents(stack)
	if err != nil {
		return ""
	}
	return contents.Potion
}
