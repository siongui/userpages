[JavaScript] Firing Order (Precedence) of Event Handler
#######################################################

:date: 2012-10-03 01:16
:modified: 2015-03-03 23:00
:tags: JavaScript, event firing order
:category: JavaScript
:summary: Use onmousedown event instead of onclick event if the event handler to
          be fired before onblur event.
:adsu: yes


There are two HTML elements, one is *div* element, and the other is *input*
element. Assume that the focus is at *input* element in the beginning. onclick_
event handler of *div* element and onblur_ event handler of *input* element are
registered. If *div* element is clicked, *input* element will lose focus. Here
comes the question:

  *Which event handler is fired first?*

     onclick_ event handler of *div* element?

     or onblur_ event handler of *input* element?

  The answer is onblur_ event handler of *input* element is fired **before**
  onclick_ event handler of *div* element.

.. adsu:: 2

* What if I want onblur_ event handler fired **after** onclick_ event handler?

    This is a question similar to that asked in reference [1]_, and also the
    question I had when I developed my web application.

There are good news and bad news for the quesiton:

  - The bad news is, according to references [2]_ and [3]_, the firing order
    (precedence) of event handlers are not defined in the HTML spec and hence
    depends on browser implementations.

  - The good news is that although onblur_ event handler is fired before
    onclick_ event handler, onmousedown_ event handler is fired **before**
    onblur_ event handler. So we can move onclick_ event handler to onmousedown_
    event of *div* element, which achieves the same result we need.

.. adsu:: 3

Summary
+++++++

1. Firing order (precedence) of event handlers depends on browser
   implementation. No spec on this.

2. Register event handler to onmousedown_ event instead of onclick_ event, if
   you want the event handler to be fired **before** onblur_ event handler.

----

References:

.. [1] `jquery - How should I fire Javascript blur event after click event that causes the blur? - Stack Overflow <http://stackoverflow.com/questions/4084780/how-should-i-fire-javascript-blur-event-after-click-event-that-causes-the-blur>`_

.. [2] `What is the event precedence in JavaScript? - Stack Overflow <http://stackoverflow.com/questions/282245/what-is-the-event-precedence-in-javascript>`_

.. [3] `Javascript event handler order - Stack Overflow <http://stackoverflow.com/questions/5143817/javascript-event-handler-order>`_


.. _onclick: http://www.w3schools.com/jsref/event_onclick.asp

.. _onblur: http://www.w3schools.com/jsref/event_onblur.asp

.. _onmousedown: http://www.w3schools.com/jsref/event_onmousedown.asp
