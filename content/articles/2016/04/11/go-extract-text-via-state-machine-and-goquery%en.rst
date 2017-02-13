[Golang] Extract Text via State Machine and goquery
###################################################

:date: 2016-04-11T20:02+08:00
:tags: Go, Golang, Commandline, DOM, goquery, Web Scrape, Go flag Package, html,
       File Input/Output
:category: Go
:summary: Extract text (i.e., footnote) in HTML_ via `state machine`_ and
          goquery_ in Go_ programming language.
:adsu: yes


Introduction
++++++++++++

Extract text (i.e., footnote) in HTML_ via `state machine`_ and goquery_ in
Golang_ (Go_ programming language).

Assume we have the following HTML:

.. show_github_file:: siongui userpages content/code/go-state-machine-get-footnote/index.html

We want to extract the text (i.e., footnote) starting from *Reference*, and
until *Updated*.

.. adsu:: 2

Install goquery_ Package
++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery


Read HTML_
++++++++++

.. show_github_file:: siongui userpages content/code/go-state-machine-get-footnote/html.go
.. adsu:: 3


Extract Text (Footnote)
+++++++++++++++++++++++

Find all children of *body* element in HTML_ document. Convert each child of
*body* element to text by `Text()`_ method. Process the text one line by one
line. If the text line starting with *Reference*, the state machine enters
*InFootnote* state, storing the text in the state machine. If the text line
starting with *Update*, the state machine leave *InFootnote* state and stop
storing the text. After all finished, output the text stored in the state
machine, which is the text we want.

.. show_github_file:: siongui userpages content/code/go-state-machine-get-footnote/footnote.go
.. adsu:: 4

Usage
+++++

Put above three files (``index.html``, ``html.go``, ``footnote.go``) together in
current directory. Run the following command:

.. code-block:: bash

  $ go run html.go footnote.go -input=index.html

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `jquery iterate over elements - Google search <https://www.google.com/search?q=jquery+iterate+over+elements>`_

.. [2] `github.com/PuerkitoBio/goquery - GoDoc <https://godoc.org/github.com/PuerkitoBio/goquery>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _state machine: https://www.google.com/search?q=state+machine
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _HTML: https://www.google.com/search?q=HTML
.. _Text(): https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Text
