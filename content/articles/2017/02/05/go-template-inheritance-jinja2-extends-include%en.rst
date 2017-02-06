Golang Template Inheritance (Python Jinja2 extends & include)
#############################################################

:date: 2017-02-05T22:33+08:00
:tags: Go, Golang, html, Jinja2, Golang template
:category: Go
:summary: Template inheritance via Go `html/template`_ package. Show how one
          tempalate extends_ from another, and also include_ templates, just as
          how we do by Python_ Jinja2_.
:og_image: http://www.unixstickers.com/image/cache/data/stickers/golang/Go-brown-side.sh-600x600.png
:adsu: yes

.. contents::

Introduction
++++++++++++

This posts will show how to do `template inheritance`_ via Go_ official
`html/template`_ package. There are other template engines in Go, such as jigo_,
pongo2_, or mustache_, but I like to use official packages as much as possible.

We will introduce `template inheritance`_ by three steps:

1. Define a base *skeleton* template that contains all the common elements of
   your site and defines **blocks** that child templates can override.

2. include_ another templates in the base *skeleton* template.

3. Define a `child template`_ that extends_ the base *skeleton* template and
   output the final *HTML* we need.


Template Hierarchy
++++++++++++++++++

Assume we have 4 templates as follows:

.. code-block:: txt

  tmpl/index.html
  tmpl/layout/layout.html
  tmpl/layout/includes/metaog.html
  tmpl/layout/includes/footer.html

- ``layout.html`` is the `base template`_ (*skeleton*)
- ``metaog.html`` and ``footer.html`` will be included in ``layout.html``
- ``index.html`` is the `child template`_ that extends ``layout.html``

.. adsu:: 2

Look at ``layout.html`` (base *skeleton* template) first:

.. show_github_file:: siongui userpages content/code/go/template-inheritance/tmpl/layout/layout.html

In line 8 and 14 of ``layout.html``, two templates(*metaog* and *footer*) are
included via `template action`_, which is roughly  the same as the Jinja2_
include_ tag.

In line 12 of ``layout.html``, a **block** named *content* is declared and will
later be overridden by child template. The **block** here is also roughly the
same as the `Jinja2 blocks`_.

Next, look at the two tempaltes included in base *skeleton*:

.. show_github_file:: siongui userpages content/code/go/template-inheritance/tmpl/layout/includes/metaog.html
.. show_github_file:: siongui userpages content/code/go/template-inheritance/tmpl/layout/includes/footer.html

Nothing special above. The syntax is quite intuitive and easy to understand.

Finally, we will see how the child template (``index.html``) extends_ the base
template (``layout.html``) and override the *content* block in the base
template:

.. show_github_file:: siongui userpages content/code/go/template-inheritance/tmpl/index.html

The effect of `template action`_ in the first line is the same as the Jinja2
`extends`_ tag, and from line 2 to last line, the *content* block is defined and
override the declaration in the base template.

.. adsu:: 3


Template Rendering
++++++++++++++++++

The following *ParseTemplateDir* function reads all the above 4 templates under
``tmpl`` directory:

.. show_github_file:: siongui userpages content/code/go/template-inheritance/template.go
.. adsu:: 4

Define the template data and render ``index.html`` template to generate final
HTML output:

.. show_github_file:: siongui userpages content/code/go/template-inheritance/template_test.go

Final HTML output:

.. code-block:: txt

  === RUN   TestTemplateToHtml

  <!doctype html>
  <html prefix="og: http://ogp.me/ns#">
  <head>
    <meta charset="utf-8">
    <title>Theory and Practice</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <meta property="og:title" content="Theory and Practice">
    <meta property="og:url" content="https://siongui.github.io/">

  </head>
  <body>


  <div>Hello World!</div>
  <div>Template Inheritance in Go html/template</div>



  <div>Powered by
    <a href="https://golang.org/">Go</a>
  </div>

  </body>
  </html>


  --- PASS: TestTemplateToHtml (0.00s)
  PASS


----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.5``

----

.. adsu:: 5

References:

.. [1] `Golang html/template versus Python Jinja2 (6) - Template Inheritance (Extends) <{filename}../../../2015/03/08/python-jinja2-vs-go-html-template-extends%en.rst>`_

.. [2] `[Golang] Example for block Action in Template package <{filename}../../../2016/03/02/go-example-for-block-action-in-template%en.rst>`_

.. [3] `golang arguments dot - Google search <https://www.google.com/search?q=golang+arguments+dot>`_

       `golang arguments dot - DuckDuckGo search <https://duckduckgo.com/?q=golang+arguments+dot>`_

       `golang arguments dot - Bing search <https://www.bing.com/search?q=golang+arguments+dot>`_

       `golang arguments dot - Yahoo search <https://search.yahoo.com/search?p=golang+arguments+dot>`_

       `golang arguments dot - Baidu search <https://www.baidu.com/s?wd=golang+arguments+dot>`_

       `golang arguments dot - Yandex search <https://www.yandex.com/search/?text=golang+arguments+dot>`_

.. [4] `[Golang] Walk All Files in Directory <{filename}../../../2016/02/04/go-walk-all-files-in-directory%en.rst>`_

.. _Go: https://golang.org/
.. _html/template: https://golang.org/pkg/html/template/
.. _extends: http://jinja.pocoo.org/docs/latest/templates/#extends
.. _include: http://jinja.pocoo.org/docs/latest/templates/#include
.. _Jinja2: http://jinja.pocoo.org/docs/dev/
.. _Python: https://www.python.org/
.. _template inheritance: http://jinja.pocoo.org/docs/latest/templates/#template-inheritance
.. _jigo: https://github.com/jmoiron/jigo
.. _pongo2: https://github.com/flosch/pongo2
.. _mustache: https://github.com/hoisie/mustache
.. _child template: http://jinja.pocoo.org/docs/latest/templates/#child-template
.. _base template: http://jinja.pocoo.org/docs/latest/templates/#base-template
.. _open graph metadata: http://ogp.me/
.. _template action: https://golang.org/pkg/text/template/#hdr-Actions
.. _Jinja2 blocks: http://jinja.pocoo.org/docs/latest/templates/#blocks
