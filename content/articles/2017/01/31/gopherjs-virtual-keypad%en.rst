[GopherJS] Virtual Keyboard
###########################

:date: 2017-01-31T10:21+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, IME, Pāli Input Method,
       Virtual Keyboard/Keypad
:category: GopherJS
:summary: Virtual kayboard/keypad in GopherJS_. Used in `Pāli Dictionary`_.
:og_image: https://pbs.twimg.com/profile_images/605816243870760960/4hP2sH_O.png
:adsu: yes


Virtual kayboard/keypad in GopherJS_ [1]_.

.. rubric:: `Demo <https://siongui.github.io/gopherjs-pali-virtual-keypad/>`_
   :class: align-center

Real world application is `Pāli Dictionary`_.
There are special characters in `romanized Pāli`_. For the convenience of input
`Pāli`_ words, users can use the virtual keyboard to type Pāli_ word without
installation of `Pāli`_ input method in computers

**Source code**:

The main logic in Go_ code is simple:

- Create input_ element for each letter in `Pāli`_ alphabet.

  * ``type`` of the input_ element is ``button``.
  * ``value`` of the input_ element is the `Pāli`_ letter.

- Listen to `click event`_ of the input_ element. In the callback_ of
  `click event`_, add the corresponding `Pāli`_ letter to the user input.

.. adsu:: 2
.. show_github_file:: siongui gopherjs-pali-virtual-keypad keypad.go

For complete source code + demo code, see [2]_.

| For AngularJS_ version, see [3]_.
| For Vue.js_ version, see [4]_.
| For Dart_ version, see [5]_.
| For plain JavaScript_ version, see [6]_.

.. adsu:: 3

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.5``
- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `GitHub - siongui/gopherjs-pali-virtual-keypad: Virtual Keyboard for Pāli input. Written in Go. <https://github.com/siongui/gopherjs-pali-virtual-keypad>`_

.. [3] `[AngularJS] Virtual Keyboard <{filename}../20/angularjs-ng-virtual-keypad%en.rst>`_
.. [4] `[Vue.js] Virtual Keyboard <{filename}../21/vuejs-virtual-keypad%en.rst>`_
.. [5] `[Dart] Virtual Keyboard <{filename}../29/dartlang-virtual-keypad%en.rst>`_
.. [6] `[JavaScript] Virtual Keyboard <{filename}../28/javascript-virtual-keypad%en.rst>`_

.. _Go: https://golang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _Vue.js: https://vuejs.org/
.. _AngularJS: https://angularjs.org/
.. _Dart: https://www.dartlang.org/
.. _Pāli Dictionary: http://dictionary.sutta.org/
.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _romanized Pāli: https://www.google.com/search?q=romanized+P%C4%81li
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _input: http://www.w3schools.com/tags/tag_input.asp
.. _click event: https://developer.mozilla.org/en/docs/Web/Events/click
.. _callback: http://javascriptissexy.com/understand-javascript-callback-functions-and-use-them/

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
