# Copilot instructions for this repo

Purpose: quick reference for Copilot sessions to build, run, and modify site correctly.

---

Build, test, lint commands

- Build (full): ./render.sh
  - Generates public/index.html, public/css/style.min.css, public/js/*min.js
  - Requires global CLI tools (install once):
    - @tokilabs/pug3-cli, sass, js-yaml, firebase-tools, svgo, png-minify, uglifycss, uglify-js, html-minifier, purgecss
  - Install example:
    npm install -g @tokilabs/pug3-cli sass js-yaml firebase-tools svgo png-minify uglifycss uglify-js html-minifier purgecss

- Serve (local):
  firebase serve --only hosting

- Deploy (production, maintainers only):
  firebase login && firebase deploy -m "<version>"

- Compress images: ./compress_images.sh

- Update deps: npm update

- Tests / Lint: none defined in package.json. No single-test command available.

---

High-level architecture

- Static SPA. Source in src/ (pug templates, sass, yaml). Rendered assets in public/.
- Build pipeline (render.sh):
  1. pug3 renders pug -> public HTML
  2. sass compiles SASS -> CSS
  3. js-yaml converts src/includes/yaml/thirdpartyservices.yml -> src/js/thirdpartyservices.js
  4. minify + purgecss produce final css/js in public/
- Hosting: Firebase Hosting (public/ directory). GitHub Actions run deploy workflows on merge/PR.
- Third-party services list driven by src/includes/yaml/thirdpartyservices.yml and rendered into public/js/thirdpartyservices.min.js

---

Key conventions

- Templates:
  - Privacy policy templates: src/includes/content/privacy_policy
  - Terms & Conditions: src/includes/content/tnc
  - Add new template: add pug file under those folders and run ./render.sh to validate

- thirdpartyservices.yml entries:
  - Required fields: name, enabled (true/false), logo (path under public/images/third_party_logos), link. Example in dev-doc.md.
  - Logo requirement: 160x160 PNG. Add file to public/images/third_party_logos/.

- build outputs expected (do not remove unless intentional): public/index.html, public/css/style.min.css, public/js/*min.js

- render.sh uses pug3-cli fork (@tokilabs/pug3-cli). Use that exact CLI when rendering locally.

- Firebase deploy guarded: maintainers only. Do not run deploy without approval.

- Repo tools: extra utilities in tools/ (e.g., tools/reviews-page-generator). Check their README before running.

---

AI / agent notes

- See AGENTS.md for agent-specific guidance used in this repo.
- Key docs for Copilot to read first: dev-doc.md, render.sh, src/includes/content/*, src/includes/yaml/thirdpartyservices.yml, package.json, tools/README files.

---

Quick workflow for edits

- Edit templates under src/
- Run ./render.sh locally to produce public/
- Verify public/ pages render correctly (open public/index.html)
- Commit source changes (do not add large generated files unless PR expects preview)

