Golang html/template versus Python Jinja2 (6) - Template Inheritance (Extends)
##############################################################################

:date: 2015-03-08 22:03
:tags: Go, Golang, html, Jinja2, Python, Golang template
:category: Go
:summary: Comparison of Python Jinja2 and Go html/template, side by side for
          easy reference - Mimic Jinja2 template inheritance in Go html/template


Template inheritance is a powerful feature of Jinja2_. Here we show how to mimic
the template *extends* of Jinja2 in Go `html/template`_.

|
|

.. list-table:: Go html/template versue Python Jinja2 - Arrays and Slices Index
   :header-rows: 1
   :class: table-syntax-diff

   * - Go `html/template`_
     - Python Jinja2_

   * - "base" template:

       .. code-block:: go

         {{define "base"}}
         <!doctype html>
         <html>
         <head>
           <meta charset="utf-8">
           <title>{{template "title" .}}</title>
         </head>
         <body>
         {{template "content" .}}
         </body>
         </html>
         {{end}}

     - "base" template:

       .. code-block:: python

         <!doctype html>
         <html>
         <head>
           <meta charset="utf-8">
           <title>{% block title %}{% endblock %}</title>
         </head>
         <body>
         {% block content %}{% endblock %}
         </body>
         </html>

   * - index.html:

       .. code-block:: go

         {{template "base" .}}
         {{define "title"}}my title{{end}}
         {{define "content"}}
         <div>hello {{.}}</div>
         {{end}}

     - index.html:

       .. code-block:: python

         {% extends "base.html" %}
         {% block title %}my title{% endblock %}
         {% block content %}
         <div>hello {{ name }}</div>
         {% endblock %}

Complete *Go html/template* source code:

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/extends/base-go.html

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/extends/index-go.html

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/extends/extends.go

Complete *Python Jinja2* source code:

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/extends/base.html

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/extends/index.html

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/extends/extends.py


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``, ``Python 2.7.8``, ``Jinja2 2.7.3``

----

*Golang html/template versus Python Jinja2* series:

.. [1] `Golang html/template versus Python Jinja2 (1) <{filename}../../02/21/python-jinja2-vs-go-html-template-1%en.rst>`_

.. [2] `Golang html/template versus Python Jinja2 (2) <{filename}../../02/24/python-jinja2-vs-go-html-template-2%en.rst>`_

.. [3] `Golang html/template versus Python Jinja2 (3) - Arrays and Slices <{filename}../05/python-jinja2-vs-go-html-template-array-slice%en.rst>`_

.. [4] `Golang html/template versus Python Jinja2 (4) - Arrays and Slices Index <{filename}../06/python-jinja2-vs-go-html-template-array-slice-index%en.rst>`_

.. [5] `Golang html/template versus Python Jinja2 (5) - Maps and Dictionaries <{filename}../07/python-jinja2-vs-go-html-template-map-dictionary%en.rst>`_

.. [6] `Golang html/template versus Python Jinja2 (6) - Template Inheritance (Extends) <{filename}python-jinja2-vs-go-html-template-extends%en.rst>`_

.. [7] `Golang html/template versus Python Jinja2 (7) - Custom Functions and Filters <{filename}../12/python-jinja2-vs-go-html-template-function-and-filter%en.rst>`_

----

References:

.. [a] `Template Inheritance - Jinja2 Documentation <http://jinja.pocoo.org/docs/dev/templates/#template-inheritance>`_

.. [b] `Nested template definitions - template - The Go Programming Language <http://golang.org/pkg/text/template/#hdr-Nested_template_definitions>`_

.. [c] `go语言：优雅的模板切割技术 <http://studygolang.com/articles/2315>`_

.. [d] `Template inheritance ? : golang <https://www.reddit.com/r/golang/comments/4b5wx5/template_inheritance/>`_

       `How to Use Template Blocks in Go 1.6 - Improving Efficiency with Technology | Joseph Spurrier <http://www.josephspurrier.com/how-to-use-template-blocks-in-go-1-6/>`_

       `Including html/template snippets: is there a better way? : golang <https://www.reddit.com/r/golang/comments/27ls5a/including_htmltemplate_snippets_is_there_a_better/>`_


.. _html/template: http://golang.org/pkg/html/template/

.. _Jinja2: http://jinja.pocoo.org/docs/dev/
