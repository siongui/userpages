[Golang] Remove Duplicates From Slice or Array
##############################################

:date: 2018-04-14T23:38+08:00
:tags: Go, Golang, Data Structure
:category: Go
:summary: Remove duplicates from slice or array in Go.
:og_image: https://research.swtch.com/godata3.png
:adsu: yes


**Question**

  How to remove duplicates from slice or array in Go?

**Solution**

  There are many methods to do this [1]_. My approach is to create a *map* [2]_
  type and for each item in the slice/array, check if the item is in the *map*.
  If the item is in the *map*, the it is duplicate. If not in the *map*, save it
  in the *map*. After finished, the *map* contains no duplicates in the original
  slice/array.

  The following is an example of above approach.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/dNHRvqnaW2H>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func RemoveDuplicatesFromSlice(s []string) []string {
  	m := make(map[string]bool)
  	for _, item := range s {
  		if _, ok := m[item]; ok {
  			// duplicate item
  			fmt.Println(item, "is a duplicate")
  		} else {
  			m[item] = true
  		}
  	}

  	var result []string
  	for item, _ := range m {
  		result = append(result, item)
  	}
  	return result
  }

  func main() {
  	var testSlice = []string{"Mike", "Matt", "Nancy", "Adam", "Jenny", "Nancy", "Carl"}
  	fmt.Println(testSlice)
  	resultSlice := RemoveDuplicatesFromSlice(testSlice)
  	fmt.Println(resultSlice)
  }

Output:

.. code-block:: txt

  [Mike Matt Nancy Adam Jenny Nancy Carl]
  Nancy is a duplicate
  [Jenny Carl Mike Matt Nancy Adam]

.. adsu:: 2

----

Tested on: `Go Playground`_

References:

.. [1] | `remove duplicates from array - Google search <https://www.google.com/search?q=remove+duplicates+from+array>`_
       | `remove duplicates from array - DuckDuckGo search <https://duckduckgo.com/?q=remove+duplicates+from+array>`_
       | `remove duplicates from array - Ecosia search <https://www.ecosia.org/search?q=remove+duplicates+from+array>`_
       | `remove duplicates from array - Qwant search <https://www.qwant.com/?q=remove+duplicates+from+array>`_
       | `remove duplicates from array - Bing search <https://www.bing.com/search?q=remove+duplicates+from+array>`_
       | `remove duplicates from array - Yahoo search <https://search.yahoo.com/search?p=remove+duplicates+from+array>`_
       | `remove duplicates from array - Baidu search <https://www.baidu.com/s?wd=remove+duplicates+from+array>`_
       | `remove duplicates from array - Yandex search <https://www.yandex.com/search/?text=remove+duplicates+from+array>`_
.. [2] `Go maps in action - The Go Blog <https://blog.golang.org/go-maps-in-action>`_

.. _Go Playground: https://play.golang.org/
