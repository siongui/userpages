Python Library xml.dom.minidom Howto (2)
########################################

:tags: DOM, html, minidom, Python, XML
:category: Python
:summary: Python XML/HTML manipulation primer of xml.dom.minidom
:adsu: yes


CREATE AND ADD A TEXT NODE
==========================

In previous post [1]_, we showed how to create a XML DOM, of which the root element is called 'html'. In this post, we show how to add a text node to the root element. Note that the text node is the leaf in the DOM trees, which means the text node has no children node or element.

.. show_github_file:: siongui userpages content/articles/2012/05/24/minidom-howto-2.py

In line 10, we get the root element of the document DOM. In line 11, we create a text node, and in line 12, we add the text node to the root element. Note that the text node can be added to any element, not only to root element.

The following is the output:

.. code-block:: bash

    <?xml version="1.0" ?><html>Hello World!</html>

It's pretty easy and straight forward to manipulate the minidom library. In the next post [3]_, we will show how to create an element and add it to the DOM trees.

----

*Python Library xml.dom.minidom Howto* series:

.. [1] `Python Library xml.dom.minidom Howto (1) <{filename}python-xml-dom-minidom-howto-1%en.rst>`_

.. [2] `Python Library xml.dom.minidom Howto (2) <{filename}python-xml-dom-minidom-howto-2%en.rst>`_

.. [3] `Python Library xml.dom.minidom Howto (3) <{filename}python-xml-dom-minidom-howto-3%en.rst>`_

.. [4] `Python Library xml.dom.minidom Howto (4) <{filename}python-xml-dom-minidom-howto-4%en.rst>`_

.. [5] `Python Library xml.dom.minidom Howto (5) <{filename}python-xml-dom-minidom-howto-5%en.rst>`_

.. [6] `Python Library xml.dom.minidom Howto (6) <{filename}python-xml-dom-minidom-howto-6%en.rst>`_

.. [7] `Python Library xml.dom.minidom Howto (7) <{filename}../27/python-xml-dom-minidom-howto-7%en.rst>`_
