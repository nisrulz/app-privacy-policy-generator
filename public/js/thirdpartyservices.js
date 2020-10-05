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
    name: "Google Play Services",
    model: "gps",
    gps: true,
    logo: "images/gps.png",
    link: {
      privacy: "https://www.google.com/policies/privacy/",
      terms: "https://policies.google.com/terms"
    }
  },
  {
    name: "AdMob",
    model: "admob",
    admob: false,
    logo: "images/admob.png",
    link: {
      privacy: "https://support.google.com/admob/answer/6128543?hl=en",
      terms: "https://developers.google.com/admob/terms"
    }
  },
  {
    name: "Google Analytics for Firebase",
    model: "firebase",
    firebase: false,
    logo: "images/firebase.png",
    link: {
      privacy: "https://firebase.google.com/policies/analytics",
      terms: "https://firebase.google.com/terms/analytics"
    }
  },
  {
    name: "Firebase Crashlytics",
    model: "crashlytics",
    crashlytics: false,
    logo: "images/firebase.png",
    link: {
      privacy: "https://firebase.google.com/support/privacy/",
      terms: "https://firebase.google.com/terms/crashlytics"
    }
  },
  {
    name: "Facebook",
    model: "facebook",
    facebook: false,
    logo: "images/facebook.png",
    link: {
      privacy: "https://www.facebook.com/about/privacy/update/printable",
      terms: "https://www.facebook.com/legal/terms/plain_text_terms"
    }
  },
  {
    name: "Fabric",
    model: "fabric",
    fabric: false,
    logo: "images/fabric.png",
    link: {
      privacy: "https://policies.google.com/privacy",
      terms: "https://fabric.io/terms/fabric"
    }
  },
  {
    name: "Matomo",
    model: "piwik",
    piwik: false,
    logo: "images/piwik.png",
    link: {
      privacy: "https://matomo.org/privacy-policy/",
      terms: "https://matomo.org/terms"
    }
  },
  {
    name: "Clicky",
    model: "clicky",
    clicky: false,
    logo: "images/clicky.png",
    link: {
      privacy: "https://clicky.com/terms#privacy",
      terms: "https://clicky.com/terms/"
    }
  },
  {
    name: "Flurry Analytics",
    model: "flurry",
    flurry: false,
    logo: "images/flurry.png",
    link: {
      privacy: "https://privacy.oath.com",
      terms:
        "https://developer.yahoo.com/flurry/legal-privacy/terms-service/flurry-analytics-terms-service.html"
    }
  },
  {
    name: "Appodeal",
    model: "appodeal",
    appodeal: false,
    logo: "images/appodeal.png",
    link: {
      privacy: "https://www.appodeal.com/home/privacy-policy/",
      terms: "https://www.appodeal.com/home/terms-of-service/"
    }
  },
  {
    name: "Fathom Analytics",
    model: "fathom",
    fathom: false,
    logo: "images/fathom.png",
    link: {
      privacy: "https://usefathom.com/privacy/",
      terms: "https://usefathom.com/terms/"
    }
  },
  {
    name: "Unity",
    model: "unity",
    unity: false,
    logo: "images/unity.png",
    link: {
      privacy: "https://unity3d.com/legal/privacy-policy",
      terms: "https://unity3d.com/legal/terms-of-service"
    }
  },
  {
    name: "SDKBOX",
    model: "sdkbox",
    sdkbox: false,
    logo: "images/sdkbox.png",
    link: {
      privacy: "https://www.sdkbox.com/privacy",
      terms: "https://www.sdkbox.com/privacy"
    }
  },
  {
    name: "GameAnalytics",
    model: "gameanalytics",
    gameanalytics: false,
    logo: "images/gameanalytics.png",
    link: {
      privacy: "https://gameanalytics.com/privacy",
      terms: "https://gameanalytics.com/terms"
    }
  },
  {
    name: "One Signal",
    model: "onesignal",
    onesignal: false,
    logo: "images/one_signal.svg",
    link: {
      privacy: "https://onesignal.com/privacy_policy",
      terms: "https://onesignal.com/tos"
    }
  },
  {
    name: "Expo",
    model: "expo",
    expo: false,
    logo: "images/expo.png",
    link: {
      privacy: "https://expo.io/privacy",
      terms: "https://expo.io/terms"
    }
  },
  {
    name: "Sentry",
    model: "sentry",
    expo: false,
    logo: "images/sentry.png",
    link: {
      privacy: "https://sentry.io/privacy/",
      terms: "https://sentry.io/terms/"
    }
  },
  {
    name: "AppLovin",
    model: "applovin",
    expo: false,
    logo: "images/applovin.png",
    link: {
      privacy: "https://www.applovin.com/privacy/",
      terms: "https://www.applovin.com/terms/"
    }
  },
  {
    name: "Vungle",
    model: "vungle",
    expo: false,
    logo: "images/vungle.png",
    link: {
      privacy: "https://vungle.com/privacy/"
    }
  },
  {
    name: "StartApp",
    model: "startapp",
    expo: false,
    logo: "images/startapp.png",
    link: {
      privacy: "https://www.startapp.com/privacy/",
      terms: "https://www.startapp.com/policy/publisher-terms/"
    }
  },
  {
    name: "AdColony",
    model: "adcolony",
    expo: false,
    logo: "images/adcolony.png",
    link: {
      privacy: "https://www.adcolony.com/privacy-policy/"
    }
  },
  {
    name: "Amplitude",
    model: "amplitude",
    expo: false,
    logo: "images/amplitude.svg",
    link: {
      privacy: "https://amplitude.com/privacy",
      terms: "https://amplitude.com/terms"
    }
  },
  {
    name: "Adjust",
    model: "adjust",
    expo: false,
    logo: "images/adjust.png",
    link: {
      privacy: "https://www.adjust.com/terms/privacy-policy",
      terms: "https://www.adjust.com/terms/general-terms-and-conditions"
    }
  },
  {
    name: "Mapbox",
    model: "mapbox",
    mapbox: false,
    logo: "images/mapbox.png",
    link: {
      privacy: "https://www.mapbox.com/legal/privacy",
      terms: "https://www.mapbox.com/legal/tos"
    }
  },
  {
    name: "Godot",
    model: "godot",
    godot: false,
    logo: "images/godot.png",
    link: {
      privacy: "https://godotengine.org/privacy-policy",
      terms: "https://godotengine.org/license"
    }
  }
];
