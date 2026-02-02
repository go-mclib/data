# Packet Bindings Update Prompt

Update Go packet bindings for a new Minecraft version.

## Steps

1. **Compute diff against the previous Minecraft version**
   - Ensure `./decompiled/current/` and `./decompiled/previous/` exist
   - If not, see `./decompiled/README.md` for decompilation instructions
   - Run: `diff -r ./decompiled/current/net/minecraft/network ./decompiled/previous/net/minecraft/network`
   - Check [Minecraft Wiki Protocol](https://minecraft.wiki/w/Java_Edition_protocol/Packets) for documentation

2. **Examine differences in the network protocol**
   - Changes may affect packet definitions or fundamental network data types
   - If core types change, discuss with maintainer - may require changes to `go-mclib/protocol`

3. **Update Go bindings**
   - Packet definitions are in `pkg/packets/`
   - Packet IDs are generated in `pkg/data/packets/`
   - Follow patterns in `pkg/packets/README.md`
   - Run `go fmt ./...` after changes

4. **Summarize changes**
   - Highlight breaking changes
   - Note new packets, removed packets, and field changes

5. **Suggest improvements** (optional)
   - Code style, API design, readability improvements

## Context

- [go-mclib/data](https://github.com/go-mclib/data): Packet bindings and game data (this repo)
- [go-mclib/protocol](https://github.com/go-mclib/protocol): Low-level primitives (VarInt, NBT, etc.)
- [go-mclib/client](https://github.com/go-mclib/client): Minecraft bot client

## References

- [Java Edition Protocol](https://minecraft.wiki/w/Java_Edition_protocol)
- [Protocol Data Types](https://minecraft.wiki/w/Java_Edition_protocol/Data_types)
- [Packet List](https://minecraft.wiki/w/Java_Edition_protocol/Packets)
