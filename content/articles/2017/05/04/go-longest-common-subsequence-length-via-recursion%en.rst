[Golang] Longest Common Subsequence Length via Recursion
########################################################

:date: 2017-05-04T03:39+08:00
:tags: Go, Golang, Algorithm, Recursion
:category: Go
:summary: Compute the length of longest common subsequence via recursion in Go.
:og_image: http://1.bp.blogspot.com/-4_oNWDrXXHk/VH3xPC1tCII/AAAAAAAAAFk/t4L8x252cQc/s1600/lcs.png
:adsu: yes

Compute the length of `longest common subsequence`_ via recursion_ in Go_.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/ettQoUkOAU>`_
      :class: align-center

**LCS Length via Recursion** in Go:

.. code-block:: go

  package lcs

  func max(a, b int) int {
  	if a > b {
  		return a
  	}
  	return b
  }

  func LCSLength(x, y []byte) int {
  	if len(x) == 0 || len(y) == 0 {
  		return 0
  	}

  	if x[len(x)-1] == y[len(y)-1] {
  		return 1 + LCSLength(x[0:len(x)-1], y[0:len(y)-1])
  	} else {
  		return max(LCSLength(x, y[0:len(y)-1]), LCSLength(x[0:len(x)-1], y))
  	}
  }

*Testing*:

.. code-block:: go

  package lcs

  import (
  	"testing"
  )

  func TestLCSLength(t *testing.T) {
  	if LCSLength([]byte(`ABCDGH`), []byte(`AEDFHR`)) != 3 {
  		t.Error("should be 3")
  	}
  	if LCSLength([]byte(`AGGTAB`), []byte(`GXTXAYB`)) != 4 {
  		t.Error("should be 4")
  	}
  }

.. adsu:: 2

----

Tested on:

- ``Ubuntu Linux 17.04``, ``Go 1.8.1``
- `Go Playground`_.

----

References:

.. [1] `Dynamic Programming | Set 4 (Longest Common Subsequence) - GeeksforGeeks <http://www.geeksforgeeks.org/dynamic-programming-set-4-longest-common-subsequence/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _longest common subsequence: https://www.google.com/search?q=longest+common+subsequence
.. _recursion: https://en.wikipedia.org/wiki/Recursion
