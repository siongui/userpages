[Golang] Iterate over All DOM Elements in HTML
##############################################

:date: 2016-04-10T20:35+08:00
:modified: 2016-04-15T21:27+08:00
:tags: Go, Golang, Commandline, DOM, Go net/html, Web Scrape, Go flag Package,
       html, reStructuredText, File Input/Output
:category: Go
:summary: Iterate over all DOM_ elements in HTML_ via Go_ programming language.
          Use `net/html`_ package to parse and iterate all elements in HTML.
          Search for `HTML links`_ and output them in reStructuredText_ format.


Introduction
++++++++++++

Iterate over all DOM_ elements in HTML_ via Golang_. Use `net/html`_ package to
parse and iterate all elements in HTML. Search for `HTML links`_ and output them
in reStructuredText_ format.

Another example of iterating over all DOM_ elements can be found in [4]_.

Install `net/html`_ package
+++++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u golang.org/x/net/html

Traverse DOM_ Tree
++++++++++++++++++

Traverse the DOM_ tree (Iterate over all elements in HTML):

.. show_github_file:: siongui userpages content/code/go-iterate-all-elements/html.go

Find `HTML links`_ and print them in reStructuredText_ format:

.. show_github_file:: siongui userpages content/code/go-iterate-all-elements/handleHtmlLink.go

Usage
+++++

Download any HTML file and pass the file path to Go_ program by ``input`` flag.
For example, if you have ``index.html`` put together with Go_ program in current
directory, run the program by the following command:

.. code-block:: bash

  $ go run html.go handleHtmlLink.go -input=index.html

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `jquery iterate over elements - Google search <https://www.google.com/search?q=jquery+iterate+over+elements>`_

.. [2] `net/html go - Google search <https://www.google.com/search?q=net/html+go>`_

       `A Simple Web Scraper in Go | Gregory Schier <http://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html>`_

       `golang.org/x/net/html GoDoc <https://godoc.org/golang.org/x/net/html>`_

.. [3] `github.com/PuerkitoBio/goquery - GoDoc <https://godoc.org/github.com/PuerkitoBio/goquery>`_

.. [4] `[Golang] getElementById via net/html Package <{filename}../15/go-getElementById-via-net-html-package%en.rst>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _net/html: https://godoc.org/golang.org/x/net/html
.. _golang.org/x/net/html: https://godoc.org/golang.org/x/net/html
.. _DOM: https://www.google.com/search?q=DOM
.. _HTML: https://www.google.com/search?q=HTML
.. _HTML links: http://www.w3schools.com/html/html_links.asp
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
