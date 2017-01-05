[GopherJS] null Test
####################

:date: 2017-01-05T08:45+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM
:category: GopherJS
:summary: Test if a variable is null_ via GopherJS_.


In JavaScript_, we test if a variable is null_ by:

.. code-block:: javascript

  if (event.state === null) {
    // do something
  }

In GopherJS_, the equivalent is:

.. code-block:: go

  if event.Get("state") == nil {
  	// do something
  }

For undefined_ test (if a variable is undefined), see [3]_.

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

.. [2] `javascript null test - Google search <https://www.google.com/search?q=javascript+null+test>`_

       `javascript null test - DuckDuckGo search <https://duckduckgo.com/?q=javascript+null+test>`_

       `javascript null test - Bing search <https://www.bing.com/search?q=javascript+null+test>`_

       `javascript null test - Yahoo search <https://search.yahoo.com/search?p=javascript+null+test>`_

       `javascript null test - Baidu search <https://www.baidu.com/s?wd=javascript+null+test>`_

       `javascript null test - Yandex search <https://www.yandex.com/search/?text=javascript+null+test>`_

       `JavaScript null check - Stack Overflow <http://stackoverflow.com/questions/16672743/javascript-null-check>`_

.. [3] `[Golang] undefined Test in GopherJS <{filename}../../../2016/02/06/go-undefined-test-in-gopherjs%en.rst>`_


.. _GopherJS: http://www.gopherjs.org/
.. _null: https://developer.mozilla.org/en/docs/Web/JavaScript/Reference/Global_Objects/null
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _undefined: https://developer.mozilla.org/en/docs/Web/JavaScript/Reference/Global_Objects/undefined

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
