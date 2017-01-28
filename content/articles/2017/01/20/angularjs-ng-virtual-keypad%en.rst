[AngularJS] Virtual Keyboard
############################

:date: 2017-01-20T20:51+08:00
:tags: AngularJS, JavaScript, IME, Pāli Input Method, Virtual Keyboard/Keypad
:category: AngularJS
:summary: Virtual kayboard/keypad via AngularJS_. Used in `Pāli Dictionary`_.
:adsu: yes


Virtual kayboard/keypad via AngularJS_.

.. rubric:: `Demo <{filename}/code/angularjs/virtual-keyboard/index.html>`_
   :class: align-center

Real world application is `Pāli Dictionary`_.
There are special characters in `romanized Pāli`_. For the convenience of input
`Pāli`_ words, users can use the virtual keyboard to type Pāli_ word without
installation of `Pāli`_ input method in computers

**Source code**:

.. show_github_file:: siongui userpages content/code/angularjs/virtual-keyboard/index.html
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/angularjs/virtual-keyboard/app.js
.. show_github_file:: siongui userpages content/code/angularjs/virtual-keyboard/keypad.css
.. adsu:: 3

To make virtual keyboard draggable, see [2]_.

To make virtual keyboard toggle-able, see [3]_.

For Vue.js_ version, see [4]_.

For plain JavaScript_ version, see [5]_.

For Dart_ version, see [6]_.

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``
- ``AngularJS 1.5.7``.

----

References:

.. [1] `pali/input.html at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/dictionary/app/partials/input.html>`_

       `pali/inputSuggest.js at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/dictionary/app/scripts/directives/inputSuggest.js>`_

.. [2] `[AngularJS] Draggable (Movable) Element <{filename}../../../2013/04/04/angularjs-draggable-movable-element%en.rst>`_

.. [3] `[AngularJS] Toggle Element without JavaScript <{filename}../../../2013/06/22/angularjs-toggle-element-without-javascript%en.rst>`_

.. [4] `[Vue.js] Virtual Keyboard <{filename}../21/vuejs-virtual-keypad%en.rst>`_
.. [5] `[JavaScript] Virtual Keyboard <{filename}../28/javascript-virtual-keypad%en.rst>`_
.. [6] `[Dart] Virtual Keyboard <{filename}../29/dartlang-virtual-keypad%en.rst>`_


.. _AngularJS: https://angularjs.org/
.. _Vue.js: https://vuejs.org/
.. _Dart: https://www.dartlang.org/
.. _Directives: https://docs.angularjs.org/guide/directive
.. _Pāli Dictionary: http://dictionary.sutta.org/
.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _romanized Pāli: https://www.google.com/search?q=romanized+P%C4%81li
.. _JavaScript: https://www.google.com/search?q=JavaScript
