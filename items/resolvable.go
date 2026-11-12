package items

import (
	"encoding/json"
	"math"
	rand "math/rand/v2"
	"slices"
	"strings"
	"sync"

	"github.com/go-mclib/data/blocks"
	"github.com/go-mclib/data/registries"
)

var contextProviders = sync.OnceValue(func() map[string]map[string]any {
	result := make(map[string]map[string]any, len(registries.ReloadableRegistryData))
	for registry, entries := range registries.ReloadableRegistryData {
		values := make(map[string]any, len(entries))
		for name, raw := range entries {
			var value any
			if json.Unmarshal(raw, &value) == nil {
				values[name] = value
			}
		}
		result[registry] = values
	}
	return result
})

// Resolve evaluates the number, or its contextual provider against a block
// state (e.g. the furnace burning a fuel). Missing or invalid references use
// fallback, as vanilla ResolvableInt does.
func (value ResolvableInt) Resolve(blockState int32, fallback int32) int32 {
	if value.Reference == "" {
		return value.Value
	}
	resolved, ok := contextNumber(value.Reference, "minecraft:context_int_provider", blockState, 0)
	if !ok || math.IsNaN(resolved) || math.IsInf(resolved, 0) {
		return fallback
	}
	return int32(max(math.MinInt32, min(math.MaxInt32, resolved)))
}

// Resolve is the floating-point counterpart of ResolvableInt.Resolve.
func (value ResolvableFloat) Resolve(blockState int32, fallback float32) float32 {
	if value.Reference == "" {
		return value.Value
	}
	resolved, ok := contextNumber(value.Reference, "minecraft:context_float_provider", blockState, 0)
	if !ok || math.IsNaN(resolved) || math.IsInf(resolved, 0) {
		return fallback
	}
	return float32(resolved)
}

func contextResource(name string) string {
	if !strings.Contains(name, ":") {
		return "minecraft:" + name
	}
	return name
}

func contextNumber(value any, registry string, state int32, depth int) (float64, bool) {
	if depth > 64 {
		return 0, false
	}
	switch value := value.(type) {
	case float64:
		return value, true
	case string:
		return contextNumber(contextProviders()[registry][contextResource(value)], registry, state, depth+1)
	case map[string]any:
		kind, _ := value["type"].(string)
		switch strings.TrimPrefix(kind, "minecraft:") {
		case "conditional":
			matches, ok := contextBlockPredicate(value["condition"], state, depth+1)
			if !ok {
				return 0, false
			}
			branch := "on_false"
			if matches {
				branch = "on_true"
			}
			return contextNumber(value[branch], registry, state, depth+1)
		case "number_dispatcher":
			cases, _ := value["cases"].([]any)
			for _, entry := range cases {
				entry, ok := entry.(map[string]any)
				if !ok {
					return 0, false
				}
				matches, ok := contextBlockPredicate(entry["condition"], state, depth+1)
				if !ok {
					return 0, false
				}
				if matches {
					return contextNumber(entry["value"], registry, state, depth+1)
				}
			}
			return contextNumber(value["default"], registry, state, depth+1)
		case "weighted_list":
			entries, _ := value["distribution"].([]any)
			total := 0.0
			for _, entry := range entries {
				entry, ok := entry.(map[string]any)
				if !ok {
					return 0, false
				}
				weight, ok := entry["weight"].(float64)
				if !ok || weight < 0 {
					return 0, false
				}
				total += weight
			}
			if total <= 0 {
				return 0, false
			}
			pick := rand.Float64() * total
			for _, entry := range entries {
				entry := entry.(map[string]any)
				pick -= entry["weight"].(float64)
				if pick < 0 {
					return contextNumber(entry["data"], registry, state, depth+1)
				}
			}
		case "div":
			left, lok := contextNumber(value["left"], registry, state, depth+1)
			right, rok := contextNumber(value["right"], registry, state, depth+1)
			if !lok || !rok || right == 0 {
				return 0, false
			}
			result := left / right
			if registry == "minecraft:context_int_provider" {
				result = math.Trunc(result)
			}
			return result, true
		}
	}
	return 0, false
}

func contextBlockPredicate(value any, state int32, depth int) (bool, bool) {
	if depth > 64 {
		return false, false
	}
	if name, ok := value.(string); ok {
		return contextBlockPredicate(contextProviders()["minecraft:predicate"][contextResource(name)], state, depth+1)
	}
	predicate, ok := value.(map[string]any)
	if !ok {
		return false, false
	}
	kind, _ := predicate["type"].(string)
	if strings.TrimPrefix(kind, "minecraft:") != "match_block" {
		return false, false
	}
	id, properties := blocks.StateProperties(state)
	if id < 0 {
		return false, true
	}
	name := blocks.Name(int32(id))
	matches := func(candidate any) bool {
		text, ok := candidate.(string)
		return ok && contextResource(text) == name
	}
	if wanted, present := predicate["blocks"]; present {
		list, _ := wanted.([]any)
		if !matches(wanted) && !slices.ContainsFunc(list, matches) {
			return false, true
		}
	}
	if wanted, ok := predicate["state"].(map[string]any); ok {
		for key, expected := range wanted {
			if properties[key] != expected {
				return false, true
			}
		}
	}
	return true, true
}
