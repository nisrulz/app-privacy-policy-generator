#!/usr/bin/env bash
cd "$(dirname "$0")/.." || exit 1

echo ""
echo "➡ Building from templates"
make build
echo ""
echo "➡ Starting Firebase local preview"
firebase serve --only hosting
echo ""
echo "✅ Done!"
