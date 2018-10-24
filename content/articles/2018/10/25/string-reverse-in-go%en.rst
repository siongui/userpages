[Golang] Reverse a String
#########################

:date: 2018-10-25T00:54+08:00
:tags: Go, Golang, String Manipulation, Locale, i18n
:category: Go
:summary: Reverse a string using *for* or *range* keyword in *Go*.
:og_image: https://www.geeksforgeeks.org/wp-content/uploads/string-reverse.jpg
:adsu: yes

Reverse a string using for_ or range_ keyword in Go_ programming language.
For example, *stressed* will become *desserts* after reversed.

range_ iteration
++++++++++++++++

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/dNN48jBPNcd>`__
   :class: align-center

.. code-block:: go

  func Reverse(s string) (rs string) {
  	for _, runeValue := range s {
  		rs = string(runeValue) + rs
  	}
  	return
  }

.. adsu:: 2


for_ loop iteration
+++++++++++++++++++

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/iEKHjXrbYWl>`__
   :class: align-center

.. code-block:: go

  import "unicode/utf8"

  func Reverse(s string) (rs string) {
  	for i, w := 0, 0; i < len(s); i += w {
  		r, width := utf8.DecodeRuneInString(s[i:])
  		w = width
  		rs = string(r) + rs
  	}
  	return
  }

----

Tested on: `The Go Playground`_

----

References:

.. adsu:: 3
.. [1] `[Golang] Iterate Over UTF-8 Strings (non-ASCII strings) <{filename}/articles/2016/02/03/go-iterate-over-utf8-non-ascii-string%en.rst>`_
.. [2] `[Golang] Calculate String Length <{filename}/articles/2018/01/24/go-calculate-string-length%en.rst>`_
.. [3] `[Golang] Get UTF-8 String Width <{filename}/articles/2016/03/23/go-utf8-string-width%en.rst>`_
.. [4] | `string reverse - Google search <https://www.google.com/search?q=string+reverse>`_
       | `string reverse - DuckDuckGo search <https://duckduckgo.com/?q=string+reverse>`_
       | `string reverse - Ecosia search <https://www.ecosia.org/search?q=string+reverse>`_
       | `string reverse - Qwant search <https://www.qwant.com/?q=string+reverse>`_
       | `string reverse - Bing search <https://www.bing.com/search?q=string+reverse>`_
       | `string reverse - Yahoo search <https://search.yahoo.com/search?p=string+reverse>`_
       | `string reverse - Baidu search <https://www.baidu.com/s?wd=string+reverse>`_
       | `string reverse - Yandex search <https://www.yandex.com/search/?text=string+reverse>`_

.. _for: https://tour.golang.org/flowcontrol/1
.. _range: https://github.com/golang/go/wiki/Range
.. _Go: https://golang.org/
.. _The Go Playground: https://play.golang.org/
