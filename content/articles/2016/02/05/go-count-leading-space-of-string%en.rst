[Golang] Count Leading Spaces of a String
#########################################

:date: 2016-02-05T03:04+08:00
:tags: Go, Golang, String Manipulation
:category: Go
:summary: Count leading spaces of a string in Go_.

Count leading spaces of a string in Golang_.

`Run code on Go Playground <https://play.golang.org/p/DuxcXVrHFH>`_

.. code-block:: go

  package main

  import "fmt"

  func countLeadingSpace(line string) int {
          i := 0
          for _, runeValue := range line {
                  if runeValue == ' ' {
                          i++
                  } else {
                          break
                  }
          }
          return i
  }

  func main() {
          fmt.Println(countLeadingSpace("0 space"))
          fmt.Println(countLeadingSpace(" 1 space"))
          fmt.Println(countLeadingSpace("  2 space"))
          fmt.Println(countLeadingSpace("   3 space"))
  }

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

References:

.. [1] `golang leading space <https://www.google.com/search?q=golang+leading+space>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
