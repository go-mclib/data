# mcdump

A standalone Java program that dumps Minecraft game data to JSON straight from
the game's own registries — the **source of truth** for data that can't be read
from the vanilla server datagen (collision shapes, entity hitboxes, hardness,
mob categories, light emission/dampening, friction, effect colors, crafting
remainders).

No Fabric, no Gradle, no remapping: the bundled vanilla **server jar already
ships with real Mojang class/method names** and all its libraries, so `dump.sh`
just puts them on the classpath, compiles `Dump.java`, and runs it headless.

## What it dumps

`./dump.sh` writes to `../generate/mcdump/`:

- **`blocks.json`** — `{ blocks: { "minecraft:stone": { protocolId, hardness, requiresTool, friction, states: [{ id, properties, shape, lightEmission, lightDampening, … }] } }, shapes: [[[minX,minY,minZ,maxX,maxY,maxZ], …], …] }`. `id` = global block-state (palette) id; `shape` indexes the deduped `shapes` table (real `BlockState.getCollisionShape().toAabbs()`).
- **`entities.json`** — `{ "minecraft:zombie": { protocolId, width, height, eyeHeight, category, trackingRange } }`.
- **`registries.json`** — every built-in registry → `{ entry name: protocolId }`.
- **`potions.json`** — each built-in potion identifier → its complete list of `MobEffectInstance` codec values, including empty-effect potions.
- **`effects.json`** — `{ "minecraft:speed": { color } }` (`MobEffect.getColor`).
- **`items.json`** — `{ "minecraft:water_bucket": { craftingRemainder } }` for every item with `Item.getCraftingRemainder`.

These feed `gen_shapes.go`, `gen_block_hardness.go`, `gen_entity_dimensions.go`,
`gen_entities.go`, `gen_potions.go`, and `gen_game_data.go`. The dumper loads vanilla tags before evaluating state
predicates. Item component values and registry reports come from vanilla server
datagen; recipes, registry definitions and tags also use the decompiled assets.

## Usage

```bash
./dump.sh                       # version defaults to ../generate/version.json; needs
                                #   ../vanilla_server_reports/server-<version>.jar
                                #   (run ../vanilla_server_reports/generate.sh first)
cd .. && go generate ./...
```

`Dump.java` and `dump.sh` are committed; the extracted jars (`.work/`) and the
dumped `*.json` are not (the JSON is reproducible; only the generated Go bindings
are committed).
