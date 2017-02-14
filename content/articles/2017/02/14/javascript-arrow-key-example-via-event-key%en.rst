JavaScript Arrow Key Example via event.key in Keyboard Event
############################################################

:date: 2017-02-14T22:32+08:00
:tags: JavaScript, Keyboard Event, DOM
:category: JavaScript
:summary: Detect `arrow keystrokes`_ via event.key_ of KeyboardEvent_
          in JavaScript_.
:adsu: yes


This morning I read the KeyboardEvent_ documentation on MDN_, and found that
`event.keyCode`_ is deprecated, so I re-write the example of detecting arrow
keys [1]_ in JavaScript_ again to use event.key_ in this post. If you need to
support old browsers, please refer to `my previous post`_.

There are three events related to keyboards: *onkeydown*, *onkeypress*,
*onkeyup*. To detect arrow keys, please use onkeydown_ (see [4]_).

In the example, the *keydown* event of document_ is listened. If users press any
key, the event handler will be triggered. In the handler, the *key* property of
of (keyboard) event will be checked to see which key is pressed.

.. rubric:: `Demo <{filename}/code/javascript/arrow-keystroke-via-event-key/index.html>`_
      :class: align-center

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/code/javascript/arrow-keystroke-via-event-key/index.html
.. adsu:: 2

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/code/javascript/arrow-keystroke-via-event-key/app.js
.. adsu:: 3

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Mozilla Firfox 51.0.1 (64-bit)``

----

References:

.. [1] `JavaScript Keyboard Event (Arrow Key Example) <{filename}../../../2012/06/25/javascript-keyboard-event-arrow-key-example%en.rst>`_
.. [2] `KeyboardEvent.key - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key>`_
.. [3] `javascript switch - Google search <https://www.google.com/search?q=javascript+switch>`_
.. [4] `keyboard events - Detecting arrow key presses in JavaScript - Stack Overflow <http://stackoverflow.com/questions/5597060/detecting-arrow-key-presses-in-javascript>`_

.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _arrow keystrokes: https://www.google.com/search?q=arrow+keystrokes
.. _arrow keys: https://www.google.com/search?q=arrow+keys
.. _event.key: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
.. _event.keyCode: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/keyCode
.. _MDN: https://developer.mozilla.org/
.. _my previous post: {filename}../../../2012/06/25/javascript-keyboard-event-arrow-key-example%en.rst
.. _KeyboardEvent: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent
.. _keyboard event: https://www.google.com/search?q=keyboard+event
.. _document: https://developer.mozilla.org/en-US/docs/Web/API/Document
.. _onkeydown: http://www.w3schools.com/jsref/event_onkeydown.asp
