#!/bin/bash

python3.11 fetch_and_generate.py

# Copy the downloaded images and generated review page to public directory
PUBLIC_DIR="../../public/"
cp -R  downloaded_images $PUBLIC_DIR
cp -R  profile_pictures $PUBLIC_DIR
cp reviews.html $PUBLIC_DIR