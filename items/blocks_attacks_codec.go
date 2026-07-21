package items

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

type blocksAttacksCodec struct{}

func (blocksAttacksCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	component, err := readBlocksAttacks(buf)
	if err != nil {
		return nil, err
	}
	return writeBlocksAttacks(component)
}

func (blocksAttacksCodec) Apply(c *Components, data []byte) error {
	component, err := readBlocksAttacks(ns.NewReader(data))
	if err != nil {
		return err
	}
	c.BlocksAttacks = component
	return nil
}

func (blocksAttacksCodec) Clear(c *Components) { c.BlocksAttacks = nil }

func (blocksAttacksCodec) Differs(c, defaults *Components) (bool, bool) {
	return !reflect.DeepEqual(c.BlocksAttacks, defaults.BlocksAttacks), c.BlocksAttacks != nil
}

func (blocksAttacksCodec) Encode(c *Components) ([]byte, error) {
	if c.BlocksAttacks == nil {
		return nil, fmt.Errorf("blocks_attacks component is nil")
	}
	return writeBlocksAttacks(c.BlocksAttacks)
}

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
	reductions := make([]DamageReduction, 0, count)
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
			types, err = readDamageTypeHolderSet(buf)
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
	bypassedBy, err := readOptionalDamageTypeHolderSet(buf)
	if err != nil {
		return nil, err
	}
	blockSound, err := readOptionalComponentSound(buf)
	if err != nil {
		return nil, err
	}
	disabledSound, err := readOptionalComponentSound(buf)
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

func writeBlocksAttacks(component *BlocksAttacks) ([]byte, error) {
	w := ns.NewWriter()
	w.WriteFloat32(ns.Float32(component.BlockDelaySeconds))
	w.WriteFloat32(ns.Float32(component.DisableCooldownScale))
	w.WriteVarInt(ns.VarInt(len(component.DamageReductions)))
	for _, reduction := range component.DamageReductions {
		w.WriteFloat32(ns.Float32(reduction.HorizontalBlockingAngle))
		w.WriteBool(ns.Boolean(reduction.Types != nil))
		if reduction.Types != nil {
			if err := writeDamageTypeHolderSet(w, reduction.Types); err != nil {
				return nil, err
			}
		}
		w.WriteFloat32(ns.Float32(reduction.Base))
		w.WriteFloat32(ns.Float32(reduction.Factor))
	}
	w.WriteFloat32(ns.Float32(component.ItemDamage.Threshold))
	w.WriteFloat32(ns.Float32(component.ItemDamage.Base))
	w.WriteFloat32(ns.Float32(component.ItemDamage.Factor))
	if err := writeOptionalDamageTypeHolderSet(w, component.BypassedBy); err != nil {
		return nil, err
	}
	if err := writeOptionalComponentSound(w, component.BlockSound, component.BlockSoundDirect, component.BlockSoundRange); err != nil {
		return nil, err
	}
	if err := writeOptionalComponentSound(w, component.DisabledSound, component.DisabledSoundDirect, component.DisabledSoundRange); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func readDamageTypeHolderSet(buf *ns.PacketBuffer) ([]string, error) {
	marker, err := buf.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if marker < 0 {
		return nil, fmt.Errorf("invalid damage type holder set marker %d", marker)
	}
	if marker == 0 {
		tag, err := buf.ReadString(maxStringLen)
		if err != nil {
			return nil, err
		}
		return []string{"#" + string(tag)}, nil
	}
	entries := registries.SynchronizedEntries["minecraft:damage_type"]
	result := make([]string, 0, marker-1)
	for range int(marker) - 1 {
		id, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}
		if id < 0 || int(id) >= len(entries) {
			return nil, fmt.Errorf("unknown damage type %d", id)
		}
		result = append(result, entries[id])
	}
	return result, nil
}

func writeDamageTypeHolderSet(w *ns.PacketBuffer, types []string) error {
	if len(types) == 1 && strings.HasPrefix(types[0], "#") {
		w.WriteVarInt(0)
		return w.WriteString(ns.String(strings.TrimPrefix(types[0], "#")))
	}
	entries := registries.SynchronizedEntries["minecraft:damage_type"]
	w.WriteVarInt(ns.VarInt(len(types) + 1))
	for _, name := range types {
		id := slices.Index(entries, name)
		if id < 0 {
			return fmt.Errorf("unknown damage type %q", name)
		}
		w.WriteVarInt(ns.VarInt(id))
	}
	return nil
}

func readOptionalDamageTypeHolderSet(buf *ns.PacketBuffer) ([]string, error) {
	present, err := buf.ReadBool()
	if err != nil || !present {
		return nil, err
	}
	return readDamageTypeHolderSet(buf)
}

func writeOptionalDamageTypeHolderSet(w *ns.PacketBuffer, types []string) error {
	w.WriteBool(ns.Boolean(types != nil))
	if types == nil {
		return nil
	}
	return writeDamageTypeHolderSet(w, types)
}

func readOptionalComponentSound(buf *ns.PacketBuffer) (consumableSound, error) {
	present, err := buf.ReadBool()
	if err != nil || !present {
		return consumableSound{}, err
	}
	return readConsumableSound(buf)
}

func writeOptionalComponentSound(w *ns.PacketBuffer, name string, direct bool, fixedRange *float64) error {
	present := name != ""
	w.WriteBool(ns.Boolean(present))
	if !present {
		return nil
	}
	return writeConsumableSound(w, name, direct, fixedRange)
}
