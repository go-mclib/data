import re
from json import load
from pathlib import Path

ROOT = Path(__file__).parent
PROTOCOL_VERSION = "772"
JSON_FILE = ROOT.parent.parent / "data" / PROTOCOL_VERSION / "packets.json"
GENERATE_INTO = ROOT

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


def load_packets_json() -> dict:
    """Load packets from JSON file."""
    with JSON_FILE.open("r") as f:
        return load(f)


def normalize_field_name(name: str) -> str:
    """Convert field name to Go naming convention."""
    # remove parentheses content, clean up special chars
    name = re.sub(r"\(.+\)", "", name)
    name = name.replace("-", " ").replace("/", " ").replace("?", "").replace(":", "")
    # remove numbers and convert to TitleCase
    name = re.sub(r"\d+", "", name)
    return "".join(word.capitalize() for word in name.split())


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


def generate_packets_go() -> None:
    """Generate Go packet definitions from JSON."""
    packets_data = load_packets_json()

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
                    packet_state=state.title(),
                    packet_bound=packet_bound,
                    packet_id=packet["id"],
                    fields="\n".join(fields) if fields else "\t// No fields",
                )

                structs.append(struct_code)
            output_file = GENERATE_INTO / FILE_MAPPINGS[(bound, state)]
            file_content = GO_FILE_TEMPLATE.format(structs="\n\n".join(structs))
            output_file.write_text(file_content)
            print(f"Generated {output_file}")


if __name__ == "__main__":
    import subprocess

    generate_packets_go()
    try:
        subprocess.run(["go", "fmt", "./..."], cwd=ROOT.parent, check=True)
        print("Formatted Go files")
    except subprocess.CalledProcessError as e:
        print(f"Warning: Failed to format Go files: {e}")
