[Golang] Counting Sundays - Problem 19 - Project Euler
######################################################

:date: 2018-10-22T02:21+08:00
:tags: Go, Golang, Algorithm, Math, Project Euler
:category: Go
:summary: Go solution to Counting Sundays
          - Problem 19 - Project Euler.
:og_image: https://3.bp.blogspot.com/-CJW2HyY0DsE/VxOKnspVDcI/AAAAAAAABvs/dnc1s9Z2lbQHpW_uOl_62YkVsUBwfMWgQCLcB/s1600/project%2Beuler%2Bproblem%2B19%2Bwith%2Bsolution.png
:adsu: yes

**Problem**: [1]_

.. raw:: html

  <p>You are given the following information, but you may prefer to do some research for yourself.</p>
  <ul><li>1 Jan 1900 was a Monday.</li>
  <li>Thirty days has September,<br />
  April, June and November.<br />
  All the rest have thirty-one,<br />
  Saving February alone,<br />
  Which has twenty-eight, rain or shine.<br />
  And on leap years, twenty-nine.</li>
  <li>A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.</li>
  </ul><p>How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?</p>

**Solution**:

  171

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/jWhYZ8bZTqj>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func monthDays(year, month int) int {
  	switch month {
  	case 1, 3, 5, 7, 8, 10, 12:
  		return 31
  	case 4, 6, 9, 11:
  		return 30
  	case 2:
  		if year%400 == 0 {
  			return 29
  		}
  		if year%100 == 0 {
  			return 28
  		}
  		if year%4 == 0 {
  			return 29
  		}
  		return 28
  	}
  	panic("not correct month")
  }

  func main() {
  	year := 1900
  	month := 1
  	day := 1
  	weekDay := 1 // 1: Monday ... 7: Sunday

  	count := 0
  	for {
  		day += 1
  		if day > monthDays(year, month) {
  			day = 1
  			month += 1
  		}
  		if month > 12 {
  			month = 1
  			year += 1
  		}

  		weekDay += 1
  		if weekDay > 7 {
  			weekDay = 1
  		}

  		// first day of the month is Sunday
  		if day == 1 && weekDay == 7 && year > 1900 {
  			//fmt.Printf("%d %d 1st is Sunday\n", year, month)
  			count += 1
  		}

  		// termination condition
  		if year == 2000 && month == 12 && day == 31 {
  			break
  		}
  	}

  	fmt.Printf("total %d Sunday", count)
  }

.. adsu:: 2

----

Test on:

- `Go Playground`_

References:

.. [1] `Counting Sundays - Problem 19 - Project Euler <https://projecteuler.net/problem=19>`_

.. _Go Playground: https://play.golang.org/
