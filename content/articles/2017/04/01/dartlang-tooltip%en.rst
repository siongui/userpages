[Dart] Tooltip
##############

:date: 2017-04-01T22:55+08:00
:tags: Dart, Tooltip, mouseenter, mouseleave, element offset, DOM,
       element position
:category: Dart
:summary: Simple *tooltip* implementation via *Dart* programming language.
:og_image: http://www.cssportal.com/blog/wp-content/uploads/2013/05/tooltip.png
:adsu: yes

Simple tooltip_ implementation via dartlang_.

.. rubric:: `Demo <https://dartpad.dartlang.org/3505fc58bacfb4dd6b8fb53dc93e0ff9>`_
   :class: align-center

The idea comes from the JavaScript implementation of my post [1]_.
The following is complete source code.

**HTML**:

.. code-block:: html


  <div class="tooltip invisible"></div>
  This is a <span data-text="hello, I am tooltip!">tooltip</span> example via
  <span data-text="programming language from Google">Dart</span>.

The texts in *data-text* attribute of the span element is the texts to be
displayed in the tooltip when the cursor hovers over the span element.

**CSS**:

.. code-block:: css

  span[data-text] {
    text-decoration: underline;
    color: blue;
    cursor: help;
  }

  .invisible {
    display: none;
  }

  .tooltip {
    position: absolute;
    border: 1px gray solid;
    border-radius: 3px;
    padding: 3px;
  }


**Dart**:

.. code-block:: dart

  import 'dart:html';

  DivElement tooltip = querySelector(".tooltip");

  void ShowTooltip(Event e) {
    Element elm = e.target;
    tooltip.style.left = elm.offsetLeft.toString() + "px";
    tooltip.style.top = (elm.offsetTop + elm.offsetHeight + 5).toString() + "px";
    tooltip.text = elm.dataset["text"];
    tooltip.classes.remove("invisible");
  }

  void HideTooltip(Event e) {
    tooltip.classes.add("invisible");
  }

  void main() {
    var spans = querySelectorAll('span[data-text]');
    for (var span in spans) {
      span.onMouseEnter.listen(ShowTooltip);
      span.onMouseLeave.listen(HideTooltip);
    }
  }

Use querySelectorAll_ to find all span elements with *data-text* attibute, and
setup corresponding mouseenter_/mouseleave_ event handlder to show/hide the
tooltip.

----

Tested on: DartPad_.

----

**References**:

.. [1] `Pure CSS Tooltip and JavaScript Implementation <{filename}../../03/04/css-only-tooltip-and-javascript-implementation%en.rst>`_
.. [2] `[Dart] Access HTML Data Attribute <{filename}../../../2015/03/01/dart-access-html-data-attribute%en.rst>`_

.. _tooltip: https://www.google.com/search?q=tooltip
.. _querySelectorAll: https://www.google.com/search?q=dartlang+querySelectorAll
.. _mouseenter: https://developer.mozilla.org/en/docs/Web/Events/mouseenter
.. _mouseleave: https://developer.mozilla.org/en/docs/Web/Events/mouseleave
.. _dartlang: https://www.dartlang.org/
.. _DartPad: https://dartpad.dartlang.org/
