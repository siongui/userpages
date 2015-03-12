Golang html/template versus Python Jinja2 (4) - Arrays and Slices Index
#######################################################################

:date: 2015-03-06 21:48
:tags: Go, Golang, html, Jinja2, Python
:category: Go
:summary: Comparison of Python Jinja2 and Go html/template, side by side for
          easy reference - Python List vs Go Arrays and Slices - Loop Index


When loop through the arrays/slices (Go_) or lists (Python_) in templates,
sometimes the loop index is needed. This post shows how to access loop index in
Python_ Jinja2_ and Go_ `html/template`_.

In Jinja2_, ``loop.index`` or ``loop.index0`` is used to access the loop index,
starting from **1** or **0**. (see [a]_)

In Go `html/template`_, one more variable is declared in ``range`` action to
access loop index, starting from **0**. (see [b]_)

|
|

.. list-table:: Go html/template versue Python Jinja2 - Arrays and Slices Index
   :header-rows: 1
   :class: table-syntax-diff

   * - Go `html/template`_
     - Python Jinja2_

   * - template:

       .. code-block:: go

         {{range $index, $link := .}}
         {{$index}}: <a href="{{$link.Href}}">{{$link.Name}}</a>
         {{end}}

     - template:

       .. code-block:: python

         {# index start from 0 #}
         {% for link in links %}
         {{loop.index0}}: <a href="{{link.href}}">{{link.name}}</a>
         {% endfor %}

       |

       .. code-block:: python

         {# index start from 1 #}
         {% for link in links %}
         {{loop.index}}: <a href="{{link.href}}">{{link.name}}</a>
         {% endfor %}

   * - template values:

       .. code-block:: go

         type Link struct {
                Name    string
                Href    string
         }

         // arrays
         var la [2]Link
         la[0] = Link{"Google", "https://www.google.com/"}
         la[1] = Link{"Facebook", "https://www.facebook.com/"}

         // slices
         var ls []Link
         ls = append(ls, Link{"Google", "https://www.google.com/"})
         ls = append(ls, Link{"Facebook", "https://www.facebook.com/"})

     - template values:

       .. code-block:: python

         links = [
           {'name': 'Google', 'href': 'https://www.google.com'},
           {'name': 'Facebook', 'href': 'https://www.facebook.com'}
         ]

Complete *Go html/template* source code:

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/html-template-example-3.go

Complete *Python Jinja2* source code:

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/jinja2-example-3.py


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``, ``Python 2.7.8``, ``Jinja2 2.7.3``

----

*Golang html/template versus Python Jinja2* series:

.. [1] `Golang html/template versus Python Jinja2 (1) <{filename}../../02/21/python-jinja2-vs-go-html-template-1%en.rst>`_

.. [2] `Golang html/template versus Python Jinja2 (2) <{filename}../../02/24/python-jinja2-vs-go-html-template-2%en.rst>`_

.. [3] `Golang html/template versus Python Jinja2 (3) - Arrays and Slices <{filename}../05/python-jinja2-vs-go-html-template-array-slice%en.rst>`_

.. [4] `Golang html/template versus Python Jinja2 (4) - Arrays and Slices Index <{filename}python-jinja2-vs-go-html-template-array-slice-index%en.rst>`_

.. [5] `Golang html/template versus Python Jinja2 (5) - Maps and Dictionaries <{filename}../07/python-jinja2-vs-go-html-template-map-dictionary%en.rst>`_

.. [6] `Golang html/template versus Python Jinja2 (6) - Template Inheritance (Extends) <{filename}../08/python-jinja2-vs-go-html-template-extends%en.rst>`_

.. [7] `Golang html/template versus Python Jinja2 (7) - Custom Functions and Filters <{filename}../12/python-jinja2-vs-go-html-template-function-and-filter%en.rst>`_

----

References:

.. [a] `For - List of Control Structures - Jinja2 Documentation <http://jinja.pocoo.org/docs/dev/templates/#for>`_

.. [b] `variables - template - The Go Programming Language <http://golang.org/pkg/text/template/#hdr-Variables>`_


.. _html/template: http://golang.org/pkg/html/template/

.. _Jinja2: http://jinja.pocoo.org/docs/dev/

.. _Go: https://golang.org/

.. _Python: https://www.python.org/
