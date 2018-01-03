[Golang] 1000-digit Fibonacci number - Problem 25 - Project Euler
#################################################################

:date: 2018-01-03T23:38+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler, Large Integer Arithmetic
:category: Go
:summary: Go solution to 1000-digit Fibonacci number
          - Problem 25 - Project Euler.
:og_image: http://thewessens.net/gnomon/ProjectEuler/ipad/problem25.gif
:adsu: yes

**Problem**: [1]_

  The Fibonacci sequence is defined by the recurrence relation:

    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.

  Hence the first 12 terms will be:

  | F1 = 1
  | F2 = 1
  | F3 = 2
  | F4 = 3
  | F5 = 5
  | F6 = 8
  | F7 = 13
  | F8 = 21
  | F9 = 34
  | F10 = 55
  | F11 = 89
  | F12 = 144

  The 12th term, F12, is the first term to contain three digits.

  What is the index of the first term in the Fibonacci sequence to contain 1000
  digits?


**Solution**:

  We use large number addition [2]_ to calcualte Fibonacci number. The answer
  is 4782.


.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/PmfGBlWx5FN>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  const MaxDigits = 1200
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

  func PositiveIntToString(a [MaxDigits]int) (result string) {
  	isLeadingZero := true
  	for i := MaxDigits - 1; i >= 0; i-- {
  		if isLeadingZero && a[i] == 0 {
  			continue
  		} else {
  			isLeadingZero = false
  			switch a[i] {
  			case 0:
  				result += "0"
  			case 1:
  				result += "1"
  			case 2:
  				result += "2"
  			case 3:
  				result += "3"
  			case 4:
  				result += "4"
  			case 5:
  				result += "5"
  			case 6:
  				result += "6"
  			case 7:
  				result += "7"
  			case 8:
  				result += "8"
  			case 9:
  				result += "9"
  			default:
  				panic("invalid digit in int array")
  			}
  		}
  	}
  	return
  }

  func main() {
  	Fn1 := MakePositiveInt(`1`)
  	Fn2 := MakePositiveInt(`1`)
  	Fn := AddPositiveInt(Fn1, Fn2)
  	index := 3

  	for len(PositiveIntToString(Fn)) < 1000 {
  		Fn2 = Fn1
  		Fn1 = Fn
  		Fn = AddPositiveInt(Fn1, Fn2)
  		index++
  	}
  	fmt.Println(PositiveIntToString(Fn))
  	fmt.Println("index: ", index)
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `1000-digit Fibonacci number - Problem 25 - Project Euler <https://projecteuler.net/problem=25>`_
.. [2] `[Golang] Large Positive Integer Addition <{filename}../../../2017/12/23/go-big-natural-number-addition%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
