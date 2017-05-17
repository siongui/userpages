[Golang] Largest Prime Factor - Problem 3 - Project Euler
#########################################################

:date: 2017-05-17T21:26+08:00
:tags: Go, Golang, Algorithm, Project Euler, Prime Number
:category: Go
:summary: Go solution to largest prime factor of 600851475143
          - Problem 3 - Project Euler.
:og_image: http://www.mathsisfun.com/numbers/images/factor-tree-48.gif
:adsu: yes

Go solution to largest prime factor - Problem 3 - Project Euler. [1]_

**Problem**:

  The prime factors of 13195 are 5, 7, 13 and 29.

  What is the largest prime factor of the number 600851475143 ?

**Solution**:

  6857 (600851475143 = 71 x 839 x 1471 x 6857)

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/Qedc4S_INB>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  // Get all prime factors of a given number n
  func PrimeFactors(n int64) (pfs []int64) {
  	// Get the number of 2s that divide n
  	for n%2 == 0 {
  		pfs = append(pfs, 2)
  		n = n / 2
  	}

  	// n must be odd at this point. so we can skip one element
  	// (note i = i + 2)
  	var i int64
  	for i = 3; i*i <= n; i = i + 2 {
  		// while i divides n, append i and divide n
  		for n%i == 0 {
  			pfs = append(pfs, i)
  			n = n / i
  		}
  	}

  	// This condition is to handle the case when n is a prime number
  	// greater than 2
  	if n > 2 {
  		pfs = append(pfs, n)
  	}

  	return
  }

  func main() {
  	fmt.Println(PrimeFactors(600851475143))
  }

The method for factoring a number comes from my another post [2]_.

.. adsu:: 2

----

References:

.. [1] `Largest prime factor - Problem 3 - Project Euler <https://projecteuler.net/problem=3>`_
.. [2] `[Golang] Get All Prime Factors of Integer Number <{filename}../09/go-find-all-prime-factors-of-integer-number%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
