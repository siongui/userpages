[Golang/GopherJS] Chrome Extension for Chinese Conversion
#########################################################

:date: 2017-04-30T22:25+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Web application, web, DOM, html,
       Conversion of Traditional and Simplified Chinese, String Manipulation,
       Chrome Extension, Traverse DOM Tree
:category: GopherJS
:summary: Build Chrome Extension with Go, compiled to JavaScript via GopherJS.
          This post show you how to build Chrome extension that convert the
          web page of Simplified Chinese website to Traditional Chinese.
:og_image: https://pbs.twimg.com/profile_images/605816243870760960/4hP2sH_O.png
:adsu: yes


Build Chrome Extension with Go_, compiled to JavaScript_ via GopherJS_.
This post show you how to use Go to build `Chrome extension`_ that converts the
web page of Simplified Chinese website to Traditional Chinese.

This shows you the great power of GopherJS_ compiler. As long as the Go packages
are written in Go and not involved with OS operation, you can compile almost
any Go code to JavaScript and run on browsers or JavaScript environment.

The following Go packages are needed to build Chrome extension for converting
Simplified Chinese web page to Traditional Chinese:

- GopherJS_: Go-to-JavaScript compiler.
- `github.com/fabioberger/chrome`_: GopherJS bindings for Chrome.
- godom_: `DOM manipulation`_ in Go.
- gojianfan_: simple converter for Traditional and Simplified Chinese

Install the packages:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs
  $ go get -u github.com/fabioberger/chrome
  $ go get -u github.com/siongui/godom
  $ go get -u github.com/siongui/gojianfan

To build the extension, we need four files: ``icon.png``, ``manifest.json``,
``click.go``, ``cc.go``. We will explain one by one.

----

**icon.png**: Icon of your extension, shown in the toolbar of Chrome browser.

----

**manifest.json**:

.. code-block:: json

  {
    "manifest_version": 2,

    "name": "Gojianfan",
    "description": "Traditional and Simplified Chinese conversion",
    "version": "0.1",

    "browser_action": {
      "default_title": "Chinese conversion",
      "default_icon": "icon.png"
    },
    "background": {
      "scripts": ["click.js"],
      "persistent": false
    },
    "permissions": [
      "activeTab"
    ]
  }

This `manifest json`_ tells Chrome the information of our extension.

- `browser_action`_: Put the specified icon in the main Chrome toolbar. When
  mouse hovers over the icon, the text in ``default_title`` is shown in tooltip.
- background_: Tell Chrome what to do if users click on the icon of our
  extension in the toolbar. In our case, if the click event happens, the
  ``click.js`` will be executed, which converts the Simplified Chinese content
  of web page in current active tab to Traditional Chinese.

.. adsu:: 2

----

**click.go**: compiled to *click.js* via GopherJS.

.. code-block:: go

  package main

  import (
  	"github.com/fabioberger/chrome"
  )

  func main() {
  	c := chrome.NewChrome()
  	c.BrowserAction.OnClicked(func(tab chrome.Tab) {
  		o := chrome.Object{
  			"file": "cc.js",
  		}
  		c.Tabs.ExecuteScript(tab.Id, o, nil)
  	})
  }

This file says if users `click on the icon`_, we will `inject the script`_
*cc.js* to the current active tab. The *cc.js* script will convert the
Simplified Chinese to Traditional Chinese in active tab.

.. adsu:: 3

----

**cc.go**: compiled to *cc.js* via GopherJS

.. code-block:: go

  package main

  import (
  	. "github.com/siongui/godom"
  	"github.com/siongui/gojianfan"
  	"strings"
  )

  var excludedElement = map[string]bool{
  	"style":    true,
  	"script":   true,
  	"noscript": true,
  	"iframe":   true,
  	"object":   true,
  }

  func traverse(elm *Object) {
  	nodeType := elm.NodeType()

  	if nodeType == 1 || nodeType == 9 {
  		// element node or document node

  		if _, in := excludedElement[strings.ToLower(elm.TagName())]; in {
  			return
  		}

  		for _, node := range elm.ChildNodes() {
  			// recursively call to traverse
  			traverse(node)
  		}
  		return
  	}

  	if nodeType == 3 {
  		// text node
  		v := strings.TrimSpace(elm.NodeValue())
  		if v != "" {
  			elm.SetNodeValue(gojianfan.S2T(v))
  		}
  		return
  	}
  }

  func main() {
  	traverse(Document)
  }

What this file does is to traverse the DOM tree, find out all the visible
Simplified Chinese texts in the web page, and convert them to Traditional
Chinese.

.. adsu:: 4

----

The source code of Chrome extension for Chinese conversion can be found at
my GitHub repo [6]_. The code in the repo implements more features and is more
complicated than the code in this post.

Tested on: ``Ubuntu Linux 17.04``, ``Go 1.8.1``, ``GopherJS 1.8-1``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)
.. [2] `GitHub - fabioberger/chrome: GopherJS Bindings for Chrome <https://github.com/fabioberger/chrome>`_
.. [3] `GitHub - siongui/godom: Make DOM manipultation in Go as similar to JavaScript as possible. (via GopherJS) <https://github.com/siongui/godom>`_
.. [4] `GitHub - siongui/gojianfan: Traditional and Simplified Chinese Conversion in Go <https://github.com/siongui/gojianfan>`_
.. [5] `Online Conversion of Traditional and Simplified Chinese in Go/GopherJS <{filename}../../02/20/online-conversion-of-traditional-and-simplified-chinese%en.rst>`_
.. [6] `GitHub - siongui/go-chrome-extension-jianfan: Chrome extension for Traditional and Simplified Chinese conversion, written in Go <https://github.com/siongui/go-chrome-extension-jianfan>`_
.. [7] `Getting Started: Building a Chrome Extension - Google Chrome <https://developer.chrome.com/extensions/getstarted>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _github.com/fabioberger/chrome: https://github.com/fabioberger/chrome
.. _gojianfan: https://github.com/siongui/gojianfan
.. _godom: https://github.com/siongui/godom
.. _GopherJS: http://www.gopherjs.org/
.. _DOM manipulation: https://www.google.com/search?q=DOM+manipulation
.. _Chrome extension: https://www.google.com/search?q=Chrome+extension
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _manifest json: https://developer.chrome.com/extensions/manifest
.. _browser_action: https://developer.chrome.com/extensions/browserAction
.. _background: https://developer.chrome.com/extensions/event_pages
.. _click on the icon: https://developer.chrome.com/extensions/browserAction#event-onClicked
.. _inject the script: https://developer.chrome.com/extensions/content_scripts#pi

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
