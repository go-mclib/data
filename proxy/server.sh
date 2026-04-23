#!/usr/bin/env bash
# download, set up, and run a Minecraft server for gomclib testing
#
# usage:
#   server.sh download <version>   fetch server jar from Mojang
#     V
#    server.sh setup <version>      ensure server is ready from template in /tmp
#      V
#     server.sh run <version>        setup + start the server
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")" && pwd)"
JARS_DIR="$ROOT_DIR/jars"
TEMPLATE_DIR="$ROOT_DIR/server_template"
SERVER_DIR="/tmp/gomclib-test-server"
SERVER_PORT=25566

download() {
    local version="$1"
    local jar="$JARS_DIR/$version.jar"

    if [[ -f "$jar" ]]; then
        echo "jar already exists: $jar"
        return 0
    fi

    echo "fetching version manifest..."
    local manifest
    manifest=$(curl -sf "https://piston-meta.mojang.com/mc/game/version_manifest_v2.json")

    local version_url
    version_url=$(python3 -c "
import json, sys
for v in json.loads(sys.stdin.read())['versions']:
    if v['id'] == sys.argv[1]:
        print(v['url']); sys.exit()
print('NOT_FOUND', file=sys.stderr); sys.exit(1)
" "$version" <<< "$manifest") || { echo "error: version '$version' not found in manifest"; return 1; }

    echo "fetching version metadata..."
    local meta
    meta=$(curl -sf "$version_url")

    local url sha1
    url=$(python3 -c "import json,sys; d=json.loads(sys.stdin.read()); print(d['downloads']['server']['url'])" <<< "$meta")
    sha1=$(python3 -c "import json,sys; d=json.loads(sys.stdin.read()); print(d['downloads']['server']['sha1'])" <<< "$meta")

    echo "downloading server jar..."
    mkdir -p "$JARS_DIR"
    curl -# -o "$jar" "$url"

    local actual
    actual=$(shasum -a 1 "$jar" | cut -d' ' -f1)
    if [[ "$actual" != "$sha1" ]]; then
        echo "error: sha1 mismatch (expected $sha1, got $actual)"
        rm -f "$jar"
        return 1
    fi

    echo "downloaded and verified: $jar"
}

setup() {
    local version="$1"
    local jar="$JARS_DIR/$version.jar"

    download "$version"

    [[ -f "$jar" ]] || { echo "error: jar not found: $jar"; return 1; }

    # clean + copy template
    if [[ -d "$SERVER_DIR" ]]; then
        echo "removing old server dir..."
        rm -rf "$SERVER_DIR"
    fi
    echo "copying template -> $SERVER_DIR"
    cp -r "$TEMPLATE_DIR" "$SERVER_DIR"
    cp "$jar" "$SERVER_DIR/server.jar"

    # adjust properties for proxy setup
    sed -i '' "s/^server-port=.*/server-port=$SERVER_PORT/" "$SERVER_DIR/server.properties"
    sed -i '' "s/^pause-when-empty-seconds=.*/pause-when-empty-seconds=0/" "$SERVER_DIR/server.properties"

    echo "server prepared at $SERVER_DIR (port $SERVER_PORT)"
}

run() {
    local version="$1"
    setup "$version"
    echo "starting server on port $SERVER_PORT..."
    cd "$SERVER_DIR"
    exec java -jar server.jar -nogui
}

case "${1:-help}" in
    download) download "${2:?usage: $0 download <version>}" ;;
    setup)    setup "${2:?usage: $0 setup <version>}" ;;
    run)      run "${2:?usage: $0 run <version>}" ;;
    *)        echo "usage: $0 {download|setup|run} <version>"; exit 1 ;;
esac
