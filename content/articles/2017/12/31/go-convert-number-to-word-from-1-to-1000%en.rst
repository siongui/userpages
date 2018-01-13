[Golang] Convert Numbers to Words From 1 to 1000
################################################

:date: 2018-01-10T03:31+08:00
:modified: 2018-01-13T22:03+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler, Type Casting, Type Conversion
:category: Go
:summary: Spell numbers from 1 to 1000 in English via Go programming language.
:og_image: http://www.mathx.net/image/read_write_numbers_with_words.jpg
:adsu: yes

**Problem**: [1]_

  Write out all the numbers from 1 to 1000 (one thousand) inclusive in words.
  For example, 342 (three hundred and forty-two) contains 23 letters and 115
  (one hundred and fifteen) contains 20 letters. The use of "and" when writing
  out numbers is in compliance with British usage.


**Solution**:

  See the following code:


.. .. rubric:: `Run Code on Go Playground <https://play.golang.org/p/yODRXUh_GmE>`__
.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/9JxAGWvigrq>`__
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
  
  func main() {
  	for i := 1; i <= 1000; i++ {
  		fmt.Println(Convert1to1000(i))
  	}
  }

.. adsu:: 2

Tested on: `Go Playground`_

----

References:

.. [1] `Number letter counts - Problem 17 - Project Euler <https://projecteuler.net/problem=17>`_
.. [2] | `number to words - Google search <https://www.google.com/search?q=number+to+words>`_
       | `number to words - DuckDuckGo search <https://duckduckgo.com/?q=number+to+words>`_
       | `number to words - Ecosia search <https://www.ecosia.org/search?q=number+to+words>`_
       | `number to words - Qwant search <https://www.qwant.com/?q=number+to+words>`_
       | `number to words - Bing search <https://www.bing.com/search?q=number+to+words>`_
       | `number to words - Yahoo search <https://search.yahoo.com/search?p=number+to+words>`_
       | `number to words - Baidu search <https://www.baidu.com/s?wd=number+to+words>`_
       | `number to words - Yandex search <https://www.yandex.com/search/?text=number+to+words>`_
.. [3] | `how to spell numbers 1-1000 in english - Google search <https://www.google.com/search?q=how+to+spell+numbers+1-1000+in+english>`_
       | `how to spell numbers 1-1000 in english - DuckDuckGo search <https://duckduckgo.com/?q=how+to+spell+numbers+1-1000+in+english>`_
       | `how to spell numbers 1-1000 in english - Ecosia search <https://www.ecosia.org/search?q=how+to+spell+numbers+1-1000+in+english>`_
       | `how to spell numbers 1-1000 in english - Qwant search <https://www.qwant.com/?q=how+to+spell+numbers+1-1000+in+english>`_
       | `how to spell numbers 1-1000 in english - Bing search <https://www.bing.com/search?q=how+to+spell+numbers+1-1000+in+english>`_
       | `how to spell numbers 1-1000 in english - Yahoo search <https://search.yahoo.com/search?p=how+to+spell+numbers+1-1000+in+english>`_
       | `how to spell numbers 1-1000 in english - Baidu search <https://www.baidu.com/s?wd=how+to+spell+numbers+1-1000+in+english>`_
       | `how to spell numbers 1-1000 in english - Yandex search <https://www.yandex.com/search/?text=how+to+spell+numbers+1-1000+in+english>`_
.. [4] | `how to write numbers in words in english - Google search <https://www.google.com/search?q=how+to+write+numbers+in+words+in+english>`_
       | `how to write numbers in words in english - DuckDuckGo search <https://duckduckgo.com/?q=how+to+write+numbers+in+words+in+english>`_
       | `how to write numbers in words in english - Ecosia search <https://www.ecosia.org/search?q=how+to+write+numbers+in+words+in+english>`_
       | `how to write numbers in words in english - Qwant search <https://www.qwant.com/?q=how+to+write+numbers+in+words+in+english>`_
       | `how to write numbers in words in english - Bing search <https://www.bing.com/search?q=how+to+write+numbers+in+words+in+english>`_
       | `how to write numbers in words in english - Yahoo search <https://search.yahoo.com/search?p=how+to+write+numbers+in+words+in+english>`_
       | `how to write numbers in words in english - Baidu search <https://www.baidu.com/s?wd=how+to+write+numbers+in+words+in+english>`_
       | `how to write numbers in words in english - Yandex search <https://www.yandex.com/search/?text=how+to+write+numbers+in+words+in+english>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
