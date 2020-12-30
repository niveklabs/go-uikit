const cacheName = "app-" + "675374d7a6eeab762a5a26f39eed910bfa445bca";

self.addEventListener("install", event => {
  console.log("installing app worker 675374d7a6eeab762a5a26f39eed910bfa445bca");
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
        "/web/js/prism.js",
        "/web/js/uikit-icons.min.js",
        "/web/js/uikit.min.js",
        "https://fonts.googleapis.com/css2?family=Fira+Code:wght@300;400;700&family=Montserrat:wght@300;400;700&display=swap",
        "https://storage.googleapis.com/murlok-github/icon-192.png",
        "https://storage.googleapis.com/murlok-github/icon-512.png",
        
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
  console.log("app worker 675374d7a6eeab762a5a26f39eed910bfa445bca is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
