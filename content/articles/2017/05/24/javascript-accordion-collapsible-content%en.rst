[JavaScript] Accordion (Collapsible Content)
############################################

:date: 2017-05-24T05:29+08:00
:tags: JavaScript, html, CSS, Accordion (Collapsible Content)
:category: JavaScript
:summary: Accordion (collapsible content) implementation via JavaScript.
:og_image: http://christineosazuwa.com/wp-content/uploads/2015/05/13-06_accordionmenu1.jpg
:adsu: yes

Accordion (collapsible content) implementation via JavaScript.
The accordion here is similar to the accordion example in `Bootstrap Collapse`_.
Click the following panels to see the demo.

.. raw:: html

  <style>
  .panel {
    margin-bottom: 1rem;
    border: 1px solid #ccc;
  }
  .panel-title {
    font-weight: bold;
    background-color: #ccc;
    padding: 0.01em 16px;
    cursor: pointer;
  }
  .panel-body {
    padding: 0.01em 16px;
    display: none;
  }
  </style>

  <div class="accordion">
    <div class="panel">
      <div class="panel-title">
        Title #1
      </div>
      <div class="panel-body">
        Content Body #1
      </div>
    </div>
    <div class="panel">
      <div class="panel-title">
        Title #2
      </div>
      <div class="panel-body">
        Content Body #2
      </div>
    </div>
    <div class="panel">
      <div class="panel-title">
        Title #3
      </div>
      <div class="panel-body">
        Content Body #3
      </div>
    </div>
  </div>


.. raw:: html

  <script>
  'use strict';

  function setupAccordion(accordion) {
    var titles = accordion.querySelectorAll(".panel-title");
    for (var i=0; i < titles.length; i++) {
      var title = titles[i];
      title.dataset.index = i;
      title.addEventListener("click", function(e) {
        var bodys = accordion.querySelectorAll(".panel-body");
        for (var j=0; j < bodys.length; j++) {
          if (j == e.target.dataset.index) {
            if (bodys[j].style.display == "block") {
              bodys[j].style.display = "none";
            } else {
              bodys[j].style.display = "block";
            }
          } else {
            bodys[j].style.display = "none";
          }
        }
      })
    }
  }

  var accordions = document.querySelectorAll(".accordion");
  for (var i=0; i < accordions.length; i++) {
    setupAccordion(accordions[i])
  }
  </script>

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <div class="accordion">
    <div class="panel">
      <div class="panel-title">
        Title #1
      </div>
      <div class="panel-body">
        Content Body #1
      </div>
    </div>
    <div class="panel">
      <div class="panel-title">
        Title #2
      </div>
      <div class="panel-body">
        Content Body #2
      </div>
    </div>
    <div class="panel">
      <div class="panel-title">
        Title #3
      </div>
      <div class="panel-body">
        Content Body #3
      </div>
    </div>
  </div>

Wrap all your panels in the element with ``.accordion`` class.
The panels consist of ``div.panel``, and there are ``div.panel-title`` and
``div.panel-body`` inside ``div.panel``.

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  function setupAccordion(accordion) {
    var titles = accordion.querySelectorAll(".panel-title");
    for (var i=0; i < titles.length; i++) {
      var title = titles[i];
      title.dataset.index = i;
      title.addEventListener("click", function(e) {
        var bodys = accordion.querySelectorAll(".panel-body");
        for (var j=0; j < bodys.length; j++) {
          if (j == e.target.dataset.index) {
            if (bodys[j].style.display == "block") {
              bodys[j].style.display = "none";
            } else {
              bodys[j].style.display = "block";
            }
          } else {
            bodys[j].style.display = "none";
          }
        }
      })
    }
  }

  var accordions = document.querySelectorAll(".accordion");
  for (var i=0; i < accordions.length; i++) {
    setupAccordion(accordions[i])
  }

Find all elements with ``.accordion`` class, and setup the event listeners of
panel titles. When users click on the title of the panel, if the body of the
panel is hidden, show the panel body and hide all other panel bodys. Otherwise
hide all panel bodys.

**CSS**:

.. code-block:: css

  .panel {
    margin-bottom: 1rem;
    border: 1px solid #ccc;
  }
  .panel-title {
    font-weight: bold;
    background-color: #ccc;
    padding: 0.01em 16px;
    cursor: pointer;
  }
  .panel-body {
    padding: 0.01em 16px;
    display: none;
  }

Nothing special in CSS code here. For demo purpose, I make CSS very simple. You
can try to add some animation if you want.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 58.0.3029.110 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] | `Accordion - Google search <https://www.google.com/search?q=Accordion>`_
       | `Accordion - DuckDuckGo search <https://duckduckgo.com/?q=Accordion>`_
       | `Accordion - Ecosia search <https://www.ecosia.org/search?q=Accordion>`_
       | `Accordion - Qwant search <https://www.qwant.com/?q=Accordion>`_
       | `Accordion - Bing search <https://www.bing.com/search?q=Accordion>`_
       | `Accordion - Yahoo search <https://search.yahoo.com/search?p=Accordion>`_
       | `Accordion - Baidu search <https://www.baidu.com/s?wd=Accordion>`_
       | `Accordion - Yandex search <https://www.yandex.com/search/?text=Accordion>`_
.. [2] `How To Create an Accordion - W3Schools <https://www.w3schools.com/howto/howto_js_accordion.asp>`_

.. _Vue.js: https://vuejs.org/
.. _Bootstrap Collapse: https://getbootstrap.com/docs/3.3/javascript/#collapse-example-accordion
