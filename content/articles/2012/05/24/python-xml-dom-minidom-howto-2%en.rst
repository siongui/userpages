Python Library xml.dom.minidom Howto (2)
########################################

:tags: DOM, howto, html, minidom, Python, tutorial, XML
:category: Python
:summary: Python XML/HTML manipulation primer of xml.dom.minidom


CREATE AND ADD A TEXT NODE
==========================

In `previous post <{filename}python-xml-dom-minidom-howto-1%en.rst>`_, we showed how to create a XML DOM, of which the root element is called 'html'. In this post, we show how to add a text node to the root element. Note that the text node is the leaf in the DOM trees, which means the text node has no children node or element.

.. show_github_file:: siongui userpages content/articles/2012/05/24/minidom-howto-2.py

In line 10, we get the root element of the document DOM. In line 11, we create a text node, and in line 12, we add the text node to the root element. Note that the text node can be added to any element, not only to root element.

The following is the output:

.. code-block:: bash

    <?xml version="1.0" ?><html>Hello World!</html>

It's pretty easy and straight forward to manipulate the minidom library. In the `next post <{filename}python-xml-dom-minidom-howto-3%en.rst>`_, we will show how to create an element and add it to the DOM trees.
