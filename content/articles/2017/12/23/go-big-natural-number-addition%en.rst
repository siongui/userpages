[Golang] Large Positive Integer Addition
########################################

:date: 2017-12-23T23:27+08:00
:modified: 2018-01-03T23:41+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler, Large Integer Arithmetic
:category: Go
:summary: Addition of big natural numbers in Go. This is for very large positive
          integers which overflows the built-in numerical type in Go.
:og_image: http://images.slideplayer.com/16/5210994/slides/slide_21.jpg
:adsu: yes

**Problem**: [1]_

  This comes from problem 13 in Project Euler. There are two large numbers:

    | 37107287533902102798797998220837590246510135740250
    | 46376937677490009712648124896970078050417018260538

  If you use strconv.Atoi_ to convert the string to int, you will get panic
  which reports overflow. So how do we add two large natural numbers in Go?

**Solution**:

  After some googling [2]_, I found the tutorial [3]_ which shows how to perform
  large integer addition. The following is Go implementation of the algorithm.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/NqZsv2eTGZQ>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  const MaxDigits = 100
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
  	a := MakePositiveInt(`37107287533902102798797998220837590246510135740250`)
  	b := MakePositiveInt(`46376937677490009712648124896970078050417018260538`)
  	c := AddPositiveInt(a, b)
  	fmt.Println(PositiveIntToString(a))
  	fmt.Println(PositiveIntToString(b))
  	fmt.Println(PositiveIntToString(c))
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `Large sum - Problem 13 - Project Euler <https://projecteuler.net/problem=13>`_
.. [2] | `big number arithmetic algorithm - Google search <https://www.google.com/search?q=big+number+arithmetic+algorithm>`_
       | `big number arithmetic algorithm - DuckDuckGo search <https://duckduckgo.com/?q=big+number+arithmetic+algorithm>`_
       | `big number arithmetic algorithm - Ecosia search <https://www.ecosia.org/search?q=big+number+arithmetic+algorithm>`_
       | `big number arithmetic algorithm - Qwant search <https://www.qwant.com/?q=big+number+arithmetic+algorithm>`_
       | `big number arithmetic algorithm - Bing search <https://www.bing.com/search?q=big+number+arithmetic+algorithm>`_
       | `big number arithmetic algorithm - Yahoo search <https://search.yahoo.com/search?p=big+number+arithmetic+algorithm>`_
       | `big number arithmetic algorithm - Baidu search <https://www.baidu.com/s?wd=big+number+arithmetic+algorithm>`_
       | `big number arithmetic algorithm - Yandex search <https://www.yandex.com/search/?text=big+number+arithmetic+algorithm>`_
.. [3] `Analysis of Algorithms: Lecture 20  <http://faculty.cse.tamu.edu/djimenez/ut/utsa/cs3343/lecture20.html>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _strconv.Atoi: https://golang.org/pkg/strconv/#Atoi
