"""
generate.py generates JSON definitions of Minecraft packets from a HTML capture of the
Minecraft Wiki page (https://minecraft.wiki/w/Java_Edition_protocol/Packets)

To use:
1. Visit https://minecraft.wiki/w/Java_Edition_protocol/Packets
2. Press Ctrl + U
3. Copy the HTML to `packets.html` in current directory (`scripts/packets.html`)
4. Run this script (e. g. from repository root, run `python scripts/import_packets.py`)

### How it works?

The packet list starts in the "Handshaking" header (so, after element `h2:has(span[id="Handshaking"])`).
Then, there is just a sequence of headers and paragraphs or tables:

```html
<h2 id="Handshaking">Handshaking</h2>
<p>Some bla bla text</p>
<h3>Clientbound</h3>
<p>there are no clientbound packets in "Handshaking" state</p>
<h3>Serverbound</h3>
<h4>Actual packet name</h4>
<table>
    <tbody>
        <tr>
            <td>Packet ID</td>
            <td>State</td>
            <td>Bound to</td>
            <td>Field name</td>
            <td>Field type</td>
            <td>Notes</td>
        </tr>
        <tr>
            <td>
                <i>protocol:</i>
                <br>
                <code>0x00</code>
                <br>
                <br>
                <i>resource:</i>
                <br>
                <code>some_registry_id</code>
            </td>
            <td>Handshaking</td>
            <td>Server</td>
            <td>Some field name</td>
            <td>
                <a href="#Type:VarInt">VarInt</a>
            </td>
            <td>
                Some bla bla notes text that can contain <a href="#some">links</a> and other stuff.
            </td>
        </tr>
    </tbody>
</table>
<h3>Clientbound</h3>
<table>...</table>
<h3>Serverbound</h3>
<p>...</p>
<table>...</table>
[...]
```

As we can see, it is mostly just a matter of using the correct selectors to find the elements and extract the text.
And maybe ignoring some elements/packets, such as "Legacy Server List Ping" packets, as they are not part of the modern protocol.

### Why not fetch the URL response directly from Python?

The Wiki is behind Cloudflare, so you have to extract the HTML manually.

## License

Output is licensed under the GNU General Public License v3.0
(per Minecraft.wiki's packets page license and https://creativecommons.org/share-your-work/licensing-considerations/compatible-licenses/)
"""

try:
    from bs4 import BeautifulSoup, Tag
except ImportError:
    raise ImportError("pip install beautifulsoup4")

from pathlib import Path
import re

ROOT = Path(__file__).parent
IMPORT_FILE = ROOT / "packets.html"
START_IMPORT_FROM_ELEMENT = 'h2:has(span[id="Handshaking"])'


def normalize_notes(text: str) -> str:
    """Collapse multiple spaces inside notes while preserving line breaks between paragraphs."""
    if not text:
        return ""
    lines = text.splitlines()
    normalized_lines = [re.sub(r"[ \t]+", " ", ln).strip() for ln in lines]
    return "\n".join(normalized_lines).strip()


def normalize_field_type(raw_type: str, target_format: str = "json") -> str:
    """
    Normalize field types from wiki format.

    Args:
        raw_type: Raw type string from wiki
        target_format: Output format - "json" for intermediate format, "go" for Go types
    """
    if not raw_type:
        return ""

    t = re.sub(r"\s+", " ", raw_type).strip()

    # drop trailing/standalone 'Enum' token(s)
    t = re.sub(r"\b[Ee]nums?\b", "", t).strip()
    t = re.sub(r"\s{2,}", " ", t).strip()

    def _prefixed_array_repl(match: re.Match) -> str:
        size = match.group(1)
        elem_type = match.group(2)
        return f"Prefixed{elem_type}Array({size})"

    t = re.sub(
        r"\bPrefixed\s+Array\s*\(\s*([0-9]+)\s*\)\s+of\s+([A-Za-z0-9]+)\b",
        _prefixed_array_repl,
        t,
    )
    t = re.sub(r"\bPrefixed\s+Array\s+of\s+([A-Za-z0-9]+)\b", r"Prefixed\1Array", t)

    def _array_size_of_repl(match: re.Match) -> str:
        size = match.group(1)
        elem_type = match.group(2)
        return f"{elem_type}Array({size})"

    t = re.sub(
        r"\bArray\s*\(\s*([0-9]+)\s*\)\s+of\s+([A-Za-z0-9]+)\b", _array_size_of_repl, t
    )
    t = re.sub(r"\bArray\s+of\s+([A-Za-z0-9]+)\b", r"\1Array", t)

    id_or_match = re.match(r"ID\s+or\s+(.+)", t, re.IGNORECASE)
    if id_or_match:
        inner_type = id_or_match.group(1).strip()
        inner_type_normalized = normalize_field_type(inner_type, "json")
        t = f"IDor{inner_type_normalized}"

    var_or_match = re.match(r"(\w+)\s+or\s+(.+)", t, re.IGNORECASE)
    if var_or_match:
        type1 = var_or_match.group(1).strip()
        type2 = var_or_match.group(2).strip()
        type1_norm = normalize_field_type(type1, "json")
        type2_norm = normalize_field_type(type2, "json")
        t = f"Or{type1_norm}{type2_norm}"

    t = re.sub(r"\b[Ee]num\b", "", t)
    t = re.sub(r"\s+", " ", t).strip()

    t = re.sub(r"\(\s*([^)]+)\s*\)", r"(\1)", t)

    t = re.sub(r"\s+", "", t)

    if target_format == "go":
        t = transform_to_go_type(t)

    return t


def transform_to_go_type(field_type: str) -> str:
    """Transform normalized field type to Go format with ns. prefix."""
    # but keep them for special handling if needed
    size_match = re.match(r"^(.+?)\((\d+)\)$", field_type)
    if size_match:
        base_type = size_match.group(1)
        field_type = base_type
    base_type = field_type.replace("-", "").replace("/", "")

    if not base_type:
        return "ns.Unknown // FIXME"
    if base_type == "Namespace":
        return "ns.Identifier"
    if base_type.lower() in ["seebelow", "see below", "(seebelow)", "(see below)"]:
        return "ns.Unknown // FIXME: See below"

    paren_match = re.match(r"^(.+)\((.+)\)$", base_type)
    if paren_match:
        main_type = paren_match.group(1)
        paren_content = paren_match.group(2).lower()
        if paren_content in ["seebelow", "see below"]:
            return f"ns.{main_type} // FIXME: See below"
        else:
            if paren_content.isdigit():
                return f"ns.{main_type} // Size: {paren_content}"
            else:
                return f"ns.{main_type} // FIXME: {paren_content}"
    id_or_match = re.match(r"^IDor(.+)$", base_type)
    if id_or_match:
        inner_type = id_or_match.group(1)
        inner_go_type = transform_to_go_type(inner_type)
        if inner_go_type.startswith("ns."):
            inner_go_type = inner_go_type[3:]
        return f"ns.Or[ns.Identifier, ns.{inner_go_type}]"

    or_match = re.match(r"^Or(.+)$", base_type)
    if or_match:
        combined = or_match.group(1)
        for boundary in [
            "TextComponent",
            "String",
            "VarInt",
            "Int",
            "Long",
            "Float",
            "Double",
            "Boolean",
            "Byte",
        ]:
            if boundary in combined and combined != boundary:
                if combined.endswith(boundary):
                    type1 = combined[: -len(boundary)]
                    type2 = boundary
                    type1_go = transform_to_go_type(type1)
                    type2_go = transform_to_go_type(type2)
                    if type1_go.startswith("ns."):
                        type1_go = type1_go[3:]
                    if type2_go.startswith("ns."):
                        type2_go = type2_go[3:]
                    return f"ns.Or[ns.{type1_go}, ns.{type2_go}]"
        return f"ns.Unknown // FIXME: Or type {combined}"
    prefixed_optional_match = re.match(r"^PrefixedOptional(.+)$", base_type)
    if prefixed_optional_match:
        inner_type = prefixed_optional_match.group(1)
        if inner_type == "ByteArray":
            return "ns.PrefixedOptional[ns.ByteArray]"
        if inner_type == "PrefixedByteArray":
            return "ns.PrefixedOptional[ns.PrefixedByteArray]"
        inner_go_type = transform_to_go_type(inner_type)
        if inner_go_type == "ns.ByteArray":
            return "ns.PrefixedOptional[ns.ByteArray]"
        if inner_go_type == "ns.PrefixedByteArray":
            return "ns.PrefixedOptional[ns.PrefixedByteArray]"
        if inner_go_type.startswith("ns."):
            inner_go_type = inner_go_type[3:]
        return f"ns.PrefixedOptional[ns.{inner_go_type}]"
    optional_match = re.match(r"^Optional(.+)$", base_type)
    if optional_match:
        inner_type = optional_match.group(1)
        if inner_type == "ByteArray" or inner_type == "PrefixedByteArray":
            return "ns.Optional[ns.ByteArray]"
        inner_go_type = transform_to_go_type(inner_type)
        if inner_go_type == "ns.ByteArray":
            return "ns.Optional[ns.ByteArray]"
        if inner_go_type.startswith("ns."):
            inner_go_type = inner_go_type[3:]
        return f"ns.Optional[ns.{inner_go_type}]"
    prefixed_array_size_match = re.match(r"^Prefixed(.+)Array\(\d+\)$", base_type)
    if prefixed_array_size_match:
        inner_type = prefixed_array_size_match.group(1)
        if inner_type == "Byte":
            return "ns.PrefixedByteArray"
        inner_go_type = transform_to_go_type(inner_type)
        if inner_go_type.startswith("ns."):
            inner_go_type = inner_go_type[3:]
        return f"ns.PrefixedArray[ns.{inner_go_type}]"
    array_size_match = re.match(r"^(.+)Array\(\d+\)$", base_type)
    if array_size_match:
        inner_type = array_size_match.group(1)
        if inner_type == "Byte":
            return "ns.ByteArray"
        inner_go_type = transform_to_go_type(inner_type)
        if inner_go_type.startswith("ns."):
            inner_go_type = inner_go_type[3:]
        return f"ns.Array[ns.{inner_go_type}]"
    prefixed_array_match = re.match(r"^Prefixed(.+)Array$", base_type)
    if prefixed_array_match:
        inner_type = prefixed_array_match.group(1)
        if inner_type == "Byte":
            return "ns.PrefixedByteArray"
        inner_go_type = transform_to_go_type(inner_type)
        if inner_go_type.startswith("ns."):
            inner_go_type = inner_go_type[3:]
        return f"ns.PrefixedArray[ns.{inner_go_type}]"
    array_match = re.match(r"^(.+)Array$", base_type)
    if array_match:
        inner_type = array_match.group(1)
        if inner_type == "Byte":
            return "ns.ByteArray"
        inner_go_type = transform_to_go_type(inner_type)
        if inner_go_type.startswith("ns."):
            inner_go_type = inner_go_type[3:]
        return f"ns.Array[ns.{inner_go_type}]"
    if not base_type or not base_type.replace("_", "").replace("-", "").isalnum():
        return f"ns.Unknown // FIXME: Invalid type '{base_type}'"

    return f"ns.{base_type}"


def import_packets_wiki() -> dict:
    """
    Returns a dictionary of all scraped packets, with the following structure:

    ```json
    {
        "configuration": {
            "clientbound": [],
            "serverbound": []
        },
        "handshaking": {
            "clientbound": [], // there are no clientbound packets in the Handshaking state
            "serverbound": [
                {
                    "id": "0x00",
                    "resource": "some_registry_id",
                    "fields": [
                        {
                            "name": "Some Field Name",
                            "type": "VarInt",
                            "notes": "Some bla bla notes text"
                        }
                    ]
                },
                {}
            ]
        },
        "login": {
            "clientbound": [],
            "serverbound": []
        },
        "play": {
            "clientbound": [],
            "serverbound": []
        },
        "status": {
            "clientbound": [],
            "serverbound": []
        }
    }
    ```
    """
    packets = {
        "configuration": {"clientbound": [], "serverbound": []},
        "handshaking": {"clientbound": [], "serverbound": []},
        "login": {"clientbound": [], "serverbound": []},
        "play": {"clientbound": [], "serverbound": []},
        "status": {"clientbound": [], "serverbound": []},
    }

    current_state = "handshaking"
    current_bound_to = "client"
    soup = BeautifulSoup(IMPORT_FILE.read_text(), "html.parser")

    start_from_element = soup.select_one(START_IMPORT_FROM_ELEMENT)
    if not start_from_element:
        raise ValueError(f"Start from element {START_IMPORT_FROM_ELEMENT} not found")

    current_packet_name = None

    for element in start_from_element.next_siblings:
        if not isinstance(element, Tag):
            continue

        if element.name == "h2":  # state
            state = map_state_from_h2(element)
            if state:
                current_state = state
                current_bound_to = "client"
                current_packet_name = None
            continue

        if element.name == "h3":  # bound
            bound = map_bound_from_h3(element)
            if bound:
                current_bound_to = "client" if bound == "clientbound" else "server"
                current_packet_name = None
            continue

        if element.name == "h4":
            current_packet_name = element.get_text(strip=True)
            continue

        if element.name == "table":
            if not current_packet_name or should_filter_tag(current_packet_name):
                current_packet_name = None
                continue

            tbody = element.find("tbody") or element
            rows = []
            if isinstance(tbody, Tag):
                rows = [r for r in tbody.find_all("tr") if isinstance(r, Tag)]

            if not rows:
                current_packet_name = None
                continue
            header_row = rows[0]
            header_cells = (
                [c for c in header_row.find_all(["td", "th"]) if isinstance(c, Tag)]
                if isinstance(header_row, Tag)
                else []
            )
            header_texts = [c.get_text(" ", strip=True).lower() for c in header_cells]
            if not (
                any("packet id" in t for t in header_texts)
                and any("state" in t for t in header_texts)
                and any("bound to" in t for t in header_texts)
            ):
                current_packet_name = None
                continue

            packet_id = ""
            packet_resource = ""
            fields = []

            inferred_bound_key = None

            table_matches_state = False
            # get index of state column from header
            state_col_index = None
            for idx, header_text in enumerate(header_texts):
                if "state" in header_text:
                    state_col_index = idx
                    break

            for probe_row in rows[1:]:
                if not isinstance(probe_row, Tag):
                    continue
                probe_cells = [
                    c for c in probe_row.find_all("td") if isinstance(c, Tag)
                ]
                # skip header-like rows
                if any(
                    (c.get_text() if hasattr(c, "get_text") else str(c)).find(
                        "Packet ID"
                    )
                    != -1
                    for c in probe_cells
                ):
                    continue
                # if we have a recognized State column and enough cells, check it.
                if state_col_index is not None and len(probe_cells) > state_col_index:
                    state_text_probe = (
                        probe_cells[state_col_index].get_text(" ", strip=True).lower()
                        if hasattr(probe_cells[state_col_index], "get_text")
                        else str(probe_cells[state_col_index]).lower()
                    )
                    if current_state in state_text_probe:
                        table_matches_state = True
                        break
                else:
                    # fallback: look for the state name anywhere in the row's text
                    row_text = probe_row.get_text(" ", strip=True).lower()
                    if current_state in row_text:
                        table_matches_state = True
                        break
            if not table_matches_state:
                current_packet_name = None
                continue

            rowspan_field = None  # track field with rowspan
            rowspan_count = 0  # track remaining rows for rowspan
            rowspan_elements = []  # track element fields for complex arrays

            for row in rows[1:]:  # skip header row
                if not isinstance(row, Tag):
                    continue
                cells = [c for c in row.find_all(["td", "th"]) if isinstance(c, Tag)]
                if not cells:
                    continue

                header_like = any(
                    (c.get_text() if hasattr(c, "get_text") else str(c)).find(
                        "Packet ID"
                    )
                    != -1
                    for c in cells
                )
                if header_like:
                    continue

                # first data row with packet metadata
                # some tables (e.g., "Status Request") have fewer than 6 cells due to colspan
                if not packet_id and len(cells) >= 1:
                    if isinstance(cells[0], Tag):
                        codes = [
                            x for x in cells[0].find_all("code") if isinstance(x, Tag)
                        ]
                    else:
                        codes = []
                    if len(codes) >= 1:
                        packet_id = codes[0].get_text(strip=True)
                    if len(codes) >= 2:
                        packet_resource = codes[1].get_text(strip=True)

                    if inferred_bound_key is None and len(cells) >= 3:
                        bound_text = (
                            cells[2].get_text(" ", strip=True).lower()
                            if hasattr(cells[2], "get_text")
                            else str(cells[2]).lower()
                        )
                        if "client" in bound_text:
                            inferred_bound_key = "clientbound"
                        elif "server" in bound_text:
                            inferred_bound_key = "serverbound"

                    # also check if this row has field data
                    if len(cells) >= 6:
                        field_name = (
                            cells[3].get_text(" ", strip=True)
                            if hasattr(cells[3], "get_text")
                            else ""
                        )
                        field_type = (
                            cells[4].get_text(" ", strip=True)
                            if hasattr(cells[4], "get_text")
                            else ""
                        )
                        field_notes = (
                            cells[5].get_text(" ", strip=True)
                            if hasattr(cells[5], "get_text")
                            else ""
                        )
                        field_notes = normalize_notes(field_notes)
                        has_rowspan = (
                            cells[3].get("rowspan")
                            if hasattr(cells[3], "get")
                            else None
                        )
                        is_array_field = "array" in field_notes.lower()

                        if has_rowspan and is_array_field:
                            try:
                                rowspan_count = (
                                    int(str(has_rowspan)) - 1
                                )  # -1 because this row counts as one
                            except (ValueError, TypeError):
                                rowspan_count = 0
                            rowspan_field = {
                                "name": field_name,
                                "type": field_type,  # this might be the first element type
                                "notes": field_notes,
                                "is_array": True,
                            }
                            rowspan_elements = []
                            if len(cells) >= 6:
                                elem_name = (
                                    cells[4].get_text(" ", strip=True)
                                    if hasattr(cells[4], "get_text") and len(cells) > 4
                                    else ""
                                )
                                elem_type = (
                                    cells[6].get_text(" ", strip=True)
                                    if hasattr(cells[6], "get_text") and len(cells) > 6
                                    else (
                                        cells[5].get_text(" ", strip=True)
                                        if hasattr(cells[5], "get_text")
                                        and len(cells) > 5
                                        else ""
                                    )
                                )
                                if elem_name and elem_type:
                                    go_field_name = "".join(
                                        word.capitalize() for word in elem_name.split()
                                    )
                                    go_field_type = transform_to_go_type(
                                        normalize_field_type(elem_type)
                                    )
                                    rowspan_elements.append(
                                        {"name": go_field_name, "type": go_field_type}
                                    )
                        elif field_name and field_type:
                            normalized_type = normalize_field_type(field_type)
                            go_type = transform_to_go_type(normalized_type)
                            fields.append(
                                {
                                    "name": field_name,
                                    "type": go_type,
                                    "notes": field_notes,
                                }
                            )

                # subsequent rows (could be array elements if we have a rowspan field)
                elif len(cells) == 3 and packet_id:
                    # 1. Array element rows (if rowspan_field is active)
                    # 2. Regular field rows (3 cells: name, type, notes)

                    if rowspan_field and rowspan_count > 0:
                        elem_name = (
                            cells[0].get_text(" ", strip=True)
                            if hasattr(cells[0], "get_text")
                            else ""
                        )
                        elem_type = (
                            cells[1].get_text(" ", strip=True)
                            if hasattr(cells[1], "get_text")
                            else ""
                        )
                        if elem_name and elem_type:
                            go_field_name = "".join(
                                word.capitalize() for word in elem_name.split()
                            )
                            go_field_type = transform_to_go_type(
                                normalize_field_type(elem_type)
                            )
                            rowspan_elements.append(
                                {"name": go_field_name, "type": go_field_type}
                            )
                        rowspan_count -= 1
                        if rowspan_count == 0:
                            if len(rowspan_elements) > 1:
                                struct_fields = []
                                for elem in rowspan_elements:
                                    struct_fields.append(
                                        f"{elem['name']} {elem['type']}"
                                    )
                                struct_def = (
                                    "struct { " + "; ".join(struct_fields) + " }"
                                )

                                if "prefixed array" in rowspan_field["notes"].lower():
                                    final_type = f"ns.PrefixedArray[{struct_def}]"
                                else:
                                    final_type = f"ns.Array[{struct_def}]"
                            elif len(rowspan_elements) == 1:
                                elem_type = rowspan_elements[0]["type"]
                                if elem_type == "ns.Byte":
                                    if "prefixed array" in rowspan_field["notes"].lower():
                                        final_type = "ns.PrefixedByteArray"
                                    else:
                                        final_type = "ns.ByteArray"
                                elif "prefixed array" in rowspan_field["notes"].lower():
                                    final_type = f"ns.PrefixedArray[{elem_type}]"
                                else:
                                    final_type = f"ns.Array[{elem_type}]"
                            else:
                                if "prefixed array" in rowspan_field["notes"].lower():
                                    final_type = "ns.PrefixedArray[ns.Unknown]"
                                else:
                                    final_type = "ns.Array[ns.Unknown]"

                            fields.append(
                                {
                                    "name": rowspan_field["name"],
                                    "type": final_type,
                                    "notes": rowspan_field["notes"],
                                }
                            )
                            rowspan_field = None
                            rowspan_elements = []
                        continue
                    else:
                        field_name = (
                            cells[0].get_text(" ", strip=True)
                            if hasattr(cells[0], "get_text")
                            else ""
                        )
                        field_type = (
                            cells[1].get_text(" ", strip=True)
                            if hasattr(cells[1], "get_text")
                            else ""
                        )
                        field_notes = (
                            cells[2].get_text(" ", strip=True)
                            if hasattr(cells[2], "get_text")
                            else ""
                        )
                        field_notes = normalize_notes(field_notes)

                        if field_name and field_type:
                            normalized_type = normalize_field_type(field_type)
                            go_type = transform_to_go_type(normalized_type)
                            fields.append(
                                {
                                    "name": field_name,
                                    "type": go_type,
                                    "notes": field_notes,
                                }
                            )

            # collect any surrounding paragraph notes for this packet table
            packet_notes = extract_surrounding_paragraphs_text(element)
            packet_notes = normalize_notes(packet_notes)

            packet_obj = {
                "name": (
                    current_packet_name.replace("[edit|edit source]", "").strip()
                    if current_packet_name
                    else ""
                ),
                "id": packet_id,
                "resource": packet_resource,
                "notes": packet_notes,
                "fields": fields,
            }

            bound_key = "clientbound" if current_bound_to == "client" else "serverbound"
            packets[current_state][bound_key].append(packet_obj)

            current_packet_name = None

    return packets


def normalize_header_text(text: str) -> str:
    return (text or "").strip().lower()


def map_state_from_h2(tag: Tag) -> str | None:
    span = tag.find("span", id=True)
    if not isinstance(span, Tag):
        return None

    if span.get("id"):
        state = normalize_header_text(str(span["id"]))
    else:
        state = normalize_header_text(tag.get_text())

    if "configuration" in state:
        return "configuration"
    if "handshak" in state:
        return "handshaking"
    if "status" in state:
        return "status"
    if "login" in state:
        return "login"
    if "play" in state:
        return "play"
    return None


def map_bound_from_h3(tag: Tag) -> str | None:
    span = tag.find("span", id=True)
    if isinstance(span, Tag) and span.get("id"):
        text = normalize_header_text(str(span["id"]))
    else:
        text = normalize_header_text(tag.get_text())

    if "clientbound" in text:
        return "clientbound"
    if "serverbound" in text:
        return "serverbound"

    return None


def should_filter_tag(name: str) -> bool:
    return normalize_header_text(name).startswith("legacy server list ping")


def extract_surrounding_paragraphs_text(table: Tag) -> str:
    if not isinstance(table, Tag):
        return ""

    before_texts = []
    for prev in table.previous_siblings:
        if not isinstance(prev, Tag):
            continue
        if prev.name in ("h2", "h3", "h4", "table"):
            break
        if prev.name == "p":
            before_texts.append(prev.get_text(" ", strip=True))
            continue
        # stop if we encounter another significant tag before hitting a header/table
        break
    before_texts.reverse()

    after_texts = []
    for nxt in table.next_siblings:
        if not isinstance(nxt, Tag):
            continue
        if nxt.name in ("h2", "h3", "h4", "table"):
            break
        if nxt.name == "p":
            after_texts.append(nxt.get_text(" ", strip=True))
            continue
        # stop if we encounter another significant tag before hitting a header/table
        break

    combined = [t for t in before_texts + after_texts if t]
    return "\n\n".join(combined)


if __name__ == "__main__":
    import json

    packets = import_packets_wiki()
    print(json.dumps(packets, indent=2))
