[Golang] Non-abundant sums - Problem 23 - Project Euler
#######################################################

:date: 2018-01-22T22:19+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler, Prime Number
:category: Go
:summary: Go solution to Non-abundant sums
          - Problem 23 - Project Euler.
:og_image: https://i.stack.imgur.com/geIb2.png
:adsu: yes

**Problem**: [1]_

  A perfect number is a number for which the sum of its proper divisors is
  exactly equal to the number. For example, the sum of the proper divisors of 28
  would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

  A number n is called deficient if the sum of its proper divisors is less than
  n and it is called abundant if this sum exceeds n.

  As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
  number that can be written as the sum of two abundant numbers is 24. By
  mathematical analysis, it can be shown that all integers greater than 28123
  can be written as the sum of two abundant numbers. However, this upper limit
  cannot be reduced any further by analysis even though it is known that the
  greatest number that cannot be expressed as the sum of two abundant numbers is
  less than this limit.

  Find the sum of all the positive integers which cannot be written as the sum
  of two abundant numbers.


**Solution**:

  4179871

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/TmItMd8upzS>`__
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
  
  func main() {
  	// Get all abundant numbers under 28123
  	var abundantNumbers []int
  	for i := 1; i <= 28123; i++ {
  		if SumOfProperDivisors(i) > i {
  			abundantNumbers = append(abundantNumbers, i)
  		}
  	}
  	//fmt.Println(abundantNumbers)
  
  	sumOfTwoAbundantNumbers := make(map[int]bool)
  	for i := 0; i < len(abundantNumbers); i++ {
  		for j := i; j < len(abundantNumbers); j++ {
  			a := abundantNumbers[i]
  			b := abundantNumbers[j]
  			if a+b <= 28123 {
  				//fmt.Println(a+b, " = ", a, " + ", b)
  				sumOfTwoAbundantNumbers[a+b] = true
  			} else {
  				break
  			}
  		}
  	}
  
  	sum := 0
  	for i := 1; i <= 28123; i++ {
  		_, ok := sumOfTwoAbundantNumbers[i]
  		if !ok {
  			//fmt.Println(i)
  			sum += i
  		}
  	}
  	fmt.Println(sum)
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `Non-abundant sums - Problem 23 - Project Euler <https://projecteuler.net/problem=23>`_
.. [2] `[Golang] Sum of the Proper Divisors (Factors) <{filename}/articles/2017/05/19/go-sum-of-proper-factors%en.rst>`_
.. [3] `Go maps in action - The Go Blog <https://blog.golang.org/go-maps-in-action>`_

.. _Go Playground: https://play.golang.org/
