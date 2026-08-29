#!/usr/bin/env bash
cd "$(dirname "$0")/.." || exit 1

make build

firebase serve --only hosting
