# go-mclib/data

Go bindings for Minecraft Java Edition **26.3 only** (protocol **777**, data version **5023**), for use with [go-mclib/protocol](https://github.com/go-mclib/protocol) in [go-mclib/client](https://github.com/go-mclib/client) and other projects.

## Features

- **Packet definitions** for all protocol states (handshake, status, login, configuration, play)
- **Generated data** from Minecraft server reports:
  - 95 registries with bidirectional lookups
  - Blocks with state calculations
  - 1,658 items with component metadata
  - Packet ID mappings
  - Vanilla collision boxes and voxel grid planes, tag-bound suffocation/motion/fluid-flow flags, sturdy faces, and contained water/lava states
- **ItemStack middleware** for typed component access
- Entity metadata consumes full vanilla 26.3 particle option payloads, including effect colors, item templates and vibration destinations, so following metadata such as held-item use flags remains aligned.
- Brewing and crafting recipes preserve result component patches; contextual fuel providers and predicates are available as server-side registry data.

## Dependency Chain

[go-mclib/protocol](https://github.com/go-mclib/protocol) <–––(requires)––– **[go-mclib/data](https://github.com/go-mclib/data)** <–––(requires)––– [go-mclib/client](https://github.com/go-mclib/client)

## Installation

```bash
go get github.com/go-mclib/data
```

The current protocol dependency also requires the unpublished `github.com/go-mclib/nbt` workspace module; provide it locally when building.

## Usage

See package documentation:

- [PACKAGES.md](./PACKAGES.md) - Registries, blocks, items
- [packets/README.md](./packets/README.md) - Packet definitions

## Updating to a New Minecraft Version

1. **Generate the server JSON reports** (downloads the vanilla server jar and runs its data generator):

   ```bash
   (cd vanilla_server_reports && ./generate.sh <version>)
   ```

2. **Dump collision, fluid and entity metadata** directly from the matching vanilla jar (run from the module root):

   ```bash
   ./mcdump/dump.sh <version>
   ```

   The dumper binds vanilla block tags before reading state predicates and identifies source/flowing fluid types directly.

3. **Regenerate data** from the server reports, dumped metadata and decompiled source (`../mcsrc/current`):

   ```bash
   go generate ./... && go fmt ./...
   ```

4. **Update packet bindings** - see `.claude/prompts/PACKETS.md`

5. **Update component decoders** - see `.claude/prompts/DATA.md`

## License

Licensed under the [MIT License](./LICENSE.md).
