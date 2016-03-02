Golang html/template versus Python Jinja2 (7) - Custom Functions and Filters
############################################################################

:date: 2015-03-12 20:15
:tags: Go, Golang, html, Jinja2, Python, Golang template
:category: Go
:summary: Comparison of Python Jinja2 and Go html/template, side by side for
          easy reference - Go template custom function vs Jinja2 custom filter
          and function.

The custom function in Go `html/template`_ is similar to Jinja2_ custom filters
in terms of syntax and functionality. The custom function can also be used in
Jinja2 template (see [d]_). This post will compare custom functions in Go
`html/template`_ with custom filters/functions in Jinja2_.

|
|

.. list-table:: Go html/template versue Python Jinja2 - Arrays and Slices Index
   :header-rows: 1
   :class: table-syntax-diff

   * - Go `html/template`_
     - Python Jinja2_

   * - custom function *gettext*

       .. code-block:: go

         func gettext(s string) string {
                 if s == "world" {
                         return "世界"
                 }
                 return s
         }

     - custom function/filter *gettext*

       .. code-block:: python

         def gettext(s):
           if s == "world":
             return u"世界"
           return s

   * - custom function in template:

       .. code-block:: go

         <span>hello {{gettext .}}</span>

       use custom function with syntax similar to Jinja2 filter:

       .. code-block:: go

         <span>hello {{. | gettext}}</span>

     - custom function in Jinja2 template: (see [d]_)

       .. code-block:: python

         <span>hello {{gettext(name)}}</span>

       custom filter in Jinja2 template: (see [e]_ [f]_)

       .. code-block:: python

         <span>hello {{name | gettext}}</span>

   * - load custom function and output: (see [a]_)

       .. code-block:: go

         var funcMap = template.FuncMap{
                "gettext": gettext,
         }

         t, _ := template.New("foo").Funcs(funcMap).Parse(tmpl)
         s := "world"
         t.Execute(os.Stdout, s)

     - load custom function and output:

       .. code-block:: python

         t = Template(tmpl)
         t.globals['gettext'] = gettext

         sys.stdout.write(t.render(name="world"))

       load custom filter and output:

       .. code-block:: python

         env = Environment()
         env.filters['gettext'] = gettext
         t = env.from_string(tmpl)

         sys.stdout.write(t.render(name="world"))


Complete *Go html/template* source code:

*Custom Function* used in template:

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/html-template-example-5.go

Complete *Python Jinja2* source code:

*Custom Function* used in template:

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/jinja2-example-5.py

*Custom Filter* used in template:

.. show_github_file:: siongui userpages content/code/python-jinja2-vs-go-html-template/jinja2-example-5_1.py


Tested on: ``Ubuntu Linux 14.10``, ``Go 1.4``, ``Python 2.7.8``, ``Jinja2 2.7.3``

----

*Golang html/template versus Python Jinja2* series:

.. [1] `Golang html/template versus Python Jinja2 (1) <{filename}../../02/21/python-jinja2-vs-go-html-template-1%en.rst>`_

.. [2] `Golang html/template versus Python Jinja2 (2) <{filename}../../02/24/python-jinja2-vs-go-html-template-2%en.rst>`_

.. [3] `Golang html/template versus Python Jinja2 (3) - Arrays and Slices <{filename}../05/python-jinja2-vs-go-html-template-array-slice%en.rst>`_

.. [4] `Golang html/template versus Python Jinja2 (4) - Arrays and Slices Index <{filename}../06/python-jinja2-vs-go-html-template-array-slice-index%en.rst>`_

.. [5] `Golang html/template versus Python Jinja2 (5) - Maps and Dictionaries <{filename}../07/python-jinja2-vs-go-html-template-map-dictionary%en.rst>`_

.. [6] `Golang html/template versus Python Jinja2 (6) - Template Inheritance (Extends) <{filename}../08/python-jinja2-vs-go-html-template-extends%en.rst>`_

.. [7] `Golang html/template versus Python Jinja2 (7) - Custom Functions and Filters <{filename}python-jinja2-vs-go-html-template-function-and-filter%en.rst>`_

----

References:

.. [a] `go - Template and custom function; panic: function not defined - Stack Overflow <http://stackoverflow.com/questions/17843311/template-and-custom-function-panic-function-not-defined>`_

.. [b] `TechnoSophos: Using Custom Template Functions in Go <http://technosophos.com/2013/11/23/using-custom-template-functions-in-go.html>`_

.. [c] Google Search: `go template function <https://www.google.com/search?q=go+template+function>`_

.. [d] `Call a python function from jinja2 - Stack Overflow <http://stackoverflow.com/questions/6036082/call-a-python-function-from-jinja2>`_

.. [e] `Custom Filters - API - Jinja2 Documentation <http://jinja.pocoo.org/docs/dev/api/#custom-filters>`_

.. [f] `google app engine - Adding a custom filter to jinja2 on GAE - Stack Overflow <http://stackoverflow.com/questions/12464095/adding-a-custom-filter-to-jinja2-on-gae>`_


.. _html/template: http://golang.org/pkg/html/template/

.. _Jinja2: http://jinja.pocoo.org/docs/dev/
