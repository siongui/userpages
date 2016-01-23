[Golang] JSONP Example (CORS) by GopherJS
#########################################

:date: 2016-01-23T02:45+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, JSON, JSONP, CORS, HTTP GET, DOM,
       Web application, web, Full-Stack Go, Full-Stack Golang
:category: GopherJS
:summary: Run Golang_ program in your browser by GopherJS_.
          Show how to make cross-domain requests (CORS_) by JSONP_ (JSON with
          Padding) technique, which allows data to be retrieved from servers of
          other domains.
          This is an example of full-stack Go_, which uses Golang_ to
          develop web applications in both front-end and backend.


Introduction
++++++++++++

It is really cool to run Go_ code in the browser. GopherJS_ is a compiler from
Go_ to JavaScript_, which makes this possible.
In this post, we will show how to make cross-domain requests (CORS_) by JSONP_
(JSON with Padding) technique, which allows data to be retrieved from servers of
other domains.
This is an example of full-stack Go_, which uses Golang_ to develop web
applications in both front-end (runs on browsers) and backend (runs on servers).

Install GopherJS_
+++++++++++++++++

Run the following command to install GopherJS_ and
`GopherJS bindings for the JavaScript DOM APIs`_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs
  $ go get -u honnef.co/go/js/dom

Source Code
+++++++++++

First we write a simple HTML for our demo:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/jsonp/index.html

A callback function whose name is *mycallback* are declared by *js.Global.Set*
method. The *mycallback* function will receive JSON_ data from the server.
Beside, a *script* element are inserted to the *head* element to make
cross-domain request.

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/jsonp/jsonp.go

The server receive the name of the callback function and data from the client.
Then encode the data in JSON_ format, and send the JSON_ data to the callback
function.

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/jsonp/server.go

Compile the Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build jsonp.go -o demo.js

Put *demo.js* together with the *index.html* and in the same directory. Modify
the path in *server.go* to the path where you place *demo.js* and *index.html*.
Run the server by:

.. code-block:: bash

  $ go run server.go

Open your browser with URL `localhost:8000 <http://localhost:8000/>`_. Open the
console of the browser, and you will see the data printed out by the callback
function.

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

.. [4] `JSONP on Google App Engine Python <{filename}../../../2015/02/20/jsonp-on-google-app-engine-python%en.rst>`_

.. [5] `golang encodeuricomponent <https://www.google.com/search?q=golang+encodeuricomponent>`_

.. [6] `escaping - Recommended way to encode/decode URLs - Stack Overflow <http://stackoverflow.com/questions/13826808/recommended-way-to-encode-decode-urls>`_

.. [7] `golang get url from request <https://www.google.com/search?q=golang+get+url+from+request>`_

.. [8] `In go's http package, how do I get the query string on a POST request? - Stack Overflow <http://stackoverflow.com/questions/15407719/in-gos-http-package-how-do-i-get-the-query-string-on-a-post-request>`_

.. [9] `[Webapp] Dart HTTP POST JSON Data to Go Server <{filename}../../../2015/02/15/dart-http-post-json-to-go-server%en.rst>`_

.. [10] `golang http write response <https://www.google.com/search?q=golang+http+write+response>`_

.. [11] `Writing Web Applications - The Go Programming Language <https://golang.org/doc/articles/wiki/#tmp_4>`_

.. [12] `http - The Go Programming Language <https://golang.org/pkg/net/http/>`_

.. [13] `encoding/json - The Go Programming Language <https://golang.org/pkg/encoding/json/>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _GopherJS bindings for the JavaScript DOM APIs: https://godoc.org/honnef.co/go/js/dom
.. _JSON: http://www.w3schools.com/json/
.. _JSONP: https://www.google.com/search?q=JSONP
.. _CORS: https://www.google.com/search?q=CORS

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
