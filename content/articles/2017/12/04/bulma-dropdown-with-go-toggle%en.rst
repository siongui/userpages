Bulma Dropdown with Go Toggle
#############################

:date: 2018-02-01T06:18+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, html,
       Bulma, dropdown menu, CSS
:category: Frontend Programming in Go
:summary: Go toggle code for Bulma *dropdown* component.
:og_image: https://cdn.scotch.io/1/r2pUTqX2QWuqxEvuibtp_SpiVz74.png
:adsu: yes

Go toggle code for Bulma dropdown_ component, compiled to JavaScript via
GopherJS.

The full code example, including JavaScript counterpart, of this post is
`on my GitHub`_.


**HTML code for Dropdown**:

.. code-block:: html

  <div class="dropdown is-active">
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

.. adsu:: 2

**Go toggle code**:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )
  
  func main() {
  	Document.AddEventListener("DOMContentLoaded", func(e Event) {
  
  		// Dropdowns
  		dds := Document.QuerySelectorAll(".dropdown:not(.is-hoverable)")
  
  		closeDropdowns := func() {
  			for _, dd := range dds {
  				dd.ClassList().Remove("is-active")
  			}
  		}
  
  		if len(dds) > 0 {
  			for _, dd := range dds {
  				dd.AddEventListener("click", func(e Event) {
  					e.StopPropagation()
  					dd.ClassList().Toggle("is-active")
  				})
  			}
  
  			Document.AddEventListener("click", func(e Event) {
  				closeDropdowns()
  			})
  		}
  
  		// Close dropdowns if ESC pressed
  		Document.AddEventListener("keydown", func(e Event) {
  			if e.KeyCode() == 27 {
  				closeDropdowns()
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

.. adsu:: 4

----

References:

.. [1] `Dropdown | Bulma: a modern CSS framework based on Flexbox <https://bulma.io/documentation/components/dropdown/>`_
.. [2] `https://bulma.io/lib/main.js?v=201801161752 <https://bulma.io/lib/main.js?v=201801161752>`_

.. _Bulma: https://bulma.io/
.. _dropdown: https://bulma.io/documentation/components/dropdown/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/024-bulma-dropdown
