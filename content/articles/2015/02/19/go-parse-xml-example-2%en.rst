[Golang] XML Parsing Example (2)
################################

:date: 2015-02-19 06:08
:tags: Go, Golang, XML, html
:category: Go
:summary: How to read XML/HTML files in Go programming language (for newbie).


following previous post [1]_, we add the attribute ``class="myClass"`` to the
*div* element in our sample XML:

.. show_github_file:: siongui userpages content/code/go-xml/example-2.xml

It is easy to read the attribute. Just add the following *struct field* in the
original struct:

.. code-block:: go

  Class	string		`xml:"class,attr"`

``"class,attr"`` means to read the attribute whose name is *class*.

.. show_github_file:: siongui userpages content/code/go-xml/parse-2.go

The output result:

.. code-block:: bash

  XMLName: { div}
  Class: myClass
  Content: Example

Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

*[Golang] XML Parsing Example* series:

.. [1] `[Golang] XML Parsing Example (1) <{filename}../17/go-parse-xml-example-1%en.rst>`_

.. [2] `[Golang] XML Parsing Example (2) <{filename}go-parse-xml-example-2%en.rst>`_

