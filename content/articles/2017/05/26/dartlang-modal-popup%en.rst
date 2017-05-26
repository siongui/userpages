[Dart] Modal (Popup)
####################

:date: 2017-05-26T21:32+08:00
:tags: Dart, html, DOM, CSS, Modal (Popup)
:category: Dart
:summary: Modal (Popup) implementation via Dart and CSS. Modal is dialog
          box/popup window that is displayed on top of the current page.
:og_image: https://dab1nmslvvntp.cloudfront.net/wp-content/uploads/2014/02/1392433498bootstrap-modal-basic.jpg
:adsu: yes

Modal (Popup) implementation via Dart and CSS. `According to w3schools`_,
modal is dialog box/popup window that is displayed on top of the current page.
The modal here is a simplified version of `Bootstrap modal`_.

.. rubric:: `Run on DartPad <https://dartpad.dartlang.org/102d99b28b485737aaf28fab9ce98b5e>`__
   :class: align-center

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
    color: black;
  }

The CSS here is to center the modal horizontally and vertically [2]_, and also
set the modal invisible in the beginning.

.. adsu:: 2

**Dart**:

.. code-block:: dart

  import 'dart:html';

  void launchModal(Event e) {
    // when user click launch button, show modal.
    ButtonElement elm = e.target;
    var selector = elm.dataset["target"];
    var modal = querySelector(selector);
    modal.style.display = "block";

    if (modal.dataset["isSetClose"] != "true") {
      // when user click button with `data-dismiss="modal"` in the modal,
      // close the modal
      var closeBtns = modal.querySelectorAll("button[data-dismiss='modal']");
      for (var closeBtn in closeBtns) {
        closeBtn.onClick.listen((e) => modal.style.display = "none");
      }

      modal.dataset["isSetClose"] = "true";
    }
  }

  void main() {
    var launchBtns = querySelectorAll("button[data-toggle='modal']");
  
    for (var launchBtn in launchBtns) {
      launchBtn.onClick.listen(launchModal);
    }
  }

First we search for button elements with *data-toggle="modal"*. If users click
on such buttons, show the modal referenced by *data-target* attribute of the
button element. Also in the click event handler of modal launching buttons, we
set the click event handler of modal closing button, which has the
*data-dismiss="modal"* attribute in the modal HTML code.

.. adsu:: 3

----

Tested on:

- DartPad_.
- ``Chromium Version 58.0.3029.110 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] `[JavaScript] Modal (Popup) <{filename}../13/javascript-modal-popup%en.rst>`_
.. [2] `html - How to center an element horizontally and vertically? - Stack Overflow <http://stackoverflow.com/questions/19461521/how-to-center-an-element-horizontally-and-vertically>`_
.. [3] `Synonyms - Dart, JavaScript, C#, Python | Dart <https://www.dartlang.org/resources/synonyms>`_

.. _dartlang: https://www.dartlang.org/
.. _DartPad: https://dartpad.dartlang.org/
.. _According to w3schools: https://www.w3schools.com/bootstrap/bootstrap_modal.asp
.. _Bootstrap modal: http://getbootstrap.com/javascript/#modals
