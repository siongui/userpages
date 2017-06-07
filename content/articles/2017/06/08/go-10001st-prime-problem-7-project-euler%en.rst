[Golang] 10001st Prime - Problem 7 - Project Euler
##################################################

:date: 2017-06-08T00:55+08:00
:tags: Go, Golang, Algorithm, Prime Number, Math, Project Euler
:category: Go
:summary: Go solution to 10001st prime
          - Problem 7 - Project Euler.
:og_image: http://www.texample.net/media/tikz/examples/PNG/eratosthenes-sieve.png
:adsu: yes

Go solution to 10001st prime - Problem 7 - Project Euler. [1]_

**Problem**:

  By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
  that the 6th prime is 13.

  What is the 10 001st prime number?


**Solution**:

  According to `prime number theorem`_, the 10001st prime number is probably
  located at somewhere around 1,000,000. We use `Sieve of Eratosthenes`_ [2]_ to
  find all prime numbers under 1,100,000. The number of primes is 10,453, and
  the 10001st prime is **104743**.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/vzAjdZHnkb>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func SieveOfEratosthenes(n int) []int {
  	// Create a boolean array "prime[0..n]" and initialize
  	// all entries it as true. A value in prime[i] will
  	// finally be false if i is Not a prime, else true.
  	integers := make([]bool, n+1)
  	for i := 2; i < n+1; i++ {
  		integers[i] = true
  	}

  	for p := 2; p*p <= n; p++ {
  		// If integers[p] is not changed, then it is a prime
  		if integers[p] == true {
  			// Update all multiples of p
  			for i := p * 2; i <= n; i += p {
  				integers[i] = false
  			}
  		}
  	}

  	// return all prime numbers <= n
  	var primes []int
  	for p := 2; p <= n; p++ {
  		if integers[p] == true {
  			primes = append(primes, p)
  		}
  	}

  	return primes
  }

  func main() {
  	primes := SieveOfEratosthenes(110000)
  	fmt.Println(len(primes))
  	if len(primes) > 10001 {
  		fmt.Println(primes[10000])
  	}
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `10001st prime - Problem 7 - Project Euler <https://projecteuler.net/problem=7>`_
.. [2] `[Golang] Sieve of Eratosthenes <{filename}../../04/17/go-sieve-of-eratosthenes%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _prime number theorem: https://en.wikipedia.org/wiki/Prime_number_theorem
.. _Sieve of Eratosthenes: https://www.google.com/search?q=Sieve+of+Eratosthenes
