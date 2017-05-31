[Dart] Visit All Sibling Elements or Nodes
##########################################

:date: 2017-05-31T21:01+08:00
:tags: Dart, html, DOM, Traverse DOM Tree
:category: Dart
:summary: DOM manipulation - traverse all sibling elements or nodes via Dart.
:og_image: https://i.stack.imgur.com/9zp3z.png
:adsu: yes

DOM manipulation - traverse all sibling elements or nodes via Dart.

Visit All Sibling Elements
++++++++++++++++++++++++++

Use nextElementSibling_ and previousElementSibling_:

.. code-block:: dart

  var sibling = elm.nextElementSibling;
  while (sibling != null) {
    // do something with the sibling element

    sibling = sibling.nextElementSibling;
  }
  sibling = elm.previousElementSibling;
  while (sibling != null) {
    // do something with the sibling element

    sibling = sibling.previousElementSibling;
  }


Visit All Sibling Nodes
+++++++++++++++++++++++

Use nextSibling_ and previousSibling_:

.. code-block:: dart

  var sibling = node.nextSibling;
  while (sibling != null) {
    // do something with the sibling node

    sibling = sibling.nextSibling;
  }
  sibling = node.previousSibling;
  while (sibling != null) {
    // do something with the sibling node

    sibling = sibling.previousSibling;
  }

.. adsu:: 2

Difference Between Element and Node
+++++++++++++++++++++++++++++++++++

In short, nodes include pure texts, comments, or elements. And elements are
something like *div*, *p*, or any other HTML tags. Elements are one kind of
nodes. For more details, see [1]_ and [2]_


Example
+++++++

See my another post ``[Dart] Tab Panel`` [3]_. We traverse the sibling elements
of tab pane to remove the ``.active`` class of other tab pane.

----

Tested on:

- DartPad_
- ``Chromium Version 58.0.3029.110 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] `javascript - Whats the difference between nextElementSibling vs nextSibling - Stack Overflow <https://stackoverflow.com/questions/31097016/whats-the-difference-between-nextelementsibling-vs-nextsibling>`_
.. [2] `HTML DOM nextElementSibling Property <https://www.w3schools.com/jsref/prop_element_nextelementsibling.asp>`_
.. [3] `[Dart] Tab Panel <{filename}../29/dartlang-tab-panel%en.rst>`_
.. [4] `[JavaScript] Visit All Sibling Elements or Nodes <{filename}../30/javascript-traverse-all-sibling-element-or-node%en.rst>`_

.. _nextElementSibling: https://developer.mozilla.org/en-US/docs/Web/API/NonDocumentTypeChildNode/nextElementSibling
.. _previousElementSibling: https://developer.mozilla.org/en-US/docs/Web/API/NonDocumentTypeChildNode/previousElementSibling
.. _nextSibling: https://developer.mozilla.org/en-US/docs/Web/API/Node/nextSibling
.. _previousSibling: https://developer.mozilla.org/en-US/docs/Web/API/Node/previousSibling
.. _DartPad: https://dartpad.dartlang.org/
