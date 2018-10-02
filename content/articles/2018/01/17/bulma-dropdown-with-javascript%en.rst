JavaScript for Bulma Dropdown
#############################

:date: 2018-01-19T05:33+08:00
:modified: 2018-10-02T19:13+08:00
:tags: CSS, html, JavaScript, Bulma, dropdown menu
:category: CSS
:summary: JavaScript code for Bulma dropdown menu.
          The code is extracted from Bulma official website.
:og_image: https://cdn.scotch.io/1/r2pUTqX2QWuqxEvuibtp_SpiVz74.png
:adsu: yes


This post is for those who do not want to write their own JavaScript code, and
just want to use the same code as that in Bulma_ official website [1]_.

The JavaScript code comes from *main.js* used in the official site [2]_.
I put the code here in the post  for easy search by Google or other search
engines.

If you do not want to toggle with JavaScript, you can use pure css solution.
See [3]_.

.. code-block:: javascript

  'use strict';
  
  document.addEventListener('DOMContentLoaded', function () {
  
    // Dropdowns
  
    var $dropdowns = getAll('.dropdown:not(.is-hoverable)');
  
    if ($dropdowns.length > 0) {
      $dropdowns.forEach(function ($el) {
        $el.addEventListener('click', function (event) {
          event.stopPropagation();
          $el.classList.toggle('is-active');
        });
      });
  
      document.addEventListener('click', function (event) {
        closeDropdowns();
      });
    }
  
    function closeDropdowns() {
      $dropdowns.forEach(function ($el) {
        $el.classList.remove('is-active');
      });
    }

    // Close dropdowns if ESC pressed
    document.addEventListener('keydown', function (event) {
      var e = event || window.event;
      if (e.keyCode === 27) {
        closeDropdowns();
      }
    });
  
    // Functions
  
    function getAll(selector) {
      return Array.prototype.slice.call(document.querySelectorAll(selector), 0);
    }
  });

**Demo**:

.. raw:: html

  <br>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">
  <div class="dropdown">
    <div class="dropdown-trigger">
      <button class="button" aria-haspopup="true" aria-controls="dropdown-menu">
        <span>Dropdown button</span>
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

  <br><br>

  <script>
  'use strict';
  
  document.addEventListener('DOMContentLoaded', function () {
  
    // Dropdowns
  
    var $dropdowns = getAll('.dropdown:not(.is-hoverable)');
  
    if ($dropdowns.length > 0) {
      $dropdowns.forEach(function ($el) {
        $el.addEventListener('click', function (event) {
          event.stopPropagation();
          $el.classList.toggle('is-active');
        });
      });
  
      document.addEventListener('click', function (event) {
        closeDropdowns();
      });
    }
  
    function closeDropdowns() {
      $dropdowns.forEach(function ($el) {
        $el.classList.remove('is-active');
      });
    }
  
    // Close dropdowns if ESC pressed
    document.addEventListener('keydown', function (event) {
      var e = event || window.event;
      if (e.keyCode === 27) {
        closeDropdowns();
      }
    });

    // Functions
  
    function getAll(selector) {
      return Array.prototype.slice.call(document.querySelectorAll(selector), 0);
    }
  });
  </script>

----

Tested on:

- ``Chromium 63.0.3239.84 on Ubuntu 17.10 (64-bit)``
- ``Bulma 0.6.2``

----

.. adsu:: 2

**References**:

.. [1] `Dropdown | Bulma: a modern CSS framework based on Flexbox <https://bulma.io/documentation/components/dropdown/>`_
.. [2] `https://bulma.io/lib/main.js?v=201801161752 <https://bulma.io/lib/main.js?v=201801161752>`_
.. [3] `Pure CSS Bulma Dropdown Toggle <{filename}/articles/2018/10/02/css-only-toggle-bulma-dropdown%en.rst>`_

.. _Bulma: http://bulma.io/
.. _Transparent navbar: https://bulma.io/documentation/components/navbar/#transparent-navbar
