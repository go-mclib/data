# Generate Go Bindings

Generate the input JSON, then run `go generate ./...` from the **data module root**.
The JSON files are reproducible and not committed — only the generated Go
bindings are.

## Generating the input JSON

Two producers, both reading the actual game:

```bash
# Run both commands from the data module root.
# 1. vanilla server datagen -> blocks/items/registries/packets/commands + tags
(cd vanilla_server_reports && ./generate.sh 26.3)

# 2. mcdump -> collision shapes, entity hitboxes, hardness, mob categories
./mcdump/dump.sh 26.3
```

## Files

Manual:

- `component_metadata.include.json`: Item component metadata;
- `entity_metadata.include.json`: Entity metadata definitions;

From the vanilla server datagen (`../vanilla_server_reports/`):

- `blocks.json`: Each block and its state enumerations mapped to their protocol ID (used to verify block-state ID calculation);
- `items/`: Each item and its components;
- `packets.json`: All network packets and their protocol IDs;
- `registries.json`: Base-game registries and their protocol IDs;

From [mcdump](../mcdump) (`mcdump/`, dumped straight from the game registries):

- `blocks.json`: per-state collision shapes, hardness, requires-correct-tool;
- `entities.json`: hitbox width/height, eye height, mob category, tracking range;
- `registries.json`: built-in registry id maps;

From decompiled assets ([go-mclib/mcsrc](https://github.com/go-mclib/mcsrc), checked out as `../mcsrc` relative to the data module root):

- `en_us.json`: English translations for all translation keys (items, blocks, UI, etc.);
- Synchronized registry entries, block/item tags, recipes and enchantments;
- Contextual integer/float providers and their predicates, including nested resource paths.
