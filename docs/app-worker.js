const cacheName = "app-" + "00cccf7b3c3c1d4ee83ed64ce90007566a158141";

self.addEventListener("install", event => {
  console.log("installing app worker 00cccf7b3c3c1d4ee83ed64ce90007566a158141");

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
  console.log("app worker 00cccf7b3c3c1d4ee83ed64ce90007566a158141 is activated");
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request).then(response => {
      return response || fetch(event.request);
    })
  );
});
