goquery - Replace HTML Link Node with reStructuredText Text Node
################################################################

:date: 2016-05-04T21:01+08:00
:tags: Go, Golang, DOM, goquery, Web Scrape, reStructuredText, Golang template,
       html
:category: Go
:summary: Replace `HTML links`_ node with `text node`_ of restructuredtext_
          format in a webpage via goquery_ in Go_ programming language.
:adsu: yes


Introduction
++++++++++++

Replace `HTML links`_ node with `text node`_ of restructuredtext_ format in a
webpage via goquery_ in Golang_ (Go_ programming language).


Install goquery_ Package
++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery


Source Code
+++++++++++

.. show_github_file:: siongui userpages content/code/goquery-replace-link-node-with-rst-text-node/link2rst.go
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/goquery-replace-link-node-with-rst-text-node/link2rst_test.go
.. adsu:: 3

----

Tested on: ``Ubuntu Linux 16.04``, ``Go 1.6.2``.

----

References:

.. [1] `github.com/PuerkitoBio/goquery - GoDoc <https://godoc.org/github.com/PuerkitoBio/goquery>`_

.. [2] `[Golang] Convert All HTML Links to reStructuredText via goquery <{filename}../../04/12/go-html-links-to-rst-via-goquery%en.rst>`_

.. [3] `[Golang] Extract Title, Image, and URL via goquery <{filename}../../03/31/go-parse-buy123-webpage-to-rst%en.rst>`_

.. [4] `html escape < - Google search <https://www.google.com/search?q=html+escape+%3C>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _HTML links: http://www.w3schools.com/html/html_links.asp
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _text node: http://www.w3schools.com/jsref/met_document_createtextnode.asp
