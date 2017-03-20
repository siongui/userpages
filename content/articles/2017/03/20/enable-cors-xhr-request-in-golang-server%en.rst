Enable Cross Domain XMLHttpRequest Request in Go Server
#######################################################

:date: 2017-03-20T22:51+08:00
:tags: Go, Golang, Go net/http, CORS
:category: Go
:summary: Enable cross-domain XHR requests in Go server via Go standard
          *net/http* package.
:og_image: https://i.ytimg.com/vi/z74PQXHnzjk/hqdefault.jpg
:adsu: yes


Howto
+++++

To enable cross-domain XHR requests in servers, we need to set the following
header in request response:

.. code-block:: txt

  Access-Control-Allow-Origin: *

The above header says all requests from any other domain are allowed.

To set the ``Access-Control-Allow-Origin`` header in Go server via standard
`net/http`_ package, we can use the following code in request handler:

.. code-block:: go

  func myHandler(w http.ResponseWriter, r *http.Request) {
  	// allow all XHR requests from other domains
  	w.Header().Set("Access-Control-Allow-Origin", "*")
  }

Just one line of code to make your Go server enable
*Cross-Origin Resource Sharing (CORS)*.

.. adsu:: 2


Complete Example
++++++++++++++++

The following is complete sample code for servers supporting CORS and running on
`Google App Engine Go Standard Environment`_

.. show_github_file:: siongui userpages content/code/go/cors-server/serverjson.go
.. adsu:: 3
.. show_github_file:: siongui userpages content/code/go/cors-server/app.yaml
.. show_github_file:: siongui userpages content/code/go/cors-server/Makefile

----

Tested on:

- ``Google App Engine Go 1.6``

----

References:

.. [1] | `go net/http set response header - Google search <https://www.google.com/search?q=go+net/http+set+response+header>`_
       | `go net/http set response header - DuckDuckGo search <https://duckduckgo.com/?q=go+net/http+set+response+header>`_
       | `go net/http set response header - Ecosia search <https://www.ecosia.org/search?q=go+net/http+set+response+header>`_
       | `go net/http set response header - Qwant search <https://www.qwant.com/?q=go+net/http+set+response+header>`_
       | `go net/http set response header - Bing search <https://www.bing.com/search?q=go+net/http+set+response+header>`_
       | `go net/http set response header - Yahoo search <https://search.yahoo.com/search?p=go+net/http+set+response+header>`_
       | `go net/http set response header - Baidu search <https://www.baidu.com/s?wd=go+net/http+set+response+header>`_
       | `go net/http set response header - Yandex search <https://www.yandex.com/search/?text=go+net/http+set+response+header>`_
.. adsu:: 4
.. [2] `Setting HTTP headers in Golang - Stack Overflow <http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang>`_

.. _net/http: https://golang.org/pkg/net/http/
.. _Google App Engine Go Standard Environment: https://cloud.google.com/appengine/docs/standard/go/
