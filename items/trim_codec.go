package items

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/go-mclib/data/registries"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

type trimCodec struct{}

func (trimCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := decodeTrimWire(buf, w); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (trimCodec) Apply(c *Components, data []byte) error {
	buf := ns.NewReader(data)
	material, err := readTrimHolder(buf, "minecraft:trim_material")
	if err != nil {
		return err
	}
	pattern, err := readTrimHolder(buf, "minecraft:trim_pattern")
	if err != nil {
		return err
	}
	c.Trim = &ArmorTrim{Material: material, Pattern: pattern}
	return nil
}

func (trimCodec) Clear(c *Components) { c.Trim = nil }

func (trimCodec) Differs(c, defaults *Components) (bool, bool) {
	return !reflect.DeepEqual(c.Trim, defaults.Trim), c.Trim != nil
}

func (trimCodec) Encode(c *Components) ([]byte, error) {
	if c.Trim == nil {
		return nil, fmt.Errorf("trim component is nil")
	}
	w := ns.NewWriter()
	if err := writeTrimHolder(w, "minecraft:trim_material", c.Trim.Material); err != nil {
		return nil, err
	}
	if err := writeTrimHolder(w, "minecraft:trim_pattern", c.Trim.Pattern); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func readTrimHolder(buf *ns.PacketBuffer, registry string) (string, error) {
	marker, err := buf.ReadVarInt()
	if err != nil {
		return "", err
	}
	if marker == 0 {
		return "", fmt.Errorf("direct %s holder is not supported", registry)
	}
	entries := registries.SynchronizedEntries[registry]
	id := int(marker) - 1
	if id < 0 || id >= len(entries) {
		return "", fmt.Errorf("invalid %s holder %d", registry, marker)
	}
	return entries[id], nil
}

func writeTrimHolder(buf *ns.PacketBuffer, registry, name string) error {
	id := slices.Index(registries.SynchronizedEntries[registry], name)
	if id < 0 {
		return fmt.Errorf("unknown %s entry %q", registry, name)
	}
	return buf.WriteVarInt(ns.VarInt(id + 1))
}
