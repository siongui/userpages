[Golang] Convert Atom to RSS
############################

:date: 2015-03-02 15:00
:tags: Go, Golang, XML, Atom, RSS
:category: Go
:summary: Convert Atom 1.0 to RSS 2.0 format in Go programming language.


This post shows how to convert `Atom 1.0`_ to `RSS 2.0`_. The
`example Atom 1.0 feed`_ comes from `kura.io`_ website. Note that only important
fields in RSS feed are copied to corresponding Atom equivalent.


Souce Code
++++++++++

.. show_github_file:: siongui userpages content/code/go-xml/atom2rss.go


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

References:

.. [1] `[Golang] XML Parsing Example (7) - Parse RSS 2.0 <{filename}../../02/27/go-parse-rss2%en.rst>`_

.. [2] `[Golang] XML Parsing Example (8) - Parse Atom 1.0 <{filename}../../02/28/go-parse-atom%en.rst>`_


.. _Atom 1.0: http://en.wikipedia.org/wiki/Atom_%28standard%29

.. _RSS 2.0: http://www.w3schools.com/rss/default.asp

.. _example Atom 1.0 feed: https://github.com/siongui/userpages/blob/master/content/code/go-xml/example-7.xml

.. _kura.io: https://kura.io/
