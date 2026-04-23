#!/usr/bin/env bash
# orchestrate a full packet capture session:
#   1. set up + start minecraft server from template
#   2. start MITM proxy (paused — captures only after macro starts)
#   3. record or replay a macro (triggered by in-game command block)
#   4. auto-shutdown and save captures when macro finishes
#
# usage:
#   capture.sh <version> record <output.json>    # record new macro
#   capture.sh <version> replay <macro.json>     # replay existing macro
#   capture.sh <version>                         # manual (no macro)
#
# the macro is triggered by a command block running /say [macro] start
# and stopped by /say [macro] stop (record only).
set -euo pipefail

PROXY_DIR="$(cd "$(dirname "$0")" && pwd)"
DATA_DIR="$(cd "$PROXY_DIR/.." && pwd)"
SERVER_DIR="/tmp/gomclib-test-server"

VERSION="${1:?usage: $0 <version> [record <out.json> | replay <macro.json>]}"
MODE="${2:-}"
# resolve macro path now, before any cd changes the working directory
MACRO_FILE="${3:-}"
if [[ -n "$MACRO_FILE" && "$MACRO_FILE" != /* ]]; then
    MACRO_FILE="$(pwd)/$MACRO_FILE"
fi
SERVER_PORT=25566
PROXY_PORT=25565

if [[ -n "$MODE" && "$MODE" != "record" && "$MODE" != "replay" ]]; then
    echo "usage: $0 <version> [record <out.json> | replay <macro.json>]"
    exit 1
fi
if [[ -n "$MODE" && -z "$MACRO_FILE" ]]; then
    echo "usage: $0 <version> $MODE <file.json>"
    exit 1
fi

# check if ports are already in use
check_port() {
    local port="$1" label="$2"
    local pid
    pid=$(lsof -ti "tcp:$port" 2>/dev/null || true)
    if [[ -n "$pid" ]]; then
        echo "error: port $port ($label) is already in use by PID $pid"
        echo "  kill it with:  kill $pid"
        return 1
    fi
}
busy=0
check_port "$SERVER_PORT" "server" || busy=1
check_port "$PROXY_PORT" "proxy"  || busy=1
if [[ $busy -ne 0 ]]; then
    exit 1
fi

SERVER_PID=""
PROXY_PID=""
MACRO_PID=""
SERVER_INPUT=""

# kill a process and all its children
killtree() {
    local pid="$1" sig="${2:-TERM}"
    pkill -"$sig" -P "$pid" 2>/dev/null  # children first
    kill -"$sig" "$pid" 2>/dev/null
}

cleanup() {
    set +e
    trap - EXIT INT TERM
    echo ""
    echo "shutting down..."
    [[ -n "$MACRO_PID" ]] && kill "$MACRO_PID" 2>/dev/null
    [[ -n "$PROXY_PID" ]] && killtree "$PROXY_PID" 9
    [[ -n "$SERVER_PID" ]] && killtree "$SERVER_PID" 9
    exec 3>&- 2>/dev/null || true
    wait 2>/dev/null
    echo "done. captures saved to $PROXY_DIR/captures/"
}
trap cleanup EXIT INT TERM

# check pynput early (before starting server/proxy) if we need it
if [[ -n "$MODE" ]]; then
    python3 -c "import pynput" 2>/dev/null || {
        echo "error: pynput is required for macros: pip install pynput"
        exit 1
    }
fi

echo "=== packet capture ==="
echo "version=$VERSION  server=:$SERVER_PORT  proxy=:$PROXY_PORT  mode=${MODE:-manual}"
echo ""

# step 1: set up server
echo "[1/4] setting up server..."
"$PROXY_DIR/server.sh" setup "$VERSION"
echo ""

# step 2: start server with stdin FIFO (so we can send commands)
echo "[2/4] starting server..."
SERVER_INPUT="$SERVER_DIR/stdin_pipe"
rm -f "$SERVER_INPUT"
mkfifo "$SERVER_INPUT"

cd "$SERVER_DIR"
# start java first (opens read end of FIFO), then open write end
# (reversed order would deadlock: write blocks until a reader exists)
java -jar server.jar -nogui < "$SERVER_INPUT" > "$SERVER_DIR/console.log" 2>&1 &
SERVER_PID=$!
exec 3>"$SERVER_INPUT"  # keep write end open so server doesn't get EOF
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
PROXY_ARGS=(-target "localhost:$SERVER_PORT" -port "$PROXY_PORT" -verbose)
if [[ -n "$MODE" ]]; then
    PROXY_ARGS+=(-paused)
fi
cd "$DATA_DIR"
go run ./proxy/cmd/proxy "${PROXY_ARGS[@]}" &
PROXY_PID=$!
sleep 2
echo "proxy PID: $PROXY_PID"
echo ""

# step 4: macro / manual
SERVER_LOG="$SERVER_DIR/console.log"
CAPTURE_TRIGGER="$PROXY_DIR/captures/.start"
MACRO_ARGS=(--server-log "$SERVER_LOG" --server-input "$SERVER_INPUT" --capture-trigger "$CAPTURE_TRIGGER")

if [[ "$MODE" == "record" ]]; then
    echo "[4/4] connect to localhost:$PROXY_PORT, then click the [macro] start button"
    echo ""
    python3 "$PROXY_DIR/input_macro.py" record "${MACRO_ARGS[@]}" "$MACRO_FILE"
    echo ""

elif [[ "$MODE" == "replay" ]]; then
    echo "[4/4] connect to localhost:$PROXY_PORT, then click the [macro] start button"
    echo ""
    python3 "$PROXY_DIR/input_macro.py" replay "${MACRO_ARGS[@]}" "$MACRO_FILE"
    echo ""

else
    echo "[4/4] connect to localhost:$PROXY_PORT"
    echo "press ctrl+c when done capturing."
    wait "$PROXY_PID" 2>/dev/null || true
fi
# cleanup trap fires on exit, stopping proxy + server
