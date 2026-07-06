# Localization Guide

This project does compile-time localization. Each language gets its own `index.html`. You only need to edit one JSON file per language. No code changes required.

---

## Adding a new language

### Step 1: Create a locale file

Copy `src/locales/en.json` to `src/locales/{code}.json` where `{code}` is your language code (e.g. `es`, `fr`, `de`, `ja`, `pt`).

```sh
cp src/locales/en.json src/locales/es.json
```

### Step 2: Translate the strings

Edit the JSON file and replace English values with your translations. Do not change the keys.

```json
// Correct: translate only the value
"wizard.step1.tagline": "Genera un documento personalizado de Política de Privacidad y Términos de Uso para tus aplicaciones",

// Wrong: do not create new keys
"wizard.step1.tagline_es": "...",  // this won't be found
```

### Step 3: Build

```sh
❯ make build
```

The Go build tool picks up all JSON files in `src/locales/` automatically. No registration needed.

To build only your new locale during development:

```sh
❯ go run ./cmd/build/ -lang es
```

Open `public/index.html` and check that your translations show up.

---

## Locale file structure

All strings live in `src/locales/{lang}.json`. Each file starts with a `_flag` key for the locale dropdown emoji. Currently available: English (`en.json`) and German (`de.json`).

Keys use dot-separated namespaces:

| Prefix | What it covers |
|--------|----------------|
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

Some strings have `{{ variableName }}` placeholders. These get replaced at runtime. Translate the surrounding text, but leave the placeholders as-is.

```json
"simple.intro": "This privacy policy applies to the {{ appName }} app for {{ platformDesc }}, together with any related services operated by {{ devOrCompanyName }}."
```

In Spanish:

```json
"simple.intro": "Esta política de privacidad se aplica a la aplicación {{ appName }} para {{ platformDesc }}, junto con cualquier servicio relacionado operado por {{ devOrCompanyName }}."
```

---

## HTML in strings

Some locale values contain HTML (`<strong>` tags, links, etc.). Vue renders these with `v-html`. You can use HTML in your translations too, but in most cases plain text works fine.

---

## Testing your translation

1. Run `❯ make build`
2. Open `public/index.html` in your browser
3. Go through the wizard and check that all UI text is translated
4. Generate a privacy policy and verify the legal text
5. Test the HTML and Markdown export buttons

---

## Translating third-party service names

Service names in [`src/includes/yaml/thirdpartyservices.yml`](src/includes/yaml/thirdpartyservices.yml) can also be translated. Add locale-specific `name_{code}` fields:

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

The app uses the locale-specific name when available and falls back to the default `name` field otherwise.

---

## Important notes

- Privacy policies and Terms & Conditions are legal documents. If you translate them, have a native-speaking legal professional review your translation before using it.
- The Vue directives (`v-if`, `v-for`) in the Pug templates are language-independent. They work the same for all languages.
- All translatable content lives in the JSON locale files. The Pug templates only have structural markup and `translate('key')` calls. Do not edit them.
