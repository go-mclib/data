#!/bin/bash
cd "$(dirname "$0")"

java -DbundlerMainClass="net.minecraft.data.Main" -jar ../proxy/server/server.jar --server --reports

ln -sf "$PWD/generated/reports/commands.json" ../pkg/data/generate/commands.json
ln -sf "$PWD/generated/reports/blocks.json" ../pkg/data/generate/blocks.json
ln -sf "$PWD/generated/reports/registries.json" ../pkg/data/generate/registries.json
ln -sf "$PWD/generated/reports/packets.json" ../pkg/data/generate/packets.json
ln -sf "$PWD/generated/reports/minecraft/components/item" ../pkg/data/generate/items
