var __themeToggle = true;
(function(){
  if (!__themeToggle) return;
  var theme = localStorage.getItem('theme');
  if (!theme) theme = window.matchMedia('(prefers-color-scheme:dark)').matches ? 'dark' : 'light';
  document.documentElement.setAttribute('data-theme', theme);
})();
