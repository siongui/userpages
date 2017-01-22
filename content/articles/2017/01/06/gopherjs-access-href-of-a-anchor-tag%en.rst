[GopherJS] Access href Value of Anchor <a> Tag
##############################################

:date: 2017-01-06T20:33+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM
:category: GopherJS
:summary: Access (raw) href_ value of anchor_ tag (<a>) via GopherJS_.
:adsu: yes


In JavaScript_, the href_ value of anchor_ element/tag (<a>) can be accessed by:

.. code-block:: javascript

  // assume your website is:
  //   https://example.com/
  // assume the a tag is:
  //   <a id="atag" href="/about/">About</a>

  var elm = document.getElementById("atag");
  // return "/about/"
  var href = elm.getAttribute("href");
  // return "https://example.com/about/"
  var href = elm.href;

In GopherJS_, the equivalent is:

.. code-block:: go

  elm := js.Global.Get("document")..Call("getElementById", "atag")
  // return "/about/"
  href := elm.Call("getAttribute", "href").String()
  // return "https://example.com/about/"
  href := elm.Get("href").String()

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

.. [2] `javascript access href - Google search <https://www.google.com/search?q=javascript+access+href>`_

       `javascript access href - DuckDuckGo search <https://duckduckgo.com/?q=javascript+access+href>`_

       `javascript access href - Bing search <https://www.bing.com/search?q=javascript+access+href>`_

       `javascript access href - Yahoo search <https://search.yahoo.com/search?p=javascript+access+href>`_

       `javascript access href - Baidu search <https://www.baidu.com/s?wd=javascript+access+href>`_

       `javascript access href - Yandex search <https://www.yandex.com/search/?text=javascript+access+href>`_

.. [3] `anchor - How to get "raw" href contents in JavaScript - Stack Overflow <http://stackoverflow.com/questions/1550901/how-to-get-raw-href-contents-in-javascript>`_


.. _GopherJS: http://www.gopherjs.org/
.. _href: http://www.w3schools.com/tags/att_a_href.asp
.. _anchor: http://www.w3schools.com/tags/tag_a.asp
.. _JavaScript: https://www.google.com/search?q=JavaScript

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
