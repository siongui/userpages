[Golang] Read Lines From URL
############################

:date: 2017-02-02T04:18+08:00
:tags: Go, Golang, String Manipulation, Read Lines
:category: Go
:summary: `Read lines`_ from web URL_ via Go_ programming language.
:adsu: yes


.. contents:: First show how to readlines_ from URL_, and then extend to read
              lines from file_ and string_.

readlines_ from URL_
++++++++++++++++++++

.. show_github_file:: siongui userpages content/code/go/readlines-from/url.go
.. adsu:: 2

Note that *LinesFromReader* func_ accepts argument of io.Reader_, which is an
interface_. Later this function will be re-used to read lines from file/string.

**Usage of UrlToLines**:

.. show_github_file:: siongui userpages content/code/go/readlines-from/url_test.go

----

readlines_ from file_
+++++++++++++++++++++

Use the following code with the *LinesFromReader* func_ in the previous section,
we can read a file line by line:

.. show_github_file:: siongui userpages content/code/go/readlines-from/file.go
.. adsu:: 3

**Usage of FileToLines**:

.. show_github_file:: siongui userpages content/code/go/readlines-from/file_test.go

----

readlines_ from string_
+++++++++++++++++++++++

Use the following code with the *LinesFromReader* func_ in the previous section,
we can read a string line by line:

.. show_github_file:: siongui userpages content/code/go/readlines-from/string.go

**Usage of StringToLines**:

.. show_github_file:: siongui userpages content/code/go/readlines-from/string_test.go

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.5``

----

References:

.. [1] `[Golang] Read Lines From File or String <{filename}../../../2016/04/06/go-readlines-from-file-or-string%en.rst>`_
.. [2] `read lines from file or url · siongui/go-rst@f0cc55b · GitHub <https://github.com/siongui/go-rst/commit/f0cc55bf0d878811956cd70d2ec99d4ee58bec15>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Read lines: https://www.google.com/search?q=Read+lines
.. _readlines: https://www.google.com/search?q=Readlines
.. _file: https://www.google.com/search?q=golang+file
.. _string: https://www.google.com/search?q=golang+string
.. _interface: https://www.google.com/search?q=golang+interface
.. _func: https://www.google.com/search?q=golang+function
.. _URL: https://www.google.com/search?q=URL
.. _io.Reader: https://golang.org/pkg/io/#Reader
