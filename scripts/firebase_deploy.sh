#!/usr/bin/env bash
cd "$(dirname "$0")/.." || exit 1

version="${1:-}"
if [ -z "$version" ]; then
  read -p "❓  Specify version?   " version
fi

make build

firebase deploy -m "$version"
