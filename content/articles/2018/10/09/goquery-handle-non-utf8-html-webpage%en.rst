goquery Handle Non-UTF8 HTML Web Page
#####################################

:date: 2018-10-09T21:58+08:00
:tags: Go, Golang, iconv command, goquery, Web Scrape, Go net/http
:category: Go
:summary: Read non-utf8 webpage with goquery if the charset of the page is
          known.
:og_image: https://i.stack.imgur.com/qc9IC.png
:adsu: yes


goquery_ handles only UTF-8_ encoded web pages. The wiki of goquery [1]_
provides a method to handle non-utf8 html pages if the character encoding
(charset) of the pages is known. The trick is to use iconv_ to convert the
encoding to utf8 first. I re-write the code in the wiki and make it more
modular.

Install goquery and `Go iconv binding`_:

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery
  $ go get -u github.com/djimenez/iconv-go

**Source code**

.. code-block:: go

  import (
  	"net/http"

  	"github.com/PuerkitoBio/goquery"
  	iconv "github.com/djimenez/iconv-go"
  )

  func NewDocumentFromNonUtf8Url(url, charset string) (doc *goquery.Document, err error) {
  	resp, err := http.Get(url)
  	if err != nil {
  		return
  	}
  	defer resp.Body.Close()

  	utfBody, err := iconv.NewReader(resp.Body, charset, "utf-8")
  	if err != nil {
  		return
  	}

  	doc, err = goquery.NewDocumentFromReader(utfBody)
  	return
  }

**Example**: Read Big5_ webpage

.. code-block:: go

  func main() {
  	doc, err := NewDocumentFromNonUtf8Url("http://shenfang.com.tw/product.htm", "big5")
  	if err != nil {
  		panic(err)
  	}

  	// do something with the doc
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 18.04``, ``Go 1.10.1``.

----

References:

.. [1] `Tips and tricks · PuerkitoBio/goquery Wiki · GitHub <https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _Big5: https://en.wikipedia.org/wiki/Big5
.. _UTF-8: https://en.wikipedia.org/wiki/UTF-8
.. _iconv: http://linux.die.net/man/1/iconv
.. _Go iconv binding: https://github.com/djimenez/iconv-go
