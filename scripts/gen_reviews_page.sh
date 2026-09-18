#!/usr/bin/env bash
set -euo pipefail

root="$(cd "$(dirname "$0")/.." && pwd)"

cd "$root/tools/reviews-page-generator"

if [ "${1:-}" = "-f" ]; then
  echo "Force-fetching from GitHub API..."
  go run . --force-fetch
else
  echo "Using cached data (use -f to force re-fetch from GitHub)"
  go run .
fi

cd "$root"
echo "Optimizing review images..."
./scripts/optimize_review_images.sh
