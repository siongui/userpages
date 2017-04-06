[Golang] BK-tree Data Structure Implementation
##############################################

:date: 2017-04-06T23:34+08:00
:tags: Go, Golang, String Manipulation, Algorithm, Data Structure
:category: Go
:summary: Go implementation of BK-tree data structure used for approximate
          string matching in a dictionary. The distance metric here is
          Levenshtein distance, the minimum number of single-character edits
          (insertions, deletions or substitutions) required to change one string
          into the other.
:adsu: yes


Go implementation of `BK-tree`_ data structure used for approximate string
matching in a dictionary. The distance metric here is `Levenshtein distance`_,
the minimum number of single-character edits (insertions, deletions or
substitutions) required to change one string into the other.

The Go implementation is based on the post [1]_. If you have no idea what
BK-tree is, read the post first.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/in71uLHiXv>`_
      :class: align-center

**BK-tree**: Implementation of Levenshtein distance is put in
the *Levenshtein distance* section below.

.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/bktree.go
.. adsu:: 2

Example of using Bk-tree:

.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/bktree_test.go
.. adsu:: 3

Levenshtein distance
++++++++++++++++++++

The `Wagner-Fischer algorithm`_: See [2]_ for more details.

.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/util.go
.. adsu:: 4
.. show_github_file:: siongui userpages content/code/go/levenshtein-distance/wf.go
.. adsu:: 5

----

Tested on:

- ``Ubuntu Linux 16.10``, ``Go 1.8``
- `Go Playground`_.

----

References:

.. [1] | `Interesting data structures: the BK-tree | Hacker News <https://news.ycombinator.com/item?id=14022424>`_
       | `Interesting data structures: the BK-tree (signal-to-noise.xyz) <http://signal-to-noise.xyz/post/bk-tree/>`_
.. [2] `[Golang] Wagner-Fischer algorithm for Edit Distance of Two Strings <{filename}../05/go-wagner-fischer-algorithm-edit-distance%en.rst>`_
.. [3] | `Go maps in action - The Go Blog <https://blog.golang.org/go-maps-in-action>`_
       | `Appending to a slice - A Tour of Go <https://tour.golang.org/moretypes/15>`_
       | `json - The Go Programming Language <https://golang.org/pkg/encoding/json/#example_Marshal>`_
       | `[Golang] Pretty Print Variable (struct, map, array, slice) <{filename}../../../2016/01/30/go-pretty-print-variable%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _BK-tree: https://www.google.com/search?q=BK-tree
.. _Wagner-Fischer algorithm: https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm
.. _Levenshtein distance: https://en.wikipedia.org/wiki/Levenshtein_distance
.. _Go Playground: https://play.golang.org/
