[Golang] Anonymous Function in Defer Statement
##############################################

:date: 2018-05-17T23:39+08:00
:tags: Go, Golang, Go defer Statement
:category: Go
:summary: Use anonymous function in Go *defer* statement.
:og_image: https://image.slidesharecdn.com/golang-150827144155-lva1-app6891/95/go-lang-25-638.jpg
:adsu: yes


**TL;DR**: Syntax for using anonymous function in Go `defer statement`_

.. code-block:: go

  defer func() {
  	// do something here after surrounding function returns
  }()


**Explanation**

I want to execute ``c <- 1`` right after surrounding function returns [1]_ [2]_.
In the beginning I thought it was easy to do this with *defer*, so I write the
following code:

.. code-block:: go

  defer c <- 1

But the code cannot be compiled successfully. The compiler complained that
*expression in defer must be function call*. So I searched and found this on
*A Tour of Go*:

  A defer statement defers the execution of a function until the surrounding
  function returns.

It looks like *defer* must be followed by a function, so I thought I can use
anonymous function in this case. Then I tried again and wrote the following
code:

.. code-block:: go

  defer func() { c <- 1 }

I still got the same complaint from compiler. After more searches, I found that
the correct way to do this is:

.. code-block:: go

  defer func() { c <- 1 }()

The *()* means call and execute the anonymous function. If you write only
*func() { c <- 1 }* without *()*, you only declare the anonymous function
without calling it. That's why I got the same error while compiling.

The following are complete example for using anonymous function in defer
statement.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/dHYMGZNbnj5>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func routine(site string, c chan int) {
  	defer func() { c <- 1 }()

  	// do something
  	fmt.Println(site)
  }

  func main() {
  	c := make(chan int)

  	sites := []string{
  		"https://www.google.com/",
  		"https://duckduckgo.com/",
  		"https://www.bing.com/"}

  	for _, site := range sites {
  		go routine(site, c)
  	}

  	// wait all goroutines to finish
  	for i := 0; i < len(sites); i++ {
  		<-c
  	}
  }

.. adsu:: 2

----

Tested on: `The Go Playground`_

References:

.. [1] `[Golang] Use Defer to Wait For Goroutines to Finish <{filename}/articles/2018/05/16/go-use-defer-to-wait-for-goroutine-to-finish%en.rst>`_
.. [2] `[Golang] Wait For Goroutine to Finish <{filename}/articles/2015/03/23/go-wait-for-goroutine-to-finish%en.rst>`_
.. [3] `Behavior of defer function in named return function : golang <https://old.reddit.com/r/golang/comments/9ysxts/behavior_of_defer_function_in_named_return/>`_
.. [4] `Higher order functions. Why not? : golang <https://old.reddit.com/r/golang/comments/a1o1d7/higher_order_functions_why_not/>`_

.. _defer statement: https://tour.golang.org/flowcontrol/12
.. _The Go Playground: https://play.golang.org/
