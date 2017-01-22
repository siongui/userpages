[Golang] Extract Title, Image, and URL via goquery
##################################################

:date: 2016-03-31T22:51+08:00
:tags: Go, Golang, Golang template, JSON, goquery, Web Scrape, reStructuredText
:category: Go
:summary: Extract title, image, and URL in `buy123 product webpage`_ via
          goquery_, and then output the info to `reStructuredText image`_.
:adsu: yes


Extract title, image, and URL in `buy123 product webpage`_ via goquery_, and
then output the info to `reStructuredText image`_.

First we examine the source code of `buy123 product webpage`_, we found that the
product info is embedded into json string in `script tag`_, type of which is
``application/ld+json``.

So we extract the json string via goquery Find_ method. Then convert it to a
struct type. Finally we use go `text/template`_ package to output the info in
`reStructuredText image`_ format.

.. code-block:: go

  package main

  import (
  	"bytes"
  	"encoding/json"
  	"github.com/PuerkitoBio/goquery"
  	"text/template"
  )

  const rstTmpl = `.. image:: {{.Image}}
    :alt: {{.Name}}
    :target: {{.Url}}
    :align: center`

  type buy123ProductInfo struct {
  	Name        string
  	Description string
  	Image       string
  	Url         string
  }

  func parseBuy123(url string) string {
  	doc, err := goquery.NewDocument(url)
  	if err != nil {
  		panic(err)
  	}

  	jsonBlob := doc.Find("script[type=\"application/ld+json\"]").Text()

  	i := buy123ProductInfo{}
  	err = json.Unmarshal([]byte(jsonBlob), &i)
  	if err != nil {
  		panic(err)
  	}

  	tmpl, err := template.New("buy123").Parse(rstTmpl)
  	if err != nil {
  		panic(err)
  	}
  	var rst bytes.Buffer
  	err = tmpl.Execute(&rst, i)
  	if err != nil {
  		panic(err)
  	}

  	return rst.String()
  }

Output:

.. code-block:: txt

  .. image:: //s3-buy123.cdn.hinet.net/images/item/GLFA9T7.png
    :alt: 6LED多功能太陽能露營燈
    :target: https://direct.buy123.com.tw/site/item/64493/6LED%E5%A4%9A%E5%8A%9F%E8%83%BD%E5%A4%AA%E9%99%BD%E8%83%BD%E9%9C%B2%E7%87%9F%E7%87%88
    :align: center

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `go template output to string <https://www.google.com/search?q=go+template+output+to+string>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://www.google.com/search?q=goquery
.. _buy123 product webpage: https://www.buy123.com.tw/site/item/64493/6LED%E5%A4%9A%E5%8A%9F%E8%83%BD%E5%A4%AA%E9%99%BD%E8%83%BD%E9%9C%B2%E7%87%9F%E7%87%88
.. _reStructuredText image: http://docutils.sourceforge.net/docs/ref/rst/directives.html#images
.. _script tag: http://www.w3schools.com/tags/tag_script.asp
.. _Find: https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Find
.. _text/template: https://golang.org/pkg/text/template/
