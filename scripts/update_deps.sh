#!/usr/bin/env bash
cd "$(dirname "$0")/.." || exit 1

echo ""
echo "➡ Updating npm dependencies"
npm update

echo ""
echo "➡ Updating Go dependencies"
go get -u ./...
go mod tidy

echo ""
echo "✅ Dependencies updated!"
