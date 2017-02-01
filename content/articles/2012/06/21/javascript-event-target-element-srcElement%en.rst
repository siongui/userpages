JavaScript Event Target Element (srcElement)
############################################

:date: 2012-06-21 15:59
:modified: 2015-04-07 09:58
:tags: JavaScript, html, DOM
:category: JavaScript
:summary: Event target element in JavaScript_ event handling.
:adsu: yes


When designing web application with JavaScript_ event handling, it is important
to *figure out what the (possible) target element is*.

Let's take mouseover_ as example, see the following description from [1]_:

  **mouseover**

    Fires when the user moves the mouse over the element you registered the
    event on *or one of its descendants*.

There is an example illustrating this situation in [2]_. And mouseenter event in
[3]_ provides a 'do not bubble' version of the mouseover event, but it is a pity
that neither Firefox or Chrome supports the event. So, when you get the event
target element (or srcElement) in the callback function of events, it is
possible the target element may be one of descendant elements. As a result, when
dealing with the callback function of events, *REMEMBER to CHECK what the TARGET
ELEMENT is*.

I wrote sample code to show how to deal with the target issue of JavaScript
event handling. Please see the following demo and code:

.. rubric:: `Demo <{filename}target.html>`_
   :class: align-center

.. adsu:: 2
.. show_github_file:: siongui userpages content/articles/2012/06/21/target.html
.. adsu:: 3
.. show_github_file:: siongui userpages content/articles/2012/06/21/target.js

.. show_github_file:: siongui userpages content/articles/2012/06/21/style.css

As you can see, only *div1* onclick event is registered. But when child elements
(i.e., *div2* or *span1*) of *div1* are clicked, the target element of the
callback function is the child element, not *div1*. The *_getParentElement*
function is provided in the sample code to show how to reach *div1* in this
case.

----

References:

.. [1] `Event - mouseover and mouseout <http://www.quirksmode.org/dom/events/mouseover.html>`_

.. [2] `mouseover and mouseout - Mouse Event - QuirksMode <http://www.quirksmode.org/js/events_mouse.html#mouseover>`_

.. [3] `mouseenter - Event compatibility tables <http://www.quirksmode.org/dom/events/index.html#t017>`_

.. [4] `Accessing the HTML element - Introduction to Events - QuirksMode <http://www.quirksmode.org/js/introevents.html#link11>`_

.. [5] `Which HTML element is the target of the event? - Event properties - QuirksMode <http://www.quirksmode.org/js/events_properties.html#target>`_

Additional References:

.. [6] `Javascript - Introduction to Events - QuirksMode <http://www.quirksmode.org/js/introevents.html>`_

.. [7] `Javascript - Event properties - QuirksMode <http://www.quirksmode.org/js/events_properties.html>`_

.. [8] `javascript - How do I prevent a parent's onclick event from firing when a child anchor is clicked? - Stack Overflow <http://stackoverflow.com/questions/1369035/how-do-i-prevent-a-parents-onclick-event-from-firing-when-a-child-anchor-is-cli>`_

.. [9] `JavaScript Event Delegation is Easier than You Think <http://www.sitepoint.com/javascript-event-delegation-is-easier-than-you-think/>`_

.. [10] `Delegating the focus and blur events - QuirksBlog <http://www.quirksmode.org/blog/archives/2008/04/delegating_the.html>`_

.. [11] `Event Delegation versus Event Handling <http://icant.co.uk/sandbox/eventdelegation/>`_

.. [12] `Event Delegation Made Easy <http://danwebb.net/2008/2/8/event-delegation-made-easy-in-jquery>`_

.. [13] `JavaScript Kit- Event Object <http://www.javascriptkit.com/jsref/event.shtml>`_

.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _mouseover: http://www.quirksmode.org/dom/events/mouseover.html
