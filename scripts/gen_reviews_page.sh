#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/../tools/reviews-page-generator"

if [ "${1:-}" = "-f" ]; then
  echo "Force-fetching from GitHub API..."
  go run . --force-fetch
else
  echo "Using cached data (use -f to force re-fetch from GitHub)"
  go run .
fi

echo "✅ Done!"
