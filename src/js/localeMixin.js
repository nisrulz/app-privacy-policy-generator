window.localeMixin = {
  data() {
    return {
      currentLocale: document.documentElement.getAttribute('lang') || 'en',
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

    _updateMeta() {
      var locale = window.__locale || {}
      if (locale['meta.title']) {
        document.title = locale['meta.title']
      }
      var metaDesc = document.querySelector('meta[name="description"]')
      if (metaDesc && locale['meta.description']) {
        metaDesc.setAttribute('content', locale['meta.description'])
      }
      var ogTitle = document.querySelector('meta[property="og:title"]')
      if (ogTitle && locale['meta.og.title']) {
        ogTitle.setAttribute('content', locale['meta.og.title'])
      }
      var ogDesc = document.querySelector('meta[property="og:description"]')
      if (ogDesc && locale['meta.og.description']) {
        ogDesc.setAttribute('content', locale['meta.og.description'])
      }
      var twitterTitle = document.querySelector('meta[name="twitter:title"]')
      if (twitterTitle && locale['meta.twitter.title']) {
        twitterTitle.setAttribute('content', locale['meta.twitter.title'])
      }
      var twitterDesc = document.querySelector('meta[name="twitter:description"]')
      if (twitterDesc && locale['meta.twitter.description']) {
        twitterDesc.setAttribute('content', locale['meta.twitter.description'])
      }
    },

    switchLocale(localeCode) {
      if (localeCode === this.currentLocale) return
      var path = localeCode === 'en' ? '/' : '/' + localeCode + '/'
      window.location.href = path
    }
  },

  mounted() {
    this._updateMeta()
  }
}
