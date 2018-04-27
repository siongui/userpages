[Golang] Quadratic primes - Problem 27 - Project Euler
######################################################

:date: 2018-04-27T22:43+08:00
:tags: Go, Golang, Algorithm, Math, Prime Number, Project Euler
:category: Go
:summary: Go solution to Quadratic primes
          - Problem 27 - Project Euler.
:og_image: https://snipcademy.com/code/img/challenges/math/quadratic-primes.svg
:adsu: yes

**Problem**: [1]_

.. raw:: html

  <script type="text/x-mathjax-config">
   MathJax.Hub.Config({
      jax: ["input/TeX", "output/HTML-CSS"],
      tex2jax: {
         inlineMath: [ ["$","$"], ["\\(","\\)"] ],
         displayMath: [ ["$$","$$"], ["\\[","\\]"] ],
         processEscapes: true
      },
      "HTML-CSS": { availableFonts: ["TeX"] }
   });
  </script>

  <script async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.4/MathJax.js?config=TeX-AMS_HTML-full,Safe"></script>

  <div class="problem_content" role="problem">

  <p>Euler discovered the remarkable quadratic formula:</p>
  <p style="text-align:center;">$n^2 + n + 41$</p>
  <p>It turns out that the formula will produce 40 primes for the consecutive integer values $0 \le n \le 39$. However, when $n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41$ is divisible by 41, and certainly when $n = 41, 41^2 + 41 + 41$ is clearly divisible by 41.</p>
  <p>The incredible formula $n^2 - 79n + 1601$ was discovered, which produces 80 primes for the consecutive values $0 \le n \le 79$. The product of the coefficients, −79 and 1601, is −126479.</p>
  <p>Considering quadratics of the form:</p>
  <blockquote>
  $n^2 + an + b$, where $|a| &lt; 1000$ and $|b| \le 1000$<br /><br /><div>where $|n|$ is the modulus/absolute value of $n$<br />e.g. $|11| = 11$ and $|-4| = 4$</div>
  </blockquote>
  <p>Find the product of the coefficients, $a$ and $b$, for the quadratic expression that produces the maximum number of primes for consecutive values of $n$, starting with $n = 0$.</p>
  </div>


**Solution**:

  (a,b,n) = (-61,971,71)

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/-CE0EQOmBlz>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

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

  func FindConsecutiveN(a, b int) int {
  	isPrime := true
  	n := 0
  	for isPrime {
  		result := n*n + a*n + b
  		isPrime = IsPrime(result)
  		if isPrime {
  			n++
  		} else {
  			return n
  		}
  	}
  	return -1
  }

  func main() {
  	maxN := 0
  	maxA := -11111
  	maxB := -11111
  	for a := -999; a < 1000; a++ {
  		for b := -1000; b < 1001; b++ {
  			n := FindConsecutiveN(a, b)
  			if n == -1 {
  				panic("n cannot be -1")
  			}
  			if n > maxN {
  				maxN = n
  				maxA = a
  				maxB = b
  				fmt.Println("current max (a,b,n)", a, b, n)
  			}
  		}
  	}
  	fmt.Println("max (a,b,n)", maxA, maxB, maxN)
  }

Output from `Go Playground`_:

.. code-block:: txt

  current max (a,b,n) -999 2 1
  current max (a,b,n) -996 997 2
  current max (a,b,n) -499 997 3
  current max (a,b,n) -325 977 4
  current max (a,b,n) -245 977 5
  current max (a,b,n) -197 983 6
  current max (a,b,n) -163 983 7
  current max (a,b,n) -131 941 8
  current max (a,b,n) -121 947 9
  current max (a,b,n) -105 967 11
  current max (a,b,n) -61 971 71
  max (a,b,n) -61 971 71

.. adsu:: 2

Tested on:

- ``Go 1.10.1``, ``Ubuntu 17.10``
- `Go Playground`_

----

References:

.. [1] `Quadratic primes - Problem 27 - Project Euler <https://projecteuler.net/problem=27>`_

.. _Go Playground: https://play.golang.org/
