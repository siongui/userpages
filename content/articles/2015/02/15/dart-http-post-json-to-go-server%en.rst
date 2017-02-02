[Webapp] Dart HTTP POST JSON Data to Go Server
##############################################

:tags: Dart, Go, Golang, HTTP POST, Web application, JSON, Makefile,
       XMLHttpRequest, Go net/http
:category: Go
:summary: Show how to pass JSON_-format data using `HTTP POST`_ between browser
          (implemented with Dart_) and web server (implemented with Go_).
:adsu: yes


This post shows how to write a web application which passes JSON_-format data
between browser (client side) and web server (server side) via `HTTP POST`_.
The code which runs on the browser is implemented with Dart_, and the code which
runs on web server is implemented with Go_. I put explanation and references
directly in the code for the convenience of fast lookup and tracing code.

To run the code, download the following files in the same directory. Modify the
path in Makefile and type ``make`` to execute. (If you do not have Dartium_,
type ``make js`` before you type ``make``). Then open your browser at
``http://localhost:8000/``.

Development Environment: ``Ubuntu Linux 14.10``, ``Dart 1.8``, ``Go 1.4``.

Server side (*Go* web server):

.. show_github_file:: siongui userpages content/code/go-dart-http-post-json/server.go
.. adsu:: 2

Server side (*HTML* template for Go_):

.. show_github_file:: siongui userpages content/code/go-dart-http-post-json/index.html

Client side (*Dart*):

.. adsu:: 3
.. show_github_file:: siongui userpages content/code/go-dart-http-post-json/app.dart

Makefile for automating the development:

.. show_github_file:: siongui userpages content/code/go-dart-http-post-json/Makefile

----

References:

.. [1] `Writing Web Applications - The Go Programming Language <https://golang.org/doc/articles/wiki/>`_

.. [2] `Handling JSON Post Request in Go <http://stackoverflow.com/questions/15672556/handling-json-post-request-in-go>`_

.. [3] `Using Dart with JSON Web Services <https://www.dartlang.org/articles/json-web-service/>`_

.. _Dart: https://www.dartlang.org/
.. _JSON: https://www.google.com/search?q=JSON
.. _Go: https://golang.org/
.. _HTTP POST: https://www.google.com/search?q=HTTP+POST
.. _Dartium: https://www.dartlang.org/tools/dartium/
