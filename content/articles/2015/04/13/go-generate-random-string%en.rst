[Golang] Generate Random String From [a-z0-9]
#############################################

:date: 2015-04-13 03:08
:tags: Go, Golang, String Manipulation
:category: Go
:summary: Generate a random string from [a-z0-9] in Golang.

Lesson learned from this exercise
+++++++++++++++++++++++++++++++++

1. Give seed value to the Seed_ func in `math/rand` package. Otherwise the same
   string is returned on every time.

2. Allocate the byte array with make_ function.

   Correct:

   .. code-block:: go

     result := make([]byte, strlen)

   Wrong:

   .. code-block:: go

     var result [strlen]byte

3. Covert the `[]byte` to string by built-in *string* keyword.

Souce Code
++++++++++

`Run code on Go Playground <https://play.golang.org/p/cIG85Za6LI>`_

.. show_github_file:: siongui userpages content/code/go-random-string/a-z0-9.go

Test
++++

.. show_github_file:: siongui userpages content/code/go-random-string/a-z0-9_test.go


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

References:

.. [1] Google Search `go random number <https://www.google.com/search?q=go+random+number>`_

.. [2] `rand - The Go Programming Language <http://golang.org/pkg/math/rand/>`_

.. [3] Google Search `go random string <https://www.google.com/search?q=go+random+string>`_

.. [4] `go - How to generate a random string of a fixed length in golang? - Stack Overflow <http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang>`_

.. [5] `What is the fastest way to generate a long random string in Go? - Stack Overflow <http://stackoverflow.com/questions/12771930/what-is-the-fastest-way-to-generate-a-long-random-string-in-go>`_

.. [6] `time - The Go Programming Language <http://golang.org/pkg/time/>`_



.. _Go: https://golang.org/

.. _Seed: http://golang.org/pkg/math/rand/#Seed

.. _make: http://tour.golang.org/moretypes/9
