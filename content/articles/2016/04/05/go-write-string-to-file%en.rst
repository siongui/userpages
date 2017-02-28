[Golang] Write String to File
#############################

:date: 2016-04-05T21:13+08:00
:modified: 2017-03-01T05:46+08:00
:tags: Go, Golang, File Input/Output
:category: Go
:summary: Go_ - write string to file via os_ Create_ method and io_ Copy_
          method.
:adsu: yes

Golang_ - write string to file via os_ Create_ method and io_ Copy_ method.

.. show_github_file:: siongui userpages content/code/go/write-string-to-file/osiocopy.go
.. adsu:: 2

Usage:

.. code-block:: go

  if err := WriteStringToFile("good.txt", "% in string\n"); err != nil {
  	panic(err)
  }

Output (*good.txt*):

.. code-block:: txt

  % in string

Caveat: Do not use fmt.Fprintf_ to write string to file. See [4]_ for more
details.

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.8``.

----

References:

.. [1] `os - The Go Programming Language <https://golang.org/pkg/os/>`_
.. [2] `io - The Go Programming Language <https://golang.org/pkg/io/>`_
.. [3] `fmt - The Go Programming Language <https://golang.org/pkg/fmt/>`_
.. [4] `[Golang] Caveat of fmt.Fprintf Use <{filename}../../../2017/03/01/go-caveat-of-fmt-Fprintf-use%en.rst>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _os: https://golang.org/pkg/os/
.. _Create: https://golang.org/pkg/os/#Create
.. _fmt.Fprintf: https://golang.org/pkg/fmt/#Fprintf
.. _io: https://golang.org/pkg/io/
.. _Copy: https://golang.org/pkg/io/#Copy
