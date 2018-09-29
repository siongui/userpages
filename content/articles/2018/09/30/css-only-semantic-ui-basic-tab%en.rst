Pure CSS Semantic UI Basic Tab
##############################

:date: 2018-09-30T02:54+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Semantic UI, Tab Panel
:category: CSS
:summary: CSS only *Semantic UI* basic tab. No *JavaScript* required.
:og_image: https://ninghao.net/files/screenshot/semantic-ui-m-08-01-tab.png
:adsu: yes

CSS only toggle `Semantic UI`_ basic tab_. No *JavaScript* required.
Click the following tabs to see the demo.

.. raw:: html

  <style>
  input[name="tab-select"] {
    display: none;
  }

  /* the magic */
  #tab1:checked ~ .ui.tab.segment.tab1,
  #tab2:checked ~ .ui.tab.segment.tab2,
  #tab3:checked ~ .ui.tab.segment.tab3 {
    display: block;
  }
  </style>

  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.0/dist/semantic.min.css">
  <div style="margin: 2rem;">

  <input type="radio" id="tab1" name="tab-select">
  <input type="radio" id="tab2" name="tab-select">
  <input type="radio" id="tab3" name="tab-select">
  <div class="ui top attached tabular menu">
    <a class="item"><label for="tab1">First</label></a>
    <a class="item"><label for="tab2">Second</label></a>
    <a class="item"><label for="tab3">Third</label></a>
  </div>
  <div class="ui bottom attached tab segment tab1">
    First
  </div>
  <div class="ui bottom attached tab segment tab2">
    Second
  </div>
  <div class="ui bottom attached tab segment tab3">
    Third
  </div>

  </div>


The basic technique is the same as *pure CSS toggle HTML element* [1]_. We need:

1 - Visible HTML *label* elements, which is the navigation tabs.

2 - Invisible HTML *input* *radio* elements, referenced by the *label* elements.
The value of *name* attribute of these elements must be the same.

3 - Tab contents, invisible in the beginning.

When users click on the visible label elements, the corresponding input radio
element is checked. We use CSS3 **:checked** selector to select corresponding
tab content, and make the selected tab content visible.

The complete source code is as follows:


**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.0/dist/semantic.min.css">

  <input type="radio" id="tab1" name="tab-select">
  <input type="radio" id="tab2" name="tab-select">
  <input type="radio" id="tab3" name="tab-select">
  <div class="ui top attached tabular menu">
    <a class="item"><label for="tab1">First</label></a>
    <a class="item"><label for="tab2">Second</label></a>
    <a class="item"><label for="tab3">Third</label></a>
  </div>
  <div class="ui bottom attached tab segment tab1">
    First
  </div>
  <div class="ui bottom attached tab segment tab2">
    Second
  </div>
  <div class="ui bottom attached tab segment tab3">
    Third
  </div>

.. adsu:: 2

**CSS**:

.. code-block:: css

  input[name="tab-select"] {
    display: none;
  }

  /* the magic */
  #tab1:checked ~ .ui.tab.segment.tab1,
  #tab2:checked ~ .ui.tab.segment.tab2,
  #tab3:checked ~ .ui.tab.segment.tab3 {
    display: block;
  }

- The first rule hides the input radio box.
- The magic is in second rule. We use **:checked** and general sibling selector
  (~) to make the user-selected tab content visible.

----

Tested on:

- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``
- ``Semantic UI 2.4.0``

----

.. adsu:: 3

References:

.. [1] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}/articles/2017/02/27/css-only-toggle-dom-element%en.rst>`_
.. [2] `Pure CSS Tab Panel <{filename}/articles/2017/05/21/css-only-tab-panel%en.rst>`_
.. [3] `Pure CSS Bulma Tabs <{filename}/articles/2018/01/30/css-only-bulma-tab-panel%en.rst>`_

.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input checkbox: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _Semantic UI: https://semantic-ui.com/
.. _tab: https://semantic-ui.com/modules/tab.html
