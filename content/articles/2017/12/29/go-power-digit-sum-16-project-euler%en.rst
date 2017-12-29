[Golang] Power digit sum - Problem 16 - Project Euler
#####################################################

:date: 2017-12-29T23:44+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler
:category: Go
:summary: Go solution to Power digit sum
          - Problem 13 - Project Euler.
:og_image: https://1.bp.blogspot.com/-ydH-Xt4GCcs/WeMutBBXP9I/AAAAAAAABnU/z6QSzNX_XUk3uin6NaTuhTsCkxV4-z0dgCLcBGAs/s1600/javamultiplex-solution-of-project-euler-problem-16-in-java.png
:adsu: yes

**Problem**: [1]_

  :math:`2^15` = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

  What is the sum of the digits of the number :math:`2^1000`?


**Solution**:

  We use large number addition [2]_ and the following formula:

  :math:`2^n = 2^{n-1} + 2^{n-1}`

  The result is:

  :math:`2^1000` = 10715086071862673209484250490600018105614048117055336074437503883703510511249361224931983788156958581275946729175531468251871452856923140435984577574698574803934567774824230985421074605062371141877954182153046474983581941267398767559165543946077062914571196477686542167660429831652624386837205668069376

  sum of digits = 1366

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/7-PMGhxiyYG>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  const MaxDigits = 350
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
  	result := MakePositiveInt("2")

  	for i := 1; i < 1000; i++ {
  		result = AddPositiveInt(result, result)
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

.. [1] `Power digit sum - Problem 16 - Project Euler <https://projecteuler.net/problem=16>`_
.. [2] `[Golang] Large Positive Integer Addition <{filename}../23/go-big-natural-number-addition%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _strconv.Atoi: https://golang.org/pkg/strconv/#Atoi
