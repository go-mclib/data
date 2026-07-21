#!/usr/bin/env bash
# build and run the standalone dumper against the vanilla server jar, writing the
# json the Go generator consumes. No Fabric, no Gradle — the server jar ships with
# real Mojang names and all its libraries, so we just put them on the classpath.
#
# usage: ./dump.sh [version]   (default: matches vanilla_server_reports)
set -euo pipefail
cd "$(dirname "$0")"

VERSION="${1:-26.2}"
SERVER_JAR="../vanilla_server_reports/server-${VERSION}.jar"
WORK=".work"
OUT="../generate/mcdump"

if [[ ! -f "$SERVER_JAR" ]]; then
	echo "server jar not found: $SERVER_JAR"
	echo "run ../vanilla_server_reports/generate.sh $VERSION first (it downloads it)"
	exit 1
fi

echo "extracting server jar + bundled libraries..."
rm -rf "$WORK"
mkdir -p "$WORK"
unzip -q "$SERVER_JAR" 'META-INF/versions/*' 'META-INF/libraries/*' -d "$WORK"

SERVER=$(ls "$WORK"/META-INF/versions/*/server-*.jar)
LIBS=$(find "$WORK/META-INF/libraries" -name '*.jar' | tr '\n' ':')
CP="$SERVER:$LIBS"

echo "compiling Dump.java..."
javac -cp "$CP" -d "$WORK" Dump.java

echo "running dumper..."
mkdir -p "$OUT"
java -cp "$WORK:$CP" Dump "$OUT"

echo "dump written to $OUT — now run 'go generate ./...' in .."
