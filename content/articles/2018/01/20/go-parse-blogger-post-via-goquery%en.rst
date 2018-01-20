[Golang] Web Scrape Blogger Post via goquery
############################################

:date: 2018-01-20T07:13+08:00
:tags: Go, Golang, goquery, Web Scrape, html, CSS selectors
:category: Go
:summary:  Fetch a public post on *Blogger* and extract data via *goquery*.
:og_image: https://jonathanmh.com/wp-content/uploads/2016/10/web-scraping-web-crawling-golang-go-goquery-825x371.png
:adsu: yes

Fetch a public post on Blogger_ and extract data via goquery_.

We will extract the following data from HTML:

- PostUrl
- Title
- TimeStamp
- Author
- Summary
- Content

The following is complete source code:

.. show_github_file:: siongui userpages content/code/go/scrape-blogger-post/parse.go
.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.9.2``.

----

References:

.. [1] `[Golang] Web Scrape Facebook Post via goquery <{filename}../../../2017/02/17/go-parse-facebook-post-via-goquery%en.rst>`_
.. [2] `GitHub - PuerkitoBio/goquery: A little like that j-thing, only in Go. <https://github.com/PuerkitoBio/goquery>`_ |godoc|
.. [3] `Tips and tricks · PuerkitoBio/goquery Wiki · GitHub <https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks>`_

.. _goquery: https://github.com/PuerkitoBio/goquery
.. _Blogger: https://www.blogger.com/

.. |godoc| image:: https://godoc.org/github.com/PuerkitoBio/goquery?status.png
   :target: https://godoc.org/github.com/PuerkitoBio/goquery
