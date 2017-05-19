[Golang] Sum of the Proper Divisors (Factors)
#############################################

:date: 2017-05-19T01:41+08:00
:tags: Go, Golang, Algorithm, Prime Number
:category: Go
:summary: Calculate the sum of all proper divisors (factors)
          of an integer number
          in Go programming language.
:og_image: https://thinklikecomputerscientist.files.wordpress.com/2014/04/12.jpg
:adsu: yes

Calculate the sum of all proper divisors (factors) of an integer number
in Go programming language.

The steps:

1. Given a natural number n, perform prime factorization of n. [3]_
2. Use the formula from [2]_ to calculate sum of all divisors (factors) of n.
3. Get sum of proper divisors by subtracting n from the sum of step 2.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/D76D1piTkQ>`__
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
  	fmt.Println(SumOfProperDivisors(220))
  	fmt.Println(SumOfProperDivisors(284))
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] | `sum of the proper divisors - Google search <https://www.google.com/search?q=sum+of+the+proper+divisors>`_
       | `sum of the proper divisors - DuckDuckGo search <https://duckduckgo.com/?q=sum+of+the+proper+divisors>`_
       | `sum of the proper divisors - Ecosia search <https://www.ecosia.org/search?q=sum+of+the+proper+divisors>`_
       | `sum of the proper divisors - Qwant search <https://www.qwant.com/?q=sum+of+the+proper+divisors>`_
       | `sum of the proper divisors - Bing search <https://www.bing.com/search?q=sum+of+the+proper+divisors>`_
       | `sum of the proper divisors - Yahoo search <https://search.yahoo.com/search?p=sum+of+the+proper+divisors>`_
       | `sum of the proper divisors - Baidu search <https://www.baidu.com/s?wd=sum+of+the+proper+divisors>`_
       | `sum of the proper divisors - Yandex search <https://www.yandex.com/search/?text=sum+of+the+proper+divisors>`_

.. [2] `Is there a formula to calculate the sum of all proper divisors of a number? - Mathematics Stack Exchange <https://math.stackexchange.com/questions/22721/is-there-a-formula-to-calculate-the-sum-of-all-proper-divisors-of-a-number>`_

.. [3] `[Golang] Get All Prime Factors of Integer Number <{filename}../09/go-find-all-prime-factors-of-integer-number%en.rst>`_
.. [4] `[Golang] Integer Exponentiation <{filename}../20/go-integer-exponentiation%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
