[Golang] Read Command-line Arguments Example
############################################

:date: 2015-02-18 02:29
:tags: Go, Golang, Commandline, Go flag Package
:category: Go
:summary: How to parse command-line arguments (flags) in Go programming language.
:adsu: yes

On \*nix system we often run the program with arguments. How do we get the
command-line arguments in Go_ program? Please see the following example:

.. show_github_file:: siongui userpages content/articles/2015/02/18/cmd.go

In the example we define a boolean flag *production*, which indicates whether
the program runs in production mode. The default value is *false*, and the usage
string is `if run in production mode`.

Now we build our program `cmd.go` by:

.. code-block:: bash

  $ go build cmd.go

A binary program named ``cmd`` will be under the same directory. Try to run the
binary without arguments:

.. code-block:: bash

  $ ./cmd
  Production Mode: false

Next run the binary with `-production=true`:

.. code-block:: bash

  $ ./cmd -production=true
  Production Mode: true

We can also read usage of the program by:

.. code-block:: bash

  $ ./cmd -h
  Usage of ./cmd:
    -production=false: if run in production mode

You can also read usage by `--help`:

.. code-block:: bash

  $ ./cmd --help
  Usage of ./cmd:
    -production=false: if run in production mode


Source code tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

References:

.. [1] `flag - The Go Programming Language <http://golang.org/pkg/flag/>`_

.. [2] `Go by Example: Command-Line Flags <https://gobyexample.com/command-line-flags>`_

.. [3] `[Golang] Parse Command Line Arguments - String Variable <{filename}../../../2016/12/21/go-parse-commandline-arguments-string-variable%en.rst>`_


.. _Go: https://golang.org/
