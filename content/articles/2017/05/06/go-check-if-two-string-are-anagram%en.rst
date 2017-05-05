[Golang] Check Whether Two Strings Are Anagram of Each Other
############################################################

:date: 2017-05-06T01:50+08:00
:tags: Go, Golang, Type Casting, Type Conversion, Algorithm
:category: Go
:summary: Check whether two strings are anagram of each other in Go programming
          language.
:og_image: https://qph.ec.quoracdn.net/main-qimg-f8bf4e32835307ad93313b0edd94adda
:adsu: yes


Check whether two strings are anagram_ of each other in Go programming language.

Check by Sorting [1]_:

1. sort both strings
2. compare the sorted strings

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/6pz7TuIU4g>`__
   :class: align-center

.. code-block:: go

  package anagram

  import (
  	"sort"
  )

  type ByRune []rune

  func (r ByRune) Len() int           { return len(r) }
  func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
  func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

  func StringToRuneSlice(s string) []rune {
  	var r []rune
  	for _, runeValue := range s {
  		r = append(r, runeValue)
  	}
  	return r
  }

  func AreAnagram(s1, s2 string) bool {
  	var r1 ByRune = StringToRuneSlice(s1)
  	var r2 ByRune = StringToRuneSlice(s2)

  	sort.Sort(r1)
  	sort.Sort(r2)

  	return string(r1) == string(r2)
  }

.. adsu:: 2

Testing:

.. code-block:: go

  package anagram

  import (
  	"testing"
  )

  func TestAreAnagram(t *testing.T) {
  	if AreAnagram("listen", "silent") != true {
  		t.Error(`"listen", "silent"`)
  	}
  	if AreAnagram("test", "ttew") != false {
  		t.Error(`"test", "ttew"`)
  	}
  	if AreAnagram("geeksforgeeks", "forgeeksgeeks") != true {
  		t.Error(`"geeksforgeeks", "forgeeksgeeks"`)
  	}
  	if AreAnagram("triangle", "integral") != true {
  		t.Error(`"triangle", "integral"`)
  	}
  	if AreAnagram("abd", "acb") != false {
  		t.Error(`"abd", "acb"`)
  	}
  }

.. adsu:: 3

Tested on:

- ``Ubuntu Linux 17.04``, ``Go 1.8.1``
- `Go Playground`_

----

References:

.. [1] `Check whether two strings are anagram of each other - GeeksforGeeks <http://www.geeksforgeeks.org/check-whether-two-strings-are-anagram-of-each-other/>`_
.. [2] `sort - The Go Programming Language <https://golang.org/pkg/sort/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _anagram: https://www.google.com/search?q=anagram
