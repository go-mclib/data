package entities

import ns "github.com/go-mclib/protocol/javaproto/wire"

// Constructors for single metadata entries: the write half of the
// MetadataEntry codec (ReadMetadata decodes the same wire values).

// ByteEntry builds a BYTE metadata entry.
func ByteEntry(index byte, v byte) MetadataEntry {
	return MetadataEntry{Index: index, Serializer: SerializerBYTE, Data: []byte{v}}
}

// VarIntEntry builds a VarInt-encoded metadata entry with the given serializer id.
func VarIntEntry(index byte, serializer int32, v int32) MetadataEntry {
	buf := ns.NewWriter()
	buf.WriteVarInt(ns.VarInt(v))
	return MetadataEntry{Index: index, Serializer: serializer, Data: buf.Bytes()}
}

// IntEntry builds a VarInt-backed INT metadata entry.
func IntEntry(index byte, v int32) MetadataEntry {
	return VarIntEntry(index, SerializerINT, v)
}

// StringEntry builds a STRING metadata entry.
func StringEntry(index byte, value string) MetadataEntry {
	buf := ns.NewWriter()
	buf.WriteString(ns.String(value))
	return MetadataEntry{Index: index, Serializer: SerializerSTRING, Data: buf.Bytes()}
}

// ComponentEntry builds a COMPONENT metadata entry.
func ComponentEntry(index byte, value ns.TextComponent) MetadataEntry {
	buf := ns.NewWriter()
	_ = value.Encode(buf)
	return MetadataEntry{Index: index, Serializer: SerializerCOMPONENT, Data: buf.Bytes()}
}

// OptionalComponentEntry builds an OPTIONAL_COMPONENT metadata entry (e.g.
// CustomName): a present flag followed by the text component when set.
func OptionalComponentEntry(index byte, value *ns.TextComponent) MetadataEntry {
	buf := ns.NewWriter()
	buf.WriteBool(value != nil)
	if value != nil {
		_ = value.Encode(buf)
	}
	return MetadataEntry{Index: index, Serializer: SerializerOPTIONAL_COMPONENT, Data: buf.Bytes()}
}

// BoolEntry builds a BOOLEAN metadata entry.
func BoolEntry(index byte, v bool) MetadataEntry {
	var b byte
	if v {
		b = 1
	}
	return MetadataEntry{Index: index, Serializer: SerializerBOOLEAN, Data: []byte{b}}
}

// FloatEntry builds a FLOAT metadata entry.
func FloatEntry(index byte, v float32) MetadataEntry {
	buf := ns.NewWriter()
	buf.WriteFloat32(ns.Float32(v))
	return MetadataEntry{Index: index, Serializer: SerializerFLOAT, Data: buf.Bytes()}
}

// RotationsEntry builds a ROTATIONS metadata entry.
func RotationsEntry(index byte, x, y, z float32) MetadataEntry {
	buf := ns.NewWriter()
	buf.WriteFloat32(ns.Float32(x))
	buf.WriteFloat32(ns.Float32(y))
	buf.WriteFloat32(ns.Float32(z))
	return MetadataEntry{Index: index, Serializer: SerializerROTATIONS, Data: buf.Bytes()}
}

// SlotEntry builds an ITEM_STACK metadata entry.
func SlotEntry(index byte, v ns.Slot) MetadataEntry {
	buf := ns.NewWriter()
	buf.WriteSlot(v)
	return MetadataEntry{Index: index, Serializer: SerializerITEM_STACK, Data: buf.Bytes()}
}
