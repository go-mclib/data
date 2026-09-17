package packets

import (
	"fmt"

	"github.com/go-mclib/data/items"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// PositionedAdvancement contains the complete holder followed by layout
// coordinates. Coordinates are present even when Display is nil.
type PositionedAdvancement struct {
	Id                  ns.Identifier
	Parent              ns.PrefixedOptional[ns.Identifier]
	Display             *AdvancementDisplay
	Requirements        [][]ns.String
	SendsTelemetryEvent ns.Boolean
	X                   ns.Float32
	Y                   ns.Float32
}

type AdvancementDisplay struct {
	Title       ns.TextComponent
	Description ns.TextComponent
	Icon        ns.Slot
	Frame       ns.VarInt
	Background  ns.PrefixedOptional[ns.Identifier]
	ShowToast   ns.Boolean
	Hidden      ns.Boolean
}

const (
	AdvancementFrameTask ns.VarInt = iota
	AdvancementFrameChallenge
	AdvancementFrameGoal
)

// AdvancementProgress is one identifier-keyed entry in the progress map.
type AdvancementProgress struct {
	Id       ns.Identifier
	Criteria []AdvancementCriterion
}

type AdvancementCriterion struct {
	Name ns.String
	// Obtained is an optional Unix timestamp in milliseconds.
	Obtained ns.PrefixedOptional[ns.Int64]
}

func readAdvancementCount(buf *ns.PacketBuffer) (int, error) {
	count, err := buf.ReadVarInt()
	if err != nil {
		return 0, err
	}
	if count < 0 {
		return 0, fmt.Errorf("negative advancement collection count %d", count)
	}
	return int(count), nil
}

func writeAdvancementCount(buf *ns.PacketBuffer, count int) error {
	if count > (1<<31)-1 {
		return fmt.Errorf("advancement collection too large: %d", count)
	}
	return buf.WriteVarInt(ns.VarInt(count))
}

func (p *S2CUpdateAdvancements) readAdvancements(buf *ns.PacketBuffer) error {
	var err error
	if p.Reset, err = buf.ReadBool(); err != nil {
		return err
	}
	count, err := readAdvancementCount(buf)
	if err != nil {
		return err
	}
	p.Added = nil
	for range count {
		var advancement PositionedAdvancement
		if err := advancement.read(buf); err != nil {
			return err
		}
		p.Added = append(p.Added, advancement)
	}
	count, err = readAdvancementCount(buf)
	if err != nil {
		return err
	}
	p.Removed = nil
	for range count {
		id, err := buf.ReadIdentifier()
		if err != nil {
			return err
		}
		p.Removed = append(p.Removed, id)
	}
	count, err = readAdvancementCount(buf)
	if err != nil {
		return err
	}
	p.Progress = nil
	for range count {
		var progress AdvancementProgress
		if progress.Id, err = buf.ReadIdentifier(); err != nil {
			return err
		}
		criterionCount, err := readAdvancementCount(buf)
		if err != nil {
			return err
		}
		for range criterionCount {
			var criterion AdvancementCriterion
			if criterion.Name, err = buf.ReadString(32767); err != nil {
				return err
			}
			if err := criterion.Obtained.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.Int64, error) {
				return b.ReadInt64()
			}); err != nil {
				return err
			}
			progress.Criteria = append(progress.Criteria, criterion)
		}
		p.Progress = append(p.Progress, progress)
	}
	p.ShowAdvancements, err = buf.ReadBool()
	return err
}

func (p *S2CUpdateAdvancements) writeAdvancements(buf *ns.PacketBuffer) error {
	if err := buf.WriteBool(p.Reset); err != nil {
		return err
	}
	if err := writeAdvancementCount(buf, len(p.Added)); err != nil {
		return err
	}
	for _, advancement := range p.Added {
		if err := advancement.write(buf); err != nil {
			return err
		}
	}
	if err := writeAdvancementCount(buf, len(p.Removed)); err != nil {
		return err
	}
	for _, id := range p.Removed {
		if err := buf.WriteIdentifier(id); err != nil {
			return err
		}
	}
	if err := writeAdvancementCount(buf, len(p.Progress)); err != nil {
		return err
	}
	for _, progress := range p.Progress {
		if err := buf.WriteIdentifier(progress.Id); err != nil {
			return err
		}
		if err := writeAdvancementCount(buf, len(progress.Criteria)); err != nil {
			return err
		}
		for _, criterion := range progress.Criteria {
			if err := buf.WriteString(criterion.Name); err != nil {
				return err
			}
			if err := criterion.Obtained.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.Int64) error {
				return b.WriteInt64(v)
			}); err != nil {
				return err
			}
		}
	}
	return buf.WriteBool(p.ShowAdvancements)
}

func (a *PositionedAdvancement) read(buf *ns.PacketBuffer) error {
	var err error
	if a.Id, err = buf.ReadIdentifier(); err != nil {
		return err
	}
	if err := a.Parent.DecodeWith(buf, func(b *ns.PacketBuffer) (ns.Identifier, error) {
		return b.ReadIdentifier()
	}); err != nil {
		return err
	}
	hasDisplay, err := buf.ReadBool()
	if err != nil {
		return err
	}
	a.Display = nil
	if hasDisplay {
		a.Display = &AdvancementDisplay{}
		if err := a.Display.read(buf); err != nil {
			return err
		}
	}
	count, err := readAdvancementCount(buf)
	if err != nil {
		return err
	}
	a.Requirements = nil
	for range count {
		criterionCount, err := readAdvancementCount(buf)
		if err != nil {
			return err
		}
		var requirement []ns.String
		for range criterionCount {
			criterion, err := buf.ReadString(32767)
			if err != nil {
				return err
			}
			requirement = append(requirement, criterion)
		}
		a.Requirements = append(a.Requirements, requirement)
	}
	if a.SendsTelemetryEvent, err = buf.ReadBool(); err != nil {
		return err
	}
	if a.X, err = buf.ReadFloat32(); err != nil {
		return err
	}
	a.Y, err = buf.ReadFloat32()
	return err
}

func (a PositionedAdvancement) write(buf *ns.PacketBuffer) error {
	if err := buf.WriteIdentifier(a.Id); err != nil {
		return err
	}
	if err := a.Parent.EncodeWith(buf, func(b *ns.PacketBuffer, v ns.Identifier) error {
		return b.WriteIdentifier(v)
	}); err != nil {
		return err
	}
	if err := buf.WriteBool(ns.Boolean(a.Display != nil)); err != nil {
		return err
	}
	if a.Display != nil {
		if err := a.Display.write(buf); err != nil {
			return err
		}
	}
	if err := writeAdvancementCount(buf, len(a.Requirements)); err != nil {
		return err
	}
	for _, requirement := range a.Requirements {
		if err := writeAdvancementCount(buf, len(requirement)); err != nil {
			return err
		}
		for _, criterion := range requirement {
			if err := buf.WriteString(criterion); err != nil {
				return err
			}
		}
	}
	if err := buf.WriteBool(a.SendsTelemetryEvent); err != nil {
		return err
	}
	if err := buf.WriteFloat32(a.X); err != nil {
		return err
	}
	return buf.WriteFloat32(a.Y)
}

func (d *AdvancementDisplay) read(buf *ns.PacketBuffer) error {
	var err error
	if d.Title, err = buf.ReadTextComponent(); err != nil {
		return err
	}
	if d.Description, err = buf.ReadTextComponent(); err != nil {
		return err
	}
	if d.Icon, err = items.ReadItemStackTemplate(buf); err != nil {
		return err
	}
	if d.Frame, err = buf.ReadVarInt(); err != nil {
		return err
	}
	flags, err := buf.ReadInt32()
	if err != nil {
		return err
	}
	d.Background = ns.PrefixedOptional[ns.Identifier]{}
	if flags&1 != 0 {
		background, err := buf.ReadIdentifier()
		if err != nil {
			return err
		}
		d.Background = ns.Some(background)
	}
	d.ShowToast = flags&2 != 0
	d.Hidden = flags&4 != 0
	return nil
}

func (d AdvancementDisplay) write(buf *ns.PacketBuffer) error {
	if err := buf.WriteTextComponent(d.Title); err != nil {
		return err
	}
	if err := buf.WriteTextComponent(d.Description); err != nil {
		return err
	}
	if err := items.WriteItemStackTemplate(buf, d.Icon); err != nil {
		return err
	}
	if err := buf.WriteVarInt(d.Frame); err != nil {
		return err
	}
	var flags ns.Int32
	if d.Background.Present {
		flags |= 1
	}
	if d.ShowToast {
		flags |= 2
	}
	if d.Hidden {
		flags |= 4
	}
	if err := buf.WriteInt32(flags); err != nil {
		return err
	}
	if d.Background.Present {
		return buf.WriteIdentifier(d.Background.Value)
	}
	return nil
}
