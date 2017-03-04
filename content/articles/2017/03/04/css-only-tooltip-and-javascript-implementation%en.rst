Pure CSS Tooltip and JavaScript Implementation
##############################################

:date: 2017-03-04T21:43+08:00
:tags: CSS, Tooltip, JavaScript, mouseenter, mouseleave, element offset, DOM,
       element position
:category: CSS
:summary: CSS_ only tooltip_, extracted from the example of MDN, and the
          corresponding JavaScript_ implementation.
:og_image: http://www.cssportal.com/blog/wp-content/uploads/2013/05/tooltip.png
:adsu: yes

.. contents:: Table of Contents

CSS Only Tooltip
++++++++++++++++

See demo first. Move the cursor (mouse pointer) to hover over the blue text with
underline:

.. raw:: html

  <style>
    span[data-descr] {
      position: relative;
      text-decoration: underline;
      color: blue;
      cursor: help;
    }
    span[data-descr]:hover::after {
      content: attr(data-descr);
      position: absolute;
      left: 0;
      top: 1.25em;
      background-color: yellow;
      border: 1px gray solid;
      border-radius: 3px;
      padding: 3px;
    }
  </style>
  <div style="text-align: center; padding: 2em;">
  This is a <span data-descr="hello, I am tooltip!">tooltip</span> example.
  </div>

I modified the example of MDN [1]_ and make it simpler. The source code is as
follows:

**HTML**:

.. code-block:: html

  This is a <span data-descr="hello, I am tooltip!">tooltip</span> example.

The texts in *data-descr* attribute of the span element is the texts to be
displayed in the tooltip when the cursor hovers over the span element.

**CSS**:

.. code-block:: css

  span[data-descr] {
      position: relative;
      text-decoration: underline;
      color: blue;
      cursor: help;
  }
  span[data-descr]:hover::after {
      content: attr(data-descr);
      position: absolute;
      left: 0;
      top: 1.25em;
      background-color: yellow;
      border: 1px gray solid;
      border-radius: 3px;
      padding: 3px;
  }

.. adsu:: 2

The tooltip is in fact an `::after`_ `pseudo-element`_, and `attr()`_ CSS
expression is used to retrieve the value of an *data-descr* attribute of the
span element, and put the value in the content of the `pseudo-element`_.


JavaScript Implementation
+++++++++++++++++++++++++

See demo first. Move the cursor (mouse pointer) to hover over the blue text with
underline:

.. raw:: html

  <style>
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
      background-color: yellow;
      border: 1px gray solid;
      border-radius: 3px;
      padding: 3px;
  }
  </style>

  <div class="tooltip invisible"></div>
  <div id="js" style="text-align: center; padding: 2em;">
  This is a <span data-text="hello, I am tooltip!">tooltip</span> example via JavaScript.
  </div>

  <script>
  var tooltip = document.querySelector(".tooltip");

  function ShowTooltip(e) {
    var elm = e.target;
    tooltip.style.left = elm.offsetLeft + 'px';
    tooltip.style.top = (elm.offsetTop + elm.offsetHeight + 5) + 'px';
    tooltip.textContent = elm.dataset.text;
    tooltip.classList.remove("invisible");
  }

  function HideTooltip(e) {
    tooltip.classList.add("invisible");
  }

  var spans = document.querySelectorAll("span[data-text]");
  for (var i = 0; i < spans.length; ++i) {
    var span = spans[i];
    span.addEventListener("mouseenter", ShowTooltip);
    span.addEventListener("mouseleave", HideTooltip);
  }
  </script>

This demo is almost the same as the CSS demo, except that it is implemented by
JavaScript_. The following is complete source code.

**HTML**: one more *div* element is inserted and invisible in the beginning.
This *div* wiil be used as the tooltip.

.. code-block:: html

  <div class="tooltip invisible"></div>
  This is a <span data-text="hello, I am tooltip!">tooltip</span> example via JavaScript.

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
      background-color: yellow;
      border: 1px gray solid;
      border-radius: 3px;
      padding: 3px;
  }

.. adsu:: 3

**JavaScript**: Use querySelectorAll_ to find all *span* elements with
*data-text* attibute, and setup corresponding mouseenter_/mouseleave_ event
handlder to show/hide the tooltip.

.. code-block:: javascript

  var tooltip = document.querySelector(".tooltip");

  function ShowTooltip(e) {
    var elm = e.target;
    tooltip.style.left = elm.offsetLeft + 'px';
    tooltip.style.top = (elm.offsetTop + elm.offsetHeight + 5) + 'px';
    tooltip.textContent = elm.dataset.text;
    tooltip.classList.remove("invisible");
  }

  function HideTooltip(e) {
    tooltip.classList.add("invisible");
  }

  var spans = document.querySelectorAll("span[data-text]");
  for (var i = 0; i < spans.length; ++i) {
    var span = spans[i];
    span.addEventListener("mouseenter", ShowTooltip);
    span.addEventListener("mouseleave", HideTooltip);
  }


Summary
+++++++

If only simple texts are in the tooltip, use CSS only solution. If you need to
retrieve the content of tooltip via XHR_ and perform complicated tasks, choose
JavaScript_ implementation.

Tested on: ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

.. adsu:: 4

References
++++++++++

.. [1] `::after (:after) - CSS | MDN <https://developer.mozilla.org/en-US/docs/Web/CSS/::after#Tooltips>`_
.. [2] | `pure css tooltip - Google search <https://www.google.com/search?q=pure+css+tooltip>`_
       | `pure css tooltip - DuckDuckGo search <https://duckduckgo.com/?q=pure+css+tooltip>`_
       | `pure css tooltip - Ecosia search <https://www.ecosia.org/search?q=pure+css+tooltip>`_
       | `pure css tooltip - Bing search <https://www.bing.com/search?q=pure+css+tooltip>`_
       | `pure css tooltip - Yahoo search <https://search.yahoo.com/search?p=pure+css+tooltip>`_
       | `pure css tooltip - Baidu search <https://www.baidu.com/s?wd=pure+css+tooltip>`_
       | `pure css tooltip - Yandex search <https://www.yandex.com/search/?text=pure+css+tooltip>`_

.. _CSS: https://www.google.com/search?q=CSS
.. _tooltip: https://www.google.com/search?q=tooltip
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _pseudo-element: https://developer.mozilla.org/en-US/docs/Web/CSS/Pseudo-elements
.. _\:\:after: https://developer.mozilla.org/en-US/docs/Web/CSS/::after
.. _attr(): https://developer.mozilla.org/en-US/docs/Web/CSS/attr
.. _XHR: https://www.google.com/search?q=javascript+xhr
.. _querySelectorAll: https://www.google.com/search?q=querySelectorAll
.. _mouseenter: https://developer.mozilla.org/en/docs/Web/Events/mouseenter
.. _mouseleave: https://developer.mozilla.org/en/docs/Web/Events/mouseleave
