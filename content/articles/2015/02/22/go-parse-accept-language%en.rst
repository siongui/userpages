[Golang] Parse Accept-Language in HTTP Request Header
#####################################################

:date: 2015-02-22 02:49
:tags: Go, Golang, Google App Engine, String Manipulation, HTTP Header, Accept-Language, Locale
:category: Go
:summary: Parse Accept-Language in HTTP Request Header in Go Programming Language.
:adsu: yes


This post is the Go_ version of my previous post
"*[Python] Parse Accept-Language in HTTP Request Header*" [1]_.
This Go_ example parses Accept-Language_ string in HTTP request header and
output (languageTag, quality) pairs:

.. show_github_file:: siongui userpages content/code/go-accept-language/parse.go

Test program for the above code:

.. show_github_file:: siongui userpages content/code/go-accept-language/parse_test.go

Makefile for automating the development:

.. show_github_file:: siongui userpages content/code/go-accept-language/Makefile

The output after run ``make``:

.. code-block:: txt

  === RUN TestParseAcceptLanguage
  --- PASS: TestParseAcceptLanguage (0.00s)
          parse_test.go:6: da, en-gb;q=0.8, en;q=0.7
          parse_test.go:14: [{da 1} {en-gb 0.8} {en 0.7}]
          parse_test.go:16: zh, en-us; q=0.8, en; q=0.6
          parse_test.go:18: [{zh 1} {en-us 0.8} {en 0.6}]
          parse_test.go:20: es-mx, es, en
          parse_test.go:22: [{es-mx 1} {es 1} {en 1}]
          parse_test.go:24: de; q=1.0, en; q=0.5
          parse_test.go:26: [{de 1} {en 0.5}]
  PASS
  ok      _/home/koan/dev/nstechdev/userpages/content/code/go-accept-language     0.003s

Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

References:

.. [1] `[Python] Parse Accept-Language in HTTP Request Header <{filename}../../../2012/10/11/python-parse-accept-language-in-http-request-header%en.rst>`_

.. [2] `Detect User Language (Locale) on Google App Engine Python <{filename}../../../2012/10/12/detect-user-language-locale-gae-python%en.rst>`_

.. [3] `strings - The Go Programming Language <http://golang.org/pkg/strings/>`_

.. [4] `go - How to split string and assign? - Stack Overflow <http://stackoverflow.com/questions/16551354/how-to-split-string-and-assign>`_

.. [5] `go - How to trim leading and trailing white spaces of a string in golang? - Stack Overflow <http://stackoverflow.com/questions/22688010/how-to-trim-leading-and-trailing-white-spaces-of-a-string-in-golang>`_

.. [6] `strconv - The Go Programming Language <http://golang.org/pkg/strconv/>`_

.. [7] `Adding elements to a slice - A Tour of Go <https://tour.golang.org/moretypes/11>`_

.. [8] `testing - The Go Programming Language <http://golang.org/pkg/testing/>`_


.. _Go: https://golang.org/

.. _Accept-Language: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
