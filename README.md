# go-mclib/data

Go bindings for Minecraft Java Edition network protocol and game data, for use with [go-mclib/protocol](https://github.com/go-mclib/protocol) in [go-mclib/client](https://github.com/go-mclib/client) and other projects.

## Features

- **Packet definitions** for all protocol states (handshake, status, login, configuration, play)
- **Generated data** from Minecraft server reports:
  - 95 registries with bidirectional lookups
  - 1,166 blocks with state calculations
  - 1,505 items with component metadata
  - Packet ID mappings
- **ItemStack middleware** for typed component access

## Dependency Chain

[go-mclib/protocol](https://github.com/go-mclib/protocol) <–––(requires)––– **[go-mclib/data](https://github.com/go-mclib/data)** <–––(requires)––– [go-mclib/client](https://github.com/go-mclib/client)

## Installation

```bash
go get github.com/go-mclib/data
```

## Usage

See package documentation:

- [pkg/data/README.md](./pkg/data/README.md) - Registries, blocks, items
- [pkg/packets/README.md](./pkg/packets/README.md) - Packet definitions

## Updating to a New Minecraft Version

1. **Regenerate data** from server reports:

   ```bash
   cd pkg/data && go generate ./... && go fmt ./...
   ```

2. **Update packet bindings** - see `.claude/prompts/PACKETS.md`

3. **Update component decoders** - see `.claude/prompts/DATA.md`

## License

Licensed under the [MIT License](./LICENSE.md).
