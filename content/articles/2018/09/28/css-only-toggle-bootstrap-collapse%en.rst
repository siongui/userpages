Pure CSS Bootstrap Collapse
###########################

:date: 2018-09-28T22:07+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Bootstrap,
       Accordion (Collapsible Content)
:category: CSS
:summary: CSS only toggle *Bootstrap* collapse. No *JavaScript* required.
:og_image: https://img.webnots.com/2017/05/Bootstrap-4-Collapse.png
:adsu: yes

CSS only toggle Bootstrap_ collapse_. No *JavaScript* required.
Click the following button to see demo.

.. raw:: html

  <style>
  #element-toggle {
    display: none;
  }
  #element-toggle:checked ~ #collapseExample {
    display: block;
  }
  </style>

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <label class="btn btn-primary" type="button" for="element-toggle">
    Collapse Trigger
  </label>
  <input id="element-toggle" type="checkbox" />

  <div class="collapse" id="collapseExample">
    <div class="card card-body">
      Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident.
    </div>
  </div>

  </div>


The above demo apply the technique of *Bootstrap* modal [2]_ to basic
*Bootstrap* collapse.

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <div style="margin: 2rem;">

  <label class="btn btn-primary" type="button" for="element-toggle">
    Collapse Trigger
  </label>
  <input id="element-toggle" type="checkbox" />

  <div class="collapse" id="collapseExample">
    <div class="card card-body">
      Anim pariatur cliche reprehenderit, enim eiusmod high life accusamus terry richardson ad squid. Nihil anim keffiyeh helvetica, craft beer labore wes anderson cred nesciunt sapiente ea proident.
    </div>
  </div>

The collapse trigger consists of visible label_ and invisible `input checkbox`_
elements. We apply ``.btn.btn-primary`` class to the label_ element to make it
look like a button. If users click on the label_ element, the visibility of the
collapse is toggled.

Note that in our pure CSS toggle, CSS animation effect during (collapsing)
transitions is not implemented.

.. adsu:: 2

**CSS**:

.. code-block:: css

  #element-toggle {
    display: none;
  }
  #element-toggle:checked ~ #collapseExample {
    display: block;
  }

Only two rules in CSS code:

- First rule make `input checkbox`_ element invisible.
- Second rule toggles the visibility of the collapse when users click on the
  label elements.

----

Tested on:

- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``
- ``Bootstrap 4.1.3``

----

.. adsu:: 3

References:

.. [1] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}/articles/2017/02/27/css-only-toggle-dom-element%en.rst>`_
.. [2] `Pure CSS Bootstrap Modal <{filename}/articles/2018/09/25/css-only-toggle-bootstrap-modal%en.rst>`_

.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input checkbox: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _Bootstrap: https://getbootstrap.com/
.. _collapse: https://getbootstrap.com/docs/4.1/components/collapse/
