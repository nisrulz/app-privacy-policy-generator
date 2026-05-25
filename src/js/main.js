const { createApp } = Vue;

const app = createApp({
  mixins: [
    window.wizardMixin,
    window.platformMixin,
    window.localeMixin,
    window.formDataMixin,
    window.generatorMixin,
    window.modalMixin,
    window.thirdPartyMixin,
    window.contentMixin,
  ],
});

app.mount("#app");
