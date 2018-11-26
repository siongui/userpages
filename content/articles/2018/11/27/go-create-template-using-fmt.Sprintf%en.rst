[Golang] Create Template Using fmt.Sprintf
##########################################

:date: 2018-11-27T05:22+08:00
:tags: Go, Golang, Golang template, reStructuredText
:category: Go
:summary: Create short pieces of templates using *fmt.Sprintf* in Go.
:og_image: http://program.okitama.org/img/gopherswrench.jpg
:adsu: yes


In Go_, there are `text/template`_ and `html/template`_ packgaes for templating.
But for short pieces of templates, I think fmt.Sprintf_ is easier to use.

For example, I want to create `reStructuredText image`_ code template according
to image url. We can use fmt.Sprintf_ as follows:

.. rubric:: `Run Code on Go Playground <https://play.golang.org/p/IPrz4EgskEJ>`__
   :class: align-center

.. code-block:: go

  package main

  import (
  	"fmt"
  )

  func rstImage(url, alt string) string {
  	rst := ""
  	rst += fmt.Sprintf(".. image:: %s\n", url)
  	rst += fmt.Sprintf("   :alt: %s\n", alt)
  	return rst
  }

  func main() {
  	rstimg := rstImage("https://golang.org/doc/gopher/pencil/gopherswrench.jpg", "gopherswrench")
  	fmt.Println(rstimg)
  }

As you can see, the code is still easy to read and modify, without the need to
define struct for the standard templating packages.

----

.. adsu:: 2

References:

.. [1] `fmt - The Go Programming Language <https://golang.org/pkg/fmt/>`_

.. _Go: https://golang.org/
.. _text/template: https://golang.org/pkg/text/template/
.. _html/template: https://golang.org/pkg/html/template/
.. _fmt.Sprintf: https://golang.org/pkg/fmt/#Sprintf
.. _reStructuredText image: http://docutils.sourceforge.net/docs/ref/rst/directives.html#images
