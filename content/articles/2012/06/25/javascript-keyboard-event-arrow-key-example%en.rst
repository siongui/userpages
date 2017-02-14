JavaScript Keyboard Event (Arrow Key Example)
#############################################

:date: 2012-06-25 13:18
:modified: 2017-02-14T22:50+08:00
:tags: JavaScript, Keyboard Event, DOM
:category: JavaScript
:summary: Detect `arrow keystrokes`_ in JavaScript_.
:adsu: yes


This post will give an example for detecting `arrow keys`_ using JavaScript_.
For general `keyboard event`_, it is very easy to achieve by extending the
example in this post (also see references below).

There are three events related to keyboards: *onkeydown*, *onkeypress*,
*onkeyup*. To detect arrow keys, please use onkeydown_ (see [2]_).

.. rubric:: `Demo <{filename}/code/javascript/keyboard-event/arrow-key.html>`_
      :class: align-center

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/code/javascript/keyboard-event/arrow-key.html
.. adsu:: 2

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/code/javascript/keyboard-event/arrow-key.js
.. adsu:: 3

According to the KeyboardEvent_ documentation on MDN_, and the event.keyCode_ is
deprecated. Should use event.key_ instead. If you need to support old browsers,
use the example in this post. Otherwise see [5]_ for new example of event.key_.

----

References:

.. [1] `Keyboard events | JavaScript Tutorial <https://javascript.info/tutorial/keyboard-events>`_
.. [2] `keyboard events - Detecting arrow key presses in JavaScript - Stack Overflow <https://stackoverflow.com/questions/5597060/detecting-arrow-key-presses-in-javascript>`_
.. [3] `JavaScript - Detecting keystrokes - QuirksMode <https://www.quirksmode.org/js/keys.html>`_
.. [4] `JavaScript Madness: Keyboard Events <http://unixpapa.com/js/key.html>`_
.. [5] `JavaScript Arrow Key Example via event.key in Keyboard Event <{filename}../../../2017/02/14/javascript-arrow-key-example-via-event-key%en.rst>`_

.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _arrow keystrokes: https://www.google.com/search?q=arrow+keystrokes
.. _arrow keys: https://www.google.com/search?q=arrow+keys
.. _keyboard event: https://www.google.com/search?q=keyboard+event
.. _onkeydown: http://www.w3schools.com/jsref/event_onkeydown.asp
.. _KeyboardEvent: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent
.. _MDN: https://developer.mozilla.org/
.. _event.key: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
.. _event.keyCode: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/keyCode
