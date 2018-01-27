Pure CSS Accordion (Collapsible Content)
########################################

:date: 2017-05-23T02:54+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS,
       Accordion (Collapsible Content)
:category: CSS
:summary: CSS only accordion (collapsible content) implementation.
:og_image: http://christineosazuwa.com/wp-content/uploads/2015/05/13-06_accordionmenu1.jpg
:adsu: yes

CSS only accordion (collapsible content) implementation.
The accordion here is similar to the accordion example in `Bootstrap Collapse`_.
Click the title of following panels to see the demo.

.. raw:: html

  <style>
  .panel {
    margin-bottom: 1rem;
    border: 1px solid #ccc;
  }
  .panel-title {
    font-weight: bold;
    background-color: #ccc;
    padding: 0.01em 16px;
  }
  .panel-title > label {
    cursor: pointer;
  }
  .panel-body {
    padding: 0.01em 16px;
    display: none;
  }

  input[name="accordion-select"] {
    display: none;
  }

  /* the magic */
  #panel-1:checked ~ .panel > .body-1,
  #panel-2:checked ~ .panel > .body-2,
  #panel-3:checked ~ .panel > .body-3 {
    display: block;
  }
  </style>

  <input type="radio" id="panel-1" name="accordion-select">
  <input type="radio" id="panel-2" name="accordion-select">
  <input type="radio" id="panel-3" name="accordion-select">
  <div class="panel">
    <div class="panel-title">
      <label for="panel-1">Title #1</label>
    </div>
    <div class="panel-body body-1">
      Content Body #1
    </div>
  </div>
  <div class="panel">
    <div class="panel-title">
      <label for="panel-2">Title #2</label>
    </div>
    <div class="panel-body body-2">
      Content Body #2
    </div>
  </div>
  <div class="panel">
    <div class="panel-title">
      <label for="panel-3">Title #3</label>
    </div>
    <div class="panel-body body-3">
      Content Body #3
    </div>
  </div>

The basic technique is the same as *Pure CSS Tab Panel* [2]_. We need:

- Visible HTML *label* elements, which is the title of panels.
- Invisible HTML *input* *radio* elements, referenced by the *label* elements.
  The value of *name* attribute of these elements must be the same.
- Panel contents (body), invisible in the beginning.

When users click on the visible label elements, the corresponding input radio
element is checked. We use CSS3 **:checked** selector to select corresponding
panel content, and make the selected panel body content visible.

The complete source code is as follows:

**HTML**:

.. code-block:: html

  <input type="radio" id="panel-1" name="accordion-select">
  <input type="radio" id="panel-2" name="accordion-select">
  <input type="radio" id="panel-3" name="accordion-select">
  <div class="panel">
    <div class="panel-title">
      <label for="panel-1">Title #1</label>
    </div>
    <div class="panel-body body-1">
      Content Body #1
    </div>
  </div>
  <div class="panel">
    <div class="panel-title">
      <label for="panel-2">Title #2</label>
    </div>
    <div class="panel-body body-2">
      Content Body #2
    </div>
  </div>
  <div class="panel">
    <div class="panel-title">
      <label for="panel-3">Title #3</label>
    </div>
    <div class="panel-body body-3">
      Content Body #3
    </div>
  </div>

.. adsu:: 2

**CSS**:

.. code-block:: css

  .panel {
    margin-bottom: 1rem;
    border: 1px solid #ccc;
  }
  .panel-title {
    font-weight: bold;
    background-color: #ccc;
    padding: 0.01em 16px;
  }
  .panel-title > label {
    cursor: pointer;
  }
  .panel-body {
    padding: 0.01em 16px;
    display: none;
  }

  input[name="accordion-select"] {
    display: none;
  }

  /* the magic */
  #panel-1:checked ~ .panel > .body-1,
  #panel-2:checked ~ .panel > .body-2,
  #panel-3:checked ~ .panel > .body-3 {
    display: block;
  }

The magic is in last rule. We use **:checked** and general sibling selector (~)
to make the user-selected panel body content visible.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 58.0.3029.110 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] `[Vue.js] Accordion (Collapsible Content) <{filename}../22/vuejs-accordion-collapsible-content%en.rst>`_
.. [2] `Pure CSS Tab Panel <{filename}../21/css-only-tab-panel%en.rst>`_

.. _Bootstrap Collapse: https://getbootstrap.com/docs/3.3/javascript/#collapse-example-accordion
