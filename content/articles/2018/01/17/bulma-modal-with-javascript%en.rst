JavaScript for Bulma Modal
##########################

:date: 2018-02-11T08:53+08:00
:tags: CSS, html, JavaScript, Bulma, Modal (Popup)
:category: CSS
:summary: JavaScript code for Bulma *modal* component.
          The code is extracted from Bulma official website.
:og_image: https://www.designil.com/wp-content/uploads/2016/10/modal-box-bulma.png
:adsu: yes


This post is for those who do not want to write their own JavaScript code, and
just want to use the same code as that in Bulma_ official website [1]_.

The JavaScript code comes from *main.js* used in the official site [2]_.
I put the code here in the post  for easy search by Google or other search
engines.

.. code-block:: javascript

  'use strict';
  
  document.addEventListener('DOMContentLoaded', function () {
  
    // Modals
  
    var rootEl = document.documentElement;
    var $modals = getAll('.modal');
    var $modalButtons = getAll('.modal-button');
    var $modalCloses = getAll('.modal-background, .modal-close, .modal-card-head .delete, .modal-card-foot .button');
  
    if ($modalButtons.length > 0) {
      $modalButtons.forEach(function ($el) {
        $el.addEventListener('click', function () {
          var target = $el.dataset.target;
          var $target = document.getElementById(target);
          rootEl.classList.add('is-clipped');
          $target.classList.add('is-active');
        });
      });
    }
  
    if ($modalCloses.length > 0) {
      $modalCloses.forEach(function ($el) {
        $el.addEventListener('click', function () {
          closeModals();
        });
      });
    }
  
    document.addEventListener('keydown', function (event) {
      var e = event || window.event;
      if (e.keyCode === 27) {
        closeModals();
      }
    });
  
    function closeModals() {
      rootEl.classList.remove('is-clipped');
      $modals.forEach(function ($el) {
        $el.classList.remove('is-active');
      });
    }
  
    // Functions
  
    function getAll(selector) {
      return Array.prototype.slice.call(document.querySelectorAll(selector), 0);
    }
  
  });

**Demo**:

.. raw:: html

  <br>
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">
  <a class="button is-primary is-large modal-button" data-target="modal-bis">Launch image modal</a>

  <div id="modal-bis" class="modal">
    <div class="modal-background"></div>
    <div class="modal-content">
      <p class="image is-4by3">
        <img src="https://bulma.io/images/placeholders/1280x960.png">
      </p>
    </div>
    <button class="modal-close is-large" aria-label="close"></button>
  </div>

  <br><br>

  <script>
  'use strict';
  
  document.addEventListener('DOMContentLoaded', function () {
  
    // Modals
  
    var rootEl = document.documentElement;
    var $modals = getAll('.modal');
    var $modalButtons = getAll('.modal-button');
    var $modalCloses = getAll('.modal-background, .modal-close, .modal-card-head .delete, .modal-card-foot .button');
  
    if ($modalButtons.length > 0) {
      $modalButtons.forEach(function ($el) {
        $el.addEventListener('click', function () {
          var target = $el.dataset.target;
          var $target = document.getElementById(target);
          rootEl.classList.add('is-clipped');
          $target.classList.add('is-active');
        });
      });
    }
  
    if ($modalCloses.length > 0) {
      $modalCloses.forEach(function ($el) {
        $el.addEventListener('click', function () {
          closeModals();
        });
      });
    }
  
    document.addEventListener('keydown', function (event) {
      var e = event || window.event;
      if (e.keyCode === 27) {
        closeModals();
      }
    });
  
    function closeModals() {
      rootEl.classList.remove('is-clipped');
      $modals.forEach(function ($el) {
        $el.classList.remove('is-active');
      });
    }
  
    // Functions
  
    function getAll(selector) {
      return Array.prototype.slice.call(document.querySelectorAll(selector), 0);
    }
  
  });
  </script>

----

Tested on:

- ``Chromium 64.0.3282.140 on Ubuntu 17.10 (64-bit)``
- ``Bulma 0.6.2``

----

.. adsu:: 2

**References**:

.. [1] `Modal | Bulma: a modern CSS framework based on Flexbox <https://bulma.io/documentation/components/modal/>`_
.. [2] `https://bulma.io/lib/main.js?v=201802091742 <https://bulma.io/lib/main.js?v=201802091742>`_

.. _Bulma: https://bulma.io/
.. _modal: https://bulma.io/documentation/components/modal/
