[Golang] Detect Browser Language Preference by window.navigator.language
########################################################################

:date: 2016-01-24T04:06+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript, Locale, i18n
:category: GopherJS
:summary: Run Golang_ program in your browser by GopherJS_. Access
          `window.navigator.language`_ of NavigatorLanguage_ API to detect
          browser language preference (user locale_).

Introduction
++++++++++++

To detect browser language preference (user locale_), one of the solution is to
access `Accept-Language`_ field in `HTTP Request Header`_. We cannot, however,
access the `Accept-Language`_ in our browser code. One of the answers in Stack
Overflow question [4]_ provides an alternative - use NavigatorLanguage_ API.
This API provides you the same information of preferred language of the user.
This post shows you how to access `window.navigator.language`_ and
`window.navigator.languages`_ in Go_.

Install GopherJS_
+++++++++++++++++

Run the following command to install GopherJS_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs

Source Code
+++++++++++

HTML file for demo: (*index.html*)

.. code-block:: html

  <!doctype html>
  <html>
  <head>
    <title>Golang - access window.navigator.language
           to detect browser language preference</title>
  </head>
  <body>
  <script src="demo.js"></script>
  </body>
  </html>

*js.Global* object is equivalent to window_ object in browsers. The *Get* method
is used to access the property of the object.

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/navigator/language.go

Compile Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build language.go -o demo.js

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

.. [4] `localization - JavaScript for detecting browser language preference - Stack Overflow <http://stackoverflow.com/questions/1043339/javascript-for-detecting-browser-language-preference>`_

.. [5] `Navigator - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/Navigator>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _window: http://www.w3schools.com/jsref/obj_window.asp
.. _Accept-Language: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
.. _HTTP Request Header: http://en.wikipedia.org/wiki/List_of_HTTP_header_fields
.. _NavigatorLanguage: https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage
.. _window.navigator.language: https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage/language
.. _window.navigator.languages: https://developer.mozilla.org/en-US/docs/Web/API/NavigatorLanguage/languages
.. _locale: https://en.wikipedia.org/wiki/Locale

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
