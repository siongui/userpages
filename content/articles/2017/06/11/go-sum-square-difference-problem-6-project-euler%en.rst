[Golang] Sum Square Difference - Problem 6 - Project Euler
##########################################################

:date: 2017-06-11T23:18+08:00
:tags: Go, Golang, Math, Project Euler
:category: Go
:summary: Go solution to sum square difference
          - Problem 6 - Project Euler.
:og_image: http://www.careerdeal.net/wp-content/uploads/2015/09/Project-Euler-problem-6-Careerdeal.png
:adsu: yes

Go solution to sum square difference - Problem 6 - Project Euler. [1]_

**Problem**:

  The sum of the squares of the first ten natural numbers is,

  :math:`1^2 + 2^2 + ... + 10^2 = 385`

  The square of the sum of the first ten natural numbers is,

  :math:`(1 + 2 + ... + 10)^2 = 55^2 = 3025`

  Hence the difference between the sum of the squares of the first ten natural
  numbers and the square of the sum is :math:`3025 − 385 = 2640`.

  Find the difference between the sum of the squares of the first one hundred
  natural numbers and the square of the sum.


**Solution**:

  25502500 - 338350 = 25164150


.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/ocruPiTbvz>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func main() {
  	var sumsquare = uint64(1)
  	var squaresum = uint64(1)
  	for i := uint64(2); i <= 100; i++ {
  		squaresum += i
  		sumsquare += i * i
  	}
  	squaresum = squaresum * squaresum
  	fmt.Print(squaresum)
  	fmt.Print(" - ")
  	fmt.Print(sumsquare)
  	fmt.Print(" = ")
  	fmt.Print(squaresum - sumsquare)
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `Sum square difference - Problem 6 - Project Euler <https://projecteuler.net/problem=6>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
