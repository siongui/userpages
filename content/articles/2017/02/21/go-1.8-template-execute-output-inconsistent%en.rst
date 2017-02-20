Go 1.8 Template Execute Output Inconsistent
###########################################

:date: 2017-02-21T02:24+08:00
:tags: Go, Golang, Golang template
:category: Go
:summary:  Bug in Go_ 1.8 `text/template`_ and `html/template`_ package
:og_image: http://www.unixstickers.com/image/cache/data/stickers/golang/Go-brown-side.sh-600x600.png
:adsu: yes


I found a bug in Go_ 1.8 template package (`text/template`_ & `html/template`_),
which output inconsistent HTML after ExecuteTemplate_. Sometimes I get correct
HTML output, and sometimes not. After I downgrade the Go version to 1.7.5,
everything works fine without problem.

The tempaltes (total 4 files) I use:

*theme/template/layout/layout.html*

.. code-block:: html

  {{- define "layout" -}}
  <!doctype html>
  <html prefix="og: http://ogp.me/ns#">
  <head>
    <meta charset="utf-8">
    <title>วัดป่าโพธิญาณ</title>
    <meta name="description" content="วัดป่าโพธิญาณ - สาขา 8 วัดหนองป่าพง">
    <meta name="keywords" content="วัดเขื่อน">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    {{template "metaog" .}}
  </head>
  <body>

  {{block "content" .}}{{end}}

  {{template "footer" .}}
  </body>
  </html>
  {{- end -}}

*theme/template/layout/includes/metaog.html*

.. code-block:: html

  {{- define "metaog" -}}
    <meta property="og:title" content="วัดป่าโพธิญาณ">
    <meta property="og:type" content="website">
    <meta property="og:description" content="วัดป่าโพธิญาณ - สาขา 8 วัดหนองป่าพง">
    <meta property="og:image" content="{{.OgImage}}">
    <meta property="og:url" content="{{.OgUrl}}">
    <meta property="og:locale" content="{{.OgLocale}}">
    <meta property="og:locale:alternate" content="en_US">
    <meta property="og:locale:alternate" content="zh_TW">
  {{- end -}}

*theme/template/layout/includes/footer.html*

.. code-block:: html

  {{- define "footer" -}}
  <div>Powered by
    <a href="https://golang.org/">Go</a>
  </div>
  {{- end -}}

*theme/template/index.html*

.. code-block:: html

  {{- template "layout" .}}
  {{define "content" -}}
  <div>Hello World!</div>
  <div>Template Inheritance in Go html/template</div>
  {{- end -}}

.. adsu:: 2

The correct HTML output should be:

.. code-block:: html

  <!doctype html>
  <html prefix="og: http://ogp.me/ns#">
  <head>
    <meta charset="utf-8">
    <title>วัดป่าโพธิญาณ</title>
    <meta name="description" content="วัดป่าโพธิญาณ - สาขา 8 วัดหนองป่าพง">
    <meta name="keywords" content="วัดเขื่อน">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta property="og:title" content="วัดป่าโพธิญาณ">
    <meta property="og:type" content="website">
    <meta property="og:description" content="วัดป่าโพธิญาณ - สาขา 8 วัดหนองป่าพง">
    <meta property="og:image" content="https://upload.wikimedia.org/wikipedia/commons/d/df/Dharma_Wheel.svg">
    <meta property="og:url" content="https://siongui.github.io/watpah/">
    <meta property="og:locale" content="th_TH">
    <meta property="og:locale:alternate" content="en_US">
    <meta property="og:locale:alternate" content="zh_TW">
  </head>
  <body>

  <div>Hello World!</div>
  <div>Template Inheritance in Go html/template</div>

  <div>Powered by
    <a href="https://golang.org/">Go</a>
  </div>
  </body>
  </html>

Sometime I get the following output in Go 1.8 template package:

.. code-block:: html

  <!doctype html>
  <html prefix="og: http://ogp.me/ns#">
  <head>
    <meta charset="utf-8">
    <title>วัดป่าโพธิญาณ</title>
    <meta name="description" content="วัดป่าโพธิญาณ - สาขา 8 วัดหนองป่าพง">
    <meta name="keywords" content="วัดเขื่อน">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta property="og:title" content="วัดป่าโพธิญาณ">
    <meta property="og:type" content="website">
    <meta property="og:description" content="วัดป่าโพธิญาณ - สาขา 8 วัดหนองป่าพง">
    <meta property="og:image" content="https://upload.wikimedia.org/wikipedia/commons/d/df/Dharma_Wheel.svg">
    <meta property="og:url" content="https://siongui.github.io/watpah/">
    <meta property="og:locale" content="th_TH">
    <meta property="og:locale:alternate" content="en_US">
    <meta property="og:locale:alternate" content="zh_TW">
  </head>
  <body>



  <div>Powered by
  <a href="https://golang.org/">Go</a>
  </div>
  </body>
  </html>

Interestingly, some one report that there is some problem with Go 1.8
sync/atomic package [2]_. Not sure whether it's the same issue or not?

----

.. adsu:: 3

References:

.. [1] `generate html from template · siongui/wat-pah-photiyan@a872d70 · GitHub <https://github.com/siongui/wat-pah-photiyan/commit/a872d70449eb143ef4b6fd8686100c5f49f42b8a>`_

.. [2] | `Cmd/compile: Go 1.8 regression: sync/atomic loop elided | Hacker News <https://news.ycombinator.com/item?id=13686863>`_
       | `cmd/compile: go1.8 regression: sync/atomic loop elided : golang <https://www.reddit.com/r/golang/comments/5v3jwv/cmdcompile_go18_regression_syncatomic_loop_elided/>`_

.. _Go: https://golang.org/
.. _html/template: https://golang.org/pkg/html/template/
.. _text/template: https://golang.org/pkg/text/template/
.. _ExecuteTemplate: https://golang.org/pkg/html/template/#Template.ExecuteTemplate
