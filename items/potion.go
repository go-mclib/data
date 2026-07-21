package items

import (
	"fmt"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

var potionContentsComponent = ns.VarInt(registries.DataComponentType.Get("minecraft:potion_contents"))

// PotionEffect contains one custom potion effect.
type PotionEffect struct {
	Effect        string
	Amplifier     int
	Duration      int
	Ambient       bool
	ShowParticles bool
	ShowIcon      bool
	Hidden        *PotionEffect
}

// PotionContentsData contains the gameplay-relevant potion component fields.
type PotionContentsData struct {
	Potion      string
	CustomColor *int32
	Effects     []PotionEffect
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
		ID: potionContentsComponent, Data: buf.Bytes(),
	})
	return s
}

func readPotionEffectDetails(buf *ns.PacketBuffer, effect string) (PotionEffect, error) {
	amplifier, err := buf.ReadVarInt()
	if err != nil {
		return PotionEffect{}, err
	}
	duration, err := buf.ReadVarInt()
	if err != nil {
		return PotionEffect{}, err
	}
	ambient, err := buf.ReadBool()
	if err != nil {
		return PotionEffect{}, err
	}
	particles, err := buf.ReadBool()
	if err != nil {
		return PotionEffect{}, err
	}
	icon, err := buf.ReadBool()
	if err != nil {
		return PotionEffect{}, err
	}
	hasHidden, err := buf.ReadBool()
	if err != nil {
		return PotionEffect{}, err
	}
	var hidden *PotionEffect
	if hasHidden {
		value, err := readPotionEffectDetails(buf, effect)
		if err != nil {
			return PotionEffect{}, err
		}
		hidden = &value
	}
	return PotionEffect{
		Effect: effect, Amplifier: int(amplifier), Duration: int(duration), Ambient: bool(ambient),
		ShowParticles: bool(particles), ShowIcon: bool(icon), Hidden: hidden,
	}, nil
}

// SuspiciousStewEffect contains one component-defined stew effect.
type SuspiciousStewEffect struct {
	Effect   string
	Duration int
}

// ReadSuspiciousStewEffects returns the stack's patched stew effects.
func ReadSuspiciousStewEffects(stack ns.Slot) ([]SuspiciousStewEffect, error) {
	for _, removed := range stack.Components.Remove {
		if int32(removed) == ComponentSuspiciousStewEffects {
			return nil, nil
		}
	}
	for _, component := range stack.Components.Add {
		if int32(component.ID) != ComponentSuspiciousStewEffects {
			continue
		}
		reader := ns.NewReader(component.Data)
		count, err := reader.ReadVarInt()
		if err != nil || count < 0 || count > 256 {
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
	return nil, nil
}

// PotionDurationScale returns the component-defined regular-effect multiplier.
func PotionDurationScale(stack ns.Slot) float64 {
	durationScale := 1.0
	if defaults := DefaultComponents(int32(stack.ItemID)); defaults != nil && defaults.PotionDurationScale != 0 {
		durationScale = defaults.PotionDurationScale
	}
	for _, component := range stack.Components.Add {
		if int32(component.ID) == ComponentPotionDurationScale {
			if value, err := ns.NewReader(component.Data).ReadFloat32(); err == nil {
				durationScale = float64(value)
			}
		}
	}
	for _, component := range stack.Components.Remove {
		if int32(component) == ComponentPotionDurationScale {
			durationScale = 1
		}
	}
	return durationScale
}

// ReadPotionContents decodes a stack's potion_contents component.
func ReadPotionContents(stack ns.Slot) (PotionContentsData, error) {
	for _, component := range stack.Components.Add {
		if component.ID != potionContentsComponent {
			continue
		}
		buf := ns.NewReader(component.Data)
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
			value := int32(color)
			contents.CustomColor = &value
		}
		count, err := buf.ReadVarInt()
		if err != nil || count < 0 || count > 256 {
			return contents, fmt.Errorf("invalid custom potion effect count %d", count)
		}
		contents.Effects = make([]PotionEffect, 0, int(count))
		for range int(count) {
			id, err := buf.ReadVarInt()
			if err != nil {
				return contents, err
			}
			effect, err := readPotionEffectDetails(buf, registries.MobEffect.ByID(int32(id)))
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
			name, err := buf.ReadString(32767)
			if err != nil {
				return contents, err
			}
			contents.CustomName = string(name)
		}
		return contents, nil
	}
	return PotionContentsData{}, nil
}

// PotionName reads the potion id from a stack's potion_contents component.
func PotionName(stack ns.Slot) string {
	contents, err := ReadPotionContents(stack)
	if err != nil {
		return ""
	}
	return contents.Potion
}
