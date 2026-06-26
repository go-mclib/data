# mcdump

A standalone Java program that dumps Minecraft game data to JSON straight from
the game's own registries — the **source of truth** for data that can't be read
from the vanilla server datagen (collision shapes, entity hitboxes, hardness,
mob categories).

No Fabric, no Gradle, no remapping: the bundled vanilla **server jar already
ships with real Mojang class/method names** and all its libraries, so `dump.sh`
just puts them on the classpath, compiles `Dump.java`, and runs it headless.

## What it dumps

`./dump.sh` writes to `../pkg/data/generate/mcdump/`:

- **`blocks.json`** — `{ blocks: { "minecraft:stone": { protocolId, hardness, requiresTool, states: [{ id, properties, shape }] } }, shapes: [[[minX,minY,minZ,maxX,maxY,maxZ], …], …] }`. `id` = global block-state (palette) id; `shape` indexes the deduped `shapes` table (real `BlockState.getCollisionShape().toAabbs()`).
- **`entities.json`** — `{ "minecraft:zombie": { protocolId, width, height, eyeHeight, category, trackingRange } }`.
- **`registries.json`** — every built-in registry → `{ entry name: protocolId }`.

These feed `gen_block_shapes.go`, `gen_block_hardness.go`, `gen_entity_hitboxes.go`,
and `gen_entities.go`. (Datapack-bound data — item component values, tags, biomes —
needs a loaded world and still comes from the vanilla server `--reports` datagen.)

## Usage

```bash
./dump.sh                       # needs ../vanilla_server_reports/server-26.2.jar
                                #   (run ../vanilla_server_reports/generate.sh first if missing)
cd ../pkg/data && go generate ./...
```

`Dump.java` and `dump.sh` are committed; the extracted jars (`.work/`) and the
dumped `*.json` are not (the JSON is reproducible; only the generated Go bindings
are committed).
