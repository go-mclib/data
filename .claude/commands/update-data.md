# Update Data Logic

Update hard-coded data logic (item components, decoders) for a new Minecraft version.

## Instructions

Read and follow `.claude/prompts/DATA.md` to update data logic for the new Minecraft version.

Key files:

- `pkg/data/items/item_stack_decoders.go` - Item component wire format
- `pkg/decoders/` - Custom packet decoders

After completing:

1. Run `go fmt ./...`
2. Run `go build ./...` to verify compilation
3. Summarize changes for the user
