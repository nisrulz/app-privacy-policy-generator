window.localeMixin = {
  data() {
    return {
      currentLocale: 'en',
      availableLocales: [
        { code: 'en', label: 'English' }
      ]
    }
  },

  methods: {
    $t(key, params) {
      var localeData = window.__locale || {}
      var text = localeData[key]
      if (text === undefined || text === null) return key
      if (!params) params = {}
      return text.replace(/\{\{(\s*)(\w+)(\s*)\}\}/g, function (match, before, name, after) {
        if (params[name] !== undefined) return params[name]
        if (this[name] !== undefined) return this[name]
        return match
      }.bind(this))
    },

    switchLocale(localeCode) {
      if (localeCode === this.currentLocale) return
      var path = localeCode === 'en' ? '/' : '/' + localeCode + '/'
      window.location.href = path
    }
  }
}
