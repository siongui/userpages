[Golang] HTML a, ul, li Element to reStructuredText
###################################################

:date: 2016-05-07T22:12+08:00
:tags: Go, Golang, DOM, Web Scrape, reStructuredText, Go net/html, html
:category: Go
:summary: Convert `HTML unordered (bulleted) list`_ and `HTML link`_ to
          restructuredtext_ format via `net/html`_ package in Go_ programming
          language.


Introduction
++++++++++++

Convert `HTML unordered (bulleted) list`_ and `HTML link`_ to restructuredtext_
format via `net/html`_ package in Golang_ (Go_ programming language).
I am not sure whether the rst output can be converted back to HTML or not.


Install `net/html`_ Package
+++++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u golang.org/x/net/html


Source Code
+++++++++++

.. show_github_file:: siongui userpages content/code/go-a-ul-li-to-rst/aulli2rst.go

.. show_github_file:: siongui userpages content/code/go-a-ul-li-to-rst/aulli2rst_test.go

Output of test:

.. code-block:: sh

  === RUN   TestHtmlAUlLiToRst
  a ul li to rst

  - item 1
  - item 2 `link 1 </>`__
  - item 3
      
    - item 3-1
    - item 3-2


  `link 2 </>`__
  --- PASS: TestHtmlAUlLiToRst (0.00s)
  PASS

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


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _HTML unordered (bulleted) list: http://www.w3schools.com/tags/tag_ul.asp
.. _HTML link: http://www.w3schools.com/html/html_links.asp
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _net/html: https://godoc.org/golang.org/x/net/html
