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
