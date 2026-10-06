window.themeToggle = (function () {
  function isDark() {
    return document.documentElement.getAttribute('data-theme') === 'dark';
  }

  function updateLogos() {
    var theme = document.documentElement.getAttribute('data-theme');
    document.querySelectorAll('img[data-theme-logo]').forEach(function (img) {
      if (!img.getAttribute('data-light-src')) {
        img.setAttribute('data-light-src', img.src);
        img.setAttribute('data-dark-src', img.src.replace(/(\.\w+)$/, '_dark$1'));
      }
      img.src = theme === 'dark' ? img.getAttribute('data-dark-src') : img.getAttribute('data-light-src');
    });
  }

  function updateGlyphs() {
    var glyph = isDark() ? '\u2600\uFE0F' : '\uD83C\uDF19';
    document.querySelectorAll('.theme-toggle').forEach(function (el) {
      el.textContent = glyph;
    });
  }

  function toggle() {
    var html = document.documentElement;
    var next = isDark() ? 'light' : 'dark';
    html.setAttribute('data-theme', next);
    localStorage.setItem('theme', next);
    updateLogos();
    updateGlyphs();
  }

  function init() {
    var theme = localStorage.getItem('theme');
    if (!theme) theme = window.matchMedia('(prefers-color-scheme:dark)').matches ? 'dark' : 'light';
    document.documentElement.setAttribute('data-theme', theme);
  }

  init();

  return { init: init, toggle: toggle, updateLogos: updateLogos, updateGlyphs: updateGlyphs };
})();
