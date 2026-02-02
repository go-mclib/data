# Server

This directory contains a vanilla Minecraft server in offline-mode at `localhost:25566`. It is used to run the proxy and capture packets.

## Running the server

```bash
java -jar server.jar nogui
```

Then, in another terminal, run the [proxy](../README.md) and connect to `localhost:25565`.
