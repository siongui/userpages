[JavaScript] Tab Panel
######################

:date: 2017-05-28T23:24+08:00
:tags: JavaScript, html, CSS, Tab Panel
:category: JavaScript
:summary: Tab panel implementation via JavaScript and CSS.
:og_image: https://s3-us-west-2.amazonaws.com/i.cdpn.io/193912.BlKxo.0e781013-9523-4f00-b75c-0c2f48b25100.png
:adsu: yes

Tab panel implementation via JavaScript and CSS.
The tab panel here is similar to `Bootstrap tab`_.
Click the following tabs to see the demo.

.. raw:: html

  <style>
  .nav-tabs {
    display: flex;
    flex-wrap: wrap;
    margin-bottom: 15px;
    border-bottom: 1px solid #ddd;
    list-style: none;
  }
  .nav-tabs > li {
    margin-bottom: 1px;
    margin-right: 1rem;
    line-height: 2rem;
  }
  .nav-tabs > li > a {
    cursor: pointer;
    text-decoration: none;
  }
  .tab-pane {
    padding: 1rem;
    display: none;
  }
  .tab-content > .active {
    display: block;
  }
  </style>

  <div style="background-color: Azure; padding: 1rem;">

  <div>

    <!-- Nav tabs -->
    <ul class="nav-tabs">
      <li><a data-target="home">Home</a></li>
      <li><a data-target="profile">Profile</a></li>
      <li><a data-target="messages">Messages</a></li>
      <li><a data-target="settings">Settings</a></li>
    </ul>

    <!-- Tab panes -->
    <div class="tab-content">
      <div class="tab-pane active" id="home">My home</div>
      <div class="tab-pane" id="profile">My profile</div>
      <div class="tab-pane" id="messages">My messages</div>
      <div class="tab-pane" id="settings">My settings</div>
    </div>

  </div>

  </div>


.. raw:: html

  <script>
  'use strict';

  function setupNavtabs(navtab) {
    var tabs = navtab.querySelectorAll("*[data-target]");
    for (var i=0; i < tabs.length; i++) {
      var tab = tabs[i];

      tab.addEventListener("click", function(e) {
        var id = e.target.dataset.target;
        var pane = document.getElementById(id);
        pane.classList.add("active");

        // remove .active class of all sibling elements
        var el = pane.nextElementSibling;
        while (el) {
          el.classList.remove("active");
          el = el.nextElementSibling;
        }
        el = pane.previousElementSibling;
        while (el) {
          el.classList.remove("active");
          el = el.previousElementSibling;
        }
      });
    }
  }

  var navtabs = document.querySelectorAll(".nav-tabs");
  for (var i=0; i < navtabs.length; i++) {
    setupNavtabs(navtabs[i]);
  }
  </script>

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <div>

    <!-- Nav tabs -->
    <ul class="nav-tabs">
      <li><a data-target="home">Home</a></li>
      <li><a data-target="profile">Profile</a></li>
      <li><a data-target="messages">Messages</a></li>
      <li><a data-target="settings">Settings</a></li>
    </ul>

    <!-- Tab panes -->
    <div class="tab-content">
      <div class="tab-pane active" id="home">My home</div>
      <div class="tab-pane" id="profile">My profile</div>
      <div class="tab-pane" id="messages">My messages</div>
      <div class="tab-pane" id="settings">My settings</div>
    </div>

  </div>

The tab has ``data-target`` attribute to indicate the *id* of the element with
tab content.

.. adsu:: 2

**CSS**:

.. code-block:: css

  .nav-tabs {
    display: flex;
    flex-wrap: wrap;
    margin-bottom: 15px;
    border-bottom: 1px solid #ddd;
    list-style: none;
  }
  .nav-tabs > li {
    margin-bottom: 1px;
    margin-right: 1rem;
    line-height: 2rem;
  }
  .nav-tabs > li > a {
    cursor: pointer;
    text-decoration: none;
  }
  .tab-pane {
    padding: 1rem;
    display: none;
  }
  .tab-content > .active {
    display: block;
  }

The first three rules use flexbox to align the tabs in one row.
The last two rules hide the tab contents except the tab content with ``.active``
class.

**JavaScript**:

.. code-block:: javascript

  'use strict';

  function setupNavtabs(navtab) {
    var tabs = navtab.querySelectorAll("*[data-target]");
    for (var i=0; i < tabs.length; i++) {
      var tab = tabs[i];

      tab.addEventListener("click", function(e) {
        var id = e.target.dataset.target;
        var pane = document.getElementById(id);
        pane.classList.add("active");

        // remove .active class of all sibling elements
        var el = pane.nextElementSibling;
        while (el) {
          el.classList.remove("active");
          el = el.nextElementSibling;
        }
        el = pane.previousElementSibling;
        while (el) {
          el.classList.remove("active");
          el = el.previousElementSibling;
        }
      });
    }
  }

  var navtabs = document.querySelectorAll(".nav-tabs");
  for (var i=0; i < navtabs.length; i++) {
    setupNavtabs(navtabs[i]);
  }

The JS code here queries for elements with ``data-target`` attribute in
element with ``.nav-tabs``. When users click on the element with ``data-target``
attribute, add ``.active`` class to the tab pane referenced by the attribute,
and remove ``.active`` class of all sibling tab panes.

.. adsu:: 3

----

Tested on:

- ``Chromium 58.0.3029.110 on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] `[Vue.js] Tab Panel <{filename}../16/vuejs-tab-panel%en.rst>`_

.. _Bootstrap tab: https://getbootstrap.com/docs/3.3/javascript/#tabs
