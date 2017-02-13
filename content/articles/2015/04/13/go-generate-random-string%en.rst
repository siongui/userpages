[Golang] Generate Random String From [a-z0-9]
#############################################

:date: 2015-04-13 03:08
:modified: 2017-02-13T10:40+08:00
:tags: Go, Golang, String Manipulation, Random Number
:category: Go
:summary: Generate a random string from [a-z0-9] in Golang.
:adsu: yes

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

.. adsu:: 2

Souce Code
++++++++++

`Run code on Go Playground <https://play.golang.org/p/cIG85Za6LI>`__

(Note that time_ package in `Go Playground`_ always returns the same value no
matter what current time is, so the generated string will always be the same on
`Go Playground`_)

.. show_github_file:: siongui userpages content/code/go/random-string/a-z0-9.go
.. adsu:: 3

Appendix
========

Another way to generate string without make_ keyword:

`Run code on Go Playground <https://play.golang.org/p/v-67s0YJ4m>`__

.. show_github_file:: siongui userpages content/code/go/random-string/rdnstr.go
.. adsu:: 4

Test
++++

.. show_github_file:: siongui userpages content/code/go/random-string/a-z0-9_test.go


Tested on:

- ``Ubuntu Linux 14.10``, ``Go 1.4``.
- ``Ubuntu Linux 16.10``, ``Go 1.7.5``.

----

References:

.. [1] Google Search `go random number <https://www.google.com/search?q=go+random+number>`_

.. [2] `rand - The Go Programming Language <http://golang.org/pkg/math/rand/>`_

.. [3] Google Search `go random string <https://www.google.com/search?q=go+random+string>`_

.. [4] `go - How to generate a random string of a fixed length in golang? - Stack Overflow <http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang>`_

.. [5] `What is the fastest way to generate a long random string in Go? - Stack Overflow <http://stackoverflow.com/questions/12771930/what-is-the-fastest-way-to-generate-a-long-random-string-in-go>`_

.. [6] `time - The Go Programming Language <http://golang.org/pkg/time/>`_
.. adsu:: 5
.. [7] `[JavaScript] Generate Random String From [a-z0-9] <{filename}../../../2017/01/14/javascript-generate-random-string%en.rst>`_


.. _Go: https://golang.org/
.. _Seed: https://golang.org/pkg/math/rand/#Seed
.. _time: https://golang.org/pkg/time/
.. _Go Playground: https://play.golang.org/
.. _make: https://tour.golang.org/moretypes/13
