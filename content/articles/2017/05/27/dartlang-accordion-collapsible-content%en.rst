[Dart] Accordion (Collapsible Content)
######################################

:date: 2017-05-27T23:14+08:00
:tags: Dart, html, CSS, Accordion (Collapsible Content)
:category: Dart
:summary: Accordion (collapsible content) implementation via Dart.
:og_image: http://christineosazuwa.com/wp-content/uploads/2015/05/13-06_accordionmenu1.jpg
:adsu: yes

Accordion (collapsible content) implementation via Dart.
The accordion here is similar to the accordion example in `Bootstrap Collapse`_.

.. rubric:: `Run on DartPad <https://dartpad.dartlang.org/c460f6d354148c8d74dae92543c7fb42>`__
   :class: align-center

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

**Dart**:

.. code-block:: dart

  import 'dart:html';

  void setupAccordion(Element accordion) {
    var titles = accordion.querySelectorAll(".panel-title");
    for (var title in titles) {
      title.dataset["index"] = titles.indexOf(title).toString();
      title.onClick.listen((Event e) {
        Element elm = e.target;
        var bodys = accordion.querySelectorAll(".panel-body");
        for (var body in bodys) {
          if (elm.dataset["index"] == bodys.indexOf(body).toString()) {
            if (body.style.display == "block") {
              body.style.display = "none";
            } else {
              body.style.display = "block";
            }
          } else {
            body.style.display = "none";
          }
        }
      });
    }
  }

  void main() {
    var accordions = querySelectorAll(".accordion");

    for (var accordion in accordions) {
      setupAccordion(accordion);
    }
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

- DartPad_.
- ``Chromium Version 58.0.3029.110 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``

----

References:

.. [1] `[JavaScript] Accordion (Collapsible Content) <{filename}../24/javascript-accordion-collapsible-content%en.rst>`_
.. [2] `Synonyms - Dart, JavaScript, C#, Python | Dart <https://www.dartlang.org/resources/synonyms>`_
.. [3] `Improving the DOM | webdev.dartlang.org <https://webdev.dartlang.org/articles/low-level-html/improving-the-dom>`_
.. [4] `Strings, Numbers, Booleans, Oh My! - Dart Tips, Episode 4 | Dart <https://www.dartlang.org/resources/dart-tips/dart-tips-ep-4>`_

.. _DartPad: https://dartpad.dartlang.org/
.. _Bootstrap Collapse: http://getbootstrap.com/javascript/#collapse
