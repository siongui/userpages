Python Library xml.dom.minidom Howto (3)
########################################

:tags: DOM, howto, html, minidom, Python, tutorial, XML
:category: Python
:summary: Python XML/HTML manipulation primer of xml.dom.minidom


ADD AN ELEMENT
==============

In previous post (`1 <{filename}python-xml-dom-minidom-howto-1%en.rst>`_ and `2 <{filename}python-xml-dom-minidom-howto-2%en.rst>`_), we showed how to create a document and add a text node to the root document. In this post, we will show how to create a new element node and add it to the DOM tree.

.. show_github_file:: siongui userpages content/articles/2012/05/24/minidom-howto-3.py

In line 11 and 12, we create a node with the tag 'demoTag' and add it to the root element of the document. The following is the output:

.. code-block:: bash

    <?xml version="1.0" ?><html><demoTag/></html>

In the `next post <{filename}python-xml-dom-minidom-howto-4%en.rst>`_, we will show how to set the attribute of the element.
