JavaScript Create DOM Element Dynamically
#########################################

:date: 2012-09-26 23:17
:modified: 2015-02-20 12:01
:tags: html, JavaScript, XML, DOM
:category: JavaScript
:summary: Create DOM elements by JavaScript.
:adsu: yes


In the previous post_, we discuss creating HTML elements in Python script on
server side. This post can be viewed as the JavaScript version of previous
post_. We will discuss creating HTML elements on client side, i.e., on user
browser, and append the elements to existing DOM tree. Just as the previous
post_, there are also two ways to create HTML DOM elements, and both ways are
conceptually the same as their Python counterparts:

  1. **innerHTML way**: Directly assign HTML code to *innerHTML* property of
     parent element.

  2. **appendChild way**: Create DOM elements by *document.createElement*
     method, and append them to parent element by *appendChild* method.

As previous post_, a example will be given here. We will create a *div* element,
and append an *anchor* element to the *div* element, then append the *div*
element to the end of *body* element.

.. adsu:: 2

**1. innerHTML way:**

Please see the following code snippet. It's very easy and intuitive. We directly
assign the HTML code to *innerHTML* property of body element.

.. code-block:: javascript

  <script type="text/javascript">
    var body = document.getElementsByTagName('body')[0];
    body.innerHTML = '<div><a href="www.google.com">Google Search</a></div>';
  </script>

.. adsu:: 3

**2. appendChild way:**

Please see the following code snippet. The *div* and *anchor* elements are
created by *document.createElement* method, and then appended to parent elements
by *appendChild* method. The result of the following code snippet is the same as
that of the above code snippet. At least there are no problems for me so far.

.. code-block:: javascript

  <script type="text/javascript">
    var body = document.getElementsByTagName('body')[0];
    var div = document.createElement('div');
    var a = document.createElement('a');
    a.href = 'www.google.com';
    div.appendChild(a);
    body.appendChild(div);
  </script>

Personally I sometimes choose one way to create elements, sometimes the other
way. My criteria for choosing is based on code readability. If one way can make
the code more readable and understandable than the other, I will choose it. Hope
this post would be helpful for those who need.


.. _post: {filename}python-create-html-element-dynamically%en.rst
