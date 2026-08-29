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

cd "$(dirname "$0")/.." || exit 1

TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT

cd public/images

size_of() {
  local bytes
  bytes=$(stat -f%z "$1" 2>/dev/null || stat -c%s "$1" 2>/dev/null)
  echo "${bytes:-0}"
}

# Render a byte count as a short human-readable string (e.g. 8286 -> "8.1 KiB")
human_size() {
  local bytes=$1
  if [ "$bytes" -ge 1048576 ]; then
    awk "BEGIN { printf \"%.2f MiB\", $bytes / 1048576 }"
  elif [ "$bytes" -ge 1024 ]; then
    awk "BEGIN { printf \"%.1f KiB\", $bytes / 1024 }"
  else
    echo "${bytes} B"
  fi
}

changed=0
total=0

# Print the size change for a file as a percentage decrease (smaller) or
# increase (larger) after compression. Files that do not change are skipped
# to keep the output readable.
report_change() {
  local file="$1" before="$2" after="$3"
  local diff pct sign label
  diff=$((before - after))
  [ "$diff" -eq 0 ] && return
  total=$((total + 1))
  changed=$((changed + 1))
  if [ "$before" -eq 0 ]; then
    pct="0.0"
  else
    pct=$(awk "BEGIN { printf \"%.1f\", ($diff * 100) / $before }")
  fi
  if [ "$diff" -gt 0 ]; then
    sign="-"; label="decrease"
  else
    sign="+"; label="increase"
  fi
  printf "  %-55s %s -> %s (%s%.1f%% %s)\n" \
    "$file" "$(human_size "$before")" "$(human_size "$after")" \
    "$sign" "${pct#-}" "$label"
}

# Compress SVG files and report the size change per file
for dir in app_graphics app_icons; do
  for file in "$dir"/*.svg; do
    [ -e "$file" ] || continue
    total=$((total + 1))
    before=$(size_of "$file")
    npx -q svgo "$file" -o "$TEMP_DIR/$(basename "$file")" > /dev/null 2>&1
    mv "$TEMP_DIR/$(basename "$file")" "$file"
    after=$(size_of "$file")
    report_change "$file" "$before" "$after"
  done
done

# Compress PNG files and report the size change per file
for file in third_party_logos/*.png; do
  [ -e "$file" ] || continue
  total=$((total + 1))
  before=$(size_of "$file")
  npx -q png-minify minify "$file" > /dev/null 2>&1
  after=$(size_of "$file")
  report_change "$file" "$before" "$after"
done

# Compress downloaded review images
for file in ../tools/reviews-page-generator/downloaded_images/*.png; do
  [ -e "$file" ] || continue
  total=$((total + 1))
  before=$(size_of "$file")
  npx -q png-minify minify "$file" > /dev/null 2>&1 || true
  after=$(size_of "$file")
  report_change "$file" "$before" "$after"
done

echo ""
echo "Compressed $total file(s): $changed changed, $((total - changed)) unchanged"

