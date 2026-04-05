#!/usr/bin/env bash
# orchestrate a full packet capture session:
#   1. set up + start minecraft server from template
#   2. start MITM proxy
#   3. wait for client connection (manual or via macro)
#   4. save captured packets on exit
#
# usage:
#   capture.sh <version> [macro.json]
#
# examples:
#   capture.sh 26.1.1                          # manual client connection
#   capture.sh 26.1.1 ../macros/full_test.json  # automated via macro
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROXY_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
DATA_DIR="$(cd "$PROXY_DIR/.." && pwd)"
SERVER_DIR="/tmp/gomclib-test-server"

VERSION="${1:?usage: $0 <version> [macro.json]}"
MACRO="${2:-}"
SERVER_PORT=25566
PROXY_PORT=25565

SERVER_PID=""
PROXY_PID=""

cleanup() {
    echo ""
    echo "shutting down..."
    if [[ -n "$PROXY_PID" ]]; then
        kill "$PROXY_PID" 2>/dev/null && wait "$PROXY_PID" 2>/dev/null
        echo "  proxy stopped"
    fi
    if [[ -n "$SERVER_PID" ]]; then
        kill "$SERVER_PID" 2>/dev/null && wait "$SERVER_PID" 2>/dev/null
        echo "  server stopped"
    fi
    echo "done. captures saved to $PROXY_DIR/captures/"
}
trap cleanup EXIT INT TERM

echo "=== packet capture ==="
echo "version=$VERSION  server=:$SERVER_PORT  proxy=:$PROXY_PORT"
echo ""

# step 1: set up server
echo "[1/4] setting up server..."
"$SCRIPT_DIR/server.sh" setup "$VERSION"
echo ""

# step 2: start server and wait for it to be ready
echo "[2/4] starting server..."
cd "$SERVER_DIR"
java -jar server.jar -nogui > "$SERVER_DIR/console.log" 2>&1 &
SERVER_PID=$!
echo "server PID: $SERVER_PID"

echo -n "waiting for server"
for i in $(seq 1 120); do
    if ! kill -0 "$SERVER_PID" 2>/dev/null; then
        echo ""
        echo "error: server exited unexpectedly. last 20 lines:"
        tail -20 "$SERVER_DIR/console.log"
        exit 1
    fi
    if grep -q "Done" "$SERVER_DIR/console.log" 2>/dev/null; then
        echo " ready!"
        break
    fi
    if [[ $i -eq 120 ]]; then
        echo ""
        echo "error: server did not start within 120s"
        tail -20 "$SERVER_DIR/console.log"
        exit 1
    fi
    echo -n "."
    sleep 1
done
echo ""

# step 3: start proxy
echo "[3/4] starting proxy..."
cd "$DATA_DIR"
go run ./proxy/cmd/proxy -target "localhost:$SERVER_PORT" -port "$PROXY_PORT" -verbose &
PROXY_PID=$!
sleep 2
echo "proxy PID: $PROXY_PID"
echo ""

# step 4: client connection
echo "[4/4] connect to localhost:$PROXY_PORT"
echo ""
echo "  tip: add --quickPlayMultiplayer localhost:$PROXY_PORT"
echo "       to your launcher profile's game arguments for auto-connect"
echo ""

if [[ -n "$MACRO" ]]; then
    echo "macro: $MACRO (starting in 5s...)"
    sleep 5
    python3 "$SCRIPT_DIR/input_macro.py" replay "$MACRO"
    echo ""
    echo "macro finished. press ctrl+c to save captures and exit."
    wait "$PROXY_PID" 2>/dev/null || true
else
    echo "press ctrl+c when done capturing."
    wait "$PROXY_PID" 2>/dev/null || true
fi
