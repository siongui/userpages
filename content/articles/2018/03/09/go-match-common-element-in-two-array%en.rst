[Golang] Intersection of Two Arrays
###################################

:date: 2018-03-09T22:29+08:00
:tags: Go, Golang, Algorithm
:category: Go
:summary: Find common elements (matches, intersection) of two arrays
          in Go programming language.
:og_image: http://4.bp.blogspot.com/-034xqHwdXwU/UqWAJEoDWKI/AAAAAAAAAHk/HyCOJbka4CU/s400/Intersection+of+two+arrays+java+coding+.jpg
:adsu: yes


Find common elements in two arrays, i.e., intersection or matches in two arrays.
There are many ways to find the intersection [1]_ [3]_. Here we will implement
the method mentioned in [2]_.

The idea is to convert one array to the data structure of key-value pairs, i.e.,
hash table. The hash table in Go is built-in map_ type. The we check if items of
the other array is in the hash table. If yes, the item is in the intersection of
the two arrays.

The following is example of above idea.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/c7wcBTNy4ht>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func Intersection(a, b []int) (c []int) {
  	m := make(map[int]bool)

  	for _, item := range a {
  		m[item] = true
  	}

  	for _, item := range b {
  		if _, ok := m[item]; ok {
  			c = append(c, item)
  		}
  	}
  	return
  }

  func main() {
  	var a = []int{1, 2, 3, 4, 5}
  	var b = []int{2, 3, 5, 7, 11}
  	fmt.Println(Intersection(a, b))
  }

The result will be `[2 3 5]`.

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] | `match common element in two array - Google search <https://www.google.com/search?q=match+common+element+in+two+array>`_
       | `match common element in two array - DuckDuckGo search <https://duckduckgo.com/?q=match+common+element+in+two+array>`_
       | `match common element in two array - Ecosia search <https://www.ecosia.org/search?q=match+common+element+in+two+array>`_
       | `match common element in two array - Qwant search <https://www.qwant.com/?q=match+common+element+in+two+array>`_
       | `match common element in two array - Bing search <https://www.bing.com/search?q=match+common+element+in+two+array>`_
       | `match common element in two array - Yahoo search <https://search.yahoo.com/search?p=match+common+element+in+two+array>`_
       | `match common element in two array - Baidu search <https://www.baidu.com/s?wd=match+common+element+in+two+array>`_
       | `match common element in two array - Yandex search <https://www.yandex.com/search/?text=match+common+element+in+two+array>`_

.. [2] `efficiency - Quick algorithm to find matches between two arrays - Software Engineering Stack Exchange <https://softwareengineering.stackexchange.com/a/223477>`_

.. [3] | `intersection of two arrays - Google search <https://www.google.com/search?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - DuckDuckGo search <https://duckduckgo.com/?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - Ecosia search <https://www.ecosia.org/search?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - Qwant search <https://www.qwant.com/?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - Bing search <https://www.bing.com/search?q=intersection+of+two+arrays>`_
       | `intersection of two arrays - Yahoo search <https://search.yahoo.com/search?p=intersection+of+two+arrays>`_
       | `intersection of two arrays - Baidu search <https://www.baidu.com/s?wd=intersection+of+two+arrays>`_
       | `intersection of two arrays - Yandex search <https://www.yandex.com/search/?text=intersection+of+two+arrays>`_

.. [4] | `Find Union and Intersection of two unsorted arrays - GeeksforGeeks <https://www.geeksforgeeks.org/find-union-and-intersection-of-two-unsorted-arrays/>`_
       | `Union and Intersection of two sorted arrays - GeeksforGeeks <https://www.geeksforgeeks.org/union-and-intersection-of-two-sorted-arrays-2/>`_

.. _Go Playground: https://play.golang.org/
.. _map: https://blog.golang.org/go-maps-in-action
