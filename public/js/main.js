/*  
    App Privacy Policy Generator: A simple web app to generate a generic 
    privacy policy for your Android/iOS apps

    Copyright (C) 2020  Nishant Srivastava

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

var app = new Vue({
  el: "#app",
  data: {
    iOrWe: "[I/We]",
    typeOfApp: "",
    typeOfAppTxt: "[open source/free/freemium/ad-supported/commercial]",
    typeOfDev: "",
    appName: "[App name]",
    appContact: "[App contact information]",
    myOrOur: "[my/our]",
    meOrUs: "[me/us]",
    atNoCost: "[at no cost]",
    retainedInfo:
      "[retained on your device and is not collected by [me/us] in any way]/[retained by us and used as described in this privacy policy]",
    devName: "",
    companyName: "",
    devOrCompanyName: "[Developer/Company name]",
    pidInfoIn: "",
    pidInfo:
      "[add whatever else you collect here, e.g. users name, address, location, pictures]",
    osType: "",
    effectiveFromDate: new Date().toISOString().slice(0, 10),
    requirementOfSystem: "system",
    thirdPartyServices: thirdPartyServicesJsonArray,
    showPrivacyModal: false,
    showTermsModal: false,
    showDisclaimerModal: false,
    hasThirdPartyServicesSelected: true,
    contentRenderType: 1,
    wizardStep: 1,
    totalWizardSteps: 3,
  },
  filters: {
    capitalize: function (value) {
      if (!value) return ''
      value = value.toString()
      return value.charAt(0).toUpperCase() + value.slice(1)
    }
  },
  methods: {
    preview: function (id) {
      this.contentRenderType = 1
    },
    nextStep: function () {
      if (this.wizardStep <= this.totalWizardSteps) {
        this.wizardStep += 1
      }
    },
    prevStep: function () {
      if (this.wizardStep >= 1) {
        this.wizardStep -= 1
      }
    },
    checkForThirdPartyServicesEnabled: function () {
      let listOfEnabledThirdPartyServices = []
      this.thirdPartyServices.forEach((item) => {
        if (item[item.model] == true) {
          listOfEnabledThirdPartyServices.push(true)
        }
      })

      return listOfEnabledThirdPartyServices.length > 0
    },
    toggleState: function (item) {
      let state = item.model

      // console.log('Item:', item.name, item.model, item[state])
      // For reactive update of the json
      // Toggle the state
      Vue.set(thirdPartyServicesJsonArray, item.model, !item[state])
    },
    getHtml: function (id, target) {
      let content = getContent(id)
      let title = getTitle(id)
      let rawHTML = getRawHTML(content, title)
      this.contentRenderType = 2
      loadInTextView(target, rawHTML)
    },
    getMarkdown: function (id, target) {
      let content = getContent(id)
      let title = getTitle(id)
      let rawHTML = getRawHTML(content, title)
      let markdown = convertHtmlToMd(rawHTML)
      this.contentRenderType = 2
      loadInTextView(target, markdown)
    },
    generate: function () {
      if (this.typeOfDev === "Individual") {
        this.devOrCompanyName = this.devName
        this.iOrWe = "I"
        this.myOrOur = "my"
        this.meOrUs = "me"
        this.retainedInfo =
          "retained on your device and is not collected by " +
          this.meOrUs +
          " in any way"
      } else if (this.typeOfDev === "Company") {
        this.devOrCompanyName = this.companyName
        this.iOrWe = "we"
        this.myOrOur = "our"
        this.meOrUs = "us"
        this.retainedInfo =
          "retained by us and used as described in this privacy policy"
      }

      if (this.typeOfApp === "Commercial") {
        this.atNoCost = ""
      } else {
        this.atNoCost = "at no cost"
      }

      if (this.pidInfoIn === "") {
        this.pidInfo = "."
      } else {
        this.pidInfo = ", including but not limited to " + this.pidInfoIn + "."
      }

      switch (this.typeOfApp) {
        case "Free":
        case "Freemium":
        case "Commercial":
          this.typeOfAppTxt = "a " + this.typeOfApp
          break
        case "Open Source":
        case "Ad Supported":
          this.typeOfAppTxt = "an " + this.typeOfApp
          break
      }

      switch (this.osType) {
        case "Android": {
          this.osType = "Android"
          this.requirementOfSystem = "system"
          break
        }
        case "iOS": {
          this.osType = "iOS"
          this.requirementOfSystem = "system"
          break
        }
        case "Android & iOS": {
          this.osType = "Android & iOS"
          this.requirementOfSystem = "both systems"
          break
        }
      }
    },
    togglePrivacyModalVisibility: function () {
      this.generate()
      this.hasThirdPartyServicesSelected = this.checkForThirdPartyServicesEnabled()
      this.contentRenderType = 1
      this.showPrivacyModal = !this.showPrivacyModal
    },
    toggleTermsModalVisibility: function () {
      this.generate()
      this.hasThirdPartyServicesSelected = this.checkForThirdPartyServicesEnabled()
      this.contentRenderType = 1
      this.showTermsModal = !this.showTermsModal
    },
    toggleDisclaimerModalVisibility: function () {
      this.showDisclaimerModal = !this.showDisclaimerModal
    },
  },
})
