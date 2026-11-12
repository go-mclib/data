#!/usr/bin/env bash
# download the vanilla server jar and run its data generator to produce the
# JSON reports (blocks, items, registries, packets, commands) and the jar's
# version.json (protocol/data version) used by the Go bindings generator in
# ../generate.
#
# usage: ./generate.sh [version]   (default: latest known, see DEFAULT_VERSION)
set -euo pipefail
cd "$(dirname "$0")"

DEFAULT_VERSION="26.3"
VERSION="${1:-$DEFAULT_VERSION}"
JAR="server-$VERSION.jar"

download_server_jar() {
    [[ -f "$JAR" ]] && { echo "server jar already present: $JAR"; return 0; }

    echo "fetching version manifest..."
    local manifest version_url meta url sha1 actual
    manifest=$(curl -sf "https://piston-meta.mojang.com/mc/game/version_manifest_v2.json")
    version_url=$(python3 -c "
import json, sys
for v in json.loads(sys.stdin.read())['versions']:
    if v['id'] == sys.argv[1]:
        print(v['url']); sys.exit()
sys.exit('version not found: ' + sys.argv[1])
" "$VERSION" <<< "$manifest")

    meta=$(curl -sf "$version_url")
    url=$(python3 -c "import json,sys; print(json.load(sys.stdin)['downloads']['server']['url'])" <<< "$meta")
    sha1=$(python3 -c "import json,sys; print(json.load(sys.stdin)['downloads']['server']['sha1'])" <<< "$meta")

    echo "downloading server jar ($VERSION)..."
    curl -# -o "$JAR" "$url"
    actual=$(shasum -a 1 "$JAR" | cut -d' ' -f1)
    [[ "$actual" == "$sha1" ]] || { echo "sha1 mismatch (expected $sha1, got $actual)"; rm -f "$JAR"; return 1; }
    echo "downloaded and verified: $JAR"
}

download_server_jar

echo "running data generator..."
java -DbundlerMainClass="net.minecraft.data.Main" -jar "$JAR" --server --reports
unzip -p "$JAR" version.json > generated/version.json

# symlink reports into the Go bindings generator's working dir
ln -sfn "$PWD/generated/reports/commands.json"   ../generate/commands.json
ln -sfn "$PWD/generated/reports/blocks.json"     ../generate/blocks.json
ln -sfn "$PWD/generated/reports/registries.json" ../generate/registries.json
ln -sfn "$PWD/generated/reports/packets.json"    ../generate/packets.json
ln -sfn "$PWD/generated/reports/minecraft/components/item" ../generate/items
ln -sfn "$PWD/generated/version.json"            ../generate/version.json

echo "reports generated; now run 'go generate ./...' in .."
