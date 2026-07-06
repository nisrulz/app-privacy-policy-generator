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

# Compress SVG files
npx -q svgo -f app_graphics "$TEMP_DIR"
mv "$TEMP_DIR"/*.svg app_graphics

npx -q svgo -f app_icons "$TEMP_DIR"
mv "$TEMP_DIR"/*.svg app_icons

# Compress PNG files
npx -q png-minify minify third_party_logos/*.png

# Compress downloaded review images
npx -q png-minify minify ../downloaded_images/*.png 2>/dev/null || true

