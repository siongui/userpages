[GopherJS] Setting Implementation via JSON and Web Storage (localStorage)
#########################################################################

:date: 2017-01-01T23:28+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM, Web Storage, JSON
:category: GopherJS
:summary: Implementation of setting feature in web application via JSON_,
          `Web Storage`_, and GopherJS_.


Implement setting feature to save user preferences in web application via JSON_,
`Web Storage`_, and GopherJS_.

.. show_github_file:: siongui userpages content/code/gopherjs/localStorage-setting/index.html

FIXME: There is bug in the code below. The ``json.Marshal`` cannot convert
``Setting`` type correctly. Still try to figure out what's going wrong.

.. show_github_file:: siongui userpages content/code/gopherjs/localStorage-setting/app.go

To see demo: use GopherJS_ to compile ``app.go`` to ``app.js``. Put
``index.html`` and ``app.js`` in the same directory. Open ``index.html`` with
your browser.

To implement setting feature via JavaScript_, see [3]_.

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

.. [2] `pali/setting.go at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/go/gopherjs/setting.go>`_

       `pali/json.go at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/go/gopherjs/json.go>`_

       `pali/dictionary.go at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/go/lib/dictionary.go>`_

       `pali/setting.html at master · siongui/pali · GitHub <https://github.com/siongui/pali/blob/master/go/theme/template/includes/setting.html>`_

.. [3] `[JavaScript] Setting Implementation via JSON and Web Storage (localStorage) <{filename}../16/javascript-implement-setting-via-json-and-localStorage%en.rst>`_

.. _GopherJS: http://www.gopherjs.org/
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _Web Storage: https://www.google.com/search?q=Web+Storage+HTML5
.. _JSON: https://www.google.com/search?q=JSON

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
