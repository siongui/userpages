[Golang] XML Parsing Example (1)
################################

:date: 2015-02-17 13:24
:modified: 2015-02-21 01:24
:tags: Go, Golang, XML, html
:category: Go
:summary: How to read XML/HTML files in Go programming language (for newbie)
          - Read element and its content.
:adsu: yes

Assume we have a XML file as follows:

.. show_github_file:: siongui userpages content/code/go-xml/example-1.xml

We would like to parse the XML file and extract the useful content. Here is how
we do in Go_ programming language:

`Run Code on Go Playground <https://play.golang.org/p/XiukTqZYv2>`_

.. show_github_file:: siongui userpages content/code/go-xml/parse-1.go

.. adsu:: 2

Put the above two files on the same directory and run the code. The ouput will
be:

.. code-block:: bash

  XMLName: { div}
  Content: Example

.. note::

  In the line:

  .. code-block:: go

    XMLName	xml.Name	`xml:"div"`

  :code:`XMLName` and :code:`xml.Name` are fixed syntax for reading ``div``
  element, you can not change them at will.

.. note::

  Also in the line:

  .. code-block:: go

    Content	string		`xml:",chardata"`

  The first letter of the variable :code:`Content` must be capital. If you use
  :code:`content`, the Go_ parser will fail to read the content in *div*
  element. You can use another name for the variable, as long as the first
  letter is capital.

.. note::

  If you replace the line:

  .. code-block:: go

    Content	string		`xml:",chardata"`

  with

  .. code-block:: go

    Content	string		`xml:",innerxml"`

  i.e., replace *,chardata* with *,innerxml*. The output result will be the same
  because the raw XML nested inside the *div* element is the same as the
  character data of the *div* element in this case.

----

Makefile for automating the development:

.. show_github_file:: siongui userpages content/code/go-xml/Makefile

Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}../19/go-parse-xml-example-2%en.rst>`_

.. [3] `[Golang] XML Parsing Example (3) <{filename}../21/go-parse-xml-example-3%en.rst>`_

.. [4] `[Golang] XML Parsing Example (4) <{filename}../24/go-parse-xml-example-4%en.rst>`_

.. [5] `[Golang] XML Parsing Example (5) - Parse OPML <{filename}../25/go-parse-opml%en.rst>`_

.. [6] `[Golang] XML Parsing Example (6) - Parse OPML Concisely <{filename}../26/go-parse-opml-concisely%en.rst>`_

.. [7] `[Golang] XML Parsing Example (7) - Parse RSS 2.0 <{filename}../27/go-parse-rss2%en.rst>`_

.. [8] `[Golang] XML Parsing Example (8) - Parse Atom 1.0 <{filename}../28/go-parse-atom%en.rst>`_

.. [9] `[Golang] Convert Atom to RSS <{filename}../../03/02/go-convert-atom-to-rss-feed%en.rst>`_

.. [10] `[Golang] Parse Web Feed - RSS and Atom <{filename}../../03/03/go-parse-web-feed-rss-atom%en.rst>`_

----

References:

.. [a] `xml - The Go Programming Language <http://golang.org/pkg/encoding/xml/>`_

.. [b] `src/encoding/xml/example_test.go - The Go Programming Language <https://golang.org/src/encoding/xml/example_test.go>`_

.. [c] `Reading XML Documents in Go <http://www.goinggo.net/2013/06/reading-xml-documents-in-go.html>`_

.. [d] `Goで任意のXMLを処理する - GolangRdyJp <http://golang.rdy.jp/2015/11/08/anyxml/>`_


.. _Go: https://golang.org/
