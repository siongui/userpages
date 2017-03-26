[JavaScript] Toggle (Show/Hide) HTML Element
############################################

:date: 2017-02-26T04:23+08:00
:tags: JavaScript, DOM, toggle, toggleable, html, CSS
:category: JavaScript
:summary: Toggle (Show/Hide) HTML_ `DOM element`_ in JavaScript_.
:og_image: http://www.javatpoint.com/images/javascript/javascript_logo.png
:adsu: yes


Toggle (Show/Hide) HTML_ `DOM element`_ in JavaScript_. See demo first. Please
click the following text to toggle the element.

.. raw:: html

  <br>
  <style>
  .invisible {
    display: none;
  }
  .menu-toggle {
    cursor: pointer;
  }
  .menu {
    width: 200px;
    height: 200px;
    background-color: yellow;
  }
  </style>

  <div style="background-color: Azure; padding: 5px;">
  <div class="menu-toggle">
    <!--
    <span class="invisible">&#9660;</span>
    <span>&#9658;</span>
    -->
    Click me to toggle
  </div>
  <div class="invisible menu">
    menu content
  </div>
  </div>

  <script>
  var menu = document.querySelector(".menu");
  var menuToggle = document.querySelector(".menu-toggle");
  menuToggle.addEventListener("click", function(e) {
    menu.classList.toggle("invisible");
  });
  </script>
  <br>

There are two HTML elements in the demo. The HTML code is as follows:

.. code-block:: html

  <div class="menu-toggle">
    Click me to toggle
  </div>
  <div class="invisible menu">
    menu content
  </div>

When users click on the first *div* element, the second *div* element will
become visible if it is invisible, or become invisible if it is visible. In
other words, the click of first *div* toggles the visibility of the second
*div*. To achieve the toggle effect, we need the following **CSS** classes:

.. code-block:: css

  .invisible {
    display: none;
  }
  .menu-toggle {
    cursor: pointer;
  }
  .menu {
    width: 200px;
    height: 200px;
    background-color: yellow;
  }

.. adsu:: 2

The class ``.invisible`` will hide the elements that have this class. In the
beginning, the second *div* has ``.invisible`` class so it is hidden.
Next we will add *click* event listener to the first *div* element. In the event
listener we will toggle the ``.invisible`` class of second *div* element.

.. code-block:: javascript

  var menu = document.querySelector(".menu");
  var menuToggle = document.querySelector(".menu-toggle");
  menuToggle.addEventListener("click", function(e) {
    menu.classList.toggle("invisible");
  });

It is easy to toggle the ``.invisible`` class of DOM element. The classList_
property of DOM element provides toggle_ method to toggle CSS class of the DOM
element. In the event listener, the ``invisible`` class of second *div* element
is toggled, so it is shown/hidden by the clicks of first *div* element.

----

Tested on:

- ``Chromium Version 55.0.2883.87 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

.. adsu:: 3

References:

.. [1] | `javascript toggle element - Google search <https://www.google.com/search?q=javascript+toggle+element>`_
       | `javascript toggle element - DuckDuckGo search <https://duckduckgo.com/?q=javascript+toggle+element>`_
       | `javascript toggle element - Ecosia search <https://www.ecosia.org/search?q=javascript+toggle+element>`_
       | `javascript toggle element - Qwant search <https://www.qwant.com/?q=javascript+toggle+element>`_
       | `javascript toggle element - Bing search <https://www.bing.com/search?q=javascript+toggle+element>`_
       | `javascript toggle element - Yahoo search <https://search.yahoo.com/search?p=javascript+toggle+element>`_
       | `javascript toggle element - Baidu search <https://www.baidu.com/s?wd=javascript+toggle+element>`_
       | `javascript toggle element - Yandex search <https://www.yandex.com/search/?text=javascript+toggle+element>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _classList: https://www.w3schools.com/jsref/prop_element_classlist.asp
.. _toggle: https://developer.mozilla.org/en/docs/Web/API/Element/classList
