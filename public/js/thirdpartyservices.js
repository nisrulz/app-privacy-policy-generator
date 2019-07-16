/*  Copyright 2017 Nishant Srivastava

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License. */

var thirdPartyServicesJsonArray = [
  {
    name: 'Google Play Services',
    model: 'gps', // this is stores the name of the model
    gps: true, // this stores the state of the model, notice the key is the same as the model above. It is required by code.
    logo: 'images/gps.png',
    link: {
      privacy: 'https://www.google.com/policies/privacy/',
      terms: ''
    }
  },
  {
    name: 'AdMob',
    model: 'admob',
    admob: false,
    logo: 'images/admob.png',
    link: {
      privacy: 'https://support.google.com/admob/answer/6128543?hl=en',
      terms: ''
    }
  },
  {
    name: 'Firebase Analytics',
    model: 'firebase',
    firebase: false,
    logo: 'images/firebase.png',
    link: {
      privacy: 'https://firebase.google.com/policies/analytics',
      terms: ''
    }
  },
  {
    name: 'Facebook',
    model: 'facebook',
    facebook: false,
    logo: 'images/facebook.png',
    link: {
      privacy: 'https://www.facebook.com/about/privacy',
      terms: ''
    }
  },
  {
    name: 'Fabric',
    model: 'fabric',
    fabric: false,
    logo: 'images/fabric.png',
    link: {
      privacy: 'https://fabric.io/privacy',
      terms: ''
    }
  },
  {
    name: 'Crashlytics',
    model: 'crashlytics',
    crashlytics: false,
    logo: 'images/crashlytics.png',
    link: {
      privacy: 'http://try.crashlytics.com/terms/privacy-policy.pdf',
      terms: ''
    }
  },
  {
    name: 'Matomo',
    model: 'piwik',
    piwik: false,
    logo: 'images/piwik.png',
    link: {
      privacy: 'https://matomo.org/privacy-policy/',
      terms: ''
    }
  },
  {
    name: 'Clicky',
    model: 'clicky',
    clicky: false,
    logo: 'images/clicky.png',
    link: {
      privacy: 'https://clicky.com/terms#privacy',
      terms: ''
    }
  },
  {
    name: 'Flurry Analytics',
    model: 'flurry',
    flurry: false,
    logo: 'images/flurry.png',
    link: {
      privacy: 'https://privacy.oath.com',
      terms:
        'https://developer.yahoo.com/flurry/legal-privacy/terms-service/flurry-analytics-terms-service.html'
    }
  },
  {
    name: 'Appodeal',
    model: 'appodeal',
    appodeal: false,
    logo: 'images/appodeal.png',
    link: {
      privacy: 'https://www.appodeal.com/home/privacy-policy/',
      terms: 'https://www.appodeal.com/home/terms-of-service/'
    }
  },
  {
    name: 'Fathom Analytics',
    model: 'fathom', // this is stores the name of the model
    fathom: false, // this stores the state of the model, notice the key is the same as the model above. It is required by code.
    logo: 'images/fathom.png',
    link: {
      privacy: 'https://usefathom.com/privacy/',
      terms: 'https://usefathom.com/terms/'
    }
  },
  {
      name: 'Unity',
      model: 'unity',
      unity: false,
      logo: 'images/unity.png',
      link: {
        privacy: 'https://unity3d.com/legal/privacy-policy',
        terms: 'https://unity3d.com/legal/terms-of-service'
      }
  },
  {
	  name: 'GameAnalytics',
	  model: 'gameanalytics',
	  gameanalytics: false,
	  logo: 'images/gameanalytics.png',
	  links: {
		  privacy: 'https://gameanalytics.com/privacy',
		  terms: 'https://gameanalytics.com/terms'
	  }
  } // Do not forget to add a ',' before adding your new item in the array
]
