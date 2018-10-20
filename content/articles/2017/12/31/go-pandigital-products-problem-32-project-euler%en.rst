[Golang] Pandigital products - Problem 32 - Project Euler
#########################################################

:date: 2018-10-21T01:09+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler, Type Casting,
       Type Conversion, String Manipulation
:category: Go
:summary: Go solution to Pandigital products
          - Problem 32 - Project Euler.
:og_image: https://3.bp.blogspot.com/-NSwrmX1HWR0/VzXmheNLUfI/AAAAAAAAB7U/CaFVbOBs8uMqjV5uDdTCe2VHGPZKPYqVQCLcB/s1600/project%2BEuler%2Bproblem%2B32%2Bwith%2Banswer.png
:adsu: yes

**Problem**: [1]_

  We shall say that an n-digit number is pandigital if it makes use of all the
  digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1
  through 5 pandigital.

  The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing
  multiplicand, multiplier, and product is 1 through 9 pandigital.

  Find the sum of all products whose multiplicand/multiplier/product identity
  can be written as a 1 through 9 pandigital.

  HINT: Some products can be obtained in more than one way so be sure to only
  include it once in your sum.

**Solution**:

  The form of the identity must be like:

  | **M * MMMM = MMMM**
  | **MM * MMM = MMMM**

  Brute-forcely check all permutations [2]_ of 1~9 in above form. We will get
  the following satisfying identity.

  | 12 * 483 = 5796
  | 18 * 297 = 5346
  | 27 * 198 = 5346
  | 28 * 157 = 4396
  | 39 * 186 = 7254
  | 42 * 138 = 5796
  | 4 * 1738 = 6952
  | 4 * 1963 = 7852
  | 48 * 159 = 7632
  | 56370

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/Ery3epswUSJ>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"strconv"
  )

  var sum int64 = 0

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
  		checkProduct(str)
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

  // MxMMMM=MMMM
  // MMxMMM=MMMM
  func checkProduct(s string) {
  	// check M x MMMM = MMMM
  	m1, _ := strconv.ParseInt(s[0:1], 10, 0)
  	m2, _ := strconv.ParseInt(s[1:5], 10, 0)
  	prd, _ := strconv.ParseInt(s[5:9], 10, 0)
  	if m1*m2 == prd {
  		fmt.Printf("%d * %d = %d\n", m1, m2, prd)
  		sum += prd
  	}

  	// check MM x MMM = MMMM
  	m1, _ = strconv.ParseInt(s[0:2], 10, 0)
  	m2, _ = strconv.ParseInt(s[2:5], 10, 0)
  	prd, _ = strconv.ParseInt(s[5:9], 10, 0)
  	if m1*m2 == prd {
  		fmt.Printf("%d * %d = %d\n", m1, m2, prd)
  		sum += prd
  	}
  }

  func main() {
  	str := "123456789"
  	permutations(str, 0, len(str))

  	fmt.Println(sum)
  }

.. adsu:: 2

----

Test on:

- ``Ubuntu 18.04``, ``Go 1.11.1``
- `Go Playground`_

References:

.. [1] `Pandigital products - Problem 32 - Project Euler <https://projecteuler.net/problem=32>`_
.. [2] `[Golang] All Permutations of Given String With Distinct Characters <{filename}/articles/2017/03/11/go-all-permutations-of-given-string-with-all-distinct-characters%en.rst>`_
.. [3] `[Golang] Lexicographic permutations - Problem 24 - Project Euler <{filename}/articles/2017/12/31/go-lexicographic-permutations-problem-24-project-euler%en.rst>`_
.. [4] `[Golang] Type Conversion between String and Integer <{filename}/articles/2017/04/14/go-string-int-type-casting%en.rst>`_

.. _Go Playground: https://play.golang.org/
