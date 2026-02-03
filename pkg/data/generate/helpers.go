package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

func loadJSON[T any](path string) T {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("failed to read %s: %v", path, err))
	}
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		panic(fmt.Sprintf("failed to parse %s: %v", path, err))
	}
	return result
}

func writeFile(path, content string) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(fmt.Sprintf("failed to create dir %s: %v", dir, err))
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		panic(fmt.Sprintf("failed to write %s: %v", path, err))
	}
	fmt.Printf("wrote %s\n", path)
}

func toGoName(id string) string {
	// minecraft:acacia_button -> AcaciaButton
	id = strings.TrimPrefix(id, "minecraft:")
	parts := strings.Split(id, "_")
	var result strings.Builder
	for _, part := range parts {
		if len(part) > 0 {
			// handle special cases like worldgen/biome_source
			subparts := strings.Split(part, "/")
			for _, sp := range subparts {
				if len(sp) > 0 {
					result.WriteString(strings.ToUpper(sp[:1]))
					result.WriteString(sp[1:])
				}
			}
		}
	}
	return result.String()
}

func toGoVarName(id string) string {
	name := toGoName(id)
	// handle names starting with numbers
	if len(name) > 0 && unicode.IsDigit(rune(name[0])) {
		name = "N" + name
	}
	return name
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
