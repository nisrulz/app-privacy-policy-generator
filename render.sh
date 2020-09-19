#!/usr/bin/env bash

# Using pug-cli to render from pug to html
# Install pug-cli: npm install -g pug-cli
pug src/index.pug --pretty --out public

# Using sass-cli to render from sass to css
# Install sass-cli: npm install -g sass
sass src/sass/style.sass public/css/style.css