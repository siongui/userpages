[Golang] Number letter counts - Problem 17 - Project Euler
##########################################################

:date: 2018-01-13T23:02+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler, Type Casting, Type Conversion
:category: Go
:summary: Go solution to Number letter counts
          - Problem 17 - Project Euler.
:og_image: http://www.mathx.net/image/read_write_numbers_with_words.jpg
:adsu: yes

**Problem**: [1]_

  If the numbers 1 to 5 are written out in words: one, two, three, four, five,
  then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

  If all the numbers from 1 to 1000 (one thousand) inclusive were written out in
  words, how many letters would be used?


  **NOTE:** Do not count spaces or hyphens. For example, 342 (three hundred and
  forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20
  letters. The use of "and" when writing out numbers is in compliance with
  British usage.


**Solution**:

  21124

  First we convert numbers to words [2]_, and then count the letters.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/tTifgbs9CPJ>`__
   :class: align-center

.. code-block:: go

  package main
  
  import (
  	"fmt"
  )
  
  var NumberToWord = map[int]string{
  	1:  "one",
  	2:  "two",
  	3:  "three",
  	4:  "four",
  	5:  "five",
  	6:  "six",
  	7:  "seven",
  	8:  "eight",
  	9:  "nine",
  	10: "ten",
  	11: "eleven",
  	12: "twelve",
  	13: "thirteen",
  	14: "fourteen",
  	15: "fifteen",
  	16: "sixteen",
  	17: "seventeen",
  	18: "eighteen",
  	19: "nineteen",
  	20: "twenty",
  	30: "thirty",
  	40: "forty",
  	50: "fifty",
  	60: "sixty",
  	70: "seventy",
  	80: "eighty",
  	90: "ninety",
  }
  
  func convert1to99(n int) (w string) {
  	if n < 20 {
  		w = NumberToWord[n]
  		return
  	}
  
  	r := n % 10
  	if r == 0 {
  		w = NumberToWord[n]
  	} else {
  		w = NumberToWord[n-r] + "-" + NumberToWord[r]
  	}
  	return
  }
  
  func convert100to999(n int) (w string) {
  	q := n / 100
  	r := n % 100
  	w = NumberToWord[q] + " " + "hundred"
  	if r == 0 {
  		return
  	} else {
  		w = w + " and " + convert1to99(r)
  	}
  	return
  }
  
  func Convert1to1000(n int) (w string) {
  	if n > 1000 || n < 1 {
  		panic("func Convert1to1000: n > 1000 or n < 1")
  	}
  
  	if n < 100 {
  		w = convert1to99(n)
  		return
  	}
  	if n == 1000 {
  		w = "one thousand"
  		return
  	}
  	w = convert100to999(n)
  	return
  }
  
  func CountLetter(s string) (n int) {
  	for _, r := range s {
  		if r == ' ' {
  			continue
  		}
  		if r == '-' {
  			continue
  		}
  		n++
  	}
  	return
  }
  
  func main() {
  	totalLetters := 0
  	for i := 1; i <= 1000; i++ {
  		totalLetters += CountLetter(Convert1to1000(i))
  	}
  	fmt.Println(totalLetters)
  }


.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `Number letter counts - Problem 17 - Project Euler <https://projecteuler.net/problem=17>`_
.. [2] `[Golang] Convert Numbers to Words From 1 to 1000 <{filename}go-convert-number-to-word-from-1-to-1000%en.rst>`_

.. _Go Playground: https://play.golang.org/
