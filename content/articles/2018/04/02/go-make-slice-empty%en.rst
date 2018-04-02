[Golang] Make Slice Empty
#########################

:date: 2018-04-02T23:36+08:00
:tags: Go, Golang, Data Structure
:category: Go
:summary: Make slice empty in Go programming language.
:og_image: https://research.swtch.com/godata3.png
:adsu: yes


After a slices contains some items and we want to empty the slice, how to write
the syntax?

Use *int* array as example, the syntax is:

.. code-block:: go

  myArray = []int{}

See the following full example on Go Playground.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/JnMftkbyVU7>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func printNumbers(numbers []int) {
  	for _, number := range numbers {
  		fmt.Println(number)
  	}
  }

  func main() {
  	// initialize int slice
  	var numbers = []int{1, 2, 4, 5}
  	printNumbers(numbers)

  	// make slice empty
  	numbers = []int{}
  	printNumbers(numbers)
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] | `golang empty slice - Google search <https://www.google.com/search?q=golang+empty+slice>`_
       | `golang empty slice - DuckDuckGo search <https://duckduckgo.com/?q=golang+empty+slice>`_
       | `golang empty slice - Ecosia search <https://www.ecosia.org/search?q=golang+empty+slice>`_
       | `golang empty slice - Qwant search <https://www.qwant.com/?q=golang+empty+slice>`_
       | `golang empty slice - Bing search <https://www.bing.com/search?q=golang+empty+slice>`_
       | `golang empty slice - Yahoo search <https://search.yahoo.com/search?p=golang+empty+slice>`_
       | `golang empty slice - Baidu search <https://www.baidu.com/s?wd=golang+empty+slice>`_
       | `golang empty slice - Yandex search <https://www.yandex.com/search/?text=golang+empty+slice>`_

.. _Go Playground: https://play.golang.org/
