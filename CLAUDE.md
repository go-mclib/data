# CLAUDE.md

The `go-mclib/data` package provides Go bindings for the Minecraft: Java Edition network protocol and game data. It uses `go-mclib/protocol` under the hood to work with primitive types and data structures, to build higher-level abstractions on top of them.

Goal: to provide packages for reading and writing Minecraft: Java Edition packets in Go.

- `./proxy` - A MITM proxy for capturing Minecraft: Java Edition packets. Captured packets can be used to generate test fixtures for validating packet serialization, see `./proxy/README.md`;
- `./pkg/packets` - Go bindings for the structure of Minecraft: Java Edition packets, see `./pkg/packets/README.md`;
- `./pkg/data` - Go bindings for the game data of Minecraft: Java Edition - stuff like blocks, items and their components, entities, etc., and methods to serialize and deserialize them to and from bytes (so they can be sent over the network). A large portion of the data is auto-generated from Minecraft server reports, see `./pkg/data/generate/README.md`;
- `./decompiled` - a sub-directory, used to decompile the Minecraft: Java Edition source code, see `./decompiled/README.md`. For latest version of source code that this repository assumes compatibility with, see `./decompiled/current/`;

High-level overview:

1. We examine the decompiled source code (see `./pkg/data/generate/README.md`), or <https://minecraft.wiki/w/Java_Edition_protocol> wiki pages, to understand the structure of the packets and the game data, how they are encoded and decoded, etc... The wiki is well-documented, but sometimes outdated, so as a last-resort, we use the actual source code of the game as source of truth;
2. We write Go bindings for the packets and the game data;
3. We create relevant unit tests. Unit tests are setup in a way where we use the `./proxy` to capture packets, then we write the raw bytes of the packets and what we expect them to look like in the Go format into the test files (for example, see `./pkg/packets/c2s_play_test.go`). We aim for this: the Go package will always be able to decode and encode the packets so that they match the captured bytes 1:1, that ensures the implementation is correct;
4. If Minecraft updates the protocol, we rinse and repeat.

## Code Style

- comments that are not part of symbol or package docstrings should start with lowercase letter, unless naming an exported symbol;
- comments should be minimal;
- comments should only document parts of the code that are less obvious;
- tests should be written in separate packages, ending with `_test`;
- after done, format the code with `go fmt ./...`;
- keep the API minimal, do not keep legacy code - let user know after making a breaking change, though;
- assume modern Go 1.25+ API, do not use older APIs;
