# Development

The web app uses 
- [VueJs](https://vuejs.org/) - For templating and reactive updates in the DOM
- [PUG-CLI](https://github.com/pugjs/pug-cli) - To convert pug templates to html and merge partials into single [`index.html`](public/index.html) file.
- [SASS-CLI](https://sass-lang.com/documentation/cli) - To convert sass templates to css and merge partials into single [`style.css`](public/css/style.css) file.
- [Firebase Hosting](https://firebase.google.com/docs/hosting/) - For Hosting

If you wish to modify the code for the webapp, then look into [`src`](src) directory.
- Privacy Policy templates are under [`src/includes/content/privacy_policy`](src/includes/content/privacy_policy)
- Terms & Conditions templates are under [`src/includes/content/tnc`](src/includes/content/tnc)

The webapp is setup in a way that it is made up of 
- html partials written in pug templating language
- css partials written in sass templating language

...all of which is compiled into a single `index.html` file which lives under [`public`](public) directory. Styles are compiled into a single `style.css` file under [`public/css`](public/css) directory, referenced directly in the `index.html` file.

To compile the code under `src` folder, simply execute the helper bash script [`render.sh`](render.sh) at the root of the repo inside a terminal:
```sh
./render.sh
```

This will generate the `index.html` and `style.css` file at their required directory path.

Open `index.html` to view the full webapp 🎉

## Contributing more 3rd Party Service's Privacy links

The webapp uses a JSON array to populate the list of 3rd party services section. If you want contribute a new 3rd party service for the section then open a Pull Request which simply adds a new JSON item to the JSON Array in [`public/js/thirdpartyservices.js`](public/js/thirdpartyservices.js)

The format is very simple and you only need to provide fill the fields as follows:

```
[
  ..
  }, // Donot forget to add a ',' before adding your new item in the array
  {
    name: "Google Play Services",
    model: "gps", // this stores the name of the model
    gps: true, // this stores the state of the model, NOTICE: Key is same as the model name above. It is required by code.
    logo: "images/gps.png",
    link: {
      privacy: "https://www.google.com/policies/privacy/",
      terms: ""
    }
  }
]
```

> Note: You also will have to drop the logo image into the [`public/images/`](public/images/) folder. Make sure the logo is **160 × 160** in size.

# Deployment

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

> Note: This can only be done by maintainers who have access to Firebase console!

- First, login by executing command: `firebase login`
- Next,To deploy to production,run: `firebase deploy -m "3.0.9"`
