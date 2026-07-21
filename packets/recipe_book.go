package packets

import (
	"fmt"

	"github.com/go-mclib/data/items"
	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

const maxRecipeDisplays = 65536

type RecipeBookEntry struct {
	DisplayID       ns.VarInt
	Display         RecipeDisplay
	Group           int32
	Category        ns.VarInt
	Requirements    []RecipeIngredient
	HasRequirements bool
	Flags           ns.Int8
}

type RecipeDisplay struct {
	Type        ns.VarInt
	Width       ns.VarInt
	Height      ns.VarInt
	Ingredients []RecipeSlotDisplay
	Template    RecipeSlotDisplay
	Base        RecipeSlotDisplay
	Addition    RecipeSlotDisplay
	Fuel        RecipeSlotDisplay
	Result      RecipeSlotDisplay
	Station     RecipeSlotDisplay
	Duration    ns.VarInt
	Experience  ns.Float32
}

type RecipeSlotDisplay struct {
	Type        ns.VarInt
	ItemID      ns.VarInt
	Count       ns.VarInt
	Tag         ns.Identifier
	ComponentID ns.VarInt
	PatternID   ns.VarInt
	Item        ns.Slot
	Children    []RecipeSlotDisplay
}

func (e *RecipeBookEntry) read(buf *ns.PacketBuffer) error {
	var err error
	if e.DisplayID, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if err := e.Display.read(buf); err != nil {
		return err
	}
	group, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	e.Group = int32(group) - 1
	if e.Category, err = buf.ReadVarInt(); err != nil {
		return err
	}
	hasRequirements, err := buf.ReadBool()
	if err != nil {
		return err
	}
	e.HasRequirements = bool(hasRequirements)
	if e.HasRequirements {
		count, err := readRecipeCount(buf, "requirements")
		if err != nil {
			return err
		}
		e.Requirements = make([]RecipeIngredient, count)
		for i := range e.Requirements {
			if err := e.Requirements[i].read(buf); err != nil {
				return err
			}
		}
	}
	e.Flags, err = buf.ReadInt8()
	return err
}

func (e RecipeBookEntry) write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(e.DisplayID); err != nil {
		return err
	}
	if err := e.Display.write(buf); err != nil {
		return err
	}
	if err := buf.WriteVarInt(ns.VarInt(e.Group + 1)); err != nil {
		return err
	}
	if err := buf.WriteVarInt(e.Category); err != nil {
		return err
	}
	hasRequirements := e.HasRequirements || e.Requirements != nil
	if err := buf.WriteBool(ns.Boolean(hasRequirements)); err != nil {
		return err
	}
	if hasRequirements {
		if err := buf.WriteVarInt(ns.VarInt(len(e.Requirements))); err != nil {
			return err
		}
		for _, requirement := range e.Requirements {
			if err := requirement.write(buf); err != nil {
				return err
			}
		}
	}
	return buf.WriteInt8(e.Flags)
}

func (d *RecipeDisplay) read(buf *ns.PacketBuffer) error {
	var err error
	if d.Type, err = buf.ReadVarInt(); err != nil {
		return err
	}
	switch d.Type {
	case 0:
		d.Ingredients, err = readSlotDisplayList(buf)
		if err == nil {
			err = d.Result.read(buf, 0)
		}
		if err == nil {
			err = d.Station.read(buf, 0)
		}
	case 1:
		if d.Width, err = buf.ReadVarInt(); err != nil {
			return err
		}
		if d.Height, err = buf.ReadVarInt(); err != nil {
			return err
		}
		d.Ingredients, err = readSlotDisplayList(buf)
		if err == nil {
			err = d.Result.read(buf, 0)
		}
		if err == nil {
			err = d.Station.read(buf, 0)
		}
	case 2:
		if err = d.readOneIngredient(buf); err == nil {
			err = d.Fuel.read(buf, 0)
		}
		if err == nil {
			err = d.Result.read(buf, 0)
		}
		if err == nil {
			err = d.Station.read(buf, 0)
		}
		if err == nil {
			d.Duration, err = buf.ReadVarInt()
		}
		if err == nil {
			d.Experience, err = buf.ReadFloat32()
		}
	case 3:
		if err = d.readOneIngredient(buf); err == nil {
			err = d.Result.read(buf, 0)
		}
		if err == nil {
			err = d.Station.read(buf, 0)
		}
	case 4:
		if err = d.Template.read(buf, 0); err == nil {
			err = d.Base.read(buf, 0)
		}
		if err == nil {
			err = d.Addition.read(buf, 0)
		}
		if err == nil {
			err = d.Result.read(buf, 0)
		}
		if err == nil {
			err = d.Station.read(buf, 0)
		}
	default:
		return fmt.Errorf("unsupported recipe display type %d", d.Type)
	}
	return err
}

func (d *RecipeDisplay) readOneIngredient(buf *ns.PacketBuffer) error {
	var ingredient RecipeSlotDisplay
	if err := ingredient.read(buf, 0); err != nil {
		return err
	}
	d.Ingredients = []RecipeSlotDisplay{ingredient}
	return nil
}

func (d RecipeDisplay) write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(d.Type); err != nil {
		return err
	}
	switch d.Type {
	case 0:
		if err := writeSlotDisplayList(buf, d.Ingredients); err != nil {
			return err
		}
		if err := d.Result.write(buf); err != nil {
			return err
		}
		return d.Station.write(buf)
	case 1:
		if err := buf.WriteVarInt(d.Width); err != nil {
			return err
		}
		if err := buf.WriteVarInt(d.Height); err != nil {
			return err
		}
		if err := writeSlotDisplayList(buf, d.Ingredients); err != nil {
			return err
		}
		if err := d.Result.write(buf); err != nil {
			return err
		}
		return d.Station.write(buf)
	case 2:
		if len(d.Ingredients) != 1 {
			return fmt.Errorf("furnace display has %d ingredients", len(d.Ingredients))
		}
		if err := d.Ingredients[0].write(buf); err != nil {
			return err
		}
		if err := d.Fuel.write(buf); err != nil {
			return err
		}
		if err := d.Result.write(buf); err != nil {
			return err
		}
		if err := d.Station.write(buf); err != nil {
			return err
		}
		if err := buf.WriteVarInt(d.Duration); err != nil {
			return err
		}
		return buf.WriteFloat32(d.Experience)
	case 3:
		if len(d.Ingredients) != 1 {
			return fmt.Errorf("stonecutter display has %d ingredients", len(d.Ingredients))
		}
		if err := d.Ingredients[0].write(buf); err != nil {
			return err
		}
		if err := d.Result.write(buf); err != nil {
			return err
		}
		return d.Station.write(buf)
	case 4:
		for _, display := range []RecipeSlotDisplay{d.Template, d.Base, d.Addition, d.Result, d.Station} {
			if err := display.write(buf); err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("unsupported recipe display type %d", d.Type)
	}
}

func (d *RecipeSlotDisplay) read(buf *ns.PacketBuffer, depth int) error {
	if depth > 32 {
		return fmt.Errorf("recipe slot display nesting exceeds 32")
	}
	var err error
	if d.Type, err = buf.ReadVarInt(); err != nil {
		return err
	}
	switch d.Type {
	case 0, 1:
		return nil
	case 2:
		child, err := readSlotDisplay(buf, depth+1)
		if err != nil {
			return err
		}
		d.Children = []RecipeSlotDisplay{child}
		return nil
	case 3:
		child, err := readSlotDisplay(buf, depth+1)
		if err != nil {
			return err
		}
		d.Children = []RecipeSlotDisplay{child}
		d.ComponentID, err = buf.ReadVarInt()
		return err
	case 4:
		d.ItemID, err = buf.ReadVarInt()
		return err
	case 5:
		d.Item, err = items.ReadItemStackTemplate(buf)
		d.ItemID, d.Count = d.Item.ItemID, d.Item.Count
		return err
	case 6:
		d.Tag, err = buf.ReadIdentifier()
		return err
	case 7, 9:
		first, err := readSlotDisplay(buf, depth+1)
		if err != nil {
			return err
		}
		second, err := readSlotDisplay(buf, depth+1)
		if err != nil {
			return err
		}
		d.Children = []RecipeSlotDisplay{first, second}
		return nil
	case 8:
		base, err := readSlotDisplay(buf, depth+1)
		if err != nil {
			return err
		}
		material, err := readSlotDisplay(buf, depth+1)
		if err != nil {
			return err
		}
		d.Children = []RecipeSlotDisplay{base, material}
		marker, err := buf.ReadVarInt()
		if err != nil {
			return err
		}
		if marker == 0 {
			return fmt.Errorf("direct trim pattern display is unsupported")
		}
		d.PatternID = marker - 1
		return nil
	case 10:
		d.Children, err = readSlotDisplayListDepth(buf, depth+1)
		return err
	default:
		return fmt.Errorf("unsupported recipe slot display type %d", d.Type)
	}
}

func (d RecipeSlotDisplay) write(buf *ns.PacketBuffer) error {
	if err := buf.WriteVarInt(d.Type); err != nil {
		return err
	}
	switch d.Type {
	case 0, 1:
		return nil
	case 2:
		if len(d.Children) != 1 {
			return fmt.Errorf("potion display has %d children", len(d.Children))
		}
		return d.Children[0].write(buf)
	case 3:
		if len(d.Children) != 1 {
			return fmt.Errorf("component-filter display has %d children", len(d.Children))
		}
		if err := d.Children[0].write(buf); err != nil {
			return err
		}
		return buf.WriteVarInt(d.ComponentID)
	case 4:
		return buf.WriteVarInt(d.ItemID)
	case 5:
		item := d.Item
		if item.IsEmpty() {
			item = ns.NewSlot(d.ItemID, d.Count)
		}
		return items.WriteItemStackTemplate(buf, item)
	case 6:
		return buf.WriteIdentifier(d.Tag)
	case 7, 9:
		if len(d.Children) != 2 {
			return fmt.Errorf("binary slot display has %d children", len(d.Children))
		}
		if err := d.Children[0].write(buf); err != nil {
			return err
		}
		return d.Children[1].write(buf)
	case 8:
		if len(d.Children) != 2 {
			return fmt.Errorf("trim display has %d children", len(d.Children))
		}
		if err := d.Children[0].write(buf); err != nil {
			return err
		}
		if err := d.Children[1].write(buf); err != nil {
			return err
		}
		return buf.WriteVarInt(d.PatternID + 1)
	case 10:
		return writeSlotDisplayList(buf, d.Children)
	default:
		return fmt.Errorf("unsupported recipe slot display type %d", d.Type)
	}
}

func readSlotDisplay(buf *ns.PacketBuffer, depth int) (RecipeSlotDisplay, error) {
	var display RecipeSlotDisplay
	err := display.read(buf, depth)
	return display, err
}

func readSlotDisplayList(buf *ns.PacketBuffer) ([]RecipeSlotDisplay, error) {
	return readSlotDisplayListDepth(buf, 0)
}

func readSlotDisplayListDepth(buf *ns.PacketBuffer, depth int) ([]RecipeSlotDisplay, error) {
	count, err := readRecipeCount(buf, "slot displays")
	if err != nil {
		return nil, err
	}
	result := make([]RecipeSlotDisplay, count)
	for i := range result {
		if err := result[i].read(buf, depth); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func writeSlotDisplayList(buf *ns.PacketBuffer, displays []RecipeSlotDisplay) error {
	if err := buf.WriteVarInt(ns.VarInt(len(displays))); err != nil {
		return err
	}
	for _, display := range displays {
		if err := display.write(buf); err != nil {
			return err
		}
	}
	return nil
}

func readRecipeCount(buf *ns.PacketBuffer, label string) (int, error) {
	count, err := buf.ReadVarInt()
	if err != nil {
		return 0, err
	}
	if count < 0 || count > maxRecipeDisplays {
		return 0, fmt.Errorf("invalid %s count %d", label, count)
	}
	return int(count), nil
}
