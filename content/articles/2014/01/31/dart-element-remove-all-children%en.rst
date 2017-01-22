[Dart] DOM Element Remove All Children
######################################

:modified: 2015-02-20 12:31
:tags: Dart, DOM
:category: Dart
:summary: DOM element remove all children in Dart programming language
:adsu: yes


Question:

  Assume that `elm` is a dom element, how to remove all children elements
  (nodes) of `elm`?

Answer:

.. code-block :: dart

  elm.children.clear();


----

References:

.. [1] `Removing all child elements from an element <https://www.dartlang.org/docs/tutorials/remove-elements/#remove-all-elem>`_

.. [2] `JavaScript Remove All Children of a DOM Element <{filename}../../../2012/09/26/javascript-remove-all-children-of-dom-element%en.rst>`_
