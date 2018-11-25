[Golang] Get Domain Name from URL
#################################

:date: 2018-01-31T04:51+08:00
:tags: Go, Golang, Go net/url
:category: Go
:summary: Get domain name from URL via Go standard *net/url* package.
:og_image: http://visionlaunch.com/wp-content/uploads/2016/09/inseev-1.jpg
:adsu: yes


Get `domain name`_ from URL via Go `net/url`_ package in standard library.

**Question**:

Assume the URL is

::

  https://siongui.github.io/pali-chanting/zh/archives.html

I want to get the domain name ``github.io`` from the URL. How to do it in Go?

**Answer**:

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/NAHTuNuYIT4>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"log"
  	"net/url"
  	"strings"
  )

  func main() {
  	u, err := url.Parse("https://siongui.github.io/pali-chanting/zh/archives.html")
  	if err != nil {
  		log.Fatal(err)
  	}
  	parts := strings.Split(u.Hostname(), ".")
  	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]
  	fmt.Println(domain)
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] | `golang get domain from url - Google search <https://www.google.com/search?q=golang+get+domain+from+url>`_
       | `golang get domain from url - DuckDuckGo search <https://duckduckgo.com/?q=golang+get+domain+from+url>`_
       | `golang get domain from url - Ecosia search <https://www.ecosia.org/search?q=golang+get+domain+from+url>`_
       | `golang get domain from url - Qwant search <https://www.qwant.com/?q=golang+get+domain+from+url>`_
       | `golang get domain from url - Bing search <https://www.bing.com/search?q=golang+get+domain+from+url>`_
       | `golang get domain from url - Yahoo search <https://search.yahoo.com/search?p=golang+get+domain+from+url>`_
       | `golang get domain from url - Baidu search <https://www.baidu.com/s?wd=golang+get+domain+from+url>`_
       | `golang get domain from url - Yandex search <https://www.yandex.com/search/?text=golang+get+domain+from+url>`_
.. [2] `Find DNS records of a domain using Golang : golang <https://old.reddit.com/r/golang/comments/a0981x/find_dns_records_of_a_domain_using_golang/>`_

.. _domain name: https://www.google.com/search?q=domain+name
.. _net/url: https://golang.org/pkg/net/url/
.. _Go Playground: https://play.golang.org/
