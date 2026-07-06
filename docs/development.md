# Development

The web app is built with:

- [Vue.js](https://vuejs.org/) for templating and reactive updates
- [Firebase Hosting](https://firebase.google.com/docs/hosting/) for hosting
- Go toolchain (`go run ./cmd/build/`) that orchestrates Less to CSS, YAML to JS, Pug to HTML (via `npx pug3`), JS minification, and cache-busting
- [pug3](https://github.com/tokilabs/pug3-cli) to convert Pug templates to HTML (the only npm `devDependency`)
- [firebase-tools](https://github.com/firebase/firebase-tools) for Firebase CLI (install globally)

---

Install build tools:

```sh
npm install
```

This gives you `pug3`, the only npm dependency. The Go build toolchain uses it via `npx`.

For image compression and deployment, install these globally:

```sh
npm install -g firebase-tools svgo png-minify
```

## Source layout

All the code you care about lives in [`src`](src):

- Privacy Policy templates: [`src/includes/content/privacy_policy`](src/includes/content/privacy_policy)
- Terms & Conditions templates: [`src/includes/content/tnc`](src/includes/content/tnc)

The app is made of Pug partials, Less stylesheets, and a YAML file for third-party service definitions. The Go build compiles everything into a single `index.html` under [`public`](public), plus minified CSS and JS files.

To build:

```sh
❯ make build
```

This generates `index.html`, `style.min.css`, `main.min.js`, `utils.min.js`, `thirdpartyservices.min.js`, and locale files.

Other useful commands:

```sh
❯ make watch            # watch for changes and rebuild automatically
❯ make reviews          # regenerate reviews page from cached data
❯ make reviews-force    # regenerate reviews page (re-fetch from GitHub)
```

You can also pass flags directly to the Go build tool:

```sh
❯ go run ./cmd/build/ -lang de        # build only a specific locale
❯ go run ./cmd/build/ -clean          # clean public/ before building
```

## Compress images

```sh
❯ ./scripts/compress_images.sh
```

## Adding a 3rd party service

The third-party services list is generated from a YAML file. To add one, open a PR that adds an entry to [`src/includes/yaml/thirdpartyservices.yml`](src/includes/yaml/thirdpartyservices.yml):

```yml
- name: Google Play Services
  enabled: false
  logo: images/third_party_logos/gps.png
  link:
    privacy: https://www.google.com/policies/privacy/
    terms: https://policies.google.com/terms
```

Drop the logo into [`public/images/third_party_logos/`](public/images/third_party_logos/). Make it 160x160 pixels.

Tip: Use [remove.bg](https://www.remove.bg/) to strip the background and [imagetools.org/trim](https://www.imagetools.org/trim) to trim excess space.

## Updating packages

```sh
npm update
```

## Serving locally

Builds and serves on port 8000:

```sh
❯ make serve
```

Use a custom port with:

```sh
❯ go run ./cmd/build/ -serve -port 9090
```

## Testing

E2E tests run against the dev server using Playwright (Chromium only):

```sh
❯ make test          # run all tests
❯ make test-ui       # run with Playwright UI mode
❯ make test-debug    # run with Playwright debug mode
```

Tests are in `tests/`. The dev server starts automatically when you run tests.

## Deploying to production

Note: Only maintainers with Firebase console access can deploy.

```sh
firebase login
❯ make firebase-deploy VERSION="3.0.9"
```

Omit `VERSION` and you will be prompted for it.

## Tools

- [Reviews Page Generator](tools/reviews-page-generator/readme.md)
