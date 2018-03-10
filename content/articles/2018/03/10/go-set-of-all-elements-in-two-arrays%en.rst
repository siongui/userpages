[Golang] Union of Two Arrays
############################

:date: 2018-03-10T22:47+08:00
:tags: Go, Golang, Algorithm
:category: Go
:summary: Find the set of all elements in two arrays in Go programming language.
:og_image: http://hideoushumpbackfreak.com/image.axd?picture=figure3-1_thumb.png
:adsu: yes


Find the set of all elements in two arrays, i.e., union of two arrays.
This is the continued work of my post yesterday [1]_, which is intersection of
two arrays.

The idea is to convert one array to the data structure of key-value pairs, i.e.,
hash table. The hash table in Go is built-in map_ type. The we check if items of
the other array is in the hash table. If not, append the item to the first
array, and return the first array after finish.

The following is example of above idea.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/aM9FZPXEBkS>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func Union(a, b []int) []int {
  	m := make(map[int]bool)

  	for _, item := range a {
  		m[item] = true
  	}

  	for _, item := range b {
  		if _, ok := m[item]; !ok {
  			a = append(a, item)
  		}
  	}
  	return a
  }

  func main() {
  	var a = []int{1, 2, 3, 4, 5}
  	var b = []int{2, 3, 5, 7, 11}
  	fmt.Println(Union(a, b))
  }

The result will be `[1 2 3 4 5 7 11]`.

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `[Golang] Intersection of Two Arrays <{filename}/articles/2018/03/09/go-match-common-element-in-two-array%en.rst>`_

.. _Go Playground: https://play.golang.org/
.. _map: https://blog.golang.org/go-maps-in-action
