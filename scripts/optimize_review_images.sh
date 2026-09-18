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
# 
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
# 
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

set -euo pipefail

cd "$(dirname "$0")/.." || exit 1

SRC="tools/reviews-page-generator/downloaded_images"
DST="public/downloaded_images"
JSON="public/reviews-data.json"

if [ ! -d "$SRC" ]; then
  echo " ✗  Missing $SRC"
  exit 1
fi

if ! command -v cwebp > /dev/null 2>&1; then
  echo " ✗  cwebp is not installed (brew install webp)"
  exit 1
fi

mkdir -p "$DST"

converted=0
for f in "$SRC"/*.png; do
  [ -e "$f" ] || continue
  name=$(basename "${f%.png}")
  if cwebp -q 85 -m 6 "$f" -o "$DST/$name.webp" > /dev/null 2>&1; then
    converted=$((converted + 1))
  fi
done

for ext in jpg jpeg gif svg webp; do
  for f in "$SRC"/*."$ext"; do
    [ -e "$f" ] || continue
    cp "$f" "$DST/"
  done
done

find "$DST" -maxdepth 1 -name '*.png' -delete

if [ -f "$JSON" ]; then
  perl -0pi -e 's{(\./downloaded_images/[A-Za-z0-9._-]+?)\.png}{$1.webp}g' "$JSON"
fi

echo "  Converted $converted PNG(s) to WebP in $DST"
