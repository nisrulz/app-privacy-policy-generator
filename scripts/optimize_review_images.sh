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
  # Repoint a .png reference only when its WebP actually exists; otherwise a
  # failed cwebp conversion would leave the JSON pointing at a file that was
  # never produced. The directory is passed via the environment because perl's
  # -i resolves its target from @ARGV before the program runs, so the directory
  # must not appear there.
  REVIEWS_IMG_DIR="$DST" perl -0pi -e '
    my $dir = $ENV{REVIEWS_IMG_DIR};
    # Do not exit from this program: with -i the output file is opened and
    # truncated at startup, and exit skips the -p loop'"'"'s implicit print,
    # which would silently leave the JSON empty. Fall back to an empty set so
    # every reference is left untouched.
    my %webp;
    if (opendir(my $dh, $dir)) {
      # Read into an explicit loop rather than map/grep over readdir: those alias
      # $_ to readdir'"'"'s buffer, and the in-place s/// then corrupts entries,
      # silently yielding a partial set (40 of 79 where 79 WebP files existed).
      # Keys are the full relative path stem, matching the capture below, which
      # excludes the trailing ".png".
      while (my $entry = readdir($dh)) {
        next unless $entry =~ /\.webp\z/;
        my $base = $entry;
        $base =~ s/\.webp\z//;
        $webp{"./downloaded_images/$base"} = 1;
      }
      closedir $dh;
    }
    s{(\./downloaded_images/[A-Za-z0-9._-]+?)\.png}{
      my $stem = $1;
      exists $webp{$stem} ? "$stem.webp" : "$stem.png";
    }ge;
  ' "$JSON"
fi

echo "  Converted $converted PNG(s) to WebP in $DST"
