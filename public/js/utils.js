/*  
    App Privacy Policy Generator: A simple web app to generate a generic 
    privacy policy for your Android/iOS apps

    Copyright 2017-Present Nishant Srivastava

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
