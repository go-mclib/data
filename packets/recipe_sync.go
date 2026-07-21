package packets

import (
	"fmt"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// RecipePropertySet contains the item registry ids accepted by a recipe input role.
type RecipePropertySet struct {
	Key   ns.Identifier
	Items []ns.VarInt
}

// RecipeIngredient is either a named item tag or an explicit item-holder set.
type RecipeIngredient struct {
	Tag   ns.Identifier
	Items []ns.VarInt
}

func (i *RecipeIngredient) read(buf *ns.PacketBuffer) error {
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	if count == 0 {
		i.Tag, err = buf.ReadIdentifier()
		return err
	}
	count--
	if count < 0 || count > 65536 {
		return fmt.Errorf("invalid recipe ingredient size %d", count)
	}
	i.Items = make([]ns.VarInt, count)
	for n := range i.Items {
		if i.Items[n], err = buf.ReadVarInt(); err != nil {
			return err
		}
	}
	return nil
}

func (i RecipeIngredient) write(buf *ns.PacketBuffer) error {
	if i.Tag != "" {
		if err := buf.WriteVarInt(0); err != nil {
			return err
		}
		return buf.WriteIdentifier(i.Tag)
	}
	if err := buf.WriteVarInt(ns.VarInt(len(i.Items) + 1)); err != nil {
		return err
	}
	for _, item := range i.Items {
		if err := buf.WriteVarInt(item); err != nil {
			return err
		}
	}
	return nil
}

// StonecutterRecipeEntry is one synchronized stonecutter input and output display.
type StonecutterRecipeEntry struct {
	Ingredient  RecipeIngredient
	ResultItem  ns.VarInt
	ResultCount ns.VarInt
}

func (e *StonecutterRecipeEntry) read(buf *ns.PacketBuffer) error {
	if err := e.Ingredient.read(buf); err != nil {
		return err
	}
	displayType, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	if displayType != 5 {
		return fmt.Errorf("unsupported stonecutter slot display type %d", displayType)
	}
	if e.ResultItem, err = buf.ReadVarInt(); err != nil {
		return err
	}
	if e.ResultCount, err = buf.ReadVarInt(); err != nil {
		return err
	}
	added, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	removed, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	if added != 0 || removed != 0 {
		return fmt.Errorf("unsupported stonecutter result component patch %d/%d", added, removed)
	}
	return nil
}

func (e StonecutterRecipeEntry) write(buf *ns.PacketBuffer) error {
	if err := e.Ingredient.write(buf); err != nil {
		return err
	}
	if err := buf.WriteVarInt(5); err != nil { // minecraft:item_stack slot display
		return err
	}
	if err := buf.WriteVarInt(e.ResultItem); err != nil {
		return err
	}
	if err := buf.WriteVarInt(e.ResultCount); err != nil {
		return err
	}
	if err := buf.WriteVarInt(0); err != nil {
		return err
	}
	return buf.WriteVarInt(0)
}
