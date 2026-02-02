#!/bin/bash
cd "$(dirname "$0")"

java -DbundlerMainClass="net.minecraft.data.Main" -jar ../packets_test/proxy/server/server.jar --server --reports

ln -sf "$PWD/generated/reports/blocks.json" ../pkg/data/generate/blocks.json
ln -sf "$PWD/generated/reports/items.json" ../pkg/data/generate/items.json
ln -sf "$PWD/generated/reports/registries.json" ../pkg/data/generate/registries.json
ln -sf "$PWD/generated/reports/packets.json" ../pkg/data/generate/packets.json
