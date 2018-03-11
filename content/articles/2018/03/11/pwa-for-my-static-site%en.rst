Progressive Web App (PWA) For My Static Website
###############################################

:date: 2018-03-11T23:18+08:00
:tags: web, Web application, Progressive Web App (PWA)
:category: Web Development
:summary: Make my static website become like a mobile app
          via Progressive Web App (PWA).
:og_image: https://speckyboy.com/wp-content/uploads/2016/10/progressive-web-applications-01.jpg
:adsu: yes


`Progressive Web App`_ is more and more popular recently [2]_. According to my
understanding, it is a way to turn your website into a mobile app, or at least
make users feel your website is a mobile app on their smartphones.

So, how to turn my static website, which is the website you are visiting now, to
a PWA? My requirement is very basic. I do not want to use any push notification
or cache, and I just want users to have my website icon on their smartphones and
click the icon to visit my website like a mobile app. After some investigations
[1]_, it seems easy to achieve. The following are steps:

1. Add *meta* tag to specify *viewport*: already done, since my website is
   responsive and this is added to the head of webpages long time ago.

2. `Web App Manifest`_: Modify the example provided in [1]_

   .. code-block:: json

    {
      "name": "Theory and Practice",
      "short_name": "T&P",
      "display": "minimal-ui",
      "start_url": "/",
      "theme_color": "#673ab6",
      "background_color": "#111111",
      "icons": [
        {
          "src": "/extra/Dharma_wheel.png",
          "sizes": "538x538",
          "type": "image/png"
        }
      ]
    }

3. Register a empty service worker: Also follow the example in [1]_, but note
   that put the *sw.js* (serviceWorker script) in the root path of your website,
   if you want to make your whole website in the app.

   The empty service worker code (saved as sw.js and put in the root path):

   .. code-block:: javascript

     /** An empty service worker! */
     self.addEventListener('fetch', function(event) {
       /** An empty fetch handler! */
     });

   Load service worker in your webpage:

   .. code-block:: html

     <script>
     navigator.serviceWorker && navigator.serviceWorker.register('/sw.js').then(function(registration) {
       console.log('Excellent, registered with scope: ', registration.scope);
     });
     </script>

The above steps is just enough to make my website behave like a app on
smartphones. If you have similar requirement, follow the steps and make your
website users more engaged!

.. adsu:: 2

----

References:

.. [1] `🎉 Migrate your site to a Progressive Web App 🐲 <https://codelabs.developers.google.com/codelabs/migrate-to-progressive-web-apps/index.html>`_
.. [2] `三巨头力捧的 PWA 有望在 2018 年迎来大爆发 - 开源中国社区 <https://www.oschina.net/news/93871/pwa-expected-major-explosion-2018>`_
.. [3] `Service Workers explained <https://flaviocopes.com/service-workers/>`_

.. _Progressive Web App: https://www.google.com/search?q=Progressive+Web+App
.. _Web App Manifest: https://www.google.com/search?q=web+app+manifest
