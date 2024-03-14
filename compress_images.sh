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

echo "Navigate to Images directory"
cd public/images

#--------------------------------- Compress SVG files ---------------------------------#
svgo -f app_graphics temp
mv temp/*.svg app_graphics

svgo -f app_icons temp
mv temp/*.svg app_icons

svgo -f social_icons temp
mv temp/*.svg social_icons

# Delete Temp directory
rm -rf temp

#--------------------------------- Compress PNG files ---------------------------------#
npx png-minify minify third_party_logos/*.png

