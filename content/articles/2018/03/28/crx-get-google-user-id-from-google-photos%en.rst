[Chrome Extension] Get Google Account Id from Google Photos
###########################################################

:date: 2018-03-28T23:19+08:00
:tags: JavaScript, Chrome Extension, JSON
:category: Chrome Extension
:summary: Chrome extension to get Google account id from Google Photos.
:og_image: https://i.stack.imgur.com/KVVmZ.png
:adsu: yes


This Chrome extension helps you get your Google account/user id from
`Google Photos`_. There is already an easy way [1]_ to get Google account id.
This post provides another way to get it.

The idea is that there is a variable *window.WIZ_global_data* in the HTML source
of https://photos.google.com/. The *"S06Grb"* field of the variable contains
Google account id. We will find the string of the variable in the HTML source
and call *JSON.parse* method to convert it to a regular variable, and read the
id from the variable.

The following is source code for above idea.

**manifest.json**:

.. code-block:: json

  {
    "manifest_version": 2,

    "name": "guserid",
    "description": "Get Google user id",
    "version": "0.1",

    "browser_action": {
      "default_title": "get Google id",
      "default_popup": "popup.html"
    },

    "permissions": [
      "tabs",
      "<all_urls>"
    ]
  }

**popup.html**:

.. code-block:: html

  <script src="popup.js"></script>

**popup.js**:

.. code-block:: javascript

  chrome.runtime.onMessage.addListener(function(request, sender) {
    document.write(request.id);
  });

  chrome.tabs.executeScript(null, {file: "getid.js"});

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

.. [1] `How to get your Google account/user ID · GitHub <https://gist.github.com/msafi/b1cb05bfab5b897c57e6>`_
.. [2] | `chrome extension read page source - Google search <https://www.google.com/search?q=chrome+extension+read+page+source>`_
       | `chrome extension read page source - DuckDuckGo search <https://duckduckgo.com/?q=chrome+extension+read+page+source>`_
       | `chrome extension read page source - Ecosia search <https://www.ecosia.org/search?q=chrome+extension+read+page+source>`_
       | `chrome extension read page source - Qwant search <https://www.qwant.com/?q=chrome+extension+read+page+source>`_
       | `chrome extension read page source - Bing search <https://www.bing.com/search?q=chrome+extension+read+page+source>`_
       | `chrome extension read page source - Yahoo search <https://search.yahoo.com/search?p=chrome+extension+read+page+source>`_
       | `chrome extension read page source - Baidu search <https://www.baidu.com/s?wd=chrome+extension+read+page+source>`_
       | `chrome extension read page source - Yandex search <https://www.yandex.com/search/?text=chrome+extension+read+page+source>`_
.. [3] `Message Passing - Google Chrome <https://developer.chrome.com/apps/messaging>`_
.. [4] `Getting the source HTML of the current page from chrome extension - Stack Overflow <https://stackoverflow.com/questions/11684454/getting-the-source-html-of-the-current-page-from-chrome-extension>`_
.. [5] `[JavaScript] Traverse DOM Tree <{filename}/articles/2017/05/01/javascript-traverse-dom-tree%en.rst>`_
.. [6] | `chrome extension window variables - Google search <https://www.google.com/search?q=chrome+extension+window+variables>`_
       | `chrome extension window variables - DuckDuckGo search <https://duckduckgo.com/?q=chrome+extension+window+variables>`_
       | `chrome extension window variables - Ecosia search <https://www.ecosia.org/search?q=chrome+extension+window+variables>`_
       | `chrome extension window variables - Qwant search <https://www.qwant.com/?q=chrome+extension+window+variables>`_
       | `chrome extension window variables - Bing search <https://www.bing.com/search?q=chrome+extension+window+variables>`_
       | `chrome extension window variables - Yahoo search <https://search.yahoo.com/search?p=chrome+extension+window+variables>`_
       | `chrome extension window variables - Baidu search <https://www.baidu.com/s?wd=chrome+extension+window+variables>`_
       | `chrome extension window variables - Yandex search <https://www.yandex.com/search/?text=chrome+extension+window+variables>`_

.. _Google Photos: https://photos.google.com/
