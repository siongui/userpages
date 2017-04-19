[Golang] Primality Test - Optimized School Method
#################################################

:date: 2017-04-19T23:35+08:00
:tags: Go, Golang, Algorithm
:category: Go
:summary: Go implementation of primality test - optimized school method.
:og_image: http://images.slideplayer.com/5/1596895/slides/slide_2.jpg
:adsu: yes

Go implementation of `primality test`_ - optimized school method. [1]_

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/gdzxMsLuv0>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func IsPrime(n int) bool {
  	// Corner cases
  	if n <= 1 {
  		return false
  	}
  	if n <= 3 {
  		return true
  	}

  	// This is checked so that we can skip
  	// middle five numbers in below loop
  	if n%2 == 0 || n%3 == 0 {
  		return false
  	}

  	for i := 5; i*i <= n; i = i + 6 {
  		if n%i == 0 || n%(i+2) == 0 {
  			return false
  		}
  	}

  	return true
  }

.. adsu:: 2

Tested on:

- `Go Playground`_

----

References:

.. [1] `Primality Test | Set 1 (Introduction and School Method) - GeeksforGeeks <http://www.geeksforgeeks.org/primality-test-set-1-introduction-and-school-method/>`_

.. [2] | `C programming examples | Programming Simplified <http://www.programmingsimplified.com/c-program-examples>`_
       | `Project Euler <https://projecteuler.net/>`_


.. [3] | `A Competitive Programmer's Handbook | Hacker News <https://news.ycombinator.com/item?id=14115826>`_
       | `Competitive Programmer's Handbook <https://cses.fi/book.html>`_

.. [4] | `Data structures and algorithms interview questions and their solutions | Hacker News <https://news.ycombinator.com/item?id=14128145>`_
       | `500 Data structures and algorithms interview questions and their solutions - Techie Delight - Quora <https://techiedelight.quora.com/500-Data-structures-and-algorithms-interview-questions-and-their-solutions>`_

.. [5] `Rosetta Code <http://rosettacode.org/>`_

.. [6] | `技术面试宝典： 很全面的算法和数据结构知识（含代码实现） - 文章 - 伯乐在线 <http://blog.jobbole.com/110835/>`_
       | `GitHub - kdn251/interviews: Everything you need to know to get the job. <https://github.com/kdn251/Interviews>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _primality test: https://www.google.com/search?q=primality+test
