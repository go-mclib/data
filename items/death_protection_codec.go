package items

import (
	"fmt"
	"reflect"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

type deathProtectionCodec struct{}

func (deathProtectionCodec) DecodeWire(buf *ns.PacketBuffer) ([]byte, error) {
	w := ns.NewWriter()
	if err := decodeDeathProtectionWire(buf, w); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (deathProtectionCodec) Apply(c *Components, data []byte) error {
	buf := ns.NewReader(data)
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	if count < 0 {
		return fmt.Errorf("negative death effect count %d", count)
	}
	effects := make([]ConsumeEffect, 0, count)
	for range int(count) {
		effect, err := readConsumeEffect(buf)
		if err != nil {
			return err
		}
		effects = append(effects, effect)
	}
	c.DeathProtection = &DeathProtection{DeathEffects: effects}
	return nil
}

func (deathProtectionCodec) Clear(c *Components) { c.DeathProtection = nil }

func (deathProtectionCodec) Differs(c, defaults *Components) (bool, bool) {
	return !reflect.DeepEqual(c.DeathProtection, defaults.DeathProtection), c.DeathProtection != nil
}

func (deathProtectionCodec) Encode(c *Components) ([]byte, error) {
	if c.DeathProtection == nil {
		return nil, fmt.Errorf("death_protection component is nil")
	}
	w := ns.NewWriter()
	w.WriteVarInt(ns.VarInt(len(c.DeathProtection.DeathEffects)))
	for _, effect := range c.DeathProtection.DeathEffects {
		if err := writeConsumeEffect(w, effect); err != nil {
			return nil, err
		}
	}
	return w.Bytes(), nil
}
