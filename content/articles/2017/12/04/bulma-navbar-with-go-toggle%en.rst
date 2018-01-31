Bulma Navbar with Go Toggle
###########################

:date: 2018-01-18T09:10+08:00
:tags: Go, Golang, GopherJS, Go to JavaScript, Frontend Programming in Go, html,
       CSS, Responsive Web Design, toggle, toggleable, Bulma, navbar
:category: Frontend Programming in Go
:summary: Go toggle code for Bulma responsive navbar, dropdown menu included.
:og_image: https://bulma.io/images/blog/dropup.png
:adsu: yes

Go toggle code for Bulma responsive `navbar`_ [1]_, dropdown menu included. The
Go code is compiled to JavaScript via GopherJS.

The full code example, including JavaScript counterpart, of this post is
`on my GitHub`_.


**HTML code for Dropdown**:

.. code-block:: html

  <nav class="navbar is-transparent">
    <div class="navbar-brand">
      <a class="navbar-item" href="https://bulma.io">
        <img src="https://bulma.io/images/bulma-logo.png" alt="Bulma: a modern CSS framework based on Flexbox" width="112" height="28">
      </a>
      <div class="navbar-burger burger" data-target="navbarExampleTransparentExample">
        <span></span>
        <span></span>
        <span></span>
      </div>
    </div>
  
    <div id="navbarExampleTransparentExample" class="navbar-menu">
      <div class="navbar-start">
        <a class="navbar-item" href="https://bulma.io/">
          Home
        </a>
        <div class="navbar-item has-dropdown is-hoverable">
          <a class="navbar-link" href="/documentation/overview/start/">
            Docs
          </a>
          <div class="navbar-dropdown is-boxed">
            <a class="navbar-item" href="/documentation/overview/start/">
              Overview
            </a>
            <a class="navbar-item" href="https://bulma.io/documentation/modifiers/syntax/">
              Modifiers
            </a>
            <a class="navbar-item" href="https://bulma.io/documentation/columns/basics/">
              Columns
            </a>
            <a class="navbar-item" href="https://bulma.io/documentation/layout/container/">
              Layout
            </a>
            <a class="navbar-item" href="https://bulma.io/documentation/form/general/">
              Form
            </a>
            <hr class="navbar-divider">
            <a class="navbar-item" href="https://bulma.io/documentation/elements/box/">
              Elements
            </a>
            <a class="navbar-item is-active" href="https://bulma.io/documentation/components/breadcrumb/">
              Components
            </a>
          </div>
        </div>
        <div class="navbar-item has-dropdown">
          <a class="navbar-link">
            Products
          </a>
          <div class="navbar-dropdown">
            <a class="navbar-item">
              Product 1
            </a>
            <a class="navbar-item">
              Product 2
            </a>
            <a class="navbar-item">
              Product 3
            </a>
          </div>
        </div>
      </div>
  
      <div class="navbar-end">
        <div class="navbar-item">
          <div class="field is-grouped">
            <p class="control">
              <a class="bd-tw-button button" data-social-network="Twitter" data-social-action="tweet" data-social-target="http://localhost:4000" target="_blank" href="https://twitter.com/intent/tweet?text=Bulma: a modern CSS framework based on Flexbox&amp;hashtags=bulmaio&amp;url=http://localhost:4000&amp;via=jgthms">
                <span class="icon">
                  <i class="fab fa-twitter"></i>
                </span>
                <span>
                  Tweet
                </span>
              </a>
            </p>
            <p class="control">
              <a class="button is-primary" href="https://github.com/jgthms/bulma/archive/0.5.1.zip">
                <span class="icon">
                  <i class="fas fa-download"></i>
                </span>
                <span>Download</span>
              </a>
            </p>
          </div>
        </div>
      </div>
    </div>
  </nav>

.. adsu:: 2

**Go toggle code**:

.. code-block:: go

  import (
  	. "github.com/siongui/godom"
  )
  
  func SetupNavbarBurgers() {
  	nbs := Document.QuerySelectorAll(".navbar-burger")
  
  	for _, nb := range nbs {
  		nb.AddEventListener("click", func(e Event) {
  			targetId := nb.Dataset().Get("target").String()
  			target := Document.GetElementById(targetId)
  
  			nb.ClassList().Toggle("is-active")
  			target.ClassList().Toggle("is-active")
  		})
  	}
  }
  
  func main() {
  	Document.AddEventListener("DOMContentLoaded", func(e Event) {
  
  		// Dropdowns in navbar
  		dds := Document.QuerySelectorAll(".navbar-item.has-dropdown:not(.is-hoverable)")
  
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
  
  		// Toggles
  		SetupNavbarBurgers()
  	})
  }

The above code use godom_ package to make the code more readable.

.. adsu:: 3

The following JavaScript code is equivalent to above Go code:

.. code-block:: javascript

  'use strict';
  
  document.addEventListener('DOMContentLoaded', function () {
  
    // Dropdowns in navbar
  
    var $dropdowns = getAll('.navbar-item.has-dropdown:not(.is-hoverable)');
  
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
  
    // Toggles
  
    var $burgers = getAll('.burger');
  
    if ($burgers.length > 0) {
      $burgers.forEach(function ($el) {
        $el.addEventListener('click', function () {
          var target = $el.dataset.target;
          var $target = document.getElementById(target);
          $el.classList.toggle('is-active');
          $target.classList.toggle('is-active');
        });
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

.. [1] `Navbar | Bulma: a modern CSS framework based on Flexbox <https://bulma.io/documentation/components/navbar/>`_
.. [2] `https://bulma.io/lib/main.js?v=201801161752 <https://bulma.io/lib/main.js?v=201801161752>`_

.. _Bulma: https://bulma.io/
.. _navbar: https://bulma.io/documentation/components/navbar/
.. _godom: https://github.com/siongui/godom
.. _on my GitHub: https://github.com/siongui/frontend-programming-in-go/tree/master/021-bulma-navbar
