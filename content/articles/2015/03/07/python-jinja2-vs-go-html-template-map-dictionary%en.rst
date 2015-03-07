Golang html/template versus Python Jinja2 (5) - Maps and Dictionaries
#####################################################################

:date: 2015-03-07 20:51
:tags: Go, Golang, html, Jinja2, Python
:category: Go
:summary: Comparison of Python Jinja2 and Go html/template, side by side for
          easy reference - Loop through Python Dictionaries vs Go Maps


This post shows how to loop through the maps_ in Go_ `html/template`_ and
dictionaries_ in Python_ Jinja2_.

In Jinja2_, ``loop.index`` or ``loop.index0`` is used to access the loop index,
starting from **1** or **0**. (see [a]_)

In Go `html/template`_, it seems that there is no way to access the loop index
while loop through the variable of ``map`` type. (see [b]_)

|
|

.. list-table:: Go html/template versue Python Jinja2 - Arrays and Slices Index
   :header-rows: 1
   :class: table-syntax-diff

   * - Go `html/template`_
     - Python Jinja2_

   * - template:

       .. code-block:: go

         {{range $name, $href := .}}
         <a href="{{$href}}">{{$name}}</a>
         {{end}}

     - template:

       .. code-block:: python

         {% for name, href in links.iteritems() %}
         <a href="{{ href }}">{{ name }}</a>
         {% endfor %}

       loop index starting from 1:

       .. code-block:: python

         {% for name, href in links.iteritems() %}
         {{ loop.index }}: <a href="{{ href }}">{{ name }}</a>
         {% endfor %}

       loop index starting from 0:

       .. code-block:: python

         {% for name, href in links.iteritems() %}
         {{ loop.index0 }}: <a href="{{ href }}">{{ name }}</a>
         {% endfor %}

   * - template values:

       .. code-block:: go

         var m = map[string]string{
             "Google": "https://www.google.com/",
             "Facebook": "https://www.facebook.com/",
         }

     - template values:

       .. code-block:: python

         links = {
           'Google': 'https://www.google.com',
           'Facebook': 'https://www.facebook.com',
         }

Complete *Go html/template* source code:

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/html-template-example-4.go

Complete *Python Jinja2* source code:

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/jinja2-example-4.py


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``, ``Python 2.7.8``, ``Jinja2 2.7.3``

----

*Golang html/template versus Python Jinja2* series:

.. [1] `Golang html/template versus Python Jinja2 (1) <{filename}../../02/21/python-jinja2-vs-go-html-template-1%en.rst>`_

.. [2] `Golang html/template versus Python Jinja2 (2) <{filename}../../02/24/python-jinja2-vs-go-html-template-2%en.rst>`_

.. [3] `Golang html/template versus Python Jinja2 (3) - Arrays and Slices <{filename}../05/python-jinja2-vs-go-html-template-array-slice%en.rst>`_

.. [4] `Golang html/template versus Python Jinja2 (4) - Arrays and Slices Index <{filename}../06/python-jinja2-vs-go-html-template-array-slice-index%en.rst>`_

.. [5] `Golang html/template versus Python Jinja2 (5) - Maps and Dictionaries <{filename}python-jinja2-vs-go-html-template-map-dictionary%en.rst>`_

----

References:

.. [a] `For - List of Control Structures - Jinja2 Documentation <http://jinja.pocoo.org/docs/dev/templates/#for>`_

.. [b] `variables - template - The Go Programming Language <http://golang.org/pkg/text/template/#hdr-Variables>`_


.. _html/template: http://golang.org/pkg/html/template/

.. _Jinja2: http://jinja.pocoo.org/docs/dev/

.. _Go: https://golang.org/

.. _Python: https://www.python.org/

.. _dictionaries: https://docs.python.org/2/tutorial/datastructures.html#dictionaries

.. _maps: https://tour.golang.org/moretypes/15
