Python Library xml.dom.minidom Howto (6)
########################################

:tags: DOM, html, minidom, Python, XML
:category: Python
:summary: Python XML/HTML manipulation primer of xml.dom.minidom


WRITE XML/HTML TO FILE
======================

In previous posts (`1 <{filename}python-xml-dom-minidom-howto-1%en.rst>`_, `2 <{filename}python-xml-dom-minidom-howto-2%en.rst>`_, `3 <{filename}python-xml-dom-minidom-howto-3%en.rst>`_, `4 <{filename}python-xml-dom-minidom-howto-4%en.rst>`_, `5 <{filename}python-xml-dom-minidom-howto-5%en.rst>`_), we show how to generate XML/HTML content using Python minidom. Here we will show how to write the content into a file. The following is the code:

.. show_github_file:: siongui userpages content/articles/2012/05/24/minidom-howto-6.py

The trick here is in line 5, and lines 19, 20, 21, 22. The library *codecs* must be imported to encode the XML file. The lines 19, 20, 21, 22 perform the task of writing the XML content into a file. It's pretty simple and easy. This again shows the power and beauty of Python.
