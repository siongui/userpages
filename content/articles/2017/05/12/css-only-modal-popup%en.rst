Pure CSS Modal (Popup)
######################

:date: 2017-05-12T03:42+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Modal (Popup)
:category: CSS
:summary: CSS only modal (popup). No JavaScript_ required.
          Modal is dialog box/popup window that is displayed on top of the
          current page.
:og_image: https://www.sololearn.com/Icons/Courses/1023.png
:adsu: yes

CSS only modal (popup) implementation.
`According to w3schools`_, modal is dialog box/popup window that is displayed on
top of the current page.
The modal here is similar to `bootstrap modal`_.

.. raw:: html

  <style>
  #element-toggle {
    display: none;
  }
  #element-toggle:not(:checked) ~ #myModal {
    display: none;
  }

  .modal {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
    background: azure;
    padding: 1rem;
    border-radius: 1rem;
    border: 1px solid gray;
  }

  .button {
    background-color: beige;
    border-radius: 4px;
    padding: 4px;
    border: 1px solid gray;
  }
  </style>

  <div style="padding: 10px;">

  <!-- Modal trigger -->
  <label class="button" for="element-toggle">Launch Modal</label>
  <input id="element-toggle" type="checkbox" />

  <!-- Modal -->
  <div class="modal" id="myModal">
    <h4>Modal Title</h4>
    <p>
       Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas turpis
       felis, tempus sit amet sollicitudin quis, aliquet ut felis.
       Nam a malesuada sem.
    </p>
    <label class="button" for="element-toggle">Close</label>
  </div>

  </div>

I apply the pure CSS toggle technique [2]_ to the Vue.js modal [1]_, and
successfully remove the JavaScript/Vue.js code in the modal/popup
implementation. Please read the references first if you are not familiar with
the CSS toggle technique or JavaScript modal.

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <!-- Modal trigger -->
  <label class="button" for="element-toggle">Launch Modal</label>
  <input id="element-toggle" type="checkbox" />

  <!-- Modal -->
  <div class="modal" id="myModal">
    <h4>Modal Title</h4>
    <p>
       Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas turpis
       felis, tempus sit amet sollicitudin quis, aliquet ut felis.
       Nam a malesuada sem.
    </p>
    <label class="button" for="element-toggle">Close</label>
  </div>

The modal trigger consists of visible *label* and invisible *input checkbox*
elements. We apply ``.button`` class to the *label* element to make it look like
a button. The **Close** button in the modal is also a *label* element with
``.button`` class. If users click on the *label* element, the visibility of the
modal is toggled.

.. adsu:: 2

**CSS**:

.. code-block:: css

  #element-toggle {
    display: none;
  }
  #element-toggle:not(:checked) ~ #myModal {
    display: none;
  }

  .modal {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
    background: azure;
    padding: 1rem;
    border-radius: 1rem;
    border: 1px solid gray;
  }

  .button {
    background-color: beige;
    border-radius: 4px;
    padding: 4px;
    border: 1px solid gray;
  }

We have four rules in above CSS code.

- First two rules toggles the visibility of the modal when users click on the
  label elements.
- ``.modal`` class centers the modal horizontally and vertically.
- ``.button`` class make the label elements look like buttons.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 58.0.3029.96 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] `[Vue.js] Modal (Popup) <{filename}../11/vuejs-modal-popup%en.rst>`_
.. [2] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}../../02/27/css-only-toggle-dom-element%en.rst>`_

.. _HTML: https://www.google.com/search?q=HTML
.. _CSS: https://www.google.com/search?q=CSS
.. _JavaScript: https://www.google.com/search?q=JavaScript
.. _DOM element: https://www.google.com/search?q=DOM+element
.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _According to w3schools: https://www.w3schools.com/bootstrap/bootstrap_modal.asp
.. _bootstrap modal: http://getbootstrap.com/javascript/#modals
