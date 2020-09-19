# App Privacy Policy Generator

A simple web app to generate a generic privacy policy for your Android/iOS apps

Check out the [web app](https://app-privacy-policy-generator.nisrulz.com/)!

### Featured in

[![AndroidDev Digest](https://img.shields.io/badge/AndroidDev%20Digest-%23133-blue.svg)](https://www.androiddevdigest.com/digest-133/)

### Show some :heart:

[![GitHub stars](https://img.shields.io/github/stars/nisrulz/app-privacy-policy-generator.svg?style=social&label=Star)](https://github.com/nisrulz/app-privacy-policy-generator) [![GitHub forks](https://img.shields.io/github/forks/nisrulz/app-privacy-policy-generator.svg?style=social&label=Fork)](https://github.com/nisrulz/app-privacy-policy-generator/fork) [![GitHub watchers](https://img.shields.io/github/watchers/nisrulz/app-privacy-policy-generator.svg?style=social&label=Watch)](https://github.com/nisrulz/app-privacy-policy-generator) [![GitHub followers](https://img.shields.io/github/followers/nisrulz.svg?style=social&label=Follow)](https://github.com/nisrulz/app-privacy-policy-generator)  
[![Twitter Follow](https://img.shields.io/twitter/follow/nisrulz.svg?style=social)](https://twitter.com/nisrulz)

## Web App Screenshot

![screenshot](/img/sc1.png)

![screenshot](/img/sc2.png)

# Contributing more 3rd Party Service's Privacy links

The webapp uses a JSON array to populate the list of 3rd party services section. If you want contribute a new 3rd party service for the section then open a Pull Request which simply adds a new JSON item to the JSON Array in [`public/js/thirdpartyservices.js`](public/js/thirdpartyservices.js)

The format is very simple and you only need to provide fill the fields as follows:

```
[
  ..
  }, // Donot forget to add a ',' before adding your new item in the array
  {
    name: "Google Play Services",
    model: "gps", // this is stores the name of the model
    gps: true, // this stores the state of the model, notice the key is the same as the model above. It is required by code.
    logo: "images/gps.png",
    link: {
      privacy: "https://www.google.com/policies/privacy/",
      terms: ""
    }
  }
]
```

> Note: You also will have to drop the logo image into the [`public/images/`](public/images/) folder. Make sure the logo is 160 × 160 in size.

# Pull Requests

I welcome and encourage all pull requests. It usually will take me within 24-48 hours to respond to any issue or request. Here are some basic rules to follow to ensure timely addition of your request:

1. If its a feature, bugfix, or anything please only change code to what you specify.
1. Please keep PR titles easy to read and descriptive of changes, this will make them easier to merge :)
1. Pull requests _must_ be made against `develop` branch. Any other branch (unless specified by the maintainers) will get rejected.
1. Check for existing [issues](https://github.com/nisrulz/app-privacy-policy-generator/issues) first, before filing an issue.
1. Have fun!

## Author & Support

This project was created by [Nishant Srivastava](https://github.com/nisrulz/nisrulz.github.io#nishant-srivastava) but hopefully developed and maintained by many others. See the [the list of contributors here](https://github.com/nisrulz/app-privacy-policy-generator/graphs/contributors).

If you appreciate my work, consider [buying me](https://www.paypal.me/nisrulz/5usd) a cup of :coffee: to keep me recharged :muscle: [[PayPal](https://www.paypal.me/nisrulz/5usd)]. A comment in the project's [GuestBook](https://github.com/nisrulz/app-privacy-policy-generator/issues/65) is appreciated :blush:

# License

    Copyright 2017 Nishant Srivastava

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
