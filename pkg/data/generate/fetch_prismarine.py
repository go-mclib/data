#!/usr/bin/env python3

"""
Fetches blockCollisionShapes.json from PrismarineJS/minecraft-data.

Usage: ./fetch_prismarine.py [version]
Default version is the highest pc version that has blockCollisionShapes.
"""

import json
import sys
import urllib.request
from pathlib import Path

BASE_URL = "https://raw.githubusercontent.com/PrismarineJS/minecraft-data/master/data"
OUT = Path(__file__).parent / "prismarine_block_collision_shapes.json"


def resolve_version() -> str:
    print("No version specified, resolving from dataPaths.json...")
    with urllib.request.urlopen(f"{BASE_URL}/dataPaths.json") as resp:
        data_paths = json.load(resp)["pc"]

    versions = [
        v["blockCollisionShapes"].split("/")[-1]
        for v in data_paths.values()
        if "blockCollisionShapes" in v
    ]
    versions.sort(key=lambda x: [int(p) for p in x.split(".")])
    return versions[-1]


def main():
    version = sys.argv[1] if len(sys.argv) > 1 else resolve_version()
    print(f"Resolved to version: {version}")

    url = f"{BASE_URL}/pc/{version}/blockCollisionShapes.json"
    print(f"Fetching {url}...")
    urllib.request.urlretrieve(url, OUT)
    print(f"Wrote {OUT.name}")


if __name__ == "__main__":
    main()
