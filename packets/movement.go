package packets

import (
	"fmt"
	"math"

	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// VecDelta is a linear displacement, or a sequence of timed displacements when
// Steps is nonempty. Each stepped displacement is relative to the previous step,
// not the initial position. Coordinates use the protocol's 1/4096 block units.
type VecDelta struct {
	X     ns.Int16
	Y     ns.Int16
	Z     ns.Int16
	Steps []VecDeltaStep
}

type VecDeltaStep struct {
	Ticks ns.VarInt
	X     ns.Int16
	Y     ns.Int16
	Z     ns.Int16
}

func readVecDelta(buf *ns.PacketBuffer) (VecDelta, ns.Boolean, error) {
	var delta VecDelta
	properties, err := buf.ReadVarInt()
	if err != nil {
		return delta, false, err
	}
	if properties < 0 {
		return delta, false, fmt.Errorf("negative movement properties %d", properties)
	}
	onGround := ns.Boolean(properties&1 != 0)
	count := properties >> 1
	if count == 0 {
		delta.X, delta.Y, delta.Z, err = readDeltaCoordinates(buf)
		return delta, onGround, err
	}
	// Append only completely decoded steps: an untrusted count never controls
	// an allocation, including when the reader does not expose its length.
	for range int(count) {
		var step VecDeltaStep
		if step.Ticks, err = buf.ReadVarInt(); err != nil {
			return delta, onGround, err
		}
		if step.X, step.Y, step.Z, err = readDeltaCoordinates(buf); err != nil {
			return delta, onGround, err
		}
		delta.Steps = append(delta.Steps, step)
	}
	return delta, onGround, nil
}

func readDeltaCoordinates(buf *ns.PacketBuffer) (x, y, z ns.Int16, err error) {
	if x, err = buf.ReadInt16(); err != nil {
		return
	}
	if y, err = buf.ReadInt16(); err != nil {
		return
	}
	z, err = buf.ReadInt16()
	return
}

func writeVecDelta(buf *ns.PacketBuffer, delta VecDelta, onGround ns.Boolean) error {
	if len(delta.Steps) > (1<<30)-1 {
		return fmt.Errorf("too many movement steps: %d", len(delta.Steps))
	}
	properties := ns.VarInt(len(delta.Steps)) << 1
	if onGround {
		properties |= 1
	}
	if err := buf.WriteVarInt(properties); err != nil {
		return err
	}
	if len(delta.Steps) == 0 {
		return writeDeltaCoordinates(buf, delta.X, delta.Y, delta.Z)
	}
	for _, step := range delta.Steps {
		if err := buf.WriteVarInt(step.Ticks); err != nil {
			return err
		}
		if err := writeDeltaCoordinates(buf, step.X, step.Y, step.Z); err != nil {
			return err
		}
	}
	return nil
}

func writeDeltaCoordinates(buf *ns.PacketBuffer, x, y, z ns.Int16) error {
	if err := buf.WriteInt16(x); err != nil {
		return err
	}
	if err := buf.WriteInt16(y); err != nil {
		return err
	}
	return buf.WriteInt16(z)
}

// PositionPath is a linear absolute position when Steps is nil. A nonnil Steps
// selects the stepped encoding and must contain at least one position. X/Y/Z
// are used only for a linear path. This packet carries no velocity.
type PositionPath struct {
	X     ns.Float64
	Y     ns.Float64
	Z     ns.Float64
	Steps []PositionStep
}

type PositionStep struct {
	X          ns.Float64
	Y          ns.Float64
	Z          ns.Float64
	TickOffset ns.VarInt
}

func (p *PositionPath) Read(buf *ns.PacketBuffer) error {
	*p = PositionPath{}
	kind, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	switch kind {
	case 0:
		p.X, p.Y, p.Z, err = readPathCoordinates(buf)
		return err
	case 1:
		count, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		if count <= 0 {
			return fmt.Errorf("invalid position path step count %d", count)
		}
		for range int(count) {
			var step PositionStep
			if step.X, step.Y, step.Z, err = readPathCoordinates(buf); err != nil {
				return err
			}
			if step.TickOffset, err = buf.ReadVarInt(); err != nil {
				return err
			}
			p.Steps = append(p.Steps, step)
		}
		return nil
	default:
		return fmt.Errorf("unknown position path type %d", kind)
	}
}

func (p PositionPath) Write(buf *ns.PacketBuffer) error {
	if p.Steps == nil {
		if err := buf.WriteVarInt(0); err != nil {
			return err
		}
		return writePathCoordinates(buf, p.X, p.Y, p.Z)
	}
	if len(p.Steps) == 0 || len(p.Steps) > math.MaxInt32 {
		return fmt.Errorf("invalid position path step count %d", len(p.Steps))
	}
	if err := buf.WriteVarInt(1); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(len(p.Steps))); err != nil {
		return err
	}
	for _, step := range p.Steps {
		if err := writePathCoordinates(buf, step.X, step.Y, step.Z); err != nil {
			return err
		}
		if err := buf.WriteVarInt(step.TickOffset); err != nil {
			return err
		}
	}
	return nil
}

func readPathCoordinates(buf *ns.PacketBuffer) (x, y, z ns.Float64, err error) {
	if x, err = buf.ReadFloat64(); err != nil {
		return
	}
	if y, err = buf.ReadFloat64(); err != nil {
		return
	}
	z, err = buf.ReadFloat64()
	return
}

func writePathCoordinates(buf *ns.PacketBuffer, x, y, z ns.Float64) error {
	if err := buf.WriteFloat64(x); err != nil {
		return err
	}
	if err := buf.WriteFloat64(y); err != nil {
		return err
	}
	return buf.WriteFloat64(z)
}
