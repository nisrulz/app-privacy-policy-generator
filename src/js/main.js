const { createApp } = Vue;

function translate(key, params) {
  var localeData = window.__locale || {}
  var text = localeData[key]
  if (text === undefined || text === null) return key
  if (!params) params = {}
  return text.replace(/\{\{(\s*)(\w+)(\s*)\}\}/g, function (match, before, name, after) {
    if (params[name] !== undefined) return params[name]
    if (this[name] !== undefined) return this[name]
    return match
  }.bind(this))
}

function _updateMeta() {
  var locale = window.__locale || {}
  if (locale['meta.title']) document.title = locale['meta.title']
  var metaDesc = document.querySelector('meta[name="description"]')
  if (metaDesc && locale['meta.description']) metaDesc.setAttribute('content', locale['meta.description'])
  var ogTitle = document.querySelector('meta[property="og:title"]')
  if (ogTitle && locale['meta.og.title']) ogTitle.setAttribute('content', locale['meta.og.title'])
  var ogDesc = document.querySelector('meta[property="og:description"]')
  if (ogDesc && locale['meta.og.description']) ogDesc.setAttribute('content', locale['meta.og.description'])
  var twitterTitle = document.querySelector('meta[name="twitter:title"]')
  if (twitterTitle && locale['meta.twitter.title']) twitterTitle.setAttribute('content', locale['meta.twitter.title'])
  var twitterDesc = document.querySelector('meta[name="twitter:description"]')
  if (twitterDesc && locale['meta.twitter.description']) twitterDesc.setAttribute('content', locale['meta.twitter.description'])
}

function _updateThemeLogo() {
  var theme = document.documentElement.getAttribute('data-theme');
  document.querySelectorAll('img[data-theme-logo]').forEach(function (img) {
    var light = img.getAttribute('data-light-src') || img.src;
    if (!img.getAttribute('data-light-src')) {
      img.setAttribute('data-light-src', img.src);
      img.setAttribute('data-dark-src', img.src.replace(/(\.\w+)$/, '_dark$1'));
    }
    img.src = theme === 'dark' ? img.getAttribute('data-dark-src') : img.getAttribute('data-light-src');
  });
}

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
  mounted: function () {
    this.$nextTick(function () {
      var theme = document.documentElement.getAttribute('data-theme');
      document.querySelectorAll('.theme-toggle').forEach(function (el) {
        el.textContent = theme === 'dark' ? '\u2600\uFE0F' : '\uD83C\uDF19';
      });
      _updateThemeLogo();
    });
  },
});

app.config.globalProperties.translate = translate
app.config.globalProperties._updateMeta = _updateMeta
app.config.globalProperties._updateThemeLogo = _updateThemeLogo
app.config.globalProperties.toggleTheme = function (e) {
  var html = document.documentElement;
  var next = html.getAttribute('data-theme') === 'dark' ? 'light' : 'dark';
  html.setAttribute('data-theme', next);
  localStorage.setItem('theme', next);
  _updateThemeLogo();
  var btn = e && e.target;
  if (btn) btn.textContent = next === 'dark' ? '\u2600\uFE0F' : '\uD83C\uDF19';
  else {
    document.querySelectorAll('.theme-toggle').forEach(function (el) {
      el.textContent = next === 'dark' ? '\u2600\uFE0F' : '\uD83C\uDF19';
    });
  }
}

app.config.globalProperties.switchLocale = function (localeCode) {
  var currentLocale = document.documentElement.getAttribute('lang') || 'en'
  if (localeCode === currentLocale) return
  var path = localeCode === 'en' ? '/' : '/' + localeCode + '/'
  window.location.assign(path)
}

app.mount("#app");