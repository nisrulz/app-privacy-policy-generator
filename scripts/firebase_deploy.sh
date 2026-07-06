#!/usr/bin/env bash
cd "$(dirname "$0")/.." || exit 1

version="${1:-}"
if [ -z "$version" ]; then
  read -p "❓  Specify version?   " version
fi
echo ""
echo "➡ Building from templates"
make build
echo ""
echo "➡ Deploying to Firebase Hosting"
firebase deploy -m "$version"
echo ""
echo "✅ Done!"
