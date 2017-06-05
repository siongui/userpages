[Golang] Amicable Numbers - Problem 21 - Project Euler
######################################################

:date: 2017-05-25T00:52+08:00
:tags: Go, Golang, Algorithm, Project Euler, Prime Number, Math
:category: Go
:summary: Go solution to amicable numbers
          - Problem 21 - Project Euler.
:og_image: http://4.bp.blogspot.com/-ziI--Wyy0yw/ULzKZdjZK5I/AAAAAAAADJ8/GBsRWfKlz38/s1600/AmicableNumbers.jpg
:adsu: yes

Go solution to amicable numbers - Problem 21 - Project Euler. [1]_

**Problem**:

  Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
  If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

  For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

  Evaluate the sum of all the amicable numbers under 10000.

**Solution**:

  [220 284 1184 1210 2620 2924 5020 5564 6232 6368]

  31626

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/bTL9gNtU9N>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  // Get all prime factors of a given number n
  func PrimeFactors(n int) (pfs []int) {
  	// Get the number of 2s that divide n
  	for n%2 == 0 {
  		pfs = append(pfs, 2)
  		n = n / 2
  	}

  	// n must be odd at this point. so we can skip one element
  	// (note i = i + 2)
  	for i := 3; i*i <= n; i = i + 2 {
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

  // return p^i
  func Power(p, i int) int {
  	result := 1
  	for j := 0; j < i; j++ {
  		result *= p
  	}
  	return result
  }

  // formula comes from https://math.stackexchange.com/a/22723
  func SumOfProperDivisors(n int) int {
  	pfs := PrimeFactors(n)

  	// key: prime
  	// value: prime exponents
  	m := make(map[int]int)
  	for _, prime := range pfs {
  		_, ok := m[prime]
  		if ok {
  			m[prime] += 1
  		} else {
  			m[prime] = 1
  		}
  	}

  	sumOfAllFactors := 1
  	for prime, exponents := range m {
  		sumOfAllFactors *= (Power(prime, exponents+1) - 1) / (prime - 1)
  	}
  	return sumOfAllFactors - n
  }

  func AmicableNumbersUnder10000() (amicables []int) {
  	for i := 3; i < 10000; i++ {
  		s := SumOfProperDivisors(i)
  		if s == i {
  			continue
  		}
  		if SumOfProperDivisors(s) == i {
  			amicables = append(amicables, i)
  		}
  	}
  	return
  }

  func main() {
  	amicables := AmicableNumbersUnder10000()
  	fmt.Println(amicables)

  	sum := 0
  	for i := 0; i < len(amicables); i++ {
  		sum += amicables[i]
  	}
  	fmt.Println(sum)
  }

The method for sum of all proper divisors (factors) comes from my another post
[2]_.

.. adsu:: 2

----

References:

.. [1] `Amicable numbers - Problem 21 - Project Euler <https://projecteuler.net/problem=21>`_
.. [2] `[Golang] Sum of the Proper Divisors (Factors) <{filename}../19/go-sum-of-proper-factors%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
