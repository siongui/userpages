[GopherJS] HTML Web History API Example
#######################################

:date: 2017-01-03T21:42+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM, HTML5 History API
:category: GopherJS
:summary: Change browser URL without reloading web pages - Example for
          `HTML history API`_ via GopherJS_.


Change browser URL without reloading the whole web page - only partially update
the content of the web page, improve the user experience of visiting your
website. This is done via `HTML history API`_ and GopherJS_.

I simplify the example in the tutorial of CSS-Tricks [4]_. In the example below,
I create three links, and if users click the link, the browser URL will change,
and update the texts correspondingly without reloading. When users move backward
or forward through the history, web page will also be partially updated.

.. rubric:: `Demo <{filename}/code/javascript/history-api/index.html>`_
     :class: align-center

.. show_github_file:: siongui userpages content/code/gopherjs/history-api/index.html

.. show_github_file:: siongui userpages content/code/gopherjs/history-api/app.go

To see demo: use GopherJS_ to compile ``app.go`` to ``app.js``. Put
``index.html`` and ``app.js`` in the same directory. Open ``index.html`` with
your browser.

For JavaScript verson, see [7]_.

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

.. [2] `html history api - Google search <https://www.google.com/search?q=html+history+api>`_

       `html history api - DuckDuckGo search <https://duckduckgo.com/?q=html+history+api>`_

       `html history api - Bing search <https://www.bing.com/search?q=html+history+api>`_

       `html history api - Yahoo search <https://search.yahoo.com/search?p=html+history+api>`_

       `html history api - Baidu search <https://www.baidu.com/s?wd=html+history+api>`_

       `html history api - Yandex search <https://www.yandex.com/search/?text=html+history+api>`_

.. [3] `Manipulating the browser history - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/History_API>`_

.. [4] `Using the HTML5 History API | CSS-Tricks <https://css-tricks.com/using-the-html5-history-api/>`_

.. [5] `[Golang] GopherJS DOM Example - Access HTML Data Attribute <{filename}../../../2016/01/12/gopherjs-dom-example-access-html-data-attribute%en.rst>`_

.. [6] `[Golang] undefined Test in GopherJS <{filename}../../../2016/02/06/go-undefined-test-in-gopherjs%en.rst>`_

.. [7] `[JavaScript] HTML Web History API Example <{filename}../04/javascript-html-web-history-api-example%en.rst>`_


.. _GopherJS: http://www.gopherjs.org/
.. _window.history: https://developer.mozilla.org/en-US/docs/Web/API/History_API
.. _HTML history API: https://www.google.com/search?q=html+history+api

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
