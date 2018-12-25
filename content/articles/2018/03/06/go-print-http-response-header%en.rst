[Golang] Get HTTP Response Header
#################################

:date: 2018-03-06T23:00+08:00
:tags: Go, Golang, Go net/http, HTTP GET, HTTP Header
:category: Go
:summary: Print HTTP response header via Go standard *net/http* package.
:og_image: https://www.drupal.org/files/project-images/http-response-headers.png
:adsu: yes


Example of printing HTTP response header via Go standard `net/http`_ package.
We will *HTTP Get* this site and print the header of HTTP response.

.. code-block:: go

  package main

  import (
  	"fmt"
  	"net/http"
  )

  func main() {
  	resp, err := http.Get("https://siongui.github.io/")
  	if err != nil {
  		panic(err)
  	}
  	defer resp.Body.Close()

  	for k, v := range resp.Header {
  		fmt.Print(k)
  		fmt.Print(" : ")
  		fmt.Println(v)
  	}
  }

The header is stored in the Header_ type in `http.Response`_ type.

Result:

.. code-block:: txt

  Age : [0]
  X-Served-By : [cache-itm18826-ITM]
  X-Cache-Hits : [0]
  Vary : [Accept-Encoding]
  Last-Modified : [Mon, 05 Mar 2018 14:43:07 GMT]
  X-Github-Request-Id : [65D8:1D40:D8AC39:E9D3C0:5A9EAA7A]
  Accept-Ranges : [bytes]
  X-Cache : [MISS]
  X-Fastly-Request-Id : [72ee00a8726c6a65793ae7fcd0ee317af97269a6]
  Server : [GitHub.com]
  Expires : [Tue, 06 Mar 2018 14:59:30 GMT]
  Access-Control-Allow-Origin : [*]
  Via : [1.1 varnish]
  X-Timer : [S1520348411.289800,VS0,VE130]
  Date : [Tue, 06 Mar 2018 15:00:11 GMT]
  Content-Type : [text/html; charset=utf-8]
  Cache-Control : [max-age=600]

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``.

----

References:

.. [1] `Package http - The Go Programming Language <https://golang.org/pkg/net/http/>`_
.. [2] `Pitfall of Golang header operation : golang <https://old.reddit.com/r/golang/comments/a99ajn/pitfall_of_golang_header_operation/>`_

.. _net/http: https://golang.org/pkg/net/http/
.. _http.Response: https://golang.org/pkg/net/http/#Response
.. _Header: https://golang.org/pkg/net/http/#Header
