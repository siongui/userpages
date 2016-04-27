goquery Get Number of Children Nodes
####################################

:date: 2016-04-27T23:01+08:00
:tags: Go, Golang, html, goquery, Web Scrape
:category: Go
:summary: Get the number of children nodes via goquery_ in Go_.


Get the number of children nodes via goquery_ in Golang_ (Go_ programming
language).

- Get the number of all children nodes - Use `Contents()`_ method
- Get the number of element children nodes - Use `Children()`_ method

Install goquery_:

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery

Source code:

.. show_github_file:: siongui userpages content/code/goquery-number-of-children/node.go

Output:

.. code-block:: txt

  5
  1

----

Tested on: ``Ubuntu Linux 16.04``, ``Go 1.6.2``.

----

References:

.. [1] `GitHub - PuerkitoBio/goquery: A little like that j-thing, only in Go. <https://github.com/PuerkitoBio/goquery>`_

       `Tips and tricks · PuerkitoBio/goquery Wiki · GitHub <https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks>`_

.. [2] `github.com/PuerkitoBio/goquery - GoDoc <https://godoc.org/github.com/PuerkitoBio/goquery>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _Contents(): https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Contents
.. _Children(): https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Children
