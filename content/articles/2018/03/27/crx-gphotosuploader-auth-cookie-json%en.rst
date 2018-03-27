[Chrome Extension] Get Authentication Cookies of gphotosuploader
################################################################

:date: 2018-03-27T22:59+08:00
:tags: JavaScript, Chrome Extension, HTTP cookie
:category: JavaScript
:summary: Chrome extension to get cookies of *gphotosuploader*.
:og_image: https://davidhancock.co/wp-content/uploads/2012/11/webkit-dev-tools.png
:adsu: yes


gphotosuploader_ is a tool that helps you upload photos/videos to Google Photos.
It uses WebDrivers Protocol to get authentication cookies and user id to upload
to Google Photos. I write a Chrome extension to get authentication cookies
instead of using WebDrivers. The following is my code:

**manifest.json**:

.. code-block:: json

  {
    "manifest_version": 2,

    "name": "gpcookies",
    "description": "Get authentication cookies of gphotosuploader",
    "version": "0.1",

    "browser_action": {
      "default_title": "export gphotosuploader auth.json",
      "default_popup": "popup.html"
    },
    "permissions": [
      "cookies",
      "tabs",
      "<all_urls>"
    ]
  }

**popup.html**:

.. code-block:: html

  <script src="popup.js"></script>

**popup.js**:

.. code-block:: javascript

  chrome.tabs.query({
    active: true,
    currentWindow: true
  }, function(tabs) {
    var tab = tabs[0];

    chrome.cookies.getAll({}, function (cookies) {

      var cookieNames = ["OTZ", "CONSENT", "SID", "APISID", "SAPISID", "HSID", "NID", "SSID"];
      var cookieDomains = [".google.com", "photos.google.com"];

      var auth = {
        "cookies": [],
        "persistantParameters": {
          "userId": ""
        }
      };

      document.write("<pre>");
      for (var i in cookies) {

        var cookie = cookies[i];
        if (cookieNames.indexOf(cookie.name) == -1) {
          continue;
        }
        if (cookieDomains.indexOf(cookie.domain) == -1) {
          continue;
        }

        var cookieAuth = {};
        cookieAuth["Name"] = cookie.name;
        cookieAuth["Value"] = cookie.value;
        cookieAuth["Domain"] = cookie.domain;
        cookieAuth["HttpOnly"] = cookie.httpOnly;
        cookieAuth["Secure"] = cookie.secure;
        cookieAuth["Path"] = cookie.path;

        auth["cookies"].push(cookieAuth);

      }
      document.write(JSON.stringify(auth, null, 2));
      document.write("</pre>");
    });
  });

----

.. adsu:: 2

References:

.. [1] `GitHub - simonedegiacomi/gphotosuploader: Unofficial Google Photos uploader and Go library <https://github.com/simonedegiacomi/gphotosuploader>`_
.. [2] `chrome.cookies - Google Chrome <https://developer.chrome.com/extensions/cookies>`_
.. [3] `webextensions-examples/list-cookies at master · mdn/webextensions-examples · GitHub <https://github.com/mdn/webextensions-examples/tree/master/list-cookies>`_

.. _gphotosuploader: https://github.com/simonedegiacomi/gphotosuploader
