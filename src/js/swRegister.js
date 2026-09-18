if ('serviceWorker' in navigator) {
  navigator.serviceWorker.register('/sw.js').catch(function (err) {
    console.warn('SW registration failed:', err);
  });
  navigator.serviceWorker.addEventListener('controllerchange', function () {
    window.location.reload();
  });
}
