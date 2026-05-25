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
      return labels.length ? labels.join(", ") : this.$t("platform.selectedNone");
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
    $tPlatform(key) {
      var locale = window.__locale || {};
      return locale[key] || key;
    },

    _setPlatformText() {
      var isMobile = this.isMobileApp;
      var isWin = this.isWindowsApp;
      var isWeb = this.isWebApp;

      var platformMobileDevices = this.$tPlatform("platform.mobileDevices");
      var platformWindowsDevices = this.$tPlatform("platform.windowsDevices");
      var platformWebBrowsers = this.$tPlatform("platform.webBrowsers");
      var platformMobileDevice = this.$tPlatform("platform.mobileDevice");
      var platformWindowsDevice = this.$tPlatform("platform.windowsDevice");
      var platformComputer = this.$tPlatform("platform.computer");
      var platformComputers = this.$tPlatform("platform.computers");
      var platformUninstallApp = this.$tPlatform("platform.uninstallApp");
      var platformCeaseWeb = this.$tPlatform("platform.ceaseWeb");
      var platformOr = this.$tPlatform("platform.or");
      var platformAnd = this.$tPlatform("platform.and");
      var platformCommaAnd = this.$tPlatform("platform.commaAnd");

      var descs = [];
      if (isMobile) descs.push(platformMobileDevices);
      if (isWin) descs.push(platformWindowsDevices);
      if (isWeb) descs.push(platformWebBrowsers);
      if (descs.length === 0) descs.push(platformMobileDevices);

      this.platformDesc = descs.length === 2
        ? descs[0] + " " + platformAnd + " " + descs[1]
        : descs.slice(0, -1).join(", ") + platformCommaAnd + descs[descs.length - 1];

      var devs = [];
      if (isMobile) devs.push(platformMobileDevice);
      if (isWin) devs.push(platformWindowsDevice);
      if (isWeb) devs.push(platformComputer);
      if (devs.length === 0) devs.push(platformMobileDevice);

      this.deviceType = devs.length === 2
        ? devs[0] + " " + platformOr + " " + devs[1]
        : devs.slice(0, -1).join(", ") + ", " + platformOr + " " + devs[devs.length - 1];

      var devPlurals = [];
      if (isMobile) devPlurals.push(platformMobileDevices);
      if (isWin) devPlurals.push(platformWindowsDevices);
      if (isWeb) devPlurals.push(platformComputers);
      if (devPlurals.length === 0) devPlurals.push(platformMobileDevices);

      this.deviceTypePlural = devPlurals.length === 2
        ? devPlurals[0] + " " + platformAnd + " " + devPlurals[1]
        : devPlurals.slice(0, -1).join(", ") + platformCommaAnd + devPlurals[devPlurals.length - 1];

      var uninstallParts = [];
      if (isMobile || isWin) uninstallParts.push(platformUninstallApp);
      if (isWeb) uninstallParts.push(platformCeaseWeb);
      this.uninstallDesc = uninstallParts.join(" " + platformOr + " ") || platformUninstallApp;

      var deviceIdMobile = this.$tPlatform("platform.deviceId.mobile");
      var deviceIdWindows = this.$tPlatform("platform.deviceId.windows");
      var deviceIdWeb = this.$tPlatform("platform.deviceId.web");
      var deviceIdMixed = this.$tPlatform("platform.deviceId.mixed");
      var osMobile = this.$tPlatform("platform.os.mobile");
      var osWindows = this.$tPlatform("platform.os.windows");
      var osWeb = this.$tPlatform("platform.os.web");
      var osMixed = this.$tPlatform("platform.os.mixed");
      var browserMobile = this.$tPlatform("platform.browser.mobile");
      var browserWindows = this.$tPlatform("platform.browser.windows");
      var browserWeb = this.$tPlatform("platform.browser.web");
      var browserMixed = this.$tPlatform("platform.browser.mixed");

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
