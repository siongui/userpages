[Vue.js] Draggable (Movable) Element
####################################

:date: 2017-02-06T21:45+08:00
:tags: Vue.js, JavaScript, DOM, draggable
:category: Vue.js
:summary: draggable_, movable (drag and drop) `HTML element`_ in Vue.js_.
:og_image: https://vuejs.org/images/logo.png
:adsu: yes


Use Vue.js_ `custom directive`_ to make `HTML element`_ draggable_.

The key points in the code:

- When users `mouse down`_ the `HTML element`_, record the position of the HTML
  element and the initial mouse position.
- When users `mouse move`_ the element, calculate the distance of initial
  position and current position, and then update the position of the element
  accordingly.
- When users release the button (`mouse up`_) over the element, remove
  registered event listener.

.. rubric:: `Demo <{filename}/code/vuejs/draggable-element/index.html>`_
   :class: align-center
.. show_github_file:: siongui userpages content/code/vuejs/draggable-element/index.html
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/vuejs/draggable-element/app.js
.. adsu:: 3

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.1.10``.

----

References:

.. [1] `Custom Directives — Vue.js <https://vuejs.org/v2/guide/custom-directive.html>`_

.. _HTML element: https://www.google.com/search?q=HTML+element
.. _draggable: https://www.google.com/search?q=draggable
.. _mouse down: https://developer.mozilla.org/en/docs/Web/Events/mousedown
.. _mouse move: https://developer.mozilla.org/en/docs/Web/Events/mousemove
.. _mouse up: https://developer.mozilla.org/en/docs/Web/Events/mouseup
.. _Vue.js: https://vuejs.org/
.. _custom directive: https://vuejs.org/v2/guide/custom-directive.html
