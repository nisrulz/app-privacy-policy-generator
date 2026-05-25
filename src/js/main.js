/*  
  App Privacy Policy Generator: A simple web app to generate a generic 
  privacy policy for your Android, iOS, and Web apps

  Copyright 2017-Present Nishant Srivastava

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
*/

const { createApp } = Vue;

const app = createApp({
  // Composition: Registering decoupled mixins
  mixins: [window.wizardMixin, window.platformMixin],

  data() {
    return {
      // === Wizard Form Inputs ===
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
      pidInfo:
        "[add whatever else you collect here, e.g. users name, address, location, pictures]",
      effectiveFromDate: new Date().toISOString().slice(0, 10),

      // === Third-Party Services ===
      thirdPartyServices: thirdPartyServicesJsonArray,
      hasThirdPartyServicesSelected: true,

      // === Privacy Policy Modal Flags ===
      showPrivacyModal: false,
      showGDPRPrivacyModal: false,
      showNoTrackingPrivacyPolicyModal: false,

      // === Other Modal Flags ===
      showTermsModal: false,
      showDisclaimerModal: false,
      showFaqModal: false,

      // === Policy Content Settings ===
      typeOfPolicy: "Simple",
      typeOfPolicyInt: 1,
      isLocationTracked: false,
      ageOfDigitalConsent: 16,
      isAIUsed: false,
    };
  },

  computed: {
    // Filter third-party services based on what links they provide
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

  methods: {
    // ==================== Utility ====================

    capitalize(value) {
      if (!value) return "";
      return value.charAt(0).toUpperCase() + value.slice(1);
    },
    isAppOpenSource() {
      return this.typeOfApp === "Open Source";
    },

    // ==================== UI Helpers ====================

    // Check whether the user has enabled at least one third-party service
    checkForThirdPartyServicesEnabled() {
      return this.thirdPartyServices.some((item) => item.enabled === true);
    },

    // Toggle a third-party service checkbox (mutates the array item in place)
    toggleState(item) {
      item.enabled = !item.enabled;
    },

    // Map policy type name to the numeric value used in v-show/v-if conditions
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

    // ==================== Content Rendering ====================

    // Render the generated content as raw HTML for copy/paste
    getHtml(id, target) {
      let content = getContent(id);
      let title = getTitle(id);
      let rawHTML = getRawHTML(content, title);
      this.contentRenderType = 2;
      loadInTextView(target, rawHTML);
    },

    // Render the generated content as Markdown for copy/paste
    getMarkdown(id, target) {
      let content = getContent(id);
      let title = getTitle(id);
      let rawHTML = getRawHTML(content, title);
      let markdown = convertHtmlToMd(rawHTML);
      this.contentRenderType = 2;
      loadInTextView(target, markdown);
    },

    // ==================== Data Generation ====================

    // Validate required fields, then compute all derived template variables.
    // Called by every modal toggle; returns true only when all checks pass.
    generate() {
      if (!this._validateRequiredFields()) return false;

      this._setDevOrCompanyName();
      this._setPidInfo();
      this._setAppTypeText();
      this._setOsRequirement();
      this._setPlatformText();
      return true;
    },

    // -- Private helpers (prefixed with _) --

    // Ensure the user filled every field the modals depend on.
    // The wizard's canAdvance already prevents reaching step 7 without these,
    // but this serves as a safety net for edge cases (e.g. direct URL access).
    _validateRequiredFields() {
      if (!this.appName.trim()) return false;
      if (!this.appContact.trim()) return false;
      if (this.typeOfDev === "Individual" && !this.devName.trim()) return false;
      if (this.typeOfDev === "Company" && !this.companyName.trim()) return false;
      return true;
    },

    // Set display name: use developer name for individuals, company name otherwise
    _setDevOrCompanyName() {
      this.devOrCompanyName =
        this.typeOfDev === "Individual" ? this.devName : this.companyName;
    },

    // PII text: empty → "." (no extra info), non-empty → ", including but not limited to ..."
    _setPidInfo() {
      this.pidInfo =
        this.pidInfoIn === ""
          ? "."
          : ", including but not limited to " + this.pidInfoIn + ".";
    },

    // "Open Source" and "Ad Supported" take "an" instead of "a" before the type
    _setAppTypeText() {
      this.typeOfAppTxt = ["Open Source", "Ad Supported"].includes(
        this.typeOfApp
      )
        ? "an " + this.typeOfApp
        : "a " + this.typeOfApp;
    },

    // ==================== Modal Toggles ====================

    // Open or close a policy modal, guarded by generate().
    // generate() returns early (false) if required fields are missing,
    // which prevents the modal from opening with placeholder text.
    _toggleModal(modalFlag) {
      if (!this.generate()) return;
      this.hasThirdPartyServicesSelected =
        this.checkForThirdPartyServicesEnabled();
      this.contentRenderType = 1;
      this[modalFlag] = !this[modalFlag];
    },

    toggleNoTrackingPrivacyPolicyModalVisibility() {
      if (!this.generate()) return;
      this.contentRenderType = 1;
      this.showNoTrackingPrivacyPolicyModal =
        !this.showNoTrackingPrivacyPolicyModal;
    },

    togglePrivacyModalVisibility() {
      this._toggleModal("showPrivacyModal");
    },
    toggleGDPRPrivacyModalVisibility() {
      this._toggleModal("showGDPRPrivacyModal");
    },
    toggleTermsModalVisibility() {
      this._toggleModal("showTermsModal");
    },

    // Disclaimer and FAQ don't need generate() — they're static content
    toggleDisclaimerModalVisibility() {
      this.showDisclaimerModal = !this.showDisclaimerModal;
    },
    toggleFaqModalVisibility() {
      this.showFaqModal = !this.showFaqModal;
    },

    // ==================== Third-Party Integration ====================

    deployFcSimple() {
      fc_deploy_simple();
    },
  },
});

app.mount("#app");
