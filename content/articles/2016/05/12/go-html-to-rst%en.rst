[Golang] HTML to reStructuredText
#################################

:date: 2016-05-12T21:56+08:00
:tags: Go, Golang, DOM, Web Scrape, reStructuredText, Go net/html, html,
       Go net/http
:category: Go
:summary: Convert HTML_ to restructuredtext_ format via `net/html`_ package
          in Go_ programming language. (Not fully supported)
:adsu: yes


Introduction
++++++++++++

Based on my previous work [7]_, convert HTML_ to restructuredtext_
in Go_ via `net/html`_ package.

The following HTML node/element is supported:

- `HTML unordered (bulleted) list`_ (*ul* *li*)
- `HTML ordered list`_ (*ol* *li*)
- `HTML link`_ (*a*)
- `HTML image`_ (*img*)
- `HTML thematic break`_ (*hr*)
- `HTML table`_ (*table*, *tbody*, *tr*, *td*)
- `HTML text node`_
- `HTML comment node`_


Install `html2rst`_ Package
+++++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u github.com/siongui/html2rst


Usage/Example
+++++++++++++

.. show_github_file:: siongui html2rst usage/example.go

----

Tested on: ``Ubuntu Linux 16.04``, ``Go 1.6.2``.

----

References:

.. [1] `html - GoDoc <https://godoc.org/golang.org/x/net/html>`_

.. [2] `goquery - Replace HTML Link Node with reStructuredText Text Node <{filename}../04/goquery-replace-html-link-node-with-rst-text-node%en.rst>`_

.. [3] `[Golang] Read Lines From File or String <{filename}../../04/06/go-readlines-from-file-or-string%en.rst>`_

.. [4] `go string concat - Google search <https://www.google.com/search?q=go+string+concat>`_

       `How to efficiently concatenate strings in Go? - Stack Overflow <http://stackoverflow.com/a/1763606>`_

.. [5] `goquery - Convert HTML Unordered List to reStructuredText <{filename}../05/goquery-html-ul-li-to-rst%en.rst>`_

.. [6] `[Golang] HTML a, ul, li Element to reStructuredText <{filename}../07/go-html-a-ul-li-to-rst%en.rst>`_

.. [7] `[Golang] HTML a, img, ul, li Element to reStructuredText <{filename}../08/go-html-a-img-ul-li-to-rst%en.rst>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _HTML unordered (bulleted) list: http://www.w3schools.com/tags/tag_ul.asp
.. _HTML ordered list: https://github.com/siongui/html2rst
.. _HTML link: http://www.w3schools.com/html/html_links.asp
.. _HTML image: http://www.w3schools.com/html/html_images.asp
.. _HTML thematic break: http://www.w3schools.com/tags/tag_hr.asp
.. _HTML table: http://www.w3schools.com/html/html_tables.asp
.. _HTML text node: https://www.google.com/search?q=html+text+node
.. _HTML comment node: https://www.google.com/search?q=html+comment+node
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _HTML: https://www.google.com/search?q=HTML
.. _net/html: https://godoc.org/golang.org/x/net/html
.. _html2rst: https://github.com/siongui/html2rst
