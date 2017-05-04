[Golang] Iterate Over UTF-8 Strings (non-ASCII strings)
#######################################################

:date: 2016-02-03T19:42+08:00
:tags: Go, Golang, String Manipulation, Locale, i18n
:category: Go
:summary: Iterate over UTF-8 or non-ASCII strings in Go_. Iterations by for_ or
          range_ keywords.
:og_image: https://image.slidesharecdn.com/firstimpressionsofgo-140616044415-phpapp01/95/first-impressions-of-go-11-638.jpg
:adsu: yes


Iterate over UTF-8 or non-ASCII strings in Golang_. Iterations by for_ or range_
keywords.

for_ loop iteration
+++++++++++++++++++

`Run Code on Go Playground <https://play.golang.org/p/_bsLyvyksJ>`__

.. code-block:: go

  package main

  import "unicode/utf8"
  import "fmt"

  func main() {
          str := "I am 字串"

          for i, w := 0, 0; i < len(str); i += w {
                  runeValue, width := utf8.DecodeRuneInString(str[i:])
                  w = width
                  fmt.Println(string(runeValue))
          }
  }

.. adsu:: 2

range_ iteration
++++++++++++++++

`Run Code on Go Playground <https://play.golang.org/p/pDudCVvtuu>`__

.. code-block:: go

  package main

  import "fmt"

  func main() {
          str := "I am 字串"

          for index, runeValue := range str {
                  fmt.Println(index, string(runeValue))
          }
  }

.. adsu:: 3

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

References:

.. [1] `go iterate over string <https://www.google.com/search?q=go+iterate+over+string>`_

.. [2] `Strings, bytes, runes and characters in Go - The Go Blog <https://blog.golang.org/strings>`_

.. [3] `utf8 - The Go Programming Language <https://golang.org/pkg/unicode/utf8/>`_

.. [4] `support non-ascii alphabet (utf8 alphabet) · siongui/go-succinct-data-structure-trie@8ad7e26 · GitHub <https://github.com/siongui/go-succinct-data-structure-trie/commit/8ad7e26db49f8df83980c71737cc9af5972bce81>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _for: https://tour.golang.org/flowcontrol/1
.. _range: https://github.com/golang/go/wiki/Range
