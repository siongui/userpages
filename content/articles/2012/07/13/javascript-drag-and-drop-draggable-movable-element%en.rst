JavaScript Drag and Drop (Draggable, Movable) Element without External Library
##############################################################################

:date: 2012-07-13 00:03
:modified: 2015-02-20 10:07
:tags: JavaScript, DOM, draggable
:category: JavaScript
:summary: Draggable, movable HTML_ element using vanilla JavaScript_.
:adsu: yes


There are many libraries like jQuery_ which provide plugins_ to make an
`DOM element`_ movable and draggable. I would like, however, to know how to make
an element draggable without external library. In [1]_, the author provides an
implementation of draggable element both using keyboard and mouse. He also gave
an post (see [3]_) about HTML5 draggable issue. His implementation:

  1) includes both keyboard and mouse, but I want mouse only.

  2) *this* keyword in the callback function refers to the DOM_ element at which
     events occur, but I want *this* keyword to refer to the object which makes
     DOM element draggable.

In [2]_, the author also gives another implementation, but I like the code
structure to be more "object-oriented" like the implementation in [1]_. So I
wrote an implementation of my own. The following is my implementation:

.. rubric:: `Demo <{filename}/code/vanilla-javascript-draggable/movable.html>`_
   :class: align-center

.. adsu:: 2
.. show_github_file:: siongui userpages content/code/vanilla-javascript-draggable/movable.html

.. show_github_file:: siongui userpages content/code/vanilla-javascript-draggable/draggable.js
.. adsu:: 3

Note that the element to be made draggable must have `CSS property`_
:code:`position: absolute;` or :code:`position: fixed;` (*position* set to
*absolute* in the JavaScript code). I put a lot of comments in the code to make
the code understandable. Hope this would be helpful for those who are
interested.

| If you need draggable elements in AngularJS_ way, see [4]_.
| If you need draggable elements in Dart_, see [5]_.

----

References:

.. [1] `Drag and drop - QuirksMode <http://www.quirksmode.org/js/dragdrop.html>`_

.. [2] `JavaScript Drag and Drop Tutorial <http://luke.breuer.com/tutorial/javascript-drag-and-drop-tutorial.aspx>`_

.. [3] `The HTML5 drag and drop disaster <http://www.quirksmode.org/blog/archives/2009/09/the_html5_drag.html>`_

.. [4] `[AngularJS] Draggable (Movable) Element <{filename}../../../2013/04/04/angularjs-draggable-movable-element%en.rst>`_

.. [5] `[Dart] Draggable (Movable) Element <{filename}../../../2015/02/17/dart-draggable-movable-element%en.rst>`_

.. [6] `[Golang] Draggable (Movable) Element by GopherJS <{filename}../../../2016/01/17/go-draggable-movable-element-by-gopherjs%en.rst>`_


.. _jQuery: http://jquery.com/
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _plugins: http://jqueryui.com/draggable/
.. _CSS property: https://www.google.com/search?q=CSS+property
.. _DOM element: http://www.w3schools.com/dom/dom_element.asp
.. _DOM: https://www.google.com/search?q=DOM
.. _AngularJS: https://angularjs.org/
.. _HTML: https://www.google.com/search?q=HTML
.. _Dart: https://www.dartlang.org/
