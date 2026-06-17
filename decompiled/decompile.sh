#!/bin/bash
set -e
cd "$(dirname "$0")"

MOJANG_LIBS=0
THIRD_PARTY_LIBS=0

for arg in "$@"; do
    case "$arg" in
        --mojang-libs) MOJANG_LIBS=1 ;;
        --third-party-libs) THIRD_PARTY_LIBS=1 ;;
        *) echo "Unknown flag: $arg" >&2; exit 1 ;;
    esac
done

java -jar vineflower.jar minecraft.jar ./current

ln -sf "$PWD/current/assets/minecraft/lang/en_us.json" ../pkg/data/generate/en_us.json

if [ "$MOJANG_LIBS" -eq 0 ] && [ "$THIRD_PARTY_LIBS" -eq 0 ]; then
    exit 0
fi

VERSION=$(jq -r '.id' ./current/version.json)
PRISM_META="https://meta.prismlauncher.org/v1"
META_FILE=$(mktemp)
ENTRIES_FILE=$(mktemp)
trap 'rm -f "$META_FILE" "$ENTRIES_FILE"' EXIT

# prefer unobfuscated metadata if available (matches download-minecraft.sh behavior)
META_SOURCE="${VERSION}_unobfuscated"
if ! curl -sfL "$PRISM_META/net.minecraft/${VERSION}_unobfuscated.json" -o "$META_FILE"; then
    META_SOURCE="$VERSION"
    curl -sfL "$PRISM_META/net.minecraft/$VERSION.json" -o "$META_FILE"
fi
echo "Using metadata: $META_SOURCE"

mkdir -p libs-cache

while read -r lib; do
    name=$(jq -r '.name' <<<"$lib")
    url=$(jq -r '.downloads.artifact.url' <<<"$lib")
    sha1=$(jq -r '.downloads.artifact.sha1' <<<"$lib")
    size=$(jq -r '.downloads.artifact.size' <<<"$lib")
    group="${name%%:*}"
    artifact=$(awk -F: '{print $2}' <<<"$name")
    version=$(awk -F: '{print $3}' <<<"$name")

    if [[ "$group" == com.mojang ]]; then
        [ "$MOJANG_LIBS" -eq 1 ] || continue
        category="mojang"
    else
        [ "$THIRD_PARTY_LIBS" -eq 1 ] || continue
        category="third-party"
    fi

    jar="libs-cache/${group}_${artifact}_${version}.jar"
    out="current/libraries/${category}/${group}/${artifact}"

    if [ ! -f "$jar" ]; then
        echo "Downloading $name..."
        curl -sfL "$url" -o "$jar"
    fi

    if [ ! -d "$out" ]; then
        echo "Decompiling $name..."
        mkdir -p "$out"
        java -jar vineflower.jar "$jar" "$out"
    fi

    jq -nc \
        --arg name "$name" \
        --arg group "$group" \
        --arg artifact "$artifact" \
        --arg version "$version" \
        --arg category "$category" \
        --arg url "$url" \
        --arg sha1 "$sha1" \
        --argjson size "$size" \
        --arg jar "$jar" \
        --arg decompiled "$out" \
        '{name:$name, group:$group, artifact:$artifact, version:$version, category:$category, url:$url, sha1:$sha1, size:$size, jar:$jar, decompiled:$decompiled}' \
        >>"$ENTRIES_FILE"
done < <(jq -c '.libraries[] | select(.downloads.artifact != null)' "$META_FILE")

jq -s \
    --arg minecraft_version "$VERSION" \
    --arg metadata_source "$META_SOURCE" \
    --arg generated_at "$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    --argjson mojang $((MOJANG_LIBS)) \
    --argjson third_party $((THIRD_PARTY_LIBS)) \
    '{minecraft_version:$minecraft_version, metadata_source:$metadata_source, generated_at:$generated_at, flags:{mojang_libs:($mojang==1), third_party_libs:($third_party==1)}, libraries:.}' \
    "$ENTRIES_FILE" >current/metadata.json

echo "Wrote current/metadata.json ($(jq '.libraries | length' current/metadata.json) libraries)"
