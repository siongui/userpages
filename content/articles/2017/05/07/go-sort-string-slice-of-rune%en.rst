[Golang] Sort String by Character
#################################

:date: 2017-05-07T03:05+08:00
:tags: Go, Golang, Go sort
:category: Go
:summary: Sort a string by character, i.e., convert the string to []rune and
          then sort the slice of runes in Go programming language.
:og_image: https://www.calhoun.io/content/images/2016/10/Let-s-Learn-Algorithms---Bubble-Sort-impl.png
:adsu: yes


Sort a string by character, i.e., convert the ``string`` to ``[]rune`` and then
sort the slice of runes in Go programming language. This problem is actually the
same as sorting a slice of runes in nature.

We use sort_ package in Go standard library to help us sort ``[]rune``,
converted from the string.

Go prior to 1.8
+++++++++++++++

The steps:

1. Convert ``string`` to ``[]rune``
2. Define a new type ``ByRune``, which is actually ``[]rune``. Implement
   ``Len``, ``Swap``, and ``Less`` methods of type ``ByRune``.
3. Call sort.Sort_ to sort the slice of runes.
4. Convert ``[]rune`` back to ``string`` and return the string.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/F3zl_-y9A8>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"sort"
  )

  type ByRune []rune

  func (r ByRune) Len() int           { return len(r) }
  func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
  func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

  func StringToRuneSlice(s string) []rune {
  	var r []rune
  	for _, runeValue := range s {
  		r = append(r, runeValue)
  	}
  	return r
  }

  func SortStringByCharacter(s string) string {
  	var r ByRune = StringToRuneSlice(s)
  	sort.Sort(r)
  	return string(r)
  }

  func main() {
  	fmt.Println(SortStringByCharacter("listen"))
  	fmt.Println(SortStringByCharacter("silent"))
  }


.. adsu:: 2


Go 1.8 or later
+++++++++++++++

The sort_ package in Go 1.8 introduces new methods for sorting slices [6]_.
We can use sort.Slice_ method directly without defining a new type.

The steps:

1. Convert ``string`` to ``[]rune``
2. Define a *less* method and call sort.Slice_ with the slice of runes and the
   *less* method as parameters.
3. Convert ``[]rune`` back to ``string`` and return the string.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/-8qDGxII5n>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"sort"
  )

  func StringToRuneSlice(s string) []rune {
  	var r []rune
  	for _, runeValue := range s {
  		r = append(r, runeValue)
  	}
  	return r
  }

  func SortStringByCharacter(s string) string {
  	r := StringToRuneSlice(s)
  	sort.Slice(r, func(i, j int) bool {
  		return r[i] < r[j]
  	})
  	return string(r)
  }

  func main() {
  	fmt.Println(SortStringByCharacter("listen"))
  	fmt.Println(SortStringByCharacter("silent"))
  }

.. adsu:: 3

Tested on:

- ``Ubuntu Linux 17.04``, ``Go 1.8.1``
- `Go Playground`_

----

References:

.. [1] | `golang sort string - Google search <https://www.google.com/search?q=golang+sort+string>`_
       | `golang sort string - DuckDuckGo search <https://duckduckgo.com/?q=golang+sort+string>`_
       | `golang sort string - Ecosia search <https://www.ecosia.org/search?q=golang+sort+string>`_
       | `golang sort string - Qwant search <https://www.qwant.com/?q=golang+sort+string>`_
       | `golang sort string - Bing search <https://www.bing.com/search?q=golang+sort+string>`_
       | `golang sort string - Yahoo search <https://search.yahoo.com/search?p=golang+sort+string>`_
       | `golang sort string - Baidu search <https://www.baidu.com/s?wd=golang+sort+string>`_
       | `golang sort string - Yandex search <https://www.yandex.com/search/?text=golang+sort+string>`_
.. [2] `string - Go sort a slice of runes? - Stack Overflow <http://stackoverflow.com/questions/18171136/go-sort-a-slice-of-runes>`_
.. [3] `Package sort - The Go Programming Language <https://golang.org/pkg/sort/>`_
.. [4] `[Golang] Check Whether Two Strings Are Anagram of Each Other <{filename}../06/go-check-if-two-string-are-anagram%en.rst>`_
.. [5] `Strings, bytes, runes and characters in Go - The Go Blog <https://blog.golang.org/strings>`_
.. [6] `sort: make sorting easier, add Slice, SliceStable, SliceIsSorted, reflect.Swapper · Issue #16721 · golang/go · GitHub <https://github.com/golang/go/issues/16721>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _sort: https://golang.org/pkg/sort/
.. _sort.Slice: https://golang.org/pkg/sort/#Slice
.. _sort.Sort: https://golang.org/pkg/sort/#Sort
