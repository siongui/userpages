goquery - Convert HTML Unordered List to reStructuredText
#########################################################

:date: 2016-05-05T23:29+08:00
:tags: Go, Golang, DOM, goquery, Web Scrape, reStructuredText, html
:category: Go
:summary: Convert `HTML unordered (bulleted) list`_ to restructuredtext_ format
          via goquery_ in Go_ programming language.
:adsu: yes


Introduction
++++++++++++

Convert `HTML unordered (bulleted) list`_ to restructuredtext_ format via
goquery_ in Golang_ (Go_ programming language).


Install goquery_ Package
++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery


Source Code
+++++++++++

.. show_github_file:: siongui userpages content/code/goquery-ul-li-to-rst/ulli2rst.go

.. show_github_file:: siongui userpages content/code/goquery-ul-li-to-rst/ulli2rst_test.go


----

Tested on: ``Ubuntu Linux 16.04``, ``Go 1.6.2``.

----

References:

.. [1] `github.com/PuerkitoBio/goquery - GoDoc <https://godoc.org/github.com/PuerkitoBio/goquery>`_

.. [2] `goquery - Replace HTML Link Node with reStructuredText Text Node <{filename}../04/goquery-replace-html-link-node-with-rst-text-node%en.rst>`_

.. [3] `[Golang] Read Lines From File or String <{filename}../../04/06/go-readlines-from-file-or-string%en.rst>`_

.. [4] `go string concat - Google search <https://www.google.com/search?q=go+string+concat>`_

       `How to efficiently concatenate strings in Go? - Stack Overflow <http://stackoverflow.com/a/1763606>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _HTML unordered (bulleted) list: http://www.w3schools.com/tags/tag_ul.asp
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
