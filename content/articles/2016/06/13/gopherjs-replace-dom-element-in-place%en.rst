[GopherJS] Replace DOM Element in Place
#######################################

:date: 2016-06-13T20:28+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM
:category: GopherJS
:summary: Replace DOM_ element in place via GopherJS_.
:adsu: yes


Replace DOM_ element in place via GopherJS_.

**Question**

Replace element of id ``foo`` with another element of id ``foo2``

**Answer**

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  element := js.Global.Get("document").Call("getElementById", "foo")
  element2 := js.Global.Get("document").Call("getElementById", "foo2")
  element.Get("parentNode").Call("replaceChild", element2, element)

See [3]_ for more information.

----

.. adsu:: 2

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] | `javascript replace element - Google search <https://www.google.com/search?q=javascript+replace+element>`_
       | `javascript replace element - DuckDuckGo search <https://duckduckgo.com/?q=javascript+replace+element>`_
       | `javascript replace element - Ecosia search <https://www.ecosia.org/search?q=javascript+replace+element>`_
       | `javascript replace element - Bing search <https://www.bing.com/search?q=javascript+replace+element>`_
       | `javascript replace element - Yahoo search <https://search.yahoo.com/search?p=javascript+replace+element>`_
       | `javascript replace element - Baidu search <https://www.baidu.com/s?wd=javascript+replace+element>`_
       | `javascript replace element - Yandex search <https://www.yandex.com/search/?text=javascript+replace+element>`_

.. [3] `How to replace DOM element in place using Javascript? - Stack Overflow <http://stackoverflow.com/questions/843680/how-to-replace-dom-element-in-place-using-javascript>`_


.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _nodeType: https://developer.mozilla.org/en/docs/Web/API/Node/nodeType

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
