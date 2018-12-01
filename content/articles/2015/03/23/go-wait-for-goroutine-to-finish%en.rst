[Golang] Wait For Goroutines to Finish
######################################

:date: 2015-03-23 22:16
:modified: 2018-05-16T22:56+08:00
:tags: Go, Golang, Goroutine, Go Channels, Go sync Package, Go defer Statement
:category: Go
:summary: Use channels to wait for all goroutines to finish
:og_image: https://cdn-images-1.medium.com/max/1200/1*GWYUFH14uOVLNHY-L1tv2w.jpeg
:adsu: yes


.. contents:: This post shows two approaches to wait for all goroutines_ to
              finish.

Channel [1]_
++++++++++++

Wait for all goroutines to finish via channels_:

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


sync.WaitGroup [4]_
+++++++++++++++++++

Wait for all goroutines to finish via sync.WaitGroup_:

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/7jSEcLb5SyD>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"sync"
  )

  var wg sync.WaitGroup

  func routine(site string) {
  	// Decrement the counter when the goroutine completes.
  	defer wg.Done()

  	// do something
  	fmt.Println(site)
  }

  func main() {
  	sites := []string{
  		"https://www.google.com/",
  		"https://duckduckgo.com/",
  		"https://www.bing.com/"}

  	for _, site := range sites {
  		// Increment the WaitGroup counter.
  		wg.Add(1)

  		go routine(site)
  	}

  	// wait all goroutines to finish
  	wg.Wait()
  }

.. adsu:: 3

----

Tested on: `The Go Playground`_

----

References:

.. [1] `Waiting for all goroutines to finish before ending main - Google Groups <https://groups.google.com/d/topic/golang-nuts/mNhXnWLFOo4>`_
.. [2] `DuckDuckGo <https://duckduckgo.com/>`_ Search: `golang wait for goroutine to finish <https://duckduckgo.com/?q=golang+wait+for+goroutine+to+finish>`__
.. [3] `Google <https://www.google.com/>`_ Search: `golang wait for goroutine to finish <https://www.google.com/search?q=golang+wait+for+goroutine+to+finish>`__
.. [4] `优雅の使用sync.WaitGroup - hawkingrei |  Blog <https://www.hawkingrei.com/blog/2017/08/26/gracefully-use-waitgroup/>`_
.. [5] `goroutine anti-pattern? : golang <https://old.reddit.com/r/golang/comments/9o4vh5/goroutine_antipattern/>`_
.. [6] `Non blocking way of notifying a goroutine, by many at once : golang <https://old.reddit.com/r/golang/comments/9zzeai/non_blocking_way_of_notifying_a_goroutine_by_many/>`_
.. [7] `Can someone explain why this program isn't faster with goroutines? : golang <https://old.reddit.com/r/golang/comments/a1wznn/can_someone_explain_why_this_program_isnt_faster/>`_

.. _channels: https://tour.golang.org/concurrency/2
.. _goroutines: https://tour.golang.org/concurrency/1
.. _sync.WaitGroup: https://golang.org/pkg/sync/#WaitGroup
.. _The Go Playground: https://play.golang.org/
