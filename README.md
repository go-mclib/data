# go-mclib/data

A repository that contains Go bindings for Minecraft packet types, for use with [go-mclib/protocol](https://github.com/go-mclib/protocol) in [go-mclib/client](https://github.com/go-mclib/client) and other projects.

## Dependency Chain

[go-mclib/protocol](https://github.com/go-mclib/protocol) <–––(requires)––– **[go-mclib/data](https://github.com/go-mclib/data)** <–––(requires)––– [go-mclib/client](https://github.com/go-mclib/client)

## Usage

Download the packet bindings for specific Minecraft protocol version.

## Updating to a new Minecraft version

Also see headers of Python script files in [`scripts/`](./scripts/) for more detailed instructions.

1. Download HTML (Ctrl/Cmd + U and copy all text) of <https://minecraft.wiki/w/Java_Edition_protocol/Packets>
2. Run `mkdir -p data/774; python scripts/import_java_packets.py > data/774/java_packets.json`, replacing `774` with the actual protocol version
3. See [`fix_packet_ids.py`](./scripts/fix_packet_ids.py): `cd scripts; java -jar -DbundlerMainClass="net.minecraft.data.Main" server.jar --all`, then `python fix_packet_ids.py 774` (again, replace `774` with the protocol ID you imported in step 2).
4. Finally, (assuming your WD is at repo root, if not `cd` to dir of this file), generate the packets as Go bindings: `mkdir -p go/774/java_packets; python scripts/generate_java_packets.py 774`. Manually inspect the files, fixing packet structures as necessary (the script is not perfect and sometimes imports the packets incorrectly - in most cases, it's enough to simply copy the definition from the protocol version before).

## License

Licensed under the [MIT License](./LICENSE.md).
