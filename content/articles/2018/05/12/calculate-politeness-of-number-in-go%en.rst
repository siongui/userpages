Calcualte Politeness of Number in Golang
########################################

:date: 2018-05-12T23:40+08:00
:tags: Go, Golang, Algorithm, Prime Number, Math
:category: Go
:summary: Calculate politeness of a number in Go, i.e., the number of ways it
          can be expressed as the sum of consecutive integers
:og_image: https://upload.wikimedia.org/wikipedia/commons/thumb/0/0f/Young_456_French.svg/640px-Young_456_French.svg.png
:adsu: yes


See this interesting problem from GeeksforGeeks [1]_, and there is detailed
description in Wikipedia.

The following comes from Wikipedia:

  The politeness of a positive number is defined as the number of ways it can be
  expressed as the sum of consecutive integers. For instance, the politeness of
  9 is 2 because it has two odd divisors, 3 and itself, and two polite
  representations

    9 = 2 + 3 + 4 = 4 + 5;

The easy way to calculate the politeness also comes from Wikipedia:

  An easy way of calculating the politeness of a positive number is that of
  decomposing the number into its prime factors, taking the powers of all prime
  factors greater than 2, adding 1 to all of them, multiplying the numbers thus
  obtained with each other and subtracting 1.

The following is implementation of above algorithm.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/jiKlmYy3770>`__
   :class: align-center

.. code-block:: go

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

  // Algorithm from wiki: An easy way of calculating the politeness of a positive
  // number is that of decomposing the number into its prime factors, taking the
  // powers of all prime factors greater than 2, adding 1 to all of them,
  // multiplying the numbers thus obtained with each other and subtracting 1.
  func CalculatePoliteness(n int) int {
  	pfs := PrimeFactors(n)

  	// key: prime
  	// value: prime exponent
  	m := make(map[int]int)
  	for _, prime := range pfs {
  		_, ok := m[prime]
  		if ok {
  			m[prime] += 1
  		} else {
  			m[prime] = 1
  		}
  	}

  	politeness := 1
  	for prime, exponent := range m {
  		if prime > 2 {
  			politeness *= (exponent + 1)
  		}
  	}
  	politeness -= 1

  	return politeness
  }

Test:

.. code-block:: go

  import (
  	"testing"
  )

  func TestCalculatePoliteness(t *testing.T) {
  	if CalculatePoliteness(9) != 2 {
  		t.Error("politeness of 9 is not 2")
  		return
  	}
  	if CalculatePoliteness(15) != 3 {
  		t.Error("politeness of 15 is not 3")
  		return
  	}
  	if CalculatePoliteness(90) != 5 {
  		t.Error("politeness of 90 is not 5")
  		return
  	}
  }

.. adsu:: 2

----

Tested on:

- `Go Playground`_
- ``Ubuntu Linux 18.04``, ``Go 1.10.2``

References:

.. [1] | `Find politeness of a number - GeeksforGeeks <https://www.geeksforgeeks.org/find-politeness-number/>`_
       | `Polite number - Wikipedia <https://en.wikipedia.org/wiki/Polite_number>`_
.. [2] `[Golang] Get All Prime Factors of Integer Number <{filename}/articles/2017/05/09/go-find-all-prime-factors-of-integer-number%en.rst>`_

.. _Go Playground: https://play.golang.org/
