[Golang] Convert All HTML Links to reStructuredText via goquery
###############################################################

:date: 2016-04-12T22:53+08:00
:tags: Go, Golang, DOM, goquery, Web Scrape, reStructuredText, Golang template,
       html
:category: Go
:summary: Convert all `HTML links`_ to restructuredtext_ in a webpage via
          goquery_ in Go_ programming language.
:adsu: yes


Introduction
++++++++++++

Convert all `HTML links`_ to restructuredtext_ in a webpage via goquery_ in
Golang_ (Go_ programming language). If you want to do the same thing via
`net/html`_ package without goquery_, see [2]_.


Install goquery_ Package
++++++++++++++++++++++++

.. code-block:: bash

  $ go get -u github.com/PuerkitoBio/goquery


Source Code
+++++++++++

Use goquery_ `Find()`_ and `Each()`_ method to iterate over all `HTML links`_,
and use `Text()`_ and `Attr()`_ method to retrieve the text and href of each
link. Print to screen via os.Stdout_ and `text/template`_ package.

.. adsu:: 2
.. show_github_file:: siongui userpages content/code/go-html-links-to-rst/link2rst.go
.. adsu:: 3


Usage
+++++

.. code-block:: bash

  $ go run link2rst.go


----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `github.com/PuerkitoBio/goquery - GoDoc <https://godoc.org/github.com/PuerkitoBio/goquery>`_

.. [2] `[Golang] Iterate over All DOM Elements in HTML <{filename}../10/go-iterate-over-all-dom-elements-in-html%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _goquery: https://github.com/PuerkitoBio/goquery
.. _HTML links: http://www.w3schools.com/html/html_links.asp
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _Text(): https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Text
.. _Attr(): https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Attr
.. _Find(): https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Find
.. _Each(): https://godoc.org/github.com/PuerkitoBio/goquery#Selection.Each
.. _net/html: https://godoc.org/golang.org/x/net/html
.. _os.Stdout: https://golang.org/pkg/os/#pkg-variables
.. _text/template: https://golang.org/pkg/text/template/
