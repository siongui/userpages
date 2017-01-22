[Golang] Parse Command Line Arguments - String Variable
#######################################################

:date: 2016-12-21T22:43+08:00
:tags: Go, Golang, Commandline, Go flag Package, Makefile
:category: Go
:summary: How to parse command-line arguments (string variable) via flag_
          package in Go_ programming language.
:adsu: yes


Pass argument from `command line`_ and used as *string* variable in Go_ program:

.. show_github_file:: siongui userpages content/code/go/command-line-argument/Makefile

.. show_github_file:: siongui userpages content/code/go/command-line-argument/demo.go

Don't forget to call `flag.Parse()`_ in your code. If you do, your variable will
always be default value.


Source code tested on: ``Ubuntu Linux 16.10``, ``Go 1.7.4``.

----

References:

.. [1] `go pass commandline argument - Google search <https://www.google.com/search?q=go+pass+commandline+argument>`_

       `go pass commandline argument - DuckDuckGo search <https://duckduckgo.com/?q=go+pass+commandline+argument>`_

       `go pass commandline argument - Bing search <https://www.bing.com/search?q=go+pass+commandline+argument>`_

       `go pass commandline argument - Yahoo search <https://search.yahoo.com/search?p=go+pass+commandline+argument>`_

       `go pass commandline argument - Baidu search <https://www.baidu.com/s?wd=go+pass+commandline+argument>`_

       `go pass commandline argument - Yandex search <https://www.yandex.com/search/?text=go+pass+commandline+argument>`_

.. [2] `flag - The Go Programming Language <https://golang.org/pkg/flag/>`_

.. [3] `Go by Example: Command-Line Flags <https://gobyexample.com/command-line-flags>`_

.. [4] `[Golang] Read Command-line Arguments Example <{filename}../../../2015/02/18/go-parse-command-line-arguments%en.rst>`_


.. _Go: https://golang.org/
.. _flag: https://golang.org/pkg/flag/
.. _command line: https://www.google.com/search?q=command+line
.. _flag.Parse(): https://golang.org/pkg/flag/#Parse
