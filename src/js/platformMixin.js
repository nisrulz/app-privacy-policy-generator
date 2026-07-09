window.platformMixin = {
  data() {
    return {
      deviceType: "",
      deviceTypePlural: "",
      platformDesc: "",
      deviceIdDesc: "",
      osDesc: "",
      browserDesc: "",
      uninstallDesc: "",
    };
  },

  computed: {
    selectedPlatformsLabel() {
      var locale = window.__locale || {};
      var labels = [];
      if (this.platforms.Android) labels.push(locale["platform.android"] || "Android");
      if (this.platforms.iOS) labels.push(locale["platform.ios"] || "iOS");
      if (this.platforms.KaiOS) labels.push(locale["platform.kaios"] || "KaiOS");
      if (this.platforms.Windows) labels.push(locale["platform.windows"] || "Windows");
      if (this.platforms.Web) labels.push(locale["platform.web"] || "Web");
      return labels.length ? labels.join(", ") : this.translate("platform.selectedNone");
    },
    isMobileApp() {
      return this.platforms.Android || this.platforms.iOS || this.platforms.KaiOS;
    },
    isWebApp() {
      return this.platforms.Web;
    },
    isWindowsApp() {
      return this.platforms.Windows;
    },
    isPhoneOs() {
      return this.isMobileApp;
    },
  },

  methods: {
    platformWord(key) {
      var locale = window.__locale || {};
      return locale[key] || key;
    },

    _setPlatformText() {
      var isMobile = this.isMobileApp;
      var isWin = this.isWindowsApp;
      var isWeb = this.isWebApp;

      var platformMobileDevices = this.platformWord("platform.mobileDevices");
      var platformWindowsDevices = this.platformWord("platform.windowsDevices");
      var platformWebBrowsers = this.platformWord("platform.webBrowsers");
      var platformMobileDevice = this.platformWord("platform.mobileDevice");
      var platformWindowsDevice = this.platformWord("platform.windowsDevice");
      var platformComputer = this.platformWord("platform.computer");
      var platformComputers = this.platformWord("platform.computers");
      var platformUninstallApp = this.platformWord("platform.uninstallApp");
      var platformCeaseWeb = this.platformWord("platform.ceaseWeb");
      var platformOr = this.platformWord("platform.or");
      var platformAnd = this.platformWord("platform.and");
      var platformCommaAnd = this.platformWord("platform.commaAnd");

      var descs = [];
      if (isMobile) descs.push(platformMobileDevices);
      if (isWin) descs.push(platformWindowsDevices);
      if (isWeb) descs.push(platformWebBrowsers);
      if (descs.length === 0) descs.push(platformMobileDevices);

      this.platformDesc = descs.length === 1
        ? descs[0]
        : descs.length === 2
          ? descs[0] + " " + platformAnd + " " + descs[1]
          : descs.slice(0, -1).join(", ") + platformCommaAnd + descs[descs.length - 1];

      var devs = [];
      if (isMobile) devs.push(platformMobileDevice);
      if (isWin) devs.push(platformWindowsDevice);
      if (isWeb) devs.push(platformComputer);
      if (devs.length === 0) devs.push(platformMobileDevice);

      this.deviceType = devs.length === 1
        ? devs[0]
        : devs.length === 2
          ? devs[0] + " " + platformOr + " " + devs[1]
          : devs.slice(0, -1).join(", ") + ", " + platformOr + " " + devs[devs.length - 1];

      var devPlurals = [];
      if (isMobile) devPlurals.push(platformMobileDevices);
      if (isWin) devPlurals.push(platformWindowsDevices);
      if (isWeb) devPlurals.push(platformComputers);
      if (devPlurals.length === 0) devPlurals.push(platformMobileDevices);

      this.deviceTypePlural = devPlurals.length === 1
        ? devPlurals[0]
        : devPlurals.length === 2
          ? devPlurals[0] + " " + platformAnd + " " + devPlurals[1]
          : devPlurals.slice(0, -1).join(", ") + platformCommaAnd + devPlurals[devPlurals.length - 1];

      var uninstallParts = [];
      if (isMobile || isWin) uninstallParts.push(platformUninstallApp);
      if (isWeb) uninstallParts.push(platformCeaseWeb);
      this.uninstallDesc = uninstallParts.join(" " + platformOr + " ") || platformUninstallApp;

      var deviceIdMobile = this.platformWord("platform.deviceId.mobile");
      var deviceIdWindows = this.platformWord("platform.deviceId.windows");
      var deviceIdWeb = this.platformWord("platform.deviceId.web");
      var deviceIdMixed = this.platformWord("platform.deviceId.mixed");
      var osMobile = this.platformWord("platform.os.mobile");
      var osWindows = this.platformWord("platform.os.windows");
      var osWeb = this.platformWord("platform.os.web");
      var osMixed = this.platformWord("platform.os.mixed");
      var browserMobile = this.platformWord("platform.browser.mobile");
      var browserWindows = this.platformWord("platform.browser.windows");
      var browserWeb = this.platformWord("platform.browser.web");
      var browserMixed = this.platformWord("platform.browser.mixed");

      if (isMobile && !isWin && !isWeb) {
        this.deviceIdDesc = deviceIdMobile;
        this.osDesc = osMobile;
        this.browserDesc = browserMobile;
      } else if (isWin && !isMobile && !isWeb) {
        this.deviceIdDesc = deviceIdWindows;
        this.osDesc = osWindows;
        this.browserDesc = browserWindows;
      } else if (isWeb && !isMobile && !isWin) {
        this.deviceIdDesc = deviceIdWeb;
        this.osDesc = osWeb;
        this.browserDesc = browserWeb;
      } else {
        this.deviceIdDesc = deviceIdMixed;
        this.osDesc = osMixed;
        this.browserDesc = browserMixed;
      }
    },
  }
};
