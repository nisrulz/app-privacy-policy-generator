function translate(key, params) {
  var localeData = window.__locale || {}
  var text = localeData[key]
  if (text === undefined || text === null) return key
  if (!params) params = {}
  return text.replace(/\{\{(\s*)(\w+)(\s*)\}\}/g, function (match, before, name, after) {
    if (params[name] !== undefined) return params[name]
    if (this[name] !== undefined) return this[name]
    return match
  }.bind(this))
}

function _updateMeta() {
  var locale = window.__locale || {}
  if (locale['meta.title']) document.title = locale['meta.title']
  var metaDesc = document.querySelector('meta[name="description"]')
  if (metaDesc && locale['meta.description']) metaDesc.setAttribute('content', locale['meta.description'])
  var ogTitle = document.querySelector('meta[property="og:title"]')
  if (ogTitle && locale['meta.og.title']) ogTitle.setAttribute('content', locale['meta.og.title'])
  var ogDesc = document.querySelector('meta[property="og:description"]')
  if (ogDesc && locale['meta.og.description']) ogDesc.setAttribute('content', locale['meta.og.description'])
  var twitterTitle = document.querySelector('meta[name="twitter:title"]')
  if (twitterTitle && locale['meta.twitter.title']) twitterTitle.setAttribute('content', locale['meta.twitter.title'])
  var twitterDesc = document.querySelector('meta[name="twitter:description"]')
  if (twitterDesc && locale['meta.twitter.description']) twitterDesc.setAttribute('content', locale['meta.twitter.description'])
}

function _updateThemeLogo() {
  var theme = document.documentElement.getAttribute('data-theme');
  document.querySelectorAll('img[data-theme-logo]').forEach(function (img) {
    var light = img.getAttribute('data-light-src') || img.src;
    if (!img.getAttribute('data-light-src')) {
      img.setAttribute('data-light-src', img.src);
      img.setAttribute('data-dark-src', img.src.replace(/(\.\w+)$/, '_dark$1'));
    }
    img.src = theme === 'dark' ? img.getAttribute('data-dark-src') : img.getAttribute('data-light-src');
  });
}

function useAppState() {
  var { reactive, computed } = Vue;

  var state = reactive({
    currentLocale: document.documentElement.getAttribute('lang') || 'en',
    availableLocales: window.availableLocalesJsonArray || [],
    iOrWe: "[I/We]",
    typeOfApp: "Free",
    typeOfAppTxt: "a Free",
    typeOfDev: "Individual",
    appName: "",
    appContact: "",
    businessAddress: "",
    euRepresentative: "",
    devName: "",
    companyName: "",
    devOrCompanyName: "[Developer/Company name]",
    pidInfoIn: "",
    pidInfo: "[add whatever else you collect here, e.g. users name, address, location, pictures]",
    effectiveFromDate: new Date().toISOString().slice(0, 10),
    thirdPartyServices: thirdPartyServicesJsonArray,
    hasThirdPartyServicesSelected: true,
    showPrivacyModal: false,
    showGDPRPrivacyModal: false,
    showNoTrackingPrivacyPolicyModal: false,
    showTermsModal: false,
    showDisclaimerModal: false,
    showFaqModal: false,
    platforms: {
      Android: true,
      iOS: false,
      KaiOS: false,
      Windows: false,
      Web: false,
    },
    typeOfPolicy: "Simple",
    typeOfPolicyInt: 1,
    isLocationTracked: false,
    ageOfDigitalConsent: 16,
    isAIUsed: false,
    hasDataDeletion: false,
  });

  state.enabledThirdPartyServicesWithPrivacy = computed(function () {
    return state.thirdPartyServices.filter(function (item) { return item.enabled && item.link && item.link.privacy; });
  });
  state.enabledThirdPartyServicesWithTerms = computed(function () {
    return state.thirdPartyServices.filter(function (item) { return item.enabled && item.link && item.link.terms; });
  });
  state.thirdPartyServicesWithPrivacyOrTerms = computed(function () {
    return state.thirdPartyServices.filter(function (item) { return (item.link && item.link.privacy) || (item.link && item.link.terms); });
  });

  return state;
}

function useWizard(state) {
  var { reactive, computed } = Vue;

  var wizard = reactive({
    wizardStep: 1,
    totalWizardSteps: 8,
    contentRenderType: 1,
    deviceType: "",
    deviceTypePlural: "",
    platformDesc: "",
    deviceIdDesc: "",
    osDesc: "",
    browserDesc: "",
    uninstallDesc: "",
  });

  wizard.canAdvance = computed(function () {
    switch (wizard.wizardStep) {
      case 2:
        return (
          String(state.appName).trim().length > 0 &&
          String(state.appContact).trim().length > 0
        );
      case 6:
        if (state.typeOfDev === "Individual")
          return String(state.devName).trim().length > 0;
        if (state.typeOfDev === "Company")
          return String(state.companyName).trim().length > 0;
        return false;
      default:
        return true;
    }
  });

  wizard.selectedPlatformsLabel = computed(function () {
    var locale = window.__locale || {};
    var labels = [];
    if (state.platforms.Android) labels.push(locale["platform.android"] || "Android");
    if (state.platforms.iOS) labels.push(locale["platform.ios"] || "iOS");
    if (state.platforms.KaiOS) labels.push(locale["platform.kaios"] || "KaiOS");
    if (state.platforms.Windows) labels.push(locale["platform.windows"] || "Windows");
    if (state.platforms.Web) labels.push(locale["platform.web"] || "Web");
    return labels.length ? labels.join(", ") : translate("platform.selectedNone");
  });

  wizard.isMobileApp = computed(function () {
    return state.platforms.Android || state.platforms.iOS || state.platforms.KaiOS;
  });
  wizard.isWebApp = computed(function () {
    return state.platforms.Web;
  });
  wizard.isWindowsApp = computed(function () {
    return state.platforms.Windows;
  });
  wizard.isPhoneOs = computed(function () {
    return wizard.isMobileApp;
  });

  wizard.preview = function () {
    wizard.contentRenderType = 1;
  };

  wizard.nextStep = function () {
    if (!wizard.canAdvance) return;
    wizard.wizardStep += 1;
  };

  wizard.prevStep = function () {
    wizard.wizardStep -= 1;
  };

  wizard.isAppOpenSource = function () {
    return state.typeOfApp === "Open Source";
  };

  wizard.generate = function () {
    if (!_validateRequiredFields()) return false;
    _setDevOrCompanyName();
    _setPidInfo();
    _setAppTypeText();
    _setPlatformText();
    return true;
  };

  function _validateRequiredFields() {
    if (!state.appName.trim()) return false;
    if (!state.appContact.trim()) return false;
    if (state.typeOfDev === "Individual" && !state.devName.trim()) return false;
    if (state.typeOfDev === "Company" && !state.companyName.trim()) return false;
    return true;
  }

  function _setDevOrCompanyName() {
    state.devOrCompanyName =
      state.typeOfDev === "Individual" ? state.devName : state.companyName;
  }

  function _setPidInfo() {
    if (state.pidInfoIn === "") {
      state.pidInfo = ".";
    } else {
      var prefix = translate("misc.pidInfoPrefix");
      var suffix = translate("misc.pidInfoSuffix");
      state.pidInfo = prefix + state.pidInfoIn + suffix;
    }
  }

  function _setAppTypeText() {
    state.typeOfAppTxt = ["Open Source", "Ad Supported"].includes(state.typeOfApp)
      ? "an " + state.typeOfApp
      : "a " + state.typeOfApp;
  }

  function platformWord(key) {
    var locale = window.__locale || {};
    return locale[key] || key;
  }

  function _joinList(items, sepTwo, sepLast) {
    if (items.length === 1) return items[0];
    if (items.length === 2) return items[0] + " " + sepTwo + " " + items[1];
    return items.slice(0, -1).join(", ") + sepLast + items[items.length - 1];
  }

  function _setPlatformText() {
    var isMobile = wizard.isMobileApp;
    var isWin = wizard.isWindowsApp;
    var isWeb = wizard.isWebApp;

    var word = platformWord;

    var descs = [];
    if (isMobile) descs.push(word("platform.mobileDevices"));
    if (isWin) descs.push(word("platform.windowsDevices"));
    if (isWeb) descs.push(word("platform.webBrowsers"));
    if (descs.length === 0) descs.push(word("platform.mobileDevices"));
    wizard.platformDesc = _joinList(descs, word("platform.and"), word("platform.commaAnd"));

    var devs = [];
    if (isMobile) devs.push(word("platform.mobileDevice"));
    if (isWin) devs.push(word("platform.windowsDevice"));
    if (isWeb) devs.push(word("platform.computer"));
    if (devs.length === 0) devs.push(word("platform.mobileDevice"));
    wizard.deviceType = _joinList(devs, word("platform.or"), ", " + word("platform.or") + " ");

    var devPlurals = [];
    if (isMobile) devPlurals.push(word("platform.mobileDevices"));
    if (isWin) devPlurals.push(word("platform.windowsDevices"));
    if (isWeb) devPlurals.push(word("platform.computers"));
    if (devPlurals.length === 0) devPlurals.push(word("platform.mobileDevices"));
    wizard.deviceTypePlural = _joinList(devPlurals, word("platform.and"), word("platform.commaAnd"));

    var uninstallParts = [];
    if (isMobile || isWin) uninstallParts.push(word("platform.uninstallApp"));
    if (isWeb) uninstallParts.push(word("platform.ceaseWeb"));
    wizard.uninstallDesc = uninstallParts.join(" " + word("platform.or") + " ") || word("platform.uninstallApp");

    if (isMobile && !isWin && !isWeb) {
      wizard.deviceIdDesc = word("platform.deviceId.mobile");
      wizard.osDesc = word("platform.os.mobile");
      wizard.browserDesc = word("platform.browser.mobile");
    } else if (isWin && !isMobile && !isWeb) {
      wizard.deviceIdDesc = word("platform.deviceId.windows");
      wizard.osDesc = word("platform.os.windows");
      wizard.browserDesc = word("platform.browser.windows");
    } else if (isWeb && !isMobile && !isWin) {
      wizard.deviceIdDesc = word("platform.deviceId.web");
      wizard.osDesc = word("platform.os.web");
      wizard.browserDesc = word("platform.browser.web");
    } else {
      wizard.deviceIdDesc = word("platform.deviceId.mixed");
      wizard.osDesc = word("platform.os.mixed");
      wizard.browserDesc = word("platform.browser.mixed");
    }
  }

  wizard.checkForThirdPartyServicesEnabled = function () {
    return state.thirdPartyServices.some(function (item) { return item.enabled === true; });
  };

  wizard.tpsName = function (item) {
    var localeKey = 'name_' + state.currentLocale;
    return item[localeKey] || item.name;
  };

  wizard.toggleState = function (item) {
    item.enabled = !item.enabled;
  };

  wizard.setTypeOfPolicyInt = function () {
    switch (state.typeOfPolicy) {
      case "Simple":
        state.typeOfPolicyInt = 1;
        break;
      case "No Tracking":
        state.typeOfPolicyInt = 2;
        break;
      case "GDPR":
        state.typeOfPolicyInt = 3;
        break;
    }
  };

  return wizard;
}

function useContent(state, wizard) {
  function _toggleModal(modalFlag) {
    if (!wizard.generate()) return;
    state.hasThirdPartyServicesSelected = wizard.checkForThirdPartyServicesEnabled();
    wizard.contentRenderType = 1;
    state[modalFlag] = !state[modalFlag];
  }

  function toggleNoTrackingPrivacyPolicyModalVisibility() {
    _toggleModal("showNoTrackingPrivacyPolicyModal");
  }

  function togglePrivacyModalVisibility() {
    _toggleModal("showPrivacyModal");
  }

  function toggleGDPRPrivacyModalVisibility() {
    _toggleModal("showGDPRPrivacyModal");
  }

  function toggleTermsModalVisibility() {
    _toggleModal("showTermsModal");
  }

  function toggleDisclaimerModalVisibility() {
    state.showDisclaimerModal = !state.showDisclaimerModal;
  }

  function toggleFaqModalVisibility() {
    state.showFaqModal = !state.showFaqModal;
  }

  function getHtml(id, target) {
    var content = getContent(id);
    var title = getTitle(id);
    var rawHTML = getRawHTML(content, title);
    wizard.contentRenderType = 2;
    loadInTextView(target, rawHTML);
  }

  function getMarkdown(id, target) {
    var content = getContent(id);
    var title = getTitle(id);
    var rawHTML = getRawHTML(content, title);
    var markdown = convertHtmlToMd(rawHTML);
    wizard.contentRenderType = 2;
    loadInTextView(target, markdown);
  }

  function deployFcSimple() {
    fc_deploy('privacy_simple_content');
  }

  return {
    toggleNoTrackingPrivacyPolicyModalVisibility: toggleNoTrackingPrivacyPolicyModalVisibility,
    togglePrivacyModalVisibility: togglePrivacyModalVisibility,
    toggleGDPRPrivacyModalVisibility: toggleGDPRPrivacyModalVisibility,
    toggleTermsModalVisibility: toggleTermsModalVisibility,
    toggleDisclaimerModalVisibility: toggleDisclaimerModalVisibility,
    toggleFaqModalVisibility: toggleFaqModalVisibility,
    getHtml: getHtml,
    getMarkdown: getMarkdown,
    deployFcSimple: deployFcSimple,
  };
}

var appState = useAppState();
var appWizard = useWizard(appState);
var appContent = useContent(appState, appWizard);

var app = Vue.createApp({
  delimiters: ['[[', ']]'],
  setup: function () {
    return Object.assign({}, Vue.toRefs(appState), Vue.toRefs(appWizard), appContent);
  },
  mounted: function () {
    this.$nextTick(function () {
      _updateMeta();
      var theme = document.documentElement.getAttribute('data-theme');
      document.querySelectorAll('.theme-toggle').forEach(function (el) {
        el.textContent = theme === 'dark' ? '\u2600\uFE0F' : '\uD83C\uDF19';
      });
      _updateThemeLogo();
    });
  },
  watch: {
    wizardStep: function () {
      this.$nextTick(function () {
        var theme = document.documentElement.getAttribute('data-theme');
        document.querySelectorAll('.theme-toggle').forEach(function (el) {
          el.textContent = theme === 'dark' ? '\u2600\uFE0F' : '\uD83C\uDF19';
        });
        _updateThemeLogo();
      });
    }
  },
});

app.config.globalProperties.translate = translate;
app.config.globalProperties._updateMeta = _updateMeta;
app.config.globalProperties._updateThemeLogo = _updateThemeLogo;
app.config.globalProperties.toggleTheme = function (e) {
  var html = document.documentElement;
  var next = html.getAttribute('data-theme') === 'dark' ? 'light' : 'dark';
  html.setAttribute('data-theme', next);
  localStorage.setItem('theme', next);
  _updateThemeLogo();
  var btn = e && e.target;
  if (btn) btn.textContent = next === 'dark' ? '\u2600\uFE0F' : '\uD83C\uDF19';
  else {
    document.querySelectorAll('.theme-toggle').forEach(function (el) {
      el.textContent = next === 'dark' ? '\u2600\uFE0F' : '\uD83C\uDF19';
    });
  }
};

app.config.globalProperties.switchLocale = function (localeCode) {
  var currentLocale = document.documentElement.getAttribute('lang') || 'en';
  if (localeCode === currentLocale) return;
  var path = localeCode === 'en' ? '/' : '/' + localeCode + '/';
  window.location.assign(path);
};

app.mount("#app");
