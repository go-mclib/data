#!/usr/bin/env bash
# download, set up, and run a Minecraft server for gomclib testing
#
# usage:
#   server.sh download <version>          fetch server jar from Mojang
#     V
#    server.sh setup <version>             ensure server is ready from template in /tmp
#      V
#     server.sh run <version>               setup + start the server
#
#   server.sh update-template <version>   boot server, let you edit the world,
#                                         then sync whitelisted files back to
#                                         server_template/ on /stop or ctrl+c
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

# files copied back verbatim from the running server dir to the template
TEMPLATE_FILES=(
    codeofconduct/en_us.txt
    eula.txt
    ops.json
    server.properties
    world/level.dat
    world/data/minecraft/game_rules.dat
    world/data/minecraft/world_gen_settings.dat
    world/dimensions/minecraft/overworld/data/minecraft/world_border.dat
)

# directories mirrored recursively (entire contents replaced in the template)
TEMPLATE_DIRS=(
    world/dimensions/minecraft/overworld/entities
    world/dimensions/minecraft/overworld/poi
    world/dimensions/minecraft/overworld/region
)

sync_to_template() {
    echo ""
    echo "syncing changes back to $TEMPLATE_DIR..."

    for f in "${TEMPLATE_FILES[@]}"; do
        if [[ -f "$SERVER_DIR/$f" ]]; then
            mkdir -p "$TEMPLATE_DIR/$(dirname "$f")"
            cp "$SERVER_DIR/$f" "$TEMPLATE_DIR/$f"
            echo "  file  $f"
        else
            echo "  skip  $f (missing)"
        fi
    done

    for d in "${TEMPLATE_DIRS[@]}"; do
        if [[ -d "$SERVER_DIR/$d" ]]; then
            rm -rf "$TEMPLATE_DIR/$d"
            mkdir -p "$TEMPLATE_DIR/$d"
            # trailing /. copies contents without the leading dir
            cp -R "$SERVER_DIR/$d/." "$TEMPLATE_DIR/$d/"
            echo "  dir   $d/"
        else
            echo "  skip  $d/ (missing)"
        fi
    done

    # revert proxy-specific tweaks applied in setup() so the template stays
    # canonical (matches a fresh vanilla server's defaults).
    if [[ -f "$TEMPLATE_DIR/server.properties" ]]; then
        sed -i '' "s/^server-port=.*/server-port=25565/" "$TEMPLATE_DIR/server.properties"
        sed -i '' "s/^pause-when-empty-seconds=.*/pause-when-empty-seconds=60/" "$TEMPLATE_DIR/server.properties"
    fi

    echo "template updated."
}

update_template() {
    local version="$1"
    setup "$version"

    local fifo="$SERVER_DIR/stdin_pipe"
    rm -f "$fifo"
    mkfifo "$fifo"

    cd "$SERVER_DIR"
    # background the server so the shell can trap ctrl+c and forward /stop.
    # without a FIFO on stdin, a backgrounded java would see EOF and exit.
    java -jar server.jar -nogui < "$fifo" &
    local server_pid=$!
    exec 4>"$fifo"  # keep write end open so server doesn't read EOF

    synced=0
    on_exit() {
        trap - EXIT INT TERM
        if [[ $synced -eq 0 ]]; then
            synced=1
            echo ""
            echo "stopping server..."
            echo "stop" >&4 2>/dev/null || true
            # wait up to 60s for clean shutdown so chunks flush to disk
            for _ in $(seq 1 600); do
                kill -0 "$server_pid" 2>/dev/null || break
                sleep 0.1
            done
            if kill -0 "$server_pid" 2>/dev/null; then
                echo "server did not exit in time, sending SIGTERM"
                kill "$server_pid" 2>/dev/null || true
                wait "$server_pid" 2>/dev/null || true
            else
                wait "$server_pid" 2>/dev/null || true
            fi
            exec 4>&- 2>/dev/null || true
            sync_to_template
        fi
    }
    trap on_exit EXIT INT TERM

    echo ""
    echo "=== template update mode ==="
    echo "connect to localhost:$SERVER_PORT, modify the world as needed."
    echo "type /stop (in-game or in this terminal) or ctrl+c to save and sync."
    echo ""

    wait "$server_pid" 2>/dev/null || true
}

case "${1:-help}" in
    download)        download "${2:?usage: $0 download <version>}" ;;
    setup)           setup "${2:?usage: $0 setup <version>}" ;;
    run)             run "${2:?usage: $0 run <version>}" ;;
    update-template) update_template "${2:?usage: $0 update-template <version>}" ;;
    *)               echo "usage: $0 {download|setup|run|update-template} <version>"; exit 1 ;;
esac
