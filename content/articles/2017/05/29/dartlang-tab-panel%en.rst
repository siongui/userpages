[Dart] Tab Panel
################

:date: 2017-05-29T06:04+08:00
:tags: Dart, html, CSS, Tab Panel
:category: Dart
:summary: Tab panel implementation via Dart and CSS.
:og_image: https://s3-us-west-2.amazonaws.com/i.cdpn.io/193912.BlKxo.0e781013-9523-4f00-b75c-0c2f48b25100.png
:adsu: yes

Tab panel implementation via Dart and CSS.
The tab panel here is similar to `Bootstrap tab`_.

.. rubric:: `Run on DartPad <https://dartpad.dartlang.org/9538286a6fbb78556559eae8510533d2>`__
   :class: align-center

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

**Dart**:

.. code-block:: dart

  import 'dart:html';

  void tabClick(Event e) {
    Element elm = e.target;
    var id = elm.dataset["target"];
    var pane = querySelector("#" + id);
    pane.classes.add("active");

    // remove .active class of all sibling elements
    var el = pane.nextElementSibling;
    while (el != null) {
      el.classes.remove("active");
      el = el.nextElementSibling;
    }
    el = pane.previousElementSibling;
    while (el != null) {
      el.classes.remove("active");
      el = el.previousElementSibling;
    }
  }

  void setupNavtab(Element navtab) {
    var tabs = navtab.querySelectorAll("*[data-target]");
    for (var tab in tabs) {
      tab.onClick.listen(tabClick);
    }
  }

  void main() {
    var navtabs = querySelectorAll(".nav-tabs");
    for (var navtab in navtabs) {
      setupNavtab(navtab);
    }
  }

The Dart code here queries for elements with ``data-target`` attribute in
element with ``.nav-tabs``. When users click on the element with ``data-target``
attribute, add ``.active`` class to the tab pane referenced by the attribute,
and remove ``.active`` class of all sibling tab panes.

.. adsu:: 3

----

Tested on:

- DartPad_.
- ``Chromium 58.0.3029.110 Built on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] `[JavaScript] Tab Panel <{filename}../28/javascript-tab-panel%en.rst>`_
.. [2] `Synonyms - Dart, JavaScript, C#, Python | Dart <https://www.dartlang.org/resources/synonyms>`_
.. [3] `Improving the DOM | webdev.dartlang.org <https://webdev.dartlang.org/articles/low-level-html/improving-the-dom>`_
.. [4] `Strings, Numbers, Booleans, Oh My! - Dart Tips, Episode 4 | Dart <https://www.dartlang.org/resources/dart-tips/dart-tips-ep-4>`_

.. _DartPad: https://dartpad.dartlang.org/
.. _Bootstrap tab: https://getbootstrap.com/docs/3.3/javascript/#tabs
