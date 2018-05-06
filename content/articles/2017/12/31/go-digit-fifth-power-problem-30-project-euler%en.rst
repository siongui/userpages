[Golang] Digit fifth powers - Problem 30 - Project Euler
########################################################

:date: 2018-05-06T23:22+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler, Type Casting, Type Conversion
:category: Go
:summary: Go solution to Digit fifth powers
          - Problem 30 - Project Euler.
:og_image: https://3.bp.blogspot.com/-xCdh7kWE6Dc/VyzTC4hCAaI/AAAAAAAAB4o/NGXhxXaZ5GUnB3NcBgK0bP0EO0It29HcgCLcB/s1600/project%2BEuler%2Bproblem30%2Bwith%2Banswer.png
:adsu: yes

**Problem**: [1]_

  Surprisingly there are only three numbers that can be written as the sum of
  fourth powers of their digits:

    | 1634 = 1^4 + 6^4 + 3^4 + 4^4
    | 8208 = 8^4 + 2^4 + 0^4 + 8^4
    | 9474 = 9^4 + 4^4 + 7^4 + 4^4

  As 1 = 1^4 is not a sum it is not included.

  The sum of these numbers is 1634 + 8208 + 9474 = 19316.

  Find the sum of all the numbers that can be written as the sum of fifth powers
  of their digits.


**Solution**:

  443839

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/TA-oxzQX4SB>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"strconv"
  )

  func IsFifthPowerEqual(n int) bool {
  	s := strconv.Itoa(n)
  	sum := 0
  	for _, digit := range s {
  		d, err := strconv.Atoi(string(digit))
  		if err != nil {
  			panic(err)
  		}
  		sum += d * d * d * d * d
  	}
  	if sum == n {
  		return true
  	}
  	return false
  }

  func main() {
  	d95 := 9 * 9 * 9 * 9 * 9
  	fmt.Println("9^5 =", d95)
  	fmt.Println("max possible: 6 * 9^5 =", 6*d95)
  	sum := 0
  	for i := 2; i < 6*d95; i++ {
  		if IsFifthPowerEqual(i) {
  			fmt.Println(i)
  			sum += i
  		}
  	}
  	fmt.Println("Sum:", sum)
  }

Result:

.. code-block:: txt

  9^5 = 59049
  max possible: 6 * 9^5 = 354294
  4150
  4151
  54748
  92727
  93084
  194979
  Sum: 443839

.. adsu:: 2

----

Test on:

- ``Ubuntu Linux 17.10``, ``Go 1.10.2``
- `Go Playground`_

References:

.. [1] `Digit fifth powers - Problem 30 - Project Euler <https://projecteuler.net/problem=30>`_

.. _Go Playground: https://play.golang.org/
