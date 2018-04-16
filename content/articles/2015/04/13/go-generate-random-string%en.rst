[Golang] Generate Random String From [a-z0-9]
#############################################

:date: 2015-04-13 03:08
:modified: 2017-03-21T01:52+08:00
:tags: Go, Golang, String Manipulation, Random Number, Go time Package
:category: Go
:summary: Generate a random string from [a-z0-9] in Golang.
:og_image: https://www.webucator.com/blog/wp-content/uploads/2015/10/random_string.png
:adsu: yes

.. contents:: Table of Contents

Lesson learned from this exercise
+++++++++++++++++++++++++++++++++

1. Give seed value to the Seed_ func in `math/rand` package. Otherwise the same
   string is returned on every time.

2. Accoding to the comment_ from `@dchapes`_, the pseudorandom number generator
   should not be re-seeded on each use, nor re-seed the global PRNG in a
   package. `@shurcooL`_ shows how to `properly seed the PRNG`_:

     If you're using global ``rand.Rand``, don't re-seed it. But if you want to
     set your own seed, create your own local instance of ``rand.Rand``.

     .. code-block:: go

       	var r *rand.Rand // Rand for this package.

       	func init() {
       		r = rand.New(rand.NewSource(time.Now().UnixNano()))
       	}

       	func RandomString(strlen int) string {
       		// use existing r, no need to re-seed it
       		r.Intn(len(chars))
       	}

     You can also `simplify this into one line`_:

     .. code-block:: go

       	// r is a source of random numbers used in this package.
       	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

3. Allocate the byte array with make_ function.

   Correct:

   .. code-block:: go

     result := make([]byte, strlen)

   Wrong:

   .. code-block:: go

     var result [strlen]byte

4. Covert the `[]byte` to string by built-in *string* keyword.

.. adsu:: 2

Souce Code
++++++++++

`Run code on Go Playground <https://play.golang.org/p/TzaeVPOYxd>`__

(Note that time_ package in `Go Playground`_ always returns the same value no
matter what current time is, so the generated string will always be the same on
`Go Playground`_)

.. show_github_file:: siongui userpages content/code/go/random-string/a-z0-9.go
.. adsu:: 3

Appendix
========

Another way to generate string without make_ keyword:

`Run code on Go Playground <https://play.golang.org/p/02EZE7kQ9z>`__

.. show_github_file:: siongui userpages content/code/go/random-string/rdnstr.go
.. adsu:: 4

Usage [10]_
+++++++++++

.. show_github_file:: siongui userpages content/code/go/random-string/a-z0-9_test.go

----

Tested on:

- ``Ubuntu Linux 14.10``, ``Go 1.4``.
- ``Ubuntu Linux 16.10``, ``Go 1.7.5``.
- ``Ubuntu Linux 16.10``, ``Go 1.8``.

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

.. [8] | `Do not reseed the global PRNG in a package - Google search <https://www.google.com/search?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - DuckDuckGo search <https://duckduckgo.com/?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Ecosia search <https://www.ecosia.org/search?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Qwant search <https://www.qwant.com/?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Bing search <https://www.bing.com/search?q=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Yahoo search <https://search.yahoo.com/search?p=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Baidu search <https://www.baidu.com/s?wd=Do+not+reseed+the+global+PRNG+in+a+package>`_
       | `Do not reseed the global PRNG in a package - Yandex search <https://www.yandex.com/search/?text=Do+not+reseed+the+global+PRNG+in+a+package>`_

.. [9] `go - Golang random number generator how to seed properly - Stack Overflow <http://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly>`_
.. [10] `How to write usage code in testing <https://github.com/siongui/userpages/commit/77cd55346752ccaa2efa44b9084e97af81b664dd#commitcomment-21401415>`_

.. _Go: https://golang.org/
.. _Seed: https://golang.org/pkg/math/rand/#Seed
.. _time: https://golang.org/pkg/time/
.. _Go Playground: https://play.golang.org/
.. _make: https://tour.golang.org/moretypes/13
.. _@dchapes: https://github.com/dchapes
.. _comment: https://github.com/siongui/userpages/commit/77cd55346752ccaa2efa44b9084e97af81b664dd#commitcomment-21400225
.. _@shurcooL: https://github.com/shurcooL
.. _properly seed the PRNG: https://github.com/siongui/userpages/commit/77cd55346752ccaa2efa44b9084e97af81b664dd#commitcomment-21401369
.. _simplify this into one line: https://github.com/siongui/userpages/commit/6dc58d0b28a615ae19c43900358bcd258c1faac6#commitcomment-21402539
