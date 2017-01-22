[Golang] Create reStructuredText Metadata via text/template Package
###################################################################

:date: 2016-04-22T22:11+08:00
:tags: Go, Golang, String Manipulation, Golang template, reStructuredText,
       Go text/width
:category: Go
:summary: Create reStructuredText_ metadata via Go_ `text/template`_ and
          `text/width`_ package.
:adsu: yes


Create reStructuredText_ metadata via Go_ `text/template`_ and `text/width`_
package.

Install `package width`_:

.. code-block:: bash

  $ go get -u golang.org/x/text/width

Source code:

.. show_github_file:: siongui userpages content/code/go-rst-metadata-template/rstmeta.go

.. show_github_file:: siongui userpages content/code/go-rst-metadata-template/rstmeta_test.go

Output:

.. code-block:: txt

  === RUN   TestRstMeta
  test title 測試標題
  ===================

  :date: 2016-04-22T20:58+08:00
  :tags: tag1, tag2
  :category: category1
  :summary: my summary


  --- PASS: TestRstMeta (0.00s)
  PASS


----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6.1``.

----

References:

.. [1] `[Golang] Get UTF-8 String Width <{filename}../../03/23/go-utf8-string-width%en.rst>`_

       `golang.org/x/text/width - GoDoc <https://godoc.org/golang.org/x/text/width>`_

.. [2] `go repeat string - Google search <https://www.google.com/search?q=go+repeat+string>`_

.. [3] `use template · twnanda/twnanda@98cb592 · GitHub <https://github.com/twnanda/twnanda/commit/98cb5927bf65d314aaf456386c95c6802496eb8e>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _package width: https://golang.org/x/text/width
.. _text/width: https://golang.org/x/text/width
.. _text/template: https://golang.org/pkg/text/template/
.. _reStructuredText: https://www.google.com/search?q=reStructuredText
