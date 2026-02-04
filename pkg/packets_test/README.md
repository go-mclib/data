# `packets_test`

Probably the most important package in this repo, as this is the one that's responsible for validating real packet captures against the Go bindings, e.g. whether they decode/encode 1:1 with the Java client implementation.

## Structure

The `packets_test.go` file contains the main `capturedPackets` map, which maps each Go representation of a packet to its corresponding wire (raw bytes) data.

Each packet is then defined in a separate file, and added to the `capturedPackets` map in the `init` function. As a result, running `go test ./... -v` in directory of this README file will run the tests and validate the captured packets against the Go bindings.
