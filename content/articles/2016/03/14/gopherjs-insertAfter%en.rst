[GopherJS] insertAfter - Insert New Node After Reference Node
#############################################################

:date: 2016-03-14T00:15+08:00
:tags: Go, Golang, GopherJS, DOM, Go to JavaScript
:category: GopherJS
:summary: insertAfter_ - Insert new node after reference node via GopherJS_.
          (opposite of insertBefore_)
:adsu: yes

insertAfter_ - Insert new node after reference node via GopherJS_.
(opposite of insertBefore_).

.. code-block:: go

  referenceNode.Get("parentNode").Call("insertBefore", newNode, referenceNode.Get("nextSibling"))

This comes from answer of [2]_.

----

Tested on: ``Ubuntu Linux 15.10``, ``Go 1.6``.

----

References:

.. [1] `GopherJS - A compiler from Go to JavaScript <http://www.gopherjs.org/>`_
       (`GitHub <https://github.com/gopherjs/gopherjs>`__,
       `GopherJS Playground <http://www.gopherjs.org/playground/>`_,
       |godoc|)

.. [2] `How to do insert After() in JavaScript without using a library? - Stack Overflow <http://stackoverflow.com/questions/4793604/how-to-do-insert-after-in-javascript-without-using-a-library>`_

.. _GopherJS: http://www.gopherjs.org/
.. _insertAfter: http://stackoverflow.com/questions/4793604/how-to-do-insert-after-in-javascript-without-using-a-library
.. _insertBefore: https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore

.. |godoc| image:: https://godoc.org/github.com/gopherjs/gopherjs/js?status.png
   :target: https://godoc.org/github.com/gopherjs/gopherjs/js
