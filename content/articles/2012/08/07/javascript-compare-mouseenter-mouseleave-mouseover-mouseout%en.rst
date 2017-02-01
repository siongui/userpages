[JavaScript] Comparison of MouseEnter MouseLeave MouseOver MouseOut
###################################################################

:date: 2012-08-07 20:07
:modified: 2015-02-26 20:08
:tags: html, JavaScript, mouseenter, mouseleave
:category: JavaScript
:summary: Compare mouseenter, mouseleave, mouseover, and mouseout events.
:adsu: yes


onmouseenter_ / onmouseleave_ event of a DOM element will be triggered **only**
when the mouse cursor *enters* / *leaves* the element:

.. rubric:: `Demo of MouseEnter & MouseLeave event <{filename}/code/javascript-mouseenter-mouseleave/mouseenterleave.html>`_
      :class: align-center

onmouseover_ / onmouseout_ event of a DOM element will be triggered **not only**
when the mouse cursor *enters* / *leaves* the element, **but also** when the
mouse cursor *enters* / *leaves* sub-elements of the element.

.. rubric:: `Demo of MouseOver & MouseOut event <{filename}/code/javascript-mouseenter-mouseleave/mouseoverout.html>`_
      :class: align-center

.. adsu:: 2

----

See also:
`[JavaScript] onMouseEnter and onMouseLeave Suppport for Old Browsers <{filename}../../10/02/javascript-mouseenter-mouseleave%en.rst>`_

----

References:

.. [1] `Mouse Events - QuirksMode <http://www.quirksmode.org/js/events_mouse.html>`_

.. [2] `javascript - What is the difference between the mouseover and mouseenter events? - Stack Overflow <http://stackoverflow.com/questions/1104344/what-is-the-difference-between-the-mouseover-and-mouseenter-events>`_


.. _onmouseenter: http://www.w3schools.com/jsref/event_onmouseenter.asp

.. _onmouseleave: http://www.w3schools.com/jsref/event_onmouseleave.asp

.. _onmouseover: http://www.w3schools.com/jsref/event_onmouseover.asp

.. _onmouseout: http://www.w3schools.com/jsref/event_onmouseout.asp
