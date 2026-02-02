#!/bin/bash
cd "$(dirname "$0")"

java -DbundlerMainClass="net.minecraft.data.Main" -jar ../packets_test/proxy/server/server.jar --server --reports

ln -sf "$PWD/generated/reports/blocks.json" ../pkg/data/json_reports/blocks.json
ln -sf "$PWD/generated/reports/items.json" ../pkg/data/json_reports/items.json
ln -sf "$PWD/generated/reports/registries.json" ../pkg/data/json_reports/registries.json
