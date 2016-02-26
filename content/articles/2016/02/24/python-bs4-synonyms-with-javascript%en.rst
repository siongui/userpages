[Python] Beautiful Soup 4 Synonyms with JavaScript
##################################################

:date: 2016-02-24T15:16+08:00
:tags: Python, Beautiful Soup, JavaScript
:category: Python
:summary: Synonyms - Python_ `Beautiful Soup 4`_ translated to JavaScript_.

Synonyms - Python_ `Beautiful Soup 4`_ translated to JavaScript_.

.. list-table:: Python_ `Beautiful Soup 4`_ vs JavaScript_
   :header-rows: 1
   :class: table-syntax-diff

   * - Python_ `Beautiful Soup 4`_
     - JavaScript_

   * - bs4 `making the soup`_ (document_ object)

       .. code-block:: python

         from bs4 import BeautifulSoup

         with open(path, 'r') as f:
             document = BeautifulSoup(f)

     - document_ object is built-in

   * - getElementById_ (bs4 find_)

       .. code-block:: python

         foo = document.find(id="foo")

     - getElementById_

       .. code-block:: javascript

         var foo = document.getElementById("foo");

   * - getElementsByTagName_ (bs4 `find_all`_)

       .. code-block:: python

         divs = document.find_all("div")
         # or
         divs = element.find_all("div")

     - getElementsByTagName_

       .. code-block:: javascript

         var divs = document.getElementsByTagName("div");
         // or
         var divs = element.getElementsByTagName("div");

   * - `Searching by CSS class`_

       .. code-block:: python

         # find first one div with class name page
         div = document.find("div", class_="page")
         # find all divs with class name page
         divs = document.find_all("div", class_="page")
         # or
         div = element.find("div", class_="page")
         divs = element.find_all("div", class_="page")

     - querySelector_ and querySelectorAll_

       .. code-block:: javascript

         // find first one div with class name page
         var div = document.querySelector("div.page");
         // find all divs with class name page
         var divs = document.querySelectorAll("div");
         // or
         var div = element.querySelector("div.page");
         var divs = element.querySelectorAll("div");

   * - Attributes_

       .. code-block:: python

         href = element.get("href")

     - getAttribute_

       .. code-block:: javascript

         var href = element.getAttribute("href");

----

Tested on: ``Ubuntu Linux 15.10``, ``Python 2.7.10``.

----

References:

.. [1] `[Python] Export PIXNET Blog to reStructuredText Files <{filename}../17/python-export-pixnet-blog-to-rst%en.rst>`_


.. _Python: https://www.python.org/
.. _JavaScript: https://www.google.com/search?q=javascript
.. _Beautiful Soup 4: https://www.google.com/search?q=Beautiful+Soup+4
.. _document: http://www.w3schools.com/jsref/dom_obj_document.asp
.. _making the soup: http://www.crummy.com/software/BeautifulSoup/bs4/doc/#making-the-soup
.. _getElementById: http://www.w3schools.com/jsref/met_doc_getelementbyid.asp
.. _find: http://www.crummy.com/software/BeautifulSoup/bs4/doc/#find
.. _getElementsByTagName: http://www.w3schools.com/jsref/met_document_getelementsbytagname.asp
.. _find_all: http://www.crummy.com/software/BeautifulSoup/bs4/doc/#a-string
.. _Searching by CSS class: http://www.crummy.com/software/BeautifulSoup/bs4/doc/#searching-by-css-class
.. _querySelector: https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
.. _querySelectorAll: https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
.. _Attributes: http://www.crummy.com/software/BeautifulSoup/bs4/doc/#attributes
.. _getAttribute: http://www.w3schools.com/jsref/met_element_getattribute.asp
