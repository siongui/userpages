[Dart] MouseEvent ClientX and ClientY Deprecated
################################################

:tags: Dart
:category: Dart
:summary: e.clientX => e.client.x & e.clientY => e.client.y


Recently I am porting JavaScript code to Dart. When I write the following code:

.. code-block:: dart

    int dx = e.clientX - initialMouseX;

It works on Dartium_. Then I compiled the code to JavaScript by dart2js_. The
compiled JavaScript code failed to run on Chromium_. After some Googling, I
found *Dart MouseEvent clientX and clientY* **deprecated** ([1]_, [2]_, [4]_).
And I also found someone got the same problem as me ([3]_). The correct way to
get *clientX* and *clientY* in Dart should be:

.. code-block:: dart

    int dx = e.client.x - initialMouseX;

Then the compiled JavaScript code runs on Chromium_ without trouble.


PS: Tested on ``Dart 1.8``

----

References:

.. [1] `MouseEvent ClientX and ClientY deprecated <https://github.com/threeDart/three.dart/issues/109>`_

.. [2] `Breaking Change: dart:html Point & Rect <https://groups.google.com/a/dartlang.org/d/topic/misc/DNgsK6Qbd6I>`_

.. [3] `JS MouseEvent unexpectedly converted to Dart MouseEvent in interop <https://code.google.com/p/dart/issues/detail?id=15216>`_

.. [4] `dart/tools/dom/templates/html/dart2js/impl_MouseEvent.darttemplate - external/dart - Git at Google <https://chromium.googlesource.com/external/dart/+/f0d085ba55f544c9338232f1ef0bbe1e08675310/dart/tools/dom/templates/html/dart2js/impl_MouseEvent.darttemplate>`_

.. _Dartium: https://www.dartlang.org/tools/dartium/

.. _dart2js: https://www.dartlang.org/tools/dart2js/

.. _Chromium: http://www.chromium.org/
