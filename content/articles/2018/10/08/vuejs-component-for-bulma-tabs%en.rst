Vue.js Component for Bulma Tabs
###############################

:date: 2018-10-08T18:44+08:00
:tags: Vue.js, JavaScript, html, CSS, Tab Panel, Bulma
:category: Vue.js
:summary: Reusable *Vue.js* component for *Bulma* tabs, and extend the original
          Bulma tabs by adding a pane to for each tab.
:og_image: https://s3-us-west-2.amazonaws.com/i.cdpn.io/9927.YVZBvd.038631b0-afac-4ad4-9405-1411877f0eda.png
:adsu: yes


Reusable Vue.js_ component_ for Bulma_ tabs_.
We extend Bulma tabs by adding a pane for each tab.
Click the following tabs to see the demo.

.. raw:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.1/css/bulma.css" integrity="sha256-zKA1Bf41O96+gJSlkn/Bh2HATW/OhwkApPlYTp3B5O8=" crossorigin="anonymous" />

  <div id="vueapp" style="margin: 2rem;">

  <bulma-tab-panel :tabpanes="mytabpanes"></bulma-tab-panel>

  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>

  <style>
    li > a:hover { text-decoration: none; }
  </style>

.. raw:: html

  <script>
  'use strict';

  Vue.component("bulma-tab-panel", {
      template: `
  <div>
    <!-- Nav tabs -->
    <div class="tabs">
      <ul>
        <li v-for="(tabpane, index) in tabpanes" v-bind:class="{ 'is-active': tabsel == index }" @click="tabsel = index"><a>{{tabpane.nav}}</a></li>
      </ul>
    </div>
    <!-- Tab panes -->
    <div class="content">
      <div v-for="(tabpane, index) in tabpanes" v-show="tabsel == index">{{tabpane.pane}}</div>
    </div>
  </div>
  `,
      props: ["tabpanes"],
      data: function() {
        return {
          tabsel: 0
        }
      }
  });

  var app = new Vue({
    el: '#vueapp',
    data: {
      mytabpanes: [
        {nav: "tab1", pane:"tab1 content"},
        {nav: "tab2", pane:"tab2 content"},
        {nav: "tab3", pane:"tab3 content"},
        {nav: "tab4", pane:"tab4 content"},
      ]
    }
  });
  </script>


The following is the source code for above demo.


**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.1/css/bulma.css" integrity="sha256-zKA1Bf41O96+gJSlkn/Bh2HATW/OhwkApPlYTp3B5O8=" crossorigin="anonymous" />

  <div id="vueapp">

  <bulma-tab-panel :tabpanes="mytabpanes"></bulma-tab-panel>

  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>


We define a Vue.js component whose name is *bulma-tab-panel*, and we pass data
of tabs to it.

|

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  Vue.component("bulma-tab-panel", {
      template: `
  <div>
    <!-- Nav tabs -->
    <div class="tabs">
      <ul>
        <li v-for="(tabpane, index) in tabpanes" v-bind:class="{ 'is-active': tabsel == index }" @click="tabsel = index"><a>{{tabpane.nav}}</a></li>
      </ul>
    </div>
    <!-- Tab panes -->
    <div class="content">
      <div v-for="(tabpane, index) in tabpanes" v-show="tabsel == index">{{tabpane.pane}}</div>
    </div>
  </div>
  `,
      props: ["tabpanes"],
      data: function() {
        return {
          tabsel: 0
        }
      }
  });

  var app = new Vue({
    el: '#vueapp',
    data: {
      mytabpanes: [
        {nav: "tab1", pane:"tab1 content"},
        {nav: "tab2", pane:"tab2 content"},
        {nav: "tab3", pane:"tab3 content"},
        {nav: "tab4", pane:"tab4 content"},
      ]
    }
  });


The data passed to the component is an array consisting of nav and content of
the tabs. we use *v-for* to create the actual HTML of the tabs in the component.

We use the variable *tabsel* to indicate current selected tab.
When users click on the tab, update *tabsel* according to the tab index and
hence show the selected tab pane according to the value of *tabsel*.

The *is-active* class of tabs will also be updated according to *tabsel*.
Set *tabsel* to ``0`` in the initialization phase to make first tab as default
tab.

You can also make the same tabs without using component, see [4]_.

----

Tested on:

- ``Bulma 0.7.1``
- ``Vue.js 2.5.17``
- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``

----

References:

.. [1] `Tabs | Bulma: a modern CSS framework based on Flexbox <https://bulma.io/documentation/components/tabs/>`_
.. adsu:: 3
.. [2] `揭密 Vue 的双向绑定 - 边城客栈 - 开源中国 <https://my.oschina.net/jamesfancy/blog/2222930>`_
.. [3] `Components Basics — Vue.js <https://vuejs.org/v2/guide/components.html>`_
.. [4] `[Vue.js] Bulma Tabs <{filename}/articles/2018/01/26/vuejs-bulma-tabs%en.rst>`_
.. [5] `List Rendering — Vue.js <https://vuejs.org/v2/guide/list.html>`_

.. _Vue.js: https://vuejs.org/
.. _component: https://vuejs.org/v2/guide/components.html
.. _Bulma: https://bulma.io/
.. _tabs: https://bulma.io/documentation/components/tabs/
