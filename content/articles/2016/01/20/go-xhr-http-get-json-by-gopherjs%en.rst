[Golang] XMLHttpRequest (XHR) HTTP GET JSON Data by GopherJS
############################################################

:date: 2016-01-20T03:30+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, HTTP GET, XMLHttpRequest, JSON,
       Go net/http
:category: GopherJS
:summary: Run Golang_ program in your browser by GopherJS_. Use XMLHttpRequest_
          (XHR) to send `HTTP GET`_ requests to retrieve JSON_ data from remote
          server.


Introduction
++++++++++++

It is really cool to run Go_ code in the browser. GopherJS_ is a compiler from
Go_ to JavaScript_, which makes this possible.
In this post, we will show how to use XMLHttpRequest_ (XHR) to send `HTTP GET`_
requests to retrieve JSON_ data from remote server.

Install GopherJS_
+++++++++++++++++

Run the following command to install GopherJS_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs

Source Code
+++++++++++

First we write a simple HTML for our demo: (*index.html*)

.. code-block:: html

  <!doctype html>
  <html>
  <head>
    <title>XHR HTTP Get by GopherJS</title>
  </head>
  <body>
  <script src="demo.js"></script>
  </body>
  </html>

The following is the JSON_ data to be retrieved by `HTTP GET`_ request:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/xhrget/sukhada.json

It is surprising easy to send `HTTP GET`_ XHR_ request: Use Golang_ built-in
`net/http`_ package! You just use *Get* method as usual, and GopherJS will take
care of all the rests for you!

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/xhrget/xhrget2.go

Compile the Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build xhrget.go -o demo.js

Put *demo.js* together with the *index.html* and *sukhada.json* in the same
directory. You need a simple HTTP server to run this demo. Use
`GopherJS serve command`_ to serve the above files, and open your browser
console to see the output.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `Bindings · gopherjs/gopherjs Wiki · GitHub <https://github.com/gopherjs/gopherjs/wiki/bindings>`_

.. [3] `xhr - GoDoc <https://godoc.org/honnef.co/go/js/xhr>`_

.. [4] `http - The Go Programming Language <https://golang.org/pkg/net/http/>`_

.. [5] `delete xhr/transport, GopherJS has its own now. · dominikh/go-js-xhr@00e3346 · GitHub <https://github.com/dominikh/go-js-xhr/commit/00e3346113aed89b501ead4e863c7c3d04fa0c5b>`_

.. [6] `go json decode array <https://www.google.com/search?q=go+json+decode+array>`_

.. [7] `How to Unmarshal a JSON Array of Arrays in Go - Fabio Berger <http://fabioberger.com/blog/2014/10/09/how-to-unmarshal-a-json-array-of-arrays-in-go/>`_

.. [8] `encoding/json - The Go Programming Language <https://golang.org/pkg/encoding/json/>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _XMLHttpRequest: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
.. _XHR: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
.. _HTTP GET: http://www.w3schools.com/tags/ref_httpmethods.asp
.. _JSON: http://www.w3schools.com/json/
.. _net/http: https://golang.org/pkg/net/http/
.. _GopherJS serve command: {filename}../10/gopherjs-serve-and-build-command-usage%en.rst

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
