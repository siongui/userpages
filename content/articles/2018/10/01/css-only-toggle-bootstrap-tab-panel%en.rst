Pure CSS Bootstrap Tab Panel
############################

:date: 2018-10-01T22:59+08:00
:tags: Pure CSS Toggle, toggle, toggleable, html, CSS, Bootstrap, Tab Panel
:category: CSS
:summary: CSS only toggle *Bootstrap* tab panel.
          No *JavaScript* required.
:og_image: http://codepen.io/wizly/pen/BlKxo/image/large.png
:adsu: yes

CSS only toggle Bootstrap_ `tab panel`_. No *JavaScript* required.
Click the following tabs to see demo.

.. raw:: html

  <style>
  input[name="nav-tab"] {
    display: none;
  }

  /* the magic: show tab panes */
  #home:checked ~ #nav-tabContent > #nav-home,
  #profile:checked ~ #nav-tabContent > #nav-profile,
  #contact:checked ~ #nav-tabContent > #nav-contact {
    display: block;
    opacity: 1;
  }
  /* the magic: make tab active */
  #home:checked ~ nav > #nav-tab > #nav-home-tab,
  #profile:checked ~ nav > #nav-tab > #nav-profile-tab,
  #contact:checked ~ nav > #nav-tab > #nav-contact-tab {
    color: #495057;
    background-color: #fff;
    border-color: #dee2e6 #dee2e6 #fff;
  }
  </style>

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <input type="radio" id="home" name="nav-tab">
  <input type="radio" id="profile" name="nav-tab">
  <input type="radio" id="contact" name="nav-tab">
  <nav>
    <div class="nav nav-tabs" id="nav-tab" role="tablist">
      <a class="nav-item nav-link" id="nav-home-tab"><label for="home">Home</label></a>
      <a class="nav-item nav-link" id="nav-profile-tab"><label for="profile">Profile</label></a>
      <a class="nav-item nav-link" id="nav-contact-tab"><label for="contact">Contact</label></a>
    </div>
  </nav>
  <div class="tab-content" id="nav-tabContent">
    <div class="tab-pane fade" id="nav-home">My Home</div>
    <div class="tab-pane fade" id="nav-profile">My Profile</div>
    <div class="tab-pane fade" id="nav-contact">My Contact</div>
  </div>

|
|

The technique comes from *Pure CSS Tab Panel* [1]_. We need:

1 - Visible HTML label_ elements, which is the navigation tabs.

2 - Invisible HTML *input* *radio* elements, referenced by the for_ attribute of
*label* elements. The value of *name* attribute of these elements must be the
same.

3 - Tab contents, invisible in the beginning.

When users click on the visible label elements, the corresponding input radio
element is checked. We use CSS3 **:checked** selector to select corresponding
tab content, and make the selected tab content visible.

The complete source code is as follows:

**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <input type="radio" id="home" name="nav-tab">
  <input type="radio" id="profile" name="nav-tab">
  <input type="radio" id="contact" name="nav-tab">
  <nav>
    <div class="nav nav-tabs" id="nav-tab" role="tablist">
      <a class="nav-item nav-link" id="nav-home-tab"><label for="home">Home</label></a>
      <a class="nav-item nav-link" id="nav-profile-tab"><label for="profile">Profile</label></a>
      <a class="nav-item nav-link" id="nav-contact-tab"><label for="contact">Contact</label></a>
    </div>
  </nav>
  <div class="tab-content" id="nav-tabContent">
    <div class="tab-pane fade" id="nav-home">My Home</div>
    <div class="tab-pane fade" id="nav-profile">My Profile</div>
    <div class="tab-pane fade" id="nav-contact">My Contact</div>
  </div>

.. adsu:: 2

**CSS**:

.. code-block:: css

  input[name="nav-tab"] {
    display: none;
  }

  /* the magic: show tab panes */
  #home:checked ~ #nav-tabContent > #nav-home,
  #profile:checked ~ #nav-tabContent > #nav-profile,
  #contact:checked ~ #nav-tabContent > #nav-contact {
    display: block;
    opacity: 1;
  }

  /* the magic: make tab active */
  #home:checked ~ nav > #nav-tab > #nav-home-tab,
  #profile:checked ~ nav > #nav-tab > #nav-profile-tab,
  #contact:checked ~ nav > #nav-tab > #nav-contact-tab {
    color: #495057;
    background-color: #fff;
    border-color: #dee2e6 #dee2e6 #fff;
  }

- The first rule hides the tab content and input radio box.
- In second rule. We use **:checked** and general sibling selector (~) to make
  the user-selected tab content visible.
- In third rule, we make the selected navigation tab look active.

----

Tested on:

- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``
- ``Bootstrap 4.1.3``

----

.. adsu:: 3

References:

.. [1] `Pure CSS Tab Panel <{filename}/articles/2017/05/21/css-only-tab-panel%en.rst>`_
.. [2] `Pure CSS Bulma Tabs <{filename}/articles/2018/01/30/css-only-bulma-tab-panel%en.rst>`_

.. _label: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
.. _for: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label#Using_the_for_attribute
.. _Bootstrap: https://getbootstrap.com/
.. _tab panel: https://getbootstrap.com/docs/4.1/components/navs/#javascript-behavior
.. _card: https://getbootstrap.com/docs/4.1/components/card/
