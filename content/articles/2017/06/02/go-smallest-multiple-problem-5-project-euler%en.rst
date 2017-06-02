[Golang] Smallest Multiple - Problem 5 - Project Euler
################################################################

:date: 2017-06-02T20:12+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler
:category: Go
:summary: Go solution to smallest multiple
          - Problem 9 - Project Euler.
:og_image: http://www.math.com/school/subject1/images/SIU3L3GL.gif
:adsu: yes

Go solution to smallest multiple - Problem 5 - Project Euler. [1]_


**Problem**:

  2520 is the smallest number that can be divided by each of the numbers from 1
  to 10 without any remainder.

  What is the smallest positive number that is evenly divisible by all of the
  numbers from 1 to 20?


**Solution**:

  LCM(1, 2, 3, ..., 20) = 18044195

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/NoSoID5vTO>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  // greatest common divisor (GCD) via Euclidean algorithm
  func GCD(a, b int) int {
  	for b != 0 {
  		t := b
  		b = a % b
  		a = t
  	}
  	return a
  }

  // find Least Common Multiple (LCD) via GCD
  func LCM(a, b int) int {
  	return a * b / GCD(a, b)
  }

  func main() {
  	result := 1
  	for j := 2; j <= 20; j++ {
  		result = LCM(result, j)
  	}
  	fmt.Println(result)
  }

.. adsu:: 2

Tested on `Go Playground`_.

----

References:

.. [1] `Smallest multiple - Problem 5 - Project Euler <https://projecteuler.net/problem=5>`_
.. [2] `[Golang] Greatest Common Divisor via Euclidean Algorithm <{filename}../../05/14/go-gcd-via-euclidean-algorithm%en.rst>`_
.. [3] `LCM Calculator - Least Common Multiple <http://www.calculatorsoup.com/calculators/math/lcm.php>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
