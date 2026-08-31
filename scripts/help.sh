#!/bin/bash

cat <<'EOF'

Build
  build              Build the project
  clean              Clean public/ directory
  compress-images    Compress images in public/images
  serve              Build and serve locally
                     Example: make serve PORT="9090"
  watch              Watch for changes and rebuild

Code Quality
  format             Format Go source, templates, and tidy modules
  check              Run Go, golden, and Playwright checks
  test-ui            Run Playwright tests in UI mode
  test-debug         Run Playwright tests in debug mode
  update-deps        Update Go dependencies

Node.js is only required for the Playwright test commands.

Data
  reviews            Generate reviews page
                     Example: make reviews FORCE="true"

Deploy
  deploy             Deploy to Firebase Hosting
                     Example: make deploy VERSION="3.0.9"
  firebase-preview   Build and preview via Firebase local server

EOF
