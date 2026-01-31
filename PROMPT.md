# PROMPT.md

Prompt to:

1. Compute diff against the previous Minecraft version;
2. Examine differences in the network protocol implementation;
3. Update the Go bindings so they are compatible with the new Minecraft version;

## Context

[go-mclib/data](https://github.com/go-mclib/data) is a repository that contains Go bindings for the Minecraft packet types. On a higher level, it is a wrapper around [go-mclib/protocol](https://github.com/go-mclib/protocol), which contains low-level methods for writing and reading packets and network structures that Minecraft uses (such as `Long`, `VarInt`, TextComponent, NBT data, etc.). This repository is then used in [go-mclib/client](https://github.com/go-mclib/client) to implement a Minecraft client in the Go programming language.

## Big-picture Goal

To create a Minecraft bot that is capable of connecting to a Minecraft server and interacting with the game world, sending chat and commands, etc.

## Problem

The Minecraft protocol is ever-evolving - each new Minecraft version changes the network protocol in some way. These changes are not officially documented anywhere - instead, we rely on third-party sources (such as the Minecraft wiki: <https://minecraft.wiki/w/Java_Edition_protocol/Packets>) and reverse-engineering the protocol from the source code of the Minecraft server and client, to adjust the Go bindings accordingly so they are compatible with the new Minecraft version. This library does NOT assume cross-compatibility between different protocol versions at the same time, instead it "lives in the moment" (if users want to use an older version, they can downgrade our package which has bindings compatible with the previous versions).

## The job of the AI agent

The job of you - the AI agent - is to:

1. Compute a diff against the previous Minecraft version;
    - it is the human's responsibility to ensure that there exists a folder for the current Minecraft version and the version before that in the `./decompiled` directory, typically named `./decompiled/current/...` and `./decompiled/previous/...`. If these folders do not exist, please point the human to `./decompiled/README.md` for instructions.
    - to compute diff of the network protocol: `diff -r ./decompiled/current/net/minecraft/network ./decompiled/previous/net/minecraft/network`
    - check also the Minecraft Wiki for third-party documentation of the packet structures: <https://minecraft.wiki/w/Java_Edition_protocol/Packets>. There is no guarantee this will be up-to-date, though it is well-documented.
2. Examine differences in the network protocol implementation;
    - these changes are unpredictable: they can not only differ in the packet definitions, but also in the fundamental network data types and logic, if this is the case, please discuss with the human before proceeding, as such changes may require a modification of the underlying `go-mclib/protocol` package;
3. Update the Go bindings so that they are compatible with the new Minecraft version. Run `go fmt` to format the bindings;
4. Summarize the changes, highlighting the ones you deem to be the most important;
5. Optional: suggest off/on-topic improvements to the codebase, in terms of code style, API design, readability, etc.;

You may use any tools/commands at your disposal, or the ones that are recommended either by this instruction set or by the human.

## The job of the human

The job of the human is to:

1. Review the changes made by the AI agent;
2. Communicate improvements and changes with the AI agent;
3. Push a new tag/version of the Go bindings, for use in our other packages;
4. Any additional tasks the human deems important to achieve the goal, at their sole discretion;

Now, please proceed with computing the diff against the previous Minecraft version and updating the Go bindings, as instructed above.
