[Golang] JSONP Server Implementation Code
#########################################

:date: 2017-03-18T04:13+08:00
:tags: Go, Golang, JSON, JSONP, CORS, Go net/http, HTTP Header
:category: Go
:summary: A Go server that returns HTTP request headers via JSONP.
:og_image: https://i.ytimg.com/vi/z74PQXHnzjk/hqdefault.jpg
:adsu: yes

My previous posts [4]_ shows how to use JSONP_ in front-end to request data from
a server in different domain. In this post, we will show how to implement the
back-end server that returns data (HTTP request headers) with Go_ programming
language.

You need to have some basic knowledge about writing web applications in Go.
Please refer to [2]_ if you have no idea about Go web applications.

The basic idea of JSONP server is as follows:

1. Get the name of front-end callback function from the query string in HTTP
   request.
2. Encode the data to be returned to the client in JSON_ format. In this case,
   it's the HTTP headers encoded in JSON.
3. Set the ``Content-Type`` as ``application/javascript`` in HTTP response.
4. return the string ``callbackName(JSONString);`` to the client, where
   *callbackName* is the name of callback in query string, and *JSONString* is
   the JSON-encoded data in step 2.

The complete implementation in Go is as follows:

.. adsu:: 2

.. show_github_file:: siongui userpages content/code/go/jsonp-server/server.go
.. adsu:: 3

The front-end code for this JSONP server is as follows:

.. show_github_file:: siongui userpages content/code/go/jsonp-server/index.go
.. adsu:: 4

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.8``
- ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] | `localization - JavaScript for detecting browser language preference - Stack Overflow <http://stackoverflow.com/a/3335420>`_
       | `GitHub - dansingerman/app-engine-headers: Google app engine application that is the server side counterpart to https://github.com/dansingerman/jQuery-Browser-Language <https://github.com/dansingerman/app-engine-headers>`_
.. [2] | `Writing Web Applications - The Go Programming Language <https://golang.org/doc/articles/wiki/>`_
       | `http - The Go Programming Language <https://golang.org/pkg/net/http/>`_
.. [3] `JSON and Go - The Go Blog <https://blog.golang.org/json-and-go>`_
.. [4] | `[JavaScript] JSONP Example <{filename}../16/javascript-jsonp-example%en.rst>`_
       | `JSONP on Google App Engine Python <{filename}../../../2015/02/20/jsonp-on-google-app-engine-python%en.rst>`_
       | `[Golang] JSONP Example (CORS) by GopherJS <{filename}../../../2016/01/23/go-jsonp-example-cors-by-gopherjs%en.rst>`_

.. _JSONP: https://www.google.com/search?q=JSONP
.. _JSON: https://www.google.com/search?q=JSON
.. _Go: https://golang.org/
