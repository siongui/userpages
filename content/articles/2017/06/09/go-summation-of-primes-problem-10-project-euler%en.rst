[Golang] Summation of Primes - Problem 10 - Project Euler
#########################################################

:date: 2017-06-09T20:44+08:00
:tags: Go, Golang, Algorithm, Prime Number, Math, Project Euler, Type Casting,
       Type Conversion
:category: Go
:summary: Go solution to summation of primes
          - Problem 10 - Project Euler.
:og_image: http://www.texample.net/media/tikz/examples/PNG/eratosthenes-sieve.png
:adsu: yes

Go solution to summation of primes - Problem 10 - Project Euler. [1]_

**Problem**:

  The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

  Find the sum of all the primes below two million.


**Solution**:

  142913828922

  We use sieve of Eratosthenes [2]_ to find out all prime number under
  2,000,000. Note that we cannot use *int* type because of integer overflow
  [3]_. Use *int64* type instead.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/GxH-IkpFPS>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func SieveOfEratosthenes(n int64) []int64 {
  	// Create a boolean array "prime[0..n]" and initialize
  	// all entries it as true. A value in prime[i] will
  	// finally be false if i is Not a prime, else true.
  	integers := make([]bool, n+1)
  	for i := int64(2); i < n+1; i++ {
  		integers[i] = true
  	}

  	for p := int64(2); p*p <= n; p++ {
  		// If integers[p] is not changed, then it is a prime
  		if integers[p] == true {
  			// Update all multiples of p
  			for i := p * 2; i <= n; i += p {
  				integers[i] = false
  			}
  		}
  	}

  	// return all prime numbers <= n
  	var primes []int64
  	for p := int64(2); p <= n; p++ {
  		if integers[p] == true {
  			primes = append(primes, p)
  		}
  	}

  	return primes
  }

  func main() {
  	primes := SieveOfEratosthenes(2000000)
  	fmt.Println(len(primes))
  	sum := int64(0)
  	for _, prime := range primes {
  		sum1 := sum + prime
  		if sum1 < sum {
  			fmt.Println(sum)
  		} else {
  			sum = sum1
  		}
  	}
  	fmt.Println(sum)
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `Summation of primes - Problem 10 - Project Euler <https://projecteuler.net/problem=10>`_
.. [2] `[Golang] Sieve of Eratosthenes <{filename}../../04/17/go-sieve-of-eratosthenes%en.rst>`_
.. [3] | `golang check integer overflow - Google search <https://www.google.com/search?q=golang+check+integer+overflow>`_
       | `golang check integer overflow - DuckDuckGo search <https://duckduckgo.com/?q=golang+check+integer+overflow>`_
       | `golang check integer overflow - Ecosia search <https://www.ecosia.org/search?q=golang+check+integer+overflow>`_
       | `golang check integer overflow - Qwant search <https://www.qwant.com/?q=golang+check+integer+overflow>`_
       | `golang check integer overflow - Bing search <https://www.bing.com/search?q=golang+check+integer+overflow>`_
       | `golang check integer overflow - Yahoo search <https://search.yahoo.com/search?p=golang+check+integer+overflow>`_
       | `golang check integer overflow - Baidu search <https://www.baidu.com/s?wd=golang+check+integer+overflow>`_
       | `golang check integer overflow - Yandex search <https://www.yandex.com/search/?text=golang+check+integer+overflow>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _Sieve of Eratosthenes: https://www.google.com/search?q=Sieve+of+Eratosthenes
