# Minecraft Data Packages

Go bindings for Minecraft protocol data including registries, blocks, block states, items, item components, entities, and translations.

## Packages

### `registries`

Contains all 95 Minecraft registries with bidirectional lookups.

```go
import "github.com/go-mclib/data/registries"

// get protocol ID for a block
id := registries.Block.Get("minecraft:stone")  // returns 1

// reverse lookup
name := registries.Block.ByID(1)  // returns "minecraft:stone"

// available registries
registries.Block          // 1286 entries
registries.Item           // 1658 entries
registries.EntityType     // 161 entries
registries.MobEffect      // 40 entries
// ... 95 total registries
```

### `blocks`

Contains block protocol IDs, block state calculations, and lookups.

```go
import "github.com/go-mclib/data/blocks"

// block ID constants
blocks.Stone          // 1
blocks.OakPlanks      // 15
blocks.DiamondBlock   // 126

// string to ID
id := blocks.ID("minecraft:oak_planks")  // 15

// ID to string
name := blocks.Name(15)  // "minecraft:oak_planks"

// calculate state ID from block + properties
stateID := blocks.StateID(blocks.OakDoor, map[string]string{
    "facing": "north",
    "half":   "lower",
    "hinge":  "left",
    "open":   "false",
    "powered": "false",
})

// reverse lookup: get block and properties from state ID
blockID, props := blocks.StateProperties(stateID)

// get default state for a block
defaultID := blocks.DefaultStateID(blocks.OakDoor)

// per-state game data dumped from the game by mcdump
blocks.LightEmission(stateID)   // BlockState.getLightEmission
blocks.LightDampening(stateID)  // BlockState.getLightDampening
blocks.Friction(stateID)        // Block.getFriction
```

### `items`

Contains item protocol IDs, lookups, default component data, and slot decoding.

```go
import "github.com/go-mclib/data/items"

// item ID constants
items.DiamondSword    // 876
items.Apple           // 918
items.IronPickaxe     // 860

// string to ID
id := items.ID("minecraft:diamond_sword")  // 876

// ID to string
name := items.Name(876)  // "minecraft:diamond_sword"

// get default components
comps := items.DefaultComponents(items.Apple)
if comps.Food != nil {
    fmt.Printf("nutrition: %d\n", comps.Food.Nutrition)  // 4
}
```

#### Stack and the Component Patch Model

In the Minecraft protocol, an item on the wire is not sent as a full component list.
Instead, the Java client uses a `DataComponentPatch` system: each item type has a
*prototype* (default components), and the wire format only carries a **patch** — a list
of component **adds** (overrides) and **removes** (deletions) relative to that prototype.
Components that match the prototype are never sent.

`items.Stack` mirrors this model. Each `Components` struct carries an internal bitset that
tracks which component IDs are explicitly *present*. `ToSlot()` only encodes present
components that differ from the item's defaults:

```go
// create a stack with only specific components (sparse patch)
stack := items.NewStackWithComponents(items.DiamondSword, 1, &items.Components{
    CustomName: &items.ItemNameComponent{Text: "Excalibur"},
    Unbreakable: true,
})
// ToSlot() encodes just these 2 components; defaults are untouched

// create a stack from defaults (full component set)
stack := items.NewStack(items.DiamondSword, 1)
stack.Components.Damage = 100
// ToSlot() encodes only Damage (the one field that differs from defaults)

// read from wire — full state reconstruction
stack, err := items.ReadSlot(buf)
// FromSlot applies the incoming patch on top of defaults;
// all components are marked as present for faithful re-encoding

// convert from/to raw Slot
stack, err := items.FromSlot(rawSlot)
rawSlot, err := stack.ToSlot()
```

Optional components are pointers (`nil` = absent). `NewStack` and `FromSlot`
track every component for patch encoding; only hand-built `Components` need
`SetPresent`, and `ClearPresent` excludes a component from the patch:

```go
stack.Components.SetPresent(items.ComponentDamage)   // include in the patch
stack.Components.ClearPresent(items.ComponentDamage) // neither add nor remove it
```

Component type constants (104 types) are generated from the registry:

```go
items.ComponentDamage         // 3
items.ComponentFood           // 23
items.ComponentEnchantments   // 13
items.MaxComponentID          // 103
```

### `entities`

Contains entity type protocol IDs, lookups, and entity metadata parsing.

```go
import "github.com/go-mclib/data/entities"

// entity type ID constants
entities.Player       // 128
entities.Zombie       // 156
entities.Creeper      // 32

// string to ID
id := entities.TypeID("minecraft:zombie")  // 156

// ID to string
name := entities.TypeName(156)  // "minecraft:zombie"
```

#### Entity Metadata

Entity metadata is read/written using the wire format: `[Index(UByte)][Type(VarInt)][Value]...[0xFF terminator]`

```go
// read metadata from packet buffer
metadata, err := entities.ReadMetadata(buf)

// access raw data by index
if data := metadata.Get(entities.EntityIndexFlags); data != nil {
    flags := data[0]
    isOnFire := flags&0x01 != 0
    isSneaking := flags&0x02 != 0
}

// set/update metadata
metadata.Set(entities.EntityIndexPose, entities.SerializerPOSE, poseData)

// write metadata back to packet buffer
err := entities.WriteMetadata(buf, metadata)
```

Single entries can be built with typed constructors (`ByteEntry`, `IntEntry`, `VarIntEntry`, `FloatEntry`, `BoolEntry`, `StringEntry`, `ComponentEntry`, `OptionalComponentEntry`, `RotationsEntry`, `SlotEntry`):

```go
metadata = append(metadata, entities.BoolEntry(entities.EntityIndexSilent, true), entities.FloatEntry(9, 20))
```

Serializer types (39 types) define how values are encoded on the wire:

```go
entities.SerializerBYTE      // 0 - single byte
entities.SerializerINT       // 1 - VarInt
entities.SerializerFLOAT     // 3 - float32
entities.SerializerSTRING    // 4 - length-prefixed string
entities.SerializerBOOLEAN   // 8 - single byte bool
entities.SerializerROTATIONS // 9 - 3x float32 (x,y,z)
entities.SerializerBLOCK_POS // 10 - packed position
```

Per-entity field indices are generated:

```go
// Entity base class (all entities inherit these)
entities.EntityIndexFlags         // 0 - on_fire|sneaking|sprinting|swimming|invisible|glowing|fall_flying
entities.EntityIndexAirSupply     // 1 - air ticks remaining
entities.EntityIndexCustomName    // 2 - optional name component
entities.EntityIndexPose          // 6 - standing/sneaking/sleeping/etc.

// Creeper-specific (inherits from Mob)
entities.CreeperIndexSwellDir     // 16 - fuse state (-1=idle, 1=fuse)
entities.CreeperIndexIsPowered    // 17 - charged creeper
entities.CreeperIndexIsIgnited    // 18 - ignited

// Player-specific (inherits from LivingEntity)
entities.PlayerIndexAdditionalHearts // 15 - absorption hearts
entities.PlayerIndexSkinParts        // 17 - visible skin parts flags
```

### `lang`

Contains English translations for Minecraft translation keys.

```go
import "github.com/go-mclib/data/lang"

// translate a key to English
text := lang.Translate("item.minecraft.iron_sword")  // "Iron Sword"
text := lang.Translate("block.minecraft.stone")      // "Stone"

// returns empty string if key not found
text := lang.Translate("invalid.key")  // ""
```

## Code Generation

The packages are generated from Minecraft server reports. To regenerate:

```bash
# from the module root
go generate ./...
go fmt ./...
```

JSON data files must be present in `generate/` (symlinked from `vanilla_server_reports/generated/reports/`).

## Block State Algorithm

Block state IDs are calculated using a right-to-left multiplier approach:

```plain
offset = Σ(property_value_index[i] × ∏(cardinality[j] for j < i))
state_id = base_id + offset
```

Properties are iterated right-to-left, where the rightmost property changes fastest. This matches Minecraft's internal state registration order.

## Lookups

`StateID` computes the ID directly from the generated property tables (no
cache); `StateProperties` and `StateBlockID` use O(log n) binary search over
the state ranges.

## Performance

Benchmarks on Apple M2 (`go test -bench=. -benchmem ./...`):

| Function | Time | Allocations |
| -------- | ---- | ----------- |
| `StateID` | ~84 ns/op | 0 |
| `StateProperties` | ~138 ns/op | 2 |

## Data Sources

Generated from Minecraft server reports and the mcdump game dump (neither
committed). To regenerate:

```bash
# generate server reports + version.json (see vanilla_server_reports/)
(cd vanilla_server_reports && ./generate.sh)
# dump game data (shapes, hardness, light, friction, effect colors, crafting remainders)
./mcdump/dump.sh

# regenerate Go code
go generate ./... && go fmt ./...
```

`ProtocolVersion`, `MinecraftVersion` and `DataVersion` (root package
`version_gen.go`) come from the server jar's `version.json`.

Source JSON files:

- `blocks.json`: 1,286 blocks, 35,723 total states, 93 unique properties
- `items/`: 1,658 items, 122 component types
- `registries.json`: 95 registries

## Testing

```bash
# from the module root
go test -v ./...
go test -bench=. -benchmem ./...
```

The test suite verifies all 29,000+ block states against the source JSON.
