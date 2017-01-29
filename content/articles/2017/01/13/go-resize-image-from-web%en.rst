[Golang] Resize Image From Web
##############################

:date: 2017-01-13T21:03+08:00
:tags: Go, Golang, Go net/http, File Input/Output, Go image/png Package
:category: Go
:summary: Read an image from web and resize it in Go_ programming language.
:adsu: yes


This post gives an example that read an png_ image from website (use Google log
as example) and resize it in half.

Install the image resizing library (`github.com/nfnt/resize`_) first:

.. code-block:: bash

  $ go get github.com/nfnt/resize

The full example:

.. adsu:: 2
.. show_github_file:: siongui userpages content/code/go/resize-image/web.go
.. adsu:: 3

For more usage of `github.com/nfnt/resize`_, visit its GitHub page. If your image
is of JPEG_ format, import `image/jpeg`_ and modify the code correspondingly.

----

Source code tested on: ``Ubuntu Linux 16.10``, ``Go 1.7.4``.

----

References:

.. [1] `golang change image size - Google search <https://www.google.com/search?q=golang+change+image+size>`_

       `golang change image size - DuckDuckGo search <https://duckduckgo.com/?q=golang+change+image+size>`_

       `golang change image size - Bing search <https://www.bing.com/search?q=golang+change+image+size>`_

       `golang change image size - Yahoo search <https://search.yahoo.com/search?p=golang+change+image+size>`_

       `golang change image size - Baidu search <https://www.baidu.com/s?wd=golang+change+image+size>`_

       `golang change image size - Yandex search <https://www.yandex.com/search/?text=golang+change+image+size>`_

.. [2] `GitHub - nfnt/resize: Pure golang image resizing <https://github.com/nfnt/resize>`_

       `png - The Go Programming Language <https://golang.org/pkg/image/png/>`_

       `http - The Go Programming Language <https://golang.org/pkg/net/http/>`_

.. [3] `Go Resizing Images - Stack Overflow <http://stackoverflow.com/questions/22940724/go-resizing-images>`_

.. [4] `GitHub - fawick/speedtest-resize: Compare various Image resize algorithms for the Go language <https://github.com/fawick/speedtest-resize>`_

.. [5] `Guess Metadata from HTML and Converted to reStructuredText <{filename}../../../2016/05/16/html-metadata-to-rst%en.rst>`_


.. _Go: https://golang.org/
.. _png: https://www.google.com/search?q=png
.. _JPEG: https://www.google.com/search?q=JPEG
.. _image/jpeg: https://golang.org/pkg/image/jpeg/
.. _github.com/nfnt/resize: https://github.com/nfnt/resize
