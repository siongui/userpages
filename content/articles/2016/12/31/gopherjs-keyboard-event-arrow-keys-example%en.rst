[GopherJS] Keyboard Event - Arrow Keys Example
##############################################

:date: 2016-12-31T21:54+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM, Keyboard Event
:category: GopherJS
:summary: Detect arrow keystrokes via GopherJS_.
:adsu: yes


Keyboard event of `arrow keys`_ via GopherJS_.

.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/arrow-keys/index.html
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/gopherjs-dom/src/arrow-keys/app.go
.. adsu:: 3

The *keyup* event of window_ object (which is `js.Global`_ in GopherJS_) is
listened. You can also bind the event listener to input_ element or textarea_
element depending on the needs of your application.


To see demo: use GopherJS_ to compile ``app.go`` to ``app.js``. Put
``index.html`` and ``app.js`` in the same directory. Open ``index.html`` with
your browser.

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Go 1.7.4``
- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)
.. adsu:: 4
.. [2] | `javascript keycode - Google search <https://www.google.com/search?q=javascript+keycode>`_
       | `javascript keycode - DuckDuckGo search <https://duckduckgo.com/?q=javascript+keycode>`_
       | `javascript keycode - Ecosia search <https://www.ecosia.org/search?q=javascript+keycode>`_
       | `javascript keycode - Bing search <https://www.bing.com/search?q=javascript+keycode>`_
       | `javascript keycode - Yahoo search <https://search.yahoo.com/search?p=javascript+keycode>`_
       | `javascript keycode - Baidu search <https://www.baidu.com/s?wd=javascript+keycode>`_
       | `javascript keycode - Yandex search <https://www.yandex.com/search/?text=javascript+keycode>`_
       | `JavaScript Event KeyCodes <http://keycode.info/>`_

.. _GopherJS: http://www.gopherjs.org/
.. _arrow keys: https://www.google.com/search?q=arrow+keys
.. _window: http://www.w3schools.com/js/js_window.asp
.. _js.Global: https://godoc.org/github.com/gopherjs/gopherjs/js#Object
.. _input: http://www.w3schools.com/tags/tag_input.asp
.. _textarea: http://www.w3schools.com/tags/tag_textarea.asp

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
