# Update Packet Bindings

Update Go packet bindings for a new Minecraft version.

## Instructions

Read and follow `.claude/prompts/PACKETS.md` to update packet bindings for the new Minecraft version.

Before starting:

1. Verify `./decompiled/current/` and `./decompiled/previous/` exist
2. Check the current Minecraft version in `go.mod` or recent commits

After completing:

1. Run `go fmt ./...`
2. Run `go build ./...` to verify compilation
3. Summarize changes for the user
