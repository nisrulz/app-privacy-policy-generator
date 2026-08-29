#!/bin/bash
set -e

PORT="${1:-8000}"
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
BIN="$ROOT/.bin/serve"

mkdir -p "$ROOT/.bin"
go build -o "$BIN" ./cmd/build

exec "$BIN" -serve -port "$PORT"
