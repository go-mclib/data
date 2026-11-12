package items_test

import (
	"testing"

	"github.com/go-mclib/data/blocks"
	"github.com/go-mclib/data/items"
	"github.com/go-mclib/nbt"
	ns "github.com/go-mclib/protocol/javaproto/wire"
)

type wireMatcher struct {
	name     string
	exact    *string
	min, max *string
}

type wirePredicate struct {
	blocks     []string // names, or a single "#tag"
	properties []wireMatcher
	hasProps   bool
	nbt        nbt.Tag
	components bool
}

func text(value string) *string { return &value }

// encodeBlockPredicates writes a vanilla AdventureModePredicate.
func encodeBlockPredicates(t *testing.T, predicates ...wirePredicate) []byte {
	t.Helper()
	w := ns.NewWriter()
	w.WriteVarInt(ns.VarInt(len(predicates)))
	for _, p := range predicates {
		w.WriteBool(p.blocks != nil)
		if p.blocks != nil {
			set := ns.IDSet{}
			if len(p.blocks) == 1 && p.blocks[0][0] == '#' {
				set = ns.IDSet{IsTag: true, TagName: ns.Identifier(p.blocks[0][1:])}
			}
			for _, name := range p.blocks {
				if !set.IsTag {
					set.IDs = append(set.IDs, ns.VarInt(blocks.ID(name)))
				}
			}
			if err := set.Encode(w); err != nil {
				t.Fatal(err)
			}
		}
		w.WriteBool(ns.Boolean(p.hasProps))
		if p.hasProps {
			w.WriteVarInt(ns.VarInt(len(p.properties)))
			for _, m := range p.properties {
				w.WriteString(ns.String(m.name))
				w.WriteBool(m.exact != nil)
				if m.exact != nil {
					w.WriteString(ns.String(*m.exact))
					continue
				}
				for _, bound := range []*string{m.min, m.max} {
					w.WriteBool(bound != nil)
					if bound != nil {
						w.WriteString(ns.String(*bound))
					}
				}
			}
		}
		w.WriteBool(p.nbt != nil)
		if p.nbt != nil {
			if err := w.WriteNBT(p.nbt); err != nil {
				t.Fatal(err)
			}
		}
		if p.components {
			w.WriteVarInt(0) // no exact components
			w.WriteVarInt(1) // one partial: any value of a component type
			w.WriteBool(false)
			w.WriteVarInt(ns.VarInt(items.ComponentCustomName))
		} else {
			w.WriteVarInt(0)
			w.WriteVarInt(0)
		}
	}
	return w.Bytes()
}

func decodedPredicates(t *testing.T, predicates ...wirePredicate) []items.BlockPredicate {
	t.Helper()
	slot, err := decodeSlot(slotWithComponent(items.ID("minecraft:diamond_pickaxe"), items.ComponentCanBreak, encodeBlockPredicates(t, predicates...)))
	if err != nil {
		t.Fatalf("decode can_break slot: %v", err)
	}
	parsed, ok := items.BlockPredicates(slot, items.ComponentCanBreak)
	if !ok {
		t.Fatal("can_break not parsed")
	}
	if _, ok := items.BlockPredicates(slot, items.ComponentCanPlaceOn); ok {
		t.Fatal("can_place_on reported on a stack without it")
	}
	return parsed
}

func state(t *testing.T, name string, properties map[string]string) int32 {
	t.Helper()
	id := blocks.ID(name)
	if properties == nil {
		return blocks.DefaultStateID(id)
	}
	full := map[string]string{}
	_, defaults := blocks.StateProperties(blocks.DefaultStateID(id))
	for k, v := range defaults {
		full[k] = v
	}
	for k, v := range properties {
		full[k] = v
	}
	s := blocks.StateID(id, full)
	if s < 0 {
		t.Fatalf("invalid state %s %v", name, properties)
	}
	return s
}

func TestBlockPredicateMatchesBlocksAndTags(t *testing.T) {
	parsed := decodedPredicates(t,
		wirePredicate{blocks: []string{"minecraft:stone", "minecraft:granite"}},
		wirePredicate{blocks: []string{"#minecraft:logs"}},
	)
	if len(parsed) != 2 {
		t.Fatalf("predicates = %+v", parsed)
	}
	direct, tagged := parsed[0], parsed[1]
	if !direct.Matches(state(t, "minecraft:stone", nil)) || !direct.Matches(state(t, "minecraft:granite", nil)) {
		t.Fatal("direct block list must match its blocks")
	}
	if direct.Matches(state(t, "minecraft:dirt", nil)) {
		t.Fatal("direct block list matched dirt")
	}
	if !tagged.Matches(state(t, "minecraft:oak_log", nil)) || tagged.Matches(state(t, "minecraft:stone", nil)) {
		t.Fatal("tag predicate mismatch")
	}
}

func TestBlockPredicateMatchesStateProperties(t *testing.T) {
	parsed := decodedPredicates(t,
		wirePredicate{blocks: []string{"minecraft:wheat"}, hasProps: true, properties: []wireMatcher{{name: "age", min: text("3"), max: text("05")}}},
		wirePredicate{hasProps: true, properties: []wireMatcher{{name: "lit", exact: text("true")}}},
		wirePredicate{hasProps: true, properties: []wireMatcher{{name: "facing", min: text("south")}}},
		wirePredicate{hasProps: true, properties: []wireMatcher{{name: "powered", max: text("false")}}},
		wirePredicate{hasProps: true, properties: []wireMatcher{{name: "age", exact: text("nope")}}},
	)
	ageRange, lit, facing, unpowered, invalid := parsed[0], parsed[1], parsed[2], parsed[3], parsed[4]
	for age, want := range map[string]bool{"2": false, "3": true, "5": true, "6": false} {
		if got := ageRange.Matches(state(t, "minecraft:wheat", map[string]string{"age": age})); got != want {
			t.Errorf("wheat age %s matched=%v, want %v", age, got, want)
		}
	}
	if !lit.Matches(state(t, "minecraft:furnace", map[string]string{"lit": "true"})) ||
		lit.Matches(state(t, "minecraft:furnace", map[string]string{"lit": "false"})) {
		t.Error("exact boolean property mismatch")
	}
	if lit.Matches(state(t, "minecraft:stone", nil)) {
		t.Error("a missing property must not match")
	}
	// facing declares north, south, west, east
	for facingValue, want := range map[string]bool{"north": false, "south": true, "east": true} {
		if got := facing.Matches(state(t, "minecraft:furnace", map[string]string{"facing": facingValue})); got != want {
			t.Errorf("facing %s matched=%v, want %v", facingValue, got, want)
		}
	}
	// booleans order false < true
	if !unpowered.Matches(state(t, "minecraft:lever", map[string]string{"powered": "false"})) ||
		unpowered.Matches(state(t, "minecraft:lever", map[string]string{"powered": "true"})) {
		t.Error("boolean range must order false before true")
	}
	if invalid.Matches(state(t, "minecraft:wheat", nil)) {
		t.Error("an invalid property value must not match")
	}
}

func TestBlockPredicateBlockEntityRequirements(t *testing.T) {
	parsed := decodedPredicates(t,
		wirePredicate{blocks: []string{"minecraft:chest"}, nbt: nbt.Compound{"Lock": nbt.String("key")}},
		wirePredicate{blocks: []string{"minecraft:chest"}, components: true},
	)
	withNBT, withComponents := parsed[0], parsed[1]
	if withNBT.NBT == nil || withNBT.Matches(state(t, "minecraft:chest", nil)) {
		t.Fatalf("NBT predicate = %+v must require block-entity data", withNBT)
	}
	if !withComponents.HasComponents || !withComponents.Matches(state(t, "minecraft:chest", nil)) {
		t.Fatalf("component matchers are ignored by adventure checks: %+v", withComponents)
	}
}

func TestBlockPredicatesRejectTruncatedWire(t *testing.T) {
	data := encodeBlockPredicates(t, wirePredicate{blocks: []string{"minecraft:stone"}})
	if _, err := decodeSlot(slotWithComponent(items.ID("minecraft:diamond_pickaxe"), items.ComponentCanBreak, data[:len(data)-1])); err == nil {
		t.Fatal("truncated can_break decoded")
	}
	stack := ns.NewSlot(ns.VarInt(items.ID("minecraft:diamond_pickaxe")), 1)
	stack.Components.Add = []ns.RawSlotComponent{{ID: ns.VarInt(items.ComponentCanBreak), Data: data[:len(data)-1]}}
	if _, ok := items.BlockPredicates(stack, items.ComponentCanBreak); ok {
		t.Fatal("malformed can_break parsed")
	}
}
