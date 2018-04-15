[Chrome Extension] Get Cookies to Access Instagram API
######################################################

:date: 2018-04-15T23:44+08:00
:tags: JavaScript, Chrome Extension, HTTP cookie, JSON, Instagram
:category: Chrome Extension
:summary: Chrome extension to get cookies to access Instagram API.
:og_image: https://media.wired.com/photos/5926e97dcefba457b079b8b6/master/w_1774,c_limit/Instagram-v051916-s.jpg
:adsu: yes

Chrome extension to get cookies, used in my repo_, to access Instagram API.

**manifest.json**:

.. code-block:: json

  {
    "manifest_version": 2,

    "name": "igcookies",
    "description": "Get cookies to access Instagram API",
    "version": "0.1",

    "browser_action": {
      "default_title": "get IG cookies",
      "default_popup": "popup.html"
    },
    "permissions": [
      "cookies",
      "tabs",
      "*://www.instagram.com/"
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
    chrome.cookies.getAll({}, function (cookies) {

      var cookieNames = ["ds_user_id", "sessionid", "csrftoken"];

      var cookieAuth = {};
      document.write("<pre>");
      for (var i in cookies) {

        var cookie = cookies[i];
        if (cookieNames.indexOf(cookie.name) == -1) {
          continue;
        }

        cookieAuth[cookie.name] = cookie.value;
      }
      document.write(JSON.stringify(cookieAuth, null, 2));
      document.write("</pre>");
    });
  });

----

.. adsu:: 2

References:

.. [1] `[Chrome Extension] Get Authentication Cookies of gphotosuploader <{filename}/articles/2018/03/27/crx-gphotosuploader-auth-cookie-json%en.rst>`_

.. _repo: https://github.com/siongui/instago
