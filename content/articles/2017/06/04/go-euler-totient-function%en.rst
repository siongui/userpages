[Golang] Euler's Totient Function
#################################

:date: 2017-06-04T01:14+08:00
:tags: Go, Golang, Algorithm, Math, Prime Number
:category: Go
:summary: Euler's totient function implementation in Go programming language.
:og_image: http://mathworld.wolfram.com/images/eps-gif/TotientFunction_800.gif
:adsu: yes

Implementation of Euler's totient function in Go programming language.

**Euler's totient function** :math:`\phi (n)` counts the positive integers that
are relatively prime to :math:`n`, i.e., the number of positive integers whose
GCD (Greatest Common Divisor) [2]_ with :math:`n` is 1.

According to above definition, the naive solution to implement :math:`\phi (n)`
is as follows:

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/crIx9y2E3M>`__
   :class: align-center

.. code-block:: go

  // greatest common divisor (GCD) via Euclidean algorithm
  func GCD(a, b uint) uint {
  	for b != 0 {
  		t := b
  		b = a % b
  		a = t
  	}
  	return a
  }

  // Euler’s Totient Function
  func Phi(n uint) uint {
  	var result uint = 1
  	var i uint
  	for i = 2; i < n; i++ {
  		if GCD(i, n) == 1 {
  			result++
  		}
  	}
  	return result
  }


According to the GeeksforGeeks post [1]_, there is a better solution based on
Euler’s product formula without floating point calculations. I port the solution
from C to Go:

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/XVY63BVSBa>`__
   :class: align-center

.. code-block:: go

  func Phi(n uint) (result uint) {
  	// Initialize result as n
  	result = n

  	// Consider all prime factors of n and subtract their
  	// multiples from result
  	var p uint
  	for p = 2; p*p <= n; p++ {
  		// Check if p is a prime factor.
  		if n%p == 0 {
  			// If yes, then update n and result
  			for n%p == 0 {
  				n /= p
  			}
  			result -= (result / p)
  		}
  	}

  	// If n has a prime factor greater than sqrt(n)
  	// (There can be at-most one such prime factor)
  	if n > 1 {
  		result -= (result / n)
  	}
  	return result
  }

Tested on:

- `Go Playground`_

----

References:

.. [1] `Euler's Totient Function - GeeksforGeeks <http://www.geeksforgeeks.org/eulers-totient-function/>`_
.. [2] `[Golang] Greatest Common Divisor via Euclidean Algorithm <{filename}../../05/14/go-gcd-via-euclidean-algorithm%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _Sieve of Eratosthenes: https://www.google.com/search?q=Sieve+of+Eratosthenes
