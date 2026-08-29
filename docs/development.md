# Development

The web app uses:

- [Vue.js](https://vuejs.org/) for templating and reactive updates
- [Firebase Hosting](https://firebase.google.com/docs/hosting/) for hosting
- A pure Go toolchain (`go run ./cmd/build/`) that compiles Less to CSS, YAML to JS, and `text/template` files to HTML. It also minifies JavaScript and adds cache-busting values.
- [firebase-tools](https://github.com/firebase/firebase-tools) for Firebase CLI (install globally)

---

For image compression and deployment, install these globally:

```sh
npm install -g firebase-tools svgo png-minify
```

## Source layout

The main application source lives in [`src`](../src). The build and review tools live in [`cmd/build`](../cmd/build) and [`tools/reviews-page-generator`](../tools/reviews-page-generator):

- Page templates (build source of truth): [`src/tpl`](src/tpl) — Go `text/template` files with `{{ }}` delimiters; Vue uses `[[ ]]` delimiters so its syntax passes through unchanged
- Less stylesheets and a YAML file for third-party service definitions

The Go build combines the templates into one `index.html` file under [`public`](public). It also writes minified CSS and JavaScript files.

To build:

```sh
make build
```

This generates `index.html`, `style.min.css`, `main.min.js`, `utils.min.js`, `thirdpartyservices.min.js`, and the locale files.

Other useful commands:

```sh
make format           # format Go source, templates, and tidy modules
make check            # run tests, vet, and build checks
make clean            # clean public/ directory
make compress-images  # compress images in public/images
make watch            # watch for changes and rebuild automatically
make reviews          # generate reviews page from cached data
make update-deps      # update Go dependencies
```

You can also clean before building:

```sh
make clean && make build
```

## Format templates

Format the Go HTML templates using [gotmplfmt](https://github.com/gohugoio/gotmplfmt):

```sh
make format
```

This formats Go source files, HTML templates, and runs `go mod tidy`.

## Compress images

```sh
make compress-images
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

## Serving locally

Builds and serves on port 8000:

```sh
make serve
```

Use a custom port with:

```sh
make serve PORT="9090"
```

## Testing

Format Go source files before you run the checks:

```sh
make format
```

Run the Go checks:

```sh
make check
```

E2E tests run against the dev server using chromedp (pure Go, Chromium only):

```sh
make serve          # start dev server in one terminal
make test           # run E2E tests in another terminal
```

Tests are in `cmd/build/e2e/`. The dev server must be running before you run tests.

## Updating dependencies

```sh
make update-deps
```

This runs `go get -u ./...` and `go mod tidy` to update all Go dependencies.

## Firebase preview

Build and preview the site locally using Firebase Hosting:

```sh
make firebase-preview
```

This runs `make build` then starts `firebase serve --only hosting`.

## Deploying to production

Note: Only maintainers with Firebase console access can deploy.

```sh
firebase login
make deploy VERSION="3.0.9"
```

Omit `VERSION` and you will be prompted for it.

## Tools

- [Reviews Page Generator](tools/reviews-page-generator/readme.md)
