# Localization Plan

**Goal**: Support generating privacy policies and T&Cs in languages other than English.

**Approach**: Compile-time per-locale builds (one HTML file per language). This aligns with the existing static-site architecture (no bundler, no hot-reload) and mirrors the existing pattern of Pug compile-time variables (`noTracking`, `gdpr`).

**Key constraint**: Legal templates (~670 lines across 4 files) need native-speaking legal review per language — this is the bottleneck, not the code.

---

## Todo List

### Phase 1 — Infrastructure (`[ ]` = not started, `[x]` = done, `[-]` = skipped)

- [ ] **1.1 Create locale JSON files**
      - `src/locales/en.json` — extract all translatable strings from templates + JS
      - `src/locales/es.json` — Spanish translations
      - One file per language; keys mirror template structure for easy lookup
- [ ] **1.2 Create `localeMixin.js`**
      - Reactively exposes `$t(key)` method
      - Loads locale JSON (sync, inlined at compile time)
      - Handles locale-aware formatting (date formats, pluralization, article logic)
- [ ] **1.3 Update `src/index.pug`**
      - Add compile-time `lang` variable
      - Conditionally load correct locale data and template includes based on `lang`
      - Update `html(lang=...)` attribute dynamically
- [ ] **1.4 Refactor `platformMixin.js`**
      - Move platform vocabulary strings (`deviceType`, `osDesc`, etc.) into locale files
      - Compute them from locale data instead of hardcoding
- [ ] **1.5 Replace hardcoded strings in wizard templates**
      - `step_2.pug` through `step_8.pug` — replace labels, placeholders, help text with `{{ $t('key') }}`
      - `step_1.pug` — welcome text, share text
- [ ] **1.6 Replace hardcoded strings in modal templates**
      - Modal titles, button labels ("HTML", "Markdown", "Preview", "Deploy")
      - `disclaimer.pug`, `faq.pug` content
- [ ] **1.7 Replace hardcoded strings in meta tags**
      - `meta/app-description.pug`, `meta/link-preview.pug`, `meta/author.pug`
- [ ] **1.8 Update `render.sh` to loop over locales**
      ```bash
      for lang in en es fr de pt ja zh; do
        npx pug3 src/index.pug --out public/$lang -O "{ lang: '$lang' }"
      done
      ```
      - Default (en) goes to `public/index.html`
      - Other locales go to `public/{lang}/index.html`
- [ ] **1.9 Handle third-party service names**
      - Add locale-specific `name` fields in YAML or create per-locale YAML files
- [ ] **1.10 Add language switcher UI**
      - Simple dropdown in the wizard header/footer
      - Navigates to `/{lang}/` path
- [ ] **1.11 Update `package.json` scripts**
      - `npm run build` — runs full multi-locale build
      - `npm run build:en` — single locale for dev

### Phase 2 — Translate Legal Templates

- [ ] **2.1 `privacy_policy/simple.pug`** — English → Spanish (or target language)
- [ ] **2.2 `privacy_policy/gdpr.pug`** — English → Spanish (or target language)
- [ ] **2.3 `privacy_policy/no_tracking.pug`** — English → Spanish (or target language)
- [ ] **2.4 `tnc/simple.pug`** — English → Spanish (or target language)
- [ ] **2.5 Legal review** — Native-speaking legal professional reviews translations
- [ ] **2.6 Repeat 2.1–2.5 for each additional language**

### Phase 3 — Polish & Verify

- [ ] **3.1 Verify all `v-if` / `v-for` conditionals still work across languages**
- [ ] **3.2 Verify `npm run build` completes cleanly for all locales**
- [ ] **3.3 Test language switcher navigates correctly**
- [ ] **3.4 Verify HTML/Markdown export works per locale**
- [ ] **3.5 Verify no HTML/CSS specificity issues when `lang` attr changes**
- [ ] **3.6 Update `AGENTS.md` with new file structure and conventions**

---

## Design Decisions

| Decision | Choice | Rationale |
|----------|--------|-----------|
| Compile-time vs runtime | **Compile-time** | Matches existing architecture; no bundler; better SEO; simpler |
| Locale format | **Flat JSON key-value** | Simplest lookup, no nested traversal needed |
| Locale key naming | **`section.component.string`** | Mirror template structure for grep-ability |
| Default locale | **`en` → `public/index.html`** | Backward-compatible URL; other locales at `/{lang}/` |
| Language switcher impl | **Simple `<select>` → `window.location`** | No SPA router; full page reload is acceptable |
| Third-party names | **YAML `name_es`, `name_fr` fields** | Keeps single source of truth for services |
| Legal template structure | **Same Pug structure, translated text only** | `v-if`/`v-for` conditionals stay identical per locale |

---

## File Changes Summary

```
NEW  src/locales/en.json          — English strings
NEW  src/locales/es.json          — Spanish strings
NEW  src/locales/fr.json          — French strings
NEW  src/js/localeMixin.js        — Translation mixin
MOD  src/index.pug                — Add lang var, conditional includes
MOD  src/js/platformMixin.js      — Move strings to locale files
MOD  src/js/main.js               — Register localeMixin
MOD  src/includes/content/wizard/step_1.pug  — Replace hardcoded strings
MOD  src/includes/content/wizard/step_2.pug  — Replace hardcoded strings
MOD  src/includes/content/wizard/step_3.pug  — Replace hardcoded strings
MOD  src/includes/content/wizard/step_4.pug  — Replace hardcoded strings
MOD  src/includes/content/wizard/step_5.pug  — Replace hardcoded strings
MOD  src/includes/content/wizard/step_6.pug  — Replace hardcoded strings
MOD  src/includes/content/wizard/step_7.pug  — Replace hardcoded strings
MOD  src/includes/content/wizard/step_8.pug  — Replace hardcoded strings
MOD  src/includes/content/faq.pug            — Replace hardcoded strings
MOD  src/includes/content/disclaimer.pug     — Replace hardcoded strings
MOD  src/includes/content/privacy_policy/simple.pug     — Template stays, strings via locale
MOD  src/includes/content/privacy_policy/gdpr.pug       — Template stays, strings via locale
MOD  src/includes/content/privacy_policy/no_tracking.pug — Template stays, strings via locale
MOD  src/includes/content/tnc/simple.pug     — Template stays, strings via locale
MOD  src/includes/meta/*.pug                 — Replace hardcoded strings
MOD  src/includes/yaml/thirdpartyservices.yml — Add locale-specific name fields
MOD  render.sh                — Loop over locales
MOD  package.json             — Add per-locale scripts
MOD  AGENTS.md                — Document new conventions
```
