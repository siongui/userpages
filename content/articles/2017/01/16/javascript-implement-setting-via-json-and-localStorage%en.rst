[JavaScript] Setting Implementation via JSON and Web Storage (localStorage)
###########################################################################

:date: 2017-01-16T02:37+08:00
:tags: JavaScript, html, DOM, Web Storage, JSON
:category: JavaScript
:summary: Implementation of setting feature in web application via JSON_,
          `Web Storage`_ (localStorage_), and JavaScript_.
:adsu: yes


Implement setting feature to save user preferences in web application via JSON_,
`Web Storage`_ (localStorage_), and JavaScript_.

.. rubric:: `Demo <{filename}/code/javascript/localStorage-setting/index.html>`_
     :class: align-center

**Source Code**:

.. show_github_file:: siongui userpages content/code/javascript/localStorage-setting/index.html
.. adsu:: 2
.. show_github_file:: siongui userpages content/code/javascript/localStorage-setting/app.js

To implement setting feature via GopherJS_, see [2]_.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] `pali/setting.go at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/go/gopherjs/setting.go>`_

       `pali/json.go at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/go/gopherjs/json.go>`_

       `pali/dictionary.go at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/go/lib/dictionary.go>`_

       `pali/setting.html at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/go/theme/template/includes/setting.html>`_

.. [2] `[GopherJS] Setting Implementation via JSON and Web Storage (localStorage) <{filename}../01/gopherjs-implement-setting-via-json-and-localStorage%en.rst>`_

.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _GopherJS: http://www.gopherjs.org/
.. _localStorage: https://www.google.com/search?q=localStorage
.. _Web Storage: https://www.google.com/search?q=Web+Storage+HTML5
.. _JSON: https://www.google.com/search?q=JSON
