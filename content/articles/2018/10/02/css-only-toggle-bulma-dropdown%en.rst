Pure CSS Bulma Dropdown Toggle
##############################

:date: 2018-10-02T19:10+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Bulma, dropdown menu
:category: CSS
:summary: CSS only toggle *Bulma* collapse. No *JavaScript* required.
:og_image: https://cdn.scotch.io/1/r2pUTqX2QWuqxEvuibtp_SpiVz74.png
:adsu: yes


CSS only toggle Bulma_ dropdown_. No *JavaScript* required. If you prefer to use
JavaScript to toggle dropdown, see [2]_.
Click the following dropdown to see the demo of pure css toggle.

|

.. raw:: html

  <style>
  #dropdown-toggle {
    display: none;
  }

  #dropdown-toggle:checked ~ #dropdown-menu {
    display: block;
  }
  </style>

.. raw:: html

  <!-- you can put the two following lines in your html head -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.1/css/bulma.css" integrity="sha256-zKA1Bf41O96+gJSlkn/Bh2HATW/OhwkApPlYTp3B5O8=" crossorigin="anonymous" />
  <script defer src="https://use.fontawesome.com/releases/v5.1.0/js/all.js"></script>

  <div class="dropdown">
    <input id="dropdown-toggle" type="checkbox" />

    <div class="dropdown-trigger">
      <button class="button" aria-haspopup="true" aria-controls="dropdown-menu">
        <label for="dropdown-toggle">Dropdown button</label>
        <span class="icon is-small">
          <i class="fas fa-angle-down" aria-hidden="true"></i>
        </span>
      </button>
    </div>
    <div class="dropdown-menu" id="dropdown-menu" role="menu">
      <div class="dropdown-content">
        <a href="#" class="dropdown-item">
          Dropdown item
        </a>
        <a class="dropdown-item">
          Other dropdown item
        </a>
        <a href="#" class="dropdown-item is-active">
          Active dropdown item
        </a>
        <a href="#" class="dropdown-item">
          Other dropdown item
        </a>
        <hr class="dropdown-divider">
        <a href="#" class="dropdown-item">
          With a divider
        </a>
      </div>
    </div>
  </div>


|

To toggle Bulma dropdown with JavaScript, add the *is-active* class to the class
attribute of the dropdown to activate the modal. The *is-active* class make the
dropdown-menu visible via *display: block*. We will use this hint in our CSS
toggle.

|

We will apply the technique of *Pure CSS Toggle HTML Element* [3]_ here. We
need:

1 - Visible HTML label_ elements, which is the dropdown-trigger.

2 - Invisible HTML input_ checkbox elements, referenced by the *label* elements
via for_ attribute.

3 - dropdown-menu, which is to be toggled and invisible in the beginning.

When users click on the visible label elements, the corresponding input checkbox
element is checked. We use CSS3 **:checked** selector to select corresponding
dropdown menu and make it visible.

The complete source code is as follows:

**HTML**:

.. code-block:: html

  <!-- you can put the two following lines in your html head -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.1/css/bulma.css" integrity="sha256-zKA1Bf41O96+gJSlkn/Bh2HATW/OhwkApPlYTp3B5O8=" crossorigin="anonymous" />
  <script defer src="https://use.fontawesome.com/releases/v5.1.0/js/all.js"></script>

  <div class="dropdown">
    <input id="dropdown-toggle" type="checkbox" />

    <div class="dropdown-trigger">
      <button class="button" aria-haspopup="true" aria-controls="dropdown-menu">
        <label for="dropdown-toggle">Dropdown button</label>
        <span class="icon is-small">
          <i class="fas fa-angle-down" aria-hidden="true"></i>
        </span>
      </button>
    </div>
    <div class="dropdown-menu" id="dropdown-menu" role="menu">
      <div class="dropdown-content">
        <a href="#" class="dropdown-item">
          Dropdown item
        </a>
        <a class="dropdown-item">
          Other dropdown item
        </a>
        <a href="#" class="dropdown-item is-active">
          Active dropdown item
        </a>
        <a href="#" class="dropdown-item">
          Other dropdown item
        </a>
        <hr class="dropdown-divider">
        <a href="#" class="dropdown-item">
          With a divider
        </a>
      </div>
    </div>
  </div>

.. adsu:: 2

**CSS**:

.. code-block:: css

  #dropdown-toggle {
    display: none;
  }

  #dropdown-toggle:checked ~ #dropdown-menu {
    display: block;
  }

- The first rule hides the input checkbox.
- The magic is in second rule. We use **:checked** and general sibling selector
  (~) to make the dropdown menu visible.

----

Tested on:

- ``Bulma 0.7.1``
- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``

----

.. adsu:: 3

References:

.. [1] `Dropdown | Bulma: a modern CSS framework based on Flexbox <https://bulma.io/documentation/components/dropdown/>`_
.. [2] `JavaScript for Bulma Dropdown <{filename}/articles/2018/01/17/bulma-dropdown-with-javascript%en.rst>`_
.. [3] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}/articles/2017/02/27/css-only-toggle-dom-element%en.rst>`_

.. _dropdown: https://bulma.io/documentation/components/dropdown/
.. _Bulma: https://bulma.io/
.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
