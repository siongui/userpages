[Golang] Get Filename Without Extension
#######################################

:date: 2018-02-25T10:06+08:00
:tags: Go, Golang, String Manipulation, Go strings Package
:category: Go
:summary: Get file name without extension in Go programming language.
:og_image: https://d2d42mpnbqmzj3.cloudfront.net/images/stories/doc-excel/extract-file-extension/doc-extract-file-extensions-4.png
:adsu: yes

Use Go standard library to get file name without extension.

- path.Ext_ method to get filename extension
- strings.TrimSuffix_ method to remove the extension from the filename.

.. code-block:: go

  import (
  	"path"
  	"strings"
  )

  func FilenameWithoutExtension(fn string) string {
  	return strings.TrimSuffix(fn, path.Ext(fn))
  }

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/-eLu5l2T2dA>`__
   :class: align-center

.. adsu:: 2

----

Tested on:

- ``Ubuntu Linux 17.10``, ``Go 1.10``.
- `The Go Playground`_

.. _path.Ext: https://golang.org/pkg/path/#Ext
.. _strings.TrimSuffix: https://golang.org/pkg/strings/#TrimSuffix
.. _The Go Playground: https://play.golang.org/
