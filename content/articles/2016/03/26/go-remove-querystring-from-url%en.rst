[Golang] Remove Query String From URL
#####################################

:date: 2016-03-26T02:54+08:00
:tags: Go, Golang, String Manipulation, Go net/url
:category: Go
:summary: Remove `query string`_ from URL_ in Go_ programming language.
:adsu: yes


Remove `query string`_ from URL_ in Golang_.

.. code-block:: go

  import "net/url"

  func StripQueryString(inputUrl string) string {
  	u, err := url.Parse(inputUrl)
  	if err != nil {
  		panic(err)
  	}
  	u.RawQuery = ""
  	return u.String()
  }

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

.. adsu:: 2

References:

.. [1] `go url remove query <https://www.google.com/search?q=go+url+remove+query>`_

.. [2] `url - The Go Programming Language <https://golang.org/pkg/net/url/>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _query string: https://en.wikipedia.org/wiki/Query_string
.. _URL: https://en.wikipedia.org/wiki/Uniform_Resource_Locator
