[Vue.js] Tab Panel
##################

:date: 2017-05-16T21:43+08:00
:tags: Vue.js, JavaScript, html, CSS, Tab Panel
:category: Vue.js
:summary: Tab panel implementation via Vue.js and CSS.
:og_image: https://s3-us-west-2.amazonaws.com/i.cdpn.io/193912.BlKxo.0e781013-9523-4f00-b75c-0c2f48b25100.png
:adsu: yes

Tab panel implementation via Vue.js and CSS.
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
  }
  </style>

  <div id="vueapp" style="background-color: Azure; padding: 1rem;">

  <div>

    <!-- Nav tabs -->
    <ul class="nav-tabs">
      <li><a v-on:click="tabsel = 'home'">Home</a></li>
      <li><a v-on:click="tabsel = 'profile'">Profile</a></li>
      <li><a v-on:click="tabsel = 'messages'">Messages</a></li>
      <li><a v-on:click="tabsel = 'settings'">Settings</a></li>
    </ul>

    <!-- Tab panes -->
    <div class="tab-content">
      <div class="tab-pane" v-show="tabsel == 'home'">My home</div>
      <div class="tab-pane" v-show="tabsel == 'profile'">My profile</div>
      <div class="tab-pane" v-show="tabsel == 'messages'">My messages</div>
      <div class="tab-pane" v-show="tabsel == 'settings'">My settings</div>
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
      tabsel: "home"
    }
  });
  </script>

The following is the source code for above demo.

**HTML**:

.. code-block:: html

  <div id="vueapp">

  <div>

    <!-- Nav tabs -->
    <ul class="nav-tabs">
      <li><a v-on:click="tabsel = 'home'">Home</a></li>
      <li><a v-on:click="tabsel = 'profile'">Profile</a></li>
      <li><a v-on:click="tabsel = 'messages'">Messages</a></li>
      <li><a v-on:click="tabsel = 'settings'">Settings</a></li>
    </ul>

    <!-- Tab panes -->
    <div class="tab-content">
      <div class="tab-pane" v-show="tabsel == 'home'">My home</div>
      <div class="tab-pane" v-show="tabsel == 'profile'">My profile</div>
      <div class="tab-pane" v-show="tabsel == 'messages'">My messages</div>
      <div class="tab-pane" v-show="tabsel == 'settings'">My settings</div>
    </div>

  </div>

  </div>

  <script src="https://unpkg.com/vue@2.3.3/dist/vue.js"></script>

We use the variable *tabsel* to indicate current selected tab. When users click
on the tab, update *tabsel* and hence show the selected tab pane according to
the value of *tabsel*.


**JavaScript**:

.. code-block:: javascript

  'use strict';

  var app = new Vue({
    el: '#vueapp',
    data: {
      tabsel: "home"
    }
  });

Set *tabsel* to ``home`` in the initialization phase to make the home tab as
default tab.

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
  }

Use flexbox to align the tabs in one row.

.. adsu:: 3

----

Tested on:

- ``Chromium 58.0.3029.96 Built on Ubuntu 17.04 (64-bit)``
- ``Vue.js 2.3.3``

----

References:

.. [1] `bootstrap/_nav.scss at v4-dev · twbs/bootstrap · GitHub <https://github.com/twbs/bootstrap/blob/v4-dev/scss/_nav.scss>`_
.. [2] `Class and Style Bindings — Vue.js <https://vuejs.org/v2/guide/class-and-style.html>`_

.. _Vue.js: https://vuejs.org/
.. _Bootstrap tab: https://getbootstrap.com/docs/3.3/javascript/#tabs
