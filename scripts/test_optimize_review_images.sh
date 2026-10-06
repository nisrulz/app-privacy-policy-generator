#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")/.." || exit 1

SCRIPT="scripts/optimize_review_images.sh"

work="$(mktemp -d)"
cleanup() {
  chmod -R u+rwx "$work" 2>/dev/null || true
  rm -rf "$work"
}
trap cleanup EXIT

fail() {
  echo " ✗  $1"
  exit 1
}

setup_root() {
  local root="$1"
  mkdir -p "$root/scripts" "$root/tools/reviews-page-generator/downloaded_images" "$root/public/downloaded_images"
  cp "$SCRIPT" "$root/scripts/"
}

stub_cwebp() {
  local stub_dir="$1" mode="$2"
  mkdir -p "$stub_dir"
  if [ "$mode" = "success" ]; then
    cat > "$stub_dir/cwebp" <<'STUB'
#!/usr/bin/env bash
out=""
while [ $# -gt 0 ]; do
  if [ "$1" = "-o" ]; then out="$2"; fi
  shift
done
printf 'fake webp' > "$out"
STUB
  else
    printf '#!/usr/bin/env bash\nexit 1\n' > "$stub_dir/cwebp"
  fi
  chmod +x "$stub_dir/cwebp"
}

root="$work/success"
setup_root "$root"
stub_cwebp "$work/stub-success" success
printf 'png' > "$root/tools/reviews-page-generator/downloaded_images/a.png"
printf '{"a":"./downloaded_images/a.png","b":"./downloaded_images/b.png"}' > "$root/public/reviews-data.json"
PATH="$work/stub-success:$PATH" bash "$root/scripts/optimize_review_images.sh" > /dev/null
grep -q '"a":"./downloaded_images/a.webp"' "$root/public/reviews-data.json" \
  || fail "success: a.png was not repointed to a.webp"
grep -q '"b":"./downloaded_images/b.png"' "$root/public/reviews-data.json" \
  || fail "success: b.png should keep its .png reference"
[ -f "$root/public/downloaded_images/a.webp" ] || fail "success: a.webp missing in DST"
[ ! -e "$root/public/downloaded_images/a.png" ] || fail "success: converted a.png should be deleted"
echo " ✓  repoints .png to .webp only when the WebP exists"

root="$work/failure"
setup_root "$root"
stub_cwebp "$work/stub-failure" failure
printf 'png' > "$root/tools/reviews-page-generator/downloaded_images/a.png"
printf 'png' > "$root/public/downloaded_images/a.png"
printf '{"a":"./downloaded_images/a.png"}' > "$root/public/reviews-data.json"
PATH="$work/stub-failure:$PATH" bash "$root/scripts/optimize_review_images.sh" > /dev/null
[ -f "$root/public/downloaded_images/a.png" ] || fail "failure: a.png should be kept when its conversion fails"
[ ! -e "$root/public/downloaded_images/a.webp" ] || fail "failure: a.webp should not exist"
grep -q '"a":"./downloaded_images/a.png"' "$root/public/reviews-data.json" \
  || fail "failure: JSON reference should stay .png"
echo " ✓  keeps the PNG and JSON reference when cwebp fails"

root="$work/locked"
setup_root "$root"
stub_cwebp "$work/stub-locked" success
printf 'png' > "$root/tools/reviews-page-generator/downloaded_images/a.png"
printf '{"a":"./downloaded_images/a.png"}' > "$root/public/reviews-data.json"
chmod 000 "$root/public/downloaded_images"
if ! PATH="$work/stub-locked:$PATH" bash "$root/scripts/optimize_review_images.sh" > /dev/null 2>&1; then
  chmod 755 "$root/public/downloaded_images"
  fail "locked: optimizer should survive an unreadable image directory"
fi
chmod 755 "$root/public/downloaded_images"
grep -q '"a":"./downloaded_images/a.png"' "$root/public/reviews-data.json" \
  || fail "locked: JSON must stay unchanged when the image directory is unreadable"
echo " ✓  leaves reviews-data.json untouched when the image directory is unreadable"
