[Golang] XML Parsing Example (3)
################################

:date: 2015-02-21 01:40
:tags: Go, Golang, XML, html
:category: Go
:summary: How to read XML/HTML files in Go programming language (for newbie)
          - Read a direct child element.


In this exmaple, we will parse a *div* element with a *span* child element,
:code:`<span>SpanText</span>`:

.. show_github_file:: siongui userpages content/code/go-xml/example-3.xml

Just as we declare a *struct* for parent *div* element, we also declare a
*struct* for the child *span* element, and add a struct field of the *span
struct* to the *div* struct.

.. show_github_file:: siongui userpages content/code/go-xml/parse-3.go


The output result:

.. code-block:: txt

  ChildSpan: {{ span} SpanText}


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}../17/go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}../19/go-parse-xml-example-2%en.rst>`_

.. [3] `[Golang] XML Parsing Example (3) <{filename}go-parse-xml-example-3%en.rst>`_

.. [4] `[Golang] XML Parsing Example (4) <{filename}../24/go-parse-xml-example-4%en.rst>`_

.. [5] `[Golang] XML Parsing Example (5) - Parse OPML <{filename}../25/go-parse-opml%en.rst>`_

.. [6] `[Golang] XML Parsing Example (6) - Parse OPML Concisely <{filename}../26/go-parse-opml-concisely%en.rst>`_

.. [7] `[Golang] XML Parsing Example (7) - Parse RSS 2.0 <{filename}../27/go-parse-rss2%en.rst>`_

.. [8] `[Golang] XML Parsing Example (8) - Parse Atom 1.0 <{filename}../28/go-parse-atom%en.rst>`_