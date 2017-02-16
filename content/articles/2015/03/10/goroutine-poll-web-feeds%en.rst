[Golang] Goroutines Poll Web Feeds
##################################

:date: 2015-03-10 21:59
:tags: Go, Golang, HTTP GET, Atom, RSS, Goroutine, Go net/http, Go time Package
:category: Go
:summary: Poll RSS/ATOM feeds with Goroutines
:adsu: yes


This post shows how to periodically poll (fetch by http get) `web feeds`_ with
goroutines_.

Souce Code
++++++++++

.. show_github_file:: siongui userpages content/code/goroutine-poll-feed/get.go
.. adsu:: 2


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

References:

.. [1] `Concurrency - A Tour of Go <http://tour.golang.org/concurrency/1>`_

.. [2] `http - The Go Programming Language <http://golang.org/pkg/net/http/>`_

.. [3] `go - Golang: Is it possible to capture a Ctrl+C signal and run a cleanup function, in a "defer" fashion? - Stack Overflow <http://stackoverflow.com/questions/11268943/golang-is-it-possible-to-capture-a-ctrlc-signal-and-run-a-cleanup-function-in>`_
.. adsu:: 3
.. [4] `Graceful stopping in Go <http://rcrowley.org/articles/golang-graceful-stop.html>`_
       (`Go 语言中实现优雅的停止程序 - 技术翻译 - 开源中国社区 <http://www.oschina.net/translate/golang-graceful-stop>`_)

.. [5] `Graceful server restart with Go <http://blog.scalingo.com/post/105609534953/graceful-server-restart-with-go>`_
       (`使用 Go 语言实现优雅的服务器重启 - 技术翻译 - 开源中国社区 <http://www.oschina.net/translate/graceful-server-restart-with-go>`_)


.. _goroutines: http://tour.golang.org/concurrency/1

.. _web feeds: http://en.wikipedia.org/wiki/Web_feed
