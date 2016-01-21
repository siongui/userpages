[Golang] XMLHttpRequest (XHR) HTTP POST JSON Data by GopherJS
#############################################################

:date: 2016-01-21T11:52+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, HTTP POST, XMLHttpRequest, JSON,
       Web application, web, Full-Stack Go, Full-Stack Golang
:category: Go
:summary: Run Golang_ program in your browser by GopherJS_. Use XMLHttpRequest_
          (XHR) to send `HTTP POST`_ requests to send JSON_ data to remote
          server. This is an example of full-stack Go_, which uses Golang_ to
          develop web applications in both front-end and backend.


Introduction
++++++++++++

It is really cool to run Go_ code in the browser. GopherJS_ is a compiler from
Go_ to JavaScript_, which makes this possible.
In this post, we will show how to use XMLHttpRequest_ (XHR) to make `HTTP POST`_
requests to send JSON_ data to remote server.
This is an example of full-stack Go_, which uses Golang_ to develop web
applications in both front-end (runs on browsers) and backend (runs on servers).

Install GopherJS_
+++++++++++++++++

Run the following command to install GopherJS_:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs

Source Code
+++++++++++

First we write a simple HTML for our demo:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/xhrpost/index.html

It is surprising easy to send `HTTP POST`_ XHR_ request: Use Golang_ built-in
`net/http`_ package! You just use *Post* method as usual, and GopherJS will take
care of all the rests for you!

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/xhrpost/xhrpost.go

The following is backend Go_ code for receiving JSON_ data from front-end by
`HTTP POST`_ method:

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/xhrpost/server.go


Compile the Go_ code to JavaScript_ by:

.. code-block:: bash

  $ gopherjs build xhrpost.go -o demo.js

Put *demo.js* together with the *index.html* and in the same directory. Modify
the path in *server.go* to the path where you place *demo.js* and *index.html*.
Run the server by:

.. code-block:: bash

  $ go run server.go

Open your browser with URL `localhost:8000 <http://localhost:8000/>`_. You will
see the console running *server.go* prints out data received from front-end.

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

.. [6] `encoding/json - The Go Programming Language <https://golang.org/pkg/encoding/json/>`_

.. [7] `[Webapp] Dart HTTP POST JSON Data to Go Server <{filename}../../../2015/02/15/dart-http-post-json-to-go-server%en.rst>`_

.. [8] `golang static file server <https://www.google.com/search?q=golang+static+file+server>`_

.. [9] `How do you serve a static html file using a go web server? - Stack Overflow <http://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server>`_

.. [10] `golang http post <https://www.google.com/search?q=golang+http+post>`_

.. [11] `rest - Go lang - How send json string in POST request - Stack Overflow <http://stackoverflow.com/questions/24455147/go-lang-how-send-json-string-in-post-request>`_

.. [12] `strings - The Go Programming Language <https://golang.org/pkg/strings/#NewReader>`_

.. [13] `bytes - The Go Programming Language <https://golang.org/pkg/bytes/#NewReader>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://en.wikipedia.org/wiki/JavaScript
.. _XMLHttpRequest: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
.. _XHR: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
.. _HTTP POST: http://www.w3schools.com/tags/ref_httpmethods.asp
.. _JSON: http://www.w3schools.com/json/
.. _net/http: https://golang.org/pkg/net/http/
.. _GopherJS serve command: {filename}../10/gopherjs-serve-and-build-command-usage%en.rst

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
