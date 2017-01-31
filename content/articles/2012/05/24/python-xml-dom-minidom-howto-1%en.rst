Python Library xml.dom.minidom Howto (1)
########################################

:modified: 2017-02-01T01:53+08:00
:tags: DOM, html, minidom, Python, XML
:category: Python
:summary: Python_ XML_/HTML_ manipulation primer of xml.dom.minidom_
:adsu: yes


CREATE A DOCUMENT
=================

Python_ provides a simple xml.dom.minidom_ library which can be used for
XML_/HTML_ manipulation. Here I want to share some of what I learned to
beginners on how to manipulate XML_/HTML_ in Python_.

First, we create a document DOM_ with root element called *html*. Here is the
sample code:

`Run code on repl.it <https://repl.it/F1jV/0>`_

.. show_github_file:: siongui userpages content/articles/2012/05/24/minidom-howto-1.py
.. adsu:: 2

In line 7 and 8, we create a dom with root element *html*. If you run the
script, the following result will show up.

.. code-block:: bash

    <?xml version="1.0" ?><html/>

In the next post [2]_, we will show how to add a text node to the root element.

----

*Python Library xml.dom.minidom Howto* series:

.. [1] `Python Library xml.dom.minidom Howto (1) <{filename}python-xml-dom-minidom-howto-1%en.rst>`_

.. [2] `Python Library xml.dom.minidom Howto (2) <{filename}python-xml-dom-minidom-howto-2%en.rst>`_

.. [3] `Python Library xml.dom.minidom Howto (3) <{filename}python-xml-dom-minidom-howto-3%en.rst>`_

.. [4] `Python Library xml.dom.minidom Howto (4) <{filename}python-xml-dom-minidom-howto-4%en.rst>`_

.. [5] `Python Library xml.dom.minidom Howto (5) <{filename}python-xml-dom-minidom-howto-5%en.rst>`_

.. [6] `Python Library xml.dom.minidom Howto (6) <{filename}python-xml-dom-minidom-howto-6%en.rst>`_

.. [7] `Python Library xml.dom.minidom Howto (7) <{filename}../27/python-xml-dom-minidom-howto-7%en.rst>`_
.. adsu:: 3

.. _Python: https://www.python.org/
.. _XML: https://www.google.com/search?q=XML
.. _HTML: https://www.google.com/search?q=HTML
.. _DOM: https://www.google.com/search?q=DOM
.. _xml.dom.minidom: https://www.google.com/search?q=xml.dom.minidom
