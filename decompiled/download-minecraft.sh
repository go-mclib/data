#!/bin/bash
set -e
cd "$(dirname "$0")"

VERSION="${1:-1.21.11}"
PRISM_META="https://meta.prismlauncher.org/v1"

echo "Downloading Minecraft $VERSION..."

# fetch version index from PrismLauncher meta (includes unobfuscated versions)
INDEX=$(curl -s "$PRISM_META/net.minecraft/index.json")

# check if unobfuscated version exists
UNOBF_VERSION="${VERSION}_unobfuscated"
HAS_UNOBF=$(echo "$INDEX" | jq -r --arg v "$UNOBF_VERSION" '.versions[] | select(.version == $v) | .version')

if [ -n "$HAS_UNOBF" ] && [ "$HAS_UNOBF" != "null" ]; then
    echo "Found unobfuscated version: $UNOBF_VERSION"
    TARGET_VERSION="$UNOBF_VERSION"
else
    echo "No unobfuscated version available, using regular version"
    TARGET_VERSION="$VERSION"
fi

# verify target version exists
VERSION_EXISTS=$(echo "$INDEX" | jq -r --arg v "$TARGET_VERSION" '.versions[] | select(.version == $v) | .version')
if [ -z "$VERSION_EXISTS" ] || [ "$VERSION_EXISTS" = "null" ]; then
    echo "Error: Version $TARGET_VERSION not found"
    echo "Available versions (latest 20):"
    echo "$INDEX" | jq -r '.versions[:20] | .[].version'
    exit 1
fi

# fetch version metadata
echo "Fetching version metadata for $TARGET_VERSION..."
VERSION_META=$(curl -s "$PRISM_META/net.minecraft/$TARGET_VERSION.json")

# extract client jar URL
CLIENT_URL=$(echo "$VERSION_META" | jq -r '.mainJar.downloads.artifact.url')

if [ -z "$CLIENT_URL" ] || [ "$CLIENT_URL" = "null" ]; then
    echo "Error: Could not find client download URL"
    exit 1
fi

# download client jar
echo "Downloading client.jar from $CLIENT_URL..."
curl -o "minecraft.jar" "$CLIENT_URL"

echo "Done! Downloaded Minecraft $TARGET_VERSION to minecraft.jar"
