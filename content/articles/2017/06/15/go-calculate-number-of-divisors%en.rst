[Golang] Calculate Number of Divisors
#####################################

:date: 2017-06-15T22:56+08:00
:tags: Go, Golang, Algorithm, Prime Number, Math
:category: Go
:summary: Given a natural number (integer), calculate the number of divisors of
          it in Go programming language.
:og_image: https://i0.wp.com/vishalkataria.in/wp-content/uploads/2016/08/Number-of-Divisors.jpeg?resize=940%2C550
:adsu: yes


Assume :math:`n` is a natural number, calculate the number of divisors of
:math:`n`.

Let the prime factorization of :math:`n` be:

  :math:`n = {p_1}^{a_1} \times {p_2}^{a_2} \times ... \times {p_k}^{a_k}`

number of divisors of n is:

  :math:`(a_1 + 1)(a_2 + 1) ... (a_k + 1)`


Prime Factorization
+++++++++++++++++++

The following code will return prime factorization of a given number n:

.. code-block:: go

  // Prime factorization of a given number
  //
  // A map is returned.
  //
  //   key of map: prime
  //   value of map: prime exponents
  //
  func PrimeFactorization(n int) (pfs map[int]int) {
  	pfs = make(map[int]int)

  	// Get the number of 2s that divide n
  	for n%2 == 0 {
  		if _, ok := pfs[2]; ok {
  			pfs[2] += 1
  		} else {
  			pfs[2] = 1
  		}
  		n = n / 2
  	}

  	// n must be odd at this point. so we can skip one element
  	// (note i = i + 2)
  	for i := 3; i*i <= n; i = i + 2 {
  		// while i divides n, append i and divide n
  		for n%i == 0 {
  			if _, ok := pfs[i]; ok {
  				pfs[i] += 1
  			} else {
  				pfs[i] = 1
  			}
  			n = n / i
  		}
  	}

  	// This condition is to handle the case when n is a prime number
  	// greater than 2
  	if n > 2 {
  		pfs[n] = 1
  	}

  	return
  }

See [2]_ [3]_ for more details.


Number of Divisors
++++++++++++++++++

Based on the result of prime factorization of previous section, we can calculate
number of divisors of a given number n as follows:

.. code-block:: go

  // Calculate number of divisors of a given number
  func NumberOfDivisors(n int) int {
  	pfs := PrimeFactorization(n)

  	num := 1
  	for _, exponents := range pfs {
  		num *= (exponents + 1)
  	}

  	return num
  }

See [1]_ for more details.


Application
+++++++++++

The problem 12 of Project Euler is about finding the number of divisors of
triangular number. See [4]_ for more details.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/v1dPoJoSAu>`__
   :class: align-center

Tested on: `Go Playground`_

----

References:

.. [1] | `calculate number of divisors - Google search <https://www.google.com/search?q=calculate+number+of+divisors>`_
       | `calculate number of divisors - DuckDuckGo search <https://duckduckgo.com/?q=calculate+number+of+divisors>`_
       | `calculate number of divisors - Ecosia search <https://www.ecosia.org/search?q=calculate+number+of+divisors>`_
       | `calculate number of divisors - Qwant search <https://www.qwant.com/?q=calculate+number+of+divisors>`_
       | `calculate number of divisors - Bing search <https://www.bing.com/search?q=calculate+number+of+divisors>`_
       | `calculate number of divisors - Yahoo search <https://search.yahoo.com/search?p=calculate+number+of+divisors>`_
       | `calculate number of divisors - Baidu search <https://www.baidu.com/s?wd=calculate+number+of+divisors>`_
       | `calculate number of divisors - Yandex search <https://www.yandex.com/search/?text=calculate+number+of+divisors>`_
.. [2] `[Golang] Get All Prime Factors of Integer Number <{filename}../../05/09/go-find-all-prime-factors-of-integer-number%en.rst>`_
.. [3] `[Golang] Sum of the Proper Divisors (Factors) <{filename}../../05/19/go-sum-of-proper-factors%en.rst>`_
.. [4] `[Golang] Highly Divisible Triangular Number - Problem 12 - Project Euler <{filename}../14/go-highly-divisible-triangular-number-problem-12-project-euler%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
