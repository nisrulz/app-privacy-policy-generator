window.generatorMixin = {
  methods: {
    capitalize(value) {
      if (!value) return "";
      return value.charAt(0).toUpperCase() + value.slice(1);
    },

    isAppOpenSource() {
      return this.typeOfApp === "Open Source";
    },

    generate() {
      if (!this._validateRequiredFields()) return false;
      this._setDevOrCompanyName();
      this._setPidInfo();
      this._setAppTypeText();
      this._setPlatformText();
      return true;
    },

    _validateRequiredFields() {
      if (!this.appName.trim()) return false;
      if (!this.appContact.trim()) return false;
      if (this.typeOfDev === "Individual" && !this.devName.trim()) return false;
      if (this.typeOfDev === "Company" && !this.companyName.trim()) return false;
      return true;
    },

    _setDevOrCompanyName() {
      this.devOrCompanyName =
        this.typeOfDev === "Individual" ? this.devName : this.companyName;
    },

    _setPidInfo() {
      if (this.pidInfoIn === "") {
        this.pidInfo = ".";
      } else {
        var prefix = this.$t("misc.pidInfoPrefix");
        var suffix = this.$t("misc.pidInfoSuffix");
        this.pidInfo = prefix + this.pidInfoIn + suffix;
      }
    },

    _setAppTypeText() {
      this.typeOfAppTxt = ["Open Source", "Ad Supported"].includes(
        this.typeOfApp
      )
        ? "an " + this.typeOfApp
        : "a " + this.typeOfApp;
    },
  },
};
