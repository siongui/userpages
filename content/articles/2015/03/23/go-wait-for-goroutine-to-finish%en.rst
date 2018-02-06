[Golang] Wait For Goroutine to Finish
#####################################

:date: 2015-03-23 22:16
:tags: Go, Golang, Goroutine, Go Channels
:category: Go
:summary: Use channels to wait for all goroutines to finish
:og_image: https://cdn-images-1.medium.com/max/1200/1*GWYUFH14uOVLNHY-L1tv2w.jpeg
:adsu: yes


This post shows how to use channels_ to wait for all goroutines_ to finish.

.. show_github_file:: siongui userpages content/code/go-wait-goroutine-finish/sync.go
.. adsu:: 2

Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``.

----

References:

.. [1] `Waiting for all goroutines to finish before ending main - Google Groups <https://groups.google.com/d/topic/golang-nuts/mNhXnWLFOo4>`_

.. [2] `DuckDuckGo <https://duckduckgo.com/>`_ Search: `golang wait for goroutine to finish <https://duckduckgo.com/?q=golang+wait+for+goroutine+to+finish>`__

.. [3] `Google <https://www.google.com/>`_ Search: `golang wait for goroutine to finish <https://www.google.com/search?q=golang+wait+for+goroutine+to+finish>`__


.. _channels: https://tour.golang.org/concurrency/2

.. _goroutines: https://tour.golang.org/concurrency/1
