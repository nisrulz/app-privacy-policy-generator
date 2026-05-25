window.platformMixin = {
  data() {
    return {
      deviceType: "mobile device",
      deviceTypePlural: "mobile devices",
      platformDesc: "mobile devices",
      deviceIdDesc: "your mobile device's unique device ID",
      osDesc: "your mobile operating system",
      browserDesc: "the type of mobile Internet browsers you use",
      uninstallDesc: "by uninstalling the Application",
    };
  },

  computed: {
    selectedPlatformsLabel() {
      var labels = [];
      if (this.platforms.Android) labels.push("Android");
      if (this.platforms.iOS) labels.push("iOS");
      if (this.platforms.KaiOS) labels.push("KaiOS");
      if (this.platforms.Windows) labels.push("Windows");
      if (this.platforms.Web) labels.push("Web");
      return labels.length ? labels.join(", ") : "None";
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
    _setPlatformText() {
      var isMobile = this.isMobileApp;
      var isWin = this.isWindowsApp;
      var isWeb = this.isWebApp;

      var descs = [];
      if (isMobile) descs.push("mobile devices");
      if (isWin) descs.push("Windows devices");
      if (isWeb) descs.push("web browsers");
      if (descs.length === 0) descs.push("mobile devices");

      this.platformDesc = descs.length === 2
        ? descs[0] + " and " + descs[1]
        : descs.slice(0, -1).join(", ") + ", and " + descs[descs.length - 1];

      var devs = [];
      if (isMobile) devs.push("mobile device");
      if (isWin) devs.push("Windows device");
      if (isWeb) devs.push("computer");
      if (devs.length === 0) devs.push("mobile device");

      this.deviceType = devs.length === 2
        ? devs[0] + " or " + devs[1]
        : devs.slice(0, -1).join(", ") + ", or " + devs[devs.length - 1];

      var devPlurals = [];
      if (isMobile) devPlurals.push("mobile devices");
      if (isWin) devPlurals.push("Windows devices");
      if (isWeb) devPlurals.push("computers");
      if (devPlurals.length === 0) devPlurals.push("mobile devices");

      this.deviceTypePlural = devPlurals.length === 2
        ? devPlurals[0] + " and " + devPlurals[1]
        : devPlurals.slice(0, -1).join(", ") + ", and " + devPlurals[devPlurals.length - 1];

      var uninstallParts = [];
      if (isMobile || isWin) uninstallParts.push("by uninstalling the Application");
      if (isWeb) uninstallParts.push("by ceasing to use the website");
      this.uninstallDesc = uninstallParts.join(" or ") || "by uninstalling the Application";

      if (isMobile && !isWin && !isWeb) {
        this.deviceIdDesc = "your mobile device's unique device ID";
        this.osDesc = "your mobile operating system";
        this.browserDesc = "the type of mobile Internet browsers you use";
      } else if (isWin && !isMobile && !isWeb) {
        this.deviceIdDesc = "your device's unique device ID";
        this.osDesc = "your Windows operating system";
        this.browserDesc = "the type of Internet browser you use";
      } else if (isWeb && !isMobile && !isWin) {
        this.deviceIdDesc = "your device's unique identifier (e.g., IP address or browser fingerprint)";
        this.osDesc = "your operating system";
        this.browserDesc = "the type of web browser you use";
      } else {
        this.deviceIdDesc = "your device's unique device ID or identifier";
        this.osDesc = "your operating system";
        this.browserDesc = "the type of browser you use";
      }
    },
  }
};
