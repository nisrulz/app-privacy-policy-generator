# Development

The web app uses

- [VueJs](https://vuejs.org/) - For templating and reactive updates in the DOM
- [Firebase Hosting](https://firebase.google.com/docs/hosting/) - For hosting the static web app
- CLI Tools
  - [pug](https://github.com/tokilabs/pug3-cli) - To convert pug templates to html and merge partials into single [`index.html`](public/index.html) file.
  - [sass](https://sass-lang.com/documentation/cli) - To convert sass templates to css and merge partials into single [`style.css`](public/css/style.css) file.
  - [js-yaml](https://github.com/nodeca/js-yaml#cli-executable) - To convert yaml templates to json and generate the [`thirdpartyservices.js`](src/js/thirdpartyservices.js) file.
  - [firebase-tools](https://github.com/firebase/firebase-tools) - To interact with Firebase as a service from command line.
  - [png-minify](https://www.npmjs.com/package/png-minify) - To compress png files.
  - [svgo](https://www.npmjs.com/package/svgo) - To compress svg files.
  - [uglifycss](https://www.npmjs.com/package/uglifycss) - To compress css files.
  - [uglify-js](https://www.npmjs.com/package/uglify-js) - To compress js files.
  - [html-minifier](https://www.npmjs.com/package/html-minifier) - To compress html files.
  - [purgecss](https://purgecss.com/CLI.html) - To compress css files.

---

Install the CLI tools using

```sh
npm install -g @tokilabs/pug3-cli sass js-yaml \
firebase-tools \
svgo png-minify \
uglifycss uglify-js html-minifier \
purgecss
```

If you wish to modify the code for the webapp, then look into [`src`](src) directory.

- Privacy Policy templates are under [`src/includes/content/privacy_policy`](src/includes/content/privacy_policy)
- Terms & Conditions templates are under [`src/includes/content/tnc`](src/includes/content/tnc)

The webapp is setup in a way that it is made up of

- html partials written in pug templating language
- css partials written in sass templating language
- js config for third party services info written in yaml templating language

...all of which is compiled into a single `index.html` file which lives under [`public`](public) directory. Styles are compiled into a single `style.css` file under [`public/css`](public/css) directory and third party services info is compiled into a single `thirdpartyservices.js` under [`public/js`](public/js) referenced directly in the `index.html` file.

To compile the code under `src` folder, simply execute the helper bash script [`render.sh`](render.sh) at the root of the repo inside a terminal:

```sh
./render.sh
```

This will generate the `index.html`, `style.min.css`, `main.min.js`, `utils.min.js` and `thirdpartyservices.min.js` files at their required directory path.

Open `index.html` to view the full web app 🎉

## Compress Images

Run

```sh
./compress_images.sh
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

### Server and test locally

To serve locally for testing, run: `firebase serve --only hosting`

### Deploy to Production

> **Warning**
> This can only be done by maintainers who have access to Firebase console!

- First, login by executing command: `firebase login`
- Next,To deploy to production,run: `firebase deploy -m "3.0.9"`

### Tools

This web app has some of its own tools.

- [Reviews Page Generator Tool](tools/reviews-page-generator/readme.md)
