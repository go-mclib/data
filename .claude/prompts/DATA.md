# Data Update Prompt

Update hard-coded data logic for a new Minecraft version.

## Scope

This prompt covers hand-written data logic that cannot be auto-generated from server reports:

- **Item component codecs**: `pkg/data/items/item_components_codec.go` (hand-written), `pkg/data/items/item_components_codec_gen.go` (auto-generated from metadata)
- **Codec registry**: `pkg/data/items/item_components_codec_registry.go`
- **Component metadata**: `pkg/data/generate/component_metadata.include.json`

## Item Component Codecs

### Architecture

Each component has a `ComponentCodec` implementing:

- `DecodeWire(buf)` - reads wire bytes into an intermediate representation
- `Apply(stack, decoded)` - applies decoded data to the typed `Components` struct
- `Clear(stack)` - resets the component to zero value
- `Differs(stack, defaults)` - checks if a component differs from defaults
- `Encode(stack, buf)` - writes the component to wire format

Codecs are registered in `item_components_codec_registry.go` via `RegisterCodec()`.

### Auto-generated codecs

Simple struct codecs (Food, Weapon, Enchantable, TooltipDisplay, UseCooldown, Fireworks) are generated from field-level metadata in `component_metadata.include.json`. To add or update these:

1. Update `wireFormat` in `component_metadata.include.json` with field definitions
2. Run `cd pkg/data && go generate ./...`
3. The generator produces codec implementations in `item_components_codec_gen.go`

### Hand-written codecs

Complex codecs (enchantments, attribute modifiers, lore, tool rules, etc.) remain hand-written in `item_components_codec.go`. When Minecraft changes these:

1. **Diff decompiled source**:

   ```bash
   diff -r ./decompiled/current/net/minecraft/world/item/component \
           ./decompiled/previous/net/minecraft/world/item/component
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
