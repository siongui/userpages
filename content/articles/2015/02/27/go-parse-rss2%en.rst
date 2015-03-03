[Golang] XML Parsing Example (7) - Parse RSS 2.0
################################################

:date: 2015-02-27 12:41
:tags: Go, Golang, XML, html, RSS
:category: Go
:summary: How to read XML/HTML files in Go programming language (for newbie)
          - Parse RSS 2.0 format.


This post shows how to parse XML of `RSS 2.0`_ format. The
`example RSS 2.0 feed`_ comes from Solidot_ website.

..
  .. show_github_file:: siongui userpages content/code/go-xml/example-6.xml

Souce Code
++++++++++

.. show_github_file:: siongui userpages content/code/go-xml/parse-6.go

.. note::

  The following struct fields in *Item* strcut have the type *template.HTML*

  .. code-block:: go

    Description template.HTML   `xml:"description"`
    Content     template.HTML   `xml:"encoded"`

  If *template.HTML* are replaced with *string*, the result is the same.


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}../17/go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}../19/go-parse-xml-example-2%en.rst>`_

.. [3] `[Golang] XML Parsing Example (3) <{filename}../21/go-parse-xml-example-3%en.rst>`_

.. [4] `[Golang] XML Parsing Example (4) <{filename}../24/go-parse-xml-example-4%en.rst>`_

.. [5] `[Golang] XML Parsing Example (5) - Parse OPML <{filename}../25/go-parse-opml%en.rst>`_

.. [6] `[Golang] XML Parsing Example (6) - Parse OPML Concisely <{filename}../26/go-parse-opml-concisely%en.rst>`_

.. [7] `[Golang] XML Parsing Example (7) - Parse RSS 2.0 <{filename}go-parse-rss2%en.rst>`_

.. [8] `[Golang] XML Parsing Example (8) - Parse Atom 1.0 <{filename}../28/go-parse-atom%en.rst>`_

.. [9] `[Golang] Convert Atom to RSS <{filename}../../03/02/go-convert-atom-to-rss-feed%en.rst>`_

.. [10] `[Golang] Parse Web Feed - RSS and Atom <{filename}../../03/03/go-parse-web-feed-rss-atom%en.rst>`_

----

Reference:

.. [a] `RSS Tutorial - w3schools.com <http://www.w3schools.com/rss/default.asp>`_

.. [b] `What is Atom 1.0 <http://www.tutorialspoint.com/rss/what-is-atom.htm>`_

.. [c] `php - ATOM to RSS feeds converter - Stack Overflow <http://stackoverflow.com/questions/16309944/atom-to-rss-feeds-converter>`_

.. [d] `RSS Syntax - w3schools.com <http://www.w3schools.com/rss/rss_syntax.asp>`_

.. [e] `RSS \<channel\> Element - w3schools.com <http://www.w3schools.com/rss/rss_channel.asp>`_

.. [f] `RSS \<item\> Element - w3schools.com <http://www.w3schools.com/rss/rss_item.asp>`_

.. [g] `xml - Difference between description and content:encoded tags in RSS2 - Stack Overflow <http://stackoverflow.com/questions/7220670/difference-between-description-and-contentencoded-tags-in-rss2>`_

.. [h] `golang xml parsing with CDATA and a colon in the keyname - Google Groups <https://groups.google.com/d/topic/golang-nuts/uBMo1BpaQCM>`_

.. [i] `[Golang] Convert Atom to RSS <{filename}../../03/02/go-convert-atom-to-rss-feed%en.rst>`_


.. _RSS 2.0: http://www.w3schools.com/rss/default.asp

.. _example RSS 2.0 feed: https://github.com/siongui/userpages/blob/master/content/code/go-xml/example-6.xml

.. _Solidot: http://www.solidot.org/
