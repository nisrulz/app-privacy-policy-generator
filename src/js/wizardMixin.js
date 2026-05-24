/*  
  App Privacy Policy Generator: A simple web app to generate a generic 
  privacy policy for your Android, iOS, and Web apps

  Copyright 2017-Present Nishant Srivastava

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
*/

window.wizardMixin = {
  data() {
    return {
      wizardStep: 1,
      totalWizardSteps: 8,
      contentRenderType: 1, // 1 = Preview, 2 = HTML/Markdown
    };
  },

  computed: {
    // Guards the Next button: blocks advance until required fields are filled
    canAdvance() {
      switch (this.wizardStep) {
        case 2:
          return !!(this.appName.trim() && this.appContact.trim());
        case 6:
          if (this.typeOfDev === "Individual") return !!this.devName.trim();
          if (this.typeOfDev === "Company") return !!this.companyName.trim();
          return false;
        default:
          return true;
      }
    },
  },

  methods: {
    // Switch content view to preview mode
    preview(id) {
      this.contentRenderType = 1;
    },

    // Advance to next step; canAdvance blocks until required fields are filled
    nextStep() {
      if (!this.canAdvance) return;
      this.wizardStep += 1;
    },

    prevStep() {
      this.wizardStep -= 1;
    },
  }
};
