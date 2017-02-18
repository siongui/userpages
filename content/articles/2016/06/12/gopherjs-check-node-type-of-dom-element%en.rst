[GopherJS] Check nodeType of DOM Element
########################################

:date: 2016-06-12T21:38+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, DOM
:category: GopherJS
:summary: Check node type (nodeType_) of DOM_ element via GopherJS_.
:adsu: yes


Check node type (nodeType_) of DOM_ element via GopherJS_.

.. code-block:: go

  import "github.com/gopherjs/gopherjs/js"

  element := js.Global.Get("document").Call("getElementById", "foo")
  nodeType := element.Get("nodeType").Int()
  if nodeType == 1 {
  	// element node
  }
  if nodeType == 3 {
  	// text node
  }
  if nodeType == 8 {
  	// comment node
  }

See [3]_ for more nodeType_.

----

.. adsu:: 2

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] | `javascript nodeType - Google search <https://www.google.com/search?q=javascript+nodeType>`_
       | `javascript nodeType - DuckDuckGo search <https://duckduckgo.com/?q=javascript+nodeType>`_
       | `javascript nodeType - Ecosia search <https://www.ecosia.org/search?q=javascript+nodeType>`_
       | `javascript nodeType - Bing search <https://www.bing.com/search?q=javascript+nodeType>`_
       | `javascript nodeType - Yahoo search <https://search.yahoo.com/search?p=javascript+nodeType>`_
       | `javascript nodeType - Baidu search <https://www.baidu.com/s?wd=javascript+nodeType>`_
       | `javascript nodeType - Yandex search <https://www.yandex.com/search/?text=javascript+nodeType>`_

.. [3] `Node.nodeType - Web APIs | MDN <https://developer.mozilla.org/en/docs/Web/API/Node/nodeType>`_

.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://www.google.com/search?q=DOM
.. _nodeType: https://developer.mozilla.org/en/docs/Web/API/Node/nodeType

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
