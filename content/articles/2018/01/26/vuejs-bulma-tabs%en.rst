[Vue.js] Bulma Tabs
###################

:date: 2018-01-26T20:30+08:00
:modified: 2018-10-08T18:46+08:00
:tags: Vue.js, JavaScript, html, CSS, Tab Panel, Bulma
:category: Vue.js
:summary: Add more features to *Bulma* tabs via *Vue.js*.
:og_image: https://s3-us-west-2.amazonaws.com/i.cdpn.io/9927.YVZBvd.038631b0-afac-4ad4-9405-1411877f0eda.png
:adsu: yes

Add more features to Bulma_ tabs_ via Vue.js_.
The tab panel here is similar to `Bootstrap tab`_.
Click the following tabs to see the demo.

.. raw:: html

  <div id="vueapp" style="margin: 2rem;">

  <!-- Nav tabs -->
  <div class="tabs">
    <ul>
      <li v-bind:class="{ 'is-active': tabsel == 'pic' }" @click="tabsel = 'pic'"><a>Pictures</a></li>
      <li v-bind:class="{ 'is-active': tabsel == 'music' }" @click="tabsel = 'music'"><a>Music</a></li>
      <li v-bind:class="{ 'is-active': tabsel == 'video' }" @click="tabsel = 'video'"><a>Videos</a></li>
      <li v-bind:class="{ 'is-active': tabsel == 'doc' }" @click="tabsel = 'doc'"><a>Documents</a></li>
    </ul>
  </div>

  <!-- Tab panes -->
  <div class="content">
    <div v-show="tabsel == 'pic'">My Pictures</div>
    <div v-show="tabsel == 'music'">My Music</div>
    <div v-show="tabsel == 'video'">My Videos</div>
    <div v-show="tabsel == 'doc'">My Documents</div>
  </div>

  </div>

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.13/dist/vue.js"></script>
  <style>
    li > a:hover { text-decoration: none; }
  </style>

.. raw:: html

  <script>
  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      tabsel: "pic"
    }
  });
  </script>


The following is the source code for above demo.


**HTML**:

.. code-block:: html

  <div id="vueapp">

  <!-- Nav tabs -->
  <div class="tabs">
    <ul>
      <li v-bind:class="{ 'is-active': tabsel == 'pic' }" @click="tabsel = 'pic'"><a>Pictures</a></li>
      <li v-bind:class="{ 'is-active': tabsel == 'music' }" @click="tabsel = 'music'"><a>Music</a></li>
      <li v-bind:class="{ 'is-active': tabsel == 'video' }" @click="tabsel = 'video'"><a>Videos</a></li>
      <li v-bind:class="{ 'is-active': tabsel == 'doc' }" @click="tabsel = 'doc'"><a>Documents</a></li>
    </ul>
  </div>

  <!-- Tab panes -->
  <div class="content">
    <div v-show="tabsel == 'pic'">My Pictures</div>
    <div v-show="tabsel == 'music'">My Music</div>
    <div v-show="tabsel == 'video'">My Videos</div>
    <div v-show="tabsel == 'doc'">My Documents</div>
  </div>

  </div>

  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.6.2/css/bulma.min.css">
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.13/dist/vue.js"></script>


We use the variable *tabsel* to indicate current selected tab.
When users click on the tab, update *tabsel* and hence show the selected tab
pane according to the value of *tabsel*.
The *is-active* class of tabs will also be updated according to *tabsel*.

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      tabsel: "pic"
    }
  });

Set *tabsel* to ``pic`` in the initialization phase to make the home tab as
default tab.

You can make the above code more reusable by Vue.js component. See [3]_ for more
details.

----

Tested on:

- ``Chromium 63.0.3239.132 on Ubuntu 17.10 (64-bit)``
- ``Vue.js 2.5.13``
- ``Bulma 0.6.2``

----

References:

.. [1] `Tabs | Bulma: a modern CSS framework based on Flexbox <https://bulma.io/documentation/components/tabs/>`_
.. adsu:: 3
.. [2] `Class and Style Bindings — Vue.js <https://vuejs.org/v2/guide/class-and-style.html>`_
.. [3] `Vue.js Component for Bulma Tabs <{filename}/articles/2018/10/08/vuejs-component-for-bulma-tabs%en.rst>`_

.. _Vue.js: https://vuejs.org/
.. _Bulma: https://bulma.io/
.. _tabs: https://bulma.io/documentation/components/tabs/
.. _Bootstrap tab: https://getbootstrap.com/docs/3.3/javascript/#tabs
