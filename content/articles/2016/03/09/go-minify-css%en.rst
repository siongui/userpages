[Golang] Minify CSS
###################

:date: 2016-03-09T03:53+08:00
:tags: Go, Golang, String Manipulation, Regular Expression, CSS,
       remove trailing newline, remove carriage return, File Input/Output,
       Minify HTML/CSS/JavaScript, Read Lines
:category: Go
:summary: `Minify CSS`_ via Go_ programming language.
:adsu: yes

`Minify CSS`_ via Golang_.

The steps:

1. Concatenate all CSS file.
2. Remove C/C++ comments [2]_ [3]_
3. Remove all leading and trailing white space of each line.

.. show_github_file:: siongui userpages content/code/go-minify-css/mincss.go
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/go-minify-css/mincss_test.go
.. adsu:: 3

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `regex match c comments <https://www.google.com/search?q=regex+match+c+comments>`_

.. [2] `Finding Comments in Source Code Using Regular Expressions – Stephen Ostermiller <http://blog.ostermiller.org/find-comment>`_

.. [3] `Comments - cppreference.com <http://en.cppreference.com/w/cpp/comment>`_

.. [4] `regexp - The Go Programming Language <https://golang.org/pkg/regexp/>`_

.. [5] `ioutil - The Go Programming Language <https://golang.org/pkg/io/ioutil/>`_

.. [6] `bufio - The Go Programming Language <https://golang.org/pkg/bufio/>`_

.. [7] `bytes - The Go Programming Language <https://golang.org/pkg/bytes/>`_

.. [8] `strings - The Go Programming Language <https://golang.org/pkg/strings/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Minify CSS: https://www.google.com/search?q=Minify+CSS
