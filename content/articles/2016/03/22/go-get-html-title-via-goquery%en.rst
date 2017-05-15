[Golang] Get HTML Title via goquery
###################################

:date: 2016-03-22T07:25+08:00
:tags: Go, Golang, Commandline, html, goquery, Web Scrape, Go flag Package
:category: Go
:summary: A simple example to read HTML title via goquery_ in Go_.
:adsu: yes


A simple example to read HTML title via goquery_ in Golang_.

Install goquery_:

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery

Source code:

.. show_github_file:: siongui userpages content/code/go-get-html-title/title.go
.. adsu:: 2

Command line usage:

.. code-block:: bash

  $ go run title.go -input=index.html

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] | `golang beautifulsoup <https://www.google.com/search?q=golang+beautifulsoup>`_
       | `What's the go equivalent of Python's Beautiful Soup (an HTML scraping library) ? : golang <https://www.reddit.com/r/golang/comments/3nyumc/whats_the_go_equivalent_of_pythons_beautiful_soup/>`_
       | `soup - Web Scraper in Go, similar to BeautifulSoup : golang <https://www.reddit.com/r/golang/comments/6b3m7n/soup_web_scraper_in_go_similar_to_beautifulsoup/>`_
       | `GitHub - anaskhan96/soup: Web Scraper in Go, similar to BeautifulSoup <https://github.com/anaskhan96/soup>`_
.. [2] | `GitHub - PuerkitoBio/goquery: A little like that j-thing, only in Go. <https://github.com/PuerkitoBio/goquery>`_
       | `Tips and tricks · PuerkitoBio/goquery Wiki · GitHub <https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks>`_
.. adsu:: 3
.. [3] `read html title · twnanda/twnanda@5d81787 · GitHub <https://github.com/twnanda/twnanda/commit/5d81787c957ae9273f78ce5f073dca47849b5ddd>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
