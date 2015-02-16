JavaScript Drag and Drop (Draggable, Movable) Element without External Library
##############################################################################

:date: 2012-07-13 00:03
:tags: JavaScript, DOM, draggable
:category: JavaScript
:summary: Draggable, movable HTML element using vanilla JavaScript.


There are many libraries like jQuery_ which provide plugins_ to make an
`DOM element`_ movable and draggable. I would like, however, to know how to make
an element draggable without external library. In [1]_, the author provides an
implementation of draggable element both using keyboard and mouse. He also gave
an post (see [3]_) about HTML5 draggable issue. His implementation:

  1) includes both keyboard and mouse, but I want mouse only.

  2) *this* keyword in the callback function refers to the DOM element at which
     events occur, but I want *this* keyword to refer to the object which makes
     DOM element draggable.

In [2]_, the author also gives another implementation, but I like the code
structure to be more "object-oriented" like the implementation in [1]_. So I
wrote an implementation of my own. The following is my implementation:

.. rubric:: `Demo <{filename}/code/vanilla-javascript-draggable/movable.html>`_
   :class: align-center

.. show_github_file:: siongui userpages content/code/vanilla-javascript-draggable/movable.html

.. show_github_file:: siongui userpages content/code/vanilla-javascript-draggable/draggable.js

Note that the element to be made draggable must have CSS property
:code:`position: absolute;` or :code:`position: fixed;` (*position* set to
*absolute* in the JavaScript code). I put a lot of comments in the code to make
the code understandable. Hope this would be helpful for those who are
interested.

----

References:

.. [1] `Drag and drop - QuirksMode <http://www.quirksmode.org/js/dragdrop.html>`_

.. [2] `JavaScript Drag and Drop Tutorial <http://luke.breuer.com/tutorial/javascript-drag-and-drop-tutorial.aspx>`_

.. [3] `The HTML5 drag and drop disaster <http://www.quirksmode.org/blog/archives/2009/09/the_html5_drag.html>`_


.. _jQuery: http://jquery.com/

.. _plugins: http://jqueryui.com/draggable/

.. _DOM element: http://www.w3schools.com/dom/dom_element.asp
