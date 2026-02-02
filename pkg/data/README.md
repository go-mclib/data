# Minecraft Data Packages

Go bindings for Minecraft protocol data including registries, blocks, block states, items, and item components.

## Packages

### `registries`

Contains all 95 Minecraft registries with bidirectional lookups.

```go
import "github.com/go-mclib/data/pkg/data/registries"

// get protocol ID for a block
id := registries.Block.Get("minecraft:stone")  // returns 1

// reverse lookup
name := registries.Block.ByID(1)  // returns "minecraft:stone"

// available registries
registries.Block          // 1166 entries
registries.Item           // 1505 entries
registries.EntityType     // 157 entries
registries.MobEffect      // 40 entries
// ... 95 total registries
```

### `blocks`

Contains block protocol IDs, block state calculations, and lookups.

```go
import "github.com/go-mclib/data/pkg/data/blocks"

// block ID constants
blocks.Stone          // 1
blocks.OakPlanks      // 15
blocks.DiamondBlock   // 126

// string to ID
id := blocks.BlockID("minecraft:oak_planks")  // 15

// ID to string
name := blocks.BlockName(15)  // "minecraft:oak_planks"

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
```

### `items`

Contains item protocol IDs, lookups, default component data, and slot decoding.

```go
import "github.com/go-mclib/data/pkg/data/items"

// item ID constants
items.DiamondSword    // 876
items.Apple           // 918
items.IronPickaxe     // 860

// string to ID
id := items.ItemID("minecraft:diamond_sword")  // 876

// ID to string
name := items.ItemName(876)  // "minecraft:diamond_sword"

// get default components
comps := items.DefaultComponents(items.Apple)
if comps.Food != nil {
    fmt.Printf("nutrition: %d\n", comps.Food.Nutrition)  // 4
}
```

#### ItemStack

`ItemStack` provides middleware over `net_structures.Slot` for typed component access:

```go
// read a slot from wire and convert to typed ItemStack
stack, err := items.ReadSlot(buf)
if !stack.IsEmpty() {
    fmt.Printf("item: %s x%d\n", items.ItemName(stack.ID), stack.Count)
    if stack.Components.Food != nil {
        fmt.Printf("nutrition: %d\n", stack.Components.Food.Nutrition)
    }
}

// create a new stack with default components
stack := items.NewStack(items.DiamondSword, 1)
stack.Components.Damage = 100  // modify durability

// write back to wire
err := stack.WriteSlot(buf)

// convert from/to raw Slot
stack, err := items.FromSlot(rawSlot)
rawSlot, err := stack.ToSlot()
```

Component type constants (104 types) are generated from the registry:

```go
items.ComponentDamage         // 3
items.ComponentFood           // 23
items.ComponentEnchantments   // 13
items.MaxComponentID          // 103
```

## Code Generation

The packages are generated from Minecraft server reports. To regenerate:

```bash
cd pkg/data
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

## Caching

`StateID` results are cached globally (default 4096 entries) for repeated lookups with the same inputs:

```go
// configure cache size
blocks.SetCacheSize(8192)  // increase cache
blocks.SetCacheSize(0)     // disable caching
blocks.ClearCache()        // clear cached entries
```

`StateProperties` uses O(log n) binary search and doesn't require caching.

## Performance

Benchmarks on Apple M2 (`go test -bench=. -benchmem ./...`):

| Function | Time | Allocations |
| -------- | ---- | ----------- |
| `StateID` (cached) | ~83 ns/op | 0 |
| `StateID` (uncached) | ~84 ns/op | 0 |
| `StateProperties` | ~138 ns/op | 2 |

## Data Sources (as of Minecraft 1.21.11)

Not committed to the repository, see [json_reports/README.md](./json_reports/README.md) for more details.

- `blocks.json`: 1,166 blocks, 29,671 total states, 92 unique properties
- `items.json`: 1,505 items, 104 component types
- `registries.json`: 95 registries

## Testing

```bash
cd pkg/data
go test -v ./...
go test -bench=. -benchmem ./...
```

The test suite verifies all 29,000+ block states against the source JSON.
