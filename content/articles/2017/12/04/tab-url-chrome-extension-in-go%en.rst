Get Current Tab URL From Chrome Extension in Go
###############################################

:date: 2018-02-02T07:34+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go,
       Chrome Extension
:category: Frontend Programming in Go
:summary: Get the URL of the current tab from Chrome extension, written in Go.
:og_image: https://i.ytimg.com/vi/vaXIknYrt4Y/maxresdefault.jpg
:adsu: yes

Build `Chrome Extension`_ with Go_, compiled to JavaScript via GopherJS_.
This post shows you how to use Go to build Chrome extension that gets the URL of
current active tab.

Besides GopherJS, you need the following packages to build extension:

- `github.com/fabioberger/chrome`_: GopherJS bindings for Chrome.
- godom_: `DOM manipulation`_ in Go.

Install the packages:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs
  $ go get -u github.com/fabioberger/chrome
  $ go get -u github.com/siongui/godom

To build the extension, we need three files: ``manifest.json``, ``popup.go``,
``popup.html``. We will explain one by one.

----

**manifest.json**:

.. code-block:: json

  {
    "manifest_version": 2,

    "name": "tab url",
    "description": "print url of current active tab",
    "version": "0.1",

    "browser_action": {
      "default_title": "tab url",
      "default_popup": "popup.html"
    },
    "permissions": [
      "tabs"
    ]
  }

This `manifest json`_ tells Chrome the information of our extension.

The above manifest says that when users click on the icon of our extension, a
popup window will show up. We need the `tabs permission`_ to access the
information of Chrome tabs.

.. adsu:: 2

----

**popup.html**:

.. code-block:: html

  <script src="popup.js"></script>

This HTML file contains nothing except including JavaScript file. We will print
the tab URL here via `document.write`_ method.

----

**popup.go**: compiled to *popup.js* via GopherJS.

.. code-block:: go

  package main

  import (
  	"github.com/fabioberger/chrome"
  	. "github.com/siongui/godom"
  )

  func main() {
  	c := chrome.NewChrome()
  	queryInfo := chrome.Object{
  		"active":        true,
  		"currentWindow": true,
  	}
  	c.Tabs.Query(queryInfo, func(tabs []chrome.Tab) {
  		tab := tabs[0]
  		Document.Write(tab.Url)
  	})
  }

The above code gets the current active tab, and then print the URL of the tab
via `document.write`_ method.

The above Go code is equivalent to the following JavaScript code:

.. code-block:: javascript

  chrome.tabs.query({'active': true, 'currentWindow': true}, function (tabs) {
      var tab = tabs[0];
      document.write(tab.url);
  });

The full code example, including JavaScript counterpart, of this post is
`on my GitHub`_.

----

.. adsu:: 3

References:

.. [1] | `chrome extension tab url - Google search <https://www.google.com/search?q=chrome+extension+tab+url>`_
       | `chrome extension tab url - DuckDuckGo search <https://duckduckgo.com/?q=chrome+extension+tab+url>`_
       | `chrome extension tab url - Ecosia search <https://www.ecosia.org/search?q=chrome+extension+tab+url>`_
       | `chrome extension tab url - Qwant search <https://www.qwant.com/?q=chrome+extension+tab+url>`_
       | `chrome extension tab url - Bing search <https://www.bing.com/search?q=chrome+extension+tab+url>`_
       | `chrome extension tab url - Yahoo search <https://search.yahoo.com/search?p=chrome+extension+tab+url>`_
       | `chrome extension tab url - Baidu search <https://www.baidu.com/s?wd=chrome+extension+tab+url>`_
       | `chrome extension tab url - Yandex search <https://www.yandex.com/search/?text=chrome+extension+tab+url>`_
.. [2] `How can I get the URL of the current tab from a Google Chrome extension? - Stack Overflow <https://stackoverflow.com/questions/1979583/how-can-i-get-the-url-of-the-current-tab-from-a-google-chrome-extension>`_
.. [3] `GopherJS - A compiler from Go to JavaScript <https://github.com/gopherjs/gopherjs>`_
       (`GopherJS Playground <https://gopherjs.github.io/playground//>`_,
       |godoc|)
.. [4] `GitHub - fabioberger/chrome: GopherJS Bindings for Chrome <https://github.com/fabioberger/chrome>`_
.. [5] `GitHub - siongui/godom: Make DOM manipultation in Go as similar to JavaScript as possible. (via GopherJS) <https://github.com/siongui/godom>`_

.. _Go: https://golang.org/
.. _Chrome Extension: https://www.google.com/search?q=Chrome+Extension
.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _github.com/fabioberger/chrome: https://github.com/fabioberger/chrome
.. _DOM manipulation: https://www.google.com/search?q=DOM+manipulation
.. _godom: https://github.com/siongui/godom
.. _manifest json: https://developer.chrome.com/extensions/manifest
.. _tabs permission: https://developer.chrome.com/extensions/declare_permissions
.. _document.write: https://developer.mozilla.org/en-US/docs/Web/API/Document/write
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/025-tab-url-chrome-extension
.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
