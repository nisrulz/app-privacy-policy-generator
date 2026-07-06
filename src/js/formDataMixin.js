window.formDataMixin = {
  data() {
    return {
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
    };
  },

  computed: {
    enabledThirdPartyServicesWithPrivacy() {
      return this.thirdPartyServices.filter(
        (item) => item.enabled && item.link?.privacy
      );
    },
    enabledThirdPartyServicesWithTerms() {
      return this.thirdPartyServices.filter(
        (item) => item.enabled && item.link?.terms
      );
    },
    thirdPartyServicesWithPrivacyOrTerms() {
      return this.thirdPartyServices.filter(
        (item) => item.link?.privacy || item.link?.terms
      );
    },
  },
};
