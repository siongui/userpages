Pass Command-line Arguments to Makefile and Go Program
######################################################

:date: 2018-05-15T22:35+08:00
:tags: Go, Golang, Commandline, Go flag Package, Makefile, Bash
:category: Bash
:summary: Pass command-line arguments to *make* and then pass then arguments to
          Go program in Makefile.
:og_image: https://www.tutorialspoint.com/makefile/images/makefile-mini-logo.jpg
:adsu: yes


Pass command line arguments to make_ first, and then pass the arguments again
to Go program in *Makefile*.

The reason to pass the same arguments twice like this is because I use
*Makefile* to manage my workflow so the arguments are passed twice.

**Makefile**:

.. code-block:: makefile

  target:
  	@go run argument.go -argument=$(arg)


**Go**: *argument.go*

.. code-block:: go

  package main

  import (
  	"flag"
  	"fmt"
  )

  func main() {
  	arg := flag.String("argument", "", "argument from Makefile")
  	flag.Parse()

  	fmt.Println("argument:", *arg)
  }


**Example**:

.. code-block:: bash

  $ make target arg=hello
  argument: hello

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 18.04``, ``Go 1.10.2``

References:

.. [1] | `makefile pass arguments - Google search <https://www.google.com/search?q=makefile+pass+arguments>`_
       | `makefile pass arguments - DuckDuckGo search <https://duckduckgo.com/?q=makefile+pass+arguments>`_
       | `makefile pass arguments - Ecosia search <https://www.ecosia.org/search?q=makefile+pass+arguments>`_
       | `makefile pass arguments - Qwant search <https://www.qwant.com/?q=makefile+pass+arguments>`_
       | `makefile pass arguments - Bing search <https://www.bing.com/search?q=makefile+pass+arguments>`_
       | `makefile pass arguments - Yahoo search <https://search.yahoo.com/search?p=makefile+pass+arguments>`_
       | `makefile pass arguments - Baidu search <https://www.baidu.com/s?wd=makefile+pass+arguments>`_
       | `makefile pass arguments - Yandex search <https://www.yandex.com/search/?text=makefile+pass+arguments>`_
.. [2] `makefile - Passing additional variables from command line to make - Stack Overflow <https://stackoverflow.com/a/2826069>`_
.. [3] `[Golang] Parse Command Line Arguments - String Variable <{filename}/articles/2016/12/21/go-parse-commandline-arguments-string-variable%en.rst>`_

.. _make: https://www.google.com/search?q=gnu+make
