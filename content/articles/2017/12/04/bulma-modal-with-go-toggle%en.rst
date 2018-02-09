Bulma Modal with Go Toggle
##########################

:date: 2018-02-10T07:52+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, html,
       Bulma, CSS, Modal (Popup)
:category: Frontend Programming in Go
:summary: Go toggle code for Bulma *modal* component.
:og_image: https://www.designil.com/wp-content/uploads/2016/10/modal-box-bulma.png
:adsu: yes

Go toggle code for Bulma_ modal_ component (version 0.6.2),
compiled to JavaScript via GopherJS.

The full code example, including JavaScript counterpart, of this post is
`on my GitHub`_.


**HTML code for modal**:

.. code-block:: html

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

.. adsu:: 2

**Go toggle code**:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )
  
  func main() {
  	Document.AddEventListener("DOMContentLoaded", func(e Event) {
  
  		// Modals
  
  		rootEl := Document.Get("documentElement")
  		modals := Document.QuerySelectorAll(".modal")
  		modalButtons := Document.QuerySelectorAll(".modal-button")
  		modalCloses := Document.QuerySelectorAll(".modal-background, .modal-close, .modal-card-head .delete, .modal-card-foot .button")
  
  		closeModals := func() {
  			rootEl.Get("classList").Call("remove", "is-clipped")
  			for _, modal := range modals {
  				modal.ClassList().Remove("is-active")
  			}
  		}
  
  		if len(modalButtons) > 0 {
  			for _, modalButton := range modalButtons {
  				modalButton.AddEventListener("click", func(e Event) {
  					targetId := modalButton.Dataset().Get("target").String()
  					target := Document.GetElementById(targetId)
  					rootEl.Get("classList").Call("add", "is-clipped")
  					target.ClassList().Add("is-active")
  				})
  			}
  		}
  
  		if len(modalCloses) > 0 {
  			for _, modalClose := range modalCloses {
  				modalClose.AddEventListener("click", func(e Event) {
  					closeModals()
  				})
  			}
  		}
  
  		// Close modals if ESC pressed
  		Document.AddEventListener("keydown", func(e Event) {
  			if e.KeyCode() == 27 {
  				closeModals()
  			}
  		})
  	})
  }

The above code use godom_ package to make the code more readable.

.. adsu:: 3

The following JavaScript code is equivalent to above Go code:

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

.. adsu:: 4

----

References:

.. [1] `Modal | Bulma: a modern CSS framework based on Flexbox <https://bulma.io/documentation/components/modal/>`_
.. [2] `https://bulma.io/lib/main.js?v=201802091742 <https://bulma.io/lib/main.js?v=201802091742>`_

.. _Bulma: https://bulma.io/
.. _modal: https://bulma.io/documentation/components/modal/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/026-bulma-modal
