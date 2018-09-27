Bulma Modal with Pure CSS Toggle
################################

:date: 2018-01-27T03:10+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Modal (Popup), Bulma
:category: CSS
:summary: CSS only toggle *Bulma* modal. No *JavaScript* required.
:og_image: https://www.designil.com/wp-content/uploads/2016/10/modal-box-bulma.png
:adsu: yes

CSS only toggle Bulma_ modal_. Click the following button to see demo.

.. raw:: html

  <style>
  #element-toggle {
    display: none;
  }
  #element-toggle:checked ~ #myModal {
    display: flex;
  }
  </style>

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">
  <div style="margin: 2rem;">

  <!-- Modal trigger -->
  <label class="button" for="element-toggle">Launch Modal</label>
  <input id="element-toggle" type="checkbox" />

  <!-- Modal -->
  <div class="modal" id="myModal">
    <div class="modal-background"></div>
    <div class="modal-content">
      <p class="image is-4by3">
        <img src="https://bulma.io/images/placeholders/1280x960.png" alt="image modal">
      </p>
    </div>
    <label class="modal-close is-large" for="element-toggle"></label>
  </div>

  </div>


The above demo apply the technique of my previous post of pure CSS modal [1]_
to Bulma modal.

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <!-- Modal trigger -->
  <label class="button" for="element-toggle">Launch Modal</label>
  <input id="element-toggle" type="checkbox" />

  <!-- Modal -->
  <div class="modal" id="myModal">
    <div class="modal-background"></div>
    <div class="modal-content">
      <p class="image is-4by3">
        <img src="https://bulma.io/images/placeholders/1280x960.png" alt="image modal">
      </p>
    </div>
    <label class="modal-close is-large" for="element-toggle"></label>
  </div>

The modal trigger consists of visible label_ and invisible `input checkbox`_
elements. We apply ``.button`` class to the label_ element to make it look like
a button. The **Close** button in the modal is also a label_ element.
If users click on the label_ element, the visibility of the modal is toggled.

.. adsu:: 2

**CSS**:

.. code-block:: css

  #element-toggle {
    display: none;
  }
  #element-toggle:checked ~ #myModal {
    display: flex;
  }

Only two rules in CSS code:

- First rule make `input checkbox`_ element invisible.
- Second rule toggles the visibility of the modal when users click on the
  label elements.

----

**You might be interested in ...**

- `[Vue.js] Bulma Modal <{filename}/articles/2018/09/27/vuejs-bulma-modal%en.rst>`_
- `JavaScript for Bulma Modal <{filename}/articles/2018/01/17/bulma-modal-with-javascript%en.rst>`_
- `Bulma Modal with Go Toggle <{filename}/articles/2017/12/04/bulma-modal-with-go-toggle%en.rst>`_

----

Tested on:

- ``Chromium 63.0.3239.132 on Ubuntu 17.10 (64-bit)``
- ``Bulma 0.6.2``

----

.. adsu:: 3

References:

.. [1] `Pure CSS Modal (Popup) <{filename}/articles/2017/05/12/css-only-modal-popup%en.rst>`_

.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input checkbox: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _Bulma: https://bulma.io/
.. _modal: https://bulma.io/documentation/components/modal/
