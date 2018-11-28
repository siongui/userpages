[Golang] Remove Leading and Trailing Empty Strings in String Array
##################################################################

:date: 2017-01-19T10:12+08:00
:tags: Go, Golang, String Manipulation
:category: Go
:summary: Remove leading and trailing empty strings in string array_/slice_ via
          Go_ programming language.
:adsu: yes

Remove leading and trailing empty strings in string array_/slice_ via Golang_.
This is similar to removing leading and trailing whitespaces in a string, but in
two dimension.

`Run code on Go Playground <https://play.golang.org/p/fcY7KGQCIo>`_

.. code-block:: go

  // Remove leading and trailing empty strings in string array
  package main

  import (
  	"fmt"
  )

  func RemoveLeadingEmptyStringsInStringArray(sa []string) []string {
  	firstNonEmptyStringIndex := 0
  	for i := 0; i < len(sa); i++ {
  		if sa[i] == "" {
  			firstNonEmptyStringIndex++
  		} else {
  			break
  		}
  	}
  	return sa[firstNonEmptyStringIndex:len(sa)]
  }

  func RemoveTrailingEmptyStringsInStringArray(sa []string) []string {
  	lastNonEmptyStringIndex := len(sa) - 1
  	for i := lastNonEmptyStringIndex; i >= 0; i-- {
  		if sa[i] == "" {
  			lastNonEmptyStringIndex--
  		} else {
  			break
  		}
  	}
  	return sa[0 : lastNonEmptyStringIndex+1]
  }

  func main() {
  	testStringArray := []string{"", "", "aa", "bb", "cc", "", "", ""}

  	trailingTrimed := RemoveTrailingEmptyStringsInStringArray(testStringArray)
  	for _, str := range trailingTrimed {
  		if str == "" {
  			fmt.Println("nil")
  		} else {
  			fmt.Println(str)
  		}
  	}

  	fmt.Println("----")
  	leadingTrailingTrimed := RemoveLeadingEmptyStringsInStringArray(trailingTrimed)
  	for _, str := range leadingTrailingTrimed {
  		if str == "" {
  			fmt.Println("nil")
  		} else {
  			fmt.Println(str)
  		}
  	}
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.7.4`` & `Go Playground`_.

----

References:

.. [1] `Go Slices: usage and internals - The Go Blog <https://blog.golang.org/go-slices-usage-and-internals>`_
.. [2] `Zero values - A Tour of Go <https://tour.golang.org/basics/12>`_
.. [3] `Avoiding partially constructed variables : golang <https://old.reddit.com/r/golang/comments/a0ueef/avoiding_partially_constructed_variables/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _slice: https://www.google.com/search?q=golang+slice
.. _array: https://www.google.com/search?q=golang+array
