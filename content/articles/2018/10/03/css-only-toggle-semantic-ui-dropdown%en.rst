Pure CSS Semantic UI Dropdown
#############################

:date: 2018-10-03T21:23+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Semantic UI,
       dropdown menu
:category: CSS
:summary: CSS only toggle *Semantic UI* dropdown. No *JavaScript* required.
:og_image: https://ninghao.net/files/screenshot/semantic-ui-m-05-05-dropdown-icon.png
:adsu: yes


CSS only toggle `Semantic UI`_ dropdown_. No *JavaScript* required.
The behavior of dropdown here is almost the same as `simple dropdown`_ in
official doc, which also requires no JavaScript. 
The difference is that `simple dropdown`_ opens when mouse hovers over, and
dropdown here opens when mouse clicks. The dropdown here is modified from
`compact dropdown`_ in official doc.

Click the following *Dropdown* text to see the demo of pure css toggle.

.. raw:: html

  <style>
  #dropdown-toggle {
    display: none;
  }

  #dropdown-toggle:checked ~ div.menu {
    display: block;
  }
  </style>

.. raw:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.0/dist/semantic.min.css">

  <div class="ui compact selection dropdown">
    <i class="dropdown icon"></i>
    <label class="text" for="dropdown-toggle">Dropdown</label>
    <input id="dropdown-toggle" type="checkbox" />
    <div class="menu">
      <div class="item">Choice 1</div>
      <div class="item">Choice 2</div>
      <div class="item">Choice 3</div>
    </div>
  </div>

|

We will apply the technique of *Pure CSS Toggle HTML Element* [1]_ here. We
need:

1 - Visible HTML label_ elements, which is the dropdown-trigger, and is
    the text *Dropdown* in our example.

2 - Invisible HTML input_ checkbox elements, referenced by the *label* elements
via for_ attribute.

3 - dropdown-menu, which is to be toggled and invisible in the beginning.

When users click on the visible label elements, the corresponding input checkbox
element is checked. We use CSS3 **:checked** selector to select corresponding
dropdown menu and make it visible.

The complete source code is as follows:

**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/semantic-ui@2.4.0/dist/semantic.min.css">

  <div class="ui compact selection dropdown">
    <i class="dropdown icon"></i>
    <label class="text" for="dropdown-toggle">Dropdown</label>
    <input id="dropdown-toggle" type="checkbox" />
    <div class="menu">
      <div class="item">Choice 1</div>
      <div class="item">Choice 2</div>
      <div class="item">Choice 3</div>
    </div>
  </div>

.. adsu:: 2

**CSS**:

.. code-block:: css

  #dropdown-toggle {
    display: none;
  }

  #dropdown-toggle:checked ~ div.menu {
    display: block;
  }

- The first rule hides the input checkbox.
- The magic is in second rule. We use **:checked** and general sibling selector
  (~) to make the dropdown menu visible.

----

Tested on:

- ``Semantic UI 2.4.0``
- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``

----

.. adsu:: 3

References:

.. [1] `Pure CSS Toggle (Show/Hide) HTML Element <{filename}/articles/2017/02/27/css-only-toggle-dom-element%en.rst>`_
.. [2] `Pure CSS Bulma Dropdown Toggle <{filename}/articles/2018/10/02/css-only-toggle-bulma-dropdown%en.rst>`_

.. _dropdown: https://semantic-ui.com/modules/dropdown.html
.. _Semantic UI: https://semantic-ui.com/
.. _simple dropdown: https://semantic-ui.com/modules/dropdown.html#simple
.. _compact dropdown: https://semantic-ui.com/modules/dropdown.html#compact
.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _input: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
