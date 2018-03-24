[Golang] Insert Line or String to File
######################################

:date: 2017-01-30T05:57+08:00
:tags: Go, Golang, String Manipulation, File Input/Output, Read Lines
:category: Go
:summary: Insert a line or string to *n-th* line of the file
          via Go_ programming language.
:og_image: https://www.howtogeek.com/wp-content/uploads/2012/01/vi-tutorial-3.png
:adsu: yes


Insert a string or line to the position of *n-th* (`0-indexed`_) line of the
file via Golang_.

The logic of the code:

- Read the file line by line into a slice of type *string* [1]_.
- `for range`_ the *string* slice to insert the string to *n-th* line of the
  file. Note that line is actually a string with trailing newline_ ``'\n'``. If
  you want to insert a line, just append ``'\n'`` to the end of the string.

The following is full source code:

.. show_github_file:: siongui userpages content/code/go/insert-line-to-file/utils.go
.. adsu:: 2

**Usage**:

Assume we have a text file named ``test.txt`` as follows:

.. code-block:: txt

  1
  2
  3
  4
  5

We want to insert a line with text ``hello world!`` to the third line (index 2)
of the file. Call ``InsertStringToFile("test.txt", "hello world!\n", 2)``

.. show_github_file:: siongui userpages content/code/go/insert-line-to-file/utils_test.go

The result is as follows

.. code-block:: txt

  1
  2
  hello world!
  3
  4
  5

.. adsu:: 3

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.5``

----

References:

.. [1] `[Golang] Read Lines From File or String <{filename}../../../2016/04/06/go-readlines-from-file-or-string%en.rst>`_
.. [2] `[Golang] Append Line/String to File <{filename}../23/go-append-text-string-to-file%en.rst>`_

.. [3] `golang insert line - Google search <https://www.google.com/search?q=golang+insert+line>`_

       `golang insert line - DuckDuckGo search <https://duckduckgo.com/?q=golang+insert+line>`_

       `golang insert line - Bing search <https://www.bing.com/search?q=golang+insert+line>`_

       `golang insert line - Yahoo search <https://search.yahoo.com/search?p=golang+insert+line>`_

       `golang insert line - Baidu search <https://www.baidu.com/s?wd=golang+insert+line>`_

       `golang insert line - Yandex search <https://www.yandex.com/search/?text=golang+insert+line>`_

.. [4] `golang insert line to file - Google search <https://www.google.com/search?q=golang+insert+line+to+file>`_

       `golang insert line to file - DuckDuckGo search <https://duckduckgo.com/?q=golang+insert+line+to+file>`_

       `golang insert line to file - Bing search <https://www.bing.com/search?q=golang+insert+line+to+file>`_

       `golang insert line to file - Yahoo search <https://search.yahoo.com/search?p=golang+insert+line+to+file>`_

       `golang insert line to file - Baidu search <https://www.baidu.com/s?wd=golang+insert+line+to+file>`_

       `golang insert line to file - Yandex search <https://www.yandex.com/search/?text=golang+insert+line+to+file>`_

.. [5] `insert sting to n-th line of file · siongui/go-rst@3bdfb1a · GitHub <https://github.com/siongui/go-rst/commit/3bdfb1a66df7137ada01005cf17002f3d8f8b24b>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _0-indexed: https://en.wikipedia.org/wiki/Zero-based_numbering
.. _newline: https://en.wikipedia.org/wiki/Newline
.. _for range: https://golang.org/doc/effective_go.html#for
