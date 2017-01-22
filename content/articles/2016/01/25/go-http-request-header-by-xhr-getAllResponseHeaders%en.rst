[Golang] Access HTTP Request Header by XHR getAllResponseHeaders()
##################################################################

:date: 2016-01-25T05:10+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, XMLHttpRequest, HTTP Header,
       Accept-Language
:category: GopherJS
:summary: Run Golang_ program in your browser by GopherJS_. Access
          `HTTP Request Header`_ by `getAllResponseHeaders()`_ method in
          XMLHttpRequest_ (XHR) request.
:adsu: yes

Introduction
++++++++++++

Access `HTTP Request Header`_ in Golang_ code running on your browser. This can
be done by `getAllResponseHeaders()`_ method in XMLHttpRequest_ (XHR) request,
as discussed in [4]_. According to my test, many important fields such as
`Accept-Language`_, however, are missing, which makes this technique useless. I
suggest you use another technique [6]_ to access HTTP request headers.

Install GopherJS_ and `XHR binding`_
++++++++++++++++++++++++++++++++++++

Run the following command to install GopherJS_ and
`GopherJS bindings for the XMLHttpRequest API`_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs
  $ go get -u honnef.co/go/js/xhr

Source Code
+++++++++++

HTML file for demo: (*index.html*)

.. code-block:: html

  <!doctype html>
  <html>
  <head>
    <title>Golang - access HTTP Request Header
           by XHR getAllResponseHeaders()</title>
  </head>
  <body>
  <script src="demo.js"></script>
  </body>
  </html>

The Golang_ code for retrieving HTTP request headers via
`getAllResponseHeaders()`_:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/http-header/getAllResponseHeaders.go

Compile Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build getAllResponseHeaders.go -o demo.js

Put *demo.js* together with the *index.html* in the same directory. You need a
simple HTTP server to run this demo. Use `GopherJS serve command`_ to serve the
above files, and open your browser console to see the output. My output is:

.. code-block:: txt

  Date: Sun, 24 Jan 2016 20:11:53 GMT
  Last-Modified: Fri, 22 Jan 2016 12:59:37 GMT
  Accept-Ranges: bytes
  Content-Length: 145
  Content-Type: text/html; charset=utf-8

----

Appendix
++++++++

If you want to use GopherJS_ native API only without `XHR binding`_, you can use
the following code:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/http-header/getAllResponseHeaders-raw.go

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

.. [3] `Package xhr provides GopherJS bindings for the XMLHttpRequest API <https://godoc.org/honnef.co/go/js/xhr>`_
       (`GitHub <https://github.com/dominikh/go-js-xhr>`__)

.. [4] `Accessing the web page's HTTP Headers in JavaScript - Stack Overflow <http://stackoverflow.com/questions/220231/accessing-the-web-pages-http-headers-in-javascript>`_

.. [5] `Using XMLHttpRequest - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest/Using_XMLHttpRequest>`_

.. [6] `[Golang] Access HTTP Request Header (Accept-Language) by JSONP <{filename}../24/go-http-request-header-by-jsonp-gopherjs%en.rst>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _Accept-Language: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
.. _HTTP Request Header: http://en.wikipedia.org/wiki/List_of_HTTP_header_fields
.. _GopherJS bindings for the XMLHttpRequest API: https://godoc.org/honnef.co/go/js/xhr
.. _XHR binding: https://godoc.org/honnef.co/go/js/xhr
.. _XMLHttpRequest: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
.. _getAllResponseHeaders(): https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest#getAllResponseHeaders()
.. _GopherJS serve command: {filename}../10/gopherjs-serve-and-build-command-usage%en.rst

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
