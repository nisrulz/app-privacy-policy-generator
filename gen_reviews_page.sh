#!/usr/bin/env bash

# App Privacy Policy Generator: A simple web app to generate a generic 
# privacy policy for your Android, iOS, and Web apps
# 
# Copyright 2017-Present Nishant Srivastava
# 
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

set -euo pipefail

cd "$(dirname "$0")/tools/reviews-page-generator"

if [ "${1:-}" = "-f" ]; then
  echo "Force-fetching from GitHub API..."
  uv run fetch_and_generate.py --force-fetch
else
  echo "Using cached data (use -f to force re-fetch from GitHub)"
  uv run fetch_and_generate.py
fi

echo ""
echo "✅ Done!"
