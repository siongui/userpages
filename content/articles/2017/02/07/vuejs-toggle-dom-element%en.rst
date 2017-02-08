[Vue.js] Toggle (Show/Hide) HTML Element
########################################

:date: 2017-02-07T09:54+08:00
:tags: Vue.js, JavaScript, DOM, toggle, toggleable
:category: Vue.js
:summary: Toggle (Show/Hide) HTML_ `DOM element`_ in Vue.js_.
:og_image: https://vuejs.org/images/logo.png
:adsu: yes


Learn to toggle (Show/Hide) HTML_ `DOM element`_ is my favorite exercise for
using a new JavaScript_ framework/library or compile-to-JavaScript languages
such as GopherJS_.

.. rubric:: `Demo <{filename}/code/vuejs/toggle-element/index.html>`_
   :class: align-center

The following is key points in the code:

- Use boolean variable ``isShow`` to control the visibility of DOM element.
- When users click on the element with class name *control* (line 15~17),
  ``isShow = !isShow`` expression will make ``isShow`` become *true* if it is
  *false*, become *false* if it is *true*.
- The element with class name *content* (line 18) will display or disappear by
  v-show_ directive and true/false of ``isShow``.

.. adsu:: 2
.. show_github_file:: siongui userpages content/code/vuejs/toggle-element/index.html
.. adsu:: 3

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``Vue.js 2.1.10``.

----

References:

.. [1] `Event Handling — Vue.js <https://vuejs.org/v2/guide/events.html>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _GopherJS: http://www.gopherjs.org/
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _Vue.js: https://vuejs.org/
.. _v-show: https://vuejs.org/v2/api/#v-show
