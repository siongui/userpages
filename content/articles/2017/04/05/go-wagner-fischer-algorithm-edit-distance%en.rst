[Golang] Wagner-Fischer algorithm for Edit Distance of Two Strings
##################################################################

:date: 2017-04-05T23:34+08:00
:tags: Go, Golang, String Manipulation, Algorithm
:category: Go
:summary: Go implementation of Wagner-Fischer algorithm that computes the edit
          distance between two strings of characters. The edit distance here is
          Levenshtein distance, the minimum number of single-character edits
          (insertions, deletions or substitutions) required to change one string
          into the other.
:adsu: yes


Go_ implementation of `Wagner-Fischer algorithm`_ that computes the edit
distance between two strings of characters. The edit distance here is
`Levenshtein distance`_, the minimum number of single-character edits
(insertions, deletions or substitutions) required to change one string into the
other.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/Z1mequP8hl>`_
      :class: align-center

The following code is utility for converting a string to rune slice, and minimum
value of three integers.

.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/util.go
.. adsu:: 2

The `Wagner-Fischer algorithm`_:

.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/wf.go
.. adsu:: 3

Testing:

.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/wf_test.go
.. adsu:: 4

For recursive implementation of Levenshtein distance, see [1]_.

----

Tested on:

- ``Ubuntu Linux 16.10``, ``Go 1.8``
- `Go Playground`_.

----

References:

.. [1] `[Golang] Levenshtein Distance Recursive Implementation <{filename}../04/go-levenshtein-distance-recursive-implementation%en.rst>`_
.. [2] `[Golang] Initialize Two Dimensional Array/Slice <{filename}../../02/01/go-initialize-two-dimensional-array-or-slice%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Wagner-Fischer algorithm: https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm
.. _Levenshtein distance: https://en.wikipedia.org/wiki/Levenshtein_distance
.. _Go Playground: https://play.golang.org/
