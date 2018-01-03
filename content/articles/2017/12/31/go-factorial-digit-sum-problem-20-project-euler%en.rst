[Golang] Factorial digit sum - Problem 20 - Project Euler
#########################################################

:date: 2018-01-02T22:30+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler, Large Integer Arithmetic
:category: Go
:summary: Go solution to Factorial digit sum
          - Problem 20 - Project Euler.
:og_image: https://benvitalenum3ers.files.wordpress.com/2014/02/factorials-1.png
:adsu: yes

**Problem**: [1]_

  n! means n × (n − 1) × ... × 3 × 2 × 1

  For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
  and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

  Find the sum of the digits in the number 100!


**Solution**:

  We use large number multiplication [2]_ to calcualte 100!

  100! = 93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000

  sum of digits = 648


.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/Pn14mMMbQXF>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"strconv"
  )

  const MaxDigits = 200
  const BASE = 10

  func MakePositiveInt(s string) (n [MaxDigits]int) {
  	// make n zero
  	for i := 0; i < MaxDigits; i++ {
  		n[i] = 0
  	}

  	for index, digit := range s {
  		i := len(s) - index - 1
  		switch digit {
  		case '0':
  			n[i] = 0
  		case '1':
  			n[i] = 1
  		case '2':
  			n[i] = 2
  		case '3':
  			n[i] = 3
  		case '4':
  			n[i] = 4
  		case '5':
  			n[i] = 5
  		case '6':
  			n[i] = 6
  		case '7':
  			n[i] = 7
  		case '8':
  			n[i] = 8
  		case '9':
  			n[i] = 9
  		default:
  			panic("invalid digit in number string")
  		}
  	}

  	return
  }

  func AddPositiveInt(a, b [MaxDigits]int) (c [MaxDigits]int) {
  	var carry, sum = 0, 0

  	// make c zero
  	for i := 0; i < MaxDigits; i++ {
  		c[i] = 0
  	}

  	for i := 0; i < MaxDigits; i++ {
  		sum = a[i] + b[i] + carry

  		if sum >= BASE {
  			carry = 1
  			sum -= BASE
  		} else {
  			carry = 0
  		}

  		c[i] = sum
  	}

  	if carry != 0 {
  		panic("overflow in addition")
  	}

  	return
  }

  func MultiplyOneDigit(a [MaxDigits]int, n int) (b [MaxDigits]int) {
  	var carry = 0

  	// make b zero
  	for i := 0; i < MaxDigits; i++ {
  		b[i] = 0
  	}

  	for i := 0; i < MaxDigits; i++ {
  		b[i] = n * a[i]

  		b[i] += carry

  		if b[i] >= BASE {
  			carry = b[i] / BASE
  			b[i] %= BASE
  		} else {
  			carry = 0
  		}
  	}

  	if carry != 0 {
  		panic("overflow in multiplication")
  	}

  	return
  }

  func ShiftLeft(a [MaxDigits]int, n int) [MaxDigits]int {
  	var i int

  	for i = MaxDigits - 1; i >= n; i-- {
  		a[i] = a[i-n]
  	}
  	for i >= 0 {
  		a[i] = 0
  		i -= 1
  	}

  	return a
  }

  func MultiplyPositiveInt(a, b [MaxDigits]int) (c [MaxDigits]int) {
  	// make c zero
  	for i := 0; i < MaxDigits; i++ {
  		c[i] = 0
  	}

  	for i := 0; i < MaxDigits; i++ {
  		tmp := MultiplyOneDigit(b, a[i])
  		tmp = ShiftLeft(tmp, i)
  		c = AddPositiveInt(c, tmp)
  	}

  	return
  }

  func PrintPositiveInt(a [MaxDigits]int) {
  	isLeadingZero := true
  	for i := MaxDigits - 1; i >= 0; i-- {
  		if isLeadingZero && a[i] == 0 {
  			continue
  		} else {
  			isLeadingZero = false
  			fmt.Print(a[i])
  		}
  	}
  	fmt.Println("\n")
  }

  func main() {
  	result := MakePositiveInt("1")
  	for i := 2; i <= 100; i++ {
  		s := strconv.Itoa(i)
  		tmp := MakePositiveInt(s)
  		result = MultiplyPositiveInt(result, tmp)
  	}
  	PrintPositiveInt(result)

  	sum := 0
  	for i := 0; i < MaxDigits; i++ {
  		sum += result[i]
  	}
  	fmt.Println(sum)
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `Factorial digit sum - Problem 20 - Project Euler <https://projecteuler.net/problem=20>`_
.. [2] `[Golang] Large Positive Integer Multiplication <{filename}go-big-natural-number-multiplication%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
