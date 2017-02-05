[Golang] Parse Web Feed - RSS and Atom
######################################

:date: 2015-03-03 18:37
:tags: Go, Golang, XML, Atom, RSS
:category: Go
:summary: Parse Atom 1.0/RSS 2.0 web feeds in Go programming language.
:adsu: yes


This post shows how to parse web feeds of `RSS 2.0`_ and `Atom 1.0`_. (The
logic will be explained below the source code)

Souce Code
++++++++++

.. show_github_file:: siongui userpages content/code/go-xml/parseFeed.go
.. adsu:: 2

Main logic of the *parseFeedContent* function:

  1. Given xml content, the function tries to parse the content in RSS format.
     If success, return the parsed result.

  2. If parsing RSS fails, then try to parse the content in Atom 1.0 format. If
     success, convert the Atom format to RSS format and return the parsed result.

  3. If both parsing RSS and Atom fails, return.

The `sample Atom 1.0 feed`_ comes from `kura.io`_ website.
The `sample RSS 2.0 feed`_ comes from Solidot_ website.
The `sample OPML xml`_ comes from my web feeds.


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

.. adsu:: 3

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}../../02/17/go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}../../02/19/go-parse-xml-example-2%en.rst>`_

.. [3] `[Golang] XML Parsing Example (3) <{filename}../../02/21/go-parse-xml-example-3%en.rst>`_

.. [4] `[Golang] XML Parsing Example (4) <{filename}../../02/24/go-parse-xml-example-4%en.rst>`_

.. [5] `[Golang] XML Parsing Example (5) - Parse OPML <{filename}../../02/25/go-parse-opml%en.rst>`_

.. [6] `[Golang] XML Parsing Example (6) - Parse OPML Concisely <{filename}../../02/26/go-parse-opml-concisely%en.rst>`_

.. [7] `[Golang] XML Parsing Example (7) - Parse RSS 2.0 <{filename}../../02/27/go-parse-rss2%en.rst>`_

.. [8] `[Golang] XML Parsing Example (8) - Parse Atom 1.0 <{filename}../../02/28/go-parse-atom%en.rst>`_

.. [9] `[Golang] Convert Atom to RSS <{filename}../02/go-convert-atom-to-rss-feed%en.rst>`_

.. [10] `[Golang] Parse Web Feed - RSS and Atom <{filename}go-parse-web-feed-rss-atom%en.rst>`_


.. _Atom 1.0: http://en.wikipedia.org/wiki/Atom_%28standard%29

.. _RSS 2.0: http://www.w3schools.com/rss/default.asp

.. _sample Atom 1.0 feed: https://github.com/siongui/userpages/blob/master/content/code/go-xml/example-7.xml

.. _kura.io: https://kura.io/

.. _sample RSS 2.0 feed: https://github.com/siongui/userpages/blob/master/content/code/go-xml/example-6.xml

.. _Solidot: http://www.solidot.org/

.. _sample OPML xml: https://github.com/siongui/userpages/blob/master/content/code/go-xml/example-5.xml
