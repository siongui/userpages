[Golang] Lexicographic permutations - Problem 24 - Project Euler
################################################################

:date: 2018-03-02T21:40+08:00
:tags: Go, Golang, Algorithm, Project Euler
:category: Go
:summary: Go solution to Lexicographic permutations
          - Problem 24 - Project Euler.
:og_image: http://images.slideplayer.com/25/7929437/slides/slide_4.jpg
:adsu: yes

**Problem**: [1]_

  A permutation is an ordered arrangement of objects. For example, 3124 is one
  possible permutation of the digits 1, 2, 3 and 4. If all of the permutations
  are listed numerically or alphabetically, we call it lexicographic order. The
  lexicographic permutations of 0, 1 and 2 are:

  .. container:: align-center

    012   021   102   120   201   210

  What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4,
  5, 6, 7, 8 and 9?


**Solution**:

  My previous post [2]_ will print the permutations of distinct characters in
  lexicographic order, so I count the prints and get the answer.

  2785960341

.. code-block:: go

  package main

  var count = 0

  // Swap the i-th byte and j-th byte of the string
  func swap(s string, i, j int) string {
  	var result []byte
  	for k := 0; k < len(s); k++ {
  		if k == i {
  			result = append(result, s[j])
  		} else if k == j {
  			result = append(result, s[i])
  		} else {
  			result = append(result, s[k])
  		}
  	}
  	return string(result)
  }

  // Function to find all Permutations of a given string str[i:n]
  // containing all distinct characters
  func permutations(str string, i, n int) {
  	// base condition
  	if i == n-1 {
  		count += 1
  		if count == 1000000 {
  			println(str)
  		}
  		return
  	}

  	// process each character of the remaining string
  	for j := i; j < n; j++ {
  		// swap character at index i with current character
  		str = swap(str, i, j)

  		// recursion for string [i+1:n]
  		permutations(str, i+1, n)

  		// backtrack (restore the string to its original state)
  		str = swap(str, i, j)
  	}
  }

  func main() {
  	str := "0123456789"
  	permutations(str, 0, len(str))
  }

.. adsu:: 2

Tested on: ``Go 1.10`` (not runable on `Go Playground`_ because process took
too long)

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/qawX86kmTFG>`__
   :class: align-center

----

References:

.. [1] `Lexicographic permutations - Problem 24 - Project Euler <https://projecteuler.net/problem=24>`_
.. [2] `[Golang] All Permutations of Given String With Distinct Characters <{filename}/articles/2017/03/11/go-all-permutations-of-given-string-with-all-distinct-characters%en.rst>`_

.. _Go Playground: https://play.golang.org/
