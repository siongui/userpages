[Golang] Succinct Trie Implementation
#####################################

:date: 2016-02-08T20:27+08:00
:tags: Go, Golang, Data Structure
:category: Go
:summary: Go_ implemantation of `succinct trie`_ ported from `Bits.js`_.


`Bits.js`_ ([1]_) implements a `succinct trie`_ for fast lookup of words in a
large dictionary and does not take much space at the same time. I have an
application which has about 200,000+ words to be encoded. It takes a lot of time
(over 30 minutes) to finish the job, so I decided to re-implement the succinct
trie in Golang_. The result [2]_ is great. It takes less than 2 minutes to
finish encoding the words. This post shows you how to use the Go_ implementation
of succinct trie.

The following code gives an example of encoding and decoding:

.. show_github_file:: siongui go-succinct-data-structure-trie example/usage.go

Note that you must include space " " in your alphabet if default alphabet is not
used. The default alphabet is *[a-z ]*.

----

References:

.. [1] `Succinct Data Structures: Cramming 80,000 words into a Javascript file. <http://stevehanov.ca/blog/?id=120>`_
          (`source code <http://www.hanovsolutions.com/trie/Bits.js>`__)

.. [2] `siongui/go-succinct-data-structure-trie: Succinct Trie in Go - GitHub <https://github.com/siongui/go-succinct-data-structure-trie>`_

.. [3] `[JavaScript] Bug in Succinct Trie Implementation of Bits.js <{filename}../02/javascript-bug-in-succinct-trie-implementation-of-bits-js%en.rst>`_


.. _succinct trie: https://www.google.com/search?q=succinct+trie
.. _Bits.js: http://www.hanovsolutions.com/trie/Bits.js
.. _trie: https://www.google.com/search?q=trie
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
