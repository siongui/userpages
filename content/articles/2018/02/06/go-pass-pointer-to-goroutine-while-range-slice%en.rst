[Golang] Pass Reference to Goroutine While range Slice
######################################################

:date: 2018-02-06T09:34+08:00
:tags: Go, Golang, Goroutine, Go Channels
:category: Go
:summary: Caveat: Passing reference/pointer to goroutine function
          while range slice in Go.
:og_image: https://cdn-images-1.medium.com/max/1200/1*GWYUFH14uOVLNHY-L1tv2w.jpeg
:adsu: yes


Something strange happened when I tried to pass references to goroutines_ while
*range* slice. Consider the following code:

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func printNumber(number *int, c chan int) {
  	fmt.Println(*number)
  	c <- 1
  }

  func main() {
  	numbers := []int{1, 2, 3}

  	c := make(chan int)
  	for _, number := range numbers {
  		go printNumber(&number, c)
  	}

  	// wait all goroutines to finish
  	for i := 0; i < len(numbers); i++ {
  		<-c
  	}
  }

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/2sc5ELbH-vX>`__
   :class: align-center

The output in the console is:

.. code-block:: txt

  3
  3
  3

I cannot figure out what happened. It supposed to be something like:

.. code-block:: txt

  1
  3
  2

But I got all 3! This is not what I want. I cannot figured out what's wrong, so
I tried some googling [2]_, and the answer in [3]_ gived me solution.

*Solution*: **use the array index instead of the value**

Change the code of *range* to the following:

.. code-block:: go

  	for index, _ := range numbers {
  		go printNumber(&numbers[index], c)
  	}

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/58cR61kRzrL>`__
   :class: align-center

And now everything works as expected! (Although I still do not know why!)

----

Tested on:

- ``Ubuntu Linux 17.10``, ``Go 1.9.3``.
- `Go Playground`_

----

References:

.. [1] `[Golang] Wait For Goroutine to Finish <{filename}/articles/2015/03/23/go-wait-for-goroutine-to-finish%en.rst>`_
.. [2] | `golang range pass same pointer - Google search <https://www.google.com/search?q=golang+range+pass+same+pointer>`_
       | `golang range pass same pointer - DuckDuckGo search <https://duckduckgo.com/?q=golang+range+pass+same+pointer>`_
       | `golang range pass same pointer - Ecosia search <https://www.ecosia.org/search?q=golang+range+pass+same+pointer>`_
       | `golang range pass same pointer - Qwant search <https://www.qwant.com/?q=golang+range+pass+same+pointer>`_
       | `golang range pass same pointer - Bing search <https://www.bing.com/search?q=golang+range+pass+same+pointer>`_
       | `golang range pass same pointer - Yahoo search <https://search.yahoo.com/search?p=golang+range+pass+same+pointer>`_
       | `golang range pass same pointer - Baidu search <https://www.baidu.com/s?wd=golang+range+pass+same+pointer>`_
       | `golang range pass same pointer - Yandex search <https://www.yandex.com/search/?text=golang+range+pass+same+pointer>`_
.. adsu:: 2
.. [3] `pointers - golang range references instead values - Stack Overflow <https://stackoverflow.com/a/29498133>`_

.. _channels: https://tour.golang.org/concurrency/2
.. _goroutines: https://tour.golang.org/concurrency/1
.. _Go Playground: https://play.golang.org/
