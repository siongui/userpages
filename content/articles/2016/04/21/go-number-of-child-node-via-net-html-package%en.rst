[Golang] Number of Child Nodes via net/html Package
###################################################

:date: 2016-04-21T20:54+08:00
:tags: Go, Golang, DOM, Go net/html, Web Scrape, html
:category: Go
:summary: Get the number of child nodes via Go_ `net/html`_ package.
:adsu: yes


Introduction
++++++++++++

Get the number of child nodes via Go_ `net/html`_ package. Equivalent to the
following JavaScript_ code:

.. code-block:: javascript

  var c = document.getElementById("foo").childNodes.length;

Install `net/html`_ package
+++++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u golang.org/x/net/html

Number of Child Nodes
+++++++++++++++++++++

.. code-block:: go

  package getelementbyid

  import (
  	"golang.org/x/net/html"
  )

  func numberOfChild(n *html.Node) int {
  	if n == nil {
  		return -1
  	}

  	count := 0
  	for c := n.FirstChild; c != nil; c = c.NextSibling {
  		count += 1
  	}

  	return count
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6.1``.

----

References:

.. [1] `jquery iterate over elements - Google search <https://www.google.com/search?q=jquery+iterate+over+elements>`_

.. [2] | `net/html go - Google search <https://www.google.com/search?q=net/html+go>`_
       | `A Simple Web Scraper in Go | Gregory Schier <http://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html>`_
       | `golang.org/x/net/html GoDoc <https://godoc.org/golang.org/x/net/html>`_

.. [3] `github.com/PuerkitoBio/goquery - GoDoc <https://godoc.org/github.com/PuerkitoBio/goquery>`_

.. [4] `[Golang] Iterate over All DOM Elements in HTML <{filename}../10/go-iterate-over-all-dom-elements-in-html%en.rst>`_

.. [5] `[Golang] getElementById via net/html Package <{filename}../15/go-getElementById-via-net-html-package%en.rst>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _net/html: https://godoc.org/golang.org/x/net/html
.. _JavaScript: https://www.google.com/search?q=JavaScript
