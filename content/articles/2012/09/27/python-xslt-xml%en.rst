[Python] Use XSL to Transform XML (XSLT)
########################################

:date: 2012-09-27 22:07
:modified: 2015-02-18 18:49
:tags: html, Python, XML, XSLT
:category: Python
:summary: XSLT (Extensible Stylesheet Language Transformations) in Python.
:adsu: yes

To transform XML_ document into another XML document (XSLT_) in Python_, we use
lxml_ library (library for processing XML and HTML in the Python language.) to
do the transformation for us.

First, we need a XML document, which is to be transformed, and a XSL_ document,
which instructs how to do the transformation. The following sample Python script
demonstrates how to do XSLT_:

.. adsu:: 2
.. show_github_file:: siongui userpages content/articles/2012/09/27/xslt.py
.. adsu:: 3

In the last line, the path of XML and XSL files are passed as arguments to the
*transform* function, which returns the content (strings) of transformed XML
document. To know the details of XSLT in lxml_ library, please see reference
[1]_. If you would like to know how to perform XSLT using JavaScript, please see
reference [2]_.

----

References:

.. [1] `XPath and XSLT with lxml <http://lxml.de/xpathxslt.html>`_

.. [2] `XSLT - On the Client <http://www.w3schools.com/Xsl/xsl_client.asp>`_

.. _XSLT: http://en.wikipedia.org/wiki/XSLT
.. _XML: https://www.google.com/search?q=XML
.. _lxml: http://lxml.de/
.. _Python: https://www.python.org/
.. _XSL: http://en.wikipedia.org/wiki/XSL
