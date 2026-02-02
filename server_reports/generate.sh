#!/bin/bash
cd "$(dirname "$0")"

java -DbundlerMainClass="net.minecraft.data.Main" -jar ../proxy/server/server.jar --server --reports
