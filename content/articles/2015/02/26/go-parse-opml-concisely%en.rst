[Golang] XML Parsing Example (6) - Parse OPML Concisely
#######################################################

:date: 2015-02-26 12:13
:tags: Go, Golang, XML, html, OPML
:category: Go
:summary: How to read XML/HTML files in Go programming language (for newbie)
          - Parse OPML format concisely.
:adsu: yes


The previous post [5]_ shows how to parse OPML_ format, this post will show how
to parse it concisely by using ``>`` in tag. First we see what the official doc
in `encoding/xml`_ says about ``>``:

.. tip::

  If the XML element contains a sub-element whose name matches
  the prefix of a tag formatted as "a" or "a>b>c", unmarshal
  will descend into the XML structure looking for elements with the
  given names, and will map the innermost elements to that struct
  field. A tag starting with ">" is equivalent to one starting
  with the field name followed by ">".

Difficult to understand above explanation without example. Now take a look at
the *struct* defined in previous example:

.. code-block:: go

  type opml struct {
        XMLName         xml.Name        `xml:"opml"`
        Version         string          `xml:"version,attr"`
        Head            head
        Body            body
  }

  type head struct {
        XMLName         xml.Name        `xml:"head"`
        Title           string          `xml:"title"`
  }

  type body struct {
        XMLName         xml.Name        `xml:"body"`
        Outlines        []outline       `xml:"outline"`
  }

We can remove the **head** and **body** *struct* and keep meaningful content by:

.. code-block:: go

  type opml struct {
        XMLName         xml.Name        `xml:"opml"`
        Version         string          `xml:"version,attr"`
        OpmlTitle       string          `xml:"head>title"`
        Outlines        []outline       `xml:"body>outline"`
  }

*Title* field in **head** struct becomes *OpmlTitle* in **opml** struct, and
*Outlines* field in **body** struct becomes *Outlines* in **opml** struct.

Complete source code for concisely parsing the OPML in [5]_:

`Run code on Go Playground <https://play.golang.org/p/ha-LIDyrOn>`_

.. show_github_file:: siongui userpages content/code/go-xml/parse-5_2.go


The output result is the same as the result in [5]_.


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}../17/go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}../19/go-parse-xml-example-2%en.rst>`_

.. [3] `[Golang] XML Parsing Example (3) <{filename}../21/go-parse-xml-example-3%en.rst>`_

.. [4] `[Golang] XML Parsing Example (4) <{filename}../24/go-parse-xml-example-4%en.rst>`_

.. [5] `[Golang] XML Parsing Example (5) - Parse OPML <{filename}../25/go-parse-opml%en.rst>`_

.. [6] `[Golang] XML Parsing Example (6) - Parse OPML Concisely <{filename}go-parse-opml-concisely%en.rst>`_

.. [7] `[Golang] XML Parsing Example (7) - Parse RSS 2.0 <{filename}../27/go-parse-rss2%en.rst>`_

.. [8] `[Golang] XML Parsing Example (8) - Parse Atom 1.0 <{filename}../28/go-parse-atom%en.rst>`_

.. [9] `[Golang] Convert Atom to RSS <{filename}../../03/02/go-convert-atom-to-rss-feed%en.rst>`_

.. [10] `[Golang] Parse Web Feed - RSS and Atom <{filename}../../03/03/go-parse-web-feed-rss-atom%en.rst>`_

----

Reference:

.. [a] `OPML <http://en.wikipedia.org/wiki/OPML>`_

.. _OPML: http://en.wikipedia.org/wiki/OPML

.. _encoding/xml: http://golang.org/pkg/encoding/xml/#Unmarshal
