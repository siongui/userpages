[Vue.js] Accordion (Collapsible Content)
########################################

:date: 2017-05-22T22:32+08:00
:tags: Vue.js, JavaScript, html, CSS, Accordion (Collapsible Content)
:category: Vue.js
:summary: Accordion (collapsible content) implementation via Vue.js.
:og_image: http://christineosazuwa.com/wp-content/uploads/2015/05/13-06_accordionmenu1.jpg
:adsu: yes

Accordion (collapsible content) implementation via Vue.js.
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
  }
  </style>

  <div id="vueapp">

  <div class="panel">
    <div class="panel-title" v-on:click="sel == 1 ? sel = 0 : sel = 1">
      Title #1
    </div>
    <div class="panel-body" v-show="sel == 1">
      Content Body #1
    </div>
  </div>
  <div class="panel">
    <div class="panel-title" v-on:click="sel == 2 ? sel = 0 : sel = 2">
      Title #2
    </div>
    <div class="panel-body" v-show="sel == 2">
      Content Body #2
    </div>
  </div>
  <div class="panel">
    <div class="panel-title" v-on:click="sel == 3 ? sel = 0 : sel = 3">
      Title #3
    </div>
    <div class="panel-body" v-show="sel == 3">
      Content Body #3
    </div>
  </div>

  </div>

  <script src="https://unpkg.com/vue@2.3.3/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      sel: 0
    }
  });
  </script>

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <div id="vueapp">

  <div class="panel">
    <div class="panel-title" v-on:click="sel == 1 ? sel = 0 : sel = 1">
      Title #1
    </div>
    <div class="panel-body" v-show="sel == 1">
      Content Body #1
    </div>
  </div>
  <div class="panel">
    <div class="panel-title" v-on:click="sel == 2 ? sel = 0 : sel = 2">
      Title #2
    </div>
    <div class="panel-body" v-show="sel == 2">
      Content Body #2
    </div>
  </div>
  <div class="panel">
    <div class="panel-title" v-on:click="sel == 3 ? sel = 0 : sel = 3">
      Title #3
    </div>
    <div class="panel-body" v-show="sel == 3">
      Content Body #3
    </div>
  </div>

  </div>

  <script src="https://unpkg.com/vue@2.3.3/dist/vue.js"></script>

We use the variable *sel* to indicate which panel is opened. If the value of
*sel* is 0, all panels are collapsed. If 1, the first panel is opened, and so
on.

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      sel: 0
    }
  });

Set *sel* to 0 in the initialization phase to make all panles collapsed.

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
  }

Nothing special in CSS code here. For demo purpose, I make CSS very simple. You
can try to add some animation if you want.

.. adsu:: 3

----

Tested on:

- ``Chromium Version 58.0.3029.110 Built on Ubuntu , running on Ubuntu 17.04 (64-bit)``
- ``Vue.js 2.3.3``

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
