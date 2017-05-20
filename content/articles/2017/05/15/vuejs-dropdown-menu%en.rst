[Vue.js] Dropdown Menu
######################

:date: 2017-05-15T21:28+08:00
:tags: Vue.js, JavaScript, html, CSS, dropdown menu
:category: Vue.js
:summary: Dropdown menu implementation via Vue.js and CSS.
:og_image: https://cdn.dribbble.com/users/33345/screenshots/1027783/dropdown-menu.png
:adsu: yes

Dropdown menu implementation via Vue.js and CSS.
The dropdown menu here is similar to `Bootstrap dropdown`_.
Click the following button to see the demo.

.. raw:: html

  <style>
  .dropdown {
    position: relative;
  }

  .dropdown-toggle {
    background-color: white;
  }

  .dropdown-menu {
    position: absolute;
    border-radius: 5px;
    border-top-color: #C9D7F1;
    border-right-color: #36C;
    border-bottom-color: #36C;
    border-left-color: #A2BAE7;
    border-style: solid;
    border-width: 1px;
    z-index: 22;
    padding: 0;
    background-color: white;
    overflow: hidden;
    font-size: small;
    font-family: Arial;
    line-height: 2em;
    word-spacing: 0;
    margin-top: 2px;
  }

  .dropdown-menu a {
    font-size: 1.25em;
    color: #00C;
    padding: .25em 2em .25em 1em;
    text-decoration: none;
    background: white;
    display: block;
    cursor: pointer;
  }

  .dropdown-menu a:hover {
    background: #00C;
    color: white;
  }
  </style>

  <div id="vueapp">

  <div class="dropdown">
    <button class="dropdown-toggle" v-on:click="isShowDrop1 = !isShowDrop1">
      Dropdown
    </button>
    <ul class="dropdown-menu" v-show="isShowDrop1">
      <li><a>Item 1</a></li>
      <li><a>Item 2</a></li>
      <li><a>Item 3</a></li>
    </ui>
  </div>

  </div>

  <script src="https://unpkg.com/vue@2.3.3/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      isShowDrop1: false
    }
  });
  </script>

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <div id="vueapp">

  <div class="dropdown">
    <button class="dropdown-toggle" v-on:click="isShowDrop1 = !isShowDrop1">
      Dropdown
    </button>
    <ul class="dropdown-menu" v-show="isShowDrop1">
      <li><a>Item 1</a></li>
      <li><a>Item 2</a></li>
      <li><a>Item 3</a></li>
    </ui>
  </div>

  </div>

  <script src="https://unpkg.com/vue@2.3.3/dist/vue.js"></script>

Wrap your **dropdown toggle** and **dropdown menu** in ``div.dropdown``.
You can use anchor element for **dropdown toggle** instead of button element.

We use the variable *isShowDrop1* to control the visibility of dropdown menu.
The demo code here is very easy to understand.

**JavaScript**:

.. code-block:: javascript

  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      isShowDrop1: false
    }
  });

Set *isShowDrop1* to false in the initialization phase to hide the dropdown menu
in the beginning.

.. adsu:: 2

**CSS**:

.. code-block:: css

  .dropdown {
    position: relative;
  }

  .dropdown-toggle {
    background-color: white;
  }

  .dropdown-menu {
    position: absolute;
    border-radius: 5px;
    border-top-color: #C9D7F1;
    border-right-color: #36C;
    border-bottom-color: #36C;
    border-left-color: #A2BAE7;
    border-style: solid;
    border-width: 1px;
    z-index: 22;
    padding: 0;
    background-color: white;
    overflow: hidden;
    font-size: small;
    font-family: Arial;
    line-height: 2em;
    word-spacing: 0;
    margin-top: 2px;
  }

  .dropdown-menu a {
    font-size: 1.25em;
    color: #00C;
    padding: .25em 2em .25em 1em;
    text-decoration: none;
    background: white;
    display: block;
    cursor: pointer;
  }

  .dropdown-menu a:hover {
    background: #00C;
    color: white;
  }

Use ``position: relative;`` in parent ``.dropdown`` class and
``position: absolute;`` in child ``.dropdown-menu`` class to align the dropdown
menu under the dropdown toogle.

Other rules is only for making the dropdown look beautiful.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 58.0.3029.96 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``
- ``Vue.js 2.3.3``

----

References:

.. [1] `[Vue.js] Modal (Popup) <{filename}../11/vuejs-modal-popup%en.rst>`_
.. [2] `[AngularJS] Dropdown Menu Using Directive <{filename}../../../2015/02/04/angularjs-dropdown-menu-using-directive%en.rst>`_

.. _Vue.js: https://vuejs.org/
.. _Bootstrap dropdown: http://getbootstrap.com/javascript/#dropdowns
