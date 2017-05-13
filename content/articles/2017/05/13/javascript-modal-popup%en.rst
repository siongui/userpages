[JavaScript] Modal (Popup)
##########################

:date: 2017-05-13T22:20+08:00
:tags: JavaScript, html, DOM, CSS, Modal (Popup)
:category: JavaScript
:summary: Modal (Popup) implementation via JavaScript and CSS. Modal is dialog
          box/popup window that is displayed on top of the current page.
:og_image: https://dab1nmslvvntp.cloudfront.net/wp-content/uploads/2014/02/1392433498bootstrap-modal-basic.jpg
:adsu: yes

Modal (Popup) implementation via JavaScript and CSS. `According to w3schools`_,
modal is dialog box/popup window that is displayed on top of the current page.
The modal here is a simplified version of `Bootstrap modal`_.

.. raw:: html

  <style>
  .modal {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
    background: azure;
    padding: 1rem;
    border-radius: 1rem;
    border: 1px solid gray;
    display: none;
  }
  </style>


  <!-- Button trigger modal -->
  <button type="button" data-toggle="modal" data-target="#myModal">
    Launch demo modal
  </button>

  <!-- Modal -->
  <div class="modal" id="myModal">
    <h4>Modal Title</h4>
    <p>
       Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas turpis
       felis, tempus sit amet sollicitudin quis, aliquet ut felis.
       Nam a malesuada sem.
    </p>
    <button type="button" data-dismiss="modal">Close</button>
  </div>


.. raw:: html

  <script>
  'use strict';

  var launchBtns = document.querySelectorAll("button[data-toggle='modal']");
  for (var i=0; i < launchBtns.length; i++) {
    var launchBtn = launchBtns[i];

    // when user click launch button, show modal.
    launchBtn.addEventListener("click", function(e) {
      var selector = e.target.dataset.target;
      var modal = document.querySelector(selector);
      modal.style.display = "block";

      if (modal.dataset.isSetClose != "true") {
        // when user click button with `data-dismiss="modal"` in the modal,
        // close the modal
        var closeBtns = modal.querySelectorAll("button[data-dismiss='modal']");
        for (var j=0; j < closeBtns.length; j++) {
          var closeBtn = closeBtns[j];
          closeBtn.addEventListener("click", function(e) {
            modal.style.display = "none";
          });
        }

        modal.dataset.isSetClose = "true";
      }
    });
  }

  </script>

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <!-- Button trigger modal -->
  <button type="button" data-toggle="modal" data-target="#myModal">
    Launch demo modal
  </button>

  <!-- Modal -->
  <div class="modal" id="myModal">
    <h4>Modal Title</h4>
    <p>
       Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas turpis
       felis, tempus sit amet sollicitudin quis, aliquet ut felis.
       Nam a malesuada sem.
    </p>
    <button type="button" data-dismiss="modal">Close</button>
  </div>

The HTML code here is simplified version of `Bootstrap modal`_, we use *button*
element with *data-toggle* and *data-target* attributes to launch the modal, and
button element with *data-dismiss* attribute in the modal to close the modal.

.. adsu:: 2

**CSS**:

.. code-block:: css

  .modal {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translateX(-50%) translateY(-50%);
    background: azure;
    padding: 1rem;
    border-radius: 1rem;
    border: 1px solid gray;
    display: none;
  }

The CSS here is to center the modal horizontally and vertically [2]_, and also
set the modal invisible in the beginning.

**JavaScript**:

.. code-block:: javascript

  'use strict';

  var launchBtns = document.querySelectorAll("button[data-toggle='modal']");
  for (var i=0; i < launchBtns.length; i++) {
    var launchBtn = launchBtns[i];

    // when user click launch button, show modal.
    launchBtn.addEventListener("click", function(e) {
      var selector = e.target.dataset.target;
      var modal = document.querySelector(selector);
      modal.style.display = "block";

      if (modal.dataset.isSetClose != "true") {
        // when user click button with `data-dismiss="modal"` in the modal,
        // close the modal
        var closeBtns = modal.querySelectorAll("button[data-dismiss='modal']");
        for (var j=0; j < closeBtns.length; j++) {
          var closeBtn = closeBtns[j];
          closeBtn.addEventListener("click", function(e) {
            modal.style.display = "none";
          });
        }

        modal.dataset.isSetClose = "true";
      }
    });
  }

First we search for button elements with *data-toggle="modal"*. If users click
on such buttons, show the modal referenced by *data-target* attribute of the
button element. Also in the click event handler of modal launching buttons, we
set the click event handler of modal closing button, which has the
*data-dismiss="modal"* attribute in the modal HTML code.

----

Tested on:

- ``Chromium Version 58.0.3029.96 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

.. adsu:: 3

References:

.. [1] | `center div horizontally and vertically - Google search <https://www.google.com/search?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - DuckDuckGo search <https://duckduckgo.com/?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Ecosia search <https://www.ecosia.org/search?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Qwant search <https://www.qwant.com/?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Bing search <https://www.bing.com/search?q=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Yahoo search <https://search.yahoo.com/search?p=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Baidu search <https://www.baidu.com/s?wd=center+div+horizontally+and+vertically>`_
       | `center div horizontally and vertically - Yandex search <https://www.yandex.com/search/?text=center+div+horizontally+and+vertically>`_

.. [2] `html - How to center an element horizontally and vertically? - Stack Overflow <http://stackoverflow.com/questions/19461521/how-to-center-an-element-horizontally-and-vertically>`_

.. _Vue.js: https://vuejs.org/
.. _According to w3schools: https://www.w3schools.com/bootstrap/bootstrap_modal.asp
.. _Bootstrap modal: http://getbootstrap.com/javascript/#modals
