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


def normalize_field_type(raw_type: str) -> str:
    if not raw_type:
        return ""

    t = re.sub(r"\s+", " ", raw_type).strip()

    # drop trailing/standalone 'Enum' token(s)
    t = re.sub(r"\b[Ee]nums?\b", "", t).strip()
    t = re.sub(r"\s{2,}", " ", t).strip()

    # Prefixed Array (N) of Type -> PrefixedTypeArray(N)
    def _prefixed_array_repl(match: re.Match) -> str:
        size = match.group(1)
        elem_type = match.group(2)
        return f"Prefixed{elem_type}Array({size})"

    t = re.sub(r"\bPrefixed\s+Array\s*\(\s*([0-9]+)\s*\)\s+of\s+([A-Za-z0-9]+)\b", _prefixed_array_repl, t)

    # Prefixed Array of Type -> PrefixedTypeArray
    t = re.sub(r"\bPrefixed\s+Array\s+of\s+([A-Za-z0-9]+)\b", r"Prefixed\1Array", t)

    # Array (N) of Type -> TypeArray(N)
    def _array_size_of_repl(match: re.Match) -> str:
        size = match.group(1)
        elem_type = match.group(2)
        return f"{elem_type}Array({size})"

    t = re.sub(r"\bArray\s*\(\s*([0-9]+)\s*\)\s+of\s+([A-Za-z0-9]+)\b", _array_size_of_repl, t)

    # Array of Type -> TypeArray
    t = re.sub(r"\bArray\s+of\s+([A-Za-z0-9]+)\b", r"\1Array", t)

    # remove all spaces
    t = re.sub(r"\s+", "", t)

    return t

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
                        field_type = normalize_field_type(field_type)
                        field_notes = (
                            cells[5].get_text(" ", strip=True)
                            if hasattr(cells[5], "get_text")
                            else ""
                        )
                        field_notes = normalize_notes(field_notes)
                        if (
                            field_name and field_type
                        ):  # only add if we have actual field data
                            fields.append(
                                {
                                    "name": field_name,
                                    "type": field_type,
                                    "notes": field_notes,
                                }
                            )

                # subsequent rows with just field data (3 cells)
                elif len(cells) == 3 and packet_id:
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
                    field_type = normalize_field_type(field_type)
                    field_notes = (
                        cells[2].get_text(" ", strip=True)
                        if hasattr(cells[2], "get_text")
                        else ""
                    )
                    field_notes = normalize_notes(field_notes)
                    if field_name and field_type:
                        fields.append(
                            {
                                "name": field_name,
                                "type": field_type,
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
