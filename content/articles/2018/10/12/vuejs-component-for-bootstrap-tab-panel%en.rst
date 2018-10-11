Vue.js Component for Bootstrap Tab Panel
########################################

:date: 2018-10-12T00:00+08:00
:tags: Vue.js, JavaScript, html, CSS, Tab Panel, Bootstrap
:category: Vue.js
:summary: Reusable *Vue.js* component to extend *Bootstrap* navigational tabs to
          create tabbable panes of local content.
:og_image: https://s3-us-west-2.amazonaws.com/i.cdpn.io/9927.YVZBvd.038631b0-afac-4ad4-9405-1411877f0eda.png
:adsu: yes


Reusable Vue.js_ component_ for Bootstrap_ `tab panel`_.
Use Vue.js to extend *Bootstrap* navigational tabs_ to create tabbable panes of
local content.
Click the following tabs to see the demo.

|

.. raw:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <div id="vueapp">
    <bootstrap-tab-panel :tabpanes="mytabpanes"></bootstrap-tab-panel>
  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>

.. raw:: html

  <script>
  'use strict';

  Vue.component("bootstrap-tab-panel", {
      template: `
  <div>
    <!-- Nav tabs -->
    <nav>
      <div class="nav nav-tabs" role="tablist">
        <a class="nav-item nav-link" role="tab"
           v-for="(tabpane, index) in tabpanes"
           v-bind:class="{ 'active': tabsel == index }"
           @click="tabsel = index">
          {{tabpane.nav}}
        </a>
      </div>
    </nav>
    <!-- Tab panes -->
    <div class="tab-content">
      <div class="tab-pane fade" role="tabpanel"
           v-for="(tabpane, index) in tabpanes"
           v-bind:class="{ 'show active': tabsel == index }">
         {{tabpane.pane}}
      </div>
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

|
|

The following is the source code for above demo.


**HTML**:

.. code-block:: html

  <!-- you can put the following line in your html head -->
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">

  <div id="vueapp">
    <bootstrap-tab-panel :tabpanes="mytabpanes"></bootstrap-tab-panel>
  </div>

  <!-- you can put the following line at the end of html body -->
  <script src="https://cdn.jsdelivr.net/npm/vue@2.5.17/dist/vue.js"></script>

|

We define a Vue.js component whose name is *bootstrap-tab-panel*, and we pass
data of tabs and panes to it.

|

.. adsu:: 2

**JavaScript**:

.. code-block:: javascript

  'use strict';

  Vue.component("bootstrap-tab-panel", {
      template: `
  <div>
    <!-- Nav tabs -->
    <nav>
      <div class="nav nav-tabs" role="tablist">
        <a class="nav-item nav-link" role="tab"
           v-for="(tabpane, index) in tabpanes"
           v-bind:class="{ 'active': tabsel == index }"
           @click="tabsel = index">
          {{tabpane.nav}}
        </a>
      </div>
    </nav>
    <!-- Tab panes -->
    <div class="tab-content">
      <div class="tab-pane fade" role="tabpanel"
           v-for="(tabpane, index) in tabpanes"
           v-bind:class="{ 'show active': tabsel == index }">
         {{tabpane.pane}}
      </div>
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

The *active* class of tabs will also be updated according to *tabsel*.
Set *tabsel* to ``0`` in the initialization phase to make first tab as default
tab.

----

Tested on:

- ``Bootstrap 4.1.3``
- ``Vue.js 2.5.17``
- ``Chromium 69.0.3497.81 on Ubuntu 18.04 (64-bit)``

----

References:

.. [1] `Vue.js Component for Bulma Tabs <{filename}/articles/2018/10/08/vuejs-component-for-bulma-tabs%en.rst>`_
.. [2] `Pure CSS Bootstrap Tab Panel <{filename}/articles/2018/10/01/css-only-toggle-bootstrap-tab-panel%en.rst>`_
.. adsu:: 3
.. [3] `Components Basics — Vue.js <https://vuejs.org/v2/guide/components.html>`_
.. [4] `List Rendering — Vue.js <https://vuejs.org/v2/guide/list.html>`_

.. _Vue.js: https://vuejs.org/
.. _component: https://vuejs.org/v2/guide/components.html
.. _Bootstrap: https://getbootstrap.com/
.. _tab panel: https://getbootstrap.com/docs/4.1/components/navs/#javascript-behavior
.. _tabs: https://getbootstrap.com/docs/4.1/components/navs/#tabs
