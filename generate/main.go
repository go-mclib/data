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
	version = loadJSON[versionJSON](filepath.Join(baseDir, "version.json"))
	registries := loadJSON[map[string]RegistryJSON](filepath.Join(baseDir, "registries.json"))
	blocks := loadJSON[map[string]BlockJSON](filepath.Join(baseDir, "blocks.json"))
	items := loadItems(baseDir)
	packets := loadJSON[PacketsJSON](filepath.Join(baseDir, "packets.json"))
	langPath := filepath.Join(baseDir, "en_us.json")

	outDir := filepath.Dir(baseDir)
	// decompiled sources live in the sibling go-mclib/mcsrc repo (../mcsrc relative to the module root)
	decompiledDir := filepath.Join(outDir, "..", "mcsrc", "current")
	datapackDir := filepath.Join(decompiledDir, "data", "minecraft")
	synchronized := loadSynchronizedEntries(datapackDir)

	// game data dumped by the standalone tool (mcdump) — collision shapes, entity
	// hitboxes, hardness, mob categories, light, friction, effect colors,
	// crafting remainders: straight from the game, no scraping.
	mcdumpDir := filepath.Join(baseDir, "mcdump")
	blockDump := loadJSON[mcdumpBlocks](filepath.Join(mcdumpDir, "blocks.json"))
	blockStates := blockDump.states()
	entityDump := loadJSON[map[string]mcdumpEntity](filepath.Join(mcdumpDir, "entities.json"))

	// generate version info
	generateVersion(filepath.Join(outDir, "version_gen.go"))

	// generate packages
	generateRegistries(registries, synchronized, filepath.Join(outDir, "registries", "registries_gen.go"))
	generateBlockStates(blocks, registries, filepath.Join(outDir, "blocks", "block_states_gen.go"))
	generateItems(items, registries, filepath.Join(baseDir, "component_metadata.include.json"), filepath.Join(outDir, "items", "items_gen.go"))
	generateComponentTypes(registries, filepath.Join(outDir, "items", "item_components_gen.go"))
	generateComponentCodecs(registries, filepath.Join(baseDir, "component_metadata.include.json"), filepath.Join(outDir, "items", "item_components_codec_gen.go"))
	generatePacketIds(packets, filepath.Join(outDir, "packetids"))
	generateLang(langPath, filepath.Join(outDir, "lang", "lang_gen.go"))
	generateEntities(entityDump, filepath.Join(outDir, "entities", "entities_gen.go"))
	generateEntityMetadata(filepath.Join(baseDir, "entity_metadata.include.json"), filepath.Join(outDir, "entities"))
	generateBlockShapes(&blockDump, blockStates, filepath.Join(outDir, "blocks", "shapes_gen.go"))
	generateEntityDimensions(entityDump, filepath.Join(outDir, "entities", "dimensions_gen.go"))
	generateRegistryData(datapackDir, synchronized, filepath.Join(outDir, "registries", "registry_data_gen.go"))
	generateTagData(filepath.Join(datapackDir, "tags"), filepath.Join(outDir, "registries", "tag_data_gen.go"))
	generateBlockHardness(&blockDump, filepath.Join(outDir, "blocks", "block_hardness_gen.go"))
	generateBlockSupport(&blockDump, blockStates, filepath.Join(outDir, "blocks", "block_support_gen.go"))
	generateMobEffectColors(filepath.Join(mcdumpDir, "effects.json"), filepath.Join(outDir, "registries", "mob_effect_colors_gen.go"))
	generateCraftingRemainders(filepath.Join(mcdumpDir, "items.json"), registries, filepath.Join(outDir, "items", "crafting_remainders_gen.go"))
	generateRecipes(filepath.Join(datapackDir, "recipe"), filepath.Join(outDir, "recipes", "recipes_gen.go"))
	generateEnchantments(filepath.Join(datapackDir, "enchantment"), filepath.Join(outDir, "enchantments", "enchantments_gen.go"))
	generatePotions(filepath.Join(mcdumpDir, "potions.json"), filepath.Join(outDir, "items", "potions_gen.go"))
	generateAdvancements(filepath.Join(datapackDir, "advancement"), filepath.Join(outDir, "advancements", "advancements_gen.go"))

	// fail loudly if any scraper produced implausibly few entries
	reportSanityChecks()

	fmt.Println("generation complete")
}
