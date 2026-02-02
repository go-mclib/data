# Data Update Prompt

Update hard-coded data logic for a new Minecraft version.

## Scope

This prompt covers hand-written data logic that cannot be auto-generated from server reports:

- **Item component decoders**: `pkg/data/items/item_stack_decoders.go`
- **Custom packet decoders**: `pkg/decoders/`

## Item Component Decoders

### What needs updating

The `decodeComponentWire()` function reads item components from wire format. When Minecraft adds, removes, or changes component structures:

1. **New components**: Add a new `case ComponentXxx:` block
2. **Changed format**: Update the read sequence in the existing case
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

### Current encoding gap

The `encodeComponent()` function only encodes simple components. Complex components like `Tool`, `Enchantments`, `Equippable`, etc. need encoding implementations to support `ItemStack.ToSlot()`.

**To add encoding for a component**:

1. Add a case in `encodeComponent()` that mirrors the decode logic
2. Update `applyComponent()` to parse the component into typed fields
3. Update `clearComponent()` to handle removal
4. Update `componentDiffers()` to detect changes from defaults

### Wire format patterns

| Pattern | Read | Write |
|---------|------|-------|
| VarInt | `buf.ReadVarInt()` | `w.WriteVarInt(v)` |
| Float | `buf.ReadFloat32()` | `w.WriteFloat32(v)` |
| String | `buf.ReadString(maxLen)` | `w.WriteString(v)` |
| NBT | `copyNBT(buf, w)` | (same) |
| Optional | `present, _ := buf.ReadBool(); if present { ... }` | (same) |
| Array | `count, _ := buf.ReadVarInt(); for i := 0; i < count; i++ { ... }` | (same) |
| HolderSet | `copyHolderSet(buf, w)` | (same) |
| SoundEvent | `copySoundEvent(buf, w)` | (same) |

## References

- [Slot Data](https://minecraft.wiki/w/Java_Edition_protocol/Slot_data)
- [Data Component Format](https://minecraft.wiki/w/Data_component_format)
- [Data Generators](https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Data_Generators)
