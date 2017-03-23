[Golang] Get Query String of URL in net/http Handler
####################################################

:date: 2017-03-24T05:11+08:00
:tags: Go, Golang, Go net/http, Go net/url
:category: Go
:summary: Get query string of URL in request handler of HTTP server via Go
          standard *net/http* Package.
:og_image: https://i.ytimg.com/vi/z74PQXHnzjk/hqdefault.jpg
:adsu: yes


Get `query string`_ of URL in request handler of HTTP server via Go standard
`net/http`_ Package.

**Question**:

Assume the URL of the HTTP request is

::

  https://example.com/?name=john

I want to get the value of ``name``, i.e., ``john`` from the URL. How to get it
in HTTP request handler of *net/http*?

**Answer**:

.. code-block:: go

  import (
  	"net/http"
  )

  func handler(w http.ResponseWriter, r *http.Request) {
  	name := r.URL.Query().Get("name")
  	// the value of name is john
  }

For more information about how to access the query string, see `net/url`_
package.

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.8``

----

References:

.. [1] | `Go sever get query string in url - Google search <https://www.google.com/search?q=Go+sever+get+query+string+in+url>`_
       | `Go sever get query string in url - DuckDuckGo search <https://duckduckgo.com/?q=Go+sever+get+query+string+in+url>`_
       | `Go sever get query string in url - Ecosia search <https://www.ecosia.org/search?q=Go+sever+get+query+string+in+url>`_
       | `Go sever get query string in url - Qwant search <https://www.qwant.com/?q=Go+sever+get+query+string+in+url>`_
       | `Go sever get query string in url - Bing search <https://www.bing.com/search?q=Go+sever+get+query+string+in+url>`_
       | `Go sever get query string in url - Yahoo search <https://search.yahoo.com/search?p=Go+sever+get+query+string+in+url>`_
       | `Go sever get query string in url - Baidu search <https://www.baidu.com/s?wd=Go+sever+get+query+string+in+url>`_
       | `Go sever get query string in url - Yandex search <https://www.yandex.com/search/?text=Go+sever+get+query+string+in+url>`_

.. [2] `[Golang] JSONP Server Implementation Code <{filename}../18/go-jsonp-server-implementation-code%en.rst>`_

.. _Go: https://golang.org/
.. _query string: https://www.google.com/search?q=query+string
.. _net/http: https://golang.org/pkg/net/http/
.. _net/url: https://golang.org/pkg/net/url/
