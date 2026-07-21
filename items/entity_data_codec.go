package items

import (
	"fmt"
	"reflect"

	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// EntityDataNBT decodes a typed entity-data component into its NBT compound.
func EntityDataNBT(data *TypedEntityData) (nbt.Compound, error) {
	if data == nil {
		return nil, fmt.Errorf("entity data is nil")
	}
	if len(data.Raw) == 0 {
		return nbt.Compound{"id": nbt.String(data.Type)}, nil
	}
	compound, err := typedEntityDataNBT(data.Raw, registries.EntityType)
	if err != nil {
		return nil, err
	}
	if id := compound.GetString("id"); id != data.Type {
		return nil, fmt.Errorf("entity data type %q does not match %q", id, data.Type)
	}
	return compound, nil
}

type entityDataCodec struct{}

func (entityDataCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := decodeTypedEntityDataWire(buf, w); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (entityDataCodec) Apply(c *Components, data []byte) error {
	compound, err := typedEntityDataNBT(data, registries.EntityType)
	if err != nil {
		return err
	}
	c.EntityData = &TypedEntityData{Type: compound.GetString("id"), Raw: append([]byte(nil), data...)}
	return nil
}

func (entityDataCodec) Clear(c *Components) { c.EntityData = nil }

func (entityDataCodec) Differs(c, defaults *Components) (bool, bool) {
	return !reflect.DeepEqual(c.EntityData, defaults.EntityData), c.EntityData != nil
}

func (entityDataCodec) Encode(c *Components) ([]byte, error) {
	if c.EntityData == nil {
		return nil, fmt.Errorf("entity_data component is nil")
	}
	typeID := registries.EntityType.Get(c.EntityData.Type)
	if typeID < 0 {
		return nil, fmt.Errorf("unknown entity type %q", c.EntityData.Type)
	}
	payload := nbt.Compound{}
	if len(c.EntityData.Raw) > 0 {
		var err error
		payload, err = typedEntityDataNBT(c.EntityData.Raw, registries.EntityType)
		if err != nil {
			return nil, err
		}
		delete(payload, "id")
	}
	w := ns.NewWriter()
	w.WriteVarInt(ns.VarInt(typeID))
	if err := nbt.NewWriterTo(w.Writer()).WriteTag(payload, "", true); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
