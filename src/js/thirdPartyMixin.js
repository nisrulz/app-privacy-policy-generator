window.thirdPartyMixin = {
  methods: {
    checkForThirdPartyServicesEnabled() {
      return this.thirdPartyServices.some((item) => item.enabled === true);
    },

    tpsName(item) {
      var localeKey = 'name_' + this.currentLocale;
      return item[localeKey] || item.name;
    },

    toggleState(item) {
      item.enabled = !item.enabled;
    },

    setTypeOfPolicyInt() {
      switch (this.typeOfPolicy) {
        case "Simple":
          this.typeOfPolicyInt = 1;
          break;
        case "No Tracking":
          this.typeOfPolicyInt = 2;
          break;
        case "GDPR":
          this.typeOfPolicyInt = 3;
          break;
      }
    },
  },
};
