[Chrome Extension] Get auth.json of gphotosuploader
###################################################

:date: 2018-03-29T23:14+08:00
:tags: JavaScript, Chrome Extension, JSON
:category: Chrome Extension
:summary: Chrome extension to get auth.json (cookies of Google Photos and Google
          account id) of gphotosuploader.
:og_image: https://1.bp.blogspot.com/-hmZf-jlsuRg/Us5Uovd3iLI/AAAAAAAAZoo/wO037s_8cBM/s728/Google+Chrome+to+encrypt+Stored+Cookies+by+default+to+enhance+browser+security.jpg
:adsu: yes


gphotosuploader_ is an unofficial Google Photos uploader written in Go. To
upload photos/videos, an authentication file called *auth.json* must be created
first. The *auth.json* contains cookies of Google Photos and your Google account
id.

The day before yesterday I wrote a Chrome extension to get cookies of Google
Photos [1]_, and yesterday I wrote another extension to get Google account id
[2]_. Now the two extensions are merged to create the content of *auth.json* of
gphotosuploader_. The following is complete source code.

**manifest.json**:

.. code-block:: json

  {
    "manifest_version": 2,

    "name": "gpauth",
    "description": "Get content of auth.json for gphotosuploader",
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

  var auth = {
    "cookies": [],
    "persistantParameters": {
      "userId": ""
    }
  };

  chrome.tabs.query({
    active: true,
    currentWindow: true
  }, function(tabs) {
    var tab = tabs[0];

    chrome.cookies.getAll({}, function (cookies) {

      var cookieNames = ["OTZ", "CONSENT", "SID", "APISID", "SAPISID", "HSID", "NID", "SSID"];
      var cookieDomains = [".google.com", "photos.google.com"];

      auth["cookies"].length = 0;
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
      chrome.tabs.executeScript(null, {file: "getid.js"});

    });
  });

  chrome.runtime.onMessage.addListener(function(request, sender) {
    auth["persistantParameters"]["userId"] = request.id;
    document.write("<pre>");
    document.write(JSON.stringify(auth, null, 2));
    document.write("</pre>");
  });

**getid.js**:

.. code-block:: javascript

  function find_WIZ_global_data(elm) {
    if (elm.nodeType == Node.ELEMENT_NODE || elm.nodeType == Node.DOCUMENT_NODE) {
      for (var i=0; i < elm.childNodes.length; i++) {
        // recursively call self
        find_WIZ_global_data(elm.childNodes[i]);
      }
    }

    if (elm.nodeType == Node.TEXT_NODE) {
      if (elm.nodeValue.startsWith("window.WIZ_global_data")) {
        var jsonString = elm.nodeValue.replace("window.WIZ_global_data = ", "");
        jsonString = jsonString.slice(0, -1);
        var wiz = JSON.parse(jsonString);
        chrome.runtime.sendMessage({id: wiz["S06Grb"]});
      }
    }
  }

  find_WIZ_global_data(document);

----

.. adsu:: 2

References:

.. [1] `[Chrome Extension] Get Authentication Cookies of gphotosuploader <{filename}/articles/2018/03/27/crx-gphotosuploader-auth-cookie-json%en.rst>`_
.. [2] `[Chrome Extension] Get Google Account Id from Google Photos <{filename}/articles/2018/03/28/crx-get-google-user-id-from-google-photos%en.rst>`_

.. _Google Photos: https://photos.google.com/
.. _gphotosuploader: https://github.com/simonedegiacomi/gphotosuploader
