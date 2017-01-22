[Dart] Online Input Method - Pali (Pāli, Pāḷi)
##############################################

:date: 2015-02-23 08:43
:tags: Dart, Pāli Input Method, IME
:category: Dart
:summary: Online Pali (Pāli, Pāḷi) Input Method using Dart programming language.
:adsu: yes


Inspired by `palipad <https://code.google.com/p/palipad/>`_
(`demo of palipad <http://palipad.googlecode.com/git/palipad.html>`_),
I write a Dart_ webapp which help users input `Pāli`_ language online.
The implementation is simple. Create an `input:text`_ element first, and then
listen to the `onKeyup event`_ of the `input:text`_ element. The last two
characters of the `input:text`_ element will be replaced according to the
following rule.

+------+-----+
| Type | For |
+======+=====+
|  AA  |  Ā  |
+------+-----+
|  aa  |  ā  |
+------+-----+
|  II  |  Ī  |
+------+-----+
|  ii  |  ī  |
+------+-----+
|  UU  |  Ū  |
+------+-----+
|  uu  |  ū  |
+------+-----+
|  "N  |  Ṅ  |
+------+-----+
|  "n  |  ṅ  |
+------+-----+
|  .M  |  Ṃ  |
+------+-----+
|  .m  |  ṃ  |
+------+-----+
|  ~N  |  Ñ  |
+------+-----+
|  ~n  |  ñ  |
+------+-----+
|  .T  |  Ṭ  |
+------+-----+
|  .t  |  ṭ  |
+------+-----+
|  .D  |  Ḍ  |
+------+-----+
|  .d  |  ḍ  |
+------+-----+
|  .N  |  Ṇ  |
+------+-----+
|  .n  |  ṇ  |
+------+-----+
|  .L  |  Ḷ  |
+------+-----+
|  .l  |  ḷ  |
+------+-----+

.. rubric:: `Demo <https://siongui.github.io/dart-online-input-method-pali/>`_ (`Fork me on Github <https://github.com/siongui/dart-online-input-method-pali>`_)
      :class: align-center

Source Code for Demo (*HTML*):

.. show_github_file:: siongui dart-online-input-method-pali pali.html

Source Code for Demo (*Dart*):

.. show_github_file:: siongui dart-online-input-method-pali pali.dart

----

References:

.. [1] `events - Respond immediately to textarea changes in Dart - Stack Overflow <http://stackoverflow.com/questions/14433156/respond-immediately-to-textarea-changes-in-dart>`_

.. [2] `dart:core.String API Docs <https://api.dartlang.org/apidocs/channels/stable/dartdoc-viewer/dart:core.String>`_

.. [3] `dart - Use static variable in function() - Stack Overflow <http://stackoverflow.com/questions/22747125/use-static-variable-in-function>`_

.. [4] `javascript - Get cursor position (in characters) within a text Input field - Stack Overflow <http://stackoverflow.com/questions/2897155/get-cursor-position-in-characters-within-a-text-input-field>`_

.. [5] `dart - TextArea cursor position - Stack Overflow <http://stackoverflow.com/questions/22797294/textarea-cursor-position>`_

.. [6] `dart - caret position in an editable div (dartlang) - Stack Overflow <http://stackoverflow.com/questions/21730134/caret-position-in-an-editable-div-dartlang>`_

.. [7] `selection - Setting and getting selectionrange/caret position in contenteditable div element (dart) - Stack Overflow <http://stackoverflow.com/questions/28477487/setting-and-getting-selectionrange-caret-position-in-contenteditable-div-element>`_

.. [8] `Pāli Input Method on Ubuntu Linux <{filename}../../../2012/05/23/pali-input-method-on-ubuntu-linux%en.rst>`_

.. [9] `[Golang] Online Input Method (Pāli) by GopherJS <{filename}../../../2016/01/12/go-online-input-method-pali-by-gopherjs%en.rst>`_

.. _Dart: https://www.dartlang.org/

.. _Pāli: http://en.wikipedia.org/wiki/Pali

.. _input\:text: http://www.w3schools.com/tags/tag_input.asp

.. _onKeyup event: http://www.w3schools.com/jsref/event_onkeyup.asp
