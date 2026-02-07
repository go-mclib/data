package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type entityDimensions struct {
	width     float64
	height    float64
	eyeHeight float64 // -1 means use default (height * 0.85)
}

func generateEntityHitboxes(entityTypeJavaPath, outPath string) {
	data, err := os.ReadFile(entityTypeJavaPath)
	if err != nil {
		fmt.Printf("skipping entity hitboxes generation: %v\n", err)
		return
	}
	src := string(data)

	// extract all register(...) blocks with their entity name, sized(), and eyeHeight()
	// pattern: register(\n   "name", ... .sized(W, H) ... optionally .eyeHeight(E) ...
	//
	// we scan for each register("name", ...) call and extract sized/eyeHeight from the
	// builder chain that follows, up to the closing ");".

	// find each register call: register(\n  "name",
	registerRe := regexp.MustCompile(`register\(\s*"([^"]+)"`)
	sizedRe := regexp.MustCompile(`\.sized\(([^,]+),\s*([^)]+)\)`)
	eyeHeightRe := regexp.MustCompile(`\.eyeHeight\(([^)]+)\)`)

	dims := make(map[string]entityDimensions)

	matches := registerRe.FindAllStringIndex(src, -1)
	for _, loc := range matches {
		// extract entity name
		nameMatch := registerRe.FindStringSubmatch(src[loc[0]:loc[1]])
		if nameMatch == nil {
			continue
		}
		name := nameMatch[1]

		// find the extent of this register() call by finding the matching ");"
		// start from after the opening "register("
		block := findRegisterBlock(src, loc[0])
		if block == "" {
			continue
		}

		// extract sized(width, height)
		sizedMatch := sizedRe.FindStringSubmatch(block)
		if sizedMatch == nil {
			continue
		}
		width := parseJavaFloat(sizedMatch[1])
		height := parseJavaFloat(sizedMatch[2])

		// extract eyeHeight if present
		eyeHeight := -1.0
		eyeMatch := eyeHeightRe.FindStringSubmatch(block)
		if eyeMatch != nil {
			eyeHeight = parseJavaFloat(eyeMatch[1])
		}

		dims["minecraft:"+name] = entityDimensions{
			width:     width,
			height:    height,
			eyeHeight: eyeHeight,
		}
	}

	fmt.Printf("entity hitboxes: extracted %d entities from EntityType.java\n", len(dims))

	// sort entity names for deterministic output
	names := make([]string, 0, len(dims))
	for name := range dims {
		names = append(names, name)
	}
	sort.Strings(names)

	// generate Go code
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("entities"))

	sb.WriteString(fmt.Sprintf("// Dimensions contains the standing hitbox dimensions for %d entity types.\n", len(dims)))
	sb.WriteString("//\n// Eye height of -1 means the entity uses the default formula: height * 0.85.\n")
	sb.WriteString("var dimensions = [...]struct {\n\tIdentifier string\n\tWidth      float32\n\tHeight     float32\n\tEyeHeight  float32\n}{\n")
	for _, name := range names {
		d := dims[name]
		eyeStr := formatFloat32(d.eyeHeight)
		sb.WriteString(fmt.Sprintf("\t{%q, %s, %s, %s},\n",
			name,
			formatFloat32(d.width),
			formatFloat32(d.height),
			eyeStr,
		))
	}
	sb.WriteString("}\n\n")

	// lookup map: name -> index
	sb.WriteString("var dimensionsByName = map[string]int{\n")
	for i, name := range names {
		sb.WriteString(fmt.Sprintf("\t%q: %d,\n", name, i))
	}
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
}

// findRegisterBlock extracts the full register(...) call starting at pos.
func findRegisterBlock(src string, pos int) string {
	depth := 0
	start := pos
	for i := pos; i < len(src); i++ {
		switch src[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				return src[start : i+1]
			}
		}
	}
	return ""
}

func parseJavaFloat(s string) float64 {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, "F")
	s = strings.TrimSuffix(s, "f")
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse Java float %q: %v", s, err))
	}
	return v
}

func formatFloat32(f float64) string {
	s := strconv.FormatFloat(f, 'f', -1, 32)
	if !strings.Contains(s, ".") {
		s += ".0"
	}
	return s
}
