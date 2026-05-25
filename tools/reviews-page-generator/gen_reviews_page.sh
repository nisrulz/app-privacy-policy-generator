#!/bin/bash
set -euo pipefail

if [ "${1:-}" = "-f" ]; then
  echo "Force-fetching from GitHub API..."
  uv run fetch_and_generate.py --force-fetch
else
  echo "Using cached data (use -f to force re-fetch from GitHub)"
  uv run fetch_and_generate.py
fi
