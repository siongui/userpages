Python Library xml.dom.minidom Howto (3)
########################################

:modified: 2017-02-01T02:28+08:00
:tags: DOM, html, minidom, Python, XML
:category: Python
:summary: Python_ XML_/HTML_ manipulation primer of xml.dom.minidom_
:adsu: yes


ADD AN ELEMENT
==============

In previous posts ([1]_ and [2]_), we showed how to create a document and add a
text node to the root document via xml.dom.minidom_. In this post, we will show
how to create a new element node and add it to the DOM_ tree.

`Run code on repl.it <https://repl.it/F1oN/0>`_

.. show_github_file:: siongui userpages content/articles/2012/05/24/minidom-howto-3.py
.. adsu:: 2

In line 11 and 12, we create a node with the tag *demoTag* and add it to the
root element of the document. The following is the output:

.. code-block:: bash

    <?xml version="1.0" ?><html><demoTag/></html>

In the next post [4]_, we will show how to set the attribute of the element.

----

.. adsu:: 3

*Python Library xml.dom.minidom Howto* series:

.. [1] `Python Library xml.dom.minidom Howto (1) <{filename}python-xml-dom-minidom-howto-1%en.rst>`_

.. [2] `Python Library xml.dom.minidom Howto (2) <{filename}python-xml-dom-minidom-howto-2%en.rst>`_

.. [3] `Python Library xml.dom.minidom Howto (3) <{filename}python-xml-dom-minidom-howto-3%en.rst>`_

.. [4] `Python Library xml.dom.minidom Howto (4) <{filename}python-xml-dom-minidom-howto-4%en.rst>`_

.. [5] `Python Library xml.dom.minidom Howto (5) <{filename}python-xml-dom-minidom-howto-5%en.rst>`_

.. [6] `Python Library xml.dom.minidom Howto (6) <{filename}python-xml-dom-minidom-howto-6%en.rst>`_

.. [7] `Python Library xml.dom.minidom Howto (7) <{filename}../27/python-xml-dom-minidom-howto-7%en.rst>`_

.. _Python: https://www.python.org/
.. _XML: https://www.google.com/search?q=XML
.. _HTML: https://www.google.com/search?q=HTML
.. _DOM: https://www.google.com/search?q=DOM
.. _xml.dom.minidom: https://www.google.com/search?q=xml.dom.minidom
