[Golang] Greatest Common Divisor via Euclidean Algorithm
########################################################

:date: 2017-05-14T21:16+08:00
:tags: Go, Golang, Algorithm, Math
:category: Go
:summary: Compute the greatest common divisor (GCD) via Euclidean algorithm
          in Go programming language.
:og_image: https://upload.wikimedia.org/wikipedia/commons/thumb/3/37/Euclid%27s_algorithm_Book_VII_Proposition_2_3.png/300px-Euclid%27s_algorithm_Book_VII_Proposition_2_3.png
:adsu: yes

Compute the greatest common divisor (GCD) via Euclidean algorithm
in Go programming language.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/tYbxNvvHpG>`__
   :class: align-center

.. code-block:: go

  // greatest common divisor (GCD) via Euclidean algorithm
  func GCD(a, b int) int {
  	for b != 0 {
  		t := b
  		b = a % b
  		a = t
  	}
  	return a
  }

Tested on:

- `Go Playground`_

----

References:

.. [1] `Euclidean algorithm - Wikipedia <https://en.wikipedia.org/wiki/Euclidean_algorithm>`_
.. adsu:: 2
.. [2] `Basic and Extended Euclidean algorithms - GeeksforGeeks <http://www.geeksforgeeks.org/basic-and-extended-euclidean-algorithms/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _Sieve of Eratosthenes: https://www.google.com/search?q=Sieve+of+Eratosthenes
