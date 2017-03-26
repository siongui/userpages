JavaScript Remove All Children of a DOM Element
###############################################

:date: 2012-09-26 04:56
:modified: 2017-03-25T22:32+08:00
:tags: html, JavaScript, DOM
:category: JavaScript
:summary: Bug-free way of removing child nodes of a DOM element in JavaScript.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


Here comes an interesting question: how to remove all child nodes of a DOM
element in JavaScript?

Assume that we have a DOM element (for example, *div* or *span*) whose name is
*elm*, the most trivial and intuitive way to remove all its child nodes is as
follows:

.. code-block:: javascript

  elm.innerHTML = '';

It looks perfectly fine, and should works without any problem. In fact, I use
this method to empty all contents of a DOM element in the beginning and it works
as expected. But one time when I used this method again in the code, it works as
expected in the latest version of Chrome, Firefox, and Opera, but failed to work
in IE8. It fails in IE8 when I did something like the following code snippet:

.. code-block:: javascript

  elm.innerHTML = '';
  elm.appendChild(childs);

I don't know why the above code will not work in IE8, so I did some search and
found better ways to remove all child nodes in references [1]_ and [2]_. The
"canonical" and bug-free way should be as follows:

.. code-block:: javascript

  while (elm.hasChildNodes()) {
    elm.removeChild(elm.lastChild);
  }

Maybe for some people the "canonical" way is obvious, but it takes me quite a
while to debug and know that sometimes the trivial way will not work. So I wrote
this post for those who have the same trouble as me.

.. adsu:: 2

----

**Appendix**

One of the methods to empty an array [5]_ in JavaScript is:

.. code-block:: javascript

  while(A.length > 0) {
    A.pop();
  }

As you can see, the idea of *empty an array* here is the same as
*remove all child nodes*. This is interesting so I put this in the appendix.

.. adsu:: 3

----

References:

.. [1] `javascript - Remove all the children DOM elements in div - Stack Overflow <http://stackoverflow.com/questions/683366/remove-all-the-children-dom-elements-in-div>`_
.. [2] `Remove all child elements of a DOM node in JavaScript - Stack Overflow <http://stackoverflow.com/questions/3955229/remove-all-child-elements-of-a-dom-node-in-javascript>`_
.. [3] `[Dart] DOM Element Remove All Children <{filename}../../../2014/01/31/dart-element-remove-all-children%en.rst>`_
.. [4] `[Golang] Remove All Child Nodes of a DOM Element by GopherJS <{filename}../../../2016/01/31/go-remove-all-children-of-dom-element-by-gopherjs%en.rst>`_
.. [5] `How do I empty an array in JavaScript? - Stack Overflow <http://stackoverflow.com/questions/1232040/how-do-i-empty-an-array-in-javascript>`_
