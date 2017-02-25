[JavaScript] Change Scope (Context) of Anonymous Function
#########################################################

:date: 2012-10-01 14:01
:modified: 2015-03-07 02:01
:tags: JavaScript, Function.prototype.bind()
:category: JavaScript
:summary: Change the scope (context) of anonymous function - change what *this*
          keyword refers to in anonymous function.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


This post is an extension of previous post "*[JavaScript] Scope (Context) of
Event Handler Function* [2]_". In this post, we will discuss how to change the
scope (context) of a JavaScript `anonymous function`_. The anonymous function in
JavaScript has wide uses, for example, as an event handler function. See the
following code snippet:

.. code-block:: javascript

  element.onclick = function() {
    console.log(this);
  };

We use an anonymous function as the onclick_ event handler of
`HTML DOM element`_. If you register the event handler as above, *this* keyword
in the anonymous event handler function `refers to the DOM element itself`_.
What if we want to change the scope (context) of the anonymous event handler
function to global scope (i.e., *window*) or other scope?

.. adsu:: 2

The answer is `Function.prototype.bind()`_ introduced in recent browsers. For
example:

.. code-block:: javascript

  element.onclick = function() {
    console.log(this);
  }.bind(window);

The above code changes the scope (context) of the anonymous function to global
scope. Another example:

.. code-block:: javascript

  element.onclick = function() {
    console.log(this);
  }.bind(this);

Will bind the scope of the anonymous function to the scope when onclick_ event
of the element is registered.

The `Function.prototype.bind()`_ is great, but for older browsers (for example,
IE8) which do not support `Function.prototype.bind()`_, we need to implement
`custom Function.prototype.bind()`_ in JavaScript code. For more details of
`Function.prototype.bind()`_, please refer to reference [1]_.

.. adsu:: 3

----

References:

.. [1] `Function.prototype.bind() - JavaScript | MDN <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function/bind>`_

.. [2] `[JavaScript] Scope (Context) of Event Handler Function <{filename}javascript-scope-context-of-event-handler%en.rst>`_


.. _Function.prototype.bind(): https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function/bind

.. _custom Function.prototype.bind(): https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function/bind#Polyfill

.. _anonymous function: http://helephant.com/2008/08/23/javascript-anonymous-functions/

.. _onclick: http://www.w3schools.com/jsref/event_onclick.asp

.. _HTML DOM element: http://www.w3schools.com/jsref/dom_obj_all.asp

.. _refers to the DOM element itself: http://www.quirksmode.org/js/events_tradmod.html
