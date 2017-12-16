[Golang] Multiples of 3 and 5 - Problem 1 - Project Euler
#########################################################

:date: 2017-12-16T22:36+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler
:category: Go
:summary: Go solution to Multiples of 3 and 5
          - Problem 1 - Project Euler.
:og_image: https://i.pinimg.com/originals/3c/fc/93/3cfc93a120afb2ea744d35fe033c17bc.jpg
:adsu: yes

**Problem**: [1]_

  If we list all the natural numbers below 10 that are multiples of 3 or 5, we
  get 3, 5, 6 and 9. The sum of these multiples is 23.

  Find the sum of all the multiples of 3 or 5 below 1000.

**Solution**:

  Brute-force method. For every natural number below 1000, check if it's a
  multiple of 3 or 5. If it is, add it to the sum.

  The sum of all the multiples of 3 or 5 below 1000 is 234168


.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/_-3G9gnMLd>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func main() {
  	sum := 0

  	for i := 1; i <= 1000; i++ {
  		if i%3 == 0 {
  			//fmt.Println(i)
  			sum += i
  			continue
  		}
  		if i%5 == 0 {
  			//fmt.Println(i)
  			sum += i
  			continue
  		}
  	}

  	fmt.Println("The sum of all the multiples of 3 or 5 below 1000 is", sum)
  }


Tested on: `Go Playground`_

----

References:

.. [1] `Multiples of 3 and 5 - Problem 1 - Project Euler <https://projecteuler.net/problem=1>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
