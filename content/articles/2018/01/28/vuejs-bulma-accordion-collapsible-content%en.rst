[Vue.js] Bulma Accordion (Collapsible Content)
##############################################

:date: 2018-01-28T03:29+08:00
:tags: Vue.js, JavaScript, html, CSS, Accordion (Collapsible Content), Bulma
:category: Vue.js
:summary: Bulma accordion (collapsible content) implementation via Vue.js.
:og_image: https://bulma.io/images/extensions/bulma-accordion.png
:adsu: yes

Bulma_ accordion (collapsible content) implementation via Vue.js_.
The accordion here is similar to the accordion example in `Bootstrap Collapse`_.
Click the following panels to see the demo.

.. raw:: html

  <div id="vueapp" style="margin: 2rem;">

  <div class="panel">
    <div class="panel-heading" v-on:click="sel == 1 ? sel = 0 : sel = 1">
      Title #1
    </div>
    <div class="panel-block" v-show="sel == 1">
      Content Body #1
    </div>
  </div>
  <div class="panel">
    <div class="panel-heading" v-on:click="sel == 2 ? sel = 0 : sel = 2">
      Title #2
    </div>
    <div class="panel-block" v-show="sel == 2">
      Content Body #2
    </div>
  </div>
  <div class="panel">
    <div class="panel-heading" v-on:click="sel == 3 ? sel = 0 : sel = 3">
      Title #3
    </div>
    <div class="panel-block" v-show="sel == 3">
      Content Body #3
    </div>
  </div>

  </div>

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.13/dist/vue.js"></script>

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

Actually there is accordion in Bulma extensions [1]_, but we are not going to
use it. I like to use only standard components provided by Bulma. Here we use
`Bulma panel component`_ to implement the accordion.
The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <div id="vueapp">

  <div class="panel">
    <div class="panel-heading" v-on:click="sel == 1 ? sel = 0 : sel = 1">
      Title #1
    </div>
    <div class="panel-block" v-show="sel == 1">
      Content Body #1
    </div>
  </div>
  <div class="panel">
    <div class="panel-heading" v-on:click="sel == 2 ? sel = 0 : sel = 2">
      Title #2
    </div>
    <div class="panel-block" v-show="sel == 2">
      Content Body #2
    </div>
  </div>
  <div class="panel">
    <div class="panel-heading" v-on:click="sel == 3 ? sel = 0 : sel = 3">
      Title #3
    </div>
    <div class="panel-block" v-show="sel == 3">
      Content Body #3
    </div>
  </div>

  </div>

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.13/dist/vue.js"></script>

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

----

Tested on:

- ``Chromium 63.0.3239.132 on Ubuntu 17.10 (64-bit)``
- ``Vue.js 2.5.13``
- ``Bulma 0.6.2``

----

References:

.. [1] `Accordion | Bulma-Extensions <https://wikiki.github.io/components/accordion/>`_
.. adsu:: 3
.. [2] `[Vue.js] Accordion (Collapsible Content) <{filename}/articles/2017/05/22/vuejs-accordion-collapsible-content%en.rst>`_

.. _Vue.js: https://vuejs.org/
.. _Bulma: https://bulma.io/
.. _Bulma panel component: https://bulma.io/documentation/components/panel/
.. _Bootstrap Collapse: https://getbootstrap.com/docs/3.3/javascript/#collapse-example-accordion
