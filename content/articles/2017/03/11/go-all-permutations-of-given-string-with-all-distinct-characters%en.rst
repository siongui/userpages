[Golang] All Permutations of Given String With Distinct Characters
##################################################################

:date: 2017-03-11T01:24+08:00
:tags: Go, Golang, String Manipulation, Algorithm
:category: Go
:summary: Find all permutations of a given string containing all distinct
          characters in Go_ programming language.
:adsu: yes


Question
++++++++

Given a string containing all distinct characters, find all permutations of it.
[2]_

PS: I came to this question via HN [1]_, and it looks interesting. So I try to
port the code from `C++`_ to Go.

Answer
++++++

Use Backtracking_ algorithm. See `original post`_ for explanation.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/TC39OY3euz>`_
      :class: align-center

.. code-block:: go

  package main

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
  		println(str)
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
  	str := "ABC"
  	permutations(str, 0, len(str))
  }

.. adsu:: 2

----

Tested on: `Go Playground`_

----

References:

.. [1] | `Data structures and algorithms problems in C++ using STL | Hacker News <https://news.ycombinator.com/item?id=13836714>`_
       | `Data structures and algorithms problems in C++ using STL <http://www.techiedelight.com/data-structures-and-algorithms-interview-questions-stl/>`_
.. [2] `Find all Permutations of a given string - Techie Delight <http://www.techiedelight.com/find-permutations-given-string/>`_
.. [3] `Strings, bytes, runes and characters in Go - The Go Blog <https://blog.golang.org/strings>`_
.. [4] `Data Structures and Algorithms Problems | Hacker News <https://news.ycombinator.com/item?id=14043891>`_
.. [5] `Common Algo Problem Solutions | Hacker News <https://news.ycombinator.com/item?id=14064698>`_
.. [6] `LeetCode Online Judge <https://leetcode.com/>`_
.. [7] | `Generating all permutations, combinations, and power set of a string (2012) | Hacker News <https://news.ycombinator.com/item?id=14272847>`_
       | `Exceptional Code: Generating all permutations, combinations, and power set of a string (or set of numbers) <http://exceptional-code.blogspot.com/2012/09/generating-all-permutations.html>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _C++: https://www.google.com/search?q=C%2B%2B
.. _backtracking: https://www.google.com/search?q=backtracking
.. _Go Playground: https://play.golang.org/
.. _original post: http://www.techiedelight.com/find-permutations-given-string/
