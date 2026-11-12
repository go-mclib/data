package advancements

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"

	ns "github.com/go-mclib/protocol/javaproto/wire"
)

// Definition is a vanilla datapack advancement definition.
type Definition struct {
	ID                  string
	Parent              string               `json:"parent"`
	Display             *Display             `json:"display"`
	Criteria            map[string]Criterion `json:"criteria"`
	Requirements        [][]string           `json:"requirements"`
	Rewards             Rewards              `json:"rewards"`
	SendsTelemetryEvent bool                 `json:"sends_telemetry_event"`
}

// Criterion is an advancement trigger and its predicates.
type Criterion struct {
	Trigger    string          `json:"trigger"`
	Conditions json.RawMessage `json:"conditions"`
}

// Display is the client-visible advancement presentation.
type Display struct {
	Icon           Icon             `json:"icon"`
	Title          ns.TextComponent `json:"title"`
	Description    ns.TextComponent `json:"description"`
	Background     string           `json:"background"`
	Frame          string           `json:"frame"`
	ShowToast      *bool            `json:"show_toast"`
	AnnounceToChat *bool            `json:"announce_to_chat"`
	Hidden         bool             `json:"hidden"`
}

// Icon is the item stack template shown in the advancement tree.
type Icon struct {
	ID         string                     `json:"id"`
	Count      int32                      `json:"count"`
	Components map[string]json.RawMessage `json:"components"`
}

// Rewards are granted when an advancement is completed.
type Rewards struct {
	Experience int32    `json:"experience"`
	Loot       []string `json:"loot"`
	Recipes    []string `json:"recipes"`
	Function   string   `json:"function"`
}

// All parses and returns every generated vanilla advancement in ID order.
func All() ([]Definition, error) {
	ids := slices.Sorted(maps.Keys(rawDefinitions))
	definitions := make([]Definition, 0, len(ids))
	for _, id := range ids {
		var definition Definition
		if err := json.Unmarshal(rawDefinitions[id], &definition); err != nil {
			return nil, fmt.Errorf("decode advancement %s: %w", id, err)
		}
		definition.ID = id
		if definition.Display != nil && definition.Display.Icon.Count < 1 {
			definition.Display.Icon.Count = 1
		}
		definitions = append(definitions, definition)
	}
	return definitions, nil
}

// CriterionNames returns criterion names in stable order.
func (d Definition) CriterionNames() []string {
	return slices.Sorted(maps.Keys(d.Criteria))
}

// NormalizedRequirements returns the explicit requirements or vanilla's all-of default.
func (d Definition) NormalizedRequirements() [][]string {
	if d.Requirements != nil {
		return d.Requirements
	}
	names := d.CriterionNames()
	result := make([][]string, len(names))
	for i, name := range names {
		result[i] = []string{name}
	}
	return result
}

// Complete reports whether the obtained criteria satisfy every requirement group.
func (d Definition) Complete(obtained map[string]int64) bool {
	for _, group := range d.NormalizedRequirements() {
		if !slices.ContainsFunc(group, func(c string) bool { _, ok := obtained[c]; return ok }) {
			return false
		}
	}
	return true
}
