[Golang] Remove Last Character of UTF-8 String
##############################################

:date: 2020-07-14T05:54+08:00
:tags: Go, Golang, String Manipulation, Locale, i18n
:category: Go
:summary: Remove last character of UTF-8 string via utf8.DecodeLastRuneInString_
          in Go.
:og_image: https://miro.medium.com/max/3188/1*xHqIcw_7V9tu39pYtJdQyA.png
:adsu: yes


Remove last character of UTF-8 string via utf8.DecodeLastRuneInString_ in Go_.
The code is modified from official example in `unicode/utf8` package [2]_.

`Run Code on Go Playground <https://play.golang.org/p/Y1yCZ4wU05t>`__

.. code-block:: go

  import "unicode/utf8"

  func RemoveLastChar(str string) string {
  	for len(str) > 0 {
  		_, size := utf8.DecodeLastRuneInString(str)
  		return str[:len(str)-size]
  	}
  	return str
  }

If you do not want to import *unicode/utf8* package, you can use *for* keyword
to range the string to remove last character. On how to range strings, see [3]_
for more reference.

----

Tested on: ``Ubuntu Linux 20.04``, ``Go 1.12.17``.

----

References:

.. [1] | `golang string get last character - Google search <https://www.google.com/search?q=golang+string+get+last+character>`_
       | `golang string get last character - DuckDuckGo search <https://duckduckgo.com/?q=golang+string+get+last+character>`_
       | `golang string get last character - Ecosia search <https://www.ecosia.org/search?q=golang+string+get+last+character>`_
       | `golang string get last character - Qwant search <https://www.qwant.com/?q=golang+string+get+last+character>`_
       | `golang string get last character - Bing search <https://www.bing.com/search?q=golang+string+get+last+character>`_
       | `golang string get last character - Yahoo search <https://search.yahoo.com/search?p=golang+string+get+last+character>`_
       | `golang string get last character - Baidu search <https://www.baidu.com/s?wd=golang+string+get+last+character>`_
       | `golang string get last character - Yandex search <https://www.yandex.com/search/?text=golang+string+get+last+character>`_

.. [2] | `go - How to get the last X Characters of a Golang String? - Stack Overflow <https://stackoverflow.com/questions/26166641/how-to-get-the-last-x-characters-of-a-golang-string>`_
       | `DecodeLastRuneInString - utf8 - The Go Programming Language <https://golang.org/pkg/unicode/utf8/#DecodeLastRuneInString>`_

.. [3] `[Golang] Iterate Over UTF-8 Strings (non-ASCII strings) <{filename}/articles/2016/02/03/go-iterate-over-utf8-non-ascii-string%en.rst>`_

.. [4] `implement Backspace and Enter keys in keypad · siongui/paligo@5423f84 · GitHub <https://github.com/siongui/paligo/commit/5423f843386263374c63604e576e0e1561364620>`_

.. _Go: https://golang.org/
.. _utf8.DecodeLastRuneInString: https://golang.org/pkg/unicode/utf8/#DecodeLastRuneInString
