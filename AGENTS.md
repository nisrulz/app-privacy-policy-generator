# AGENTS.md

## Project Overview

Static web app (Vue.js + Pug) that generates privacy policies and terms & conditions for Android/iOS/Web apps. Single `index.html` output — no bundler, no hot-reload. Built via Go toolchain (`cmd/build/`) which orchestrates `pug3` via `npx`, handles YAML→JS, Less compilation (with built-in compression), JS minification, and cache-busting natively.

## Source Layout

- `src/index.pug` — entry point (includes all partials); `lang`, `flycricket`, `noTracking`, `gdpr` compile-time variables; `data-theme="light"` on `<html>`
- `src/locales/` — locale JSON files (one per language); `en.json` is the source of truth; `de.json` also exists; `_flag` key for each locale
- `src/includes/content.pug` — includes all content: wizard, privacy policies, T&C, disclaimer, FAQ, affiliate
- `src/includes/mixins.pug` — shared Pug mixins: `modal()`, `outputButtons()`, `attribution()`, `flycricketDeploy()`, `wizardProgress()`, `wizardMobileLogo()`, `wizardSidebar()`
- `src/includes/meta.pug` — aggregates meta partials: `base.pug`, `app-description.pug`, `author.pug`, `webapp-support.pug`, `favicon.pug`, `link-preview.pug`
- `src/includes/css.pug` — CSS includes (inline normalize.css, preload sidebar image, stylesheet link)
- `src/includes/license-header.pug` — GPL license comment block
- `src/includes/content/privacy_policies.pug` — wraps privacy policy includes with compile-time conditionals (`noTracking`, `gdpr`)
- `src/includes/content/tncs.pug` — wraps T&C include
- `src/includes/content/wizard.pug` — wraps all 8 wizard step includes
- `src/includes/content/privacy_policy/` — privacy templates: `gdpr.pug`, `simple.pug`, `no_tracking.pug`; all text via `translate('key')`
- `src/includes/content/tnc/` — T&C template: `simple.pug`; all text via `translate('key')`
- `src/includes/content/wizard/` — wizard steps: `step_1.pug`–`step_8.pug`; all labels via `translate('key')`
- `src/includes/content/faq.pug` — FAQ modal with Vue-controlled visibility; text via `translate('key')`
- `src/includes/content/disclaimer.pug` — Disclaimer modal; text via `translate('key')`
- `src/includes/content/affiliate/` — affiliate integrations (Flycricket)
- `src/js/main.js` — Vue app entry point: creates app with 8 mixins, registers `translate()`, `_updateMeta()`, `switchLocale()` globally
- `src/js/localeMixin.js` — Vue mixin: `_updateMeta()` on mount
- `src/js/wizardMixin.js` — Vue mixin: wizard navigation (`nextStep`/`prevStep`), `canAdvance` guard, `totalWizardSteps: 8`, `contentRenderType`
- `src/js/platformMixin.js` — Vue mixin: platform-aware text descriptors, `_setPlatformText()`, computed `isMobileApp`/`isWebApp`/`isWindowsApp`
- `src/js/formDataMixin.js` — Vue mixin: all form state (`appName`, `appContact`, `platforms`, `typeOfPolicy`, `thirdPartyServices`, etc.), `availableLocales`, computed filtered services
- `src/js/generatorMixin.js` — Vue mixin: `generate()`, validation (`_validateRequiredFields`), derived text setters (`_setDevOrCompanyName`, `_setPidInfo`, `_setAppTypeText`)
- `src/js/modalMixin.js` — Vue mixin: modal toggle methods (`togglePrivacyModalVisibility`, `toggleGDPRPrivacyModalVisibility`, etc.)
- `src/js/thirdPartyMixin.js` — Vue mixin: third-party service helpers (`tpsName()` with locale-aware fallback, `toggleState`, `setTypeOfPolicyInt`)
- `src/js/contentMixin.js` — Vue mixin: HTML/Markdown export (`getHtml`/`getMarkdown`), Flycricket deploy (`deployFcSimple`)
- `src/js/utils.js` — utility helpers (`convertHtmlToMd`, `getRawHTML`, `getContent`, `getTitle`, `loadInTextView`)
- `src/js/flycricket.js` — Flycricket form submission helpers (`fc_deploy_simple`, `fc_deploy_notracking`, `fc_deploy_gdpr`)
- `src/includes/yaml/thirdpartyservices.yml` — 3rd-party service definitions (source of truth); supports locale-aware `name_{code}` fields; JS is auto-generated during build
- `src/includes/vendor/` — vendored third-party assets: `vue.global.prod.js`, `to-markdown.min.js`
- `src/less/style.less` — Less entry point (`@import`s 6 partials); compiled to CSS by Go build
- `src/less/_variables.less` — Less variables (`@color-primary`, `@font-sans`, `@border-radius`, etc.)
- `src/less/_base.less` — Base layout, typography, hero, cards, footer
- `src/less/_scrollbar.less` — Custom scrollbar styles
- `src/less/_overrides.less` — Label, button, input, modal, checkbox overrides
- `src/less/_bg-pattern.less` — `.bg-pattern` SVG data URI background
- `src/less/_dark.less` — Dark theme overrides
- `public/site.webmanifest` — PWA manifest (tracked, edit directly)
- `public/sw.js` — service worker with offline caching (tracked, edit directly)
- `public/404.html` — custom 404 page (tracked, edit directly)
- `public/robots.txt` — robots exclusion (tracked, edit directly)
- `public/humans.txt` — author/tech credits (tracked, edit directly)
- `public/logo.svg` — SVG logo (tracked, edit directly)
- `public/images/third_party_logos/` — third-party service logo images (tracked, add via `add-thirdparty-service` skill)
- `public/images/app_graphics/` — app graphics (side_image.png, 404.svg, etc.)
- `public/images/app_icons/` — UI icons (disclaimer.svg, etc.)
- `public/favicon*`, `public/apple-touch-icon.png`, `public/android-chrome-*` — favicons and PWA icons
- `public/reviews.html` — generated reviews page (committed for convenience)
- `public/reviews-data.json` — generated reviews data (committed for convenience)
- `public/profile_pictures/` — reviewer profile pictures (gitignored, generated by reviews tool)
- `cmd/build/*.go` — Go build toolchain source files (edit to modify build pipeline); includes `serve.go` (HTTP server for local dev)

## Build

```sh
npm install
npm run build
```

Compiles `src/` → `public/index.html`, `public/css/style.min.css`, `public/js/*.min.js`. Build outputs are tracked.

Build pipeline:
1. Go compiles `src/less/style.less` → `public/css/style.min.css` via `toakleaf/less.go`
2. Go parses `src/includes/yaml/thirdpartyservices.yml` → `public/tmp/thirdpartyservices.js`
3. Go builds locales registry from `src/locales/` → `public/js/locales.min.js`
4. `pug3` renders `src/index.pug` → `public/index.html`
5. Go concatenates all 8 mixins + `main.js` → minified `public/js/main.min.js`; separately minifies `utils.min.js`, `thirdpartyservices.min.js`, `flycricket.min.js`
6. Go copies vendor assets (Vue, to-markdown, Ko-fi image) to `public/`
7. Per-locale: `pug3` renders HTML with `lang` override; Go generates `locale.min.js`
8. Cache-busting: `?v=<md5>` appended to all CSS/JS references in HTML files

Dev: `make serve` (builds then serves `public/` on port 8000 via Go HTTP server).

## Code Conventions

- **No comments** in source code unless documenting a complex legal rationale
- **Commit style**: `type(scope): Description`
- **Vue.js 3** is self-hosted (served from `public/js/vendor/vue.global.prod.js`); state is split across 8 mixin files merged via `mixins: [...]` in `main.js`
- **Pug** templates use `{{ variable }}` interpolation (Vue.js mustache syntax), not Pug's native interpolation
- **Translation**: use `translate('key')` in templates (global method set in `main.js`), not `$t('key')`
- **Computed properties** in Vue mixin `computed` blocks — prefer over inline string matching
- **Derived text** (e.g. `deviceType`, `platformDesc`) computed inside `generate()` following existing pattern, not watchers
- **Compile-time Pug variables** set in `index.pug`: `flycricket`, `noTracking`, `gdpr` (booleans control template inclusion)
- **`data-theme`** attribute on `<html>` element; dark mode support via `_dark.less` and `toggleTheme()` in `main.js`
- **Go 1.25** — toolchain uses `github.com/tdewolff/minify/v2` for JS and `gopkg.in/yaml.v3` for YAML parsing

## Testing

No test framework. Verify by running `npm run build` and visually inspecting `public/index.html`.

## Deployment

```sh
firebase login
firebase deploy -m "<version>"
```

CI via GitHub Actions (`.github/workflows/`): production deploy on push to `master`, preview deploy on PRs from same repo. Only maintainers can deploy.

## Skills

- `.agents/skills/add-thirdparty-service/` — Load this skill when adding a new 3rd-party service entry.
- `.agents/skills/add-localization/` — Load this skill when adding a new language translation.

## Quick Workflow for Edits

1. Edit templates under `src/`
2. Run `npm run build` to produce `public/`
3. Verify `public/` renders correctly (open `public/index.html`)
4. Commit source and build output changes

## Contribution Rules

- Bug fixes only for PRs (features discussed via issue first)
- New 3rd-party service: load the `add-thirdparty-service` skill or manually add entry to `src/includes/yaml/thirdpartyservices.yml` + logo (160×160) to `public/images/third_party_logos/`
- `cmd/build/` uses `npx` to resolve `pug3` from devDependencies (`@tokilabs/pug3-cli` is the required pug fork)
- Build outputs expected (do not remove unless intentional): `public/index.html`, `public/css/style.min.css`, `public/js/*.min.js`, `public/js/vendor/`, `public/images/vendor/`
- `public/sw.js`, `public/site.webmanifest`, `public/404.html`, `public/robots.txt`, `public/humans.txt`, `public/logo.svg` are tracked source files (not generated)
- `public/reviews.html` and `public/reviews-data.json` are generated by `tools/reviews-page-generator/` but tracked for convenience
