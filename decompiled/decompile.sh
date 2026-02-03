#!/bin/bash
cd "$(dirname "$0")"

java -jar vineflower.jar minecraft.jar ./current

ln -sf "$PWD/current/assets/minecraft/lang/en_us.json" ../pkg/data/generate/en_us.json