JavaScript Keyboard Event (Arrow Key Example)
#############################################

:date: 2012-06-25 13:18
:modified: 2015-02-18 15:24
:tags: JavaScript, Keyboard Event
:category: JavaScript
:summary: Detect arrow keystrokes in JavaScript.
:adsu: yes


This post will give an example for detecting arrow keys using JavaScript. For
general keyboard event, it is very easy to achieve by extending the example in
this post (also see references below).

There are three events related to keyboards: *onkeydown*, *onkeypress*,
*onkeyup*. To detect arrow keys, please use onkeydown (see [2]_).

.. rubric:: `Demo <{filename}/code/javascript-keyboard-event/arrow-key.html>`_
      :class: align-center

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/code/javascript-keyboard-event/arrow-key.html

.. adsu:: 2

Source Code for Demo (*JavaScript*):

.. show_github_file:: siongui userpages content/code/javascript-keyboard-event/arrow-key.js

----

References:

.. [1] `Keyboard events | JavaScript Tutorial <http://javascript.info/tutorial/keyboard-events>`_

.. [2] `Detecting arrow key presses in JavaScript <http://stackoverflow.com/questions/5597060/detecting-arrow-key-presses-in-javascript>`_

.. [3] `Detecting keystrokes <http://www.quirksmode.org/js/keys.html>`_

.. [4] `JavaScript Madness: Keyboard Events <http://unixpapa.com/js/key.html>`_
