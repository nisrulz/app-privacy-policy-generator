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
  "/images/app_graphics/side_image.webp",
  "/images/app_icons/disclaimer.svg"
];

self.addEventListener("install", (event) => {
  self.skipWaiting();
  event.waitUntil(
    caches.open(CACHE_NAME).then((cache) =>
      Promise.allSettled(
        PRECACHE_URLS.map((url) =>
          cache.add(url).catch(() => {})
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
    ).then(() => self.clients.claim())
  );
});

self.addEventListener("fetch", (event) => {
  const { request } = event;
  const url = new URL(request.url);

  if (url.origin !== location.origin) return;
  if (request.method !== "GET") return;
  if (url.pathname.startsWith("/cdn-cgi/")) return;

  if (request.mode === "navigate") {
    event.respondWith(networkFirst(request));
  } else {
    event.respondWith(staleWhileRevalidate(request));
  }
});

async function staleWhileRevalidate(request) {
  const cached = await caches.match(request);

  const fetchPromise = fetch(request).then(async (response) => {
    try {
      if (response.ok) {
        const cache = await caches.open(CACHE_NAME);
        await cache.put(request, response.clone());
      }
    } catch (_) {}
    return response;
  }).catch(() => null);

  return cached || (await fetchPromise) || new Response("", { status: 504 });
}

async function networkFirst(request) {
  // Normalize root URL to index.html for cache matching
  const url = new URL(request.url);
  const cacheKey = url.pathname === "/" ? "/index.html" : request;

  try {
    const response = await fetch(request);
    if (response.ok && request.method === "GET") {
      const cache = await caches.open(CACHE_NAME);
      await cache.put(cacheKey, response.clone());
    }
    return response;
  } catch {
    const cached = await caches.match(cacheKey);
    if (cached) return cached;
    const offlineShell = await caches.match("/index.html");
    if (offlineShell) return offlineShell;
    return new Response("Offline", { status: 503 });
  }
}
