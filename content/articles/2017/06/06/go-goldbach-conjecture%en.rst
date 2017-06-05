[Golang] Goldbach's conjecture
##############################

:date: 2017-06-06T00:18+08:00
:tags: Go, Golang, Algorithm, Prime Number, Math
:category: Go
:summary: *Goldbach's conjecture* - Every even integer greater than 2 can be
          written as the sum of two primes. Given a positive even integer, write
          a Go program to find the two primes.
:og_image: http://www.mathsisgoodforyou.com/images/imagesforworksheets/goldbachsletter.gif
:adsu: yes

`Goldbach's conjecture`_ - Every even integer greater than 2 can be written as
the sum of two primes.

Given a positive even integer, write a Go program to find the two primes.

Strategy:

- Find all primes under the given number by sieve of Eratosthenes. [4]_
- Find the two primes in previous step that sum up to the given number.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/Z53Fd-iAMX>`__
   :class: align-center

.. code-block:: go

  // Find all primes <= n
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

  // Given a positive even integer n,
  // return the two primes that sum up to n.
  func Goldbach(n int) (p1, p2 int) {
  	if n < 3 {
  		return
  	}
  	if n%2 == 1 {
  		return
  	}

  	primes := SieveOfEratosthenes(n)

  	m := make(map[int]bool)
  	for i := 0; i < len(primes); i++ {
  		m[primes[i]] = true
  	}

  	for p, _ := range m {
  		if _, ok := m[n-p]; ok {
  			p1 = p
  			p2 = n - p
  			return
  		}
  	}
  	return
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `A list of practical projects that anyone can solve in any programming language | Hacker News <https://news.ycombinator.com/item?id=14481941>`_
.. [2] | `Prolog Problems - Prolog Site <https://sites.google.com/site/prologsite/prolog-problems>`_
       | `2. Arithmetic - Prolog Site <https://sites.google.com/site/prologsite/prolog-problems/2>`_
.. [3] `Program for Goldbach’s Conjecture (Two Primes with given Sum) - GeeksforGeeks <http://www.geeksforgeeks.org/program-for-goldbachs-conjecture-two-primes-with-given-sum/>`_
.. [4] `[Golang] Sieve of Eratosthenes <{filename}../../04/17/go-sieve-of-eratosthenes%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _Goldbach's conjecture: https://www.google.com/search?q=Goldbach's+conjecture
