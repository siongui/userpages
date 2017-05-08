[Golang] Anagram Check by Characters Count
##########################################

:date: 2017-05-08T16:43+08:00
:tags: Go, Golang, Algorithm
:category: Go
:summary: Check whether two strings are anagram of each other by characters
          count in Go programming language.
:og_image: https://qph.ec.quoracdn.net/main-qimg-f8bf4e32835307ad93313b0edd94adda
:adsu: yes


Check whether two strings are anagram_ of each other in Go programming language.

Check by characters count:

1. setup ``map`` structures to count the characters of both strings [2]_.
2. compare if the two map structures are equal [3]_ via reflect.DeepEqual_.

Another way to check anagram is by sorting, see [1]_.

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/1HotD5pN-6>`__
   :class: align-center

.. code-block:: go

  package anagram

  import (
  	"reflect"
  )

  func GetCharCount(s string) (c map[rune]int) {
  	c = make(map[rune]int)
  	for _, runeValue := range s {
  		if _, ok := c[runeValue]; ok {
  			c[runeValue] += 1
  		} else {
  			c[runeValue] = 1
  		}
  	}
  	return
  }

  func AreAnagram(s1, s2 string) bool {
  	c1 := GetCharCount(s1)
  	c2 := GetCharCount(s2)

  	return reflect.DeepEqual(c1, c2)
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

.. [1] `[Golang] Check Whether Two Strings Are Anagram of Each Other <{filename}../06/go-check-if-two-string-are-anagram%en.rst>`_
.. [2] `Go maps in action - The Go Blog <https://blog.golang.org/go-maps-in-action>`_
.. [3] `go - How to compare struct, slice, map are equal? - Stack Overflow <http://stackoverflow.com/questions/24534072/how-to-compare-struct-slice-map-are-equal>`_
.. [4] | `go map equal - Google search <https://www.google.com/search?q=go+map+equal>`_
       | `go map equal - DuckDuckGo search <https://duckduckgo.com/?q=go+map+equal>`_
       | `go map equal - Ecosia search <https://www.ecosia.org/search?q=go+map+equal>`_
       | `go map equal - Qwant search <https://www.qwant.com/?q=go+map+equal>`_
       | `go map equal - Bing search <https://www.bing.com/search?q=go+map+equal>`_
       | `go map equal - Yahoo search <https://search.yahoo.com/search?p=go+map+equal>`_
       | `go map equal - Baidu search <https://www.baidu.com/s?wd=go+map+equal>`_
       | `go map equal - Yandex search <https://www.yandex.com/search/?text=go+map+equal>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Go Playground: https://play.golang.org/
.. _anagram: https://www.google.com/search?q=anagram
.. _reflect.DeepEqual: https://golang.org/pkg/reflect/#DeepEqual
