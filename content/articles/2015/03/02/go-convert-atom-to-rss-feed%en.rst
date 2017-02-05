[Golang] Convert Atom to RSS
############################

:date: 2015-03-02 15:00
:tags: Go, Golang, XML, Atom, RSS
:category: Go
:summary: Convert Atom 1.0 to RSS 2.0 format in Go programming language.
:adsu: yes


This post shows how to convert `Atom 1.0`_ to `RSS 2.0`_. The
`example Atom 1.0 feed`_ comes from `kura.io`_ website. Note that only important
fields in RSS feed are copied to corresponding Atom equivalent.


Souce Code
++++++++++

`Run code on Go Playground <https://play.golang.org/p/fMzOUkeVzV>`_

.. show_github_file:: siongui userpages content/code/go-xml/atom2rss.go
.. adsu:: 2

Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}../../02/17/go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}../../02/19/go-parse-xml-example-2%en.rst>`_

.. [3] `[Golang] XML Parsing Example (3) <{filename}../../02/21/go-parse-xml-example-3%en.rst>`_

.. [4] `[Golang] XML Parsing Example (4) <{filename}../../02/24/go-parse-xml-example-4%en.rst>`_

.. [5] `[Golang] XML Parsing Example (5) - Parse OPML <{filename}../../02/25/go-parse-opml%en.rst>`_
.. adsu:: 3
.. [6] `[Golang] XML Parsing Example (6) - Parse OPML Concisely <{filename}../../02/26/go-parse-opml-concisely%en.rst>`_

.. [7] `[Golang] XML Parsing Example (7) - Parse RSS 2.0 <{filename}../../02/27/go-parse-rss2%en.rst>`_

.. [8] `[Golang] XML Parsing Example (8) - Parse Atom 1.0 <{filename}../../02/28/go-parse-atom%en.rst>`_

.. [9] `[Golang] Convert Atom to RSS <{filename}go-convert-atom-to-rss-feed%en.rst>`_

.. [10] `[Golang] Parse Web Feed - RSS and Atom <{filename}../03/go-parse-web-feed-rss-atom%en.rst>`_


.. _Atom 1.0: http://en.wikipedia.org/wiki/Atom_%28standard%29

.. _RSS 2.0: http://www.w3schools.com/rss/default.asp

.. _example Atom 1.0 feed: https://github.com/siongui/userpages/blob/master/content/code/go-xml/example-7.xml

.. _kura.io: https://kura.io/
