Pure CSS Semantic UI Modal
##########################

:date: 2018-01-31T01:29+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Modal (Popup),
       Semantic UI
:category: CSS
:summary: CSS only toggle *Semantic UI* modal. No *JavaScript* required.
:og_image: https://i.stack.imgur.com/2wmjd.png
:adsu: yes

CSS only toggle `Semantic UI`_ modal_. No *JavaScript* required.
Click the following button to see demo.

.. raw:: html

  <style>
  #element-toggle {
    display: none;
  }
  #element-toggle:checked ~ #myModal {
    display: block;
  }
  </style>

  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/semantic-ui@2.2.14/dist/semantic.min.css">
  <div style="margin: 2rem;">

  <!-- Modal trigger -->
  <label class="ui button" for="element-toggle">Launch Modal</label>
  <input id="element-toggle" type="checkbox" />

  <!-- Modal -->
  <div class="ui modal" id="myModal">
    <div class="header">Header</div>
    <div class="content">
      <p>Line 1 of Modal Content</p>
      <p>Line 2 of Modal Content</p>
      <p>Line 3 of Modal Content</p>
    </div>
    <div class="actions">
      <label class="ui ok button" for="element-toggle">OK</label>
    </div>
  </div>

  </div>


The above demo apply the technique of Bulma modal [1]_ to Semantic UI modal.

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <!-- Modal trigger -->
  <label class="ui button" for="element-toggle">Launch Modal</label>
  <input id="element-toggle" type="checkbox" />

  <!-- Modal -->
  <div class="ui modal" id="myModal">
    <div class="header">Header</div>
    <div class="content">
      <p>Line 1 of Modal Content</p>
      <p>Line 2 of Modal Content</p>
      <p>Line 3 of Modal Content</p>
    </div>
    <div class="actions">
      <label class="ui ok button" for="element-toggle">OK</label>
    </div>
  </div>

The modal trigger consists of visible label_ and invisible `input checkbox`_
elements. We apply ``.button`` class to the label_ element to make it look like
a button. The **OK** button in the modal is also a label_ element.
If users click on the label_ element, the visibility of the modal is toggled.

.. adsu:: 2

**CSS**:

.. code-block:: css

  #element-toggle {
    display: none;
  }
  #element-toggle:checked ~ #myModal {
    display: block;
  }

Only two rules in CSS code:

- First rule make `input checkbox`_ element invisible.
- Second rule toggles the visibility of the modal when users click on the
  label elements.

----

Tested on:

- ``Chromium 63.0.3239.132 on Ubuntu 17.10 (64-bit)``
- ``Semantic UI 2.2.14``

----

.. adsu:: 3

References:

.. [1] `Bulma Modal with Pure CSS Toggle <{filename}/articles/2018/01/27/css-only-toggle-bulma-modal%en.rst>`_

.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input checkbox: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _Semantic UI: https://semantic-ui.com/
.. _modal: https://semantic-ui.com/modules/modal.html
