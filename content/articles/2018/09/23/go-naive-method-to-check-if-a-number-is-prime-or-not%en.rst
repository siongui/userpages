[Golang] Naive Method for Primality Test
########################################

:date: 2018-09-23T02:32+08:00
:tags: Go, Golang, Math, Algorithm, Prime Number
:category: Algorithm
:summary: Naive method for primality test in Go: Given a natural number n, if n
          is divisible by any number from 2 to square root of n, then n is
          composite. Otherwise n is prime.
:og_image: http://images.slideplayer.com/5/1596895/slides/slide_2.jpg
:adsu: yes


Naive method for primality test in Go: Given a natural number n, if n is
divisible by any number from 2 to square root of n, then n is composite.
Otherwise n is prime.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/NLzvyKmbUQb>`__
   :class: align-center

.. code-block:: go

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

Tested on: `Go Playground`_

----

.. adsu:: 2

References:

.. [1] `[Golang] Primality Test - Optimized School Method <{filename}/articles/2017/04/19/go-primality-test-optimized-school-method%en.rst>`_
.. [2] `Lemoine's Conjecture - GeeksforGeeks <https://www.geeksforgeeks.org/lemoines-conjecture/>`_
.. [3] `Lemoine’s Conjecture <{filename}/articles/2018/04/21/lemoine-conjecture%en.rst>`_
.. [4] `[Golang] Sieve of Eratosthenes <{filename}/articles/2017/04/17/go-sieve-of-eratosthenes%en.rst>`_
.. [5] | `primality test - Google search <https://www.google.com/search?q=primality+test>`_
       | `primality test - DuckDuckGo search <https://duckduckgo.com/?q=primality+test>`_
       | `primality test - Ecosia search <https://www.ecosia.org/search?q=primality+test>`_
       | `primality test - Qwant search <https://www.qwant.com/?q=primality+test>`_
       | `primality test - Bing search <https://www.bing.com/search?q=primality+test>`_
       | `primality test - Yahoo search <https://search.yahoo.com/search?p=primality+test>`_
       | `primality test - Baidu search <https://www.baidu.com/s?wd=primality+test>`_
       | `primality test - Yandex search <https://www.yandex.com/search/?text=primality+test>`_

.. _Go Playground: https://play.golang.org/
