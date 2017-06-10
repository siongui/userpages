[Golang] Longest Collatz Sequence - Problem 14 - Project Euler
##############################################################

:date: 2017-06-10T21:31+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler
:category: Go
:summary: Go solution to longest Collatz sequence
          - Problem 14 - Project Euler.
:og_image: https://martin-thoma.com/images/2013/05/collatz-graph.png
:adsu: yes

Go solution to longest Collatz sequence - Problem 14 - Project Euler. [1]_

**Problem**:

  The following iterative sequence is defined for the set of positive integers:

    | n → n/2 (n is even)
    | n → 3n + 1 (n is odd)

  Using the rule above and starting with 13, we generate the following sequence:

    13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

  It can be seen that this sequence (starting at 13 and finishing at 1) contains
  10 terms. Although it has not been proved yet (Collatz Problem), it is thought
  that all starting numbers finish at 1.

  Which starting number, under one million, produces the longest chain?

  NOTE: Once the chain starts the terms are allowed to go above one million.


**Solution**:

  starting number: 837799, chain length: 525

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/Ha3kbkW2eo>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func CollatzSequence(n uint) []uint {
  	var result []uint
  	result = append(result, n)
  	for n != 1 {
  		if n%2 == 0 {
  			n /= 2
  		} else {
  			n = 3*n + 1
  		}
  		result = append(result, n)
  	}
  	return result
  }

  func main() {
  	number := uint(1)
  	max := 1
  	for n := uint(1); n < 1000000; n++ {
  		length := len(CollatzSequence(n))
  		if length > max {
  			number = n
  			max = length
  		}
  	}
  	fmt.Println(number)
  	fmt.Println(max)
  }


.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `Longest Collatz sequence - Problem 14 - Project Euler <https://projecteuler.net/problem=14>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
