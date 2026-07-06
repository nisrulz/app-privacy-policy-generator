---
name: add-localization
description: >
  Add or update a language translation in the privacy policy generator.
  Trigger when user says: "add {language} translation", "add locale",
  "translate to {language}", "add localization for {language}", or when
  a user wants to contribute a new language.
---

# Add Localization

## Quick Steps

### 1. Create locale JSON from English

```sh
cp src/locales/en.json src/locales/{code}.json
```

Use the ISO 639-1 language code (e.g., `es`, `fr`, `de`, `ja`, `pt`, `zh`).

### 2. Translate all string values

Edit the new JSON file. **Only translate values, never keys.**

✅ Correct:
```json
"wizard.step1.tagline": "Genera un documento personalizado..."
```

❌ Wrong — don't change keys:
```json
"wizard.step1.tagline_es": "Genera un documento personalizado..."
```

### 3. Preserve `{{ variable }}` placeholders

Keep placeholders like `{{ appName }}`, `{{ platformDesc }}`, `{{ deviceType }}` unchanged in the translated text.

### 4. Build

```sh
make build
```

The Go build tool automatically detects all JSON files in `src/locales/` and generates per-locale HTML + JS.

Open `public/index.html` — the language switcher dropdown now includes the new language.

## Validation Checklist

- [ ] JSON file is valid (no trailing commas, no duplicate keys)
- [ ] All `{{ variable }}` placeholders preserved
- [ ] UI labels, placeholders, help text translated
- [ ] Legal templates translated (requires native-speaking legal review)
- [ ] FAQ and disclaimer content translated
- [ ] Platform vocabulary translated (device types, OS, browser strings)
- [ ] Language appears in the dropdown
- [ ] `npm run build` succeeds

## Important Notes

- **Legal text needs professional review**: Have a native-speaking lawyer review legal template translations.
- **Don't translate third-party service names** in `thirdpartyservices.yml` — only the `name` field in the locale JSON if you add locale-specific names.
- **Meta tags** (`meta.*` keys) are updated dynamically via JS after page load.
