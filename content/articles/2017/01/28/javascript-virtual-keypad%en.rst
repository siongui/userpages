[JavaScript] Virtual Keyboard
#############################

:date: 2017-01-28T04:19+08:00
:tags: JavaScript, IME, Pāli Input Method, Virtual Keyboard/Keypad
:category: JavaScript
:summary: Virtual kayboard/keypad in vanilla JavaScript_.
          Used in `Pāli Dictionary`_.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


Virtual kayboard/keypad in vanilla JavaScript_ [1]_ (plain JavaScript_ without
any additional library/framework like jQuery_, Vue.js_, or AngularJS_).

.. rubric:: `Demo <{filename}/code/javascript/virtual-keyboard/index.html>`_
            (`JSFiddle <https://jsfiddle.net/w3qysaax/>`__)
   :class: align-center

Real world application is `Pāli Dictionary`_.
There are special characters in `romanized Pāli`_. For the convenience of input
`Pāli`_ words, users can use the virtual keyboard to type Pāli_ word without
installation of `Pāli`_ input method in computers

**Source code**:

.. show_github_file:: siongui userpages content/code/javascript/virtual-keyboard/index.html

The logic in JavaScript_ code is simple:

- Create input_ element for each letter in `Pāli`_ alphabet.

  * ``type`` of the input_ element is ``button``.
  * ``value`` of the input_ element is the `Pāli`_ letter.

- Listen to `click event`_ of the input_ element. In the callback_ of
  `click event`_, add the corresponding `Pāli`_ letter to the user input.

.. adsu:: 2
.. show_github_file:: siongui userpages content/code/javascript/virtual-keyboard/app.js
.. show_github_file:: siongui userpages content/code/javascript/virtual-keyboard/keypad.css
.. adsu:: 3

| For AngularJS_ version, see [3]_.
| For Vue.js_ version, see [4]_.
| For Dart_ version, see [5]_.
| For GopherJS_ version, see [6]_.

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `javascript - What is VanillaJS? - Stack Overflow <http://stackoverflow.com/questions/20435653/what-is-vanillajs>`_
.. [2] `javascript main - Google search <https://www.google.com/search?q=javascript+main>`_

       `javascript main - DuckDuckGo search <https://duckduckgo.com/?q=javascript+main>`_

       `javascript main - Bing search <https://www.bing.com/search?q=javascript+main>`_

       `javascript main - Yahoo search <https://search.yahoo.com/search?p=javascript+main>`_

       `javascript main - Baidu search <https://www.baidu.com/s?wd=javascript+main>`_

       `javascript main - Yandex search <https://www.yandex.com/search/?text=javascript+main>`_

       `Why doesn't JavaScript need a main() function? - Stack Overflow <http://stackoverflow.com/questions/9015836/why-doesnt-javascript-need-a-main-function>`_
.. [3] `[AngularJS] Virtual Keyboard <{filename}../20/angularjs-ng-virtual-keypad%en.rst>`_
.. [4] `[Vue.js] Virtual Keyboard <{filename}../21/vuejs-virtual-keypad%en.rst>`_
.. [5] `[Dart] Virtual Keyboard <{filename}../29/dartlang-virtual-keypad%en.rst>`_
.. [6] `[GopherJS] Virtual Keyboard <{filename}../31/gopherjs-virtual-keypad%en.rst>`_

.. _jQuery: https://jquery.com/
.. _Vue.js: https://vuejs.org/
.. _AngularJS: https://angularjs.org/
.. _Dart: https://www.dartlang.org/
.. _GopherJS: http://www.gopherjs.org/
.. _Pāli Dictionary: http://dictionary.sutta.org/
.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _romanized Pāli: https://www.google.com/search?q=romanized+P%C4%81li
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _input: http://www.w3schools.com/tags/tag_input.asp
.. _click event: https://developer.mozilla.org/en/docs/Web/Events/click
.. _callback: http://javascriptissexy.com/understand-javascript-callback-functions-and-use-them/
