package entities

import (
	"fmt"

	"github.com/go-mclib/data/items"
	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// Rotations represents 3D rotation angles (x, y, z in degrees).
type Rotations struct {
	X, Y, Z float32
}

// Position represents a block position (packed 64-bit format).
type Position struct {
	X, Y, Z int32
}

// Metadata represents raw entity metadata entries.
// Use this for passthrough or when you don't need typed access.
type Metadata []MetadataEntry

// ReadMetadata reads entity metadata from a packet buffer.
// Wire format: [Index(UByte)][Type(VarInt)][Value]...[0xFF terminator]
func ReadMetadata(buf *ns.PacketBuffer) (Metadata, error) {
	var entries Metadata

	for {
		index, err := buf.ReadByte()
		if err != nil {
			return nil, err
		}
		if index == 0xFF {
			break
		}

		serializerID, err := buf.ReadVarInt()
		if err != nil {
			return nil, err
		}

		data, err := readSerializerValue(buf, int32(serializerID))
		if err != nil {
			return nil, fmt.Errorf("reading metadata index %d (serializer %d): %w", index, serializerID, err)
		}

		entries = append(entries, MetadataEntry{
			Index:      index,
			Serializer: int32(serializerID),
			Data:       data,
		})
	}

	return entries, nil
}

// WriteMetadata writes entity metadata to a packet buffer.
func WriteMetadata(buf *ns.PacketBuffer, entries Metadata) error {
	for _, entry := range entries {
		buf.WriteByte(entry.Index)
		buf.WriteVarInt(ns.VarInt(entry.Serializer))
		if _, err := buf.Write(entry.Data); err != nil {
			return err
		}
	}
	buf.WriteByte(0xFF) // terminator
	return nil
}

// Get returns the data for a specific metadata index, or nil if not present.
func (m Metadata) Get(index byte) []byte {
	for _, entry := range m {
		if entry.Index == index {
			return entry.Data
		}
	}
	return nil
}

// Set sets or adds a metadata entry.
func (m *Metadata) Set(index byte, serializerID int32, data []byte) {
	for i, entry := range *m {
		if entry.Index == index {
			(*m)[i].Serializer = serializerID
			(*m)[i].Data = data
			return
		}
	}
	*m = append(*m, MetadataEntry{
		Index:      index,
		Serializer: serializerID,
		Data:       data,
	})
}

// wireCopier copies one metadata value of a given wire type from buf to w.
type wireCopier func(buf, w *ns.PacketBuffer) error

// wireCopiers maps serializer wire types to their copy handlers.
var wireCopiers = map[string]wireCopier{
	"byte":        func(buf, w *ns.PacketBuffer) error { return w.CopyInt8(buf) },
	"varint":      copyVarIntValue,
	"varlong":     func(buf, w *ns.PacketBuffer) error { return w.CopyVarLong(buf) },
	"float32":     copyFloat32Value,
	"string":      copyStringValue,
	"bool":        func(buf, w *ns.PacketBuffer) error { return w.CopyBool(buf) },
	"nbt":         copyNBTValue,
	"slot":        copySlotValue,
	"position":    copyPositionValue,
	"rotations":   copyNValues(3, copyFloat32Value),
	"vector3f":    copyNValues(3, copyFloat32Value),
	"quaternionf": copyNValues(4, copyFloat32Value),
	// type, profession, level (all VarInt)
	"villager_data": copyNValues(3, copyVarIntValue),
	// varint with 0 = absent, 1+ = value+1
	"optional_varint":     copyVarIntValue,
	"optional_nbt":        copyOptionalValue(copyNBTValue),
	"optional_position":   copyOptionalValue(copyPositionValue),
	"optional_uuid":       copyOptionalValue(func(buf, w *ns.PacketBuffer) error { return w.CopyUUID(buf) }),
	"optional_global_pos": copyOptionalValue(copyGlobalPos),
	"particle":            CopyParticle,
	"particle_list":       copyParticleList,
	"id_or_inline":        copyPaintingVariant,
	"resolvable_profile":  copyResolvableProfile,
	"empty":               func(buf, w *ns.PacketBuffer) error { return nil },
}

// readSerializerValue reads a value based on serializer type.
func readSerializerValue(buf *ns.PacketBuffer, serializerID int32) ([]byte, error) {
	wireType, ok := serializerWireTypes[serializerID]
	if !ok {
		return nil, fmt.Errorf("unknown serializer ID %d", serializerID)
	}

	copier, ok := wireCopiers[wireType]
	if !ok {
		return nil, fmt.Errorf("unhandled wire type %q for serializer %d", wireType, serializerID)
	}

	w := ns.NewWriter()
	if err := copier(buf, w); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func copyVarIntValue(buf, w *ns.PacketBuffer) error {
	return w.CopyVarInt(buf)
}

func copyFloat32Value(buf, w *ns.PacketBuffer) error {
	return w.CopyFloat32(buf)
}

func copyStringValue(buf, w *ns.PacketBuffer) error {
	return w.CopyString(buf, 32767)
}

func copyNBTValue(buf, w *ns.PacketBuffer) error {
	return nbt.Copy(w.Writer(), buf.Reader(), true)
}

func copyPositionValue(buf, w *ns.PacketBuffer) error {
	return w.CopyPosition(buf)
}

// copyNValues returns a copier that applies fn n times.
func copyNValues(n int, fn wireCopier) wireCopier {
	return func(buf, w *ns.PacketBuffer) error {
		for range n {
			if err := fn(buf, w); err != nil {
				return err
			}
		}
		return nil
	}
}

// copyOptionalValue returns a copier for a bool-prefixed optional value.
func copyOptionalValue(fn wireCopier) wireCopier {
	return func(buf, w *ns.PacketBuffer) error {
		present, err := buf.ReadBool()
		if err != nil {
			return err
		}
		w.WriteBool(present)
		if present {
			return fn(buf, w)
		}
		return nil
	}
}

// copySlotValue reads a slot with the items decoder, then re-encodes it to bytes.
func copySlotValue(buf, w *ns.PacketBuffer) error {
	slot, err := buf.ReadSlot(items.Decoder())
	if err != nil {
		return err
	}
	return w.WriteSlot(slot)
}

// copyGlobalPos copies a dimension identifier followed by a position.
func copyGlobalPos(buf, w *ns.PacketBuffer) error {
	if err := copyStringValue(buf, w); err != nil {
		return err
	}
	return copyPositionValue(buf, w)
}

// CopyParticle copies a particle from buf to w.
func CopyParticle(buf, w *ns.PacketBuffer) error {
	particleType, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(particleType)

	// Registry-dispatched ParticleOptions.STREAM_CODEC in vanilla 26.2.
	// Even consumers that only need flags must consume every particle payload:
	// otherwise its bytes become bogus metadata entries or an early terminator.
	switch name := registries.ParticleType.ByID(int32(particleType)); name {
	case "":
		return fmt.Errorf("unknown particle type %d", particleType)
	case "minecraft:block", "minecraft:block_marker", "minecraft:falling_dust", "minecraft:dust_pillar", "minecraft:block_crumble", "minecraft:shriek":
		return copyVarIntValue(buf, w)
	case "minecraft:entity_effect", "minecraft:tinted_leaves", "minecraft:flash", "minecraft:geyser", "minecraft:geyser_plume":
		return w.CopyInt32(buf)
	case "minecraft:dust", "minecraft:effect", "minecraft:instant_effect", "minecraft:geyser_base", "minecraft:geyser_poof":
		if err := w.CopyInt32(buf); err != nil {
			return err
		}
		return copyFloat32Value(buf, w)
	case "minecraft:dragon_breath", "minecraft:sculk_charge":
		return copyFloat32Value(buf, w)
	case "minecraft:dust_color_transition":
		for range 2 {
			if err := w.CopyInt32(buf); err != nil {
				return err
			}
		}
		return copyFloat32Value(buf, w)
	case "minecraft:trail":
		for range 3 {
			value, err := buf.ReadFloat64()
			if err != nil {
				return err
			}
			if err := w.WriteFloat64(value); err != nil {
				return err
			}
		}
		if err := w.CopyInt32(buf); err != nil {
			return err
		}
		return copyVarIntValue(buf, w)
	case "minecraft:item":
		stack, err := items.ReadItemStackTemplate(buf)
		if err != nil {
			return err
		}
		return items.WriteItemStackTemplate(w, stack)
	case "minecraft:vibration":
		source, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		w.WriteVarInt(source)
		switch registries.PositionSourceType.ByID(int32(source)) {
		case "minecraft:block":
			if err := copyPositionValue(buf, w); err != nil {
				return err
			}
		case "minecraft:entity":
			if err := copyVarIntValue(buf, w); err != nil {
				return err
			}
			if err := copyFloat32Value(buf, w); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown vibration position source %d", source)
		}
		return copyVarIntValue(buf, w)
	default:
		return nil // Registered SimpleParticleType: no option payload.
	}
}

func copyParticleList(buf, w *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	if count < 0 {
		return fmt.Errorf("negative particle count %d", count)
	}
	w.WriteVarInt(count)
	for range int(count) {
		if err := CopyParticle(buf, w); err != nil {
			return err
		}
	}
	return nil
}

// copyPaintingVariant copies a painting variant: either a registry ID or inline data.
func copyPaintingVariant(buf, w *ns.PacketBuffer) error {
	typeID, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(typeID)
	if typeID != 0 {
		return nil
	}

	// inline: asset_id, width, height
	if err := copyStringValue(buf, w); err != nil {
		return err
	}
	for range 2 {
		if err := copyVarIntValue(buf, w); err != nil {
			return err
		}
	}
	// optional title and author (NBT)
	for range 2 {
		if err := copyOptionalValue(copyNBTValue)(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func copyResolvableProfile(buf, w *ns.PacketBuffer) error {
	// optional name, optional uuid
	if err := copyOptionalValue(copyStringValue)(buf, w); err != nil {
		return err
	}
	if err := copyOptionalValue(func(buf, w *ns.PacketBuffer) error { return w.CopyUUID(buf) })(buf, w); err != nil {
		return err
	}
	// properties: name, value, optional signature
	propCount, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	w.WriteVarInt(propCount)
	for range int(propCount) {
		if err := copyStringValue(buf, w); err != nil {
			return err
		}
		if err := copyStringValue(buf, w); err != nil {
			return err
		}
		if err := copyOptionalValue(copyStringValue)(buf, w); err != nil {
			return err
		}
	}
	return nil
}
