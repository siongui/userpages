[Golang] Capture and Handle Ctrl+C Event
########################################

:date: 2015-03-11 20:59
:tags: Go, Golang, Signal
:category: Go
:summary: Run cleanup function when receiving ^C signal.


For database programs, we need to close the database if users press `Ctrl+C`_
to terminate the program. This post shows how to capture Ctrl+C event and run
the handler in Go_.


Souce Code
++++++++++

.. show_github_file:: siongui userpages content/code/go-ctrl+c/ctrl+c.go



Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

References:

.. [1] `go - Golang: Is it possible to capture a Ctrl+C signal and run a cleanup function, in a "defer" fashion? - Stack Overflow <http://stackoverflow.com/questions/11268943/golang-is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in>`_

.. [2] `Go by Example: Signals <https://gobyexample.com/signals>`_

.. [3] `Channels - A Tour of Go <https://tour.golang.org/concurrency/2>`_


.. _Ctrl+C: http://en.wikipedia.org/wiki/Control-C

.. _Go: https://golang.org/
