[Golang] reStructuredText Package
#################################

:date: 2016-01-23T22:35+08:00
:tags: Go, Golang, reStructuredText
:category: Go
:summary: Packages of reStructuredText_ (to HTML_) implementation in Go_.
:adsu: yes


Packages of reStructuredText_ (to HTML_) implementation in Golang_. After some
googling, I found [3]_ and [4]_, it seems only [3]_ provides rst-to-HTML
conversion, so I tried [3]_ only.

Install the package:

.. code-block:: bash

  $ go get -u github.com/hhatto/gorst

Test Code
+++++++++

Test program:

.. show_github_file:: siongui userpages content/code/go-reStructuredText/hhatto.go

Test rst file:

.. show_github_file:: siongui userpages content/code/go-reStructuredText/test.rst

Output of Test Code
+++++++++++++++++++

.. code-block:: txt

  <p><a href="https://siongui.github.io">my blog </a></p>

  <hr />

  <p><em>italic</em></p>

  <p><strong>bold</strong></p>

  <p><code>hello</code></p>

  <h2>Heading</h2>


----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.5.3``.

----

References:

.. [1] `Add support for native Go implementation of reStructuredText (reST) · Issue #1436 · spf13/hugo · GitHub <https://github.com/spf13/hugo/issues/1436>`_

.. [2] `GoのreStructuredTextパーサgorstをリリースしました | hexacosa.net <http://www.hexacosa.net/blog/detail/172/>`_

.. [3] `hhatto/gorst · GitHub <https://github.com/hhatto/gorst>`_ (Go implementation of reStructuredText)

.. [4] `demizer/go-rst · GitHub <https://github.com/demizer/go-rst>`_ (Parse reStructuredText with Go.)

.. [5] `Docutils text processing framework <https://www.google.com/search?q=Docutils+text+processing+framework>`_

.. [6] `reStructuredText tool support - Stack Overflow <http://stackoverflow.com/questions/2746692/restructuredtext-tool-support>`_

.. [7] `The Docutils Subversion Repository <http://docutils.sourceforge.net/docs/dev/repository.html>`_

.. [8] `Docutils: Documentation Utilities / Code /trunk/docutils/tools/rst2html5.py <http://sourceforge.net/p/docutils/code/HEAD/tree/trunk/docutils/tools/rst2html5.py>`_

.. [9] `Docutils: Documentation Utilities / Code /trunk/docutils/docutils/core.py <http://sourceforge.net/p/docutils/code/HEAD/tree/trunk/docutils/docutils/core.py#l328>`_

.. [10] `The Docutils Publisher <http://docutils.sourceforge.net/docs/api/publisher.html>`_

.. [11] `PEP 258 -- Docutils Design Specification <http://docutils.sourceforge.net/docs/peps/pep-0258.html>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
.. _HTML: https://www.google.com/search?q=HTML
