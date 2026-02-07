package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	baseDir := filepath.Dir(os.Args[0])
	if len(os.Args) > 1 {
		baseDir = os.Args[1]
	}

	// load JSON data
	registries := loadJSON[map[string]RegistryJSON](filepath.Join(baseDir, "registries.json"))
	blocks := loadJSON[map[string]BlockJSON](filepath.Join(baseDir, "blocks.json"))
	items := loadJSON[map[string]ItemJSON](filepath.Join(baseDir, "items.json"))
	packets := loadJSON[PacketsJSON](filepath.Join(baseDir, "packets.json"))
	langPath := filepath.Join(baseDir, "en_us.json")

	outDir := filepath.Dir(baseDir)
	genDir := filepath.Dir(outDir) // go back one level to generate/

	// generate version info
	generateVersion(filepath.Join(outDir, "version_gen.go"))

	// generate packages
	generateRegistries(registries, filepath.Join(outDir, "registries", "registries_gen.go"))
	generateBlocks(registries, filepath.Join(outDir, "blocks", "blocks_gen.go"))
	generateBlockStates(blocks, registries, filepath.Join(outDir, "blocks", "block_states_gen.go"))
	generateItems(items, registries, filepath.Join(outDir, "items", "items_gen.go"))
	generateComponentTypes(registries, filepath.Join(outDir, "items", "item_components_gen.go"))
	generateComponentCodecs(registries, filepath.Join(genDir, "generate", "component_metadata.include.json"), filepath.Join(outDir, "items", "item_components_codec_gen.go"))
	generatePacketIds(packets, filepath.Join(outDir, "packet_ids"))
	generateLang(langPath, filepath.Join(outDir, "lang", "lang_gen.go"))
	generateEntities(registries, filepath.Join(outDir, "entities", "entities_gen.go"))
	generateEntityMetadata(filepath.Join(baseDir, "entity_metadata.include.json"), filepath.Join(outDir, "entities"))
	generateBlockShapes(blocks, filepath.Join(baseDir, "prismarine_block_collision_shapes.json"), filepath.Join(outDir, "hitboxes", "blocks", "block_shapes_gen.go"))

	fmt.Println("generation complete")
}
