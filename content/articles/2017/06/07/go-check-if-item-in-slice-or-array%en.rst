[Golang] Check If Item in Slice or Array
########################################

:date: 2017-06-07T00:27+08:00
:tags: Go, Golang, Data Structure
:category: Go
:summary: Check if an item is in slice or array in Go programming language.
:og_image: https://research.swtch.com/godata3.png
:adsu: yes


**Question**

  How to check if an item is in a slice or array in Go?

**Solution**

  There are many methods to do this [1]_. My approach is to create a *map* [2]_
  type and save the items of slice/array in the map, and then check if the item
  is in the *map* or not.

  The following is an example of above approach.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/gk-otCgvwC>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func main() {
  	// initialize a slice of int
  	slice := []int{2, 3, 5, 7, 11, 13}

  	// save the items of slice in map
  	m := make(map[int]bool)
  	for i := 0; i < len(slice); i++ {
  		m[slice[i]] = true
  	}

  	// check if 5 is in slice by checking map
  	if _, ok := m[5]; ok {
  		fmt.Println("5 is in slice")
  	} else {
  		fmt.Println("5 is not in slice")
  	}

  	// check if 6 is in slice by checking map
  	if _, ok := m[6]; ok {
  		fmt.Println("6 is in slice")
  	} else {
  		fmt.Println("6 is not in slice")
  	}
  }

To see real-world application, see reference [3]_. The prime numbers are
returned in a slice. Then the prime numbers are saved in the map, and we can
check if a number is prime or not by checking if the number is in the map.

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] | `golang check if item in slice - Google search <https://www.google.com/search?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - DuckDuckGo search <https://duckduckgo.com/?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Ecosia search <https://www.ecosia.org/search?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Qwant search <https://www.qwant.com/?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Bing search <https://www.bing.com/search?q=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Yahoo search <https://search.yahoo.com/search?p=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Baidu search <https://www.baidu.com/s?wd=golang+check+if+item+in+slice>`_
       | `golang check if item in slice - Yandex search <https://www.yandex.com/search/?text=golang+check+if+item+in+slice>`_
.. [2] `Go maps in action - The Go Blog <https://blog.golang.org/go-maps-in-action>`_
.. [3] `[Golang] Goldbach's conjecture <{filename}../06/go-goldbach-conjecture%en.rst>`_
.. [4] `📣 [Data Structures with Go]: Data Structures and Algorithms Example: Array, Binary Tree, Linked List, Sort (Quick, Merge), Stack, Queue (my Midterm & Final study notes for Data Structure course) (Feel Free to Contribute) : golang <https://old.reddit.com/r/golang/comments/9y4xu1/data_structures_with_go_data_structures_and/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _Goldbach's conjecture: https://www.google.com/search?q=Goldbach's+conjecture
