[JavaScript] HTML Web History API Example
#########################################

:date: 2017-01-04T03:58+08:00
:tags: JavaScript, DOM, html, HTML5 History API, Web application
:category: JavaScript
:summary: Change browser URL without reloading web pages - Example for how to
          use `HTML history API`_.


Change browser URL without reloading the whole web page - only partially update
the content of the web page, improve the user experience of visiting your
website. This is done via `HTML history API`_.

I simplify the example in the tutorial of CSS-Tricks [3]_. In the example below,
I create three links, and if users click the link, the browser URL will change,
and update the texts correspondingly without reloading. When users move backward
or forward through the history, web page will also be partially updated.

.. rubric:: `Demo <{filename}/code/javascript/history-api/index.html>`_
     :class: align-center

.. show_github_file:: siongui userpages content/code/javascript/history-api/index.html

.. show_github_file:: siongui userpages content/code/javascript/history-api/app.js

For GopherJS_ version, see [4]_.

----

Tested on:

- ``Ubuntu Linux 16.10``
- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `html history api - Google search <https://www.google.com/search?q=html+history+api>`_

       `html history api - DuckDuckGo search <https://duckduckgo.com/?q=html+history+api>`_

       `html history api - Bing search <https://www.bing.com/search?q=html+history+api>`_

       `html history api - Yahoo search <https://search.yahoo.com/search?p=html+history+api>`_

       `html history api - Baidu search <https://www.baidu.com/s?wd=html+history+api>`_

       `html history api - Yandex search <https://www.yandex.com/search/?text=html+history+api>`_

.. [2] `Manipulating the browser history - Web APIs | MDN <https://developer.mozilla.org/en-US/docs/Web/API/History_API>`_

.. [3] `Using the HTML5 History API | CSS-Tricks <https://css-tricks.com/using-the-html5-history-api/>`_

.. [4] `[GopherJS] HTML Web History API Example <{filename}../03/gopherjs-html-web-history-api-example%en.rst>`_


.. _HTML history API: https://www.google.com/search?q=html+history+api
.. _GopherJS: http://www.gopherjs.org/
