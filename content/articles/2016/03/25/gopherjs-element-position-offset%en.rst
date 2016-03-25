[GopherJS] Element Position (Offset)
####################################

:date: 2016-03-25T20:21+08:00
:tags: Go, Golang, GopherJS, DOM, element offset, element position,
       Go to JavaScript
:category: GopherJS
:summary: Get DOM_ element position (offset) via GopherJS_.


Get DOM_ element position (offset) via GopherJS_.

Install GopherJS_ first:

.. code-block:: bash

  $ go get -u github.com/gopherjs/gopherjs

Source code for element position (offset):

.. code-block:: go

  package gojs

  import "github.com/gopherjs/gopherjs/js"

  type DOMClientRect struct {
  	Bottom string
  	Height string
  	Left   string
  	Right  string
  	Top    string
  	Width  string
  }

  func GetPosition(elm *js.Object) DOMClientRect {
  	rect := elm.Call("getBoundingClientRect")
  	return DOMClientRect{
  		Bottom: rect.Get("bottom").String(),
  		Height: rect.Get("height").String(),
  		Left:   rect.Get("left").String(),
  		Right:  rect.Get("right").String(),
  		Top:    rect.Get("top").String(),
  		Width:  rect.Get("width").String(),
  	}
  }

This comes from answer of Stack Overflow question [3]_.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `GitHub - siongui/gopherjs-utils: useful collections of utilites (functions) for front end (browser) development via GopherJS <https://github.com/siongui/gopherjs-utils>`_

.. [3] `javascript element position <https://www.google.com/search?q=javascript+element+position>`_

       `javascript - Retrieve the position (X,Y) of an HTML element - Stack Overflow <http://stackoverflow.com/questions/442404/retrieve-the-position-x-y-of-an-html-element>`_

.. [4] `[AngularJS] Get Element Offset (Position) <{filename}../../../2013/05/12/angularjs-get-element-offset-position%en.rst>`_

.. [5] `JavaScript DOM Element Position (Scroll Position Included) <{filename}../../../2012/07/01/javascript-dom-element-position-scroll-included%en.rst>`_


.. _GopherJS: http://www.gopherjs.org/
.. _DOM: https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
