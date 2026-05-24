/*  
  App Privacy Policy Generator: A simple web app to generate a generic 
  privacy policy for your Android, iOS, and Web apps

  Copyright 2017-Present Nishant Srivastava

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
*/

window.platformMixin = {
  data() {
    return {
      // === Platform-Aware Text ===
      // These defaults assume Mobile App; generate() overrides for Web/Both
      platformType: "Mobile App",
      deviceType: "mobile device",
      deviceTypePlural: "mobile devices",
      platformDesc: "mobile devices",
      deviceIdDesc: "your mobile device's unique device ID",
      osDesc: "your mobile operating system",
      browserDesc: "the type of mobile Internet browsers you use",
      uninstallDesc: "by uninstalling the Application",
      osType: "Android",
      requirementOfSystem: "system",
    };
  },

  computed: {
    // Platform type matchers used across all templates
    isMobileApp() {
      return this.platformType === "Mobile App";
    },
    isWebApp() {
      return this.platformType === "Web App";
    },
    isBothPlatforms() {
      return this.platformType === "Both";
    },
  },

  methods: {
    // Map selected OS combo to the system requirement phrase used in T&C
    _setOsRequirement() {
      switch (this.osType) {
        case "Android & iOS":
          this.requirementOfSystem = "both systems";
          break;
        default:
          this.requirementOfSystem = "system";
      }
    },

    // Derive device/language vocabulary from the platform type selection.
    // All templates reference these variables so the text reads naturally
    // whether the app targets mobile, web, or both.
    _setPlatformText() {
      switch (this.platformType) {
        case "Web App":
          this.deviceType = "device";
          this.deviceTypePlural = "devices";
          this.platformDesc = "web browsers";
          this.deviceIdDesc = "your device's unique identifier (e.g., IP address or browser fingerprint)";
          this.osDesc = "your operating system";
          this.browserDesc = "the type of web browser you use";
          this.uninstallDesc = "by ceasing to use the website";
          break;
        case "Both":
          this.deviceType = "mobile device or computer";
          this.deviceTypePlural = "mobile devices and computers";
          this.platformDesc = "mobile devices and web browsers";
          this.deviceIdDesc = "your device's unique device ID or identifier";
          this.osDesc = "your operating system";
          this.browserDesc = "the type of browser you use";
          this.uninstallDesc = "by uninstalling the Application or ceasing to use the website";
          break;
        default:
          this.deviceType = "mobile device";
          this.deviceTypePlural = "mobile devices";
          this.platformDesc = "mobile devices";
          this.deviceIdDesc = "your mobile device's unique device ID";
          this.osDesc = "your mobile operating system";
          this.browserDesc = "the type of mobile Internet browsers you use";
          this.uninstallDesc = "by uninstalling the Application";
          break;
      }
    },
  }
};
