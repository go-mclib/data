from json import load
from pathlib import Path


ROOT = Path(__file__).parent
PROTOCOL_VERSION = "772"
JSON_FILE = ROOT.parent / "data" / PROTOCOL_VERSION / "packets.json"


def load_packets_json() -> dict:
    with JSON_FILE.open("r") as f:
        return load(f)


def generate_packets_go() -> None:
    packets = load_packets_json()


if __name__ == "__main__":
    generate_packets_go()
