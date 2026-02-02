# Data Update Prompt

Update hard-coded data logic for a new Minecraft version.

## Scope

This prompt covers hand-written data logic that cannot be auto-generated from server reports:

- **Item component encoders/decoders**: `pkg/data/items/item_stack_codecs.go`
- **Custom packet decoders**: `pkg/decoders/`

## Item Component Codecs

### What needs updating

The codec functions handle item component wire format:

- `decodeComponentWire()` - reads components from wire format
- `encodeComponentWire()` - writes components to wire format

When Minecraft adds, removes, or changes component structures:

1. **New components**: Add a new `case ComponentXxx:` block to both functions
2. **Changed format**: Update the read/write sequence in the existing case
3. **Removed components**: Remove the case (or keep for backwards compat)

### How to find changes

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

### Typed component encoding

For components that have typed Go struct fields in `Components`:

1. Update `applyComponent()` to parse wire bytes into typed fields
2. Update `encodeComponent()` to encode typed fields back to wire bytes
3. Update `clearComponent()` to handle removal
4. Update `componentDiffers()` to detect changes from defaults

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
