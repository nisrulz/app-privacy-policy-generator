/*  
	App Privacy Policy Generator: A simple web app to generate a generic 
	privacy policy for your Android/iOS apps

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
      iOrWe: "[I/We]",
      typeOfApp: "Free",
      typeOfAppTxt: "[open source/free/freemium/ad-supported/commercial]",
      typeOfDev: "Individual",
      appName: "",
      appContact: "",
      devName: "",
      companyName: "",
      devOrCompanyName: "[Developer/Company name]",
      pidInfoIn: "",
      pidInfo:
        "[add whatever else you collect here, e.g. users name, address, location, pictures]",
      osType: "Android",
      effectiveFromDate: new Date().toISOString().slice(0, 10),
      requirementOfSystem: "system",

      // Third-party services array and UI flags
      thirdPartyServices: thirdPartyServicesJsonArray,
      showPrivacyModal: false,
      showGDPRPrivacyModal: false,
      showTermsModal: false,
      showDisclaimerModal: false,
      showFaqModal: false,
      showNoTrackingPrivacyPolicyModal: false,
      hasThirdPartyServicesSelected: true,

      // UI wizard state
      contentRenderType: 1, // 1 = Preview, 2 = HTML/Markdown
      wizardStep: 1,
      totalWizardSteps: 7,
      typeOfPolicy: "Simple",
      typeOfPolicyInt: 1,
      isLocationTracked: false,
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

  methods: {
    capitalize(value) {
      if (!value) return "";
      return value.charAt(0).toUpperCase() + value.slice(1);
    },

    // Preview view switch
    preview(id) {
      this.contentRenderType = 1;
    },

    // Set numeric policy type
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

    // Step navigator
    nextStep() {
      this.wizardStep += 1;
    },
    prevStep() {
      this.wizardStep -= 1;
    },

    // Check enabled 3rd party services
    checkForThirdPartyServicesEnabled() {
      return this.thirdPartyServices.some((item) => item.enabled === true);
    },

    // For reactive update of the JSON - toggle state
    toggleState(item) {
      item.enabled = !item.enabled;
    },

    // Render HTML content
    getHtml(id, target) {
      let content = getContent(id);
      let title = getTitle(id);
      let rawHTML = getRawHTML(content, title);
      this.contentRenderType = 2;
      loadInTextView(target, rawHTML);
    },

    // Render Markdown content
    getMarkdown(id, target) {
      let content = getContent(id);
      let title = getTitle(id);
      let rawHTML = getRawHTML(content, title);
      let markdown = convertHtmlToMd(rawHTML);
      this.contentRenderType = 2;
      loadInTextView(target, markdown);
    },

    // Generate output data based on inputs
    generate() {
      // Dev/Company Name
      this.devOrCompanyName =
        this.typeOfDev === "Individual" ? this.devName : this.companyName;

      // PID Info
      this.pidInfo =
        this.pidInfoIn === ""
          ? "."
          : ", including but not limited to " + this.pidInfoIn + ".";

      // App Type Text
      this.typeOfAppTxt = ["Open Source", "Ad Supported"].includes(
        this.typeOfApp
      )
        ? "an " + this.typeOfApp
        : "a " + this.typeOfApp;

      // OS system requirement
      switch (this.osType) {
        case "Android":
        case "iOS":
          this.requirementOfSystem = "system";
          break;
        case "Android & iOS":
          this.requirementOfSystem = "both systems";
          break;
      }
    },

    // Toggle modals & refresh data
    toggleNoTrackingPrivacyPolicyModalVisibility() {
      this.generate();
      this.contentRenderType = 1;
      this.showNoTrackingPrivacyPolicyModal =
        !this.showNoTrackingPrivacyPolicyModal;
    },
    togglePrivacyModalVisibility() {
      this.generate();
      this.hasThirdPartyServicesSelected =
        this.checkForThirdPartyServicesEnabled();
      this.contentRenderType = 1;
      this.showPrivacyModal = !this.showPrivacyModal;
    },
    toggleGDPRPrivacyModalVisibility() {
      this.generate();
      this.hasThirdPartyServicesSelected =
        this.checkForThirdPartyServicesEnabled();
      this.contentRenderType = 1;
      this.showGDPRPrivacyModal = !this.showGDPRPrivacyModal;
    },
    toggleTermsModalVisibility() {
      this.generate();
      this.hasThirdPartyServicesSelected =
        this.checkForThirdPartyServicesEnabled();
      this.contentRenderType = 1;
      this.showTermsModal = !this.showTermsModal;
    },
    toggleDisclaimerModalVisibility() {
      this.showDisclaimerModal = !this.showDisclaimerModal;
    },
    toggleFaqModalVisibility() {
      this.showFaqModal = !this.showFaqModal;
    },
    fc_deploy_simple() {
      fc_deploy_simple();
    },
    isAppOpenSource() {
      return this.typeOfApp === "Open Source";
    },
  },
});

// Mount the Vue app
app.mount("#app");
