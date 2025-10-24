#!/usr/bin/env python3
# Usage: python scripts/generate_java_packets.py [protocol_version]
import re
import argparse
from json import load
from pathlib import Path

ROOT = Path(__file__).parent.parent
DATA_DIR = ROOT / "data"
GO_DIR = ROOT / "go"

FILE_MAPPINGS = {
    ("serverbound", "handshaking"): "c2s_handshaking.go",
    ("serverbound", "status"): "c2s_status.go",
    ("serverbound", "configuration"): "c2s_configuration.go",
    ("serverbound", "login"): "c2s_login.go",
    ("serverbound", "play"): "c2s_play.go",
    ("clientbound", "handshaking"): "s2c_handshaking.go",
    ("clientbound", "status"): "s2c_status.go",
    ("clientbound", "configuration"): "s2c_configuration.go",
    ("clientbound", "login"): "s2c_login.go",
    ("clientbound", "play"): "s2c_play.go",
}

GO_FILE_TEMPLATE = """package packets

import (
	jp "github.com/go-mclib/protocol/java_protocol"
	ns "github.com/go-mclib/protocol/net_structures"
)

{structs}"""

GO_STRUCT_TEMPLATE = """// {struct_name} represents "{packet_name}".
{packet_notes}
//
// https://minecraft.wiki/w/Java_Edition_protocol/Packets#{minecraft_wiki_anchor}
var {struct_name} = jp.NewPacket(jp.State{packet_state}, jp.{packet_bound}, {packet_id})

type {struct_name}Data struct {{
{fields}
}}"""

GO_FIELD_TEMPLATE = """	// {field_notes}
	{field_name} {field_type}"""


def load_packets_json(protocol_version: str) -> dict:
    """Load packets from JSON file."""
    json_file = DATA_DIR / protocol_version / "java_packets.json"
    with json_file.open("r") as f:
        return load(f)


def normalize_field_name(name: str) -> str:
    """Convert field name to Go naming convention."""
    # remove parentheses content, clean up special chars
    name = re.sub(r"\(.+\)", "", name)
    name = name.replace("-", " ").replace("/", " ").replace("?", "").replace(":", "")
    # Convert to TitleCase while preserving numbers
    words = name.split()
    result = []
    for word in words:
        # If the word contains both letters and numbers, capitalize the first letter
        # Otherwise just capitalize normally
        if any(c.isdigit() for c in word) and any(c.isalpha() for c in word):
            # Find first letter and capitalize it
            new_word = ""
            first_letter_found = False
            for c in word:
                if c.isalpha() and not first_letter_found:
                    new_word += c.upper()
                    first_letter_found = True
                else:
                    new_word += c
            result.append(new_word)
        else:
            result.append(word.capitalize())
    final_name = "".join(result)

    # If the name starts with a number, prefix it with 'Field' to make it a valid Go identifier
    if final_name and final_name[0].isdigit():
        final_name = "Field" + final_name

    return final_name


def normalize_resource_name(packet_resource: str, packet_name: str) -> str:
    """Convert resource name to Go struct name part."""
    # extract state suffix if present in packet name
    state_suffix = ""
    state_match = re.search(r"\s*\(([^)]+)\)\s*$", packet_name)
    if state_match:
        state = state_match.group(1).lower()
        if state in ["play", "configuration", "login", "status"]:
            state_suffix = state.capitalize()

    if packet_resource and packet_resource not in ["", "unknown"]:
        base_name = "".join(
            word.capitalize() for word in packet_resource.replace("_", " ").split()
        )
    else:
        clean_name = re.sub(r"\s*\([^)]*\)\s*$", "", packet_name)
        base_name = "".join(word.capitalize() for word in clean_name.split())

    return base_name + state_suffix


def format_packet_notes(notes: str) -> str:
    """Format packet notes for Go comments."""
    if not notes:
        return "//"
    lines = notes.strip().split("\n")
    formatted = ["//"] + [f"// > {line}" for line in lines]
    return "\n".join(formatted)


def generate_packets_go(protocol_version: str) -> None:
    """Generate Go packet definitions from JSON."""
    packets_data = load_packets_json(protocol_version)
    generate_into = GO_DIR / protocol_version / "java_packets"

    # Ensure output directory exists
    generate_into.mkdir(parents=True, exist_ok=True)

    for state, bounds in packets_data.items():
        for bound, packets in bounds.items():
            if not packets:
                continue

            structs = []

            for packet in packets:
                fields = []
                for field in packet["fields"]:
                    field_name = normalize_field_name(field["name"])
                    field_type = field["type"]  # type is already Go-ready from JSON
                    # Handle inline struct syntax - make it multiline for readability
                    if "struct {" in field_type and ";" in field_type:
                        # Convert inline struct to multiline format
                        field_type = field_type.replace("struct { ", "struct {\n\t\t")
                        field_type = field_type.replace("; ", "\n\t\t")
                        field_type = field_type.replace(" }", "\n\t}")
                        # Remove trailing semicolon before closing brace if present
                        field_type = field_type.replace(";\n\t}", "\n\t}")
                    field_notes = field["notes"].replace("\n", " ").strip()

                    if field_name and field_type:
                        fields.append(
                            GO_FIELD_TEMPLATE.format(
                                field_notes=field_notes,
                                field_name=field_name,
                                field_type=field_type,
                            )
                        )
                packet_bound = "C2S" if bound == "serverbound" else "S2C"
                struct_name = packet_bound + normalize_resource_name(
                    packet["resource"], packet["name"]
                )

                wiki_anchor = packet["name"].title().replace(" ", "_")
                packet_notes = format_packet_notes(packet["notes"])

                struct_code = GO_STRUCT_TEMPLATE.format(
                    struct_name=struct_name,
                    packet_name=packet["name"],
                    packet_notes=packet_notes,
                    minecraft_wiki_anchor=wiki_anchor,
                    packet_state="Handshake" if state == "handshaking" else state.title(),
                    packet_bound=packet_bound,
                    packet_id=packet["id"],
                    fields="\n".join(fields) if fields else "\t// No fields",
                )

                structs.append(struct_code)
            output_file = generate_into / FILE_MAPPINGS[(bound, state)]
            file_content = GO_FILE_TEMPLATE.format(structs="\n\n".join(structs))
            output_file.write_text(file_content)
            print(f"Generated {output_file}")


def get_all_protocol_versions() -> list[str]:
    """Get all protocol versions from the data directory."""
    versions = []
    if DATA_DIR.exists():
        for path in DATA_DIR.iterdir():
            if path.is_dir() and path.name.isdigit():
                versions.append(path.name)
    return sorted(versions)


if __name__ == "__main__":
    import subprocess

    parser = argparse.ArgumentParser(description="Generate Go packet definitions from JSON")
    parser.add_argument(
        "protocol_version",
        nargs="?",
        help="Protocol version to generate (e.g., 772). If not provided, generates for all versions.",
    )
    args = parser.parse_args()

    if args.protocol_version:
        versions = [args.protocol_version]
    else:
        versions = get_all_protocol_versions()
        if not versions:
            print("No protocol versions found in data directory")
            exit(1)
        print(f"Generating for all versions: {', '.join(versions)}")

    for version in versions:
        print(f"\n=== Processing protocol version {version} ===")
        try:
            generate_packets_go(version)
            go_output_dir = GO_DIR / version
            if go_output_dir.exists():
                subprocess.run(["go", "fmt", "./..."], cwd=go_output_dir, check=True)
                print(f"Formatted Go files for version {version}")
        except FileNotFoundError as e:
            print(f"Warning: Skipping version {version} - {e}")
        except subprocess.CalledProcessError as e:
            print(f"Warning: Failed to format Go files for version {version}: {e}")
