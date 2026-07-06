# Development

The web app uses

- [VueJs](https://vuejs.org/) - For templating and reactive updates in the DOM
- [Firebase Hosting](https://firebase.google.com/docs/hosting/) - For hosting the static web app
- Go toolchain (`go run ./cmd/build/`) — orchestrates Less → CSS compilation, YAML → JS conversion, Pug → HTML rendering (via `npx pug3`), JS minification, and cache-busting
- [pug3](https://github.com/tokilabs/pug3-cli) - To convert pug templates to html and merge partials (only remaining npm `devDependency`)
- [firebase-tools](https://github.com/firebase/firebase-tools) - To interact with Firebase as a service from command line (install globally)

---

Install build tools:

```sh
npm install
```

This installs `pug3` (the only npm dependency — used by the Go build toolchain via `npx`).

Additional tools (image compression, deployment) require global install:

```sh
npm install -g firebase-tools svgo png-minify
```

If you wish to modify the code for the webapp, then look into the [`src`](src) directory.

- Privacy Policy templates are under [`src/includes/content/privacy_policy`](src/includes/content/privacy_policy)
- Terms & Conditions templates are under [`src/includes/content/tnc`](src/includes/content/tnc)

The webapp is setup in a way that it is made up of

- html partials written in pug templating language
- css partials written in Less templating language
- js config for third party services info written in yaml templating language

...all of which is compiled into a single `index.html` file which lives under [`public`](public) directory. Styles are compiled into a single `style.css` file under [`public/css`](public/css) directory and third party services info is compiled into a single `thirdpartyservices.js` under [`public/js`](public/js) referenced directly in the `index.html` file.

To compile the code under `src` folder, run:

```sh
npm run build
```

Or directly invoke the Go build tool:

```sh
make build
```

This will generate the `index.html`, `style.min.css`, `main.min.js`, `utils.min.js`, `thirdpartyservices.min.js`, and locale files at their required directory path.

Open `public/index.html` to view the full web app 🎉

## Compress Images

Run

```sh
./scripts/compress_images.sh
```

## Contributing more 3rd Party Service's links

The web app uses a JSON array to populate the list of 3rd party services section. This JSON array is generated from another yaml file.

If you want contribute a new 3rd party service for the section then open a Pull Request which simply adds a new item in the YAML file [`src/includes/yaml/thirdpartyservices.yml`](src/includes/yaml/thirdpartyservices.yml)

The format is very simple and you only need to provide the below fields:

```yml
- name: Google Play Services
  enabled: false
  logo: images/third_party_logos/gps.png
  link:
    privacy: https://www.google.com/policies/privacy/
    terms: https://policies.google.com/terms
```

Additionally, you also will have to drop the logo image into the [`public/images/third_party_logos/`](public/images/third_party_logos/) folder. Make sure the logo is **160 × 160** in size.

- To remove background from the logo image, [you can use this online tool](https://www.remove.bg/).

- To trim excess space around the logo image, [you can use this online tool](https://www.imagetools.org/trim).

## Update

To update packages in the project, run: `npm update`

## Deployment

### Serve locally

Builds and serves on port 8000:

```sh
make serve
```

Or simply:

```sh
npm run serve
```

This runs the Go build then starts a Go HTTP server serving `public/`. Pass a custom port with:

```sh
go run ./cmd/build/ -serve -port 9090
```

### Deploy to Production

> **Warning**
> This can only be done by maintainers who have access to Firebase console!

- First, login by executing command: `firebase login`
- Next, to deploy to production, run: `firebase deploy -m "3.0.9"`

### Tools

This web app has some of its own tools.

- [Reviews Page Generator Tool](tools/reviews-page-generator/readme.md)
