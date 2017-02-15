goquery querySelector
#####################

:date: 2017-02-15T23:08+08:00
:tags: Go, Golang, DOM, goquery, Web Scrape, html, Go interface
:category: Go
:summary: Implement querySelector_ in goquery_
:og_image: https://jonathanmh.com/wp-content/uploads/2016/10/web-scraping-web-crawling-golang-go-goquery-825x371.png
:adsu: yes


Introduction
++++++++++++

The Find_ function in goquery_ returns all matched elements in `*Selection`_,
but sometimes we may need only first matched element in `*Selection`_, just like
querySelector_ does, so I write a *QuerySelector* function, which accepts CSS
selector as arguemnt and outputs the Selection object which contains only first
matched element.


Install goquery_ Package
++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery


Source Code
+++++++++++

Implementation
==============

.. code-block:: go

  import (
  	"github.com/PuerkitoBio/goquery"
  )

  type object interface {
  	Find(string) *goquery.Selection
  }

  func QuerySelector(s object, selector string) *goquery.Selection {
  	return s.Find(selector).First()
  }

.. adsu:: 2

Usage
=====

.. code-block:: go

  import (
  	"github.com/PuerkitoBio/goquery"
  )

  func main() {
  	doc, err := goquery.NewDocument("https://siongui.github.io/")
  	if err != nil {
  		panic(err)
  	}

  	s := QuerySelector(doc, "meta")
  	// do something with s, which contains only one matched element
  }

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.7.5``.

----

.. adsu:: 3

References:

.. [1] `GitHub - PuerkitoBio/goquery: A little like that j-thing, only in Go. <https://github.com/PuerkitoBio/goquery>`_ |godoc|
.. [2] `Tips and tricks · PuerkitoBio/goquery Wiki · GitHub <https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _querySelector: https://www.google.com/search?q=querySelector
.. _Find: https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Find
.. _*Selection: https://godoc.org/github.com/PuerkitoBio/goquery#Selection

.. |godoc| image:: https://godoc.org/github.com/PuerkitoBio/goquery?status.png
   :target: https://godoc.org/github.com/PuerkitoBio/goquery
