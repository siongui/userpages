[Golang] Get UTF-8 String Width
###############################

:date: 2016-03-23T21:46+08:00
:tags: Go, Golang, String Manipulation, Locale, i18n
:category: Go
:summary: Get UTF-8_ string width (width of English letter is 1, width of
          `CJK character`_ is 2) in Go_ programming language.


Get UTF-8_ string width (width of English letter is 1, width of `CJK character`_
is 2) in Golang_.

Install `package width`_:

.. code-block:: bash

  $ go get -u golang.org/x/text/width

Source code:

.. show_github_file:: siongui userpages content/code/go-char-width/width.go

.. show_github_file:: siongui userpages content/code/go-char-width/width_test.go

----

References:

.. [1] `[Golang] Iterate Over UTF-8 Strings (non-ASCII strings) <{filename}../../02/03/go-iterate-over-utf8-non-ascii-string%en.rst>`_

.. [2] `golang get character width <https://www.google.com/search?q=golang+get+character+width>`_

       `width - GoDoc - The Go Programming Language <https://golang.org/x/text/width>`_

.. [3] `python get character width <https://www.google.com/search?q=python+get+character+width>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _package width: https://golang.org/x/text/width
.. _UTF-8: https://en.wikipedia.org/wiki/UTF-8
.. _CJK character: https://en.wikipedia.org/wiki/CJK_characters
