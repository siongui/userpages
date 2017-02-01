[Golang] Minify HTML
####################

:date: 2016-03-10T05:15+08:00
:tags: Go, Golang, String Manipulation, Regular Expression, html,
       remove trailing newline, remove carriage return, File Input/Output,
       Minify HTML/CSS/JavaScript, Read Lines
:category: Go
:summary: `Minify HTML`_ via Go_ programming language.
:adsu: yes

`Minify HTML`_ via Golang_.

The steps:

1. Remove HTML comments [2]_
2. Remove all leading and trailing white space of each line.
3. Pad a single space to the line if its length > 0.

.. show_github_file:: siongui userpages content/code/go-minify-html/minhtml.go
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/go-minify-html/minhtml_test.go
.. adsu:: 3

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `regex html comments <https://www.google.com/search?q=regex+html+comments>`_

.. [2] `php - RegExp to strip HTML comments - Stack Overflow <http://stackoverflow.com/a/1084759>`_

.. [3] `regexp - The Go Programming Language <https://golang.org/pkg/regexp/>`_

.. [4] `ioutil - The Go Programming Language <https://golang.org/pkg/io/ioutil/>`_

.. [5] `bufio - The Go Programming Language <https://golang.org/pkg/bufio/>`_

.. [6] `bytes - The Go Programming Language <https://golang.org/pkg/bytes/>`_

.. [7] `strings - The Go Programming Language <https://golang.org/pkg/strings/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Minify HTML: https://www.google.com/search?q=Minify+HTML
