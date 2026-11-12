package items

import (
	"fmt"

	ns "github.com/go-mclib/protocol/javaproto/wire"

	"github.com/go-mclib/data/registries"
)

func readBlocksAttacks(buf *ns.PacketBuffer) (*BlocksAttacks, error) {
	blockDelay, err := buf.ReadFloat32()
	if err != nil {
		return nil, err
	}
	disableScale, err := buf.ReadFloat32()
	if err != nil {
		return nil, err
	}
	count, err := buf.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if count < 0 {
		return nil, fmt.Errorf("negative damage reduction count %d", count)
	}
	reductions := makeCapped[DamageReduction](count)
	for range int(count) {
		angle, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		hasTypes, err := buf.ReadBool()
		if err != nil {
			return nil, err
		}
		var types []string
		if hasTypes {
			types, err = readHolderSet(buf, "damage type", damageTypes.ByID)
			if err != nil {
				return nil, err
			}
		}
		base, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		factor, err := buf.ReadFloat32()
		if err != nil {
			return nil, err
		}
		reductions = append(reductions, DamageReduction{
			HorizontalBlockingAngle: float64(angle), Types: types, Base: float64(base), Factor: float64(factor),
		})
	}
	threshold, err := buf.ReadFloat32()
	if err != nil {
		return nil, err
	}
	base, err := buf.ReadFloat32()
	if err != nil {
		return nil, err
	}
	factor, err := buf.ReadFloat32()
	if err != nil {
		return nil, err
	}
	bypassedBy, err := readOptionalHolderSet(buf, "damage type", damageTypes.ByID)
	if err != nil {
		return nil, err
	}
	blockSound, err := readOptionalSoundHolder(buf)
	if err != nil {
		return nil, err
	}
	disabledSound, err := readOptionalSoundHolder(buf)
	if err != nil {
		return nil, err
	}
	return &BlocksAttacks{
		BlockDelaySeconds: float64(blockDelay), DisableCooldownScale: float64(disableScale), DamageReductions: reductions,
		ItemDamage: DamageSpec{Threshold: float64(threshold), Base: float64(base), Factor: float64(factor)}, BypassedBy: bypassedBy,
		BlockSound: blockSound.name, BlockSoundDirect: blockSound.direct, BlockSoundRange: blockSound.fixedRange,
		DisabledSound: disabledSound.name, DisabledSoundDirect: disabledSound.direct, DisabledSoundRange: disabledSound.fixedRange,
	}, nil
}

func writeBlocksAttacks(w *ns.PacketBuffer, component *BlocksAttacks) error {
	w.WriteFloat32(ns.Float32(component.BlockDelaySeconds))
	w.WriteFloat32(ns.Float32(component.DisableCooldownScale))
	w.WriteVarInt(ns.VarInt(len(component.DamageReductions)))
	for _, reduction := range component.DamageReductions {
		w.WriteFloat32(ns.Float32(reduction.HorizontalBlockingAngle))
		w.WriteBool(ns.Boolean(reduction.Types != nil))
		if reduction.Types != nil {
			if err := writeHolderSet(w, "damage type", reduction.Types, damageTypes.Get); err != nil {
				return err
			}
		}
		w.WriteFloat32(ns.Float32(reduction.Base))
		w.WriteFloat32(ns.Float32(reduction.Factor))
	}
	w.WriteFloat32(ns.Float32(component.ItemDamage.Threshold))
	w.WriteFloat32(ns.Float32(component.ItemDamage.Base))
	w.WriteFloat32(ns.Float32(component.ItemDamage.Factor))
	if err := writeOptionalHolderSet(w, "damage type", component.BypassedBy, damageTypes.Get); err != nil {
		return err
	}
	if err := writeOptionalSoundHolder(w, component.BlockSound, component.BlockSoundDirect, component.BlockSoundRange); err != nil {
		return err
	}
	return writeOptionalSoundHolder(w, component.DisabledSound, component.DisabledSoundDirect, component.DisabledSoundRange)
}

var damageTypes = registries.SynchronizedEntries["minecraft:damage_type"]

func readOptionalSoundHolder(buf *ns.PacketBuffer) (soundHolder, error) {
	present, err := buf.ReadBool()
	if err != nil || !present {
		return soundHolder{}, err
	}
	return readSoundHolder(buf)
}

func writeOptionalSoundHolder(w *ns.PacketBuffer, name string, direct bool, fixedRange *float64) error {
	present := name != ""
	w.WriteBool(ns.Boolean(present))
	if !present {
		return nil
	}
	return writeSoundHolder(w, name, direct, fixedRange)
}
