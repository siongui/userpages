Golang html/template versus Python Jinja2 (1)
#############################################

:date: 2015-02-21 02:51
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

   * - Go
     - Python

   * - .. code-block:: go

         import "html/template"

     - .. code-block:: python

         import jinja2

   * - .. code-block:: go

         // assume template and go source code are in the same directory.

     - .. code-block:: python

         # Tell Jinja2 where the template folder is
         # Template files and Python scripts are in the same directory in this example.
         JINJA_ENVIRONMENT = jinja2.Environment(
             loader=jinja2.FileSystemLoader(os.path.dirname(__file__)),
             extensions=['jinja2.ext.autoescape'],
             autoescape=True)

   * - .. code-block:: go

         t, _ := template.ParseFiles("index.html")

     - .. code-block:: python

         template = JINJA_ENVIRONMENT.get_template('index.html')

   * - .. code-block:: go

         // w => http.ResponseWriter
         t.Execute(w, template_values)

     - .. code-block:: python

         # self.response.write() => write to client browser
         self.response.write(template.render(template_values))


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``, ``Python 2.7.8``

----

*Golang html/template versus Python Jinja2* series:

.. [1] `Golang html/template versus Python Jinja2 (1) <{filename}python-jinja2-vs-go-html-template-1%en.rst>`_

.. [2] `Golang html/template versus Python Jinja2 (2) <{filename}../24/python-jinja2-vs-go-html-template-2%en.rst>`_

----

References:

.. [a] `html/template - Golang <http://golang.org/pkg/html/template/>`_

.. [b] `Jinja2 (The Python Template Engine) <http://jinja.pocoo.org/>`_

.. [c] `Using Templates - Google App Engine for Python <https://cloud.google.com/appengine/docs/python/gettingstartedpython27/templates>`_

.. [d] `Go HTML Templates: Not Jinja2 <http://blog.ojrac.com/go-html-templates-not-jinja2.html>`_

.. [e] `Go HTML Templates: Applying Data <http://blog.ojrac.com/go-html-templates-applying-data.html>`_

.. [f] `Go HTML Templates: Functions and Flow <http://blog.ojrac.com/go-html-templates-functions-and-flow.html>`_

.. [g] `pongo2: Django-syntax like template-engine for Go <https://github.com/flosch/pongo2>`_

.. [h] `go语言：优雅的模板切割技术 <http://studygolang.com/articles/2315>`_
