[Golang] Fetch DOM Style Object Properties via goquery
######################################################

:date: 2017-02-28T21:22+08:00
:tags: Go, Golang, goquery, Web Scrape, html, DOM, CSS selectors
:category: Go
:summary:  Web scraping - fetch `HTML DOM Style Object`_ properties via goquery_
           in Go_ programming language.
:og_image: https://jonathanmh.com/wp-content/uploads/2016/10/web-scraping-web-crawling-golang-go-goquery-825x371.png
:adsu: yes


Fetch properties of `HTML DOM Style Object`_ via goquery_ - example of Golang_
`web scraping`_.

The idea of `web scraping`_ is not difficult if the webpage is not rendered by
JavaScript_. View source of the webpage, find out where the data is, and apply
correct `CSS selector`_ to extract the HTML element which contains the data we
need. The following Go_ code fetches the properties of `HTML DOM Style Object`_
in W3Schools_.

.. code-block:: go

  package codegen

  import (
  	"github.com/PuerkitoBio/goquery"
  )

  func GetAllStyleProperty() (cssprops []string, err error) {
  	url := "https://www.w3schools.com/jsref/dom_obj_style.asp"

  	doc, err := goquery.NewDocument(url)
  	if err != nil {
  		return
  	}

  	csspropslinks := doc.Find(".w3-table-all").First().Find("a")
  	csspropslinks.Each(func(i int, s *goquery.Selection) {
  		cssprops = append(cssprops, s.Text())
  	})

  	return
  }

.. adsu:: 2

Print the fetched properties on screen.

.. code-block:: go

  package codegen

  import (
  	"testing"
  )

  func TestGetAllStyleProperty(t *testing.T) {
  	cssprops, err := GetAllStyleProperty()
  	if err != nil {
  		t.Error(err)
  	}

  	for _, prop := range cssprops {
  		println(prop)
  	}
  }

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.8``.

----

.. adsu:: 3

References:

.. [1] `GitHub - PuerkitoBio/goquery: A little like that j-thing, only in Go. <https://github.com/PuerkitoBio/goquery>`_ |godoc|
.. [2] `Tips and tricks · PuerkitoBio/goquery Wiki · GitHub <https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks>`_
.. [3] `HTML DOM Style object - W3Schools <https://www.w3schools.com/jsref/dom_obj_style.asp>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _web scraping: https://www.google.com/search?q=web+scraping
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _HTML DOM Style Object: https://www.w3schools.com/jsref/dom_obj_style.asp
.. _CSS selector: https://www.google.com/search?q=CSS+selector
.. _W3Schools: https://www.w3schools.com/

.. |godoc| image:: https://godoc.org/github.com/PuerkitoBio/goquery?status.png
   :target: https://godoc.org/github.com/PuerkitoBio/goquery
