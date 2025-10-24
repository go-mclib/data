#!/usr/bin/env python3
"""
Fixed packet IDs in java_packets.json of `./data`,
according to packets.json report (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Data_Generators#Packets_report)

Sometimes the Wiki lacks behind, which is understandable because reverse engineering the protocol is a very time
consuming process that takes a lot of effort from the Minecraft community.
This script fixes the IDs automatically, so go-mclib can work on newer Minecraft versions faster.

Quickstart:
1. Download official server.jar and put it in the root of this dir (e. g.: ./scripts/server.jar).
2. Dump the packets.json: `java -jar -DbundlerMainClass="net.minecraft.data.Main" server.jar --all`.
3. Find packets.json in `generated/reports/packets.json` (should be symlinked to `./packets.json`).
4. Run `fix_packet_ids.py <protocol_version>` - this will update packet IDs in `./data` to corrent ones.
"""

import argparse
import json
from pathlib import Path

ROOT = Path(__file__).parent.parent
CORRECT_PACKET_IDS = ROOT / "scripts" / "packets.json"
DATA_DIR = ROOT / "data"


def main():
    parser = argparse.ArgumentParser(description="Fix packet IDs in java_packets.json")
    parser.add_argument("protocol_id", type=str, help="Protocol version ID")
    args = parser.parse_args()

    # correct
    with open(CORRECT_PACKET_IDS, "r") as f:
        correct_packets = json.load(f)

    # to fix
    java_packets_path = DATA_DIR / args.protocol_id / "java_packets.json"
    with open(java_packets_path, "r") as f:
        java_packets = json.load(f)

    # update packet IDs
    missing_mappings = []
    updated_count = 0

    for state, directions in correct_packets.items():
        if state not in java_packets:
            continue

        for direction, packets in directions.items():
            if direction not in java_packets[state]:
                continue

            # mapping from resource name to protocol_id
            for packet_key, packet_data in packets.items():
                resource = packet_key.removeprefix("minecraft:")
                protocol_id = packet_data["protocol_id"]

                found = False
                for java_packet in java_packets[state][direction]:
                    if java_packet.get("resource") == resource:
                        java_packet["id"] = f"0x{f'{protocol_id:02x}'.upper()}"
                        updated_count += 1
                        found = True
                        break

                if not found:
                    missing_mappings.append(f"{state}/{direction}/{resource}")

    # write
    with open(java_packets_path, "w") as f:
        json.dump(java_packets, f, indent=2)

    print(f"Updated {updated_count} packet(s) in {java_packets_path}")

    # report missing
    if missing_mappings:
        print(f"\nMissing mappings ({len(missing_mappings)}):")
        for mapping in missing_mappings:
            print(f"  - {mapping}")


if __name__ == "__main__":
    main()
