package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
)

// versionJSON is the server jar's version.json, extracted by
// vanilla_server_reports/generate.sh.
type versionJSON struct {
	Name            string `json:"name"`
	ProtocolVersion int    `json:"protocol_version"`
	DataVersion     int    `json:"world_version"`
}

// version is the Minecraft version being generated, loaded first in main.
var version versionJSON

// lowCountWarnings accumulates scrapers that produced implausibly few entries.
var lowCountWarnings []string

// checkMinCount records a sanity-check failure when a scraper produced fewer
// entries than expected. This catches the silent failure mode where Mojang
// refactors the decompiled source layout and a Java-scraping parser quietly
// matches nothing (e.g. "extracted 0 entities") but generation still completes.
// Recorded failures are reported together by reportSanityChecks.
func checkMinCount(label string, got, min int) {
	if got < min {
		lowCountWarnings = append(lowCountWarnings, fmt.Sprintf("%s: scraped %d entries (expected >= %d)", label, got, min))
	}
}

// reportSanityChecks prints any accumulated low-count failures and exits
// non-zero, so a broken scraper fails generation loudly instead of silently
// shipping near-empty data.
func reportSanityChecks() {
	if len(lowCountWarnings) == 0 {
		return
	}
	fmt.Fprintln(os.Stderr, "\n!!! SANITY CHECK FAILED: a parser likely broke (decompiled source layout may have changed):")
	for _, w := range lowCountWarnings {
		fmt.Fprintf(os.Stderr, "  - %s\n", w)
	}
	fmt.Fprintln(os.Stderr, "")
	os.Exit(1)
}

// generatedFileHeader returns the standard header for generated Go files.
func generatedFileHeader(pkg string) string {
	return fmt.Sprintf(`// Code generated for Minecraft %s (Protocol %d); DO NOT EDIT.

package %s

`, version.Name, version.ProtocolVersion, pkg)
}

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
	var result strings.Builder
	for part := range strings.SplitSeq(id, "_") {
		// handle special cases like worldgen/biome_source
		for sp := range strings.SplitSeq(part, "/") {
			if len(sp) > 0 {
				result.WriteString(strings.ToUpper(sp[:1]))
				result.WriteString(sp[1:])
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

// loadItems loads item component data from the per-item report files
// (items/<item_name>.json).
func loadItems(baseDir string) map[string]ItemJSON {
	itemsDir := filepath.Join(baseDir, "items")
	entries, err := os.ReadDir(itemsDir)
	if err != nil {
		panic(fmt.Sprintf("failed to read items directory %s: %v", itemsDir, err))
	}
	items := make(map[string]ItemJSON, len(entries))
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		name := "minecraft:" + strings.TrimSuffix(entry.Name(), ".json")
		items[name] = loadJSON[ItemJSON](filepath.Join(itemsDir, entry.Name()))
	}
	return items
}

func sortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	return slices.Sorted(maps.Keys(m))
}

// datapackEntry is one JSON file of a vanilla datapack directory.
type datapackEntry struct {
	ID   string // "minecraft:" + path relative to the walked dir, without .json
	Path string
}

// datapackEntries walks dir recursively and returns its JSON files as
// "minecraft:<relative path>" identifiers, sorted by identifier like vanilla's
// resource listing (which also decides synchronized registry protocol IDs).
func datapackEntries(dir string) []datapackEntry {
	var entries []datapackEntry
	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
		if err != nil || entry.IsDir() || filepath.Ext(path) != ".json" {
			return err
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		id := "minecraft:" + strings.TrimSuffix(filepath.ToSlash(rel), ".json")
		entries = append(entries, datapackEntry{ID: id, Path: path})
		return nil
	})
	if err != nil {
		panic(fmt.Sprintf("failed to read datapack dir %s: %v", dir, err))
	}
	slices.SortFunc(entries, func(a, b datapackEntry) int { return cmp.Compare(a.ID, b.ID) })
	return entries
}

// generateVersion creates version_gen.go with the version constants.
func generateVersion(outPath string) {
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("data"))
	fmt.Fprintf(&sb, `// ProtocolVersion is the Minecraft protocol version this package was generated for.
const ProtocolVersion = %d

// MinecraftVersion is the Minecraft game version this package was generated for.
const MinecraftVersion = %q

// DataVersion is the world data version saved data is written with.
const DataVersion = %d
`, version.ProtocolVersion, version.Name, version.DataVersion)

	writeFile(outPath, sb.String())
}
