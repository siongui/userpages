[Bash] List Go Source Files Excluding Test Files
################################################

:date: 2018-03-23T23:29+08:00
:tags: Bash, Commandline, find command, grep command, Go, Golang, Makefile
:category: Bash
:summary: List *.go* source files excluding *_test.go* files via *find* and
          *grep* command in *bash*.
:og_image: https://i.ytimg.com/vi/WASzwAOpolg/maxresdefault.jpg
:adsu: yes


List *.go* source files, excluding *_test.go* files via find_ and grep_ command
in Bash_.

.. code-block:: bash

  $ find *.go | grep -v _test.go

Sometimes we just want to run one of the test files, but need all *.go* files to
run the single test file. We can run as follows in our Makefile_.

.. code-block:: makefile

  ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

  default:
  	@go test -v $(ALL_GO_SOURCES) oneofmy_test.go

.. adsu:: 2

----

References:

.. [1] `go - The Go Programming Language <https://golang.org/cmd/go/>`_

.. _find: https://linux.die.net/man/1/find
.. _grep: http://linuxcommand.org/lc3_man_pages/grep1.html
.. _Bash: https://www.google.com/search?q=Bash
.. _Makefile: https://www.google.com/search?q=Makefile
