[JavaScript] Show Note on Mouse Hovering Over Text
##################################################

:date: 2018-01-04T04:31+08:00
:modified: 2018-01-05T00:44+08:00
:tags: Tooltip, JavaScript, mouseenter, mouseleave, element offset, DOM,
       element position
:category: JavaScript
:summary: Show annotatoin (note) on mouse hovering over text. Used to help users
          read ancient Buddhist texts.
:og_image: https://acidmartin.files.wordpress.com/2010/12/tooltip-js.gif
:adsu: yes


See demo first. Move the cursor (mouse pointer) to hover over the italic text
with underline:
(demo texts from here_)

| 　　我聽到這樣：
| 　　有一次， *世尊* 住在舍衛城祇樹林給孤獨園。
| 　　在那裡，世尊召喚 *比丘* 們：「比丘們！」

.. raw:: html

  <style>
  em {
    text-decoration: underline;
    cursor: pointer;
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

  <script>
  var notes = {
    "世尊": "Note of 世尊",
    "比丘": "Note of 比丘"
  };

  // create and append invisible tooltip to DOM tree
  var tooltip = document.createElement("div");
  tooltip.classList.add("tooltip");
  tooltip.classList.add("invisible");
  document.querySelector("body").appendChild(tooltip);

  // event handler for mouseenter event of italic text
  function ShowTooltip(e) {
    var elm = e.target;
    var key = elm.textContent;
    if (notes.hasOwnProperty(key)) {
      tooltip.textContent = notes[key];
      tooltip.style.left = elm.getBoundingClientRect().left + window.pageXOffset + 'px';
      tooltip.style.top = (elm.getBoundingClientRect().top + window.pageYOffset + elm.offsetHeight + 5) + 'px';
      tooltip.classList.remove("invisible");
    }
  }

  // event handler for mouseleave event of italic text
  function HideTooltip(e) {
    tooltip.classList.add("invisible");
  }

  // select all italic texts and attach mouseenter and mouseleave event handler
  var ems = document.querySelectorAll("em");
  for (var i = 0; i < ems.length; ++i) {
    var em = ems[i];
    em.addEventListener("mouseenter", ShowTooltip);
    em.addEventListener("mouseleave", HideTooltip);
  }
  </script>


To achieve the effect of the demo, you need to wrap the texts with notes in the
*em* element (i.e., make texts italic) in the HTML.

Next, add the following rules to your CSS code:

**CSS**:

.. code-block:: css

  em {
    text-decoration: underline;
    cursor: pointer;
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

1. First rule says add underline to the text in the *em* element, and also make
   the mouse cursor become pointer on hover over.
2. Second rule, as the name suggests, is used to make element invisible.
3. Third rule is used to style the tooltip box that shows the note of the text.

.. adsu:: 2

Then add the following to your JavaScript code:

**JavaScript**:

.. code-block:: javascript

  var notes = {
    "世尊": "Note of 世尊",
    "比丘": "Note of 比丘"
  };

  // create and append invisible tooltip to DOM tree
  var tooltip = document.createElement("div");
  tooltip.classList.add("tooltip");
  tooltip.classList.add("invisible");
  document.querySelector("body").appendChild(tooltip);

  // event handler for mouseenter event of italic text
  function ShowTooltip(e) {
    var elm = e.target;
    var key = elm.textContent;
    if (notes.hasOwnProperty(key)) {
      tooltip.textContent = notes[key];
      tooltip.style.left = elm.getBoundingClientRect().left + window.pageXOffset + 'px';
      tooltip.style.top = (elm.getBoundingClientRect().top + window.pageYOffset + elm.offsetHeight + 5) + 'px';
      tooltip.classList.remove("invisible");
    }
  }

  // event handler for mouseleave event of italic text
  function HideTooltip(e) {
    tooltip.classList.add("invisible");
  }

  // select all italic texts and attach mouseenter and mouseleave event handler
  var ems = document.querySelectorAll("em");
  for (var i = 0; i < ems.length; ++i) {
    var em = ems[i];
    em.addEventListener("mouseenter", ShowTooltip);
    em.addEventListener("mouseleave", HideTooltip);
  }

In the JavaScript code:

1. First use object literal syntax to define notes of texts.
2. create and append a tooltip (*div* element) to the HTML *body*. The tooltip
   is used to show note on mouse hovering over the italic text. The tooltip is
   invisible in the beginning.
3. Define *mouseenter* and *mouseleave* event handler to texts wrapped in *em*
   element. In mouseenter handler we show the note of the text, and in
   mouseleave handler we make the note of the text invisible.
4. Use querySelectorAll_ to find all *em* elements, and attach corresponding
   mouseenter_/mouseleave_ event handlder to show/hide notes of texts.

----

Tested on: ``Chromium Version 63.0.3239.84 (Official Build) Built on Ubuntu , running on Ubuntu 17.10 (64-bit)``

.. adsu:: 3

----

**References**:

.. [1] `Pure CSS Tooltip and JavaScript Implementation <{filename}../../../2017/03/04/css-only-tooltip-and-javascript-implementation%en.rst>`_
.. [2] `JavaScript DOM Element Position (Scroll Position Included) <{filename}../../../2012/07/01/javascript-dom-element-position-scroll-included%en.rst>`_

.. _CSS: https://www.google.com/search?q=CSS
.. _tooltip: https://www.google.com/search?q=tooltip
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _querySelectorAll: https://www.google.com/search?q=querySelectorAll
.. _mouseenter: https://developer.mozilla.org/en/docs/Web/Events/mouseenter
.. _mouseleave: https://developer.mozilla.org/en/docs/Web/Events/mouseleave
.. _here: http://agama.buddhason.org/SN/SN0011.htm
