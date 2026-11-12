# Data Update Prompt

Update hard-coded data logic for a new Minecraft version.

## Scope

This prompt covers hand-written data logic that cannot be auto-generated from server reports:

- **Item component codecs**: `pkg/data/items/item_components_codec.go` (hand-written), `pkg/data/items/item_components_codec_gen.go` (auto-generated from metadata)
- **Codec registry**: `pkg/data/items/item_components_codec_registry.go`
- **Component metadata**: `pkg/data/generate/component_metadata.include.json`

## Item Component Codecs

### Architecture

Each component has a `componentCodec` implementing:

- `DecodeWire(buf)` - reads wire bytes into an intermediate representation
- `Apply(stack, decoded)` - applies decoded data to the typed `Components` struct
- `Clear(stack)` - resets the component to zero value
- `Differs(stack, defaults)` - checks if a component differs from defaults
- `Encode(stack, buf)` - writes the component to wire format

Codecs are registered from `init` in `item_components_codec_registry.go` and `component_payloads.go` via `registerCodec(id, codec)`; `registerCodec` panics when an ID is registered twice. Components that only need their wire bytes preserved use `registerPassthrough(id, decode)`, where `decode` copies the payload (`decodeVarIntWire`, `decodeNBTWire`, `decodeHolderSetWire`, ...).

### Auto-generated codecs

Each entry in `component_metadata.include.json` describes one component:

- `goField` / `goType`: the `Components` field and its type. Default item literals in `items_gen.go` are written by the per-component writers in `componentLiteralWriters` (`generate/gen_items.go`), which take the field name from `goField`.
- `wireType`: `varint`, `float32`, `identifier`, `empty` (plus `allowZero` for pointer VarInt fields such as `MapID`) generate simple typed codecs; `struct` with `wireFormat` generates a full struct codec (Food, Weapon, Enchantable, TooltipDisplay, UseCooldown, VillagerFood); `manual` means items registers the codec by hand, so the generator emits nothing for it.
- `passthrough: true` with a basic `wireType` (`varint`, `bool`, `identifier`, `empty`, `int32`, `nbt`, `holderSet`, `slotList`, `slot`) generates a `registerPassthrough` call.

Registry components with a `/` in their name that have no metadata entry are generated as VarInt entity-variant passthroughs; add an entry (e.g. `"wireType": "manual"`) to opt out. To add or update generated codecs:

1. Update `component_metadata.include.json` (and `componentLiteralWriters` if the component appears in vanilla item defaults)
2. Run `cd pkg/data && go generate ./...`
3. The generator produces registrations and codec implementations in `item_components_codec_gen.go`

### Hand-written codecs

Complex codecs (enchantments, attribute modifiers, lore, tool rules, etc.) remain hand-written in `item_components_codec.go`. When Minecraft changes these:

1. **Diff decompiled source**:

   ```bash
   diff -r ../mcsrc/current/net/minecraft/world/item/component \
           ../mcsrc/previous/net/minecraft/world/item/component
   ```

2. **Check Minecraft Wiki**:
   - [Slot Data](https://minecraft.wiki/w/Java_Edition_protocol/Slot_data)
   - [Data Component Format](https://minecraft.wiki/w/Data_component_format)

3. **Check component registry** for new IDs:

   ```bash
   diff ./vanilla_server_reports/generated/reports/registries.json \
        ./vanilla_server_reports_previous/generated/reports/registries.json
   ```

### Wire format patterns

| Pattern    | Read                                                           | Write                  |
| ---------- | -------------------------------------------------------------- | ---------------------- |
| VarInt     | `buf.ReadVarInt()`                                             | `w.WriteVarInt(v)`     |
| Float      | `buf.ReadFloat32()`                                            | `w.WriteFloat32(v)`    |
| String     | `buf.ReadString(maxLen)`                                       | `w.WriteString(v)`     |
| NBT        | `copyNBT(buf, w)`                                              | (same)                 |
| Optional   | `present, _ := buf.ReadBool(); if present { ... }`             | (same)                 |
| Array      | `count, _ := buf.ReadVarInt(); for i := range int(count) {...}`| (same)                 |
| HolderSet  | `copyHolderSet(buf, w)`                                        | (same)                 |
| SoundEvent | `copySoundEvent(buf, w)`                                       | (same)                 |

## References

- [Slot Data](https://minecraft.wiki/w/Java_Edition_protocol/Slot_data)
- [Data Component Format](https://minecraft.wiki/w/Data_component_format)
- [Data Generators](https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Data_Generators)
