/*  
  App Privacy Policy Generator: A simple web app to generate a generic 
  privacy policy for your Android, iOS, and Web apps

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

const { createApp } = Vue;

const app = createApp({
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
      osType: "Android",
      effectiveFromDate: new Date().toISOString().slice(0, 10),
      requirementOfSystem: "system",

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

      // === UI State ===
      contentRenderType: 1, // 1 = Preview, 2 = HTML/Markdown
      wizardStep: 1,
      totalWizardSteps: 7,

      // === Policy Content Settings ===
      typeOfPolicy: "Simple",
      typeOfPolicyInt: 1,
      isLocationTracked: false,
      ageOfDigitalConsent: 16,
      isAIUsed: false,

      // === Platform-Aware Text (set by generate()) ===
      // These defaults assume Mobile App; generate() overrides for Web/Both
      platformType: "Mobile App",
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

    // Guards the Next button: blocks advance until required fields are filled
    canAdvance() {
      switch (this.wizardStep) {
        case 2:
          return !!(this.appName.trim() && this.appContact.trim());
        case 5:
          if (this.typeOfDev === "Individual") return !!this.devName.trim();
          if (this.typeOfDev === "Company") return !!this.companyName.trim();
          return false;
        default:
          return true;
      }
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

    // Switch content view to preview mode
    preview(id) {
      this.contentRenderType = 1;
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

    // ==================== Wizard Navigation ====================

    // Advance to next step; canAdvance blocks until required fields are filled
    nextStep() {
      if (!this.canAdvance) return;
      this.wizardStep += 1;
    },
    prevStep() {
      this.wizardStep -= 1;
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
