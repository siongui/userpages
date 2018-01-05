[JavaScript] Tooltip
####################

:date: 2018-01-06T03:00+08:00
:tags: Tooltip, JavaScript, mouseenter, mouseleave, element offset, DOM, CSS,
       element position
:category: JavaScript
:summary: Simple tooltip implementation via JavaScript.
:og_image: http://freefrontend.com/assets/img/css-tooltips/automation-tooltips-with-simple-data-attributes.png
:adsu: yes


See demo first. Move the cursor (mouse pointer) to hover over the blue text with
underline:

.. raw:: html

  <div style="text-align: center; padding: 2em;">

  This is a <span data-descr="hello, I am tooltip!">tooltip</span> example via
  <span data-descr="programming language">JavaScript</span>.

  </div>

.. raw:: html

  <style>
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
      background-color: yellow;
      border: 1px gray solid;
      border-radius: 3px;
      padding: 3px;
  }
  </style>

.. raw:: html

  <script>
  // create and append invisible tooltip to DOM tree
  var tooltip = document.createElement("div");
  tooltip.classList.add("tooltip");
  tooltip.classList.add("invisible");
  document.querySelector("body").appendChild(tooltip);

  // event handler for mouseenter event of elements with data-descr attribute
  function ShowTooltip(e) {
    var elm = e.target;
    tooltip.textContent = elm.dataset.descr;
    tooltip.style.left = elm.getBoundingClientRect().left + window.pageXOffset + 'px';
    tooltip.style.top = (elm.getBoundingClientRect().top + window.pageYOffset + elm.offsetHeight + 5) + 'px';
    tooltip.classList.remove("invisible");
  }

  // event handler for mouseleave event of elements with data-descr attribute
  function HideTooltip(e) {
    tooltip.classList.add("invisible");
  }

  // select all elements with data-descr attribute and attach mouseenter and mouseleave event handler
  var els = document.querySelectorAll("*[data-descr]");
  for (var i = 0; i < els.length; ++i) {
    var el = els[i];
    el.addEventListener("mouseenter", ShowTooltip);
    el.addEventListener("mouseleave", HideTooltip);
  }
  </script>


To achieve the effect of the demo, first wrap the text with proper HTML tag. In
the example of demo, we wrap the text **tooltip** and **JavaScript** with *span*
tag. Then put the explanation of the text in the *data-descr* attribute of the
tag, just like the following HTML code used in the demo:

**HTML**:

.. code-block:: html

  This is a <span data-descr="hello, I am tooltip!">tooltip</span> example via
  <span data-descr="programming language">JavaScript</span>.


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
      background-color: yellow;
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

Then add the following to your JavaScript code:

**JavaScript**:

.. code-block:: javascript

  // create and append invisible tooltip to DOM tree
  var tooltip = document.createElement("div");
  tooltip.classList.add("tooltip");
  tooltip.classList.add("invisible");
  document.querySelector("body").appendChild(tooltip);

  // event handler for mouseenter event of elements with data-descr attribute
  function ShowTooltip(e) {
    var elm = e.target;
    tooltip.textContent = elm.dataset.descr;
    tooltip.style.left = elm.getBoundingClientRect().left + window.pageXOffset + 'px';
    tooltip.style.top = (elm.getBoundingClientRect().top + window.pageYOffset + elm.offsetHeight + 5) + 'px';
    tooltip.classList.remove("invisible");
  }

  // event handler for mouseleave event of elements with data-descr attribute
  function HideTooltip(e) {
    tooltip.classList.add("invisible");
  }

  // select all elements with data-descr attribute and attach mouseenter and mouseleave event handler
  var els = document.querySelectorAll("*[data-descr]");
  for (var i = 0; i < els.length; ++i) {
    var el = els[i];
    el.addEventListener("mouseenter", ShowTooltip);
    el.addEventListener("mouseleave", HideTooltip);
  }

In the JavaScript code:

1. Create and append a tooltip (*div* element) to the HTML *body*. The tooltip
   is used to show description on mouse hovering over the text. The tooltip is
   invisible in the beginning.
2. Define *mouseenter* and *mouseleave* event handler to texts wrapped in the
   element with *data-descr* attribute. In mouseenter handler we show the
   description of the text in the tooltip, and in mouseleave handler we make the
   tooltip invisible.
3. Use querySelectorAll_ to find all elements with *data-descr* attribute, and
   attach corresponding mouseenter_/mouseleave_ event handlder to show/hide
   description of texts.

----

.. adsu:: 3

Tested on: ``Chromium Version 63.0.3239.84 (Official Build) Built on Ubuntu , running on Ubuntu 17.10 (64-bit)``

----

**References**:

.. [1] `Pure CSS Tooltip and JavaScript Implementation <{filename}../../../2017/03/04/css-only-tooltip-and-javascript-implementation%en.rst>`_
.. [2] `JavaScript DOM Element Position (Scroll Position Included) <{filename}../../../2012/07/01/javascript-dom-element-position-scroll-included%en.rst>`_
.. [3] `[JavaScript] Show Note on Mouse Hovering Over Text <{filename}../../../2018/01/04/javascript-show-annotation-on-mouse-hover-over-text%en.rst>`_

.. _CSS: https://www.google.com/search?q=CSS
.. _tooltip: https://www.google.com/search?q=tooltip
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _querySelectorAll: https://www.google.com/search?q=querySelectorAll
.. _mouseenter: https://developer.mozilla.org/en/docs/Web/Events/mouseenter
.. _mouseleave: https://developer.mozilla.org/en/docs/Web/Events/mouseleave
.. _here: http://agama.buddhason.org/SN/SN0011.htm
