#!/usr/bin/env bash
# Parse the decompiled Minecraft source (mcsrc) and emit a packet IR JSON
# describing every packet's wire fields. Stage 1 of packet-binding generation;
# the Go generator (pkg/data/generate/gen_packets.go) consumes the IR.
#
# Unlike mcdump this needs no server jar — it parses .java source with JavaParser,
# so it only needs the JavaParser jar (fetched once) + gson (already vendored).
#
# usage: ./gen.sh [mcsrc-dir]   (default: ../../mcsrc/current)
set -euo pipefail
cd "$(dirname "$0")"

MCSRC="${1:-../../mcsrc/current}"
OUT="../pkg/data/generate/packets_ir.json"
WORK=".work"
LIB=".lib"

JP_VERSION="3.26.4"
JP_JAR="$LIB/javaparser-core-${JP_VERSION}.jar"
JP_URL="https://repo1.maven.org/maven2/com/github/javaparser/javaparser-core/${JP_VERSION}/javaparser-core-${JP_VERSION}.jar"
GSON_JAR=$(ls ../vanilla_server_reports/libraries/com/google/code/gson/gson/*/gson-*.jar 2>/dev/null | sort -V | tail -1)

if [[ ! -d "$MCSRC/net/minecraft/network/protocol" ]]; then
	echo "mcsrc protocol dir not found under: $MCSRC"
	exit 1
fi
if [[ -z "${GSON_JAR:-}" ]]; then
	echo "gson jar not found; run ../vanilla_server_reports/generate.sh first"
	exit 1
fi

mkdir -p "$LIB"
if [[ ! -f "$JP_JAR" ]]; then
	echo "fetching JavaParser ${JP_VERSION}..."
	curl -fsSL "$JP_URL" -o "$JP_JAR"
fi

CP="$JP_JAR:$GSON_JAR"
echo "compiling PacketDump.java..."
rm -rf "$WORK"; mkdir -p "$WORK"
javac -cp "$CP" -d "$WORK" PacketDump.java

echo "parsing packets under $MCSRC..."
java -cp "$WORK:$CP" PacketDump "$MCSRC" "$OUT"

echo "IR written to $OUT"
