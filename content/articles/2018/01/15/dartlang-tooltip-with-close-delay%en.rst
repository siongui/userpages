[Dart] Tooltip with Close Delay
###############################

:date: 2018-01-15T08:07+08:00
:tags: Dart, Tooltip, mouseenter, mouseleave, element offset, DOM, CSS
       element position
:category: Dart
:summary: Show tooltip when the cursor hovers over the text, and close the
          tooltip with delay if the cursor is not in the tooltip
          via *Dart* programming language.
:og_image: http://www.cssportal.com/blog/wp-content/uploads/2013/05/tooltip.png
:adsu: yes


See demo first.
Move the cursor (mouse pointer) to hover over the blue text with underline:

.. rubric:: `Demo <https://dartpad.dartlang.org/509a9fa17e355c230c6fa3a6b896ef8a>`_
   :class: align-center

If the cursor leaves the blue text and does not enter the tooltip within 250ms,
the tooltip will be closed. If the cursor leaves the blue text and is in the
tooltip after 250ms, the tooltip will still be visible without being closed.

The demo here is not the same as the implementation of simple tooltip in my
previous post [2]_. In my previous post, the tooltip will be closed right after
the cursor leaves the blue text without any delay.

To achieve the effect of the demo, first wrap the text with proper HTML tag. In
the example of demo, we wrap the text **tooltip** and **Dart** with *span* tag.
Then put the explanation of the text in the *data-descr* attribute of the tag,
just like the following HTML code used in the demo:

**HTML**:

.. code-block:: html

  This is a <span data-descr="hello, I am tooltip!">tooltip</span> example via
  <span data-descr="programming language from Google">Dart</span>.

Next, add the following rules to your CSS code:

**CSS**:

.. code-block:: css

  *[data-descr] {
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

1. First rule says add underline to the text wrapped in the element with
   *data-descr* attribute, make the text blue, and make the mouse cursor become
   help icon on hover over.
2. Second rule, as the name suggests, is used to make element invisible.
3. Third rule is used to style the tooltip box that shows the description of the
   text.

.. adsu:: 2

Then add the following to your Dart code:

**Dart**:

.. code-block:: dart

  import 'dart:html';
  import 'dart:async';

  DivElement tooltip;

  // indicate if the mouse cursor is in the tooltip
  bool isCursorInTooltip = false;

  // when cursor leaves the text, the delay time to close the tooltip if the
  // cursor is not in the tooltip. (milisecond)
  const delay = const Duration(milliseconds: 250);

  // event handler for mouseenter event of elements with data-descr attribute
  void ShowTooltip(Event e) {
    Element elm = e.target;
    tooltip.style.left = elm.offsetLeft.toString() + "px";
    tooltip.style.top = (elm.offsetTop + elm.offsetHeight + 5).toString() + "px";
    tooltip.text = elm.dataset["descr"];
    tooltip.classes.remove("invisible");
  }

  // event handler for mouseleave event of elements with data-descr attribute
  void HideTooltip(Event e) {
    new Timer(delay, () {
      if (!isCursorInTooltip) {
        tooltip.classes.add("invisible");
      }
    });
  }

  void main() {
    // create and append invisible tooltip to DOM tree
    tooltip = document.createElement("div");
    tooltip.classes.add("tooltip");
    tooltip.classes.add("invisible");
    tooltip.onMouseEnter.listen((Event e) {
      isCursorInTooltip = true;
    });
    tooltip.onMouseLeave.listen((Event e) {
      isCursorInTooltip = false;
      tooltip.classes.add("invisible");
    });
    document.body.append(tooltip);

    // select all elements with data-descr attribute
    // and attach mouseenter and mouseleave event handler
    var els = querySelectorAll('*[data-descr]');
    for (var el in els) {
      el.onMouseEnter.listen(ShowTooltip);
      el.onMouseLeave.listen(HideTooltip);
    }
  }

In the Dart code:

1. Create and append a tooltip (*div* element) to the HTML *body*. The tooltip
   is used to show description on mouse hovering over the text. The tooltip is
   invisible in the beginning.
2. Attach *mouseenter* and *mouseleave* event handler to the tooltip. In the
   event handler, set the variable *isCursorInTooltip* accordingly. Also close
   the tooltip if the cursor leaves the tooltip.
3. Define *mouseenter* and *mouseleave* event handler to texts wrapped in the
   element with *data-descr* attribute. In *mouseenter* handler we show the
   description of the text in the tooltip, and in *mouseleave* handler we make
   the tooltip invisible if the cursor is not in the tooltip after 250ms.
4. Use querySelectorAll_ to find all elements with *data-descr* attribute, and
   attach corresponding mouseenter_/mouseleave_ event handlder to show/hide
   description of texts.

.. adsu:: 3

This post is actually the Dart translation of JavaScript implememtation [1]_.

----

Tested on: DartPad_.

----

**References**:

.. [1] `[JavaScript] Tooltip with Close Delay <{filename}../../../2018/01/08/javascript-tooltip-with-close-delay%en.rst>`_
.. [2] `[Dart] Tooltip <{filename}../../../2017/04/01/dartlang-tooltip%en.rst>`_
.. [3] `How do I do the equivalent of setTimeout + clearTimeout in Dart? - Stack Overflow <https://stackoverflow.com/questions/27063312/how-do-i-do-the-equivalent-of-settimeout-cleartimeout-in-dart>`_

.. _tooltip: https://www.google.com/search?q=tooltip
.. _querySelectorAll: https://www.google.com/search?q=dartlang+querySelectorAll
.. _mouseenter: https://developer.mozilla.org/en/docs/Web/Events/mouseenter
.. _mouseleave: https://developer.mozilla.org/en/docs/Web/Events/mouseleave
.. _dartlang: https://www.dartlang.org/
.. _DartPad: https://dartpad.dartlang.org/
