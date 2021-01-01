const cacheName = "app-" + "4fa0d38235f90dc97c9d18f776584445d4c9e122";

self.addEventListener("install", event => {
  console.log("installing app worker 4fa0d38235f90dc97c9d18f776584445d4c9e122");
  self.skipWaiting();

  event.waitUntil(
    caches.open(cacheName).then(cache => {
      return cache.addAll([
        "/",
        "/app.css",
        "/app.js",
        "/manifest.webmanifest",
        "/wasm_exec.js",
        "/web/app.wasm",
        "/web/css/app.css",
        "/web/css/prism.css",
        "/web/css/uikit.min.css",
        "/web/images/android-chrome-192x192.png",
        "/web/images/android-chrome-512x512.png",
        "/web/images/apple-touch-icon.png",
        "/web/js/prism.js",
        "/web/js/uikit-icons.min.js",
        "/web/js/uikit.min.js",
        "https://fonts.googleapis.com/css2?family=Fira+Code:wght@300;400;700&family=Montserrat:wght@300;400;700&display=swap",
        
      ]);
    })
  );
});

self.addEventListener("activate", event => {
  event.waitUntil(
    caches.keys().then(keyList => {
      return Promise.all(
        keyList.map(key => {
          if (key !== cacheName) {
            return caches.delete(key);
          }
        })
      );
    })
  );
  console.log("app worker 4fa0d38235f90dc97c9d18f776584445d4c9e122 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
