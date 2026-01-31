#!/bin/bash
cd "$(dirname "$0")"

claude < <(cat PROMPT.md)
