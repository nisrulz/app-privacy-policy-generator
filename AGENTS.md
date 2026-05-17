# AGENTS.md

## Project Overview

Static web app (Vue.js + Pug + Sass) that generates privacy policies and terms & conditions for Android/iOS/Web apps. Single `index.html` output — no bundler, no hot-reload. Compiled via CLI tools (`pug3`, `sass`, `js-yaml`, `uglify-js`, `uglifycss`, `purgecss`).

## Source Layout

- `src/index.pug` — entry point (includes all partials)
- `src/includes/content/privacy_policy/` — privacy templates: `gdpr.pug`, `simple.pug`, `no_tracking.pug`
- `src/includes/content/tnc/` — T&C template: `simple.pug`
- `src/includes/content/wizard/` — wizard steps: `step_1.pug`–`step_7.pug`
- `src/js/main.js` — Vue app: data, computed properties, `generate()` function
- `src/js/utils.js` — utility helpers
- `src/includes/yaml/thirdpartyservices.yml` — 3rd-party service definitions (source of truth); JS is auto-generated during build
- `src/sass/` — Sass partials compiled to a single stylesheet

## Build

```sh
npm install
npm run build
```

Compiles `src/` → `public/index.html`, `public/css/style.min.css`, `public/js/*.min.js`. Edit source files only (under `src/`), never edit `public/` files directly.

Build pipeline:
1. `pug3` renders `src/index.pug` → `public/index.html`
2. `sass` compiles `src/sass/` → `public/css/style.css`
3. `js-yaml` converts `src/includes/yaml/thirdpartyservices.yml` → temp JS (intermediate file, not tracked)
4. `uglifycss` + `purgecss` minify CSS → `public/css/style.min.css`
5. `uglifyjs` minifies JS → `public/js/*.min.js`

Additional global tools (image compression, deployment):
```sh
npm install -g firebase-tools svgo png-minify
```

Dev: open `public/index.html` in browser after render. No dev server.

## npm Scripts

| Command | Action |
|---------|--------|
| `npm run build` | Render templates and minify outputs |
| `npm run serve` | Start Firebase local dev server |
| `npm run compress-images` | Optimize SVG and PNG images |

## Code Conventions

- **No comments** in source code unless documenting a complex legal rationale
- **Commit style**: `type(scope): Description` — e.g. `fix(privacy): Add data breach notification`, `feat(platform): Add platform type selection`
- **Vue** is loaded via CDN (not npm); all state is on a single Vue instance in `main.js`
- **Pug** templates use `{{ variable }}` interpolation (Vue mustache syntax), not Pug's native interpolation
- **Computed properties** for conditional logic (e.g. `isMobileApp`, `isWebApp`, `isBothPlatforms`) — prefer over inline string matching in templates
- **Derived text** (e.g. `deviceType`, `platformDesc`) computed inside `generate()` following existing pattern, not watchers

## Key Files

| File | Purpose |
|------|---------|
| `src/js/main.js` | Vue app — data, computed, `generate()`, validation |
| `src/includes/content/privacy_policy/gdpr.pug` | GDPR/CCPA-compliant policy |
| `src/includes/content/privacy_policy/simple.pug` | Simple/standard policy |
| `src/includes/content/privacy_policy/no_tracking.pug` | No-tracking policy |
| `src/includes/content/tnc/simple.pug` | Terms & Conditions |
| `src/includes/content/wizard/step_4.pug` | Platform type selector |
| `src/includes/content/wizard/step_2.pug` | EU Representative field |

## Testing

No test framework. Verify by running `npm run build` and visually inspecting `public/index.html`.

## Deployment

```sh
firebase login
firebase deploy -m "<version>"
```

CI via GitHub Actions (.github/workflows/). Only maintainers can deploy. Firebase deploy is guarded — do not run deploy without approval.

## Branch Strategy

- `master` — stable, deployed to production
- Feature/fix branches off `master`, PRs merge back

## Skills

- `.agents/skills/add-thirdparty-service/` — Load this skill when adding a new 3rd-party service entry.

## Quick Workflow for Edits

1. Edit templates under `src/`
2. Run `npm run build` to produce `public/`
3. Verify `public/` renders correctly (open `public/index.html`)
4. Commit source changes only

Key docs to read first: `dev-doc.md`, `render.sh`, `src/includes/content/*`, `src/includes/yaml/thirdpartyservices.yml`, `package.json`, `tools/` README files.

## Contribution Rules

- Bug fixes only for PRs (features discussed via issue first)
- New 3rd-party service: load the `add-thirdparty-service` skill or manually add entry to `src/includes/yaml/thirdpartyservices.yml` + logo (160×160) to `public/images/third_party_logos/`
- Do NOT commit `public/index.html`, `public/js/main.min.js`, or other generated files
- `render.sh` uses `npx` to resolve all CLI tools from devDependencies (`@tokilabs/pug3-cli` is the required pug fork)
- Build outputs expected (do not remove unless intentional): `public/index.html`, `public/css/style.min.css`, `public/js/*.min.js`
