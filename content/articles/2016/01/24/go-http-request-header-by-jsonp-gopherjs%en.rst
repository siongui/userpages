[Golang] Access HTTP Request Header (Accept-Language) by JSONP
##############################################################

:date: 2016-01-24T20:17+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, Locale, i18n, JSONP, JSON,
       CORS, HTTP Header, Accept-Language
:category: GopherJS
:summary: Run Golang_ program in your browser by GopherJS_. Access
          `Accept-Language`_ field in `HTTP Request Header`_ by JSONP_ to detect
          browser language preference (user locale_).
:adsu: yes

Introduction
++++++++++++

To detect browser language preference (user locale_), one of the solution is to
access `Accept-Language`_ field in `HTTP Request Header`_. We cannot, however,
access the `Accept-Language`_ field in our browser code directly. One of the
answers in Stack Overflow question [4]_ provides an `cloud service`_ on
`Google App Engine`_ to return you the HTTP request headers via JSONP_. In this
post, we will show you how to write Go_ code which runs on browsers to retrieve
HTTP request headers via JSONP_ by GopherJS_ and its `DOM binding`_.

Note that you can get similar information to detect user locale_ by alternative
solution - browser NavigatorLanguage_ API. This API provides you the same
information of preferred language of the user. See [5]_ for more details.

Install GopherJS_
+++++++++++++++++

Run the following command to install GopherJS_ and
`GopherJS bindings for the JavaScript DOM APIs`_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs
  $ go get -u honnef.co/go/js/dom

Source Code
+++++++++++

HTML file for demo: (*index.html*)

.. code-block:: html

  <!doctype html>
  <html>
  <head>
    <title>Golang - access HTTP Request Header (Accept-Language) by JSONP
           to detect browser language preference</title>
  </head>
  <body>
  <script src="demo.js"></script>
  </body>
  </html>

.. adsu:: 2

The Golang_ code for retrieving HTTP request headers via JSONP:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/http-header/jsonp.go
.. adsu:: 3

Compile Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build jsonp.go -o demo.js

Put *demo.js* together with the *index.html* in the same directory. Open the
*index.html* with your browser, and then open the browser console to see the
result.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``,
``Chromium Version 47.0.2526.106 Ubuntu 15.10 (64-bit)``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_

.. [3] `dom - GopherJS bindings for the JavaScript DOM APIs <https://godoc.org/honnef.co/go/js/dom>`_
       (`GitHub <https://github.com/dominikh/go-js-dom>`__)
.. adsu:: 4
.. [4] `localization - JavaScript for detecting browser language preference - Stack Overflow <http://stackoverflow.com/questions/1043339/javascript-for-detecting-browser-language-preference>`_

.. [5] `[Golang] Detect Browser Language Preference by window.navigator.language <{filename}go-detect-browser-language-preference%en.rst>`_

.. [6] `Navigator - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/Navigator>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _window: http://www.w3schools.com/jsref/obj_window.asp
.. _Accept-Language: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
.. _HTTP Request Header: http://en.wikipedia.org/wiki/List_of_HTTP_header_fields
.. _NavigatorLanguage: https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage
.. _locale: https://en.wikipedia.org/wiki/Locale
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _DOM binding: https://godoc.org/honnef.co/go/js/dom
.. _cloud service: http://ajaxhttpheaders.appspot.com/
.. _Google App Engine: https://cloud.google.com/appengine/docs
.. _JSONP: https://www.google.com/search?q=JSONP

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
