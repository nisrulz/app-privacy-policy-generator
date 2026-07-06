# Localization Guide

This project supports **compile-time localization**. Each language produces its own `index.html` file. Translators only need to edit one JSON file per language — **no code changes required**.

---

## Quick Start: Adding a New Language

### 1. Create a locale file

Copy `src/locales/en.json` to `src/locales/{code}.json` where `{code}` is your language code (e.g. `es`, `fr`, `de`, `ja`, `pt`).

```sh
cp src/locales/en.json src/locales/es.json
```

### 2. Translate the strings

Edit `src/locales/es.json` and replace all English values with translations. **Do not change the keys** — only translate the values on the right side of the `:`.

Example:
```json
// ✅ Correct — translate only the value
"wizard.step1.tagline": "Genera un documento personalizado de Política de Privacidad y Términos de Uso para tus aplicaciones",

// ❌ Wrong — do not change the key
"wizard.step1.tagline": "Genera un documento personalizado...",  // ← key is correct
"wizard.step1.tagline_es": "Genera un documento personalizado...", // ← WRONG, new key won't be found
```

### 3. Build

```sh
npm run build
```

The Go build tool (`cmd/build/`) automatically detects all JSON files in `src/locales/` and generates per-locale HTML + JS. No manual registration needed.

Open `public/index.html` and verify your translations appear correctly.

---

## Locale File Structure

All translatable strings are in `src/locales/{lang}.json`. Keys use dot-separated namespaces:

| Prefix | Content |
|--------|---------|
| `wizard.stepN.*` | Wizard step labels, placeholders, help text |
| `simple.*` | Simple Privacy Policy body text |
| `gdpr.*` | GDPR Privacy Policy body text |
| `noTracking.*` | No Tracking Privacy Policy body text |
| `tnc.*` | Terms & Conditions body text |
| `faq.*` | FAQ questions and answers |
| `disclaimer.*` | Disclaimer modal text |
| `platform.*` | Platform vocabulary (device types, OS, browser) |
| `modal.*` | Modal titles, buttons |
| `meta.*` | Page title, meta descriptions |
| `misc.*` | Miscellaneous utility strings |

### String interpolation

Some strings contain `{{ variableName }}` placeholders. These are replaced with user-provided values at runtime. **Translate the surrounding text, but keep the `{{ variableName }}` placeholders unchanged.**

Example:
```json
"simple.intro": "This privacy policy applies to the {{ appName }} app for {{ platformDesc }}, together with any related services operated by {{ devOrCompanyName }}."
```

In Spanish:
```json
"simple.intro": "Esta política de privacidad se aplica a la aplicación {{ appName }} para {{ platformDesc }}, junto con cualquier servicio relacionado operado por {{ devOrCompanyName }}."
```

---

## HTML Formatting in Strings

Some locale values contain HTML in the source JSON (e.g., `<strong>` tags, links). These are rendered using `v-html` in Vue. When translating, you can use HTML markup if needed, but in most cases plain text is sufficient.

---

## Testing Your Translation

1. Run `npm run build`
2. Open `public/index.html` in your browser
3. Go through the wizard steps and verify all UI text is translated
4. Generate a privacy policy and verify the legal text
5. Test the HTML and Markdown export buttons

---

## Third-Party Service Names

Third-party service names in `src/includes/yaml/thirdpartyservices.yml` can also be translated. Add locale-specific `name_{code}` fields to any entry:

```yaml
- name: Google Play Services
  name_es: Servicios de Google Play
  name_fr: Services Google Play
  enabled: false
  logo: images/third_party_logos/gps.png
  link:
    privacy: https://www.google.com/policies/privacy/
    terms: https://policies.google.com/terms
```

The app automatically uses the locale-specific name when available, falling back to the default `name` field.

---

## Important Notes

- **Legal templates require professional review**: Privacy policies and Terms & Conditions are legal documents. If you translate them, have your translation reviewed by a native-speaking legal professional before using it.
- **Keep structure identical**: The `v-if`, `v-for`, and other Vue directives in the Pug templates are language-independent. They will work the same for all languages.
- **Do not edit the Pug templates**: All translatable content is in the JSON locale files. The Pug templates only contain structural markup and `translate('key')` calls.
