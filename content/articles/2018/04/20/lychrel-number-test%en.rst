Lychrel Number Test
###################

:date: 2018-04-20T23:44+08:00
:tags: Go, Golang, Type Casting, Type Conversion, Algorithm, Math
:category: Algorithm
:summary: Test if a number (at least 2-digit) is Lychrel number or not under
          limited iterations in Go.
:og_image: http://datagenetics.com/blog/october12015/g.png
:adsu: yes


Test if a number (at least 2-digit) is `Lychrel number`_ or not under limited
iterations in Go.

I saw this question [1] from GeeksforGeeks. In short, If we reverse the digits
of a natural number (at least 2-digit) and add the reversed number back to the
original number, this is one iteration of the process. If we repeatedly do this,
and the resulting number cannot form a palindrome_, then this number is called
`Lychrel number`_.

The following is the implementation of Lychrel number test under limited
iterations in Go.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/Sw2LS2WYbHb>`__
   :class: align-center

.. code-block:: go

  package lychrel

  import (
  	"fmt"
  	"strconv"
  )

  func Reverse(s string) (rs string) {
  	for _, r := range s {
  		rs = string(r) + rs
  	}
  	return
  }

  func IsDecimalPalindromeNumber(n int64) bool {
  	if n < 0 {
  		n = -n
  	}

  	s := strconv.FormatInt(n, 10)
  	bound := (len(s) / 2) + 1
  	for i := 0; i < bound; i++ {
  		if s[i] != s[len(s)-1-i] {
  			return false
  		}
  	}
  	return true
  }

  func LychrelNumberTest(n int64, maxIteration int) bool {
  	s := strconv.FormatInt(n, 10)
  	if len(s) < 2 {
  		panic("input number must be at least 2 digits!")
  	}
  	if maxIteration < 0 {
  		panic("negative iteration number")
  	}

  	fmt.Println("Lychrel Number Test", n)
  	nextN := n
  	for i := 0; i < maxIteration; i++ {
  		// reverse nextN
  		nstr := strconv.FormatInt(nextN, 10)
  		rnstr := Reverse(nstr)
  		rn, _ := strconv.ParseInt(rnstr, 10, 64)
  		// add the reverse back to nextN
  		nextN = nextN + rn

  		fmt.Println("iteration", i+1)
  		fmt.Println(nextN)
  		if IsDecimalPalindromeNumber(nextN) {
  			return true
  		}
  	}
  	return false
  }


Example:

.. code-block:: go

  package lychrel

  import (
  	"testing"
  )

  func TestReverse(t *testing.T) {
  	if Reverse("你好嗎") != "嗎好你" {
  		t.Error("你好嗎")
  		return
  	}
  }

  func TestLychrelNumberTest(t *testing.T) {
  	LychrelNumberTest(56, 20)
  	LychrelNumberTest(87, 20)
  }

.. adsu:: 2

Tested on:

- ``Ubuntu Linux 17.10``, ``Go 1.10.1``
- `Go Playground`_

----

References:

.. [1] `Lychrel Number Implementation - GeeksforGeeks <https://www.geeksforgeeks.org/lychrel-number-implementation/>`_
.. [2] `[Golang] Check if Integer is Palindromic Number <{filename}/articles/2017/04/15/go-check-if-integer-number-is-palindrome%en.rst>`_
.. [3] `How to reverse a string in Go? - Stack Overflow <https://stackoverflow.com/a/4965535>`_
.. [4] `strconv - The Go Programming Language <https://golang.org/pkg/strconv/>`_

.. _Go Playground: https://play.golang.org/
.. _palindrome: https://www.google.com/search?q=palindrome+number
.. _Lychrel number: https://www.google.com/search?q=Lychrel+number
