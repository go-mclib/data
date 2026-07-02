# packetdump

Stage 1 of packet-binding generation. Parses the decompiled Minecraft source
(`mcsrc`) with [JavaParser](https://javaparser.org/) and emits a
language-agnostic **wire-field IR** (`../pkg/data/generate/packets_ir.json`)
describing every clientbound/serverbound packet's fields in order.

The Go generator (Stage 2, `pkg/packets/gen_from_ir.go`, run with
`go run gen_from_ir.go`) consumes the IR and maps each token to an `ns.*` type to
emit structs + `Read`/`Write` into the `pkg/packets/pgen` package. Packet **IDs**
still come from the vanilla server report (`packets.json`); this tool only
recovers **structure**, which the report doesn't carry. `ID()`/`State()`/`Bound()`
are left to the existing `pkg/packets/generate.go` boilerplate pass.

Unlike `mcdump` this needs no server jar — it reads `.java` source, so it only
needs the JavaParser jar (fetched once into `.lib/`) and gson (already vendored
under `vanilla_server_reports/libraries`).

## Usage

```bash
./gen.sh                      # parses ../../mcsrc/current
./gen.sh /path/to/mcsrc       # alternate source tree
```

## How it reads packets

Minecraft defines packet codecs two ways; both are parsed:

- **`StreamCodec.composite(codec, accessor, …, ::new)`** — declarative. Each
  `(codec, accessor)` pair → one field; the accessor (`Packet::field`) gives the
  name, the codec expression gives the wire type.
- **`Packet.codec(::write, ::new)`** — imperative. The `(ByteBuf input)`
  constructor is walked statement by statement (`this.x = input.readVarInt()` or
  `this(input.readX(), …)` delegating to a record's canonical ctor).
- `StreamCodec.unit(...)` → empty packet.

It also resolves referenced types recursively: **cross-file** (follow imports to
the type's `.java`, parse its `STREAM_CODEC` or `(ByteBuf)` reader ctor, emit
under `types` as `ref:`; the alias table bounds recursion, e.g. `ItemStack`→
`slot`), **enums** (→ `enum:Name`, wire = varint), **wrapped reads**
(`GameType.byId(input.readByte())` → the inner read), and **bitfield collapse**
(`int flags = input.readByte(); this.x = (flags & 1)…` keeps only the
buffer-reading statements → one `flags` field, not N derived booleans).

## IR schema

```json
{ "packets": [
  { "clazz": "ClientboundAddEntityPacket", "name": "add_entity",
    "state": "play", "bound": "S2C", "codec": "packet.codec", "status": "ok",
    "fields": [ { "name": "id", "type": "varint" }, { "name": "uuid", "type": "uuid" }, … ] }
] }
```

Wire-type tokens are generic: `varint`, `varlong`, `bool`, `i8/u8/i16/i32/i64`,
`f32/f64`, `string`, `bytes`, `uuid`, `identifier`, `position`, `nbt`,
`textcomponent`, `slot`, `vec3d`, `enum:<Name>`, `list<T>`, `optional<T>`,
`map<K,V>`. Anything not yet mapped is emitted as **`raw:<JavaExpr>`** so gaps
are visible, never silently wrong.

`status` per packet: `ok` (every field resolved), `partial` (≥1 `raw:` token —
structure captured, some sub-type needs the Stage-2 alias table), `unparsed`
(codec shape not recognized, e.g. manual bitfields).

## Coverage (26.2)

Stage 1 — 223 packets: **168 ok**, 43 partial, 12 unparsed; 26 sub-types
resolved (file-local + cross-file) and emitted under `types`.

Stage 2 — **161 packets generate** into a compiling, round-trip-tested `pgen`
package; 62 are skipped (each listed with the offending token). The generator
maps complex fields via an **alias table** to existing `ns.*` types with a
uniform `Decode`/`Encode` (`ClientboundLevelChunkPacketData`→`ns.ChunkData`,
`ClientboundLightUpdatePacketData`→`ns.LightData`, `GameProfile`, `IDSet`, …),
and special-cases `slot` (`buf.ReadSlot(items.Decoder())`) and `nbt` (inline
streaming reader). The remaining skip tail is the hand-maintain bucket:

1. **`map<…>`** (tags, recipes, `set_time` clock updates) — no `PrefixedMap`
   helper exists yet.
2. **Signed-chat / custom-payload / cookies** — manual codecs (`raw:read`,
   `MANUAL-CTOR`, discarded-payload wrappers).
3. **Item/recipe/advancement structures** (`AdvancementHolder`, `RecipeDisplay`,
   full `ItemStack`, `Dialog`, `NumberFormatTypes`) — partial sub-types with a
   deep unresolved field.
4. A few niche scalars (`vec3i`, `soundevent` id-or-inline, `BlockHitResult`).

Improve Stage 1 coverage by growing the `BYTEBUF` / `READ` / `TYPES` tables in
`PacketDump.java`; improve Stage 2 by growing the `scalars` and `aliases` tables
in `gen_from_ir.go` (each entry unlocks several packets).
