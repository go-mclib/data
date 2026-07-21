package items

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

type consumableCodec struct{}

var consumeAnimations = []string{
	"none", "eat", "drink", "block", "bow", "trident", "crossbow", "spyglass", "toot_horn", "brush", "bundle", "spear",
}

func (consumableCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := decodeConsumableWire(buf, w); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (consumableCodec) Apply(c *Components, data []byte) error {
	buf := ns.NewReader(data)
	seconds, err := buf.ReadFloat32()
	if err != nil {
		return err
	}
	animationID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	if animationID < 0 || int(animationID) >= len(consumeAnimations) {
		return fmt.Errorf("invalid consume animation %d", animationID)
	}
	consumeSound, err := readConsumableSound(buf)
	if err != nil {
		return err
	}
	hasParticles, err := buf.ReadBool()
	if err != nil {
		return err
	}
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	if count < 0 {
		return fmt.Errorf("negative consume effect count %d", count)
	}
	effects := make([]ConsumeEffect, 0, count)
	for range int(count) {
		effect, err := readConsumeEffect(buf)
		if err != nil {
			return err
		}
		effects = append(effects, effect)
	}
	c.Consumable = &Consumable{
		ConsumeSeconds:   float64(seconds),
		Animation:        consumeAnimations[animationID],
		Sound:            consumeSound.name,
		SoundDirect:      consumeSound.direct,
		SoundRange:       consumeSound.fixedRange,
		HasParticles:     bool(hasParticles),
		OnConsumeEffects: effects,
	}
	return nil
}

func (consumableCodec) Clear(c *Components) { c.Consumable = nil }

func (consumableCodec) Differs(c, defaults *Components) (bool, bool) {
	return !reflect.DeepEqual(c.Consumable, defaults.Consumable), c.Consumable != nil
}

func (consumableCodec) Encode(c *Components) ([]byte, error) {
	if c.Consumable == nil {
		return nil, fmt.Errorf("consumable component is nil")
	}
	animationID := -1
	for id, animation := range consumeAnimations {
		if c.Consumable.Animation == animation {
			animationID = id
			break
		}
	}
	if animationID < 0 {
		return nil, fmt.Errorf("unknown consume animation %q", c.Consumable.Animation)
	}
	w := ns.NewWriter()
	w.WriteFloat32(ns.Float32(c.Consumable.ConsumeSeconds))
	w.WriteVarInt(ns.VarInt(animationID))
	if err := writeConsumableSound(w, c.Consumable.Sound, c.Consumable.SoundDirect, c.Consumable.SoundRange); err != nil {
		return nil, err
	}
	w.WriteBool(ns.Boolean(c.Consumable.HasParticles))
	w.WriteVarInt(ns.VarInt(len(c.Consumable.OnConsumeEffects)))
	for _, effect := range c.Consumable.OnConsumeEffects {
		if err := writeConsumeEffect(w, effect); err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}

type consumableSound struct {
	name       string
	direct     bool
	fixedRange *float64
}

func readConsumableSound(buf *ns.PacketBuffer) (consumableSound, error) {
	var sound ns.SoundEvent
	if err := sound.Decode(buf); err != nil {
		return consumableSound{}, err
	}
	if sound.Direct {
		result := consumableSound{name: string(sound.SoundID), direct: true}
		if sound.FixedRange.Present {
			value := float64(sound.FixedRange.Value)
			result.fixedRange = &value
		}
		return result, nil
	}
	name := registries.SoundEvent.ByID(int32(sound.RegistryID))
	if name == "" {
		return consumableSound{}, fmt.Errorf("unknown sound event %d", sound.RegistryID)
	}
	return consumableSound{name: name}, nil
}

func writeConsumableSound(w *ns.PacketBuffer, name string, direct bool, fixedRange *float64) error {
	id := registries.SoundEvent.Get(name)
	if !direct && id >= 0 {
		return ns.RegistrySoundEvent(id).Encode(w)
	}
	sound := ns.DirectSoundEvent(ns.Identifier(name))
	if fixedRange != nil {
		sound.FixedRange = ns.Some(ns.Float32(*fixedRange))
	}
	return sound.Encode(w)
}

func readConsumeEffect(buf *ns.PacketBuffer) (ConsumeEffect, error) {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return ConsumeEffect{}, err
	}
	effect := ConsumeEffect{Probability: 1}
	switch typeID {
	case 0:
		effect.Type = "minecraft:apply_effects"
		count, err := buf.ReadVarInt()
		if err != nil {
			return ConsumeEffect{}, err
		}
		if count < 0 {
			return ConsumeEffect{}, fmt.Errorf("negative status effect count %d", count)
		}
		for range int(count) {
			status, err := readStatusEffect(buf)
			if err != nil {
				return ConsumeEffect{}, err
			}
			effect.Effects = append(effect.Effects, status)
		}
		probability, err := buf.ReadFloat32()
		if err != nil {
			return ConsumeEffect{}, err
		}
		effect.Probability = float64(probability)
	case 1:
		effect.Type = "minecraft:remove_effects"
		effect.EffectIDs, err = readEffectHolderSet(buf)
		if err != nil {
			return ConsumeEffect{}, err
		}
	case 2:
		effect.Type = "minecraft:clear_all_effects"
	case 3:
		effect.Type = "minecraft:teleport_randomly"
		diameter, err := buf.ReadFloat32()
		if err != nil {
			return ConsumeEffect{}, err
		}
		effect.Diameter = float64(diameter)
	case 4:
		effect.Type = "minecraft:play_sound"
		sound, soundErr := readConsumableSound(buf)
		if soundErr != nil {
			return ConsumeEffect{}, soundErr
		}
		effect.Sound, effect.SoundDirect, effect.SoundRange = sound.name, sound.direct, sound.fixedRange
	default:
		return ConsumeEffect{}, fmt.Errorf("unknown consume effect type %d", typeID)
	}
	return effect, nil
}

func writeConsumeEffect(w *ns.PacketBuffer, effect ConsumeEffect) error {
	switch effect.Type {
	case "minecraft:apply_effects":
		w.WriteVarInt(0)
		w.WriteVarInt(ns.VarInt(len(effect.Effects)))
		for _, status := range effect.Effects {
			if err := writeStatusEffect(w, status); err != nil {
				return err
			}
		}
		w.WriteFloat32(ns.Float32(effect.Probability))
	case "minecraft:remove_effects":
		w.WriteVarInt(1)
		if err := writeEffectHolderSet(w, effect.EffectIDs); err != nil {
			return err
		}
	case "minecraft:clear_all_effects":
		w.WriteVarInt(2)
	case "minecraft:teleport_randomly":
		w.WriteVarInt(3)
		w.WriteFloat32(ns.Float32(effect.Diameter))
	case "minecraft:play_sound":
		w.WriteVarInt(4)
		if err := writeConsumableSound(w, effect.Sound, effect.SoundDirect, effect.SoundRange); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown consume effect type %q", effect.Type)
	}
	return nil
}

func readStatusEffect(buf *ns.PacketBuffer) (StatusEffect, error) {
	effectID, err := buf.ReadVarInt()
	if err != nil {
		return StatusEffect{}, err
	}
	name := registries.MobEffect.ByID(int32(effectID))
	if name == "" {
		return StatusEffect{}, fmt.Errorf("unknown mob effect %d", effectID)
	}
	return readStatusEffectDetails(buf, name)
}

func readStatusEffectDetails(buf *ns.PacketBuffer, name string) (StatusEffect, error) {
	amplifier, err := buf.ReadVarInt()
	if err != nil {
		return StatusEffect{}, err
	}
	duration, err := buf.ReadVarInt()
	if err != nil {
		return StatusEffect{}, err
	}
	ambient, err := buf.ReadBool()
	if err != nil {
		return StatusEffect{}, err
	}
	showParticles, err := buf.ReadBool()
	if err != nil {
		return StatusEffect{}, err
	}
	showIcon, err := buf.ReadBool()
	if err != nil {
		return StatusEffect{}, err
	}
	hasHidden, err := buf.ReadBool()
	if err != nil {
		return StatusEffect{}, err
	}
	status := StatusEffect{
		ID:            name,
		Amplifier:     int32(amplifier),
		Duration:      int32(duration),
		Ambient:       bool(ambient),
		ShowParticles: bool(showParticles),
		ShowIcon:      bool(showIcon),
	}
	if hasHidden {
		hidden, err := readStatusEffectDetails(buf, name)
		if err != nil {
			return StatusEffect{}, err
		}
		status.Hidden = &hidden
	}
	return status, nil
}

func writeStatusEffect(w *ns.PacketBuffer, status StatusEffect) error {
	id := registries.MobEffect.Get(status.ID)
	if id < 0 {
		return fmt.Errorf("unknown mob effect %q", status.ID)
	}
	w.WriteVarInt(ns.VarInt(id))
	return writeStatusEffectDetails(w, status)
}

func writeStatusEffectDetails(w *ns.PacketBuffer, status StatusEffect) error {
	w.WriteVarInt(ns.VarInt(status.Amplifier))
	w.WriteVarInt(ns.VarInt(status.Duration))
	w.WriteBool(ns.Boolean(status.Ambient))
	w.WriteBool(ns.Boolean(status.ShowParticles))
	w.WriteBool(ns.Boolean(status.ShowIcon))
	w.WriteBool(ns.Boolean(status.Hidden != nil))
	if status.Hidden != nil {
		return writeStatusEffectDetails(w, *status.Hidden)
	}
	return nil
}

func readEffectHolderSet(buf *ns.PacketBuffer) ([]string, error) {
	marker, err := buf.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if marker == 0 {
		tag, err := buf.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		return []string{"#" + string(tag)}, nil
	}
	ids := make([]string, 0, marker-1)
	for range int(marker) - 1 {
		id, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		name := registries.MobEffect.ByID(int32(id))
		if name == "" {
			return nil, fmt.Errorf("unknown mob effect %d", id)
		}
		ids = append(ids, name)
	}
	return ids, nil
}

func writeEffectHolderSet(w *ns.PacketBuffer, effects []string) error {
	if len(effects) == 1 && strings.HasPrefix(effects[0], "#") {
		w.WriteVarInt(0)
		return w.WriteString(ns.String(strings.TrimPrefix(effects[0], "#")))
	}
	w.WriteVarInt(ns.VarInt(len(effects) + 1))
	for _, name := range effects {
		id := registries.MobEffect.Get(name)
		if id < 0 {
			return fmt.Errorf("unknown mob effect %q", name)
		}
		w.WriteVarInt(ns.VarInt(id))
	}
	return nil
}
