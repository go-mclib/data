# Minecraft Packet Capture Proxy

A MITM (Man-in-the-Middle) proxy for capturing Minecraft Java Edition packets. Captured packets can be used to generate test fixtures for validating packet serialization.

## Requirements

- Go 1.25+
- A Minecraft server running in **offline mode** (online-mode=false in server.properties)
- A Minecraft client

## Minimal Usage

Run the vanilla server locally in offline-mode (no encryption) at `localhost:25566`. Then, start the proxy:

```bash
go run ./cmd/proxy -target localhost:25566 [options]
```

Connect your Minecraft client to `localhost:25565`. The proxy will:

1. Accept the client connection
2. Connect to the real server at localhost:25566
3. Forward all packets between client and server
4. Log packet information (with `-verbose`)
5. Save captured packets when the session ends

### Options

| Flag | Default | Description |
| ---- | ------- | ----------- |
| `-help` | false | Show help message |
| `-target` | (required) | Target server address (host:port) |
| `-port` | 25565 | Port to listen on for client connections |
| `-output` | captures | Directory to save captured packets |
| `-verbose` | false | Enable verbose logging of all packets |
| `-state` | (all) | Comma-separated states to capture (e.g., `login,play`) |
| `-packetId` | (all) | Comma-separated packet IDs to capture (e.g., `0x00,0x01`) |

### Filtering Examples

Capture only login and configuration packets:

```bash
go run ./cmd/proxy -target localhost:25566 -state login,configuration
```

Capture only specific packet IDs across all states:

```bash
go run ./cmd/proxy -target localhost:25566 -packetId 0x00,0x02,0x03
```

Combine filters to capture specific packets in specific states:

```bash
go run ./cmd/proxy -target localhost:25566 -state play -packetId 0x00
```

## Output Format

Captured packets are saved as JSON files in the output directory:

```plain
captures/
  session_1_20260201_143052.json
  session_2_20260201_143215.json
```

Each file contains an array of captured packets:

```json
[
  {
    "direction": "c2s",
    "state": "handshake",
    "packet_id": "0x00",
    "raw_data": "060974657374686f7374633909c0020402"
  },
  {
    "direction": "s2c",
    "state": "login",
    "packet_id": "0x03",
    "raw_data": "ff01"
  }
]
```

### Fields

| Field | Description |
| ----- | ----------- |
| `direction` | `c2s` (client to server) or `s2c` (server to client) |
| `state` | Protocol state: `handshake`, `status`, `login`, `configuration`, or `play` |
| `packet_id` | Hex packet ID within the current state (e.g., `0x00`) |
| `raw_data` | Hex-encoded packet payload (excludes packet ID) |

## Decoding Captured Packets

Use the decode command to pretty-print captured packets:

```bash
go run ./cmd/decode captures/session_*.json
```

Example output:

```plain
Decoding 5 packets from captures/session_1_20260201_143052.json

// [0] c2s 0x00
C2SIntention {
  ProtocolVersion VarInt = 774
  ServerAddress String = "localhost"
  ServerPort Uint16 = 25565
  Intent VarInt = 2
}

// [1] c2s 0x00
C2SHello {
  Name String = "Player123"
  PlayerUuid UUID = 550e8400-e29b-41d4-a716-446655440000
}

// [2] WARNING: unknown packet 0x07 in configuration_s2c
```

## Using Captures for Tests

1. Run the proxy and connect a client to generate captures
2. Copy the capture file to `packets_test/fixtures/`:

   ```bash
   cp captures/session_1_*.json ../packets_test/fixtures/
   ```

3. Run the tests:

   ```bash
   go test ./packets_test/ -v
   ```

The test framework will:

- Load all JSON files from `packets_test/fixtures/`
- Parse each packet using the appropriate packet struct
- Verify round-trip serialization (read then write should produce identical bytes)

## Protocol States

The proxy automatically tracks state transitions:

```plain
Handshake ─┬─> Status (intent=1) ─> [ping/pong]
           │
           └─> Login (intent=2) ─> Configuration ─> Play
```

State transitions are detected by:

- **Handshake → Status/Login**: Intent field in C2S Handshake packet (0x00)
- **Login → Configuration**: S2C Login Success packet (0x02)
- **Configuration → Play**: S2C/C2S Finish Configuration packet (0x03)

Compression is enabled when the server sends S2C Set Compression (0x03) during login.

## Limitations

- **Offline mode only**: The proxy cannot intercept encrypted connections. Online-mode servers use encryption after the login handshake, which prevents packet inspection.
- **No packet modification**: The proxy forwards packets unmodified. It's designed for capture, not injection.
- **Single protocol version**: Designed for the latest protocol version. Other versions may have different packet IDs.

## Verbose Output

With `-verbose`, the proxy logs each packet in stdout:

```plain
2026/02/01 14:30:52 [proxy] c2s: state=handshake id=0x00 len=18
2026/02/01 14:30:52 [proxy] state -> login
2026/02/01 14:30:52 [proxy] c2s: state=login id=0x00 len=19
2026/02/01 14:30:52 [proxy] s2c: state=login id=0x03 len=2
2026/02/01 14:30:52 [proxy] compression threshold -> 256
2026/02/01 14:30:52 [proxy] s2c: state=login id=0x02 len=38
2026/02/01 14:30:52 [proxy] state -> configuration
```

## Troubleshooting

### "connection refused" when connecting to target

- Verify the Minecraft server is running on the specified port
- Check firewall settings

### Client shows "Encryption not supported" or similar

- The target server must have `online-mode=false` in server.properties
- Restart the server after changing this setting

### Empty or missing capture files

- Ensure the output directory is writable
- Check for error messages in the proxy output
- If using filters, ensure at least some packets match

### Packets not parsing correctly in tests

- All packets must correctly parse in the Go bindings/tests - if this is not the case, update either `go-mclib/protocol` or `go-mclib/data` so it parses correctly
- Verify you're testing against the correct Minecraft version
