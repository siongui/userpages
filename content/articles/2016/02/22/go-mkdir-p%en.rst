[Golang] mkdir -p
#################

:date: 2016-02-22T05:55+08:00
:tags: Go, Golang, Bash, Commandline
:category: Go
:summary: `mkdir -p`_ command in Go_


`mkdir -p`_ command in Golang_

.. code-block:: go

  import "os"

  func Mkdirp(dirpath string) {
          os.MkdirAll(dirpath, 0775)
  }

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `golang mkdir p <https://www.google.com/search?q=golang+mkdir+p>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _mkdir -p: http://linux.die.net/man/1/mkdir
