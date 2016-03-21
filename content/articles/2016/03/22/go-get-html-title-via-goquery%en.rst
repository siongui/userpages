[Golang] Get HTML Title via goquery
###################################

:date: 2016-03-22T07:25+08:00
:tags: Go, Golang, Commandline, html
:category: Go
:summary: A simple example to read HTML title via goquery_ in Go_.


A simple example to read HTML title via goquery_ in Golang_.

Install goquery_:

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery

Source code:

.. show_github_file:: siongui userpages content/code/go-get-html-title/title.go

Command line usage:

.. code-block:: bash

  $ go run title.go -input=index.html

----

References:

.. [1] `golang beautifulsoup <https://www.google.com/search?q=golang+beautifulsoup>`_

       `What's the go equivalent of Python's Beautiful Soup (an HTML scraping library) ? : golang <https://www.reddit.com/r/golang/comments/3nyumc/whats_the_go_equivalent_of_pythons_beautiful_soup/>`_

.. [2] `GitHub - PuerkitoBio/goquery: A little like that j-thing, only in Go. <https://github.com/PuerkitoBio/goquery>`_

       `Tips and tricks · PuerkitoBio/goquery Wiki · GitHub <https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
