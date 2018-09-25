Pure CSS Bootstrap Modal
########################

:date: 2018-09-25T21:44+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Modal (Popup), Bootstrap
:category: CSS
:summary: CSS only toggle *Bootstrap* modal. No *JavaScript* required.
:og_image: https://scripts-cdn.softpedia.com/screenshots/Bootstrap-Modal_1.png
:adsu: yes

CSS only toggle Bootstrap_ modal_. No *JavaScript* required.
Click the following button to see demo.

.. raw:: html

  <style>
  #element-toggle {
    display: none;
  }
  #element-toggle:checked ~ #exampleModal {
    display: block;
  }
  </style>

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <div style="margin: 2rem;">

  <!-- Pseudo-Button trigger modal -->
  <label type="button" class="btn btn-primary" for="element-toggle">
    Launch demo modal
  </label>
  <input id="element-toggle" type="checkbox" />

  <!-- Modal -->
  <div class="modal fade show" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="exampleModalLabel">Modal title</h5>
          <label for="element-toggle" type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </label>
        </div>
        <div class="modal-body">
          ...
        </div>
        <div class="modal-footer">
          <label for="element-toggle" type="button" class="btn btn-secondary" data-dismiss="modal">Close</label>
        </div>
      </div>
    </div>
  </div>

  </div>


The above demo apply the technique of *Semantic UI* modal [2]_ to *Bootstrap*
modal.

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <!-- Pseudo-Button trigger modal -->
  <label type="button" class="btn btn-primary" for="element-toggle">
    Launch demo modal
  </label>
  <input id="element-toggle" type="checkbox" />

  <!-- Modal -->
  <div class="modal fade show" id="exampleModal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="exampleModalLabel">Modal title</h5>
          <label for="element-toggle" type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </label>
        </div>
        <div class="modal-body">
          ...
        </div>
        <div class="modal-footer">
          <label for="element-toggle" type="button" class="btn btn-secondary" data-dismiss="modal">Close</label>
        </div>
      </div>
    </div>
  </div>

The modal trigger consists of visible label_ and invisible `input checkbox`_
elements. We apply ``.btn`` class to the label_ element to make it look like a
button. The **Close** and **×** button in the modal is also a label_ element. If
users click on the label_ element, the visibility of the modal is toggled.

.. adsu:: 2

**CSS**:

.. code-block:: css

  #element-toggle {
    display: none;
  }
  #element-toggle:checked ~ #exampleModal {
    display: block;
  }

Only two rules in CSS code:

- First rule make `input checkbox`_ element invisible.
- Second rule toggles the visibility of the modal when users click on the
  label elements.

----

Tested on:

- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``
- ``Bootstrap 4.1.3``

----

.. adsu:: 3

References:

.. [1] `Bulma Modal with Pure CSS Toggle <{filename}/articles/2018/01/27/css-only-toggle-bulma-modal%en.rst>`_
.. [2] `Pure CSS Semantic UI Modal <{filename}/articles/2018/01/31/css-only-toggle-semantic-ui-modal%en.rst>`_

.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input checkbox: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _Bootstrap: https://getbootstrap.com/
.. _modal: https://getbootstrap.com/docs/4.1/components/modal/
