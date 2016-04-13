[Golang] HTML Table to reStructuredText list-table via goquery
##############################################################

:date: 2016-04-14T07:33+08:00
:tags: Go, Golang, DOM, goquery, Web Scrape, html, reStructuredText
:category: Go
:summary: Convert `HTML table`_ to reStructuredText_ `list-table`_ via goquery_
          in Go_ programming language.


Introduction
++++++++++++

Convert `HTML table`_ to reStructuredText_ `list-table`_ via goquery_ in Golang_
(Go_ programming language). For unrobust implementation via Go_ `net/html`_
package, see [5]_. For Python_ `Beautiful Soup 4`_ (bs4_) implementation, see
[4]_.

Install goquery_
++++++++++++++++

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery

`HTML table`_ to reStructuredText_ `list-table`_
++++++++++++++++++++++++++++++++++++++++++++++++

.. show_github_file:: siongui userpages content/code/go-html-table-to-rst/goquery.go

Usage:

.. show_github_file:: siongui userpages content/code/go-html-table-to-rst/goquery_test.go

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `jquery iterate over elements - Google search <https://www.google.com/search?q=jquery+iterate+over+elements>`_

.. [2] `net/html go - Google search <https://www.google.com/search?q=net/html+go>`_

       `A Simple Web Scraper in Go | Gregory Schier <http://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html>`_

       `golang.org/x/net/html GoDoc <https://godoc.org/golang.org/x/net/html>`_

.. [3] `github.com/PuerkitoBio/goquery - GoDoc <https://godoc.org/github.com/PuerkitoBio/goquery>`_

.. [4] `[Python] Convert HTML Table to reStructuredText list-table <{filename}../../02/28/python-convert-html-table-to-rst-list-table%en.rst>`_

.. [5] `[Golang] Unrobust HTML Table to reStructuredText list-table <{filename}../13/go-unrobust-html-table-to-rst-list-table%en.rst>`_

.. [6] `html table to rst list-table · twnanda/twnanda@e022835 · GitHub <https://github.com/twnanda/twnanda/commit/e022835fdddd3282588f38304c649ad71d73476b>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _net/html: https://godoc.org/golang.org/x/net/html
.. _golang.org/x/net/html: https://godoc.org/golang.org/x/net/html
.. _DOM: https://www.google.com/search?q=DOM
.. _HTML: https://www.google.com/search?q=HTML
.. _HTML links: http://www.w3schools.com/html/html_links.asp
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _Python: https://www.python.org/
.. _list-table: http://docutils.sourceforge.net/docs/ref/rst/directives.html#list-table
.. _bs4: http://www.crummy.com/software/BeautifulSoup/bs4/doc/
.. _Beautiful Soup 4: http://www.crummy.com/software/BeautifulSoup/bs4/doc/
.. _HTML table: http://www.w3schools.com/html/html_tables.asp
