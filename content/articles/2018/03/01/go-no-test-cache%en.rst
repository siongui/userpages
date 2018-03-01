[Golang] Disable Test Cache
###########################

:date: 2018-03-01T18:46+08:00
:tags: Go, Golang, Commandline, Go testing
:category: Go
:summary: Disable caching in Go test.
:og_image: https://newrelic.com/assets/pages/golang/go-mascot.svg
:adsu: yes


Not sure when Go starts to cache the test results. When I tried to run test with
HTTP requests, the results are all the same after successful test. After some
Google searches, I found `the answer in the reply`_ of the issue [1]_. The test
caching can be disabled by the following setting:

  **GOCACHE=off**

If you use command line, you can run the following command:

.. code-block:: bash

  $ export GOCACHE=off

Or add the above command in the Makefile if you use *make* in your workflow.

.. adsu:: 2

----

Tested on: ``Ubuntu Linux 17.10``, ``Go 1.10``.

----

References:

.. [1] `cmd/dist: no obvious way to disable test caching · Issue #22758 · golang/go · GitHub <https://github.com/golang/go/issues/22758>`_

.. _the answer in the reply: https://github.com/golang/go/issues/22758#issuecomment-344921828
