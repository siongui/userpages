[Golang] Calculate Least Common Multiple (LCM) by GCD
#####################################################

:date: 2017-06-03T03:02+08:00
:tags: Go, Golang, Algorithm, Math, Variadic Function, Recursion
:category: Go
:summary: Find least common multiple (LCM) by greatest common divisor (GCD)
          via Go programming language.
:og_image: http://www.math.com/school/subject1/images/SIU3L3GL.gif
:adsu: yes


Find least common multiple (LCM) by greatest common divisor (GCD) via Go
programming language.

We use the following formula to calculate LCM:

.. container:: align-center

  For :math:`a, b \in \mathbb{N}`, :math:`a * b = LCM(a,b) * GCD(a,b)`

To calculate GCD, see my previous post [1]_.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/SmzvkDjYlb>`__
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

  // find Least Common Multiple (LCM) via GCD
  func LCM(a, b int, integers ...int) int {
  	result := a * b / GCD(a, b)

  	for i := 0; i < len(integers); i++ {
  		result = LCM(result, integers[i])
  	}

  	return result
  }

  func main() {
  	fmt.Println(LCM(10, 15))
  	fmt.Println(LCM(10, 15, 20))
  	fmt.Println(LCM(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
  }

If you are not familiar with variadic function of Go, see my previous post [2]_.

.. adsu:: 2

Tested on `Go Playground`_.

----

References:

.. [1] `[Golang] Greatest Common Divisor via Euclidean Algorithm <{filename}../../05/14/go-gcd-via-euclidean-algorithm%en.rst>`_
.. [2] `[Golang] Variadic Function Example - addEventListener <{filename}../../03/12/go-variadic-function-example-addEventListener%en.rst>`_
.. [3] `[Golang] Smallest Multiple - Problem 5 - Project Euler <{filename}../02/go-smallest-multiple-problem-5-project-euler%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
