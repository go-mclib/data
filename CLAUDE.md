# CLAUDE.md

`go-mclib/data` provides Go bindings for Minecraft: Java Edition network protocol and game data, built on `go-mclib/protocol`.

## Folders

- `./pkg/packets` — Go bindings for packet structures (see `./pkg/packets/README.md`)
- `./pkg/data` — game data bindings (blocks, items, entities, etc.) with serialization; auto-generated (see `./pkg/data/generate/README.md`)
- `./vanilla_server_reports` — downloads the vanilla server jar and runs its data generator → blocks/items/registries/packets/commands/tags JSON
- `./mcdump` — standalone Java tool that boots the server jar's registries and dumps collision shapes, entity hitboxes, hardness, mob categories straight from the game (replaced PrismarineJS and the decompiled-source scrapers)

Decompiled Minecraft source (for reading protocol internals) lives in the sibling [go-mclib/mcsrc](https://github.com/go-mclib/mcsrc) repo — checked out as `../mcsrc`, latest version in `../mcsrc/current/`.

## Workflow

1. Examine the decompiled source in `../mcsrc/current/`; diff against a previous version (e.g. `diff -r ../mcsrc/<old>/net/minecraft/network ../mcsrc/current/net/minecraft/network`) to find protocol changes.
2. Regenerate the input JSON: `cd vanilla_server_reports && ./generate.sh <version>` and `cd mcdump && ./dump.sh <version>`.
3. Regenerate the Go bindings: `cd pkg/data && go generate ./... && go fmt ./...`
4. Write/adjust hand-written packet bindings in `./pkg/packets` per the diffed source.
5. Repeat when Minecraft updates the protocol.

## Testing

There are no packet round-trip fixtures in this repo (the old proxy-captured corpus was removed — it went stale on every version bump and was not maintained). Game-data correctness is covered by the unit tests under `./pkg/data` (block-state IDs, item IDs, registries, packet IDs validated against the authoritative server reports). End-to-end protocol validation lives as integration/smoke tests in [go-mclib/client](https://github.com/go-mclib/client) (join a vanilla server, exchange a chat message, disconnect).
