Pure CSS Toggle Centered Element Width
######################################

:date: 2017-03-29T22:47+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS
:category: CSS
:summary: CSS_ only toggle the width of centered HTML_ `DOM element`_. No
          JavaScript_ required.
:og_image: http://www.unm.edu/~tbeach/IT145/week08/images/boxmodel.gif
:adsu: yes


.. role:: raw-html(raw)
   :format: html

CSS_ only toggle the width of centered HTML_ `DOM element`_. No JavaScript_
required. See demo first.
Please click the following text to toggle the width of centered div:

.. raw:: html

  <style>
  label[for=element-toggle] {
    cursor: pointer;
    color: red;
  }
  #element-toggle {
    display: none;
  }
  #toggled-element {
    margin: auto;
    width: 100%;
    background-color: aqua;
    padding: 1rem;
  }
  #element-toggle:not(:checked) ~ #toggled-element {
    width: 60%;
  }
  </style>

  <div>

  <label for="element-toggle">click me to toggle width</label>
  <input id="element-toggle" type="checkbox" />
  <br><br>
  <div id="toggled-element">centered div element</div>

  </div>


This demo uses the technique of CSS toggling in [3]_.
There are three main HTML elements in the demo:

.. code-block:: html

  <label for="element-toggle">click me to toggle width</label>
  <input id="element-toggle" type="checkbox" />

  <div id="toggled-element">centered div element</div>

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
    color: red;
  }
  #element-toggle {
    display: none;
  }
  #toggled-element {
    margin: auto;
    width: 100%;
    background-color: aqua;
    padding: 1rem;
  }
  #element-toggle:not(:checked) ~ #toggled-element {
    width: 60%;
  }

There are three rules in above CSS_ code. The last rule is the most important
part in this trick. It says if the element with *id=element-toggle* (i.e., the
input checkbox) is not checked, the width of adjacent element with
*id=toggled-element* (*div* element in this case) becomes *60%*. If the input
checkbox is checked, this rule is not applied to the *div* element, which makes
the *div* element ``width: 100%;``. The last CSS rule toggles the width of *div*
according to whether the checkbox is checked or not.

To center a block element horizontally, use ``margin: auto;`` [2]_.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 56.0.2924.76 Built on Ubuntu , running on Ubuntu 16.10 (64-bit)``

----

References:

.. [1] | `center element css - Google search <https://www.google.com/search?q=center+element+css>`_
       | `center element css - DuckDuckGo search <https://duckduckgo.com/?q=center+element+css>`_
       | `center element css - Ecosia search <https://www.ecosia.org/search?q=center+element+css>`_
       | `center element css - Qwant search <https://www.qwant.com/?q=center+element+css>`_
       | `center element css - Bing search <https://www.bing.com/search?q=center+element+css>`_
       | `center element css - Yahoo search <https://search.yahoo.com/search?p=center+element+css>`_
       | `center element css - Baidu search <https://www.baidu.com/s?wd=center+element+css>`_
       | `center element css - Yandex search <https://www.yandex.com/search/?text=center+element+css>`_

.. [2] `CSS Layout - Horizontal & Vertical Align <https://www.w3schools.com/css/css_align.asp>`_

.. [3] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}../../02/27/css-only-toggle-dom-element%en.rst>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
