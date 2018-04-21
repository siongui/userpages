Lemoine’s Conjecture
####################

:date: 2018-04-21T23:33+08:00
:tags: Go, Golang, Math, Algorithm, Prime Number
:category: Algorithm
:summary: 2n + 1 = p + 2q always has a solution in primes p and q (not
          necessarily distinct) for n > 2. Write a Go program to find p and q
          for given odd number greater than 5.
:og_image: 
:adsu: yes


`Lemoine’s Conjecture`_ says any odd integer greater than 5 can be represented
as the sum of an odd prime number and an even semiprime. Another statement which
is suitable for programming is that 2n + 1 = p + 2q always has a solution in
primes p and q (not necessarily distinct) for n > 2. We will write a Go program
to find p and q for given odd number greater than 5.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/67PP7CD6xIR>`__
   :class: align-center

.. code-block:: go

  package lemoine

  import (
  	"fmt"
  )

  func IsPrime(n int) bool {
  	if n < 2 {
  		return false
  	}

  	for i := 2; i*i <= n; i++ {
  		if n%i == 0 {
  			return false
  		}
  	}
  	return true
  }

  func Lemoine(n int) {
  	if n <= 5 || n%2 == 0 {
  		panic("n must be greater than 5 and must be odd")
  	}

  	m := make(map[int]int)

  	for q := 1; q <= n/2; q++ {
  		p := n - 2*q

  		if IsPrime(p) && IsPrime(q) {
  			m[p] = q
  		}
  	}

  	for p, q := range m {
  		fmt.Println(n, "=", p, "+ ( 2 *", q, ")")
  	}
  }


Example:

.. code-block:: go

  package lemoine

  import (
  	"testing"
  )

  func TestIsPrime(t *testing.T) {
  	if IsPrime(97) != true {
  		t.Error("97 fail")
  	}
  	if IsPrime(98) != false {
  		t.Error("98 fail")
  	}
  }

  func TestLemoine(t *testing.T) {
  	Lemoine(39)
  	Lemoine(7)
  	Lemoine(101)
  }

.. adsu:: 2

Tested on:

- ``Ubuntu Linux 17.10``, ``Go 1.10.1``
- `Go Playground`_

----

References:

.. [1] `Lemoine's Conjecture - GeeksforGeeks <https://www.geeksforgeeks.org/lemoines-conjecture/>`_
.. [2] `[Golang] Sieve of Eratosthenes <{filename}/articles/2017/04/17/go-sieve-of-eratosthenes%en.rst>`_

.. _Go Playground: https://play.golang.org/
.. _Lemoine's conjecture: https://www.google.com/search?q=Lemoine's+Conjecture
