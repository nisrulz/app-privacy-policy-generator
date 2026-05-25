/*  
  App Privacy Policy Generator — Service Worker
  Caches static assets for offline use and reduced server load.
*/

const CACHE_NAME = "app-privacy-policy-v2";

const PRECACHE_URLS = [
  "/index.html",
  "/reviews.html",
  "/404.html",
  "/css/style.min.css",
  "/js/main.min.js",
  "/js/utils.min.js",
  "/js/thirdpartyservices.min.js",
  "/js/locale.min.js",
  "/js/flycricket.min.js",
  "/js/vendor/vue.global.prod.js",
  "/js/vendor/to-markdown.min.js",
  "/favicon-32x32.png",
  "/favicon-16x16.png",
  "/android-chrome-192x192.png",
  "/android-chrome-512x512.png",
  "/apple-touch-icon.png",
  "/site.webmanifest",
  "/images/vendor/kofi1.png",
  "/images/app_graphics/side_image.png",
  "/images/app_icons/disclaimer.svg"
];

self.addEventListener("install", (event) => {
  event.waitUntil(
    caches.open(CACHE_NAME).then((cache) =>
      Promise.allSettled(
        PRECACHE_URLS.map((url) =>
          cache.add(url).catch(() => { })
        )
      )
    )
  );
});

self.addEventListener("activate", (event) => {
  event.waitUntil(
    caches.keys().then((names) =>
      Promise.allSettled(
        names
          .filter((name) => name !== CACHE_NAME)
          .map((name) => caches.delete(name))
      )
    )
  );
});

self.addEventListener("fetch", (event) => {
  const { request } = event;
  const url = new URL(request.url);

  if (url.origin !== location.origin) return;

  if (request.mode === "navigate") {
    event.respondWith(networkFirst(request));
  } else {
    event.respondWith(cacheFirst(request));
  }
});

async function cacheFirst(request) {
  const cached = await caches.match(request);
  if (cached) return cached;

  if (request.method !== "GET") return fetch(request);

  try {
    const response = await fetch(request);
    if (response.ok) {
      const cache = await caches.open(CACHE_NAME);
      cache.put(request, response.clone());
    }
    return response;
  } catch {
    const fallback = await caches.match(request);
    return fallback || new Response("", { status: 404 });
  }
}

async function networkFirst(request) {
  // Normalize root and locale paths to index.html for cache matching
  const url = new URL(request.url);
  const path = url.pathname;
  const isLocaleRoot = /^\/[a-z]{2}(\/index\.html)?$/.test(path) && path !== "/index.html";
  const cacheKey = path === "/" || isLocaleRoot ? (path.replace(/\/$/, "") + "/index.html") : request;

  if (request.method !== "GET") return fetch(request);

  try {
    const response = await fetch(request);
    if (response.ok) {
      const cache = await caches.open(CACHE_NAME);
      cache.put(cacheKey, response.clone());
    }
    return response;
  } catch {
    const cached = await caches.match(cacheKey);
    return cached || new Response("Offline", { status: 503 });
  }
}
