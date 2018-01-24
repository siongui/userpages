[Golang] Calculate String Length
################################

:date: 2018-01-24T08:49+08:00
:tags: Go, Golang, String Manipulation, Locale, i18n
:category: Go
:summary: Calculate the length of string (UTF-8) in *Go*.
          Iterate over the UTF-8 string by *for* or *range* keyword.
:og_image: https://i.ytimg.com/vi/CRGLPCH18-E/maxresdefault.jpg
:adsu: yes

Calculate the length of string (UTF-8) in *Go* programming language.
Iterate over the UTF-8 string by for_ or range_ keyword
to calculate string length.


range_ iteration
++++++++++++++++

`Run Code on Go Playground <https://play.golang.org/p/j1zLFIODXzu>`__

.. code-block:: go

  func StringLength(s string) int {
  	l := 0
  	for _, runeValue := range s {
  		runeValue = runeValue
  		l++
  	}
  	return l
  }

.. adsu:: 2


for_ loop iteration
+++++++++++++++++++

`Run Code on Go Playground <https://play.golang.org/p/lydOnJ-izrR>`__

.. code-block:: go

  import "unicode/utf8"

  func StringLength(s string) int {
  	l := 0
  	for i, w := 0, 0; i < len(s); i += w {
  		_, width := utf8.DecodeRuneInString(s[i:])
  		w = width
  		l++
  	}
  	return l
  }

----

Tested on: `The Go Playground`_

----

References:

.. [1] | `calculate string length - Google search <https://www.google.com/search?q=calculate+string+length>`_
       | `calculate string length - DuckDuckGo search <https://duckduckgo.com/?q=calculate+string+length>`_
       | `calculate string length - Ecosia search <https://www.ecosia.org/search?q=calculate+string+length>`_
       | `calculate string length - Qwant search <https://www.qwant.com/?q=calculate+string+length>`_
       | `calculate string length - Bing search <https://www.bing.com/search?q=calculate+string+length>`_
       | `calculate string length - Yahoo search <https://search.yahoo.com/search?p=calculate+string+length>`_
       | `calculate string length - Baidu search <https://www.baidu.com/s?wd=calculate+string+length>`_
       | `calculate string length - Yandex search <https://www.yandex.com/search/?text=calculate+string+length>`_
.. [2] `Calculate String Length Online <http://string-functions.com/length.aspx>`_
.. [3] | `golang detect thai script - Google search <https://www.google.com/search?q=golang+detect+thai+script>`_
       | `golang detect thai script - DuckDuckGo search <https://duckduckgo.com/?q=golang+detect+thai+script>`_
       | `golang detect thai script - Ecosia search <https://www.ecosia.org/search?q=golang+detect+thai+script>`_
       | `golang detect thai script - Qwant search <https://www.qwant.com/?q=golang+detect+thai+script>`_
       | `golang detect thai script - Bing search <https://www.bing.com/search?q=golang+detect+thai+script>`_
       | `golang detect thai script - Yahoo search <https://search.yahoo.com/search?p=golang+detect+thai+script>`_
       | `golang detect thai script - Baidu search <https://www.baidu.com/s?wd=golang+detect+thai+script>`_
       | `golang detect thai script - Yandex search <https://www.yandex.com/search/?text=golang+detect+thai+script>`_
.. adsu:: 3
.. [4] `unicode - The Go Programming Language <https://golang.org/pkg/unicode/>`_
.. [5] `unicode - How do I count Japanese words in Go-lang - Stack Overflow <https://stackoverflow.com/questions/24576659/how-do-i-count-japanese-words-in-go-lang>`_
.. [6] `GitHub - veer66/mapkha: Thai word segmentation program in Go <https://github.com/veer66/mapkha>`_
.. [7] `GitHub - blevesearch/segment: A Go library for performing Unicode Text Segmentation as described in Unicode Standard Annex #29 <https://github.com/blevesearch/segment>`_
.. [8] `[Golang] Iterate Over UTF-8 Strings (non-ASCII strings) <{filename}/articles/2016/02/03/go-iterate-over-utf8-non-ascii-string%en.rst>`_

.. _for: https://tour.golang.org/flowcontrol/1
.. _range: https://github.com/golang/go/wiki/Range
.. _The Go Playground: https://play.golang.org/
