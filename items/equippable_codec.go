package items

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

var equipmentSlots = []string{"mainhand", "offhand", "feet", "legs", "chest", "head", "body", "saddle"}

func readEquippable(buf *ns.PacketBuffer) (*Equippable, error) {
	slotID, err := buf.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if slotID < 0 || int(slotID) >= len(equipmentSlots) {
		return nil, fmt.Errorf("unknown equipment slot %d", slotID)
	}
	equipSound, err := readSoundHolder(buf)
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
	allowedEntities, err := readOptionalHolderSet(buf, "entity type", registries.EntityType.ByID)
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
	shearingSound, err := readSoundHolder(buf)
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

func writeEquippable(w *ns.PacketBuffer, component *Equippable) error {
	slotID := slices.Index(equipmentSlots, component.Slot)
	if slotID < 0 {
		return fmt.Errorf("unknown equipment slot %q", component.Slot)
	}
	w.WriteVarInt(ns.VarInt(slotID))
	equipSound := cmp.Or(component.EquipSound, "minecraft:item.armor.equip_generic")
	if err := writeSoundHolder(w, equipSound, component.EquipSoundDirect, component.EquipSoundRange); err != nil {
		return err
	}
	writeOptionalIdentifier(w, component.AssetID)
	writeOptionalIdentifier(w, component.CameraOverlay)
	if err := writeOptionalHolderSet(w, "entity type", component.AllowedEntities, registries.EntityType.Get); err != nil {
		return err
	}
	w.WriteBool(ns.Boolean(component.Dispensable))
	w.WriteBool(ns.Boolean(component.Swappable))
	w.WriteBool(ns.Boolean(component.DamageOnHurt))
	w.WriteBool(ns.Boolean(component.EquipOnInteract))
	w.WriteBool(ns.Boolean(component.CanBeSheared))
	shearingSound := cmp.Or(component.ShearingSound, "minecraft:item.shears.snip")
	return writeSoundHolder(w, shearingSound, component.ShearingSoundDirect, component.ShearingSoundRange)
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
