[Golang] String startswith and endswith
#######################################

:date: 2017-03-17T01:55+08:00
:tags: Go, Golang, String Manipulation
:category: Go
:summary: Go equivalent of Python string startswith and endswith methods.
:og_image: https://talks.golang.org/2014/go4java/img/gopher.jpg
:adsu: yes

Go equivalent of Python string startswith_ and endswith_ methods.

Actually strings_ package in Go standard library already provides these two
methods, but with different names. HasPrefix_ is Go's string startswith, and
HasSuffix_ is Go's string endswith.

Go String starts with
+++++++++++++++++++++

.. code-block:: go

  import "strings"

  // true
  strings.HasPrefix("Gopher", "Go")

`Run Example on Go Playground <https://play.golang.org/p/46wlxCiDVG>`__

.. adsu:: 2

Go String ends with
+++++++++++++++++++

.. code-block:: go

  import "strings"

  // true
  strings.HasSuffix("Amigo", "go")

`Run Example on Go Playground <https://play.golang.org/p/tae9j8T5MC>`__

----

References:

.. [1] `func HasPrefix - strings - The Go Programming Language <https://golang.org/pkg/strings/#HasPrefix>`_
.. [2] `func HasSuffix - strings - The Go Programming Language <https://golang.org/pkg/strings/#HasSuffix>`_

.. _startswith: https://www.tutorialspoint.com/python/string_startswith.htm
.. _endswith: https://www.tutorialspoint.com/python/string_endswith.htm
.. _strings: https://golang.org/pkg/strings/
.. _HasPrefix: https://golang.org/pkg/strings/#HasPrefix
.. _HasSuffix: https://golang.org/pkg/strings/#HasSuffix
