[Golang] Levenshtein Distance Recursive Implementation
######################################################

:date: 2017-04-04T04:43+08:00
:modified: 2017-04-05T23:41+08:00
:tags: Go, Golang, String Manipulation, Algorithm
:category: Go
:summary: Recursive implementation of Levenshtein distance algorithm in Go
          programming language. Levenshtein distance (LD) is a measure of the
          similarity between two strings, the minimum number of single-character
          edits (insertions, deletions or substitutions) required to change one
          string into the other.
:adsu: yes


I read post of the BK-tree referred by HN [1]_, and in the post, an interesting
algorithm called Levenshtein distance [2]_ is used. Levenshtein distance (LD) is
a measure of the similarity between two strings, the minimum number of
single-character edits (insertions, deletions or substitutions) required to
change one string into the other.

The recursive implementation of LD in Wikipedia is not very difficult, so I
port the code from C to Go as an exercise.

If you want a more efficient algorithm (`Wagner-Fischer algorithm`_), see [4]_.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/rOgRYC5QFu>`_
      :class: align-center

.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/util.go
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/recursive.go
.. adsu:: 3
.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/ld_test.go
.. adsu:: 4

----

Tested on:

- ``Ubuntu Linux 16.10``, ``Go 1.8``
- `Go Playground`_.

----

References:

.. [1] | `Interesting data structures: the BK-tree | Hacker News <https://news.ycombinator.com/item?id=14022424>`_
       | `Interesting data structures: the BK-tree (signal-to-noise.xyz) <http://signal-to-noise.xyz/post/bk-tree/>`_
.. [2] `Levenshtein distance - Wikipedia <https://en.wikipedia.org/wiki/Levenshtein_distance>`_
.. [3] `[Golang] Iterate Over UTF-8 Strings (non-ASCII strings) <{filename}../../../2016/02/03/go-iterate-over-utf8-non-ascii-string%en.rst>`_
.. [4] `[Golang] Wagner-Fischer algorithm for Edit Distance of Two Strings <{filename}../05/go-wagner-fischer-algorithm-edit-distance%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _Wagner-Fischer algorithm: https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm
