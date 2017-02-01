Python Create HTML Element Dynamically on Server Side
#####################################################

:date: 2012-09-26 15:09
:modified: 2015-02-20 11:39
:tags: html, Python, XML, DOM
:category: Python
:summary: Serve DOM elements to clients dynamically on Google App Engine Python
:adsu: yes


For web application, sometimes we need to create HTML element on the server side
(for example, on `Google App Engine Python`_) and serve the webpage to the
client side (user browser). There are two different ways to create HTML elements
in Python script on server side:

  1. **Write HTML code directly**: This means you write the HTML code directly,
     and serve the code directly to user client browser.

  2. **Use Python xml.dom.minidom library**: Create DOM elements using Python
     xml.dom.minidom_ library, and serve the HTML code by the use of *toxml()*
     function in the minidom_ library. This method is similar to JavaScript DOM
     manipulation and gives programmers consistent programming style on both
     server side and client side development.

The following example will give demonstration on how to create HTML elements
in both ways.

In this example, we will try to create a *div* element, and append an *anchor*
element to the *div* element.

.. adsu:: 2

**1. Write HTML code directly**

It's very easy and requires no explaination. Just write the HTML code as you did
while editing a *HTML* document file, put the code in the python string and
serve it to user client browser:

.. code-block:: python

  HTMLcode = '<div><a href="www.google.com">Google Search</a></div>'
  # serve the HTML code to client on Google App Engine Python using webapp2
  self.response.out.write(HTMLcode)

**2. Use Python xml.dom.minidom library**

If you have no idea of what Python *minidom* is, please refer to series of
introduction to *minidom* library in this blog [1]_. The result of following
code snippet is exactly the same as that of the above code snippet, i.e.,
*HTMLcode* above is the same as *div.toxml()* below.

.. code-block:: python

  import xml.dom.minidom

  impl = xml.dom.minidom.getDOMImplementation()
  dom = impl.createDocument(None, 'div', None)
  div = dom.documentElement
  a = dom.createElement('a')
  a.setAttribute('href', 'www.google.com')
  a.appendChild(dom.createTextNode('Google Search'))
  div.appendChild(a)

  # serve th DOM to client on Google App Engine Python using webapp2
  self.response.out.write(div.toxml())

.. adsu:: 3

Personally I think generating HTML code with *mindom* library is somewhat
tedious, and I will write the HTML code directly if I want to make the code more
concise and readable. Generating HTML code with *minidom*, however, can make you
feel you are coding with JavaScript and makes your code style consistent either
on server side or client side scripting.

I hope this post would be helpful for those who want to develop web application
with Python on server side. For readers who want to know how to create elements
on client side (user browser) by JavaScript, please refer to [2]_.

----

References:

.. [1] `Python Library xml.dom.minidom Howto <{tag}minidom>`_

.. [2] `JavaScript Create DOM Element Dynamically <{filename}javascript-create-dom-element-dynamically%en.rst>`_

.. _Google App Engine Python: https://cloud.google.com/appengine/docs/python/

.. _xml.dom.minidom: https://docs.python.org/2/library/xml.dom.minidom.html

.. _minidom: https://docs.python.org/2/library/xml.dom.minidom.html
