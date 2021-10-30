const cacheName = "app-" + "b5d712cfceed16edfd98a9f19ead5bf4aaa947f5";

self.addEventListener("install", event => {
  console.log("installing app worker b5d712cfceed16edfd98a9f19ead5bf4aaa947f5");

  event.waitUntil(
    caches.open(cacheName).
      then(cache => {
        return cache.addAll([
          "/go-app-website",
          "/go-app-website/app.css",
          "/go-app-website/app.js",
          "/go-app-website/manifest.webmanifest",
          "/go-app-website/wasm_exec.js",
          "/go-app-website/web/app.wasm",
          "https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css",
          "https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js",
          "https://storage.googleapis.com/murlok-github/icon-192.png",
          "https://storage.googleapis.com/murlok-github/icon-512.png",
          
        ]);
      }).
      then(() => {
        self.skipWaiting();
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
  console.log("app worker b5d712cfceed16edfd98a9f19ead5bf4aaa947f5 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
