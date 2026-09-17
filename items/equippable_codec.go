package items

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

type equippableCodec struct{}

var equipmentSlots = []string{"mainhand", "offhand", "feet", "legs", "chest", "head", "body", "saddle"}

func (equippableCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	component, err := readEquippable(buf)
	if err != nil {
		return nil, err
	}
	return writeEquippable(component)
}

func (equippableCodec) Apply(c *Components, data []byte) error {
	component, err := readEquippable(ns.NewReader(data))
	if err != nil {
		return err
	}
	c.Equippable = component
	return nil
}

func (equippableCodec) Clear(c *Components) { c.Equippable = nil }

func (equippableCodec) Differs(c, defaults *Components) (bool, bool) {
	return !reflect.DeepEqual(c.Equippable, defaults.Equippable), c.Equippable != nil
}

func (equippableCodec) Encode(c *Components) ([]byte, error) {
	if c.Equippable == nil {
		return nil, fmt.Errorf("equippable component is nil")
	}
	return writeEquippable(c.Equippable)
}

func readEquippable(buf *ns.PacketBuffer) (*Equippable, error) {
	slotID, err := buf.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if slotID < 0 || int(slotID) >= len(equipmentSlots) {
		return nil, fmt.Errorf("unknown equipment slot %d", slotID)
	}
	equipSound, err := readConsumableSound(buf)
	if err != nil {
		return nil, err
	}
	assetID, err := readOptionalIdentifier(buf)
	if err != nil {
		return nil, err
	}
	cameraOverlay, err := readOptionalIdentifier(buf)
	if err != nil {
		return nil, err
	}
	allowedEntities, err := readOptionalEntityTypeHolderSet(buf)
	if err != nil {
		return nil, err
	}
	dispensable, err := buf.ReadBool()
	if err != nil {
		return nil, err
	}
	swappable, err := buf.ReadBool()
	if err != nil {
		return nil, err
	}
	damageOnHurt, err := buf.ReadBool()
	if err != nil {
		return nil, err
	}
	equipOnInteract, err := buf.ReadBool()
	if err != nil {
		return nil, err
	}
	canBeSheared, err := buf.ReadBool()
	if err != nil {
		return nil, err
	}
	shearingSound, err := readConsumableSound(buf)
	if err != nil {
		return nil, err
	}
	return &Equippable{
		Slot:       equipmentSlots[slotID],
		EquipSound: equipSound.name, EquipSoundDirect: equipSound.direct, EquipSoundRange: equipSound.fixedRange,
		AssetID: assetID, CameraOverlay: cameraOverlay, AllowedEntities: allowedEntities,
		Dispensable: bool(dispensable), Swappable: bool(swappable), DamageOnHurt: bool(damageOnHurt),
		EquipOnInteract: bool(equipOnInteract), CanBeSheared: bool(canBeSheared),
		ShearingSound: shearingSound.name, ShearingSoundDirect: shearingSound.direct, ShearingSoundRange: shearingSound.fixedRange,
	}, nil
}

func writeEquippable(component *Equippable) ([]byte, error) {
	w := ns.NewWriter()
	slotID := slices.Index(equipmentSlots, component.Slot)
	if slotID < 0 {
		return nil, fmt.Errorf("unknown equipment slot %q", component.Slot)
	}
	w.WriteVarInt(ns.VarInt(slotID))
	equipSound := component.EquipSound
	if equipSound == "" {
		equipSound = "minecraft:item.armor.equip_generic"
	}
	if err := writeConsumableSound(w, equipSound, component.EquipSoundDirect, component.EquipSoundRange); err != nil {
		return nil, err
	}
	writeOptionalIdentifier(w, component.AssetID)
	writeOptionalIdentifier(w, component.CameraOverlay)
	if err := writeOptionalEntityTypeHolderSet(w, component.AllowedEntities); err != nil {
		return nil, err
	}
	w.WriteBool(ns.Boolean(component.Dispensable))
	w.WriteBool(ns.Boolean(component.Swappable))
	w.WriteBool(ns.Boolean(component.DamageOnHurt))
	w.WriteBool(ns.Boolean(component.EquipOnInteract))
	w.WriteBool(ns.Boolean(component.CanBeSheared))
	shearingSound := component.ShearingSound
	if shearingSound == "" {
		shearingSound = "minecraft:item.shears.snip"
	}
	if err := writeConsumableSound(w, shearingSound, component.ShearingSoundDirect, component.ShearingSoundRange); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func readOptionalIdentifier(buf *ns.PacketBuffer) (string, error) {
	present, err := buf.ReadBool()
	if err != nil || !present {
		return "", err
	}
	value, err := buf.ReadString(maxStringLen)
	return string(value), err
}

func writeOptionalIdentifier(w *ns.PacketBuffer, value string) {
	w.WriteBool(ns.Boolean(value != ""))
	if value != "" {
		w.WriteString(ns.String(value))
	}
}

func readOptionalEntityTypeHolderSet(buf *ns.PacketBuffer) ([]string, error) {
	present, err := buf.ReadBool()
	if err != nil || !present {
		return nil, err
	}
	return readEntityTypeHolderSet(buf)
}

func readEntityTypeHolderSet(buf *ns.PacketBuffer) ([]string, error) {
	marker, err := buf.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if marker < 0 {
		return nil, fmt.Errorf("invalid entity type holder set marker %d", marker)
	}
	if marker == 0 {
		tag, err := buf.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		return []string{"#" + string(tag)}, nil
	}
	result := make([]string, 0, marker-1)
	for range int(marker) - 1 {
		id, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		name := registries.EntityType.ByID(int32(id))
		if name == "" {
			return nil, fmt.Errorf("unknown entity type %d", id)
		}
		result = append(result, name)
	}
	return result, nil
}

func writeOptionalEntityTypeHolderSet(w *ns.PacketBuffer, types []string) error {
	w.WriteBool(ns.Boolean(types != nil))
	if types == nil {
		return nil
	}
	return writeEntityTypeHolderSet(w, types)
}

func writeEntityTypeHolderSet(w *ns.PacketBuffer, types []string) error {
	if len(types) == 1 && strings.HasPrefix(types[0], "#") {
		w.WriteVarInt(0)
		w.WriteString(ns.String(strings.TrimPrefix(types[0], "#")))
		return nil
	}
	w.WriteVarInt(ns.VarInt(len(types) + 1))
	for _, name := range types {
		id := registries.EntityType.Get(name)
		if id < 0 {
			return fmt.Errorf("unknown entity type %q", name)
		}
		w.WriteVarInt(ns.VarInt(id))
	}
	return nil
}
