# CLAUDE.md

`go-mclib/data` provides Go bindings for Minecraft: Java Edition network protocol and game data, built on `go-mclib/protocol`.

## Folders

- `./packets` — Go bindings for packet structures (see `./packets/README.md`)
- `./blocks`, `./items`, `./entities`, `./registries`, `./lang`, `./packetids` — game data bindings with serialization; auto-generated (see `./generate/README.md`)
- `./chunks`, `./storage`, `./slots` — hand-written runtime helpers (chunk containers, world storage, player inventory slot layout)
- `./vanilla_server_reports` — downloads the vanilla server jar and runs its data generator → blocks/items/registries/packets/commands/tags JSON
- `./mcdump` — standalone Java tool that boots the server jar's registries and dumps collision shapes, entity hitboxes, hardness, mob categories, light, friction, effect colors and crafting remainders straight from the game (replaced PrismarineJS, the decompiled-source scrapers and hand-kept Go tables)

Decompiled Minecraft source (for reading protocol internals) lives in the sibling [go-mclib/mcsrc](https://github.com/go-mclib/mcsrc) repo — checked out as `../mcsrc`, latest version in `../mcsrc/current/`.

## Workflow

1. Examine the decompiled source in `../mcsrc/current/`; diff against a previous version (e.g. `diff -r ../mcsrc/<old>/net/minecraft/network ../mcsrc/current/net/minecraft/network`) to find protocol changes.
2. Regenerate the input JSON: `cd vanilla_server_reports && ./generate.sh <version>` and `cd mcdump && ./dump.sh <version>`.
3. Regenerate the Go bindings from the module root: `go generate ./... && go fmt ./...`
4. Write/adjust hand-written packet bindings in `./packets` per the diffed source.
5. Repeat when Minecraft updates the protocol.

## Testing

There are no packet round-trip fixtures in this repo (the old proxy-captured corpus was removed — it went stale on every version bump and was not maintained). Game-data correctness is covered by the unit tests in the data packages (block-state IDs, item IDs, registries, packet IDs validated against the authoritative server reports). End-to-end protocol validation lives as integration/smoke tests in [go-mclib/client](https://github.com/go-mclib/client) (join a vanilla server, exchange a chat message, disconnect).
