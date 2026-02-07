# `packets_test`

Probably the most important package in this repo, as this is the one that's responsible for validating real packet captures against the Go bindings, e.g. whether they decode/encode correctly compared to the Java client implementation.

## Structure

The `packets_test.go` file contains the main `capturedPackets` map, which maps each Go representation of a packet to its corresponding wire (raw bytes) data.

Each packet is then defined in a separate file, and added to the `capturedPackets` map in the `init` function. As a result, running `go test ./... -v` in directory of this README file will run the tests and validate the captured packets against the Go bindings.

## Validation Approach

Each packet is validated in two phases:

1. **Decode correctness** — the captured bytes are decoded into a packet struct,
   then both the decoded and expected structs are re-encoded through our encoder
   and compared at the byte level. This validates that our decoder handles real
   server traffic correctly, while normalizing differences like nil vs empty slices
   and component ordering (the server may encode components in a different order
   than our encoder, which always uses ascending component ID order).

2. **Encoder round-trip** — the expected struct is encoded, decoded back, and
   re-encoded. The two encoded byte sequences must be identical, proving that
   our encoder produces deterministic, valid output.
