[Dart] Draggable (Movable) Element
##################################

:date: 2015-02-17 03:54
:tags: Dart, DOM, draggable
:category: Dart
:summary: Draggable, movable (drag and drop) HTML element using Dart programming language.
:adsu: yes


This post is Dart_ version of my previous posts [1]_ and [2]_.

If you need draggable element using pure (vanilla) JavaScript, see [1]_.

If you need draggable element in AngularJS_ way, see [2]_.

Dartium_ is needed for the demo.

.. rubric:: `Demo <{filename}/code/dart-draggable-element/dart-draggable-element.html>`_
   :class: align-center

Development Environment: ``Ubuntu Linux 14.10``, ``Dart 1.8``.

Source Code for Demo (*HTML*):

.. show_github_file:: siongui userpages content/code/dart-draggable-element/dart-draggable-element.html

Source Code for Demo (*Dart*):

.. show_github_file:: siongui userpages content/code/dart-draggable-element/app.dart

Makefile for automating the development:

.. show_github_file:: siongui userpages content/code/dart-draggable-element/Makefile

.. note::

  Dart MouseEvent *clientX* and *clientY* are deprecated. If you use *clientX*
  or *clientY*, the compiled JavaScript will not work. see [3]_ for detail.

----

References:

.. [1] `JavaScript Drag and Drop (Draggable, Movable) Element without External Library <{filename}../../../2012/07/13/javascript-drag-and-drop-draggable-movable-element%en.rst>`_

.. [2] `[AngularJS] Draggable (Movable) Element <{filename}../../../2013/04/04/angularjs-draggable-movable-element%en.rst>`_

.. [3] `[Dart] MouseEvent ClientX and ClientY Deprecated <{filename}../16/dart-MouseEvent-clientX-clientY-deprecated%en.rst>`_

.. [4] `[Golang] Draggable (Movable) Element by GopherJS <{filename}../../../2016/01/17/go-draggable-movable-element-by-gopherjs%en.rst>`_


.. _AngularJS: https://angularjs.org/

.. _Dart: https://www.dartlang.org/

.. _Dartium: https://www.dartlang.org/tools/dartium/
