# Development

The web app uses [VueJs](https://vuejs.org/) and is deployed to [Firebase Hosting](https://firebase.google.com/docs/hosting/).

In order to work with the web app and Firebase hosting you would need to setup Firebase CLI.

### Setup Firebase CLI

The instructions to set it up differ based on the operating system, which is explained [here](https://firebase.google.com/docs/cli).

However you can also use the below command to setup/update the Firebase CLI, which will take care of everything (including detecting the OS you are running):

```bash
curl -sL firebase.tools | bash
```

### Test locally

To serve locally for testing, run: `firebase serve --only hosting`

### Deploy to Production

> Note: This can only be done by maintainers who have access to Firebase console!

- First, login by executing command: `firebase login`
- Next,To deploy to production,run: `firebase deploy -m "3.0.2"`

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
