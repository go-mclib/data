# Decompiled root

This is the root directory for decompiled Minecraft client folders. Keep the decompiled sources in `.gitignore`!

Decompiling the source code is helpful to reveal the internals of the Minecraft client/server networking code, which allows us to pipe the diffs between versions to LLMs and develop the Go bindings faster. For this purpose, we are using a [Claude Code agent](./AGENTS.md).

## Decompiling

This repository uses [MinecraftDecompiler](https://github.com/MaxPixelStudios/MinecraftDecompiler) under the hood. You must download the latest .jar to this directory, see `download.sh`. See also <https://minecraft.wiki/w/Tutorial:See_Minecraft%27s_code>
