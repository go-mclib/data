package main

import (
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

// EntityMetadataJSON represents the entity_metadata.include.json structure.
type EntityMetadataJSON struct {
	Serializers map[string]SerializerJSON `json:"serializers"`
	Entities    map[string]EntityJSON     `json:"entities"`
}

type SerializerJSON struct {
	Name     string `json:"name"`
	WireType string `json:"wireType"`
}

type EntityJSON struct {
	Parent string      `json:"parent"`
	Fields []FieldJSON `json:"fields"`
}

type FieldJSON struct {
	Index       int    `json:"index"`
	Name        string `json:"name"`
	Serializer  int    `json:"serializer"`
	GoField     string `json:"goField"`
	GoType      string `json:"goType"`
	Description string `json:"description,omitempty"`
	Passthrough bool   `json:"passthrough,omitempty"`
}

func generateEntityMetadata(includePath, outDir string) {
	metadata := loadJSON[EntityMetadataJSON](includePath)

	generateSerializers(metadata.Serializers, filepath.Join(outDir, "entity_metadata_serializers_gen.go"))
	generateEntityMetadataTypes(metadata, filepath.Join(outDir, "entity_metadata_gen.go"))
}

func generateSerializers(serializers map[string]SerializerJSON, outPath string) {
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("entities"))

	sb.WriteString("// Entity metadata serializer type IDs.\n")
	sb.WriteString("const (\n")

	// sort by ID (numeric)
	ids := make([]int, 0, len(serializers))
	for idStr := range serializers {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(fmt.Errorf("serializer id %q: %w", idStr, err))
		}
		ids = append(ids, id)
	}
	slices.Sort(ids)

	for _, id := range ids {
		s := serializers[strconv.Itoa(id)]
		fmt.Fprintf(&sb, "\tSerializer%s = %d\n", s.Name, id)
	}
	sb.WriteString(")\n\n")

	sb.WriteString("// serializerWireTypes maps serializer IDs to wire types.\n")
	sb.WriteString("var serializerWireTypes = map[int32]string{\n")
	for _, id := range ids {
		s := serializers[strconv.Itoa(id)]
		fmt.Fprintf(&sb, "\t%d: %q,\n", id, s.WireType)
	}
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
}

func generateEntityMetadataTypes(metadata EntityMetadataJSON, outPath string) {
	var sb strings.Builder
	sb.WriteString(generatedFileHeader("entities"))

	// generate field index constants per entity
	sb.WriteString("// Entity metadata field indices.\n")

	for _, name := range sortedKeys(metadata.Entities) {
		entity := metadata.Entities[name]
		if len(entity.Fields) == 0 {
			continue
		}

		fmt.Fprintf(&sb, "\n// %s metadata indices\n", name)
		sb.WriteString("const (\n")
		for _, f := range entity.Fields {
			fmt.Fprintf(&sb, "\t%sIndex%s = %d\n", name, f.GoField, f.Index)
		}
		sb.WriteString(")\n")
	}

	sb.WriteString("\n// MetadataEntry represents a single entity metadata entry.\n")
	sb.WriteString("type MetadataEntry struct {\n")
	sb.WriteString("\tIndex      byte\n")
	sb.WriteString("\tSerializer int32\n")
	sb.WriteString("\tData       []byte // raw wire data\n")
	sb.WriteString("}\n")

	writeFile(outPath, sb.String())
}
