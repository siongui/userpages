[Golang] Largest Palindrome Product - Problem 4 - Project Euler
###############################################################

:date: 2017-04-24T00:29+08:00
:tags: Go, Golang, Algorithm, Project Euler
:category: Go
:summary: Go solution to largest palindrome made from the product of two 3-digit
          numbers. - Problem 4 - Project Euler.
:og_image: http://4.bp.blogspot.com/-htBpelT-eGw/UaSk7mpOz5I/AAAAAAAAAtU/Pq8LgZ4ZZeA/s640/pal.jpg
:adsu: yes

Go solution to largest palindrome product - Problem 4 - Project Euler. [2]_

**Problem**:

  A palindromic number reads the same both ways. The largest palindrome made from
  the product of two 2-digit numbers is 9009 = 91 × 99.

  Find the largest palindrome made from the product of two 3-digit numbers.

**Solution**:

  906609 = 913 x 993

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/gPupKocOXk>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"strconv"
  )

  func IsDecimalPalindromeNumber(n int64) bool {
  	if n < 0 {
  		n = -n
  	}

  	s := strconv.FormatInt(n, 10)
  	bound := (len(s) / 2) + 1
  	for i := 0; i < bound; i++ {
  		if s[i] != s[len(s)-1-i] {
  			return false
  		}
  	}
  	return true
  }

  func main() {
  	currentMax := 0
  	maxI := 0
  	maxJ := 0
  	for i := 999; i >= 100; i-- {
  		for j := 999; j >= 100; j-- {
  			product := i * j
  			if currentMax > product {
  				break
  			}
  			if IsDecimalPalindromeNumber(int64(product)) {
  				currentMax = product
  				maxI = i
  				maxJ = j
  				continue
  			}
  		}
  	}
  	fmt.Println("answer: ", currentMax)
  	fmt.Println("i: ", maxI)
  	fmt.Println("j: ", maxJ)
  }

.. adsu:: 2

----

References:

.. [1] `[Golang] Check if Integer is Palindromic Number <{filename}../15/go-check-if-integer-number-is-palindrome%en.rst>`_

.. [2] | `Largest palindrome product - Problem 4 - Project Euler <https://projecteuler.net/index.php?section=problems&id=4>`_
       | `Project Euler <https://projecteuler.net/>`_

.. [3] | `C programming examples | Programming Simplified <http://www.programmingsimplified.com/c-program-examples>`_
       | `GeeksforGeeks | A computer science portal for geeks <http://www.geeksforgeeks.org/>`_

.. [4] | `A Competitive Programmer's Handbook | Hacker News <https://news.ycombinator.com/item?id=14115826>`_
       | `Competitive Programmer's Handbook <https://cses.fi/book.html>`_

.. [5] | `Data structures and algorithms interview questions and their solutions | Hacker News <https://news.ycombinator.com/item?id=14128145>`_
       | `500 Data structures and algorithms interview questions and their solutions - Techie Delight - Quora <https://techiedelight.quora.com/500-Data-structures-and-algorithms-interview-questions-and-their-solutions>`_

.. [6] `Rosetta Code <http://rosettacode.org/>`_

.. [7] | `技术面试宝典： 很全面的算法和数据结构知识（含代码实现） - 文章 - 伯乐在线 <http://blog.jobbole.com/110835/>`_
       | `GitHub - kdn251/interviews: Everything you need to know to get the job. <https://github.com/kdn251/Interviews>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _palindrome: https://www.google.com/search?q=palindrome+number
