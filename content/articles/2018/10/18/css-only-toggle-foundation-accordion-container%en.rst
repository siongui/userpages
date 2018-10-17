Pure CSS Foundation Accordion Container
#######################################

:date: 2018-10-18T03:20+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS,
       Foundation Front-end Framework, Accordion (Collapsible Content)
:category: CSS
:summary: CSS only toggle *Foundation* accordion (collapsible content).
          No *JavaScript* required.
:og_image: https://foundationstacks.com/depot/images/accordion.svg
:adsu: yes

CSS only toggle Foundation_ accordion_ container. No *JavaScript* required.
Click the following texts of items to see demo.

.. raw:: html

  <style>
  input[name="accordion-select"], .panel-block {
    display: none;
  }

  /* the magic */
  #panel-1:checked ~ .accordion > .accordion-item > .content-1,
  #panel-2:checked ~ .accordion > .accordion-item > .content-2,
  #panel-3:checked ~ .accordion > .accordion-item > .content-3 {
    display: block;
  }
  </style>

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/foundation-sites@6.5.0-rc.4/dist/css/foundation.min.css" integrity="sha256-VEEaOnBKVRoYPn4AID/tY/XKVxKEqXstoo/xZ6nemak= sha384-D46t32f421/hB30qwnim2pIcisNN5GU9+6m2Mfnd3dKpTSFidZLa08/1StEiCFId sha512-WkgzH8VKemDfwrp18r+wgbx+oHXOkfd2kJs7ocAXdGDgonXDXh88E90IRtRZRXtO0IHprxYHYlY14h+wyTsUDA==" crossorigin="anonymous">

  <input type="radio" id="panel-1" name="accordion-select">
  <input type="radio" id="panel-2" name="accordion-select">
  <input type="radio" id="panel-3" name="accordion-select">

  <ul class="accordion" data-accordion>
    <li class="accordion-item" data-accordion-item>
      <a class="accordion-title"><label for="panel-1">Accordion 1</label></a>
      <div class="accordion-content content-1" data-tab-content>
        <p>Panel 1. Lorem ipsum dolor</p>
        <a href="#">Nowhere to Go</a>
      </div>
    </li>
    <li class="accordion-item" data-accordion-item>
      <a class="accordion-title"><label for="panel-2">Accordion 2</label></a>
      <div class="accordion-content content-2" data-tab-content>
        <textarea></textarea>
        <button class="button">I do nothing!</button>
      </div>
    </li>
    <li class="accordion-item" data-accordion-item>
      <a class="accordion-title"><label for="panel-3">Accordion 3</label></a>
      <div class="accordion-content content-3" data-tab-content>
        Type your name!
        <input type="text"></input>
      </div>
    </li>
  </ul>


The basic technique is the same as *Pure CSS Tab Panel* [3]_. We need:

1 - Visible HTML label_ elements, which is the item title of accordion tab.

2 - Invisible HTML *input* *radio* elements, referenced by the *label* elements.
The value of *name* attribute of these elements must be the same.

3 - Content of accordion tab, invisible in the beginning.

When users click on the visible label elements, which are accordion tab title,
the corresponding input radio element is checked. We use CSS3 **:checked**
selector to select corresponding accordion tab content, and make the selected
content visible.

.. adsu:: 2

The complete source code is as follows:

**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/foundation-sites@6.5.0-rc.4/dist/css/foundation.min.css" integrity="sha256-VEEaOnBKVRoYPn4AID/tY/XKVxKEqXstoo/xZ6nemak= sha384-D46t32f421/hB30qwnim2pIcisNN5GU9+6m2Mfnd3dKpTSFidZLa08/1StEiCFId sha512-WkgzH8VKemDfwrp18r+wgbx+oHXOkfd2kJs7ocAXdGDgonXDXh88E90IRtRZRXtO0IHprxYHYlY14h+wyTsUDA==" crossorigin="anonymous">

  <input type="radio" id="panel-1" name="accordion-select">
  <input type="radio" id="panel-2" name="accordion-select">
  <input type="radio" id="panel-3" name="accordion-select">

  <ul class="accordion" data-accordion>
    <li class="accordion-item" data-accordion-item>
      <a class="accordion-title"><label for="panel-1">Accordion 1</label></a>
      <div class="accordion-content content-1" data-tab-content>
        <p>Panel 1. Lorem ipsum dolor</p>
        <a href="#">Nowhere to Go</a>
      </div>
    </li>
    <li class="accordion-item" data-accordion-item>
      <a class="accordion-title"><label for="panel-2">Accordion 2</label></a>
      <div class="accordion-content content-2" data-tab-content>
        <textarea></textarea>
        <button class="button">I do nothing!</button>
      </div>
    </li>
    <li class="accordion-item" data-accordion-item>
      <a class="accordion-title"><label for="panel-3">Accordion 3</label></a>
      <div class="accordion-content content-3" data-tab-content>
        Type your name!
        <input type="text"></input>
      </div>
    </li>
  </ul>

.. adsu:: 3

**CSS**:

.. code-block:: css

  input[name="accordion-select"], .panel-block {
    display: none;
  }

  /* the magic */
  #panel-1:checked ~ .accordion > .accordion-item > .content-1,
  #panel-2:checked ~ .accordion > .accordion-item > .content-2,
  #panel-3:checked ~ .accordion > .accordion-item > .content-3 {
    display: block;
  }

The magic is in last rule. We use **:checked** and general sibling selector (~)
to make the user-selected accordion tab content visible.

----

Tested on:

- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``
- ``Foundation for Sites 6.5.0-rc.4``

----

.. adsu:: 4

References:

.. [1] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}/articles/2017/02/27/css-only-toggle-dom-element%en.rst>`_
.. [2] `Pure CSS Bootstrap Modal <{filename}/articles/2018/09/25/css-only-toggle-bootstrap-modal%en.rst>`_
.. [3] `Pure CSS Tab Panel <{filename}/articles/2017/05/21/css-only-tab-panel%en.rst>`_
.. [4] `Pure CSS Accordion (Collapsible Content) <{filename}/articles/2017/05/23/css-only-accordion-collapsible-content%en.rst>`_
.. [5] `[Vue.js] Bulma Accordion (Collapsible Content) <{filename}/articles/2018/01/28/vuejs-bulma-accordion-collapsible-content%en.rst>`_
.. [6] `Pure CSS Bulma Accordion (Collapsible Content) <{filename}/articles/2018/01/29/css-only-bulma-accordion-collapsible-content%en.rst>`_
.. [7] `Pure CSS Semantic UI Standard Accordion <{filename}/articles/2018/02/07/css-only-semantic-ui-standard-accordion%en.rst>`_
.. [8] `Pure CSS Bootstrap Accordion <{filename}/articles/2018/09/29/css-only-toggle-bootstrap-accordion%en.rst>`_
.. [9] `Evaluating CSS Frameworks — Bootstrap vs Bulma vs Foundation vs Milligram vs Pure vs Semantic vs UIKit <https://codeburst.io/evaluating-css-frameworks-bulma-vs-foundation-vs-milligram-vs-pure-vs-semantic-vs-uikit-503883bd25a3>`_

.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _Foundation: https://github.com/zurb/foundation-sites
.. _accordion: https://foundation.zurb.com/sites/docs/accordion.html
