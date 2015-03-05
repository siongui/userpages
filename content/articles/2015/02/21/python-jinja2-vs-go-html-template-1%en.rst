Golang html/template versus Python Jinja2 (1)
#############################################

:date: 2015-02-21 02:51
:modified: 2015-03-05 14:36
:tags: Go, Golang, html, Jinja2, Python
:category: Go
:summary: Comparison of Python Jinja2 and Go html/template, side by side for
          easy reference - Load Templates

|
|
|

.. list-table:: Go html/template versue Python Jinja2 - Load Templates
   :header-rows: 1
   :class: table-syntax-diff

   * - Go `html/template`_
     - Python Jinja2_

   * - .. code-block:: go

         import "html/template"

     - .. code-block:: python

         import jinja2

   * - .. code-block:: go

         // assume template and go source code are in the same directory.

     - .. code-block:: python

         # Tell Jinja2 where the template folder is
         # Template files and Python scripts are in the same directory in this example.
         import os

         JINJA_ENVIRONMENT = jinja2.Environment(
             loader=jinja2.FileSystemLoader(os.path.dirname(__file__)),
             extensions=['jinja2.ext.autoescape'],
             autoescape=True)

   * - .. code-block:: go

         t, _ := template.ParseFiles("index.html")

     - .. code-block:: python

         template = JINJA_ENVIRONMENT.get_template('index.html')

   * - in *index.html*:

       .. code-block:: html

         Hello {{ . }}

     - in *index.html*:

       .. code-block:: html

         Hello {{ name }}

   * - setup tempalate value in Go:

       .. code-block:: go

         template_values := "World"

     - setup tempalate value in Python:

       .. code-block:: python

         template_values = {'name': 'World'}

   * - use go `net/http`_ to serve webpage

       .. code-block:: go

         // w http.ResponseWriter
         t.Execute(w, template_values)

     - use webapp2_ to serve webpage

       .. code-block:: python

         # self.response.write() => write to client browser
         self.response.write(template.render(template_values))

       use web.py_ to serve webpage

       .. code-block:: python

         return template.render(template_values)


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``, ``Python 2.7.8``, ``Google App Engine Python SDK 1.9.18``

----

*Golang html/template versus Python Jinja2* series:

.. [1] `Golang html/template versus Python Jinja2 (1) <{filename}python-jinja2-vs-go-html-template-1%en.rst>`_

.. [2] `Golang html/template versus Python Jinja2 (2) <{filename}../24/python-jinja2-vs-go-html-template-2%en.rst>`_

----

References:

.. [a] `html/template - Golang <http://golang.org/pkg/html/template/>`_

.. [b] `Jinja2 (The Python Template Engine) <http://jinja.pocoo.org/>`_

.. [c] `jinja2.FileSystemLoader <http://jinja.pocoo.org/docs/dev/api/#jinja2.FileSystemLoader>`_

.. [d] `Using Templates - Google App Engine for Python <https://cloud.google.com/appengine/docs/python/gettingstartedpython27/templates>`_

.. [e] `Go HTML Templates: Not Jinja2 <http://blog.ojrac.com/go-html-templates-not-jinja2.html>`_

.. [f] `Go HTML Templates: Applying Data <http://blog.ojrac.com/go-html-templates-applying-data.html>`_

.. [g] `Go HTML Templates: Functions and Flow <http://blog.ojrac.com/go-html-templates-functions-and-flow.html>`_

.. [h] `pongo2: Django-syntax like template-engine for Go <https://github.com/flosch/pongo2>`_

.. [i] `go语言：优雅的模板切割技术 <http://studygolang.com/articles/2315>`_


.. _html/template: http://golang.org/pkg/html/template/

.. _webapp2: https://cloud.google.com/appengine/docs/python/tools/webapp2

.. _Jinja2: http://jinja.pocoo.org/docs/dev/

.. _net/http: http://golang.org/pkg/net/http/

.. _web.py: http://webpy.org/
