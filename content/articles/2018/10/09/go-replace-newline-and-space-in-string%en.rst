[Golang] Replace Space and Newline in String
############################################

:date: 2018-10-09T22:26+08:00
:tags: Go, Golang, Regular Expression, String Manipulation
:category: Go
:summary: Use Go *regexp* package to repace spaces and newline in the string.
:og_image: https://regex101.com/preview.png
:adsu: yes


Use Go_ regexp_ package to repace spaces and newline in the string with only one
space.

Problem
+++++++

  We have a string as follows:

  .. code-block:: go

    `CONDUIT 
                    FITTINGS`

  There are spaces and newline in the string. We want the string to become

  .. code-block:: go

    `CONDUIT FITTINGS`

  How to do it?


Solution
++++++++

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/uaEY0sjBImv>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  	"regexp"
  )

  var str = `CONDUIT 
                  FITTINGS`

  func TrimSpaceNewlineInString(s string) string {
  	re := regexp.MustCompile(` +\r?\n +`)
  	return re.ReplaceAllString(s, " ")
  }

  func main() {
  	fmt.Println(str)
  	fmt.Println(TrimSpaceNewlineInString(str))
  }

.. adsu:: 2

Trick: If you are not sure whether the regular expression is written correctly,
use [2]_ to test it online first.

**Output**

.. code-block:: txt

  CONDUIT 
                  FITTINGS
  CONDUIT FITTINGS

Tested on: `Go Playground`_

----

References:

.. [1] `Golang: Issues replacing newlines in a string from a text file - Stack Overflow <https://stackoverflow.com/questions/34112382/golang-issues-replacing-newlines-in-a-string-from-a-text-file>`_
.. [2] `Online regex tester and debugger: PHP, PCRE, Python, Golang and JavaScript <https://regex101.com/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _regexp: https://golang.org/pkg/regexp/
.. _Go Playground: https://play.golang.org/
