[Golang] Use Defer to Wait For Goroutines to Finish
###################################################

:date: 2018-05-16T23:35+08:00
:tags: Go, Golang, Goroutine, Go Channels, Go defer Statement
:category: Go
:summary: Use Go *defer* statement to elegantly wait for all goroutines to
          finish.
:og_image: https://image.slidesharecdn.com/golang-150827144155-lva1-app6891/95/go-lang-25-638.jpg
:adsu: yes


This post shows how to use `defer statement`_ to elegantly wait for all
goroutines to finish.

Let's start with bad practice. Look at the following code example, which waits
for all goroutines to finish:

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/0lDZVtmpL1e>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func routine(site string, c chan int) {
  	// do something
  	fmt.Println(site)

  	c <- 1
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

In the end of the goroutine_, The line ``c <- 1`` sends integer 1 to channel_ to
notify that the goroutine is finished. This way looks good, if your goroutine
return only at the end of the func. More often than not, the goroutine will
return at multiple places inside the func. For example, consider the following
goroutine:

.. code-block:: go

  func routine(c chan int) {
  	b, err := ioutil.ReadFile("myfile.txt")
  	if err != nil {
  		fmt.Println(err)
  		c <- 1
  		return
  	}

  	mystruct := MyStruct{}
  	err = json.Unmarshal(b, &mystrcut)
  	if err != nil {
  		fmt.Println(err)
  		c <- 1
  		return
  	}

  	fmt.Println(mystruct)
  	c <- 1
  	return
  }

As you can see from above example, there are three *return* statement in the
goroutine, and you have to write ``c <- 1`` three times right before the
*return*. This is not a good practice because it is easy to forget to write
``c <- 1`` if you add more *return* for handling more possible errors in the
code.

A good practice and more elegant way to do this is to use `defer statement`_.
According to the description in *A Tour of Go*:

  A defer statement defers the execution of a function until the surrounding
  function returns.

So we can use *defer* to write ``c <- 1`` only once as follows:

.. code-block:: go

  func routine(c chan int) {
  	defer func() { c <- 1 }()

  	b, err := ioutil.ReadFile("myfile.txt")
  	if err != nil {
  		fmt.Println(err)
  		return
  	}

  	mystruct := MyStruct{}
  	err = json.Unmarshal(b, &mystrcut)
  	if err != nil {
  		fmt.Println(err)
  		return
  	}

  	fmt.Println(mystruct)
  	return
  }

No matter how many *return* in your goroutine, the ``c <- 1`` is guaranteed to
be executed right after the function returns. To use *defer* is a better
practice because it makes your code more readable and you will not forget to add
``c <- 1`` if you add more *return* in the function.

The following is complete code example of good practice:

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

In the example of this post, we use channel_ to wait for all goroutines to
finish, you can also use sync.WaitGroup_ to do this. See [3]_ for more
information.

.. adsu:: 2

----

Tested on: `The Go Playground`_

----

References:

.. [1] `Defer - A Tour of Go <https://tour.golang.org/flowcontrol/12>`_
.. [2] `Defer, Panic, and Recover - The Go Blog <https://blog.golang.org/defer-panic-and-recover>`_
.. [3] `[Golang] Wait For Goroutine to Finish <{filename}/articles/2015/03/23/go-wait-for-goroutine-to-finish%en.rst>`_

.. _defer statement: https://tour.golang.org/flowcontrol/12
.. _channel: https://tour.golang.org/concurrency/2
.. _goroutine: https://tour.golang.org/concurrency/1
.. _sync.WaitGroup: https://golang.org/pkg/sync/#WaitGroup
.. _The Go Playground: https://play.golang.org/
