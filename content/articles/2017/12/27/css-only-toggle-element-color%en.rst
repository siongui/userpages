Pure CSS Toggle (Change) Color of HTML Element
##############################################

:date: 2017-12-27T23:09+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS
:category: CSS
:summary: CSS_ only toggle the color of HTML_ `DOM element`_. No
          JavaScript_ required.
:og_image: https://www.gavick.com/documentation/wp-content/uploads/2013/09/colors-table.png
:adsu: yes


.. role:: raw-html(raw)
   :format: html

CSS_ only toggle the color of HTML_ `DOM element`_. No JavaScript_ required.
See demo first.
Please click the following text to toggle the color of the div element:

.. raw:: html

  <style>
  label[for=element-toggle] {
    cursor: pointer;
    color: blue;
  }
  #element-toggle {
    display: none;
  }
  #toggled-element {
    padding: 1rem;
  }
  #element-toggle:checked ~ #toggled-element {
    color: red;
  }
  </style>

  <div>

  <label for="element-toggle">click me to toggle color</label>
  <input id="element-toggle" type="checkbox" />
  <br><br>
  <div id="toggled-element">My color will be toggled</div>

  </div>

The color is black in the beginning, if you click on the text, the color will be
changed to red. If you click again, the color will be changed back to black, and
so on.

This demo uses the technique of CSS toggling in [1]_.
There are three main HTML elements in the demo:

.. code-block:: html

  <label for="element-toggle">click me to toggle color</label>
  <input id="element-toggle" type="checkbox" />

  <div id="toggled-element">My color will be toggled</div>

1. A visible HTML label_ element
2. A invisible HTML input_ checkbox element, referenced by the above label_
   element via for_ attribute.
3. The element to be width-toggled. It is a *div* element in the demo.

.. adsu:: 2

When users click on the visible label element, the invisible input checkbox
element becomes checked/unchecked by the click. Then the following CSS rules are
used to toggle the width of the element according to whether the checkbox is
checked or not.

.. code-block:: css

  label[for=element-toggle] {
    cursor: pointer;
    color: blue;
  }
  #element-toggle {
    display: none;
  }
  #toggled-element {
    padding: 1rem;
  }
  #element-toggle:checked ~ #toggled-element {
    color: red;
  }

There are four rules in above CSS_ code. The last rule is the most important
part in this trick. It says if the element with *id=element-toggle* (i.e., the
input checkbox) is checked, the color of adjacent element *id=toggled-element*
(*div* element in this case) becomes *red*. If the input checkbox is not
checked, this rule is not applied to the *div* element, which makes the *div*
element original color. The last CSS rule toggles the color of *div* according
to whether the checkbox is checked or not.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 63.0.3239.84 (Official Build) Built on Ubuntu , running on Ubuntu 17.10 (64-bit)``

----

References:

.. [1] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}../../02/27/css-only-toggle-dom-element%en.rst>`_
.. [2] `Pure CSS Toggle Centered Element Width <{filename}../../../2017/03/29/css-only-toggle-centered-element-width%en.rst>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
