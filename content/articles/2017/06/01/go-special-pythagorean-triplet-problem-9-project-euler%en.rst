[Golang] Special Pythagorean Triplet - Problem 9 - Project Euler
################################################################

:date: 2017-06-01T08:08+08:00
:tags: Go, Golang, Algorithm, Project Euler
:category: Go
:summary: Go solution to special Pythagorean triplet
          - Problem 9 - Project Euler.
:og_image: https://www.mathsisfun.com/geometry/images/triangle-3-4-5-pyth.gif
:adsu: yes

Go solution to special Pythagorean triplet - Problem 9 - Project Euler. [1]_


**Problem**:

  A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

  :math:`a^2` + :math:`b^2` = :math:`c^2`

  For example, :math:`3^2` + :math:`4^2` = 9 + 16 = 25 = :math:`5^2`.

  There exists exactly one Pythagorean triplet for which a + b + c = 1000.

  Find the product abc.


**Solution**:

  200 * 375 * 425 = 31875000

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/kPRzHdzFJg>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func findTriplet() (a, b, c int) {
  	for a = 1; a < 500; a++ {
  		for b = a + 1; b < 1000; b++ {
  			c = 1000 - a - b
  			if b > c {
  				break
  			}
  			if a*a+b*b == c*c {
  				return
  			}
  		}
  	}

  	// cannot find. return all 0
  	a = 0
  	b = 0
  	c = 0
  	return
  }

  func main() {
  	a, b, c := findTriplet()
  	fmt.Println(a, b, c)
  	fmt.Println(a * b * c)
  }

.. adsu:: 2

Tested on `Go Playground`_.

----

References:

.. [1] `Special Pythagorean triplet - Problem 9 - Project Euler <https://projecteuler.net/problem=9>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
