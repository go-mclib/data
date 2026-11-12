package items

import (
	"fmt"
	"maps"
	"slices"
	"strconv"

	"github.com/go-mclib/data/blocks"
	"github.com/go-mclib/data/registries"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// maxPartialComponentPredicates is vanilla DataComponentPredicate.STREAM_CODEC's list limit.
const maxPartialComponentPredicates = 64

// BlockPredicate is one vanilla BlockPredicate of an adventure-mode
// can_break / can_place_on component (AdventureModePredicate).
type BlockPredicate struct {
	// Blocks lists block names, or a single "#tag" entry; nil matches any block.
	Blocks []string
	// Properties must all match the block state; nil matches any state.
	Properties []StatePropertyMatcher
	// NBT is the block-entity data the block must contain, or nil.
	NBT nbt.Tag
	// HasComponents reports data-component matchers on the block entity. Like
	// vanilla BlockPredicate.matches(BlockInWorld), adventure checks ignore them.
	HasComponents bool
}

// StatePropertyMatcher is a vanilla StatePropertiesPredicate entry: an exact
// value, or an inclusive range with optional bounds.
type StatePropertyMatcher struct {
	Name  string
	Exact bool
	Value string  // exact value
	Min   *string // range lower bound, nil when open
	Max   *string // range upper bound, nil when open
}

// BlockPredicates parses the adventure-mode block predicate component
// (ComponentCanBreak or ComponentCanPlaceOn) explicitly set on stack. It
// reports false when the stack does not carry the component or its data is
// malformed.
func BlockPredicates(stack ns.Slot, component int32) ([]BlockPredicate, bool) {
	for _, raw := range stack.Components.Add {
		if int32(raw.ID) != component {
			continue
		}
		predicates, err := readBlockPredicates(ns.NewReader(raw.Data))
		return predicates, err == nil
	}
	return nil, false
}

// Matches reports whether a block state satisfies the predicate the way
// vanilla adventure mode tests a block in the world. A predicate requiring
// block-entity NBT never matches: the state alone cannot satisfy it.
func (p BlockPredicate) Matches(state int32) bool {
	if p.NBT != nil {
		return false
	}
	blockID, properties := blocks.StateProperties(state)
	if blockID < 0 {
		return false
	}
	if p.Blocks != nil && !blockInHolderSet(blocks.Name(blockID), p.Blocks) {
		return false
	}
	for _, matcher := range p.Properties {
		if !matcher.matches(blockID, state, properties) {
			return false
		}
	}
	return true
}

func blockInHolderSet(name string, entries []string) bool {
	if tag, ok := singleTagRef(entries); ok {
		return slices.Contains(registries.TagData["minecraft:block"][tag], name)
	}
	return slices.Contains(entries, name)
}

func (m StatePropertyMatcher) matches(blockID, state int32, properties map[string]string) bool {
	if _, ok := properties[m.Name]; !ok {
		return false
	}
	if m.Exact {
		order, ok := comparePropertyValue(blockID, state, properties, m.Name, m.Value)
		return ok && order == 0
	}
	if m.Min != nil {
		if order, ok := comparePropertyValue(blockID, state, properties, m.Name, *m.Min); !ok || order < 0 {
			return false
		}
	}
	if m.Max != nil {
		if order, ok := comparePropertyValue(blockID, state, properties, m.Name, *m.Max); !ok || order > 0 {
			return false
		}
	}
	return true
}

// comparePropertyValue compares the state's value of a property with value,
// in vanilla Property order (false < true, integers numerically, enums by
// declaration). ok is false when value is not a valid value of the property.
func comparePropertyValue(blockID, state int32, properties map[string]string, name, value string) (int, bool) {
	actual := properties[name]
	if actual == "true" || actual == "false" {
		if value != "true" && value != "false" {
			return 0, false
		}
		return boolOrder(actual) - boolOrder(value), true
	}
	if _, err := strconv.Atoi(actual); err == nil {
		// vanilla IntegerProperty parses with Integer.parseInt
		parsed, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return 0, false
		}
		value = strconv.FormatInt(parsed, 10)
	}
	// Other values are listed in declaration order, so state IDs that differ
	// only in this property order the same way as its values.
	other := maps.Clone(properties)
	other[name] = value
	expected := blocks.StateID(blockID, other)
	if expected < 0 {
		return 0, false
	}
	switch {
	case state < expected:
		return -1, true
	case state > expected:
		return 1, true
	}
	return 0, true
}

func boolOrder(value string) int {
	if value == "true" {
		return 1
	}
	return 0
}

// decodeBlockPredicatesWire copies a vanilla AdventureModePredicate.
func decodeBlockPredicatesWire(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	return copyVarIntPrefixedList(buf, w, copyBlockPredicate)
}

func copyBlockPredicate(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := copyOptionalHolderSet(buf, w); err != nil {
		return err
	}
	present, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(present)
	if present {
		if err := copyVarIntPrefixedList(buf, w, copyPropertyMatcher); err != nil {
			return err
		}
	}
	if err := copyOptionalNBT(buf, w); err != nil {
		return err
	}
	// DataComponentMatchers: exact typed components, then partial predicates
	if err := copyVarIntPrefixedList(buf, w, copyTypedComponent); err != nil {
		return err
	}
	count, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	if count < 0 || count > maxPartialComponentPredicates {
		return fmt.Errorf("invalid component predicate count %d", count)
	}
	w.WriteVarInt(count)
	for range int(count) {
		if err := copyComponentPredicate(buf, w); err != nil {
			return err
		}
	}
	return nil
}

func copyPropertyMatcher(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	if err := w.CopyString(buf, maxStringLen); err != nil { // property name
		return err
	}
	exact, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(exact)
	if exact {
		return w.CopyString(buf, maxStringLen)
	}
	if err := copyOptionalString(buf, w); err != nil { // min
		return err
	}
	return copyOptionalString(buf, w) // max
}

// copyTypedComponent copies a TypedDataComponent: component type id + value.
func copyTypedComponent(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	id, err := buf.ReadVarInt()
	if err != nil {
		return err
	}
	data, err := decodeComponentWire(buf, id)
	if err != nil {
		return err
	}
	w.WriteVarInt(id)
	_, err = w.Write(data)
	return err
}

// copyComponentPredicate copies a DataComponentPredicate.Single: either a
// predicate type with its NBT-encoded value, or a bare component type
// (an "any value" check).
func copyComponentPredicate(buf *ns.PacketBuffer, w *ns.PacketBuffer) error {
	predicateType, err := buf.ReadBool()
	if err != nil {
		return err
	}
	w.WriteBool(predicateType)
	if err := w.CopyVarInt(buf); err != nil { // registry id
		return err
	}
	if predicateType {
		return copyNBT(buf, w)
	}
	return nil
}

func readBlockPredicates(buf *ns.PacketBuffer) ([]BlockPredicate, error) {
	count, err := buf.ReadVarInt()
	if err != nil {
		return nil, err
	}
	if count < 0 {
		return nil, fmt.Errorf("negative block predicate count %d", count)
	}
	var predicates []BlockPredicate
	for range int(count) {
		predicate, err := readBlockPredicate(buf)
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, predicate)
	}
	return predicates, nil
}

func readBlockPredicate(buf *ns.PacketBuffer) (BlockPredicate, error) {
	var predicate BlockPredicate
	var err error
	predicate.Blocks, err = readOptionalHolderSet(buf, "block", blocks.Name)
	if err != nil {
		return predicate, err
	}
	present, err := buf.ReadBool()
	if err != nil {
		return predicate, err
	}
	if present {
		count, err := buf.ReadVarInt()
		if err != nil {
			return predicate, err
		}
		if count < 0 {
			return predicate, fmt.Errorf("negative state property count %d", count)
		}
		predicate.Properties = []StatePropertyMatcher{}
		for range int(count) {
			matcher, err := readPropertyMatcher(buf)
			if err != nil {
				return predicate, err
			}
			predicate.Properties = append(predicate.Properties, matcher)
		}
	}
	if present, err = buf.ReadBool(); err != nil {
		return predicate, err
	}
	if present {
		if predicate.NBT, err = buf.ReadNBT(); err != nil {
			return predicate, err
		}
	}
	discard := ns.NewWriter()
	exact, err := buf.ReadVarInt()
	if err != nil {
		return predicate, err
	}
	if exact < 0 {
		return predicate, fmt.Errorf("negative component count %d", exact)
	}
	for range int(exact) {
		if err := copyTypedComponent(buf, discard); err != nil {
			return predicate, err
		}
	}
	partial, err := buf.ReadVarInt()
	if err != nil {
		return predicate, err
	}
	if partial < 0 || partial > maxPartialComponentPredicates {
		return predicate, fmt.Errorf("invalid component predicate count %d", partial)
	}
	for range int(partial) {
		if err := copyComponentPredicate(buf, discard); err != nil {
			return predicate, err
		}
	}
	predicate.HasComponents = exact > 0 || partial > 0
	return predicate, nil
}

func readPropertyMatcher(buf *ns.PacketBuffer) (StatePropertyMatcher, error) {
	var matcher StatePropertyMatcher
	name, err := buf.ReadString(maxStringLen)
	if err != nil {
		return matcher, err
	}
	matcher.Name = string(name)
	exact, err := buf.ReadBool()
	if err != nil {
		return matcher, err
	}
	if exact {
		value, err := buf.ReadString(maxStringLen)
		matcher.Exact, matcher.Value = true, string(value)
		return matcher, err
	}
	if matcher.Min, err = readOptionalPredicateString(buf); err != nil {
		return matcher, err
	}
	matcher.Max, err = readOptionalPredicateString(buf)
	return matcher, err
}

func readOptionalPredicateString(buf *ns.PacketBuffer) (*string, error) {
	present, err := buf.ReadBool()
	if err != nil || !present {
		return nil, err
	}
	value, err := buf.ReadString(maxStringLen)
	if err != nil {
		return nil, err
	}
	text := string(value)
	return &text, nil
}
