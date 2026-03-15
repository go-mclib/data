package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func generateEntities(registries map[string]RegistryJSON, entityTypeJavaPath, outPath string) {
	entityRegistry := registries["minecraft:entity_type"]

	// extract MobCategory from EntityType.java
	categories := extractMobCategories(entityTypeJavaPath)

	var sb strings.Builder
	sb.WriteString(generatedFileHeader("entities"))

	// mob category map
	sb.WriteString("var entityCategory = map[string]string{\n")
	for _, name := range sortedKeys(entityRegistry.Entries) {
		cat := categories[name]
		if cat == "" {
			cat = "misc"
		}
		sb.WriteString(fmt.Sprintf("\t%q: %q,\n", name, cat))
	}
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
}

// extractMobCategories parses EntityType.java to extract the MobCategory for each entity.
func extractMobCategories(javaPath string) map[string]string {
	data, err := os.ReadFile(javaPath)
	if err != nil {
		fmt.Printf("skipping mob categories: %v\n", err)
		return nil
	}
	src := string(data)

	registerRe := regexp.MustCompile(`register\(\s*"([^"]+)"`)
	categoryRe := regexp.MustCompile(`MobCategory\.(\w+)`)

	categories := make(map[string]string)
	matches := registerRe.FindAllStringIndex(src, -1)
	for _, loc := range matches {
		nameMatch := registerRe.FindStringSubmatch(src[loc[0]:loc[1]])
		if nameMatch == nil {
			continue
		}
		name := "minecraft:" + nameMatch[1]

		block := findRegisterBlock(src, loc[0])
		if block == "" {
			continue
		}

		catMatch := categoryRe.FindStringSubmatch(block)
		if catMatch == nil {
			continue
		}
		categories[name] = strings.ToLower(catMatch[1])
	}

	fmt.Printf("mob categories: extracted %d entities from EntityType.java\n", len(categories))
	return categories
}
