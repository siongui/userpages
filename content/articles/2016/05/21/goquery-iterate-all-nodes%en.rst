goquery - Iterate over All DOM Elements in HTML
###############################################

:date: 2016-05-21T21:12+08:00
:tags: Go, Golang, DOM, goquery, Web Scrape, html
:category: Go
:summary: Iterate over all DOM nodes/elements via goquery_ in Go_ programming
          language.
:adsu: yes


Introduction
++++++++++++

Iterate over all DOM nodes/elements via goquery_ in Golang_
(Go_ programming language).

The trick is to use ``Find("*")`` to access all nodes in DOM tree.


Install goquery_ Package
++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery


Source Code
+++++++++++

.. show_github_file:: siongui userpages content/code/goquery-iterate-all-element/node.go

----

Tested on: ``Ubuntu Linux 16.04``, ``Go 1.6.2``.

----

References:

.. [1] `github.com/PuerkitoBio/goquery - GoDoc <https://godoc.org/github.com/PuerkitoBio/goquery>`_

.. [2] `[Golang] Iterate over All DOM Elements in HTML <{filename}../../04/10/go-iterate-over-all-dom-elements-in-html%en.rst>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
