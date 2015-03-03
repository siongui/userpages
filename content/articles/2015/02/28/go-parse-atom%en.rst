[Golang] XML Parsing Example (8) - Parse Atom 1.0
#################################################

:date: 2015-02-28 09:41
:tags: Go, Golang, XML, html, Atom
:category: Go
:summary: How to read XML/HTML files in Go programming language (for newbie)
          - Parse Atom 1.0 format.


This post shows how to parse XML of `Atom 1.0`_ format. The
`example Atom 1.0 feed`_ comes from `kura.io`_ website.

..
  .. show_github_file:: siongui userpages content/code/go-xml/example-7.xml

Souce Code
++++++++++

.. show_github_file:: siongui userpages content/code/go-xml/parse-7.go


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}../17/go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}../19/go-parse-xml-example-2%en.rst>`_

.. [3] `[Golang] XML Parsing Example (3) <{filename}../21/go-parse-xml-example-3%en.rst>`_

.. [4] `[Golang] XML Parsing Example (4) <{filename}../24/go-parse-xml-example-4%en.rst>`_

.. [5] `[Golang] XML Parsing Example (5) - Parse OPML <{filename}../25/go-parse-opml%en.rst>`_

.. [6] `[Golang] XML Parsing Example (6) - Parse OPML Concisely <{filename}../26/go-parse-opml-concisely%en.rst>`_

.. [7] `[Golang] XML Parsing Example (7) - Parse RSS 2.0 <{filename}../27/go-parse-rss2%en.rst>`_

.. [8] `[Golang] XML Parsing Example (8) - Parse Atom 1.0 <{filename}go-parse-atom%en.rst>`_

.. [9] `[Golang] Convert Atom to RSS <{filename}../../03/02/go-convert-atom-to-rss-feed%en.rst>`_

.. [10] `[Golang] Parse Web Feed - RSS and Atom <{filename}../../03/03/go-parse-web-feed-rss-atom%en.rst>`_

----

Reference:

.. [a] `What is Atom 1.0 <http://www.tutorialspoint.com/rss/what-is-atom.htm>`_

.. [b] `php - ATOM to RSS feeds converter - Stack Overflow <http://stackoverflow.com/questions/16309944/atom-to-rss-feeds-converter>`_

.. [c] `Atom (standard) - Wikipedia, the free encyclopedia <http://en.wikipedia.org/wiki/Atom_%28standard%29>`_

.. [d] `src/encoding/xml/ - The Go Programming Language <http://golang.org/src/encoding/xml/>`_

.. [e] `[Golang] Convert Atom to RSS <{filename}../../03/02/go-convert-atom-to-rss-feed%en.rst>`_


.. _Atom 1.0: http://en.wikipedia.org/wiki/Atom_%28standard%29

.. _example Atom 1.0 feed: https://github.com/siongui/userpages/blob/master/content/code/go-xml/example-7.xml

.. _kura.io: https://kura.io/
