window.modalMixin = {
  methods: {
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

    toggleDisclaimerModalVisibility() {
      this.showDisclaimerModal = !this.showDisclaimerModal;
    },

    toggleFaqModalVisibility() {
      this.showFaqModal = !this.showFaqModal;
    },
  },
};
