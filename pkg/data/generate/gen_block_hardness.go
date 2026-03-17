package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type blockProps struct {
	hardness            float32
	requiresCorrectTool bool
}

// generateBlockHardness parses Blocks.java to extract hardness and requiresCorrectToolForDrops.
func generateBlockHardness(decompiledDir, outPath string) {
	blocksJava := filepath.Join(decompiledDir, "net", "minecraft", "world", "level", "block", "Blocks.java")
	data, err := os.ReadFile(blocksJava)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: cannot read Blocks.java for hardness data: %v\n", err)
		return
	}

	props := parseBlockProps(string(data))
	if len(props) == 0 {
		fmt.Fprintf(os.Stderr, "warning: no block hardness data found in Blocks.java\n")
		return
	}

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("blocks"))
	sb.WriteString("// BlockHardness returns the destroy time (hardness) for a block by registry name.\n")
	sb.WriteString("// Returns -1 for unbreakable blocks, 0 for instant break, -2 if not found.\n")
	sb.WriteString("func BlockHardness(name string) float32 {\n")
	sb.WriteString("\tv, ok := blockHardness[name]\n")
	sb.WriteString("\tif !ok { return -2 }\n")
	sb.WriteString("\treturn v\n")
	sb.WriteString("}\n\n")

	sb.WriteString("// BlockRequiresCorrectTool returns whether the block needs the correct tool to mine at full speed.\n")
	sb.WriteString("func BlockRequiresCorrectTool(name string) bool {\n")
	sb.WriteString("\treturn blockRequiresCorrectTool[name]\n")
	sb.WriteString("}\n\n")

	names := make([]string, 0, len(props))
	for name := range props {
		names = append(names, name)
	}
	sort.Strings(names)

	sb.WriteString("var blockHardness = map[string]float32{\n")
	for _, name := range names {
		sb.WriteString(fmt.Sprintf("\t%q: %s,\n", "minecraft:"+name, formatFloat32(float64(props[name].hardness))))
	}
	sb.WriteString("}\n\n")

	// only emit blocks that require correct tool (saves space)
	sb.WriteString("var blockRequiresCorrectTool = map[string]bool{\n")
	for _, name := range names {
		if props[name].requiresCorrectTool {
			sb.WriteString(fmt.Sprintf("\t%q: true,\n", "minecraft:"+name))
		}
	}
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
	fmt.Printf("block hardness: extracted %d blocks\n", len(props))
}

// parseBlockProps extracts block name -> props from Blocks.java source.
// Also resolves helper methods (e.g. logProperties, leavesProperties) that
// contain .strength() calls referenced from register() statements.
func parseBlockProps(src string) map[string]*blockProps {
	result := make(map[string]*blockProps)

	registerRe := regexp.MustCompile(`register\(\s*\n?\s*"([a-z_]+)"`)
	strengthRe := regexp.MustCompile(`\.strength\(\s*(-?[0-9.]+)F?`)
	indestructibleRe := regexp.MustCompile(`\.indestructible\(\)`)
	requiresToolRe := regexp.MustCompile(`\.requiresCorrectToolForDrops\(\)`)

	// first pass: extract helper methods that return BlockBehaviour.Properties
	// pattern: private static ... methodName(...) { return BlockBehaviour.Properties... .strength(X)... }
	helpers := parseHelperMethods(src, strengthRe, indestructibleRe, requiresToolRe)

	matches := registerRe.FindAllStringSubmatchIndex(src, -1)

	for i, match := range matches {
		name := src[match[2]:match[3]]

		start := match[0]
		var end int
		if i+1 < len(matches) {
			end = matches[i+1][0]
		} else {
			end = len(src)
		}

		semiIdx := strings.Index(src[start:end], ";")
		if semiIdx == -1 {
			semiIdx = end - start
		}
		stmt := src[start : start+semiIdx]

		bp := &blockProps{}

		if indestructibleRe.MatchString(stmt) {
			bp.hardness = -1
		} else if sm := strengthRe.FindStringSubmatch(stmt); sm != nil {
			val := strings.TrimSuffix(sm[1], "F")
			val = strings.TrimSuffix(val, "f")
			if f, err := strconv.ParseFloat(val, 32); err == nil {
				bp.hardness = float32(f)
			}
		} else {
			// try resolving helper method calls
			for helperName, helperProps := range helpers {
				if strings.Contains(stmt, helperName+"(") {
					bp.hardness = helperProps.hardness
					if helperProps.requiresCorrectTool {
						bp.requiresCorrectTool = true
					}
					break
				}
			}
		}

		if requiresToolRe.MatchString(stmt) {
			bp.requiresCorrectTool = true
		}
		result[name] = bp
	}

	return result
}

// parseHelperMethods extracts private/static methods in Blocks.java that build
// BlockBehaviour.Properties and contain .strength() calls.
func parseHelperMethods(src string, strengthRe, indestructibleRe, requiresToolRe *regexp.Regexp) map[string]*blockProps {
	helpers := make(map[string]*blockProps)

	// match: private static ... methodName(
	methodRe := regexp.MustCompile(`private\s+static\s+\S+\s+(\w+)\s*\(`)
	methodMatches := methodRe.FindAllStringSubmatchIndex(src, -1)

	for _, mm := range methodMatches {
		methodName := src[mm[2]:mm[3]]
		// find the method body: from the opening { to the matching }
		bodyStart := strings.Index(src[mm[0]:], "{")
		if bodyStart == -1 {
			continue
		}
		bodyStart += mm[0]
		depth := 1
		bodyEnd := bodyStart + 1
		for bodyEnd < len(src) && depth > 0 {
			switch src[bodyEnd] {
			case '{':
				depth++
			case '}':
				depth--
			}
			bodyEnd++
		}
		body := src[bodyStart:bodyEnd]

		bp := &blockProps{}
		if indestructibleRe.MatchString(body) {
			bp.hardness = -1
		} else if sm := strengthRe.FindStringSubmatch(body); sm != nil {
			val := strings.TrimSuffix(sm[1], "F")
			val = strings.TrimSuffix(val, "f")
			if f, err := strconv.ParseFloat(val, 32); err == nil {
				bp.hardness = float32(f)
			}
		}
		bp.requiresCorrectTool = requiresToolRe.MatchString(body)

		if bp.hardness != 0 || bp.requiresCorrectTool {
			helpers[methodName] = bp
		}
	}

	return helpers
}
