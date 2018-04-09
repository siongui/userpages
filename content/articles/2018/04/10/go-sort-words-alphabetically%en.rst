[Golang] Sort Words Alphabetically
##################################

:date: 2018-04-10T03:29+08:00
:tags: Go, Golang, Go sort
:category: Go
:summary: Sort a list of words alphabetically in Go.
:og_image: https://image.slidesharecdn.com/ftdgolang-140524152550-phpapp01/95/golang-42-638.jpg
:adsu: yes


Go standard library provides sort.Strings_ to help us sort words alphabetically.
The following example comes from Go official site.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/NkFCfMKOLY_3>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"sort"
  )

  func main() {
  	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
  	sort.Strings(s)
  	fmt.Println(s)
  }

.. adsu:: 2

----

References:

.. [1] | `sort words alphabetically - Google search <https://www.google.com/search?q=sort+words+alphabetically>`_
       | `sort words alphabetically - DuckDuckGo search <https://duckduckgo.com/?q=sort+words+alphabetically>`_
       | `sort words alphabetically - Ecosia search <https://www.ecosia.org/search?q=sort+words+alphabetically>`_
       | `sort words alphabetically - Qwant search <https://www.qwant.com/?q=sort+words+alphabetically>`_
       | `sort words alphabetically - Bing search <https://www.bing.com/search?q=sort+words+alphabetically>`_
       | `sort words alphabetically - Yahoo search <https://search.yahoo.com/search?p=sort+words+alphabetically>`_
       | `sort words alphabetically - Baidu search <https://www.baidu.com/s?wd=sort+words+alphabetically>`_
       | `sort words alphabetically - Yandex search <https://www.yandex.com/search/?text=sort+words+alphabetically>`_
.. [2] `[Golang] Sort String by Character <{filename}/articles/2017/05/07/go-sort-string-slice-of-rune%en.rst>`_

.. _sort.Strings: https://golang.org/pkg/sort/#Strings
