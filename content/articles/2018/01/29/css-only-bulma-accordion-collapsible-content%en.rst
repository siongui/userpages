Pure CSS Bulma Accordion (Collapsible Content)
##############################################

:date: 2018-01-29T03:49+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Bulma,
       Accordion (Collapsible Content)
:category: CSS
:summary: CSS only Bulma accordion (collapsible content) implementation.
:og_image: https://bulma.io/images/extensions/bulma-accordion.png
:adsu: yes

CSS only Bulma_ accordion (collapsible content) implementation.
The accordion here is similar to the accordion example in `Bootstrap Collapse`_.
Click the title of following panels to see the demo.

.. raw:: html

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">
  <style>
  input[name="accordion-select"], .panel-block {
    display: none;
  }
  label.panel-heading {
    cursor: pointer;
    display: block;
  }

  /* the magic */
  #panel-1:checked ~ .panel > .body-1,
  #panel-2:checked ~ .panel > .body-2,
  #panel-3:checked ~ .panel > .body-3 {
    display: flex;
  }
  </style>

  <div style="margin: 2rem;">

  <input type="radio" id="panel-1" name="accordion-select">
  <input type="radio" id="panel-2" name="accordion-select">
  <input type="radio" id="panel-3" name="accordion-select">
  <div class="panel">
    <label class="panel-heading" for="panel-1">Title #1</label>
    <div class="panel-block body-1">
      Content Body #1
    </div>
  </div>
  <div class="panel">
    <label class="panel-heading" for="panel-2">Title #2</label>
    <div class="panel-block body-2">
      Content Body #2
    </div>
  </div>
  <div class="panel">
    <label class="panel-heading" for="panel-3">Title #3</label>
    <div class="panel-block body-3">
      Content Body #3
    </div>
  </div>

  </div>


Actually there is `accordion in Bulma extensions`_, but we are not going to use
it. I like to use only standard components provided by Bulma. Here we use
`Bulma panel component`_ to implement the accordion.

The basic technique is the same as *Pure CSS Tab Panel* [3]_. We need:

1 - Visible HTML *label* elements, which is the title of panels.

2 - Invisible HTML *input* *radio* elements, referenced by the *label* elements.
The value of *name* attribute of these elements must be the same.

3 - Panel contents (body), invisible in the beginning.

When users click on the visible label elements, the corresponding input radio
element is checked. We use CSS3 **:checked** selector to select corresponding
panel content, and make the selected panel body content visible.

The complete source code is as follows:

**HTML**:

.. code-block:: html

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">

  <input type="radio" id="panel-1" name="accordion-select">
  <input type="radio" id="panel-2" name="accordion-select">
  <input type="radio" id="panel-3" name="accordion-select">
  <div class="panel">
    <label class="panel-heading" for="panel-1">Title #1</label>
    <div class="panel-block body-1">
      Content Body #1
    </div>
  </div>
  <div class="panel">
    <label class="panel-heading" for="panel-2">Title #2</label>
    <div class="panel-block body-2">
      Content Body #2
    </div>
  </div>
  <div class="panel">
    <label class="panel-heading" for="panel-3">Title #3</label>
    <div class="panel-block body-3">
      Content Body #3
    </div>
  </div>

.. adsu:: 2

**CSS**: Note that you must put Bulma CSS code before the following code.

.. code-block:: css

  input[name="accordion-select"], .panel-block {
    display: none;
  }
  label.panel-heading {
    cursor: pointer;
    display: block;
  }

  /* the magic */
  #panel-1:checked ~ .panel > .body-1,
  #panel-2:checked ~ .panel > .body-2,
  #panel-3:checked ~ .panel > .body-3 {
    display: flex;
  }

The magic is in last rule. We use **:checked** and general sibling selector (~)
to make the user-selected panel body content visible.

.. adsu:: 3

----

Tested on:

- ``Chromium 63.0.3239.132 on Ubuntu 17.10 (64-bit)``
- ``Bulma 0.6.2``

----

References:

.. [1] `[Vue.js] Bulma Accordion (Collapsible Content) <{filename}/articles/2018/01/28/vuejs-bulma-accordion-collapsible-content%en.rst>`_
.. [2] `Pure CSS Accordion (Collapsible Content) <{filename}/articles/2017/05/23/css-only-accordion-collapsible-content%en.rst>`_
.. [3] `Pure CSS Tab Panel <{filename}/articles/2017/05/21/css-only-tab-panel%en.rst>`_

.. _Bulma: https://bulma.io/
.. _accordion in Bulma extensions: https://wikiki.github.io/components/accordion/
.. _Bulma panel component: https://bulma.io/documentation/components/panel/
.. _Bootstrap Collapse: https://getbootstrap.com/docs/3.3/javascript/#collapse-example-accordion
