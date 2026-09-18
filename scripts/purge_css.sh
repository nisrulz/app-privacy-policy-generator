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

BULMA="public/css/vendor/bulma.min.css"

if [ ! -f "$BULMA" ]; then
  echo " ✗  Missing $BULMA"
  exit 1
fi

size_of() {
  local bytes
  bytes=$(stat -f%z "$1" 2>/dev/null || stat -c%s "$1" 2>/dev/null)
  echo "${bytes:-0}"
}

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

before=$(size_of "$BULMA")

npx -y purgecss@6 \
  --css "$BULMA" \
  --content "src/tpl/**/*.html" "src/js/*.js" "public/index.html" "public/de/index.html" "public/reviews.html" \
  --safelist "is-active" "is-link" "is-primary" "is-info" "is-light" "is-small" "is-warning" "is-danger" "is-success" "is-disabled" "is-hidden-mobile" "is-hidden-tablet" "is-hidden-desktop" "is-hidden-touch" \
  --output "public/css/vendor/"

after=$(size_of "$BULMA")

printf "  %-45s %s -> %s\n" "$BULMA" "$(human_size "$before")" "$(human_size "$after")"
