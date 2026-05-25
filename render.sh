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

echo ""
echo " 👨🏻‍💻 Starting..."
echo ""

LANGS="en"
BASE_DIR="public"

# --- Shared: Sass ---
npx -q sass src/sass/style.sass "$BASE_DIR/css/style.css"
echo " ✅  STEP 1: RENDER SASS > CSS"

# --- Shared: YAML → thirdparty JS ---
TS_JS_FILE=$(mktemp /tmp/thirdpartyservices.XXXXXX.js)
LOCALE_JS_FILE=$(mktemp /tmp/locale.XXXXXX.js)
trap 'rm -f "$TS_JS_FILE" "$LOCALE_JS_FILE"' EXIT
YAML_DIR="src/includes/yaml"
echo "/*  
    App Privacy Policy Generator: A simple web app to generate a generic 
    privacy policy for your Android, iOS, and Web apps

    Copyright 2017-Present Nishant Srivastava

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
 " >"$TS_JS_FILE"
TEMP_JSON_FILE="$YAML_DIR/thirdpartyservices.json"
touch $TEMP_JSON_FILE
npx -q js-yaml $YAML_DIR/thirdpartyservices.yml >$TEMP_JSON_FILE
JSON_DATA=$(cat $TEMP_JSON_FILE)
echo "var thirdPartyServicesJsonArray = $JSON_DATA" >>"$TS_JS_FILE"
rm $TEMP_JSON_FILE
echo " ✅  STEP 2: RENDER YAML > JSON"

# --- Shared: First pug render (for PurgeCSS reference) ---
npx -q pug3 src/index.pug --out "$BASE_DIR" --silent
echo " ✅  STEP 3: RENDER PUG > HTML (reference for PurgeCSS)"

# --- Shared: Minify CSS + PurgeCSS ---
npx -q uglifycss "$BASE_DIR/css/style.css" --output "$BASE_DIR/css/style.min.css"
rm -f "$BASE_DIR/css/style.css" "$BASE_DIR/css/style.css.map"
npx -q purgecss --css "$BASE_DIR/css/style.min.css" \
--content "$BASE_DIR/index.html" \
--output "$BASE_DIR/css/style.min.css"
echo " ✅  STEP 4: Minify CSS"

# --- Shared: Minify JS (except locale) ---
npx -q uglifyjs "$TS_JS_FILE" --output "$BASE_DIR/js/thirdpartyservices.min.js"
npx -q uglifyjs src/js/wizardMixin.js src/js/platformMixin.js src/js/localeMixin.js src/js/main.js --output "$BASE_DIR/js/main.min.js"
npx -q uglifyjs src/js/utils.js --output "$BASE_DIR/js/utils.min.js"
npx -q uglifyjs src/js/flycricket.js --output "$BASE_DIR/js/flycricket.min.js"
echo " ✅  STEP 5: Minify JS"

# --- Shared: Copy vendor assets ---
mkdir -p "$BASE_DIR/js/vendor"
cp src/includes/vendor/vue.global.prod.js "$BASE_DIR/js/vendor/vue.global.prod.js"
cp src/includes/vendor/to-markdown.min.js "$BASE_DIR/js/vendor/to-markdown.min.js"
mkdir -p "$BASE_DIR/images/vendor"
cp src/includes/vendor/kofi1.png "$BASE_DIR/images/vendor/kofi1.png"
echo " ✅  STEP 6: Copy vendor assets"

# --- Per-locale: Render HTML + locale JS ---
for LANG in $LANGS; do
  OUT_DIR="$BASE_DIR"
  if [ "$LANG" != "en" ]; then
    OUT_DIR="$BASE_DIR/$LANG"
  fi

  echo "window.__locale = $(cat src/locales/$LANG.json);" > "$LOCALE_JS_FILE"
  mkdir -p "$OUT_DIR/js"
  npx -q uglifyjs "$LOCALE_JS_FILE" --output "$OUT_DIR/js/locale.min.js"

  # Render pug with lang variable override (sed then restore)
  sed -i.bak "s/- var lang = '[a-z]*'/- var lang = '$LANG'/" src/index.pug
  npx -q pug3 src/index.pug --out "$OUT_DIR" --silent
  mv src/index.pug.bak src/index.pug

  echo " ✅  STEP 7.$LANG: RENDER PUG > HTML ($LANG)"
done

echo ""
echo " 📈  STEP 8: Optimized website is ready"
