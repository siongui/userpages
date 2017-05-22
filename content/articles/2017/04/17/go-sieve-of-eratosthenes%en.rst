[Golang] Sieve of Eratosthenes
##############################

:date: 2017-04-17T16:43+08:00
:tags: Go, Golang, Algorithm, Prime Number
:category: Go
:summary: Go implementation of *Sieve of Eratosthenes*.
:og_image: http://www.texample.net/media/tikz/examples/PNG/eratosthenes-sieve.png
:adsu: yes

Go implementation of `Sieve of Eratosthenes`_.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/66tBiUdvy2>`__
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
  	fmt.Println(SieveOfEratosthenes(20))
  	fmt.Println(SieveOfEratosthenes(30))
  	fmt.Println(SieveOfEratosthenes(40))
  }

.. adsu:: 2

Tested on:

- `Go Playground`_

----

References:

.. [1] | `Sieve of Eratosthenes - GeeksforGeeks <http://www.geeksforgeeks.org/sieve-of-eratosthenes/>`_

.. [2] | `C programming examples | Programming Simplified <http://www.programmingsimplified.com/c-program-examples>`_
       | `Project Euler <https://projecteuler.net/>`_

.. [3] | `A Competitive Programmer's Handbook | Hacker News <https://news.ycombinator.com/item?id=14115826>`_
       | `Competitive Programmer's Handbook <https://cses.fi/book.html>`_

.. [4] | `Data structures and algorithms interview questions and their solutions | Hacker News <https://news.ycombinator.com/item?id=14128145>`_
       | `500 Data structures and algorithms interview questions and their solutions - Techie Delight - Quora <https://techiedelight.quora.com/500-Data-structures-and-algorithms-interview-questions-and-their-solutions>`_
       | `Algorithms and Data Structures Problem Set | Hacker News <https://news.ycombinator.com/item?id=14385924>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _Sieve of Eratosthenes: https://www.google.com/search?q=Sieve+of+Eratosthenes
