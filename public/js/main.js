function escapeHtml (string) {
  var entityMap = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#39;',
    '/': '&#x2F;'
  }
  return String(string).replace(/[&<>"'\/]/g, function (s) {
    return entityMap[s]
  })
}

function generateHTML () {
  var codeblock = document.getElementById('preview')
  var head = '<!DOCTYPE html><html> <head> <meta charset="utf-8"> <meta name="viewport" content="width=device-width"> <title>Privacy Policy</title> <style>body{font-family: "Helvetica Neue", Helvetica, Arial, sans-serif; padding:1em;}</style></head> <body>'
  var end = '</body></html>'
  var html = head + codeblock.innerHTML + end
  html = escapeHtml(html)
  document.getElementById('html').innerHTML = html
}

var vueapp = new Vue({
  el: '#vueapp',
  data: {
    iOrWe: '[I | we]',
    typeOfApp: '[open source | free | freemium | ad-supported | commercial]',
    typeOfDev: '[Individual or Company]',
    appName: '[App Name]',
    myOrOur: '[my|our]',
    meOrUs: '[me|us]',
    pidInfo: '[add whatever else you collect here, e.g. users name | address | location | pictures]',
    atNoCost: '[at no cost]',
    devOrCompanyName: '[Dev Name or Company Name]',
    retainedInfo: '[retained on your device and is not collected by [me|us] in any way]|[will be retained by us and used as described in this privacy policy].',
    devName: '',
    companyName: '',
    devOrCompanyName: '',
    pidInfoIn: '',
    typeOfAppTxt: ''
  },
  methods: {
    generate: function () {
      if (this.typeOfDev == 'Individual') {
        this.devOrCompanyName = this.devName
        this.iOrWe = 'I'
        this.myOrOur = 'my'
        this.meOrUs = 'me'
        this.retainedInfo = 'retained on your device and is not collected by ' + this.meOrUs + ' in any way'
      }else if (this.typeOfDev == 'Company') {
        this.devOrCompanyName = this.companyName
        this.iOrWe = 'we'
        this.myOrOur = 'our'
        this.meOrUs = 'us'
        this.retainedInfo = 'will be retained by us and used as described in this privacy policy.'
      }

      if (this.typeOfApp == 'Commercial') {
        this.atNoCost = ''
      }else {
        this.atNoCost = 'at no cost'
      }

      if (this.pidInfoIn === '') {
        this.pidInfo = '.'
      } else {
        this.pidInfo = ', including but not limited to ' + this.pidInfoIn + '.'
      }

      switch (this.typeOfApp) {
        case 'Free':
        case 'Freemium':
        case 'Commercial':
          this.typeOfAppTxt = 'a ' + this.typeOfApp
          break
        case 'Open Source':
        case 'Ad Supported':
          this.typeOfAppTxt = 'an ' + this.typeOfApp
          break
      }

      generateHTML()
    }
  }
})
