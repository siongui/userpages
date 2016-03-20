[Golang] Example for block Action in Template package
#####################################################

:date: 2016-03-02T15:30+08:00
:tags: Go, Golang, Golang template
:category: Go
:summary: Example for `block action`_ in Go_ `text/template`_ and
          `html/template`_ packages.


Example for `block action`_ in Golang_ `text/template`_ and `html/template`_
packages.

`Run Code on Go Playground <https://play.golang.org/p/WfJ6Yha9Ew>`_

.. code-block:: go

  package main

  import (
          "os"
          "text/template"
  )

  var layout = `
  <!DOCTYPE html>
  <html>
  <head>
    <meta charset="utf-8">
    <title>{{block "title" .}}{{end}}</title>
  </head>
  <body>
  </body>
  </html>
  `

  var index = `
  {{define "title"}}My Site Name{{end}}
  `

  func main() {
          layoutTmpl, _ := template.New("layout").Parse(layout)
          indexTmpl, _ := layoutTmpl.Parse(index)
          indexTmpl.Execute(os.Stdout, nil)
  }


----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `Go 1.6 is released - The Go Blog <https://blog.golang.org/go1.6>`_

.. [2] `new template example program <https://github.com/golang/example/tree/master#template-godoc>`_

.. [3] `{{block}} action <https://golang.org/pkg/text/template/#hdr-Actions>`_

.. [4] `Templates - Go 1.6 Release Notes - The Go Programming Language <https://golang.org/doc/go1.6#template>`_

.. [5] `Template inheritance ? : golang <https://www.reddit.com/r/golang/comments/4b5wx5/template_inheritance/>`_

       `How to Use Template Blocks in Go 1.6 - Improving Efficiency with Technology | Joseph Spurrier <http://www.josephspurrier.com/how-to-use-template-blocks-in-go-1-6/>`_

       `Including html/template snippets: is there a better way? : golang <https://www.reddit.com/r/golang/comments/27ls5a/including_htmltemplate_snippets_is_there_a_better/>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _block action: https://golang.org/pkg/text/template/#hdr-Actions
.. _text/template: https://golang.org/pkg/text/template/
.. _html/template: https://golang.org/pkg/html/template/
