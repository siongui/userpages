[Vue.js] Keyboard Event (Arrow Key Example)
###########################################

:date: 2017-12-06T22:05+08:00
:tags: Vue.js, JavaScript, html, Keyboard Event
:category: Vue.js
:summary: Detect arrow keys pressed by users via Vue.js.
:og_image: https://image.slidesharecdn.com/introtovue-170720154735/95/intro-to-vuejs-23-638.jpg?cb=1500565888
:adsu: yes


See demo first. Focus_ on the following *textarea* element and press arrow keys
on your keyboard. You will see the keyCode_ of the arrow key pressed by you.

.. raw:: html

  <div id="vueapp">
    <textarea @keyup.arrow-keys="show">{{ keypressed }}</textarea>
  </div>

  <script src="https://unpkg.com/vue@2.5.9/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  Vue.config.keyCodes = {
    "arrow-keys": [37, 38, 39, 40]
  };

  var app = new Vue({
    el: '#vueapp',
    data: {
      keypressed: ""
    },
    methods: {
      show: function (event) {
        this.keypressed += event.keyCode;
      }
    }
  });
  </script>

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <div id="vueapp">
    <textarea @keyup.arrow-keys="show">{{ keypressed }}</textarea>
  </div>

  <script src="https://unpkg.com/vue@2.5.9/dist/vue.js"></script>

We use `key modifiers`_ [3]_ to listen to the keyup event of textarea element.
When textarea element is focused and users press any arrow key, the keyCode of
the arrow key will be shown in the textarea.

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  Vue.config.keyCodes = {
    "arrow-keys": [37, 38, 39, 40]
  };

  var app = new Vue({
    el: '#vueapp',
    data: {
      keypressed: ""
    },
    methods: {
      show: function (event) {
        this.keypressed += event.keyCode;
      }
    }
  });

There are four arrow keys on the keyboard. We need to define custom key alias
for the arrow keys in the `keyCodes in Vue.config`_. The keyCodes of arrow keys
are from 37~40, so we use array of numbers to define key alias.

Note that the following syntax in `Vue.config` does not work:

.. code-block:: javascript

  arrowKeys: [37, 38, 39, 40]

The correct syntax is:

.. code-block:: javascript

  "arrow-keys": [37, 38, 39, 40]

When users press arrow key in textarea element, the *show* method will be
executed and the keyCode_ of the corresponding arrow key will be appended to the
textarea.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 62.0.3202.94 (Official Build) Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``
- ``Vue.js 2.5.9``

----

References:

.. [1] `JavaScript Keyboard Event (Arrow Key Example) <{filename}../../../2012/06/25/javascript-keyboard-event-arrow-key-example%en.rst>`_
.. [2] `JavaScript Arrow Key Example via event.key in Keyboard Event <{filename}../../02/14/javascript-arrow-key-example-via-event-key%en.rst>`_
.. [3] `Key Modifiers - Event Handling — Vue.js <https://vuejs.org/v2/guide/events.html#Key-Modifiers>`_
.. [4] `keyCodes - Global Config - API - Vue.js <https://vuejs.org/v2/api/#keyCodes>`_

.. _Vue.js: https://vuejs.org/
.. _Focus: https://www.google.com/search?q=focus+javascript
.. _keyCode: https://www.google.com/search?q=keycode+javascript
.. _key modifiers: https://vuejs.org/v2/guide/events.html#Key-Modifiers
.. _keyCodes in Vue.config: https://vuejs.org/v2/api/#keyCodes
