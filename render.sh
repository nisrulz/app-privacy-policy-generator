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

echo "⚙️ STEP 1: RENDER PUG > HTML"

# Using pug-cli to render from pug to html
# Specifically a fork
# - NPM: https://www.npmjs.com/package/@anduh/pug-cli
# - Github: https://github.com/Anduh/pug-cli
# because the original is not maintained anymore!
# Install pug-cli: npm install -g @anduh/pug-cli

pug3 src/index.pug --pretty --out public

echo "⚙️ STEP 2: RENDER SASS > CSS"
# Using sass-cli to render from sass to css
# Install sass-cli: npm install -g sass
sass src/sass/style.sass public/css/style.css

echo ""
echo "⚙️ STEP 3: RENDER YAML > JSON"

# Using js-yaml to convert yaml to json
# Install js-yaml: npm install -g js-yaml
JS_FILE="public/js/thirdpartyservices.js"
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
