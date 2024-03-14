#!/usr/bin/env bash

# App Privacy Policy Generator: A simple web app to generate a generic 
# privacy policy for your Android/iOS apps
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
# Using pug-cli to render from pug to html
# Specifically a fork
# - NPM: https://www.npmjs.com/package/@anduh/pug-cli
# - Github: https://github.com/Anduh/pug-cli
# because the original is not maintained anymore!
# Install pug-cli: npm install -g @anduh/pug-cli
pug3 src/index.pug --out public --silent
echo " ✅  STEP 1: RENDER PUG > HTML"

echo ""
# Using sass-cli to render from sass to css
sass src/sass/style.sass public/css/style.css
echo " ✅  STEP 2: RENDER SASS > CSS"

echo ""
# Using js-yaml to convert yaml to json
JS_FILE="src/js/thirdpartyservices.js"
YAML_DIR="src/includes/yaml"
echo "/*  
    App Privacy Policy Generator: A simple web app to generate a generic 
    privacy policy for your Android/iOS apps

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
 " >$JS_FILE
TEMP_JSON_FILE="$YAML_DIR/thirdpartyservices.json"
touch $TEMP_JSON_FILE
js-yaml $YAML_DIR/thirdpartyservices.yml >$TEMP_JSON_FILE
JSON_DATA=$(cat $TEMP_JSON_FILE)
echo "var thirdPartyServicesJsonArray = $JSON_DATA" >>$JS_FILE
rm $TEMP_JSON_FILE
echo " ✅  STEP 3: RENDER YAML > JSON"

echo ""
uglifycss public/css/style.css --output public/css/style.min.css 
rm public/css/style.css public/css/style.css.map
echo " ✅  STEP 4: Minify CSS"

echo ""
purgecss --css public/css/style.min.css \
--content public/index.html \
--output public/css/style.min.css
echo " ✅  STEP 5: Cleanup unused CSS"

echo ""
uglifyjs src/js/thirdpartyservices.js --output public/js/thirdpartyservices.min.js
uglifyjs src/js/main.js --output public/js/main.min.js
uglifyjs src/js/utils.js --output public/js/utils.min.js
uglifyjs src/js/flycricket.js --output public/js/flycricket.min.js
echo " ✅  STEP 6: Minify JS"

echo ""
html-minifier public/index.html \
--collapse-whitespace --keep-closing-slash --remove-comments \
--output public/index.html
echo " ✅  STEP 7: Minify HTML"