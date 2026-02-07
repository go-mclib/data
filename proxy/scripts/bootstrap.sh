#!/bin/bash
# Resets the test world from the template for reproducible packet captures.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROXY_DIR="$SCRIPT_DIR/.."
SERVER_DIR="$PROXY_DIR/server"
TEMPLATE_DIR="$PROXY_DIR/testdata/world-template"

if [ ! -d "$TEMPLATE_DIR" ]; then
    echo "Error: world template not found at $TEMPLATE_DIR"
    echo "Create the template first (see proxy/testdata/world-setup.md)"
    exit 1
fi

echo "Resetting test world from template..."

rm -rf "$SERVER_DIR/world"
mkdir -p "$SERVER_DIR/world/region"
cp "$TEMPLATE_DIR/level.dat" "$SERVER_DIR/world/"
cp "$TEMPLATE_DIR/region/"*.mca "$SERVER_DIR/world/region/"

echo "World reset."
echo ""
echo "Next steps:"
echo "  1. Start server:  cd $SERVER_DIR && java -jar server.jar nogui"
echo "  2. Start proxy:   cd $PROXY_DIR && go run ./cmd/proxy -target localhost:25566"
echo "  3. Connect client to localhost:25565 as GoMclib"
echo "  4. Perform actions (see proxy/testdata/world-setup.md)"
echo "  5. Ctrl+C the proxy to save capture"
echo "  6. Rename capture:  mv captures/session_*.json captures/<name>.json"
echo "  7. Regenerate:  cd ../../pkg/packets_test && go run generate.go"
