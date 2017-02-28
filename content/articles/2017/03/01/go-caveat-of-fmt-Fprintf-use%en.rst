[Golang] Caveat of fmt.Fprintf Use
##################################

:date: 2017-03-01T06:19+08:00
:tags: Go, Golang, File Input/Output
:category: Go
:summary: Do not use fmt.Fprintf_ to write to file
:adsu: yes

Do not use fmt.Fprintf_ to write to file.

The following code writes string to file via fmt.Fprintf_:

.. show_github_file:: siongui userpages content/code/go/write-string-to-file/osfmt.go
.. adsu:: 2

The above code is ok if there is no ``%`` in the string. What if ``%`` in the
string?

.. code-block:: go

  if err := BadWriteStringToFile("bad.txt", "% in string\n"); err != nil {
  	panic(err)
  }

The output (*bad.txt*) is:

.. code-block:: txt

  %!i(MISSING)n string

This is not the result we want, so what is the good way to write string to file?
The answer is use io.Copy_:

.. show_github_file:: siongui userpages content/code/go/write-string-to-file/osiocopy.go
.. adsu:: 3

Again write the same string to file:

.. code-block:: go

  if err := WriteStringToFile("good.txt", "% in string\n"); err != nil {
  	panic(err)
  }

The output (*good.txt*):

.. code-block:: txt

  % in string

This is the result we want!

----

Tested on: ``Ubuntu Linux 16.10``, ``Go 1.8``.

----

References:

.. [1] `os - The Go Programming Language <https://golang.org/pkg/os/>`_
.. [2] `io - The Go Programming Language <https://golang.org/pkg/io/>`_
.. [3] `fmt - The Go Programming Language <https://golang.org/pkg/fmt/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _os: https://golang.org/pkg/os/
.. _Create: https://golang.org/pkg/os/#Create
.. _fmt: https://golang.org/pkg/fmt/
.. _Fprintf: https://golang.org/pkg/fmt/#Fprintf
.. _fmt.Fprintf: https://golang.org/pkg/fmt/#Fprintf
.. _io: https://golang.org/pkg/io/
.. _Copy: https://golang.org/pkg/io/#Copy
.. _io.Copy: https://golang.org/pkg/io/#Copy
