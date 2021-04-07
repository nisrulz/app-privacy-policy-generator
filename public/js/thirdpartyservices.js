/*  
    App Privacy Policy Generator: A simple web app to generate a generic 
    privacy policy for your Android/iOS apps

    Copyright 2017-Present Nishant Srivastava

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
 
var thirdPartyServicesJsonArray = [
  {
    "name": "Google Play Services",
    "model": "gps",
    "gps": true,
    "logo": "images/third_party_logos/gps.png",
    "link": {
      "privacy": "https://www.google.com/policies/privacy/",
      "terms": "https://policies.google.com/terms"
    }
  },
  {
    "name": "AdMob",
    "model": "admob",
    "admob": false,
    "logo": "images/third_party_logos/admob.png",
    "link": {
      "privacy": "https://support.google.com/admob/answer/6128543?hl=en",
      "terms": "https://developers.google.com/admob/terms"
    }
  },
  {
    "name": "Google Analytics for Firebase",
    "model": "firebase",
    "firebase": false,
    "logo": "images/third_party_logos/firebase.png",
    "link": {
      "privacy": "https://firebase.google.com/policies/analytics",
      "terms": "https://firebase.google.com/terms/analytics"
    }
  },
  {
    "name": "Firebase Crashlytics",
    "model": "crashlytics",
    "crashlytics": false,
    "logo": "images/third_party_logos/firebase.png",
    "link": {
      "privacy": "https://firebase.google.com/support/privacy/",
      "terms": "https://firebase.google.com/terms/crashlytics"
    }
  },
  {
    "name": "Facebook",
    "model": "facebook",
    "facebook": false,
    "logo": "images/third_party_logos/facebook.png",
    "link": {
      "privacy": "https://www.facebook.com/about/privacy/update/printable",
      "terms": "https://www.facebook.com/legal/terms/plain_text_terms"
    }
  },
  {
    "name": "Fabric",
    "model": "fabric",
    "fabric": false,
    "logo": "images/third_party_logos/fabric.png",
    "link": {
      "privacy": "https://policies.google.com/privacy",
      "terms": "https://fabric.io/terms/fabric"
    }
  },
  {
    "name": "Matomo",
    "model": "piwik",
    "piwik": false,
    "logo": "images/third_party_logos/piwik.png",
    "link": {
      "privacy": "https://matomo.org/privacy-policy/",
      "terms": "https://matomo.org/terms"
    }
  },
  {
    "name": "Clicky",
    "model": "clicky",
    "clicky": false,
    "logo": "images/third_party_logos/clicky.png",
    "link": {
      "privacy": "https://clicky.com/terms#privacy",
      "terms": "https://clicky.com/terms/"
    }
  },
  {
    "name": "Flurry Analytics",
    "model": "flurry",
    "flurry": false,
    "logo": "images/third_party_logos/flurry.png",
    "link": {
      "privacy": "https://privacy.oath.com",
      "terms": "https://developer.yahoo.com/flurry/legal-privacy/terms-service/flurry-analytics-terms-service.html"
    }
  },
  {
    "name": "Appodeal",
    "model": "appodeal",
    "appodeal": false,
    "logo": "images/third_party_logos/appodeal.png",
    "link": {
      "privacy": "https://www.appodeal.com/home/privacy-policy/",
      "terms": "https://www.appodeal.com/home/terms-of-service/"
    }
  },
  {
    "name": "Fathom Analytics",
    "model": "fathom",
    "fathom": false,
    "logo": "images/third_party_logos/fathom.png",
    "link": {
      "privacy": "https://usefathom.com/privacy/",
      "terms": "https://usefathom.com/terms/"
    }
  },
  {
    "name": "Unity",
    "model": "unity",
    "unity": false,
    "logo": "images/third_party_logos/unity.png",
    "link": {
      "privacy": "https://unity3d.com/legal/privacy-policy",
      "terms": "https://unity3d.com/legal/terms-of-service"
    }
  },
  {
    "name": "SDKBOX",
    "model": "sdkbox",
    "sdkbox": false,
    "logo": "images/third_party_logos/sdkbox.png",
    "link": {
      "privacy": "https://www.sdkbox.com/privacy",
      "terms": "https://www.sdkbox.com/privacy"
    }
  },
  {
    "name": "GameAnalytics",
    "model": "gameanalytics",
    "gameanalytics": false,
    "logo": "images/third_party_logos/gameanalytics.png",
    "link": {
      "privacy": "https://gameanalytics.com/privacy",
      "terms": "https://gameanalytics.com/terms"
    }
  },
  {
    "name": "One Signal",
    "model": "onesignal",
    "onesignal": false,
    "logo": "images/third_party_logos/one_signal.svg",
    "link": {
      "privacy": "https://onesignal.com/privacy_policy",
      "terms": "https://onesignal.com/tos"
    }
  },
  {
    "name": "Expo",
    "model": "expo",
    "expo": false,
    "logo": "images/third_party_logos/expo.png",
    "link": {
      "privacy": "https://expo.io/privacy",
      "terms": "https://expo.io/terms"
    }
  },
  {
    "name": "Sentry",
    "model": "sentry",
    "sentry": false,
    "logo": "images/third_party_logos/sentry.png",
    "link": {
      "privacy": "https://sentry.io/privacy/",
      "terms": "https://sentry.io/terms/"
    }
  },
  {
    "name": "AppLovin",
    "model": "applovin",
    "applovin": false,
    "logo": "images/third_party_logos/applovin.png",
    "link": {
      "privacy": "https://www.applovin.com/privacy/",
      "terms": "https://www.applovin.com/terms/"
    }
  },
  {
    "name": "Vungle",
    "model": "vungle",
    "vungle": false,
    "logo": "images/third_party_logos/vungle.png",
    "link": {
      "privacy": "https://vungle.com/privacy/"
    }
  },
  {
    "name": "StartApp",
    "model": "startapp",
    "startapp": false,
    "logo": "images/third_party_logos/startapp.png",
    "link": {
      "privacy": "https://www.startapp.com/privacy/",
      "terms": "https://www.startapp.com/policy/publisher-terms/"
    }
  },
  {
    "name": "AdColony",
    "model": "adcolony",
    "adcolony": false,
    "logo": "images/third_party_logos/adcolony.png",
    "link": {
      "privacy": "https://www.adcolony.com/privacy-policy/"
    }
  },
  {
    "name": "Amplitude",
    "model": "amplitude",
    "amplitude": false,
    "logo": "images/third_party_logos/amplitude.svg",
    "link": {
      "privacy": "https://amplitude.com/privacy",
      "terms": "https://amplitude.com/terms"
    }
  },
  {
    "name": "Adjust",
    "model": "adjust",
    "adjust": false,
    "logo": "images/third_party_logos/adjust.png",
    "link": {
      "privacy": "https://www.adjust.com/terms/privacy-policy",
      "terms": "https://www.adjust.com/terms/general-terms-and-conditions"
    }
  },
  {
    "name": "Mapbox",
    "model": "mapbox",
    "mapbox": false,
    "logo": "images/third_party_logos/mapbox.png",
    "link": {
      "privacy": "https://www.mapbox.com/legal/privacy",
      "terms": "https://www.mapbox.com/legal/tos"
    }
  },
  {
    "name": "Godot",
    "model": "godot",
    "godot": false,
    "logo": "images/third_party_logos/godot.png",
    "link": {
      "privacy": "https://godotengine.org/privacy-policy",
      "terms": "https://godotengine.org/license"
    }
  },
  {
    "name": "Segment",
    "model": "segment",
    "sentry": false,
    "logo": "images/third_party_logos/segment.png",
    "link": {
      "privacy": "https://segment.com/legal/privacy/",
      "terms": "https://segment.com/legal/terms/"
    }
  },
  {
    "name": "Mixpanel",
    "model": "mixpanel",
    "mixpanel": false,
    "logo": "images/third_party_logos/mixpanel.png",
    "link": {
      "privacy": "https://mixpanel.com/legal/privacy-policy/",
      "terms": "https://mixpanel.com/legal/terms-of-use/"
    }
  }
]
