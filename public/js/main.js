/*  Copyright 2017 Nishant Srivastava

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at
    
       http://www.apache.org/licenses/LICENSE-2.0
    
    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License. */


var vueapp = new Vue({
  el: '#vueapp',
  data: {
    iOrWe: '[I/We]',
    typeOfApp: '',
    typeOfAppTxt: '[open source/free/freemium/ad-supported/commercial]',
    typeOfDev: '',
    appName: '[App Name]',
    myOrOur: '[my/our]',
    meOrUs: '[me/us]',
    atNoCost: '[at no cost]',
    retainedInfo: '[retained on your device and is not collected by [me/us] in any way]/[retained by us and used as described in this privacy policy]',
    devName: '',
    companyName: '',
    devOrCompanyName: '[Developer/Company name]',
    pidInfoIn: '',
    pidInfo: '[add whatever else you collect here, e.g. users name, address, location, pictures]',
    gps: true,
    admob: false,
    firebase_analytics: false,
    piwik: false,
    fabric: false,
    clicky: false,
    crashlytics:false,
    osType: '',
    requirementOfSystem :'system'
  },
  filters: {
    capitalize: function (value) {
      if (!value) return ''
      value = value.toString()
      return value.charAt(0).toUpperCase() + value.slice(1)
    }
  },
  methods: {
    selectAllText: function (id) {
      selectText(id)
    },
    generateHTML: function (id, filename) {
      let content = getContent(id)
      let title = getTitle(id)
      let rawHTML = getRawHTML(content, title)
      downloadHTML(filename, rawHTML)
    },
    generateMD: function (id, filename) {
      let content = getContent(id)
      let title = getTitle(id)
      let rawHTML = getRawHTML(content, title)
      let markdown = convertHtmlToMd(rawHTML)
      downloadMD(filename, markdown)
    },
    generate: function () {
      if (this.typeOfDev === 'Individual') {
        this.devOrCompanyName = this.devName
        this.iOrWe = 'I'
        this.myOrOur = 'my'
        this.meOrUs = 'me'
        this.retainedInfo = 'retained on your device and is not collected by ' + this.meOrUs + ' in any way'
      } else if (this.typeOfDev === 'Company') {
        this.devOrCompanyName = this.companyName
        this.iOrWe = 'we'
        this.myOrOur = 'our'
        this.meOrUs = 'us'
        this.retainedInfo = 'retained by us and used as described in this privacy policy'
      }

      if (this.typeOfApp === 'Commercial') {
        this.atNoCost = ''
      } else {
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

      switch (this.osType) {
        case 'Android':{
        this.osType = 'Android'
        this.requirementOfSystem = 'system'
        break
        }
        case 'iOS':{
          this.osType = 'iOS'
          this.requirementOfSystem = 'system'
          break
        }
        case 'Android & iOS':{
          this.osType = 'Android & iOS'
          this.requirementOfSystem = 'both systems'
          break
        }
      }

      document.getElementById('privacy_tab').click()
    }
  }
})

function convertHtmlToMd (html) {
  let markdown = toMarkdown(html, {converters: [{
    filter: 'div',
    replacement: function (content) { return content }
  }]})
  return markdown
}

function getRawHTML (content, title) {
  let html = "<!DOCTYPE html>\n\
<html>\n\
<head>\n\
  <meta charset='utf-8'>\n\
  <meta name='viewport' content='width=device-width'>\n\
  <title>" + title + "</title>\n\
  <style> body { font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif; padding:1em; } </style>\n\
</head>\n\
<body>\n\
" + content + "\n\
</body>\n\
</html>\n\
  "
  return html
}

function getContent (id) {
  var content = document.getElementById(id)
  return content.innerHTML
}

function getTitle (id) {
  var content = document.getElementById(id)
  var title = content.getElementsByTagName('h2')[0]
  return title.innerHTML
}

function selectText (containerId) {
  var range
  if (document.selection) {
    range = document.body.createTextRange()
    range.moveToElementText(document.getElementById(containerId))
    range.select()
  } else if (window.getSelection) {
    range = document.createRange()
    range.selectNode(document.getElementById(containerId))
    window.getSelection().addRange(range)
  }
}

function download (filename, text, format) {
  format += ';charset=utf-8,'
  var pom = document.createElement('a')
  pom.setAttribute('href', format + encodeURIComponent(text))
  pom.setAttribute('download', filename)

  if (document.createEvent) {
    var event = document.createEvent('MouseEvents')
    event.initEvent('click', true, true)
    pom.dispatchEvent(event)
  } else {
    pom.click()
  }
}

function downloadHTML (filename, content) {
  filename += '.html'
  download(filename, content, 'data:text/html')
}

function downloadMD (filename, content) {
  filename += '.md'
  download(filename, content, 'data:text/markdown')
}