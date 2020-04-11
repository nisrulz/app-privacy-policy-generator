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

function convertHtmlToMd(html) {
  let markdown = toMarkdown(html, {
    converters: [{
      filter: 'div',
      replacement: function (content) {
        return content
      }
    }]
  })
  return markdown
}

function getRawHTML(content, title) {
  let html =
    "<!DOCTYPE html>\n\
    <html>\n\
    <head>\n\
      <meta charset='utf-8'>\n\
      <meta name='viewport' content='width=device-width'>\n\
      <title>" +
    title +
    "</title>\n\
      <style> body { font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif; padding:1em; } </style>\n\
    </head>\n\
    <body>\n\
    " +
    content +
    '\n\
    </body>\n\
    </html>\n\
      '
  return html
}

function getContent(id) {
  var content = document.getElementById(id)
  return content.innerHTML
}

function getTitle(id) {
  var content = document.getElementById(id)
  var title = content.getElementsByTagName('strong')[0]
  return title.innerHTML
}

function loadInTextView(id, content) {
  document.getElementById(id).value = content
}