[Golang] XML Parsing Example (4)
################################

:date: 2015-02-24 10:09
:tags: Go, Golang, XML, html
:category: Go
:summary: How to read XML/HTML files in Go programming language (for newbie)
          - Read multiple direct child elements.


In this exmaple, we will parse a *div* element with multiple *span* child
elements:

.. show_github_file:: siongui userpages content/code/go-xml/example-4.xml

To parse multiple *span* child elements, the struct field in *div* struct of
previous example [3]_:

.. code-block:: go

    ChildSpan	span

becomes:

.. code-block:: go

    SpanList	[]span		`xml:"span"`

The type :code:`span` becomes :code:`[]span`, and **`xml:"span`** is added to
indicate the tag name of child elements.

.. show_github_file:: siongui userpages content/code/go-xml/parse-4.go


The output result:

.. code-block:: txt

  {{ div} [{SpanText1} {SpanText2} {SpanText3}]}

.. note::

  If you uncomment the following line:

  .. code-block:: go

    //	XMLName		xml.Name	`xml:"span"`

  The output will be:

  .. code-block:: txt

    {{ div} [{{ span} SpanText1} {{ span} SpanText2} {{ span} SpanText3}]}



Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}../17/go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}../19/go-parse-xml-example-2%en.rst>`_

.. [3] `[Golang] XML Parsing Example (3) <{filename}../21/go-parse-xml-example-3%en.rst>`_

.. [4] `[Golang] XML Parsing Example (4) <{filename}go-parse-xml-example-4%en.rst>`_
