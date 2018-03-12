[Golang] Read JPEG Image Date via Exif Metadata
###############################################

:date: 2016-03-16T20:29+08:00
:tags: Go, Golang, File Input/Output
:category: Go
:summary: Read JPEG_ image date from Exif_ metadata in Go_ programming language.
:adsu: yes

Read JPEG_ image date from Exif_ metadata in Golang_.

We use goexif_ package to read Exif_ metadata. Install it by:

.. code-block:: bash

  $ go get -u github.com/rwcarlsen/goexif/exif

Source code:

.. show_github_file:: siongui userpages content/code/go-jpg-exif/date.go
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/go-jpg-exif/date_test.go

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `rwcarlsen/goexif: Decode embedded EXIF meta data from image files. <https://github.com/rwcarlsen/goexif>`_

.. [2] | `[Golang] Walk All Files in Directory <{filename}../../02/04/go-walk-all-files-in-directory%en.rst>`_
       | `List directory in go - Stack Overflow <http://stackoverflow.com/questions/14668850/list-directory-in-go>`_
.. adsu:: 3
.. [3] `golang read jpeg metadata <https://www.google.com/search?q=golang+read+jpeg+metadata>`_

.. [4] `python read jpeg metadata <https://www.google.com/search?q=python+read+jpeg+metadata>`_

.. [5] `Reading video stream metadata with go : golang <https://www.reddit.com/r/golang/comments/83mfkm/reading_video_stream_metadata_with_go/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Exif: https://www.google.com/search?q=EXIF
.. _JPEG: https://www.google.com/search?q=jpeg
.. _goexif: https://github.com/rwcarlsen/goexif
