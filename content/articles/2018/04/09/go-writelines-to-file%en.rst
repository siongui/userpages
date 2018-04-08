[Golang] Write Lines to File
############################

:date: 2018-04-09T06:26+08:00
:tags: Go, Golang, String Manipulation, File Input/Output, Go strings Package
:category: Go
:summary: Write a list of strings to file in Go programming language.
:og_image: https://i.ytimg.com/vi/pkVQzanZY9w/hqdefault.jpg
:adsu: yes


Write a list of strings (lines) to file in Go. First use `strings.Join`_ to
convert lines to a single string. Then use `ioutil.WriteFile`_ to write the
single string to file.

.. code-block:: go

  import (
  	"io/ioutil"
  	"strings"
  )

  func WriteLinesToFile(lines []string, filename string) (err error) {
  	return ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")+"\n"), 0644)
  }

If you want to know how to read lines from file, see [1]_.

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10.1``.

----

References:

.. [1] `[Golang] Read Lines From File or String <{filename}/articles/2016/04/06/go-readlines-from-file-or-string%en.rst>`_

.. _strings.Join: https://golang.org/pkg/strings/#Join
.. _ioutil.WriteFile: https://golang.org/pkg/io/ioutil/#WriteFile
