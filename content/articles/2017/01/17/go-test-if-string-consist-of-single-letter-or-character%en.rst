[Golang] Test if a String Consists of Single Letter or Character
################################################################

:date: 2017-01-17T05:28+08:00
:tags: Go, Golang, String Manipulation, Go strings Package
:category: Go
:summary: Test if a string consists of single character or letter in Go_
          programming language.
:adsu: yes

Test if a string consists of single character or letter in Golang_.

`Run code on Go Playground <https://play.golang.org/p/OcHJchHbID>`_

.. code-block:: go

  package main

  import (
  	"fmt"
  	"strings"
  )

  func IsConsistOfSingleCharacter(s, c string) bool {
  	return strings.TrimLeft(s, c) == ""
  }

  func main() {
  	fmt.Println(IsConsistOfSingleCharacter("#####", "#"))
  	fmt.Println(IsConsistOfSingleCharacter("aabaa", "a"))
  	fmt.Println(IsConsistOfSingleCharacter("哈哈哈哈", "哈"))
  	fmt.Println(IsConsistOfSingleCharacter("哈哈嘻哈哈", "哈"))
  }

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.7.4`` & Go Playground.

----

References:

.. [1] `func TrimLeft - strings - The Go Programming Language <https://golang.org/pkg/strings/#TrimLeft>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
