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
  check              Run tests, vet, and build checks
  test               Run E2E tests (requires make serve running)
  update-deps        Update Go dependencies

Data
  reviews            Generate reviews page
                     Example: make reviews FORCE="true"

Deploy
  deploy             Deploy to Firebase Hosting
                     Example: make deploy VERSION="3.0.9"
  firebase-preview   Build and preview via Firebase local server

EOF
