[Golang] Download HTML From URL
###############################

:date: 2016-03-19T03:05+08:00
:tags: Go, Golang, HTTP GET, Commandline, File Input/Output, Go flag Package,
       html, Go net/http
:category: Go
:summary: Download and save HTML file from given URL via Go_. Do nothing if the
          HTML file already locally exists.


Download and save HTML file from given URL via Golang_. Do nothing if the HTML
file already locally exists.

.. show_github_file:: siongui userpages content/code/go-save-url-html/download.go

.. show_github_file:: siongui userpages content/code/go-save-url-html/Makefile

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `golang if file exists <https://www.google.com/search?q=golang+if+file+exists>`_

       `How to check if a file exists in Go? - Stack Overflow <http://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go>`_

.. [2] `http - The Go Programming Language <https://golang.org/pkg/net/http/>`_

.. [3] `golang copy io reader <https://www.google.com/search?q=golang+copy+io+reader>`_

       `func Copy - io - The Go Programming Language <https://golang.org/pkg/io/#Copy>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
