# Development

The web app uses

- [VueJs](https://vuejs.org/) - For templating and reactive updates in the DOM
- [Pug-CLI](https://github.com/pugjs/pug-cli) - To convert pug templates to html and merge partials into single [`index.html`](public/index.html) file.
- [SASS-CLI](https://sass-lang.com/documentation/cli) - To convert sass templates to css and merge partials into single [`style.css`](public/css/style.css) file.
- [js-yaml-CLI](https://github.com/nodeca/js-yaml#cli-executable) - To convert yaml templates to json and generate the [`thirdpartyservices.js`](public/js/thirdpartyservices.js) file.
- [Firebase Hosting](https://firebase.google.com/docs/hosting/) - For Hosting

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

This will generate the `index.html`, `style.css` and `thirdpartyservices.js` files at their required directory path.

Open `index.html` to view the full webapp 🎉

## Contributing more 3rd Party Service's links

The webapp uses a JSON array to populate the list of 3rd party services section. This JSON array is generated from another yaml file.

If you want contribute a new 3rd party service for the section then open a Pull Request which simply adds a new item in the YAML file [`src/includes/yaml/thirdpartyservices.yml`](src/includes/yaml/thirdpartyservices.yml)

The format is very simple and you only need to provide the below fields:

```yml
- name: Google Play Services
  model: gps
  gps: false
  logo: images/third_party_logos/gps.png
  link:
    privacy: https://www.google.com/policies/privacy/
    terms: https://policies.google.com/terms
```

> Note: The value of model is the same as the the field below it. i.e `model: gps` and `gps: false`. This is required by the code. Make sure this is maintained.

Additionally, you also will have to drop the logo image into the [`public/images/third_party_logos/`](public/images/third_party_logos/) folder. Make sure the logo is **160 × 160** in size.

### Deployment

In order to deploy with the web app and Firebase hosting you would need to setup Firebase CLI.

### Setup Firebase CLI

The instructions to set it up differ based on the operating system, which is explained [here](https://firebase.google.com/docs/cli).

However you can also use the below command to setup/update the Firebase CLI, which will take care of everything (including detecting the OS you are running):

```bash
curl -sL firebase.tools | bash
```

### Server and test locally

To serve locally for testing, run: `firebase serve --only hosting`

### Deploy to Production

> *Warning*
> This can only be done by maintainers who have access to Firebase console!

- First, login by executing command: `firebase login`
- Next,To deploy to production,run: `firebase deploy -m "3.0.9"`
