[Golang] Set Difference of Two Arrays
#####################################

:date: 2018-03-14T21:13+08:00
:tags: Go, Golang, Algorithm, Set Operation
:category: Go
:summary: Find set differecne of two arrays, i.e., the elements in one array but
          not in the other, in Go.
:og_image: https://cdn.programiz.com/sites/tutorial2program/files/set-difference.jpg
:adsu: yes


Find the elements in one array but not in the other, i.e., set difference of two
arrays. In mathematical term:

  :math:`A-B=\{x\in A \ and \ x \notin B\}`

The idea is to convert the array B to the data structure of key-value pairs,
i.e., hash table. The hash table in Go is built-in map_ type. Then we check if
items in array A is in the hash table. If not, append the item to the difference
array, and return the difference array after finish.

The following is the implementation of above idea.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/ZnXDliHPB1m>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  // Set Difference: A - B
  func Difference(a, b []int) (diff []int) {
  	m := make(map[int]bool)

  	for _, item := range b {
  		m[item] = true
  	}

  	for _, item := range a {
  		if _, ok := m[item]; !ok || len(m) == 0 {
  			diff = append(diff, item)
  		}
  	}
  	return
  }

  func main() {
  	var a = []int{1, 2, 3, 4, 5}
  	var b = []int{2, 3, 5, 7, 11}
  	fmt.Println(Difference(a, b))
  }

The result will be `[1 4]`.

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `[Golang] Intersection of Two Arrays <{filename}/articles/2018/03/09/go-match-common-element-in-two-array%en.rst>`_
.. [2] `[Golang] Union of Two Arrays <{filename}/articles/2018/03/10/go-set-of-all-elements-in-two-arrays%en.rst>`_

.. _Go Playground: https://play.golang.org/
.. _map: https://blog.golang.org/go-maps-in-action
