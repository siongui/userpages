[Dart] Virtual Keyboard
#######################

:date: 2017-01-29T02:33+08:00
:tags: Dart, IME, Pāli Input Method, Virtual Keyboard/Keypad
:category: Dart
:summary: Virtual kayboard/keypad in Dart_ programming language.
          Used in `Pāli Dictionary`_.
:adsu: yes


Virtual kayboard/keypad in Dart_ programming language.

.. rubric:: `Demo <https://dartpad.dartlang.org/77c941818715c4f4b4b2986212fd0af0>`_
   :class: align-center

Real world application is `Pāli Dictionary`_.
There are special characters in `romanized Pāli`_. For the convenience of input
`Pāli`_ words, users can use the virtual keyboard to type Pāli_ word without
installation of `Pāli`_ input method in computers

**Source code**:

.. show_github_file:: siongui userpages content/code/javascript/virtual-keyboard/index.html

The logic in Dart_ code is simple:

- Create input_ element for each letter in `Pāli`_ alphabet.

  * ``type`` of the input_ element is ``button``.
  * ``value`` of the input_ element is the `Pāli`_ letter.

- Listen to `click event`_ of the input_ element. In the callback_ of
  `click event`_, add the corresponding `Pāli`_ letter to the user input.

.. adsu:: 2
.. show_github_file:: siongui userpages content/code/dart/virtual-keyboard/app.dart
.. show_github_file:: siongui userpages content/code/javascript/virtual-keyboard/keypad.css
.. adsu:: 3

For AngularJS_ version, see [1]_.

For Vue.js_ version, see [2]_.

For plain JavaScript_ version, see [3]_.

----

Tested on: DartPad_.

----

References:

.. [1] `[AngularJS] Virtual Keyboard <{filename}../20/angularjs-ng-virtual-keypad%en.rst>`_
.. [2] `[Vue.js] Virtual Keyboard <{filename}../21/vuejs-virtual-keypad%en.rst>`_
.. [3] `[JavaScript] Virtual Keyboard <{filename}../28/javascript-virtual-keypad%en.rst>`_

.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _Vue.js: https://vuejs.org/
.. _AngularJS: https://angularjs.org/
.. _Dart: https://www.dartlang.org/
.. _DartPad: https://dartpad.dartlang.org/
.. _Pāli Dictionary: http://dictionary.sutta.org/
.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _romanized Pāli: https://www.google.com/search?q=romanized+P%C4%81li
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _input: https://api.dartlang.org/stable/dart-html/InputElement-class.html
.. _click event: https://www.google.com/search?q=dart+onclick+event
.. _callback: https://www.google.com/search?q=dart+callback
