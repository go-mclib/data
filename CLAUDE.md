# CLAUDE.md

`go-mclib/data` provides Go bindings for Minecraft: Java Edition network protocol and game data, built on `go-mclib/protocol`.

## Folders

- `./proxy` — MITM proxy for capturing packets, used to generate test fixtures (see `./proxy/README.md`)
- `./pkg/packets` — Go bindings for packet structures (see `./pkg/packets/README.md`)
- `./pkg/data` — game data bindings (blocks, items, entities, etc.) with serialization; largely auto-generated from server reports (see `./pkg/data/generate/README.md`)

Decompiled Minecraft source (source of truth, used by the generator) lives in the sibling [go-mclib/mcsrc](https://github.com/go-mclib/mcsrc) repo — checked out as `../mcsrc`, latest version in `../mcsrc/current/`.

## Workflow

1. Examine decompiled source or <https://minecraft.wiki/w/Java_Edition_protocol> (wiki may be outdated — source is the source of truth)
2. Write Go bindings for packets and game data
3. Unit tests compare decoded/encoded packets against raw bytes captured by `./proxy`, ensuring 1:1 match
4. Repeat when Minecraft updates the protocol
