[Golang] Check If The Rune is Chinese Character
###############################################

:date: 2018-10-11T04:26+08:00
:tags: Go, Golang, String Manipulation, Locale, i18n, Chinese Pinyin
:category: Go
:summary: Given a rune value, check if the rune is a Chinese character.
:og_image: https://i.pinimg.com/236x/77/ae/80/77ae804cd46b2b43d8971f18ec7015b7.jpg
:adsu: yes


Given a rune_ value, check if the rune is a Chinese character.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/ypZpFhp23nt>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"unicode"
  )

  func main() {
  	str := "hello, 世界"
  	for _, runeValue := range str {
  		if unicode.Is(unicode.Han, runeValue) {
  			fmt.Println(string(runeValue), "is a Chinese character")
  		} else {
  			fmt.Println(string(runeValue), "is not a Chinese character")
  		}
  	}
  }

.. adsu:: 2

**Output**

.. code-block:: txt

  h is not a Chinese character
  e is not a Chinese character
  l is not a Chinese character
  l is not a Chinese character
  o is not a Chinese character
  , is not a Chinese character
    is not a Chinese character
  世 is a Chinese character
  界 is a Chinese character

----

Tested on: `Go Playground`_

----

References:

.. [1] `[Golang] Iterate Over UTF-8 Strings (non-ASCII strings) <{filename}/articles/2016/02/03/go-iterate-over-utf8-non-ascii-string%en.rst>`_
.. [2] `unicode - The Go Programming Language <https://golang.org/pkg/unicode/>`_
.. [3] `[Golang] First Letter of Chinese Character Pinyin <{filename}/articles/2017/05/05/go-chinese-character-pinyin-first-letter%en.rst>`_
.. [4] `[Golang] Calculate String Length <{filename}/articles/2018/01/24/go-calculate-string-length%en.rst>`_
.. [5] `[Golang] Get UTF-8 String Width <{filename}/articles/2016/03/23/go-utf8-string-width%en.rst>`_
.. [6] `[Golang] Sort String by Character <{filename}/articles/2017/05/07/go-sort-string-slice-of-rune%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _for: https://tour.golang.org/flowcontrol/1
.. _range: https://github.com/golang/go/wiki/Range
.. _Go Playground: https://play.golang.org/
.. _rune: https://blog.golang.org/strings
